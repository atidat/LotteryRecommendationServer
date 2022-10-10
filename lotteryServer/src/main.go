package main

import (
	"LotteryServer/src/dao"
	"LotteryServer/src/handler"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Start LotteryRecommendation...")
	r := gin.Default()
	handler.Register(r)
	dao.DBInit()
	dao.CacheInit()
	_ = r.Run()
}
