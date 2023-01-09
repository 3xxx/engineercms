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

// @Title upload videodata
// @Description upload video
// @Success 200 {object} models.uploadvideo
// @Failure 400 Invalid page supplied
// @Failure 404 video not found
// @router /uploadvideo [get]
func (c *VideoController) UploadVideo() {
	// beego.Info(c.Ctx.Input.UserAgent())
	u := c.Ctx.Input.UserAgent()
	// re := regexp.MustCompile("Trident")
	// loc := re.FindStringIndex(u)
	// loc[0] > 1
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		c.TplName = "photo/upload_video.tpl"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "photo/upload_video.tpl"
	}
}

// @Title post wx video
// @Description post video
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 video not found
// @router /uploadvideodata [post]
//
// func (c *VideoController) UploadVideoData() {
// 	photopath, err := web.AppConfig.String("photopath")
// 	if err != nil {
// 		logs.Error(err)
// 	}
// 	cwd, _ := os.Getwd()
// 	replacephotopath := strings.Replace(cwd, "\\", "/", -1)

// 	DiskDirectory := replacephotopath + photopath
// 	logs.Info(DiskDirectory)
// 	//获取上传的文件
// 	_, h, err := c.GetFile("input-ke-2[]")
// 	if err != nil {
// 		logs.Error(err)
// 	}
// 	fileSuffix := path.Ext(h.Filename)
// 	// random_name
// 	nanoname := strconv.FormatInt(time.Now().UnixNano(), 10)
// 	newname := nanoname + fileSuffix // + "_" + filename
// 	small_newname := nanoname + "_small" + fileSuffix
// 	year, month, _ := time.Now().Date()
// 	err = os.MkdirAll(DiskDirectory+"/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
// 	if err != nil {
// 		logs.Error(err)
// 	}
// 	var imagepath, new_imagepath, Url string
// 	var filesize int64
// 	if h != nil {
// 		//保存附件
// 		imagepath = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + newname
// 		logs.Info(imagepath)
// 		new_imagepath = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + small_newname
// 		Url = photopath + strconv.Itoa(year) + month.String() + "/"
// 		logs.Info(Url)
// 		err = c.SaveToFile("input-ke-2[]", imagepath) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
// 		if err != nil {
// 			logs.Error(err)
// 			c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "title": "", "original": "", "data": "文件保存错误！"}
// 			c.ServeJSON()
// 			return
// 		}
// 		filesize, _ = FileSize(imagepath)
// 		filesize = filesize / 1000.0

// 		//*****压缩图片***
// 		file, err := os.Open(imagepath)
// 		if err != nil {
// 			// log.Fatal(err)
// 			logs.Error(err)
// 		}
// 		defer file.Close()
// 		var img image.Image
// 		var typeImage int
// 		// ext := filepath.Ext(imagepath)
// 		if strings.EqualFold(fileSuffix, ".jpg") || strings.EqualFold(fileSuffix, ".jpeg") {
// 			img, err = jpeg.Decode(file)
// 			if err != nil {
// 				// log.Fatal(err)
// 				logs.Error(err)
// 			}
// 			typeImage = 0
// 		} else if strings.EqualFold(fileSuffix, ".png") {
// 			// decode png into image.Image
// 			img, err = png.Decode(file)
// 			if err != nil {
// 				// log.Fatal(err)
// 				logs.Error(err)
// 			}
// 			typeImage = 1
// 		} else if strings.EqualFold(fileSuffix, ".gif") {
// 			img, err = gif.Decode(file)
// 			if err != nil {
// 				// log.Fatal(err)
// 				logs.Error(err)
// 			}
// 			typeImage = 2
// 		}

// 		// file.Close()

// 		// resize to width 1000 using Lanczos resampling
// 		// and preserve aspect ratio
// 		m := resize.Resize(1000, 0, img, resize.Lanczos3)
// 		// m := resize.Thumbnail(1000, 0, img, resize.Lanczos3)

// 		out, err := os.Create(new_imagepath)
// 		defer out.Close()
// 		if err != nil {
// 			logs.Error(err)
// 		}
// 		if typeImage == 0 {
// 			err = jpeg.Encode(out, m, &jpeg.Options{Quality: 80})
// 			if err != nil {
// 				logs.Error(err)
// 			}
// 		} else {
// 			err = png.Encode(out, m)
// 			if err != nil {
// 				logs.Error(err)
// 			}
// 		}

// 		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": Url + small_newname, "title": "111", "original": "demo.jpg"}
// 		c.ServeJSON()
// 	} else {
// 		c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "title": "", "original": ""}
// 		c.ServeJSON()
// 	}
// }

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
