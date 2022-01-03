//提供iprole和用户登录操作功能
//成果操作
package controllers

import (
	"encoding/json"
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	// beego "github.com/beego/beego/v2/adapter"
	// "github.com/beego/beego/v2/adapter/context"
	"database/sql"
	"github.com/3xxx/flow"
	"github.com/beego/beego/v2/adapter/httplib"
	// "log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ProdController struct {
	web.Controller
}

type ProductLink struct {
	Id             int64
	Code           string
	Title          string
	Label          string
	Relevancy      []RelevancyProj
	Uid            int64
	Principal      string
	ProjectId      int64
	Content        string
	Created        time.Time
	Updated        time.Time
	Views          int64
	Pdflink        []PdfLink
	Attachmentlink []AttachmentLink
	Articlecontent []ArticleContent
	DocState       flow.DocState
	ProdDoc        models.ProductDocument
}
type PdfLink struct {
	Id        int64
	Title     string
	Link      string
	ActIndex  string
	FileSize  int64
	Downloads int64
	Created   time.Time
	Updated   time.Time
}

type AttachmentLink struct {
	Id       int64
	Title    string
	Link     string
	FileSize int64
	// Suffix    string
	Downloads int64
	Created   time.Time
	Updated   time.Time
}

type ArticleContent struct {
	Id        int64
	Title     string
	Subtext   string
	ProductId int64
	Content   string
	Link      string
	// Views   int64
	Created time.Time
	Updated time.Time
}

//后端分页的数据结构
type prodTableserver struct {
	Rows  []ProductLink `json:"rows"`
	Page  int64         `json:"page"`
	Total int64         `json:"total"` //string或int64都行！
}

//后端分页的数据结构
type prodTableserver2 struct {
	Rows  []*models.ProductAttachment `json:"rows"`
	Page  int64                       `json:"page"`
	Total int64                       `json:"total"` //string或int64都行！
}

// var db *sql.DB

// func init() {
// 	driver, connStr := "mysql", "travis@/flow?charset=utf8&parseTime=true"
// 	tdb := fatal1(sql.Open(driver, connStr)).(*sql.DB)
// 	// flow.RegisterDB(tdb)
// 	if tdb == nil {
// 		log.Fatal("given database handle is `nil`")
// 	}
// 	db = tdb
// }

