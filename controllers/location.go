// Location Trip
package controllers

import (
	// "encoding/json"
	"github.com/3xxx/engineercms/controllers/utils"
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	// beego "github.com/beego/beego/v2/adapter"
	// "sort"
	"strconv"
	// "strings"
	// "time"
)

type LocationController struct {
	web.Controller
}

// type wxuser struct {
// 	Name string `json:"name"`
// }

// @Title post location by projectid
// @Description post location by projectid
// @Param title query string  true "The title of locationpart"
// @Param describe query string  true "The describe of locationpart"
// @Param sort query string  true "The sort of locationpart"
// @Param userid query string true "The userid of location"
// @Param id path string true "The projectid of project"
// @Success 200 {object} models.CreateLocation
// @Failure 400 Invalid page supplied
// @Failure 404 pas not found
// @router /addlocationpart/:id [post]
// 添加一个定位组
func (c *LocationController) AddLocationPart() {
	//content去验证
	app_version := c.GetString("app_version")
	accessToken, _, _, err := utils.GetAccessToken(app_version)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": err}
		c.ServeJSON()
	}
	var user models.User
	// var err error
	openID := c.GetSession("openID")
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			logs.Error(err)
		}
	} else {
		c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		c.ServeJSON()
		return
		// user.Id = 9
	}

	pid := c.Ctx.Input.Param(":id")
	//id转成64为
	projectid, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		logs.Error(err)
	}

	title := c.GetString("title")
	describe := c.GetString("describe")
	sort := c.GetString("sort")
	sortint, err := strconv.Atoi(sort)
	if err != nil {
		logs.Error(err)
	}
	// 进行敏感字符验证
	content := title + describe
	// errcode, errmsg, err := utils.MsgSecCheck(accessToken, content)
	suggest, label, err := utils.MsgSecCheck(2, 2, accessToken, openID.(string), content)
	// errcode, errmsg, err := utils.MsgSecCheck(2, 2, accessToken, openID.(string), content)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": err}
		c.ServeJSON()
	} else if suggest == "pass" {
		// } else if errcode != 87014 {
		//根据项目id添加定位组
		var location models.Location
		// 添加location
		location = models.Location{
			ProjectID: projectid,
			UserID:    user.Id,
			Title:     title,
			Describe:  describe,
			Sort:      sortint,
		}
		Id, err := models.CreateLocation(location)

		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"data": "WRONG", "info": "添加定位组数据错误"}
			c.ServeJSON()
		} else if Id != 0 {
			c.Data["json"] = map[string]interface{}{"data": "OK", "info": "SUCCESS", "id": Id}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"data": "data already exist", "info": "数据已存在"}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 3, "info": suggest, "data": label}
		c.ServeJSON()
	}
}

// @Title post locationnavigate by locationid
// @Description post locationnavigate by locationid
// @Param title query string true "The title of location"
// @Param label query string true "The label of locationnavigate"
// @Param location query string true "The location of location"
// @Param address query string false "The address of location"
// @Param lat query string false "The lat of location"
// @Param lng query string false "The lng of location"
// @Param id path string true "The locationid of location"
// @Success 200 {object} models.CreateLocation
// @Failure 400 Invalid page supplied
// @Failure 404 pas not found
// @router /addlocationnavigate/:id [post]
// 添加一个定位
func (c *LocationController) AddLocationNavigate() {
	//content去验证
	app_version := c.GetString("app_version")
	accessToken, _, _, err := utils.GetAccessToken(app_version)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": err}
		c.ServeJSON()
	}
	// var err error
	openID := c.GetSession("openID")
	if openID != nil {
		_, err := models.GetUserByOpenID(openID.(string))
		if err != nil {
			logs.Error(err)
		}
	} else {
		c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		c.ServeJSON()
		return
		// user.Id = 9
	}

	locid := c.Ctx.Input.Param(":id")
	//id转成uint
	locidint, err := strconv.Atoi(locid)
	if err != nil {
		logs.Error(err)
	}
	locationid := uint(locidint)

	title := c.GetString("title")
	label := c.GetString("label")
	location := c.GetString("location")
	address := c.GetString("address")
	sort := c.GetString("sort")
	var sortint int
	if sort != "" {
		sortint, err = strconv.Atoi(sort)
		if err != nil {
			logs.Error(err)
		}
	}
	lat := c.GetString("lat")
	latfloat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		logs.Error(err)
	}
	lng := c.GetString("lng")
	lngfloat, err := strconv.ParseFloat(lng, 64)
	if err != nil {
		logs.Error(err)
	}

	// 进行敏感字符验证
	content := address
	// errcode, errmsg, err := utils.MsgSecCheck(accessToken, content)
	suggest, outlabel, err := utils.MsgSecCheck(2, 2, accessToken, openID.(string), content)
	// errcode, errmsg, err := utils.MsgSecCheck(2, 2, accessToken, openID.(string), content)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": err}
		c.ServeJSON()
	} else if suggest == "pass" {
		//根据项目id添加出差
		// var locationnavigate models.LocationNavigate
		// 添加location
		locationnavigate := models.LocationNavigate{
			LocationID: locationid,
			Title:      title,
			Label:      label,
			Location:   location,
			Address:    address,
			Lat:        latfloat,
			Lng:        lngfloat,
			Sort:       sortint,
		}
		Id, err := models.CreateLocationNavigate(locationnavigate)
		// beego.Info(Id)
		// beego.Info(err)
		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"data": "WRONG", "info": "添加出差数据错误"}
			c.ServeJSON()
		} else if Id != 0 {
			c.Data["json"] = map[string]interface{}{"data": "OK", "info": "SUCCESS", "id": Id}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"data": "data already exist", "info": "数据已存在"}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"info": suggest, "data": outlabel}
		c.ServeJSON()
	}
}

// @Title get location by userid
// @Description get location by userid
// @Param id path string true "The projectid of location"
// @Success 200 {object} models.GetLocationPage
// @Failure 400 Invalid page supplied
// @Failure 404 location not found
// @router /getlocation/:id [get]
// 根据项目id列出有关的location
func (c *LocationController) GetLocation() {
	pid := c.Ctx.Input.Param(":id")
	//id转成64为
	projectid, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(projectid)
	// projectid = 25001
	//查出未过期的与本人有关的出差活动
	location, err := models.GetAllLocation(projectid)
	if err != nil {
		logs.Error(err)
	}
	c.Data["json"] = location //map[string]interface{}{"userId": 1, "avatorUrl": "Filename"}
	c.ServeJSON()
}

// @Title get location by locationid
// @Description get location by locationid
// @Param id path string true "The id of location"
// @Success 200 {object} models.GetLocationPage
// @Failure 400 Invalid page supplied
// @Failure 404 location not found
// @router /getlocationbyid/:id [get]
// 根据locationid查出location
func (c *LocationController) GetLocationById() {
	var err error
	openID := c.GetSession("openID")
	if openID == nil {
		// c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		// c.ServeJSON()
		// return
		// user.Id = 9
	}

	bid := c.Ctx.Input.Param(":id")
	//id转成64为
	locationid, err := strconv.ParseInt(bid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	//查出出差活动
	location, err := models.GetLocationById(locationid)
	if err != nil {
		logs.Error(err)
	}

	c.Data["json"] = location //map[string]interface{}{"userId": 1, "avatorUrl": "Filename"}
	c.ServeJSON()
}
