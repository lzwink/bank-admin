package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var (
	o orm.Ormer
)

func init() {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	// set default database
	bankSource := beego.AppConfig.String("username") +
		":" + beego.AppConfig.String("password") +
		"@tcp(" + beego.AppConfig.String("host") + ")/" +
		beego.AppConfig.String("database") + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", bankSource, 10, 30)

	orm.RegisterModel(new(Users))

	//default model
	o = orm.NewOrm()
	o.Using("default")
}
