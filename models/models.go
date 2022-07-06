package models

import (
	"github.com/astaxie/beego/orm"
)

type UserLogin struct {
	Account  string `json:"account"` // 首字母必须大写，不然无法被JSON包解析
	Password string `json:"password"`
}

type UserRegister struct {
	Name     string
	Account  string `json:"account"` // 首字母必须大写，不然无法被JSON包解析
	Password string `json:"password"`
}

func init() {
	orm.RegisterModel(new(UserLogin)) // 注册模型，所有的模型都要在初始化函数中注册后才能使用
}
