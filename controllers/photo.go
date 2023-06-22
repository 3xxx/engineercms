package controllers

import (
	// "encoding/json"
	"github.com/3xxx/engineercms/controllers/utils/filetil"
	"github.com/3xxx/engineercms/models"
	// "github.com/3xxx/engineercms/vendor/github.com/minio/minio-go/v7"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	// beego "github.com/beego/beego/v2/adapter"

	"io"
	// "encoding/json"
	// "net/http"
	// "github.com/beego/beego/v2/adapter/httplib"
	// "github.com/beego/beego/v2/server/web"
	// "bytes"
	// "io/ioutil"
	// "mime/multipart"
	// "bytes"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var minioClient *minio.Client // 超级惨！！！下面的minioClient不能用:=，否则重新定义了啊！！！

func init() {
	// http://3.1.1.38:59735/api/v1/buckets/engineercms/objects/download?prefix=Mi3lub/kuJznnIHogYznp7Dor4TlrqHooagyMDIyMTEyMC5kb2N4
	// http://3.1.1.38:59735/api/v1/buckets/engineercms/objects/upload?prefix=Lw==
	useminio, err := web.AppConfig.String("useminio")
	if err != nil {
		logs.Error(err)
	}
	minio_endpoint, err := web.AppConfig.String("minio_endpoint")
	if err != nil {
		logs.Error(err)
	}
	if useminio == "true" && minio_endpoint != "" {
		endpoint := minio_endpoint //"3.1.1.156:9000"

		minio_accessKeyID, err := web.AppConfig.String("minio_accessKeyID")
		if err != nil {
			logs.Error(err)
		}
		minio_secretAccessKey, err := web.AppConfig.String("minio_secretAccessKey")
		if err != nil {
			logs.Error(err)
		}
		minio_useSSL, err := web.AppConfig.String("minio_useSSL")
		if err != nil {
			logs.Error(err)
		}

		var useSSL bool
		accessKeyID := minio_accessKeyID         //"FjtR8D8nvAQuAG3h"
		secretAccessKey := minio_secretAccessKey //"M3NSzXVmuuArai7xBMuZIKOgOXuKovxi"
		if minio_useSSL == "true" {
			useSSL = true
		} else if minio_useSSL == "false" {
			useSSL = false
		} else {
			useSSL = false
		}
		// Initialize minio client object.
		// 超级惨！！！下面的minioClient不能用:=，否则重新定义了啊！！！
		minioClient, err = minio.New(endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: useSSL,
		})
		if err != nil {
			logs.Error(err)
		}

		log.Printf("%#v\n", minioClient) // minioClient is now setup
	}

	// buckets, err := minioClient.ListBuckets(context.Background())
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for _, bucket := range buckets {
	// 	fmt.Println(bucket)
	// }

	// ctx, cancel := context.WithCancel(context.Background())

	// defer cancel()

	// objectCh := minioClient.ListObjects(ctx, "engineercms", minio.ListObjectsOptions{
	// 	Prefix:    "1",
	// 	Recursive: true,
	// })
	// for object := range objectCh {
	// 	if object.Err != nil {
	// 		fmt.Println(object.Err)
	// 		return
	// 	}
	// 	fmt.Println(object)
	// }
}

type PhotoController struct {
	web.Controller
}

// @Title upload photo page
// @Description upload photo page
// @Success 200 {object} models.UploadPhoto
// @Failure 400 Invalid page supplied
// @Failure 404 photo not found
// @router /uploadphoto [get]
func (c *PhotoController) UploadPhoto() {
	logs.Info(c.Ctx.Input.UserAgent())
	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		c.TplName = "photo/upload_photo.tpl"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "photo/upload_photo.tpl"
	}
}

