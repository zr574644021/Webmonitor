package serverManage

import (
	"github.com/astaxie/beego/orm"

	"errors"
)

type Department struct {
	Id   int
	Name string `orm:"unique"` //部门名
}

func (d *Department) Add() (id int64, err error) {
	if id, err = orm.NewOrm().Insert(d); err != nil {
		return 0, err
	}
	return id, nil
}

func (d *Department) Remove() (err error) {
	qs := orm.NewOrm().QueryTable("department")
	if d.Id != 0 {
		qs = qs.Filter("id", d.Id)
	}
	if d.Name != "" {
		qs = qs.Filter("name", d.Name)
	}
	if num,err := qs.Count();err != nil {
		return err
	} else if num != 1 {
		return errors.New("parameter error")
	}
	if _,err := qs.Delete();err != nil  {
		return err
	}
	return nil
}

func (d *Department) SetName(name string) (err error) {
	o := orm.NewOrm()
	if d.Id != 0 {
		d.Name = name
		if _,err =o.Update(d,"name");err != nil {
			return err
		}
	} else if d.Name != "" {
		if err := o.QueryTable("department").Filter("name",name).One(d,"name");err != nil {
			return err
		}
		d.Name = name
		if _,err = o.Update(d);err != nil {
			return err
		}
	}
	return nil
}

// input a parameter,
// get all message
func (d *Department) Query() error {
	qs := orm.NewOrm().QueryTable("department")

	if d.Id != 0 {
		qs = qs.Filter("id", d.Id)
	}
	if d.Name != "" {
		qs = qs.Filter("name", d.Name)
	}
	if num ,err := qs.Count(); err != nil {
		return err
	} else if num != 1 {
		return errors.New("query data is not one")
	}
	if err := qs.One(d); err != nil {
		return err
	}
	return nil
}

func Dp_Get_All() (count int64, dps []Department, err error) {
	if count, err = orm.NewOrm().QueryTable("department").All(&dps); err != nil {
		return 0, nil, err
	}
	return count, dps, nil
}
