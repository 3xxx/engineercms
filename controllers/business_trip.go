// Business Trip
package controllers

import (
	"encoding/json"
	"github.com/3xxx/engineercms/controllers/utils"
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"sort"
	"strconv"
	"strings"
	"time"
)

type BusinessController struct {
	web.Controller
}

type wxuser struct {
	Name string `json:"name"`
}

// @Title post business by projectid
// @Description post business by projectid
// @Param content query string  true "The content of release"
// @Param location query string true "The location of business"
// @Param lat query string false "The lat of location"
// @Param lng query string false "The lng of location"
// @Param startDate query string false "The startDate of business"
// @Param endDate query string false "The endDate of business"
// @Param projecttitle query string false "The projecttitle of business"
// @Param drivername   query string false "The drivername   of business"
// @Param subsidy      query string false "The subsidy      of business"
// @Param carfare      query string false "The carfare      of business"
// @Param hotelfee     query string false "The hotelfee     of business"
// @Param users        query string false "The users        of business"
// @Param articleshow query string false "The larticleshow of business"
// @Param id path string true "The projectid of project"
// @Param title query string false "The title of article"
// @Param articlecontent query string false "The content of article"
// @Success 200 {object} models.CreateBusiness
// @Failure 400 Invalid page supplied
// @Failure 404 pas not found
// @router /addbusiness/:id [post]
// 添加一个活动
func (c *BusinessController) AddBusiness() {
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
	// _, _, uid, _, isLogin := checkprodRole(c.Ctx)
	// if !isLogin {
	// route := c.Ctx.Request.URL.String()
	// c.Data["Url"] = route
	// c.Redirect("/roleerr?url="+route, 302)
	// c.Data["json"] = "未登陆"
	// c.ServeJSON()
	// return
	// }
	// user, err := models.GetUserByUsername(uname)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// beego.Info("ok")

	// location: 东璟花园
	// lat: 23.122552
	// lng: 113.361518
	// startDate: 2021/12/15
	// endDate  : 2021/12/18
	// projecttitle: 1
	// drivername  : 2
	// subsidy     : 3
	// carfare     : 4
	// hotelfee    : 5
	// users       : [{"index":1,"name":"6","showtag":true},{"index":2,"name":"7","showtag":true}]
	// articleshow
	pid := c.Ctx.Input.Param(":id")
	//id转成64为
	projectid, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		logs.Error(err)
	}

	location := c.GetString("location")
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
	startdate := c.GetString("startDate")
	datearray := strings.Split(startdate, "/")
	year := datearray[0]
	month := datearray[1]
	if len(month) == 1 {
		month = "0" + month
	}
	day := datearray[2]
	if len(day) == 1 {
		day = "0" + day
	}
	startdate = year + "-" + month + "-" + day
	const base_format = "2006-01-02"
	StartDate, err := time.Parse(base_format, startdate)
	if err != nil {
		logs.Error(err)
	}

	enddate := c.GetString("endDate")
	datearray = strings.Split(enddate, "/")
	year = datearray[0]
	month = datearray[1]
	if len(month) == 1 {
		month = "0" + month
	}
	day = datearray[2]
	if len(day) == 1 {
		day = "0" + day
	}
	enddate = year + "-" + month + "-" + day
	EndDate, err := time.Parse(base_format, enddate)
	if err != nil {
		logs.Error(err)
	}
	EndDate = EndDate.AddDate(0, 0, 1)

	projecttitle := c.GetString("projecttitle")
	drivername := c.GetString("drivername")
	subsidy := c.GetString("subsidy")
	subsidyint, err := strconv.Atoi(subsidy)
	if err != nil {
		logs.Error(err)
	}
	carfare := c.GetString("carfare")
	carfareint, err := strconv.Atoi(carfare)
	if err != nil {
		logs.Error(err)
	}
	hotelfee := c.GetString("hotelfee")
	hotelfeeint, err := strconv.Atoi(hotelfee)
	if err != nil {
		logs.Error(err)
	}
	users := c.GetString("users")
	m := []wxuser{}
	err = json.Unmarshal([]byte(users), &m)
	if err != nil {
		logs.Error(err)
	}

	articleshow := c.GetString("articleshow")
	// 2020/11/15 17:08:02.488 [I] [business_trip.go:60] [{"index":1,"name":"6","showta
	// g":true},{"index":2,"name":"61","showtag":true}]
	// 进行敏感字符验证
	content := projecttitle + drivername + users

	suggest, label, err := utils.MsgSecCheck(2, 2, accessToken, openID.(string), content)
	// logs.Info(suggest)
	// logs.Info(label)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"code": 4, "info": suggest, "data": err}
		c.ServeJSON()
	} else if suggest == "pass" {
		//根据项目id添加出差
		var business models.Business
		// 添加business
		business = models.Business{
			ProjectID: projectid,
			UserID:    user.Id, //uid,
			// ArticleID:    aid,
			StartDate:    StartDate,
			EndDate:      EndDate,
			Location:     location,
			Lat:          latfloat,
			Lng:          lngfloat,
			Projecttitle: projecttitle,
			Drivername:   drivername,
			Subsidy:      subsidyint,
			Carfare:      carfareint,
			Hotelfee:     hotelfeeint,
			// Worktime:  worktime.Hours(),
			// Overtime:  0,
		}
		Id, err := models.CreateBusiness(business)
		// beego.Info(Id)
		// beego.Info(err)
		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"data": "WRONG", "info": "添加出差数据错误"}
			c.ServeJSON()
		} else if Id != 0 {
			// 添加自己
			var businessuser models.BusinessUser
			businessuser.UserID = user.Id //uid
			businessuser.BusinessID = Id
			_, err = models.CreateUserBusiness(businessuser)
			if err != nil {
				logs.Error(err)
				c.Data["json"] = map[string]interface{}{"data": "WRONG", "info": "添加同行人错误"}
				c.ServeJSON()
			}
			// 循环添加出差同行人员与出差的关联表
			if len(m) > 0 {
				for _, v := range m {
					// beego.Info(v.Name)
					tripuser := models.GetUserByNickname(v.Name)
					// beego.Info(tripuser.Id)
					// var businessuser models.BusinessUser
					businessuser.UserID = tripuser.Id
					businessuser.BusinessID = Id
					_, err = models.CreateUserBusiness(businessuser)
					if err != nil {
						logs.Error(err)
						c.Data["json"] = map[string]interface{}{"data": "WRONG", "info": "添加同行人错误"}
						c.ServeJSON()
					}
				}
			}
			// 如果文章时打开的，则添加文章并关联到出差
			if articleshow == "true" {
				_, err := models.AddArticle("title", "content", int64(Id))
				if err != nil {
					logs.Error(err)
					c.Data["json"] = map[string]interface{}{"data": "WRONG", "info": "添加文章错误"}
					c.ServeJSON()
				}
			}
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

// @Title post update business by businessid
// @Description post update business by businessid
// @Param content query string  true "The content of release"
// @Param location query string true "The location of business"
// @Param lat query string false "The lat of location"
// @Param lng query string false "The lng of location"
// @Param startDate query string false "The startDate of business"
// @Param endDate query string false "The endDate of business"
// @Param projecttitle query string false "The projecttitle of business"
// @Param drivername   query string false "The drivername   of business"
// @Param subsidy      query string false "The subsidy      of business"
// @Param carfare      query string false "The carfare      of business"
// @Param hotelfee     query string false "The hotelfee     of business"
// @Param users        query string false "The users        of business"
// @Param articleshow query string false "The larticleshow of business"
// @Param id path string true "The id of business"
// @Param title query string false "The title of article"
// @Param articlecontent query string false "The content of article"
// @Success 200 {object} models.UpdateBusiness
// @Failure 400 Invalid page supplied
// @Failure 404 pas not found
// @router /updatebusiness/:id [post]
// 用户修改business
func (c *BusinessController) UpdateBusiness() {
	//content去验证
	app_version := c.GetString("app_version")
	accessToken, _, _, err := utils.GetAccessToken(app_version)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": err}
		c.ServeJSON()
	}

	openID := c.GetSession("openID")
	if openID != nil {

	} else {
		c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		c.ServeJSON()
		return
		// user.Id = 9
	}

	bid := c.Ctx.Input.Param(":id")
	//id转成uint为
	businessidint, err := strconv.Atoi(bid)
	if err != nil {
		logs.Error(err)
	}
	businessid := uint(businessidint)

	location := c.GetString("location")
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
	startdate := c.GetString("startDate")
	datearray := strings.Split(startdate, "/")
	year := datearray[0]
	month := datearray[1]
	if len(month) == 1 {
		month = "0" + month
	}
	day := datearray[2]
	if len(day) == 1 {
		day = "0" + day
	}
	startdate = year + "-" + month + "-" + day
	const base_format = "2006-01-02"
	StartDate, err := time.Parse(base_format, startdate)
	if err != nil {
		logs.Error(err)
	}

	enddate := c.GetString("endDate")
	datearray = strings.Split(enddate, "/")
	year = datearray[0]
	month = datearray[1]
	if len(month) == 1 {
		month = "0" + month
	}
	day = datearray[2]
	if len(day) == 1 {
		day = "0" + day
	}
	enddate = year + "-" + month + "-" + day
	EndDate, err := time.Parse(base_format, enddate)
	if err != nil {
		logs.Error(err)
	}
	EndDate = EndDate.AddDate(0, 0, 1) //在编辑活动的显示那里，将enddate减少一天。

	projecttitle := c.GetString("projecttitle")
	drivername := c.GetString("drivername")
	subsidy := c.GetString("subsidy")
	subsidyint, err := strconv.Atoi(subsidy)
	if err != nil {
		logs.Error(err)
	}
	carfare := c.GetString("carfare")
	carfareint, err := strconv.Atoi(carfare)
	if err != nil {
		logs.Error(err)
	}
	hotelfee := c.GetString("hotelfee")
	hotelfeeint, err := strconv.Atoi(hotelfee)
	if err != nil {
		logs.Error(err)
	}
	users := c.GetString("users")
	m := []wxuser{}
	err = json.Unmarshal([]byte(users), &m)
	if err != nil {
		logs.Error(err)
	}

	// articleshow := c.GetString("articleshow")
	// 2020/11/15 17:08:02.488 [I] [business_trip.go:60] [{"index":1,"name":"6","showta
	// g":true},{"index":2,"name":"61","showtag":true}]
	// 进行敏感字符验证
	content := projecttitle + drivername + users
	suggest, label, err := utils.MsgSecCheck(2, 2, accessToken, openID.(string), content)
	// errcode, errmsg, err := utils.MsgSecCheck(2, 2, accessToken, openID.(string), content)
	// errcode, errmsg, err := utils.MsgSecCheck(accessToken, content)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": err}
		c.ServeJSON()
	} else if suggest == "pass" {
		//根据项目id添加出差
		var business models.Business
		// 添加business
		business = models.Business{
			// ArticleID:    aid,
			ID:           businessid,
			StartDate:    StartDate,
			EndDate:      EndDate,
			Location:     location,
			Lat:          latfloat,
			Lng:          lngfloat,
			Projecttitle: projecttitle,
			Drivername:   drivername,
			Subsidy:      subsidyint,
			Carfare:      carfareint,
			Hotelfee:     hotelfeeint,
			// Worktime:  worktime.Hours(),
			// Overtime:  0,
		}
		err := models.UpdateBusiness(business)
		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"data": "WRONG", "info": "更新出差数据错误"}
			c.ServeJSON()
		} else {
			// 循环添加出差同行人员与出差的关联表
			if len(m) > 0 {
				// 查出数据库中关联的users
				businessusers, err := models.GetBusinessUsers(businessid)
				if err != nil {
					logs.Error(err)
				}

				for _, v := range m {
					var nameisold bool
					for _, w := range businessusers {
						if v.Name == w.NickNames.Nickname {
							nameisold = true
							break
						}
					}
					// beego.Info(v.Name)
					if nameisold == false {
						tripuser := models.GetUserByNickname(v.Name)
						// beego.Info(tripuser.Id)
						var businessuser models.BusinessUser
						businessuser.UserID = tripuser.Id
						businessuser.BusinessID = businessid
						_, err = models.CreateUserBusiness(businessuser)
						if err != nil {
							logs.Error(err)
							c.Data["json"] = map[string]interface{}{"data": "WRONG", "info": "添加同行人错误"}
							c.ServeJSON()
						}
					}
				}

				for _, w := range businessusers {
					var databasehasname bool
					for _, v := range m {
						if v.Name == w.NickNames.Nickname {
							databasehasname = true
							break
						}
					}
					// beego.Info(v.Name)
					if databasehasname == false {
						err = models.DeleteUserBusiness(w.NickNames.Id, businessid)
						if err != nil {
							logs.Error(err)
							c.Data["json"] = map[string]interface{}{"data": "WRONG", "info": "删除同行人错误"}
							c.ServeJSON()
						}
					}
				}
			}
			// 如果文章时打开的，则添加文章并关联到出差
			c.Data["json"] = map[string]interface{}{"data": "OK", "info": "SUCCESS"}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"info": suggest, "data": label}
		c.ServeJSON()
	}
}