//根据项目侧栏id查看这个id下的成果页面，table中的数据填充用GetProducts
//任何一级目录下都可以放成果
func (c *ProdController) GetProjProd() {
	username, role, uid, isadmin, isLogin := checkprodRole(c.Ctx)
	// if !isLogin {
	// 	c.Data["json"] = "未登陆"
	// 	c.ServeJSON()
	// 	return
	// }
	useridstring := strconv.FormatInt(uid, 10)
	id := c.Ctx.Input.Param(":id")
	//id转成64位
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	//取项目本身
	category, err := models.GetProj(idNum)
	if err != nil {
		logs.Error(err)
	}

	var topprojectid int64
	if category.ParentId != 0 { //如果不是根目录
		parentidpath := strings.Replace(strings.Replace(category.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		topprojectid, err = strconv.ParseInt(patharray[0], 10, 64)
		if err != nil {
			logs.Error(err)
		}
	} else {
		topprojectid = category.Id
	}

	projectuser, err := models.GetProjectUser(topprojectid)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(projectuser.Id)
	// beego.Info(uid)
	if uid != 0 && projectuser.Id == uid || isadmin {
		c.Data["RoleAdd"] = "true"
		c.Data["RoleNewDwg"] = "true"
		c.Data["RoleFlow"] = "true"
		c.Data["RoleUpdate"] = "true"
		c.Data["RoleDelete"] = "true"
		c.Data["RoleGet"] = "true"
	} else {
		//2.取得侧栏目录路径——路由id
		//2.1 根据id取得路由
		var projurls string
		proj, err := models.GetProj(idNum)
		if err != nil {
			logs.Error(err)
		}
		if proj.ParentId == 0 { //如果是项目根目录
			projurls = "/" + strconv.FormatInt(proj.Id, 10)
		} else {
			// projurls = "/" + strings.Replace(proj.ParentIdPath, "-", "/", -1) + "/" + strconv.FormatInt(proj.Id, 10)
			projurls = "/" + strings.Replace(strings.Replace(proj.ParentIdPath, "#", "/", -1), "$", "", -1) + strconv.FormatInt(proj.Id, 10)
		}

		if res, _ := e.Enforce(useridstring, projurls+"/", "POST", ".1"); res {
			// beego.Info("posttrue")
			c.Data["RoleAdd"] = "true"
			c.Data["RoleNewDwg"] = "true"
			c.Data["RoleFlow"] = "true"
		} else {
			c.Data["RoleAdd"] = "false"
			c.Data["RoleNewDwg"] = "false"
			c.Data["RoleFlow"] = "false"
		}
		if res2, _ := e.Enforce(useridstring, projurls+"/", "PUT", ".1"); res2 {
			c.Data["RoleUpdate"] = "true"
		} else {
			c.Data["RoleUpdate"] = "false"
		}
		if res3, _ := e.Enforce(useridstring, projurls+"/", "DELETE", ".1"); res3 {
			c.Data["RoleDelete"] = "true"
		} else {
			c.Data["RoleDelete"] = "false"
		}
		if res4, _ := e.Enforce(useridstring, projurls+"/", "GET", ".1"); res4 {
			c.Data["RoleGet"] = "true"
		} else {
			c.Data["RoleGet"] = "false"
		}
	}
	c.Data["Id"] = id
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	site := c.Ctx.Input.Site()
	port := strconv.Itoa(c.Ctx.Input.Port())
	if port == "80" {
		c.Data["Site"] = site
	} else {
		c.Data["Site"] = site + ":" + port
	}
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = isLogin
	c.Data["Uid"] = uid
	// var categories []*models.ProjCategory
	// var err error
	//id转成64为
	// idNum, err := strconv.ParseInt(id, 10, 64)
	// if err != nil {
	// 	logs.Error(err)
	// }
	//取项目本身
	// category, err := models.GetProj(idNum)
	// if err != nil {
	// 	logs.Error(err)
	// }
	//取项目所有子孙
	// categories, err := models.GetProjectsbyPid(idNum)
	// if err != nil {
	// 	logs.Error(err)
	// }
	//算出最大级数
	// grade := make([]int, 0)
	// for _, v := range categories {
	// 	grade = append(grade, v.Grade)
	// }
	// height := intmax(grade[0], grade[1:]...)

	// c.Data["json"] = root
	// c.ServeJSON()
	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	// elapsed := time.Since(start)
	// beego.Info(elapsed)
	//取出流程flow类型
	//要么等点击按钮的时候，用ajax获取
	// flowtypelist
	// flowgrouplist
	// flowaccesscontextlist
	if matched == true {
		c.TplName = "mobile/mproject_products.tpl"
	} else {
		c.TplName = "project_products.tpl"
	}
}

type RelevancyProj struct {
	Id        int64  `form:"-"`
	ProductId int64  `orm:"null"` //成果编号
	Relevancy string `orm:"null"` //关联成果编号
	ProjectId int64  //目录id
}

// @Title get projproducts...
// @Description get projproducts..
// @Param id path string  true "The id of projproducts"
// @Param searchText query string false "The searchText of projproducts"
// @Param page query string false "The page of projproducts"
// @Param limit query string false "The size of page"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /project/products/:id [get]
//取得某个侧栏id下的成果给table
//这里增加项目同步ip的获得成果，设置连接超时。
//专门做一个接口provideproducts,由
func (c *ProdController) GetProducts() {
	var err error
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	var idNum int64
	if id != "" {
		//id转成64为
		idNum, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			logs.Error(err)
		}
	}
	searchText := c.GetString("searchText")
	limit := c.GetString("limit")
	limit1, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("pageNo")
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

	//根据项目id取得所有成果
	//2019-09-21这里要修改成联合查询，得到成果下的附件和文章，以及成果的flowdocument
	//***           ******
	//***           ****
	//***           **
	products, err := models.GetProductsPage(idNum, limit1, offset, 0, searchText)
	if err != nil {
		logs.Error(err)
	}

	// products, err := models.GetProductAttachment(idNum, limit1, offset)
	// if err != nil {
	// 	logs.Error(err)
	// }

	link := make([]ProductLink, 0)
	Attachslice := make([]AttachmentLink, 0)
	Pdfslice := make([]PdfLink, 0)
	Articleslice := make([]ArticleContent, 0)
	// tx, _ := db.Begin()
	var tx *sql.Tx
	// db.Close()
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
				// attacharr[0].Suffix = path.Ext(v.FileName)
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
			articlearr[0].Content = x.Content //返回的content是空，因为真要返回内容，会很影响速度，也没必要
			articlearr[0].Link = "/project/product/article"
			Articleslice = append(Articleslice, articlearr...)
		}
		linkarr[0].Articlecontent = Articleslice
		Articleslice = make([]ArticleContent, 0)

		//取得关联
		relevancies, err := models.GetRelevancy(w.Id)
		if err != nil {
			logs.Error(err)
		}
		relevancies1 := make([]RelevancyProj, 0)
		if len(relevancies) > 0 {
			for _, tt := range relevancies {
				relevancies2 := make([]RelevancyProj, 1)
				relevancies2[0].Relevancy = tt.Relevancy.Relevancy
				relevancies2[0].ProjectId = tt.Product.ProjectId
				relevancies1 = append(relevancies1, relevancies2...)
			}
			linkarr[0].Relevancy = relevancies1
			// relevancies1 = make([]models.Relevancy, 0)
		}
		//取出和本成果编号相同的关联
		relevancyproduct, err := models.GetRelevancybyName(w.Code)
		if err != nil {
			logs.Error(err)
		}
		for _, vv := range relevancyproduct {
			relevancies2 := make([]RelevancyProj, 1)
			relevancies2[0].Relevancy = vv.Product.Code
			relevancies2[0].ProjectId = vv.Product.ProjectId
			relevancies1 = append(relevancies1, relevancies2...)
		}
		// else if len(relevancies) == 0 {
		// 	//循环所有relevancies,以,号分割，如果相等prodcode,则返回
		// 	relevancies3, err := models.GetAllRelevancies()
		// 	if err != nil {
		// 		logs.Error(err)
		// 	}
		// if len(relevancies)>0{}
		// for _, vv := range relevancies3 {
		// 	array := strings.Split(vv.Relevancy, ",")
		// 	// beego.Info(array)
		// 	for _, ww := range array {
		// 		if ww == w.Code {
		// 			relevancies2 := make([]models.Relevancy, 1)
		// 			// v.ProductId查出prodcode
		// 			prod, err := models.GetProd(vv.ProductId)
		// 			if err != nil {
		// 				logs.Error(err)
		// 			} else {
		// 				// beego.Info(ww)        //20171228
		// 				// beego.Info(prod.Code) //20171231
		// 				relevancies2[0].Relevancy = prod.Code
		// 				relevancies1 = append(relevancies1, relevancies2...)
		// 			}
		// 			break
		// 		}
		// 	}
		// }
		linkarr[0].Relevancy = relevancies1
		relevancies1 = make([]RelevancyProj, 0)

		//这里去查flow表格里文档状态
		//默认关闭flow流程，不查询成果状态
		openflow, err := web.AppConfig.String("openflow")
		if err != nil {
			logs.Error(err)
		}
		if openflow == "true" {
			proddoc, err := models.GetProductDocument(w.Id)
			if err != nil {
				logs.Error(err)
			} else {
				document, err := flow.Documents.Get(tx, flow.DocTypeID(proddoc.DocTypeId), flow.DocumentID(proddoc.DocumentId))
				if err != nil {
					logs.Error(err)
				} else {
					linkarr[0].DocState = document.State
				}
				linkarr[0].ProdDoc = proddoc
			}
		}
		link = append(link, linkarr...)
	}

	count, err := models.GetProductsCount(idNum, searchText)
	if err != nil {
		logs.Error(err)
	}
	table := prodTableserver{link, page1, count}
	// table := prodTableserver2{products, 1, 20}
	c.Data["json"] = table
	// c.Data["json"] = table //products
	c.ServeJSON()
	// c.Data["json"] = root
	// c.ServeJSON()
}

