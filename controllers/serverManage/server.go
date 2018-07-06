package serverManage

import (
	"WebMonitor-mas/models/class/serverManage"
	"WebMonitor-mas/controllers"
	"errors"
	"encoding/json"
	"fmt"
)

type ServerController struct {
	controllers.MainController
}

func (s *ServerController) serverGet() (server serverManage.Server) {
	json.Unmarshal(s.Ctx.Input.RequestBody, &server)
	return
}

func (s *ServerController) getManagerId(m *serverManage.Manager) (manager *serverManage.Manager,err error) {
	if num,managers,err := m.Query();err != nil {
		return nil,err
	} else {
		if num > 1 {
			s.Data["json"] = controllers.ResultJson(9004)
			s.ServeJSON()
		}	else if num == 0 {
			return nil,nil
		} else {
			return &managers[0],nil
		}
	}
	return nil,errors.New("remove server outnumber one")
}

// @router /servermanage/server/getall [post]
func (s *ServerController) GetServer() {
	if controllers.LoginCheck(&s.MainController) {
		num, servers, err := serverManage.ServerGet()
		if err != nil {
			result := controllers.ResultJson(10001)
			result["error"] = err.Error()
			s.Data["json"] = result
			s.ServeJSON()
		}
		result := controllers.ResultJson(4003)
		result["servers"] = servers
		result["num"] = num
		s.Data["json"] = result
		s.ServeJSON()
	}
}

//
// @router /servermanage/server/add [post]
func (s *ServerController) AddServer() {
	if controllers.LoginCheck(&s.MainController) {
		server := s.serverGet()
		var err error
		if server.Manager,err = s.getManagerId(server.Manager);err != nil {
			result := controllers.ResultJson(9004)
			result["error"] = err.Error()
			s.Data["json"] = result
			s.ServeJSON()
			return
		}
		if id, err := server.Add(); err != nil {
			result := controllers.ResultJson(10000)
			result["error"] = err.Error()
			s.Data["json"] = result
			s.ServeJSON()
			return
		} else {
			result := controllers.ResultJson(4003)
			result["id"] = id
			s.Data["json"] = result
			s.ServeJSON()
			return
		}
	}
}

//
// @router /servermanage/server/remove [post]
func (s *ServerController) RemoveServer() {
	if controllers.LoginCheck(&s.MainController) {
		//server := s.serverGet()
		var server[] serverManage.Server
		json.Unmarshal(s.Ctx.Input.RequestBody, &server)
		fmt.Println(server)
		var err error
		for i := 0; i < len(server); i++ {
			if server[i].Manager,err = s.getManagerId(server[i].Manager);err != nil {
				result := controllers.ResultJson(9004)
				result["error"] = err.Error()
				s.Data["json"] = result
				s.ServeJSON()
				return
			}
			if err := server[i].Remove(); err != nil {
				result := controllers.ResultJson(10002)
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
// @router /servermanage/server/query [post]
func (s *ServerController) QueryServer() {
	if controllers.LoginCheck(&s.MainController) {
		server := s.serverGet()
		server.Manager.Dp = &serverManage.Department{}
		if server.Manager.Name != ""{
			var err error
			if server.Manager,err = s.getManagerId(server.Manager);err != nil {
				result := controllers.ResultJson(9004)
				result["error"] = err.Error()
				s.Data["json"] = result
				s.ServeJSON()
				return
			}
		}
		if num,servers,err := server.Query();err != nil {
			result := controllers.ResultJson(10004)
			result["error"] = err.Error()
			s.Data["json"] = result
		} else {
			result := controllers.ResultJson(4003)
			result["num"] = num
			result["servers"] = servers
			s.Data["json"] = result
		}
		s.ServeJSON()
	}
}

//
// must input a true server_id
// @router /servermanage/server/set [post]
func (s *ServerController) SetServer() {
	if controllers.LoginCheck(&s.MainController) {
		server := s.serverGet()
		if err := server.Set(); err != nil {
			result := controllers.ResultJson(10003)
			result["error"] = err.Error()
			s.Data["json"] = result
			s.ServeJSON()
			return
		} else {
			s.Data["json"] = controllers.ResultJson(4003)
			s.ServeJSON()
			return
		}
	}
}