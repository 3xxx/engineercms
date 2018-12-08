//成果里的附件操作
package controllers

import (
	// "code.google.com/p/mahonia"
	"encoding/json"
	"github.com/3xxx/engineercms/controllers/utils"
	"github.com/3xxx/engineercms/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego/logs"
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
	} //else {

	//}
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
			// linkarr[0].Suffix=path.Ext(v.FileName)
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
	_, _, uid, _, _ := checkprodRole(c.Ctx)

	meritbasic, err := models.GetMeritBasic()
	if err != nil {
		beego.Error(err)
	}
	var news string
	var cid, topprojectid int64
	var attachment, parentidpath, parentidpath1 string
	var filepath, DiskDirectory, Url string
	var catalog models.PostMerit
	// if role == 1 {
	//解析表单
	pid := c.Input().Get("pid")
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
	//根据pid查出项目id
	proj, err := models.GetProj(pidNum)
	if err != nil {
		beego.Error(err)
	}

	//根据proj的parentIdpath——这个已经有了专门函数，下列可以简化！
	//由proj id取得url
	// Url, DiskDirectory, err = GetUrlPath(proj.Id)
	// if err != nil {
	// 	beego.Error(err)
	// }
	if proj.ParentIdPath != "" { //如果不是根目录
		parentidpath = strings.Replace(strings.Replace(proj.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 = strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		topprojectid, err = strconv.ParseInt(patharray[0], 10, 64)
		if err != nil {
			beego.Error(err)
		}
		// patharray := strings.Split(proj.ParentIdPath1, "-")
		//pid转成64位
		meritNum, err := strconv.ParseInt(patharray[0], 10, 64)
		if err != nil {
			beego.Error(err)
		}
		meritproj, err := models.GetProj(meritNum)
		if err != nil {
			beego.Error(err)
		}
		catalog.ProjectNumber = meritproj.Code
		catalog.ProjectName = meritproj.Title

		if len(patharray) > 1 {
			//pid转成64位
			meritNum1, err := strconv.ParseInt(patharray[1], 10, 64)
			if err != nil {
				beego.Error(err)
			}
			meritproj1, err := models.GetProj(meritNum1)
			if err != nil {
				beego.Error(err)
			}
			catalog.DesignStage = meritproj1.Title
		}

		if len(patharray) > 2 {
			//pid转成64位
			meritNum2, err := strconv.ParseInt(patharray[2], 10, 64)
			if err != nil {
				beego.Error(err)
			}
			meritproj2, err := models.GetProj(meritNum2)
			if err != nil {
				beego.Error(err)
			}
			catalog.Section = meritproj2.Title
		}

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
				DiskDirectory = "./attachment/" + proj1.Code + proj1.Title
				Url = "/attachment/" + proj1.Code + proj1.Title
			} else {
				filepath = proj1.Title
				DiskDirectory = DiskDirectory + "/" + filepath
				Url = Url + "/" + filepath
			}
		}
		DiskDirectory = DiskDirectory + "/" + proj.Title //加上自身
		Url = Url + "/" + proj.Title
	} else { //如果是根目录
		DiskDirectory = "./attachment/" + proj.Code + proj.Title //加上自身
		Url = "/attachment/" + proj.Title
		catalog.ProjectNumber = proj.Code
		catalog.ProjectName = proj.Title
		topprojectid = proj.Id
	}
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	if h != nil {
		//保存附件
		attachment = h.Filename

		// filepath = DiskDirectory + "/" + h.Filename
		// f.Close() // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		//将附件的编号和名称写入数据库
		_, filename1, filename2, _, _, _, _ := Record(attachment)
		// filename1, filename2 := SubStrings(attachment)
		//当2个文件都取不到filename1的时候，数据库里的tnumber的唯一性检查出错。
		if filename1 == "" {
			filename1 = filename2 //如果编号为空，则用文件名代替，否则多个编号为空导致存入数据库唯一性检查错误
		}
		code := filename1
		title := filename2
		//存入成果数据库
		//如果编号重复，则不写入，只返回Id值。
		//根据id添加成果code, title, label, principal, content string, projectid int64
		prodId, err := models.AddProduct(code, title, prodlabel, prodprincipal, uid, pidNum, topprojectid)
		if err != nil {
			beego.Error(err)
		}
		//改名，替换文件名中的#和斜杠
		filename2 = strings.Replace(filename2, "#", "号", -1)
		filename2 = strings.Replace(filename2, "/", "-", -1)
		FileSuffix := path.Ext(h.Filename)

		//成果写入postmerit表，准备提交merit*********
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
		cid, err, news = models.AddPostMerit(catalog)
		if err != nil {
			beego.Error(err)
		} else {
			link1 := Url + "/" + filename1 + filename2 + FileSuffix //附件链接地址
			filepath = DiskDirectory + "/" + filename1 + filename2 + FileSuffix
			_, err = models.AddCatalogLink(cid, link1)
			if err != nil {
				beego.Error(err)
			}
			data := news
			c.Ctx.WriteString(data)
		}
		//生成提交merit的清单结束*******************

		//把成果id作为附件的parentid，把附件的名称等信息存入附件数据库
		//如果附件名称相同，则覆盖上传，但数据库不追加
		attachmentname := filename1 + filename2 + FileSuffix
		_, err = models.AddAttachment(attachmentname, filesize, 0, prodId)
		if err != nil {
			beego.Error(err)
		} else {
			//存入文件夹
			err = c.SaveToFile("file", filepath) //.Join("attachment", attachment)) //存文件    WaterMark(filepath)    //给文件加水印
			if err != nil {
				beego.Error(err)
			}
			c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "title": attachment, "original": attachment, "url": Url + "/" + attachment}
			c.ServeJSON()
		}
	}
	// } else {
	// route := c.Ctx.Request.URL.String()
	// c.Data["Url"] = route
	// c.Redirect("/roleerr?url="+route, 302)
	// c.Redirect("/roleerr", 302)
	// return
	// }
	// c.TplName = "topic_one_add.tpl" //不加这句上传出错，虽然可以成功上传
	// c.Redirect("/topic", 302)
	// success : 0 | 1,           // 0 表示上传失败，1 表示上传成功
	//    message : "提示的信息，上传成功或上传失败及错误信息等。",
	//    url     : "图片地址"        // 上传成功时才返回
}

