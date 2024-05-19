package controllers

import (
	// beego "github.com/beego/beego/v2/adapter"
	// ole "github.com/go-ole/go-ole"
	// "github.com/go-ole/go-ole/oleutil"
	// "log"
	// "math"
	"os"
	// "os/exec"
	"path"
	"strconv"
	"time"
	// "hydrocms/models"
	// "encoding/base64"
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
	"bytes"
	"crypto/rand"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"math/big"
	"os/exec"
	// "reflect"
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

// type ParseFC struct {
// 	Num    int     `json:"num"`
// 	Name   string  `json:"name"`
// 	Symbol string  `json:"symbol"`
// 	Unit   string  `json:"unit"`
// 	Value  float64 `json:"value"`
// 	Remark string  `json:"remark"`
// }

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
// 模型展示页面
func (c *FreeCADController) FreeCAD() {
	// var err error
	// var state State2
	// var tags [1]string
	// tags[0] = strconv.Itoa(100)
	// //递归生成目录json
	// root := FreeCADList{1, "0001-0100", tags, false, "0001-0100", false, state, []*FreeCADList{}}
	// // makemathtreejson(cates, categories, products, &root)
	// // c.Data["json"] = root
	// freecadmodels, err := models.GetFreecadModels(0, 0, "")
	// if err != nil {
	// 	logs.Error(err)
	// }
	// makefreecadtreejson(freecadmodels, &root)
	// c.Data["json"] = root //data
	// // 初始展示模型，随机
	// if len(freecadmodels) > 0 {
	// 	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(freecadmodels))))
	// 	number := n.String()             //转成string
	// 	num, err := strconv.Atoi(number) //string转int
	// 	if err != nil {
	// 		logs.Error(err)
	// 	}
	// 	c.Data["ModelId"] = freecadmodels[num].ID
	// } else {
	// 	c.Data["ModelId"] = 1
	// }
	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		// c.TplName = "freecad/freecad.tpl"
		c.TplName = "freecad/freecad.html"
		// c.TplName = "freecad/index.html"//蜗壳，保留
	} else {
		// c.TplName = "freecad/freecad.tpl"
		c.TplName = "freecad/freecad.html"
	}
}

// @Title get freecad model update by id...
// @Description get freecad model update by id...
// @Param id query string true "The id of freecad"
// @Success 200 {object} models.UserTemple
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /update [get]
// 模型展示页面
func (c *FreeCADController) Update() {
	_, _, uid, isadmin, _ := checkprodRole(c.Ctx)
	if uid == 0 {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "用户未登录！"}
		c.ServeJSON()
		return
	}
	id := c.GetString("id")
	//id转成uint为
	idint, err := strconv.Atoi(id)
	if err != nil {
		logs.Error(err)
	}
	modelid := uint(idint)
	freecad, err := models.GetFreecadModel(modelid)
	if err != nil {
		logs.Error(err)
	}
	if freecad.UserID != uid && !isadmin {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "非本人或管理员！"}
		c.ServeJSON()
		return
	}
	c.Data["ID"] = freecad.ID
	c.Data["FreeCADModel"] = freecad

	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		c.TplName = "freecad/freecad_update.tpl"
	} else {
		c.TplName = "freecad/freecad_update.tpl"
	}
}

// @Title post freecad model update by id...
// @Description post freecad model update by id...
// @Param id query string true "The id of freecad"
// @Param description query string false "The description of freecad"
// @Param indicators query string false "The indicators of freecad"
// @Success 200 {object} models.FreeCAD
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /updatefreecad [post]
// 编辑 freecad模型的指标和描述
func (c *FreeCADController) UpdateFreeCAD() {
	_, _, uid, isadmin, _ := checkprodRole(c.Ctx)
	if uid == 0 {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "用户未登录！"}
		c.ServeJSON()
		return
	}
	id := c.GetString("id")
	//id转成uint为
	idint, err := strconv.Atoi(id)
	if err != nil {
		logs.Error(err)
	}
	modelid := uint(idint)
	freecad, err := models.GetFreecadModel(modelid)
	if err != nil {
		logs.Error(err)
	}
	if freecad.UserID != uid && !isadmin {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "非本人或管理员！"}
		c.ServeJSON()
		return
	}

	description := c.GetString("description")
	indicators := c.GetString("indicators")

	err = models.UpdateFreecadModel(modelid, description, indicators)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "更新数据库错误！"}
	} else {
		c.Data["json"] = map[string]interface{}{"info": "修改成功！", "state": "SUCCESS", "data": freecad.TitleB}
		c.ServeJSON()
	}
}

