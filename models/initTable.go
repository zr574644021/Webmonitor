package models

import (
	"github.com/astaxie/beego/orm"
	"WebMonitor-mas/models/class"
	"WebMonitor-mas/models/class/serverManage"
)

//初始化
func init() {
	orm.RegisterModel(new(class.User),new(class.LoginRecord),new(class.UrlList),new(class.UrlVisitRecord),new(class.UrlVisitError))
	orm.RegisterModel(new(class.CarrierList),new(class.DnsList),new(class.DnsVisitError),new(class.DnsVisitRecord))
	orm.RegisterModel(new(serverManage.Server),new(serverManage.Manager),new(serverManage.Department))
}

