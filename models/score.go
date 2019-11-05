package models

import (
	"errors"
	"log"
)

type Score struct {
	Id               int `orm:"pk"`
	UserName         string
	EventOneTarget   float64
	EventTwoTarget   float64
	EventThreeTarget float64
	EventFourTarget  float64
	EventOneReal     float64
	EventTwoReal     float64
	EventThreeReal   float64
	EventFourReal    float64
	Score            float64
}

func (s *Score) TableName() string {
	return "score"
}

func (s *Score) AddTargetScore(userName string, eOne float64, eTwo float64, eThree float64, eFour float64) error {
	score := Score{
		UserName:         userName,
		EventOneTarget:   eOne,
		EventTwoTarget:   eTwo,
		EventThreeTarget: eThree,
		EventFourTarget:  eFour,
	}
	created, _, err := o.ReadOrCreate(&score, "UserName")
	if err != nil {
		log.Println("创建得分目标错误：", err.Error())
	}
	if !created {
		log.Println("该用户已填写得分目标")
		err = errors.New("已填写")
	}
	return err
}

func (s *Score) AddRealScore(userName string, eOne float64, eTwo float64, eThree float64, eFour float64) error {
	score := Score{UserName: userName}
	err := o.Read(&score, "user_name")
	if err != nil {
		log.Println("查询用户得分信息错误")
		return err
	}
	score.EventOneReal = eOne
	score.EventTwoReal = eTwo
	score.EventThreeReal = eThree
	score.EventFourReal = eFour
	a := (eOne / score.EventOneTarget) * 0.4
	b := ((eTwo / score.EventTwoTarget) + (eThree / score.EventThreeTarget)) * 0.15
	c := (eFour / score.EventFourTarget) * 0.3
	score.Score = a + b + c
	_, err = o.Update(&score, "event_one_real", "event_two_real", "event_three_real", "event_four_real", "score")
	if err != nil {
		log.Println("更新数据错误")
	}
	return err
}

func (s *Score) GetScoreByUserName(userName string) (Score, error) {
	score := Score{UserName: userName}
	err := o.Read(&score, "user_name")
	if err != nil {
		log.Println("查找错误：", err.Error())
	}
	return score, err
}
