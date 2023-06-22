package controllers

import (
	// "fmt"
	// // beego "github.com/beego/beego/v2/adapter"
	// ole "github.com/go-ole/go-ole"
	// "github.com/go-ole/go-ole/oleutil"
	// "log"
	// "math"
	"os"
	// "os/exec"
	"path"
	"strconv"
	// "time"
	// // "hydrocms/models"
	// // "encoding/base64"
	"encoding/json"
	// "errors"
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"regexp"
	// // "github.com/PuerkitoBio/goquery"
	// "github.com/gorilla/websocket"
	// // "io"
	// "github.com/dlclark/regexp2"
	// "net/http"
	// "regexp"
	"strings"
	// "sync"
	// "bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	// "io"
)

type FreeCADController struct {
	web.Controller
}

type FreeCADList struct {
	Id         uint           `json:"id"`
	Number     string         `json:"number"` //分级目录代码
	Tags       [1]string      `json:"tags"`   //显示员工数量，如果定义为数值[1]int，则无论如何都显示0，所以要做成字符
	LazyLoad   bool           `json:"lazyLoad"`
	Text       string         `json:"text"`
	Selectable bool           `json:"selectable"`
	State      State2         `json:"state"`
	Nodes      []*FreeCADList `json:"nodes"`
}

type State2 struct {
	Checked  bool `json:"checked"`
	Disabled bool `json:"disabled"`
	Expanded bool `json:"expanded"`
	Selected bool `json:"selected"`
}

type FreeCADMenu struct {
	Id       uint           `json:"id"`
	Key      string         `json:"key"`  //
	Tags     [1]string      `json:"tags"` //显示数量
	LazyLoad bool           `json:"lazyLoad"`
	Title    string         `json:"title"`
	Children []*FreeCADMenu `json:"children"`
}

// type Nodes2 struct {
// 	Text string `json:"text"`
// }

// @Title get mathcad temples by keyword...
// @Description get mathcad temples by keyword...
// @Param keyword query string false "The keyword of mathtemples"
// @Param page query string false "The page of mathtemples"
// @Param limit query string false "The size of page"
// @Success 200 {object} models.UserTemple
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /freecad [get]
// 模型侧栏+模型展示页面
func (c *FreeCADController) FreeCAD() {
	var err error
	// keyword := c.GetString("keyword")

	// limit := c.GetString("limit")
	// limit1, err := strconv.Atoi(limit)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// page := c.GetString("pageNo")
	// page1, err := strconv.Atoi(page)
	// if err != nil {
	// 	logs.Error(err)
	// }

	// var offset int
	// if page1 <= 1 {
	// 	offset = 0
	// } else {
	// 	offset = (page1 - 1) * limit1
	// }

	// jsonData := "[{\"text\":\"Parent 1\",\"nodes\":[{\"text\":\"Child 1\"},{\"text\":\"Child 2\"}]},{\"text\":\"Parent 2\"}]"
	// str := []byte(jsonData)
	// str := []byte(`[{"id":1,"text":"0001-0100","selectable":false,"nodes":[{"id":2,"text":"FC0001带补强圈岔管","selectable":true},{"id":3,"text":"FC0002相贯线岔管","selectable":true}]},{"id":4,"text":"0101-0200","selectable":false}]`)
	// var data []FreeCADList
	// if err := json.Unmarshal(str, &data); err == nil {
	// 	logs.Info(data)
	// } else {
	// 	logs.Error(err)
	// }

	// type FreeCADList struct {
	// 	Id         uint          `json:"id"`
	// 	Number     string        `json:"number"` //分级目录代码
	// 	Tags       [1]string     `json:"tags"`   //显示员工数量，如果定义为数值[1]int，则无论如何都显示0，所以要做成字符
	// 	LazyLoad   bool          `json:"lazyLoad"`
	// 	Text       string        `json:"text"`
	// 	Selectable bool          `json:"selectable"`
	// 	State      State2        `json:"state"`
	// 	Nodes      []FreeCADList `json:"nodes"`
	// }
	var state State2
	var tags [1]string
	tags[0] = strconv.Itoa(100)
	//递归生成目录json
	root := FreeCADList{1, "0001-0100", tags, false, "0001-0100", false, state, []*FreeCADList{}}
	// makemathtreejson(cates, categories, products, &root)
	// c.Data["json"] = root
	freecadmodels, err := models.GetFreecadModels(0, 0, "")
	if err != nil {
		logs.Error(err)
	}
	makefreecadtreejson(freecadmodels, &root)

	c.Data["json"] = root //data
	// 初始展示模型，随机
	if len(freecadmodels) > 0 {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(freecadmodels))))
		number := n.String()             //转成string
		num, err := strconv.Atoi(number) //string转int
		if err != nil {
			logs.Error(err)
		}
		c.Data["ModelId"] = freecadmodels[num].ID
	} else {
		c.Data["ModelId"] = 1
	}

	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		c.TplName = "freecad/freecad.tpl"
	} else {
		c.TplName = "freecad/freecad.tpl"
	}
}

