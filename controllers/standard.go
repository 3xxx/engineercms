package controllers

import (
	"errors"
	"fmt"
	"github.com/3xxx/engineercms/models"
	"github.com/PuerkitoBio/goquery"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"

	"context"
	//"github.com/PuerkitoBio/goquery"
	"encoding/json"
	// beego "github.com/beego/beego/v2/adapter"
	"github.com/elastic/go-elasticsearch/v8"
	// "github.com/elastic/go-elasticsearch/v8/esapi"
	// "github.com/3xxx/engineercms/controllers/utils"
	"github.com/google/go-tika/tika"
	"github.com/tealeg/xlsx"
	"github.com/xuri/excelize/v2"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type StandardController struct {
	web.Controller
}

type Standardmore struct {
	Id           int64
	Number       string
	Title        string
	NumTitle     string
	Uname        string //换成用户名
	CategoryName string //换成规范类别GB、DL……
	// Content       string `orm:"sie(5000)"`
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

// 显示所有规范
func (c *StandardController) GetStandard() {
	c.Data["IsStandard"] = true //这里修改到ListAllPosts()
	// c.TplName = "standard.tpl"
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	// uname, _, _ := checkRoleread(c.Ctx) //login里的
	// c.Data["Uname"] = uname
	standards, err := models.GetAllStandards()
	if err != nil {
		logs.Error(err.Error)
	} else {
		c.Data["json"] = standards
		c.ServeJSON()
	}
}

// 修改规范
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
	id := c.GetString("cid")
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	number := c.GetString("number")
	title := c.GetString("title")
	route := c.GetString("route")
	uname := c.GetString("uname")
	var uid int64
	// uname查询出uid
	if uname != "" {
		user, err := models.GetUserByUsername(uname)
		if err != nil {
			logs.Error(err)
		}
		uid = user.Id
	}
	err = models.UpdateStandard(idNum, uid, number, title, route)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = "修改出错"
		c.ServeJSON()
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

// 删除规范
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
			logs.Error(err)
		}
		//循环删除成果
		//根据成果id取得所有附件
		err = models.DeleteStandard(idNum)
		if err != nil {
			logs.Error(err)
			c.Data["json"] = "删除出错"
			c.ServeJSON()
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	}
}

func (c *StandardController) Index() {
	c.Data["IsStandard"] = true
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
		c.TplName = "standard.tpl" //"mobile/mstandard.tpl"
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
		logs.Error(err.Error)
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

// @Title get standard elasticsearch web
// @Description get standard elasticsearch web
// @Success 200 {object} models.GetElastic
// @Failure 400 Invalid page supplied
// @Failure 404 Elastic not found
// @router /getelasticstandard [get]
// 进入全文搜索主页面
func (c *StandardController) GetElasticStandard() {
	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		c.TplName = "standard/index_standard.html"
	} else {
		c.TplName = "standard/index_standard.html"
	}
}

