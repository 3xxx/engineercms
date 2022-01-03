package controllers

import (
	// "fmt"
	// beego "github.com/beego/beego/v2/adapter"
	// "github.com/beego/beego/v2/adapter/logs"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/pagination"
	// "github.com/tealeg/xlsx"
	// "os"
	// "path"
	// "path/filepath"
	"github.com/3xxx/engineercms/models"
	// "regexp"
	"strconv"
	"strings"
)

type WikiController struct {
	web.Controller
}

func (c *WikiController) Get() { //这个给爬虫用。而为了配合pagenate，用后面的listall()
	// username, role := checkprodRole(c.Ctx)
	// roleint, err := strconv.Atoi(role)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// if role == "1" {
	// 	c.Data["IsAdmin"] = true
	// } else if roleint > 1 && roleint < 5 {
	// 	c.Data["IsLogin"] = true
	// } else {
	// 	c.Data["IsAdmin"] = false
	// 	c.Data["IsLogin"] = false
	// }
	// c.Data["Username"] = username
	c.Data["IsWiki"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid

	c.TplName = "wiki.tpl"
	// c.Data["IsLogin"] = checkAccount(c.Ctx)

	wikis, err := models.GetAllWikis(false) //这里传入空字符串
	if err != nil {
		logs.Error(err.Error)
	} else {
		count := len(wikis)
		count1 := strconv.Itoa(count)
		count2, err := strconv.ParseInt(count1, 10, 64)
		if err != nil {
			logs.Error(err)
		}
		postsPerPage := 20
		paginator := pagination.SetPaginator(c.Ctx, postsPerPage, count2)
		wikis, err = models.ListWikisByOffsetAndLimit(paginator.Offset(), postsPerPage)
		if err != nil {
			logs.Error(err)
		}
		c.Data["paginator"] = paginator
		c.Data["Wikis"] = wikis
		c.Data["Length"] = count
	}
	var replycounts int64
	for _, v1 := range wikis {
		replycounts = replycounts + v1.ReplyCount
	}
	c.Data["ReplyCounts"] = replycounts

	users, ucount := models.Getuserlist(1, 2000, "Id")
	if c.IsAjax() {
		c.Data["json"] = &map[string]interface{}{"total": ucount, "rows": &users}
		c.ServeJSON()
		return
	} else {
		// tree := this.GetTree()
		// this.Data["tree"] = &tree
		c.Data["Users"] = &users
		// this.TplName = "user.tpl"
		c.Data["UsersCounts"] = ucount
		// if this.GetTemplatetype() != "easyui" {
		// this.Layout = this.GetTemplatetype() + "/public/layout.tpl"
		// }
		// this.TplName = this.GetTemplatetype() + "/rbac/user.tpl"
	}
	//var err error
	//	c.Data["Wiki"], err = models.GetAllWikis()
	//	if err != nil {
	//		logs.Error(err)
	//	}
}

//根据用户名查看wiki
func (c *WikiController) Viewbyuname() {
	// username, role := checkprodRole(c.Ctx)
	// roleint, err := strconv.Atoi(role)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// if role == "1" {
	// 	c.Data["IsAdmin"] = true
	// } else if roleint > 1 && roleint < 5 {
	// 	c.Data["IsLogin"] = true
	// } else {
	// 	c.Data["IsAdmin"] = false
	// 	c.Data["IsLogin"] = false
	// }
	// c.Data["Username"] = username
	c.Data["IsWiki"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid

	c.TplName = "wiki_uname.tpl"

	// rolename, _ = strconv.Atoi(role)
	// c.Data["Uname"] = uname
	uname := c.GetString("uname")
	wiki, _ := models.Getwikisbyuname(uname) //由uname取出项目
	c.Data["Wikis"] = wiki
}

func (c *WikiController) Add() { //参考下面的 modify,这个add是wiki/add
	// username, role := checkprodRole(c.Ctx)
	// roleint, err := strconv.Atoi(role)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// if role == "1" {
	// 	c.Data["IsAdmin"] = true
	// } else if roleint > 1 && roleint < 5 {
	// 	c.Data["IsLogin"] = true
	// } else {
	// 	c.Data["IsAdmin"] = false
	// 	c.Data["IsLogin"] = false
	// }
	// c.Data["Username"] = username
	c.Data["IsWiki"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	// var rolename int
	// var uname string
	//2.如果登录或ip在允许范围内，进行访问权限检查
	// uname, role := checkprodRole(c.Ctx)
	// beego.Info(role)
	// rolename, _ = strconv.Atoi(role)
	// c.Data["Uname"] = uname

	if !islogin { //&& uname != category.Author
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	// c.Data["IsWiki"] = true
	c.TplName = "wiki_add.tpl"
}

//这个提交添加wiki的方法
func (c *WikiController) AddWiki() {
	c.Data["IsWiki"] = true
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid

	title := c.GetString("title")
	content := c.GetString("content")
	if !islogin { //&& uname != category.Author
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	id, err := models.AddWikiOne(title, content, username)
	if err != nil {
		logs.Error(err)
	} else {
		c.Data["json"] = id
		c.ServeJSON()
	}
	// c.Redirect("/wiki", 302)
}

//这个是微信小程序添加wiki的方法
func (c *WikiController) AddPic() {
	// c.Data["IsWiki"] = true
	// username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	// c.Data["Username"] = username
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	// c.Data["IsAdmin"] = isadmin
	// c.Data["IsLogin"] = islogin
	// c.Data["Uid"] = uid

	title := c.GetString("title")
	content := c.GetString("content")
	content = "<p>" + content + "</p>"
	imagesurl := c.GetString("images")
	array := strings.Split(imagesurl, ",")
	for _, v := range array {
		content = content + "<p><img src='" + v + "'></p>"
	}
	// if !islogin { //&& uname != category.Author
	// 	// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/roleerr?url="+route, 302)
	// 	// c.Redirect("/roleerr", 302)
	// 	return
	// }
	// <p><img src="/attachment/wiki/2018January/1515026559287477900.jpg" title="Snap2.jpg" alt="Snap2.jpg" class="fr-fic fr-dii"></p>
	id, err := models.AddWikiOne(title, content, "test")
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERR", "id": id}
		c.ServeJSON()
	} else {
		// c.Data["json"] = id
		c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "id": id}
		c.ServeJSON()
	}
	// c.Redirect("/wiki", 302)
}

func (c *WikiController) Wiki_many_addbaidu() { //一对多模式
	// var uname string
	//解析表单
	tid := c.GetString("tid") //这句没用。教程里漏了这句，导致修改总是变成添加文章
	title := c.GetString("title")
	content := c.GetString("content")
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		logs.Error(err)
	}
	// var attachment string
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		// attachment := h.Filename
		// beego.Info(attachment)
		// path =  + categoryproj.Number + categoryproj.Title + "/" + categoryphase.Title + "/" + categoryspec.Title + "/" + category + "/" + h.Filename
		path = "./attachment/wiki/" + h.Filename
		// path := c.GetString("url")  //存文件的路径
		// path = path[3:]
		// path = "./attachment" + "/" + h.Filename
		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			logs.Error(err)
		}
		filesize, _ = FileSize(path)
		filesize = filesize / 1000.0
	}
	//获取用户名
	// uname, _ := checkprodRole(c.Ctx)
	// // rolename, _ = strconv.Atoi(role)
	// c.Data["Uname"] = uname

	// username, _, _, _, _ := checkprodRole(c.Ctx)
	// c.Data["Username"] = username

	// route := "/attachment/" + categoryproj.Number + categoryproj.Title + "/" + categoryphase.Title + "/" + categoryspec.Title + "/" + category + "/" + h.Filename
	// route := "/attachment/wiki/" + h.Filename
	//wikiid := c.GetString("wikiid")
	// var wikiid int64
	if len(tid) == 0 {
		// wikiid, err := models.AddWikiMany(title, uname, content, attachment)
		//这里返回wikiid，并存入attachment表中
		// if err != nil { //如果发生错误，返回错误，并获取该文章的wikiid
		// 	logs.Error(err)
		// }
		// cid := strconv.FormatInt(wikiid, 10)
		// filesize := strconv.FormatInt(filesize, 10)
		// err = models.AddAttachment(attachment, filesize, path, route, cid, uname)
		// if err != nil {
		// 	logs.Error(err)
		// }
	} else { //用这种结合的方式不好，因为uploader先上传附件
		err = models.ModifyWiki(tid, title, content)
		if err != nil {
			logs.Error(err) //return multi rows
		}
		// filesize := strconv.FormatInt(filesize, 10)
		// err = models.AddAttachment(attachment, filesize, path, route, tid, uname)
		// if err != nil {
		// 	logs.Error(err)
		// }
	}
	c.TplName = "wiki_add.tpl" //不加这句上传出错，虽然可以成功上传
	// c.Redirect("/wiki", 302)
}

