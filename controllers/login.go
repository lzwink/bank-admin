package controllers

type LoginController struct {
	BaseController
}

func (ctx *LoginController) Login() {
	userName := ctx.XssFilter(ctx.GetString("userName"))
	pwd := ctx.XssFilter(ctx.GetString("pwd"))
	userInfo, err := userModel.CheckUserPwd(userName, pwd)
	if err != nil {
		ctx.JsonEncode(100, "failed", nil, 0)
	}
	ctx.SetSession("real_name", userInfo.RealName)
	ctx.SetSession("user_name", userInfo.UserName)
	ctx.SetSession("user_id", userInfo.Id)
	ctx.JsonEncode(0, "success", nil, 0)
}

func (ctx *LoginController) Logout() {
	ctx.DelSession("real_name")
	ctx.DelSession("user_name")
	ctx.DelSession("user_id")
	ctx.JsonEncode(0, "success", nil, 0)
}

func (ctx *LoginController) Check() {
	realName := ctx.GetSession("real_name")
	if realName == nil {
		ctx.JsonEncode(100, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", realName, 0)
}