//向服务器保存dwg文件
func (c *AttachController) SaveDwgfile() {

	id := c.Input().Get("id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据附件id取得附件的prodid，路径
	attachment, err := models.GetAttachbyId(idNum)
	if err != nil {
		beego.Error(err)
	}

	product, err := models.GetProd(attachment.ProductId)
	if err != nil {
		beego.Error(err)
	}
	//由proj id取得文件路径
	_, diskdirectory, err := GetUrlPath(product.ProjectId)
	if err != nil {
		beego.Error(err)
	}

	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	if h != nil {
		//保存附件
		filepath := diskdirectory + "/" + attachment.FileName
		// f.Close() // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		//存入文件夹
		err = c.SaveToFile("file", filepath) //.Join("attachment", attachment)) //存文件    WaterMark(filepath)    //给文件加水印
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "title": attachment, "original": attachment}
			c.ServeJSON()
		}
	}
}

//向某个侧栏id下添加成果——用于第二种添加，多附件模式
func (c *AttachController) AddAttachment2() {
	//取得客户端用户名
	// v := c.GetSession("uname")
	// var user models.User
	// var err error
	// if v != nil {
	// 	uname := v.(string)
	// 	user, err = models.GetUserByUsername(uname)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// }
	_, _, uid, _, _ := checkprodRole(c.Ctx)

	meritbasic, err := models.GetMeritBasic()
	if err != nil {
		beego.Error(err)
	}
	var catalog models.PostMerit
	var news string
	var cid int64

	// _, role := checkprodRole(c.Ctx)
	// if role == 1 {
	//解析表单
	pid := c.Input().Get("pid")
	// beego.Info(pid)
	//pid转成64为
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
	prodcode := c.Input().Get("prodcode") //和上面那个区别仅仅在此处而已
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
	Number, Name, DesignStage, Section, err := GetProjTitleNumber(pidNum)
	if err != nil {
		beego.Error(err)
	}
	catalog.ProjectNumber = Number
	catalog.ProjectName = Name
	catalog.DesignStage = DesignStage
	catalog.Section = Section

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
		path = DiskDirectory + "/" + h.Filename
		// f.Close()// 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		//存入成果数据库
		//如果编号重复，则不写入，值返回Id值。
		//根据id添加成果code, title, label, principal, content string, projectid int64
		prodId, err := models.AddProduct(prodcode, prodname, prodlabel, prodprincipal, uid, pidNum, topprojectid)
		if err != nil {
			beego.Error(err)
		}

		//成果写入postmerit表，准备提交merit*********
		catalog.Tnumber = prodcode
		catalog.Name = prodname
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
		cid, err, news = models.AddPostMerit(catalog)
		if err != nil {
			beego.Error(err)
		} else {
			link1 := Url + "/" + attachment //附件链接地址
			_, err = models.AddCatalogLink(cid, link1)
			if err != nil {
				beego.Error(err)
			}
			data := news
			c.Ctx.WriteString(data)
		}
		//生成提交merit的清单结束*******************

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
	// } else {
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/roleerr?url="+route, 302)
	// 	// c.Redirect("/roleerr", 302)
	// 	return
	// }
}

