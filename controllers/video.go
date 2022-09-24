package controllers

import (
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	// beego "github.com/beego/beego/v2/adapter"
	// "io/ioutil"
	// "net"
	// "net/http"
	// "net/url"
	// "os"
	// "path"
	"strconv"
	// "strings"
	// "time"
	"regexp"
)

// CMSVideo API
type VideoController struct {
	web.Controller
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
// 获得视频，如果是admin角色，则查询全部
func (c *VideoController) GetUserVideo() {
	//解析表单
	pid := c.Ctx.Input.Param(":id")
	// pid := web.AppConfig.String("wxcatalogid") //"26159" //c.GetString("pid")
	//pid转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		logs.Error(err)
	}

	// v := c.GetSession("uname")
	// var user models.User
	// var userid, roleid string
	// var err error
	// if v != nil { //如果登录了
	// 	user, err = models.GetUserByUsername(v.(string))
	// 	if err != nil {
	// 		logs.Error(err)
	// 	}
	//查询admin角色的id
	//重新获取roleid
	// role, err := models.GetRoleByRolename("admin")
	// if err != nil {
	// 	logs.Error(err)
	// }
	// userid = strconv.FormatInt(user.Id, 10)
	// roleid = strconv.FormatInt(role.Id, 10)

	// } else {
	// 	c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
	// 	c.ServeJSON()
	// 	return
	// }

	searchText := c.GetString("searchText")

	limit := c.GetString("limit")
	if limit == "" {
		limit = "15"
	}
	limit1, err := strconv.Atoi(limit)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("pageNo")
	page1, err := strconv.Atoi(page)
	if err != nil {
		logs.Error(err)
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
		logs.Error(err)
	}

	videos, err := models.GetUserVideo(proj.Id, limit1, offset, searchText)
	if err != nil {
		logs.Error(err)
	}
	count, err := models.GetUserVideoCount(proj.Id, searchText)
	if err != nil {
		logs.Error(err)
	}
	// c.Data["json"] = carts
	c.Data["json"] = map[string]interface{}{"page": page1, "total": count, "rows": videos}

	// c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "mycarts": carts}
	c.ServeJSON()
}

type MonthVideoList struct {
	YearMonth string //time.Time
	// videolists []videolist
	VideoLists []models.Video
}

// @Title get video
// @Description get video
// @Param page query string true "The page for photos list"
// @Param limit query string true "The limit of page for photos list"
// @Success 200 {object} models.GetAttachbyId
// @Failure 400 Invalid page supplied
// @Failure 404 pdf not found
// @router /video [get]
func (c *VideoController) Video() {
	var yearmonth, videoyearmonth string
	// limit := "5"
	limit := c.GetString("limit")
	if limit == "" {
		limit = "10"
	}
	limit1, err := strconv.Atoi(limit)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("page")
	if page == "" {
		page = "0"
	}
	page1, err := strconv.Atoi(page)
	if err != nil {
		logs.Error(err)
	}
	var offset int
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	videolist, err := models.GetVideoData(limit1, offset)
	if err != nil {
		logs.Error(err)
	}

	// logs.Info(videolist)
	videolist2 := make([]models.Video, 0)
	monthvideolist := make([]MonthVideoList, 0)
	const layoutyearmonth = "2006-01"
	for j := 0; j < len(videolist); j++ {
		videoyearmonth = videolist[j].Created.Format(layoutyearmonth)
		if yearmonth == "" || yearmonth == videoyearmonth {
			yearmonth = videoyearmonth
			photoarr := make([]models.Video, 1)
			photoarr[0] = videolist[j]
			// photoarr[0].Url = videolist[j].Url
			videolist2 = append(videolist2, photoarr...)
		} else {
			yearmonth = videoyearmonth
			monthphotoarr := make([]MonthVideoList, 1)
			monthphotoarr[0].YearMonth = videolist[j-1].Created.Format(layoutyearmonth)
			monthphotoarr[0].VideoLists = videolist2
			monthvideolist = append(monthvideolist, monthphotoarr...)

			videolist2 = make([]models.Video, 0)
			photoarr := make([]models.Video, 1)
			photoarr[0] = videolist[j]
			// photoarr[0].Url = videolist[j].Url
			videolist2 = append(videolist2, photoarr...)
		}
	}
	// 最后一部分
	monthphotoarr := make([]MonthVideoList, 1)
	monthphotoarr[0].YearMonth = yearmonth
	monthphotoarr[0].VideoLists = videolist2
	monthvideolist = append(monthvideolist, monthphotoarr...)
	// logs.Info(monthvideolist)
	// c.Data["videolist"] = photodata
	c.Data["MonthVideoList"] = monthvideolist
	u := c.Ctx.Input.UserAgent()
	// re := regexp.MustCompile("Trident")
	// loc := re.FindStringIndex(u)
	// loc[0] > 1
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		// logs.Info("移动端~")
		c.TplName = "video/video.tpl"
	} else {
		// logs.Info("电脑端！")
		c.TplName = "video/video.tpl"
	}
}

