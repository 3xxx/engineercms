// @APIVersion 1.0.0
// @Title EngineerCMS API
// @Description ECMS has every tool to get any job done, so codename for the new ECMS APIs.
// @Contact 504284@qq.com
package routers

import (
	"github.com/3xxx/engineercms/conf"
	"github.com/3xxx/engineercms/controllers"
	// "githsub.com/3xxx/engineercms/controllers/checkin"
	// beego "github.com/beego/beego/v2/adapter"
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	// "github.com/beego/beego/v2/adapter/context"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/filter/cors" //这个坑
	// "github.com/3xxx/engineercms/controllers"
	// "github.com/3xxx/engineercms/controllers/utils"
	// "github.com/3xxx/engineercms/models"
	// "strconv"
)

// ***********mindoc********
func rt(req *http.Request) (*http.Response, error) {
	log.Printf("request received. url=%s", req.URL)
	// req.Header.Set("Host", "httpbin.org") // <--- I set it here as well
	defer log.Printf("request complete. url=%s", req.URL)

	return http.DefaultTransport.RoundTrip(req)
}

// roundTripper makes func signature a http.RoundTripper
type roundTripper func(*http.Request) (*http.Response, error)

func (f roundTripper) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }

type CorsTransport struct {
	http.RoundTripper
}

func (t *CorsTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	// refer:
	// - https://stackoverflow.com/questions/31535569/golang-how-to-read-response-body-of-reverseproxy/31536962#31536962
	// - https://gist.github.com/simon-cj/b4da0b2bca793ec3b8a5abe04c8fca41
	resp, err = t.RoundTripper.RoundTrip(req)
	logs.Debug(resp)
	if err != nil {
		return nil, err
	}
	resp.Header.Del("Access-Control-Request-Method")
	resp.Header.Set("Access-Control-Allow-Origin", "*")
	return resp, nil
}

// ***********mindoc********

// var FilterFunc = func(ctx *context.Context) {
// 	// userName := ctx.Input.Session("userName")
// 	// if userName == nil {
// 	// 	ctx.Redirect()
// 	// }
// 	v := ctx.Input.CruSession.Get("uname") //用来获取存储在服务器端中的数据??。
// 	// beego.Info(v)                          //qin.xc
// 	var user models.User
// 	var role string
// 	var uid int64
// 	var err error
// 	if v != nil { //如果登录了
// 		uname := v.(string)
// 		user, err = models.GetUserByUsername(uname)
// 		if err != nil {
// 			beego.Error(err)
// 		} else {
// 			uid := user.Id
// 			role = user.Role
// 		}
// 	} else { //如果没登录
// 		role = "anonymous"
// 	}
// }

// var FilterAdmin = func(ctx *context.Context) {
// 	// userName := ctx.Input.Session("userName")
// 	// if userName == nil {
// 	// 	ctx.Redirect()
// 	// }
// 	v := ctx.Input.CruSession.Get("uname") //用来获取存储在服务器端中的数据??。
// 	// beego.Info(v)                          //qin.xc
// 	var user models.User
// 	var err error
// 	if v != nil { //如果登录了
// 		uname := v.(string)
// 		user, err = models.GetUserByUsername(uname)
// 		if err != nil {
// 			beego.Error(err)
// 		} else {
// 			role, err := models.GetRoleByRolename("admin")
// 			if err != nil {
// 				beego.Error(err)
// 			}
// 			userid := strconv.FormatInt(user.Id, 10)
// 			roleid := strconv.FormatInt(role.Id, 10)
// 			isadmin = e.HasRoleForUser(userid, "role_"+roleid)
// 		}
// 	} else { //如果没登录
// 		role = "anonymous"
// 	}
// }

var FilterAdmin = func(ctx *context.Context) {
	_, _, _, isadmin, _ := controllers.CheckprodRole(ctx)
	if !isadmin {
		ctx.Redirect(301, "/login")
	}
}