// @Title post wx photo img
// @Description post photo
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 photo not found
// @router /uploadphotodata [post]
// 上传图片
func (c *PhotoController) UploadPhotoData() {
	useminio, err := web.AppConfig.String("useminio")
	if err != nil {
		logs.Error(err)
	}
	year, month, _ := time.Now().Date()
	const layout = "2006-01-02"
	const layoutyearmonth = "2006-01"
	photolist := make([]models.PhotoData, 0)
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

		photolist := make([]models.PhotoData, 0)

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
			// fmt.Println("Successfully uploaded bytes: ", uploadInfo)
			fmt.Println("etag: ", uploadInfo.ETag)

			if filetil.IsImageExt(files[i].Filename) && !strings.Contains(files[i].Filename, "_small") {
				photoarr := make([]models.PhotoData, 1)
				temptime := time.Now()

				photoarr[0].YearMonth = temptime.Format(layoutyearmonth)
				photoarr[0].YearMonthDay = temptime.Format(layout)
				photoarr[0].Url = "http://" + minio_endpoint + "/" + minio_bucketname + "/" + strconv.Itoa(year) + month.String() + "/" + files[i].Filename
				photoarr[0].CreatedAt = temptime
				photolist = append(photolist, photoarr...)
			}
		}
		// err = models.AddPhotoData(photolist)
		// if err != nil {
		// 	logs.Error(err)
		// 	c.Data["json"] = map[string]interface{}{"state": "ERROR", "errNo": 1, "info": "插入photo数据失败！", "data": "插入photo数据失败！", "msg": "插入photo数据失败！"}
		// 	c.ServeJSON()
		// } else {
		// 	c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "errNo": 0, "info": "插入photo数据成功！", "data": "插入photo数据成功！", "msg": "插入photo数据成功！"}
		// 	c.ServeJSON()
		// }
	} else {
		photopath, err := web.AppConfig.String("photopath")
		if err != nil {
			logs.Error(err)
		}
		cwd, _ := os.Getwd()
		replacephotopath := strings.Replace(cwd, "\\", "/", -1)

		DiskDirectory := replacephotopath + photopath
		// logs.Info(DiskDirectory)
		//获取上传的文件
		files, err := c.GetFiles("input-ke-2[]")
		// func (c *Controller) GetFiles(key string) ([]*multipart.FileHeader, error) {
		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "title": "", "original": "", "data": "获取上传文件错误！"}
			c.ServeJSON()
			return
		}
		for i, _ := range files {
			//	//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			// func (fh *FileHeader) Open() (File, error)
			// type File interface {
			// 	io.Reader
			// 	io.ReaderAt
			// 	io.Seeker
			// 	io.Closer
			// }

			defer file.Close()
			if err != nil {
				logs.Error(err)
				return
			}
			fileSuffix := path.Ext(files[i].Filename)
			// random_name
			nanoname := strconv.FormatInt(time.Now().UnixNano(), 10)
			newname := nanoname + fileSuffix // + "_" + filename
			small_newname := nanoname + "_small" + fileSuffix
			// year, month, _ := time.Now().Date()
			err = os.MkdirAll(DiskDirectory+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
			if err != nil {
				logs.Error(err)
			}
			var imagepath, new_imagepath string
			//保存附件
			imagepath = DiskDirectory + strconv.Itoa(year) + month.String() + "/" + newname
			// logs.Info(imagepath)
			new_imagepath = DiskDirectory + strconv.Itoa(year) + month.String() + "/" + small_newname
			Url := photopath + strconv.Itoa(year) + month.String() + "/"
			//create destination file making sure the path is writeable.
			dst, err := os.Create(imagepath)
			defer dst.Close()
			if err != nil {
				logs.Error(err)
				return
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				logs.Error(err)
				c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "title": "", "original": "", "data": "文件保存错误！"}
				c.ServeJSON()
				return
			}

			//*****压缩图片***
			file, err = os.Open(imagepath)
			// func Open(name string) (file *File, err error)
			if err != nil {
				logs.Error(err)
			}
			defer file.Close()
			var img image.Image
			var typeImage int
			// ext := filepath.Ext(imagepath)
			if strings.EqualFold(fileSuffix, ".jpg") || strings.EqualFold(fileSuffix, ".jpeg") {
				img, err = jpeg.Decode(file)
				// func Decode(r io.Reader) (image.Image, error)
				if err != nil {
					// log.Fatal(err)
					logs.Error(err)
				}
				typeImage = 0
			} else if strings.EqualFold(fileSuffix, ".png") {
				// decode png into image.Image
				img, err = png.Decode(file)
				if err != nil {
					// log.Fatal(err)
					logs.Error(err)
				}
				typeImage = 1
			} else if strings.EqualFold(fileSuffix, ".gif") {
				img, err = gif.Decode(file)
				if err != nil {
					// log.Fatal(err)
					logs.Error(err)
				}
				typeImage = 2
			}

			// file.Close()

			// resize to width 1000 using Lanczos resampling
			// and preserve aspect ratio
			m := resize.Resize(1000, 0, img, resize.Lanczos3)
			// m := resize.Thumbnail(1000, 0, img, resize.Lanczos3)

			out, err := os.Create(new_imagepath)
			defer out.Close()
			if err != nil {
				logs.Error(err)
			}
			if typeImage == 0 {
				err = jpeg.Encode(out, m, &jpeg.Options{Quality: 80})
				if err != nil {
					logs.Error(err)
				}
			} else {
				err = png.Encode(out, m)
				if err != nil {
					logs.Error(err)
				}
			}

			if filetil.IsImageExt(files[i].Filename) && !strings.Contains(files[i].Filename, "_small") {
				photoarr := make([]models.PhotoData, 1)
				temptime := time.Now()

				photoarr[0].YearMonth = temptime.Format(layoutyearmonth)
				photoarr[0].YearMonthDay = temptime.Format(layout)
				photoarr[0].Url = Url + newname
				// photoarr[0].CreatedAt = temptime
				photolist = append(photolist, photoarr...)
			}
		}
	}
	err = models.AddPhotoData(photolist)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "errNo": 1, "info": "插入photo数据失败！", "data": "插入photo数据失败！", "msg": "插入photo数据失败！"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "errNo": 0, "info": "插入photo数据成功！", "data": photolist, "msg": "插入photo数据成功！"}
		c.ServeJSON()
	}
}

