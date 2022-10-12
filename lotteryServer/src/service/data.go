package service

import (
	"LotteryServer/src/dao"
	"LotteryServer/src/manager"
	"LotteryServer/src/model"
	"LotteryServer/src/utils"

	"github.com/sirupsen/logrus"
)

/*
description: normal analysis
details: 获取N期往期数据，分析每一位置上，1~33数据出现的频率，挑出每一位置上最少出现的数字
parameter:
*/
func NormalAnalysis(data []model.DoubleColorBall) (error, []int, int) {
	logrus.Info("[Normal Analysis]")
	blueSlots, redSlots := manager.FormatAndMergeRawBallData(&data)
	recReds, recBlue := manager.CalcFrequency(blueSlots, redSlots)
	logrus.Info("Recommendation")
	logrus.Info(recReds, recBlue)
	return nil, recReds, recBlue
}

func GetDoubleColorNewestData() *model.TblDoubleColorHist {
	return dao.GetDoubleColorNewest()
}

func BatchInsertDoubleColorHist(data []model.DoubleColorBall) {
	dao.BatchInsertDoubleColorHistDB(data)
}

func BatchUpdateDoubleColorCnt(data []model.DoubleColorBall) {
	redBallsRaw, blueBallRaw := dao.GetDoubleColorCnts()
	redBallsCnts, blueBallsCnts := redBallsRaw.RSlots, blueBallRaw.BSlots
	for _, datus := range data {
		redBalls := utils.ConvertRedBallStr(datus.Red) // redBalls含有6数字个数字，每个数字[1~33]
		blueBalls := utils.ConvertRedBallStr(datus.Blue)
		for _, redBall := range redBalls {
			redBallsCnts[redBall]++
		}
		for _, blueBall := range blueBalls {
			blueBallsCnts[blueBall]++
		}
	}
	dao.UpdateDoubleColorCntData(redBallsRaw, blueBallRaw)
}

func doubleColorRedBallInc(handlers []*model.DoubleColorRed, ind, tar int) {
	if tar < 1 || tar > 33 {
		logrus.Errorf("red ooooops! ==> %d", tar)
		return
	}
	handler := handlers[ind]
	handler.Add([]int{tar})
}

func doubleColorBlueBallInc(handler *model.DoubleColorBlue, tar int) {
	if tar < 1 || tar > 16 {
		logrus.Errorf("blue ooooops! ==> %d", tar)
		return
	}
	handler.Add([]int{tar})
}
