package models

/*
数据库操作
*/
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	o orm.Ormer
)

func init() {
	orm.Debug = true
	err := orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		log.Println("数据库驱动注册错误：", err.Error())
	}
	// set default database
	bankSource := beego.AppConfig.String("username") +
		":" + beego.AppConfig.String("password") +
		"@tcp(" + beego.AppConfig.String("host") + ")/" +
		beego.AppConfig.String("database") + "?charset=utf8"
	err = orm.RegisterDataBase("default", "mysql", bankSource, 10, 30)
	if err != nil {
		log.Println("数据库注册错误：", err.Error())
	}

	orm.RegisterModel(new(Users))
	orm.RegisterModel(new(Sign))
	orm.RegisterModel(new(Draw))

	//default model
	o = orm.NewOrm()
	err = o.Using("default")
	if err != nil {
		log.Println("默认数据库实例化错误：", err.Error())
	}
}