type MonthPhotoList struct {
	YearMonth string
	// PhotoLists []PhotoList
	PhotoLists []models.PhotoData
}

type PhotoList struct {
	YearMonth    string
	YearMonthDay string
	Url          string
	ThumbnailUrl string
	PhotoName    string
}

// @Title get photo
// @Description get photo
// @Success 200 {object} models.GetAttachbyId
// @Failure 400 Invalid page supplied
// @Failure 404 pdf not found
// @router /photo_back [get]
// 纯文件夹检索，作废，保留
func (c *PhotoController) Photo_back() {
	// var yearmonthday, yearmonth string
	// cwd, _ := os.Getwd()
	// logs.Info(cwd) //D:\gowork\src\github.com\3xxx\engineercms

	// photopath, err := web.AppConfig.String("photopath")
	// if err != nil {
	// 	logs.Error(err)
	// }
	// // replacephotopath, err := web.AppConfig.String("replacephotopath")
	// // if err != nil {
	// // 	logs.Error(err)
	// // }
	// replacephotopath := strings.Replace(cwd, "\\", "/", -1)
	// logs.Info(replacephotopath)
	// workpath := cwd + photopath
	// filelist, err := filetil.ScanFiles(workpath)
	// photolist := make([]PhotoList, 0)
	// monthphotolist := make([]MonthPhotoList, 0)
	// for _, v := range filelist {
	// 	if filetil.IsImageExt(v.Name) {
	// 		photoarr := make([]PhotoList, 1)
	// 		temptime := time.Unix(v.ModTime, 0)
	// 		const layout = "2006-01-02"
	// 		const layoutyearmonth = "2006-01"
	// 		// logs.Info(temptime.Format(layout)) //2022-05-03
	// 		// logs.Info(temptime.Year())
	// 		// logs.Info(temptime.Month())
	// 		if yearmonthday != temptime.Format(layout) {
	// 			yearmonthday = temptime.Format(layout)
	// 			photoarr[0].YearMonth = temptime.Format(layoutyearmonth)
	// 			photoarr[0].YearMonthDay = temptime.Format(layout)
	// 			photoarr[0].PhotoUrl = strings.Replace(v.Path, replacephotopath, "", -1)
	// 			photolist = append(photolist, photoarr...)
	// 		}
	// 		// D:/gowork/src/github.com/3xxx/engineercms
	// 	}
	// }

	// photolist2 := make([]PhotoList, 0)
	// for j := len(photolist) - 1; j > -1; j-- {
	// 	// for j, w := range photolist {
	// 	if yearmonth == "" || yearmonth == photolist[j].YearMonth {
	// 		yearmonth = photolist[j].YearMonth
	// 		photoarr := make([]PhotoList, 1)
	// 		photoarr[0] = photolist[j]
	// 		photolist2 = append(photolist2, photoarr...)
	// 	} else {
	// 		yearmonth = photolist[j].YearMonth
	// 		monthphotoarr := make([]MonthPhotoList, 1)
	// 		monthphotoarr[0].YearMonth = photolist[j+1].YearMonth
	// 		monthphotoarr[0].PhotoLists = photolist2
	// 		monthphotolist = append(monthphotolist, monthphotoarr...)

	// 		photolist2 = make([]PhotoList, 0)
	// 		photoarr := make([]PhotoList, 1)
	// 		photoarr[0] = photolist[j]
	// 		photolist2 = append(photolist2, photoarr...)
	// 	}
	// }
	// // 最后一部分
	// monthphotoarr := make([]MonthPhotoList, 1)
	// monthphotoarr[0].YearMonth = yearmonth
	// monthphotoarr[0].PhotoLists = photolist2
	// monthphotolist = append(monthphotolist, monthphotoarr...)

	// c.Data["PhotoList"] = photolist
	// c.Data["MonthPhotoList"] = monthphotolist
	// u := c.Ctx.Input.UserAgent()
	// // re := regexp.MustCompile("Trident")
	// // loc := re.FindStringIndex(u)
	// // loc[0] > 1
	// matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// if matched == true {
	// 	// logs.Info("移动端~")
	// 	c.TplName = "photo/photo.tpl"
	// } else {
	// 	// logs.Info("电脑端！")
	// 	c.TplName = "photo/photo.tpl"
	// }
}

