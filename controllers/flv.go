package controllers

import (
	"github.com/astaxie/beego"
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
