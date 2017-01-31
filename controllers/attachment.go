package controllers

import (
	"encoding/json"
	"engineercms/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/httplib"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type AttachController struct {
	beego.Controller
}

// type AttachmentLink struct {
// 	Id    int64
// 	Title string
// 	Link  string
// }

//取得某个成果id下的附件(除去pdf)给table
func (c *AttachController) GetAttachments() {
	id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	c.Data["Id"] = id
	var idNum int64
	var err error
	var Url string
	if id != "" {
		//id转成64为
		idNum, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//由成果id（后台传过来的行id）取得侧栏目录id
		prod, err := models.GetProd(idNum)
		if err != nil {
			beego.Error(err)
		}
		//由proj id取得url
		Url, _, err = GetUrlPath(prod.ProjectId)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(Url)
	} else {

	}
	//根据成果id取得所有附件
	Attachments, err := models.GetAttachments(idNum)
	if err != nil {
		beego.Error(err)
	}
	//对成果进行循环
	//赋予url
	//如果是一个成果，直接给url;如果大于1个，则是数组:这个在前端实现
	// http.ServeFile(ctx.ResponseWriter, ctx.Request, filePath)
	link := make([]AttachmentLink, 0)
	for _, v := range Attachments {
		// fileext := path.Ext(v.FileName)
		if path.Ext(v.FileName) != ".pdf" && path.Ext(v.FileName) != ".PDF" {
			linkarr := make([]AttachmentLink, 1)
			linkarr[0].Id = v.Id
			linkarr[0].Title = v.FileName
			linkarr[0].FileSize = v.FileSize
			linkarr[0].Downloads = v.Downloads
			linkarr[0].Created = v.Created
			linkarr[0].Updated = v.Updated
			linkarr[0].Link = Url + "/" + v.FileName
			link = append(link, linkarr...)
		}
	}
	c.Data["json"] = link
	c.ServeJSON()
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	// c.Data["json"] = root
	// c.ServeJSON()
}

//取得同步成果下的附件列表
func (c *AttachController) GetsynchAttachments() {
	id := c.Input().Get("id")
	site := c.Input().Get("site")
	// beego.Info(site)
	link := make([]AttachmentLink, 0)
	var attachmentlink []AttachmentLink
	//	通过如下接口可以设置请求的超时时间和数据读取时间
	jsonstring, err := httplib.Get(site+"project/product/providesynchattachment?id="+id).SetTimeout(100*time.Second, 30*time.Second).String() //.ToJSON(&productlink)
	if err != nil {
		beego.Error(err)
	}
	//json字符串解析到结构体，以便进行追加
	err = json.Unmarshal([]byte(jsonstring), &attachmentlink)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(productlink)
	link = append(link, attachmentlink...)

	c.Data["json"] = link //products
	c.ServeJSON()
}

//提供给同步用的某个成果id下的附件(除去pdf)给table
func (c *AttachController) ProvideAttachments() {
	id := c.Input().Get("id")
	site := c.Ctx.Input.Site() + ":" + strconv.Itoa(c.Ctx.Input.Port())
	c.Data["Id"] = id
	var idNum int64
	var err error
	var Url string
	if id != "" {
		//id转成64为
		idNum, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//由成果id（后台传过来的行id）取得侧栏目录id
		prod, err := models.GetProd(idNum)
		if err != nil {
			beego.Error(err)
		}
		//由proj id取得url
		Url, _, err = GetUrlPath(prod.ProjectId)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(Url)
	} else {

	}
	//根据成果id取得所有附件
	Attachments, err := models.GetAttachments(idNum)
	if err != nil {
		beego.Error(err)
	}
	//对成果进行循环
	//赋予url
	//如果是一个成果，直接给url;如果大于1个，则是数组:这个在前端实现
	// http.ServeFile(ctx.ResponseWriter, ctx.Request, filePath)
	link := make([]AttachmentLink, 0)
	for _, v := range Attachments {
		// fileext := path.Ext(v.FileName)
		if path.Ext(v.FileName) != ".pdf" && path.Ext(v.FileName) != ".PDF" {
			linkarr := make([]AttachmentLink, 1)
			linkarr[0].Id = v.Id
			linkarr[0].Title = v.FileName
			linkarr[0].FileSize = v.FileSize
			linkarr[0].Downloads = v.Downloads
			linkarr[0].Created = v.Created
			linkarr[0].Updated = v.Updated
			linkarr[0].Link = site + Url + "/" + v.FileName
			link = append(link, linkarr...)
		}
	}
	c.Data["json"] = link
	c.ServeJSON()
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	// c.Data["json"] = root
	// c.ServeJSON()
}

//取得某个成果id下的所有附件(包含pdf和文章)给table
//用于编辑，这个要改，不要显示文章？
//自从文章采用link后，是否可以一同删除？
func (c *AttachController) GetAllAttachments() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	var idNum int64
	var err error
	var Url string
	//id转成64为
	idNum, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//由成果id（后台传过来的行id）取得侧栏目录id
	prod, err := models.GetProd(idNum)
	if err != nil {
		beego.Error(err)
	}
	//由proj id取得url
	Url, _, err = GetUrlPath(prod.ProjectId)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(Url)

	//根据成果id取得所有附件
	Attachments, err := models.GetAttachments(idNum)
	if err != nil {
		beego.Error(err)
	}
	//对成果进行循环
	//赋予url
	//如果是一个成果，直接给url;如果大于1个，则是数组:这个在前端实现
	// http.ServeFile(ctx.ResponseWriter, ctx.Request, filePath)
	link := make([]AttachmentLink, 0)
	for _, v := range Attachments {
		// fileext := path.Ext(v.FileName)
		linkarr := make([]AttachmentLink, 1)
		linkarr[0].Id = v.Id
		linkarr[0].Title = v.FileName
		linkarr[0].FileSize = v.FileSize
		linkarr[0].Downloads = v.Downloads
		linkarr[0].Created = v.Created
		linkarr[0].Updated = v.Updated
		linkarr[0].Link = Url + "/" + v.FileName
		link = append(link, linkarr...)
	}
	//根据成果id取得所有文章——经过论证，不能列出文章，否则会导致使用文字的id删除了对应id的附件。
	// Articles, err := models.GetArticles(idNum)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// for _, x := range Articles {
	// 	linkarr := make([]AttachmentLink, 1)
	// 	linkarr[0].Id = x.Id
	// 	linkarr[0].Title = x.Subtext
	// 	linkarr[0].Created = x.Created
	// 	linkarr[0].Updated = x.Updated
	// 	linkarr[0].Link = "/project/product/article/" + strconv.FormatInt(x.Id, 10)
	// 	link = append(link, linkarr...)
	// }
	c.Data["json"] = link
	c.ServeJSON()
	// c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	// c.Data["json"] = root
	// c.ServeJSON()
}