// @Title get business by userid
// @Description get business by userid
// @Param id path string true "The projectid of business"
// @Success 200 {object} models.GetBusinessPage
// @Failure 400 Invalid page supplied
// @Failure 404 business not found
// @router /getbusiness/:id [get]
// 根据项目id和用户id列出有关的business
// 不应该过滤项目id！！
func (c *BusinessController) GetBysiness() {
	var user models.User
	var err error
	_, _, uid, isadmin, islogin := checkprodRole(c.Ctx)
	openID := c.GetSession("openID")
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			logs.Error(err)
		}
	} else if isadmin || islogin {
		user.Id = uid
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
	// beego.Info(projectid)
	// projectid = 25001
	//查出未过期的与本人有关的出差活动
	business, err := models.GetAllBusiness(projectid, user.Id)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(business)
	// local, _ := time.LoadLocation("Local")
	// starttime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-10-23 15:01:01", local)
	// endtime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-10-24 15:05:01", local)
	// now := time.Now()
	// fmt.Println(now)
	// fmt.Println(endtime, endtime.After(now))
	// fmt.Println(starttime, starttime.Before(now))
	c.Data["json"] = business //map[string]interface{}{"userId": 1, "avatorUrl": "Filename"}
	c.ServeJSON()
}

// @Title get business by businessid
// @Description get business by businessid
// @Param id path string true "The id of business"
// @Success 200 {object} models.GetBusinessPage
// @Failure 400 Invalid page supplied
// @Failure 404 business not found
// @router /getbusinessbyid/:id [get]
// 根据businessid查出business
func (c *BusinessController) GetBysinessById() {
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
	businessid, err := strconv.ParseInt(bid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	//查出出差活动
	business, err := models.GetBusinessById(businessid)
	if err != nil {
		logs.Error(err)
	}
	// enddate要减一天
	business.EndDate = business.EndDate.AddDate(0, 0, -1)
	c.Data["json"] = business //map[string]interface{}{"userId": 1, "avatorUrl": "Filename"}
	c.ServeJSON()
}

// @Title post checkin person
// @Description post person
// @Param userid query string true "The userid for person"
// @Param businessid query string true "The businessid for Business"
// @Param lat query string true "The businessid of check"
// @Param lng query string true "The businessid of check"
// @Param photoUrl query string true "The photoUrl of check"
// @Param year query string true "The businessid of check"
// @Param month query string true "The businessid of check"
// @Param day query string true "The businessid of check"
// @Param location query string true "The location of check"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /businesscheck [post]
// 打卡记录写入数据库
func (c *BusinessController) BusinessCheck() {
	userid := c.GetString("userid")
	//pid转成64为
	UserId, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	businessid := c.GetString("businessid")
	businessidint, err := strconv.Atoi(businessid)
	if err != nil {
		logs.Error(err)
	}
	BusinessId := uint(businessidint)

	lat := c.GetString("lat")
	//string到float64
	Lat, err := strconv.ParseFloat(lat, 64)
	lng := c.GetString("lng")
	Lng, err := strconv.ParseFloat(lng, 64)

	// startdate := c.GetString("startDate")
	// beego.Info(startdate)
	const base_format = "2006-01-02"
	// StartDate, err := time.Parse(base_format, startdate)
	// if err != nil {
	// 	logs.Error(err)
	// }
	PhotoUrl := c.GetString("photoUrl")
	location := c.GetString("location")
	//根据userid取出user和avatorUrl
	useravatar, err := models.GetUserAvatorUrl(UserId)
	if err != nil {
		logs.Error(err)
	}
	var photo string
	if len(useravatar) != 0 {
		wxsite, err := web.AppConfig.String("wxreqeustsite")
		if err != nil {
			logs.Error(err)
		}
		photo = wxsite + useravatar[0].UserAvatar.AvatarUrl
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
	SelectDate, err := time.Parse(base_format, year+"-"+month+"-"+day)
	if err != nil {
		logs.Error(err)
	}
	_, err = models.BusinessCheck(BusinessId, UserId, Lat, Lng, PhotoUrl, location, SelectDate)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"code": 2, "message": PhotoUrl}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": 1, "avatorUrl": photo}
		c.ServeJSON()
	}
}

