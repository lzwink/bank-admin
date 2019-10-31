package models

import (
	"github.com/astaxie/beego/orm"
	"log"
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

func (u *Users) GetUserById(id int) (Users, error) {
	result := Users{Id: id}
	err := o.Read(&result)
	if err == orm.ErrNoRows {
		log.Println("根据id查询用户错误")
	} else if err == orm.ErrMissPK {
		log.Println("根据id查询用户错误: 找不到主键")
	}
	return result, err
}

func (u *Users) GetUserByName(userName string) (Users, error) {
	result := Users{UserName: userName}
	err := o.Read(&result, "UserName")
	if err == orm.ErrNoRows {
		log.Println("根据用户名查询用户错误")
	} else if err == orm.ErrMissPK {
		log.Println("根据用户名查询用户错误：找不到主键")
	}
	return result, err
}

func (u *Users) GetUserByGroupId(groupId int) ([]Users, error) {
	result := make([]Users, 0)
	objectTable := o.QueryTable(new(Users))
	fieldStr := make([]string, 0)
	_, err := objectTable.RelatedSel().Filter("group_id", groupId).All(&result, fieldStr...)
	if err != nil {
		log.Println("GetUserByGroupId error: " + err.Error())
	}
	return result, err
}

func (u *Users) GetUsersByCombatGroupId(combatGroupId int) ([]Users, error) {
	result := make([]Users, 0)
	objectTable := o.QueryTable(new(Users))
	fieldStr := make([]string, 0)
	_, err := objectTable.RelatedSel().Filter("combat_group_id", combatGroupId).All(&result, fieldStr...)
	if err != nil {
		log.Println("GetUserByCombatGroupId error: " + err.Error())
	}
	return result, err
}

func (u *Users) CheckUserPwd(userName string, userPwd string) error {
	result := Users{UserName: userName, Pwd: userPwd}
	err := o.Read(&result, "UserName", "Pwd")
	if err != nil {
		log.Println("CheckUserPwd error: " + err.Error())
		return err
	}
	return err
}

func (u *Users) UpdateUserPwd(userId int, userName string, oldPwd string, newPwd string) error {
	user := Users{Id: userId, UserName: userName, Pwd: oldPwd}
	err := o.Read(&user, "UserName", "Pwd")
	if err == nil {
		user.Pwd = newPwd
		if _, err := o.Update(&user, "Pwd"); err == nil {
			return err
		}
		log.Println("更新用户密码错误：", err.Error())
	}
	log.Println("用户原密码错误：", err.Error())
	return err
}
