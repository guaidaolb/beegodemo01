package controllers

import (
	"encoding/xml"
	"fmt"

	"github.com/astaxie/beego"
)

type GoodsController struct {
	beego.Controller
}

func (c *GoodsController) Get() {

	//获取cookie
	username := c.Ctx.GetCookie("username")
	c.Data["username"] = username

	c.Data["title"] = "你好beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "goods.tpl"
}

type Product struct {
	Title   string `form:"title" xml:"title"`
	Content string `form:"content" xml:"content"`
}

func (c *GoodsController) Xml() {

	p := Product{}
	// str := string(c.Ctx.Input.RequestBody)
	// beego.Info(str)
	// c.Ctx.WriteString(str)
	var err error
	if err1 := xml.Unmarshal(c.Ctx.Input.RequestBody, &p); err1 != nil {
		c.Data["json"] = err.Error()
		c.ServeJSON()
	} else {
		fmt.Printf("%#v", p)
		c.Data["json"] = p
		c.ServeJSON()
	}

}