type BusinessCheckDate struct {
	Year  string `json:"year"`
	Month string `json:"month"`
	Day   string `json:"day"`
}

// @Title get checkin check
// @Description get check
// @Param userid query string true "The userId of check"
// @Param businessid query string true "The activityId of check"
// @Param year query string true "The year of check"
// @Param month query string true "The month of check"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getbusinesscheck [get]
// 取得当月的打卡记录
func (c *BusinessController) GetBusinessCheck() {
	userid := c.GetString("userid")
	//pid转成64为
	UserId, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	businessid := c.GetString("businessid")
	BusinessId, err := strconv.ParseInt(businessid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
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
	SelectMonth2 := SelectMonth1.AddDate(0, 1, 0)

	data, err := models.GetBusinessCheck(BusinessId, UserId, SelectMonth1, SelectMonth2)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"code": 2, "data": nil}
		c.ServeJSON()
	} else {
		// beego.Info(data)
		checkdataslice := make([]BusinessCheckDate, 0)
		for _, v := range data {
			checkdata := make([]BusinessCheckDate, 1)
			checkdata[0].Year = v.SelectDate.Format("2006")
			checkdata[0].Month = v.SelectDate.Format("01")
			checkdata[0].Day = v.SelectDate.Format("02")
			checkdataslice = append(checkdataslice, checkdata...)
		}
		c.Data["json"] = map[string]interface{}{"code": 1, "data": checkdataslice}
		c.ServeJSON()
	}
}