//取出项目下所有成果——这个修改
func (c *ProdController) GetProjProducts() {
	id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	c.Data["Id"] = id
	var idNum int64
	var err error
	if id != "" {
		//id转成64为
		idNum, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			logs.Error(err)
		}
	} else {

	}
	//根据项目id取得项目下所有成果
	_, products, err := models.GetProjProducts(idNum, 1)
	if err != nil {
		logs.Error(err)
	}
	//根据项目侧栏id取得项目下所有成果
	// products, err := models.GetProducts(idNum)
	// if err != nil {
	// 	logs.Error(err)
	// }

	//由proj id取得url
	// Url, _, err := GetUrlPath(idNum)
	// if err != nil {
	// 	logs.Error(err)
	// }

	link := make([]ProductLink, 0)
	Attachslice := make([]AttachmentLink, 0)
	Pdfslice := make([]PdfLink, 0)
	Articleslice := make([]ArticleContent, 0)
	for _, w := range products {
		//根据product的projid取得url
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

		//取得关联
		relevancies, err := models.GetRelevancy(w.Id)
		if err != nil {
			logs.Error(err)
		}
		relevancies1 := make([]RelevancyProj, 0)
		if len(relevancies) > 0 {
			for _, tt := range relevancies {
				relevancies2 := make([]RelevancyProj, 1)
				relevancies2[0].Relevancy = tt.Relevancy.Relevancy
				relevancies1 = append(relevancies1, relevancies2...)
			}
			linkarr[0].Relevancy = relevancies1
			relevancies1 = make([]RelevancyProj, 0)
		} else if len(relevancies) == 0 {
			//循环所有relevancies,以,号分割，如果相等prodcode,则返回
			relevancies3, err := models.GetAllRelevancies()
			if err != nil {
				logs.Error(err)
			}
			// if len(relevancies)>0{}
			for _, vv := range relevancies3 {
				array := strings.Split(vv.Relevancy, ",")
				// beego.Info(array)
				for _, ww := range array {
					if ww == w.Code {
						relevancies2 := make([]RelevancyProj, 1)
						// v.ProductId查出prodcode
						prod, err := models.GetProd(vv.ProductId)
						if err != nil {
							logs.Error(err)
						} else {
							// beego.Info(ww)        //20171228
							// beego.Info(prod.Code) //20171231
							relevancies2[0].Relevancy = prod.Code
							relevancies1 = append(relevancies1, relevancies2...)
						}
						break
					}
				}
			}
			linkarr[0].Relevancy = relevancies1
			// relevancies1 = make([]models.Relevancy, 0)
		}

		link = append(link, linkarr...)
	}

	c.Data["json"] = link //products
	c.ServeJSON()
	// c.Data["json"] = root
	// c.ServeJSON()
}

