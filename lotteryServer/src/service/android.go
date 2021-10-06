package service

import "LotteryServer/src/model"

func AndroidRBBallsData(reds []int, blue int) model.AndroidApp {
	return model.AndroidApp{
		RedBalls: reds,
		BlueBall: blue,
	}
}