package serverManage

import (
	"github.com/astaxie/beego/orm"
	"errors"
	"fmt"
)

type Server struct {
	Id           int
	Area         string                        // 区域
	HardPosition string                        // 物理位置
	Name         string                        // 服务器名
	HardWare     string                        // 硬件环境
	Manager      *Manager `orm:"null;rel(fk)"` // 管理员
}

func (s *Server) Add() (id int64, err error) {
	if id, err = orm.NewOrm().Insert(s); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Server) Remove() (err error) {
	var num int64
	qs := orm.NewOrm().QueryTable("server")
	if s.Id != 0 {
		qs = qs.Filter("id",s.Id)
	}
	if s.Area != "" {
		qs = qs.Filter("id",s.Id)
	}
	if s.HardPosition != "" {
		qs = qs.Filter("id",s.Id)
	}
	if s.HardWare != "" {
		qs = qs.Filter("id",s.Id)
	}
	if s.Name != "" {
		qs = qs.Filter("id",s.Id)
	}
	if s.Manager != nil && s.Manager.Id != 0 {
		qs = qs.Filter("id",s.Id)
	}
	if num, err = qs.Count(); err != nil {
		return err
	}
	if num > 1 {
		return errors.New("remove server outnumber one")
	}
	if _, err = qs.Delete(); err != nil {
		return err
	}
	return nil
}

func (s *Server) Set() (err error) {
	s.Manager.Dp = &Department{}
	if num, manages, err := s.Manager.Query(); err != nil {
		return err
	} else if num > 1 {
		return errors.New("can't get only one data")
	} else {
		s.Manager = &manages[0]
	}
	var server Server

	if err := orm.NewOrm().QueryTable("server").
		Filter("id", s.Id).RelatedSel().One(&server); err != nil {
	}

	if s.Id != 0 {
		server.Id = s.Id
	}
	if s.Name != "" {
		server.Name = s.Name
	}
	if s.HardPosition != "" {
		server.HardPosition = s.HardPosition
	}
	if s.Area != "" {
		server.Area = s.Area
	}
	if s.HardWare != "" {
		server.HardWare = s.HardWare
	}
	if s.Manager != nil {
		var manager Manager
		server.Manager = &manager
		server.Manager.Id = s.Manager.Id
	}
	if _, err = orm.NewOrm().Update(&server); err != nil {
		return err
	}
	return nil
}

// if you input a member,
// will get one value in database top
func (s *Server) Query() (num int64, servers []Server, err error) {
	qs := orm.NewOrm().QueryTable("server")
	if s.Id != 0 {
		qs = qs.Filter("id", s.Id)
	}
	if s.Name != "" {
		qs = qs.Filter("name", s.Name)
	}
	if s.HardWare != "" {
		qs = qs.Filter("hard_ware", s.HardWare)
	}
	if s.Area != "" {
		qs = qs.Filter("area", s.Area)
	}
	if s.HardPosition != "" {
		qs = qs.Filter("hard_position", s.HardPosition)
	}
	if s.Manager.Id != 0 {
		// haven't test
		qs = qs.Filter("manager_id", s.Manager)
	}
	if s.Manager.Dp.Name != "" {
		qs = qs.Filter("", s.Manager.Dp.Name)
	}
	if num, err = qs.RelatedSel().All(&servers); err != nil {
		return 0, nil, err
	}
	fmt.Println(servers[0].Manager)
	fmt.Println(servers[0].Manager.Dp)
	return num, servers, nil
}

func ServerGet() (number int64, servers []Server, err error) {
	if number, err = orm.NewOrm().
		QueryTable("server").
		RelatedSel().All(&servers); err != nil {
		return 0, nil, err
	}
	return number, servers, nil
}
