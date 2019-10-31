package models

/*
用户抽奖记录表
*/
import (
	"errors"
	"log"
)

type Draw struct {
	Id         int `orm:"pk"`
	UserId     int
	AwardId    int
	CreateTime string
}

func (d *Draw) TableName() string {
	return "draw"
}

func (d *Draw) CreateDraw(userId int, awardId int, createTime string) error {
	// 获取用户当前信息
	userModel := Users{}
	user, err := userModel.GetUserById(userId)
	if err != nil {
		log.Println("抽奖前查询用户信息错误，错误用户的id为：", userId, " error：", err.Error())
		return err
	}
	// 查看是否已抽奖
	if user.IsDraw == 1 {
		log.Println("Id为", userId, "的用户该周已抽奖")
		err = errors.New("已抽奖")
		return err
	}
	// 添加抽奖记录
	draw := Draw{UserId: userId, CreateTime: createTime, AwardId: awardId}
	_, err = o.Insert(&draw)
	if err != nil {
		log.Println("插入抽奖记录错误：", err.Error())
	}
	return err
}
