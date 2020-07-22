package routers

import (
	"beegodemo01/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/goods", &controllers.GoodsController{})
	beego.Router("/goods/xml", &controllers.GoodsController{}, "post:Xml")

	beego.Router("/article", &controllers.ArticleController{})
	beego.Router("/article/add", &controllers.ArticleController{}, "get:AddArticle")
	beego.Router("/article/edit", &controllers.ArticleController{}, "get:EditArticle")

	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/add", &controllers.UserController{}, "get:AddUser")
	beego.Router("/user/doAdd", &controllers.UserController{}, "post:DoAddUser")

	beego.Router("/api/:id", &controllers.ApiController{})

	beego.Router("/cms_:id([0-9]+).html", &controllers.CmsController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/doLogin", &controllers.LoginController{}, "post:DoLogin")

	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/doRegister", &controllers.RegisterController{}, "post:DoRegister")

}