//取得某个成果id下的附件中的pdf给table
func (c *AttachController) GetPdfs() {
	id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	c.Data["Id"] = id
	var idNum int64
	var err error
	var Url string
	if id != "" {
		//id转成64为
		idNum, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//由成果id（后台传过来的行id）取得侧栏目录id
		prod, err := models.GetProd(idNum)
		if err != nil {
			beego.Error(err)
		}
		//由proj id取得url
		Url, _, err = GetUrlPath(prod.ProjectId)
		if err != nil {
			beego.Error(err)
		}
		beego.Info(Url)
	} else {

	}
	//根据成果id取得所有附件
	Attachments, err := models.GetAttachments(idNum)
	if err != nil {
		beego.Error(err)
	}
	//对成果进行循环
	//赋予url
	//如果是一个成果，直接给url;如果大于1个，则是数组:这个在前端实现
	// http.ServeFile(ctx.ResponseWriter, ctx.Request, filePath)
	link := make([]AttachmentLink, 0)
	for _, v := range Attachments {
		if path.Ext(v.FileName) == ".pdf" || path.Ext(v.FileName) == ".PDF" {
			linkarr := make([]AttachmentLink, 1)
			linkarr[0].Id = v.Id
			linkarr[0].Title = v.FileName
			linkarr[0].FileSize = v.FileSize
			linkarr[0].Downloads = v.Downloads
			linkarr[0].Created = v.Created
			linkarr[0].Updated = v.Updated
			linkarr[0].Link = Url + "/" + v.FileName
			link = append(link, linkarr...)
		}
	}
	c.Data["json"] = link
	c.ServeJSON()
}

