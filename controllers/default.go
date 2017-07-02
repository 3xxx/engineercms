package controllers

import (
	"github.com/astaxie/beego"
	"strings"
	"time"
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
	const lll = "2006-01-02"
	// convdate := time.Now().Format(lll)
	// t1, err := time.Parse(lll, convdate) //这里t1要是用t1:=就不是前面那个t1了
	// if err != nil {
	// 		beego.Error(err)
	// 	}
	date := time.Now()
	convdate := string(date.Format(lll))

	f, _, err := c.GetFile("uploadfile")
	// beego.Info(h) //这里 filename是路径，所以不能以filename作为保存的文件名。坑！！
	defer f.Close()

	if err != nil {
		beego.Error(err)
	} else {
		c.SaveToFile("uploadfile", "./static/upload/"+convdate+".data") // 保存位置在 static/upload, 没有文件夹要先创建
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
