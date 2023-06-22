package controllers

import (
	// "bytes"
	"encoding/json"
	"fmt"
	// beego "github.com/beego/beego/v2/adapter"
	"github.com/pborman/uuid"
	"image"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	// "hydrocms/models"
	"encoding/base64"
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/nfnt/resize"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	// "path/filepath"
	"bytes"
	"github.com/3xxx/engineercms/controllers/utils"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// CMSWX froala API
type FroalaController struct {
	web.Controller
}

type UploadimgFroala struct {
	url      string
	title    string
	original string
	state    string
	// "url": fmt.Sprintf("/static/upload/%s", filename),
	// "title": "demo.jpg",
	// "original": header.Filename,
	// "state": "SUCCESS"
}

// 下面这个没用
func (c *FroalaController) ControllerFroala() {
	op := c.GetString("action")
	key := c.GetString("key") //这里进行判断各个页面，如果是addtopic，如果是addcategory
	switch op {
	case "config": //这里还是要优化成conf/config.json
		// $CONFIG = json_decode(preg_replace("/\/\*[\s\S]+?\*\//", "", file_get_contents("config.json")), true);
		// $action = $_GET['action'];
		// switch ($action) {
		//     case 'config':
		//         $result =  json_encode($CONFIG);
		// var configJson []byte // 当客户端请求/controller?action=config 返回的json内容
		file, err := os.Open("conf/config.json")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer file.Close()
		// fi, err := os.Open("d:/config.json")
		// if err != nil {
		// 	panic(err)
		// }
		// defer fi.Close()
		fd, err := ioutil.ReadAll(file)
		src := string(fd)
		re, _ := regexp.Compile("\\/\\*[\\S\\s]+?\\*\\/") //参考php的$CONFIG = json_decode(preg_replace("/\/\*[\s\S]+?\*\//", "", file_get_contents("config.json")), true);
		//将php中的正则移植到go中，需要将/ \/\*[\s\S]+?\*\/  /去掉前后的/，然后将\改成2个\\
		//参考//去除所有尖括号内的HTML代码，并换成换行符
		// re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
		// src = re.ReplaceAllString(src, "\n")
		//当把<和>换成/*和*\时，斜杠/和*之间加双斜杠\\才行。
		src = re.ReplaceAllString(src, "")
		tt := []byte(src)
		// buf := bytes.NewBuffer(nil)
		// buf.ReadFrom(file)
		// configJson = buf.Bytes()
		// w.Write(configJson)
		var r interface{}
		json.Unmarshal(tt, &r) //这个byte要解码
		c.Data["json"] = r
		c.ServeJSON()
		//下面这段是测试用的
		// b := []byte(`{
		//             "imageActionName": "uploadimage",
		//             "imageFieldName": "upfile",
		//             "imageMaxSize": 2048000,
		//             "imageAllowFiles": [".png", ".jpg", ".jpeg", ".gif", ".bmp"],
		//             "imageCompressEnable": true,
		//             "imageCompressBorder": 1600,
		//             "imageInsertAlign": "none",
		//             "imageUrlPrefix": "",
		//             "imagePathFormat": "/static/upload/{yyyy}{mm}{dd}/{time}{rand:6}"
		//       }`)
		// var r interface{}
		// json.Unmarshal(b, &r)
		// c.Data["json"] = r
		// c.ServeJSON()
	case "uploadimage", "uploadfile", "uploadvideo":
		switch key {
		case "wiki": //添加wiki
			//保存上传的图片
			_, h, err := c.GetFile("upfile")
			if err != nil {
				logs.Error(err)
			}
			var filesize int64
			fileSuffix := path.Ext(h.Filename)
			newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
			year, month, _ := time.Now().Date()
			err = os.MkdirAll("./attachment/wiki/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
			if err != nil {
				logs.Error(err)
			}
			path1 := "./attachment/wiki/" + strconv.Itoa(year) + month.String() + "/" + newname //h.Filename
			Url := "/attachment/wiki/" + strconv.Itoa(year) + month.String() + "/"
			err = c.SaveToFile("upfile", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
			if err != nil {
				logs.Error(err)
			}
			filesize, _ = FileSize(path1)
			filesize = filesize / 1000.0
			c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": Url + newname, "title": h.Filename, "original": h.Filename}
			c.ServeJSON()
		default:
			//解析表单
			pid := c.GetString("pid")
			// beego.Info(pid)
			//pid转成64为
			pidNum, err := strconv.ParseInt(pid, 10, 64)
			if err != nil {
				logs.Error(err)
			}
			//根据proj的parentIdpath
			Url, DiskDirectory, err := GetUrlPath(pidNum)
			if err != nil {
				logs.Error(err)
			}
			// beego.Info(DiskDirectory)
			//获取上传的文件
			_, h, err := c.GetFile("upfile")
			if err != nil {
				logs.Error(err)
			}
			// beego.Info(h.Filename)
			fileSuffix := path.Ext(h.Filename)
			// random_name
			newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
			// err = ioutil.WriteFile(path1+newname+".jpg", ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
			// if err != nil {
			// 	logs.Error(err)
			// }
			year, month, _ := time.Now().Date()
			err = os.MkdirAll(DiskDirectory+"/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
			if err != nil {
				logs.Error(err)
			}
			var path string
			var filesize int64
			if h != nil {
				//保存附件
				path = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + newname
				Url = "/" + Url + "/" + strconv.Itoa(year) + month.String() + "/"
				err = c.SaveToFile("upfile", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
				if err != nil {
					logs.Error(err)
				}
				filesize, _ = FileSize(path)
				filesize = filesize / 1000.0
				c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": Url + newname, "title": h.Filename, "original": h.Filename}
				c.ServeJSON()
			} else {
				c.Data["json"] = map[string]interface{}{"state": "ERROR", "url": "", "title": "", "original": ""}
				c.ServeJSON()
			}
		}
	case "uploadscrawl":
		number := c.GetString("number")

		name := c.GetString("name")
		err := os.MkdirAll("./attachment/"+number+name, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
		if err != nil {
			logs.Error(err)
		}
		path1 := "./attachment/" + number + name + "/"
		//保存上传的图片
		//upfile为base64格式文件，转成图片保存
		ww := c.GetString("upfile")
		ddd, _ := base64.StdEncoding.DecodeString(ww)           //成图片文件并把文件写入到buffer
		newname := strconv.FormatInt(time.Now().Unix(), 10)     // + "_" + filename
		err = ioutil.WriteFile(path1+newname+".jpg", ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
		if err != nil {
			logs.Error(err)
		}
		// var filesize int64
		// filesize, _ = FileSize(path1)
		// filesize = filesize / 1000.0
		c.Data["json"] = map[string]interface{}{
			"state":    "SUCCESS",
			"url":      "/attachment/" + number + name + "/" + newname + ".jpg",
			"title":    newname + ".jpg",
			"original": newname + ".jpg",
		}
		c.ServeJSON()
	case "listimage":
		type List struct {
			Url string `json:"url"`
			// Source string
			// State  string
		}
		type Listimage struct {
			State string `json:"state"` //这些第一个字母要大写，否则不出结果
			List  []List `json:"list"`
			Start int    `json:"start"`
			Total int    `json:"total"`
			// Name        string
			// Age         int
			// Slices      []string //slice
			// Mapstring   map[string]string
			// StructArray []List            //结构体的切片型
			// MapStruct   map[string][]List //map:key类型是string或struct，value类型是切片，切片的类型是string或struct
			//	Desks  List
		}
		// var m map[string]string = make(map[string]string)
		// m["Go"] = "No.1"
		// m["Java"] = "No.2"
		// m["C"] = "No.3"
		// fmt.Println(m)
		list := []List{
			{"/static/upload/1.jpg"},
			{"/static/upload/2.jpg"},
			// {"upload/1.jpg", "http://a.com/1.jpg", "SUCCESS"},
			// {"upload/2.jpg", "http://b.com/2.jpg", "SUCCESS"},
		}
		// var mm map[string][]List = make(map[string][]List)
		// mm["Go"] = list
		// mm["Java"] = list
		// fmt.Println(mm)
		listimage := Listimage{"SUCCESS", list, 1, 21}
		// beego.Info(listimage){SUCCESS [{/static/upload/1.jpg} {/static/upload/2.jpg}] 1 21}
		// fmt.Println(listimage)
		// b, _ := json.Marshal(listimage)
		// mystruct := { ... }
		// c.Data["jsonp"] = listimage
		// beego.Info(string(b)){"State":"SUCCESS","List":[{"Url":"/static/upload/1.jpg"},{"Url":"/static/upload/2.jpg"}],"Start":1,"Total":21}
		// c.ServeJSONP()
		c.Data["json"] = listimage
		c.ServeJSON()
		// c.Data["json"] = map[string]interface{}{"State":"SUCCESS","List":[{"Url":"/static/upload/1.jpg"},{"Url":"/static/upload/2.jpg"}],"Start":1,"Total":21}
		// 需要支持callback参数,返回jsonp格式
		// {
		//     "state": "SUCCESS",
		//     "list": [{
		//         "url": "upload/1.jpg"
		//     }, {
		//         "url": "upload/2.jpg"
		//     }, ],
		//     "start": 20,
		//     "total": 100
		// }
	case "catchimage":
		type List struct {
			Url    string `json:"url"`
			Source string `json:"source"`
			State  string `json:"state"`
		}
		type Catchimage struct {
			State string `json:"state"` //这些第一个字母要大写，否则不出结果
			List  []List `json:"list"`
			// Start int
			// Total int
			// Name        string
			// Age         int
			// Slices      []string //slice
			// Mapstring   map[string]string
			// StructArray []List            //结构体的切片型
			// MapStruct   map[string][]List //map:key类型是string或struct，value类型是切片，切片的类型是string或struct
			//	Desks  List
		}
		// var m map[string]string = make(map[string]string)
		// m["Go"] = "No.1"
		// m["Java"] = "No.2"
		// m["C"] = "No.3"
		// fmt.Println(m)
		list := []List{
			{"/static/upload/1.jpg", "https://pic2.zhimg.com/7c4a389acaa008a6d1fe5a0083c86975_b.png", "SUCCESS"},
			{"/static/upload/2.jpg", "https://pic2.zhimg.com/7c4a389acaa008a6d1fe5a0083c86975_b.png", "SUCCESS"},
			// {"upload/1.jpg", "http://a.com/1.jpg", "SUCCESS"},
			// {"upload/2.jpg", "http://b.com/2.jpg", "SUCCESS"},
		}
		// var mm map[string][]List = make(map[string][]List)
		// mm["Go"] = list
		// mm["Java"] = list
		// fmt.Println(mm)
		catchimage := Catchimage{"SUCCESS", list}
		// beego.Info(catchimage){SUCCESS [{/static/upload/1.jpg} {/static/upload/2.jpg}] 1 21}
		// fmt.Println(catchimage)
		// b, _ := json.Marshal(catchimage)
		// mystruct := { ... }
		// c.Data["jsonp"] = catchimage
		// beego.Info(string(b)){"State":"SUCCESS","List":[{"Url":"/static/upload/1.jpg"},{"Url":"/static/upload/2.jpg"}],"Start":1,"Total":21}
		// c.ServeJSONP()
		c.Data["json"] = catchimage
		c.ServeJSON()

		file, header, err := c.GetFile("source") // r.FormFile("upfile")
		// beego.Info(header.Filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		filename := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(header.Filename)
		err = os.MkdirAll(path.Join("static", "upload"), 0775)
		if err != nil {
			panic(err)
		}
		outFile, err := os.Create(path.Join("static", "upload", filename))
		if err != nil {
			panic(err)
		}
		defer outFile.Close()
		io.Copy(outFile, file)
	}
}

// 这个没用
func UploadImg(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("upfile")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	filename := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(header.Filename)
	err = os.MkdirAll(path.Join("static", "upload"), 0775)
	if err != nil {
		panic(err)
	}
	outFile, err := os.Create(path.Join("static", "upload", filename))
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	io.Copy(outFile, file)
	b, err := json.Marshal(map[string]string{
		"url":      fmt.Sprintf("/static/upload/%s", filename), //保存后的文件路径
		"title":    "",                                         //文件描述，对图片来说在前端会添加到title属性上
		"original": header.Filename,                            //原始文件名
		"state":    "SUCCESS",                                  //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
	w.Write(b)
}

// 添加文章里的图片上传
func (c *FroalaController) UploadImg() {
	// 登录
	_, _, _, _, islogin := checkprodRole(c.Ctx)
	if !islogin {
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "data": "用户未登录！"}
		c.ServeJSON()
	}
	//解析表单
	pid := c.GetString("pid")
	// beego.Info(pid)
	//pid转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	//根据proj的parentIdpath
	Url, DiskDirectory, err := GetUrlPath(pidNum)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(DiskDirectory)
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(h.Filename)
	fileSuffix := path.Ext(h.Filename)
	// random_name
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	// err = ioutil.WriteFile(path1+newname+".jpg", ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
	// if err != nil {
	// 	logs.Error(err)
	// }
	year, month, _ := time.Now().Date()
	err = os.MkdirAll(DiskDirectory+"/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		logs.Error(err)
	}
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		path = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + newname
		Url = "/" + Url + "/" + strconv.Itoa(year) + month.String() + "/"
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			logs.Error(err)
		}
		filesize, _ = FileSize(path)
		filesize = filesize / 1000.0
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": Url + newname, "title": "111", "original": "demo.jpg"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "title": "", "original": ""}
		c.ServeJSON()
	}
}

// @Title post wx artile img by catalogId
// @Description post article img by catalogid
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /uploadwximg [post]
// 微信wx添加文章里的图片上传_独立上传图片模式
func (c *FroalaController) UploadWxImg() {
	//解析表单
	pid, err := web.AppConfig.String("wxcatalogid") //"26159" //c.GetString("pid")
	if err != nil {
		logs.Error(err)
	}
	//pid转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	//根据proj的parentIdpath
	Url, DiskDirectory, err := GetUrlPath(pidNum)
	if err != nil {
		logs.Error(err)
	}
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		logs.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	// random_name
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	year, month, _ := time.Now().Date()
	err = os.MkdirAll(DiskDirectory+"/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		logs.Error(err)
	}
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		path = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + newname
		Url = "/" + Url + "/" + strconv.Itoa(year) + month.String() + "/"
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			logs.Error(err)
		}
		filesize, _ = FileSize(path)
		filesize = filesize / 1000.0
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": Url + newname, "title": "111", "original": "demo.jpg"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "title": "", "original": ""}
		c.ServeJSON()
	}
}

// @Title post wx artile img by catalogId
// @Description post article img by catalogid
// @Param projectid query string true "The projectid of wxeditor"
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /uploadwxeditorimg [post]
// 微信wx添加文章里的图片上传_小程序富文本里的上传图片
func (c *FroalaController) UploadWxEditorImg() {
	var ProjectId int64
	var err error
	projectid := c.GetString("projectid")
	if projectid != "" {
		ProjectId, err = strconv.ParseInt(projectid, 10, 64)
		if err != nil {
			logs.Error(err)
		}
	}
	//解析表单
	// pid := web.AppConfig.String("wxcatalogid") //"26159" //c.GetString("pid")
	// //pid转成64为
	// pidNum, err := strconv.ParseInt(pid, 10, 64)
	// if err != nil {
	// 	logs.Error(err)
	// }
	//根据proj的parentIdpath
	Url, DiskDirectory, err := GetUrlPath(ProjectId)
	if err != nil {
		logs.Error(err)
	}
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		logs.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	// random_name
	nanoname := strconv.FormatInt(time.Now().UnixNano(), 10)
	newname := nanoname + fileSuffix // + "_" + filename
	small_newname := nanoname + "_small" + fileSuffix
	year, month, _ := time.Now().Date()
	err = os.MkdirAll(DiskDirectory+"/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		logs.Error(err)
	}
	var imagepath, new_imagepath string
	var filesize int64
	if h != nil {
		//保存附件
		imagepath = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + newname
		new_imagepath = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + small_newname
		Url = "/" + Url + "/" + strconv.Itoa(year) + month.String() + "/"
		err = c.SaveToFile("file", imagepath) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			logs.Error(err)
		}
		filesize, _ = FileSize(imagepath)
		filesize = filesize / 1000.0

		//*****压缩图片***
		file, err := os.Open(imagepath)
		if err != nil {
			// log.Fatal(err)
			logs.Error(err)
		}
		defer file.Close()
		var img image.Image
		var typeImage int
		// ext := filepath.Ext(imagepath)
		if strings.EqualFold(fileSuffix, ".jpg") || strings.EqualFold(fileSuffix, ".jpeg") {
			img, err = jpeg.Decode(file)
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

		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": Url + small_newname, "title": "111", "original": "demo.jpg"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "title": "", "original": ""}
		c.ServeJSON()
	}
}

// @Title post wx artile img by catalogId
// @Description post article img by catalogid
// @Param id path string  true "The id of project"
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /uploadwximgs/:id [post]
// 微信wx添加文章里的图片上传——这个鲁班宝，但这个id更通用
func (c *FroalaController) UploadWxImgs() {
	//解析表单
	pid := c.Ctx.Input.Param(":id")
	// pid := web.AppConfig.String("wxcatalogid") //"26159" //c.GetString("pid")
	//pid转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	//根据proj的parentIdpath
	Url, DiskDirectory, err := GetUrlPath(pidNum)
	if err != nil {
		logs.Error(err)
	}
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		logs.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	// random_name
	nanoname := strconv.FormatInt(time.Now().UnixNano(), 10)
	newname := nanoname + fileSuffix // + "_" + filename
	small_newname := nanoname + "_small" + fileSuffix
	year, month, _ := time.Now().Date()
	err = os.MkdirAll(DiskDirectory+"/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		logs.Error(err)
	}
	var imagepath, new_imagepath string
	var filesize int64
	if h != nil {
		//保存附件
		imagepath = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + newname
		new_imagepath = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + small_newname
		Url = "/" + Url + "/" + strconv.Itoa(year) + month.String() + "/"
		err = c.SaveToFile("file", imagepath) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			logs.Error(err)
		}
		filesize, _ = FileSize(imagepath)
		filesize = filesize / 1000.0

		//*****压缩图片***
		file, err := os.Open(imagepath)
		if err != nil {
			// log.Fatal(err)
			logs.Error(err)
		}
		defer file.Close()
		var img image.Image
		var typeImage int
		// ext := filepath.Ext(imagepath)
		if strings.EqualFold(fileSuffix, ".jpg") || strings.EqualFold(fileSuffix, ".jpeg") {
			img, err = jpeg.Decode(file)
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
		// 用https://github.com/gographics/imagick  进行裁剪图片 https://www.cnblogs.com/Torrance/p/8877039.html
		// resize to width 1000 using Lanczos resampling
		// and preserve aspect ratio
		// m := resize.Resize(1000, 0, img, resize.Lanczos3)
		m := resize.Resize(500, 0, img, resize.Lanczos3)
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
		// 开始图片鉴别
		app_version := c.GetString("app_version")
		accessToken, _, _, err := utils.GetAccessToken(app_version)
		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": err}
			c.ServeJSON()
			return
		}

		// 把上传文件转为byte
		// file, header, err := ctx.Request.FormFile("file")
		// defer file.Close()
		// if err != nil {
		//     return nil, err
		// }

		// buf := bytes.NewBuffer(nil)
		// if _, err := io.Copy(buf, file); err != nil {
		//     return nil, err
		// }
		// fmt.Println(buf.Bytes())

		// 将os.File输出转换为String
		var buf bytes.Buffer
		io.Copy(&buf, out)
		// asString := string(buf.Bytes())

		errcode, errmsg, err := utils.ImgSecCheck(buf.Bytes(), accessToken)
		// beego.Info(errcode)
		// beego.Info(errmsg)
		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"info": err, "state": "ERROR", "data": errmsg}
			c.ServeJSON()
		} else if errcode != 87014 {
			c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "state": "SUCCESS", "link": Url + small_newname, "title": small_newname, "original": "demo.jpg"}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "文件为空！"}
		c.ServeJSON()
	}
}

