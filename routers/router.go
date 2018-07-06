package routers

import (
	"WebMonitor-mas/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"WebMonitor-mas/controllers/serverManage"
)

func init() {
	beego.Include(&controllers.MainController{})
	beego.Include(&controllers.MonitorController{})
	beego.Include(&controllers.DNSController{})
	beego.Include(&serverManage.ServerController{})
	//解决跨域问题
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
}