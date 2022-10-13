package main

import (
	"LotteryServer/src/dao"
	"LotteryServer/src/handler"
	"LotteryServer/src/service"
	"fmt"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Start LotteryRecommendation...")
	r := gin.Default()
	handler.Register(r)
	pprof.Register(r)
	dao.DBInit()
	dao.CacheInit()

	v := service.ParseDoubleColorRawContent("../ha.json")
	fmt.Printf("parsed is %v", v)

	_ = r.Run()
}
