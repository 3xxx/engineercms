package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	// "github.com/astaxie/beego/httplib"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "projects.tpl"
}

func (c *MainController) Test() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "test.tpl"
}

func (c *MainController) Slide() {
	c.TplName = "slide.tpl"
}

func (c *MainController) Postdata() {
	f, h, err := c.GetFile("uploadfile1")
	beego.Info(h) //这里 filename是路径，所以不能以filename作为保存的文件名。坑！！
	defer f.Close()
	if err != nil {
		beego.Error(err)
	} else {
		c.SaveToFile("uploadfile1", "./static/upload/1.txt") // 保存位置在 static/upload, 没有文件夹要先创建
		c.Ctx.WriteString("ok")
	}
}

func (c *MainController) Register() {
	// flash := beego.NewFlash()
	token := c.Input().Get("token")
	//是否重复提交
	if c.IsSubmitAgain(token) {
		c.Redirect("/registerpage", 302)
		return
	}
}

func (c *MainController) IsSubmitAgain(token string) bool {
	cotoken := c.Ctx.GetCookie("token")
	if token == "" || len(token) == 0 || token != cotoken || strings.Compare(cotoken, token) != 0 {
		return true
	}
	return false
}
