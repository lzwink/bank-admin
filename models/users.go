package models

/*
用户基本信息表
*/

import (
	"log"
)

type Users struct {
	Id          int `orm:"pk"`
	UserName    string
	RealName    string
	Pwd         string
	AuthGroupId int
	GroupId     int
	Photo       string
	IsDraw      int
	IsSign      int
	IsPriority  int
	OpponentId  int
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

// 根据用户账户称获取单个用户
func (u *Users) GetUserByUserName(userName string) (Users, error) {
	result := Users{UserName: userName}
	err := o.Read(&result, "UserName")
	if err != nil {
		log.Println("根据用户名称获取单个用户错误：", err.Error())
	}
	return result, err
}

// 根据用户真实姓名获取用户（可能为多个或一个）
func (u *Users) GetUserByRealName(realName string) ([]Users, error) {
	result := make([]Users, 0)
	objectTable := o.QueryTable(new(Users))
	fieldStr := make([]string, 0)
	_, err := objectTable.RelatedSel().Filter("real_name", realName).All(&result, fieldStr...)
	if err != nil {
		log.Println("根据用户真实姓名查找用户错误：", err.Error())
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

// 验证用户名密码
func (u *Users) CheckUserPwd(userName string, userPwd string) (Users, error) {
	result := Users{UserName: userName, Pwd: userPwd}
	err := o.Read(&result, "UserName", "Pwd")
	if err != nil {
		log.Println("验证用户名密码错误: " + err.Error())
	}
	return result, err
}

// 更新用户名密码
func (u *Users) UpdateUserPwd(userId int, oldPwd string, newPwd string) error {
	user := Users{Id: userId, Pwd: oldPwd}
	err := o.Read(&user, "Id", "Pwd")
	if err != nil {
		log.Println("用户原密码错误：", err.Error())
		return err
	}
	user.Pwd = newPwd
	_, err = o.Update(&user, "Pwd")
	if err != nil {
		log.Println("更新用户密码错误：", err.Error())
	}
	return err
}

// 根据用户id绑定对手id
func (u *Users) UpdateUserOpp(userId int, oppId int) error {
	// 更新当前用户信息
	user := Users{Id: userId}
	err := o.Read(&user)
	if err != nil {
		log.Println("查找用户信息错误：", err.Error())
		return err
	}
	user.OpponentId = oppId
	_, err = o.Update(&user, "OpponentId")
	if err != nil {
		log.Println("更新用户对手信息错误：", err.Error())
	}
	// 更新对手用户信息
	userOpp := Users{Id: oppId}
	err = o.Read(&userOpp)
	if err != nil {
		log.Println("查找对手id错误：", err.Error())
		return err
	}
	userOpp.OpponentId = userId
	_, err = o.Update(&userOpp, "OpponentId")
	if err != nil {
		log.Println("更新对手记录字段错误：", err.Error())
	}
	return err
}

// 每日更新签到数据
func (u *Users) CronSignData() error {
	user := Users{}
	cnt, err := o.QueryTable(new(Users)).Count()
	if err != nil {
		log.Println("统计用户数量错误：", err.Error())
		return err
	}
	for i := 1; i < int(cnt)-2; i++ {
		user.Id = i
		user.IsSign = 2
		_, err = o.Update(&user, "IsSign")
		if err != nil {
			log.Println("每日更新用户签到信息错误：", err.Error())
			return err
		}
	}
	return err
}

// 每周更新抽奖数据
func (u *Users) CronDrawData() error {
	user := Users{}
	cnt, err := o.QueryTable(new(Users)).Count()
	if err != nil {
		log.Println("统计用户数量错误：", err.Error())
		return err
	}
	for i := 1; i < int(cnt)-2; i++ {
		user.Id = i
		user.IsDraw = 2
		_, err = o.Update(&user, "IsDraw")
		if err != nil {
			log.Println("每周更新用户抽奖信息错误：", err.Error())
			return err
		}
	}
	return err
}
