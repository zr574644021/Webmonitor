webpackJsonp([8],{520:function(e,t,a){var r=a(200)(a(538),a(567),null,null);e.exports=r.exports},530:function(e,t,a){"use strict";a.d(t,"b",function(){return r}),a.d(t,"d",function(){return i}),a.d(t,"c",function(){return l}),a.d(t,"a",function(){return n});var r=function(e){var t=new Date(1e3*e);return t.getFullYear()+"-"+(t.getMonth()+1<10?"0"+(t.getMonth()+1):t.getMonth()+1)+"-"+t.getDate()+" "+t.getHours()+":"+t.getMinutes()+":"+t.getSeconds()},i=function(e,t){for(var a=0;a<t.length;a++)e.push({value:t[a].WebName})},l=function(e){return function(t){return 0===t.value.toLowerCase().indexOf(e.toLowerCase())}},n=function(e,t,a,r,i){for(var l=t.length,n=0;n<l;n++)t[n].Time>r&&t[n].Time<i&&e.push(t[n])}},538:function(e,t,a){"use strict";Object.defineProperty(t,"__esModule",{value:!0});var r=a(201),i=a.n(r),l=a(530);t.default={data:function(){return{tableData:[],bufferData:[],httpMonitor:!1,message:"开始检测",timer:0,first_time:0,last_time:0,webName:[],select:{WebName:"",MonitorTime:""},paging:{total:0,pageSize:25,page:1}}},methods:{handleCurrentChange:function(e){this.paging.page=e},getData:function(){var e=this;this.is_search?(this.bufferData=[],a.i(l.a)(this.bufferData,this.tableData,this.paging.total,this.first_time,this.last_time),this.paging.total=this.bufferData.length,this.tableData=this.bufferData):this.$axios.get("/dnsmonitor/error/all").then(function(t){if(7002===t.data.status)e.$message({message:"您未登录！",type:"error"}),e.$router.push({path:"/login"});else{var r=t.data.dnserror;e.tableData=r,e.paging.total=r.length;for(var i=0;i<r.length;i++)e.tableData[i].Id=i+1,e.tableData[i].Time=a.i(l.b)(e.tableData[i].Time);0===e.webName.length&&e.loadName(e.webName,e.tableData)}})},search:function(){var e=i()({},this.select);""===e.WebName&&""===e.MonitorTime||null===e.MonitorTime&&""===e.WebName?(this.is_search=!1,this.getData()):(this.is_search=!0,this.first_time=e.MonitorTime[0],this.last_time=e.MonitorTime[1],this.getData())},querySearch:function(e,t){var r=this.webName;t(e?r.filter(a.i(l.c)(e)):r)},loadName:function(){for(var e=0;e<this.tableData.length;e++){for(var t=0,a=0;a<this.webName.length;a++)if(this.webName[a].value===this.tableData[e].Dns.WebName){t=1;break}0===t&&this.webName.push({value:this.tableData[e].Dns.WebName})}}},created:function(){this.getData()}}},567:function(e,t){e.exports={render:function(){var e=this,t=e.$createElement,a=e._self._c||t;return a("div",{staticClass:"table"},[a("div",{staticClass:"crumbs"},[a("el-breadcrumb",{attrs:{separator:"/"}},[a("el-breadcrumb-item",[a("i",{staticClass:"el-icon-menu"}),e._v(" 域名解析")]),e._v(" "),a("el-breadcrumb-item",[e._v("错误历史")])],1)],1),e._v(" "),a("div",{staticClass:"handle-box"},[a("el-autocomplete",{staticClass:"inline-input",attrs:{placeholder:"网站名称","fetch-suggestions":e.querySearch,"trigger-on-focus":!1},model:{value:e.select.WebName,callback:function(t){e.$set(e.select,"WebName",t)},expression:"select.WebName"}}),e._v(" "),a("el-date-picker",{attrs:{"value-format":"yyyy-MM-dd hh:mm:ss",type:"datetimerange",align:"right","start-placeholder":"开始日期","end-placeholder":"结束日期","default-time":["00:00:00","24:00:00"]},model:{value:e.select.MonitorTime,callback:function(t){e.$set(e.select,"MonitorTime",t)},expression:"select.MonitorTime"}}),e._v(" "),a("el-button",{attrs:{type:"primary",icon:"search"},on:{click:e.search}},[e._v("搜索")])],1),e._v(" "),a("el-table",{staticStyle:{width:"100%"},attrs:{data:e.tableData.slice((e.paging.page-1)*e.paging.pageSize,e.paging.page*e.paging.pageSize),border:""}},[a("el-table-column",{attrs:{prop:"Id",label:"ID"}}),e._v(" "),a("el-table-column",{attrs:{prop:"Dns.WebName",label:"网站名称"}}),e._v(" "),a("el-table-column",{attrs:{prop:"Dns.Url",label:"网址"}}),e._v(" "),a("el-table-column",{attrs:{prop:"Carrier.Carrier",label:"运营商"}}),e._v(" "),a("el-table-column",{attrs:{prop:"Carrier.Address",label:"地址"}}),e._v(" "),a("el-table-column",{attrs:{prop:"Carrier.DnsIp",label:"DNSIP"}}),e._v(" "),a("el-table-column",{attrs:{prop:"ErrorMsg",label:"异常信息"}}),e._v(" "),a("el-table-column",{attrs:{prop:"ErrorIp",label:"异常IP"}}),e._v(" "),a("el-table-column",{attrs:{prop:"Time",label:"时间"}})],1),e._v(" "),a("div",{staticClass:"pagination"},[a("el-pagination",{attrs:{layout:"total, prev, pager, next",total:e.paging.total,"page-size":e.paging.pageSize},on:{"current-change":e.handleCurrentChange}})],1)],1)},staticRenderFns:[]}}});