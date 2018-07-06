package class

import "github.com/astaxie/beego/orm"

type CarrierList struct {
	Id		int
	Carrier string	//运营商
	Address string	//地区
	DnsIp   string  //DNS服务器IP
}

//获取所有DNS表记录
func GetCarrierAll() (int64,[]CarrierList,error){
	var carrierlist []CarrierList
	i,err := orm.NewOrm().QueryTable("carrier_list").All(&carrierlist)
	if err != nil {
		return 0,nil,err
	}
	return i,carrierlist,nil
}


//预留DNS服务器的增删查改