package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"

	c.Ctx.WriteString("用户中心")
}
func (c *UserController) AddUser() {

	c.TplName = "user.html"

}
func (c *UserController) DoAddUser() {
	//c.Ctx.WriteString("用户中心")
	username := c.GetString("username")
	password := c.GetString("password")
	c.Ctx.WriteString("欢迎用户:" + username + "...你的密码是:" + password)
}
