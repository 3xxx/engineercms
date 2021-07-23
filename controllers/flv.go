package controllers

import (
	"github.com/3xxx/engineercms/controllers/utils"
	"github.com/astaxie/beego"
	"strconv"
)

// CMSFLV API
type FlvController struct {
	beego.Controller
}

// @Title getFlv
// @Description get admin page
// @Success 200 {object} success
// @Failure 400 Invalid page
// @Failure 404 page not found
// @router /flvlist [get]
func (c *FlvController) GetFlvList() {
	token := c.Input().Get("token")
	site := c.Ctx.Input.Site() + ":" + strconv.Itoa(c.Ctx.Input.Port())
	if token == "" {
		// 	c.Redirect("http://localhost:8080/v1/sso/ssologin?service="+site+c.Ctx.Request.URL.String()+"token="+token, 301)
		// } else {
		token = c.Ctx.GetCookie("token")
		//
	}
	username, err := utils.CheckToken(token)
	beego.Info(username)
	if err != nil {
		c.Redirect("http://localhost:8080/v1/sso/ssologin?service="+site+c.Ctx.Request.URL.String(), 301)
	} else {
		c.Ctx.SetCookie("token", token, "3600", "/")
	}
	// username, err := utils.CheckToken(token)
	c.TplName = "flvlist.tpl"
}

// @Title getFlv
// @Description get admin page
// @Param filepath query string true "The mp4 file path"
// @Success 200 {object} success
// @Failure 400 Invalid page
// @Failure 404 page not found
// @router / [get]
func (c *FlvController) Get() {
	mp4link := c.Input().Get("filepath")
	c.Data["Mp4Link"] = mp4link
	c.TplName = "flv.tpl"
}