//向一个成果id下追加附件
func (c *AttachController) UpdateAttachment() {
	// _, role := checkprodRole(c.Ctx)
	// if role == 1 {
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
		path = DiskDirectory + "/" + h.Filename // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
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
	// } else {
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/roleerr?url="+route, 302)
	// 	// c.Redirect("/roleerr", 302)
	// 	return
	// }
}

//删除附件——这个用于针对删除一个附件
func (c *AttachController) DeleteAttachment() {
	// _, role := checkprodRole(c.Ctx)
	// if role == 1 {
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
		} else if DiskDirectory != "" {
			path := DiskDirectory + "/" + attach.FileName
			err = os.Remove(path)
			if err != nil {
				beego.Error(err)
			}
			err = models.DeleteAttachment(idNum)
			if err != nil {
				beego.Error(err)
			}
		}
	}
	c.Data["json"] = "ok"
	c.ServeJSON()
	// } else {
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/roleerr?url="+route, 302)
	// 	// c.Redirect("/roleerr", 302)
	// 	return
	// }
}

// 下载附件——这个仅是测试
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

//根据权限查看附件/downloadattachment?id=
func (c *AttachController) DownloadAttachment() {
	// logs := logs.NewLogger()
	// logs.EnableFuncCallDepth(true)
	// logs.SetLogger("multifile", `{"filename":"log/engineercms.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info"]}`)
	// v := c.GetSession("pwd")
	// beego.Info("v.(string)")
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//4.取得客户端用户名
	var projurl string
	// v := c.GetSession("uname")
	// if v != nil {
	// 	uname = v.(string)
	// 	c.Data["Uname"] = v.(string)
	// 	user, err := models.GetUserByUsername(uname)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// 	useridstring = strconv.FormatInt(user.Id, 10)
	// }
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	useridstring := strconv.FormatInt(uid, 10)
	var usersessionid string //客户端sesssionid
	if islogin {
		usersessionid = c.Ctx.Input.Cookie("hotqinsessionid")
		//服务端sessionid怎么取出
		// v := c.GetSession("uname")
		// beego.Info(v.(string))
	}
	id := c.Input().Get("id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
		utils.FileLogs.Error(c.Ctx.Input.IP() + " 转换id " + err.Error())
	}

	//根据附件id取得附件的prodid，路径
	attachment, err := models.GetAttachbyId(idNum)
	if err != nil {
		beego.Error(err)
		utils.FileLogs.Error(c.Ctx.Input.IP() + " 查询附件 " + err.Error())
	}

	product, err := models.GetProd(attachment.ProductId)
	if err != nil {
		beego.Error(err)
		utils.FileLogs.Error(c.Ctx.Input.IP() + " 用附件查询成果 " + err.Error())
	}

	//根据projid取出路径
	proj, err := models.GetProj(product.ProjectId)
	if err != nil {
		beego.Error(err)
		utils.FileLogs.Error(err.Error())
	}
	if proj.ParentIdPath == "" || proj.ParentIdPath == "$#" {
		projurl = "/" + strconv.FormatInt(proj.Id, 10) + "/"
	} else {
		// projurl = "/" + strings.Replace(proj.ParentIdPath, "-", "/", -1) + "/" + strconv.FormatInt(proj.Id, 10) + "/"
		projurl = "/" + strings.Replace(strings.Replace(proj.ParentIdPath, "#", "/", -1), "$", "", -1) + strconv.FormatInt(proj.Id, 10) + "/"
	}
	//由proj id取得url
	fileurl, _, err := GetUrlPath(product.ProjectId)
	if err != nil {
		beego.Error(err)
		utils.FileLogs.Error(c.Ctx.Input.IP() + " 查询成果路径 " + err.Error())
	}
	fileext := path.Ext(attachment.FileName)
	switch fileext {
	case ".JPG", ".jpg", ".png", ".PNG", ".bmp", ".BMP", ".mp4", ".MP4":
		c.Ctx.Output.Download(fileurl + "/" + attachment.FileName)
	case ".dwg", ".DWG":
		//beego.Info(c.Ctx.Input.Site())
		if e.Enforce(useridstring, projurl, c.Ctx.Request.Method, fileext) || isadmin { //+ strconv.Itoa(c.Ctx.Input.Port())
			// dwglink, err := url.ParseRequestURI(c.Ctx.Input.Site() + ":" + "/" + fileurl + "/" + attachment.FileName)
			dwglink, err := url.ParseRequestURI(c.Ctx.Input.Scheme() + "://" + c.Ctx.Input.IP() + ":" + strconv.Itoa(c.Ctx.Input.Port()) + "/" + fileurl + "/" + attachment.FileName)
			if err != nil {
				beego.Error(err)
				utils.FileLogs.Error(c.Ctx.Input.IP() + " 获取dwg路径 " + err.Error())
			}
			// beego.Info(dwglink)
			// beego.Info(usersessionid)
			c.Data["Id"] = id
			c.Data["DwgLink"] = dwglink
			c.Data["Sessionid"] = usersessionid
			c.TplName = "dwg.tpl"
		} else {
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			// c.Redirect("/roleerr", 302)
			return
		}
		// case ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx":
		// c.Ctx.Output.Download(fileurl + "/" + attachment.FileName)
		//这里缺少权限设置！！！！！！！！！！！
		// beego.Info(useridstring)
		// beego.Info(isadmin)
	default:
		if e.Enforce(useridstring, projurl, c.Ctx.Request.Method, fileext) || isadmin {
			// http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)//这样写下载的文件名称不对
			// c.Redirect(url+"/"+attachment.FileName, 302)
			c.Ctx.Output.Download(fileurl + "/" + attachment.FileName)
			utils.FileLogs.Info(username + " " + "download" + " " + fileurl + "/" + attachment.FileName)
		} else {
			utils.FileLogs.Info(c.Ctx.Input.IP() + "want " + "download" + " " + fileurl + "/" + attachment.FileName)
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			// c.Redirect("/roleerr", 302)
			return
		}
	}
	// config := make(map[string]interface{})
	// config["filename"] = "e:/golang/go_pro/logs/logcollect.log"
	// config["filename"] = "e://golang//go_pro//logs//logcollect.log"
	// config["filename"] = "e:/golang/go_pro/logs/logcollect.log"
	// config["level"] = logs.LevelDebug
	// configStr, err := json.Marshal(config)
	// if err != nil {
	// 	fmt.Println("marshal failed, err:", err)
	// 	return
	// }
	// logs.SetLogger(logs.AdapterFile, string(configStr))
	// utils.FileLogs.Warn("this is a warn, my name is %s", map[string]int{"key": 2016})
	// utils.FileLogs.Critical("oh,crash")
	// logs.Close()
	// 日志有以下7个级别                  对应的方法
	// LevelEmergency = 0      --> logs.Emergency()
	// LevelAlert = 1          --> logs.Alert()
	// LevelCritical = 2       --> logs.Critical()
	// LevelError = 3          --> logs.Error()
	// LevelWarning = 4        --> logs.Warning()
	// LevelNotice = 5         --> logs.Notice()
	// LevelInformational = 6  --> logs.Informational()
	// LevelDebug = 7          --> logs.Debug()
	// utils.FileLogs.Info("this is a file log with info.")
	// utils.FileLogs.Debug("this is a file log with debug.")
	// utils.FileLogs.Alert("this is a file log with alert.")
	// utils.FileLogs.Error("this is a file log with error.")
	// utils.FileLogs.Trace("this is a file log with trace.")
}