// @Title get photo
// @Description get photo
// @Param page query string true "The page for photos list"
// @Param limit query string true "The limit of page for photos list"
// @Success 200 {object} models.GetAttachbyId
// @Failure 400 Invalid page supplied
// @Failure 404 pdf not found
// @router /photo [get]
func (c *PhotoController) Photo() {
	var yearmonth string
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

	photolist, err := models.GetPhotoData(limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	// logs.Info(photolist)
	photolist2 := make([]models.PhotoData, 0)
	monthphotolist := make([]MonthPhotoList, 0)
	// for j := len(photolist) - 1; j > -1; j-- {
	for j := 0; j < len(photolist); j++ {
		// for j, w := range photolist {
		if yearmonth == "" || yearmonth == photolist[j].YearMonth {
			yearmonth = photolist[j].YearMonth
			photoarr := make([]models.PhotoData, 1)
			photoarr[0] = photolist[j]
			if !strings.Contains(photolist[j].Url, "http://") {
				photoarr[0].Url = strings.Replace(photolist[j].Url, ".jpg", "_small.jpg", -1)
			}
			photolist2 = append(photolist2, photoarr...)
		} else {
			yearmonth = photolist[j].YearMonth
			monthphotoarr := make([]MonthPhotoList, 1)
			monthphotoarr[0].YearMonth = photolist[j-1].YearMonth
			monthphotoarr[0].PhotoLists = photolist2
			monthphotolist = append(monthphotolist, monthphotoarr...)

			photolist2 = make([]models.PhotoData, 0)
			photoarr := make([]models.PhotoData, 1)
			photoarr[0] = photolist[j]
			if !strings.Contains(photolist[j].Url, "http://") {
				photoarr[0].Url = strings.Replace(photolist[j].Url, ".jpg", "_small.jpg", -1)
			} else {
				photoarr[0].Url = photolist[j].Url
			}
			photolist2 = append(photolist2, photoarr...)
		}
	}
	// 最后一部分
	monthphotoarr := make([]MonthPhotoList, 1)
	monthphotoarr[0].YearMonth = yearmonth
	monthphotoarr[0].PhotoLists = photolist2
	monthphotolist = append(monthphotolist, monthphotoarr...)

	// c.Data["PhotoList"] = photodata
	c.Data["MonthPhotoList"] = monthphotolist
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
		c.TplName = "photo/photo.tpl"
	} else {
		// logs.Info("电脑端！")
		c.TplName = "photo/photo.tpl"
	}
}

