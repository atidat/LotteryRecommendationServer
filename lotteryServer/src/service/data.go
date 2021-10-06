package service

import (
	"LotteryServer/src/manager"
	"LotteryServer/src/model"
	"github.com/sirupsen/logrus"
)

/*
	description: normal analysis
	details: 获取N期往期数据，分析每一位置上，1~33数据出现的频率，挑出每一位置上最少出现的数字
	parameter:
*/
func NormalAnalysis(data []model.RBBall) (error, []int, int) {
	logrus.Info("[Normal Analysis]")
	blueSlots, redSlots := manager.FormatAndMergeRawBallData(&data)
	recReds, recBlue := manager.CalcFrequency(blueSlots, redSlots)
	logrus.Info("Recommendation")
	logrus.Info(recReds, recBlue)
	return nil, recReds, recBlue
}


/*
6个红球槽位
[1,4,7,2,2,7]
[6,22,5,15,8,1]


slots = [6][33]{{

}}

1个蓝球槽位

*/