//目前有文章中的图片、成果中文档的预览、onlyoffice中的文档协作、pdf中的附件路径等均采用绝对路径型式
//文章中的附件呢？
//default中的pdf页面中的{{.pdflink}}，绝对路径
// type Session struct {
// 	Session int
// }
//attachment/路径/附件名称
func (c *AttachController) Attachment() {
	// beego.Info("v111.(string)")
	// c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	// http.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	// var sess Session
	// json.Unmarshal(c.Ctx.Input.RequestBody, &sess)
	//beego.Info(c.Ctx.Input.RequestBody)
	// var ob models.Object
	//    json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	// beego.Info(sess.Session)
	//如果url带了sessionid,就能取到uid等信息
	var useridstring string
	_, _, uid, isadmin, _ := checkprodRole(c.Ctx)
	// if uid != 0 {
	useridstring = strconv.FormatInt(uid, 10)
	// beego.Info(useridstring)
	// beego.Info(isadmin)
	// } else {
	// 	//根据url携带的sessionid
	// 	v := c.GetSession("uname")
	// 	if v != nil {
	// 		user, err := models.GetUserByUsername(v.(string))
	// 		if err != nil {
	// 			beego.Error(err)
	// 		} else {
	// 			useridstring = strconv.FormatInt(user.Id, 10)
	// 			if user.Role == "1" {
	// 				isadmin = true
	// 			} else {
	// 				isadmin = false
	// 			}
	// 		}
	// 	}
	// 	beego.Info(useridstring)
	// 	beego.Info(isadmin)
	// }
	// }
	//1.url处理中文字符路径，[1:]截掉路径前面的/斜杠
	// filePath := path.Base(c.Ctx.Request.RequestURI)
	filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:]) //attachment/SL2016测试添加成果/A/FB/1/Your First Meteor Application.pdf
	if err != nil {
		beego.Error(err)
	}
	if strings.Contains(filePath, "?hotqinsessionid=") {
		filePathtemp := strings.Split(filePath, "?")
		filePath = filePathtemp[0]
		// beego.Info(filePath)
	}
	// beego.Info(filePath)
	fileext := path.Ext(filePath)
	filepath1 := path.Dir(filePath)
	array := strings.Split(filepath1, "/")
	// beego.Info(strings.Split(filepath1, "/"))
	//查出所有项目
	var pid int64
	proj, err := models.GetProjects()
	//循环，项目编号+名称=array[1]则其id
	for _, v := range proj {
		if v.Code+v.Title == array[1] {
			pid = v.Id
			break
		}
	}

	var projurl models.Project
	projurls := "/" + strconv.FormatInt(pid, 10)
	//id作为parentid+array[2]
	if len(array) > 1 {
		for i := 2; i < len(array); i++ {
			// beego.Info(i)
			// beego.Info(array[i])
			// beego.Info(projurl.Id)
			if i == 2 {
				// beego.Info(pid)
				projurl, err = models.GetProjbyParentidTitle(pid, array[i])
				if err != nil {
					beego.Error(err)
				}
			} else {
				projurl, err = models.GetProjbyParentidTitle(projurl.Id, array[i])
				if err != nil {
					beego.Error(err)
				}
			}
			projurls = projurls + "/" + strconv.FormatInt(projurl.Id, 10)
		}
	}

	switch fileext {
	case ".JPG", ".jpg", ".png", ".PNG", ".bmp", ".BMP", ".mp4", ".MP4":
		http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)
	// case ".dwg", ".DWG":
	// http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)
	// case ".doc", ".docx", ".xls", ".xlsx", ".ppt", ".pptx":
	// beego.Info("ok")
	// http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)
	// beego.Info(filePath)
	// beego.Info(useridstring)
	//这里缺少权限设置！！！！！！！！！！！
	default:
		if e.Enforce(useridstring, projurls+"/", c.Ctx.Request.Method, fileext) || isadmin {
			http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath) //这样写下载的文件名称不对
			// c.Redirect(url+"/"+attachment.FileName, 302)
			// c.Ctx.Output.Download(fileurl + "/" + attachment.FileName)
		} else {
			// beego.Info(useridstring)
			route := c.Ctx.Request.URL.String()
			c.Data["Url"] = route
			c.Redirect("/roleerr?url="+route, 302)
			// c.Redirect("/roleerr", 302)
			return
		}
	}
	// switch fileext {
	// case ".pdf", ".PDF", ".JPG", ".jpg", ".png", ".PNG", ".bmp", ".BMP", ".mp4", ".MP4":
	// 	if role < 5 {
	// 		http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)
	// 	} else {
	// 		route := c.Ctx.Request.URL.String()
	// 		c.Data["Url"] = route
	// 		c.Redirect("/roleerr?url="+route, 302)
	// 		return
	// 	}
	// default:
	// 	if role <= 2 {
	// 		http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)
	// 	} else {
	// 		route := c.Ctx.Request.URL.String()
	// 		c.Data["Url"] = route
	// 		c.Redirect("/roleerr?url="+route, 302)
	// 		return
	// 	}
	// }
}

