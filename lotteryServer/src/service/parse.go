package service

import (
	"LotteryServer/src/model"
	"encoding/json"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func ParseDoubleColorRawContent(rawFile string) *model.OfficialDoubleColor {
	raws, err := ioutil.ReadFile(rawFile)
	if err != nil {
		logrus.Warnf("open double-color-ball file failed: %s", err.Error())
		return nil
	}
	var dcContent model.OfficialDoubleColor
	err = json.Unmarshal(raws, &dcContent)
	if err != nil {
		logrus.Warnf("parse double-color-ball raw content failed: %s", err.Error())
		return nil
	}
	return &dcContent
}