func init() {
	//运行跨域请求
	//在http请求的响应流头部加上如下信息
	//rw.Header().Set("Access-Control-Allow-Origin", "*")
	// beego.InsertFilter("?=/article/*", beego.BeforeRouter, FilterFunc)
	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Token", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	// $.ajax({
	// 	url: ‘http://192.168.1.221:8080/v1/login/doLogin’,
	// 	type: ‘Post’,
	// 	dataType: ‘json’,
	// 	headers: {
	// 		Token:“wesdf2346dfioerjl2423mFADFOEJFds23SDFF”, //所有请求都带token 登陆就会得到token
	// 	},
	// 	data: {UserName: “retertsssssssssssssss”,UserId: ‘ertert22222222222’},
	// 	success: function(res){
	// 		console.log(res)
	// 	},
	// 	error: function(e) {
	// 		console.log(“errr==”)
	// 	}
	// });

	// var FilterUser = func(ctx *context.Context) {
	// 	// v := ctx.Input.CruSession.Get("uname")
	// 	// v := ctx.Input.Session("uname")
	// 	authString := ctx.GetCookie("token")
	// 	if authString == "" {
	// 		authString = ctx.Input.Query("token")
	// 		beego.Info(authString)
	// 	}
	// 	userid, username, _ := utils.LubanCheckToken(authString)
	// 	beego.Info(userid)
	// 	beego.Info(username)

	// 	// uname = v.(string)//uid---v.(int)
	// 	// if v == nil {
	// 	site := ctx.Input.Site() + ":" + strconv.Itoa(ctx.Input.Port())
	// 	if userid == "" {
	// 		ctx.Redirect(302, "https://www.54lby.com/sso/login?redirect_url="+site+ctx.Request.URL.String())
	// 	}
	// 	// if err != nil {
	// 	// 	// ctx.Redirect(302, "http://localhost:8080/v1/sso/ssologin?service="+site+ctx.Request.URL.String())
	// 	// 	ctx.Redirect(302, "https://www.54lby.com/sso/login?redirect_url="+site+ctx.Request.URL.String())
	// 	// }
	// }

	//自动化文档
	ns :=
		web.NewNamespace("/v1",
			// beego.NSBefore(FilterFunc),
			web.NSNamespace("/admin",
				// beego.NSBefore(FilterAdmin),20210724把admin都调整过来
				web.NSInclude(
					&controllers.AdminController{},
					&controllers.FlowController{},
					// &controllers.UserController{},
					// &controllers.AttachController{},20200626调整
					// &controllers.LoginController{},
					// &controllers.CustomerCookieCheckerController{},
				),
			),
			web.NSNamespace("/wx",
				// beego.NSBefore(FilterUser), //20210501
				web.NSInclude(
					&controllers.ArticleController{},
					&controllers.FroalaController{},
					&controllers.UserController{},
					&controllers.RegistController{},
					&controllers.LoginController{},
					&controllers.ReplyController{},
					&controllers.SearchController{},
					&controllers.AttachController{},
					&controllers.MainController{},
					&controllers.StandardController{},
					&controllers.DiaryController{},
					&controllers.PayController{},
					&controllers.FinanceController{},
					&controllers.PhotoController{},
					&controllers.VideoController{},
					&controllers.BusinessController{},
					&controllers.LocationController{},
					&controllers.SupaMapusController{},
				),
			),
			web.NSNamespace("/share",
				web.NSInclude(
					&controllers.ShareController{},
				),
			),
			web.NSNamespace("/adminlog",
				web.NSInclude(
					&controllers.AdminLogController{},
				),
			),
			web.NSNamespace("/project",
				web.NSInclude(
					&controllers.ProjController{},
					&controllers.ProdController{},
				),
			),
			web.NSNamespace("/checkin",
				web.NSInclude(
					&controllers.CheckController{},
				),
			),
			web.NSNamespace("/bbs",
				web.NSInclude(
					&controllers.BbsController{},
				),
			),
			web.NSNamespace("/todo",
				web.NSInclude(
					&controllers.TodoController{},
				),
			),
			web.NSNamespace("/onlyoffice",
				web.NSInclude(
					&controllers.OnlyController{},
				),
			),
			web.NSNamespace("/fileinput",
				web.NSInclude(
					&controllers.FileinputController{},
				),
			),
			web.NSNamespace("/pdfcpu",
				// beego.NSBefore(FilterUser),
				web.NSInclude(
					&controllers.PdfCpuController{},
				),
			),
			web.NSNamespace("/flv",
				// beego.NSBefore(FilterUser),
				web.NSInclude(
					&controllers.FlvController{},
				),
			),
			web.NSNamespace("/cart",
				// beego.NSBefore(FilterUser),
				web.NSInclude(
					&controllers.CartController{},
				),
			),
			web.NSNamespace("/mathcad",
				// beego.NSBefore(FilterUser),
				web.NSInclude(
					&controllers.MathcadController{},
					&controllers.WsMathcadController{},
				),
			),
			web.NSNamespace("/ansys",
				// beego.NSBefore(FilterUser),
				web.NSInclude(
					&controllers.AnsysController{},
				),
			),
			web.NSNamespace("/excel",
				// beego.NSBefore(FilterUser),
				web.NSInclude(
					&controllers.ExcelController{},
					&controllers.WsExcelcalController{},
				),
			),
			web.NSNamespace("/chat",
				web.NSInclude(
					&controllers.ChatController{},
				),
			),
			web.NSNamespace("/elastic",
				web.NSInclude(
					&controllers.ElasticController{},
				),
			),
			web.NSNamespace("/simwe",
				web.NSInclude(
					&controllers.SimweController{},
				),
			),
			web.NSNamespace("/freecad",
				web.NSInclude(
					&controllers.FreeCADController{},
				),
			),
			// beego.NSNamespace("/suggest",
			// 	beego.NSInclude(
			// 		&controllers.SearchController{},
			// 	),
			// ),
		)
	web.AddNamespace(ns)

	web.Router("/debug/pprof", &controllers.ProfController{})
	web.Router("/debug/pprof/:app([\\w]+)", &controllers.ProfController{})

	web.Router("/test", &controllers.MainController{}, "*:Test")
	web.Router("/mapus", &controllers.MainController{}, "*:Mapus")
	web.Router("/autodesk", &controllers.MainController{}, "*:Autodesk")
	// web.Router("/usermanage", &controllers.MainController{}, "*:UserManage")
	// web.Router("/.well-known/pki-validation/AC9A20F9BD09F18D247337AABC67BC06.txt", &controllers.AdminController{}, "*:Testdown")
	web.Router("/.well-known/pki-validation/*", &controllers.AdminController{}, "*:Testdown")

	web.Router("/doctree", &controllers.OnlyController{}, "*:GetTree")
	//升级数据库
	web.Router("/updatedatabase", &controllers.MainController{}, "*:UpdateDatabase")
	//删除数据表和字段测试
	web.Router("/modifydatabase", &controllers.MainController{}, "*:ModifyDatabase")

	web.Router("/url-to-callback", &controllers.OnlyController{}, "*:UrltoCallback")
	//cms中预览office回调
	web.Router("/officeviewcallback", &controllers.ProdController{}, "*:OfficeViewCallback")

	// 访问onlyoffice页面需要登录
	// web.InsertFilter("/onlyoffice", web.BeforeRouter, FilterUser)

	// web.Router("/onlyoffice/post", &controllers.OnlyController{}, "post:PostOnlyoffice")
	web.Router("/onlyoffice", &controllers.OnlyController{}, "get:Get")
	//table获取所有数据给上面界面使用
	web.Router("/onlyoffice/getdata", &controllers.OnlyController{}, "*:GetData")
	//添加一个文档
	web.Router("/onlyoffice/addattachment", &controllers.OnlyController{}, "post:AddOnlyAttachment")
	//在onlyoffice中打开文档协作
	// 协作页面下载文档需要登录
	// web.InsertFilter("/onlyoffice/:id:string", web.BeforeRouter, FilterUser)
	web.Router("/onlyoffice/:id:string", &controllers.OnlyController{}, "get:OnlyOffice") //20230603
	//cms中预览office
	web.Router("/officeview/:id:string", &controllers.ProdController{}, "*:OfficeView")

	//删除
	web.Router("/onlyoffice/deletedoc", &controllers.OnlyController{}, "*:DeleteDoc")
	//修改
	web.Router("/onlyoffice/updatedoc", &controllers.OnlyController{}, "*:UpdateDoc")
	//onlyoffice页面下载doc
	web.Router("/attachment/onlyoffice/*", &controllers.OnlyController{}, "get:DownloadDoc")
	//文档管理页面下载doc
	web.Router("/onlyoffice/download/:id:string", &controllers.OnlyController{}, "get:Download")
	// web.Router("/onlyoffice/changes", &controllers.OnlyController{}, "post:ChangesUrl")
	//*****onlyoffice document权限
	//添加用户和角色权限
	web.Router("/onlyoffice/addpermission", &controllers.OnlyController{}, "post:Addpermission")
	//取得文档的用户和角色权限列表
	web.Router("/onlyoffice/getpermission", &controllers.OnlyController{}, "get:Getpermission")

	web.Router("/role/test", &controllers.RoleController{}, "*:Test")
	web.Router("/1/slide", &controllers.MainController{}, "*:Slide")
	// web.Router("/postdata", &controllers.MainController{}, "*:Postdata")
	//文档
	web.Router("/doc/ecms", &controllers.MainController{}, "get:Getecmsdoc")
	web.Router("/doc/meritms", &controllers.MainController{}, "get:Getmeritmsdoc")
	web.Router("/doc/hydrows", &controllers.MainController{}, "get:Gethydrowsdoc")

	//api接口
	web.Router("/api/ecms", &controllers.MainController{}, "get:Getecmsapi")
	web.Router("/api/meritms", &controllers.MainController{}, "get:Getmeritmsapi")
	//根据app.conf里的设置，显示首页
	web.Router("/", &controllers.MainController{}, "get:Get")
	web.Router("/cms", &controllers.IndexController{}, "get:Cms")
	//显示首页
	web.Router("/index", &controllers.IndexController{}, "*:GetIndex")
	//首页放到onlyoffice
	// web.Router("/", &controllers.OnlyController{}, "get:Get")

	//显示右侧页面框架
	web.Router("/index/user", &controllers.IndexController{}, "*:GetUser")
	//这里显示用户查看主人日程
	web.Router("/calendar", &controllers.IndexController{}, "*:Calendar")
	//会议室预定日程
	web.Router("/meetingroom", &controllers.IndexController{}, "*:MeetingroomCalendar")
	web.Router("/meetingroom/searchcalendar", &controllers.IndexController{}, "*:SearchCalendar")

	//车辆预定日程
	web.Router("/car", &controllers.IndexController{}, "*:GetCarCalendar")
	//订餐
	web.Router("/order", &controllers.IndexController{}, "*:GetOrderCalendar")
	//考勤
	web.Router("/attendance", &controllers.IndexController{}, "*:GetAttendanceCalendar")

	//首页搜索项目或成果
	web.Router("/index/searchproject", &controllers.SearchController{}, "*:SearchProject")
	// web.Router("/index/searchproduct", &controllers.SearchController{}, "*:SearchProduct")

	//后台
	web.Router("/admin", &controllers.AdminController{})
	// web.Router("/admincategory", &controllers.AdminController{}, "*:GetAdminCategory")
	//显示对应侧栏id的右侧界面
	web.Router("/admin/:id:string", &controllers.AdminController{}, "*:Admin")

	//批量添加首页轮播图片
	web.Router("/admin/base/addcarousel", &controllers.AdminController{}, "*:AddCarousel")
	//获取首页轮播图片填充表格
	web.Router("/admin/base/carousel", &controllers.AdminController{}, "*:Carousel")
	//删除首页轮播图片
	web.Router("/admin/base/deletecarousel", &controllers.AdminController{}, "*:DeleteCarousel")

	//根据数字id查询类别或目录分级表
	web.Router("/admin/category/?:id:string", &controllers.AdminController{}, "*:Category")
	//根据名字查询目录分级表_这里应该放多一个/category路径下
	web.Router("/admin/categorytitle", &controllers.AdminController{}, "*:CategoryTitle")
	//添加目录类别
	web.Router("/admin/category/addcategory", &controllers.AdminController{}, "*:AddCategory")
	//修改目录类别
	web.Router("/admin/category/updatecategory", &controllers.AdminController{}, "*:UpdateCategory")
	//删除目录类
	web.Router("/admin/category/deletecategory", &controllers.AdminController{}, "*:DeleteCategory")

	//添加IP地址段
	web.Router("/admin/ipsegment/addipsegment", &controllers.AdminController{}, "*:AddIpsegment")
	//修改IP地址段
	web.Router("/admin/ipsegment/updateipsegment", &controllers.AdminController{}, "*:UpdateIpsegment")
	//删除IP地址段
	web.Router("/admin/ipsegment/deleteipsegment", &controllers.AdminController{}, "*:DeleteIpsegment")

	//查询所有ip地址段
	web.Router("/admin/ipsegment", &controllers.AdminController{}, "*:Ipsegment")

	//添加日历
	web.Router("/admin/calendar/addcalendar", &controllers.AdminController{}, "*:AddCalendar")
	//查询
	web.Router("/admin/calendar", &controllers.AdminController{}, "*:Calendar")
	//修改
	web.Router("/admin/calendar/updatecalendar", &controllers.AdminController{}, "*:UpdateCalendar")
	//删除
	web.Router("/admin/calendar/deletecalendar", &controllers.AdminController{}, "*:DeleteCalendar")
	//拖曳事件
	web.Router("/admin/calendar/dropcalendar", &controllers.AdminController{}, "*:DropCalendar")
	//resize事件
	web.Router("/admin/calendar/resizecalendar", &controllers.AdminController{}, "*:ResizeCalendar")
	//搜索事件
	web.Router("/admin/calendar/searchcalendar", &controllers.AdminController{}, "*:SearchCalendar")

	//显示项目目录:注意这个方法放在projcontroller中
	web.Router("/admin/project/getprojectcate/:id:string", &controllers.ProjController{}, "*:GetProjectCate")
	//添加子节点
	web.Router("/admin/project/addprojectcate", &controllers.ProjController{}, "*:AddProjectCate")
	//修改节点名
	web.Router("/admin/project/updateprojectcate", &controllers.ProjController{}, "*:UpdateProjectCate")
	//删除节点和其下子节点——和删除项目一样
	web.Router("/admin/project/deleteprojectcate", &controllers.ProjController{}, "*:DeleteProject")

	//根据项目id查询项目同步ip
	web.Router("/admin/project/synchip/:id:string", &controllers.AdminController{}, "*:SynchIp")
	//添加项目同步ip:注意这是在admincontrollers中
	web.Router("/admin/project/addsynchip", &controllers.AdminController{}, "*:AddsynchIp")
	//修改项目同步ip:注意这是在admincontrollers中
	web.Router("/admin/project/updatesynchip", &controllers.AdminController{}, "*:UpdatesynchIp")
	//删除项目同步ip:注意这是在admincontrollers中
	web.Router("/admin/project/deletesynchip", &controllers.AdminController{}, "*:DeletesynchIp")
	//后台部门结构
	//填充部门表格数据
	web.Router("/admin/department", &controllers.AdminController{}, "*:Department")
	//根据数字id查询类别或目录分级表
	web.Router("/admin/department/?:id:string", &controllers.AdminController{}, "*:Department")
	//根据名字查询目录分级表
	web.Router("/admin/departmenttitle", &controllers.AdminController{}, "*:DepartmentTitle")
	//添加目录类别
	web.Router("/admin/department/adddepartment", &controllers.AdminController{}, "*:AddDepartment")
	//修改目录类别
	web.Router("/admin/department/updatedepartment", &controllers.AdminController{}, "*:UpdateDepartment")
	//删除目录类
	web.Router("/admin/department/deletedepartment", &controllers.AdminController{}, "*:DeleteDepartment")

	//***后台用户管理
	web.Router("/jsoneditor", &controllers.AdminController{}, "get:Jsoneditor")

	//如果后面不带id，则显示所有用户
	// web.Router("/admin/user/?:id:string", &controllers.UserController{}, "*:User")
	// //添加用户
	// web.Router("/admin/user/adduser", &controllers.UserController{}, "*:AddUser")
	// //导入用户
	// web.Router("/admin/user/importusers", &controllers.UserController{}, "*:ImportUsers")

	// //修改用户
	// web.Router("/admin/user/updateuser", &controllers.UserController{}, "*:UpdateUser")
	// //删除用户
	// web.Router("/admin/user/deleteuser", &controllers.UserController{}, "*:DeleteUser")

	//新建角色
	web.Router("/admin/role/post", &controllers.RoleController{}, "post:Post")
	web.Router("/admin/role/update", &controllers.RoleController{}, "put:Update")
	web.Router("/admin/role/delete", &controllers.RoleController{}, "post:Delete")
	web.Router("/admin/role/get/?:id:string", &controllers.RoleController{}, "get:Get")
	//添加用户角色
	web.Router("/admin/role/userrole", &controllers.RoleController{}, "post:UserRole")
	//添加角色对项目目录文件操作权限
	web.Router("/admin/role/permission", &controllers.RoleController{}, "post:RolePermission")
	//查询角色对项目目录文件操作的权限
	web.Router("/admin/role/getpermission", &controllers.RoleController{}, "get:GetRolePermission")

	//meritbasic表格数据填充
	web.Router("/admin/merit/meritbasic", &controllers.AdminController{}, "*:MeritBasic")
	//修改meritbasic表格数据
	web.Router("/admin/merit/updatemeritbasic", &controllers.AdminController{}, "*:UpdateMeritBasic")
	//未提交和已提交的merit的成果列表
	web.Router("/admin/merit/meritlist/:id:int", &controllers.AdminController{}, "*:GetPostMerit")
	//提交成果给merit
	web.Router("/admin/merit/sendmeritlist", &controllers.AdminController{}, "post:SendMeritlist")
	//删除成果merit
	web.Router("/admin/merit/deletemeritlist", &controllers.AdminController{}, "post:DeleteMeritlist")
	//回退merit已提交成果给未提交
	web.Router("/admin/merit/downmeritlist", &controllers.AdminController{}, "post:DownMeritlist")

	//查看附件列表
	web.Router("/admin/merit/meritlist/attachment/:id:string", &controllers.AdminController{}, "get:CatalogAttachment")
	//修改附件
	web.Router("/admin/merit/meritlist/modify", &controllers.AdminController{}, "post:ModifyCatalog")
	//修改附件
	web.Router("/admin/merit/meritlist/modifylink", &controllers.AdminController{}, "post:ModifyLink")

	//用户修改自己密码
	web.Router("/user", &controllers.UserController{}, "get:GetUserByUsername")
	//用户登录后查看自己的资料
	web.Router("/user/getuserbyusername", &controllers.UserController{}, "get:GetUserByUsername")
	//用户产看自己的table中数据填充
	web.Router("/usermyself", &controllers.UserController{}, "get:Usermyself")

	web.Router("/login", &controllers.LoginController{}, "get:Login")
	//页面登录提交用户名和密码
	web.Router("/post", &controllers.LoginController{}, "post:Post")
	//弹框登录提交用户名和密码
	web.Router("/loginpost", &controllers.LoginController{}, "post:LoginPost")
	web.Router("/logout", &controllers.LoginController{}, "get:Logout")
	web.Router("/loginerr", &controllers.LoginController{}, "get:Loginerr")
	web.Router("/roleerr", &controllers.UserController{}, "*:Roleerr") //显示权限不够

	web.Router("/regist", &controllers.RegistController{})
	web.Router("/regist/checkuname", &controllers.RegistController{}, "post:CheckUname")
	web.Router("/regist/getuname", &controllers.RegistController{}, "*:GetUname")
	//项目列表界面
	web.Router("/project", &controllers.ProjController{}, "*:Get")
	//table获取所有项目数据给上面界面使用_后续扩展按标签获取
	web.Router("/project/getprojects", &controllers.ProjController{}, "*:GetProjects")

	//侧栏懒加载下级
	web.Router("/project/getprojcate", &controllers.ProjController{}, "*:GetProjCate")
	//添加项目，应该是project/addproj,delproj,updateproj
	web.Router("/project/addproject", &controllers.ProjController{}, "*:AddProject")
	//添加项目，根据项目模板
	web.Router("/project/addprojtemplet", &controllers.ProjController{}, "*:AddProjTemplet")

	//修改项目
	web.Router("/project/updateproject", &controllers.ProjController{}, "*:UpdateProject")
	//删除项目
	web.Router("/project/deleteproject", &controllers.ProjController{}, "*:DeleteProject")
	//项目时间轴
	web.Router("/project/gettimeline/:id([0-9]+)", &controllers.ProjController{}, "get:ProjectTimeline")
	web.Router("/project/:id([0-9]+)/timeline", &controllers.ProjController{}, "get:Timeline")
	// web.Router("/project/timeline/:id([0-9]+)", &controllers.ProjController{}, "get:Timeline")

	//根据项目id进入一个具体项目的侧栏
	web.Router("/project/:id([0-9]+)", &controllers.ProjController{}, "*:GetProject")
	//进入项目日历
	web.Router("/project/getcalendar/:id([0-9]+)", &controllers.ProjController{}, "*:GetCalendar")
	//取得日历数据
	web.Router("/project/calendar/:id([0-9]+)", &controllers.ProjController{}, "*:Calendar")
	//添加日历
	web.Router("/project/calendar/addcalendar/:id([0-9]+)", &controllers.ProjController{}, "*:AddCalendar")
	//修改
	web.Router("/project/calendar/updatecalendar", &controllers.ProjController{}, "*:UpdateCalendar")
	//删除
	web.Router("/project/calendar/deletecalendar", &controllers.ProjController{}, "*:DeleteCalendar")
	//拖曳事件
	web.Router("/project/calendar/dropcalendar", &controllers.ProjController{}, "*:DropCalendar")
	//resize事件
	web.Router("/project/calendar/resizecalendar", &controllers.ProjController{}, "*:ResizeCalendar")
	//日历中传照片
	web.Router("/project/calendar/uploadimage", &controllers.ProjController{}, "*:UploadImage")

	//点击侧栏，根据id返回json数据给导航条
	web.Router("/project/navbar/:id:string", &controllers.ProjController{}, "*:GetProjNav")

	//根据项目侧栏id显示这个id下的成果界面——作废，用上面GetProject界面
	web.Router("/project/:id:string/:id:string", &controllers.ProdController{}, "*:GetProjProd")

	//给上面那个页面提供table所用的json数据
	web.Router("/project/products/:id:string", &controllers.ProdController{}, "*:GetProducts")
	//点击项目名称，显示这个项目下所有成果
	web.Router("/project/allproducts/:id:string", &controllers.ProjController{}, "*:GetProjProducts")
	//这个对应上面的页面进行表格内数据填充，用的是prodcontrollers中的方法
	web.Router("/project/products/all/:id:string", &controllers.ProdController{}, "*:GetProjProducts")
	//给上面那个页面提供项目同步ip的json数据
	web.Router("/project/synchproducts/:id:string", &controllers.ProdController{}, "*:GetsynchProducts")

	//对外提供同步成果接口json数据
	web.Router("/project/providesynchproducts", &controllers.ProdController{}, "*:ProvidesynchProducts")

	//向项目侧栏id下添加成果_这个没用到，用下面addattachment
	// web.Router("/project/product/addproduct", &controllers.ProdController{}, "post:AddProduct")
	//编辑成果信息
	web.Router("/project/product/updateproduct", &controllers.ProdController{}, "post:UpdateProduct")
	//删除成果
	web.Router("/project/product/deleteproduct", &controllers.ProdController{}, "post:DeleteProduct")
	//取得成果中所有附件中的非pdf附件列表
	web.Router("/project/product/attachment/:id:string", &controllers.AttachController{}, "*:GetAttachments")
	//取得同步成果中所有附件中的非pdf附件列表
	web.Router("/project/product/synchattachment", &controllers.AttachController{}, "*:GetsynchAttachments")
	//提供接口：同步成果中所有附件中的非pdf附件列表
	web.Router("/project/product/providesynchattachment", &controllers.AttachController{}, "*:ProvideAttachments")

	//取得成果中所有附件中——包含pdf，非pdf，文章
	web.Router("/project/product/allattachments/:id:string", &controllers.AttachController{}, "*:GetAllAttachments")

	//取得成果中所有附件的pdf列表
	web.Router("/project/product/pdf/:id:string", &controllers.AttachController{}, "*:GetPdfs")
	//取得同步成果中所有pdf列表
	web.Router("/project/product/synchpdf", &controllers.AttachController{}, "*:GetsynchPdfs")
	//提供同步成果中所有pdf列表
	web.Router("/project/product/providesynchpdf", &controllers.AttachController{}, "*:ProvidePdfs")
	web.Router("/pdf", &controllers.AttachController{}, "*:Pdf")

	//根据文章id取得成果中的文章
	web.Router("/project/product/article/:id:string", &controllers.ArticleController{}, "*:GetArticle")
	//根据成果id取得成果的所有文章列表_注意articles是复数
	web.Router("/project/product/articles/:id:string", &controllers.ArticleController{}, "*:GetArticles")
	//取得同步成果的所有文章列表_注意articles是复数
	web.Router("/project/product/syncharticles", &controllers.ArticleController{}, "*:GetsynchArticles")
	//根据成果id取得成果的所有文章列表_注意articles是复数
	web.Router("/project/product/providesyncharticles", &controllers.ArticleController{}, "*:ProvideArticles")
	//文章进行编辑页面
	web.Router("/project/product/modifyarticle/:id:string", &controllers.ArticleController{}, "*:ModifyArticle")
	//文章进行编辑
	web.Router("/project/product/updatearticle", &controllers.ArticleController{}, "*:UpdateArticle")
	//删除文章
	web.Router("/project/product/deletearticle", &controllers.ArticleController{}, "*:DeleteArticle")

	//查看一个成果
	// web.Router("/project/product/?:id:string"

	//向成果里添加附件：批量一对一模式
	web.Router("/project/product/addattachment", &controllers.AttachController{}, "post:AddAttachment")
	//dwg写入服务器
	web.Router("/project/product/savedwgfile", &controllers.AttachController{}, "post:SaveDwgfile")
	//新建dwg文件
	web.Router("/project/product/newdwg", &controllers.AttachController{}, "post:NewDwg")

	//向成果里添加附件：多附件模式
	web.Router("/project/product/addattachment2", &controllers.AttachController{}, "post:AddAttachment2")
	//编辑附件列表：向成果里追加附件：多附件模式
	web.Router("/project/product/updateattachment", &controllers.AttachController{}, "post:UpdateAttachment")
	//删除附件列表中的附件
	web.Router("/project/product/deleteattachment", &controllers.AttachController{}, "post:DeleteAttachment")

	//向成果里添加文章
	//向侧栏下添加文章
	web.Router("/project/product/addarticle", &controllers.ArticleController{}, "post:AddArticle")
	//在某个项目里搜索成果
	// web.Router("/search", &controllers.SearchController{}, "get:Get")
	// web.Router("/project/product/search", &controllers.SearchController{}, "get:SearchProjProducts")
	//搜索某个项目——没意义
	web.Router("/projects/search", &controllers.SearchController{}, "get:SearchProjects")

	//********gante
	//项目进度甘特图
	web.Router("/projectgant", &controllers.ProjGantController{}, "get:Get")
	//数据填充
	// web.Router("/projectgant/getprojgants", &controllers.ProjGantController{}, "get:GetProjGants")
	//添加项目进度
	web.Router("/projectgant/addprojgant", &controllers.ProjGantController{}, "post:AddProjGant")
	//导入项目进度数据
	web.Router("/projectgant/importprojgant", &controllers.ProjGantController{}, "post:ImportProjGant")

	//修改项目进度
	web.Router("/projectgant/updateprojgant", &controllers.ProjGantController{}, "post:UpdateProjGant")
	//删除项目进度
	web.Router("/projectgant/deleteprojgant", &controllers.ProjGantController{}, "post:DeleteProjGant")
	//关闭项目进度
	web.Router("/projectgant/closeprojgant", &controllers.ProjGantController{}, "post:CloseProjGant")

	//附件下载"/attachment/*", &controllers.AttachController{}
	// web.InsertFilter("/attachment/*", web.BeforeRouter, controllers.ImageFilter)
	//根据附件绝对地址下载
	web.Router("/attachment/*", &controllers.AttachController{}, "get:Attachment")
	//根据附件id号，判断权限下载
	web.Router("/downloadattachment", &controllers.AttachController{}, "get:DownloadAttachment")

	//上面用attachment.ImageFilter是不行的，必须是package.func
	//首页轮播图片的权限
	web.Router("/attachment/carousel/*.*", &controllers.AttachController{}, "get:GetCarousel")

	//ue富文本编辑器
	web.Router("/controller", &controllers.UeditorController{}, "*:ControllerUE")
	//添加文章——froala上传插入的图片
	web.Router("/uploadimg", &controllers.FroalaController{}, "*:UploadImg")
	//添加wiki——froala上传插入的图片
	web.Router("/uploadwikiimg", &controllers.FroalaController{}, "*:UploadWikiImg")
	//添加文章——froala上传插入的视频
	web.Router("/uploadvideo", &controllers.FroalaController{}, "*:UploadVideo")

	//添加日历
	web.Router("/index/carcalendar/addcalendar", &controllers.IndexController{}, "*:AddCarCalendar")
	//查询
	web.Router("/index/carcalendar", &controllers.IndexController{}, "*:CarCalendar")
	//修改
	web.Router("/index/carcalendar/updatecalendar", &controllers.IndexController{}, "*:UpdateCarCalendar")
	//删除
	web.Router("/index/carcalendar/deletecalendar", &controllers.IndexController{}, "*:DeleteCarCalendar")
	//拖曳事件
	web.Router("/index/carcalendar/dropcalendar", &controllers.IndexController{}, "*:DropCarCalendar")
	//resize事件
	web.Router("/index/carcalendar/resizecalendar", &controllers.IndexController{}, "*:ResizeCarCalendar")

	//添加日历
	web.Router("/index/meetingroomcalendar/addcalendar", &controllers.IndexController{}, "*:AddMeetCalendar")
	//查询
	web.Router("/index/meetingroomcalendar", &controllers.IndexController{}, "*:MeetCalendar")
	//修改
	web.Router("/index/meetingroomcalendar/updatecalendar", &controllers.IndexController{}, "*:UpdateMeetCalendar")
	//删除
	web.Router("/index/meetingroomcalendar/deletecalendar", &controllers.IndexController{}, "*:DeleteMeetCalendar")
	//拖曳事件
	web.Router("/index/meetingroomcalendar/dropcalendar", &controllers.IndexController{}, "*:DropMeetCalendar")
	//resize事件
	web.Router("/index/meetingroomcalendar/resizecalendar", &controllers.IndexController{}, "*:ResizeMeetCalendar")

	//wiki页面
	web.Router("/wiki", &controllers.WikiController{}) //get
	//发表文章界面
	web.Router("/wiki/add", &controllers.WikiController{}, "get:Add")
	//发表文章提交
	web.Router("/wiki/addwiki", &controllers.WikiController{}, "post:AddWiki")

	//查看一个文章
	web.Router("/wiki/view/", &controllers.WikiController{}, "get:View")
	//删除wiki
	web.Router("/wiki/delete", &controllers.WikiController{}, "post:Delete")
	//修改一个文章
	// web.Router("/wiki/modify", &controllers.WikiController{}, "get:Modify")
	web.AutoRouter(&controllers.WikiController{})
	//添加删除wiki的评论
	web.Router("/reply/addwiki", &controllers.ReplyController{}, "post:AddWiki")
	web.Router("/reply/deletewiki", &controllers.ReplyController{}, "get:DeleteWiki")

	// web.InsertFilter("/attachment/standard/", web.BeforeRouter, FilterUser)
	// web.SetStaticPath("/attachment/standard/", "attachment/standard/")
	//这个有哦何用？
	web.SetStaticPath("/attachment/wiki", "attachment/wiki")
	web.SetStaticPath("/swagger", "swagger")
	// *全匹配方式 //匹配 /download/ceshi/file/api.json :splat=file/api.json
	web.Router("/searchwiki", &controllers.SearchController{}, "get:SearchWiki")

	//规范管理
	web.Router("/standard", &controllers.StandardController{}, "get:Index")
	web.Router("/standard/search", &controllers.StandardController{}, "get:Search")
	web.Router("/standard/importexcel", &controllers.StandardController{}, "post:ImportExcel")
	web.Router("/standard/standard_one_addbaidu", &controllers.StandardController{}, "post:Standard_one_addbaidu")
	web.Router("/standard/importlibrary", &controllers.StandardController{}, "post:ImportLibrary")
	//显示规范所有
	web.Router("/standard/getstandard", &controllers.StandardController{}, "get:GetStandard")
	//修改规范库
	web.Router("/standard/updatestandard", &controllers.StandardController{}, "post:UpdateStandard")
	//删除规范库
	web.Router("/standard/deletestandard", &controllers.StandardController{}, "post:DeleteStandard")

	//显示有效库所有
	web.Router("/standard/valid", &controllers.StandardController{}, "get:Valid")
	//删除有效库中选中
	web.Router("/standard/deletevalid", &controllers.StandardController{}, "post:DeleteValid")

	//对标
	web.Router("/legislation", &controllers.LegislationController{}, "*:Index")
	web.Router("/legislation/checklist", &controllers.LegislationController{}, "*:Checklist")
	//上传文档，提供解析和替换（增加）标准号
	web.Router("/legislation/fileinput", &controllers.LegislationController{}, "get:FileInput")

	//设代日志列表页面
	web.Router("/diary", &controllers.DiaryController{}, "get:Get") //get
	//具体一个日志页面
	web.Router("/getwxdiary2/:id", &controllers.DiaryController{}, "get:GetWxDiary2")
	//转换为word
	// web.Router("/htmltodoc", &controllers.DiaryController{}, "get:HtmlToDoc")

	//设代财务登记列表页面
	web.Router("/finance", &controllers.FinanceController{}, "get:Get") //get
	//具体一个财务登记页面
	web.Router("/getwxfinance2/:id", &controllers.FinanceController{}, "get:GetWxFinance2")

	//微信小程序
	//小程序发表文章提交
	// web.Router("/wx/addwxarticle", &controllers.ArticleController{}, "post:AddWxArticle")
	//小程序上传图片，返回地址
	// web.Router("/wx/uploadwximg", &controllers.FroalaController{}, "*:UploadWxImg")
	//小程序获得文章列表，分页
	// web.Router("/wx/getwxarticles", &controllers.ArticleController{}, "*:GetWxArticles")
	//小程序根据文章id返回文章数据
	// web.Router("/wx/getwxarticle/:id:string", &controllers.ArticleController{}, "*:GetWxArticle")
	//微信登录
	// web.Router("/wx/wxlogin", &controllers.LoginController{}, "*:WxLogin")

	web.Any("/cors-anywhere", func(ctx *context.Context) {
		u, _ := url.PathUnescape(ctx.Input.Query("url"))
		logs.Error("ReverseProxy: ", u)
		if len(u) > 0 && strings.HasPrefix(u, "http") {
			if strings.TrimRight(conf.BaseUrl, "/") == ctx.Input.Site() {
				ctx.Redirect(302, u)
			} else {
				target, _ := url.Parse(u)
				logs.Debug("target: ", target)

				proxy := &httputil.ReverseProxy{
					Transport: roundTripper(rt),
					Director: func(req *http.Request) {
						req.Header = ctx.Request.Header
						req.URL.Scheme = target.Scheme
						req.URL.Host = target.Host
						req.URL.Path = target.Path
						req.Header.Set("Host", target.Host)
					},
				}

				http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
				proxy.ServeHTTP(ctx.ResponseWriter, ctx.Request)
			}
		} else {
			ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
			ctx.Output.Body([]byte("400 Bad Request"))
		}
	})

	web.Router("/mindoc", &controllers.HomeController{}, "*:Index")

	web.Router("/mindoclogin", &controllers.AccountController{}, "*:Login")
	web.Router("/dingtalk_login", &controllers.AccountController{}, "*:DingTalkLogin")
	web.Router("/workweixin-login", &controllers.AccountController{}, "*:WorkWeixinLogin")
	web.Router("/workweixin-callback", &controllers.AccountController{}, "*:WorkWeixinLoginCallback")
	web.Router("/workweixin-bind", &controllers.AccountController{}, "*:WorkWeixinLoginBind") //urlfor params must key-value pair
	web.Router("/workweixin-ignore", &controllers.AccountController{}, "*:WorkWeixinLoginIgnore")
	web.Router("/qrlogin/:app", &controllers.AccountController{}, "*:QRLogin")
	web.Router("/mindoclogout", &controllers.AccountController{}, "*:Logout")
	web.Router("/register", &controllers.AccountController{}, "*:Register")
	web.Router("/find_password", &controllers.AccountController{}, "*:FindPassword")
	web.Router("/valid_email", &controllers.AccountController{}, "post:ValidEmail")
	web.Router("/captcha", &controllers.AccountController{}, "*:Captcha")

	web.Router("/manager", &controllers.ManagerController{}, "*:Index")
	web.Router("/manager/users", &controllers.ManagerController{}, "*:Users")
	web.Router("/manager/users/edit/:id", &controllers.ManagerController{}, "*:EditMember")
	web.Router("/manager/member/create", &controllers.ManagerController{}, "post:CreateMember")
	web.Router("/manager/member/delete", &controllers.ManagerController{}, "post:DeleteMember")
	web.Router("/manager/member/update-member-status", &controllers.ManagerController{}, "post:UpdateMemberStatus")
	web.Router("/manager/member/change-member-role", &controllers.ManagerController{}, "post:ChangeMemberRole")
	web.Router("/manager/books", &controllers.ManagerController{}, "*:Books")
	web.Router("/manager/books/edit/:key", &controllers.ManagerController{}, "*:EditBook")
	web.Router("/manager/books/delete", &controllers.ManagerController{}, "*:DeleteBook")

	web.Router("/manager/comments", &controllers.ManagerController{}, "*:Comments")
	web.Router("/manager/setting", &controllers.ManagerController{}, "*:Setting")
	web.Router("/manager/books/token", &controllers.ManagerController{}, "post:CreateToken")
	web.Router("/manager/books/transfer", &controllers.ManagerController{}, "post:Transfer")
	web.Router("/manager/books/open", &controllers.ManagerController{}, "post:PrivatelyOwned")

	web.Router("/manager/attach/list", &controllers.ManagerController{}, "*:AttachList")
	web.Router("/manager/attach/detailed/:id", &controllers.ManagerController{}, "*:AttachDetailed")
	web.Router("/manager/attach/delete", &controllers.ManagerController{}, "post:AttachDelete")
	web.Router("/manager/label/list", &controllers.ManagerController{}, "get:LabelList")
	web.Router("/manager/label/delete/:id", &controllers.ManagerController{}, "post:LabelDelete")

	//web.Router("/manager/config",  &controllers.ManagerController{}, "*:Config")

	web.Router("/manager/team", &controllers.ManagerController{}, "*:Team")
	web.Router("/manager/team/create", &controllers.ManagerController{}, "POST:TeamCreate")
	web.Router("/manager/team/edit", &controllers.ManagerController{}, "POST:TeamEdit")
	web.Router("/manager/team/delete", &controllers.ManagerController{}, "POST:TeamDelete")

	web.Router("/manager/team/member/list/:id", &controllers.ManagerController{}, "*:TeamMemberList")
	web.Router("/manager/team/member/add", &controllers.ManagerController{}, "POST:TeamMemberAdd")
	web.Router("/manager/team/member/delete", &controllers.ManagerController{}, "POST:TeamMemberDelete")
	web.Router("/manager/team/member/change_role", &controllers.ManagerController{}, "POST:TeamChangeMemberRole")
	web.Router("/manager/team/member/search", &controllers.ManagerController{}, "*:TeamSearchMember")

	web.Router("/manager/team/book/list/:id", &controllers.ManagerController{}, "*:TeamBookList")
	web.Router("/manager/team/book/add", &controllers.ManagerController{}, "POST:TeamBookAdd")
	web.Router("/manager/team/book/delete", &controllers.ManagerController{}, "POST:TeamBookDelete")
	web.Router("/manager/team/book/search", &controllers.ManagerController{}, "*:TeamSearchBook")

	web.Router("/manager/itemsets", &controllers.ManagerController{}, "*:Itemsets")
	web.Router("/manager/itemsets/edit", &controllers.ManagerController{}, "post:ItemsetsEdit")
	web.Router("/manager/itemsets/delete", &controllers.ManagerController{}, "post:ItemsetsDelete")

	web.Router("/setting", &controllers.SettingController{}, "*:Index")
	web.Router("/setting/password", &controllers.SettingController{}, "*:Password")
	web.Router("/setting/upload", &controllers.SettingController{}, "*:Upload")

	web.Router("/book", &controllers.BookController{}, "*:Index")
	web.Router("/book/:key/dashboard", &controllers.BookController{}, "*:Dashboard")
	web.Router("/book/:key/setting", &controllers.BookController{}, "*:Setting")
	web.Router("/book/:key/users", &controllers.BookController{}, "*:Users")
	web.Router("/book/:key/release", &controllers.BookController{}, "post:Release")
	web.Router("/book/:key/sort", &controllers.BookController{}, "post:SaveSort")
	web.Router("/book/:key/teams", &controllers.BookController{}, "*:Team")

	web.Router("/book/create", &controllers.BookController{}, "*:Create")
	web.Router("/book/itemsets/search", &controllers.BookController{}, "*:ItemsetsSearch")

	web.Router("/book/users/create", &controllers.BookMemberController{}, "post:AddMember")
	web.Router("/book/users/change", &controllers.BookMemberController{}, "post:ChangeRole")
	web.Router("/book/users/delete", &controllers.BookMemberController{}, "post:RemoveMember")
	web.Router("/book/users/import", &controllers.BookController{}, "post:Import")
	web.Router("/book/users/copy", &controllers.BookController{}, "post:Copy")

	web.Router("/book/setting/save", &controllers.BookController{}, "post:SaveBook")
	web.Router("/book/setting/open", &controllers.BookController{}, "post:PrivatelyOwned")
	web.Router("/book/setting/transfer", &controllers.BookController{}, "post:Transfer")
	web.Router("/book/setting/upload", &controllers.BookController{}, "post:UploadCover")
	web.Router("/book/setting/delete", &controllers.BookController{}, "post:Delete")

	web.Router("/book/team/add", &controllers.BookController{}, "POST:TeamAdd")
	web.Router("/book/team/delete", &controllers.BookController{}, "POST:TeamDelete")
	web.Router("/book/team/search", &controllers.BookController{}, "*:TeamSearch")

	//管理文章的路由
	web.Router("/manage/blogs", &controllers.BlogController{}, "*:ManageList")
	web.Router("/manage/blogs/setting/?:id", &controllers.BlogController{}, "*:ManageSetting")
	web.Router("/manage/blogs/edit/?:id", &controllers.BlogController{}, "*:ManageEdit")
	web.Router("/manage/blogs/delete", &controllers.BlogController{}, "post:ManageDelete")
	web.Router("/manage/blogs/upload", &controllers.BlogController{}, "post:Upload")
	web.Router("/manage/blogs/attach/:id", &controllers.BlogController{}, "post:RemoveAttachment")

	//读文章的路由
	web.Router("/blogs", &controllers.BlogController{}, "*:List")
	web.Router("/blog-attach/:id:int/:attach_id:int", &controllers.BlogController{}, "get:Download")
	web.Router("/blog-:id([0-9]+).html", &controllers.BlogController{}, "*:Index")

	//模板相关接口
	web.Router("/api/template/get", &controllers.TemplateController{}, "get:Get")
	web.Router("/api/template/list", &controllers.TemplateController{}, "post:List")
	web.Router("/api/template/add", &controllers.TemplateController{}, "post:Add")
	web.Router("/api/template/remove", &controllers.TemplateController{}, "post:Delete")

	web.Router("/api/attach/remove/", &controllers.DocumentController{}, "post:RemoveAttachment")
	web.Router("/api/:key/edit/?:id", &controllers.DocumentController{}, "*:Edit")
	web.Router("/api/upload", &controllers.DocumentController{}, "post:Upload")
	web.Router("/api/:key/create", &controllers.DocumentController{}, "post:Create")
	web.Router("/api/:key/delete", &controllers.DocumentController{}, "post:Delete")
	web.Router("/api/:key/content/?:id", &controllers.DocumentController{}, "*:Content")
	web.Router("/api/:key/compare/:id", &controllers.DocumentController{}, "*:Compare")
	web.Router("/api/search/user/:key", &controllers.MindocSearchController{}, "*:User")

	web.Router("/history/get", &controllers.DocumentController{}, "get:History")
	web.Router("/history/delete", &controllers.DocumentController{}, "*:DeleteHistory")
	web.Router("/history/restore", &controllers.DocumentController{}, "*:RestoreHistory")

	web.Router("/docs/:key", &controllers.DocumentController{}, "*:Index")
	web.Router("/docs/:key/check-password", &controllers.DocumentController{}, "post:CheckPassword")

	web.Router("/docs/:key/:id", &controllers.DocumentController{}, "*:Read")

	web.Router("/docs/:key/search", &controllers.DocumentController{}, "post:Search")
	web.Router("/export/:key", &controllers.DocumentController{}, "*:Export")
	web.Router("/qrcode/:key.png", &controllers.DocumentController{}, "get:QrCode")

	web.Router("/attach_files/:key/:attach_id", &controllers.DocumentController{}, "get:DownloadAttachment")

	web.Router("/comment/create", &controllers.CommentController{}, "post:Create")
	web.Router("/comment/delete", &controllers.CommentController{}, "post:Delete")
	web.Router("/comment/lists", &controllers.CommentController{}, "get:Lists")
	web.Router("/comment/index", &controllers.CommentController{}, "*:Index")

	web.Router("/mindocsearch", &controllers.MindocSearchController{}, "get:Index")

	web.Router("/tag/:key", &controllers.LabelController{}, "get:Index")
	web.Router("/tags", &controllers.LabelController{}, "get:List")

	web.Router("/items", &controllers.ItemsetsController{}, "get:Index")
	web.Router("/items/:key", &controllers.ItemsetsController{}, "get:List")

}

// 那么Session在何时创建呢？当然还是在服务器端程序运行的过程中创建的，
// 不同语言实现的应用程序有不同创建Session的方法，
// 而在Java中是通过调用HttpServletRequest的getSession方法（使用true作为参数）创建的。
// 在创建了Session的同时，服务器会为该Session生成唯一的Session id，
// 而这个Session id在随后的请求中会被用来重新获得已经创建的Session；
// 在Session被创建之后，就可以调用Session相关的方法往Session中增加内容了，
// 而这些内容只会保存在服务器中，发到客户端的只有Session id；
// 当客户端再次发送请求的时候，会将这个Session id带上，
// 服务器接受到请求之后就会依据Session id找到相应的Session，从而再次使用之
