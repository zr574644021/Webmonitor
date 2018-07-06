webpackJsonp([9],{528:function(e,a,t){var r=t(200)(t(546),t(568),null,null);e.exports=r.exports},546:function(e,a,t){"use strict";Object.defineProperty(a,"__esModule",{value:!0});var r=t(201),o=t.n(r);a.default={data:function(){return{tableData:[],multipleSelection:[],exportLoading:!1,select:{Area:"",Manager:{Id:0,Name:"",Dp:{Id:0,Name:""}}},options:[],deletes:[],paging:{total:0,pageSize:15,page:1},editFormVisible:!1,editLoading:!1,editFormRules:{Name:[{required:!0,message:"请输入服务器名称",trigger:"blur"}],Area:[{required:!0,message:"请输入区域位置",trigger:"blur"}],HardPosition:[{required:!0,message:"请输入物理位置",trigger:"blur"}],Manager:{Name:[{required:!0,message:"请输入管理员姓名",trigger:"blur"}],PhoneNumber:[{required:!0,message:"请输入管理员电话",trigger:"blur"}]}},editForm:{Id:0,Name:"",Area:"",HardPosition:"",HardWare:"",Manager:{Id:0,Name:"",PhoneNumber:"",Dp:{Id:0,Name:""}}},addFormVisible:!1,addLoading:!1,addFormRules:{Name:[{required:!0,message:"请输入服务器名称",trigger:"blur"}],HardPosition:[{required:!0,message:"请输入物理位置",trigger:"blur"}],Area:[{required:!0,message:"请输入区域位置",trigger:"blur"}],ManagerName:[{required:!0,message:"请输入管理员姓名",trigger:"blur"}],ManagerPhoneNumber:[{required:!0,message:"请输入管理员电话",trigger:"blur"}],Options:[{required:!0,message:"请选择责任部门",trigger:"blur"}]},addForm:{Id:0,Name:"",Area:"",HardPosition:"",HardWare:"",Manager:{Id:0,Name:"",PhoneNumber:"",Dp:{Id:0,Name:""}}}}},created:function(){this.getData()},methods:{handleCurrentChange:function(e){this.paging.page=e},getData:function(){var e=this;this.$axios.post("/servermanage/server/getall").then(function(a){switch(a.data.status){case 7002:e.$message({message:"您未登录！",type:"error"}),e.$router.push({path:"/login"});break;case 4003:var t=a.data.servers;console.log(t),e.tableData=t,e.paging.total=t.length,0===e.options.length&&e.$axios.post("/servermanage/department/getall").then(function(a){switch(a.data.status){case 7002:e.$message({message:"您未登录!",type:"error"}),e.$router.push({path:"/login"});break;case 4003:for(var t=a.data.departments,r=0;r<t.length;r++)e.options.push({label:t[r].Id,value:t[r].Name});break;case 8e3:e.$message({message:"信息查找失败!",type:"error"})}});break;case 10001:e.$message({message:"查找信息异常！",type:"error"})}})},handleEdit:function(e,a){this.editFormVisible=!0,this.editForm=o()({},a)},handleAdd:function(){this.addFormVisible=!0,this.addForm={Name:"",Area:"",HardPosition:"",HardWare:"",Manager:{Id:0,Name:"",PhoneNumber:"",Dp:{Id:0,Name:""}}}},search:function(){""===this.select.Area&&""===this.select.Manager.Name&&""===this.select.Manager.Dp.Name?(this.is_search=!1,this.querySearch()):(this.is_search=!0,this.querySearch())},querySearch:function(){var e=this;if(this.is_search){var a=o()({},this.select);console.log(a),this.$axios.post("/servermanage/server/query",a).then(function(a){switch(console.log(a.data),a.data.status){case 7002:e.$message({message:"您未登录！",type:"error"}),e.$router.push({path:"/login"});break;case 4003:e.tableData=a.data.servers,e.paging.total=a.data.servers.length;break;case 9004:e.$message({message:"管理员不存在!",type:"error"});break;case 10004:e.$message({message:"查询失败!",type:"error"});break;default:e.tableData=[]}})}else this.getData()},handleDelete:function(e,a){var t=this;this.editForm=o()({},a),this.deletes[0]=this.editForm,this.$axios.post("/servermanage/server/remove",this.deletes).then(function(e){var a=e.data;switch(a.status){case 4001:t.$message.error("删除时"+a.err+"失败");break;case 5004:t.$message.success("删除成功!");break;case 5005:t.$message.error("删除失败!");break;case 7002:t.$message({message:"您未登录！",type:"error"}),t.$router.push({path:"/login"})}t.getData()})},editSubmit:function(){var e=this;this.$refs.editForm.validate(function(a){a&&e.$confirm("确认提交吗？","提示",{}).then(function(){e.editLoading=!0;for(var a=o()({},e.editForm),t=0;t<e.options.length;t++)if(a.Manager.Dp.Name===e.options[t].value){a.Manager.Dp.Id=e.options[t].label;break}e.$axios.post("/servermanage/server/set",a).then(function(a){e.editLoading=!1;var t=a.data;switch(console.log(t.status),t.status){case 4003:e.$message({message:"编辑成功",type:"success"});break;case 9004:e.$message({message:"管理员不存在",type:"error"});break;case 10004:e.$message({message:"编辑失败",type:"error"});break;case 7002:e.$message({message:"您未登录！",type:"error"}),e.$router.push({path:"/login"})}e.$refs.editForm.resetFields(),e.editFormVisible=!1,e.getData()})})})},addSubmit:function(){var e=this;this.$refs.addForm.validate(function(a){a&&e.$confirm("确认提交吗？","提示",{}).then(function(){e.addLoading=!0;for(var a=o()({},e.addForm),t=0;t<e.options.length;t++)if(a.Manager.Dp.Name===e.options[t].value){a.Manager.Dp.Id=e.options[t].label;break}e.$axios.post("/servermanage/server/add",a).then(function(a){switch(e.addLoading=!1,a.data.status){case 4003:e.$message({message:"添加成功",type:"success"});break;case 4001:e.$message({message:"管理员不存在",type:"error"});break;case 1e4:e.$message({message:"添加失败",type:"error"});break;case 7002:e.$message({message:"您未登录！",type:"error"}),e.$router.push({path:"/login"})}e.$refs.addForm.resetFields(),e.addFormVisible=!1,e.getData()})})})},delAll:function(){var e=this;console.log(this.multipleSelection),this.$axios.post("/servermanage/server/remove",this.multipleSelection).then(function(a){switch(a.data.status){case 4001:self.$message.error("删除失败!");break;case 5004:self.$message.success("删除成功");break;case 5005:self.$message.success("删除失败");break;case 7002:e.$message({message:"您未登录！",type:"error"}),e.$router.push({path:"/login"})}e.getData()}),this.multipleSelection=[]},handleSelectionChange:function(e){this.multipleSelection=e},exportAllTable:function(){var e=this;this.exportLoading=!0,t.e(15).then(function(){e.exportLoading=!1;var a=t(532),r=a.export_json_to_excel,o=["ID","服务器名","区域","物理位置","硬件环境","责任部门","管理员","联系电话"],s=["Id","Name","Area","HardPosition","HardWare","Dp_Name","Manager_Name","Manager_PhoneNumber"],l=e.transfer(e.tableData);r(o,e.formatJson(s,l),"机房服务器清单")}.bind(null,t)).catch(t.oe)},exportTable:function(){var e=this;this.$confirm("是否导出所选的"+this.multipleSelection.length+"条数据?","提示",{}).then(function(){t.e(15).then(function(){var a=t(532),r=a.export_json_to_excel,o=["ID","服务器名","区域","物理位置","硬件环境","责任部门","管理员","联系电话"],s=["Id","Name","Area","HardPosition","HardWare","Dp_Name","Manager_Name","Manager_PhoneNumber"],l=e.transfer(e.multipleSelection);r(o,e.formatJson(s,l),"机房服务器清单")}.bind(null,t)).catch(t.oe)})},transfer:function(e){for(var a=[],t=0;t<e.length;t++)a.push({Id:t,Name:e[t].Name,Area:e[t].Area,HardPosition:e[t].HardPosition,HardWare:e[t].HardWare,Dp_Name:e[t].Manager.Dp.Name,Manager_Name:e[t].Manager.Name,Manager_PhoneNumber:e[t].Manager.PhoneNumber});return a},formatJson:function(e,a){return console.log(a),a.map(function(a){return e.map(function(e){return a[e]})})}}}},568:function(e,a){e.exports={render:function(){var e=this,a=e.$createElement,t=e._self._c||a;return t("div",{staticClass:"table"},[t("div",{staticClass:"crumbs"},[t("el-breadcrumb",{attrs:{separator:"/"}},[t("el-breadcrumb-item",[t("i",{staticClass:"el-icon-menu"}),e._v(" 服务器信息管理")]),e._v(" "),t("el-breadcrumb-item",[e._v("服务器信息表")])],1)],1),e._v(" "),t("div",{staticClass:"handle-box"},[t("el-button",{staticClass:"handle-del mr10",attrs:{type:"primary",icon:"delete"},on:{click:e.delAll}},[e._v("批量删除")]),e._v(" "),t("el-input",{staticClass:"handle-input mr10",attrs:{placeholder:"区域"},model:{value:e.select.Area,callback:function(a){e.$set(e.select,"Area",a)},expression:"select.Area"}}),e._v(" "),t("el-select",{attrs:{clearable:"",placeholder:"请选择责任部门"},model:{value:e.select.Manager.Dp.Name,callback:function(a){e.$set(e.select.Manager.Dp,"Name",a)},expression:"select.Manager.Dp.Name"}},e._l(e.options,function(e){return t("el-option",{key:e.value,attrs:{label:e.value,value:e.value}})})),e._v(" "),t("el-input",{staticClass:"handle-input mr10",attrs:{placeholder:"管理员"},model:{value:e.select.Manager.Name,callback:function(a){e.$set(e.select.Manager,"Name",a)},expression:"select.Manager.Name"}}),e._v(" "),t("el-button",{attrs:{type:"primary",icon:"search"},on:{click:e.search}},[e._v("搜索")]),e._v(" "),t("el-button",{attrs:{type:"primary"},on:{click:e.handleAdd}},[e._v("新增")]),e._v(" "),t("el-button",{attrs:{type:"primary",loading:e.exportLoading},on:{click:e.exportAllTable}},[e._v("全部导出")])],1),e._v(" "),t("el-table",{staticStyle:{width:"100%"},attrs:{data:e.tableData.slice((e.paging.page-1)*e.paging.pageSize,e.paging.page*e.paging.pageSize),border:""},on:{"selection-change":e.handleSelectionChange}},[t("el-table-column",{attrs:{type:"selection",width:"55"}}),e._v(" "),t("el-table-column",{attrs:{type:"index",label:"ID",width:"60"}}),e._v(" "),t("el-table-column",{attrs:{prop:"Name",label:"服务器名"}}),e._v(" "),t("el-table-column",{attrs:{prop:"Area",label:"区域"}}),e._v(" "),t("el-table-column",{attrs:{prop:"HardPosition",label:"物理位置"}}),e._v(" "),t("el-table-column",{attrs:{prop:"HardWare",label:"硬件环境"}}),e._v(" "),t("el-table-column",{attrs:{prop:"Manager.Dp.Name",label:"数据责任部门"}}),e._v(" "),t("el-table-column",{attrs:{prop:"Manager.Name",label:"服务器管理员"}}),e._v(" "),t("el-table-column",{attrs:{prop:"Manager.PhoneNumber",label:"联系电话"}}),e._v(" "),t("el-table-column",{attrs:{label:"操作",width:"180"},scopedSlots:e._u([{key:"default",fn:function(a){return[t("el-button",{attrs:{size:"small"},on:{click:function(t){e.handleEdit(a.$index,a.row)}}},[e._v("编辑")]),e._v(" "),t("el-button",{attrs:{size:"small",type:"danger"},on:{click:function(t){e.handleDelete(a.$index,a.row)}}},[e._v("删除")])]}}])})],1),e._v(" "),t("br"),e._v(" "),t("el-button",{staticStyle:{"margin-top":"4px",display:"inline-block"},attrs:{type:"primary",size:"small"},on:{click:function(a){e.exportTable()}}},[e._v("导出Excel")]),e._v(" "),t("div",{staticClass:"pagination"},[t("el-pagination",{attrs:{layout:"total, prev, pager, next",total:e.paging.total,"page-size":e.paging.pageSize},on:{"current-change":e.handleCurrentChange}})],1),e._v(" "),t("el-dialog",{attrs:{title:"编辑",visible:e.editFormVisible},on:{"update:visible":function(a){e.editFormVisible=a}}},[t("el-form",{ref:"editForm",attrs:{model:e.editForm,"label-width":"80px",rules:e.editFormRules}},[t("el-form-item",{attrs:{label:"服务器名",prop:"Name"}},[t("el-input",{attrs:{"auto-complete":"off"},model:{value:e.editForm.Name,callback:function(a){e.$set(e.editForm,"Name",a)},expression:"editForm.Name"}})],1),e._v(" "),t("el-form-item",{attrs:{label:"区域",prop:"Area"}},[t("el-input",{attrs:{"auto-complete":"off"},model:{value:e.editForm.Area,callback:function(a){e.$set(e.editForm,"Area",a)},expression:"editForm.Area"}})],1),e._v(" "),t("el-form-item",{attrs:{label:"物理位置",prop:"HardPosition"}},[t("el-input",{attrs:{"auto-complete":"off"},model:{value:e.editForm.HardPosition,callback:function(a){e.$set(e.editForm,"HardPosition",a)},expression:"editForm.HardPosition"}})],1),e._v(" "),t("el-form-item",{attrs:{label:"硬件环境"}},[t("el-input",{attrs:{"auto-complete":"off"},model:{value:e.editForm.HardWare,callback:function(a){e.$set(e.editForm,"HardWare",a)},expression:"editForm.HardWare"}})],1),e._v(" "),t("el-form-item",{attrs:{label:"责任部门"}},[t("el-select",{attrs:{filerable:"",placeholder:"请选择责任部门"},model:{value:e.editForm.Manager.Dp.Name,callback:function(a){e.$set(e.editForm.Manager.Dp,"Name",a)},expression:"editForm.Manager.Dp.Name"}},e._l(e.options,function(e){return t("el-option",{key:e.value,attrs:{label:e.value,value:e.value}})}))],1),e._v(" "),t("el-form-item",{attrs:{label:"管理员",prop:"ManagerName"}},[t("el-input",{attrs:{"auto-complete":"off"},model:{value:e.editForm.Manager.Name,callback:function(a){e.$set(e.editForm.Manager,"Name",a)},expression:"editForm.Manager.Name"}})],1),e._v(" "),t("el-form-item",{attrs:{label:"联系电话",prop:"ManagerPhoneNumber"}},[t("el-input",{attrs:{"auto-complete":"off"},model:{value:e.editForm.Manager.PhoneNumber,callback:function(a){e.$set(e.editForm.Manager,"PhoneNumber",a)},expression:"editForm.Manager.PhoneNumber"}})],1)],1),e._v(" "),t("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[t("el-button",{nativeOn:{click:function(a){e.editFormVisible=!1}}},[e._v("取消")]),e._v(" "),t("el-button",{attrs:{type:"primary",loading:e.editLoading},nativeOn:{click:function(a){return e.editSubmit(a)}}},[e._v("提交")])],1)],1),e._v(" "),t("el-dialog",{attrs:{title:"新增",visible:e.addFormVisible},on:{"update:visible":function(a){e.addFormVisible=a}}},[t("el-form",{ref:"addForm",attrs:{model:e.addForm,"label-width":"80px",rules:e.addFormRules}},[t("el-form-item",{attrs:{label:"服务器名",prop:"Name"}},[t("el-input",{attrs:{"auto-complete":"off"},model:{value:e.addForm.Name,callback:function(a){e.$set(e.addForm,"Name",a)},expression:"addForm.Name"}})],1),e._v(" "),t("el-form-item",{attrs:{label:"区域",prop:"Area"}},[t("el-input",{attrs:{"auto-complete":"off"},model:{value:e.addForm.Area,callback:function(a){e.$set(e.addForm,"Area",a)},expression:"addForm.Area"}})],1),e._v(" "),t("el-form-item",{attrs:{label:"物理位置",prop:"HardPosition"}},[t("el-input",{attrs:{"auto-complete":"off"},model:{value:e.addForm.HardPosition,callback:function(a){e.$set(e.addForm,"HardPosition",a)},expression:"addForm.HardPosition"}})],1),e._v(" "),t("el-form-item",{attrs:{label:"硬件环境"}},[t("el-input",{attrs:{"auto-complete":"off"},model:{value:e.addForm.HardWare,callback:function(a){e.$set(e.addForm,"HardWare",a)},expression:"addForm.HardWare"}})],1),e._v(" "),t("el-form-item",{attrs:{label:"责任部门"}},[t("el-select",{attrs:{filerable:"",placeholder:"请选择责任部门"},model:{value:e.addForm.Manager.Dp.Name,callback:function(a){e.$set(e.addForm.Manager.Dp,"Name",a)},expression:"addForm.Manager.Dp.Name"}},e._l(e.options,function(e){return t("el-option",{key:e.value,attrs:{label:e.value,value:e.value}})}))],1),e._v(" "),t("el-form-item",{attrs:{label:"管理员",prop:"Manager.Name"}},[t("el-input",{attrs:{"auto-complete":"off"},model:{value:e.addForm.Manager.Name,callback:function(a){e.$set(e.addForm.Manager,"Name",a)},expression:"addForm.Manager.Name"}})],1),e._v(" "),t("el-form-item",{attrs:{label:"联系电话",prop:"Manager.PhoneNumber"}},[t("el-input",{attrs:{"auto-complete":"off"},model:{value:e.addForm.Manager.PhoneNumber,callback:function(a){e.$set(e.addForm.Manager,"PhoneNumber",a)},expression:"addForm.Manager.PhoneNumber"}})],1)],1),e._v(" "),t("div",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[t("el-button",{nativeOn:{click:function(a){e.addFormVisible=!1}}},[e._v("取消")]),e._v(" "),t("el-button",{attrs:{type:"primary",loading:e.addLoading},nativeOn:{click:function(a){return e.addSubmit(a)}}},[e._v("提交")])],1)],1)],1)},staticRenderFns:[]}}});