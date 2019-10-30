package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hsw409328/bluemonday"
	"html"
	"strings"
)

var (
	removeDivArr = []string{"</div>", "</DIV>"}
	xssFilter    = bluemonday.UGCPolicy()
)

type BaseController struct {
	beego.Controller
}

func (ctx *BaseController) JsonEncode(code int, error_msg string, data interface{}, count int) {
	ctx.Data["json"] = map[string]interface{}{"code": code, "error_msg": error_msg, "msg": error_msg, "data": data, "count": count}
	ctx.ServeJSON()
	ctx.StopRun()
}

func (ctx *BaseController) XssFilter(str string) string {
	str = ctx.DivFilter(str)
	return html.EscapeString(xssFilter.Sanitize(str))
}

func (ctx *BaseController) DivFilter(str string) string {
	// 仅删除结束标签，开始标签xssFilter会删除
	for _, v := range removeDivArr {
		str = strings.Replace(str, v, "", -1)
	}
	return str
}
