package main

import (
	"time"
	"net/http"
	"log"
	_"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
	"encoding/json"
	"tingting-sync/var"
	"tingting-sync/models"
)

func main() {
	RunTimer()
}
func RunTimer() {
	GetDashboardTimer := time.NewTicker(20 * time.Second)
	for {
		select {
		case <-GetDashboardTimer.C:
			GetDashboardData()
		}
	}
}
func GetDashboardData() {
	go func() {
         //今日新增用户
		db,err := utils.OpenConnection()
		if err!=nil{
			fmt.Println(err.Error())
		}
		defer db.Close()
		var number_newuser int64
		db.Raw("select count(*) from members where to_days(ts) = to_days(now())").Count(&number_newuser)

         //今日购买专辑
         var number_today_buy_albums int64
         db.Raw("select count(*) from orders  where to_days(ts) = to_days(now())").Count(&number_today_buy_albums)

         //今日成交额
         var money_today int64
         db.Raw("select sum(value) from orders where to_days(ts) = to_days(now())").Count(&money_today)

         //今日上传录音
         var number_today_record int64
         db.Raw("select count(*) from ")

         //月内购买专辑数

         //月内成交额

         //近期热门

         //今日新增分类数

         //今日新增专辑数

         //今日新增书本数

         //今日新增章节数

	}()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}