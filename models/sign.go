package models

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

func (s *Sign) CreateSign(userId int, signTime string) (err error) {
	sign := Sign{UserId: userId, SignTime: signTime}
	created, _, err := o.ReadOrCreate(&sign, "UserId", "SignTime")
	if err != nil {
		log.Println("查询历史签到错误：", err.Error())
		return err
	}
	if !created {
		log.Println("Id为", userId, "的用户当日已签到")
		err = errors.New("已签到")
		return err
	}
	return err
}
