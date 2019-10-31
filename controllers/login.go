package controllers

type LoginController struct {
	BaseController
}

func (ctx *LoginController) Login() {
	name := ctx.XssFilter(ctx.GetString("name"))
	pwd := ctx.XssFilter(ctx.GetString("pwd"))
	err := userModel.CheckUserPwd(name, pwd)
	if err != nil {
		ctx.JsonEncode(100, "failed", nil, 0)
	}
	ctx.SetSession("user_name", name)
	ctx.JsonEncode(0, "success", nil, 0)
}

func (ctx *LoginController) Logout() {
	ctx.DelSession("user_name")
	ctx.JsonEncode(0, "success", nil, 0)
}

func (ctx *LoginController) Check() {
	userName := ctx.GetSession("user_name")
	if userName == nil {
		ctx.JsonEncode(100, "failed", nil, 0)
	}
	ctx.JsonEncode(0, "success", nil, 0)
}
