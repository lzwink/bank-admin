package models

import "log"

type Users struct {
	Id       int
	UserName string
}

func (u *Users) TableName() string {
	return "users"
}
func (u *Users) GetAllUsers() []Users {
	var result []Users
	objectTable := o.QueryTable(new(Users))
	fieldStr := []string{}
	_, err := objectTable.RelatedSel().All(&result, fieldStr...)
	if err != nil {
		log.Println("error: " + err.Error())
	}
	return result
}