// @Title post wx user avatar
// @Description post user avatar
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /uploadwxavatar [post]
// 微信wx添加用户头像上传
func (c *FroalaController) UploadWxAvatar() {
	var user models.User
	var err error
	openID := c.GetSession("openID")
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			logs.Error(err)
		}
	}
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		logs.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	// random_name
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	err = os.MkdirAll("./static/avatar/", 0777)                          //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		logs.Error(err)
	}
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		path = "./static/avatar/" + newname
		Url := "/static/avatar/"
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"state": "ERROR", "photo": "", "title": "", "original": ""}
			c.ServeJSON()
		} else {
			filesize, _ = FileSize(path)
			filesize = filesize / 1000.0
			_, err = models.AddUserAvator(user.Id, Url+newname)
			if err != nil {
				logs.Error(err)
			}
			wxsite, err := web.AppConfig.String("wxreqeustsite")
			if err != nil {
				logs.Error(err)
			}
			// c.Data["json"] = map[string]interface{}{"errNo": 1, "msg": "success", "photo": wxsite + Url + newname, "title": newname, "original": newname}
			// c.ServeJSON()
			c.Ctx.WriteString(wxsite + Url + newname)
		}
	} else {
		c.Data["json"] = map[string]interface{}{"errNo": 0, "state": "ERROR", "photo": "", "title": "", "original": ""}
		c.ServeJSON()
	}
}

