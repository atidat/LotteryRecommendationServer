package dao

import (
	"LotteryServer/src/model"
	"LotteryServer/src/utils"

	"github.com/go-redis/redis"
)

var cacheCli = &redis.Client{}

func connCfgInfo() redis.Options {
	return redis.Options{
		Addr:     "127.0.0.1:3306",
		Password: "123123",
	}
}

func CacheInit() {
	cacheCli = redis.NewClient(connCfgInfo())
}

func cachePing() error {
	_, err := cacheCli.Ping().Result()
	return err
}

func BatchInsertDoubleColorHistCache(data *[]model.DoubleColorBall) {
	// key：redhist
	// value：["", "", ..., ""]
	hists := make([]string, len(data))
	for i, datum := range data {
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
	for _, datum := range data {
		redsCnt, blueCnt := utils.ConvertDoubleColorCSVStr(datum)
		for _, redCnt := range redsCnt {
			redCnts[redCnt]++
		}
		blueCnts[blueCnt]++
	}

	stockRedCnts, stockBlueCnts := FetchDoubleCntCache()
	stockRedCnts += redCnts
	stockBlueCnts += blueCnts
	cacheCli.Set(model.CacheDoubleColorRedsCnt, utils.Convert2RedBallStr(stockRedCnts))
	cacheCli.Set(model.CacheDoubleColorBluesCnt, utils.Convert2BlueBallStr(stockBlueCnts))
}

func FetchDoubleCntCache() (model.RedArray, model.BlueArray) {
	reds := cacheCli.Get(model.CacheDoubleColorRedsCnt)
	blues := cacheCli.Get(model.CacheDoubleColorBluesCnt)
	return utils.ConvertRedBallStr(reds), utils.ConvertBlueBallStr(blues)
}
