package utils

import (
	"LotteryServer/src/model"
	"strconv"
	"strings"
)

func ConvertRedBallStr(reds string) model.RedArray {
	return *(*model.RedArray)(ballStr2Array(reds))
}

func ConvertBlueBallStr(blues string) model.BlueArray {
	return *(*model.BlueArray)(ballStr2Array(blues))
}

func ballStr2Array(s string) []int {
	ballStrs := strings.Split(s, ",")
	balls := make([]int, len(ballStrs))
	for k, ball := range ballStrs {
		balls[k], _ = strconv.Atoi(ball)
	}
	return balls
}

func Convert2RedBallStr(reds model.RedArray) string {
	return ballArray2Str(reds)
}

func Convert2BlueBallStr(blues model.BlueArray) string {
	return ballArray2Str(blues)
}

func ballArray2Str(balls interface{}) string {
	str := ""
	for _, ball := range balls.([]int) {
		str += strconv.Itoa(ball) + ","
	}
	return str[:len(str)-1]
}

func AcculateRedBalls(balls *model.RedArray, incres model.RedArray) {
	for ind, incre := range incres {
		(*balls)[ind] += incre
	}
}

func AcculateBlueBalls(balls *model.BlueArray, incres model.BlueArray) {
	for ind, incre := range incres {
		balls[ind] += incre
	}
}
