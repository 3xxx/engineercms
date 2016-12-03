package controllers

import (
	"engineercms/models"
	"github.com/astaxie/beego"
	"strconv"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "admin.tpl"
}

func (c *AdminController) Admin() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	switch id {
	case "00": //日历
		c.TplName = "admin_calendar.tpl"
	case "01": //目录设置

		c.TplName = "admin_category.tpl"
	case "02": //爬虫设置
		c.TplName = "admin_spider.tpl"
	case "03": //项目权限
		c.TplName = "admin_project.tpl"
	case "04": //账号设置
		c.TplName = "admin_project.tpl"
	case "05": //
		c.TplName = "admin_project.tpl"
	case "06": //
		c.TplName = "admin_project.tpl"
	}
}

//根据数字id或空查询分类，如果有pid，则查询下级，如果pid为空，则查询类别
func (c *AdminController) Category() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	// var categories []*models.AdminCategory
	var err error
	if id == "" { //如果id为空，则查询类别
		id = "0"
	}
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	categories, err := models.GetAdminCategory(idNum)
	if err != nil {
		beego.Error(err)
	}

	c.Data["json"] = categories
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

//根据名称title查询分级表
func (c *AdminController) CategoryTitle() {
	// title := c.Ctx.Input.Param(":id")
	title := c.Input().Get("title")
	// beego.Info(title)
	categories, err := models.GetAdminCategoryTitle(title)
	// beego.Info(categories)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = categories
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

func (c *AdminController) AddCategory() {
	// pid := c.Ctx.Input.Param(":id")
	pid := c.Input().Get("pid")
	title := c.Input().Get("title")
	code := c.Input().Get("code")
	grade := c.Input().Get("grade")
	//pid转成64为
	var pidNum int64
	var err error
	if pid != "" {
		pidNum, err = strconv.ParseInt(pid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	} else {
		pidNum = 0
	}
	gradeNum, err := strconv.Atoi(grade)
	if err != nil {
		beego.Error(err)
	}
	_, err = models.AddAdminCategory(pidNum, title, code, gradeNum)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

func (c *AdminController) UpdateCategory() {
	// pid := c.Ctx.Input.Param(":id")
	cid := c.Input().Get("cid")
	title := c.Input().Get("title")
	code := c.Input().Get("code")
	grade := c.Input().Get("grade")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	gradeNum, err := strconv.Atoi(grade)
	if err != nil {
		beego.Error(err)
	}
	err = models.UpdateAdminCategory(cidNum, title, code, gradeNum)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}
