package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/hsw409328/bluemonday"
	"gopkg.in/mgo.v2/bson"
	"html"
	"strconv"
	"strings"
	"time"
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
func (ctx *BaseController) InterfaceToStr(inter interface{}) (s string) {
	tempStr := ""
	switch inter.(type) {
	case nil:
		tempStr = ""
		break
	case string:
		tempStr = inter.(string)
		break
	case float64:
		tempStr = strconv.FormatFloat(inter.(float64), 'f', -1, 64)
		break
	case int64:
		tempStr = strconv.FormatInt(inter.(int64), 10)
		break
	case int:
		tempStr = strconv.Itoa(inter.(int))
		break
	case bool:
		tempStr = strconv.FormatBool(inter.(bool))
	case bson.ObjectId:
		tempStr = inter.(bson.ObjectId).Hex()
	case []interface{}:
		tempStr, _ = ctx.JsonToString(inter)
	case []int:
		tempStr, _ = ctx.JsonToString(inter)
	case []int64:
		tempStr, _ = ctx.JsonToString(inter)
	case []float32:
		tempStr, _ = ctx.JsonToString(inter)
	case []float64:
		tempStr, _ = ctx.JsonToString(inter)
	case map[string]interface{}:
		tempStr, _ = ctx.JsonToString(inter)
	case map[string]string:
		tempStr, _ = ctx.JsonToString(inter)
	case time.Time:
		tempStr = inter.(time.Time).String()
	default:
		tempStr = "Error! Not Found Type!"
	}
	return tempStr
}
func (ctx *BaseController) JsonToString(inter interface{}) (string, error) {
	by, err := json.Marshal(inter)
	if err != nil {
		return "", err
	} else {
		return string(by), nil
	}
}
