package controllers

import (
	"encoding/json"
	"github.com/3xxx/engineercms/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"time"
	// "os"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"path"
	"regexp"
	"strconv"
	"strings"
)

// CMSWX article API
type ArticleController struct {
	beego.Controller
}

//取得某个成果id下的文章给table
func (c *ArticleController) GetArticles() {
	pid := c.Ctx.Input.Param(":id")
	c.Data["Id"] = pid
	var pidNum int64
	var err error
	if pid != "" {
		//id转成64为
		pidNum, err = strconv.ParseInt(pid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//根据成果id取得所有文章
		Articles, err := models.GetArticles(pidNum)
		if err != nil {
			beego.Error(err)
		}
		//查出成果编号，名称和作者
		prod, err := models.GetProd(pidNum)
		if err != nil {
			beego.Error(err)
		}
		//对文章进行循环
		//赋予url
		//如果是一个文章，直接给url;如果大于1个，则是数组:这个在前端实现
		// http.ServeFile(ctx.ResponseWriter, ctx.Request, filePath)
		Articleslice := make([]ArticleContent, 0)
		for _, v := range Articles {
			articlearr := make([]ArticleContent, 1)
			articlearr[0].Id = v.Id
			articlearr[0].Title = prod.Title
			articlearr[0].Subtext = v.Subtext
			articlearr[0].ProductId = v.ProductId
			articlearr[0].Content = v.Content
			articlearr[0].Created = v.Created
			articlearr[0].Updated = v.Updated
			articlearr[0].Link = "/project/product/article/" + strconv.FormatInt(v.Id, 10)
			Articleslice = append(Articleslice, articlearr...)
		}
		c.Data["json"] = Articleslice
		c.ServeJSON()
	} else {

	}
}

//取得同步文章列表
func (c *ArticleController) GetsynchArticles() {
	id := c.Input().Get("id")
	site := c.Input().Get("site")
	// beego.Info(site)
	link := make([]ArticleContent, 0)
	var Articleslice []ArticleContent
	//	通过如下接口可以设置请求的超时时间和数据读取时间
	jsonstring, err := httplib.Get(site+"project/product/providesyncharticles?id="+id).SetTimeout(100*time.Second, 30*time.Second).String() //.ToJSON(&productlink)
	if err != nil {
		beego.Error(err)
	}
	//json字符串解析到结构体，以便进行追加
	err = json.Unmarshal([]byte(jsonstring), &Articleslice)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(productlink)
	link = append(link, Articleslice...)

	c.Data["json"] = link //products
	c.ServeJSON()
}

//提供同步文章列表
func (c *ArticleController) ProvideArticles() {
	pid := c.Input().Get("id")
	site := c.Ctx.Input.Site() + ":" + strconv.Itoa(c.Ctx.Input.Port())
	c.Data["Id"] = pid
	var pidNum int64
	var err error
	if pid != "" {
		//id转成64为
		pidNum, err = strconv.ParseInt(pid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//根据成果id取得所有文章
		Articles, err := models.GetArticles(pidNum)
		if err != nil {
			beego.Error(err)
		}
		//查出成果编号，名称和作者
		prod, err := models.GetProd(pidNum)
		if err != nil {
			beego.Error(err)
		}
		//对文章进行循环
		//赋予url
		//如果是一个文章，直接给url;如果大于1个，则是数组:这个在前端实现
		// http.ServeFile(ctx.ResponseWriter, ctx.Request, filePath)
		Articleslice := make([]ArticleContent, 0)
		for _, v := range Articles {
			articlearr := make([]ArticleContent, 1)
			articlearr[0].Id = v.Id
			articlearr[0].Title = prod.Title
			articlearr[0].Subtext = v.Subtext
			articlearr[0].ProductId = v.ProductId
			articlearr[0].Content = v.Content
			articlearr[0].Created = v.Created
			articlearr[0].Updated = v.Updated
			articlearr[0].Link = site + "/project/product/article/" + strconv.FormatInt(v.Id, 10)
			Articleslice = append(Articleslice, articlearr...)
		}
		c.Data["json"] = Articleslice
		c.ServeJSON()
	} else {

	}
}

//提供给成果列表的table中json数据
// func (c *ArticleController) GetArticles() {
// 	id := c.Ctx.Input.Param(":id")
// 	if id == "" {
// 		//显示全部
// 		Articles, err := models.GetArticles()
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		c.Data["json"] = Articles
// 		c.ServeJSON()
// 	} else {
// 		//根据标签查询
// 	}
// }

//根据id查看，查出文章
func (c *ArticleController) GetArticle() {
	id := c.Ctx.Input.Param(":id")

	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	useridstring := strconv.FormatInt(uid, 10)
	// var categories []*models.ProjCategory
	var err error
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	Article, err := models.GetArticle(idNum)
	if err != nil {
		beego.Error(err)
	}
	//查出成果编号，名称和作者
	prod, err := models.GetProd(Article.ProductId)
	if err != nil {
		beego.Error(err)
	}

	//2.取得侧栏目录路径——路由id
	//2.1 根据id取得路由
	var projurls string
	proj, err := models.GetProj(prod.ProjectId)
	if err != nil {
		beego.Error(err)
	}
	if proj.ParentId == 0 { //如果是项目根目录
		projurls = "/" + strconv.FormatInt(proj.Id, 10)
	} else {
		// projurls = "/" + strings.Replace(proj.ParentIdPath, "-", "/", -1) + "/" + strconv.FormatInt(proj.Id, 10)
		projurls = "/" + strings.Replace(strings.Replace(proj.ParentIdPath, "#", "/", -1), "$", "", -1) + strconv.FormatInt(proj.Id, 10)
	}

	//上一篇和下一篇
	//根据项目id取得所有成果
	products, err := models.GetProducts(proj.Id)
	if err != nil {
		beego.Error(err)
	}

	// Attachslice := make([]AttachmentLink, 0)
	Articles := make([]*models.Article, 0)
	for _, w := range products {
		//取到每个成果的附件（模态框打开）；pdf、文章——新窗口打开
		//循环成果
		//每个成果取到所有附件
		//一个附件则直接打开/下载；2个以上则打开模态框
		//取得文章
		Articles1, err := models.GetArticles(w.Id)
		if err != nil {
			beego.Error(err)
		}
		Articles = append(Articles, Articles1...)
	}
	// count := len(Articles)
	// count1 := strconv.Itoa(count)
	// count2, err := strconv.ParseInt(count1, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// if p == "" {//这里是用于在页面打开pdf附件的时候，显示地址是/pdf?id=628，id是附件id
	// var p1 string
	for i, v := range Articles {
		if v.Id == idNum { //idnumb是附件的id
			// p1 = strconv.Itoa(i + 1)
			// PdfLink = Url + "/" + v.FileName
			//上一篇
			if i > 0 {
				c.Data["NextArticleId"] = strconv.FormatInt(Articles[i-1].Id, 10)
				//由文章id取得prodtitle
				//查出成果编号，名称和作者
				articleprod, err := models.GetProd(Articles[i-1].ProductId)
				if err != nil {
					beego.Error(err)
				}
				c.Data["NextArticleTitle"] = articleprod.Title
				c.Data["Next"] = true
			} else {
				c.Data["Next"] = false
			}
			if i < len(Articles)-1 {
				c.Data["PreArticleId"] = strconv.FormatInt(Articles[i+1].Id, 10)
				//由文章id取得prodtitle
				//查出成果编号，名称和作者
				articleprod, err := models.GetProd(Articles[i+1].ProductId)
				if err != nil {
					beego.Error(err)
				}
				c.Data["PreArticleTitle"] = articleprod.Title
				c.Data["Pre"] = true
			} else {
				c.Data["Pre"] = false
			}
			break
		}
	}
	// if e.Enforce(useridstring, projurls+"/", "POST", ".1") {
	// 	c.Data["RoleAdd"] = "true"
	// } else {
	// 	c.Data["RoleAdd"] = "false"
	// }
	if e.Enforce(useridstring, projurls+"/", "PUT", ".1") || uid == prod.Uid {
		c.Data["RoleUpdate"] = true
	} else {
		c.Data["RoleUpdate"] = false
	}
	if e.Enforce(useridstring, projurls+"/", "DELETE", ".1") || uid == prod.Uid {
		c.Data["RoleDelete"] = true
	} else {
		c.Data["RoleDelete"] = false
	}
	// c.Data["productid"] = prod.Uid//文章作者id
	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		beego.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		c.TplName = "marticle.tpl"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "article.tpl"
	}
	c.Data["article"] = Article
	c.Data["product"] = prod
}

type WxArticle struct {
	Id          int64     `json:"id",form:"-"`
	Title       string    `json:"title"`
	Subtext     string    `json:"subtext",orm:"sie(20)"`
	Author      string    `json:"author"`
	ImgUrl      string    `json:"imgUrl"`
	ImgUrls     []Img     `json:"imgUrls"`
	Content     string    `json:"html",orm:"sie(5000)"`
	LeassonType int       `json:"leassonType"`
	ProductId   int64     `orm:"null"`
	Views       int64     `orm:"default(0)"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     string    `json:"time",orm:"auto_now_add;type(datetime)"`
	Word        string    `json:"word",orm:"sie(5000)"`
}

type Img struct {
	Src  string `json:"src"`
	Name string `json:"name"`
}

//后端分页的数据结构
type prodWxTableserver struct {
	Rows  []WxArticle `json:"rows"`
	Page  int64       `json:"page"`
	Total int64       `json:"total"` //string或int64都行！
}

// @Title get wx artiles list
// @Description get articles by page
// @Param page query string  true "The page for articles list"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getwxarticles [get]
//小程序取得所有文章列表，分页
func (c *ArticleController) GetWxArticles() {
	// id := c.Ctx.Input.Param(":id")
	id := beego.AppConfig.String("wxcatalogid") //"26159" //25002珠三角设代日记id26159
	wxsite := beego.AppConfig.String("wxreqeustsite")
	limit := "5"
	limit1, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	page := c.Input().Get("page")
	page1, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	var idNum int64
	//id转成64为
	idNum, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	var offset int64
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}
	// articleslice := make([]*WxArticle, 0)
	// getwxarticles(idNum, limit1, offset, articleslice, wxsite)
	//根据项目id取得所有成果
	//bug_如果5个成果里都没有文章，则显示文章失败；如果多个文章，会超过5个
	products, err := models.GetProductsPage(idNum, limit1, offset, "")
	if err != nil {
		beego.Error(err)
	}
	Articleslice := make([]WxArticle, 0)
	for _, w := range products {
		//取得文章
		Articles, err := models.GetWxArticles(w.Id)
		if err != nil {
			beego.Error(err)
		}
		for _, x := range Articles {
			//取到文章里的图片地址
			slice2 := make([]Img, 0)
			var r io.Reader = strings.NewReader(string(x.Content))
			doc, err := goquery.NewDocumentFromReader(r)
			if err != nil {
				beego.Error(err)
			}
			doc.Find("img").Each(func(i int, s *goquery.Selection) {
				sel, _ := s.Attr("src")
				aa := make([]Img, 1)
				aa[0].Src = sel
				aa[0].Name = path.Base(sel)
				slice2 = append(slice2, aa...)
			})
			articlearr := make([]WxArticle, 1)
			articlearr[0].Id = x.Id
			articlearr[0].Title = w.Title
			articlearr[0].Subtext = x.Subtext
			articlearr[0].Author = w.Principal
			if len(slice2) > 0 {
				articlearr[0].ImgUrl = wxsite + slice2[0].Src
			} else {
				articlearr[0].ImgUrl = wxsite + "/static/img/go.jpg"
			}
			articlearr[0].LeassonType = 1
			articlearr[0].ProductId = x.ProductId
			Articleslice = append(Articleslice, articlearr...)
		}
	}
	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "articles": Articleslice}
	c.ServeJSON()
}

// @Title get wx artiles list
// @Description get articles by page
// @Param id path string  true "The id of project"
// @Param page query string  true "The page for articles list"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getwxarticless/:id [get]
//小程序取得所有文章列表，分页
func (c *ArticleController) GetWxArticless() {
	id := c.Ctx.Input.Param(":id")
	// id := beego.AppConfig.String("wxcatalogid") //"26159" //25002珠三角设代日记id26159
	wxsite := beego.AppConfig.String("wxreqeustsite")
	limit := "5"
	limit1, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	page := c.Input().Get("page")
	page1, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	var idNum int64
	//id转成64为
	idNum, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	var offset int64
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}
	// articleslice := make([]*WxArticle, 0)
	// getwxarticles(idNum, limit1, offset, articleslice, wxsite)
	//根据项目id取得所有成果
	//bug_如果5个成果里都没有文章，则显示文章失败；如果多个文章，会超过5个
	products, err := models.GetProductsPage(idNum, limit1, offset, "")
	if err != nil {
		beego.Error(err)
	}
	Articleslice := make([]WxArticle, 0)
	for _, w := range products {
		//取得文章
		Articles, err := models.GetWxArticles(w.Id)
		if err != nil {
			beego.Error(err)
		}
		for _, x := range Articles {
			//取到文章里的图片地址
			slice2 := make([]Img, 0)
			var r io.Reader = strings.NewReader(string(x.Content))
			doc, err := goquery.NewDocumentFromReader(r)
			if err != nil {
				beego.Error(err)
			}
			doc.Find("img").Each(func(i int, s *goquery.Selection) {
				sel, _ := s.Attr("src")
				aa := make([]Img, 1)
				aa[0].Src = sel
				aa[0].Name = path.Base(sel)
				slice2 = append(slice2, aa...)
			})
			articlearr := make([]WxArticle, 1)
			articlearr[0].Id = x.Id
			articlearr[0].Title = w.Title
			articlearr[0].Subtext = x.Subtext
			articlearr[0].Author = w.Principal
			articlearr[0].Views = x.Views
			if len(slice2) > 0 {
				articlearr[0].ImgUrl = wxsite + slice2[0].Src
			} else {
				articlearr[0].ImgUrl = wxsite + "/static/img/go.jpg"
			}
			articlearr[0].LeassonType = 1
			articlearr[0].ProductId = x.ProductId
			Articleslice = append(Articleslice, articlearr...)
		}
	}
	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "articles": Articleslice}
	c.ServeJSON()
}

//递归取得至少5篇文章——逻辑上无法实现
//仅供研究，没有用上
func getwxarticles(id, limit, offset int64, articleslice []*WxArticle, wxsite string) {
	//bug_如果5个成果里都没有文章，则显示文章失败；如果多个文章，会超过5个
	products, err := models.GetProductsPage(id, limit, offset, "")
	if err != nil {
		beego.Error(err)
	}
	// Articleslice := make([]WxArticle, 0)
	for _, w := range products {
		//取得文章
		Articles, err := models.GetWxArticles(w.Id)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(Articles)
		for _, x := range Articles {
			//取到文章里的图片地址
			slice2 := make([]Img, 0)
			var r io.Reader = strings.NewReader(string(x.Content))
			doc, err := goquery.NewDocumentFromReader(r)
			if err != nil {
				beego.Error(err)
			}
			doc.Find("img").Each(func(i int, s *goquery.Selection) {
				sel, _ := s.Attr("src")
				aa := make([]Img, 1)
				aa[0].Src = sel
				aa[0].Name = path.Base(sel)
				slice2 = append(slice2, aa...)
			})
			articlearr := make([]*WxArticle, 1)
			// beego.Info(x.Id)
			articlearr[0].Id = x.Id //不知为何这里出错？？
			// beego.Info(x.Id)
			articlearr[0].Title = w.Title
			articlearr[0].Subtext = x.Subtext
			articlearr[0].Author = w.Principal
			if len(slice2) > 0 {
				articlearr[0].ImgUrl = wxsite + slice2[0].Src
			} else {
				articlearr[0].ImgUrl = wxsite + "/static/img/go.jpg"
			}
			// articlearr[0].Content = x.Content
			articlearr[0].LeassonType = 1
			articlearr[0].ProductId = x.ProductId
			// articlearr[0].Views = x.Views
			// articlearr[0].Created = x.Created
			// articlearr[0].Updated = x.Updated
			articleslice = append(articleslice, articlearr...)
		}
		// linkarr[0].Articlecontent = Articleslice
		// Articleslice = make([]ArticleContent, 0)
		// link = append(link, linkarr...)
	}
	if len(articleslice) < 5 {
		getwxarticles(id, 1, offset+1, articleslice, wxsite)
	}
}

// @Title get wx artile by articleId
// @Description get article by articleid
// @Param id path string  true "The id of article"
// @Success 200 {object} models.GetArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /getwxarticle/:id [get]
//根据id查看一篇微信文章
func (c *ArticleController) GetWxArticle() {
	// username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	// c.Data["Username"] = username
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	// c.Data["IsAdmin"] = isadmin
	// c.Data["IsLogin"] = islogin
	// c.Data["Uid"] = uid
	// useridstring := strconv.FormatInt(uid, 10)
	id := c.Ctx.Input.Param(":id")
	wxsite := beego.AppConfig.String("wxreqeustsite")
	// beego.Info(id)
	var err error
	var photo string
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	Article, err := models.GetArticle(idNum)
	if err != nil {
		beego.Error(err)
	}
	//查出成果编号，名称和作者
	prod, err := models.GetProd(Article.ProductId)
	if err != nil {
		beego.Error(err)
	}
	// slice2 := make([]img, 0)
	// var r io.Reader = strings.NewReader(string(Article.Content))
	// doc, err := goquery.NewDocumentFromReader(r)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// doc.Find("p img").Each(func(i int, s *goquery.Selection) {
	// 	sel, _ := s.Attr("src")
	// 	httpssel := "https://zsj.itdos.com" + sel
	// 	aa := make([]img, 1)
	// 	aa[0].src = "https://zsj.itdos.com" + sel
	// 	aa[0].name = path.Base(sel)
	// 	slice2 = append(slice2, aa...)
	// })
	// if len(slice2) == 0 {
	// 	slice2 = make([]img, 1)
	// 	slice2[0].src = "https://zsj.itdos.com/static/img/go.jpg"
	// }
	//取到文章里的图片地址
	slice2 := make([]Img, 0)
	var slice3 string
	var r io.Reader = strings.NewReader(string(Article.Content))
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		beego.Error(err)
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		sel, _ := s.Attr("src")
		aa := make([]Img, 1)
		aa[0].Src = wxsite + sel
		aa[0].Name = path.Base(sel)
		slice2 = append(slice2, aa...)
	})
	doc.Each(func(i int, s *goquery.Selection) {
		word := s.Text()
		slice3 = slice3 + word
	})

	nameRune := []rune(slice3)

	if len(nameRune) >= 20 {
		slice3 = string(nameRune[0:20])
	}
	// articlearr := make([]WxArticle, 1)
	//第一个图片 wxsite +
	if len(slice2) > 0 {
		photo = slice2[0].Src
	} else {
		photo = wxsite + "/static/img/go.jpg"
	}

	content := strings.Replace(Article.Content, "/attachment/", wxsite+"/attachment/", -1)
	const lll = "2006-01-02 15:04"
	articletime := Article.Updated.Format(lll)
	// beego.Info(slice2)
	wxArticle := &WxArticle{
		Id:          Article.Id,
		Title:       prod.Title,
		Subtext:     Article.Subtext,
		Author:      prod.Principal,
		ImgUrl:      photo,
		ImgUrls:     slice2,
		Content:     content, //Article.Content,
		LeassonType: 1,
		ProductId:   Article.ProductId,
		Views:       Article.Views,
		Created:     Article.Created,
		Updated:     articletime,
		//给小程序用的图片和文本
		Word: slice3,
		// Photo: photo,
	}
	//以下都没用了
	// dataList = [
	//   {
	//     "id": 5,
	//     "imgUrl": "https://mmbiz.qpic.cn/mmbiz_jpg/UMDZj3FARlE2U5aRjrb6daAWWvXQkh7hskQEEX9Hnc9HpicK8j7WHCIpxJcXibIn9CHqM6u1XzUsnwgXaBeES3Aw/0?wx_fmt=jpeg",
	//     "title": prod.Title"教程（5）有趣又简单的挤娃娃画",
	//     "leassonType": 1,
	//     "html":
	// "time":
	// }
	// ]

	//2.取得侧栏目录路径——路由id
	//2.1 根据id取得路由
	// var projurls string
	proj, err := models.GetProj(prod.ProjectId)
	if err != nil {
		beego.Error(err)
	}
	// if proj.ParentId == 0 { //如果是项目根目录
	// 	projurls = "/" + strconv.FormatInt(proj.Id, 10)
	// } else {
	// 	// projurls = "/" + strings.Replace(proj.ParentIdPath, "-", "/", -1) + "/" + strconv.FormatInt(proj.Id, 10)
	// 	projurls = "/" + strings.Replace(strings.Replace(proj.ParentIdPath, "#", "/", -1), "$", "", -1) + strconv.FormatInt(proj.Id, 10)
	// }

	//上一篇和下一篇
	//根据项目id取得所有成果
	products, err := models.GetProducts(proj.Id)
	if err != nil {
		beego.Error(err)
	}

	// Attachslice := make([]AttachmentLink, 0)
	Articles := make([]*models.Article, 0)
	for _, w := range products {
		//取得文章
		Articles1, err := models.GetArticles(w.Id)
		if err != nil {
			beego.Error(err)
		}
		Articles = append(Articles, Articles1...)
	}
	// count := len(Articles)
	// count1 := strconv.Itoa(count)
	// count2, err := strconv.ParseInt(count1, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// if p == "" {//这里是用于在页面打开pdf附件的时候，显示地址是/pdf?id=628，id是附件id
	// var p1 string
	for i, v := range Articles {
		if v.Id == idNum { //idnumb是附件的id
			// p1 = strconv.Itoa(i + 1)
			// PdfLink = Url + "/" + v.FileName
			//上一篇
			if i > 0 {
				c.Data["NextArticleId"] = strconv.FormatInt(Articles[i-1].Id, 10)
				//由文章id取得prodtitle
				//查出成果编号，名称和作者
				articleprod, err := models.GetProd(Articles[i-1].ProductId)
				if err != nil {
					beego.Error(err)
				}
				c.Data["NextArticleTitle"] = articleprod.Title
				c.Data["Next"] = true
			} else {
				c.Data["Next"] = false
			}
			if i < len(Articles)-1 {
				c.Data["PreArticleId"] = strconv.FormatInt(Articles[i+1].Id, 10)
				//由文章id取得prodtitle
				//查出成果编号，名称和作者
				articleprod, err := models.GetProd(Articles[i+1].ProductId)
				if err != nil {
					beego.Error(err)
				}
				c.Data["PreArticleTitle"] = articleprod.Title
				c.Data["Pre"] = true
			} else {
				c.Data["Pre"] = false
			}
			break
		}
	}
	// if e.Enforce(useridstring, projurls+"/", "POST", ".1") {
	// 	c.Data["RoleAdd"] = "true"
	// } else {
	// 	c.Data["RoleAdd"] = "false"
	// }
	// if e.Enforce(useridstring, projurls+"/", "PUT", ".1") {
	// 	c.Data["RoleUpdate"] = "true"
	// } else {
	// 	c.Data["RoleUpdate"] = "false"
	// }
	// if e.Enforce(useridstring, projurls+"/", "DELETE", ".1") {
	// 	c.Data["RoleDelete"] = "true"
	// } else {
	// 	c.Data["RoleDelete"] = "false"
	// }
	// // c.Data["productid"] = prod.Uid//文章作者id
	// u := c.Ctx.Input.UserAgent()
	// matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// if matched == true {
	// 	// beego.Info("移动端~")
	// 	c.TplName = "marticle.tpl"
	// } else {
	// 	// beego.Info("电脑端！")
	// 	c.TplName = "article.tpl"
	// }
	// c.Data["article"] = Article
	c.Data["json"] = wxArticle
	c.ServeJSON()
	// c.Data["product"] = prod
}

//向某个侧栏id下添加文章
func (c *ArticleController) AddArticle() {
	_, _, uid, _, _ := checkprodRole(c.Ctx)

	meritbasic, err := models.GetMeritBasic()
	if err != nil {
		beego.Error(err)
	}
	var catalog models.PostMerit
	var news string
	var cid int64

	pid := c.Input().Get("pid")
	code := c.Input().Get("code")
	title := c.Input().Get("title")
	subtext := c.Input().Get("subtext")
	label := c.Input().Get("label")
	principal := c.Input().Get("principal")
	relevancy := c.Input().Get("relevancy")
	content := c.Input().Get("content")
	// c.Data["Id"] = id
	// beego.Info(subtext)
	//id转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据pid查出项目id
	proj, err := models.GetProj(pidNum)
	if err != nil {
		beego.Error(err)
	}
	parentidpath := strings.Replace(strings.Replace(proj.ParentIdPath, "#$", "-", -1), "$", "", -1)
	parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
	patharray := strings.Split(parentidpath1, "-")
	topprojectid, err := strconv.ParseInt(patharray[0], 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据项目id添加成果code, title, label, principal, content string, projectid int64
	Id, err := models.AddProduct(code, title, label, principal, uid, pidNum, topprojectid)
	if err != nil {
		beego.Error(err)
	}

	//*****添加成果关联信息
	if relevancy != "" {
		_, err = models.AddRelevancy(Id, relevancy)
		if err != nil {
			beego.Error(err)
		}
	}
	//*****添加成果关联信息结束

	//成果写入postmerit表，准备提交merit*********
	Number, Name, DesignStage, Section, err := GetProjTitleNumber(pidNum)
	if err != nil {
		beego.Error(err)
	}
	catalog.ProjectNumber = Number
	catalog.ProjectName = Name
	catalog.DesignStage = DesignStage
	catalog.Section = Section

	catalog.Tnumber = code
	catalog.Name = title
	catalog.Count = 1
	catalog.Drawn = meritbasic.Nickname
	catalog.Designd = meritbasic.Nickname
	catalog.Author = meritbasic.Username
	catalog.Drawnratio = 0.4
	catalog.Designdratio = 0.4

	const lll = "2006-01-02"
	convdate := time.Now().Format(lll)
	t1, err := time.Parse(lll, convdate) //这里t1要是用t1:=就不是前面那个t1了
	if err != nil {
		beego.Error(err)
	}
	catalog.Datestring = convdate
	catalog.Date = t1

	catalog.Created = time.Now() //.Add(+time.Duration(hours) * time.Hour)
	catalog.Updated = time.Now() //.Add(+time.Duration(hours) * time.Hour)

	catalog.Complex = 1
	catalog.State = 0
	//生成提交merit的清单结束*******************

	//将文章添加到成果id下
	aid, err := models.AddArticle(subtext, content, Id)
	if err != nil {
		beego.Error(err)
	} else {
		//生成提交merit的清单*******************
		cid, err, news = models.AddPostMerit(catalog)
		if err != nil {
			beego.Error(err)
		} else {
			link1 := "/project/product/article/" + strconv.FormatInt(aid, 10) //附件链接地址
			_, err = models.AddCatalogLink(cid, link1)
			if err != nil {
				beego.Error(err)
			}
			data := news
			c.Ctx.WriteString(data)
		}
		//生成提交merit的清单结束*******************
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
	// } else {
	// route := c.Ctx.Request.URL.String()
	// c.Data["Url"] = route
	// c.Redirect("/roleerr?url="+route, 302)
	// c.Redirect("/roleerr", 302)
	// return
	// }
}

// @Title post wx artile by catalogId
// @Description post article by catalogid
// @Param title query string  true "The title of article"
// @Param content query string  true "The content of article"
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /addwxarticle [post]
//向设代日记id下添加微信小程序文章
func (c *ArticleController) AddWxArticle() {
	JSCODE := c.Input().Get("code")
	// beego.Info(JSCODE)
	APPID := "wx7f77b90a1a891d93"
	SECRET := "f58ca4f28cbb52ccd805d66118060449"
	requestUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=" + APPID + "&secret=" + SECRET + "&js_code=" + JSCODE + "&grant_type=authorization_code"
	resp, err := http.Get(requestUrl)
	if err != nil {
		beego.Error(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		beego.Error(err)
		// return
	}
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		beego.Error(err)
		// return
	}
	// beego.Info(data)
	var openID string
	// var sessionKey string
	if _, ok := data["session_key"]; !ok {
		errcode := data["errcode"]
		errmsg := data["errmsg"].(string)
		// return
		c.Data["json"] = map[string]interface{}{"errNo": errcode, "msg": errmsg, "data": "session_key 不存在"}
		c.ServeJSON()
	} else {
		// var unionId string
		openID = data["openid"].(string)
		// sessionKey = data["session_key"].(string)
		// unionId = data["unionid"].(string)
		// beego.Info(openID)
		// beego.Info(sessionKey)
	}
	// _, _, _, _, islogin := checkprodRole(c.Ctx)
	// if !islogin {
	// 	return
	// }
	//设代日记26159
	//项目id25002
	// code := c.Input().Get("code")
	pid := beego.AppConfig.String("wxcatalogid") //"26159"
	// title := c.Input().Get("title")
	// subtext := c.Input().Get("subtext")
	// label := c.Input().Get("label")
	// principal := c.Input().Get("principal")
	// relevancy := c.Input().Get("relevancy")
	// content := c.Input().Get("content")
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	content = "<p style='font-size: 18px;'>" + openID + "~" + content + "</p>" //<span style="font-size: 18px;">这个字体到底是多大才好看</span>
	imagesurl := c.Input().Get("images")
	array := strings.Split(imagesurl, ",")
	for _, v := range array {
		content = content + "<p><img src='" + v + "'></p>"
	}
	// id, err := models.AddWikiOne(title, content, "test")
	// if err != nil {
	// 	beego.Error(err)
	// 	c.Data["json"] = map[string]interface{}{"info": "ERR", "id": id}
	// 	c.ServeJSON()
	// } else {
	// 	// c.Data["json"] = id
	// 	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "id": id}
	// 	c.ServeJSON()
	// }

	//id转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据pid查出项目id
	proj, err := models.GetProj(pidNum)
	if err != nil {
		beego.Error(err)
	}
	parentidpath := strings.Replace(strings.Replace(proj.ParentIdPath, "#$", "-", -1), "$", "", -1)
	parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
	patharray := strings.Split(parentidpath1, "-")
	topprojectid, err := strconv.ParseInt(patharray[0], 10, 64)
	if err != nil {
		beego.Error(err)
	}

	code := time.Now().Format("2006-01-02 15:04")
	code = strings.Replace(code, "-", "", -1)
	code = strings.Replace(code, " ", "", -1)
	code = strings.Replace(code, ":", "", -1)
	//根据项目id添加成果code, title, label, principal, content string, projectid int64
	Id, err := models.AddProduct(code, title, "wx", "wxuser", 0, pidNum, topprojectid)
	if err != nil {
		beego.Error(err)
	}

	//将文章添加到成果id下
	aid, err := models.AddArticle(title, content, Id)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERR", "id": aid}
		c.ServeJSON()
	} else {
		// c.Data["json"] = id
		c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "id": aid}
		c.ServeJSON()
	}
}

// @Title post wx artile by catalogId
// @Description post article by catalogid
// @Param id query string  true "The id of project"
// @Param title query string  true "The title of article"
// @Param content query string  true "The content of article"
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /addwxarticles/:id [post]
//向设代日记id下添加微信小程序文章
func (c *ArticleController) AddWxArticles() {
	JSCODE := c.Input().Get("code")
	// beego.Info(JSCODE)
	APPID := "wx05b480781ac8f4c8"                //这里一定要修改啊
	SECRET := "0ddfd84afc74f3bc5b5db373a4090dc5" //这里一定要修改啊
	requestUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=" + APPID + "&secret=" + SECRET + "&js_code=" + JSCODE + "&grant_type=authorization_code"
	resp, err := http.Get(requestUrl)
	if err != nil {
		beego.Error(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		beego.Error(err)
		// return
	}
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		beego.Error(err)
		// return
	}
	// beego.Info(data)
	var openID string
	// var sessionKey string
	if _, ok := data["session_key"]; !ok {
		errcode := data["errcode"]
		errmsg := data["errmsg"].(string)
		// return
		c.Data["json"] = map[string]interface{}{"errNo": errcode, "msg": errmsg, "data": "session_key 不存在"}
		c.ServeJSON()
	} else {
		// var unionId string
		openID = data["openid"].(string)
		// sessionKey = data["session_key"].(string)
		// unionId = data["unionid"].(string)
		// beego.Info(openID)
		// beego.Info(sessionKey)
	}
	// _, _, _, _, islogin := checkprodRole(c.Ctx)
	// if !islogin {
	// 	return
	// }
	//设代日记26159
	//项目id25002
	// code := c.Input().Get("code")
	pid := c.Ctx.Input.Param(":id")
	// pid := beego.AppConfig.String("wxcatalogid") //"26159"
	// title := c.Input().Get("title")
	// subtext := c.Input().Get("subtext")
	// label := c.Input().Get("label")
	// principal := c.Input().Get("principal")
	// relevancy := c.Input().Get("relevancy")
	// content := c.Input().Get("content")
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	content = "<p style='font-size: 18px;'>" + openID + "~" + content + "</p>" //<span style="font-size: 18px;">这个字体到底是多大才好看</span>
	imagesurl := c.Input().Get("images")
	array := strings.Split(imagesurl, ",")
	for _, v := range array {
		content = content + "<p><img src='" + v + "'></p>"
	}
	// id, err := models.AddWikiOne(title, content, "test")
	// if err != nil {
	// 	beego.Error(err)
	// 	c.Data["json"] = map[string]interface{}{"info": "ERR", "id": id}
	// 	c.ServeJSON()
	// } else {
	// 	// c.Data["json"] = id
	// 	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "id": id}
	// 	c.ServeJSON()
	// }

	//id转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据pid查出项目id
	proj, err := models.GetProj(pidNum)
	if err != nil {
		beego.Error(err)
	}
	parentidpath := strings.Replace(strings.Replace(proj.ParentIdPath, "#$", "-", -1), "$", "", -1)
	parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
	patharray := strings.Split(parentidpath1, "-")
	topprojectid, err := strconv.ParseInt(patharray[0], 10, 64)
	if err != nil {
		beego.Error(err)
	}

	code := time.Now().Format("2006-01-02 15:04")
	code = strings.Replace(code, "-", "", -1)
	code = strings.Replace(code, " ", "", -1)
	code = strings.Replace(code, ":", "", -1)
	//根据项目id添加成果code, title, label, principal, content string, projectid int64
	Id, err := models.AddProduct(code, title, "wx", "wxuser", 0, pidNum, topprojectid)
	if err != nil {
		beego.Error(err)
	}

	//将文章添加到成果id下
	aid, err := models.AddArticle(title, content, Id)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERR", "id": aid}
		c.ServeJSON()
	} else {
		// c.Data["json"] = id
		c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "id": aid}
		c.ServeJSON()
	}
}

//向成果id下添加文章——这个没用，上面那个已经包含了
func (c *ArticleController) AddProdArticle() {
	// _, role := checkprodRole(c.Ctx)
	// if role == 1 {
	// id := c.Ctx.Input.Param(":id")
	pid := c.Input().Get("pid")
	subtext := c.Input().Get("subtext")
	content := c.Input().Get("content")
	//id转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//将文章添加到成果id下
	_, err = models.AddArticle(subtext, content, pidNum)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
	// } else {
	// route := c.Ctx.Request.URL.String()
	// c.Data["Url"] = route
	// c.Redirect("/roleerr?url="+route, 302)
	// // c.Redirect("/roleerr", 302)
	// return
	// }
}

//修改文章页面
func (c *ArticleController) ModifyArticle() {
	// _, _, uid, isadmin, _ := checkprodRole(c.Ctx)
	// if !isadmin {
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/roleerr?url="+route, 302)
	// 	// c.Redirect("/roleerr", 302)
	// 	return
	// }
	//这里再添加一次验证才行！！！
	id := c.Ctx.Input.Param(":id")

	var err error
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	Article, err := models.GetArticle(idNum)
	if err != nil {
		beego.Error(err)
	}
	//查出成果编号，名称和作者
	prod, err := models.GetProd(Article.ProductId)
	if err != nil {
		beego.Error(err)
	}
	c.Data["product"] = prod
	c.Data["article"] = Article
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "article_modify.tpl"
}

//编辑 成果id
func (c *ArticleController) UpdateArticle() {
	pid := c.Input().Get("aid")
	// beego.Info(aid)
	subtext := c.Input().Get("subtext")
	// beego.Info(subtext)
	// content := c.Input().Get("content")
	// content := c.Input().Get("editorValue")
	content := c.Input().Get("content")
	// beego.Info(content)
	//id转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//将文章添加到成果id下
	err = models.UpdateArticle(pidNum, subtext, content)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
		// c.Redirect("/project/product/article/"+pid, 302) //回到修改后的文章
	}
	// } else {
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/roleerr?url="+route, 302)
	// 	// c.Redirect("/roleerr", 302)
	// 	return
	// }
}

//根据文章id删除文章_没删除文章中的图片
func (c *ArticleController) DeleteArticle() {
	// _, role := checkprodRole(c.Ctx)
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role == "1" {
		// id := c.Ctx.Input.Param(":id")
		pid := c.Input().Get("pid")
		//id转成64为
		pidNum, err := strconv.ParseInt(pid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		err = models.DeleteArticle(pidNum)
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	} else {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
}
