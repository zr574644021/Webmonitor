package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"],
		beego.ControllerComments{
			Method: "AddDepartment",
			Router: `/servermanage/department/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"],
		beego.ControllerComments{
			Method: "GetDepartment",
			Router: `/servermanage/department/getall`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"],
		beego.ControllerComments{
			Method: "RemoveDepartment",
			Router: `/servermanage/department/remove`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"],
		beego.ControllerComments{
			Method: "SetDepartment",
			Router: `/servermanage/department/set`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"],
		beego.ControllerComments{
			Method: "AddManager",
			Router: `/servermanage/manager/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"],
		beego.ControllerComments{
			Method: "GetManager",
			Router: `/servermanage/manager/getall`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"],
		beego.ControllerComments{
			Method: "QueryManager",
			Router: `/servermanage/manager/query`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"],
		beego.ControllerComments{
			Method: "RemoveManager",
			Router: `/servermanage/manager/remove`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"],
		beego.ControllerComments{
			Method: "SetManager",
			Router: `/servermanage/manager/set`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"],
		beego.ControllerComments{
			Method: "AddServer",
			Router: `/servermanage/server/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"],
		beego.ControllerComments{
			Method: "GetServer",
			Router: `/servermanage/server/getall`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"],
		beego.ControllerComments{
			Method: "QueryServer",
			Router: `/servermanage/server/query`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"],
		beego.ControllerComments{
			Method: "RemoveServer",
			Router: `/servermanage/server/remove`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers/serverManage:ServerController"],
		beego.ControllerComments{
			Method: "SetServer",
			Router: `/servermanage/server/set`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
