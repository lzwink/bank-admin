package models

import (
	"errors"
	"log"
)

type Award struct {
	Id        int `orm:"pk"`
	AwardName string
	Num       int
	Surplus   int
}

func (a *Award) TableName() string {
	return "award"
}

// 抽中奖品后减少库存值
func (a *Award) SubNumById(id int) error {
	award := Award{Id: id}
	err := o.Read(&award)
	if err != nil {
		return err
	}
	if award.Surplus == 0 {
		log.Println("此奖品已抽完！")
		err = errors.New("此奖品已抽完！")
		return err
	}
	award.Surplus = award.Surplus - 1
	_, err = o.Update(&award, "Surplus")
	if err != nil {
		log.Println("更新奖品剩余数量错误：", err.Error())
	}
	return err
}

// 根据奖品id查询剩余数量
func (a *Award) GetSurplusById(id int) (int, error) {
	award := Award{Id: id}
	err := o.Read(&award)
	if err != nil {
		log.Println("根据奖品id查询剩余数量错误: ", err.Error())
		return 0, err
	}
	return award.Surplus, err
}

// 根据奖品id查询名称
func (a *Award) GetNameById(id int) (string, error) {
	award := Award{Id: id}
	err := o.Read(&award)
	if err != nil {
		log.Println("根据奖品id查询名称错误: ", err.Error())
		return "", err
	}
	return award.AwardName, err
}

// 每日更新库存数据
func (a *Award) CronData() (err error) {
	award := Award{}
	dataList := []int{3, 8, 15, 100}
	for k, v := range dataList {
		award.Id = k + 1
		award.Surplus = v
		_, err = o.Update(&award, "Surplus")
		if err != nil {
			log.Println("每日奖品数量更新错误：", err.Error())
		}
	}
	return err
}
