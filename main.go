package main

import (
	_ "test_server/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
