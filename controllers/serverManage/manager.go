package serverManage

import (
	"WebMonitor-mas/models/class/serverManage"

	"WebMonitor-mas/controllers"
	"github.com/astaxie/beego"
	"encoding/json"
)

func (s *ServerController) managerGet() (manager serverManage.Manager) {
	//var managerMessage serverManage.Manager
	json.Unmarshal(s.Ctx.Input.RequestBody, &manager)
	return
}

//
// @router /servermanage/manager/getall [post]
func (s *ServerController) GetManager() {
	if controllers.LoginCheck(&s.MainController) {
		if num, managers, err := serverManage.Mg_Get_All(); err != nil {
			result := controllers.ResultJson(9001)
			result["error"] = err
			s.Data["json"] = result
			s.ServeJSON()
			return
		} else {
			result := controllers.ResultJson(4003)
			result["num"] = num
			result["managers"] = managers
			s.Data["json"] = result
			s.ServeJSON()
			return
		}
	}
}

//
// @router /servermanage/manager/add [post]
func (s *ServerController) AddManager() {
	if controllers.LoginCheck(&s.MainController) {
		var manager serverManage.Manager
		var department serverManage.Department
		managerAdd := s.managerGet()
		manager.Name = managerAdd.Name
		manager.PhoneNumber = managerAdd.PhoneNumber
		department = *managerAdd.Dp

		if err := department.Query();err != nil {
			beego.Debug("department.Query", manager.Name, manager.PhoneNumber, department.Name, err)
			s.Data["json"] = controllers.ResultJson(8000)
			s.ServeJSON()
			return
		}
		manager.Dp = &department
		if _, err := manager.Add(); err != nil {
			beego.Debug("manager.Add", manager.Name, manager.PhoneNumber, department.Name, err)
			s.Data["json"] = controllers.ResultJson(9000)
			s.ServeJSON()
			return
		}

		s.Data["json"] = controllers.ResultJson(4003)
		s.ServeJSON()
		return
	}
}

//
// @router /servermanage/manager/remove  [post]
func (s *ServerController) RemoveManager() {
	if controllers.LoginCheck(&s.MainController) {
		var manager[] serverManage.Manager
		json.Unmarshal(s.Ctx.Input.RequestBody, &manager)
		for i:=0; i<len(manager);i++  {
			if err := manager[i].Remove(); err != nil {
				result := controllers.ResultJson(9002)
				result["error"] = err.Error()
				s.Data["json"] = result
			} else {
				s.Data["json"] = controllers.ResultJson(4003)
			}
			s.ServeJSON()
		}
	}
}

//
// @router /servermanage/manager/query [post]
func (s *ServerController) QueryManager() {
	if controllers.LoginCheck(&s.MainController) {
		//var  managerQuery
		var manager serverManage.Manager
		manaQuery := s.managerGet()
		manager.Name = manaQuery.Name
		manager.Dp = manaQuery.Dp
		if err :=  manager.Dp.Query();err != nil {
			result := controllers.ResultJson(8000)
			result["error"] = err.Error()
			s.Data["json"] = result
		}
		if num,managers,err := manager.Query();err != nil {
			result := controllers.ResultJson(9001)
			result["error"] = err.Error()
			s.Data["json"] = result
		} else {
			if num == 0 {
				s.Data["json"] = controllers.ResultJson(9001)
			} else {
				result := controllers.ResultJson(4003)
				result["num"] = num
				result["managers"] = managers
				s.Data["json"] = result
			}
		}
		s.ServeJSON()
	}
}

// @router /servermanage/manager/set [post]
func (s *ServerController) SetManager() {
	if controllers.LoginCheck(&s.MainController) {
		var need_set serverManage.Manager
		manager := s.managerGet()

		need_set.Name = manager.Name
		need_set.PhoneNumber = manager.PhoneNumber
		var department serverManage.Department
		department = *manager.Dp
		if err := department.Query();err != nil {
			result := controllers.ResultJson(9003)
			result["error"] = err.Error()
			s.Data["json"] = result
		}

		need_set.Dp = &department
		if id,err := manager.Set(&need_set);err != nil {
			result := controllers.ResultJson(9003)
			result["error"] = err.Error()
			s.Data["json"] = result
		} else {
			result := controllers.ResultJson(4003)
			result["id"] = id
			s.Data["json"] = result
		}
		s.ServeJSON()
	}
}
