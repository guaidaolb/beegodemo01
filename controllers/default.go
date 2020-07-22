package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	//设置cookie
	c.Ctx.SetCookie("username", "zhansan", 5)

	//通过.string获取的返回一个值
	mysqluser1 := beego.AppConfig.String("mysqluser") //root
	//通过.strings获取的返回一个切片
	mysqluser2 := beego.AppConfig.Strings("mysqluser") //[root]
	beego.Info(mysqluser1)
	beego.Info(mysqluser2)
	admin := beego.AppConfig.Strings("admin")
	beego.Info(admin)

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