// @Title get search
// @Description get earch
// @Param q query string false "The query="
// @Param a formData string false "The after..."
// @Success 200 {object} models.GetSearch
// @Failure 400 Invalid page supplied
// @Failure 404 Search not found
// @router /elasticsearch [get]
// Search returns results matching a query, paginated by after.
func (c *StandardController) ElasticSearch() { //(*SearchResults, error)
	// 3. Search for the indexed documents
	// Build the request body.
	query := c.GetString("q")
	after := c.GetStrings("a")
	//beego.Info(after)
	//if query == "" {
	//	query = "目标"
	//}
	// *************
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	var results SearchResults
	//Search returns results matching a query, paginated by after.
	res, err := es.Search(
		es.Search.WithIndex(indexName),
		es.Search.WithBody(buildQuery(query, after...)),
	)

	defer res.Body.Close()
	//log.Printf(res.String())

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			//return //&results, err
			c.Data["json"] = map[string]interface{}{"status": 1, "info": "ERR"}
			c.ServeJSON()
		}
		//return //&results, fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
		c.Data["json"] = map[string]interface{}{"status": 2, "info": "ERR"}
		c.ServeJSON()
	}

	type envelopeResponse struct {
		Took int
		Hits struct {
			Total struct {
				Value int
			}
			Hits []struct {
				ID         string          `json:"_id"`
				Source     json.RawMessage `json:"_source"`
				Highlights json.RawMessage `json:"highlight"`
				Sort       []interface{}   `json:"sort"`
			}
		}
	}

	var r envelopeResponse
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		//return //&results, err
		c.Data["json"] = map[string]interface{}{"status": 3, "info": "ERR"}
		c.ServeJSON()
	}

	results.Total = r.Hits.Total.Value

	if len(r.Hits.Hits) < 1 {
		results.Hits = []*Hit{}
		//return //&results, nil
		c.Data["json"] = &results //map[string]interface{}{"status": 1, "info": "SUCCESS", "id": aid}
		c.ServeJSON()
	}

	for _, hit := range r.Hits.Hits {
		var h Hit
		h.ID = hit.ID
		h.Sort = hit.Sort
		//如果是office文件，则/officeview/id，如果是pdf文件，则/pdf?id=
		//根据id（attatchment id）判断附件类型
		//h.URL = strings.Join([]string{baseURL, h.ID, ""}, "/")
		// h.URL = strings.Join([]string{baseURL, h.ID}, "?id=")
		h.URL = strings.Join([]string{baseURL, h.ID, "&keyword=" + query}, "")

		if err := json.Unmarshal(hit.Source, &h); err != nil {
			//return //&results, err
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"status": 5, "info": "ERR"}
			c.ServeJSON()
		}

		// 对body字段进行缩减
		//myString.substr(start, 318)
		//beego.Info(h.Body)
		h.Body = SubString(h.Body, 0, 300)

		if len(hit.Highlights) > 0 {
			if err := json.Unmarshal(hit.Highlights, &h.Highlights); err != nil {
				//return //&results, err
				c.Data["json"] = map[string]interface{}{"status": 6, "info": "ERR"}
				c.ServeJSON()
			}
		}

		results.Hits = append(results.Hits, &h)
	}
	// beego.Info(&results)
	//return &results, nil
	c.Data["json"] = &results //map[string]interface{}{"status": 1, "info": "SUCCESS", "id": aid}
	c.ServeJSON()
}

// @Title get search
// @Description get earch
// @Param keyword query string false "The query="
// @Param search_after formData string false "paginated by earchpage..."
// @Success 200 {object} models.GetSearch
// @Failure 400 Invalid page supplied
// @Failure 404 Search not found
// @router /wxelasticsearch [get]
// Search returns results matching a query, paginated by after.
func (c *StandardController) WxElasticSearch() { //(*SearchResults, error)
	// 3. Search for the indexed documents
	// Build the request body.
	query := c.GetString("keyword")
	after := c.GetStrings("search_after")
	//beego.Info(after)
	//if query == "" {
	//	query = "目标"
	//}
	// *************
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	var results SearchResults
	//Search returns results matching a query, paginated by after.
	res, err := es.Search(
		es.Search.WithIndex(indexName),
		es.Search.WithBody(buildQuery(query, after...)),
	)

	defer res.Body.Close()
	//log.Printf(res.String())

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			//return //&results, err
			c.Data["json"] = map[string]interface{}{"status": 1, "info": "ERR"}
			c.ServeJSON()
		}
		//return //&results, fmt.Errorf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
		c.Data["json"] = map[string]interface{}{"status": 2, "info": "ERR"}
		c.ServeJSON()
	}

	type envelopeResponse struct {
		Took int
		Hits struct {
			Total struct {
				Value int
			}
			Hits []struct {
				ID         string          `json:"_id"`
				Source     json.RawMessage `json:"_source"`
				Highlights json.RawMessage `json:"highlight"`
				Sort       []interface{}   `json:"sort"`
			}
		}
	}

	var r envelopeResponse
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		//return //&results, err
		c.Data["json"] = map[string]interface{}{"status": 3, "info": "ERR"}
		c.ServeJSON()
	}

	results.Total = r.Hits.Total.Value

	if len(r.Hits.Hits) < 1 {
		// results.Hits = []*Hit{}
		//return //&results, nil
		c.Data["json"] = []*Hit{} //map[string]interface{}{"status": 1, "info": "SUCCESS", "id": aid}
		c.ServeJSON()
	}

	wxsite, err := web.AppConfig.String("wxreqeustsite")
	if err != nil {
		logs.Error(err)
	}
	var wxsearchresults []*Hit
	for _, hit := range r.Hits.Hits {
		var h Hit
		h.ID = hit.ID
		h.Sort = hit.Sort
		//如果是office文件，则/officeview/id，如果是pdf文件，则/pdf?id=
		//根据id（attatchment id）判断附件类型
		//h.URL = strings.Join([]string{baseURL, h.ID, ""}, "/")
		// h.URL = strings.Join([]string{baseURL, h.ID}, "?id=")
		h.URL = strings.Join([]string{baseURL, h.ID, "&keyword=" + query}, "")

		if err := json.Unmarshal(hit.Source, &h); err != nil {
			//return //&results, err
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"status": 5, "info": "ERR"}
			c.ServeJSON()
		}

		// 对body字段进行缩减
		//myString.substr(start, 318)
		//beego.Info(h.Body)
		h.Body = SubString(h.Body, 0, 300)

		if len(hit.Highlights) > 0 {
			if err := json.Unmarshal(hit.Highlights, &h.Highlights); err != nil {
				//return //&results, err
				c.Data["json"] = map[string]interface{}{"status": 6, "info": "ERR"}
				c.ServeJSON()
			}
		}

		h.ImageURL = wxsite + h.ImageURL

		// results.Hits = append(results.Hits, &h)
		wxsearchresults = append(wxsearchresults, &h)
	}
	// beego.Info(&results)
	//return &results, nil
	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "searchers": &wxsearchresults}
	// c.Data["json"] = //map[string]interface{}{"status": 1, "info": "SUCCESS", "id": aid}
	c.ServeJSON()
}