// @Title post wx user avatar
// @Description post user avatar
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /uploadappreciationphoto [post]
// 小程序wx添加用户赞赏码上传
func (c *FroalaController) UploadAppreciationPhoto() {
	var user models.User
	var err error
	openID := c.GetSession("openID")
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			logs.Error(err)
		}
	}
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		logs.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	// random_name
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	err = os.MkdirAll("./static/appreciation/", 0777)                    //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		logs.Error(err)
	}
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		path = "./static/appreciation/" + newname
		Url := "/static/appreciation/"
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"state": "ERROR", "photo": "", "title": "", "original": ""}
			c.ServeJSON()
		} else {
			filesize, _ = FileSize(path)
			filesize = filesize / 1000.0
			_, err = models.AddUserAppreciation(user.Id, Url+newname)
			if err != nil {
				logs.Error(err)
			}
			wxsite, err := web.AppConfig.String("wxreqeustsite")
			if err != nil {
				logs.Error(err)
			}
			// c.Data["json"] = map[string]interface{}{"errNo": 1, "msg": "success", "photo": wxsite + Url + newname, "title": newname, "original": newname}
			// c.ServeJSON()
			c.Ctx.WriteString(wxsite + Url + newname)
		}
	} else {
		c.Data["json"] = map[string]interface{}{"errNo": 0, "state": "ERROR", "photo": "", "title": "", "original": ""}
		c.ServeJSON()
	}
}

