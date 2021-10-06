package handler

import (
	"LotteryServer/src/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)


func Register(r *gin.Engine) {
	//curl -H "Content-Type:application/json" -X GET 'http://localhost:8080/api/v1/healthcheck'
	r.GET("/api/v1/healthcheck", HealthCheck())

	//curl -H "Content-Type:application/json" -X GET 'http://localhost:8080/api/v1/recommendation/doublecolor/normal?history=100'
	r.GET("/api/v1/recommendation/doublecolor/normal", GetTopNormalRec())

	//curl -H "Content-Type:application/json" -X GET 'http://localhost:8080/api/v1/recommendation/doublecolor/normal/batch?top=3&history=100'
	r.GET("/api/v1/recommendation/doublecolor/normal/batch", GetTopNNormalRec())
}


func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(200, "pong")
	}
}


func GetTopNormalRec() gin.HandlerFunc {
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
		if  err != nil {
			logrus.Errorf("get topN normal convert failed: %s-%s", err.Error())
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		err, histData := service.GetHistoryData(hist)
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
func GetTopNNormalRec() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "")
	}
}