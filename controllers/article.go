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
	// "net/http"
	"database/sql"
	"github.com/3xxx/flow"
	"path"
	// "regexp"
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
	if e.Enforce(useridstring, projurls+"/", "PUT", ".1") || uid == prod.Uid && islogin || isadmin {
		c.Data["RoleUpdate"] = true
	} else {
		c.Data["RoleUpdate"] = false
	}
	if e.Enforce(useridstring, projurls+"/", "DELETE", ".1") || uid == prod.Uid && islogin || isadmin {
		c.Data["RoleDelete"] = true
	} else {
		c.Data["RoleDelete"] = false
	}
	// c.Data["productid"] = prod.Uid//文章作者id
	// u := c.Ctx.Input.UserAgent()
	// matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// if matched == true {
	// beego.Info("移动端~")
	// c.TplName = "marticle.tpl"
	// } else {
	// beego.Info("电脑端！")
	c.TplName = "article.tpl"
	// }
	//根据userid取出user和avatorUrl
	useravatar, err := models.GetUserAvatorUrl(prod.Uid)
	if err != nil {
		beego.Error(err)
	}
	var photo string
	if len(useravatar) != 0 {
		photo = useravatar[0].UserAvatar.AvatarUrl
		// beego.Info(photo)
	} else {
		photo = "/static/img/go.jpg"
	}
	c.Data["article"] = Article
	c.Data["product"] = prod
	c.Data["avatar"] = photo
}

type WxArticle struct {
	Id      int64  `json:"id",form:"-"`
	Title   string `json:"title"`
	Subtext string `json:"subtext",orm:"sie(20)"`
	Author  string `json:"author"`
	// UserId      in64      `json:"userid"`
	AppreciationUrl string    `json:"appreciationurl"` //作者赞赏码
	IsArticleMe     bool      `json:"isArticleMe"`
	ImgUrl          string    `json:"imgUrl"`
	ImgUrls         []Img     `json:"imgUrls"`
	Content         string    `json:"html",orm:"sie(5000)"`
	LeassonType     int       `json:"leassonType"`
	ProductId       int64     `orm:"null"`
	Views           int64     `orm:"default(0)"`
	Created         time.Time `orm:"auto_now_add;type(datetime)"`
	Updated         string    `json:"time",orm:"auto_now_add;type(datetime)"`
	Word            string    `json:"word",orm:"sie(5000)"`
	LikeNum         int       `json:"likeNum"`
	Liked           bool      `json:"liked"`
	Comment         []Comment `json:"comment"`
	CommentNum      int       `json:"commentNum"`
	DocState        flow.DocState
	DocNotification flow.Notification
	ProdDoc         models.ProductDocument
}

//微信小程序里质量流程的单个具体文章
type WxArticleFlowDetail struct {
	Id      int64  `json:"id",form:"-"`
	Title   string `json:"title"`
	Subtext string `json:"subtext",orm:"sie(20)"`
	Author  string `json:"author"`
	// UserId      in64      `json:"userid"`
	AppreciationUrl string    `json:"appreciationurl"` //作者赞赏码
	IsArticleMe     bool      `json:"isArticleMe"`
	ImgUrl          string    `json:"imgUrl"`
	ImgUrls         []Img     `json:"imgUrls"`
	Content         string    `json:"html",orm:"sie(5000)"`
	LeassonType     int       `json:"leassonType"`
	ProductId       int64     `orm:"null"`
	Views           int64     `orm:"default(0)"`
	Created         time.Time `orm:"auto_now_add;type(datetime)"`
	Updated         string    `json:"time",orm:"auto_now_add;type(datetime)"`
	Word            string    `json:"word",orm:"sie(5000)"`
	LikeNum         int       `json:"likeNum"`
	Liked           bool      `json:"liked"`
	Comment         []Comment `json:"comment"`
	CommentNum      int       `json:"commentNum"`
	// DocState        flow.DocState
	// DocNotification flow.Notification
	// ProdDoc         models.ProductDocument

	Document *flow.Document
	Action   []flow.DocAction
	History  []*flow.DocEventsHistory
	Text     string
	//这里增加一个映射product
	Product models.Product
}

