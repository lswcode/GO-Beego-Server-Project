package routers

import (
	"beego_server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	// beego.Router("/login", &controllers.UserAuthController{}) // 一个固定的路由，一个控制器，根据不同的请求方式，自动对应控制器中不同的方法
	beego.Router("/login", &controllers.UserAuthController{}, "Post:Login")
	beego.Router("/register", &controllers.UserAuthController{}, "Post:Register")
}
