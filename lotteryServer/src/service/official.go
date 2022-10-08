package service

import (
	"LotteryServer/src/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

/*
return: [{Red: "01,03,21,17,07,04", Blue: "12"}, ..., {Red: "12,33,14,07,24, 31", Blue: "09"}]
*/
func FetchHistoryData(hist int) (error, []model.DoubleColorBall) {
	/* 向彩票官网拉数据
		ping www.cwl.gov.cn
		61.132.231.47

	http://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=ssq&issueCount=999&issueStart=&issueEnd=&dayStart=&dayEnd=
	*/
	url := fmt.Sprintf("http://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=ssq&issueCount=%d&issueStart=&issueEnd=&dayStart=&dayEnd=", hist)
	res, err := http.Get(url)
	if err != nil {
		return err, nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	offRes := &model.OfficialRes{}
	err = json.Unmarshal(body, offRes)
	if err != nil {
		logrus.Errorf("convert official result failed: %s", err.Error())
		return err, nil
	}

	rawData := make([]model.DoubleColorBall, 0)
	for _, offres := range offRes.Result {
		rawData = append(rawData, model.DoubleColorBall{
			Red:  offres.(map[string]interface{})["red"].(string),
			Blue: offres.(map[string]interface{})["blue"].(string),
			Time: offres.(map[string]interface{})["date"].(string),
		})
	}
	return nil, rawData
}

func CompareDoubleColorData(hist *[]model.DoubleColorBall, last string) int {
	var cnt int
	if last != "" {
		for ; cnt < len(*hist); cnt++ {
			if (*hist)[cnt].Time[:10] <= last {
				break
			}
			(*hist)[cnt].Time = (*hist)[cnt].Time[:10]
		}

		if cnt < 1 {
			return cnt
		}
	} else {
		cnt = len(*hist)
		for i := 0; i < cnt; i++ {
			(*hist)[i].Time = (*hist)[i].Time[:10]
		}
	}
	return cnt
}