func (c *WikiController) View() {
	// username, role := checkprodRole(c.Ctx)
	// roleint, err := strconv.Atoi(role)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// if role == "1" {
	// 	c.Data["IsAdmin"] = true
	// } else if roleint > 1 && roleint < 5 {
	// 	c.Data["IsLogin"] = true
	// } else {
	// 	c.Data["IsAdmin"] = false
	// 	c.Data["IsLogin"] = false
	// }
	// c.Data["Username"] = username
	c.Data["IsWiki"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid

	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	// c.Data["IsWiki"] = true
	// uname, _ := checkprodRole(c.Ctx)
	// rolename, _ = strconv.Atoi(role)
	// c.Data["Uname"] = uname
	c.TplName = "wiki_view.tpl"
	wiki, err := models.GetWiki(c.Ctx.Input.Param("0"))
	if err != nil {
		logs.Error(err)
		c.Redirect("/", 302)
		return
	}
	c.Data["Wiki"] = wiki
	// beego.Info(wiki)

	// c.Data["Attachment"] = attachment
	c.Data["Tid"] = c.Ctx.Input.Param("0") //教程中用的是圆括号，导致错误wiki.go:52: cannot call non-function c.Controller.Ctx.Input.Params (type map[string]string)
	//教程第8章开头有修改
	wid := c.Ctx.Input.Param("0")
	widNum, err := strconv.ParseInt(wid, 10, 64)
	replies, err := models.GetAllWikiReplies(widNum)
	if err != nil {
		logs.Error(err)
		return
	}
	c.Data["Replies"] = replies
}

//修改wiki页面
func (c *WikiController) Modify() { //这个也要登陆验证
	// username, role := checkprodRole(c.Ctx)
	// roleint, err := strconv.Atoi(role)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// if role == "1" {
	// 	c.Data["IsAdmin"] = true
	// } else if roleint > 1 && roleint < 5 {
	// 	c.Data["IsLogin"] = true
	// } else {
	// 	c.Data["IsAdmin"] = false
	// 	c.Data["IsLogin"] = false
	// }
	// c.Data["Username"] = username
	c.Data["IsWiki"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	tid := c.GetString("tid")
	// beego.Info(tid)
	//2.取得文章的作者
	wiki, err := models.GetWiki(tid)
	if err != nil {
		logs.Error(err)
		c.Redirect("/", 302)
		return
	}
	// var rolename int
	// var uname string
	//2.如果登录或ip在允许范围内，进行访问权限检查
	// uname, role := checkprodRole(c.Ctx)
	// rolename, _ = strconv.Atoi(role)
	// c.Data["Uname"] = uname
	roleint, err := strconv.Atoi(role)
	if err != nil {
		logs.Error(err)
	}
	if roleint > 2 && username != wiki.Author { //
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		return
	}

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	c.TplName = "wiki_modify.html"

	c.Data["Wiki"] = wiki
	c.Data["Tid"] = tid
	c.Data["IsWiki"] = true
}

func (c *WikiController) Post() { //这个post属于wiki_modify.html提交修改。
	//解析表单
	tid := c.GetString("tid") //教程里漏了这句，导致修改总是变成添加文章
	title := c.GetString("title")
	//其实这里只修改title, tnumber,和content
	content := c.GetString("content")
	err := models.ModifyWiki(tid, title, content)
	// }
	if err != nil {
		logs.Error(err)
	} else {
		c.Data["json"] = tid
		c.ServeJSON()
	}
	// c.Redirect("/wiki/view/"+tid, 302) //回到修改后的文章
}

//删除文章
func (c *WikiController) Delete() { //应该显示警告
	url := c.GetString("url")
	c.Data["IsWiki"] = true
	//2.取得文章的作者
	wiki, err := models.GetWiki(c.GetString("tid"))
	if err != nil {
		logs.Error(err)
		c.Redirect("/", 302)
		return
	}
	// var rolename int
	// var uname string
	//2.如果登录或ip在允许范围内，进行访问权限检查
	// uname, role := checkprodRole(c.Ctx)
	username, role, _, _, _ := checkprodRole(c.Ctx)
	c.Data["role"] = role
	roleint, err := strconv.Atoi(role)
	if err != nil {
		logs.Error(err)
	}
	// rolename, _ = strconv.Atoi(role)
	// beego.Info(rolename)=5
	// beego.Info(wiki.Author)=127.0.0.1
	// c.Data["Uname"] = uname
	if roleint > 2 && username != wiki.Author { //
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	err = models.DeletWiki(c.GetString("tid")) //(c.Ctx.Input.Param("0"))
	if err != nil {
		logs.Error(err)
	} else { //没有这个返回值，会出现错误提示。wsasend: An established connection was aborted by the software in your host machine
		data := wiki.Title
		c.Ctx.WriteString(data)
	}
	c.Redirect(url, 302) //这里增加wiki
}

//删除文章中的附件，保持页面不跳转怎么办？
func (c *WikiController) DeleteAttachment() { //应该显示警告
	//2.取得文章的作者
	wiki, err := models.GetWiki(c.GetString("tid"))
	if err != nil {
		logs.Error(err)
		c.Redirect("/", 302)
		return
	}
	// var rolename int
	// var uname string
	//2.如果登录或ip在允许范围内，进行访问权限检查
	// uname, role := checkprodRole(c.Ctx)
	username, role, _, _, _ := checkprodRole(c.Ctx)
	roleint, err := strconv.Atoi(role)
	if err != nil {
		logs.Error(err)
	}
	// rolename, _ = strconv.Atoi(role)
	// c.Data["Uname"] = uname
	if roleint > 2 && username != wiki.Author { //
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// Tid := c.Ctx.Input.Param("0")
	Tid := c.GetString("tid")
	// beego.Info(Tid)
	err = models.DeletAttachment(c.GetString("aid")) //(c.Ctx.Input.Param("0"))
	if err != nil {
		logs.Error(err)
	}
	op := c.GetString("op")
	switch op {
	case "modify":
		c.Redirect("/wiki/modify?tid="+Tid, 302)
	default:
		c.Redirect("/wiki/view/"+Tid, 302) //这里增加wiki
	}
}

// func FileSize(file string) (int64, error) {
// 	f, e := os.Stat(file)
// 	if e != nil {
// 		return 0, e
// 	}
// 	return f.Size(), nil
// }

// func (c *WikiController) ListAllWikis() {
// 	c.Data["IsWiki"] = true
// 	c.TplName = "wiki.tpl"
// 	c.Data["IsLogin"] = checkAccount(c.Ctx)
// 	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
// 	defer sess.SessionRelease(c.Ctx.ResponseWriter)
// 	v := sess.Get("uname")
// 	if v != nil {
// 		c.Data["Uname"] = v.(string)
// 	}
// 	wikis, err := models.GetAllWikis(false)
// 	if err != nil {
// 		logs.Error(err)
// 	}
// 	count := len(wikis)
// 	count1 := strconv.Itoa(count)
// 	count2, err := strconv.ParseInt(count1, 10, 64)
// 	if err != nil {
// 		logs.Error(err)
// 	}
// 	postsPerPage := 20
// 	paginator := pagination.SetPaginator(c.Ctx, postsPerPage, count2)
// 	wikis, err = models.ListPostsByOffsetAndLimit(paginator.Offset(), postsPerPage)
// 	if err != nil {
// 		logs.Error(err)
// 	}
// 	c.Data["Wikis"] = wikis
// 	c.Data["paginator"] = paginator
// }
