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
			Router: `/wxlogin/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MainController"],
		beego.ControllerComments{
			Method: "WxPdf",
			Router: `/wxpdf/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MainController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:MainController"],
		beego.ControllerComments{
			Method: "WxStandardPdf",
			Router: `/wxstandardpdf/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ReplyController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "AddWxLike",
			Router: `/addwxlike/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ReplyController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "AddWxRelease",
			Router: `/addwxrelease/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ReplyController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:ReplyController"],
		beego.ControllerComments{
			Method: "DeleteWxRelease",
			Router: `/deletewxrelease/:id`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:SearchController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:SearchController"],
		beego.ControllerComments{
			Method: "SearchWxDrawings",
			Router: `/searchwxdrawings`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"] = append(beego.GlobalControllerRouter["github.com/3xxx/engineercms/controllers:StandardController"],
		beego.ControllerComments{
			Method: "SearchWxStandards",
			Router: `/searchwxstandards`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
