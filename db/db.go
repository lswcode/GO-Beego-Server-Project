package db

import (
	_ "beego_server/models" // 导入项目内部的包时，使用项目名开头
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 第一步 获取app.conf中的数据库相关配置
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	database := beego.AppConfig.String("database")

	// -------------------------------------------------------
	// 第二步  连接数据库
	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=Local", username, password, host, port, database) // 最后一个参数是操作的数据库名字，在app.conf中配置
	orm.RegisterDriver("mysql", orm.DRMySQL)                                                                          // 第一个参数是数据库名称，第二个参数是数据库类型，使用orm包中的mysql类型
	connectErr := orm.RegisterDataBase("default", "mysql", datasource, 30, 30)                                        // 开始连接数据库
	// 第一个参数是自定义的数据库的别名，第二个参数是orm.RegisterDriver中的第一个参数，第三个参数是数据库连接参数
	// ORM必须注册一个别名为default的数据库，作为默认使用
	// 第四个和第五个参数可以省略，设置最大空闲连接和最大数据库连接
	if connectErr != nil {
		logs.Error("数据库连接失败:", connectErr)
	}

	// 第三步 在main.go启动db.go
	// -------------------------------------------------------
	// beego的orm运行参数
	name := "default" // 操作的数据库的别名
	force := false    // 强制创建表，如果表重名，先删除旧的再创建表，一般设置为false
	// go run beego_rom.go orm sqlall   显示orm创建的表时对应的sql语句
	verbose := true                                        // 打印执行的过程，即自动执行 go run beego_rom.go orm sqlall
	ormRunSyncdbErr := orm.RunSyncdb(name, force, verbose) // 设置orm运行参数

	if ormRunSyncdbErr != nil {
		logs.Error("ORM配置失败:", ormRunSyncdbErr)
		panic(ormRunSyncdbErr)
	}

}
