package models

/*
用户签到表
*/

import (
	"errors"
	"log"
	"time"
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
		log.Println("创建签到记录错误：", err.Error())
	}
	if !created {
		log.Println("Id为", userId, "的用户当日已签到")
		err = errors.New("已签到")
	}
	return err
}

// 根据用户id查询已签到天数
func (s *Sign) CheckSignDays(userId int) (int, error) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday := weekStartDate.Format("2006-01-02")
	result := make([]Sign, 0)
	objectTable := o.QueryTable(new(Sign))
	fieldStr := make([]string, 0)
	_, err := objectTable.RelatedSel().Filter("sign_time__gte", weekMonday).All(&result, fieldStr...)
	if err != nil {
		log.Println("查询签到记录错误：", err.Error())
	}
	return len(result), err
}