//获取项目同步ip数据
func (c *ProdController) GetsynchProducts() {
	id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	c.Data["Id"] = id
	var idNum int64
	var err error
	if id != "" {
		//id转成64为
		idNum, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			logs.Error(err)
		}
	} else {

	}
	//取目录本身
	proj, err := models.GetProj(idNum)
	if err != nil {
		logs.Error(err)
	}
	var projid int64
	//根据目录id取出项目id，以便得到同步ip
	var parentidpath, parentidpath1 string
	// array := strings.Split(proj.ParentIdPath, "-")
	if proj.ParentIdPath != "" { //如果不是根目录
		// array := strings.Split(proj.ParentIdPath, "-")
		parentidpath = strings.Replace(strings.Replace(proj.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 = strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		//pid转成64位
		projid, err = strconv.ParseInt(patharray[0], 10, 64)
		// beego.Info(projid)
		if err != nil {
			logs.Error(err)
		}
	} else {
		projid = proj.Id
		// beego.Info(projid)
	}
	// projid, err := strconv.ParseInt(array[0], 10, 64)
	// if err != nil {
	// 	logs.Error(err)
	// }

	//取得项目同步ip——进行循环ip——根据parenttitlepath和自身title查询ip上分级id——获取这个id下成果
	projsynchip, err := models.GetAdminSynchIp(projid)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(projsynchip)
	link := make([]ProductLink, 0)
	for _, v := range projsynchip {
		var productlink []ProductLink
		//	通过如下接口可以设置请求的超时时间和数据读取时间
		jsonstring, err := httplib.Get("http://"+v.SynchIp+":"+v.Port+"/project/providesynchproducts?parenttitlepath="+proj.ParentTitlePath+"&title="+proj.Title).SetTimeout(100*time.Second, 30*time.Second).String() //.ToJSON(&productlink)
		if err != nil {
			logs.Error(err)
		}
		//json字符串解析到结构体，以便进行追加
		err = json.Unmarshal([]byte(jsonstring), &productlink)
		if err != nil {
			logs.Error(err)
		}
		// beego.Info(productlink)
		link = append(link, productlink...)
	}

	c.Data["json"] = link //products
	c.ServeJSON()
}

//对外提供成果数据接口
func (c *ProdController) ProvidesynchProducts() {
	parenttitlepath := c.GetString("parenttitlepath")
	title := c.GetString("title")
	//由上面2个条件查询到目录id
	proj, err := models.GetProjbyParenttitlepath(parenttitlepath, title)
	if err != nil {
		logs.Error(err)
	}
	//用户请求的地址和端口
	site := c.Ctx.Input.Site() + ":" + strconv.Itoa(c.Ctx.Input.Port())
	c.Data["Id"] = proj.Id
	// var idNum int64
	// var err error
	// if id != "" {
	// 	//id转成64为
	// 	idNum, err := strconv.ParseInt(id, 10, 64)
	// 	if err != nil {
	// 		logs.Error(err)
	// 	}
	// } else {

	// }
	//根据项目id取得所有成果
	products, err := models.GetProducts(proj.Id)
	if err != nil {
		logs.Error(err)
	}
	//由proj id取得url
	Url, _, err := GetUrlPath(proj.Id)
	if err != nil {
		logs.Error(err)
	}

	link := make([]ProductLink, 0)
	Attachslice := make([]AttachmentLink, 0)
	Pdfslice := make([]PdfLink, 0)
	Articleslice := make([]ArticleContent, 0)
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
				attacharr[0].Link = site + Url
				Attachslice = append(Attachslice, attacharr...)
			} else if path.Ext(v.FileName) == ".pdf" || path.Ext(v.FileName) == ".PDF" {
				pdfarr := make([]PdfLink, 1)
				pdfarr[0].Id = v.Id
				pdfarr[0].Title = v.FileName
				pdfarr[0].Link = site + Url
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
			articlearr[0].Link = site + "/project/product/article"
			Articleslice = append(Articleslice, articlearr...)
		}
		linkarr[0].Articlecontent = Articleslice
		Articleslice = make([]ArticleContent, 0)
		link = append(link, linkarr...)
	}

	c.Data["json"] = link //products
	c.ServeJSON()
}