//首页轮播图片给予任何权限
func (c *AttachController) GetCarousel() {
	//1.url处理中文字符路径，[1:]截掉路径前面的/斜杠
	// filePath := path.Base(ctx.Request.RequestURI)
	filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:]) //  attachment/SL2016测试添加成果/A/FB/1/Your First Meteor Application.pdf
	if err != nil {
		beego.Error(err)
	}
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)
}

//返回文件大小
func FileSize(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

//根据侧栏id返回附件url和文件夹路径
func GetUrlPath(id int64) (Url, DiskDirectory string, err error) {
	var parentidpath, parentidpath1 string
	proj, err := models.GetProj(id)
	if err != nil {
		beego.Error(err)
		return "", "", err
	} else {
		//根据proj的parentIdpath
		var path string
		if proj.ParentIdPath != "" { //如果不是根目录
			// patharray := strings.Split(proj.ParentIdPath, "-")
			parentidpath = strings.Replace(strings.Replace(proj.ParentIdPath, "#$", "-", -1), "$", "", -1)
			parentidpath1 = strings.Replace(parentidpath, "#", "", -1)
			patharray := strings.Split(parentidpath1, "-")
			for _, v := range patharray {
				//pid转成64为
				idNum, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					beego.Error(err)
				}
				proj1, err := models.GetProj(idNum)
				if err != nil {
					beego.Error(err)
				} else {
					if proj1.ParentId == 0 { //如果是项目名称，则加上项目编号
						DiskDirectory = "./attachment/" + proj1.Code + proj1.Title
						Url = "attachment/" + proj1.Code + proj1.Title
					} else {
						path = proj1.Title
						DiskDirectory = DiskDirectory + "/" + path
						Url = Url + "/" + path
					}
				}
			}
			DiskDirectory = DiskDirectory + "/" + proj.Title //加上自身
			Url = Url + "/" + proj.Title
		} else { //如果是根目录
			DiskDirectory = "./attachment/" + proj.Code + proj.Title //加上自身
			Url = "attachment/" + proj.Code + proj.Title
		}
		return Url, DiskDirectory, err
	}
}

