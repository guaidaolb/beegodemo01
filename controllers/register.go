package controllers

import (
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "register.tpl"
}

func (c *RegisterController) DoRegister() {
	c.TplName = "success.tpl"
}