// @Title get photo
// @Description get photo
// @Param page query string true "The page for photos list"
// @Param limit query string true "The limit of page for photos list"
// @Success 200 {object} models.GetAttachbyId
// @Failure 400 Invalid page supplied
// @Failure 404 pdf not found
// @router /getmonthphotodata [get]
// 触底追加数据
func (c *PhotoController) GetMonthPhotoData() {
	var yearmonth string
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

	photolist, err := models.GetPhotoData(limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	// logs.Info(photolist)
	photolist2 := make([]models.PhotoData, 0)
	monthphotolist := make([]MonthPhotoList, 0)
	// for j := len(photolist) - 1; j > -1; j-- {
	for j := 0; j < len(photolist); j++ {
		// for j, w := range photolist {
		if yearmonth == "" || yearmonth == photolist[j].YearMonth {
			yearmonth = photolist[j].YearMonth
			photoarr := make([]models.PhotoData, 1)
			photoarr[0] = photolist[j]
			if !strings.Contains(photolist[j].Url, "http://") {
				photoarr[0].Url = strings.Replace(photolist[j].Url, ".jpg", "_small.jpg", -1)
			} else {
				photoarr[0].Url = photolist[j].Url
			}
			photolist2 = append(photolist2, photoarr...)
		} else {
			yearmonth = photolist[j].YearMonth
			monthphotoarr := make([]MonthPhotoList, 1)
			monthphotoarr[0].YearMonth = photolist[j-1].YearMonth
			monthphotoarr[0].PhotoLists = photolist2
			monthphotolist = append(monthphotolist, monthphotoarr...)

			photolist2 = make([]models.PhotoData, 0)
			photoarr := make([]models.PhotoData, 1)
			photoarr[0] = photolist[j]
			if !strings.Contains(photolist[j].Url, "http://") {
				photoarr[0].Url = strings.Replace(photolist[j].Url, ".jpg", "_small.jpg", -1)
			} else {
				photoarr[0].Url = photolist[j].Url
			}
			photolist2 = append(photolist2, photoarr...)
		}
	}
	// 最后一部分
	monthphotoarr := make([]MonthPhotoList, 1)
	monthphotoarr[0].YearMonth = yearmonth
	monthphotoarr[0].PhotoLists = photolist2
	monthphotolist = append(monthphotolist, monthphotoarr...)

	c.Data["json"] = monthphotolist
	c.ServeJSON()
}

// @Title get photo
// @Description get photo
// @Success 200 {object} models.GetAttachbyId
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /createphotodata [get]
// 图片数据写入数据库
func (c *PhotoController) CreatePhotoData() {
	cwd, _ := os.Getwd()
	// logs.Info(cwd)
	photopath, err := web.AppConfig.String("photopath")
	if err != nil {
		logs.Error(err)
	}
	// replacephotopath, err := web.AppConfig.String("replacephotopath")
	// if err != nil {
	// 	logs.Error(err)
	// }
	replacephotopath := strings.Replace(cwd, "\\", "/", -1)
	// logs.Info(replacephotopath)
	workpath := cwd + photopath
	filelist, err := filetil.ScanFiles(workpath)
	photolist := make([]models.PhotoData, 0)
	for _, v := range filelist {
		if filetil.IsImageExt(v.Name) && !strings.Contains(v.Name, "_small") {
			photoarr := make([]models.PhotoData, 1)
			temptime := time.Unix(v.ModTime, 0) // 使用time.Unix()函数可以将时间戳转换为时间格式
			const layout = "2006-01-02"
			const layoutyearmonth = "2006-01"
			// logs.Info(temptime.Format(layout)) //2022-05-03
			// logs.Info(temptime.Year())
			// logs.Info(temptime.Month())
			// if yearmonthday != temptime.Format(layout) {
			// yearmonthday = temptime.Format(layout)
			photoarr[0].YearMonth = temptime.Format(layoutyearmonth)
			photoarr[0].YearMonthDay = temptime.Format(layout)
			photoarr[0].Url = strings.Replace(v.Path, replacephotopath, "", -1)
			photoarr[0].CreatedAt = temptime
			photolist = append(photolist, photoarr...)
			// }
			// D:/gowork/src/github.com/3xxx/engineercms
		}
	}
	err = models.AddPhotoData(photolist)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "errNo": 1, "info": "插入photo数据失败！", "data": "插入photo数据失败！", "msg": "插入photo数据失败！"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "errNo": 0, "info": "插入photo数据成功！", "data": "插入photo数据成功！", "msg": "插入photo数据成功！"}
		c.ServeJSON()
	}
}

