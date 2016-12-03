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

	//项目列表界面
	beego.Router("/project", &controllers.ProjController{}, "*:Get")
	//table获取所有项目数据给上面界面使用_后续扩展按标签获取
	beego.Router("/project/getprojects", &controllers.ProjController{}, "*:GetProjects")
	//添加项目，应该是project/addproj,delproj,updateproj
	beego.Router("/project/addproject", &controllers.ProjController{}, "*:AddProject")

	//根据项目id进入一个具体项目的侧栏
	beego.Router("/project/:id:string", &controllers.ProjController{}, "*:GetProject")
	//根据项目侧栏id显示这个id下的成果界面
	beego.Router("/project/:id:string/:id:string", &controllers.ProjController{}, "*:GetProjProd")
	//给上面那个页面提供table所用的json数据
	beego.Router("/project/products/:id:string", &controllers.ProjController{}, "*:GetProducts")
	//向项目侧栏id下添加成果project/addprod
	beego.Router("/project/addproduct", &controllers.ProjController{}, "post:AddProduct")
	//查看一个成果
	// beego.Router("/project/product/?:id:string"
	beego.Router("/controller", &controllers.UeditorController{}, "*:ControllerUE")
}
