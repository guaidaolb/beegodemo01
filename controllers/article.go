package controllers

import (
	"beegodemo01/models"
	"crypto/md5"
	"fmt"
	"io"
	"time"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.Ctx.WriteString("大威天龙")

	now := time.Now()
	fmt.Println(now)
	c.Data["now"] = now

	c.Data["title"] = "我是一个中国人"
	c.Data["html"] = "<h2>我爱我的祖国</h2>"

	userinfo := make(map[string]interface{})
	userinfo["username"] = "zhansan"
	userinfo["age"] = 20
	c.Data["userinfo"] = userinfo
	userinfo["hobby"] = map[int]string{
		1: "吃饭",
		2: "睡觉",
		3: "看片",
	}
	t := time.Now().Unix()
	fmt.Println(t) //1595393407
	// c.Data["unix"] = 1389058332
	// c.Data["unix"] = time.Now().UnixNano()
	c.Data["unix"] = int(t)

	var str = "中华人民共和国"
	slice := []rune(str)
	fmt.Println(slice) //[20013 21326 20154 27665 20849 21644 22269]
	// c.Data["str"] = string(slice[:4])
	c.Data["str"] = string(slice[2:6])

	c.Data["date"] = "1986-11-20 02:23:21"

	h := md5.New()
	io.WriteString(h, "19891127")
	fmt.Printf("%x\n", h.Sum(nil)) //71de3abc1b32b65dba38aaab946bac86

	data := []byte("19891127")
	fmt.Printf("%x\n", md5.Sum(data)) //71de3abc1b32b65dba38aaab946bac86

	fmt.Println(models.Md5("19861120")) //341BE97D9AFF90C9978347F66F945B77
	c.TplName = "article.tpl"

}
func (c *ArticleController) AddArticle() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.Ctx.WriteString("世尊地藏")
}
func (c *ArticleController) EditArticle() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.Ctx.WriteString("我一眼就看出你不是人")
}
