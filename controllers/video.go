package controllers

import (
	"github.com/3xxx/engineercms/models"
	"github.com/astaxie/beego"
	// "io/ioutil"
	// "net"
	// "net/http"
	// "net/url"
	// "os"
	// "path"
	"strconv"
	// "strings"
	// "time"
)

// CMSVideo API
type VideoController struct {
	beego.Controller
}

// @Title get uservideolist
// @Description get uservideolist
// @Param id path string true "The id of project"
// @Param searchText query string false "The searchText of uservideo"
// @Param pageNo query string true "The page for uservideo list"
// @Param limit query string true "The limit of page for uservideo list"
// @Success 200 {object} models.Create
// @Failure 400 Invalid page supplied
// @Failure 404 cart not found
// @router /getuservideo/:id [get]
//根据用户id获得借阅记录，如果是admin角色，则查询全部
func (c *VideoController) GetUserVideo() {
	//解析表单
	pid := c.Ctx.Input.Param(":id")
	// pid := beego.AppConfig.String("wxcatalogid") //"26159" //c.Input().Get("pid")
	//pid转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	// v := c.GetSession("uname")
	// var user models.User
	// var userid, roleid string
	// var err error
	// if v != nil { //如果登录了
	// 	user, err = models.GetUserByUsername(v.(string))
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
		//查询admin角色的id
		//重新获取roleid
		// role, err := models.GetRoleByRolename("admin")
		// if err != nil {
		// 	beego.Error(err)
		// }
		// userid = strconv.FormatInt(user.Id, 10)
		// roleid = strconv.FormatInt(role.Id, 10)

	// } else {
	// 	c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
	// 	c.ServeJSON()
	// 	return
	// }

	searchText := c.Input().Get("searchText")

	limit := c.Input().Get("limit")
	if limit == "" {
		limit = "15"
	}
	limit1, err := strconv.Atoi(limit)
	if err != nil {
		beego.Error(err)
	}
	page := c.Input().Get("pageNo")
	page1, err := strconv.Atoi(page)
	if err != nil {
		beego.Error(err)
	}
	var offset int
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

				//根据pid查出项目id
	proj, err := models.GetProj(pidNum)
	if err != nil {
		beego.Error(err)
	}

	videos, err := models.GetUserVideo(proj.Id, limit1, offset, searchText)
	if err != nil {
		beego.Error(err)
	}
	count, err := models.GetUserVideoCount(proj.Id, searchText)
	if err != nil {
		beego.Error(err)
	}
	// c.Data["json"] = carts
	c.Data["json"] = map[string]interface{}{"page": page1, "total": count, "rows": videos}

	// c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "mycarts": carts}
	c.ServeJSON()
}