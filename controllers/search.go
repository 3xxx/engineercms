package controllers

import (
	"engineercms/models"
	"github.com/astaxie/beego"
	"path"
	"strconv"
)

type SearchController struct {
	beego.Controller
}

//显示搜索也
//用模态框弹窗吧_比较麻烦
func (c *SearchController) Get() { //search用的是get方法
	pid := c.Input().Get("productid")
	key := c.Input().Get("keyword")
	if pid != "" {
		c.Data["IsProduct"] = true
	} else {
		c.Data["IsProject"] = true
	}
	c.Data["Pid"] = pid
	c.Data["Key"] = key
	c.TplName = "searchs.tpl"
}

//搜索项目
func (c *SearchController) SearchProject() { //search用的是get方法
	key := c.Input().Get("keyword")
	if key != "" {
		searchs, err := models.SearchProject(key)
		if err != nil {
			beego.Error(err.Error)
		} else {
			c.Data["json"] = searchs
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = "关键字为空！"
		c.ServeJSON()
	}
}

//首页里的搜索所有成果
func (c *SearchController) SearchProduct() { //search用的是get方法
	key := c.Input().Get("keyword")
	if key != "" {
		products, err := models.SearchProduct(key)
		if err != nil {
			beego.Error(err.Error)
		}
		//由product取得proj
		//取目录本身
		// proj, err := models.GetProj(products.ProjectId)
		// if err != nil {
		// 	beego.Error(err)
		// }
		//根据目录id取出项目id，以便得到同步ip
		// array := strings.Split(proj.ParentIdPath, "-")
		// projid, err := strconv.ParseInt(array[0], 10, 64)
		// if err != nil {
		// 	beego.Error(err)
		// }
		//由proj id取得url

		// beego.Info(Url)
		link := make([]ProductLink, 0)
		Attachslice := make([]AttachmentLink, 0)
		Pdfslice := make([]PdfLink, 0)
		Articleslice := make([]ArticleContent, 0)
		for _, w := range products {
			// Url, _, err := GetUrlPath(w.ProjectId)
			// if err != nil {
			// 	beego.Error(err)
			// }
			//取到每个成果的附件（模态框打开）；pdf、文章——新窗口打开
			//循环成果
			//每个成果取到所有附件
			//一个附件则直接打开/下载；2个以上则打开模态框
			Attachments, err := models.GetAttachments(w.Id)
			if err != nil {
				beego.Error(err)
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
			linkarr[0].Content = w.Content
			linkarr[0].Created = w.Created
			linkarr[0].Updated = w.Updated
			linkarr[0].Views = w.Views
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
				beego.Error(err)
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
		c.Data["json"] = link
		c.ServeJSON()
	} else {
		c.Data["json"] = "关键字为空！"
		c.ServeJSON()
	}
}

//在某个项目里搜索成果：全文搜索，article全文，编号，名称，关键字，作者……
func (c *SearchController) SearchProjProducts() {
	pid := c.Input().Get("productid")
	var pidNum int64
	var err error
	if pid != "" {
		pidNum, err = strconv.ParseInt(pid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	key := c.Input().Get("keyword")
	if key != "" {
		products, err := models.SearchProjProduct(pidNum, key) //这里要将侧栏所有id进行循环
		if err != nil {
			beego.Error(err.Error)
		} else {
			link := make([]ProductLink, 0)
			Attachslice := make([]AttachmentLink, 0)
			Pdfslice := make([]PdfLink, 0)
			Articleslice := make([]ArticleContent, 0)
			for _, w := range products {
				Url, _, err := GetUrlPath(w.ProjectId)
				if err != nil {
					beego.Error(err)
				}
				//取到每个成果的附件（模态框打开）；pdf、文章——新窗口打开
				//循环成果
				//每个成果取到所有附件
				//一个附件则直接打开/下载；2个以上则打开模态框
				Attachments, err := models.GetAttachments(w.Id)
				if err != nil {
					beego.Error(err)
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
				linkarr[0].Content = w.Content
				linkarr[0].Created = w.Created
				linkarr[0].Updated = w.Updated
				linkarr[0].Views = w.Views
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
					beego.Error(err)
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
			c.Data["json"] = link
			c.ServeJSON()
		}
	}
}

//未修改
func (c *SearchController) SearchProjects() {
	pid := c.Input().Get("productid")
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	key := c.Input().Get("keyword")
	if key != "" {
		products, err := models.SearchProjProduct(pidNum, key) //这里要将侧栏所有id进行循环
		if err != nil {
			beego.Error(err.Error)
		} else {
			c.Data["json"] = products
			c.ServeJSON()
		}
	}
}

//搜索wiki
func (c *SearchController) SearchWiki() { //search用的是get方法
	tid := c.Input().Get("wikiname")
	if tid != "" {
		c.Data["IsWiki"] = true
		// c.Data["IsSearch"] = true
		c.Data["IsLogin"] = checkAccount(c.Ctx)
		c.TplName = "searchwiki.tpl"
		Searchs, err := models.SearchWikis(tid, false)
		if err != nil {
			beego.Error(err.Error)
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
