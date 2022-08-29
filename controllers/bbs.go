//小程序日历公告
package controllers

import (
	"github.com/3xxx/engineercms/controllers/utils"
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"strconv"
	"strings"
	"time"
)

// CMSBBS API
type BbsController struct {
	web.Controller
}

// @Title post bbs
// @Description post bbs
// @Param userId query string true "The userId for bbs"
// @Param desc query string true "The description for bbs"
// @Param year query string true "The year for bbs"
// @Param month query string true "The month for bbs"
// @Param day query string true "The day for bbs"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 bbs not found
// @router /bbs [post]
// 新增公告写入数据库
func (c *BbsController) Bbs() {
	// var user models.User
	// var err error
	openID := c.GetSession("openID")
	if openID == nil {
		c.Data["json"] = map[string]interface{}{"code": 6, "info": "用户未登录"}
		c.ServeJSON()
		return
	}

	userid := c.GetString("userId")
	desc1 := c.GetString("desc1")
	desc2 := c.GetString("desc2")
	desc3 := c.GetString("desc3")
	desc4 := c.GetString("desc4")
	//pid转成64为
	UserId, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	year := c.GetString("year")
	month := c.GetString("month")
	if len(month) == 1 {
		month = "0" + month
	}
	day := c.GetString("day")
	if len(day) == 1 {
		day = "0" + day
	}
	const base_format = "2006-01-02"
	SelectDate, err := time.Parse(base_format, year+"-"+month+"-"+day)
	if err != nil {
		logs.Error(err)
	}
	// var array []string
	// data, err := models.GetBbs(SelectDate)
	// if err != nil {
	// 	logs.Error(err)
	// 	// arrayold := [4]string{"", "", "", ""}
	// } else {
	// 	array = strings.Split(data.Desc, "&#")
	// 	// for _, v := range array {
	// 	// }
	// 	if desc1 == "" && array[0] != "" {
	// 		desc1 = array[0]
	// 	}
	// 	if desc2 == "" && array[1] != "" {
	// 		desc2 = array[1]
	// 	}
	// 	if desc3 == "" && array[2] != "" {
	// 		desc3 = array[2]
	// 	}
	// 	if desc4 == "" && array[3] != "" {
	// 		desc4 = array[3]
	// 	}
	// }

	desc := desc1 + "&#" + desc2 + "&#" + desc3 + "&#" + desc4
	// 进行敏感字符验证
	app_version := c.GetString("app_version")
	accessToken, _, _, err := utils.GetAccessToken(app_version)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"code": 5, "info": "ERROR", "data": err}
		c.ServeJSON()
		return
	}
	suggest, label, err := utils.MsgSecCheck(2, 2, accessToken, openID.(string), desc)
	logs.Info(suggest)
	logs.Info(label)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"code": 4, "info": suggest, "data": err}
		c.ServeJSON()
	} else if suggest == "pass" {
		_, err = models.BbsBbs(UserId, desc, SelectDate)
		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"code": 2, "message": err}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"code": 1, "message": desc}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"code": 3, "info": suggest, "data": label}
		c.ServeJSON()
	}
}

type BbsDate struct {
	Year  string `json:"year"`
	Month string `json:"month"`
	Day   string `json:"day"`
}

// @Title get checkin bbs
// @Description get bbs
// @Param year query string true "The year for bbs"
// @Param month query string true "The month for bbs"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /bbsgetbbs [get]
// 取得当月的打卡记录
func (c *BbsController) BbsGetBbs() {
	const base_format = "2006-01-02"
	year := c.GetString("year")
	month := c.GetString("month")
	if len(month) == 1 {
		month = "0" + month
	}
	SelectMonth1, err := time.Parse(base_format, year+"-"+month+"-01")
	if err != nil {
		logs.Error(err)
	}
	SelectMonth2 := SelectMonth1.AddDate(0, 1, -1)

	data, err := models.BbsGetBbs(SelectMonth1, SelectMonth2)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"code": 2, "data": nil}
		c.ServeJSON()
	} else {
		// beego.Info(data)
		bbsdataslice := make([]BbsDate, 0)
		for _, v := range data {
			bbsdata := make([]BbsDate, 1)
			bbsdata[0].Year = v.SelectDate.Format("2006")
			bbsdata[0].Month = v.SelectDate.Format("01")
			bbsdata[0].Day = v.SelectDate.Format("02")
			bbsdataslice = append(bbsdataslice, bbsdata...)
		}
		c.Data["json"] = map[string]interface{}{"code": 1, "data": bbsdataslice}
		c.ServeJSON()
	}
}

// @Title get bbs
// @Description get bbs
// @Param year query string true "The year for bbs"
// @Param month query string true "The month for bbs"
// @Param day query string true "The day for bbs"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 bbs not found
// @router /getbbs [get]
// 取得选定日期的记录
func (c *BbsController) GetBbs() {
	const base_format = "2006-01-02"
	year := c.GetString("year")
	month := c.GetString("month")
	if len(month) == 1 {
		month = "0" + month
	}
	day := c.GetString("day")
	if len(day) == 1 {
		day = "0" + day
	}
	SelectDate, err := time.Parse(base_format, year+"-"+month+"-"+day)
	if err != nil {
		logs.Error(err)
	}

	data, err := models.GetBbs(SelectDate)
	if err != nil {
		logs.Error(err)
		array := [4]string{"", "", "", ""}
		c.Data["json"] = map[string]interface{}{"code": 2, "data": array}
		c.ServeJSON()
	} else {
		array := strings.Split(data.Desc, "&#")
		// for _, v := range array {

		// }
		c.Data["json"] = map[string]interface{}{"code": 1, "data": array}
		c.ServeJSON()
	}
}