//根据id返回项目编号，项目名称，项目阶段，项目专业
func GetProjTitleNumber(id int64) (ProjectNumber, ProjectName, DesignStage, Section string, err error) {
	var parentidpath, parentidpath1 string
	proj, err := models.GetProj(id)
	if err != nil {
		beego.Error(err)
		return "", "", "", "", err
	} else {
		//根据proj的parentIdpath
		if proj.ParentIdPath != "" { //如果不是根目录
			parentidpath = strings.Replace(strings.Replace(proj.ParentIdPath, "#$", "-", -1), "$", "", -1)
			parentidpath1 = strings.Replace(parentidpath, "#", "", -1)
			patharray := strings.Split(parentidpath1, "-")
			// patharray := strings.Split(proj.ParentIdPath, "-")
			//pid转成64位
			meritNum, err := strconv.ParseInt(patharray[0], 10, 64)
			if err != nil {
				beego.Error(err)
			}
			meritproj, err := models.GetProj(meritNum)
			if err != nil {
				beego.Error(err)
			}
			ProjectNumber = meritproj.Code
			ProjectName = meritproj.Title
			if len(patharray) > 1 {
				//pid转成64位
				meritNum1, err := strconv.ParseInt(patharray[1], 10, 64)
				if err != nil {
					beego.Error(err)
				}
				meritproj1, err := models.GetProj(meritNum1)
				if err != nil {
					beego.Error(err)
				}
				DesignStage = meritproj1.Title
			}

			if len(patharray) > 2 {
				//pid转成64位
				meritNum2, err := strconv.ParseInt(patharray[2], 10, 64)
				if err != nil {
					beego.Error(err)
				}
				meritproj2, err := models.GetProj(meritNum2)
				if err != nil {
					beego.Error(err)
				}
				Section = meritproj2.Title
			}

		} else { //如果是根目录
			ProjectNumber = proj.Code
			ProjectName = proj.Title
		}
		return ProjectNumber, ProjectName, DesignStage, Section, err
	}
}