//向某个侧栏id下添加成果——这个没用，用attachment里的addattachment
func (c *ProdController) AddProduct() {
	//取得客户端用户名
	// v := c.GetSession("uname")
	// var user models.User
	// var err error
	// if v != nil {
	// 	uname := v.(string)
	// 	user, err = models.GetUserByUsername(uname)
	// 	if err != nil {
	// 		logs.Error(err)
	// 	}
	// }
	_, _, uid, _, _ := checkprodRole(c.Ctx)

	id := c.Ctx.Input.Param(":id")
	pid := c.GetString("pid")
	code := c.GetString("code")
	title := c.GetString("title")
	label := c.GetString("label")
	principal := c.GetString("principal")
	// content := c.GetString("content")
	// beego.Info(id)
	c.Data["Id"] = id
	//id转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	//根据pid查出项目id
	proj, err := models.GetProj(pidNum)
	if err != nil {
		logs.Error(err)
	}
	var topprojectid int64
	if proj.ParentIdPath != "" {
		parentidpath := strings.Replace(strings.Replace(proj.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		topprojectid, err = strconv.ParseInt(patharray[0], 10, 64)
		if err != nil {
			logs.Error(err)
		}
	} else {
		topprojectid = proj.Id
	}
	//根据id添加成果code, title, label, principal, content string, projectid int64
	_, err = models.AddProduct(code, title, label, principal, uid, pidNum, topprojectid)
	if err != nil {
		logs.Error(err)
	}
	c.Data["json"] = "ok"
	c.ServeJSON()
	// c.Data["json"] = root
	// c.ServeJSON()
}

//编辑成果信息
func (c *ProdController) UpdateProduct() {
	_, _, uid, isadmin, isLogin := checkprodRole(c.Ctx)
	if !isLogin {
		// route := c.Ctx.Request.URL.String()
		// c.Data["Url"] = route
		// c.Redirect("/roleerr?url="+route, 302)
		c.Data["json"] = "未登陆"
		c.ServeJSON()
		return
	}
	useridstring := strconv.FormatInt(uid, 10)
	id := c.GetString("pid")
	//id转成64位
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	// 根据成果id取得项目id
	product, err := models.GetProd(idNum)
	if err != nil {
		logs.Error(err)
	}
	//取项目本身
	category, err := models.GetProj(product.ProjectId)
	if err != nil {
		logs.Error(err)
	}
	//根据项目顶级id取得项目下所有成果
	var topprojectid int64
	if category.ParentId != 0 { //如果不是根目录
		parentidpath := strings.Replace(strings.Replace(category.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		topprojectid, err = strconv.ParseInt(patharray[0], 10, 64)
		if err != nil {
			logs.Error(err)
		}
	} else {
		topprojectid = category.Id
	}

	projectuser, err := models.GetProjectUser(topprojectid)
	if err != nil {
		logs.Error(err)
	}

	var UpdatePermission bool
	if projectuser.Id == uid || isadmin {
		UpdatePermission = true
	} else {
		//2.取得侧栏目录路径——路由id
		//2.1 根据id取得路由
		var projurls string
		// proj, err := models.GetProj(idNum)
		// if err != nil {
		// 	logs.Error(err)
		// }
		if category.ParentId == 0 { //如果是项目根目录
			projurls = "/" + strconv.FormatInt(category.Id, 10)
		} else {
			// projurls = "/" + strings.Replace(proj.ParentIdPath, "-", "/", -1) + "/" + strconv.FormatInt(proj.Id, 10)
			projurls = "/" + strings.Replace(strings.Replace(category.ParentIdPath, "#", "/", -1), "$", "", -1) + strconv.FormatInt(category.Id, 10)
		}
		// beego.Info(useridstring)
		if res, _ := e.Enforce(useridstring, projurls+"/", "PUT", ".1"); res {
			UpdatePermission = true
			// beego.Info(UpdatePermission)
		}
	}
	if UpdatePermission {
		code := c.GetString("code")
		title := c.GetString("title")
		label := c.GetString("label")
		principal := c.GetString("principal")
		relevancy := c.GetString("relevancy")

		//根据id添加成果code, title, label, principal, content string, projectid int64
		err = models.UpdateProduct(idNum, code, title, label, principal)
		if err != nil {
			logs.Error(err)
		}
		//*****更新关联信息
		//如果和数据库的不一致，则先删除数据库，然后再存储
		//取得关联
		relevancies, err := models.GetRelevancy(idNum)
		if err != nil {
			logs.Error(err)
		}
		var relevancy1 string
		if len(relevancies) != 0 {
			for i, w := range relevancies {
				if i == 0 {
					relevancy1 = w.Relevancy.Relevancy
				} else {
					relevancy1 = relevancy1 + "," + w.Relevancy.Relevancy
				}
			}
		}
		//*****添加成果关联信息
		if relevancy1 != relevancy && relevancy != "" {
			// 删除关联信息
			err = models.DeleteRelevancy(idNum)
			if err != nil {
				logs.Error(err)
			}
			array := strings.Split(relevancy, ",")
			for _, v := range array {
				_, err = models.AddRelevancy(idNum, v)
				if err != nil {
					logs.Error(err)
				}
			}
		}
		//*****添加成果关联信息结束
		c.Data["json"] = "ok"
		c.ServeJSON()
	} else {
		c.Data["json"] = "非管理员、非本人、未赋予权限"
		c.ServeJSON()
	}
	// } else {
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/roleerr?url="+route, 302)
	// 	// c.Redirect("/roleerr", 302)
	// 	return
	// }
}

//删除成果，包含成果里的附件。删除附件用attachment中的
func (c *ProdController) DeleteProduct() {
	_, _, uid, isadmin, isLogin := checkprodRole(c.Ctx)
	if !isLogin {
		// route := c.Ctx.Request.URL.String()
		// c.Data["Url"] = route
		// c.Redirect("/roleerr?url="+route, 302)
		c.Data["json"] = "未登陆"
		c.ServeJSON()
		return
	}
	useridstring := strconv.FormatInt(uid, 10)
	ids := c.GetString("ids")
	// beego.Info(ids)
	array := strings.Split(ids, ",")
	//id转成64位
	idNum, err := strconv.ParseInt(array[0], 10, 64)
	if err != nil {
		logs.Error(err)
	}
	//取项目本身
	category, err := models.GetProj(idNum)
	if err != nil {
		logs.Error(err)
	}
	//根据项目顶级id取得项目下所有成果
	var topprojectid int64
	if category.ParentId != 0 { //如果不是根目录
		parentidpath := strings.Replace(strings.Replace(category.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		topprojectid, err = strconv.ParseInt(patharray[0], 10, 64)
		if err != nil {
			logs.Error(err)
		}
	} else {
		topprojectid = category.Id
	}

	projectuser, err := models.GetProjectUser(topprojectid)
	if err != nil {
		logs.Error(err)
	}

	var DeletePermission bool
	if projectuser.Id == uid || isadmin {
		DeletePermission = true
	} else {
		//2.取得侧栏目录路径——路由id
		//2.1 根据id取得路由
		var projurls string
		proj, err := models.GetProj(idNum)
		if err != nil {
			logs.Error(err)
		}
		if proj.ParentId == 0 { //如果是项目根目录
			projurls = "/" + strconv.FormatInt(proj.Id, 10)
		} else {
			// projurls = "/" + strings.Replace(proj.ParentIdPath, "-", "/", -1) + "/" + strconv.FormatInt(proj.Id, 10)
			projurls = "/" + strings.Replace(strings.Replace(proj.ParentIdPath, "#", "/", -1), "$", "", -1) + strconv.FormatInt(proj.Id, 10)
		}

		if res, _ := e.Enforce(useridstring, projurls+"/", "DELETE", ".1"); res {
			DeletePermission = true
		}
	}
	if DeletePermission {
		// ids := c.GetString("ids")
		// array := strings.Split(ids, ",")
		for _, v := range array {
			// pid = strconv.FormatInt(v1, 10)
			//id转成64位
			idNum, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				logs.Error(err)
			}
			//循环删除成果
			//根据成果id取得所有附件
			attachments, err := models.GetAttachments(idNum)
			if err != nil {
				logs.Error(err)
			}
			for _, w := range attachments {
				//取得附件的成果id——再取得成果的项目目录id——再取得路径
				attach, err := models.GetAttachbyId(w.Id)
				if err != nil {
					logs.Error(err)
				}
				prod, err := models.GetProd(attach.ProductId)
				if err != nil {
					logs.Error(err)
				}
				//根据proj的id
				_, DiskDirectory, err := GetUrlPath(prod.ProjectId)
				if err != nil {
					logs.Error(err)
				} else if DiskDirectory != "" {
					path := DiskDirectory + "/" + attach.FileName
					//删除附件
					err = os.Remove(path)
					if err != nil {
						logs.Error(err)
					}
					//删除附件数据表
					err = models.DeleteAttachment(w.Id)
					if err != nil {
						logs.Error(err)
					}
				}
			}
			//删除文章，文章中的图片无法删除
			//取得成果id下所有文章
			articles, err := models.GetArticles(idNum)
			if err != nil {
				logs.Error(err)
			}
			//删除文章表
			for _, z := range articles {
				//删除文章数据表
				err = models.DeleteArticle(z.Id)
				if err != nil {
					logs.Error(err)
				}
			}
			err = models.DeleteProduct(idNum) //删除成果数据表
			if err != nil {
				logs.Error(err)
			} else {
				c.Data["json"] = "ok"
				c.ServeJSON()
			}
		}
	} else {
		c.Data["json"] = "非管理员、非本人、未赋予删除权限"
		c.ServeJSON()
	}
	// } else {
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/roleerr?url="+route, 302)
	// 	// c.Redirect("/roleerr", 302)
	// 	return
	// }
}

// {
//   "id": 33,
//   "text": "test3",
//   "nodes": [
//     {
//       "id": 34,
//       "text": "项目建议书",
//       "nodes": [
//         {
//           "id": 36,
//           "text": "综合",
//           "nodes": [
//             {
//               "id": 40,
//               "text": "设计大纲",
//               "nodes": []
//             },
//             {
//               "id": 41,
//               "text": "计算书",
//               "nodes": []
//             }
//           ]
//         },
// type FileNode struct {
// 	Name      string      `json:"name"`
// 	Path      string      `json:"path"`
// 	FileNodes []*FileNode `json:"children"`
// }

// func walkback(path string, info os.FileInfo, node *FileNode) {
// 	// 列出当前目录下的所有目录、文件
// 	files := listFiles(path)
// 	// 遍历这些文件
// 	for _, filename := range files {
// 		// 拼接全路径
// 		fpath := filepath.Join(path, filename)
// 		// 构造文件结构
// 		fio, _ := os.Lstat(fpath)
// 		// 将当前文件作为子节点添加到目录下
// 		child := FileNode{filename, fpath, []*FileNode{}}
// 		node.FileNodes = append(node.FileNodes, &child)
// 		// 如果遍历的当前文件是个目录，则进入该目录进行递归
// 		if fio.IsDir() {
// 			walk(fpath, fio, &child)
// 		}
// 	}
// 	return
// }
