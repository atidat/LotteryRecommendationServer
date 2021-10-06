package service

import (
	"LotteryServer/src/model"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)


func GetHistoryData(hist int) (error, []model.RBBall) {
	/* 向彩票官网拉数据
		ping www.cwl.gov.cn
		61.132.231.47

	http://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=ssq&issueCount=999&issueStart=&issueEnd=&dayStart=&dayEnd=
	*/
	url := fmt.Sprintf("http://www.cwl.gov.cn/cwl_admin/front/cwlkj/search/kjxx/findDrawNotice?name=ssq&issueCount=%d&issueStart=&issueEnd=&dayStart=&dayEnd=", hist)
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return err, nil
	}

	body, err := ioutil.ReadAll(res.Body)

	offRes := &model.OfficialRes{}
	err = json.Unmarshal(body, offRes)
	if err != nil {
		logrus.Errorf("convert official result failed: %s", err.Error())
		return err, nil
	}

	rawData := make([]model.RBBall, 0)
	for _, offres := range offRes.Result {
		rawData = append(rawData, model.RBBall{
			Red: offres.(map[string]interface{})["red"].(string),
			Blue: offres.(map[string]interface{})["blue"].(string),
		})
	}
	return nil, rawData
}
