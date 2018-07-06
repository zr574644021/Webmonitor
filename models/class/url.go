package class

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UrlList struct {
	Id            int
	Url           string  `orm:"unique;index"` //网站链接
	WebName       string  `orm:"unique"`       //网站名称
	SleepTime     float64 //睡眠时间(分钟)
	LastVisitTime int64   `orm:"null"`       //上次监测时间
	Status        int     `orm:"default(0)"` //软删除 0:未删除 1:已删除
}

type UrlVisitRecord struct {
	Id           int
	ResponseTime float64  //响应时间
	Time         int64    //监测时间
	Url          *UrlList `orm:"rel(fk)"` //外键->UrlList表
}

type UrlVisitError struct {
	Id          int
	Time        int64    //监测时间
	ErrorStatus string   //状态码
	Url         *UrlList `orm:"rel(fk)"`
}

type StatusTime struct {
	Status string
	Time   float64
}

// 获取所有需要监测的记录
func GetUrlAll() (int64, []UrlList, error) {
	var urllists []UrlList
	i, err := orm.NewOrm().QueryTable("url_list").
		Filter("status", 0).All(&urllists)
	//筛选只有0的，1为软删除
	if err != nil {
		return 0, nil, err
	}
	return i, urllists, nil
}

// 获取最近的错误记录,can select limit max number
func GetErrorAll() (int64, *[]UrlVisitError, error) {
	var urlerror []UrlVisitError
	o := orm.NewOrm()
	i, err := o.QueryTable("url_visit_error").OrderBy("-time").RelatedSel().All(&urlerror)
	if err != nil {
		return 0, nil, err
	}

	return i, &urlerror, nil
}

func GetUrlRecordAll(url string, begin, end int64) *[]UrlVisitRecord {
	var urlrecord []UrlVisitRecord
	var urllist UrlList
	o := orm.NewOrm()
	err1 := o.QueryTable("url_list").
		Filter("status", 0).
		Filter("url", url).
		One(&urllist)
	if err1 != nil {
		return nil
	}
	_, err2 := o.QueryTable("url_visit_record").
		Filter("url_id", &urllist).
		Filter("time__gte", begin).
		Filter("time__lte", end).
		All(&urlrecord)
	if err2 != nil {
		return nil
	}
	return &urlrecord
}

// 判断是否已完成休眠
func (c *UrlList) FinishSleep() bool {
	now := time.Now()
	// 将int64的数据转化为时间戳
	lasttime := time.Unix(c.LastVisitTime, 0)
	// 将两个时间戳相减，判断是否大于等于休眠时间
	if now.Sub(lasttime).Minutes() >= (c.SleepTime - 1) {
		return true
	}
	return false
}

// 将监测的记录存入url_list表中(好像没有使用)
func (c *UrlVisitRecord) SaveVisit() bool {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		beego.Error("save visit record is ", err)
		return false
	}
	return true
}

func (c *UrlList) UpdateTime() {
	c.LastVisitTime = time.Now().Unix()
	if _, err := orm.NewOrm().Update(c); err != nil {
		beego.Error("save update time error is ", err)
		return
	}
}

// 保存超时监测记录
func (c *UrlVisitError) SaveVisit() {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		beego.Error("save visit record is ", err)
		return
	}
	return
}

func (c *UrlList) WebUrlAdd() bool {
	o := orm.NewOrm()
	if _, err := o.Insert(c); err != nil {
		var urllist UrlList
		if err := o.QueryTable("url_list").
			Filter("url", c.Url).
			Filter("status", 1).One(&urllist); err != nil {
			beego.Error("web url add(in geting) error is ", err)
			return false
		}
		urllist.SleepTime = c.SleepTime
		urllist.WebName = c.WebName
		urllist.Status = 0
		if _, err := o.Update(&urllist); err != nil {
			beego.Error("web url add error is", err)
		}
		beego.Error("web url add error is ", err)
		return true
	}
	return true
}