// @Title get photodetail
// @Description get photodetail
// @Success 200 {object} models.GetAttachbyId
// @Failure 400 Invalid page supplied
// @Failure 404 pdf not found
// @router /photodetail_back [get]
// 纯文件夹检索，作废
func (c *PhotoController) PhotoDetail_back() {
	// keywords := c.GetString("keywords")
	// // logs.Info(keywords)
	// cwd, _ := os.Getwd()
	// // logs.Info(cwd)
	// photopath, err := web.AppConfig.String("photopath")
	// if err != nil {
	// 	logs.Error(err)
	// }
	// replacephotopath, err := web.AppConfig.String("replacephotopath")
	// if err != nil {
	// 	logs.Error(err)
	// }
	// replacephotopath := strings.Replace(cwd, "\\", "/", -1)
	// logs.Info(replacephotopath)
	// workpath := cwd + photopath //"/attachment/pass-001mathcad云计算书/"
	// filelist, err := filetil.ScanFiles(workpath)
	// photolist := make([]PhotoList, 0)
	// for _, v := range filelist {
	// 	if filetil.IsImageExt(v.Name) {
	// 		photoarr := make([]PhotoList, 1)
	// 		temptime := time.Unix(v.ModTime, 0)
	// 		const layout = "2006-01-02"
	// 		if temptime.Format(layout) == keywords {
	// 			photoarr[0].YearMonthDay = temptime.Format(layout)
	// 			photoarr[0].PhotoUrl = strings.Replace(v.Path, replacephotopath, "", -1)
	// 			photolist = append(photolist, photoarr...)
	// 		}
	// 	}
	// }
	// c.Data["PhotoList"] = photolist
	// u := c.Ctx.Input.UserAgent()
	// matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// if matched == true {
	// 	// beego.Info("移动端~")
	// 	c.TplName = "photo/photodetail.tpl"
	// } else {
	// 	// beego.Info("电脑端！")
	// 	c.TplName = "photo/photodetail.tpl"
	// }
}

// @Title get photodetail
// @Description get photodetail
// @Success 200 {object} models.GetAttachbyId
// @Failure 400 Invalid page supplied
// @Failure 404 pdf not found
// @router /photodetail [get]
func (c *PhotoController) PhotoDetail() {
	keywords := c.GetString("keywords")
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
	photodata, err := models.GetDayPhotoData(keywords, limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	photolists := make([]PhotoList, 0)
	var small_newname string
	for _, v := range photodata {
		photolistarr := make([]PhotoList, 1)
		photolistarr[0].Url = v.Url

		fileSuffix := path.Ext(v.Url)
		if !strings.Contains(v.Url, "http://") {
			small_newname = strings.Replace(v.Url, fileSuffix, "_small"+fileSuffix, -1)
		} else {
			small_newname = v.Url
		}

		photolistarr[0].ThumbnailUrl = small_newname //strings.Replace(v.Url, ".jpg", "_small.jpg", -1)
		photolistarr[0].YearMonth = v.YearMonth
		photolistarr[0].YearMonthDay = v.YearMonthDay
		photolists = append(photolists, photolistarr...)
	}

	// key转成日期或用photodata[0].Create-1，后者不好减一天，因为有的时间并没有相隔24小时
	var timeLayoutStr = "2006-01-02"                    //go中的时间格式化必须是这个时间
	daytime, err := time.Parse(timeLayoutStr, keywords) //string转time
	// logs.Info(daytime)
	oneday, err := time.ParseDuration("24h")
	if err != nil {
		logs.Error(err)
	}
	nextdaytime := daytime.Add(oneday)
	// logs.Info(nextdaytime)
	prevphotodata, err := models.GetPrevPhotoData(daytime)
	if err != nil {
		logs.Error(err)
	}
	if len(prevphotodata) == 0 {
		c.Data["prev"] = "没有了"
	} else {
		c.Data["prev"] = prevphotodata[0].YearMonthDay
	}

	nextphotodata, err := models.GetNextPhotoData(nextdaytime)
	if err != nil {
		logs.Error(err)
	}
	if len(nextphotodata) == 0 {
		c.Data["next"] = "没有了"
	} else {
		c.Data["next"] = nextphotodata[0].YearMonthDay
	}

	c.Data["PhotoList"] = photolists

	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		c.TplName = "photo/photodetail.tpl"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "photo/photodetail.tpl"
	}
}

