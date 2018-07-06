package class

import (
	"WebMonitor-mas/util"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type DnsList struct {
	Id            int
	Url           string  `orm:"unique"` //网站链接
	Ip            string  //网站IP
	WebName       string  `orm:"unique"` //网站名称
	SleepTime     float64 //睡眠时间(分钟)
	LastVisitTime int64   `orm:"null"`       //上次监测时间
	Status        int     `orm:"default(0)"` //软删除 0:未删除 1:已删除
}

type DnsVisitRecord struct {
	Id          int
	ErrorMsg    string       //异常信息
	ResolveIp   string       //解析IP
	MonitorTime int64        `orm:"null"`       //监测时间
	Status      int          `orm:"default(0)"` //状态信息 0:解析正常 1:解析异常 2:解析超时
	Dns         *DnsList     `orm:"rel(fk)"`
	Carrier     *CarrierList `orm:"rel(fk)"`
}

type DnsVisitError struct {
	Id       int
	Time     int64        //异常时间
	ErrorIp  string       //异常IP
	ErrorMsg string       //异常信息
	Dns      *DnsList     `orm:"rel(fk)"`
	Carrier  *CarrierList `orm:"rel(fk)"`
}

//URL解析
func DnsResolve(url, IP string, res chan string) {
	remsg, n, _ := util.Send(IP+":53", url) //地址解析为IP
	if n < 4 {
		close(res)
		return
	}
	ip := strconv.Itoa(int(remsg[n-4])) + "." + strconv.Itoa(int(remsg[n-3])) + "." +
		strconv.Itoa(int(remsg[n-2])) + "." + strconv.Itoa(int(remsg[n-1]))
	res <- ip
	close(res)
}

//获取所有域名解析异常记录
func GetDnsErrorAll() (int64, []DnsVisitError, error) {
	var dnsError []DnsVisitError
	o := orm.NewOrm()
	i, err := o.QueryTable("dns_visit_error").OrderBy("-time").RelatedSel().All(&dnsError)
	if err != nil {
		return 0, nil, err
	}
	return i, dnsError, nil
}

//获取所有需要解析的域名记录
func GetDnsAll() (int64, []DnsList, error) {
	var dnslist []DnsList
	i, err := orm.NewOrm().QueryTable("dns_list").
		Filter("status", 0).All(&dnslist) //筛选删除标志为0的记录
	if err != nil {
		return 0, nil, err
	}
	return i, dnslist, nil
}

//获取最新监测记录
func GetDnsVistAll() (int64, []DnsVisitRecord, error) {
	var dnsVisit []DnsVisitRecord
	i, err := orm.NewOrm().QueryTable("dns_visit_record").OrderBy("dns_id", "carrier_id").RelatedSel().All(&dnsVisit)
	if err != nil {
		return 0, nil, err
	}
	return i, dnsVisit, nil
}

//添加域名
func (c *DnsList) DnsUrlAdd() bool {
	o := orm.NewOrm()
	if _, err := o.Insert(c); err != nil {
		var dnsList DnsList
		if err := o.QueryTable("dns_list").
			Filter("url", c.Url).
			Filter("status", 1).One(&dnsList); err != nil {
			beego.Error("dnsUrl add(in geting) error is ", err)
			return false
		}
		dnsList.SleepTime = c.SleepTime
		dnsList.WebName = c.WebName
		dnsList.Ip = c.Ip
		dnsList.Status = 0
		if _, err := o.Update(&dnsList); err != nil {
			beego.Error("dns url add error is", err)
		}
		beego.Error("dnsUrl add error is ", err)
		return true
	}
	return true
}

//查询域名,网站名称,网站链接
func DnsUrlQuerry(webName, Url string) (bool, []DnsList) {
	o := orm.NewOrm()
	var dnslist []DnsList
	if webName == "" && Url != "" { //通过链接查询记录
		if err := o.QueryTable("dns_list").Filter("status", 0).Filter("url", Url).
			One(&dnslist); err != nil {
			beego.Error("DnsUrlFind find DnsList by url error is ", err)
			return false, nil
		} else {
			return true, dnslist
		}
	} else if webName != "" && Url == "" { //通过网站名称查询记录
		if err := o.QueryTable("dns_list").Filter("status", 0).Filter("web_name", webName).
			One(&dnslist); err != nil {
			beego.Error("DnsUrlFind find DnsList by webname error is ", err)
			return false, nil
		} else {
			return true, dnslist
		}
	} else if webName != "" && Url != "" { //通过网站名称和链接查询记录
		if err := o.QueryTable("dns_list").Filter("status", 0).Filter("url", Url).
			Filter("web_name", webName).One(&dnslist); err != nil {
			beego.Error("DnsUrlFind find DnsList by url,webname error is ", err)
			return false, nil
		} else {
			return true, dnslist
		}
	} else {
		if _, list, err := GetDnsAll(); err != nil {
			beego.Error("DnsUrlFind find DnsList error is ", err)
			return false, nil
		} else {
			return true, list
		}
	}
}

//更新域名
func (c *DnsList) DnsUrlUpdate() bool {
	var dnsList DnsList
	o := orm.NewOrm()
	if err := o.QueryTable("dns_list").Filter("id", c.Id).One(&dnsList); err != nil {
		beego.Error("dnsurlupdate get dnslist error is ", err)
		return false
	} else {
		c.LastVisitTime = dnsList.LastVisitTime
		if _, err := o.Update(c); err != nil {
			beego.Error("dnsurlupdate update dnslist error is ", err)
			return false
		}
	}
	return true
}

//批量删除域名
func DnsUrlDelete(url []string) (bool, string) {
	o := orm.NewOrm()
	var flag bool
	var errName string
	var dnslist DnsList
	for _, u := range url {
		if err := o.QueryTable("dns_list").Filter("url", u).One(&dnslist); err == nil {
			if _, err := o.QueryTable("dns_list").Filter("url", u).Update(orm.Params{"status": 1}); err == nil {
				flag = true
			} else {
				beego.Error("DnsUrlDelete update DnsList error is ", err)
				errName = dnslist.WebName
				flag = false
				return flag, errName
			}
		}
	}
	return flag, errName
}

//更新最后次检测时间
func (c *DnsList) UpdateTime() {
	c.LastVisitTime = time.Now().Unix()
	if _, err := orm.NewOrm().Update(c); err != nil {
		beego.Error("save update time error is ", err)
		return
	}
}

//保存异常信息
func (c *DnsVisitError) SaveError() bool {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		beego.Error("save error is ", err)
		return false
	}
	return true
}

//判断是否已完成休眠
func (c *DnsList) FinishSleep() bool {
	now := time.Now()
	lasttime := time.Unix(c.LastVisitTime, 0)
	//将两个时间戳相减，判断是否大于等于休眠时间
	if now.Sub(lasttime).Minutes() >= (c.SleepTime - 1) {
		return true
	}
	return false
}

func (c *DnsVisitRecord) SaveVist() bool {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		beego.Error("save error is ", err)
		return false
	}
	return true
}

func DeleteVist() bool {
	_, dnslist, err := GetDnsAll()
	if err != nil {
		return false
	}
	for _, u := range dnslist {
		orm.NewOrm().QueryTable("dns_visit_record").Filter("dns_id", u.Id).Delete()
	}
	return true
}
