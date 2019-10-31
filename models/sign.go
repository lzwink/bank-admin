package models

/*
用户签到表
*/

import (
	"errors"
	"log"
)

type Sign struct {
	Id       int `orm:"pk"`
	UserId   int
	SignTime string
}

func (s *Sign) TableName() string {
	return "sign"
}

// 创建签到记录
func (s *Sign) CreateSign(userId int, signTime string) (err error) {
	sign := Sign{UserId: userId, SignTime: signTime}
	created, _, err := o.ReadOrCreate(&sign, "UserId", "SignTime")
	if err != nil {
		log.Println("未查询到指定日期的签到记录：", err.Error())
	}
	if !created {
		log.Println("Id为", userId, "的用户当日已签到")
		err = errors.New("已签到")
	}
	return err
}
