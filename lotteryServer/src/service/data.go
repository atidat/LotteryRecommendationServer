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

func BatchInsertDoubleColorHist(data []model.DoubleColorBall)  {
	dao.BatchInsertDoubleColorHistData(data)
}

func BatchUpdateDoubleColorCnt(data []model.DoubleColorBall) {
	redBallsCnts, blueBallCnt := dao.GetDoubleColorCnts()
	for _, datus := range data {
		redBalls := utils.ConvertRedBallStr(datus.Red) // redBalls含有6数字个数字，每个数字[1~33]
		blueBalls := utils.ConvertRedBallStr(datus.Blue)
		for i := 0; i < len(redBallsCnts); i++ {
			doubleColorRedBallInc(redBallsCnts, i, redBalls[i])
		}
		doubleColorBlueBallInc(blueBallCnt, blueBalls[0])
	}
	dao.UpdateDoubleColorCntData(redBallsCnts, blueBallCnt)
}

func doubleColorRedBallInc(handlers []*model.DoubleColorRed, ind, tar int) {
	if tar < 1 || tar > 33 {
		logrus.Errorf("red ooooops! ==> %d", tar)
		return
	}
	handler := handlers[ind]
	switch tar {
	case 1: handler.Add1()
	case 2: handler.Add2()
	case 3: handler.Add3()
	case 4: handler.Add4()
	case 5: handler.Add5()
	case 6: handler.Add6()
	case 7: handler.Add7()
	case 8: handler.Add8()
	case 9: handler.Add9()
	case 10: handler.Add10()
	case 11: handler.Add11()
	case 12: handler.Add12()
	case 13: handler.Add13()
	case 14: handler.Add14()
	case 15: handler.Add15()
	case 16: handler.Add16()
	case 17: handler.Add17()
	case 18: handler.Add18()
	case 19: handler.Add19()
	case 20: handler.Add20()
	case 21: handler.Add21()
	case 22: handler.Add22()
	case 23: handler.Add23()
	case 24: handler.Add24()
	case 25: handler.Add25()
	case 26: handler.Add26()
	case 27: handler.Add27()
	case 28: handler.Add28()
	case 29: handler.Add29()
	case 30: handler.Add30()
	case 31: handler.Add31()
	case 32: handler.Add32()
	case 33: handler.Add33()
	}
}

func doubleColorBlueBallInc(handler *model.TblDoubleColorBlue, tar int) {
	if tar < 1 || tar > 16 {
		logrus.Errorf("blue ooooops! ==> %d", tar)
		return
	}
	switch tar {
	case 1: handler.Add1()
	case 2: handler.Add2()
	case 3: handler.Add3()
	case 4: handler.Add4()
	case 5: handler.Add5()
	case 6: handler.Add6()
	case 7: handler.Add7()
	case 8: handler.Add8()
	case 9: handler.Add9()
	case 10: handler.Add10()
	case 11: handler.Add11()
	case 12: handler.Add12()
	case 13: handler.Add13()
	case 14: handler.Add14()
	case 15: handler.Add15()
	case 16: handler.Add16()
	}
}