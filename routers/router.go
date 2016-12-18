package routers

import (
	"engineercms/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	//后台
	beego.Router("/admin", &controllers.AdminController{})
	// beego.Router("/admincategory", &controllers.AdminController{}, "*:GetAdminCategory")
	//显示对应侧栏id的右侧界面
	beego.Router("/admin/:id:string", &controllers.AdminController{}, "*:Admin")

	//根据数字id查询类别或目录分级表
	beego.Router("/admin/category/?:id:string", &controllers.AdminController{}, "*:Category")
	//根据名字查询目录分级表
	beego.Router("/admin/categorytitle", &controllers.AdminController{}, "*:CategoryTitle")
	//添加目录类别
	beego.Router("/admin/category/addcategory", &controllers.AdminController{}, "*:AddCategory")
	//修改目录类别
	beego.Router("/admin/category/updatecategory", &controllers.AdminController{}, "*:UpdateCategory")

	//添加IP地址段
	beego.Router("/admin/ipsegment/addipsegment", &controllers.AdminController{}, "*:AddIpsegment")
	//修改IP地址段
	beego.Router("/admin/ipsegment/updateipsegment", &controllers.AdminController{}, "*:UpdateIpsegment")
	//查询所有ip地址段
	beego.Router("/admin/ipsegment", &controllers.AdminController{}, "*:Ipsegment")

	//项目列表界面
	beego.Router("/project", &controllers.ProjController{}, "*:Get")
	//table获取所有项目数据给上面界面使用_后续扩展按标签获取
	beego.Router("/project/getprojects", &controllers.ProjController{}, "*:GetProjects")
	//添加项目，应该是project/addproj,delproj,updateproj
	beego.Router("/project/addproject", &controllers.ProjController{}, "*:AddProject")

	//根据项目id进入一个具体项目的侧栏
	beego.Router("/project/:id:string", &controllers.ProjController{}, "*:GetProject")
	//点击侧栏，根据id返回json数据给导航条
	beego.Router("/project/navbar/:id:string", &controllers.ProjController{}, "*:GetProjNav")

	//根据项目侧栏id显示这个id下的成果界面——作废，用上面GetProject界面
	beego.Router("/project/:id:string/:id:string", &controllers.ProdController{}, "*:GetProjProd")

	//给上面那个页面提供table所用的json数据
	beego.Router("/project/products/:id:string", &controllers.ProdController{}, "*:GetProducts")

	//向项目侧栏id下添加成果project/addprod
	beego.Router("/project/addproduct", &controllers.ProdController{}, "post:AddProduct")
	//取得成果中所有附件的pdf列表
	beego.Router("/project/product/attachment/:id:string", &controllers.AttachController{}, "*:GetAttachments")
	//取得成果中所有附件的pdf列表
	beego.Router("/project/product/pdf/:id:string", &controllers.AttachController{}, "*:GetPdfs")
	//根据文章id取得成果中的文章
	beego.Router("/project/product/article/:id:string", &controllers.ArticleController{}, "*:GetArticle")
	//根据成果id取得成果的所有文章列表
	beego.Router("/project/product/articles/:id:string", &controllers.ArticleController{}, "*:GetArticles")

	//查看一个成果
	// beego.Router("/project/product/?:id:string"

	//向成果里添加附件：批量一对一模式
	beego.Router("/project/product/addattachment", &controllers.AttachController{}, "post:AddAttachment")
	//向成果里添加附件：多附件模式
	beego.Router("/project/product/addattachment2", &controllers.AttachController{}, "post:AddAttachment2")
	//向成果里添加文章
	beego.Router("/project/product/addarticle", &controllers.ArticleController{}, "post:AddArticle")

	//附件下载"/attachment/*", &controllers.AttachController{}
	// beego.InsertFilter("/attachment/*", beego.BeforeRouter, controllers.ImageFilter)
	beego.Router("/attachment/*", &controllers.AttachController{}, "get:DownloadAttachment")
	//上面用attachment.ImageFilter是不行的，必须是package.func
	//ue富文本编辑器
	beego.Router("/controller", &controllers.UeditorController{}, "*:ControllerUE")

}