// @Title post freecad model delete by id...
// @Description post freecad model delete by id...
// @Param id query string true "The id of freecad"
// @Success 200 {object} models.FreeCAD
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /deletefreecad [post]
// 根据id删除freecad
func (c *ArticleController) DeleteFreeCAD() {
	// _, role := checkprodRole(c.Ctx)
	_, _, _, isadmin, _ := checkprodRole(c.Ctx)
	if !isadmin {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "非管理员！"}
		c.ServeJSON()
	}

	id := c.GetString("id")
	//id转成uint为
	idint, err := strconv.Atoi(id)
	if err != nil {
		logs.Error(err)
	}
	modelid := uint(idint)
	err = models.DeleteFreeCAD(modelid)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "删除数据出错！"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "state": "SUCCESS", "data": "删除完成！"}
		c.ServeJSON()
	}
}

// @Title get freecad models by keyword...
// @Description get freecad models by keyword...
// @Param keyword query string false "The keyword of freecad models"
// @Param page query string false "The page of freecad models"
// @Param limit query string false "The size of page"
// @Success 200 {object} models.UserTemple
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /freecaddata [get]
// 模型展示页面数据
func (c *FreeCADController) FreeCADData() {
	limit := c.GetString("limit")
	if limit == "" {
		limit = "9"
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

	fcmodellist, err := models.GetFreecadModels(limit1, offset, "")
	if err != nil {
		logs.Error(err)
	}

	c.Data["json"] = fcmodellist
	c.ServeJSON()
}

// @Title get online3dview
// @Description getthreejs
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /online3dview [get]
// freecad页面
func (c *FreeCADController) Online3dview() {
	c.TplName = "online3dview/3dview.html"
}

// @Title get online3dviewembed
// @Description getthreejs
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /online3dviewembed [get]
// freecad页面
func (c *FreeCADController) Online3dviewembed() {
	c.TplName = "online3dview/3dview_embed.html"
}

// @Title get online3deditor
// @Description online3deditor
// @Param model query string true "The model url"
// @Param id query string true "The model id"
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /online3deditor [get]
// freecad页面
func (c *FreeCADController) Oline3dEditor() {
	username, _, _, _, islogin := checkprodRole(c.Ctx)
	if !islogin {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "用户未登录！"}
		c.ServeJSON()
		return
	}
	model := c.GetString("model")
	c.Data["Model"] = model
	id := c.GetString("id")
	//id转成uint为
	idint, err := strconv.Atoi(id)
	if err != nil {
		logs.Error(err)
	}
	modelid := uint(idint)
	freecad, err := models.GetFreecadModel(modelid)
	if err != nil {
		logs.Error(err)
	}
	c.Data["FreeCADModel"] = freecad
	c.Data["DxfSvgUrl"] = strings.Replace(freecad.DxfSvgUrl, "./", "/", -1)
	c.Data["RenderImgUrl"] = strings.Replace(freecad.RenderImgUrl, "./", "/", -1)
	c.Data["ID"] = id

	c.Data["IsLogin"] = islogin
	c.Data["UserName"] = username
	c.TplName = "freecad/3deditor.html"
}

// @Title get freecad parameter
// @Description get freecad parameter
// @Param id query string true "The model id"
// @Success 200 {object} models.FreecadInputs
// @Failure 400 Invalid page supplied
// @Failure 404 model not found
// @router /getfcparameter/:id [get]
// freecad页面
func (c *FreeCADController) GetFCParameter() {
	id := c.Ctx.Input.Param(":id")
	//id转成uint为
	idint, err := strconv.Atoi(id)
	if err != nil {
		logs.Error(err)
	}
	modelid := uint(idint)
	// 从路径中获取这个modelid，再传给页面中的ajax获取模型txt文件
	c.Data["ModelId"] = modelid

	freecadinputs, err := models.GetFCModelInput(modelid)
	if err != nil {
		logs.Info(err)
	}

	c.Data["json"] = freecadinputs
	c.ServeJSON()
}

