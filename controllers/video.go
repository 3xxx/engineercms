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
	"context"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/minio/minio-go/v7"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
	"path"
	"strconv"
	"strings"
	// "github.com/minio/minio-go/v7/pkg/credentials"
	"bytes"
	"io"
	"regexp"
	"time"
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
		c.TplName = "video/upload_video.tpl"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "video/upload_video.tpl"
	}
}

// @Title post wx video data
// @Description post video
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 video not found
// @router /uploadvideodata [post]
// 上传图片
func (c *PhotoController) UploaVideoData() {
	useminio, err := web.AppConfig.String("useminio")
	if err != nil {
		logs.Error(err)
	}
	year, month, _ := time.Now().Date()
	if useminio == "true" {
		minio_endpoint, err := web.AppConfig.String("minio_endpoint")
		if err != nil {
			logs.Error(err)
		}
		files, err := c.GetFiles("input-ke-2[]")
		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "title": "", "original": "", "data": "获取上传文件错误！"}
			c.ServeJSON()
			return
		}
		minio_bucketname, err := web.AppConfig.String("minio_bucketname")
		if err != nil {
			logs.Error(err)
		}
		videolist := make([]models.Video, 0)
		// GetFiles return multi-upload files
		for i, _ := range files {
			//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				logs.Error(err)
				return
			}

			uploadInfo, err := minioClient.PutObject(context.Background(), minio_bucketname, strconv.Itoa(year)+month.String()+"/"+files[i].Filename, file, -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
			if err != nil {
				fmt.Println(err)
				return
			}
			// 提取封面
			fileSuffix := path.Ext(files[i].Filename)
			filenameOnly := strings.TrimSuffix(files[i].Filename, fileSuffix)
			reader := ExampleReadFrameAsJpeg("http://"+minio_endpoint+"/"+minio_bucketname+"/"+strconv.Itoa(year)+month.String()+"/"+files[i].Filename, 1)
			img, err := imaging.Decode(reader)
			if err != nil {
				logs.Error(err)
			}
			err = imaging.Save(img, "http://"+minio_endpoint+"/"+minio_bucketname+"/"+strconv.Itoa(year)+month.String()+"/"+filenameOnly+".jpg")
			if err != nil {
				logs.Error(err)
			}

			// fmt.Println("Successfully uploaded bytes: ", uploadInfo)
			fmt.Println("etag: ", uploadInfo.ETag)

			videoarr := make([]models.Video, 1)
			temptime := time.Now()
			videoarr[0].Url = "http://" + minio_endpoint + "/" + minio_bucketname + "/" + strconv.Itoa(year) + month.String() + "/" + files[i].Filename
			videoarr[0].CoverUrl = "http://" + minio_endpoint + "/" + minio_bucketname + "/" + strconv.Itoa(year) + month.String() + "/" + filenameOnly + ".jpg"
			videoarr[0].Created = temptime
			videoarr[0].Updated = temptime
			videolist = append(videolist, videoarr...)
		}
		err = models.AddVideoData(videolist)
		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"state": "ERROR", "errNo": 1, "info": "插入video数据失败！", "data": "插入video数据失败！", "msg": "插入video数据失败！"}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "errNo": 0, "info": "插入video数据成功！", "data": "插入video数据成功！", "msg": "插入video数据成功！"}
			c.ServeJSON()
		}
	} else {
		//获取上传的文件
		files, err := c.GetFiles("input-ke-2[]")
		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "title": "", "original": "", "data": "获取上传文件错误！"}
			c.ServeJSON()
			return
		}
		videolist := make([]models.Video, 0)
		for i, _ := range files {
			//	//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				logs.Error(err)
				return
			}
			fileSuffix := path.Ext(files[i].Filename)
			// filenameOnly := strings.TrimSuffix(files[i].Filename, fileSuffix)
			// random_name
			nanoname := strconv.FormatInt(time.Now().UnixNano(), 10)
			newname := nanoname + fileSuffix // + "_" + filename

			cwd, _ := os.Getwd()
			replacevideopath := strings.Replace(cwd, "\\", "/", -1)

			DiskDirectory := replacevideopath + "/attachment/video/"
			// year, month, _ := time.Now().Date()
			err = os.MkdirAll(DiskDirectory+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
			if err != nil {
				logs.Error(err)
				return
			}

			imagepath := DiskDirectory + strconv.Itoa(year) + month.String() + "/" + newname

			//create destination file making sure the path is writeable.
			dst, err := os.Create(imagepath)
			defer dst.Close()
			if err != nil {
				logs.Error(err)
				return
			}
			//copy the uploaded file to the destination file  //golang实现http表单大文件流式上传服务端代码
			if _, err := io.Copy(dst, file); err != nil {
				logs.Error(err)
				c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "title": "", "original": "", "data": "文件保存错误！"}
				c.ServeJSON()
				return
			}

			// reader := ExampleReadFrameAsJpeg(file, 5)
			// logs.Info(imagepath)
			reader := ExampleReadFrameAsJpeg(imagepath, 1)
			img, err := imaging.Decode(reader)
			if err != nil {
				logs.Error(err)
			}
			err = imaging.Save(img, DiskDirectory+strconv.Itoa(year)+month.String()+"/"+nanoname+".jpg")
			if err != nil {
				logs.Error(err)
			}

			videoarr := make([]models.Video, 1)
			temptime := time.Now()
			videoarr[0].Url = "/attachment/video/" + strconv.Itoa(year) + month.String() + "/" + newname
			videoarr[0].CoverUrl = "/attachment/video/" + strconv.Itoa(year) + month.String() + "/" + nanoname + ".jpg"
			videoarr[0].Created = temptime
			videoarr[0].Updated = temptime
			videolist = append(videolist, videoarr...)
		}
		// 写入数据库
		err = models.AddVideoData(videolist)
		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "info": "写入数据库出错！！"}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "data": videolist, "info": "上传成功！！"}
			c.ServeJSON()
		}
	}
}

