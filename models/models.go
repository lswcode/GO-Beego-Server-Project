package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int       `orm:"pk;auto"`                     // pk:主键，auto:自增
	Name     string    `orm:"unique"`                      // string类型字段默认映射为 varchar(255)
	Account  string    `orm:"index;unique"`                //  index: 为单个字段增加索引
	Password string    `orm:"column(user_password)"`       // 为字段设置自定义名称
	Created  time.Time `orm:"auto_now_add;type(datetime)"` // 使用datetime类型
}

func (u *User) TableName() string { // 给结构体添加方法，来自定义表名
	return "beego_user"
}

func init() {
	orm.RegisterModel(new(User)) // 注册模型，所有的模型都要在初始化函数中注册后才能使用
}