// @Title get mathcad temples by keyword...
// @Description get mathcad temples by keyword...
// @Param keyword query string false "The keyword of mathtemples"
// @Param page query string false "The page of mathtemples"
// @Param limit query string false "The size of page"
// @Success 200 {object} models.UserTemple
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /freecadmenu [get]
// 获取模板列表数据-antd tree用的数据结构
func (c *FreeCADController) FreeCADMenu() {
	// jsonData := "[{\"text\":\"Parent 1\",\"nodes\":[{\"text\":\"Child 1\"},{\"text\":\"Child 2\"}]},{\"text\":\"Parent 2\"}]"
	// str := []byte(jsonData)
	// str := []byte(`[{"id":1,"title":"0001-0100","key":"1","children":[{"id":2,"title":"FC0001带法兰岔管","key":"2"},{"id":3,"title":"FC0002相贯线岔管","key":"3"}]},{"id":4,"title":"0101-0200","key":"4"}]`)
	// var data []FreeCADMenu
	// if err := json.Unmarshal(str, &data); err == nil {
	// 	logs.Info(data)
	// } else {
	// 	logs.Error(err)
	// }

	var tags [1]string
	tags[0] = strconv.Itoa(100)
	//递归生成目录json
	root := FreeCADMenu{0, "0", tags, false, "0001-0100", []*FreeCADMenu{}}

	freecadmodels, err := models.GetFreecadModels(0, 0, "")
	if err != nil {
		logs.Error(err)
	}
	makefreecadmenujson(freecadmodels, &root)

	// 初始展示模型，随机
	if len(freecadmodels) > 0 {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(freecadmodels))))
		number := n.String()             //转成string
		num, err := strconv.Atoi(number) //string转int
		if err != nil {
			logs.Error(err)
		}
		c.Data["ModelId"] = freecadmodels[num].ID
	} else {
		c.Data["ModelId"] = 1
	}
	c.Data["json"] = root
	c.ServeJSON()
}

// @Title get mathcad temples by keyword...
// @Description get mathcad temples by keyword...
// @Param id path string true "The id of project"
// @Success 200 {object} models.UserTemple
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /getfreecad/:id [get]
// 模型展示子页面
func (c *FreeCADController) GetFreeCAD() {
	id := c.Ctx.Input.Param(":id")
	//id转成uint为
	idint, err := strconv.Atoi(id)
	if err != nil {
		logs.Error(err)
	}
	modelid := uint(idint)
	// 从路径中获取这个modelid，再传给页面中的ajax获取模型txt文件
	c.Data["ModelId"] = modelid

	freecadmodel, err := models.GetFreecadModel(modelid)
	if err != nil {
		logs.Info(err)
	}
	c.Data["Title"] = freecadmodel.TitleB

	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		c.TplName = "freecad/freecad_model.tpl"
	} else {
		c.TplName = "freecad/freecad_model.tpl"
	}
}

