package models

/*
用户基本信息表
*/

import (
	"log"
	rand2 "math/rand"
	"time"
)

type Users struct {
	Id            int `orm:"pk"`
	UserName      string
	Pwd           string
	AuthGroupId   int
	GroupId       int
	CombatGroupId int
	Photo         string
	IsDraw        int
	IsSign        int
}

func (u *Users) TableName() string {
	return "users"
}

// 获取全部用户信息
func (u *Users) GetAllUsers() ([]Users, error) {
	result := make([]Users, 0)
	objectTable := o.QueryTable(new(Users))
	fieldStr := make([]string, 0)
	_, err := objectTable.RelatedSel().All(&result, fieldStr...)
	if err != nil {
		log.Println("获取全部用户信息错误: " + err.Error())
	}
	return result, err
}

// 根据用户id获取单个用户
func (u *Users) GetUserById(id int) (Users, error) {
	result := Users{Id: id}
	err := o.Read(&result)
	if err != nil {
		log.Println("根据用户id获取单个用户错误：", err.Error())
	}
	return result, err
}

// 根据用户名称获取单个用户
func (u *Users) GetUserByName(userName string) (Users, error) {
	result := Users{UserName: userName}
	err := o.Read(&result, "UserName")
	if err != nil {
		log.Println("根据用户名称获取单个用户错误：", err.Error())
	}
	return result, err
}

// 根据用户id更新用户抽奖状态
func (u *Users) UpdateUserIsDraw(userId int) error {
	result := Users{Id: userId}
	err := o.Read(&result)
	if err != nil {
		log.Println("根据用户Id读取信息错误：", err.Error())
		return err
	}
	result.IsDraw = 1
	_, err = o.Update(&result, "IsDraw")
	if err != nil {
		log.Println("更新用户抽奖状态错误：", err.Error())
	}
	return err
}

// 根据用户id更新用户签到状态
func (u *Users) UpdateUserIsSign(userId int) error {
	result := Users{Id: userId}
	err := o.Read(&result)
	if err != nil {
		log.Println("根据用户Id读取信息错误：", err.Error())
		return err
	}
	result.IsSign = 1
	_, err = o.Update(&result, "IsSign")
	if err != nil {
		log.Println("更新用户签到状态错误：", err.Error())
	}
	return err
}

// 根据分组编号获取用户列表
func (u *Users) GetUserByGroupId(groupId int) ([]Users, error) {
	result := make([]Users, 0)
	objectTable := o.QueryTable(new(Users))
	fieldStr := make([]string, 0)
	_, err := objectTable.RelatedSel().Filter("group_id", groupId).All(&result, fieldStr...)
	if err != nil {
		log.Println("根据分组编号获取用户列表错误: " + err.Error())
	}
	return result, err
}

// 根据对抗编号获取用户列表
func (u *Users) GetUsersByCombatGroupId(combatGroupId int) ([]Users, error) {
	result := make([]Users, 0)
	objectTable := o.QueryTable(new(Users))
	fieldStr := make([]string, 0)
	_, err := objectTable.RelatedSel().Filter("combat_group_id", combatGroupId).All(&result, fieldStr...)
	if err != nil {
		log.Println("根据对抗编号获取用户列表错误: " + err.Error())
	}
	return result, err
}

// 验证用户名密码
func (u *Users) CheckUserPwd(userName string, userPwd string) error {
	result := Users{UserName: userName, Pwd: userPwd}
	err := o.Read(&result, "UserName", "Pwd")
	if err != nil {
		log.Println("验证用户名密码错误: " + err.Error())
	}
	return err
}

// 更新用户名密码
func (u *Users) UpdateUserPwd(userId int, userName string, oldPwd string, newPwd string) error {
	user := Users{Id: userId, UserName: userName, Pwd: oldPwd}
	err := o.Read(&user, "UserName", "Pwd")
	if err != nil {
		log.Println("用户原密码错误：", err.Error())
		return err
	}
	user.Pwd = newPwd
	_, err = o.Update(&user, "Pwd")
	if err != nil {
		log.Println("更新用户密码错误：", err.Error())
		return err
	}
	return err
}

// 批量生成用户账户信息
//func (u *Users) CreateUsersList() {
//	userName := u.RandomStr(6)
//	pwd := u.RandomStr(6)
//
//
//}
//
//func (u *Users) RandomStr(l int) string {
//	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
//	bytes := []byte(str)
//	result := make([]byte, 0)
//	r := rand2.New(rand2.NewSource(time.Now().UnixNano()))
//	for i := 0; i < l; i++ {
//		result = append(result, bytes[r.Intn(len(bytes))])
//	}
//	return string(result)
//}
