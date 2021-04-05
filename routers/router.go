// @APIVersion 1.0.0
// @Title EngineerCMS API
// @Description ECMS has every tool to get any job done, so codename for the new ECMS APIs.
// @Contact 504284@qq.com
package routers

import (
	"github.com/3xxx/engineercms/controllers"
	// "githsub.com/3xxx/engineercms/controllers/checkin"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/plugins/cors"
	// "github.com/3xxx/engineercms/controllers"
)

// var FilterFunc = func(ctx *context.Context) {
// 	userName := ctx.Input.Session("userName")
// 	if userName == nil {
// 		ctx.Redirect()
// 	}
// 	v := ctx.Input.CruSession.Get("uname") //用来获取存储在服务器端中的数据??。
// 	// beego.Info(v)                          //qin.xc
// 	var user models.User
// 	var err error
// 	if v != nil { //如果登录了
// 		uname = v.(string)
// 		user, err = models.GetUserByUsername(uname)
// 		if err != nil {
// 			beego.Error(err)
// 		} else {
// 			uid = user.Id
// 			role = user.Role
// 		}
// 	} else { //如果没登录
// 		role = "anonymous"
// 	}
// }

func init() {
	//运行跨域请求
	//在http请求的响应流头部加上如下信息
	//rw.Header().Set("Access-Control-Allow-Origin", "*")
	// beego.InsertFilter("?=/article/*", beego.BeforeRouter, FilterFunc)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
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
	var FilterUser = func(ctx *context.Context) {
		// v := ctx.Input.CruSession.Get("uname")
		v := ctx.Input.Session("uname")
		// uname = v.(string)//uid---v.(int)
		if v == nil {
			ctx.Redirect(302, "/login")
		}
	}

	//自动化文档
	ns :=
		beego.NewNamespace("/v1",
			// beego.NSBefore(FilterFunc),
			// beego.NSBefore(auth),
			beego.NSNamespace("/admin",
				beego.NSInclude(
					&controllers.AdminController{},
					&controllers.FlowController{},
					// &controllers.AttachController{},20200626调整
					// &controllers.LoginController{},
					// &controllers.CustomerCookieCheckerController{},
				),
			),
			beego.NSNamespace("/wx",
				beego.NSInclude(
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
					&controllers.VideoController{},
					&controllers.BusinessController{},
					&controllers.LocationController{},
				),
			),
			beego.NSNamespace("/share",
				beego.NSInclude(
					&controllers.ShareController{},
				),
			),
			beego.NSNamespace("/adminlog",
				beego.NSInclude(
					&controllers.AdminLogController{},
				),
			),
			beego.NSNamespace("/project",
				beego.NSInclude(
					&controllers.ProjController{},
					&controllers.ProdController{},
				),
			),
			beego.NSNamespace("/checkin",
				beego.NSInclude(
					&controllers.CheckController{},
				),
			),
			beego.NSNamespace("/bbs",
				beego.NSInclude(
					&controllers.BbsController{},
				),
			),
			beego.NSNamespace("/todo",
				beego.NSInclude(
					&controllers.TodoController{},
				),
			),
			beego.NSNamespace("/onlyoffice",
				beego.NSInclude(
					&controllers.OnlyController{},
				),
			),
			beego.NSNamespace("/fileinput",
				beego.NSInclude(
					&controllers.FileinputController{},
				),
			),
			beego.NSNamespace("/pdfcpu",
				beego.NSBefore(FilterUser),
				beego.NSInclude(
					&controllers.PdfCpuController{},
				),
			),
			beego.NSNamespace("/flv",
				// beego.NSBefore(FilterUser),
				beego.NSInclude(
					&controllers.FlvController{},
				),
			),
			beego.NSNamespace("/cart",
				// beego.NSBefore(FilterUser),
				beego.NSInclude(
					&controllers.CartController{},
				),
			),
			// beego.NSNamespace("/cms",
			// 	beego.NSInclude(
			// 		&controllers.CMSController{},
			// 	),
			// ),
			// beego.NSNamespace("/suggest",
			// 	beego.NSInclude(
			// 		&controllers.SearchController{},
			// 	),
			// ),
		)
	beego.AddNamespace(ns)

	beego.Router("/debug/pprof", &controllers.ProfController{})
	beego.Router("/debug/pprof/:app([\\w]+)", &controllers.ProfController{})

	beego.Router("/test", &controllers.MainController{}, "*:Test")
	beego.Router("/autodesk", &controllers.MainController{}, "*:Autodesk")
	// beego.Router("/usermanage", &controllers.MainController{}, "*:UserManage")
	// beego.Router("/.well-known/pki-validation/AC9A20F9BD09F18D247337AABC67BC06.txt", &controllers.AdminController{}, "*:Testdown")
	beego.Router("/.well-known/pki-validation/*", &controllers.AdminController{}, "*:Testdown")

	beego.Router("/doctree", &controllers.OnlyController{}, "*:GetTree")
	//升级数据库
	beego.Router("/updatedatabase", &controllers.MainController{}, "*:UpdateDatabase")
	//删除数据表和字段测试
	beego.Router("/modifydatabase", &controllers.MainController{}, "*:ModifyDatabase")

	beego.Router("/url-to-callback", &controllers.OnlyController{}, "*:UrltoCallback")
	//cms中预览office回调
	beego.Router("/officeviewcallback", &controllers.OnlyController{}, "*:OfficeViewCallback")

	// 访问onlyoffice页面需要登录
	// beego.InsertFilter("/onlyoffice", beego.BeforeRouter, FilterUser)

	// beego.Router("/onlyoffice/post", &controllers.OnlyController{}, "post:PostOnlyoffice")
	beego.Router("/onlyoffice", &controllers.OnlyController{}, "get:Get")
	//table获取所有数据给上面界面使用
	beego.Router("/onlyoffice/getdata", &controllers.OnlyController{}, "*:GetData")
	//添加一个文档
	beego.Router("/onlyoffice/addattachment", &controllers.OnlyController{}, "post:AddOnlyAttachment")
	//在onlyoffice中打开文档协作
	// 协作页面下载文档需要登录
	// beego.InsertFilter("/onlyoffice/:id:string", beego.BeforeRouter, FilterUser)
	beego.Router("/onlyoffice/:id:string", &controllers.OnlyController{}, "*:OnlyOffice")
	//cms中预览office
	beego.Router("/officeview/:id:string", &controllers.OnlyController{}, "*:OfficeView")

	//删除
	beego.Router("/onlyoffice/deletedoc", &controllers.OnlyController{}, "*:DeleteDoc")
	//修改
	beego.Router("/onlyoffice/updatedoc", &controllers.OnlyController{}, "*:UpdateDoc")
	//onlyoffice页面下载doc
	beego.Router("/attachment/onlyoffice/*", &controllers.OnlyController{}, "get:DownloadDoc")
	//文档管理页面下载doc
	beego.Router("/onlyoffice/download/:id:string", &controllers.OnlyController{}, "get:Download")
	// beego.Router("/onlyoffice/changes", &controllers.OnlyController{}, "post:ChangesUrl")
	//*****onlyoffice document权限
	//添加用户和角色权限
	beego.Router("/onlyoffice/addpermission", &controllers.OnlyController{}, "post:Addpermission")
	//取得文档的用户和角色权限列表
	beego.Router("/onlyoffice/getpermission", &controllers.OnlyController{}, "get:Getpermission")

	beego.Router("/role/test", &controllers.RoleController{}, "*:Test")
	beego.Router("/1/slide", &controllers.MainController{}, "*:Slide")
	// beego.Router("/postdata", &controllers.MainController{}, "*:Postdata")
	//文档
	beego.Router("/doc/ecms", &controllers.MainController{}, "get:Getecmsdoc")
	beego.Router("/doc/meritms", &controllers.MainController{}, "get:Getmeritmsdoc")
	beego.Router("/doc/hydrows", &controllers.MainController{}, "get:Gethydrowsdoc")

	//api接口
	beego.Router("/api/ecms", &controllers.MainController{}, "get:Getecmsapi")
	beego.Router("/api/meritms", &controllers.MainController{}, "get:Getmeritmsapi")
	//根据app.conf里的设置，显示首页
	beego.Router("/", &controllers.MainController{}, "get:Get")
	beego.Router("/cms", &controllers.IndexController{}, "get:Cms")
	//显示首页
	beego.Router("/index", &controllers.IndexController{}, "*:GetIndex")
	//首页放到onlyoffice
	// beego.Router("/", &controllers.OnlyController{}, "get:Get")

	//显示右侧页面框架
	beego.Router("/index/user", &controllers.IndexController{}, "*:GetUser")
	//这里显示用户查看主人日程
	beego.Router("/calendar", &controllers.IndexController{}, "*:Calendar")
	//会议室预定日程
	beego.Router("/meetingroom", &controllers.IndexController{}, "*:MeetingroomCalendar")
	beego.Router("/meetingroom/searchcalendar", &controllers.IndexController{}, "*:SearchCalendar")

	//车辆预定日程
	beego.Router("/car", &controllers.IndexController{}, "*:GetCarCalendar")
	//订餐
	beego.Router("/order", &controllers.IndexController{}, "*:GetOrderCalendar")
	//考勤
	beego.Router("/attendance", &controllers.IndexController{}, "*:GetAttendanceCalendar")

	//首页搜索项目或成果
	beego.Router("/index/searchproject", &controllers.SearchController{}, "*:SearchProject")
	// beego.Router("/index/searchproduct", &controllers.SearchController{}, "*:SearchProduct")

	//后台
	beego.Router("/admin", &controllers.AdminController{})
	// beego.Router("/admincategory", &controllers.AdminController{}, "*:GetAdminCategory")
	//显示对应侧栏id的右侧界面
	beego.Router("/admin/:id:string", &controllers.AdminController{}, "*:Admin")

	//批量添加首页轮播图片
	beego.Router("/admin/base/addcarousel", &controllers.AdminController{}, "*:AddCarousel")
	//获取首页轮播图片填充表格
	beego.Router("/admin/base/carousel", &controllers.AdminController{}, "*:Carousel")
	//删除首页轮播图片
	beego.Router("/admin/base/deletecarousel", &controllers.AdminController{}, "*:DeleteCarousel")

	//根据数字id查询类别或目录分级表
	beego.Router("/admin/category/?:id:string", &controllers.AdminController{}, "*:Category")
	//根据名字查询目录分级表_这里应该放多一个/category路径下
	beego.Router("/admin/categorytitle", &controllers.AdminController{}, "*:CategoryTitle")
	//添加目录类别
	beego.Router("/admin/category/addcategory", &controllers.AdminController{}, "*:AddCategory")
	//修改目录类别
	beego.Router("/admin/category/updatecategory", &controllers.AdminController{}, "*:UpdateCategory")
	//删除目录类
	beego.Router("/admin/category/deletecategory", &controllers.AdminController{}, "*:DeleteCategory")

	//添加IP地址段
	beego.Router("/admin/ipsegment/addipsegment", &controllers.AdminController{}, "*:AddIpsegment")
	//修改IP地址段
	beego.Router("/admin/ipsegment/updateipsegment", &controllers.AdminController{}, "*:UpdateIpsegment")
	//删除IP地址段
	beego.Router("/admin/ipsegment/deleteipsegment", &controllers.AdminController{}, "*:DeleteIpsegment")

	//查询所有ip地址段
	beego.Router("/admin/ipsegment", &controllers.AdminController{}, "*:Ipsegment")

	//添加日历
	beego.Router("/admin/calendar/addcalendar", &controllers.AdminController{}, "*:AddCalendar")
	//查询
	beego.Router("/admin/calendar", &controllers.AdminController{}, "*:Calendar")
	//修改
	beego.Router("/admin/calendar/updatecalendar", &controllers.AdminController{}, "*:UpdateCalendar")
	//删除
	beego.Router("/admin/calendar/deletecalendar", &controllers.AdminController{}, "*:DeleteCalendar")
	//拖曳事件
	beego.Router("/admin/calendar/dropcalendar", &controllers.AdminController{}, "*:DropCalendar")
	//resize事件
	beego.Router("/admin/calendar/resizecalendar", &controllers.AdminController{}, "*:ResizeCalendar")
	//搜索事件
	beego.Router("/admin/calendar/searchcalendar", &controllers.AdminController{}, "*:SearchCalendar")

	//显示项目目录:注意这个方法放在projcontroller中
	beego.Router("/admin/project/getprojectcate/:id:string", &controllers.ProjController{}, "*:GetProjectCate")
	//添加子节点
	beego.Router("/admin/project/addprojectcate", &controllers.ProjController{}, "*:AddProjectCate")
	//修改节点名
	beego.Router("/admin/project/updateprojectcate", &controllers.ProjController{}, "*:UpdateProjectCate")
	//删除节点和其下子节点——和删除项目一样
	beego.Router("/admin/project/deleteprojectcate", &controllers.ProjController{}, "*:DeleteProject")

	//根据项目id查询项目同步ip
	beego.Router("/admin/project/synchip/:id:string", &controllers.AdminController{}, "*:SynchIp")
	//添加项目同步ip:注意这是在admincontrollers中
	beego.Router("/admin/project/addsynchip", &controllers.AdminController{}, "*:AddsynchIp")
	//修改项目同步ip:注意这是在admincontrollers中
	beego.Router("/admin/project/updatesynchip", &controllers.AdminController{}, "*:UpdatesynchIp")
	//删除项目同步ip:注意这是在admincontrollers中
	beego.Router("/admin/project/deletesynchip", &controllers.AdminController{}, "*:DeletesynchIp")
	//后台部门结构
	//填充部门表格数据
	beego.Router("/admin/department", &controllers.AdminController{}, "*:Department")
	//根据数字id查询类别或目录分级表
	beego.Router("/admin/department/?:id:string", &controllers.AdminController{}, "*:Department")
	//根据名字查询目录分级表
	beego.Router("/admin/departmenttitle", &controllers.AdminController{}, "*:DepartmentTitle")
	//添加目录类别
	beego.Router("/admin/department/adddepartment", &controllers.AdminController{}, "*:AddDepartment")
	//修改目录类别
	beego.Router("/admin/department/updatedepartment", &controllers.AdminController{}, "*:UpdateDepartment")
	//删除目录类
	beego.Router("/admin/department/deletedepartment", &controllers.AdminController{}, "*:DeleteDepartment")

	//***后台用户管理
	beego.Router("/jsoneditor", &controllers.AdminController{}, "get:Jsoneditor")

	//如果后面不带id，则显示所有用户
	beego.Router("/admin/user/?:id:string", &controllers.UserController{}, "*:User")
	//添加用户
	beego.Router("/admin/user/adduser", &controllers.UserController{}, "*:AddUser")
	//导入用户
	beego.Router("/admin/user/importusers", &controllers.UserController{}, "*:ImportUsers")

	//修改用户
	beego.Router("/admin/user/updateuser", &controllers.UserController{}, "*:UpdateUser")
	//删除用户
	beego.Router("/admin/user/deleteuser", &controllers.UserController{}, "*:DeleteUser")

	//新建角色
	beego.Router("/admin/role/post", &controllers.RoleController{}, "post:Post")
	beego.Router("/admin/role/update", &controllers.RoleController{}, "put:Update")
	beego.Router("/admin/role/delete", &controllers.RoleController{}, "post:Delete")
	beego.Router("/admin/role/get/?:id:string", &controllers.RoleController{}, "get:Get")
	//添加用户角色
	beego.Router("/admin/role/userrole", &controllers.RoleController{}, "post:UserRole")
	//添加角色对项目目录文件操作权限
	beego.Router("/admin/role/permission", &controllers.RoleController{}, "post:RolePermission")
	//查询角色对项目目录文件操作的权限
	beego.Router("/admin/role/getpermission", &controllers.RoleController{}, "get:GetRolePermission")

	//meritbasic表格数据填充
	beego.Router("/admin/merit/meritbasic", &controllers.AdminController{}, "*:MeritBasic")
	//修改meritbasic表格数据
	beego.Router("/admin/merit/updatemeritbasic", &controllers.AdminController{}, "*:UpdateMeritBasic")
	//未提交和已提交的merit的成果列表
	beego.Router("/admin/merit/meritlist/:id:int", &controllers.AdminController{}, "*:GetPostMerit")
	//提交成果给merit
	beego.Router("/admin/merit/sendmeritlist", &controllers.AdminController{}, "post:SendMeritlist")
	//删除成果merit
	beego.Router("/admin/merit/deletemeritlist", &controllers.AdminController{}, "post:DeleteMeritlist")
	//回退merit已提交成果给未提交
	beego.Router("/admin/merit/downmeritlist", &controllers.AdminController{}, "post:DownMeritlist")

	//查看附件列表
	beego.Router("/admin/merit/meritlist/attachment/:id:string", &controllers.AdminController{}, "get:CatalogAttachment")
	//修改附件
	beego.Router("/admin/merit/meritlist/modify", &controllers.AdminController{}, "post:ModifyCatalog")
	//修改附件
	beego.Router("/admin/merit/meritlist/modifylink", &controllers.AdminController{}, "post:ModifyLink")

	//用户修改自己密码
	beego.Router("/user", &controllers.UserController{}, "get:GetUserByUsername")
	//用户登录后查看自己的资料
	beego.Router("/user/getuserbyusername", &controllers.UserController{}, "get:GetUserByUsername")
	//用户产看自己的table中数据填充
	beego.Router("/usermyself", &controllers.UserController{}, "get:Usermyself")

	beego.Router("/login", &controllers.LoginController{}, "get:Login")
	//页面登录提交用户名和密码
	beego.Router("/post", &controllers.LoginController{}, "post:Post")
	//弹框登录提交用户名和密码
	beego.Router("/loginpost", &controllers.LoginController{}, "post:LoginPost")
	beego.Router("/logout", &controllers.LoginController{}, "get:Logout")
	beego.Router("/loginerr", &controllers.LoginController{}, "get:Loginerr")
	beego.Router("/roleerr", &controllers.UserController{}, "*:Roleerr") //显示权限不够

	beego.Router("/regist", &controllers.RegistController{})
	beego.Router("/regist/checkuname", &controllers.RegistController{}, "post:CheckUname")
	beego.Router("/regist/getuname", &controllers.RegistController{}, "*:GetUname")
	//项目列表界面
	beego.Router("/project", &controllers.ProjController{}, "*:Get")
	//table获取所有项目数据给上面界面使用_后续扩展按标签获取
	beego.Router("/project/getprojects", &controllers.ProjController{}, "*:GetProjects")

	//侧栏懒加载下级
	beego.Router("/project/getprojcate", &controllers.ProjController{}, "*:GetProjCate")
	//添加项目，应该是project/addproj,delproj,updateproj
	beego.Router("/project/addproject", &controllers.ProjController{}, "*:AddProject")
	//添加项目，根据项目模板
	beego.Router("/project/addprojtemplet", &controllers.ProjController{}, "*:AddProjTemplet")

	//修改项目
	beego.Router("/project/updateproject", &controllers.ProjController{}, "*:UpdateProject")
	//删除项目
	beego.Router("/project/deleteproject", &controllers.ProjController{}, "*:DeleteProject")
	//项目时间轴
	beego.Router("/project/:id([0-9]+)/gettimeline", &controllers.ProjController{}, "get:ProjectTimeline")
	beego.Router("/project/:id([0-9]+)/timeline", &controllers.ProjController{}, "get:Timeline")

	//根据项目id进入一个具体项目的侧栏
	beego.Router("/project/:id([0-9]+)", &controllers.ProjController{}, "*:GetProject")
	//进入项目日历
	beego.Router("/project/:id([0-9]+)/getcalendar", &controllers.ProjController{}, "*:GetCalendar")
	//取得日历数据
	beego.Router("/project/:id([0-9]+)/calendar", &controllers.ProjController{}, "*:Calendar")
	//添加日历
	beego.Router("/project/:id([0-9]+)/calendar/addcalendar", &controllers.ProjController{}, "*:AddCalendar")
	//修改
	beego.Router("/project/calendar/updatecalendar", &controllers.ProjController{}, "*:UpdateCalendar")
	//删除
	beego.Router("/project/calendar/deletecalendar", &controllers.ProjController{}, "*:DeleteCalendar")
	//拖曳事件
	beego.Router("/project/calendar/dropcalendar", &controllers.ProjController{}, "*:DropCalendar")
	//resize事件
	beego.Router("/project/calendar/resizecalendar", &controllers.ProjController{}, "*:ResizeCalendar")
	//日历中传照片
	beego.Router("/project/calendar/uploadimage", &controllers.ProjController{}, "*:UploadImage")

	//点击侧栏，根据id返回json数据给导航条
	beego.Router("/project/navbar/:id:string", &controllers.ProjController{}, "*:GetProjNav")

	//根据项目侧栏id显示这个id下的成果界面——作废，用上面GetProject界面
	beego.Router("/project/:id:string/:id:string", &controllers.ProdController{}, "*:GetProjProd")

	//给上面那个页面提供table所用的json数据
	beego.Router("/project/products/:id:string", &controllers.ProdController{}, "*:GetProducts")
	//点击项目名称，显示这个项目下所有成果
	beego.Router("/project/allproducts/:id:string", &controllers.ProjController{}, "*:GetProjProducts")
	//这个对应上面的页面进行表格内数据填充，用的是prodcontrollers中的方法
	beego.Router("/project/products/all/:id:string", &controllers.ProdController{}, "*:GetProjProducts")
	//给上面那个页面提供项目同步ip的json数据
	beego.Router("/project/synchproducts/:id:string", &controllers.ProdController{}, "*:GetsynchProducts")

	//对外提供同步成果接口json数据
	beego.Router("/project/providesynchproducts", &controllers.ProdController{}, "*:ProvidesynchProducts")

	//向项目侧栏id下添加成果_这个没用到，用下面addattachment
	// beego.Router("/project/product/addproduct", &controllers.ProdController{}, "post:AddProduct")
	//编辑成果信息
	beego.Router("/project/product/updateproduct", &controllers.ProdController{}, "post:UpdateProduct")
	//删除成果
	beego.Router("/project/product/deleteproduct", &controllers.ProdController{}, "post:DeleteProduct")
	//取得成果中所有附件中的非pdf附件列表
	beego.Router("/project/product/attachment/:id:string", &controllers.AttachController{}, "*:GetAttachments")
	//取得同步成果中所有附件中的非pdf附件列表
	beego.Router("/project/product/synchattachment", &controllers.AttachController{}, "*:GetsynchAttachments")
	//提供接口：同步成果中所有附件中的非pdf附件列表
	beego.Router("/project/product/providesynchattachment", &controllers.AttachController{}, "*:ProvideAttachments")

	//取得成果中所有附件中——包含pdf，非pdf，文章
	beego.Router("/project/product/allattachments/:id:string", &controllers.AttachController{}, "*:GetAllAttachments")

	//取得成果中所有附件的pdf列表
	beego.Router("/project/product/pdf/:id:string", &controllers.AttachController{}, "*:GetPdfs")
	//取得同步成果中所有pdf列表
	beego.Router("/project/product/synchpdf", &controllers.AttachController{}, "*:GetsynchPdfs")
	//提供同步成果中所有pdf列表
	beego.Router("/project/product/providesynchpdf", &controllers.AttachController{}, "*:ProvidePdfs")
	beego.Router("/pdf", &controllers.AttachController{}, "*:Pdf")

	//根据文章id取得成果中的文章
	beego.Router("/project/product/article/:id:string", &controllers.ArticleController{}, "*:GetArticle")
	//根据成果id取得成果的所有文章列表_注意articles是复数
	beego.Router("/project/product/articles/:id:string", &controllers.ArticleController{}, "*:GetArticles")
	//取得同步成果的所有文章列表_注意articles是复数
	beego.Router("/project/product/syncharticles", &controllers.ArticleController{}, "*:GetsynchArticles")
	//根据成果id取得成果的所有文章列表_注意articles是复数
	beego.Router("/project/product/providesyncharticles", &controllers.ArticleController{}, "*:ProvideArticles")
	//文章进行编辑页面
	beego.Router("/project/product/modifyarticle/:id:string", &controllers.ArticleController{}, "*:ModifyArticle")
	//文章进行编辑
	beego.Router("/project/product/updatearticle", &controllers.ArticleController{}, "*:UpdateArticle")
	//删除文章
	beego.Router("/project/product/deletearticle", &controllers.ArticleController{}, "*:DeleteArticle")

	//查看一个成果
	// beego.Router("/project/product/?:id:string"

	//向成果里添加附件：批量一对一模式
	beego.Router("/project/product/addattachment", &controllers.AttachController{}, "post:AddAttachment")
	//dwg写入服务器
	beego.Router("/project/product/savedwgfile", &controllers.AttachController{}, "post:SaveDwgfile")
	//新建dwg文件
	beego.Router("/project/product/newdwg", &controllers.AttachController{}, "post:NewDwg")

	//向成果里添加附件：多附件模式
	beego.Router("/project/product/addattachment2", &controllers.AttachController{}, "post:AddAttachment2")
	//编辑附件列表：向成果里追加附件：多附件模式
	beego.Router("/project/product/updateattachment", &controllers.AttachController{}, "post:UpdateAttachment")
	//删除附件列表中的附件
	beego.Router("/project/product/deleteattachment", &controllers.AttachController{}, "post:DeleteAttachment")

	//向成果里添加文章
	//向侧栏下添加文章
	beego.Router("/project/product/addarticle", &controllers.ArticleController{}, "post:AddArticle")
	//在某个项目里搜索成果
	// beego.Router("/search", &controllers.SearchController{}, "get:Get")
	// beego.Router("/project/product/search", &controllers.SearchController{}, "get:SearchProjProducts")
	//搜索某个项目——没意义
	beego.Router("/projects/search", &controllers.SearchController{}, "get:SearchProjects")

	//********gante
	//项目进度甘特图
	beego.Router("/projectgant", &controllers.ProjGantController{}, "get:Get")
	//数据填充
	// beego.Router("/projectgant/getprojgants", &controllers.ProjGantController{}, "get:GetProjGants")
	//添加项目进度
	beego.Router("/projectgant/addprojgant", &controllers.ProjGantController{}, "post:AddProjGant")
	//导入项目进度数据
	beego.Router("/projectgant/importprojgant", &controllers.ProjGantController{}, "post:ImportProjGant")

	//修改项目进度
	beego.Router("/projectgant/updateprojgant", &controllers.ProjGantController{}, "post:UpdateProjGant")
	//删除项目进度
	beego.Router("/projectgant/deleteprojgant", &controllers.ProjGantController{}, "post:DeleteProjGant")
	//关闭项目进度
	beego.Router("/projectgant/closeprojgant", &controllers.ProjGantController{}, "post:CloseProjGant")

	//附件下载"/attachment/*", &controllers.AttachController{}
	// beego.InsertFilter("/attachment/*", beego.BeforeRouter, controllers.ImageFilter)
	//根据附件绝对地址下载
	beego.Router("/attachment/*", &controllers.AttachController{}, "get:Attachment")
	//根据附件id号，判断权限下载
	beego.Router("/downloadattachment", &controllers.AttachController{}, "get:DownloadAttachment")

	//上面用attachment.ImageFilter是不行的，必须是package.func
	//首页轮播图片的权限
	beego.Router("/attachment/carousel/*.*", &controllers.AttachController{}, "get:GetCarousel")

	//ue富文本编辑器
	beego.Router("/controller", &controllers.UeditorController{}, "*:ControllerUE")
	//添加文章——froala上传插入的图片
	beego.Router("/uploadimg", &controllers.FroalaController{}, "*:UploadImg")
	//添加wiki——froala上传插入的图片
	beego.Router("/uploadwikiimg", &controllers.FroalaController{}, "*:UploadWikiImg")
	//添加文章——froala上传插入的视频
	beego.Router("/uploadvideo", &controllers.FroalaController{}, "*:UploadVideo")

	//添加日历
	beego.Router("/index/carcalendar/addcalendar", &controllers.IndexController{}, "*:AddCarCalendar")
	//查询
	beego.Router("/index/carcalendar", &controllers.IndexController{}, "*:CarCalendar")
	//修改
	beego.Router("/index/carcalendar/updatecalendar", &controllers.IndexController{}, "*:UpdateCarCalendar")
	//删除
	beego.Router("/index/carcalendar/deletecalendar", &controllers.IndexController{}, "*:DeleteCarCalendar")
	//拖曳事件
	beego.Router("/index/carcalendar/dropcalendar", &controllers.IndexController{}, "*:DropCarCalendar")
	//resize事件
	beego.Router("/index/carcalendar/resizecalendar", &controllers.IndexController{}, "*:ResizeCarCalendar")

	//添加日历
	beego.Router("/index/meetingroomcalendar/addcalendar", &controllers.IndexController{}, "*:AddMeetCalendar")
	//查询
	beego.Router("/index/meetingroomcalendar", &controllers.IndexController{}, "*:MeetCalendar")
	//修改
	beego.Router("/index/meetingroomcalendar/updatecalendar", &controllers.IndexController{}, "*:UpdateMeetCalendar")
	//删除
	beego.Router("/index/meetingroomcalendar/deletecalendar", &controllers.IndexController{}, "*:DeleteMeetCalendar")
	//拖曳事件
	beego.Router("/index/meetingroomcalendar/dropcalendar", &controllers.IndexController{}, "*:DropMeetCalendar")
	//resize事件
	beego.Router("/index/meetingroomcalendar/resizecalendar", &controllers.IndexController{}, "*:ResizeMeetCalendar")

	//wiki页面
	beego.Router("/wiki", &controllers.WikiController{}) //get
	//发表文章界面
	beego.Router("/wiki/add", &controllers.WikiController{}, "get:Add")
	//发表文章提交
	beego.Router("/wiki/addwiki", &controllers.WikiController{}, "post:AddWiki")

	//查看一个文章
	beego.Router("/wiki/view/", &controllers.WikiController{}, "get:View")
	//删除wiki
	beego.Router("/wiki/delete", &controllers.WikiController{}, "post:Delete")
	//修改一个文章
	// beego.Router("/wiki/modify", &controllers.WikiController{}, "get:Modify")
	beego.AutoRouter(&controllers.WikiController{})
	//添加删除wiki的评论
	beego.Router("/reply/addwiki", &controllers.ReplyController{}, "post:AddWiki")
	beego.Router("/reply/deletewiki", &controllers.ReplyController{}, "get:DeleteWiki")

	// beego.InsertFilter("/attachment/standard/", beego.BeforeRouter, FilterUser)
	// beego.SetStaticPath("/attachment/standard/", "attachment/standard/")
	//这个有哦何用？
	beego.SetStaticPath("/attachment/wiki", "attachment/wiki")
	beego.SetStaticPath("/swagger", "swagger")
	// *全匹配方式 //匹配 /download/ceshi/file/api.json :splat=file/api.json
	beego.Router("/searchwiki", &controllers.SearchController{}, "get:SearchWiki")

	//规范管理
	beego.Router("/standard", &controllers.StandardController{}, "get:Index")
	beego.Router("/standard/search", &controllers.StandardController{}, "get:Search")
	beego.Router("/standard/importexcel", &controllers.StandardController{}, "post:ImportExcel")
	beego.Router("/standard/standard_one_addbaidu", &controllers.StandardController{}, "post:Standard_one_addbaidu")
	beego.Router("/standard/importlibrary", &controllers.StandardController{}, "post:ImportLibrary")
	//显示规范所有
	beego.Router("/standard/getstandard", &controllers.StandardController{}, "get:GetStandard")
	//修改规范库
	beego.Router("/standard/updatestandard", &controllers.StandardController{}, "post:UpdateStandard")
	//删除规范库
	beego.Router("/standard/deletestandard", &controllers.StandardController{}, "post:DeleteStandard")

	//显示有效库所有
	beego.Router("/standard/valid", &controllers.StandardController{}, "get:Valid")
	//删除有效库中选中
	beego.Router("/standard/deletevalid", &controllers.StandardController{}, "post:DeleteValid")

	//对标
	beego.Router("/legislation", &controllers.LegislationController{}, "*:Index")
	beego.Router("/legislation/checklist", &controllers.LegislationController{}, "*:Checklist")
	//上传文档，提供解析和替换（增加）标准号
	beego.Router("/legislation/fileinput", &controllers.LegislationController{}, "get:FileInput")

	//设代日志列表页面
	beego.Router("/diary", &controllers.DiaryController{}, "get:Get") //get
	//具体一个日志页面
	beego.Router("/getwxdiary2/:id", &controllers.DiaryController{}, "get:GetWxDiary2")
	//转换为word
	// beego.Router("/htmltodoc", &controllers.DiaryController{}, "get:HtmlToDoc")

	//设代财务登记列表页面
	beego.Router("/finance", &controllers.FinanceController{}, "get:Get") //get
	//具体一个财务登记页面
	beego.Router("/getwxfinance2/:id", &controllers.FinanceController{}, "get:GetWxFinance2")

	//微信小程序
	//小程序发表文章提交
	// beego.Router("/wx/addwxarticle", &controllers.ArticleController{}, "post:AddWxArticle")
	//小程序上传图片，返回地址
	// beego.Router("/wx/uploadwximg", &controllers.FroalaController{}, "*:UploadWxImg")
	//小程序获得文章列表，分页
	// beego.Router("/wx/getwxarticles", &controllers.ArticleController{}, "*:GetWxArticles")
	//小程序根据文章id返回文章数据
	// beego.Router("/wx/getwxarticle/:id:string", &controllers.ArticleController{}, "*:GetWxArticle")
	//微信登录
	// beego.Router("/wx/wxlogin", &controllers.LoginController{}, "*:WxLogin")

	beego.Router("/mindoc", &controllers.HomeController{}, "*:Index")

	beego.Router("/mindoclogin", &controllers.AccountController{}, "*:Login")
	beego.Router("/logout", &controllers.AccountController{}, "*:Logout")
	beego.Router("/register", &controllers.AccountController{}, "*:Register")
	beego.Router("/find_password", &controllers.AccountController{}, "*:FindPassword")
	beego.Router("/valid_email", &controllers.AccountController{}, "post:ValidEmail")
	beego.Router("/captcha", &controllers.AccountController{}, "*:Captcha")

	beego.Router("/manager", &controllers.ManagerController{}, "*:Index")
	beego.Router("/manager/users", &controllers.ManagerController{}, "*:Users")
	beego.Router("/manager/users/edit/:id", &controllers.ManagerController{}, "*:EditMember")
	beego.Router("/manager/member/create", &controllers.ManagerController{}, "post:CreateMember")
	beego.Router("/manager/member/delete", &controllers.ManagerController{}, "post:DeleteMember")
	beego.Router("/manager/member/update-member-status", &controllers.ManagerController{}, "post:UpdateMemberStatus")
	beego.Router("/manager/member/change-member-role", &controllers.ManagerController{}, "post:ChangeMemberRole")
	beego.Router("/manager/books", &controllers.ManagerController{}, "*:Books")
	beego.Router("/manager/books/edit/:key", &controllers.ManagerController{}, "*:EditBook")
	beego.Router("/manager/books/delete", &controllers.ManagerController{}, "*:DeleteBook")

	beego.Router("/manager/comments", &controllers.ManagerController{}, "*:Comments")
	beego.Router("/manager/setting", &controllers.ManagerController{}, "*:Setting")
	beego.Router("/manager/books/token", &controllers.ManagerController{}, "post:CreateToken")
	beego.Router("/manager/books/transfer", &controllers.ManagerController{}, "post:Transfer")
	beego.Router("/manager/books/open", &controllers.ManagerController{}, "post:PrivatelyOwned")

	beego.Router("/manager/attach/list", &controllers.ManagerController{}, "*:AttachList")
	beego.Router("/manager/attach/detailed/:id", &controllers.ManagerController{}, "*:AttachDetailed")
	beego.Router("/manager/attach/delete", &controllers.ManagerController{}, "post:AttachDelete")
	beego.Router("/manager/label/list", &controllers.ManagerController{}, "get:LabelList")
	beego.Router("/manager/label/delete/:id", &controllers.ManagerController{}, "post:LabelDelete")

	//beego.Router("/manager/config",  &controllers.ManagerController{}, "*:Config")

	beego.Router("/manager/team", &controllers.ManagerController{}, "*:Team")
	beego.Router("/manager/team/create", &controllers.ManagerController{}, "POST:TeamCreate")
	beego.Router("/manager/team/edit", &controllers.ManagerController{}, "POST:TeamEdit")
	beego.Router("/manager/team/delete", &controllers.ManagerController{}, "POST:TeamDelete")

	beego.Router("/manager/team/member/list/:id", &controllers.ManagerController{}, "*:TeamMemberList")
	beego.Router("/manager/team/member/add", &controllers.ManagerController{}, "POST:TeamMemberAdd")
	beego.Router("/manager/team/member/delete", &controllers.ManagerController{}, "POST:TeamMemberDelete")
	beego.Router("/manager/team/member/change_role", &controllers.ManagerController{}, "POST:TeamChangeMemberRole")
	beego.Router("/manager/team/member/search", &controllers.ManagerController{}, "*:TeamSearchMember")

	beego.Router("/manager/team/book/list/:id", &controllers.ManagerController{}, "*:TeamBookList")
	beego.Router("/manager/team/book/add", &controllers.ManagerController{}, "POST:TeamBookAdd")
	beego.Router("/manager/team/book/delete", &controllers.ManagerController{}, "POST:TeamBookDelete")
	beego.Router("/manager/team/book/search", &controllers.ManagerController{}, "*:TeamSearchBook")

	beego.Router("/manager/itemsets", &controllers.ManagerController{}, "*:Itemsets")
	beego.Router("/manager/itemsets/edit", &controllers.ManagerController{}, "post:ItemsetsEdit")
	beego.Router("/manager/itemsets/delete", &controllers.ManagerController{}, "post:ItemsetsDelete")

	beego.Router("/setting", &controllers.SettingController{}, "*:Index")
	beego.Router("/setting/password", &controllers.SettingController{}, "*:Password")
	beego.Router("/setting/upload", &controllers.SettingController{}, "*:Upload")

	beego.Router("/book", &controllers.BookController{}, "*:Index")
	beego.Router("/book/:key/dashboard", &controllers.BookController{}, "*:Dashboard")
	beego.Router("/book/:key/setting", &controllers.BookController{}, "*:Setting")
	beego.Router("/book/:key/users", &controllers.BookController{}, "*:Users")
	beego.Router("/book/:key/release", &controllers.BookController{}, "post:Release")
	beego.Router("/book/:key/sort", &controllers.BookController{}, "post:SaveSort")
	beego.Router("/book/:key/teams", &controllers.BookController{}, "*:Team")

	beego.Router("/book/create", &controllers.BookController{}, "*:Create")
	beego.Router("/book/itemsets/search", &controllers.BookController{}, "*:ItemsetsSearch")

	beego.Router("/book/users/create", &controllers.BookMemberController{}, "post:AddMember")
	beego.Router("/book/users/change", &controllers.BookMemberController{}, "post:ChangeRole")
	beego.Router("/book/users/delete", &controllers.BookMemberController{}, "post:RemoveMember")
	beego.Router("/book/users/import", &controllers.BookController{}, "post:Import")
	beego.Router("/book/users/copy", &controllers.BookController{}, "post:Copy")

	beego.Router("/book/setting/save", &controllers.BookController{}, "post:SaveBook")
	beego.Router("/book/setting/open", &controllers.BookController{}, "post:PrivatelyOwned")
	beego.Router("/book/setting/transfer", &controllers.BookController{}, "post:Transfer")
	beego.Router("/book/setting/upload", &controllers.BookController{}, "post:UploadCover")
	beego.Router("/book/setting/delete", &controllers.BookController{}, "post:Delete")

	beego.Router("/book/team/add", &controllers.BookController{}, "POST:TeamAdd")
	beego.Router("/book/team/delete", &controllers.BookController{}, "POST:TeamDelete")
	beego.Router("/book/team/search", &controllers.BookController{}, "*:TeamSearch")

	//管理文章的路由
	beego.Router("/manage/blogs", &controllers.BlogController{}, "*:ManageList")
	beego.Router("/manage/blogs/setting/?:id", &controllers.BlogController{}, "*:ManageSetting")
	beego.Router("/manage/blogs/edit/?:id", &controllers.BlogController{}, "*:ManageEdit")
	beego.Router("/manage/blogs/delete", &controllers.BlogController{}, "post:ManageDelete")
	beego.Router("/manage/blogs/upload", &controllers.BlogController{}, "post:Upload")
	beego.Router("/manage/blogs/attach/:id", &controllers.BlogController{}, "post:RemoveAttachment")

	//读文章的路由
	beego.Router("/blogs", &controllers.BlogController{}, "*:List")
	beego.Router("/blog-attach/:id:int/:attach_id:int", &controllers.BlogController{}, "get:Download")
	beego.Router("/blog-:id([0-9]+).html", &controllers.BlogController{}, "*:Index")

	//模板相关接口
	beego.Router("/api/template/get", &controllers.TemplateController{}, "get:Get")
	beego.Router("/api/template/list", &controllers.TemplateController{}, "post:List")
	beego.Router("/api/template/add", &controllers.TemplateController{}, "post:Add")
	beego.Router("/api/template/remove", &controllers.TemplateController{}, "post:Delete")

	beego.Router("/api/attach/remove/", &controllers.DocumentController{}, "post:RemoveAttachment")
	beego.Router("/api/:key/edit/?:id", &controllers.DocumentController{}, "*:Edit")
	beego.Router("/api/upload", &controllers.DocumentController{}, "post:Upload")
	beego.Router("/api/:key/create", &controllers.DocumentController{}, "post:Create")
	beego.Router("/api/:key/delete", &controllers.DocumentController{}, "post:Delete")
	beego.Router("/api/:key/content/?:id", &controllers.DocumentController{}, "*:Content")
	beego.Router("/api/:key/compare/:id", &controllers.DocumentController{}, "*:Compare")
	beego.Router("/api/search/user/:key", &controllers.MindocSearchController{}, "*:User")

	beego.Router("/history/get", &controllers.DocumentController{}, "get:History")
	beego.Router("/history/delete", &controllers.DocumentController{}, "*:DeleteHistory")
	beego.Router("/history/restore", &controllers.DocumentController{}, "*:RestoreHistory")

	beego.Router("/docs/:key", &controllers.DocumentController{}, "*:Index")
	beego.Router("/docs/:key/:id", &controllers.DocumentController{}, "*:Read")
	beego.Router("/docs/:key/search", &controllers.DocumentController{}, "post:Search")
	beego.Router("/export/:key", &controllers.DocumentController{}, "*:Export")
	beego.Router("/qrcode/:key.png", &controllers.DocumentController{}, "get:QrCode")

	beego.Router("/attach_files/:key/:attach_id", &controllers.DocumentController{}, "get:DownloadAttachment")

	beego.Router("/comment/create", &controllers.CommentController{}, "post:Create")
	beego.Router("/comment/lists", &controllers.CommentController{}, "get:Lists")
	beego.Router("/comment/index", &controllers.CommentController{}, "*:Index")

	beego.Router("/mindocsearch", &controllers.MindocSearchController{}, "get:Index")

	beego.Router("/tag/:key", &controllers.LabelController{}, "get:Index")
	beego.Router("/tags", &controllers.LabelController{}, "get:List")

	beego.Router("/items", &controllers.ItemsetsController{}, "get:Index")
	beego.Router("/items/:key", &controllers.ItemsetsController{}, "get:List")

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