// 后端分页的数据结构
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
// 根据用户输入的关键字，查询规范
// 搜索规范或者图集的名称或编号
// 20170704：linumber没有用。因为用category+编号+年份比较好
func (c *StandardController) Search() { //search用的是post方法
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

	searchText := c.GetString("searchText")
	if searchText == "" {
		c.Data["json"] = map[string]interface{}{"code": "OK", "msg": "关键字为空"}
		c.ServeJSON()
		return
	}
	if searchText == "allstandard" {
		standards, err := models.GetAllStandards()
		if err != nil {
			logs.Error(err.Error)
		} else {
			c.Data["json"] = standards
			c.ServeJSON()
		}
	} else {
		//搜索名称
		// Results1, err := models.SearchStandardsName(searchText, false)
		// if err != nil {
		// 	logs.Error(err.Error)
		// }
		//搜索编号
		// Results2, err := models.SearchStandardsNumber(searchText, false)
		// if err != nil {
		// 	logs.Error(err.Error)
		// }
		// Results1 = append(Results1, Results2...)
		Results1, err := models.GetUserStandard(limit1, offset, searchText)
		if err != nil {
			logs.Error(err.Error)
		}
		//由categoryid查categoryname
		aa := make([]Standardmore, len(Results1))
		//由standardnumber查librarynumber
		for i, v := range Results1 {
			//由standardnumber正则得到编号50268和分类GB
			Category, _, Number := SplitStandardFileNumber(v.Number)
			// logs.Info(Category)
			// logs.Info(Number)
			//由分类和编号查有效版本库中的编号
			library, err := models.SearchLibraryNumber(Category, Number)
			if err != nil {
				logs.Error(err.Error)
			}
			// logs.Info(library)
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
			logs.Error(err)
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
	// 	logs.Error(err.Error)
	// } else {
	// 	c.Data["Standards"] = standards
	// 	c.Data["Length"] = len(standards) //得到总记录数
	// }
}

// @Title get standardpdf
// @Description get standardpdf
// @Success 200 {object} models.Create
// @Failure 400 Invalid page supplied
// @Failure 404 cart not found
// @router /standardpdf [get]
// 手机端使用的下载方式
func (c *StandardController) StandardPdf() { //search用的是post方法
	_, _, _, isadmin, islogin := checkprodRole(c.Ctx)
	if !isadmin && !islogin {
		// beego.Info(!islogin)
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
	}
	pdflink := c.GetString("file") // 这个只是一个route路径
	// logs.Info(path.Base(pdflink))// 1
	id := path.Base(pdflink)
	// id := c.Ctx.Input.Param(":id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	//根据id取得规范的路径
	standard, err := models.GetStandard(idNum)
	if err != nil {
		logs.Error(err)
	}
	// logs.Info(standard.Route)// /attachment/standard/GB/GBT 555-2008水库除险加固计算案例.pdf
	fileurl := strings.Replace(standard.Route, "/attachment/", "attachment/", -1)
	// http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, standard.Route)
	// logs.Info(path.Base(standard.Route)) // GBT 555-2008水库除险加固计算案例.pdf
	filename := path.Base(fileurl)

	// c.Data["PdfLink"] = pdflink
	// id := c.Ctx.Input.Param(":id")
	// c.Data["Id"] = id
	c.Data["fileName"] = filename
	c.TplName = "pdf/web/viewer.html" //?file=\"" + pdflink + "\""
}

// @Title dowload standardpdf
// @Description get standardpdf by id
// @Param id path string  true "The id of standardpdf"
// @Success 200 {object} models.GetAttachbyId
// @Failure 400 Invalid page supplied
// @Failure 404 pdf not found
// @router /downloadstandard/:id [get]
// 文件流方式下载pdf规范
func (c *StandardController) DownloadStandard() {
	// id := c.GetString("id")
	id := c.Ctx.Input.Param(":id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	//根据id取得规范的路径
	standard, err := models.GetStandard(idNum)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(standard.Route)
	fileurl := strings.Replace(standard.Route, "/attachment/", "attachment/", -1)
	// http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, standard.Route)

	filename := path.Base(fileurl)
	fileext := path.Ext(filename)
	matched, err := regexp.MatchString("\\.*[m|M][c|C][d|D]", fileext)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(matched)
	if matched {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "不能下载mcd文件!"}
		c.ServeJSON()
		return
	}

	filepath := strings.Replace(standard.Route, "/attachment/", "./attachment/", -1)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filepath)

	// body, err := ioutil.ReadFile(filepath)
	// if err != nil {
	// 	logs.Error(err)
	// }

	// var rw = c.Ctx.ResponseWriter
	// //允许访问所有域
	// rw.Header().Set("Access-Control-Allow-Origin", "*")

	// rw.Header().Set(`content-type`, `application/octet-stream;charset=utf-8`) // 指明response的返回对象是文件流
	// rw.Header().Set("Content-Disposition", "attachment;filename=\""+filename+"\"")
	// rw.Header().Set("Range", strconv.Itoa(65536*16))

	// // 下载的字节范围
	// var _, _, totalByte int
	// //if (c.Ctx.Request.Headerrequest != null && request.getHeader("range") != null) {
	// if c.Ctx.Request.Header.Get("range") != "" {

	// 	totalByte = len(body)

	// 	rw.Header().Set("Status Code", "206")
	// } else {
	// 	// 正常下载
	// 	// 文件总大小
	// 	totalByte = len(body)
	// 	// 下载起始位置
	// 	_ = 0
	// 	// 下载结束位置
	// 	_ = totalByte - 1
	// 	// 返回http状态
	// 	rw.Header().Set("Accept-Ranges", "bytes")
	// 	rw.Header().Set("Status Code", "200")
	// }

	// length := 65536 * 16
	// // 响应头
	// rw.Header().Set("Accept-Ranges", "bytes")

	// rw.Header().Set("length", strconv.Itoa(length))

	// rw.Write(body)
	// rw.Flush()
	// c.Ctx.Output.Download(fileurl)
	// ioutil.ReadFile(filename string) ([]byte, error)
	// 设置Content-Type
	// "Content-Disposition", "attachment;filename=" + URLEncoder.encode(file.getName(), "UTF-8")
	// 断点续传
	//range := c.Ctx.Request.Header.Get("range")//.replaceAll("[^0-9\\-]", "").split("-");
	// 文件总大小
	// 下载起始位置
	//startByte = Integer.parseInt(range[0]);
	// 下载结束位置
	//if (range.length > 1) {
	//    endByte = Integer.parseInt(range[1]);
	//} else {
	//    endByte = totalByte - 1;
	//}
	// 返回http状态
	// 需要下载字节数
	//int length = endByte - startByte + 1;

	// 需要下载字节数
	// int length = endByte - startByte + 1;
	// c.Ctx.Output.Body(body) //流stream的方式
	// b, _ := a.DrawToBytes(text, 32) //背景的大小
	// beego.Info(text)
	// w http.ResponseWriter, r *http.Request
	// io.Copy(c.Ctx.ResponseWriter, b) // stream实现了io.reader接口
	// 响应内容
	//bis.skip(startByte);
	//int len = 0;
	//byte[] buff = new byte[1024 * 64];
	//while ((len = bis.read(buff, 0, buff.length)) != -1) {
	//    if (length <= len) {
	//        bos.write(buff, 0, length);
	//        break;
	//    } else {
	//        length -= len;
	//        bos.write(buff, 0, len);
	//    }
	//}
	// var bs = bytes.NewBufferString("data:ceshi\n\n")
	//rw.Header().Set("Content-Range", "bytes "+startByte+"-"+endByte+"/"+totalByte)
}

// RangeSize sets the default range size to 1MB
const (
	RangeSize    int64 = 1024 * 1024
	ThreadAmount int   = 32
)

// RangeHeader defines the part of file to download.
type RangeHeader struct {
	StartPos int64
	EndPos   int64
}

func (h *RangeHeader) String() string {
	return fmt.Sprintf("bytes=%d-%d", h.StartPos, h.EndPos)
}

// Fetcher downloads file from URL.
type Fetcher struct {
	URL    string
	Pieces []RangeHeader
}

// retrieveAll downloads the file completely.
func (f *Fetcher) retrieveAll(w io.Writer) (int64, error) {
	resp, err := http.Get(f.URL)
	if err != nil {
		return 0, err
	}

	n, err := io.Copy(w, resp.Body)
	return n, err
}

// retrievePartial downloads part of the file.
func (f *Fetcher) retrievePartial(pieceN int, w io.WriterAt) (n int, err error) {
	if pieceN < 0 || pieceN >= len(f.Pieces) {
		return 0, errors.New("Unspported index")
	}
	s := f.Pieces[pieceN]

	// make HTTP Range request to get file from server
	req, err := http.NewRequest(http.MethodGet, f.URL, nil)
	if err != nil {
		return
	}
	req.Header.Set("Range", s.String())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// read data from response and write it
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	n, err = w.WriteAt(data, s.StartPos)
	return
}

func splitSize(length int64) (size int64) {
	// less than 1KB
	if length <= 1024 {
		return 1024
	}

	size = length / int64(ThreadAmount)
	if length%32 != 0 {
		size = size + 1
	}

	return
}

// @Title get wx standards list
// @Description get standards by page
// @Param keyword query string true "The keyword of standards"
// @Param searchpage query string true "The page for drawings list"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 standards not found
// @router /searchwxstandards [get]
// 小程序取得规范列表，分页_plus
func (c *StandardController) SearchWxStandards() {
	//search用的是post方法
	// wxsite := web.AppConfig.String("wxreqeustsite")
	limit := "5"
	limit1, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("searchpage")
	page1, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		logs.Error(err)
	}

	var offset int64
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	key := c.GetString("keyword")

	if key != "" {
		//搜索名称
		Results1, err := models.SearchStandardsNamePage(limit1, offset, key, false)
		if err != nil {
			logs.Error(err.Error)
		}
		//搜索编号
		Results2, err := models.SearchStandardsNumberPage(limit1, offset, key, false)
		if err != nil {
			logs.Error(err.Error)
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
			library, err := models.SearchLibraryNumber(Category, Number)
			if err != nil {
				logs.Error(err.Error)
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
// @Param id path string true "The id of standardpdf"
// @Param token query string true "The token of user"
// @Success 200 {object} models.GetAttachbyId
// @Failure 400 Invalid page supplied
// @Failure 404 pdf not found
// @router /wxstandardpdf/:id [get]
// 小程序下载规范
func (c *StandardController) WxStandardPdf() {
	// 20220908测试通过
	// token := c.GetString("token")
	// logs.Info(token)
	// username, err := utils.CheckToken(token)
	// if err != nil {
	// 	logs.Error(err)
	// 	c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "token错误！"}
	// 	return
	// } else if username == "" {
	// 	c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "用户未登录或未授权！"}
	// 	return
	// }
	// logs.Info(username)
	// id := c.GetString("id")
	id := c.Ctx.Input.Param(":id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	//根据id取得规范的路径
	standard, err := models.GetStandard(idNum)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(standard.Route)
	fileurl := strings.Replace(standard.Route, "/attachment/", "attachment/", -1)
	// http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, standard.Route)

	filename := path.Base(fileurl)
	fileext := path.Ext(filename)
	matched, err := regexp.MatchString("\\.*[m|M][c|C][d|D]", fileext)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(matched)
	if matched {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "不能下载mcd文件!"}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.Download(fileurl)
}

// 显示所有有效库
func (c *StandardController) Valid() {
	//search用的是post方法
	// name := c.GetString("name")
	c.Data["IsStandard"] = true //
	c.TplName = "standard.tpl"
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	// uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	// c.Data["Uname"] = uname
	//搜索名称
	valids, err := models.GetAllValids()
	if err != nil {
		logs.Error(err.Error)
	}
	c.Data["json"] = valids //这里必须要是c.Data["json"]，其他c.Data["Data"]不行
	c.ServeJSON()

	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/test.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "valid:")
	logs.Close()
}

// 删除有效库中选中
func (c *StandardController) DeleteValid() {
	//search用的是post方法
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
			logs.Error(err)
		}
		//循环删除成果
		//根据成果id取得所有附件
		err = models.DeleteValid(idNum)
		if err != nil {
			logs.Error(err)
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

// 上传excel文件，导入到规范数据库，用于批量导入规范文件
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
		logs.Error(err)
	}
	// var attachment string
	var path string
	// var filesize int64
	if h != nil {
		//保存附件
		// attachment = h.Filename
		// beego.Info(attachment)
		path = "./attachment/" + h.Filename
		// path := c.GetString("url")  //存文件的路径
		// path = path[3:]
		// path = "./attachment" + "/" + h.Filename
		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("excel", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			logs.Error(err)
		}
	}
	if err != nil {
		logs.Error(err)
	}
	var standard models.Standard
	//读出excel内容写入数据库
	// excelFileName := path                    //"/home/tealeg/foo.xlsx"
	xlFile, err := xlsx.OpenFile(path) //excelFileName
	if err != nil {
		logs.Error(err)
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
			// standard.Content = row.Cells[j+4].String()
			standard.Route = row.Cells[j+5].String()
			_, err = models.SaveStandard(standard)

			if err != nil {
				logs.Error(err)
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

// 上传excel文件，导入到有效版本数据库
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
		logs.Error(err)
	}
	// var attachment string
	var path string
	// var filesize int64
	if h != nil {
		//保存附件
		// attachment = h.Filename
		// beego.Info(attachment)
		path = "./attachment/" + h.Filename
		// path := c.GetString("url")  //存文件的路径
		// path = path[3:]
		// path = "./attachment" + "/" + h.Filename
		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("excel2", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			logs.Error(err)
		}
	}
	var library models.Library
	//读出excel内容写入数据库
	// excelFileName := path//"/home/tealeg/foo.xlsx"
	// xlFile, err := xlsx.OpenFile(path) //excelFileName
	f, err := excelize.OpenFile(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	// defer func() {
	// 	if err := f.Close(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()
	// func (f *File) GetActiveSheetIndex() int
	// func (f *File) GetSheetList() []string
	// func (f *File) GetSheetName(index int) string

	// if err != nil {
	// 	logs.Error(err)
	// } else {
	sheetindexint := f.GetActiveSheetIndex()
	sheetname := f.GetSheetName(sheetindexint)
	rows, err := f.GetRows(sheetname)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i, row := range rows {
		if i != 0 {
			// for j, colCell := range row {
			logs.Info(row)
			j := 0
			library.Number = row[j]
			library.Title = row[j+1]
			library.Category = row[j+2]
			library.LiNumber = row[j+3]
			library.Year = row[j+4]
			if len(row) > 5 {
				library.Execute = row[j+5]
			}
			library.Created = time.Now()
			library.Updated = time.Now()
			_, err = models.SaveLibrary(library)
			if err != nil {
				logs.Error(err)
			}
			// fmt.Print(colCell, "\t")
		}
	}
	// if err = rows.Close(); err != nil {
	// 	fmt.Println(err)
	// }
	// for _, sheet := range xlFile.Sheets {
	// 	for i, row := range sheet.Rows {
	// 		if i != 0 {
	// 			// 这里要判断单元格列数，如果超过单元格使用范围的列数，则出错for j := 2; j < 7; j += 5 {
	// 			j := 0
	// 			library.Number = row.Cells[j].String()
	// 			if err != nil {
	// 				logs.Error(err)
	// 			}
	// 			library.Title = row.Cells[j+1].String()
	// 			if err != nil {
	// 				logs.Error(err)
	// 			}
	// 			library.Category = row.Cells[j+2].String()
	// 			if err != nil {
	// 				logs.Error(err)
	// 			}
	// 			library.LiNumber = row.Cells[j+3].String()
	// 			if err != nil {
	// 				logs.Error(err)
	// 			}
	// 			library.Year = row.Cells[j+4].String()
	// 			if err != nil {
	// 				logs.Error(err)
	// 			}
	// 			library.Execute = row.Cells[j+5].String()
	// 			if err != nil {
	// 				logs.Error(err)
	// 			}
	// 			library.Created = time.Now()
	// 			library.Updated = time.Now()
	// 			_, err = models.SaveLibrary(library)
	// 			if err != nil {
	// 				logs.Error(err)
	// 			}
	// 		}
	// 		// for _, cell := range row.Cells {这里要继续循环cells，不能为空，即超出单元格使用范围
	// 		// 	fmt.Printf("%s\n", cell.String())
	// 		// }
	// 	}
	// }
	c.Data["json"] = map[string]interface{}{"state": "SUCCESS"}
	c.ServeJSON()
	// }

	// c.TplName = "standard.tpl"
	// c.Redirect("/standard", 302)
}

func (c *StandardController) Standard_one_addbaidu() {
	//一对一模式
	_, _, uid, _, _ := checkprodRole(c.Ctx)
	var standard models.Standard
	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		logs.Error(err)
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
				logs.Error(err)
			}
		}
		path = "./attachment/standard/" + category + "/" + h.Filename
		// path := c.GetString("url")  //存文件的路径
		// path = path[3:]
		// path = "./attachment" + "/" + h.Filename
		// f.Close()   // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("file", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			logs.Error(err)
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
		logs.Error(err)
	} else {
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "title": "111", "original": "demo.jpg", "url": standard.Route}
		c.ServeJSON()
	}
}

// @Title upload file to standard html
// @Description get upload file to standard html
// @Success 200 {object} models.Elastic
// @Failure 400 Invalid page supplied
// @Failure 404 Page not found
// @router /uploadstandard [get]
// 进入上传主页面
func (c *StandardController) UploadStandard() {
	_, _, _, _, isLogin := checkprodRole(c.Ctx)
	if !isLogin {
		// route := c.Ctx.Request.URL.String()
		// c.Data["Url"] = route
		// c.Redirect("/roleerr?url="+route, 302)
		c.Data["json"] = "未登陆"
		c.ServeJSON()
		return
	}
	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		c.TplName = "standard/upload_standard.tpl"
	} else {
		c.TplName = "standard/upload_standard.tpl"
	}
}