// @Title post wx video by catalogId
// @Description post video by catalogid
// @Param id path string true "The id of project"
// @Param desc query string true "The descript of video"
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 article not found
// @router /uploadwxvideo/:id [post]
// 微信wx添加视频
func (c *FroalaController) UploadWxVideo() {
	var user models.User
	var err error
	openID := c.GetSession("openID")
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			logs.Error(err)
		}
	} else {
		// c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		// c.ServeJSON()
		// return //本地调试的时候，由于小程序无法登陆服务器，所以这里要注释掉。
	}
	c.Ctx.Request.Body = http.MaxBytesReader(c.Ctx.ResponseWriter, c.Ctx.Request.Body, MAX_UPLOAD_SIZE)
	if err := c.Ctx.Request.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		// sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}
	//解析表单
	content := c.GetString("desc")

	pid := c.Ctx.Input.Param(":id")
	// pid := web.AppConfig.String("wxcatalogid") //"26159" //c.GetString("pid")
	//pid转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	//根据proj的parentIdpath
	Url, DiskDirectory, err := GetUrlPath(pidNum)
	if err != nil {
		logs.Error(err)
	}
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		logs.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	// random_name
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	year, month, _ := time.Now().Date()
	err = os.MkdirAll(DiskDirectory+"/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		logs.Error(err)
	}
	var path string
	// var filesize int64
	if h != nil {
		//保存附件
		path = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + newname
		Url = "/" + Url + "/" + strconv.Itoa(year) + month.String() + "/"
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			logs.Error(err)
			return
		}
		//根据pid查出项目id
		proj, err := models.GetProj(pidNum)
		if err != nil {
			logs.Error(err)
		}
		//写入数据表
		vid, err := models.CreateVideo(proj.Id, user.Id, content, newname, Url+newname)
		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"state": "ERROR", "info": "ERR", "id": vid}
			c.ServeJSON()
		} else {
			c.Data["json"] = vid
			// c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "info": "SUCCESS", "link": Url + newname, "id": vid}
			c.ServeJSON()
		}
		// filesize, _ = FileSize(path)
		// filesize = filesize / 1000.0
	}
}

