package models

type UsersGroup struct {
	Id   int `orm:"pk"`
	Pid  int
	Name string
}

func (u *UsersGroup) TableName() string {
	return "users_group"
}
