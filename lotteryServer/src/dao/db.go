package dao

import (
	"LotteryServer/src/model"
	"LotteryServer/src/utils"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db = &gorm.DB{}

func DBInit() {
	db = connectDatabase(connCfgInfo())
	if db == nil {
		return
	}
	if !createTable() {
		return
	}
}

/*
为普通用户创建数据库：
1、先以root用户登录
2、执行grant all privileges on ${database}.${table} to '${user_name}'@'%'
在${database}.${table}上，为所有IP的${user_name}授权所有权限
*/

func connCfgInfo() string {
	// TODO 后续从文件读取，文件要加密
	connCfg := model.DBConn{
		User:   "lottery",
		Pawd:   "123",
		Host:   "127.0.0.1",
		Port:   3306,
		DbName: "business",
	}
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		connCfg.User, connCfg.Pawd, connCfg.Host, connCfg.Port, connCfg.DbName)
}

func createTable() bool {
	mig := db.Migrator()
	if !mig.HasTable(&model.TblDoubleColorHist{}) {
		if err := mig.CreateTable(&model.TblDoubleColorHist{}); err != nil {
			logrus.Errorf("create table double color hist failed: %s", err.Error())
			return false
		}
	}
	if !mig.HasTable(&model.TblDoubleColorTotal{}) {
		if err := mig.CreateTable(&model.TblDoubleColorTotal{}); err != nil {
			logrus.Errorf("create table double color total failed: %s", err.Error())
			return false
		}
	}
	return true
}

func connectDatabase(conn string) *gorm.DB {
	// TODO 测试暂用 lottery@任意IP:123
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	if err != nil {
		logrus.Errorf("connect database failed: %s", err.Error())
		return nil
	}
	return db
}

func GetDoubleColorNewest() *model.TblDoubleColorHist {
	dbnew := &model.TblDoubleColorHist{}
	db.Last(dbnew)
	return dbnew
}

// data由近到远排序
func BatchInsertDoubleColorHistDB(data []model.DoubleColorBall) {
	dcHists := make([]model.TblDoubleColorHist, len(data))
	for i := 0; i < len(data); i++ {
		dcHists[i] = model.TblDoubleColorHist{
			Date: data[i].Time,
			Red:  data[i].Red,
			Blue: data[i].Blue,
		}
	}
	rowsAff := db.CreateInBatches(dcHists, len(dcHists)).RowsAffected
	if rowsAff == 0 {
		return
	}

	// 需要统计成功写入的条数，然后在更新到Total表里
	// 默认策略写入成功的都是离当前时间最近的
	toReds, toBlues := GetDoubleColorCnts()
	for _, datum := range data[:rowsAff] {
		reds, blues := utils.ConvertRedBallStr(datum.Red), utils.ConvertBlueBallStr(datum.Blue)
		utils.AcculateRedBalls(&(toReds.RSlots), reds)
		utils.AcculateBlueBalls(&(toBlues.BSlots), blues)
	}
	totals := &model.TblDoubleColorTotal{utils.Convert2RedBallStr(toReds.RSlots), utils.Convert2BlueBallStr(toBlues.BSlots)}
	db.Updates(totals)
}

func GetDoubleColorCnts() (*model.DoubleColorRed, *model.DoubleColorBlue) {
	balls := &model.TblDoubleColorTotal{}
	db.First(balls)
	redBallsCnt := utils.ConvertRedBallStr(balls.Reds)
	blueBallsCnt := utils.ConvertBlueBallStr(balls.Blues)
	return &model.DoubleColorRed{redBallsCnt}, &model.DoubleColorBlue{blueBallsCnt}
}

func UpdateDoubleColorCntData(redBallsCnts *model.DoubleColorRed, blueBallCnt *model.DoubleColorBlue) {
	redsStr, bluesStr := utils.Convert2RedBallStr(*&redBallsCnts.RSlots), utils.Convert2BlueBallStr(*&blueBallCnt.BSlots)
	db.Updates(model.TblDoubleColorTotal{redsStr, bluesStr})
}