//编码转换
// l3, err3 := url.Parse(c.Ctx.Request.RequestURI[1:])
// 	if err3 != nil {
// 		beego.Error(err3)
// 	}
// 	beego.Info(l3.Query())
// 	beego.Info(l3.Query().Encode())
// 	beego.Info(url.QueryEscape("初步设计"))
// 	dec := mahonia.NewDecoder("utf8") //定义转换乱码
// 	//保留	fmt.Println(dec.ConvertString(j))   //转成utf-8
// 	beego.Info(dec.ConvertString(url.QueryEscape("初步设计")))
// 	// rd := dec.NewReader(resp.Body)
// 	// doc, err := goquery.NewDocumentFromReader(rd)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	bn := []byte("\xbf\xc6\xd1\xa7\xC3\xF1\xD6\xF7\xCF\xDC\xD5\xFE")
// 	beego.Info(bn)
// 	sn := string(bn)
// 	beego.Info(sn)
// 	cbn, err1, _, _ := ConvertGB2312(bn)
// 	if err1 != nil {
// 		beego.Error(err1)
// 	}
// 	beego.Info(cbn)
// 	csn, err2, _, _ := ConvertGB2312String(c.Ctx.Request.RequestURI[1:])
// 	if err2 != nil {
// 		beego.Error(err2)
// 	}
// 	beego.Info(csn)
// 	//根据路由path.Dir——再转成数组strings.Split——查出项目id——加上名称——查出下级id
// 	beego.Info(path.Dir(filePath))
// 	beego.Info(c.Ctx.Request.RequestURI[1:])
