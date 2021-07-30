package controllers

import (
	"github.com/3xxx/engineercms/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/tealeg/xlsx"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type StandardController struct {
	beego.Controller
}

type Standardmore struct {
	Id            int64
	Number        string
	Title         string
	NumTitle      string
	Uname         string //换成用户名
	CategoryName  string //换成规范类别GB、DL……
	Content       string `orm:"sie(5000)"`
	Route         string
	Link          string
	ActIndex      string
	Created       time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"index","auto_now;type(datetime)"`
	Views         int64     `orm:"index"`
	LibraryNumber string    //规范有效版本库中的编号
	LibraryTitle  string
	LiNumber      string //完整编号
}

//显示所有规范
func (c *StandardController) GetStandard() {
	c.Data["IsStandard"] = true //这里修改到ListAllPosts()
	// c.TplName = "standard.tpl"
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	// uname, _, _ := checkRoleread(c.Ctx) //login里的
	// c.Data["Uname"] = uname
	standards, err := models.GetAllStandards()
	if err != nil {
		beego.Error(err.Error)
	} else {
		c.Data["json"] = standards
		c.ServeJSON()
	}
}

//修改规范
func (c *StandardController) UpdateStandard() {
	_, _, _, isadmin, isLogin := checkprodRole(c.Ctx)
	if !isLogin {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		c.Data["json"] = "未登陆"
		c.ServeJSON()
		return
	} else if !isadmin {
		c.Data["json"] = "非管理员"
		c.ServeJSON()
		return
	}
	id := c.Input().Get("cid")
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	number := c.Input().Get("number")
	title := c.Input().Get("title")
	route := c.Input().Get("route")
	uname := c.Input().Get("uname")
	var uid int64
	// uname查询出uid
	if uname != "" {
		user, err := models.GetUserByUsername(uname)
		if err != nil {
			beego.Error(err)
		}
		uid = user.Id
	}
	err = models.UpdateStandard(idNum, uid, number, title, route)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = "修改出错"
		c.ServeJSON()
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

//删除规范
func (c *StandardController) DeleteStandard() {
	_, _, _, isadmin, isLogin := checkprodRole(c.Ctx)
	if !isLogin {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		c.Data["json"] = "未登陆"
		c.ServeJSON()
		return
	} else if !isadmin {
		c.Data["json"] = "非管理员"
		c.ServeJSON()
		return
	}
	ids := c.GetString("ids")
	array := strings.Split(ids, ",")
	// beego.Info(array)
	for _, v := range array {
		// pid = strconv.FormatInt(v1, 10)
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//循环删除成果
		//根据成果id取得所有附件
		err = models.DeleteStandard(idNum)
		if err != nil {
			beego.Error(err)
			c.Data["json"] = "删除出错"
			c.ServeJSON()
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	}
}

func (c *StandardController) Index() { //
	c.Data["IsStandard"] = true
	u := c.Ctx.Input.UserAgent()
	// re := regexp.MustCompile("Trident")
	// loc := re.FindStringIndex(u)
	// loc[0] > 1
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		beego.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		c.TplName = "mobile/mstandard.tpl"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "standard.tpl"
	}

	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	// rolename, _ = strconv.Atoi(role)
	// c.Data["Uname"] = uname
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	standards, err := models.GetAllStandards() //这里传入空字符串
	if err != nil {
		beego.Error(err.Error)
	} else {
		// c.Data["Standards"] = standards   //这个没用吧
		c.Data["Length"] = len(standards) //得到总记录数
	}

	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/test.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP())
	logs.Close()
}

//后端分页的数据结构
type StandardTableserver struct {
	Rows  []Standardmore `json:"rows"`
	Page  int            `json:"page"`
	Total int64          `json:"total"` //string或int64都行！
}

// @Title get standardlist
// @Description get standardlist
// @Param searchText query string false "The searchText of standard"
// @Param pageNo query string true "The page for standard list"
// @Param limit query string true "The limit of page for standard list"
// @Success 200 {object} models.Create
// @Failure 400 Invalid page supplied
// @Failure 404 cart not found
// @router /search [get]
//根据用户输入的关键字，查询规范
//搜索规范或者图集的名称或编号
//20170704：linumber没有用。因为用category+编号+年份比较好
func (c *StandardController) Search() { //search用的是post方法
	limit := c.Input().Get("limit")
	if limit == "" {
		limit = "15"
	}
	limit1, err := strconv.Atoi(limit)
	if err != nil {
		beego.Error(err)
	}
	page := c.Input().Get("pageNo")
	page1, err := strconv.Atoi(page)
	if err != nil {
		beego.Error(err)
	}
	var offset int
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	searchText := c.Input().Get("searchText")
	if searchText == "" {
		c.Data["json"] = map[string]interface{}{"code": "OK", "msg": "关键字为空"}
		c.ServeJSON()
		return
	}
	if searchText == "allstandard" {
		standards, err := models.GetAllStandards()
		if err != nil {
			beego.Error(err.Error)
		} else {
			c.Data["json"] = standards
			c.ServeJSON()
		}
	} else {
		//搜索名称
		// Results1, err := models.SearchStandardsName(searchText, false)
		// if err != nil {
		// 	beego.Error(err.Error)
		// }
		//搜索编号
		// Results2, err := models.SearchStandardsNumber(searchText, false)
		// if err != nil {
		// 	beego.Error(err.Error)
		// }
		// Results1 = append(Results1, Results2...)
		Results1, err := models.GetUserStandard(limit1, offset, searchText)
		if err != nil {
			beego.Error(err.Error)
		}
		//由categoryid查categoryname
		aa := make([]Standardmore, len(Results1))
		//由standardnumber查librarynumber
		for i, v := range Results1 {
			//由standardnumber正则得到编号50268和分类GB
			Category, _, Number := SplitStandardFileNumber(v.Number)
			//由分类和编号查有效版本库中的编号
			library, err := models.SearchLiabraryNumber(Category, Number)
			if err != nil {
				beego.Error(err.Error)
			}
			aa[i].Id = v.Id
			aa[i].Number = v.Number //`orm:"unique"`
			aa[i].Title = v.Title
			aa[i].Uname = v.UserName //换成用户名
			aa[i].Route = v.Route
			aa[i].Created = v.Created
			aa[i].Updated = v.Updated
			aa[i].Views = v.Views
			if library != nil {
				aa[i].LibraryNumber = library.Number //规范有效版本库中的编号
				aa[i].LibraryTitle = library.Title
				aa[i].LiNumber = library.Category + " " + library.Number + "-" + library.Year //完整编号
			} else {
				aa[i].LiNumber = "No LibraryNumber Match Find!"
				aa[i].LibraryTitle = ""
				aa[i].LibraryNumber = ""
			}
		}
		count, err := models.GetUserStandardCount(searchText)
		if err != nil {
			beego.Error(err)
		}
		table := StandardTableserver{aa, page1, count}
		c.Data["json"] = table //这里必须要是c.Data["json"]，其他c.Data["Data"]不行
		c.ServeJSON()
	}

	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/test.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "SearchStandardsName:" + searchText)
	logs.Close()
	// standards, err := models.GetAllStandards() //这里传入空字符串
	// if err != nil {
	// 	beego.Error(err.Error)
	// } else {
	// 	c.Data["Standards"] = standards
	// 	c.Data["Length"] = len(standards) //得到总记录数
	// }
}

// @Title get standardpdf
// @Description get standardpdf
// @Param file query string true "The link of standardpdf"
// @Success 200 {object} models.Create
// @Failure 400 Invalid page supplied
// @Failure 404 cart not found
// @router /standardpdf [get]
func (c *StandardController) StandardPdf() { //search用的是post方法
	_, _, _, isadmin, islogin := checkprodRole(c.Ctx)
	if !isadmin && !islogin {
		// beego.Info(!islogin)
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
	}
	pdflink := c.Input().Get("file")
	c.Data["PdfLink"] = pdflink
	c.TplName = "web/viewer.html"
}

// @Title get wx standards list
// @Description get standards by page
// @Param keyword query string  true "The keyword of standards"
// @Param searchpage query string  true "The page for drawings list"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 standards not found
// @router /searchwxstandards [get]
//小程序取得规范列表，分页_plus
func (c *StandardController) SearchWxStandards() { //search用的是post方法
	// wxsite := beego.AppConfig.String("wxreqeustsite")
	limit := "5"
	limit1, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	page := c.Input().Get("searchpage")
	page1, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	var offset int64
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	key := c.Input().Get("keyword")

	if key != "" {
		//搜索名称
		Results1, err := models.SearchStandardsNamePage(limit1, offset, key, false)
		if err != nil {
			beego.Error(err.Error)
		}
		//搜索编号
		Results2, err := models.SearchStandardsNumberPage(limit1, offset, key, false)
		if err != nil {
			beego.Error(err.Error)
		}
		// Standards := make([]*Standard, 0)
		Results1 = append(Results1, Results2...)
		//由categoryid查categoryname
		aa := make([]Standardmore, len(Results1))
		//由standardnumber查librarynumber
		for i, v := range Results1 {
			//由userid查username
			//由standardnumber正则得到编号50268和分类GB
			Category, _, Number := SplitStandardFileNumber(v.Number)
			//由分类和编号查有效版本库中的编号
			library, err := models.SearchLiabraryNumber(Category, Number)
			if err != nil {
				beego.Error(err.Error)
			}
			aa[i].Id = v.Id
			aa[i].Number = v.Number //`orm:"unique"`
			aa[i].Title = v.Number + v.Title
			// aa[i].NumTitle = v.Number + v.Title
			// aa[i].Route = v.Route
			aa[i].Link = "data:image/svg+xml,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20width%3D%2232%22%20height%3D%2232%22%3E%3Crect%20fill%3D%22%234CAF50%22%20x%3D%220%22%20y%3D%220%22%20width%3D%22100%25%22%20height%3D%22100%25%22%3E%3C%2Frect%3E%3Ctext%20fill%3D%22%23FFF%22%20x%3D%2250%25%22%20y%3D%2250%25%22%20text-anchor%3D%22middle%22%20font-size%3D%2216%22%20font-family%3D%22Verdana%2C%20Geneva%2C%20sans-serif%22%20alignment-baseline%3D%22middle%22%3E%E6%A0%87%3C%2Ftext%3E%3C%2Fsvg%3E" //wxsite + "/static/img/go.jpg"
			aa[i].ActIndex = "standard"
			aa[i].Created = v.Created
			aa[i].Updated = v.Updated
			aa[i].Views = v.Views
			if library != nil {
				aa[i].LibraryNumber = library.Number //规范有效版本库中的编号
				aa[i].LibraryTitle = library.Title
				aa[i].LiNumber = library.Category + " " + library.Number + "-" + library.Year //完整编号
			} else {
				aa[i].LiNumber = "No LibraryNumber Match Find!"
				aa[i].LibraryTitle = ""
				aa[i].LibraryNumber = ""
			}
		}
		c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "searchers": aa}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"info": "关键字为空"}
		c.ServeJSON()
	}

	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/test.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "SearchStandardsName:" + key)
	logs.Close()
}

// @Title dowload wx standardpdf
// @Description get wx standardpdf by id
// @Param id path string  true "The id of standardpdf"
// @Success 200 {object} models.GetAttachbyId
// @Failure 400 Invalid page supplied
// @Failure 404 pdf not found
// @router /wxstandardpdf/:id [get]
// 小程序查询规范
func (c *StandardController) WxStandardPdf() {
	// id := c.Input().Get("id")
	id := c.Ctx.Input.Param(":id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据id取得规范的路径
	standard, err := models.GetStandard(idNum)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(standard.Route)
	fileurl := strings.Replace(standard.Route, "/attachment/", "attachment/", -1)
	// http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, standard.Route)

	filename := path.Base(fileurl)
	fileext := path.Ext(filename)
	matched, err := regexp.MatchString("\\.*[m|M][c|C][d|D]", fileext)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(matched)
	if matched {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "不能下载mcd文件!"}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.Download(fileurl)
}

//显示所有有效库
func (c *StandardController) Valid() { //search用的是post方法
	// name := c.Input().Get("name")
	c.Data["IsStandard"] = true //
	c.TplName = "standard.tpl"
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	// uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	// c.Data["Uname"] = uname
	//搜索名称
	valids, err := models.GetAllValids()
	if err != nil {
		beego.Error(err.Error)
	}
	c.Data["json"] = valids //这里必须要是c.Data["json"]，其他c.Data["Data"]不行
	c.ServeJSON()

	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/test.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "valid:")
	logs.Close()
}

//删除有效库中选中
func (c *StandardController) DeleteValid() { //search用的是post方法
	_, _, _, isadmin, isLogin := checkprodRole(c.Ctx)
	if !isLogin {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		c.Data["json"] = "未登陆"
		c.ServeJSON()
		return
	} else if !isadmin {
		c.Data["json"] = "非管理员"
		c.ServeJSON()
		return
	}
	ids := c.GetString("ids")
	array := strings.Split(ids, ",")
	// beego.Info(array)
	for _, v := range array {
		// pid = strconv.FormatInt(v1, 10)
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//循环删除成果
		//根据成果id取得所有附件
		err = models.DeleteValid(idNum)
		if err != nil {
			beego.Error(err)
			c.Data["json"] = "删除出错"
			c.ServeJSON()
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	}
	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/test.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "valid:")
	logs.Close()
}

//上传excel文件，导入到规范数据库，用于批量导入规范文件
func (c *StandardController) ImportExcel() {
	_, _, _, isadmin, isLogin := checkprodRole(c.Ctx)
	if !isLogin {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		c.Data["json"] = "未登陆"
		c.ServeJSON()
		return
	} else if !isadmin {
		c.Data["json"] = "非管理员"
		c.ServeJSON()
		return
	}
	//获取上传的文件
	_, h, err := c.GetFile("excel")
	if err != nil {
		beego.Error(err)
	}
	// var attachment string
	var path string
	// var filesize int64
	if h != nil {
		//保存附件
		// attachment = h.Filename
		// beego.Info(attachment)
		path = "./attachment/" + h.Filename
		// path := c.Input().Get("url")  //存文件的路径
		// path = path[3:]
		// path = "./attachment" + "/" + h.Filename
		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("excel", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
	}
	if err != nil {
		beego.Error(err)
	}
	var standard models.Standard
	//读出excel内容写入数据库
	// excelFileName := path                    //"/home/tealeg/foo.xlsx"
	xlFile, err := xlsx.OpenFile(path) //excelFileName
	if err != nil {
		beego.Error(err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			// 这里要判断单元格列数，如果超过单元格使用范围的列数，则出错for j := 2; j < 7; j += 5 {
			j := 0
			standard.Number = row.Cells[j].String()
			standard.Title = row.Cells[j+1].String()
			// Uname, err := row.Cells[j+2].String()
			// user := models.GetUserByUsername(Uname)
			// standard.Uid = user.Id
			// Category, err := row.Cells[j+3].String()
			// category, _ := models.GetCategoryName(Category)
			// standard.CategoryId = category.Id
			standard.Created = time.Now()
			standard.Updated = time.Now()
			standard.Content = row.Cells[j+4].String()
			standard.Route = row.Cells[j+5].String()
			_, err = models.SaveStandard(standard)

			if err != nil {
				beego.Error(err)
			}
			// }
			// for _, cell := range row.Cells {这里要继续循环cells，不能为空，即超出单元格使用范围
			// 	fmt.Printf("%s\n", cell.String())

			// }
		}
	}
	c.TplName = "standard.tpl"
	c.Redirect("/standard", 302)
}

//上传excel文件，导入到有效版本数据库
func (c *StandardController) ImportLibrary() {
	_, _, _, isadmin, isLogin := checkprodRole(c.Ctx)
	if !isLogin {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		c.Data["json"] = "未登陆"
		c.ServeJSON()
		return
	} else if !isadmin {
		c.Data["json"] = "非管理员"
		c.ServeJSON()
		return
	}
	//获取上传的文件
	_, h, err := c.GetFile("excel2")
	if err != nil {
		beego.Error(err)
	}
	// var attachment string
	var path string
	// var filesize int64
	if h != nil {
		//保存附件
		// attachment = h.Filename
		// beego.Info(attachment)
		path = "./attachment/" + h.Filename
		// path := c.Input().Get("url")  //存文件的路径
		// path = path[3:]
		// path = "./attachment" + "/" + h.Filename
		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("excel2", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
	}
	var library models.Library
	//读出excel内容写入数据库
	// excelFileName := path//"/home/tealeg/foo.xlsx"
	xlFile, err := xlsx.OpenFile(path) //excelFileName
	if err != nil {
		beego.Error(err)
	} else {
		for _, sheet := range xlFile.Sheets {
			for i, row := range sheet.Rows {
				if i != 0 {
					// 这里要判断单元格列数，如果超过单元格使用范围的列数，则出错for j := 2; j < 7; j += 5 {
					j := 0
					library.Number = row.Cells[j].String()
					if err != nil {
						beego.Error(err)
					}
					library.Title = row.Cells[j+1].String()
					if err != nil {
						beego.Error(err)
					}
					library.Category = row.Cells[j+2].String()
					if err != nil {
						beego.Error(err)
					}
					library.LiNumber = row.Cells[j+3].String()
					if err != nil {
						beego.Error(err)
					}
					library.Year = row.Cells[j+4].String()
					if err != nil {
						beego.Error(err)
					}
					library.Execute = row.Cells[j+5].String()
					if err != nil {
						beego.Error(err)
					}
					library.Created = time.Now()
					library.Updated = time.Now()
					_, err = models.SaveLibrary(library)
					if err != nil {
						beego.Error(err)
					}
				}
				// for _, cell := range row.Cells {这里要继续循环cells，不能为空，即超出单元格使用范围
				// 	fmt.Printf("%s\n", cell.String())
				// }
			}
		}
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS"}
		c.ServeJSON()
	}

	// c.TplName = "standard.tpl"
	// c.Redirect("/standard", 302)
}

func (c *StandardController) Standard_one_addbaidu() { //一对一模式
	_, _, uid, _, _ := checkprodRole(c.Ctx)
	var standard models.Standard
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	//2016-4-23这里将文件的分类强制变为2位，那么GB 122-2016与GBT 122-2016就是一个文件了。
	//是否应该增加一个返回值，将真实的GBT返回来。
	category, categoryname, fileNumber, year, fileName, _ := SplitStandardName(h.Filename)
	var path string
	// var filesize int64
	if h != nil {
		//保存附件
		if category != "" {
			err := os.MkdirAll("./attachment/standard/"+category, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
			if err != nil {
				beego.Error(err)
			}
		}
		path = "./attachment/standard/" + category + "/" + h.Filename
		// path := c.Input().Get("url")  //存文件的路径
		// path = path[3:]
		// path = "./attachment" + "/" + h.Filename
		// f.Close()   // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
		// filesize, _ = FileSize(path)
		// filesize = filesize / 1000.0
	}
	//纯英文下没有取到汉字字符，所以没有名称
	if fileName == "" {
		fileName = fileNumber
	}

	// uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	// c.Data["Uname"] = uname
	if category != "Atlas" {
		standard.Number = categoryname + " " + fileNumber + "-" + year
		standard.Title = fileName
	} else {
		standard.Number = fileNumber
		standard.Title = fileName
	}
	//这里增加Category
	standard.Category = categoryname //2016-7-16这里改为GBT这种，空格前的名字
	standard.Created = time.Now()
	standard.Updated = time.Now()
	standard.Uid = uid
	standard.Route = "/attachment/standard/" + category + "/" + h.Filename
	_, err = models.SaveStandard(standard)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "title": "111", "original": "demo.jpg", "url": standard.Route}
		c.ServeJSON()
	}
}
