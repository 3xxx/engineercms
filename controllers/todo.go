package controllers

import (
	"github.com/3xxx/engineercms/controllers/utils"
	// "crypto/md5"
	// "encoding/hex"
	// "encoding/json"
	"github.com/3xxx/engineercms/models"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego/logs"
	// "net"
	// "net/http"
	// "net/url"
	// "path"
	"strconv"
	// "strings"
	// "fmt"
	// "reflect"
	// "sort"
	// "time"
)

// CMSTODO API
type TodoController struct {
	beego.Controller
}

// @Title post todo
// @Description post todo
// @Param name query string true "The name for todo"
// @Param projectid query string true "The projectid of todo"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /create [post]
func (c *TodoController) Create() {
	var userid int64
	openID := c.GetSession("openID")
	if openID != nil {
		user, err := models.GetUserByOpenID(openID.(string))
		if err != nil {
			beego.Error(err)
		}
		userid = user.Id
	} else {
		c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		c.ServeJSON()
		return
		// user.Id = 9
	}

	name := c.Input().Get("name")
	projectid := c.Input().Get("projectid")
	ProjectId, err := strconv.ParseInt(projectid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	// 进行敏感字符验证
	app_version := c.Input().Get("app_version")
	accessToken, _, _, err := utils.GetAccessToken(app_version)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": err}
		c.ServeJSON()
		return
	}
	errcode, errmsg, err := utils.MsgSecCheck(accessToken, name)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": err}
		c.ServeJSON()
	} else if errcode != 87014 {
		todoid, err := models.TodoCreate(ProjectId, name, userid)
		if err != nil {
			beego.Error(err)
			c.Data["json"] = map[string]interface{}{"message": "写入数据库出错"}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "message": "ok", "todoid": todoid}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": errmsg}
		c.ServeJSON()
	}
}

// @Title get todolist
// @Description get tolist
// @Param projectid query string true "The projectid of todo"
// @Param page query string false "The page of todo"
// @Param limit query string false "The size of todo"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /gettodo [get]
//取出所有未完成待办
func (c *TodoController) GetTodo() {
	var offset, limit1, page1 int
	var err error
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 20
	} else {
		limit1, err = strconv.Atoi(limit)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		// limit1 = 10
		page1 = 1
	} else {
		page1, err = strconv.Atoi(page)
		if err != nil {
			beego.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}
	projectid := c.Input().Get("projectid")
	ProjectId, err := strconv.ParseInt(projectid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	todos, err := models.GetTodoUser(ProjectId, limit1, offset)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = todos //map[string]interface{}{"userId": 1, "avatorUrl": "Filename"}
	c.ServeJSON()
}

// @Title update todolist
// @Description update tolist
// @Param todoid query string false "The id of todo"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /updatetodo [post]
// 更新待办状态
func (c *TodoController) UpdateTodo() {
	openID := c.GetSession("openID")
	if openID != nil {
	} else {
		c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		c.ServeJSON()
		return
		// user.Id = 9
	}
	id := c.Input().Get("todoid")
	//pid转成64为
	todoid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	err = models.UpdateTodo(todoid)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"message": "更新数据库错误"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"todoid": id, "message": "ok"}
		c.ServeJSON()
	}
}

// @Title delete todolist
// @Description delete tolist
// @Param todoid query string false "The id of todo"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /deletetodo [post]
// 删除待办
func (c *TodoController) DeleteTodo() {
	id := c.Input().Get("todoid")
	//pid转成64为
	todoid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	err = models.DeleteTodo(todoid)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"message": "删除数据错误"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"todoid": id, "message": "ok"}
		c.ServeJSON()
	}
}
