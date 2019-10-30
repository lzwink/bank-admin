package models

import "log"

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
		log.Println("GetAllUsers error: " + err.Error())
	}
	return result, err
}

func (u *Users) GetUserById(id int) (Users, error) {
	result := Users{}
	objectTable := o.QueryTable(new(Users))
	err := objectTable.RelatedSel().Filter("id", id).One(&result)
	if err != nil {
		log.Println("GetUserById error: " + err.Error())
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

func (u *Users) GetUserByCombatGroupId(combatGroupId int) ([]Users, error) {
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
	result := Users{}
	objectTable := o.QueryTable(new(Users))
	err := objectTable.RelatedSel().Filter("user_name", userName).Filter("pwd", userPwd).One(&result)
	if err != nil {
		log.Println("CheckUserPwd error: " + err.Error())
		return err
	}
	return err
}

//func (u *Users) UpdateUserPwd(oldPwd string, newPwd string) error {
//
//}
