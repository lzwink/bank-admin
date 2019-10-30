package routers

import (
	"github.com/astaxie/beego"
	"project1/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