// @Title get video
// @Description get video
// @Param page query string true "The page for video list"
// @Param limit query string true "The limit of page for video list"
// @Success 200 {object} models.GetAttachbyId
// @Failure 400 Invalid page supplied
// @Failure 404 pdf not found
// @router /getmonthvideodata [get]
// 触底追加数据
func (c *VideoController) GetMonthVideoData() {
	var yearmonth, videoyearmonth string
	// limit := "5"
	limit := c.GetString("limit")
	if limit == "" {
		limit = "10"
	}
	limit1, err := strconv.Atoi(limit)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("page")
	if page == "" {
		page = "0"
	}
	page1, err := strconv.Atoi(page)
	if err != nil {
		logs.Error(err)
	}
	var offset int
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	videolist, err := models.GetVideoData(limit1, offset)
	if err != nil {
		logs.Error(err)
	}

	// logs.Info(videolist)
	videolist2 := make([]models.Video, 0)
	monthvideolist := make([]MonthVideoList, 0)
	const layoutyearmonth = "2006-01"
	for j := 0; j < len(videolist); j++ {
		videoyearmonth = videolist[j].Created.Format(layoutyearmonth)
		if yearmonth == "" || yearmonth == videoyearmonth {
			yearmonth = videoyearmonth
			photoarr := make([]models.Video, 1)
			photoarr[0] = videolist[j]
			// photoarr[0].Url = videolist[j].Url
			videolist2 = append(videolist2, photoarr...)
		} else {
			yearmonth = videoyearmonth
			monthphotoarr := make([]MonthVideoList, 1)
			monthphotoarr[0].YearMonth = videolist[j-1].Created.Format(layoutyearmonth)
			monthphotoarr[0].VideoLists = videolist2
			monthvideolist = append(monthvideolist, monthphotoarr...)

			videolist2 = make([]models.Video, 0)
			photoarr := make([]models.Video, 1)
			photoarr[0] = videolist[j]
			// photoarr[0].Url = videolist[j].Url
			videolist2 = append(videolist2, photoarr...)
		}
	}
	// 最后一部分
	monthphotoarr := make([]MonthVideoList, 1)
	monthphotoarr[0].YearMonth = yearmonth
	monthphotoarr[0].VideoLists = videolist2
	monthvideolist = append(monthvideolist, monthphotoarr...)

	c.Data["json"] = monthvideolist
	c.ServeJSON()
}

// @Title get videodetail
// @Description get videodetail
// @Param id path string true "The id of video"
// @Success 200 {object} models.GetVideobyId
// @Failure 400 Invalid page supplied
// @Failure 404 pdf not found
// @router /videodetail/:id [get]
func (c *VideoController) VideoDetail() {
	// key转成日期或用photodata[0].Create-1，后者不好减一天，因为有的时间并没有相隔24小时
	// var timeLayoutStr = "2006-01-02"                    //go中的时间格式化必须是这个时间
	// daytime, err := time.Parse(timeLayoutStr, keywords) //string转time
	// oneday, err := time.ParseDuration("24h")
	// if err != nil {
	// 	logs.Error(err)
	// }
	// nextdaytime := daytime.Add(oneday)
	// prevphotodata, err := models.GetPrevPhotoData(daytime)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// if len(prevphotodata) == 0 {
	// 	c.Data["prev"] = "没有了"
	// } else {
	// 	c.Data["prev"] = prevphotodata[0].YearMonthDay
	// }
	// nextphotodata, err := models.GetNextPhotoData(nextdaytime)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// if len(nextphotodata) == 0 {
	// 	c.Data["next"] = "没有了"
	// } else {
	// 	c.Data["next"] = nextphotodata[0].YearMonthDay
	// }
	id := c.Ctx.Input.Param(":id")
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
		return
	}
	video, err := models.GetVideobyId(idNum)
	if err != nil {
		logs.Error(err)
	}
	// logs.Info(video)
	c.Data["Mp4Link"] = video

	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		c.TplName = "video/videodetail.tpl"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "video/videodetail.tpl"
	}
}
