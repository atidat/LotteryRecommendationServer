package dao

import (
	"LotteryServer/src/model"

	"github.com/go-redis/redis"
)

var cacheCli = &redis.Client{}

func init() {
	cacheCli = redis.NewClient(redis.Options)
}

func cachePing() error {
	_, err := cacheCli.Ping().Result()
	return err
}

func BatchInsertDoubleColorHistCache(data []model.DoubleColorBall) {
	// key：redhist/日期
	// value：["", "", ..., ""]
	cacheCli.Lpush("redhist", )
}
