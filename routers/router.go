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
	beego.Router("/GetOpponentList", &controllers.UsersController{}, "get:GetOpponentList")
	beego.Router("/ChooseOpponent", &controllers.UsersController{}, "get:ChooseOpponent")
	beego.Router("/AddTargetScore", &controllers.UsersController{}, "get:AddTargetScore")
	beego.Router("/AddRealScore", &controllers.UsersController{}, "get:AddRealScore")
	beego.Router("/GetLoginUser", &controllers.UsersController{}, "get:GetLoginUser")
	beego.Router("/GetUserScore", &controllers.UsersController{}, "get:GetUserScore")
	beego.Router("/GetOppScore", &controllers.UsersController{}, "get:GetOppScore")
	beego.Router("/GetOppInfo", &controllers.UsersController{}, "get:GetOppInfo")
	beego.Router("/GetAllUsersGroup", &controllers.UsersController{}, "get:GetAllUsersGroup")
	beego.Router("/GetUsersByGroupId", &controllers.UsersController{}, "get:GetUsersByGroupId")
	beego.Router("/GetAllInfoById", &controllers.UsersController{}, "get:GetAllInfoById")

	beego.Router("/Login", &controllers.LoginController{}, "get:Login")
	beego.Router("/Check", &controllers.LoginController{}, "get:Check")
	beego.Router("/Logout", &controllers.LoginController{}, "get:Logout")

	beego.Router("/CreateSign", &controllers.SignController{}, "get:CreateSign")

	beego.Router("/CreateDraw", &controllers.DrawController{}, "get:CreateDraw")

	beego.Router("/UpForm", &controllers.UploadController{}, "get:UpForm")
	beego.Router("/UpFile", &controllers.UploadController{}, "post:UpFile")
}