// @Title post bootstrapfileinput
// @Description post file by BootstrapFileInput
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 page not found
// @router /standard/upload [post]
// 上传
func (c *StandardController) Upload() {
	// 取得用户名
	_, _, uid, _, isLogin := checkprodRole(c.Ctx)
	if !isLogin {
		// route := c.Ctx.Request.URL.String()
		// c.Data["Url"] = route
		// c.Redirect("/roleerr?url="+route, 302)
		c.Data["json"] = "未登陆"
		c.ServeJSON()
		return
	}
	//获取上传的文件
	_, h, err := c.GetFile("input-ke-2[]")
	// beego.Info(h.Filename)自动将英文括号改成了_下划线
	if err != nil {
		// logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "获取文件错误！"}
		c.ServeJSON()
		return
	}
	fileSuffix := path.Ext(h.Filename)
	if fileSuffix != ".DOC" && fileSuffix != ".doc" && fileSuffix != ".DOCX" && fileSuffix != ".docx" && fileSuffix != ".pdf" && fileSuffix != ".PDF" {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "文件类型错误，请上传doc或pdf"}
		c.ServeJSON()
		return
	}

	category, categoryname, fileNumber, year, fileName, _ := SplitStandardName(h.Filename)
	var filepath, article_body string
	var standard models.Standard
	// var filesize int64
	if h == nil {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "文件为空！"}
		c.ServeJSON()
		return
	}
	//纯英文下没有取到汉字字符，所以没有名称
	if fileName == "" {
		fileName = fileNumber
	}

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
	sid, err := models.SaveStandard(standard)
	if err != nil {
		// logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "存入数据库错误！"}
		c.ServeJSON()
		return
	} else {
		// 如果文件存在，则返回
		if PathisExist(filepath) {
			c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "文件已存在！"}
			c.ServeJSON()
			return
		}

		//保存附件
		if category != "" {
			err := os.MkdirAll("./attachment/standard/"+category, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
			if err != nil {
				logs.Error(err)
			}
		}
		filepath = "./attachment/standard/" + category + "/" + h.Filename
		err = c.SaveToFile("input-ke-2[]", filepath) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			// logs.Error(err)
			c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "文件保存错误！"}
			c.ServeJSON()
			return
		}
		//filesize, _ = FileSize(filepath)
		//filesize = filesize / 1000.0
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "title": h.Filename, "original": h.Filename, "url": standard.Route}
		c.ServeJSON()
	}

	cwd, _ := os.Getwd()

	filepath = strings.Replace(filepath, "./attachment/", "/attachment/", -1)
	f, err := os.Open(cwd + filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// fmt.Println(f.Name())
	client := tika.NewClient(nil, "http://localhost:9998")
	body, err := client.Parse(context.Background(), f)
	// body, err := client.Detect(context.Background(), f) //application/pdf
	// fmt.Println(err)
	// fmt.Println(body)

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		log.Fatalln(err)
	}

	dom.Find("p").Each(func(i int, selection *goquery.Selection) {
		//if selection.Text() != " " || selection.Text() != "\n" {
		//	fmt.Println(selection.Text())
		text := strings.Replace(selection.Text(), "\n", "", -1)
		text = strings.Replace(text, " ", "", -1)
		article_body = article_body + text //selection.Text()
		//}
	})

	now := time.Now()
	year_2, month, day := now.Date()
	today_str := fmt.Sprintf("%04d-%02d-%02d", year_2, month, day)
	rand.Seed(time.Now().Unix())
	// 提取pdf第一页作为封面
	datapath := cwd + "/static/pdf/mutool.exe"
	// beego.Info(datapath)
	// workpath := "/attachment/standard/" + category + "/"

	// filename := "01"
	// inameexp := ".pdf"workpath + filename + onameexp
	onameexp := ".png"
	// mutool convert -O width=200 -F png -o output.png 01.pdf 1
	arg := []string{"convert", "-O", "width=300", "-F", "png", "-o", cwd + "/static/images/" + fileName + onameexp, cwd + filepath, "1"}
	fmt.Println("------------", arg)
	cmd := exec.Command(datapath, arg...)
	//记录开始时间
	// start := time.Now()

	err = cmd.Start()
	if err != nil {
		fmt.Printf("err: %v", err)
	}

	doc := &Document{
		//是productid还是attachmentid
		ID: strconv.FormatInt(sid, 10),
		// ImageURL:  "/static/images/" + strconv.Itoa(rand.Intn(8)) + "s.jpg", //1s.jpg
		ImageURL:  "/static/images/" + fileName + "1" + onameexp,
		Published: today_str, //fmt.Sprintf("%04d-%02d-%02d", jYear, jMonth, jDay),
		Title:     categoryname + " " + fileNumber + "-" + year + fileName,
		Body:      article_body,
	}

	err = Createitem(indexName, doc) //这个indexName是全局变量
	if err != nil {
		// log.Fatalln(err)
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "加入全文检索elastic错误！"}
		c.ServeJSON()
		return
	}
}

// 全文检索删除，根据id

// 全文检索更新
