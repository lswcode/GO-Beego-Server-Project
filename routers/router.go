package routers

import (
	"test_server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.TestController{})
	beego.Router("/abc", &controllers.TestController{})
}
