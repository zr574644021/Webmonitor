package controllers

import (
	"encoding/json"

	"WebMonitor-mas/models/class"

	"github.com/astaxie/beego"
)

type dnsList struct {
	Id        int
	Url       string
	Ip        string
	WebName   string
	SleepTime float64
}

type querryDns struct {
	WebName string
	Url     string
}

/*type resultList struct {
	DnsList class.DnsList
	CarrierName		string
	CarrierAddress	string
	CarrierIP		string
	ResolveIP 		string
	ErrorMsg		string
	time			time.Time
	status 			int //状态，0：解析正常，1：解析超时，2：解析IP异常
}*/

type DNSController struct {
	class.DnsList
	MainController
}

//@router /dnsmonitor/getresult [get]
func (c *DNSController) DnsMonitorNow() {
	if LoginCheck(&c.MainController) {
		var dnsVisit []class.DnsVisitRecord
		_, dnsVisit, err := class.GetDnsVistAll()
		if err != nil {
			c.Data["json"] = ResultJson(4004)
			c.ServeJSON()
			return
		}
		result := ResultJson(4003)
		result["record"] = dnsVisit
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
}

/*func(c* DNSController) DnsResolveMonitor() {
	if LoginCheck(&c.MainController) {
		var record []resultList
		a, dnsAllList, err_1 := class.GetDnsAll()
		b, carrierList, err_2 := class.GetCarrierAll()
		if err_1 != nil || err_2 != nil {
			c.Data["json"] = ResultJson(4004)
			c.ServeJSON()
			return
		}
		record = make([]resultList,a*b)
		var wait sync.WaitGroup
		for i, dns := range dnsAllList {
			wait.Add(1)
				go func(i int,dnslist class.DnsList) {
					for _, car := range carrierList {
						c1 := make(chan string)
						var dnsrecord resultList
						go class.DnsResolve(dnslist.Url,car.DnsIp,c1)
						select {
						case resIP := <-c1:
							if resIP != dnslist.Ip {//IP是否正常
								dnsrecord.DnsList = dnslist
								dnsrecord.ErrorMsg = "解析IP异常"
								dnsrecord.time = time.Now()
								dnsrecord.CarrierName = car.Carrier
								dnsrecord.CarrierAddress = car.Address
								dnsrecord.CarrierIP = car.DnsIp
								dnsrecord.ResolveIP = resIP
								dnsrecord.status = 2
							}else {
								dnsrecord.DnsList = dnslist
								dnsrecord.time = time.Now()
								dnsrecord.CarrierName = car.Carrier
								dnsrecord.CarrierAddress = car.Address
								dnsrecord.CarrierIP = car.DnsIp
								dnsrecord.ResolveIP = resIP
								dnsrecord.status = 0
							}
						case <- time.After(time.Second * 3)://3S默认为超时
							dnsrecord.DnsList = dnslist
							dnsrecord.ErrorMsg = "解析请求超时"
							dnsrecord.time = time.Now()
							dnsrecord.CarrierName = car.Carrier
							dnsrecord.CarrierAddress = car.Address
							dnsrecord.CarrierIP = car.DnsIp
							dnsrecord.status = 1
						}
						record[i] = dnsrecord
						//close(c1)
					}
						wait.Done()
				}(i,dns)
		}
		wait.Wait()
		result := ResultJson(4003)
		result["record"] = record
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
}*/

//@router /dnsmonitor/getdnsall [get]
func (c *DNSController) GetDnsAll() {
	if LoginCheck(&c.MainController) { //检查是否登陆
		i, dnsall, err := class.GetDnsAll() //获得所有需要解析的记录
		if err != nil {                     //获取失败
			beego.Error("get dns all error is ", err)
			c.Data["json"] = ResultJson(4004) //数据库没有记录
			c.ServeJSON()
			return
		}
		var j int64
		record := make([]dnsList, i) //创建DNS数组，长度为i
		for j = 0; j < i; j++ {
			record[j].Id = dnsall[j].Id
			record[j].Url = dnsall[j].Url
			record[j].WebName = dnsall[j].WebName
			record[j].Ip = dnsall[j].Ip
			record[j].SleepTime = dnsall[j].SleepTime
		}
		result := ResultJson(4002) //请求成功
		result["record"] = record
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
}

//@router /dnsmonitor/error/all [get]
func (c *DNSController) GetDnsError() {
	if LoginCheck(&c.MainController) {
		_, dnserror, err := class.GetDnsErrorAll()
		if err != nil {
			c.Data["json"] = ResultJson(4002) //请求成功
			c.ServeJSON()
			return
		}
		result := ResultJson(4003) //获取数据成功
		result["dnserror"] = dnserror
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
}

//@router /dnsmonitor/query [post]
func (c *DNSController) QuerryDns() {
	if LoginCheck(&c.MainController) {
		var querryDns querryDns
		json.Unmarshal(c.Ctx.Input.RequestBody, &querryDns)
		//fmt.Println(querryDns)
		WebName := querryDns.WebName
		Url := querryDns.Url
		flag, list := class.DnsUrlQuerry(WebName, Url)
		if flag {
			result := ResultJson(4003) //获取数据成功
			result["record"] = list
			c.Data["json"] = result
			c.ServeJSON()
			return
		} else {
			result := ResultJson(4004) //数据库没有记录
			c.Data["json"] = result
			c.ServeJSON()
			return
		}
	}
}

//@router /dnsmonitor/add [post]
func (c *DNSController) AddDns() {
	if LoginCheck(&c.MainController) {
		var dnsList class.DnsList
		json.Unmarshal(c.Ctx.Input.RequestBody, &dnsList)
		c.WebName = dnsList.WebName
		c.Url = dnsList.Url
		c.Ip = dnsList.Ip
		c.SleepTime = dnsList.SleepTime
		if CheckInputString(c.Url, c.WebName) && c.SleepTime != 0 { //输入检查是否为空
			if c.DnsUrlAdd() {
				result := ResultJson(6000) //添加成功
				c.Data["json"] = result
				c.ServeJSON()
				return
			} else {
				result := ResultJson(6001) //添加失败
				c.Data["json"] = result
				c.ServeJSON()
				return
			}
		}
		result := ResultJson(4001)
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
}

//@router /dnsmonitor/update [post]
func (c *DNSController) DnsUpdate() {
	if LoginCheck(&c.MainController) {
		var dnsList class.DnsList
		json.Unmarshal(c.Ctx.Input.RequestBody, &dnsList)
		c.Id = dnsList.Id
		c.Url = dnsList.Url
		c.WebName = dnsList.WebName
		c.Ip = dnsList.Ip
		c.SleepTime = dnsList.SleepTime
		if CheckInputString(c.Url, c.WebName, c.Ip) && c.Id != 0 && c.SleepTime != 0 { //输入检查是否为空
			flag := c.DnsUrlUpdate()
			if flag {
				c.Data["json"] = ResultJson(6002) //修改成功
				c.ServeJSON()
				return
			}
			c.Data["json"] = ResultJson(6003) //修改失败
			c.ServeJSON()
			return
		}
		c.Data["json"] = ResultJson(4001)
		c.ServeJSON()
		return
	}
}

//@router /dnsmonitor/del [post]
func (c *DNSController) DnsDelete() {
	if LoginCheck(&c.MainController) {
		var urls []string
		json.Unmarshal(c.Ctx.Input.RequestBody, &urls)
		if len(urls) == 0 {
			c.Data["json"] = ResultJson(4001)
		} else {
			flag, errName := class.DnsUrlDelete(urls)
			if flag {
				c.Data["json"] = ResultJson(6004) //删除成功
			} else {
				result := ResultJson(6005)
				result["err"] = errName
				c.Data["json"] = result //删除失败
				c.ServeJSON()
				return
			}
		}
		//软删除网站
		c.ServeJSON()
		return
	}
}
