package controllers

import (
	"encoding/json"
	"engineercms/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"time"
	// "os"
	"strconv"
	// "strings"
	// "path"
)

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
	_, role := checkprodRole(c.Ctx)
	c.Data["role"] = role
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
	c.Data["article"] = Article
	c.Data["product"] = prod
	c.TplName = "article.tpl"
}

//向某个侧栏id下添加文章
func (c *ArticleController) AddArticle() {
	meritbasic, err := models.GetMeritBasic()
	if err != nil {
		beego.Error(err)
	}
	var catalog models.PostMerit
	var news string
	var cid int64

	_, role := checkprodRole(c.Ctx)
	if role == 1 {
		// id := c.Ctx.Input.Param(":id")
		pid := c.Input().Get("pid")
		code := c.Input().Get("code")
		title := c.Input().Get("title")
		subtext := c.Input().Get("subtext")
		label := c.Input().Get("label")
		principal := c.Input().Get("principal")
		content := c.Input().Get("content")
		// c.Data["Id"] = id
		// beego.Info(subtext)
		//id转成64为
		pidNum, err := strconv.ParseInt(pid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//根据项目id添加成果code, title, label, principal, content string, projectid int64
		Id, err := models.AddProduct(code, title, label, principal, "", pidNum)
		if err != nil {
			beego.Error(err)
		}

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
	} else {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
}

//向成果id下添加文章——这个没用，上面那个已经包含了
func (c *ArticleController) AddProdArticle() {
	_, role := checkprodRole(c.Ctx)
	if role == 1 {
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
	} else {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
}

//编辑 成果id
func (c *ArticleController) UpdateArticle() {
	_, role := checkprodRole(c.Ctx)
	if role == 1 {
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
		err = models.UpdateArticle(pidNum, subtext, content)
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

//根据文章id删除文章_没删除文章中的图片
func (c *ArticleController) DeleteArticle() {
	_, role := checkprodRole(c.Ctx)
	if role == 1 {
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