//取得同步成果下的pdf列表
func (c *AttachController) GetsynchPdfs() {
	id := c.Input().Get("id")
	site := c.Input().Get("site")
	// beego.Info(site)
	link := make([]PdfLink, 0)
	var pdflink []PdfLink
	//	通过如下接口可以设置请求的超时时间和数据读取时间
	jsonstring, err := httplib.Get(site+"project/product/providesynchpdf?id="+id).SetTimeout(100*time.Second, 30*time.Second).String() //.ToJSON(&productlink)
	if err != nil {
		beego.Error(err)
	}
	//json字符串解析到结构体，以便进行追加
	err = json.Unmarshal([]byte(jsonstring), &pdflink)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(productlink)
	link = append(link, pdflink...)

	c.Data["json"] = link //products
	c.ServeJSON()
}

//提供给同步用的pdf列表数据
func (c *AttachController) ProvidePdfs() {
	id := c.Input().Get("id")
	site := c.Ctx.Input.Site() + ":" + strconv.Itoa(c.Ctx.Input.Port())

	c.Data["Id"] = id
	var idNum int64
	var err error
	var Url string
	if id != "" {
		//id转成64为
		idNum, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//由成果id（后台传过来的行id）取得侧栏目录id
		prod, err := models.GetProd(idNum)
		if err != nil {
			beego.Error(err)
		}
		//由proj id取得url
		Url, _, err = GetUrlPath(prod.ProjectId)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(Url)
	} else {

	}
	//根据成果id取得所有附件
	Attachments, err := models.GetAttachments(idNum)
	if err != nil {
		beego.Error(err)
	}
	//对成果进行循环
	//赋予url
	//如果是一个成果，直接给url;如果大于1个，则是数组:这个在前端实现
	// http.ServeFile(ctx.ResponseWriter, ctx.Request, filePath)
	link := make([]AttachmentLink, 0)
	for _, v := range Attachments {
		if path.Ext(v.FileName) == ".pdf" || path.Ext(v.FileName) == ".PDF" {
			linkarr := make([]AttachmentLink, 1)
			linkarr[0].Id = v.Id
			linkarr[0].Title = v.FileName
			linkarr[0].FileSize = v.FileSize
			linkarr[0].Downloads = v.Downloads
			linkarr[0].Created = v.Created
			linkarr[0].Updated = v.Updated
			linkarr[0].Link = site + Url + "/" + v.FileName
			link = append(link, linkarr...)
		}
	}
	c.Data["json"] = link
	c.ServeJSON()
}

