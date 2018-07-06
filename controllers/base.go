package controllers

import (
	"WebMonitor-mas/models/class"
	"fmt"
	"net/http"
	"time"
	"net"
)
var Client = &http.Client{}

func init() {
	Client = &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				c, err := net.DialTimeout(netw, addr, time.Second*3)
				if err != nil {
					fmt.Println("dail timeout", err)
					return nil, err
				}
				return c, nil

			},
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * 2,
		},
	}
}



func ResultJson(errcode int) (result map[string]interface{}) {
	result = make(map[string]interface{})
	result["status"] = errcode
	return result
}

func LoginCheck(c *MainController) bool {
	username := c.GetSession("username")
	password := c.GetSession("password")    //密文形式的密码
	if username == nil || password == nil { //尝试获取session,获取不到则返回7002(未登录)
		c.Data["json"] = ResultJson(7002)
		c.ServeJSON()
		return false
	} else {
		var user class.User
		user.UserName = username.(string) //GetSession返回的为interface{} 需要使用.(string)强制转换成string类型
		user.Password = password.(string)
		if user.Login_matching_crypte() { //验证session中的用户名和密码是否正确
			return true
		} else {
			c.Data["json"] = ResultJson(7002) //错误返回7002(未登录)
			c.ServeJSON()
		}
	}
	return false
}

//测试是否输入了参数
func CheckInputString(input ...string) bool {
	for _, input_ := range input {
		if input_ == "" {
			return false
		}
	}
	return true
}

func (c *MainController) GetIntCheck(key string) (int, error) {
	value, err := c.GetInt(key)
	if err != nil {
		c.Data["json"] = ResultJson(4001)
		c.ServeJSON()
		return 0, err
	}
	return value, nil
}

func HttpGetTime(url string, get chan class.StatusTime) {
	//将计算好的时间差和状态码放入通道
	var statustime class.StatusTime
	time1 := time.Now() //获取请求前的时间
	res, err := Client.Get("http://" + url)
	time2 := time.Now() //获取请求后的时间
	if err != nil {
		//if get request error.content of error input in status
		statustime.Status = fmt.Sprintf("%s", err)
		get <- statustime
		close(get)
		return
	}

	statustime.Time = time2.Sub(time1).Seconds()
	statustime.Status = res.Status
	get <- statustime
	close(get)
}
