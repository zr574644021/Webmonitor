package serverManage

import (
	"WebMonitor-mas/models/class/serverManage"
	"WebMonitor-mas/controllers"
	"WebMonitor-mas/controllers/base"

	"encoding/json"
)

type departmentSet struct {
	Id int
	Name string
	NewName string
}

func (s *ServerController) departmentGet() (department serverManage.Department) {
	json.Unmarshal(s.Ctx.Input.RequestBody, &department)
	return
}

//
// @router /servermanage/department/add [post]
func (s *ServerController) AddDepartment() {
	if controllers.LoginCheck(&s.MainController) {
		var department,departmentAdd serverManage.Department
		json.Unmarshal(s.Ctx.Input.RequestBody, &departmentAdd)
		department = departmentAdd
		if id, err := department.Add(); err != nil {
			result := controllers.ResultJson(8001)
			result["error"] = base.ErrorTo_zh_CN(err)
			s.Data["json"] = result
		} else {
			result := controllers.ResultJson(4003)
			result["id"] = id
			s.Data["json"] = result
		}
		s.ServeJSON()
	}
}

//
// @router /servermanage/department/remove [post]
func (s *ServerController) RemoveDepartment() {
	if controllers.LoginCheck(&s.MainController) {
		department := s.departmentGet()
		json.Unmarshal(s.Ctx.Input.RequestBody, &department)
		if err := department.Remove(); err != nil {
			result := base.ResultJson(9004)
			result["error"] = err.Error()
			s.Data["json"] = result
		} else {
			s.Data["json"] = base.ResultJson(4003)
		}
		s.ServeJSON()
	}
}

//
// @router /servermanage/department/set [post]
func (s *ServerController) SetDepartment() {
	if controllers.LoginCheck(&s.MainController) {
		//department := s.departmentGet()
		var department serverManage.Department
		var departmentAdd departmentSet
		json.Unmarshal(s.Ctx.Input.RequestBody, &departmentAdd)
		// use the name to set data
		department.Id = departmentAdd.Id
		department.Name = departmentAdd.Name
		name := departmentAdd.NewName
		if err := department.SetName(name);err != nil {
			result := controllers.ResultJson(8002)
			result["error"] = err.Error()
			s.Data["json"] = result
		} else {
			s.Data["json"] = controllers.ResultJson(4003)
		}
		s.ServeJSON()
	}
}

//
// @router /servermanage/department/getall [post]
func (s *ServerController) GetDepartment() {
	if controllers.LoginCheck(&s.MainController) {
		if num, departments, err := serverManage.Dp_Get_All(); err != nil {
			result := controllers.ResultJson(8000)
			result["error"] = err.Error()
			s.Data["json"] = result
			s.ServeJSON()
			return
		} else {
			result := controllers.ResultJson(4003)
			result["num"] = num
			result["departments"] = departments
			s.Data["json"] = result
			s.ServeJSON()
			return
		}

	}
}
