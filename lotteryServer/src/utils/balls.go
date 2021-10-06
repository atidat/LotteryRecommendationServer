package utils

import (
	"strconv"
	"strings"
)

func ConvertRedBallStr(reds string) []int {
	return ballStr2Array(reds)
}

func ConvertBlueBallStr(blue string) []int {
	return ballStr2Array(blue)
}

func ballStr2Array(s string) []int {
	ballStrs := strings.Split(s, ",")
	balls := make([]int, len(ballStrs), len(ballStrs))
	for k, _ := range ballStrs {
		balls[k], _ = strconv.Atoi(ballStrs[k])
	}
	return balls
}