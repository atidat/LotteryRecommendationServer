package manager

import (
	"LotteryServer/src/model"
	"LotteryServer/src/utils"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

/*
	return: 红/蓝球各个槽位各数字出现的次数
*/
func FormatAndMergeRawBallData(data *[]model.RBBall) (*[][]int, *[]int) {
	blueSlots := []int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}
	redSlots := [][]int{
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
		{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0},
	}

	for _, datus := range *data {
		reds := utils.ConvertRedBallStr(datus.Red)
		blue := utils.ConvertBlueBallStr(datus.Blue)
		for ind, red := range reds {
			redSlots[ind][red-1]++
		}
		blueSlots[blue[0]-1]++
	}
	logrus.Info(redSlots)
	logrus.Info(blueSlots)
	return &redSlots, &blueSlots
}

func CalcFrequency(reds *[][]int, blue *[]int) ([]int, int) {
	//pickReds[i]:存储reds[i]中相同大小的索引
	pickReds := [][]int{}
	for _, red := range *reds {
		pickRed := utils.MinEleInds(&red)
		pickReds = append(pickReds, *pickRed)
	}
	pickBlue := utils.MinEleInds(blue)

	logrus.Info(pickReds)
	logrus.Info(pickBlue)

	rand.Seed(time.Now().Unix())
	finalReds := []int{}
	for i := 0; i < len(pickReds); i++ {
		ind := rand.Intn(len(pickReds[i]))
		finalReds = append(finalReds, pickReds[i][ind]+1)
	}

	finalBlue := -1
	if len(*pickBlue) > 0 {
		ind := rand.Intn(len(*pickBlue))
		finalBlue = (*pickBlue)[ind]+1
	}
	return finalReds, finalBlue
}