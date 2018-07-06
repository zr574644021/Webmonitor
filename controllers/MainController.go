package controllers

import (
	"encoding/json"

	"WebMonitor-mas/models/class"

	"github.com/astaxie/beego"
)

type Session struct {
	Session int
}

type MainController struct {
	class.User
	beego.Controller
}

type Login struct {
	UserName string
	PassWord string
}

//@router / [get]
func (c *MainController) LoginGet() {
	c.TplName = "index.html"
}

//@router /login [*]
func (c *MainController) LoginPost() {
	var login Login
	json.Unmarshal(c.Ctx.Input.RequestBody, &login) //json格式解析
	c.UserName = login.UserName
	c.Password = login.PassWord
	if c.Login_matching() { //验证session中的用户名和密码是否正确
		c.SetSession("username", c.UserName)
		c.SetSession("password", class.Encrypt(c.Password))
		c.Data["json"] = ResultJson(7004) //返回登陆成功
		c.ServeJSON()
		c.Login_record("", c.Ctx.Request.RemoteAddr) //写入登陆记录
		return
	} else {
		c.Data["json"] = ResultJson(7000) //账号或密码错误返回7000
		c.ServeJSON()
		c.Login_record(c.Password, c.Ctx.Request.RemoteAddr) //将错误的密码记录
		return
	}

}

//@router /logout [*]
func (c *MainController) LogoutPost() {
	username := c.GetSession("username")
	password := c.GetSession("password")    //密文形式的密码
	if username == nil || password == nil { //尝试获取session,获取不到则返回7002(未登录)
		c.Data["json"] = ResultJson(7002)
		c.ServeJSON()
		return
	} else {
		c.DelSession("username")
		c.DelSession("password")
		c.Data["json"] = ResultJson(7005)
		c.ServeJSON()
		return
	}
}
