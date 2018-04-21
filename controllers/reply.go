package controllers

import (
	"github.com/3xxx/engineercms/models"
	"github.com/astaxie/beego"
)

type ReplyController struct {
	beego.Controller
}

//添加文章评论
func (c *ReplyController) Add() {
	// if !checkAccount(c.Ctx) {
	// 	c.Redirect("/login", 302)
	// 	return
	// }
	// //由文章Id获取文章author
	// role, _ := checkRole(c.Ctx)
	// if role != "1" { //or（||或） 文章的作者不等于user
	// 	c.Redirect("/roleerr", 302)
	// 	return
	// }
	tid := c.Input().Get("tid") //tid是文章id
	// content := c.Input().Get("editorValue")
	//获取用户名，如果已经登录，则nickname用登录名
	//4.取得客户端用户名
	var uname string
	v := c.GetSession("uname")
	// var role, userrole int
	if v != nil {
		uname = v.(string)
		c.Data["Uname"] = v.(string)
		// user, err := models.GetUserByUsername(uname)
		// if err != nil {
		// 	beego.Error(err)
		// }
		// userrole = user.Role
	} else {
		uname = c.Input().Get("nickname")
		// userrole = 5
	}

	err := models.AddTopicReply(tid, uname, c.Input().Get("editorValue"))
	if err != nil {
		beego.Error(err)
	}
	op := c.Input().Get("op")
	switch op {
	case "b":
		c.Redirect("/topic/view_b/"+tid, 302)
	case "c":
		c.Redirect("/wiki/view/"+tid, 302)
	default:
		c.Redirect("/topic/view/"+tid, 302)
	}
	// if op == "b" {
	// 	c.Redirect("/topic/view_b/"+tid, 302)
	// } else {
	// 	c.Redirect("/topic/view/"+tid, 302)
	// }
}

//删除文章评论
func (c *ReplyController) Delete() {
	if !checkAccount(c.Ctx) {
		return
	}
	tid := c.Input().Get("tid")
	err := models.DeleteTopicReply(c.Input().Get("rid"))
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/topic/view/"+tid, 302)
}

//添加wiki评论
func (c *ReplyController) AddWiki() {
	// if !checkAccount(c.Ctx) {
	// 	c.Redirect("/login", 302)
	// 	return
	// }
	// //由文章Id获取文章author
	// role, _ := checkRole(c.Ctx)
	// if role != "1" { //or（||或） 文章的作者不等于user
	// 	c.Redirect("/roleerr", 302)
	// 	return
	// }
	tid := c.Input().Get("tid") //tid是文章id
	// content := c.Input().Get("editorValue")
	//获取用户名，如果已经登录，则nickname用登录名
	//4.取得客户端用户名
	var uname string
	v := c.GetSession("uname")
	// var role, userrole int
	if v != nil {
		uname = v.(string)
		c.Data["Uname"] = v.(string)
		// user, err := models.GetUserByUsername(uname)
		// if err != nil {
		// 	beego.Error(err)
		// }
		// userrole = user.Role
	} else {
		uname = c.Input().Get("nickname")
		// userrole = 5
	}

	err := models.AddWikiReply(tid, uname, c.Input().Get("editorValue"))
	if err != nil {
		beego.Error(err)
	}
	op := c.Input().Get("op")
	switch op {
	case "b":
		c.Redirect("/topic/view_b/"+tid, 302)
	case "c":
		c.Redirect("/wiki/view/"+tid, 302)
	default:
		c.Redirect("/topic/view/"+tid, 302)
	}
	// if op == "b" {
	// 	c.Redirect("/topic/view_b/"+tid, 302)
	// } else {
	// 	c.Redirect("/topic/view/"+tid, 302)
	// }
}

//删除wiki评论
func (c *ReplyController) DeleteWiki() {
	if !checkAccount(c.Ctx) {
		return
	}
	tid := c.Input().Get("tid")
	err := models.DeleteWikiReply(c.Input().Get("rid"))
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/wiki/view/"+tid, 302)
}
