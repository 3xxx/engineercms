package controllers

import (
	"github.com/3xxx/engineercms/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	// beego "github.com/beego/beego/v2/adapter"
	"io"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type SearchController struct {
	web.Controller
}

// 搜索项目
func (c *SearchController) SearchProject() { //search用的是get方法
	key := c.GetString("keyword")
	if key != "" {
		searchs, err := models.SearchProject(key)
		if err != nil {
			logs.Error(err.Error)
		} else {
			c.Data["json"] = searchs
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = "关键字为空！"
		c.ServeJSON()
	}
}

// @Title get a project's products search list
// @Description get a project's products search by page
// @Param keyword query string false "The keyword of products"
// @Param productid query string false "The id of project"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 drawings not found
// @router /searchprojectproduct [get]
// 显示特定项目的搜索页面
func (c *SearchController) SearchProjectProduct() { //search用的是get方法
	c.Data["IsProject"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["IsProjects"] = true
	// beego.Info(c.Ctx.Input.IP())
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid

	pid := c.GetString("productid")
	key := c.GetString("keyword")
	if pid != "" {
		c.Data["IsProduct"] = true
	} else {
		c.Data["IsProject"] = true
	}
	c.Data["Pid"] = pid
	c.Data["Key"] = key
	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		c.TplName = "mobile/msearchs.tpl"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "searchs.tpl"
	}
}

// @Title get a project's products search list
// @Description get a project's products search by page
// @Param keyword query string true "The keyword of products"
// @Param searchText query string false "The searchText of products"
// @Param pageNo query string false "The pageNo of drawings list"
// @Param limit query string false "The limit of products list"
// @Param productid query string true "The id of project"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 drawings not found
// @router /searchprojectproductdata [get]
// 分页
// 在某个项目里搜索成果：全文搜索，article全文，编号，名称，关键字，作者……
func (c *SearchController) SearchProjectProductData() {
	// limit := "15"
	limit := c.GetString("limit")
	// beego.Info(limit)
	limit1, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("pageNo")
	// beego.Info(page)
	page1, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	var offset int64
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	pid := c.GetString("productid")
	var pidNum int64
	// var err error
	if pid != "" {
		pidNum, err = strconv.ParseInt(pid, 10, 64)
		if err != nil {
			logs.Error(err)
		}
	}
	key := c.GetString("keyword")
	searchText := c.GetString("searchText")
	// if searchText != "" {
	count, products, err := models.SearchProjProduct(pidNum, limit1, offset, key, searchText) //这里要将侧栏所有id进行循环
	if err != nil {
		logs.Error(err.Error)
	} else {
		link := make([]ProductLink, 0)
		Attachslice := make([]AttachmentLink, 0)
		Pdfslice := make([]PdfLink, 0)
		Articleslice := make([]ArticleContent, 0)
		for _, w := range products {
			Url, _, err := GetUrlPath(w.ProjectId)
			if err != nil {
				logs.Error(err)
			}
			//取到每个成果的附件（模态框打开）；pdf、文章——新窗口打开
			//循环成果
			//每个成果取到所有附件
			//一个附件则直接打开/下载；2个以上则打开模态框
			Attachments, err := models.GetAttachments(w.Id)
			if err != nil {
				logs.Error(err)
			}
			//对成果进行循环
			//赋予url
			//如果是一个成果，直接给url;如果大于1个，则是数组:这个在前端实现
			// http.ServeFile(ctx.ResponseWriter, ctx.Request, filePath)
			linkarr := make([]ProductLink, 1)
			linkarr[0].Id = w.Id
			linkarr[0].Code = w.Code
			linkarr[0].Title = w.Title
			linkarr[0].Label = w.Label
			linkarr[0].Uid = w.Uid
			linkarr[0].Principal = w.Principal
			linkarr[0].ProjectId = w.ProjectId
			// linkarr[0].Content = w.Content
			linkarr[0].Created = w.Created
			linkarr[0].Updated = w.Updated
			// linkarr[0].Views = w.Views
			for _, v := range Attachments {
				// fileext := path.Ext(v.FileName)
				if path.Ext(v.FileName) != ".pdf" && path.Ext(v.FileName) != ".PDF" {
					attacharr := make([]AttachmentLink, 1)
					attacharr[0].Id = v.Id
					attacharr[0].Title = v.FileName
					attacharr[0].Link = Url
					Attachslice = append(Attachslice, attacharr...)
				} else if path.Ext(v.FileName) == ".pdf" || path.Ext(v.FileName) == ".PDF" {
					pdfarr := make([]PdfLink, 1)
					pdfarr[0].Id = v.Id
					pdfarr[0].Title = v.FileName
					pdfarr[0].Link = Url
					Pdfslice = append(Pdfslice, pdfarr...)
				}
			}
			linkarr[0].Pdflink = Pdfslice
			linkarr[0].Attachmentlink = Attachslice
			Attachslice = make([]AttachmentLink, 0) //再把slice置0
			Pdfslice = make([]PdfLink, 0)           //再把slice置0
			// link = append(link, linkarr...)
			//取得文章
			Articles, err := models.GetArticles(w.Id)
			if err != nil {
				logs.Error(err)
			}
			for _, x := range Articles {
				articlearr := make([]ArticleContent, 1)
				articlearr[0].Id = x.Id
				articlearr[0].Content = x.Content
				articlearr[0].Link = "/project/product/article"
				Articleslice = append(Articleslice, articlearr...)
			}
			linkarr[0].Articlecontent = Articleslice
			Articleslice = make([]ArticleContent, 0)
			link = append(link, linkarr...)
		}

		table := prodTableserver{link, page1, count}

		c.Data["json"] = table
		c.ServeJSON()
	}
}

// @Title get allprojects' products search list
// @Description get allprojects' products search by page
// @Param keyword query string false "The keyword of products"
// @Param productid query string false "The id of project"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 drawings not found
// @router /searchproduct [get]
// 显示所有项目的搜索结果页面，分页
func (c *SearchController) SearchProduct() {
	c.Data["IsProject"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["IsProjects"] = true
	// beego.Info(c.Ctx.Input.IP())
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid

	pid := c.GetString("productid")
	key := c.GetString("keyword")
	if pid != "" {
		c.Data["IsProduct"] = true
	} else {
		c.Data["IsProject"] = true
	}
	c.Data["Pid"] = pid
	c.Data["Key"] = key
	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		c.TplName = "mobile/msearchs_allproducts.tpl"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "searchs_allproducts.tpl"
	}
}

// @Title get all projects' products search list
// @Description get all projects' products search by page
// @Param keyword query string true "The keyword of products"
// @Param searchText query string false "The searchText of products"
// @Param limit query string false "The limit of products list"
// @Param pageNo query string false "The page of products list"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 drawings not found
// @router /searchproductdata [get]
// 搜索所有项目的成果数据
func (c *SearchController) SearchProductData() {
	var offset, limit1, page1 int64
	var err error
	limit := c.GetString("limit")
	if limit == "" {
		limit1 = 15
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			logs.Error(err)
		}
	}
	page := c.GetString("pageNo")
	if page == "" {
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			logs.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	key := c.GetString("keyword")
	searchText := c.GetString("searchText")
	if key != "" {
		// products, err := models.SearchProduct(key)
		products, err := models.SearchProductPage(limit1, offset, key, searchText)
		if err != nil {
			logs.Error(err.Error)
		}

		link := make([]ProductLink, 0)
		Attachslice := make([]AttachmentLink, 0)
		Pdfslice := make([]PdfLink, 0)
		Articleslice := make([]ArticleContent, 0)
		for _, w := range products {
			// Url, _, err := GetUrlPath(w.ProjectId)
			// if err != nil {
			// 	logs.Error(err)
			// }
			//取到每个成果的附件（模态框打开）；pdf、文章——新窗口打开
			//循环成果
			//每个成果取到所有附件
			//一个附件则直接打开/下载；2个以上则打开模态框
			Attachments, err := models.GetAttachments(w.Id)
			if err != nil {
				logs.Error(err)
			}
			//对成果进行循环
			//赋予url
			//如果是一个成果，直接给url;如果大于1个，则是数组:这个在前端实现
			// http.ServeFile(ctx.ResponseWriter, ctx.Request, filePath)
			linkarr := make([]ProductLink, 1)
			linkarr[0].Id = w.Id
			linkarr[0].Code = w.Code
			linkarr[0].Title = w.Title
			linkarr[0].Label = w.Label
			linkarr[0].Uid = w.Uid
			linkarr[0].Principal = w.Principal
			linkarr[0].ProjectId = w.ProjectId
			// linkarr[0].Content = w.Content
			linkarr[0].Created = w.Created
			linkarr[0].Updated = w.Updated
			linkarr[0].ProjectId = w.ProjectId
			linkarr[0].TopProjectId = w.TopProjectId
			// linkarr[0].Views = w.Views
			for _, v := range Attachments {
				// fileext := path.Ext(v.FileName)
				if path.Ext(v.FileName) != ".pdf" && path.Ext(v.FileName) != ".PDF" {
					attacharr := make([]AttachmentLink, 1)
					attacharr[0].Id = v.Id
					attacharr[0].Title = v.FileName
					// attacharr[0].Link = Url
					Attachslice = append(Attachslice, attacharr...)
				} else if path.Ext(v.FileName) == ".pdf" || path.Ext(v.FileName) == ".PDF" {
					pdfarr := make([]PdfLink, 1)
					pdfarr[0].Id = v.Id
					pdfarr[0].Title = v.FileName
					// pdfarr[0].Link = Url
					Pdfslice = append(Pdfslice, pdfarr...)
				}
			}
			linkarr[0].Pdflink = Pdfslice
			linkarr[0].Attachmentlink = Attachslice
			Attachslice = make([]AttachmentLink, 0) //再把slice置0
			Pdfslice = make([]PdfLink, 0)           //再把slice置0
			// link = append(link, linkarr...)
			//取得文章
			Articles, err := models.GetArticles(w.Id)
			if err != nil {
				logs.Error(err)
			}
			for _, x := range Articles {
				articlearr := make([]ArticleContent, 1)
				articlearr[0].Id = x.Id
				articlearr[0].Content = x.Content
				articlearr[0].Link = "/project/product/article"
				Articleslice = append(Articleslice, articlearr...)
			}
			linkarr[0].Articlecontent = Articleslice
			Articleslice = make([]ArticleContent, 0)
			link = append(link, linkarr...)
		}
		count, err := models.SearchProductCount(key, "")
		if err != nil {
			logs.Error(err.Error)
		}
		table := prodTableserver{link, page1, count}
		c.Data["json"] = table
		c.ServeJSON()
	} else {
		c.Data["json"] = "关键字为空！"
		c.ServeJSON()
	}
}

// @Title get wx drawings list
// @Description get drawings by page
// @Param keyword query string false "The keyword of drawings"
// @Param projectid query string false "The projectid of drawings"
// @Param searchpage query string true "The page for drawings list"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 drawings not found
// @router /searchwxdrawings [get]
// 小程序取得所有图纸列表，分页_plus
func (c *SearchController) SearchWxDrawings() {
	// wxsite := web.AppConfig.String("wxreqeustsite")
	limit := "8"
	limit1, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("searchpage")
	page1, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	var offset int64
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	pid := c.GetString("projectid")
	var pidNum int64
	if pid != "" {
		pidNum, err = strconv.ParseInt(pid, 10, 64)
		if err != nil {
			logs.Error(err)
		}
	}

	key := c.GetString("keyword")
	var products []*models.Product
	if key != "" {
		if pidNum == 0 { //搜索所有成果
			products, err = models.SearchProductPage(limit1, offset, key, "")
			if err != nil {
				logs.Error(err.Error)
			}
		} else {
			products, err = models.SearchProjProductPage(pidNum, limit1, offset, key, "")
			if err != nil {
				logs.Error(err.Error)
			}
		}
	} else {
		products, err = models.SearchProjProductPage(pidNum, limit1, offset, key, "")
		if err != nil {
			logs.Error(err.Error)
		}
		// c.Data["json"] = map[string]interface{}{"info": "关键字为空"}
		// c.ServeJSON()
	}
	Pdfslice := make([]PdfLink, 0)
	for _, w := range products {
		//取到每个成果的附件（模态框打开）；pdf、文章——新窗口打开
		//循环成果
		//每个成果取到所有附件
		//一个附件则直接打开/下载；2个以上则打开模态框
		Attachments, err := models.GetAttachments(w.Id)
		if err != nil {
			logs.Error(err)
		}
		//对成果进行循环
		//赋予url
		for _, v := range Attachments {
			if path.Ext(v.FileName) == ".pdf" || path.Ext(v.FileName) == ".PDF" {
				pdfarr := make([]PdfLink, 1)
				pdfarr[0].Id = v.Id
				pdfarr[0].Title = v.FileName
				// if pidNum == 25002 { //图
				pdfarr[0].Link = "data:image/svg+xml,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20width%3D%2232%22%20height%3D%2232%22%3E%3Crect%20fill%3D%22%233F51B5%22%20x%3D%220%22%20y%3D%220%22%20width%3D%22100%25%22%20height%3D%22100%25%22%3E%3C%2Frect%3E%3Ctext%20fill%3D%22%23FFF%22%20x%3D%2250%25%22%20y%3D%2250%25%22%20text-anchor%3D%22middle%22%20font-size%3D%2216%22%20font-family%3D%22Verdana%2C%20Geneva%2C%20sans-serif%22%20alignment-baseline%3D%22middle%22%3E%E5%9B%BE%3C%2Ftext%3E%3C%2Fsvg%3E" //wxsite + "/static/img/go.jpg" //当做微信里的src来用
				pdfarr[0].ActIndex = "drawing"
				// } else { //纪
				// pdfarr[0].Link = "data:image/svg+xml,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20width%3D%2232%22%20height%3D%2232%22%3E%3Crect%20fill%3D%22%23009688%22%20x%3D%220%22%20y%3D%220%22%20width%3D%22100%25%22%20height%3D%22100%25%22%3E%3C%2Frect%3E%3Ctext%20fill%3D%22%23FFF%22%20x%3D%2250%25%22%20y%3D%2250%25%22%20text-anchor%3D%22middle%22%20font-size%3D%2216%22%20font-family%3D%22Verdana%2C%20Geneva%2C%20sans-serif%22%20alignment-baseline%3D%22middle%22%3E%E7%BA%AA%3C%2Ftext%3E%3C%2Fsvg%3E" //wxsite + "/static/img/go.jpg" //当做微信里的src来用
				// pdfarr[0].ActIndex = "other"
				// }
				pdfarr[0].Created = v.Created
				// timeformatdate, _ := time.Parse(datetime, thisdate)
				// const lll = "2006-01-02 15:04"
				pdfarr[0].Updated = v.Updated //.Format(lll)
				Pdfslice = append(Pdfslice, pdfarr...)
			}
		}

		//取得文章
		// Articles, err := models.GetArticles(w.Id)
		// if err != nil {
		// 	logs.Error(err)
		// }
		// for _, x := range Articles {
		// 	articlearr := make([]ArticleContent, 1)
		// 	articlearr[0].Id = x.Id
		// 	articlearr[0].Content = x.Content //返回的content是空，因为真要返回内容，会很影响速度，也没必要
		// 	articlearr[0].Link = "/project/product/article"
		// 	Articleslice = append(Articleslice, articlearr...)
		// }
		// linkarr[0].Articlecontent = Articleslice
	}
	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "searchers": Pdfslice}
	c.ServeJSON()
	// var user models.User
	// //取出用户openid
	// JSCODE := c.GetString("code")
	// if JSCODE != "" {
	// 	APPID := web.AppConfig.String("wxAPPID2")
	// 	SECRET := web.AppConfig.String("wxSECRET2")
	// 	app_version := c.GetString("app_version")
	// 	if app_version == "3" {
	// 		APPID = web.AppConfig.String("wxAPPID3")
	// 		SECRET = web.AppConfig.String("wxSECRET3")
	// 	}
	// 	// APPID := "wx7f77b90a1a891d93"
	// 	// SECRET := "f58ca4f28cbb52ccd805d66118060449"
	// 	requestUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=" + APPID + "&secret=" + SECRET + "&js_code=" + JSCODE + "&grant_type=authorization_code"
	// 	resp, err := http.Get(requestUrl)
	// 	if err != nil {
	// 		logs.Error(err)
	// 		return
	// 	}
	// 	defer resp.Body.Close()
	// 	if resp.StatusCode != 200 {
	// 		logs.Error(err)
	// 	}
	// 	var data map[string]interface{}
	// 	err = json.NewDecoder(resp.Body).Decode(&data)
	// 	if err != nil {
	// 		logs.Error(err)
	// 	}
	// 	var openID string
	// 	if _, ok := data["session_key"]; !ok {
	// 		errcode := data["errcode"]
	// 		errmsg := data["errmsg"].(string)
	// 		c.Data["json"] = map[string]interface{}{"errNo": errcode, "msg": errmsg, "data": "session_key 不存在"}
	// 	} else {
	// 		openID = data["openid"].(string)
	// 		user, err = models.GetUserByOpenID(openID)
	// 		if err != nil {
	// 			logs.Error(err)
	// 		}
	// 	}
	// }
	// var userid int64
	// if user.Nickname != "" {
	// 	userid = user.Id
	// } else {
	// 	userid = 0
	// }
}

// @Title get wx pdf list
// @Description get pdf by page
// @Param keyword query string  false "The keyword of pdf"
// @Param projectid query string  false "The projectid of pdf"
// @Param searchpage query string  true "The page for pdf list"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 pdf not found
// @router /getwxpdflist [get]
// 小程序取得某个项目目录下的所有pdf，分页
// 结合搜索来用
func (c *SearchController) GetWxPdfList() {
	// wxsite := web.AppConfig.String("wxreqeustsite")
	limit := "6"
	limit1, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("searchpage")
	page1, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	var offset int64
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	key := c.GetString("keyword")
	var products []*models.Product

	pid := c.GetString("projectid")
	// beego.Info(pid)
	var pidNum int64
	if pid != "" && key == "" {
		pidNum, err = strconv.ParseInt(pid, 10, 64)
		if err != nil {
			logs.Error(err)
		}
		products, err = models.SearchProjProductPage(pidNum, limit1, offset, "", "")
		if err != nil {
			logs.Error(err.Error)
		}
	} else if pid != "" && key != "" {
		pidNum, err = strconv.ParseInt(pid, 10, 64)
		if err != nil {
			logs.Error(err)
		}
		products, err = models.SearchProjProductPage(pidNum, limit1, offset, key, "")
		if err != nil {
			logs.Error(err.Error)
		}
	} else {
		c.Data["json"] = map[string]interface{}{"info": "关键字或项目为空"}
		c.ServeJSON()
		return
	}
	Pdfslice := make([]PdfLink, 0)
	for _, w := range products {
		Attachments, err := models.GetAttachments(w.Id)
		if err != nil {
			logs.Error(err)
		}
		//对成果进行循环
		//赋予url
		for _, v := range Attachments {
			if path.Ext(v.FileName) == ".pdf" || path.Ext(v.FileName) == ".PDF" {
				pdfarr := make([]PdfLink, 1)
				pdfarr[0].Id = v.Id
				pdfarr[0].Title = v.FileName
				if pidNum == 25002 { //图
					pdfarr[0].Link = "data:image/svg+xml,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20width%3D%2232%22%20height%3D%2232%22%3E%3Crect%20fill%3D%22%233F51B5%22%20x%3D%220%22%20y%3D%220%22%20width%3D%22100%25%22%20height%3D%22100%25%22%3E%3C%2Frect%3E%3Ctext%20fill%3D%22%23FFF%22%20x%3D%2250%25%22%20y%3D%2250%25%22%20text-anchor%3D%22middle%22%20font-size%3D%2216%22%20font-family%3D%22Verdana%2C%20Geneva%2C%20sans-serif%22%20alignment-baseline%3D%22middle%22%3E%E5%9B%BE%3C%2Ftext%3E%3C%2Fsvg%3E" //wxsite + "/static/img/go.jpg" //当做微信里的src来用
					pdfarr[0].ActIndex = "drawing"
				} else { //纪
					pdfarr[0].Link = "data:image/svg+xml,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20width%3D%2232%22%20height%3D%2232%22%3E%3Crect%20fill%3D%22%23009688%22%20x%3D%220%22%20y%3D%220%22%20width%3D%22100%25%22%20height%3D%22100%25%22%3E%3C%2Frect%3E%3Ctext%20fill%3D%22%23FFF%22%20x%3D%2250%25%22%20y%3D%2250%25%22%20text-anchor%3D%22middle%22%20font-size%3D%2216%22%20font-family%3D%22Verdana%2C%20Geneva%2C%20sans-serif%22%20alignment-baseline%3D%22middle%22%3E%E7%BA%AA%3C%2Ftext%3E%3C%2Fsvg%3E" //wxsite + "/static/img/go.jpg" //当做微信里的src来用
					pdfarr[0].ActIndex = "other"
				}
				pdfarr[0].Created = v.Created
				// timeformatdate, _ := time.Parse(datetime, thisdate)
				// const lll = "2006-01-02 15:04"
				pdfarr[0].Updated = v.Updated //.Format(lll)
				Pdfslice = append(Pdfslice, pdfarr...)
			}
		}
	}
	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "searchers": Pdfslice}
	c.ServeJSON()
}

