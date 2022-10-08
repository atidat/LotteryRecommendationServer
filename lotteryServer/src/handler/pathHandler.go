package handler

import (
	"LotteryServer/src/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Register(r *gin.Engine) {
	//curl -H "Content-Type:application/json" -X GET 'http://localhost:8080/api/v1/healthcheck'
	r.GET("/api/v1/healthcheck", healthCheck())

	//curl -H "Content-Type:application/json" -X GET 'http://localhost:8080/api/v1/recommendation/doublecolor/normal?history=100'
	r.GET("/api/v1/recommendation/doublecolor/normal", getTopNormalRec())

	//curl -H "Content-Type:application/json" -X GET 'http://localhost:8080/api/v1/recommendation/doublecolor/normal/batch?top=3&history=100'
	r.GET("/api/v1/recommendation/doublecolor/normal/batch", getTopNNormalRec())

	//curl -H "Content-Type:application/json" -X POST -d '' 'http://localhost:8080/api/v1/doublecolor/datarepo'
	r.POST("/api/v1/doublecolor/datarepo", updateDoubleColorDataRepo())
}

func healthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "pong")
	}
}

func getTopNormalRec() gin.HandlerFunc {
	return func(c *gin.Context) {
		logrus.Info("topN recommends in")

		// hist最多只支持100期数据
		histStr, histExist := c.GetQuery("history")
		if !histExist {
			logrus.Error("get topN normal parse failed")
			c.String(http.StatusInternalServerError, "parse parameter failed")
			return
		}
		hist, err := strconv.Atoi(histStr)
		if err != nil {
			logrus.Errorf("get topN normal convert failed: %s-%s", err.Error())
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		err, histData := service.FetchHistoryData(hist)
		if err != nil {
			logrus.Errorf("get topN normal history failed: %s", err.Error())
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		err, reds, blue := service.NormalAnalysis(histData)
		if err != nil {
			logrus.Errorf("get topN normal analysis failed: %s", err.Error())
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		appData := service.AndroidRBBallsData(reds, blue)
		c.JSON(http.StatusOK, appData)
	}
}

// To be continued
func getTopNNormalRec() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "")
	}
}

func updateDoubleColorDataRepo() gin.HandlerFunc {
	return func(c *gin.Context) {
		err, histData := service.FetchHistoryData(100)
		if err != nil {
			logrus.Errorf("get topN normal history failed: %s", err.Error())
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		logrus.Info("get history data done")

		dbLastDate := service.GetDoubleColorNewestData().Date
		logrus.Infof("last date in db: %s", dbLastDate)

		cnt := service.CompareDoubleColorData(&histData, dbLastDate)
		if cnt == 0 {
			c.String(http.StatusOK, "update successfully")
			return
		}
		logrus.Info("get newest data done")

		service.BatchInsertDoubleColorHist(histData[:cnt])
		logrus.Info("batch insert done")
		service.BatchUpdateDoubleColorCnt(histData[:cnt])
		logrus.Info("batch update done")
		c.String(http.StatusOK, "update successfully")
	}
}