// func ExampleReadFrameAsJpeg(inFileName io.Reader, frameNum int) io.Reader {
func ExampleReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	logs.Info(inFileName)
	buf := bytes.NewBuffer(nil)
	// err := ffmpeg.Input("pipe:0", ffmpeg.KwArgs{"format": "mp4"}).
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		// WithInput(inFileName).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf

	// fd, err := os.Open("./file.mp3")
	// if err != nil {
	// 	panic(err)
	// }
	// err = ffmpeg_go.Input("pipe:", ffmpeg_go.KwArgs{"format": "mp3"}).
	// 	WithInput(fd).
	// 	Output("./playlist.m3u8", ffmpeg_go.KwArgs{"c:a": "aac", "b:a": "320k", "hls_list_size": "0", "hls_time": "20"}).
	// 	OverWriteOutput().ErrorToStdOut().Run()

	// telegramFile, err := r.bot.GetFile(&msg.VideoNote.File)

	// if err != nil {
	// 	// handling error
	// }
	// audio := bytes.NewBuffer(nil)
	// err = ffmpeg.Input(
	// 	"pipe:0", ffmpeg.KwArgs{"format": "mp4"},
	// ).
	// 	Output("pipe:1", ffmpeg.KwArgs{"format": "wav"}).WithInput(telegramFile).WithOutput(audio).Run()
	// if err != nil {
	// 	// handling error
	// }

	// buf = bytes.NewBuffer([]byte)
	// err := ffmpeg.Input("pipe:", ffmpeg.KwArgs{}).
	// 	Output(outfileName, ffmpeg.KwArgs{}).
	// 	OverWriteOutput().
	// 	WithInput(buf).Run()

	// reader, err := header.Open()
	// if err != nil {
	// 	return
	// }

	// err = ffmpeg.Input("pipe:").WithInput(reader).
	// 	Overlay(ffmpeg.Input("./logo/logo.png"), "").
	// 	Output("lucene.m3u8", ffmpeg.KwArgs{"c:v": "libx265", "hls_time": "12", "hls_list_size": "0", "hls_segment_filename": "file%d.ts"}).
	// 	ErrorToStdOut().
	// 	OverWriteOutput().
	// 	Run()
	// log.Println("ffmpeg process1 done")
	// if err != nil {
	// 	panic(err)
	// }
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
