package main

import (
	_ "bank-admin/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
	// TODO: 定时任务，每天更新用户签到状态
	// TODO: 定时任务，每周一更新用户抽奖状态
}
