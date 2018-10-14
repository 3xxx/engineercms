package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"],
		beego.ControllerComments{
			Method: "Category",
			Router: `/category/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"],
		beego.ControllerComments{
			Method: "AddCategory",
			Router: `/category/addcategory`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminController"],
		beego.ControllerComments{
			Method: "CategoryTitle",
			Router: `/categorytitle`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminLogController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminLogController"],
		beego.ControllerComments{
			Method: "ErrLog",
			Router: `/errlog`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminLogController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:AdminLogController"],
		beego.ControllerComments{
			Method: "InfoLog",
			Router: `/infolog`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "AddWxArticle",
			Router: `/addwxarticle`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "AddWxArticles",
			Router: `/addwxarticles/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetWxArticle",
			Router: `/getwxarticle/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetWxArticles",
			Router: `/getwxarticles`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ArticleController"],
		beego.ControllerComments{
			Method: "GetWxArticless",
			Router: `/getwxarticless/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"],
		beego.ControllerComments{
			Method: "UploadWxImg",
			Router: `/uploadwximg`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:FroalaController"],
		beego.ControllerComments{
			Method: "UploadWxImgs",
			Router: `/uploadwximgs/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LoginController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:LoginController"],
		beego.ControllerComments{
			Method: "WxLogin",
			Router: `/wxlogin`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