// 从MINIO获取文件然后返回前台文件流
// minioObject, getObjErr := minioClient.GetObject(ConfigUtil.MinioBucket, key, minio.GetObjectOptions{})
// if getObjErr != nil {
//    c.JSON(http.StatusOK, gin.H{"success": false, "info": "获取文件信息异常！" + getObjErr.Error()})
//    return
// }

// buf := bytes.NewBuffer(nil)
// io.Copy(buf, minioObject)

// c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", originName))
// c.Data(http.StatusOK, "application/octet-stream", buf.Bytes())

// func DownloadFileHandler(c *gin.Context) {
// 	taskid := c.Query("taskid")
// 	fmt.Println(taskid)
// 	id, err := strconv.Atoi(taskid)
// 	if err != nil {
// 	}
// 	err, data := logic.Downloadfile(id)
// 	if err != nil {
// 	}
// 	fmt.Println(len(data))
// 	// 设置返回头并返回数据
// 	c.Header("Content-Type", "application/octet-stream")
// 	c.Header("Content-Disposition", "attachement;filename=\""+"你的文件名.jpg"+"\"")
// 	c.Data(http.StatusOK, "application/octet-stream", data)
// }

// logic层：
// 在转化[]byte 这块思路一直有点乱。
// 方法也很多，这里总结了四种方法去返回[]byte数据
// func Downloadfile(taskid int) (err error, data []byte) {
// 	ctx := context.Background()
// 	filenames, err := mysql.FindTaskFileName(taskid)
// 	if err != nil {
// 		return err, nil
// 	}
// 	fmt.Println(filenames)
// 	object, err := miniocli.MinIOClient.GetObject(ctx, miniocli.BucketPicture, filenames[0], minio.GetObjectOptions{})
// 	if err != nil {
// 		logger.Lg.Error("MinIOClient.GetObject failed ...", zap.Error(err))
// 		return err, nil
// 	}
// 	fmt.Printf("object:%v\n", object)
// 	//方法1⃣️：
// 	//bowl := bytes.Buffer{}
// 	//num, err := bowl.ReadFrom(object)
// 	//fmt.Println(num)
// 	//return nil, bowl.Bytes()
// 	//方法2⃣️：
// 	//var chunk []byte
// 	//buf := make([]byte, 1024)
// 	//for {
// 	//	//从file读取到buf中
// 	//	//Read reads up to len(b) bytes into b.
// 	//	//It returns the number of bytes read (0 <= n <= len(b)) and any error encountered.
// 	//	//Returns io.EOF upon end of file.
// 	//	n, err := object.Read(buf)
// 	//	if err != nil && err != io.EOF {
// 	//		fmt.Println("read buf fail", err)
// 	//		return err, nil
// 	//	}
// 	//	//说明读取结束
// 	//	if n == 0 {
// 	//		fmt.Println("n=0")
// 	//		break
// 	//	}
// 	//	//读取到最终的缓冲区中
// 	//	chunk = append(chunk, buf[:n]...)
// 	//}
// 	//方法3⃣️：
// 	//r := bufio.NewReader(object)
// 	//var chunk []byte
// 	//
// 	//buf := make([]byte, 1024)
// 	//
// 	//for {
// 	//	n, err := r.Read(buf)
// 	//	if err != nil && err != io.EOF {
// 	//		panic(err)
// 	//	}
// 	//	if 0 == n {
// 	//		break
// 	//	}
// 	//	//fmt.Println(string(buf))
// 	//	chunk = append(chunk, buf...)
// 	//}
// 	//方法4⃣️：
// 	chunk, err := ioutil.ReadAll(object)
// 	if err != nil {
// 		fmt.Println("read to fd fail", err)
// 		return err, nil
// 	}
// 	return nil, chunk

// http://127.0.0.1:53202/api/v1/buckets/engineercms/objects/download?prefix=U25hcDI1LmpwZw==
// http://127.0.0.1:53202/api/v1/buckets/engineercms/objects/download?preview=true&prefix=U25hcDI1LmpwZw==

// http://127.0.0.1:9000/engineercms/Snap25.jpg?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=3ZKBL4HG9RH3SRTOLR1V%2F20230108%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20230108T015415Z&X-Amz-Expires=604800&X-Amz-Security-Token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NLZXkiOiIzWktCTDRIRzlSSDNTUlRPTFIxViIsImV4cCI6MTY3MzE0NjgzNCwicGFyZW50IjoibWluaW9hZG1pbiJ9.jklS00Uh3bSapNXlCJVkCp4hl67xRwVcVn4SzwCA9cjlPJ7yHbW_eSQ7vpK_hmdwXYxgZyJfRSTa8loEjpYx5Q&X-Amz-SignedHeaders=host&versionId=null&X-Amz-Signature=e358e6278bcaf0fb2f57d4e9b725f553946570772ed49b3e93e756df63b7bc4a

