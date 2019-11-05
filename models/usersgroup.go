package models

import "log"

type UsersGroup struct {
	Id   int `orm:"pk"`
	Pid  int
	Name string
}

func (u *UsersGroup) TableName() string {
	return "users_group"
}

func (u *UsersGroup) GetAllUsersGroup() ([]UsersGroup, error) {
	result := make([]UsersGroup, 0)
	objectTable := o.QueryTable(new(UsersGroup))
	fieldStr := make([]string, 0)
	_, err := objectTable.RelatedSel().All(&result, fieldStr...)
	if err != nil {
		log.Println("获取用户分组列表错误: " + err.Error())
	}
	return result, err
}
