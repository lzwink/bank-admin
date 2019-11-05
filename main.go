package main

import (
	"bank-admin/models"
	_ "bank-admin/routers"
	"github.com/astaxie/beego"
	"log"
	"strings"
	"time"
)

func main() {
	userModel := models.Users{}
	awardModel := models.Award{}
	go func() {
		for {
			t := time.Now()
			if int(t.Weekday()) == 1 {
				log.Println("开始执行每周更新 【开始】：", time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
				userModel.CronDrawData()
				log.Println("开始执行每周更新 【结束】：", time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
				time.Sleep(167 * time.Hour)
			} else {
				time.Sleep(1 * time.Minute)
			}
		}
	}()
	go func() {
		for {
			timeLayout := "2006-01-02 15:04:05"
			dateTimeStr := time.Unix(time.Now().Unix(), 0).Format(timeLayout)
			dateArr := strings.Split(dateTimeStr, " ")
			timeArr := strings.Split(dateArr[1], ":")
			if (timeArr[0] == "00") && (timeArr[1] == "00") {
				log.Println("开始执行每日更新 【开始】：", time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
				userModel.CronSignData()
				awardModel.CronData()
				log.Println("开始执行每日更新 【结束】：", time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
				time.Sleep(23 * time.Hour)
			} else {
				time.Sleep(1 * time.Second)
			}
		}
	}()
	beego.Run()

}
