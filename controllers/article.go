package controllers

import (
	// "encoding/json"
	"engineercms/models"
	"github.com/astaxie/beego"
	// "os"
	"strconv"
	// "strings"
)

type ArticController struct {
	beego.Controller
}

//提供给文章列表的table中json数据
func (c *ArticController) GetArticles() {
	id := c.Ctx.Input.Param(":id")
	if id == "" {
		//显示全部
		Articles, err := models.GetArticles()
		if err != nil {
			beego.Error(err)
		}
		c.Data["json"] = Articles
		c.ServeJSON()
	} else {
		//根据标签查询
	}
}

//根据ip查看项目，查出项目目录
func (c *ArticController) GetArticle() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	// var categories []*models.ProjCategory
	var err error
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//取项目本身
	category, err := models.GetProj(idNum)
	if err != nil {
		beego.Error(err)
	}
	//取项目所有子孙
	// categories, err := models.GetArticlesbyPid(idNum)
	// if err != nil {
	// 	beego.Error(err)
	// }
	//算出最大级数
	// grade := make([]int, 0)
	// for _, v := range categories {
	// 	grade = append(grade, v.Grade)
	// }
	// height := intmax(grade[0], grade[1:]...)
	//递归生成目录json
	root := FileNode{category.Id, category.Title, []*FileNode{}}
	walk(category.Id, &root)
	// beego.Info(root)
	// data, _ := json.Marshal(root)
	c.Data["json"] = root //data
	// c.ServeJSON()
	c.Data["Category"] = category
	c.TplName = "Article.tpl"
}
