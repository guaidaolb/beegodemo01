package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.tpl"
}

func (c *LoginController) DoLogin() {
	//c.Redirect("/", 302)
	c.Ctx.Redirect(302, "/")
}
