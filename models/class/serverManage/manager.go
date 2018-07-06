package serverManage

import (
	"github.com/astaxie/beego/orm"
	"errors"
)

type Manager struct {
	Id          int
	Name        string      `orm:"unique"`       // 管理员姓名
	PhoneNumber string                           // 电话号码
	Dp          *Department `orm:"null;rel(fk)"` //部门外键
	Delete      bool        `orm:"default(false)"`
}

func managerFilter(qs *orm.QuerySeter,manager *Manager) {
	if manager.Id != 0 {
		*qs = (*qs).Filter("id", manager.Id)
	}
	if manager.Name != "" {
		*qs = (*qs).Filter("name", manager.Name)
	}
	if manager.PhoneNumber != "" {
		*qs = (*qs).Filter("phone_number", manager.PhoneNumber)
	}
	if manager.Dp.Id != 0 {
		(*qs).Filter("dp_id", manager.Dp)
	}
}

// add a manager
func (m *Manager) Add() (id int64, err error) {
	if id, err = orm.NewOrm().Insert(m); err != nil {
		return 0, err
	}
	return id, nil
}

// remove the manager.
// if don't delete one data, return a error
func (m *Manager) Remove() (err error) {
	var num int64
	qs := orm.NewOrm().QueryTable("manager")
	managerFilter(&qs,m)
	if num, err = qs.Count(); err != nil {
		return err
	} else {
		if num > 1 {
			return errors.New("remove server outnumber one")
		}
		if num, err = qs.Delete(); err != nil {
			return err
		}
		return nil
	}
}

// parameter is you need set data.
// m is to select a (lot of) data.
func (m *Manager) Set(need_update_m *Manager) (id int64, err error) {
	qs := orm.NewOrm().QueryTable("manager")
	managerFilter(&qs,m)
	if err := qs.One(m); err != nil {
		return 0, err
	}

	if need_update_m.Name!= "" {
		m.Name = need_update_m.Name
	}
	if need_update_m.PhoneNumber != "" {
		m.PhoneNumber = need_update_m.PhoneNumber
	}
	if need_update_m.Dp != nil {
		m.Dp = need_update_m.Dp
	}

	if id, err = orm.NewOrm().Update(m); err != nil {
		return 0, err
	}
	return id, nil
}


// input data of manager to query it
func (m *Manager) Query() (num int64, managers []Manager, err error) {
	qs := orm.NewOrm().QueryTable("manager")
	if m.Id != 0 {
		qs = qs.Filter("id", m.Id)
	}
	if m.Name != "" {
		qs = qs.Filter("name", m.Name)
	}
	if m.PhoneNumber != "" {
		qs = qs.Filter("phone_number", m.PhoneNumber)
	}
	if m.Dp != nil && m.Dp.Id != 0 {
		qs = qs.Filter("dp_id", m.Dp)
	}
	if num, err = qs.RelatedSel().All(&managers); err != nil {
		return 0, nil, err
	}
	return num, managers, nil
}

func Mg_Get_All() (number int64, managers []Manager, err error) {
	if number, err = orm.NewOrm().
		QueryTable("manager").Filter("delete",0).
		RelatedSel().All(&managers); err != nil {
		return 0, nil, err
	}
	return number, managers, nil

}
