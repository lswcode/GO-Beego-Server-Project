package main

import (
	_ "beego_server/routers"

	"github.com/astaxie/beego"

	_ "beego_server/db" // 启动数据库

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	orm.RunCommand() // 运行orm，数据库使用orm的前提

	// 开启cors跨域
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//允许访问所有源
		AllowAllOrigins: true,
		//可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		//其中Options跨域复杂请求预检
		AllowMethods: []string{"*"},
		//指的是允许的Header的种类
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "content-type"},
		//公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin"},
		//如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))

	logs.SetLogger(logs.AdapterMultiFile, // 日志分类别，多文件写入，各个级别不同的日志生成单独的文件
		`{"filename":"logs/beego_server.log","separate":
		 [ "error", "warning","notice","info","debug"]}`)

	beego.Run() // 所有的配置都要写在beego.Run()执行之前
}
