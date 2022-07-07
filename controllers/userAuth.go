package controllers

import (
	"encoding/json"

	"beego_server/models"
	"beego_server/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type UserAuthController struct {
	beego.Controller // 这个是必须写的，所有的请求响应处理，都继承自beego.Controller这个内置的结构体
}

func (UA *UserAuthController) Login() {
	// 获取Ajax数据
	// u.Ctx.Request 所有请求信息
	// u.Ctx.Request.Header 请求头
	// u.Ctx.Request.Host 请求的主机
	// u.Ctx.Request.Method 请求的方法
	// u.Ctx.Request.URL 请求URL

	// account := UA.GetString("account") // beego会自动解析请求中的query参数(使用?拼接在url中的数据)，直接使用GetString/GetInt获取即可
	// password := UA.GetString("password")
	logs.Info("完整的请求信息", UA.Ctx.Request)
	// ----------------------------------------------------------------

	// 使用GetString/GetInt无法获取body请求体中的数据，只能获取axios中的query格式参数

	var user models.UserLogin
	data := UA.Ctx.Input.RequestBody // 使用Ctx.Input.RequestBody获取Body参数
	logs.Debug("解析body获得data", data)
	err := json.Unmarshal(data, &user) // 将json数据封装并保存到user对象中
	if err != nil {
		logs.Info("json.Unmarshal is err:", err.Error())
	}
	logs.Debug("保存后的结构体", user.Account, user.Password)

	afterMD5 := utils.Md5(user.Password)
	logs.Notice("原密码", user.Password)
	logs.Notice("加密后的密码", afterMD5)

	res := map[string]interface{}{ // 返回键值对格式给前端，map或结构体都可以，推荐使用map，因为结构体的字段名一般都是大写，需要使用tag标签进行转换
		"code": 200, "msg": "登陆成功",
	}

	UA.Data["json"] = res // 表示以json格式发送数据给前端，会设置响应体头中的content-type为application/json
	UA.ServeJSON()        // 发送前必须将map类型转换成JSON格式
}

func (UA *UserAuthController) Register() {
	var user models.UserRegister
	data := UA.Ctx.Input.RequestBody
	logs.Debug("解析body获得data", data)
	err := json.Unmarshal(data, &user) // 将json数据封装并保存到user对象中
	if err != nil {
		logs.Info("json.Unmarshal is err:", err.Error())
	}

	logs.Info("原密码", user.Password)

	afterMD5 := utils.Md5(user.Password)
	logs.Info("加密后的密码", afterMD5)

	res := map[string]interface{}{ // 返回键值对格式给前端，map或结构体都可以，推荐使用map，因为结构体的字段名一般都是大写，需要使用tag标签进行转换
		"code": 200, "msg": "注册成功",
	}

	UA.Data["json"] = res // 表示以json格式发送数据给前端，会设置响应体头中的content-type为application/json
	UA.ServeJSON()        // 发送前必须将map类型转换成JSON格式
}
