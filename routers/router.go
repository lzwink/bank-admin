package routers

import (
	"bank-admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.UsersController{})
	beego.Router("/GetAllUsers", &controllers.UsersController{}, "get:GetAllUsers")
	beego.Router("/GetUserById", &controllers.UsersController{}, "post:GetUserById")

	beego.Router("/Login", &controllers.LoginController{}, "get:Login")
	//beego.Router("/Check",&controllers.LoginController{},"get:Check")
	beego.Router("/Logout", &controllers.LoginController{}, "get:Logout")
}
