package controllers

import (
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("api接口---" + id)
}
