package routers

import (
	"bank-admin/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.UsersController{})
	beego.Router("/GetAllUsers", &controllers.UsersController{}, "get:GetAllUsers")
	beego.Router("/GetUserById", &controllers.UsersController{}, "get:GetUserById")
	beego.Router("/UpdateUserPwd", &controllers.UsersController{}, "get:UpdateUserPwd")

	beego.Router("/Login", &controllers.LoginController{}, "get:Login")
	beego.Router("/Check", &controllers.LoginController{}, "get:Check")
	beego.Router("/Logout", &controllers.LoginController{}, "get:Logout")

	beego.Router("/CreateSign", &controllers.SignController{}, "get:CreateSign")

	beego.Router("/CreateDraw", &controllers.DrawController{}, "get:CreateDraw")
}
