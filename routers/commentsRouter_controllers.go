package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:DNSController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:DNSController"],
		beego.ControllerComments{
			Method: "AddDns",
			Router: `/dnsmonitor/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:DNSController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:DNSController"],
		beego.ControllerComments{
			Method: "DnsDelete",
			Router: `/dnsmonitor/del`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:DNSController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:DNSController"],
		beego.ControllerComments{
			Method: "GetDnsError",
			Router: `/dnsmonitor/error/all`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:DNSController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:DNSController"],
		beego.ControllerComments{
			Method: "GetDnsAll",
			Router: `/dnsmonitor/getdnsall`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:DNSController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:DNSController"],
		beego.ControllerComments{
			Method: "DnsMonitorNow",
			Router: `/dnsmonitor/getresult`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:DNSController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:DNSController"],
		beego.ControllerComments{
			Method: "QuerryDns",
			Router: `/dnsmonitor/query`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:DNSController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:DNSController"],
		beego.ControllerComments{
			Method: "DnsUpdate",
			Router: `/dnsmonitor/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:MainController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:MainController"],
		beego.ControllerComments{
			Method: "LoginGet",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:MainController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:MainController"],
		beego.ControllerComments{
			Method: "LoginPost",
			Router: `/login`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:MainController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:MainController"],
		beego.ControllerComments{
			Method: "LogoutPost",
			Router: `/logout`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"],
		beego.ControllerComments{
			Method: "AutoMonitorAdd",
			Router: `/httpmonitor/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"],
		beego.ControllerComments{
			Method: "AutoMonitorDelete",
			Router: `/httpmonitor/del`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"],
		beego.ControllerComments{
			Method: "GetError",
			Router: `/httpmonitor/error/all`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"],
		beego.ControllerComments{
			Method: "HttpErrorDelete",
			Router: `/httpmonitor/error/del`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"],
		beego.ControllerComments{
			Method: "QueryErrorRecord",
			Router: `/httpmonitor/error/query`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"],
		beego.ControllerComments{
			Method: "GetUrlAll",
			Router: `/httpmonitor/geturlall`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"],
		beego.ControllerComments{
			Method: "HttpMonitorNow",
			Router: `/httpmonitor/now`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"],
		beego.ControllerComments{
			Method: "HttpMonitorQuery",
			Router: `/httpmonitor/query`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"],
		beego.ControllerComments{
			Method: "QuerryHttp",
			Router: `/httpmonitor/record/query`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"] = append(beego.GlobalControllerRouter["WebMonitor-mas/controllers:MonitorController"],
		beego.ControllerComments{
			Method: "AutoMonitorUpDate",
			Router: `/httpmonitor/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