//向某个侧栏id下添加成果——用于第一种批量添加一对一模式
func (c *AttachController) AddAttachment() {
	if checkprodRole(c.Ctx) == 1 {
		//解析表单
		pid := c.Input().Get("pid")
		// beego.Info(pid)
		//pid转成64为
		pidNum, err := strconv.ParseInt(pid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		prodlabel := c.Input().Get("prodlabel")
		prodprincipal := c.Input().Get("prodprincipal")
		size := c.Input().Get("size")
		filesize, err := strconv.ParseInt(size, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		filesize = filesize / 1000.0
		proj, err := models.GetProj(pidNum)
		if err != nil {
			beego.Error(err)
		}
		//根据proj的parentIdpath——这个已经有了专门函数，下列可以简化！
		var path, DiskDirectory, Url string
		if proj.ParentIdPath != "" { //如果不是根目录
			patharray := strings.Split(proj.ParentIdPath, "-")
			for _, v := range patharray {
				//pid转成64位
				idNum, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					beego.Error(err)
				}
				proj1, err := models.GetProj(idNum)
				if err != nil {
					beego.Error(err)
				}
				if proj1.ParentId == 0 { //如果是项目名称，则加上项目编号
					DiskDirectory = ".\\attachment\\" + proj1.Code + proj1.Title
					Url = "/attachment/" + proj1.Code + proj1.Title
				} else {
					path = proj1.Title
				}
				DiskDirectory = DiskDirectory + "\\" + path
				Url = Url + "/" + path
			}
			DiskDirectory = DiskDirectory + "\\" + proj.Title //加上自身
			Url = Url + "/" + proj.Title
			// beego.Info(DiskDirectory)
			// beego.Info(Url)
		} else { //如果是根目录
			DiskDirectory = ".\\attachment\\" + proj.Code + proj.Title //加上自身
			Url = "/attachment/" + proj.Title
		}
		//获取上传的文件
		_, h, err := c.GetFile("file")
		if err != nil {
			beego.Error(err)
		}
		var attachment string
		// var filesize int64
		if h != nil {
			//保存附件
			attachment = h.Filename
			// beego.Info(attachment)
			// path = ".\\attachment\\" + categoryproj.Number + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + h.Filename
			path = DiskDirectory + "\\" + h.Filename
			// path := c.Input().Get("url")  //存文件的路径
			// path = path[3:]
			// path = "./attachment" + "/" + h.Filename
			// f.Close() // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
			// filesize, _ = FileSize(path)
			// filesize = filesize / 1000.0
			//将附件的编号和名称写入数据库
			_, filename1, filename2, _, _, _, _ := Record(attachment)
			// filename1, filename2 := SubStrings(attachment)
			//当2个文件都取不到filename1的时候，数据库里的tnumber的唯一性检查出错。
			// beego.Info(filename1)
			// beego.Info(filename2)
			if filename1 == "" {
				filename1 = filename2 //如果编号为空，则用文件名代替，否则多个编号为空导致存入数据库唯一性检查错误
			}
			code := filename1
			title := filename2
			//存入成果数据库
			//如果编号重复，则不写入，值返回Id值。
			//根据id添加成果code, title, label, principal, content string, projectid int64
			prodId, err := models.AddProduct(code, title, prodlabel, prodprincipal, "", pidNum)
			if err != nil {
				beego.Error(err)
			}
			//把成果id作为附件的parentid，把附件的名称等信息存入附件数据库
			//如果附件名称相同，则覆盖上传，但数据库不追加
			_, err = models.AddAttachment(attachment, filesize, 0, prodId)
			if err != nil {
				beego.Error(err)
			} else {
				//存入文件夹
				err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
				if err != nil {
					beego.Error(err)
				}
				c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "title": attachment, "original": attachment, "url": Url + "/" + attachment}
				c.ServeJSON()
			}
		}
	} else {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// c.TplName = "topic_one_add.tpl" //不加这句上传出错，虽然可以成功上传
	// c.Redirect("/topic", 302)
	// success : 0 | 1,           // 0 表示上传失败，1 表示上传成功
	//    message : "提示的信息，上传成功或上传失败及错误信息等。",
	//    url     : "图片地址"        // 上传成功时才返回
}

//向某个侧栏id下添加成果——用于第二种添加，多附件模式
func (c *AttachController) AddAttachment2() {
	if checkprodRole(c.Ctx) == 1 {
		//解析表单
		pid := c.Input().Get("pid")
		// beego.Info(pid)
		//pid转成64为
		pidNum, err := strconv.ParseInt(pid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		prodcode := c.Input().Get("prodcode")
		prodname := c.Input().Get("prodname")
		prodlabel := c.Input().Get("prodlabel")
		prodprincipal := c.Input().Get("prodprincipal")
		size := c.Input().Get("size")
		filesize, err := strconv.ParseInt(size, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		filesize = filesize / 1000.0
		// proj, err := models.GetProj(pidNum)
		// if err != nil {
		// 	beego.Error(err)
		// }
		//根据proj的Id
		Url, DiskDirectory, err := GetUrlPath(pidNum)
		if err != nil {
			beego.Error(err)
		}
		//获取上传的文件
		_, h, err := c.GetFile("file")
		if err != nil {
			beego.Error(err)
		}
		var path, attachment string
		// var filesize int64
		if h != nil {
			//保存附件
			attachment = h.Filename
			// beego.Info(attachment)
			// path = ".\\attachment\\" + categoryproj.Number + categoryproj.Title + "\\" + categoryphase.Title + "\\" + categoryspec.Title + "\\" + category + "\\" + h.Filename
			path = DiskDirectory + "\\" + h.Filename
			// path := c.Input().Get("url")  //存文件的路径
			// path = path[3:]
			// path = "./attachment" + "/" + h.Filename
			// f.Close()// 关闭上传的文件，不然的话会出现临时文件不能清除的情况
			// filesize, _ = FileSize(path)
			// filesize = filesize / 1000.0
			//存入成果数据库
			//如果编号重复，则不写入，值返回Id值。
			//根据id添加成果code, title, label, principal, content string, projectid int64
			prodId, err := models.AddProduct(prodcode, prodname, prodlabel, prodprincipal, "", pidNum)
			if err != nil {
				beego.Error(err)
			}
			//把成果id作为附件的parentid，把附件的名称等信息存入附件数据库
			//如果附件名称相同，则覆盖上传，但数据库不追加
			_, err = models.AddAttachment(attachment, filesize, 0, prodId)
			if err != nil {
				beego.Error(err)
			} else {
				err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
				if err != nil {
					beego.Error(err)
				}
				c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "title": attachment, "original": attachment, "url": Url + "/" + attachment}
				c.ServeJSON()
			}
		}
	} else {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
}

//向一个成果id下追加附件
func (c *AttachController) UpdateAttachment() {
	if checkprodRole(c.Ctx) == 1 {
		//解析表单
		pid := c.Input().Get("pid")
		size := c.Input().Get("size")
		filesize, err := strconv.ParseInt(size, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		filesize = filesize / 1000.0
		// beego.Info(pid)
		//pid转成64为
		pidNum, err := strconv.ParseInt(pid, 10, 64)
		if err != nil {
			beego.Error(err)
		}

		prod, err := models.GetProd(pidNum)
		if err != nil {
			beego.Error(err)
		}
		//根据proj的Id
		Url, DiskDirectory, err := GetUrlPath(prod.ProjectId)
		if err != nil {
			beego.Error(err)
		}
		//获取上传的文件
		_, h, err := c.GetFile("file")
		if err != nil {
			beego.Error(err)
		}
		var path, attachment string
		// var filesize int64
		if h != nil {
			//保存附件
			attachment = h.Filename
			path = DiskDirectory + "\\" + h.Filename // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
			// filesize, _ = FileSize(path)
			// filesize = filesize / 1000.0
			//把成果id作为附件的parentid，把附件的名称等信息存入附件数据库
			//如果附件名称相同，则覆盖上传，但数据库不追加
			_, err = models.AddAttachment(attachment, filesize, 0, pidNum)
			if err != nil {
				beego.Error(err)
			} else {
				err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
				if err != nil {
					beego.Error(err)
				}
				c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "title": attachment, "original": attachment, "url": Url + "/" + attachment}
				c.ServeJSON()
			}
		}
	} else {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
}

//删除附件——这个用于针对删除一个附件
func (c *AttachController) DeleteAttachment() {
	if checkprodRole(c.Ctx) == 1 {
		//解析表单
		ids := c.GetString("ids")
		array := strings.Split(ids, ",")
		for _, v := range array {
			// pid = strconv.FormatInt(v1, 10)
			//id转成64位
			idNum, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				beego.Error(err)
			}
			//取得附件的成果id——再取得成果的项目目录id——再取得路径
			attach, err := models.GetAttachbyId(idNum)
			if err != nil {
				beego.Error(err)
			}
			prod, err := models.GetProd(attach.ProductId)
			if err != nil {
				beego.Error(err)
			}
			//根据proj的id
			_, DiskDirectory, err := GetUrlPath(prod.ProjectId)
			if err != nil {
				beego.Error(err)
			}
			path := DiskDirectory + "\\" + attach.FileName
			err = os.Remove(path)
			if err != nil {
				beego.Error(err)
			}
			err = models.DeleteAttachment(idNum)
			if err != nil {
				beego.Error(err)
			}
		}
		c.Data["json"] = "ok"
		c.ServeJSON()
	} else {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
}

// 下载附件
// func (c *AttachController) ImageFilter() {
func ImageFilter(ctx *context.Context) {
	// token := path.Base(ctx.Request.RequestURI)

	// split token and file ext
	// var filePath string
	// if i := strings.IndexRune(token, '.'); i == -1 {
	// 	return
	// } else {
	// 	filePath = token[i+1:]
	// 	token = token[:i]
	// }

	// decode token to file path
	// var image models.Image
	// if err := image.DecodeToken(token); err != nil {
	// 	beego.Info(err)
	// 	return
	// }

	// file real path
	// filePath = GetUrlPath(&image) + filePath
	//1.url处理中文字符路径，[1:]截掉路径前面的/斜杠
	filePath := path.Base(ctx.Request.RequestURI)
	// filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:]) //  attachment/SL2016测试添加成果/A/FB/1/Your First Meteor Application.pdf
	// if err != nil {
	// 	beego.Error(err)
	// }
	// if x-send on then set header and http status
	// fall back use proxy serve file
	// if setting.ImageXSend {
	// 	ext := filepath.Ext(filePath)
	// 	ctx.Output.ContentType(ext)
	// 	ctx.Output.Header(setting.ImageXSendHeader, "/"+filePath)
	// 	ctx.Output.SetStatus(200)
	// } else {
	// direct serve file use go
	http.ServeFile(ctx.ResponseWriter, ctx.Request, filePath)
	// http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)
	// }
}

func (c *AttachController) DownloadAttachment() {
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//4.取得客户端用户名
	var uname string
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	var role, userrole int
	if v != nil {
		uname = v.(string)
		c.Data["Uname"] = v.(string)
		user, err := models.GetUserByUsername(uname)
		if err != nil {
			beego.Error(err)
		}
		userrole = user.Role
	} else {
		userrole = 5
	}
	// uname := v.(string) //ck.Value
	// 4.取出用户的权限等级
	// role, _ := checkRole(c.Ctx) //login里的
	// 5.进行逻辑分析：
	// rolename, err := strconv.ParseInt(role, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// username := c.Input().Get("username")
	iprole := Getiprole(c.Ctx.Input.IP())
	if iprole <= userrole {
		role = iprole
	} else {
		role = userrole
	}
	// beego.Info(c.Ctx.Input.IP())
	beego.Info(role)
	//1.url处理中文字符路径，[1:]截掉路径前面的/斜杠
	// filePath := path.Base(ctx.Request.RequestURI)
	filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:]) //  attachment/SL2016测试添加成果/A/FB/1/Your First Meteor Application.pdf
	if err != nil {
		beego.Error(err)
	}
	fileext := path.Ext(filePath)
	switch fileext {
	case ".pdf", ".PDF":
		if role <= 3 {
			http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)
		} else {
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			// c.Redirect("/roleerr", 302)
			return
			// c.Data["json"] = "权限不够！"
			// c.ServeJSON()
		}
	default:
		if role <= 2 {
			http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)
		} else {
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			// c.Redirect("/roleerr", 302)
			return
			// c.Data["json"] = "权限不够！"
			// c.ServeJSON()
		}
	}
}

