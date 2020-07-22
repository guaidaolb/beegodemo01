package main

import (
	"beegodemo01/models"
	_ "beegodemo01/routers"

	"github.com/astaxie/beego"
)

func main() {

	beego.AddFuncMap("hello", models.Hello)
	beego.AddFuncMap("unixToDate", models.UnixToDate)
	beego.AddFuncMap("dateToUnix", models.DateToUnix)

	//配置redis数据库的连接地址
	beego.AppConfig.Set("redisuser", "redisadmin")

	beego.BConfig.WebConfig.Session.SessionProvider = "redis"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "127.0.0.1:6379"

	beego.Run()

}