// @Title get mathcad temples by keyword...
// @Description get mathcad temples by keyword...
// @Param id path string true "The id of project"
// @Success 200 {object} models.UserTemple
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /freecaddata/:id [get]
// 获取模板列表数据
func (c *FreeCADController) FreeCADData() {
	var err error
	// file, err := os.Open("D:/gowork/src/github.com/3xxx/engineercms/attachment/freecad/001.txt")
	// if err != nil {
	// 	fmt.Println("文件打开失败 = ", err)
	// 	return
	// }
	// defer file.Close()              // 关闭文本流
	// reader := bufio.NewReader(file) // 读取文本数据
	// for {
	// 	str, err := reader.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	logs.Info(str)
	// }
	// logs.Info("文件读取结束")

	// f, err := os.Open("./attachment/freecad/001.txt")
	// if err != nil {
	// 	logs.Error(err)
	// 	return
	// }
	// logs.Info(f)
	// defer f.Close()

	// scanner := bufio.NewScanner(f)
	// logs.Info(scanner)
	// scanner.Split(bufio.ScanLines)

	// // 这是我们的缓冲区
	// var lines []string

	// for scanner.Scan() {
	// 	lines = append(lines, scanner.Text())
	// 	logs.Info(scanner.Text())
	// }

	// var line2 string
	// fmt.Println("read lines:")
	// for _, line := range lines {
	// 	fmt.Println(line)
	// 	line2 = line2 + line
	// }
	id := c.Ctx.Input.Param(":id")
	//id转成uint为
	idint, err := strconv.Atoi(id)
	if err != nil {
		logs.Error(err)
	}
	modelid := uint(idint)
	freecadmodel, err := models.GetFreecadModel(modelid)
	if err != nil {
		logs.Info(err)
	}

	// file, err := os.Open("./attachment/freecad/001.txt")
	file, err := os.Open(freecadmodel.Path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	bytesread, err := file.Read(buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("bytes read: ", bytesread)
	// fmt.Println("bytestream to string: ", string(buffer))

	// r := bufio.NewReader(f)
	// for {
	// 	line, err := r.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	} else if err != nil {
	// 		fmt.Printf("error reading file %s", err)
	// 		break
	// 	}
	// 	logs.Info(line)
	// 	fmt.Print(line)
	// 	line2 = line2 + line
	// }
	// logs.Info(line2)

	// jsonData := []byte(line2)
	jsonData := []byte(buffer)
	var v interface{}
	json.Unmarshal(jsonData, &v)
	data := v.(map[string]interface{})

	// for k, v := range data {
	// 	switch v := v.(type) {
	// 	case string:
	// 		// fmt.Println(k, v, "(string)")
	// 	case float64:
	// 		// fmt.Println(k, v, "(float64)")
	// 	case []interface{}:
	// 		// fmt.Println(k, "(array):")
	// 		// for i, u := range v {
	// 		// fmt.Println("    ", i, u)
	// 		// }
	// 	default:
	// 		// fmt.Println(k, v, "(unknown)")
	// 	}
	// }

	c.Data["json"] = data
	c.ServeJSON()
}

// @Title upload mathcadtemple
// @Description get mathcad
// @Success 200 {object} models.GetMathcadPage
// @Failure 400 Invalid page supplied
// @Failure 404 Page not found
// @router /uploadmodel [get]
// 打开上传界面
func (c *FreeCADController) UploadModel() {
	c.TplName = "freecad/uploadmodel.tpl"
}

// @Title post bootstrapfileinput
// @Description post file by BootstrapFileInput
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 page not found
// @router /uploadmodelfile [post]
// 上传mathcad模板文件，解析模板文件的输入输出存到数据库中
func (c *FreeCADController) UploadModelFile() {
	// 取得用户名
	// username, _, uid, _, _ := checkprodRole(c.Ctx)
	//获取上传的文件
	_, h, err := c.GetFile("input-ke-2[]")
	// beego.Info(h.Filename)自动将英文括号改成了_下划线
	if err != nil {
		logs.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	if fileSuffix != ".TXT" && fileSuffix != ".txt" && fileSuffix != ".STP" && fileSuffix != ".stp" && fileSuffix != ".STEP" && fileSuffix != ".step" && fileSuffix != ".FCStd" {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "文件类型错误", "pdflink": ""}
		c.ServeJSON()
		return
	}

	// random_name
	// newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	// year, month, _ := time.Now().Date()
	err = os.MkdirAll("./attachment/freecad/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		logs.Error(err)
	}
	var filesize int64

	if h != nil {
		//保存附件——要防止重名覆盖！！！！先判断是否存在！！！
		// filepath := "./attachment/mathcad/" + username + "/" + h.Filename
		filepath := "./attachment/freecad/" + h.Filename
		Url := "/attachment/freecad/"
		// 如果文件存在，则返回
		if PathisExist(filepath) {
			c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "文件已存在！"}
			c.ServeJSON()
			return
		}
		err = c.SaveToFile("input-ke-2[]", filepath) //.Join("attachment", attachment)) //存文件
		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "文件保存错误！"}
			c.ServeJSON()
			return
		}
		filesize, _ = FileSize(filepath)
		filesize = filesize / 1000.0

		// 解析mcdx,pdf直接保存
		if fileSuffix == ".txt" || fileSuffix == ".stp" || fileSuffix == ".step" || fileSuffix == ".FCStd" {
			_, modelnumber, modeltitle, version := FreecadName(h.Filename)
			// logs.Info(modelnumber)
			// logs.Info(modeltitle)
			// logs.Info(version)
			//根据id
			_, err := models.AddFreecadModel(9, modelnumber, h.Filename, modeltitle, filepath, version)
			if err != nil {
				logs.Error(err)
			}
		}
		var fileinput Fileinput
		fileinput.InitialPreview = []string{Url + h.Filename}
		// fileinput.InitialPreviewDownloadUrl = []string{Url + newname}
		// var config PreviewConfig
		config := make([]PreviewConfig, 1)
		config[0].Caption = h.Filename
		config[0].DownloadUrl = Url + h.Filename
		config[0].Size = filesize
		config[0].Key = Url + h.Filename
		config[0].Url = Url + h.Filename
		fileinput.InitialPreviewConfig = config
		c.Data["json"] = fileinput
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "文件上传为空", "pdflink": ""}
		// c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "title": "", "original": ""}
		c.ServeJSON()
	}
}