// @Title get turbin
// @Description getthreejs
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /turbin [get]
// freecad页面
func (c *FreeCADController) Turbin() {
	c.TplName = "freecad/turbin.html" //蜗壳，保留
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

// @Title get freecad model by id...
// @Description get freecad model by id...
// @Param id path string true "The id of freecad model"
// @Success 200 {object} models.FreecadModel
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /freecadmodel/:id [get]
// 获取模型数据
func (c *FreeCADController) FreeCADModel() {
	var err error
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

	c.Data["json"] = freecadmodel
	c.ServeJSON()
}

func (c *FreeCADController) FreeCADModel_back() {
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
		logs.Info(err)
		return
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		logs.Info(err)
		return
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	bytesread, err := file.Read(buffer)
	if err != nil {
		logs.Info(err)
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
	_, _, _, _, islogin := checkprodRole(c.Ctx)
	if !islogin {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "用户未登录！"}
		c.ServeJSON()
		return
	}
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
	_, _, uid, _, _ := checkprodRole(c.Ctx)
	//获取上传的文件
	_, h, err := c.GetFile("input-ke-2[]")
	// beego.Info(h.Filename)自动将英文括号改成了_下划线
	if err != nil {
		logs.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	if fileSuffix != ".FCStd" && fileSuffix != ".fcstd" && fileSuffix != ".STEP" && fileSuffix != ".step" && fileSuffix != ".glb" && fileSuffix != ".obj" && fileSuffix != ".jpg" && fileSuffix != ".png" && fileSuffix != ".gif" && fileSuffix != ".bmp" && fileSuffix != ".svg" {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "文件类型错误！fcstd/stp/glb/obj/jpg/png/bmp/svg", "flink": ""}
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

	cwd, _ := os.Getwd()
	workpath := cwd + "/attachment/freecad/" //+ username + "/" + newname + "/"

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

		_, modelnumber, modeltitle, version := FreecadName(h.Filename)
		// 解析fcstd
		if fileSuffix == ".FCStd" || fileSuffix == ".fcstd" {
			//根据id
			fcId, err := models.AddFreecadModel(uid, modelnumber, h.Filename, modeltitle, filepath, version)
			if err != nil {
				logs.Error(err)
			}
			parsefcpath, err := web.AppConfig.String("parsefcpath")
			if err != nil {
				logs.Error(err)
			}
			convertfcpath, err := web.AppConfig.String("convertfcpath")
			if err != nil {
				logs.Error(err)
			}

			arg := []string{"-i", workpath + h.Filename}
			fmt.Println("-----parse--arg-------", arg)
			cmd := exec.Command(parsefcpath, arg...)
			//记录开始时间
			// start := time.Now()
			// err = cmd.Start()
			// if err != nil {
			// 	//fmt.Println(err)
			// 	fmt.Printf("err: %v", err)
			// }
			// err = cmd.Wait() //Wait等待command退出，他必须和Start一起使用，如果命令能够顺利执行完并顺利退出则返回nil，否则的话便会返回error，其中Wait会是放掉所有与cmd命令相关的资源
			// if err != nil {
			// 	fmt.Printf("err plot: %v", err)
			// }
			// //记录结束时间差
			// elapsed := time.Since(start)
			// fmt.Printf("elapsed plot: %s\n", elapsed)
			// var out bytes.Buffer
			// buf, err := cmd.Output() //运行命令并返回其标准输出
			buf, err := cmd.CombinedOutput()
			// _, err = cmd.CombinedOutput()
			if err != nil {
				logs.Info(err)
			}
			// fmt.Println("buf plot:", buf)
			// fmt.Println("buf plot:", string(buf))

			// 实例化结构体
			// parsfc := ParseFC{}
			var parsfc []models.FreecadInputs
			str, err := GbkToUtf8(buf)
			if err != nil {
				logs.Error(err)
			}
			err = json.Unmarshal(str, &parsfc)
			// err = json.Unmarshal(buf, &parsfc)
			if err != nil {
				logs.Info(err)
			}
			// 写入数据库
			for _, v := range parsfc {
				// fmt.Println(reflect.TypeOf(v.InputValue))
				_, err = models.CreateFCModelInput(fcId, v)
				if err != nil {
					logs.Error(err)
				}
			}
			// for i, v := range parsfc {
			// 	logs.Info(i, v)
			// 	logs.Info(v.Value)
			// }
			// str1 := strings.Replace(strings.Replace(string(str), "[[", "", -1), "]]", "", -1)
			// // Enter = 回车+换行(\r\n)
			// // 理解：
			// // \n是换行，英文是New line
			// // \r是回车，英文是Carriage return
			// str1 = strings.Replace(str1, "\r", "", -1)
			// str1 = strings.Replace(str1, "\n", "", -1)
			// str1_arr := strings.Split(str1, "], [")
			// for _, v := range str1_arr {
			// 	str2_arr := strings.Split(v, ", ")
			// 	for _, w := range str2_arr {
			// 		// float, err := strconv.ParseFloat(w, 64)
			// 		// if err != nil {
			// 		// 	logs.Error(err)
			// 		// }
			// 		logs.Info(w)
			// 	}
			// }
			// filenameall := path.Base(filename)
			filesuffix := path.Ext(h.Filename)
			// fileprefix := filenameall[0:len(filenameall) - len(filesuffix)]
			fileprefix := strings.TrimSuffix(h.Filename, filesuffix)
			// 导出gltf会自动导出相应的bin文件，会带颜色。如果仅导出glb文件，则是一个独立的文件，不带颜色。但是fc里直接导出glb，则会带颜色
			arg = []string{"-i", workpath + h.Filename, "-o", workpath + fileprefix + ".gltf"}
			fmt.Println("-----convertfc--arg-------", arg)
			cmd = exec.Command(convertfcpath, arg...)
			_, err = cmd.CombinedOutput()
			// _, err = cmd.CombinedOutput()
			if err != nil {
				logs.Info(err)
			}
		} else if fileSuffix == ".JPG" || fileSuffix == ".jpg" || fileSuffix == ".png" || fileSuffix == ".gif" || fileSuffix == ".svg" {
			// 上传的是封面文件
			_, err = models.AddFreecadModelFace(uid, modelnumber, modeltitle, filepath, version)
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
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "文件上传为空", "fclink": ""}
		// c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "title": "", "original": ""}
		c.ServeJSON()
	}
}

// @Title post freecad model parameter
// @Description post freecad model parameter json by ajax
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 page not found
// @router /editfcmodel [post]
// 提交修改后的参数
func (c *FreeCADController) EditFCModel() {
	username, _, _, _, islogin := checkprodRole(c.Ctx)
	if !islogin {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "用户未登录！"}
		c.ServeJSON()
		return
	}

	inputdata := c.GetString("inputdata")
	var freecadinputs []models.FreecadInputs
	err := json.Unmarshal([]byte(inputdata), &freecadinputs)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(templeinputs)
	modelid := c.GetString("modelid")
	//id转成uint为
	modelidint, err := strconv.Atoi(modelid)
	if err != nil {
		logs.Error(err)
	}
	modeliduint := uint(modelidint)

	// freecadmodelinput, err := models.GetFCModelInput(modeliduint)
	// if err != nil {
	// 	logs.Info(err)
	// }

	freecadmodel, err := models.GetFreecadModel(modeliduint)
	if err != nil {
		logs.Info(err)
	}

	var parameter string
	for i, v := range freecadinputs {
		logs.Info(v.InputValue)
		if i == 0 {
			parameter = strconv.FormatFloat(v.InputValue, 'f', -1, 64)
		} else {
			parameter = parameter + " " + strconv.FormatFloat(v.InputValue, 'f', -1, 64)
		}
	}

	convertfcpath, err := web.AppConfig.String("convertfcpath")
	if err != nil {
		logs.Error(err)
	}

	newname := strconv.FormatInt(time.Now().UnixNano(), 10)
	// year, month, _ := time.Now().Date()
	err = os.MkdirAll("./attachment/freecad/"+username+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		logs.Error(err)
	}
	cwd, _ := os.Getwd()
	filepath := cwd + strings.Replace(strings.Replace(freecadmodel.Path, "./", "/", -1), "/", "\\", -1)
	userfilepath := cwd + "\\attachment\\freecad\\" + username + "\\"
	userfileurl := "/attachment/freecad/" + username + "/"
	// 导出gltf会自动导出相应的bin文件，会带颜色。如果仅导出glb文件，则是一个独立的文件，不带颜色。但是fc里直接导出glb，则会带颜色
	arg := []string{"-i", filepath, "-o", userfilepath + newname + ".glb", "-p", parameter}
	fmt.Println("-----convertfc--arg-------", arg)
	cmd := exec.Command(convertfcpath, arg...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	// cmd.Stderr = os.Stderr
	cmd.Stderr = &stderr
	err = cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\n err:\n%s\n", outStr, errStr)
	if err != nil {
		fmt.Printf("cmd.Run() failed with %s\n", err)
	}
	// err = cmd.Start()
	// if err != nil {
	// 	fmt.Printf("err: %v", err)
	// }

	// err = cmd.Wait() //Wait等待command退出，他必须和Start一起使用，如果命令能够顺利执行完并顺利退出则返回nil，否则的话便会返回error，其中Wait会释放掉所有与cmd命令相关的资源
	// if err != nil {
	// 	fmt.Printf("err: %v", err)
	// }

	// buf, err := cmd.Output() //运行命令并返回其标准输出
	// fmt.Println("buf:", buf)
	// // _, err = cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// _, err = cmd.CombinedOutput()
	// if err != nil {
	// 	logs.Info(err)
	// }

	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "data": "修改成功！", "glblink": userfileurl + newname + ".glb"}
	c.ServeJSON()
}