// @Title post wx videoCover by videoid
// @Description post videoCover by videoid
// @Param id query string true "The id of video"
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 article not found
// @router /uploadwxvideocover/:id [post]
// 微信wx添加视频封面
func (c *FroalaController) UploadWxVideoCover() {
	//解析表单
	vid := c.Ctx.Input.Param(":id")
	// pidNum, err := strconv.ParseInt(id, 10, 64)
	// if err != nil {
	// 	logs.Error(err)
	// }
	//根据proj的parentIdpath
	// Url, DiskDirectory, err := GetUrlPath(pidNum)
	// if err != nil {
	// 	logs.Error(err)
	// }

	// vid := c.GetString("id")
	//pid转成64为
	vidNum, err := strconv.ParseInt(vid, 10, 64)
	if err != nil {
		logs.Error(err)
	}

	video, err := models.GetVideobyId(vidNum)
	if err != nil {
		logs.Error(err)
	}

	fileSuffix := path.Ext(video.Name)
	// var filenameOnly string
	filenameOnly := strings.TrimSuffix(video.Name, fileSuffix)
	viedeocoverpath := "." + path.Dir(video.Url) + "/"

	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		logs.Error(err)
	}
	if h != nil {
		//保存附件
		// path = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + newname
		// Url = "/" + Url + "/" + strconv.Itoa(year) + month.String() + "/"
		err = c.SaveToFile("file", viedeocoverpath+filenameOnly+".jpg") //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			logs.Error(err)
			return
		}
		//更新数据表
		err := models.UpdateVideo(video.Id, path.Dir(video.Url)+"/"+filenameOnly+".jpg")
		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"state": "ERROR", "info": "ERR", "id": vid}
			c.ServeJSON()
		} else {
			c.Data["json"] = "SUCCESS"
			// c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "info": "SUCCESS", "link": Url + newname, "id": vid}
			c.ServeJSON()
		}
	}
}

