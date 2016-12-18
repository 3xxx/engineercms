package controllers

import (
	// "encoding/json"
	"engineercms/models"
	"github.com/astaxie/beego"
	// "os"
	"strconv"
	// "strings"
	// "path"
)

type ArticleController struct {
	beego.Controller
}

//取得某个成果id下的文章给table
//还没改
func (c *ArticleController) GetArticles() {
	pid := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	c.Data["Id"] = pid
	var pidNum int64
	var err error
	if pid != "" {
		//id转成64为
		pidNum, err = strconv.ParseInt(pid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//由成果id（后台传过来的行id）取得侧栏目录id
		prod, err := models.GetProd(pidNum)
		if err != nil {
			beego.Error(err)
		}
		//根据成果id取得所有附件
		Articles, err := models.GetArticles(pidNum)
		if err != nil {
			beego.Error(err)
		}
		//对成果进行循环
		//赋予url
		//如果是一个成果，直接给url;如果大于1个，则是数组:这个在前端实现
		// http.ServeFile(ctx.ResponseWriter, ctx.Request, filePath)
		content := make([]ProductLink, 0)
		contentarr := make([]ProductLink, 1)
		contentarr[0].Content = prod.Content
		contentarr[0].Code = prod.Code
		contentarr[0].Title = prod.Title
		contentarr[0].Label = prod.Label
		contentarr[0].Uid = prod.Uid
		contentarr[0].Principal = prod.Principal
		contentarr[0].Views = prod.Views
		for _, v := range Articles {
			// fileext := path.Ext(v.FileName)
			contentarr[0].Id = v.Id
			contentarr[0].Created = v.Created
			contentarr[0].Updated = v.Updated
			content = append(content, contentarr...)
		}
		c.Data["json"] = content
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
	// id := c.Ctx.Input.Param(":id")
	pid := c.Input().Get("pid")
	code := c.Input().Get("code")
	title := c.Input().Get("title")
	label := c.Input().Get("label")
	principal := c.Input().Get("principal")
	content := c.Input().Get("content")
	// c.Data["Id"] = id
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
	//将文章添加到成果id下
	_, err = models.AddArticle(content, Id)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}