// 取得freecadmodel名称，日期和版本
func FreecadName(filenameWithSuffix string) (Suffix, FileNumber, FileName, Version string) {
	// beego.Info("文件名：", filenameWithSuffix)
	FileSuffix := path.Ext(filenameWithSuffix) //只留下后缀名
	LengthSuffix := len([]rune(FileSuffix))
	Suffix = SubString(FileSuffix, 1, LengthSuffix-1)

	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, FileSuffix) //只留下文件名，无后缀

	// beego.Info("文件全名：", filenameOnly) //filenameOnly= mai
	//这个测试一个字符串是否符合一个表达式。
	//    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	//    fmt.Println(match)
	//上面我们是直接使用字符串，但是对于一些其他的正则任务，你需要使用 Compile 一个优化的 Regexp 结构体。
	// r, _ := regexp.Compile(`[[:upper:]]{2}[0-9]+[[:upper:]\.0-9]+[-][0-9]+[-][0-9]+[\p{Han} \(\)\/~]`)
	//这个结构体有很多方法。这里是类似我们前面看到的一个匹配测试。
	// fmt.Println(r.MatchString(filenameOnly))
	lengthname := len([]rune(filenameOnly))
	vindex := UnicodeIndex(filenameOnly, "_v") // 查找"_v"这个字符的位置
	if vindex == 0 {
		vindex = UnicodeIndex(filenameOnly, "_V")
	}
	// beego.Info(vindex)
	// beego.Info("文件名：", filenameOnly)
	if vindex == 0 {
		Version = "0.0.0"
	} else {
		Version = SubString(filenameOnly, vindex+2, lengthname-vindex-1)
		filenameOnly = SubString(filenameOnly, 0, vindex)
		// beego.Info("文件名：", filenameOnly)
		lengthname = len([]rune(filenameOnly))
	}

	// 查找连续2个的大写字母
	// reg := regexp.MustCompile(`[[:upper:]]{2}`)
	// fmt.Printf("大写字母%q\n", reg.FindAllString(filenameOnly, -1))
	// ["H" "G"]
	blankloc := UnicodeIndex(filenameOnly, " ") // 查找空格这个字符的位置
	if blankloc == 0 {                          //如果没有空格,                                                   //如果没有空格，则用正则表达式获取编号
		re, _ := regexp.Compile("[^a-zA-Z0-9-.~]")
		loc := re.FindStringIndex(filenameOnly)
		if loc != nil { //如果有编号——如果没文件名？？？？？
			FileNumber = SubString(filenameOnly, 0, loc[0])
			// logs.Info("文件编号：", FileNumber)
			FileName = SubString(filenameOnly, loc[0], lengthname-loc[0])
			// logs.Info("文件名：", FileName)
		} else { //如果没有编号
			FileNumber = filenameOnly
			// logs.Info("文件编号：", FileNumber)
			FileName = filenameOnly
			// logs.Info("文件名：", filenameOnly)
		}
	} else { //如果有空格
		re, _ := regexp.Compile("[^a-zA-Z0-9-.~]")
		loc := re.FindStringIndex(filenameOnly)
		if loc != nil { //如果有编号
			FileNumber = SubString(filenameOnly, 0, loc[0])
			// logs.Info("文件编号：", FileNumber)
			FileName = SubString(filenameOnly, loc[0], lengthname-loc[0])
			// logs.Info("文件名：", FileName)
		} else { //如果没有编号
			FileNumber = filenameOnly
			// logs.Info("文件编号：", FileNumber)
			FileName = filenameOnly
			// logs.Info("文件名：", filenameOnly)
		}
	}

	return Suffix, FileNumber, FileName, Version
}

