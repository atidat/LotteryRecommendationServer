package main

import (
	"LotteryServer/src/handler"
	"LotteryServer/src/dao"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Start LotteryRecommendation...")
	r := gin.Default()
	handler.Register(r)
	dao.DBInit()
	_ = r.Run()
}
