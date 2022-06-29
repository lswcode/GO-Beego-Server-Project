package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	beego.Controller
}

func (t *TestController) Get() {
	t.Ctx.WriteString("测试test")
}