//查询
func HttpUrlQuerry(webName, Url string) (bool, []UrlList) {
	o := orm.NewOrm()
	var urlList []UrlList
	if webName == "" && Url != "" { //通过链接查询记录
		if err := o.QueryTable("url_list").Filter("status", 0).Filter("url", Url).
			One(&urlList); err != nil {
			beego.Error("HttpUrlFind find UrlList by url error is ", err)
			return false, nil
		} else {
			return true, urlList
		}
	} else if webName != "" && Url == "" { //通过网站名称查询记录
		if err := o.QueryTable("url_list").Filter("status", 0).Filter("web_name", webName).
			One(&urlList); err != nil {
			beego.Error("HttpUrlFind find UrlList by webname error is ", err)
			return false, nil
		} else {
			return true, urlList
		}
	} else if webName != "" && Url != "" { //通过网站名称和链接查询记录
		if err := o.QueryTable("url_list").Filter("status", 0).Filter("url", Url).
			Filter("web_name", webName).One(&urlList); err != nil {
			beego.Error("HttpUrlFind find UrlList by url,webname error is ", err)
			return false, nil
		} else {
			return true, urlList
		}
	} else {
		if _, list, err := GetUrlAll(); err != nil {
			beego.Error("DnsUrlFind find DnsList error is ", err)
			return false, nil
		} else {
			return true, list
		}
	}
}

// 更新网站url、名称，或监测间隔时间
func WebUrlUpdate(id int, url, webname string, sleeptime float64) bool {
	var urllist UrlList
	o := orm.NewOrm()
	if err := o.QueryTable("url_list").
		Filter("id", id).One(&urllist); err != nil {
		beego.Error("weburlupdate get urllist error is ", err)
		return false
	} else {
		urllist.Url = url
		// 当休眠时间为-1时，不修改休眠时间
		if sleeptime != float64(-1) {
			urllist.SleepTime = sleeptime
		}
		// 当网站名称为空时，不修改网站名
		if webname != "" {
			urllist.WebName = webname
		}
		if _, err := o.Update(&urllist); err != nil {
			beego.Error("weburlupdate update urllist error is ", err)
			return false
		}
	}
	return true
}

func WebUrlDelete(url string) bool {
	var urllist UrlList
	o := orm.NewOrm()
	if err := o.QueryTable("url_list").
		Filter("url", url).Filter("status", 0).One(&urllist); err != nil {
		beego.Error("weburldelete get urllist error is ", err)
		return false
	} else {
		urllist.Status = 1
		// 软删除，Delete属性为删除的时候为1
		if _, err := o.Update(&urllist); err != nil {
			beego.Error("weburldelete update urllist error is ", err)
			return false
		}
	}
	return true
}

// 通过URL获取对应网站的请求的休眠时间(-1为查询错误)
func GetSleepTime(url string) (float64, error) {
	var urllist UrlList
	err := orm.NewOrm().QueryTable("url_list").
		Filter("status", 0).Filter("url", url).
		One(&urllist, "sleep_time")
	if err != nil {
		return -1, err
	}
	return urllist.SleepTime, nil
}

func GetIntervalOneRecord(url string, begin, end int64) *UrlVisitRecord {
	var urllistrecord UrlVisitRecord
	var urllist UrlList
	o := orm.NewOrm()
	err1 := o.QueryTable("url_list").
		Filter("url", url).One(&urllist)
	if err1 != nil {
		return nil
	}
	err2 := o.QueryTable("url_visit_record").
		Filter("url_id", &urllist).Filter("time__gte", begin).Filter("time__lte", end).One(&urllistrecord, "response_time", "time")
	if err2 != nil {
		return nil
	}
	return &urllistrecord
}

func QueryUrlErrorRecord(url string, begin, end int64) (int64, *[]UrlVisitError) {
	var urllisterror []UrlVisitError
	var urllist UrlList
	o := orm.NewOrm()
	err1 := o.QueryTable("url_list").
		Filter("url", url).One(&urllist)
	if err1 != nil {
		return 0, nil
	}
	num, err2 := o.QueryTable("url_visit_record").
		Filter("url_id", &urllist).
		Filter("time__gte", begin).
		Filter("time__lte", end).
		All(&urllisterror)
	if err2 != nil {
		return 0, nil
	}
	return num, &urllisterror
}

// delete url record
func UrlErrorDelete(url string) (int64, error) {
	var urllist UrlList
	var num int64
	err := orm.NewOrm().
		QueryTable("url_list").
		Filter("url", url).
		One(&urllist)
	if err != nil {
		return 0, err
	}
	var urlrecord UrlVisitRecord
	urlrecord.Url = &urllist
	num, err = orm.NewOrm().
		QueryTable("url_visit_error").
		Filter("url_id", &urllist).Delete()
	if err != nil {
		return 0, err
	}
	return num, nil
}