type Comment struct {
	Id           int64  `json:"id"`
	Content      string `json:"content"`
	Isme         bool   `json:"is_me",default:"false"`
	Avatar       string `json:"avatar"`
	Username     string `json:"username"`
	Publish_time string `json:"publish_time"`
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
// @Param limit query string  true "The limit of page for articles list"
// @Param skey query string  true "The skey for user"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getwxarticles [get]
// 小程序取得所有文章列表，分页_珠三角设代用
// 作废
func (c *ArticleController) GetWxArticles() {
	// id := c.Ctx.Input.Param(":id")
	id := beego.AppConfig.String("wxcatalogid") //"26159" //25002珠三角设代日记id26159
	wxsite := beego.AppConfig.String("wxreqeustsite")
	// limit := "5"
	limit := c.Input().Get("limit")
	if limit == "" {
		limit = "6"
	}
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

	// var user models.User
	// JSCODE := c.Input().Get("code")
	// if JSCODE != "" {
	// 	APPID := beego.AppConfig.String("wxAPPID")
	// 	SECRET := beego.AppConfig.String("wxSECRET")
	// 	app_version := c.Input().Get("app_version")
	// 	if app_version == "3" {
	// 		APPID = beego.AppConfig.String("wxAPPID3")
	// 		SECRET = beego.AppConfig.String("wxSECRET3")
	// 	}
	// 	// APPID := "wx7f77b90a1a891d93"
	// 	// SECRET := "f58ca4f28cbb52ccd805d66118060449"
	// 	requestUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=" + APPID + "&secret=" + SECRET + "&js_code=" + JSCODE + "&grant_type=authorization_code"
	// 	resp, err := http.Get(requestUrl)
	// 	if err != nil {
	// 		beego.Error(err)
	// 		return
	// 	}
	// 	defer resp.Body.Close()
	// 	if resp.StatusCode != 200 {
	// 		beego.Error(err)
	// 	}
	// 	var data map[string]interface{}
	// 	err = json.NewDecoder(resp.Body).Decode(&data)
	// 	if err != nil {
	// 		beego.Error(err)
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
	// 			beego.Error(err)
	// 		}
	// 	}
	// }
	// var userid int64
	// if user.Nickname != "" {
	// 	userid = user.Id
	// } else {
	// 	userid = 0
	// }
	//根据项目id取得所有成果
	//bug_如果5个成果里都没有文章，则显示文章失败；如果多个文章，会超过5个
	products, err := models.GetProductsPage(idNum, limit1, offset, 0, "")
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
// @Param limit query string  true "The limit of page for articles list"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getwxarticless/:id [get]
//小程序取得我的文章列表，分页_plus_通用_含文章状态
func (c *ArticleController) GetWxArticless() {
	id := c.Ctx.Input.Param(":id")
	// id := beego.AppConfig.String("wxcatalogid") //"26159" //25002珠三角设代日记id26159
	wxsite := beego.AppConfig.String("wxreqeustsite")
	// limit := "5"
	limit := c.Input().Get("limit")
	if limit == "" {
		limit = "6"
	}
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

	var userid int64
	openID := c.GetSession("openID")
	if openID == nil {
		userid = 0
	} else {
		user, err := models.GetUserByOpenID(openID.(string))
		if err != nil {
			beego.Error(err)
		}
		userid = user.Id
	}

	//根据项目id取得所有成果
	//bug_如果5个成果里都没有文章，则显示文章失败；如果多个文章，会超过5个
	products, err := models.GetProductsPage(idNum, limit1, offset, userid, "")
	if err != nil {
		beego.Error(err)
	}
	var tx *sql.Tx
	var document *flow.Document
	Articleslice := make([]WxArticle, 0)
	articlearr := make([]WxArticle, 1)
	for _, w := range products {
		//取得文章
		Articles, err := models.GetWxArticles(w.Id)
		if err != nil {
			beego.Error(err)
		}

		//这里去查flow表格里文档状态
		proddoc, err := models.GetProductDocument(w.Id)
		if err != nil {
			beego.Error(err)
		} else {
			document, err = flow.Documents.Get(tx, flow.DocTypeID(proddoc.DocTypeId), flow.DocumentID(proddoc.DocumentId))
			if err != nil {
				beego.Error(err)
			} else {
				articlearr[0].DocState = document.State
			}
			articlearr[0].ProdDoc = proddoc
			// linkarr[0].DocState = document.State
			// linkarr[0].ProdDoc = proddoc
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

			//取得文章所有点赞
			likes, err := models.GetAllTopicLikes(x.Id)
			if err != nil {
				beego.Error(err)
			}

			//取得文章所有评论
			comments, err := models.GetAllTopicReplies(x.Id)
			if err != nil {
				beego.Error(err)
			}

			articlearr[0].Id = x.Id
			articlearr[0].Title = w.Title
			articlearr[0].Subtext = x.Subtext
			articlearr[0].Author = w.Principal
			articlearr[0].Views = x.Views
			articlearr[0].LikeNum = len(likes)
			articlearr[0].CommentNum = len(comments)
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
// @Param page query string  true "The page for articles list"
// @Param limit query string  true "The limit of page for articles list"
// @Param dsid query string  true "The id of docstate"
// @Param dtid query string  true "The id of doctype"
// @Param acid query string  true "The id of accesscontext"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getwxarticletype [get]
//小程序根据flow文档的doctype获取登录者的文档_待提交的文档列表
func (c *ArticleController) GetWxArticleType() {
	// id := c.Ctx.Input.Param(":id")
	wxsite := beego.AppConfig.String("wxreqeustsite")
	limit := c.Input().Get("limit")
	if limit == "" {
		limit = "12"
	}
	limit1, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	page := c.Input().Get("page")
	page1, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	var offset int64
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	var user models.User
	var uID int64
	openID := c.GetSession("openID")
	if openID == nil {
		// userid = 0
		user.Username = "qin.xc"
	} else {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			beego.Error(err)
			user.Username = "qin.xc"
		}
		// userid = user.Id
	}
	//根据uid获取gid
	flowuser, err := flow.Users.GetByName(user.Username)
	if err != nil {
		beego.Error(err)
	} else {
		uID = int64(flowuser.ID)
	}
	// beego.Info(uID)
	//当前用户所在的用户组
	singletongroup, err := flow.Users.SingletonGroupOf(flow.UserID(uID))
	if err != nil {
		beego.Error(err)
	}
	// gid := c.Input().Get("gid")
	// gID, err := strconv.ParseInt(gid, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	//查询flowdoc——根据uid，acid和stateid，查出flowdocid，然后查表proddocument，
	// 得到文档productid
	dtid := c.Input().Get("dtid")
	dtID, err := strconv.ParseInt(dtid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	dsid := c.Input().Get("dsid")
	dsID, err := strconv.ParseInt(dsid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	docstate, err := flow.DocStates.Get(flow.DocStateID(dsID))
	if err != nil {
		beego.Error(err)
	}
	acid := c.Input().Get("acid")
	acID, err := strconv.ParseInt(acid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	// var tx *sql.Tx
	// var document *flow.Document
	Articleslice := make([]WxArticle, 0)
	articlearr := make([]WxArticle, 1)
	documentslistinput := flow.DocumentsListInput{
		DocTypeID:       flow.DocTypeID(dtID),       // Documents of this type are listed; required
		AccessContextID: flow.AccessContextID(acID), // Access context from within which to list; required
		GroupID:         singletongroup.ID,          // List documents created by this (singleton) group
		DocStateID:      flow.DocStateID(dsID),      // 忽略List documents currently in this state
		//CtimeStarting:   time.Now(), // List documents created after this time
		//CtimeBefore:     time.Now(), // List documents created before this time
		//TitleContains:   string,     // List documents whose title contains the given text; expensive operation
		//RootOnly:        bool,       // List only root (top-level) documents
	}
	documents, err := flow.Documents.List(&documentslistinput, offset, limit1)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"err": err, "data": "查询失败!"}
		c.ServeJSON()
	}
	for _, v := range documents {
		// beego.Info(v.ID)
		productdoc, err := models.GetDocumentProduct(int64(v.ID))
		if err != nil {
			beego.Error(err)
		} else {
			// for _, w := range products {
			product, err := models.GetProd(productdoc.ProductId)
			if err != nil {
				beego.Error(err)
			}
			//取得文章
			Articles, err := models.GetWxArticles(productdoc.ProductId)
			if err != nil {
				beego.Error(err)
			}

			articlearr[0].DocState.Name = docstate.Name //document.State
			articlearr[0].DocState.ID = docstate.ID
			articlearr[0].ProdDoc = productdoc

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

				//取得文章所有点赞
				likes, err := models.GetAllTopicLikes(x.Id)
				if err != nil {
					beego.Error(err)
				}

				//取得文章所有评论
				comments, err := models.GetAllTopicReplies(x.Id)
				if err != nil {
					beego.Error(err)
				}

				articlearr[0].Id = x.Id
				articlearr[0].Title = product.Title
				articlearr[0].Subtext = x.Subtext
				articlearr[0].Author = product.Principal
				articlearr[0].Views = x.Views
				articlearr[0].LikeNum = len(likes)
				articlearr[0].CommentNum = len(comments)
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
	}
	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "articles": Articleslice}
	c.ServeJSON()
}

//递归取得至少5篇文章——逻辑上无法实现
//仅供研究，没有用上
func getwxarticles(id, limit, offset int64, articleslice []*WxArticle, wxsite string) {
	//bug_如果5个成果里都没有文章，则显示文章失败；如果多个文章，会超过5个
	products, err := models.GetProductsPage(id, limit, offset, 0, "")
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
// @Param skey path string  true "The skey of user"
// @Success 200 {object} models.GetArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /getwxarticle/:id [get]
//根据id查看一篇微信文章
func (c *ArticleController) GetWxArticle() {
	var openID string
	var user models.User
	var err error
	var isArticleMe bool
	var wxArticle *WxArticle
	openid := c.GetSession("openID")
	if openid != nil {
		openID = openid.(string)
		user, err = models.GetUserByOpenID(openid.(string))
		if err != nil {
			beego.Error(err)
		}
	}

	id := c.Ctx.Input.Param(":id")
	if id == "" {
		return
	}
	wxsite := beego.AppConfig.String("wxreqeustsite")
	var photo string
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
		return
	}
	Article, err := models.GetArticle(idNum)
	if err != nil {
		beego.Error(err)
		return
	}
	//查出成果编号，名称和作者
	prod, err := models.GetProd(Article.ProductId)
	if err != nil {
		beego.Error(err)
	}
	var tx *sql.Tx
	var document *flow.Document

	//这里去查flow表格里文档状态
	proddoc, err := models.GetProductDocument(Article.ProductId)
	if err != nil {
		beego.Error(err)
	} else {
		document, err = flow.Documents.Get(tx, flow.DocTypeID(proddoc.DocTypeId), flow.DocumentID(proddoc.DocumentId))
		if err != nil {
			beego.Error(err)
		} else {
			wxArticle = &WxArticle{
				DocState: document.State,
			}
		}
		wxArticle = &WxArticle{
			ProdDoc: proddoc,
		}
		// linkarr[0].DocState = document.State
		// linkarr[0].ProdDoc = proddoc
	}
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
	type Duration int64
	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)
	hours := 8
	const lll = "2006-01-02 15:04"
	articletime := Article.Updated.Add(time.Duration(hours) * time.Hour).Format(lll)

	//取得文章所有点赞，检查有无自己的点赞
	likes, err := models.GetAllTopicLikes(idNum)
	if err != nil {
		beego.Error(err)
	}
	var liked = false
	for _, v1 := range likes {
		if v1.OpenID == openID {
			liked = true
		}
	}
	//取得文章所有评论，是自己的则isme为true
	comments, err := models.GetAllTopicReplies(idNum)
	if err != nil {
		beego.Error(err)
	}
	commentslice := make([]Comment, 0)
	for _, v1 := range comments {
		comment := make([]Comment, 1)
		comment[0].Id = v1.Id
		comment[0].Content = v1.Content
		comment[0].Avatar = v1.Avatar
		comment[0].Username = v1.Username
		comment[0].Publish_time = v1.Created
		if v1.OpenID == openID {
			comment[0].Isme = true
		} else {
			comment[0].Isme = false
		}
		commentslice = append(commentslice, comment...)
	}

	if prod.Uid == user.Id { //20191122修改
		isArticleMe = true
	}
	//根据prod.Uid，查询作者的赞赏码
	userappreciationurl, err := models.GetUserAppreciationUrl(prod.Uid)
	if err != nil {
		beego.Error(err)
	}
	var userappreciationphoto string
	if len(userappreciationurl) != 0 {
		wxsite := beego.AppConfig.String("wxreqeustsite")
		userappreciationphoto = wxsite + userappreciationurl[0].UserAppreciation.AppreciationUrl
	}
	// beego.Info(isArticleMe)
	wxArticle = &WxArticle{
		Id:              Article.Id,
		Title:           prod.Title,
		Subtext:         Article.Subtext,
		Author:          prod.Principal,
		AppreciationUrl: userappreciationphoto,
		// UserId:      prod.Uid,
		IsArticleMe: isArticleMe,
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
		LikeNum:    len(likes),
		Liked:      liked,
		Comment:    commentslice,
		CommentNum: len(comments),
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
	c.Data["json"] = wxArticle
	c.ServeJSON()
}

// @Title get wx artile by articleId
// @Description get article by articleid
// @Param id path string  true "The id of article"
// @Param skey path string  true "The skey of user"
// @Param dtid query string  true "The id of doctype"
// @Param docid query string  true "The id of doc"
// @Param acid query string  true "The id of accesscontext"
// @Success 200 {object} models.GetArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /getwxarticleflow/:id [get]
//根据id查看一篇微信文章，带action的权限判断
func (c *ArticleController) GetWxArticleFlow() {
	var tx *sql.Tx
	var openID string
	var user models.User
	var err error
	var isArticleMe bool
	var articledetail *WxArticleFlowDetail
	openid := c.GetSession("openID")
	if openid != nil {
		openID = openid.(string)
		user, err = models.GetUserByOpenID(openid.(string))
		if err != nil {
			beego.Error(err)
		}
	}

	//查出accessid和uid获取groupid
	//查出登录用户所在group
	// var uID int64
	flowuser, err := flow.Users.GetByName(user.Username)
	if err != nil {
		beego.Error(err)
	}
	// else {
	// 	uID = int64(flowuser.ID)
	// }
	// beego.Info(uID)
	//当前用户所在的用户组
	// singletongroup, err := flow.Users.SingletonGroupOf(flow.UserID(uID))
	// if err != nil {
	// 	beego.Error(err)
	// }

	acid := c.Input().Get("acid")
	acID, err := strconv.ParseInt(acid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	dtid := c.Input().Get("dtid")
	dtID, err := strconv.ParseInt(dtid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	docid := c.Input().Get("docid")
	docID, err := strconv.ParseInt(docid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	document, err := flow.Documents.Get(tx, flow.DocTypeID(dtID), flow.DocumentID(docID))
	if err != nil {
		beego.Error(err)
	}
	//列出actions
	//列出符合条件的actions
	TransitionMap, err := flow.DocTypes.Transitions(flow.DocTypeID(dtID), document.State.ID)
	if err != nil {
		beego.Error(err)
	}

	//数组模式
	//这里增加判断用户有无这个action的权利grouphaspermission(ac,gid,dtid,daid)
	articleaction := make([]flow.DocAction, 0)
	if _, ok := TransitionMap[document.State.ID]; ok {
		//存在
		for _, value := range TransitionMap[document.State.ID].Transitions {
			// beego.Info(key)
			// beego.Info(value.Upon.ID)
			beego.Info(flowuser.ID)
			beego.Info(flow.DocTypeID(dtID))
			beego.Info(value.Upon.ID)
			userhaspermission, err := flow.AccessContexts.UserHasPermission(flow.AccessContextID(acID), flowuser.ID, flow.DocTypeID(dtID), value.Upon.ID)
			if err != nil {
				beego.Error(err)
			}
			beego.Info(userhaspermission)
			if userhaspermission {
				articledetailarr := make([]flow.DocAction, 1)
				articledetailarr[0].ID = value.Upon.ID
				articledetailarr[0].Name = value.Upon.Name
				articleaction = append(articleaction, articledetailarr...)
			}
		}
	}
	// articledetail[0].Document = document
	//查出历史记录
	doceventshistory, err := flow.DocEvents.DocEventsHistory(flow.DocTypeID(dtID), flow.DocumentID(docID))
	if err != nil {
		beego.Error(err)
	}
	// articledetail[0].History = doceventshistory

	id := c.Ctx.Input.Param(":id")
	wxsite := beego.AppConfig.String("wxreqeustsite")
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

	//这里去查flow表格里文档状态
	// proddoc, err := models.GetProductDocument(Article.ProductId)
	// if err != nil {
	// 	beego.Error(err)
	// } else {
	// 	document, err = flow.Documents.Get(tx, flow.DocTypeID(proddoc.DocTypeId), flow.DocumentID(proddoc.DocumentId))
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// 	wxArticle = &WxArticle{
	// 		DocState: document.State,
	// 		ProdDoc:  proddoc,
	// 	}
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
	//第一个图片 wxsite +
	if len(slice2) > 0 {
		photo = slice2[0].Src
	} else {
		photo = wxsite + "/static/img/go.jpg"
	}

	content := strings.Replace(Article.Content, "/attachment/", wxsite+"/attachment/", -1)
	type Duration int64
	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)
	hours := 8
	const lll = "2006-01-02 15:04"
	articletime := Article.Updated.Add(time.Duration(hours) * time.Hour).Format(lll)

	//取得文章所有点赞，检查有无自己的点赞
	likes, err := models.GetAllTopicLikes(idNum)
	if err != nil {
		beego.Error(err)
	}
	var liked = false
	for _, v1 := range likes {
		if v1.OpenID == openID {
			liked = true
		}
	}
	//取得文章所有评论，是自己的则isme为true
	comments, err := models.GetAllTopicReplies(idNum)
	if err != nil {
		beego.Error(err)
	}
	commentslice := make([]Comment, 0)
	for _, v1 := range comments {
		comment := make([]Comment, 1)
		comment[0].Id = v1.Id
		comment[0].Content = v1.Content
		comment[0].Avatar = v1.Avatar
		comment[0].Username = v1.Username
		comment[0].Publish_time = v1.Created
		if v1.OpenID == openID {
			comment[0].Isme = true
		} else {
			comment[0].Isme = false
		}
		commentslice = append(commentslice, comment...)
	}

	if prod.Uid == user.Id { //20191122修改
		isArticleMe = true
	}
	//根据prod.Uid，查询作者的赞赏码
	userappreciationurl, err := models.GetUserAppreciationUrl(prod.Uid)
	if err != nil {
		beego.Error(err)
	}
	var userappreciationphoto string
	if len(userappreciationurl) != 0 {
		wxsite := beego.AppConfig.String("wxreqeustsite")
		userappreciationphoto = wxsite + userappreciationurl[0].UserAppreciation.AppreciationUrl
	}

	articledetail = &WxArticleFlowDetail{
		Id:              Article.Id,
		Title:           prod.Title,
		Subtext:         Article.Subtext,
		Author:          prod.Principal,
		AppreciationUrl: userappreciationphoto,
		// UserId:      prod.Uid,
		IsArticleMe: isArticleMe,
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
		LikeNum:    len(likes),
		Liked:      liked,
		Comment:    commentslice,
		CommentNum: len(comments),
		Document:   document,
		Action:     articleaction,
		History:    doceventshistory,
		// Text     string
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
	c.Data["json"] = articledetail
	c.ServeJSON()
}

// @Title get wx artiles list
// @Description get articles by page
// @Param page query string  true "The page for articles list"
// @Param limit query string  true "The limit of page for articles list"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getlistarticles [get]
// 小程序取得所有文章列表，分页_首页用
// 这个应该要作废，用在哪里？
func (c *ArticleController) GetListArticles() {
	// id := c.Ctx.Input.Param(":id")
	id := beego.AppConfig.String("wxcatalogid") //"26159" //25002珠三角设代日记id26159
	// wxsite := beego.AppConfig.String("wxreqeustsite")
	// limit := "6"
	limit := c.Input().Get("limit")
	if limit == "" {
		limit = "6"
	}
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
	var user models.User
	var userid int64
	if user.Nickname != "" {
		userid = user.Id
	} else {
		userid = 0
	}
	//根据项目id取得所有成果
	//bug_如果5个成果里都没有文章，则显示文章失败；如果多个文章，会超过5个
	products, err := models.GetProductsPage(idNum, limit1, offset, userid, "")
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
				articlearr[0].ImgUrl = slice2[0].Src
			} else {
				articlearr[0].ImgUrl = "/static/img/go.jpg"
			}
			articlearr[0].Content = x.Content
			articlearr[0].LeassonType = 1
			articlearr[0].ProductId = x.ProductId
			Articleslice = append(Articleslice, articlearr...)
		}
	}
	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "articles": Articleslice}
	c.ServeJSON()
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
	var topprojectid int64
	if proj.ParentIdPath != "" {
		parentidpath := strings.Replace(strings.Replace(proj.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		topprojectid, err = strconv.ParseInt(patharray[0], 10, 64)
		if err != nil {
			beego.Error(err)
		}
	} else {
		topprojectid = proj.Id
	}
	//根据项目id添加成果code, title, label, principal, content string, projectid int64
	Id, err := models.AddProduct(code, title, label, principal, uid, pidNum, topprojectid)
	if err != nil {
		beego.Error(err)
	}

	//*****添加成果关联信息
	if relevancy != "" {
		array := strings.Split(relevancy, ",")
		for _, v := range array {
			_, err = models.AddRelevancy(Id, v)
			if err != nil {
				beego.Error(err)
			}
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
	// c.Redirect("/roleerr?url="+route, 301)
	// c.Redirect("/roleerr", 301)
	// return
	// }
}

// @Title post wx artile by catalogId
// @Description post article by catalogid
// @Param title query string true "The title of article"
// @Param content query string true "The content of article"
// @Param skey query string false "The skey of user"
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /addwxarticle [post]
//作废！向设代日记id下添加微信小程序文章_珠三角设代plus用_
//这个是文字图片分开方式，用下面这个
func (c *ArticleController) AddWxArticle() {
	// var userrole string
	// var user models.User
	// var err error
	// var iprole int
	// if v != nil { //如果登录了
	// 	islogin = true
	// 	uname = v.(string)
	// JSCODE := c.Input().Get("code")
	// // beego.Info(JSCODE)
	// APPID := beego.AppConfig.String("wxAPPID")
	// SECRET := beego.AppConfig.String("wxSECRET")
	// // APPID := "wx7f77b90a1a891d93"
	// // SECRET := "f58ca4f28cbb52ccd805d66118060449"
	// requestUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=" + APPID + "&secret=" + SECRET + "&js_code=" + JSCODE + "&grant_type=authorization_code"
	// resp, err := http.Get(requestUrl)
	// if err != nil {
	// 	beego.Error(err)
	// 	// return
	// }
	// defer resp.Body.Close()
	// if resp.StatusCode != 200 {
	// 	beego.Error(err)
	// 	// return
	// }
	// var data map[string]interface{}
	// err = json.NewDecoder(resp.Body).Decode(&data)
	// if err != nil {
	// 	beego.Error(err)
	// 	// return
	// }
	// // beego.Info(data)
	// var openID string
	var user models.User
	var err error
	openID := c.GetSession("openID")
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			beego.Error(err)
		}
	}
	// // var sessionKey string
	// if _, ok := data["session_key"]; !ok {
	// 	errcode := data["errcode"]
	// 	errmsg := data["errmsg"].(string)
	// 	// return
	// 	c.Data["json"] = map[string]interface{}{"errNo": errcode, "msg": errmsg, "data": "session_key 不存在"}
	// 	c.ServeJSON()
	// } else {
	// 	// var unionId string
	// 	openID = data["openid"].(string)
	pid := beego.AppConfig.String("wxcatalogid") //"26159"
	// pid := c.Ctx.Input.Param(":id")
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	content = "<p style='font-size: 18px;'>" + content + "</p>" //<span style="font-size: 18px;">这个字体到底是多大才好看</span>
	imagesurl := c.Input().Get("images")
	array := strings.Split(imagesurl, ",")
	for _, v := range array {
		content = content + "<p><img src='" + v + "'></p>"
	}
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
	var topprojectid int64
	if proj.ParentIdPath != "" {
		parentidpath := strings.Replace(strings.Replace(proj.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		topprojectid, err = strconv.ParseInt(patharray[0], 10, 64)
		if err != nil {
			beego.Error(err)
		}
	} else {
		topprojectid = proj.Id
	}
	code := time.Now().Format("2006-01-02 15:04")
	code = strings.Replace(code, "-", "", -1)
	code = strings.Replace(code, " ", "", -1)
	code = strings.Replace(code, ":", "", -1)
	//根据项目id添加成果code, title, label, principal, content string, projectid int64
	Id, err := models.AddProduct(code, title, "wx", user.Nickname, user.Id, pidNum, topprojectid)
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
// @Param title query string true "The title of article"
// @Param content query string true "The content of article"
// @Param skey query string false "The skey of user"
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /addwxeditorarticle [post]
// 向设代日记id下添加微信小程序文章_珠三角设代plus用_editor方式
// 作废
func (c *ArticleController) AddWxEditorArticle() {
	var user models.User
	var err error
	openID := c.GetSession("openID")
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			beego.Error(err)
		}
	} else {
		c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		c.ServeJSON()
		return
	}

	pid := beego.AppConfig.String("wxcatalogid") //"26159"
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	// content = "<p style='font-size: 18px;'>" + content + "</p>" //<span style="font-size: 18px;">这个字体到底是多大才好看</span>
	// imagesurl := c.Input().Get("images")
	// array := strings.Split(imagesurl, ",")
	// for _, v := range array {
	// 	content = content + "<p><img src='" + v + "'></p>"
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
	var topprojectid int64
	if proj.ParentIdPath != "" {
		parentidpath := strings.Replace(strings.Replace(proj.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		topprojectid, err = strconv.ParseInt(patharray[0], 10, 64)
		if err != nil {
			beego.Error(err)
		}
	} else {
		topprojectid = proj.Id
	}
	code := time.Now().Format("2006-01-02 15:04")
	code = strings.Replace(code, "-", "", -1)
	code = strings.Replace(code, " ", "", -1)
	code = strings.Replace(code, ":", "", -1)
	//根据项目id添加成果code, title, label, principal, content string, projectid int64
	Id, err := models.AddProduct(code, title, "wx", user.Nickname, user.Id, pidNum, topprojectid)
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

// @Title post wx artile by articleid
// @Description post article by articleid
// @Param id query string true "The id of article"
// @Param title query string true "The title of article"
// @Param content query string true "The content of article"
// @Param skey query string false "The skey of user"
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /updatewxeditorarticle [post]
// 编辑设代日记id下微信小程序文章_珠三角设代plus用_editor方式
func (c *ArticleController) UpdateWxEditorArticle() {
	// pid := beego.AppConfig.String("wxcatalogid") //"26159"
	//hotqinsessionid携带过来后，用下面的方法获取用户登录存储在服务端的session
	openid := c.GetSession("openID")
	if openid == nil {
		return
	}

	id := c.Input().Get("id")
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	//将content中的http://ip/去掉
	wxsite := beego.AppConfig.String("wxreqeustsite")
	content = strings.Replace(content, wxsite, "", -1)
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	//取得文章
	Article, err := models.GetArticle(idNum)
	if err != nil {
		beego.Error(err)
	} else {
		// 更新成果名称
		err = models.UpdateProductTtile(Article.ProductId, title)
		if err != nil {
			beego.Error(err)
		}
		//更新文章
		err = models.UpdateArticle(idNum, title, content)
		if err != nil {
			beego.Error(err)
			c.Data["json"] = map[string]interface{}{"info": "ERR", "id": id}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "id": id}
			c.ServeJSON()
		}
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
// 向设代日记id下添加微信小程序文章——通用：包括青少儿书画用
func (c *ArticleController) AddWxArticles() {
	var user models.User
	var err error
	openID := c.GetSession("openID")
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			beego.Error(err)
		}
	} else {
		c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		c.ServeJSON()
		return //本地调试的时候，由于小程序无法登陆服务器，所以这里要注释掉。
	}

	pid := c.Ctx.Input.Param(":id")

	title := c.Input().Get("title")
	content := c.Input().Get("content")
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
	var topprojectid int64
	if proj.ParentIdPath != "" {
		parentidpath := strings.Replace(strings.Replace(proj.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		topprojectid, err = strconv.ParseInt(patharray[0], 10, 64)
		if err != nil {
			beego.Error(err)
		}
	} else {
		topprojectid = proj.Id
	}
	code := time.Now().Format("2006-01-02 15:04")
	code = strings.Replace(code, "-", "", -1)
	code = strings.Replace(code, " ", "", -1)
	code = strings.Replace(code, ":", "", -1)
	//根据项目id添加成果code, title, label, principal, content string, projectid int64
	Id, err := models.AddProduct(code, title, "wx", user.Nickname, user.Id, pidNum, topprojectid)
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
// @Param id path string  true "The id of project"
// @Param title query string  true "The title of article"
// @Param content query string  true "The content of article"
// @Param content query string  true "The content of article"
// @Param dtid query string  true "The id of doctype"
// @Param acid query string  true "The id of accesscontext"
// @Param gid query string  true "The id of group"
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /addwxarticleflow/:id [post]
// 添加微信文章并执初始化一个流程
func (c *ArticleController) AddWxArticleFlow() {
	var user models.User
	var err error
	openID := c.GetSession("openID")
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			beego.Error(err)
		}
	} else {
		// c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		// c.ServeJSON()
		// return //本地调试的时候，由于小程序无法登陆服务器，所以这里要注释掉。
		user.Id = 5
		user.Username = "qin.xc"
	}

	pid := c.Ctx.Input.Param(":id")

	title := c.Input().Get("title")
	content := c.Input().Get("content")
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
	var topprojectid int64
	if proj.ParentIdPath != "" { //如果不是根目录
		parentidpath := strings.Replace(strings.Replace(proj.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		topprojectid, err = strconv.ParseInt(patharray[0], 10, 64)
		if err != nil {
			beego.Error(err)
		}
	} else {
		topprojectid = proj.Id
	}
	code := time.Now().Format("2006-01-02 15:04")
	code = strings.Replace(code, "-", "", -1)
	code = strings.Replace(code, " ", "", -1)
	code = strings.Replace(code, ":", "", -1)
	//根据项目id添加成果code, title, label, principal, content string, projectid int64
	Id, err := models.AddProduct(code, title, "wx", user.Nickname, user.Id, pidNum, topprojectid)
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
		dtid := c.Input().Get("dtid")
		dtID, err := strconv.ParseInt(dtid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		acid := c.Input().Get("acid")
		acID, err := strconv.ParseInt(acid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		gid := c.Input().Get("gid")
		gID, err := strconv.ParseInt(gid, 10, 64)
		if err != nil {
			beego.Error(err)
		}

		// err = wxFlowNext(dtID, daID, docID, messageID, groupIds, user)
		docid, err := wxFlowDoc(dtID, acID, gID, Id, title)
		if err != nil {
			beego.Error(err)
			c.Data["json"] = map[string]interface{}{"info": "ERR", "err": err, "data": "用户未登录!"}
			c.ServeJSON()
		}
		// c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "err": nil, "data": "写入成功!"}
		// c.ServeJSON()
		c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "err": nil, "id": aid, "docid": docid, "dtid": dtid}
		c.ServeJSON()
	}
}

// 提供给上面addwxarticleflow里用的初始化一个流程
func wxFlowDoc(dtID, acID, gID, prodID int64, articletitle string) (docid int64, err error) {
	// c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	var tx *sql.Tx
	//查询预先定义的doctype流程类型
	// dtid := c.Input().Get("dtid")
	// dtID, err := strconv.ParseInt(dtid, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// acid := c.Input().Get("acid")
	// acID, err := strconv.ParseInt(acid, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// gid := c.Input().Get("gid")
	// gID, err := strconv.ParseInt(gid, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// name := c.Input().Get("docname")
	//与flowdoc接口的区别就是由文章id取得prodid
	// data := c.Input().Get("docdata")
	// //docid是文章id，要取得成果productid,可以是数组
	// articleid, err := strconv.ParseInt(data, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// article, err := models.GetArticle(articleid)
	// if err != nil {
	// 	beego.Error(err)
	// }
	//prodID转string
	data := strconv.FormatInt(prodID, 10)
	docNewInput := flow.DocumentsNewInput{
		DocTypeID:       flow.DocTypeID(dtID),       //属于图纸设计类型的流程
		AccessContextID: flow.AccessContextID(acID), //所有用户权限符合这个contex的要求
		GroupID:         flow.GroupID(gID),          //groupId,初始状态下的用户组，必须是个人用户组（一个用户也可以成为一个独特的组，因为用户无法赋予角色，所以必须将用户放到组里）
		Title:           articletitle,               //这个文件的名称
		Data:            data,                       //文件的描述
	}
	// flow.Documents.New(tx, &docNewInput)
	documentid, err := flow.Documents.New(tx, &docNewInput)
	if err != nil {
		beego.Error(err)
		return 0, err
		// c.Data["json"] = map[string]interface{}{"info": "ERR", "err": err, "data": "写入失败!"}
		// c.ServeJSON()
	} else {
		//可以传另外一个参数productid过来，然后在这里写入productdocument表格
		//DocTypeID
		_, err = models.AddProductDocument(dtID, int64(documentid), prodID)
		if err != nil {
			beego.Error(err)
			return 0, err
		}
		return int64(documentid), err
		// c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "err": "ok", "data": "写入成功!"}
		// c.ServeJSON()
	}
}

// 下面这个作废，用上面那个flowdoc添加一个初始文档流程
func wxFlowNext(dtID, daID, docID, messageID int64, groupIds []flow.GroupID, user models.User) error {
	// c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", c.Ctx.Request.Header.Get("Origin"))
	var tx *sql.Tx //用这个nil，后面就不用commit了吧，都在flow里commit了。
	var tx1 *sql.Tx
	// var user models.User
	// var err error
	// openid := c.GetSession("openID")
	// if openid != nil {
	// 	user, err = models.GetUserByOpenID(openid.(string))
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	var uID int64
	flowuser, err := flow.Users.GetByName(user.Username) //user.Username
	if err != nil {
		beego.Error(err)
		// uID = 5
	} else {
		uID = int64(flowuser.ID)
	}

	// dtid := c.Input().Get("dtid")
	// dtID, err := strconv.ParseInt(dtid, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// daid := c.Input().Get("daid")
	// daID, err := strconv.ParseInt(daid, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// docid := c.Input().Get("docid")
	// docID, err := strconv.ParseInt(docid, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// messageid := c.Input().Get("messageid")
	// messageID, err := strconv.ParseInt(messageid, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }

	//根据docid取得document
	document, err := flow.Documents.Get(tx, flow.DocTypeID(dtID), flow.DocumentID(docID))
	if err != nil {
		beego.Error(err)
	}
	//根据document取得workflow
	myWorkflow, err := flow.Workflows.GetByDocType(document.DocType.ID)
	if err != nil {
		beego.Error(err)
	}

	//当前用户所在的用户组
	singletongroup, err := flow.Users.SingletonGroupOf(flow.UserID(uID))
	if err != nil {
		beego.Error(err)
	}
	//接受用户组
	// gid := make([]string, 0, 2)
	// c.Ctx.Input.Bind(&gid, "gid")

	// gid := make([]string, 1, 2)
	// gid[0] = c.Input().Get("gid")
	// var groupIds []flow.GroupID
	// for _, v := range gid {
	// 	gID, err := strconv.ParseInt(v, 10, 64)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// 	groupIds = append(groupIds, flow.GroupID(gID))
	// }
	// text := c.Input().Get("text")
	// if text == "" {
	text := "no comments"
	// }
	//建立event
	docEventInput := flow.DocEventsNewInput{
		DocTypeID:   flow.DocTypeID(dtID),
		DocumentID:  flow.DocumentID(docID),
		DocStateID:  document.State.ID,      //document state must be this state，文档的现状状态
		DocActionID: flow.DocActionID(daID), //Action performed by `Group`; required,由用户组执行的操作
		GroupID:     singletongroup.ID,      //Group (user) who performed the action that raised this event; required，执行引发此事件的操作的组(用户)
		Text:        text,                   //Any comments or notes; required，
	}
	deID, err := flow.DocEvents.New(tx, &docEventInput)
	if err != nil {
		beego.Error(err)
	}

	myDocEvent, err := flow.DocEvents.Get(flow.DocEventID(deID))
	if err != nil {
		beego.Error(err)
	} else {
		beego.Info(myDocEvent)
	}
	//这里要将邮箱对应的信息改为已读unread改为false
	newDocStateId, err := myWorkflow.ApplyEvent(tx1, myDocEvent, groupIds)
	if err != nil {
		return err
		beego.Error(err)
		// c.Data["json"] = map[string]interface{}{"info": "ERR", "err": err, "data": "写入失败!"}
		// c.ServeJSON()
	} else {
		//修改邮件为已读，即——已处理
		err = flow.Mailboxes.SetStatusByUser(tx, flowuser.ID, flow.MessageID(messageID), false)
		if err != nil {
			return err
			beego.Error(err)
		}
		beego.Info("newDocStateId=", newDocStateId, err)
		return nil
		// c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "err": nil, "data": "写入成功!"}
		// c.ServeJSON()
	}
	return err
	// } else {
	// 	c.Data["json"] = map[string]interface{}{"info": "ERR", "err": err, "data": "用户未登录!"}
	// 	c.ServeJSON()
	// }
}

// 向成果id下添加文章——这个没用，上面那个已经包含了
func (c *ArticleController) AddProdArticle() {
	_, _, _, _, islogin := checkprodRole(c.Ctx)
	if !islogin {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 301)
		// c.Redirect("/roleerr", 301)
		return
	}
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
	// c.Redirect("/roleerr?url="+route, 301)
	// // c.Redirect("/roleerr", 301)
	// return
	// }
}

// 修改文章页面
func (c *ArticleController) ModifyArticle() {
	// _, _, uid, isadmin, _ := checkprodRole(c.Ctx)
	// if !isadmin {
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/roleerr?url="+route, 301)
	// 	// c.Redirect("/roleerr", 301)
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

// 编辑 成果id
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
		// c.Redirect("/project/product/article/"+pid, 301) //回到修改后的文章
	}
	// } else {
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/roleerr?url="+route, 301)
	// 	// c.Redirect("/roleerr", 301)
	// 	return
	// }
}

// 根据文章id删除文章_没删除文章中的图片
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
		c.Redirect("/roleerr?url="+route, 301)
		// c.Redirect("/roleerr", 301)
		return
	}
}

// @Title post wx artile by articleId
// @Description post article by catalogid
// @Param id query string  true "The id of article"
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /deletewxarticle [post]
// 根据文章id删除文章_没删除文章中的图片
func (c *ArticleController) DeleteWxArticle() {
	var openID string
	openid := c.GetSession("openID")
	if openid == nil {
		c.Data["json"] = "没有注册，未查到openid"
		c.ServeJSON()
	} else {
		openID = openid.(string)
		user, err := models.GetUserByOpenID(openID)
		if err != nil {
			beego.Error(err)
			c.Data["json"] = "未查到openid对应的用户"
			c.ServeJSON()
		} else {
			//判断是否具备admin角色
			role, err := models.GetRoleByRolename("admin")
			if err != nil {
				beego.Error(err)
			}
			uid := strconv.FormatInt(user.Id, 10)
			roleid := strconv.FormatInt(role.Id, 10)
			if e.HasRoleForUser(uid, "role_"+roleid) {
				id := c.Input().Get("id")
				//id转成64为
				idNum, err := strconv.ParseInt(id, 10, 64)
				if err != nil {
					beego.Error(err)
				}
				err = models.DeleteArticle(idNum)
				if err != nil {
					beego.Error(err)
					c.Data["json"] = "delete wrong"
					c.ServeJSON()
				} else {
					c.Data["json"] = "ok"
					c.ServeJSON()
				}
			} else {
				c.Data["json"] = "不是管理员"
				c.ServeJSON()
			}
		}
	}
}

// @Title get wx userarticles count
// @Description get userarticles by projid
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articles not found
// @router /getuserarticle [get]
//小程序取得我的文章列表，分页_plus_通用_含文章状态
func (c *ArticleController) GetUserArticle() {
	c.TplName = "article_pie.tpl"
}

// @Title get wx userarticles count
// @Description get userarticles by projid
// @Param id path string  true "The id of project"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articles not found
// @router /getwxuserarticles/:id [get]
//小程序取得我的文章列表，分页_plus_通用_含文章状态
func (c *ArticleController) GetWxUserArticles() {
	id := c.Ctx.Input.Param(":id")
	// var idNum int64
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	// var userid int64
	// openID := c.GetSession("openID")
	// if openID == nil {
	// 	userid = 0
	// } else {
	// 	user, err := models.GetUserByOpenID(openID.(string))
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// 	userid = user.Id
	// }
	userarticles, err := models.GetWxUserArticles(idNum)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "data": userarticles}
	c.ServeJSON()
}
