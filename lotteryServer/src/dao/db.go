package dao

import (
	"LotteryServer/src/model"
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
	connCfg := model.Conn{
		User: "lottery",
		Pawd: "123",
		Host: "127.0.0.1",
		Port: 3306,
		DbName: "business",
	}
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		connCfg.User ,connCfg.Pawd, connCfg.Host, connCfg.Port, connCfg.DbName)
}


func createTable() bool {
	mig := db.Migrator()
	if !mig.HasTable(&model.TblDoubleColorHist{}) {
		if err := mig.CreateTable(&model.TblDoubleColorHist{}); err != nil {
			logrus.Errorf("create table double color hist failed: %s", err.Error())
			return false
		}
	}
	if !mig.HasTable(&model.TblDoubleColorRed1{}) {
		if err := mig.CreateTable(&model.TblDoubleColorRed1{}); err != nil {
			logrus.Errorf("create table double red1 balls failed: %s", err.Error())
			return false
		}
	}
	if !mig.HasTable(&model.TblDoubleColorRed2{}) {
		if err := mig.CreateTable(&model.TblDoubleColorRed2{}); err != nil {
			logrus.Errorf("create table double red2 balls failed: %s", err.Error())
			return false
		}
	}
	if !mig.HasTable(&model.TblDoubleColorRed3{}) {
		if err := mig.CreateTable(&model.TblDoubleColorRed3{}); err != nil {
			logrus.Errorf("create table double red3 balls failed: %s", err.Error())
			return false
		}
	}
	if !mig.HasTable(&model.TblDoubleColorRed4{}) {
		if err := mig.CreateTable(&model.TblDoubleColorRed4{}); err != nil {
			logrus.Errorf("create table double red4 balls failed: %s", err.Error())
			return false
		}
	}
	if !mig.HasTable(&model.TblDoubleColorRed5{}) {
		if err := mig.CreateTable(&model.TblDoubleColorRed5{}); err != nil {
			logrus.Errorf("create table double red5 balls failed: %s", err.Error())
			return false
		}
	}
	if !mig.HasTable(&model.TblDoubleColorRed6{}) {
		if err := mig.CreateTable(&model.TblDoubleColorRed6{}); err != nil {
			logrus.Errorf("create table double red6 balls failed: %s", err.Error())
			return false
		}
	}
	if !mig.HasTable(&model.TblDoubleColorBlue{}) {
		if err := mig.CreateTable(&model.TblDoubleColorBlue{}); err != nil {
			logrus.Errorf("create table double blue balls failed: %s", err.Error())
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

func BatchInsertDoubleColorHistData(data []model.DoubleColorBall) {
	dcHists := make([]model.TblDoubleColorHist, len(data))
	for i := 0; i < len(data); i++ {
		dcHists[i] = model.TblDoubleColorHist{
			Date: data[i].Time,
			Red: data[i].Red,
			Blue: data[i].Blue,
		}
	}
	db.CreateInBatches(dcHists, len(dcHists))
}

func GetDoubleColorCnts() ([]*model.DoubleColorRed, *model.TblDoubleColorBlue) {
	red1, red2 := &model.TblDoubleColorRed1{}, &model.TblDoubleColorRed2{}
	red3, red4 := &model.TblDoubleColorRed3{}, &model.TblDoubleColorRed4{}
	red5, red6 := &model.TblDoubleColorRed5{}, &model.TblDoubleColorRed6{}
	blue:= &model.TblDoubleColorBlue{}
	db.First(red1); db.First(red2); db.First(red3)
	db.First(red4);	db.First(red5);	db.First(red6)
	db.First(blue)
	return []*model.DoubleColorRed{
		&(red1.DoubleColorRed), &(red1.DoubleColorRed),
		&(red1.DoubleColorRed), &(red1.DoubleColorRed),
		&(red1.DoubleColorRed)}, blue
}

func UpdateDoubleColorCntData(redBallsCnts []*model.DoubleColorRed, blueBallCnt *model.TblDoubleColorBlue) {
	db.Updates(redBallsCnts[0])
	db.Updates(redBallsCnts[1])
	db.Updates(redBallsCnts[2])
	db.Updates(redBallsCnts[3])
	db.Updates(redBallsCnts[4])
	db.Updates(redBallsCnts[5])
	db.Updates(blueBallCnt)
}