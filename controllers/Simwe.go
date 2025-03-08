package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

// CMSSIMWE API
type SimweController struct {
	web.Controller
}

// @Title get simansys
// @Description getsimansys
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /simansys [get]
// simansys页面
func (c *SimweController) SimAnsys() {
	c.TplName = "simwe/simansys.tpl"
}

// @Title get simwe
// @Description getsimwe
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /simwe [get]
// simwe页面
func (c *SimweController) Simwe() {
	c.TplName = "simwe/simwe.tpl"
}

// @Title get threejs
// @Description getthreejs
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /threejs [get]
// simwe页面
func (c *SimweController) Threejs() {
	c.TplName = "simwe/threejs.tpl"
}
