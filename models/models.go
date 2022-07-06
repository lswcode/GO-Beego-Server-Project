package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id   int
	Name string
	Age  int
	Addr string
}

func init() {
	orm.RegisterModel(new(User)) // 注册模型，所有的模型都要在初始化函数中注册后才能使用
}
