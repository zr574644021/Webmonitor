package main

import (
	"WebMonitor-mas/controllers"
	_ "WebMonitor-mas/initData"
	"WebMonitor-mas/models/class"
	_ "WebMonitor-mas/routers"
	_"WebMonitor-mas/controllers"
	"time"
	"github.com/astaxie/beego"
)

func main() {
	//go httpMonitor()
	//go dnsMonitor()
	beego.Run()
}

//URL解析IP
func dnsMonitor() {
	for {
		_, dnsAllList, _ := class.GetDnsAll()
		_, carrierList, _ := class.GetCarrierAll()
		class.DeleteVist()
		for _, dns := range dnsAllList {
			//wg := sync.WaitGroup{}
			//wg.Add(1)
			//go func(dns class.DnsList,wg sync.WaitGroup) {
			for _, car := range carrierList {
				var dnsError class.DnsVisitError
				var dnsVisit class.DnsVisitRecord
				go func(dnsList class.DnsList, carList class.CarrierList) {
					if dns.FinishSleep() {
						c1 := make(chan string)
						go class.DnsResolve(dnsList.Url, carList.DnsIp, c1)
						select {
						case resIP := <-c1:
							if resIP != dnsList.Ip { //IP是否正常
								dnsVisit.ErrorMsg, dnsError.ErrorMsg = "解析IP异常", "解析IP异常"
								dnsVisit.MonitorTime, dnsError.Time = time.Now().Unix(), time.Now().Unix()
								dnsVisit.ResolveIp, dnsError.ErrorIp = resIP, resIP
								dnsVisit.Dns, dnsError.Dns = &dnsList, &dnsList
								dnsVisit.Carrier, dnsError.Carrier = &carList, &carList
								dnsVisit.Status = 1
								dnsError.SaveError()
								dnsVisit.SaveVist()
								dnsList.UpdateTime() //更新检测时间
							} else {
								dnsList.UpdateTime() //更新检测时间
								dnsVisit.MonitorTime = time.Now().Unix()
								dnsVisit.Dns = &dnsList
								dnsVisit.Carrier = &carList
								dnsVisit.Status = 0
								dnsVisit.ResolveIp = resIP
								dnsVisit.SaveVist()
							}
						case <-time.After(time.Second * 3): //3S默认为超时
							//dnsVisit.ErrorMsg = /*, dnserror.ErrorMsg 	= "解析请求超时"		,*/ "解析请求超时"
							//dnsVisit.MonitorTime = /*, dnserror.Time 		= time.Now().Unix() ,*/ time.Now().Unix()
							//dnsVisit.Dns = /*, dnserror.Dns 		= &dnslist			,*/ &dnsList
							//dnsVisit.Carrier = /*, dnserror.Carrier 	= &carlist			,*/ &carList
							//dnsVisit.Status = 2
							//dnsList.UpdateTime() //更新检测时间
							////dnserror.SaveError()
							//dnsVisit.SaveVist()
						}
					}
				}(dns, car)
			}

		}
		time.Sleep(5 * time.Minute)
	}
}

func httpMonitor() {
	for {
		var i int64
		n, urllistall, _ := class.GetUrlAll() //获取所有需要监测的记录
		for i = 0; i < n; i++ {
			//并发请求所有记录
			go func(urllist class.UrlList) {
				statustime := make(chan class.StatusTime)
				//var visitrecord class.UrlVisitRecord
				var visiterror class.UrlVisitError
				if urllist.FinishSleep() { //判断是否休眠完成
					go controllers.HttpGetTime(urllist.Url, statustime)
					select {
					case get := <-statustime: //尝试取出时间
						if status := get.Status; status == "200 OK" {
							//visitrecord.Url = &urllist
							//visitrecord.Time = time.Now().Unix()
							//visitrecord.ResponseTime = get.Time
							//visitrecord.SaveVisit() //将数据存入记录表
							urllist.UpdateTime() //更新上次访问时间
							if get.Time > 2 {
								visiterror.Url = &urllist
								visiterror.ErrorStatus = "访问时间过长"
								visiterror.Time = time.Now().Unix()
								visiterror.SaveVisit()
							}
						} else {
							urllist.UpdateTime() //更新上次访问时间
							visiterror.Time = time.Now().Unix()
							visiterror.Url = &urllist
							visiterror.ErrorStatus = get.Status
							visiterror.SaveVisit()
						}
					case <-time.After(5 * time.Second): //超时的情况
						visiterror.ErrorStatus = "OutTime"
						visiterror.Time = time.Now().Unix()
						visiterror.Url = &urllist
						visiterror.SaveVisit() //将数据存入异常表
						urllist.UpdateTime()   //更新上次访问时间
					}
				}

				//visiterror = nil
				//urllist = nil
			}(urllistall[i])
		}
		time.Sleep(8 * time.Second)
	}
}