package dao

import (
	"LotteryServer/src/model"
	"LotteryServer/src/utils"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var cacheCli = &redis.Client{}

func cacheConnCfgInfo() *redis.Options {
	return &redis.Options{
		Addr:     "192.168.1.4:6379",
		Password: "",
	}
}

func CacheInit() {
	cacheCli = redis.NewClient(cacheConnCfgInfo())
	if cachePing() != nil {
		logrus.Warn("redis init failed")
	} else {
		logrus.Debug("redis init failed")
	}
}

func cachePing() error {
	_, err := cacheCli.Ping().Result()
	return err
}

func BatchInsertDoubleColorHistCache(data *[]model.DoubleColorBall) {
	// key：redhist
	// value：["", "", ..., ""]
	hists := make([]string, len(*data))
	for i, datum := range *data {
		hists[i] = datum.Time + "_" + datum.Red + "_" + datum.Blue
	}
	cacheCli.LPush(model.CacheDoubleColorHist, hists)
}

func BatchFetchDoubleColorHistCache() []model.DoubleColorBall {
	histLen := cacheCli.LLen(model.CacheDoubleColorHist).Val()
	histTotal := cacheCli.LRange(model.CacheDoubleColorHist, 0, histLen).Val()
	return utils.ParseDoubleColorHists(histTotal)
}

func BatchUpdateDoubleColorCnt(data *[]model.DoubleColorBall) /*bool*/ {
	redCnts, blueCnts := model.RedArray{}, model.BlueArray{}
	for _, datum := range *data {
		redsCnt, blueCnt := utils.ConvertDoubleColorCSVStr(datum)
		for _, redCnt := range redsCnt {
			redCnts[redCnt]++
		}
		blueCnts[blueCnt]++
	}

	stockRedCnts, stockBlueCnts := FetchDoubleCntCache()
	(&stockRedCnts).Add(redCnts)
	(&stockBlueCnts).Add(blueCnts)
	cacheCli.Set(model.CacheDoubleColorRedsCnt, utils.Convert2RedBallStr(stockRedCnts), -1)
	cacheCli.Set(model.CacheDoubleColorBluesCnt, utils.Convert2BlueBallStr(stockBlueCnts), -1)
}

func FetchDoubleCntCache() (model.RedArray, model.BlueArray) {
	reds := cacheCli.Get(model.CacheDoubleColorRedsCnt)
	blues := cacheCli.Get(model.CacheDoubleColorBluesCnt)
	return utils.ConvertRedBallStr(reds.Val()), utils.ConvertBlueBallStr(blues.Val())
}