// 递归构造项目树状目录_带成果数量_只显示项目层和下面第一层
func makefreecadtreejson(freecadmodels []models.FreecadModel, nodes *FreeCADList) {
	// 遍历目录
	for _, v := range freecadmodels { //0001-0100
		var tags [1]string
		var state State2
		// 将当前名和id作为子节点添加到目录下
		child := FreeCADList{v.ID, v.Number, tags, false, v.Number + v.TitleB, true, state, []*FreeCADList{}}
		nodes.Nodes = append(nodes.Nodes, &child)
	}
	return
}

// 递归构造项目树状目录_带成果数量_只显示项目层和下面第一层
func makefreecadmenujson(freecadmodels []models.FreecadModel, nodes *FreeCADMenu) {
	// 遍历目录
	for _, v := range freecadmodels { //0001-0100
		var tags [1]string
		// 将当前名和id作为子节点添加到目录下
		child := FreeCADMenu{v.ID, strconv.Itoa(int(v.ID)), tags, false, v.Number + v.TitleB, []*FreeCADMenu{}}
		nodes.Children = append(nodes.Children, &child)
	}
	return
}

// type FreeCADList struct {
// 	Id         int64         `json:"id"`
// 	Code       string        `json:"code"` //分级目录代码
// 	Tags       [1]string     `json:"tags"` //显示员工数量，如果定义为数值[1]int，则无论如何都显示0，所以要做成字符
// 	LazyLoad   bool          `json:"lazyLoad"`
// 	Text       string        `json:"text"`
// 	Selectable bool          `json:"selectable"`
// 	State      State2        `json:"state"`
// 	Nodes      []FreeCADList `json:"nodes"`
// }
