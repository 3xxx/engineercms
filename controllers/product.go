package controllers

import (
	// "encoding/json"
	"engineercms/models"
	"github.com/astaxie/beego"
	// "os"
	"path"
	"strconv"
	"time"
	// "strings"
)

type ProdController struct {
	beego.Controller
}

type ProductLink struct {
	Id             int64
	Code           string
	Title          string
	Label          string
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
}

type AttachmentLink struct {
	Id        int64
	Title     string
	Link      string
	FileSize  int64
	Downloads int64
	Created   time.Time
	Updated   time.Time
}

type PdfLink struct {
	Id        int64
	Title     string
	Link      string
	FileSize  int64
	Downloads int64
	Created   time.Time
	Updated   time.Time
}
type ArticleContent struct {
	Id      int64
	Content string
	// Link    string
	// Views   int64
	Created time.Time
	Updated time.Time
}

//根据项目侧栏id查看这个id下的成果页面，数据用GetProducts
//任何一级目录下都可以放成果
func (c *ProdController) GetProjProd() {
	id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	c.Data["Id"] = id
	role := Getiprole(c.Ctx.Input.IP())
	c.Data["role"] = role
	// var categories []*models.ProjCategory
	// var err error
	//id转成64为
	// idNum, err := strconv.ParseInt(id, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	//取项目本身
	// category, err := models.GetProj(idNum)
	// if err != nil {
	// 	beego.Error(err)
	// }
	//取项目所有子孙
	// categories, err := models.GetProjectsbyPid(idNum)
	// if err != nil {
	// 	beego.Error(err)
	// }
	//算出最大级数
	// grade := make([]int, 0)
	// for _, v := range categories {
	// 	grade = append(grade, v.Grade)
	// }
	// height := intmax(grade[0], grade[1:]...)

	// c.Data["json"] = root
	// c.ServeJSON()
	c.TplName = "project_products.tpl"
}

//取得某个侧栏id下的成果给table
func (c *ProdController) GetProducts() {
	id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	c.Data["Id"] = id
	var idNum int64
	var err error
	if id != "" {
		//id转成64为
		idNum, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//由成果id（后台传过来的行id）取得侧栏目录id
		// prod, err := models.GetProd(idNum)
		// if err != nil {
		// 	beego.Error(err)
		// }
		// //由proj id取得url
		// Url, _, err = GetUrlPath(prod.ProjectId)
		// if err != nil {
		// 	beego.Error(err)
		// }
		// beego.Info(Url)
	} else {

	}
	//根据id取得所有成果
	products, err := models.GetProducts(idNum)
	if err != nil {
		beego.Error(err)
	}
	//由proj id取得url
	Url, _, err := GetUrlPath(idNum)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(Url)
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
			// articlearr[0].Link = Url
			Articleslice = append(Articleslice, articlearr...)
		}
		linkarr[0].Articlecontent = Articleslice
		Articleslice = make([]ArticleContent, 0)
		link = append(link, linkarr...)
	}

	c.Data["json"] = link //products
	c.ServeJSON()
	// c.Data["json"] = root
	// c.ServeJSON()
}

//向某个侧栏id下添加成果
func (c *ProdController) AddProduct() {
	id := c.Ctx.Input.Param(":id")
	pid := c.Input().Get("pid")
	code := c.Input().Get("code")
	title := c.Input().Get("title")
	label := c.Input().Get("label")
	principal := c.Input().Get("principal")
	content := c.Input().Get("content")
	// beego.Info(id)
	c.Data["Id"] = id
	//id转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据id添加成果code, title, label, principal, content string, projectid int64
	_, err = models.AddProduct(code, title, label, principal, content, pidNum)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = "ok"
	c.ServeJSON()
	// c.Data["json"] = root
	// c.ServeJSON()
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
