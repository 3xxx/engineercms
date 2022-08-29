package controllers

import (
	// "github.com/3xxx/engineercms/controllers/utils"
	// "github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	// beego "github.com/beego/beego/v2/adapter"
	// "strconv"
	// "strings"
)

// CMSFLV API
type FlvController struct {
	web.Controller
}

// @Title getFlv
// @Description get admin page
// @Param xxl_sso_token query string false "The tokenText"
// @Success 200 {object} success
// @Failure 400 Invalid page
// @Failure 404 page not found
// @router /flvlist [get]
func (c *FlvController) GetFlvList() {
	// var token string
	// token = c.Ctx.GetCookie("token")
	// site := c.Ctx.Input.Site() + ":" + strconv.Itoa(c.Ctx.Input.Port())
	// if token != "" {
	// 	urlarray := strings.Split(c.Ctx.Request.URL.String(), "?")
	// 	if len(urlarray) > 1 {
	// 		c.Redirect(strings.Split(c.Ctx.Request.URL.String(), "?")[0], 302)
	// 	} else {
	// 		userid, username, usernickname, err := utils.LubanCheckToken(token)
	// 		if err != nil {
	// 			logs.Error(err)
	// 		}
	// 		c.Ctx.SetCookie("token", token, "3600", "/")
	// 		c.SetSession("uname", username)
	// 		c.SetSession("userid", userid)
	// 		c.SetSession("usernickname", usernickname)
	// 	}
	// } else {
	// 	token = c.GetString("xxl_sso_token")
	// 	if token == "" {
	// 		c.Redirect("https://www.54lby.com/sso/login?redirect_url="+site+c.Ctx.Request.URL.String(), 302)
	// 	} else {
	// 		userid, username, usernickname, err := utils.LubanCheckToken(token)
	// 		if err != nil {
	// 			logs.Error(err)
	// 		}
	// 		c.Ctx.SetCookie("token", token, "3600", "/")
	// 		c.SetSession("uname", username)
	// 		c.SetSession("userid", userid)
	// 		c.SetSession("usernickname", usernickname)
	// 		urlarray := strings.Split(c.Ctx.Request.URL.String(), "?")
	// 		if len(urlarray) > 1 {
	// 			c.Redirect(strings.Split(c.Ctx.Request.URL.String(), "?")[0], 302)
	// 		}
	// 	}
	// }

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
	mp4link := c.GetString("filepath")
	c.Data["Mp4Link"] = mp4link
	c.TplName = "flv.tpl"
}
