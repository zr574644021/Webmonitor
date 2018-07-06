package controllers

import (
	"encoding/json"
	"sync"
	"time"

	"WebMonitor-mas/models/class"

	"github.com/astaxie/beego"
)

type UrlRecord struct {
	Url          string
	Responsetime float64
	Time
}

type Time struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
}

type UrlNow struct {
	Id           int
	Url          string
	WebName      string
	SleepTime    float64
	ResponseTime float64
	Status       string
	ResCode      int
}

type urllist struct {
	Id        int
	Url       string
	WebName   string
	SleepTime float64
}

type MonitorController struct {
	class.UrlList
	MainController
}

//@router /httpmonitor/geturlall [get]
func (c *MonitorController) GetUrlAll() {
	if LoginCheck(&c.MainController) { // check whether the user is login
		i, urlall, err := class.GetUrlAll() //get monitor url all and count(urlall)
		if err != nil {                     //get monitor url all error
			beego.Error("get url all error is ", err)
			c.Data["json"] = ResultJson(4004)
			c.ServeJSON()
			return
		}
		var j int64
		record := make([]urllist, i) //create a array of urllist, list is i
		for j = 0; j < i; j++ {
			record[j].Id = urlall[j].Id
			record[j].Url = urlall[j].Url
			record[j].WebName = urlall[j].WebName
			record[j].SleepTime = urlall[j].SleepTime
		}
		result := ResultJson(4002)
		result["record"] = record
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
}

//@router /httpmonitor/query [post]
func (c *MonitorController) HttpMonitorQuery() {
	if LoginCheck(&c.MainController) {
		statustime := make(chan class.StatusTime)
		url := c.GetString("url")
		if CheckInputString(url) {
			//open a go func to test outtime
			go HttpGetTime(url, statustime)
			select {
			//try to get data in chan
			case get := <-statustime:
				result := ResultJson(4003)
				result["time"] = get
				c.Data["json"] = result
				c.ServeJSON()
				return
				//get data in chan outtime
			case <-time.After(5 * time.Second):
				c.Data["json"] = ResultJson(-1)
				c.ServeJSON()
				return
			}
			close(statustime)
		}
		c.Data["json"] = ResultJson(4001)
		c.ServeJSON()
		return
	}
}

//@router /httpmonitor/now [get]
func (c *MonitorController) HttpMonitorNow() {
	if LoginCheck(&c.MainController) {
		var record []UrlNow
		//get url all
		i, urllists, err := class.GetUrlAll()
		if err != nil {
			c.Data["json"] = ResultJson(4004)
			c.ServeJSON()
			return
		}
		//make a chan get all returning data
		record = make([]UrlNow, i)
		var wg sync.WaitGroup
		for i, urllist := range urllists {
			wg.Add(1) //add a go func
			go func(i int, urlList class.UrlList) {
				statustime := make(chan class.StatusTime)
				go HttpGetTime(urlList.Url, statustime)
				var node UrlNow
				select {
				case get := <-statustime:
					node.ResponseTime = get.Time
					node.Status = get.Status
					node.ResCode = 200 //响应成功
				case <-time.After(2 * time.Second):
					node.ResponseTime = 0
					node.Status = "OutTime"
					node.ResCode = 400 //响应失败
				}
				node.Id = urlList.Id
				node.SleepTime = urlList.SleepTime
				node.WebName = urlList.WebName
				node.Url = urlList.Url
				record[i] = node
				//close(statustime)
				wg.Done() // finish a go func
			}(i, urllist)
		}
		wg.Wait() // if have a go func not finish ,wait it
		result := ResultJson(4003)
		result["record"] = record
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
}

//@router /httpmonitor/error/all [get]
func (c *MonitorController) GetError() {
	if LoginCheck(&c.MainController) {
		_, urlerror, err := class.GetErrorAll()
		if err != nil {
			c.Data["json"] = ResultJson(4002)
			c.ServeJSON()
			return
		}
		result := ResultJson(4003)
		result["urlerror"] = *urlerror
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
}

//@router /httpmonitor/error/query [post]
func (c *MonitorController) QueryErrorRecord() {
	if LoginCheck(&c.MainController) {
		url := c.GetString("url")
		time1year, err1 := c.GetIntCheck("time1year")
		time1month, err2 := c.GetIntCheck("time1month")
		time1day, err3 := c.GetIntCheck("time1day")
		time1hour, err4 := c.GetInt("time1hour")
		time2year, err5 := c.GetIntCheck("time2year")
		time2month, err6 := c.GetIntCheck("time2month")
		time2day, err7 := c.GetIntCheck("time2day")
		time2hour, err8 := c.GetInt("time2hour")
		//check whether year,month,day was input
		if err1 != nil || err2 != nil || err3 != nil || err5 != nil || err6 != nil || err7 != nil {
			return
		}
		//if the hour is not input,they will be 0
		if err4 != nil || err8 != nil {
			time1hour = 0
			time2hour = 0
		}

		//create time.Time object ,it have operation of time
		time1 := time.Date(time1year, time.Month(time1month), time1day, time1hour, 0, 0, 0, time.Local)
		time2 := time.Date(time2year, time.Month(time2month), time2day, time2hour, 0, 0, 0, time.Local)
		var result map[string]interface{}
		if i, errorrecords := class.QueryUrlErrorRecord(url, time1.Unix(), time2.Unix()); errorrecords != nil {
			result = ResultJson(4003)
			result["errorrecords"] = *errorrecords
			result["count"] = i
		} else {
			result = ResultJson(4004)
		}
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
}

//@router /httpmonitor/add [post]
func (c *MonitorController) AutoMonitorAdd() {
	if LoginCheck(&c.MainController) {
		//var err error
		var urlList urllist
		json.Unmarshal(c.Ctx.Input.RequestBody, &urlList)
		c.Url = urlList.Url
		c.WebName = urlList.WebName
		c.SleepTime = urlList.SleepTime
		if c.SleepTime == 0 {
			//not found sleeptime
			c.Data["json"] = ResultJson(4001)
			c.ServeJSON()
			return
		}
		//check get string true
		if CheckInputString(c.Url, c.WebName) {
			//add url record to mysql
			if c.WebUrlAdd() {
				//success
				c.Data["json"] = ResultJson(5000)
				c.ServeJSON()
				return
			} else {
				//false
				c.Data["json"] = ResultJson(5001)
				c.ServeJSON()
				return
			}
		}
		c.Data["json"] = ResultJson(4001)
		c.ServeJSON()
		return
	}
}

//@router /httpmonitor/update [post]
func (c *MonitorController) AutoMonitorUpDate() {
	if LoginCheck(&c.MainController) { //确认是否登陆
		var urlList urllist
		json.Unmarshal(c.Ctx.Input.RequestBody, &urlList)
		c.Id = urlList.Id
		c.Url = urlList.Url
		c.WebName = urlList.WebName
		c.SleepTime = urlList.SleepTime
		//check whether getstring is ""
		if CheckInputString(c.Url, c.WebName) {
			if c.SleepTime == 0 {
				//未输入seleeptime即不修改休眠时间
				if !class.WebUrlUpdate(c.Id, c.Url, c.WebName, -1) {
					c.Data["json"] = ResultJson(5003)
					c.ServeJSON()
					return
				}
			}
			//update mysql success
			if !class.WebUrlUpdate(c.Id, c.Url, c.WebName, c.SleepTime) {
				c.Data["json"] = ResultJson(5003)
				c.ServeJSON()
				return
			}
			c.Data["json"] = ResultJson(5002)
			c.ServeJSON()
			return
		}
		c.Data["json"] = ResultJson(4001)
		c.ServeJSON()
		return
	}
}

//@router /httpmonitor/del [post]
func (c *MonitorController) AutoMonitorDelete() {
	if LoginCheck(&c.MainController) {
		//url := c.GetString("url")
		var urls []string
		json.Unmarshal(c.Ctx.Input.RequestBody, &urls)
		for i, _ := range urls {
			if urls[i] == "" {
				result := ResultJson(4001)
				result["err"] = urls[i]
				c.Data["json"] = result
				c.ServeJSON()
				return
			}
		}
		//软删除网站
		for i, _ := range urls {
			if !class.WebUrlDelete(urls[i]) {
				c.Data["json"] = ResultJson(5005)
				c.ServeJSON()
				return
			}
		}
		c.Data["json"] = ResultJson(5004)
		c.ServeJSON()
		return
	}

}

//@router /httpmonitor/record/query [post]
func (c *MonitorController) QuerryHttp() {
	if LoginCheck(&c.MainController) {
		var querryDns querryDns
		json.Unmarshal(c.Ctx.Input.RequestBody, &querryDns)
		WebName := querryDns.WebName
		Url := querryDns.Url
		flag, list := class.HttpUrlQuerry(WebName, Url)
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

/*func (c *MonitorController) Queryrecord() {
	if LoginCheck(&c.MainController) {
		url := c.GetString("url")
		time1year, err1 := c.GetIntCheck("time1year")
		time1month, err2 := c.GetIntCheck("time1month")
		time1day, err3 := c.GetIntCheck("time1day")
		time1hour, err4 := c.GetInt("time1hour")
		time2year, err5 := c.GetIntCheck("time2year")
		time2month, err6 := c.GetIntCheck("time2month")
		time2day, err7 := c.GetIntCheck("time2day")
		time2hour, err8 := c.GetInt("time2hour")
		//check whether year,month,day was input
		if err1 != nil || err2 != nil || err3 != nil || err5 != nil || err6 != nil || err7 != nil {
			return
		}
		//if the hour is not input,they will be 0
		if err4 != nil || err8 != nil {
			time1hour = 0
			time2hour = 0
		}
		//create time.Time object ,it have operation of time
		time1 := time.Date(time1year, time.Month(time1month), time1day, time1hour, 0, 0, 0, time.Local)
		time2 := time.Date(time2year, time.Month(time2month), time2day, time2hour, 0, 0, 0, time.Local)
		//get sleeptime of the url
		sleeptime, err := class.GetSleepTime(url)
		if err != nil {
			c.Data["json"] = ResultJson(4004)
			c.ServeJSON()
			return
		}
		//calculate interval between time1 to time2
		timeinterval := time2.Sub(time1).Minutes()
		//calculate interval 1/100 time of time all
		frequency := int(timeinterval) / 100
		//judge whether time interval greater than sleeptime * 100
		if frequency < int(sleeptime) {
			//get url record all of the url between time1 and time2
			urlrecord := class.GetUrlRecordAll(url, time1.Unix(), time2.Unix())
			//handel when urlrecord is nil
			if urlrecord == nil {
				c.Data["json"] = ResultJson(4004)
				c.ServeJSON()
				return
			}
			urlrecords := make([]UrlRecord, len(*urlrecord))
			for j, value := range *urlrecord {
				urlrecords[j].Responsetime = value.ResponseTime
				//structure year.month.day.:hour:minute
				urlrecords[j].Year = time.Unix(int64(value.Time), 0).Year()
				urlrecords[j].Month = int(time.Unix(int64(value.Time), 0).Month())
				urlrecords[j].Day = time.Unix(int64(value.Time), 0).Day()
				urlrecords[j].Hour = time.Unix(int64(value.Time), 0).Hour()
				urlrecords[j].Minute = time.Unix(int64(value.Time), 0).Minute()
			}
			result := ResultJson(4003)
			result["urlrecord"] = urlrecords
			c.Data["json"] = result
			c.ServeJSON()
			return
		} else {
			urlrecordschan := make(chan UrlRecord, 100)
			//go func() {} () control
			var wg sync.WaitGroup
			for i := 0; i < 100; i++ {
				wg.Add(1) //add a go func wait record
				go func(i int) {
					var urlrecord_ UrlRecord
					timeadd, err1 := time.ParseDuration(strconv.Itoa(i*frequency+int(sleeptime)) + "m")
					timesub, err2 := time.ParseDuration(strconv.Itoa(i*frequency-int(sleeptime)) + "m")

					if err1 != nil || err2 != nil {
						beego.Error("time parse duration error is ", err1, err2)
						wg.Done() //finish a go func
						return
					}
					//get a url record between timeadd and timesub
					urlrecord := class.GetIntervalOneRecord(
						url,
						time1.Add(timesub).Unix(),
						time1.Add(timeadd).Unix())
					if urlrecord != nil {
						urlrecord_.Responsetime = urlrecord.ResponseTime
						//take a string.it likes year.month.day:hour:minute
						urlrecord_.Year = time.Unix(int64(urlrecord.Time), 0).Year()
						urlrecord_.Month = int(time.Unix(int64(urlrecord.Time), 0).Month())
						urlrecord_.Day = time.Unix(int64(urlrecord.Time), 0).Day()
						urlrecord_.Hour = time.Unix(int64(urlrecord.Time), 0).Hour()
						urlrecord_.Minute = time.Unix(int64(urlrecord.Time), 0).Minute()
						urlrecordschan <- urlrecord_
					}
					wg.Done() // finish a go func() {} ()
				}(i)
			}
			wg.Wait() //if have a go func() {} () can't finish in this place
			n := len(urlrecordschan)
			urlrecords := make([]UrlRecord, n)
			for i := 0; i < n; i++ {
				urlrecords[i] = <-urlrecordschan
			}
			close(urlrecordschan)
			result := ResultJson(4003)
			result["urlrecord"] = urlrecords
			c.Data["json"] = result
			c.ServeJSON()
			return
		}
	}
}*/

//router /httpmonitor/record/del [post]
//func (c *MonitorController) HttpRecordDelete() {
//	if LoginCheck(&c.MainController) {
//		url := c.GetString("url")
//		if url == "" {
//			c.Data["json"] = ResultJson(4001)
//			c.ServeJSON()
//			return
//		}
//		num, err := class.UrlRecordDelete(url)
//		if err != nil {
//			result := ResultJson(4002)
//			c.Data["json"] = result
//			c.ServeJSON()
//			return
//		}
//		result := ResultJson(4003)
//		result["num"] = num
//		c.Data["json"] = result
//		c.ServeJSON()
//		return
//	}
//
//}

//@router /httpmonitor/error/del [post]
func (c *MonitorController) HttpErrorDelete() {
	if LoginCheck(&c.MainController) {
		var urldelete []string
		var urls []string
		json.Unmarshal(c.Ctx.Input.RequestBody, &urls)
		//urls := c.GetStrings("urls")
		result := make(map[string]interface{})
		for _, url := range urls {
			if url == "" {
				c.Data["json"] = ResultJson(4001)
				c.ServeJSON()
				return
			}
			num, err := class.UrlErrorDelete(url)
			if err != nil {
				result := ResultJson(4002)
				c.Data["json"] = result
				c.ServeJSON()
				return
			}
			urldelete = append(urldelete, url)
			result["num"] = result["num"].(int64) + num
			result["urldelete"] = urldelete
		}
		result = ResultJson(4003)
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
}