// @Title post wx photo img
// @Description post photo
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 photo not found
// @router /uploadphotodataminio [post]
// 上传图片到minios
// func (c *PhotoController) UploadPhotoData() {
// 	//获取上传的文件
// 	files, err := c.GetFiles("input-ke-2[]")
// 	if err != nil {
// 		logs.Error(err)
// 		c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "title": "", "original": "", "data": "获取上传文件错误！"}
// 		c.ServeJSON()
// 		return
// 	}

// 	useminio, err := web.AppConfig.String("useminio")
// 	if err != nil {
// 		logs.Error(err)
// 	}

// 	minio_bucketname, err := web.AppConfig.String("minio_bucketname")
// 	if err != nil {
// 		logs.Error(err)
// 	}

// 	// GetFiles return multi-upload files
// 	for i, _ := range files {
// 		//for each fileheader, get a handle to the actual file
// 		file, err := files[i].Open()
// 		defer file.Close()
// 		if err != nil {
// 			logs.Error(err)
// 			return
// 		}

// 		uploadInfo, err := minioClient.PutObject(context.Background(), minio_bucketname, files[i].Filename, file, -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 		// fmt.Println("Successfully uploaded bytes: ", uploadInfo)
// 		fmt.Println("etag: ", uploadInfo.ETag)
// 	}

// 	c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": "", "title": "111", "original": "demo.jpg"}
// 	c.ServeJSON()
// 	// file2, ok := files.(*os.File)

// 	// FileHeader.Open()

// 	// _, header , _ := c.Request.FormFile("upload")
// 	// out, _ := os.Open(uploadphoto.Filename)

// 	// files, fileHeader, _ := c.Request.FormFile("file")
// 	// byteData := make([]byte, h.Size)
// 	// files.Read(byteData)

// 	// file, _, err := c.Ctx.Request.FormFile("input-ke-2[]")
// 	// if err != nil {
// 	// 	logs.Error(err)
// 	// }
// 	// f, err := os.OpenFile(tofile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
// 	// if err != nil {
// 	// 	logs.Error(err)
// 	// }
// 	// defer f.Close()

// 	// file2, err := os.Open(h.Filename)
// 	// if err != nil {
// 	// 	logs.Error(err)
// 	// 	return
// 	// }
// }

// Upload an object.
// _, err = c.PutObject(context.Background(), "bucket", "object", bytes.NewBuffer([]byte("abcd")), 3, minio.PutObjectOptions{})

// b, err := ioutil.ReadFile("TEST_FILE")
// func ReadFile(filename string) ([]byte, error)
// if err != nil {
//   panic(err)
// }
// in := bytes.NewReader(b)
// name := fmt.Sprintf("TEST_%d", x)
// _, err := minioClient.PutObject(bucketName, name, in, "application/octet-stream")

// content := bytes.NewReader([]byte("Hello again"))
// _, err = minioClient.PutObject("testbucket", "my-encrypted-object.txt", content, 11, minio.PutObjectOptions{UserMetadata: metadata})
// func NewReader(b []byte) *Reader

// r := strings.NewReader("foo")
// _, err := minioClient.PutObject("mybukkit", "myfile/file1", r, 3, minio.PutObjectOptions{})
// minioClient.putObject('bucket', fileName, file.buffer, (err, etag) => {
//   if (err) throw err
//   return res.json(fileName)
// })

// for p := range files {
//   fInfo, err := os.Stat(p)
//   if err != nil {
//   	fmt.Printf("failed to stat file, path %s, %v\n", p, err)
//   	return
//   }
//   f, err := os.Open(p)
//   if err != nil {
//   	fmt.Printf("failed to open file, path %s, %v\n", p, err)
//   	return
//   }
//   n, err := mc.PutObject(
//   	bName,
//   	toUpload+strconv.Itoa(i),
//   	f,
//   	fInfo.Size(),
//   	minio.PutObjectOptions{},
//   )
