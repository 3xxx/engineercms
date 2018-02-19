// @APIVersion 1.0.0
// @Title mobile API
// @Description mobile has every tool to get any job done, so codename for the new mobile APIs.
// @Contact astaxie@gmail.com
package routers

import (
	"engineercms/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/test", &controllers.MainController{}, "*:Test")
	beego.Router("/url-to-callback", &controllers.OnlyController{}, "*:UrltoCallback")
	// beego.Router("/onlyoffice/post", &controllers.OnlyController{}, "post:PostOnlyoffice")
	beego.Router("/onlyoffice", &controllers.OnlyController{}, "get:Get")
	//table获取所有数据给上面界面使用
	beego.Router("/onlyoffice/data", &controllers.OnlyController{}, "*:GetData")
	//添加一个文档
	beego.Router("/onlyoffice/addattachment", &controllers.OnlyController{}, "post:AddOnlyAttachment")
	//在onlyoffice中打开文档协作
	beego.Router("/onlyoffice/:id:string", &controllers.OnlyController{}, "*:OnlyOffice")
	//删除
	beego.Router("/onlyoffice/deletedoc", &controllers.OnlyController{}, "*:DeleteDoc")
	//修改
	beego.Router("/onlyoffice/updatedoc", &controllers.OnlyController{}, "*:UpdateDoc")

	//下载doc
	beego.Router("/attachment/onlyoffice/*", &controllers.OnlyController{}, "get:DownloadDoc")

	beego.Router("/role/test", &controllers.RoleController{}, "*:Test")
	beego.Router("/1/slide", &controllers.MainController{}, "*:Slide")
	beego.Router("/postdata", &controllers.MainController{}, "*:Postdata")
	//文档
	beego.Router("/doc/ecms", &controllers.MainController{}, "get:Getecmsdoc")
	beego.Router("/doc/meritms", &controllers.MainController{}, "get:Getmeritmsdoc")
	beego.Router("/doc/hydrows", &controllers.MainController{}, "get:Gethydrowsdoc")

	//api接口
	beego.Router("/api/ecms", &controllers.MainController{}, "get:Getecmsapi")
	beego.Router("/api/meritms", &controllers.MainController{}, "get:Getmeritmsapi")

	beego.Router("/", &controllers.MainController{}, "get:Get")

	beego.Router("/pdf", &controllers.MainController{}, "*:Pdf")
	//显示首页
	beego.Router("/index", &controllers.IndexController{}, "*:GetIndex")
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
	beego.Router("/index/searchproduct", &controllers.SearchController{}, "*:SearchProduct")

	//后台
	beego.Router("/admin", &controllers.AdminController{})
	// beego.Router("/admincategory", &controllers.AdminController{}, "*:GetAdminCategory")
	//显示对应侧栏id的右侧界面
	beego.Router("/admin/:id:string", &controllers.AdminController{}, "*:Admin")

	//批量添加首页轮播图片
	beego.Router("/admin/base/addcarousel", &controllers.AdminController{}, "*:AddCarousel")
	//获取首页轮播图片填充表格
	beego.Router("/admin/base/carousel", &controllers.AdminController{}, "*:Carousel")

	//根据数字id查询类别或目录分级表
	beego.Router("/admin/category/?:id:string", &controllers.AdminController{}, "*:Category")
	//根据名字查询目录分级表
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
	//添加角色权限
	beego.Router("/admin/role/permission", &controllers.RoleController{}, "post:RolePermission")
	//查询角色权限
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

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/loginerr", &controllers.LoginController{}, "get:Loginerr")
	beego.Router("/roleerr", &controllers.UserController{}, "*:Roleerr") //显示权限不够

	//项目列表界面
	beego.Router("/project", &controllers.ProjController{}, "*:Get")
	//table获取所有项目数据给上面界面使用_后续扩展按标签获取
	beego.Router("/project/getprojects", &controllers.ProjController{}, "*:GetProjects")
	//添加项目，应该是project/addproj,delproj,updateproj
	beego.Router("/project/addproject", &controllers.ProjController{}, "*:AddProject")
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

	//根据项目侧栏id显示这个id下的成果界面—
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
	beego.Router("/search", &controllers.SearchController{}, "get:Get")
	beego.Router("/project/product/search", &controllers.SearchController{}, "get:SearchProjProducts")
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

	beego.Router("/attachment/*", &controllers.AttachController{}, "get:DownloadAttachment")
	//根据权限下载附件
	beego.Router("/attachment", &controllers.AttachController{}, "get:Attachment")

	//上面用attachment.ImageFilter是不行的，必须是package.func
	//首页轮播图片的权限
	beego.Router("/attachment/carousel/*.*", &controllers.AttachController{}, "get:GetCarousel")

	//ue富文本编辑器
	beego.Router("/controller", &controllers.UeditorController{}, "*:ControllerUE")

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
	beego.SetStaticPath("/attachment/wiki", "attachment/wiki")
	// *全匹配方式 //匹配 /download/ceshi/file/api.json :splat=file/api.json
	beego.Router("/searchwiki", &controllers.SearchController{}, "get:SearchWiki")

}