//返回文件大小
func FileSize(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

//根据id返回附件url和文件夹路径
func GetUrlPath(id int64) (Url, DiskDirectory string, err error) {
	proj, err := models.GetProj(id)
	if err != nil {
		beego.Error(err)
	}
	//根据proj的parentIdpath
	var path string
	if proj.ParentIdPath != "" { //如果不是根目录
		patharray := strings.Split(proj.ParentIdPath, "-")
		for _, v := range patharray {
			//pid转成64为
			idNum, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				beego.Error(err)
			}
			proj1, err := models.GetProj(idNum)
			if err != nil {
				beego.Error(err)
			}
			if proj1.ParentId == 0 { //如果是项目名称，则加上项目编号
				DiskDirectory = ".\\attachment\\" + proj1.Code + proj1.Title
				Url = "/attachment/" + proj1.Code + proj1.Title
			} else {
				path = proj1.Title
				DiskDirectory = DiskDirectory + "\\" + path
				Url = Url + "/" + path
			}
		}
		DiskDirectory = DiskDirectory + "\\" + proj.Title //加上自身
		Url = Url + "/" + proj.Title
		// beego.Info(DiskDirectory)
		// beego.Info(Url)
	} else { //如果是根目录
		DiskDirectory = ".\\attachment\\" + proj.Code + proj.Title //加上自身
		Url = "/attachment/" + proj.Code + proj.Title
	}
	return Url, DiskDirectory, err
}