// @Title post freecad model parameter batch
// @Description post freecad model parameter batch json by ajax
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 page not found
// @router /editfcmodelbatch [post]
// 上传mathcad模板文件，解析模板文件的输入输出存到数据库中
func (c *FreeCADController) EditFCModelBatch() {
	username, _, _, _, islogin := checkprodRole(c.Ctx)
	if !islogin {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "用户未登录！"}
		c.ServeJSON()
		return
	}
	logs.Info(username)
	inputdata := c.GetString("inputdata")
	// logs.Info(inputdata)
	var parameter_arry [][]string
	err := json.Unmarshal([]byte(inputdata), &parameter_arry)
	if err != nil {
		logs.Info(err)
	}
	// logs.Info(&wo)

	modelid := c.GetString("modelid")
	//id转成uint为
	modelidint, err := strconv.Atoi(modelid)
	if err != nil {
		logs.Error(err)
	}
	modeliduint := uint(modelidint)

	freecadmodel, err := models.GetFreecadModel(modeliduint)
	if err != nil {
		logs.Info(err)
	}

	convertfcpath, err := web.AppConfig.String("convertfcpath")
	if err != nil {
		logs.Error(err)
	}
	err = os.MkdirAll("./attachment/freecad/"+username+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		logs.Error(err)
	}
	cwd, _ := os.Getwd()
	filepath := cwd + strings.Replace(strings.Replace(freecadmodel.Path, "./", "/", -1), "/", "\\", -1)
	userfilepath := cwd + "\\attachment\\freecad\\" + username + "\\"
	userfileurl := "/attachment/freecad/" + username + "/"

	// 二维数组转置
	logs.Info(parameter_arry)
	parameter := transpose(parameter_arry)
	logs.Info(parameter)
	var parameter_str string
	var glblink_batch []string
	for _, v := range parameter {
		for j, w := range v {
			if j == 0 {
				parameter_str = w
			} else {
				parameter_str = parameter_str + " " + w
			}
		}
		logs.Info(parameter_str)
		newname := strconv.FormatInt(time.Now().UnixNano(), 10)
		arg := []string{"-i", filepath, "-o", userfilepath + newname + ".glb", "-p", parameter_str}
		fmt.Println("-----convertfc--arg-----", arg)
		cmd := exec.Command(convertfcpath, arg...)

		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		// cmd.Stderr = os.Stderr
		cmd.Stderr = &stderr
		err = cmd.Run()
		outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
		fmt.Printf("out:\n%s\n err:\n%s\n", outStr, errStr)
		if err != nil {
			fmt.Printf("cmd.Run() failed with %s\n", err)
		}
		glblink_batch = append(glblink_batch, userfileurl+newname+".glb")
	}

	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "data": "修改成功！", "glblink_batch": glblink_batch}
	c.ServeJSON()
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
	if blankloc == 0 {                          //如果没有空格,//如果没有空格，则用正则表达式获取编号
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

func stringToTwoDimensionalArray(s string) [][]int {
	rows := strings.Split(s, ",")
	if len(rows) == 0 {
		return nil
	}

	result := make([][]int, len(rows))
	for i, row := range result {
		if i >= len(rows)-1 {
			result[i] = []int{}
		} else {
			rowArr := make([]int, len(row))
			for j, value := range row {
				rowArr[j] = int(value)
			}
			result[i] = rowArr
		}
	}

	return result
}

// #GbkToUtf8 方法
// import (
// 	"bytes"
// 	"golang.org/x/text/encoding/simplifiedchinese"
// 	"golang.org/x/text/transform"
// 	"io/ioutil"
// )

func GbkToUtf8(str []byte) (b []byte, err error) {
	r := transform.NewReader(bytes.NewReader(str), simplifiedchinese.GBK.NewDecoder())
	b, err = ioutil.ReadAll(r)
	if err != nil {
		return
	}
	return
}

// 二维数组转置
func transpose(A [][]string) [][]string {
	B := make([][]string, len(A[0]))
	for i := 0; i < len(A[0]); i++ {
		B[i] = make([]string, len(A))
		for j := 0; j < len(A); j++ {
			B[i][j] = A[j][i]
		}
	}
	return B
}

// ————————————————
// 版权声明：本文为CSDN博主「hjx_dou」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
// 原文链接：https://blog.csdn.net/hjx_dou/article/details/108105895

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
