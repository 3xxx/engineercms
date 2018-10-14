package controllers

import (
	// "bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/pborman/uuid"
	// "image/png"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	// "hydrocms/models"
	"encoding/base64"
	// "engineercms/models"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// CMSWX froala API
type FroalaController struct {
	beego.Controller
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

//下面这个没用
func (c *FroalaController) ControllerFroala() {
	op := c.Input().Get("action")
	key := c.Input().Get("key") //这里进行判断各个页面，如果是addtopic，如果是addcategory
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
				beego.Error(err)
			}
			var filesize int64
			fileSuffix := path.Ext(h.Filename)
			newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
			year, month, _ := time.Now().Date()
			err = os.MkdirAll("./attachment/wiki/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
			if err != nil {
				beego.Error(err)
			}
			path1 := "./attachment/wiki/" + strconv.Itoa(year) + month.String() + "/" + newname //h.Filename
			Url := "/attachment/wiki/" + strconv.Itoa(year) + month.String() + "/"
			err = c.SaveToFile("upfile", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
			if err != nil {
				beego.Error(err)
			}
			filesize, _ = FileSize(path1)
			filesize = filesize / 1000.0
			c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": Url + newname, "title": h.Filename, "original": h.Filename}
			c.ServeJSON()
		default:
			//解析表单
			pid := c.Input().Get("pid")
			// beego.Info(pid)
			//pid转成64为
			pidNum, err := strconv.ParseInt(pid, 10, 64)
			if err != nil {
				beego.Error(err)
			}
			//根据proj的parentIdpath
			Url, DiskDirectory, err := GetUrlPath(pidNum)
			if err != nil {
				beego.Error(err)
			}
			// beego.Info(DiskDirectory)
			//获取上传的文件
			_, h, err := c.GetFile("upfile")
			if err != nil {
				beego.Error(err)
			}
			beego.Info(h.Filename)
			fileSuffix := path.Ext(h.Filename)
			// random_name
			newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
			// err = ioutil.WriteFile(path1+newname+".jpg", ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
			// if err != nil {
			// 	beego.Error(err)
			// }
			year, month, _ := time.Now().Date()
			err = os.MkdirAll(DiskDirectory+"/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
			if err != nil {
				beego.Error(err)
			}
			var path string
			var filesize int64
			if h != nil {
				//保存附件
				path = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + newname
				Url = "/" + Url + "/" + strconv.Itoa(year) + month.String() + "/"
				err = c.SaveToFile("upfile", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
				if err != nil {
					beego.Error(err)
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
		number := c.Input().Get("number")

		name := c.Input().Get("name")
		err := os.MkdirAll("./attachment/"+number+name, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
		if err != nil {
			beego.Error(err)
		}
		path1 := "./attachment/" + number + name + "/"
		//保存上传的图片
		//upfile为base64格式文件，转成图片保存
		ww := c.Input().Get("upfile")
		ddd, _ := base64.StdEncoding.DecodeString(ww)           //成图片文件并把文件写入到buffer
		newname := strconv.FormatInt(time.Now().Unix(), 10)     // + "_" + filename
		err = ioutil.WriteFile(path1+newname+".jpg", ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
		if err != nil {
			beego.Error(err)
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

//这个没用
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

//添加文章里的图片上传
func (c *FroalaController) UploadImg() {
	//解析表单
	pid := c.Input().Get("pid")
	// beego.Info(pid)
	//pid转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据proj的parentIdpath
	Url, DiskDirectory, err := GetUrlPath(pidNum)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(DiskDirectory)
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(h.Filename)
	fileSuffix := path.Ext(h.Filename)
	// random_name
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	// err = ioutil.WriteFile(path1+newname+".jpg", ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
	// if err != nil {
	// 	beego.Error(err)
	// }
	year, month, _ := time.Now().Date()
	err = os.MkdirAll(DiskDirectory+"/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		beego.Error(err)
	}
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		path = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + newname
		Url = "/" + Url + "/" + strconv.Itoa(year) + month.String() + "/"
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
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
//微信wx添加文章里的图片上传
func (c *FroalaController) UploadWxImg() {
	//解析表单
	pid := beego.AppConfig.String("wxcatalogid") //"26159" //c.Input().Get("pid")
	//pid转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据proj的parentIdpath
	Url, DiskDirectory, err := GetUrlPath(pidNum)
	if err != nil {
		beego.Error(err)
	}
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	// random_name
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	year, month, _ := time.Now().Date()
	err = os.MkdirAll(DiskDirectory+"/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		beego.Error(err)
	}
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		path = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + newname
		Url = "/" + Url + "/" + strconv.Itoa(year) + month.String() + "/"
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
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
// @Param id query string  true "The id of project"
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /uploadwximgs/:id [post]
//微信wx添加文章里的图片上传
func (c *FroalaController) UploadWxImgs() {
	//解析表单
	pid := c.Ctx.Input.Param(":id")
	// pid := beego.AppConfig.String("wxcatalogid") //"26159" //c.Input().Get("pid")
	//pid转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据proj的parentIdpath
	Url, DiskDirectory, err := GetUrlPath(pidNum)
	if err != nil {
		beego.Error(err)
	}
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	// random_name
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	year, month, _ := time.Now().Date()
	err = os.MkdirAll(DiskDirectory+"/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		beego.Error(err)
	}
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		path = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + newname
		Url = "/" + Url + "/" + strconv.Itoa(year) + month.String() + "/"
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
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

//添加wiki里的图片上传
func (c *FroalaController) UploadWikiImg() {
	//保存上传的图片
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	var filesize int64
	fileSuffix := path.Ext(h.Filename)
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	year, month, _ := time.Now().Date()
	err = os.MkdirAll("./attachment/wiki/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		beego.Error(err)
	}
	path1 := "./attachment/wiki/" + strconv.Itoa(year) + month.String() + "/" + newname //h.Filename
	Url := "/attachment/wiki/" + strconv.Itoa(year) + month.String() + "/"
	err = c.SaveToFile("file", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
	if err != nil {
		beego.Error(err)
	}
	filesize, _ = FileSize(path1)
	filesize = filesize / 1000.0
	c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "link": Url + newname, "title": "111", "original": "demo.jpg"}
	c.ServeJSON()
}

//添加文章里的视频上传
func (c *FroalaController) UploadVideo() {
	//解析表单
	pid := c.Input().Get("pid")
	// beego.Info(pid)
	//pid转成64为
	pidNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据proj的parentIdpath
	Url, DiskDirectory, err := GetUrlPath(pidNum)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(DiskDirectory)
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(h.Filename)
	fileSuffix := path.Ext(h.Filename)
	// random_name
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	// err = ioutil.WriteFile(path1+newname+".jpg", ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
	// if err != nil {
	// 	beego.Error(err)
	// }
	year, month, _ := time.Now().Date()
	err = os.MkdirAll(DiskDirectory+"/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		beego.Error(err)
	}
	var path string
	var filesize int64
	if h != nil {
		//保存附件
		path = DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + newname
		Url = "/" + Url + "/" + strconv.Itoa(year) + month.String() + "/"
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
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