// 添加wiki里的图片上传
func (c *FroalaController) UploadWikiImg() {
	//保存上传的图片
	_, h, err := c.GetFile("file")
	if err != nil {
		logs.Error(err)
	}
	// var filesize int64
	fileSuffix := path.Ext(h.Filename)
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	year, month, _ := time.Now().Date()
	err = os.MkdirAll("./attachment/wiki/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		logs.Error(err)
	}
	path1 := "./attachment/wiki/" + strconv.Itoa(year) + month.String() + "/" + newname //h.Filename
	Url := "/attachment/wiki/" + strconv.Itoa(year) + month.String() + "/"
	err = c.SaveToFile("file", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
	if err != nil {
		logs.Error(err)
	}
	// filesize, _ = FileSize(path1)
	// filesize = filesize / 1000.0
	c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": Url + newname, "title": "111", "original": "demo.jpg"}
	c.ServeJSON()
}

// 添加文章里的视频上传
func (c *FroalaController) UploadVideo() {
	//解析表单
	pid := c.GetString("pid")
	// beego.Info(pid)
	//pid转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	//根据proj的parentIdpath
	Url, DiskDirectory, err := GetUrlPath(pidNum)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(DiskDirectory)
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(h.Filename)
	fileSuffix := path.Ext(h.Filename)
	// random_name
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	// err = ioutil.WriteFile(path1+newname+".jpg", ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
	// if err != nil {
	// 	logs.Error(err)
	// }
	year, month, _ := time.Now().Date()
	err = os.MkdirAll(DiskDirectory+"/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		logs.Error(err)
	}
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		path = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + newname
		Url = "/" + Url + "/" + strconv.Itoa(year) + month.String() + "/"
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			logs.Error(err)
		}
		filesize, _ = FileSize(path)
		filesize = filesize / 1000.0
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": Url + newname, "title": "111", "original": "demo.jpg"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "title": "", "original": ""}
		c.ServeJSON()
	}
}

//下面这个保留
// func (c *FroalaController) UploadImg() { //对应这个路由 beego.Router("/controller", &controllers.FroalaController{}, "post:UploadImage")
// 	file, header, err := c.GetFile("file") // r.FormFile("upfile")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()
// 	filename := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(header.Filename)
// 	err = os.MkdirAll(path.Join("static", "upload"), 0775)
// 	if err != nil {
// 		panic(err)
// 	}
// 	outFile, err := os.Create(path.Join("static", "upload", filename))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer outFile.Close()
// 	io.Copy(outFile, file)
// 	c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": "/static/upload/" + filename, "title": "111", "original": "demo.jpg"}
// 	c.ServeJSON()
// }