type WxProduct struct {
	Id      int64
	Title   string
	Type    string
	Link    string
	Subtext string
	Author  string
	Created time.Time
	Updated time.Time
}

// @Title get wx drawings list
// @Description get drawings by page
// @Param keyword query string false "The keyword of drawings"
// @Param projectid query string false "The projectid of drawings"
// @Param searchpage query string true "The page for drawings list"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 drawings not found
// @router /searchwxproducts [get]
// 小程序取得所有图纸列表，分页_plus
func (c *SearchController) SearchWxProducts() {
	wxsite, err := web.AppConfig.String("wxreqeustsite")
	if err != nil {
		logs.Error(err)
	}
	limit := "8"
	limit1, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("searchpage")
	page1, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	var offset int64
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	pid := c.GetString("projectid")
	var pidNum int64
	if pid != "" {
		pidNum, err = strconv.ParseInt(pid, 10, 64)
		if err != nil {
			logs.Error(err)
		}
	}

	key := c.GetString("keyword")
	var products []*models.Product
	if key != "" {
		if pidNum == 0 { //搜索所有成果
			products, err = models.SearchProductPage(limit1, offset, key, "")
			if err != nil {
				logs.Error(err.Error)
			}
		} else {
			products, err = models.SearchProjProductPage(pidNum, limit1, offset, key, "")
			if err != nil {
				logs.Error(err.Error)
			}
		}
	} else {
		products, err = models.SearchProjProductPage(pidNum, limit1, offset, key, "")
		if err != nil {
			logs.Error(err.Error)
		}
		// c.Data["json"] = map[string]interface{}{"info": "关键字为空"}
		// c.ServeJSON()
	}
	productslice := make([]WxProduct, 0)
	for _, w := range products {
		//取到每个成果的附件（模态框打开）；pdf、文章——新窗口打开
		//循环成果
		//每个成果取到所有附件
		//一个附件则直接打开/下载；2个以上则打开模态框
		Attachments, err := models.GetAttachments(w.Id)
		if err != nil {
			logs.Error(err)
		}
		//对成果进行循环
		//赋予url
		for _, v := range Attachments {
			// if path.Ext(v.FileName) == ".pdf" || path.Ext(v.FileName) == ".PDF" {
			wxproductarr := make([]WxProduct, 1)
			wxproductarr[0].Id = v.Id
			wxproductarr[0].Title = v.FileName
			// if pidNum == 25002 { //图
			// wxproductarr[0].ActIndex = "drawing"
			switch path.Ext(v.FileName) {
			case ".pdf", ".PDF":
				wxproductarr[0].Type = "pdf"
				wxproductarr[0].Link = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMDAiIGhlaWdodD0iMTAwIj48cmVjdCBmaWxsPSIjMDA5Njg4IiB4PSIwIiB5PSIwIiB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIj48L3JlY3Q+PHRleHQgZmlsbD0iI0ZGRiIgeD0iNTAlIiB5PSI1MCUiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGZvbnQtc2l6ZT0iNDQiIGZvbnQtZmFtaWx5PSJzYW5zLXNlcmlmIiBhbGlnbm1lbnQtYmFzZWxpbmU9Im1pZGRsZSI+UERGPC90ZXh0Pjwvc3ZnPg=="
			case ".doc", ".DOC":
				wxproductarr[0].Type = "doc"
				wxproductarr[0].Link = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMDAiIGhlaWdodD0iMTAwIj48cmVjdCBmaWxsPSIjRkY1NzIyIiB4PSIwIiB5PSIwIiB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIj48L3JlY3Q+PHRleHQgZmlsbD0iI0ZGRiIgeD0iNTAlIiB5PSI1MCUiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGZvbnQtc2l6ZT0iNTAiIGZvbnQtZmFtaWx5PSJzYW5zLXNlcmlmIiBhbGlnbm1lbnQtYmFzZWxpbmU9Im1pZGRsZSI+VzwvdGV4dD48L3N2Zz4="
			case ".docx", ".DOCX":
				wxproductarr[0].Type = "docx"
				wxproductarr[0].Link = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMDAiIGhlaWdodD0iMTAwIj48cmVjdCBmaWxsPSIjRkY1NzIyIiB4PSIwIiB5PSIwIiB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIj48L3JlY3Q+PHRleHQgZmlsbD0iI0ZGRiIgeD0iNTAlIiB5PSI1MCUiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGZvbnQtc2l6ZT0iNTAiIGZvbnQtZmFtaWx5PSJzYW5zLXNlcmlmIiBhbGlnbm1lbnQtYmFzZWxpbmU9Im1pZGRsZSI+VzwvdGV4dD48L3N2Zz4="
			case ".xls", ".XLS":
				wxproductarr[0].Type = "xls"
				wxproductarr[0].Link = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMDAiIGhlaWdodD0iMTAwIj48cmVjdCBmaWxsPSIjNzk1NTQ4IiB4PSIwIiB5PSIwIiB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIj48L3JlY3Q+PHRleHQgZmlsbD0iI0ZGRiIgeD0iNTAlIiB5PSI1MCUiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGZvbnQtc2l6ZT0iNTAiIGZvbnQtZmFtaWx5PSJzYW5zLXNlcmlmIiBhbGlnbm1lbnQtYmFzZWxpbmU9Im1pZGRsZSI+WDwvdGV4dD48L3N2Zz4="
			case ".xlsx", ".XLSX":
				wxproductarr[0].Type = "xlsx"
				wxproductarr[0].Link = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMDAiIGhlaWdodD0iMTAwIj48cmVjdCBmaWxsPSIjNzk1NTQ4IiB4PSIwIiB5PSIwIiB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIj48L3JlY3Q+PHRleHQgZmlsbD0iI0ZGRiIgeD0iNTAlIiB5PSI1MCUiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGZvbnQtc2l6ZT0iNTAiIGZvbnQtZmFtaWx5PSJzYW5zLXNlcmlmIiBhbGlnbm1lbnQtYmFzZWxpbmU9Im1pZGRsZSI+WDwvdGV4dD48L3N2Zz4="
			case ".ppt", ".PPT":
				wxproductarr[0].Type = "ppt"
				wxproductarr[0].Link = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMDAiIGhlaWdodD0iMTAwIj48cmVjdCBmaWxsPSIjMDA5Njg4IiB4PSIwIiB5PSIwIiB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIj48L3JlY3Q+PHRleHQgZmlsbD0iI0ZGRiIgeD0iNTAlIiB5PSI1MCUiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGZvbnQtc2l6ZT0iNTAiIGZvbnQtZmFtaWx5PSJzYW5zLXNlcmlmIiBhbGlnbm1lbnQtYmFzZWxpbmU9Im1pZGRsZSI+UDwvdGV4dD48L3N2Zz4="
			case ".pptx", ".PPTX":
				wxproductarr[0].Type = "pptx"
				wxproductarr[0].Link = "data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIxMDAiIGhlaWdodD0iMTAwIj48cmVjdCBmaWxsPSIjMDA5Njg4IiB4PSIwIiB5PSIwIiB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIj48L3JlY3Q+PHRleHQgZmlsbD0iI0ZGRiIgeD0iNTAlIiB5PSI1MCUiIHRleHQtYW5jaG9yPSJtaWRkbGUiIGZvbnQtc2l6ZT0iNTAiIGZvbnQtZmFtaWx5PSJzYW5zLXNlcmlmIiBhbGlnbm1lbnQtYmFzZWxpbmU9Im1pZGRsZSI+UDwvdGV4dD48L3N2Zz4="
			}
			wxproductarr[0].Created = v.Created
			// timeformatdate, _ := time.Parse(datetime, thisdate)
			// const lll = "2006-01-02 15:04"
			wxproductarr[0].Updated = v.Updated //.Format(lll)
			productslice = append(productslice, wxproductarr...)
			// }
		}

		//取得文章
		// Articles, err := models.GetArticles(w.Id)
		// if err != nil {
		// 	logs.Error(err)
		// }
		Articles, err := models.GetWxArticles(w.Id)
		if err != nil {
			logs.Error(err)
		}
		for _, x := range Articles {
			wxproductarr := make([]WxProduct, 1)
			wxproductarr[0].Id = x.Id
			wxproductarr[0].Title = w.Title
			wxproductarr[0].Type = "isArticle"
			wxproductarr[0].Created = x.Created
			wxproductarr[0].Updated = x.Updated
			wxproductarr[0].Subtext = x.Subtext
			wxproductarr[0].Author = w.Principal
			//取到文章里的图片地址
			slice2 := make([]string, 0)
			var r io.Reader = strings.NewReader(string(x.Content))
			doc, err := goquery.NewDocumentFromReader(r)
			if err != nil {
				logs.Error(err)
			}
			doc.Find("img").Each(func(i int, s *goquery.Selection) {
				sel, _ := s.Attr("src")
				aa := make([]string, 1)
				aa[0] = sel
				// aa[0].Name = path.Base(sel)
				slice2 = append(slice2, aa...)
			})
			if len(slice2) > 0 {
				wxproductarr[0].Link = wxsite + slice2[0]
			} else {
				wxproductarr[0].Link = wxsite + "/static/img/go.jpg"
			}
			// wxproductarr[0].Link = "data:image/svg+xml,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20width%3D%2232%22%20height%3D%2232%22%3E%3Crect%20fill%3D%22%233F51B5%22%20x%3D%220%22%20y%3D%220%22%20width%3D%22100%25%22%20height%3D%22100%25%22%3E%3C%2Frect%3E%3Ctext%20fill%3D%22%23FFF%22%20x%3D%2250%25%22%20y%3D%2250%25%22%20text-anchor%3D%22middle%22%20font-size%3D%2216%22%20font-family%3D%22Verdana%2C%20Geneva%2C%20sans-serif%22%20alignment-baseline%3D%22middle%22%3E%E5%9B%BE%3C%2Ftext%3E%3C%2Fsvg%3E" //wxsite + "/static/img/go.jpg" //当做微信里的src来用
			productslice = append(productslice, wxproductarr...)
		}
	}
	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "searchers": productslice}
	c.ServeJSON()
}

// 未修改
func (c *SearchController) SearchProjects() {
	limit := "15"
	limit1, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("searchpage")
	page1, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	var offset int64
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}
	pid := c.GetString("productid")
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	key := c.GetString("keyword")
	if key != "" {
		_, products, err := models.SearchProjProduct(pidNum, limit1, offset, key, "") //这里要将侧栏所有id进行循环
		if err != nil {
			logs.Error(err.Error)
		} else {
			c.Data["json"] = products
			c.ServeJSON()
		}
	}
}

// 搜索wiki
func (c *SearchController) SearchWiki() { //search用的是get方法
	tid := c.GetString("wikiname")
	if tid != "" {
		c.Data["IsWiki"] = true
		// c.Data["IsSearch"] = true
		c.Data["IsLogin"] = checkAccount(c.Ctx)
		c.TplName = "searchwiki.tpl"
		Searchs, err := models.SearchWikis(tid, false)
		if err != nil {
			logs.Error(err.Error)
		} else {
			c.Data["Searchs"] = Searchs
		}
	} else {
		c.Data["IsWiki"] = true
		// c.Data["IsSearch"] = true
		c.Data["IsLogin"] = checkAccount(c.Ctx)
		c.TplName = "searchwiki.tpl"
	}
}