// ********************统计**************
// @Title get businessmothcheckin
// @Description get businessmonthcheck
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /businessmonthchecksum [get]
func (c *BusinessController) BusinessMonthCheckSum() {
	c.TplName = "businesscheck/check.tpl"
	c.Data["IsMonthCheck"] = true
}

// @Title get businessmothcheckin
// @Description get businessmonthcheck
// @Param id path string true "The projectid of business"
// @Param page query string false "The page of check"
// @Param limit query string false "The size of check"
// @Param year query string true "The year of check"
// @Param month query string true "The month of check"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /businessmonthcheck/:id [get]
// 月度考勤统计
func (c *BusinessController) BusinessMonthCheck() {
	var offset, limit1, page1 int
	var err error
	limit := c.GetString("limit")
	if limit == "" {
		limit1 = 500
	} else {
		limit1, err = strconv.Atoi(limit)
		if err != nil {
			logs.Error(err)
		}
	}
	page := c.GetString("page")
	if page == "" {
		page1 = 1
	} else {
		page1, err = strconv.Atoi(page)
		if err != nil {
			logs.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	//当月天数
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
	SelectMonth2 := SelectMonth1.AddDate(0, 1, 0)
	//建立一个动态月日数组
	days := SelectMonth2.Sub(SelectMonth1) / 24
	dayss := days.Hours()
	daysss := strconv.FormatFloat(dayss, 'f', -1, 64) //float64转string
	dayssss, err := strconv.Atoi(daysss)              //string转int

	business, err := models.GetBusinessCheckUser(SelectMonth1, SelectMonth2, limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	pid := c.Ctx.Input.Param(":id")
	projectid, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	s := []map[int]interface{}{}
	for _, w := range business {
		if len(w.BusinessCheckins) > 0 && w.ProjectID == projectid {
			var checkmap = make(map[int]interface{}, dayssss+1)
			var checkmap3 = make(map[int]interface{}, dayssss+1)
			for i := 0; i <= dayssss; i++ {
				if i == 0 {
					checkmap3[0] = "项目名称：" + w.Projecttitle
				} else if i == 1 {
					checkmap3[1] = "出差地点：" + w.Location
				} else if i == 2 {
					checkmap3[2] = "司机/车牌号：" + w.Drivername
				} else if i == 3 {
					checkmap3[3] = "补贴标准：" + strconv.Itoa(w.Subsidy)
				} else if i == 4 {
					checkmap3[4] = "交通费：" + strconv.Itoa(w.Carfare)
				} else if i == 5 {
					checkmap3[5] = "住宿费：" + strconv.Itoa(w.Hotelfee)
				} else {
					checkmap3[i] = ""
				}
			}
			// 在s数组前面插入
			s = append(s, checkmap3)

			// if len(w.BusinessCheckins) > 0 {
			// map去重
			strMap := make(map[string]string)
			for _, v := range w.BusinessCheckins {
				strMap[v.Users.Nickname] = v.Users.Nickname
			}
			//strMap为：{"slice"："slice","int"："int","string"："string","boolean"：boolean"}
			//如果想将map转换为slice，可利用数组的append函数
			// var secondStr []string
			// for _, value := range strMap {
			// 	secondStr = append(secondStr, value)
			// }
			for _, value := range strMap { //循环人名
				for i := 0; i <= dayssss; i++ {
					if i == 0 {
						checkmap[0] = value
					} else {
						checkmap[i] = ""
					}
				}

				for _, vv := range w.BusinessCheckins { //循环这个人的考勤
					if value == vv.Users.Nickname {
						day := vv.SelectDate.Format("02")
						dayint, err := strconv.Atoi(day)
						if err != nil {
							logs.Error(err)
						}
						checkmap[dayint] = "1" //vv.CheckTime.Add(8 * h) //vv.Location
					}
				}
				//查出一个用户这个月的打卡记录
				// data, err := models.CheckGetCheck(ActivityId, v.Checkin.UserId, SelectMonth1, SelectMonth2)
				// if err != nil {
				// 	logs.Error(err)
				// } else {
				// 	for _, w := range data {
				// 		day := w.SelectDate.Format("02")
				// 		dayint, err := strconv.Atoi(day)
				// 		if err != nil {
				// 			logs.Error(err)
				// 		}
				// 		checkmap[dayint] = "1"
				// 	}
				// }

				//排序map，按key
				var keys []int
				for k := range checkmap {
					keys = append(keys, k)
				}
				sort.Ints(keys)
				var checkmap2 = make(map[int]interface{}, dayssss+1)
				for _, k := range keys {
					checkmap2[k] = checkmap[k]
				}
				s = append(s, checkmap2)
			}
		}
	}
	b, err := json.Marshal(s)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"code": 2, "data": nil}
		c.ServeJSON()
	} else {
		c.Ctx.WriteString(string(b))
	}
}

// @Title get businessmothcheckin
// @Description get businessmonthcheck
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /businessmonthchecksum2 [get]
// 报税用
func (c *BusinessController) BusinessMonthCheckSum2() {
	c.TplName = "businesscheck/check2.tpl"
	c.Data["IsMonthCheck"] = true
}

// @Title get businessmothcheckin
// @Description get businessmonthcheck
// @Param id path string true "The projectid of business"
// @Param page query string false "The page of check"
// @Param limit query string false "The size of check"
// @Param year query string true "The year of check"
// @Param month query string true "The month of check"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /businessmonthcheck2/:id [get]
// 月度考勤统计-报税用
func (c *BusinessController) BusinessMonthCheck2() {
	// h, _ := time.ParseDuration("1h")
	var offset, limit1, page1 int
	var err error
	var city, province string
	limit := c.GetString("limit")
	if limit == "" {
		limit1 = 500
	} else {
		limit1, err = strconv.Atoi(limit)
		if err != nil {
			logs.Error(err)
		}
	}
	page := c.GetString("page")
	if page == "" {
		// limit1 = 10
		page1 = 1
	} else {
		page1, err = strconv.Atoi(page)
		if err != nil {
			logs.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	//当月天数
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
	SelectMonth2 := SelectMonth1.AddDate(0, 1, 0)
	//建立一个动态月日数组
	days := SelectMonth2.Sub(SelectMonth1) / 24
	dayss := days.Hours()
	daysss := strconv.FormatFloat(dayss, 'f', -1, 64) //float64转string
	dayssss, err := strconv.Atoi(daysss)              //string转int

	business, err := models.GetBusinessCheckUser(SelectMonth1, SelectMonth2, limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(SelectMonth1)
	// beego.Info(SelectMonth2)
	// beego.Info(business)
	// c.Data["json"] = business
	// c.ServeJSON()
	pid := c.Ctx.Input.Param(":id")
	//id转成64为
	projectid, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	s := []map[int]interface{}{}

	// 将姓名过滤
	strMap := make(map[string]string)
	for _, w := range business {
		if len(w.BusinessCheckins) > 0 && w.ProjectID == projectid {
			for _, v := range w.BusinessCheckins {
				strMap[v.Users.Nickname] = v.Users.Nickname
			}
		}
	}
	// beego.Info(strMap)
	//活动
	for _, value := range strMap { //循环人名
		var checkmap = make(map[int]interface{}, dayssss+1)
		for i := 0; i <= dayssss; i++ {
			if i == 0 {
				checkmap[0] = value
			} else {
				checkmap[i] = ""
			}
		}
		for _, w := range business {
			if len(w.BusinessCheckins) > 0 && w.ProjectID == projectid {
				// strMap := make(map[string]string)
				// for _, v := range w.BusinessCheckins {
				// 	strMap[v.Users.Nickname] = v.Users.Nickname
				// }
				for _, vv := range w.BusinessCheckins { //循环这个人的考勤
					if value == vv.Users.Nickname {
						day := vv.SelectDate.Format("02")
						dayint, err := strconv.Atoi(day)
						if err != nil {
							logs.Error(err)
						}
						// logs.Info(vv.Location)
						provinceindex := UnicodeIndex(vv.Location, "省") // 查找格这个字符的位置
						// logs.Info(provinceindex)
						if provinceindex == 0 {
							provinceindex = UnicodeIndex(vv.Location, "自治区")
							province = SubString(vv.Location, 0, provinceindex+1+2)
						} else {
							province = SubString(vv.Location, 0, provinceindex+1)
						}
						if province == "广东省" {
							province = ""
						}
						cityindex := UnicodeIndex(vv.Location, "市")
						// logs.Info(cityindex)
						if cityindex == 0 {
							cityindex = UnicodeIndex(vv.Location, "自治州")
							city = SubString(vv.Location, provinceindex+1, cityindex-provinceindex+2)
						} else {
							city = SubString(vv.Location, provinceindex+1, cityindex-provinceindex)
						}
						// logs.Info(city)
						if province == "" && city == "" {
							city = vv.Location
						} else {
							city = province + city
						}
						if w.Subsidy > 100 {
							checkmap[dayint] = strconv.Itoa(w.Subsidy-100) + "-" + city //vv.CheckTime.Add(8 * h) //vv.Location
						}
					}
				}
			}
		}
		//排序map，按key
		var keys []int
		for k := range checkmap {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		var checkmap2 = make(map[int]interface{}, dayssss+1)
		for _, k := range keys {
			checkmap2[k] = checkmap[k]
		}
		s = append(s, checkmap2)
	}
	b, err := json.Marshal(s)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"code": 2, "data": nil}
		c.ServeJSON()
	} else {
		c.Ctx.WriteString(string(b))
	}
}

// @Title get businessmothcheckin
// @Description get businessmonthcheck
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /businessmonthchecksum3 [get]
// 报销用
func (c *BusinessController) BusinessMonthCheckSum3() {
	c.TplName = "businesscheck/check3.tpl"
	c.Data["IsMonthCheck"] = true
}

// @Title get businessmothcheckin
// @Description get businessmonthcheck
// @Param id path string true "The projectid of business"
// @Param page query string false "The page of check"
// @Param limit query string false "The size of check"
// @Param year query string true "The year of check"
// @Param month query string true "The month of check"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /businessmonthcheck3/:id [get]
// 月度考勤统计-报销用
func (c *BusinessController) BusinessMonthCheck3() {
	// h, _ := time.ParseDuration("1h")
	var offset, limit1, page1 int
	var city, province string
	var err error
	limit := c.GetString("limit")
	if limit == "" {
		limit1 = 500
	} else {
		limit1, err = strconv.Atoi(limit)
		if err != nil {
			logs.Error(err)
		}
	}
	page := c.GetString("page")
	if page == "" {
		// limit1 = 10
		page1 = 1
	} else {
		page1, err = strconv.Atoi(page)
		if err != nil {
			logs.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	//当月天数
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
	SelectMonth2 := SelectMonth1.AddDate(0, 1, 0)
	//建立一个动态月日数组
	days := SelectMonth2.Sub(SelectMonth1) / 24
	dayss := days.Hours()
	daysss := strconv.FormatFloat(dayss, 'f', -1, 64) //float64转string
	dayssss, err := strconv.Atoi(daysss)              //string转int

	business, err := models.GetBusinessCheckUser(SelectMonth1, SelectMonth2, limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(SelectMonth1)
	// beego.Info(SelectMonth2)
	// beego.Info(business)
	// c.Data["json"] = business
	// c.ServeJSON()
	pid := c.Ctx.Input.Param(":id")
	//id转成64为
	projectid, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	s := []map[int]interface{}{}

	// 将姓名过滤
	strMap := make(map[string]string)
	for _, w := range business {
		if len(w.BusinessCheckins) > 0 && w.ProjectID == projectid {
			for _, v := range w.BusinessCheckins {
				strMap[v.Users.Nickname] = v.Users.Nickname
			}
		}
	}

	for _, value := range strMap { //循环人名
		var checkmap = make(map[int]interface{}, dayssss+1)
		for i := 0; i <= dayssss; i++ {
			if i == 0 {
				checkmap[0] = value
			} else {
				checkmap[i] = ""
			}
		}
		for _, w := range business {
			if len(w.BusinessCheckins) > 0 && w.ProjectID == projectid {
				// strMap := make(map[string]string)
				// for _, v := range w.BusinessCheckins {
				// 	strMap[v.Users.Nickname] = v.Users.Nickname
				// }
				for _, vv := range w.BusinessCheckins { //循环这个人的考勤
					if value == vv.Users.Nickname {
						day := vv.SelectDate.Format("02")
						dayint, err := strconv.Atoi(day)
						if err != nil {
							logs.Error(err)
						}
						provinceindex := UnicodeIndex(vv.Location, "省") // 查找格这个字符的位置
						// logs.Info(provinceindex)
						if provinceindex == 0 {
							provinceindex = UnicodeIndex(vv.Location, "自治区")
							province = SubString(vv.Location, 0, provinceindex+1+2)
						} else {
							province = SubString(vv.Location, 0, provinceindex+1)
						}
						if province == "广东省" {
							province = ""
						}
						cityindex := UnicodeIndex(vv.Location, "市")
						// logs.Info(cityindex)
						if cityindex == 0 {
							cityindex = UnicodeIndex(vv.Location, "自治州")
							city = SubString(vv.Location, provinceindex+1, cityindex-provinceindex+2)
						} else {
							city = SubString(vv.Location, provinceindex+1, cityindex-provinceindex)
						}
						// beego.Info(city)
						if province == "" && city == "" {
							city = vv.Location
						} else {
							city = province + city
						}
						if w.Drivername == "" {
							w.Drivername = "司机"
						}
						checkmap[dayint] = w.Projecttitle + "-" + city + "-" + w.Drivername + "-" + strconv.Itoa(w.Carfare) + "-" + strconv.Itoa(w.Hotelfee) + "-" + strconv.Itoa(w.Subsidy) //vv.CheckTime.Add(8 * h) //vv.Location
					}
				}
			}
		}
		//排序map，按key
		var keys []int
		for k := range checkmap {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		var checkmap2 = make(map[int]interface{}, dayssss+1)
		for _, k := range keys {
			checkmap2[k] = checkmap[k]
		}
		s = append(s, checkmap2)
	}

	b, err := json.Marshal(s)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"code": 2, "data": nil}
		c.ServeJSON()
	} else {
		c.Ctx.WriteString(string(b))
	}
}

// @Title get businessmothcheckin
// @Description get businessmonthcheck
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /businessmonthchecksum4 [get]
func (c *BusinessController) BusinessMonthCheckSum4() {
	c.TplName = "businesscheck/check4.tpl"
	c.Data["IsMonthCheck"] = true
}

// @Title get businessmothcheckin
// @Description get businessmonthcheck
// @Param id path string true "The projectid of business"
// @Param page query string false "The page of check"
// @Param limit query string false "The size of check"
// @Param year query string true "The year of check"
// @Param month query string true "The month of check"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /businessmonthcheck4/:id [get]
// 月度考勤统计——项目+姓名放同一行
func (c *BusinessController) BusinessMonthCheck4() {
	var offset, limit1, page1 int
	var err error
	limit := c.GetString("limit")
	if limit == "" {
		limit1 = 500
	} else {
		limit1, err = strconv.Atoi(limit)
		if err != nil {
			logs.Error(err)
		}
	}
	page := c.GetString("page")
	if page == "" {
		page1 = 1
	} else {
		page1, err = strconv.Atoi(page)
		if err != nil {
			logs.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	//当月天数
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
	SelectMonth2 := SelectMonth1.AddDate(0, 1, 0)
	//建立一个动态月日数组
	// days := SelectMonth2.Sub(SelectMonth1) / 24
	// dayss := days.Hours()
	// daysss := strconv.FormatFloat(dayss, 'f', -1, 64) //float64转string
	// dayssss, err := strconv.Atoi(daysss)              //string转int

	business, err := models.GetBusinessCheckUser(SelectMonth1, SelectMonth2, limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	pid := c.Ctx.Input.Param(":id")
	projectid, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	s := []map[int]interface{}{}
	for _, w := range business {
		if len(w.BusinessCheckins) > 0 && w.ProjectID == projectid {
			var checkmap = make(map[int]interface{}, 32)
			// var checkmap3 = make(map[int]interface{}, 32) //dayssss+2
			// for i := 0; i <= 32; i++ {
			// 	if i == 0 {
			// 		checkmap3[0] = "项目名称：" + w.Projecttitle
			// 	} else if i == 1 {
			// 		checkmap3[1] = "出差地点：" + w.Location
			// 	} else if i == 2 {
			// 		checkmap3[2] = "司机/车牌号：" + w.Drivername
			// 	} else if i == 3 {
			// 		checkmap3[3] = "补贴标准：" + strconv.Itoa(w.Subsidy)
			// 	} else if i == 4 {
			// 		checkmap3[4] = "交通费：" + strconv.Itoa(w.Carfare)
			// 	} else if i == 5 {
			// 		checkmap3[5] = "住宿费：" + strconv.Itoa(w.Hotelfee)
			// 	} else {
			// 		checkmap3[i] = ""
			// 	}
			// }
			// 在s数组前面插入
			// s = append(s, checkmap3)

			// if len(w.BusinessCheckins) > 0 {
			strMap := make(map[string]string)
			for _, v := range w.BusinessCheckins {
				strMap[v.Users.Nickname] = v.Users.Nickname
			}
			//strMap为：{"slice"："slice","int"："int","string"："string","boolean"：boolean"}
			//如果想将map转换为slice，可利用数组的append函数
			// var secondStr []string
			// for _, value := range strMap {
			// 	secondStr = append(secondStr, value)
			// }
			for _, value := range strMap { //循环人名
				for i := 0; i <= 32; i++ {
					if i == 0 {
						checkmap[0] = value
					} else if i == 32 {
						checkmap[i] = "项目名称：" + w.Projecttitle
					} else {
						checkmap[i] = ""
					}
				}

				for _, vv := range w.BusinessCheckins { //循环这个人的考勤
					if value == vv.Users.Nickname {
						day := vv.SelectDate.Format("02")
						dayint, err := strconv.Atoi(day)
						if err != nil {
							logs.Error(err)
						}
						checkmap[dayint] = "1" //vv.CheckTime.Add(8 * h) //vv.Location
					}
				}
				//查出一个用户这个月的打卡记录
				// data, err := models.CheckGetCheck(ActivityId, v.Checkin.UserId, SelectMonth1, SelectMonth2)
				// if err != nil {
				// 	logs.Error(err)
				// } else {
				// 	for _, w := range data {
				// 		day := w.SelectDate.Format("02")
				// 		dayint, err := strconv.Atoi(day)
				// 		if err != nil {
				// 			logs.Error(err)
				// 		}
				// 		checkmap[dayint] = "1"
				// 	}
				// }

				//排序map，按key
				var keys []int
				for k := range checkmap {
					keys = append(keys, k)
				}
				sort.Ints(keys)
				var checkmap2 = make(map[int]interface{}, 32)
				for _, k := range keys {
					checkmap2[k] = checkmap[k]
				}
				s = append(s, checkmap2)
			}
		}
	}
	b, err := json.Marshal(s)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"code": 2, "data": nil}
		c.ServeJSON()
	} else {
		c.Ctx.WriteString(string(b))
	}
}

// @Title get businessmothcheckin
// @Description get businessmonthcheck
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /businessmonthchecksum5 [get]
// 报销用
func (c *BusinessController) BusinessMonthCheckSum5() {
	c.TplName = "businesscheck/check5.tpl"
	c.Data["IsMonthCheck"] = true
}

// @Title get businessmothcheckin
// @Description get businessmonthcheck
// @Param id path string true "The projectid of business"
// @Param page query string false "The page of check"
// @Param limit query string false "The size of check"
// @Param year query string true "The year of check"
// @Param month query string true "The month of check"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /businessmonthcheck5/:id [get]
// 月度考勤统计-报销用
func (c *BusinessController) BusinessMonthCheck5() {
	// h, _ := time.ParseDuration("1h")
	var offset, limit1, page1 int
	var err error
	limit := c.GetString("limit")
	if limit == "" {
		limit1 = 500
	} else {
		limit1, err = strconv.Atoi(limit)
		if err != nil {
			logs.Error(err)
		}
	}
	page := c.GetString("page")
	if page == "" {
		// limit1 = 10
		page1 = 1
	} else {
		page1, err = strconv.Atoi(page)
		if err != nil {
			logs.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	//当月天数
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
	SelectMonth2 := SelectMonth1.AddDate(0, 1, 0)
	//建立一个动态月日数组
	days := SelectMonth2.Sub(SelectMonth1) / 24
	dayss := days.Hours()
	daysss := strconv.FormatFloat(dayss, 'f', -1, 64) //float64转string
	dayssss, err := strconv.Atoi(daysss)              //string转int

	business, err := models.GetBusinessCheckUser(SelectMonth1, SelectMonth2, limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(SelectMonth1)
	// beego.Info(SelectMonth2)
	// beego.Info(business)
	// c.Data["json"] = business
	// c.ServeJSON()
	pid := c.Ctx.Input.Param(":id")
	//id转成64为
	projectid, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	s := []map[int]interface{}{}

	// 将姓名过滤
	strMap := make(map[string]string)
	for _, w := range business {
		if len(w.BusinessCheckins) > 0 && w.ProjectID == projectid {
			for _, v := range w.BusinessCheckins {
				strMap[v.Users.Nickname] = v.Users.Nickname
			}
		}
	}

	for _, value := range strMap { //循环人名
		var checkmap = make(map[int]interface{}, dayssss+1)
		for i := 0; i <= dayssss; i++ {
			if i == 0 {
				checkmap[0] = value
			} else {
				checkmap[i] = ""
			}
		}
		for _, w := range business {
			if len(w.BusinessCheckins) > 0 && w.ProjectID == projectid {
				// strMap := make(map[string]string)
				// for _, v := range w.BusinessCheckins {
				// 	strMap[v.Users.Nickname] = v.Users.Nickname
				// }
				for _, vv := range w.BusinessCheckins { //循环这个人的考勤
					if value == vv.Users.Nickname {
						day := vv.SelectDate.Format("02")
						dayint, err := strconv.Atoi(day)
						if err != nil {
							logs.Error(err)
						}
						checkmap[dayint] = "1"
					}
				}
			}
		}
		//排序map，按key
		var keys []int
		for k := range checkmap {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		var checkmap2 = make(map[int]interface{}, dayssss+1)
		for _, k := range keys {
			checkmap2[k] = checkmap[k]
		}
		s = append(s, checkmap2)
	}

	b, err := json.Marshal(s)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"code": 2, "data": nil}
		c.ServeJSON()
	} else {
		c.Ctx.WriteString(string(b))
	}
}
