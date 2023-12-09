package controllers

import (
	// "crypto/md5"
	// "encoding/hex"
	// "encoding/json"
	"github.com/3xxx/engineercms/models"
	// beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	// "github.com/beego/beego/v2/adapter/httplib"
	// "github.com/beego/beego/v2/adapter/logs"
	// "net"
	// "net/http"
	// "net/url"
	"path"
	"strconv"
	"strings"
	"time"
	// "io/ioutil"
	"github.com/PuerkitoBio/goquery"
	"io"
	// "log"

	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	// "github.com/unidoc/unioffice/schema/soo/wml"
)

// CMSWXDIARY API
type DiaryController struct {
	web.Controller
}

// 日志列表页
func (c *DiaryController) Get() {
	c.TplName = "wxdiaries.tpl"
}

// 日志页面
func (c *DiaryController) GetWxDiary2() {
	c.TplName = "wxdiary.tpl"
	var err error
	id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	wxsite, err := web.AppConfig.String("wxreqeustsite")
	if err != nil {
		logs.Error(err)
	}

	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(idNum)
	Diary, err := models.GetDiary(idNum)
	if err != nil {
		logs.Error(err)
	}

	content := strings.Replace(Diary.Content, "/attachment/", wxsite+"/attachment/", -1)
	type Duration int64
	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)
	// hours := 8
	// const lll = "2006-01-02 15:04"
	// diarytime := Diary.Updated.Add(time.Duration(hours) * time.Hour).Format(lll)
	wxDiary := &models.Diary{
		Id:        Diary.Id,
		Title:     Diary.Title,
		Diarydate: Diary.Diarydate,
		Content:   content, //Diary.Content,
		// LeassonType: 1,
		Views:   Diary.Views,
		Created: Diary.Created,
		// Updated:     diarytime,
	}
	c.Data["Diary"] = wxDiary
}

// @Title post wx diary by catalogId
// @Description post diary by projectid
// @Param projectid query string true "The projectid of diary"
// @Param title query string  true "The title of diary"
// @Param diarydate query string  true "The diarydate of diary"
// @Param diaryactivity query string  true "The diaryactivity of diary"
// @Param diaryweather query string  true "The diaryweather of diary"
// @Param content query string  true "The content of diary"
// @Param skey query string  true "The skey of user"
// @Success 200 {object} models.AddDiary
// @Failure 400 Invalid page supplied
// @Failure 404 Diary not found
// @router /addwxdiary [post]
// 向设代日记id下添加微信小程序文章_珠三角设代plus用_
func (c *DiaryController) AddWxDiary() {
	var user models.User
	var err error
	openID := c.GetSession("openID")
	// beego.Info(openID)
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			logs.Error(err)
		} else {
			// beego.Info(user)
			// pid := web.AppConfig.String("wxdiaryprojectid") //"26159"
			pid := c.GetString("projectid")
			title := c.GetString("title")
			diarydate := c.GetString("diarydate")

			array := strings.Split(diarydate, "-")
			//当月天数
			const base_format = "2006-01-02"
			year := array[0]
			month := array[1]
			if len(month) == 1 {
				month = "0" + month
			}
			day := array[2]
			if len(day) == 1 {
				day = "0" + day
			}
			// diarydate2, err := time.Parse(base_format, year+"-"+month+"-"+day)
			// if err != nil {
			// 	logs.Error(err)
			// }
			diarydate2 := year + "-" + month + "-" + day
			// beego.Info(diarydate2)
			diaryactivity := c.GetString("diaryactivity")
			diaryweather := c.GetString("diaryweather")

			content := c.GetString("content")
			content = "<p style='font-size: 16px;'>分部：" + diaryactivity + "；</p><p style='font-size: 16px;'>天气：" + diaryweather + "；</p><p style='font-size: 16px;'>记录：" + user.Nickname + "；</p>" + content //<span style="font-size: 18px;">这个字体到底是多大才好看</span>

			//id转成64为
			pidNum, err := strconv.ParseInt(pid, 10, 64)
			if err != nil {
				logs.Error(err)
			}
			//添加日志
			aid, err := models.AddDiary(title, content, diarydate2, pidNum, user.Id)
			if err != nil {
				logs.Error(err)
				c.Data["json"] = map[string]interface{}{"status": 0, "info": "ERR", "id": aid}
				c.ServeJSON()
			} else {
				c.Data["json"] = map[string]interface{}{"status": 1, "info": "SUCCESS", "id": aid}
				c.ServeJSON()
			}
		}
	} else {
		return
	}
	// //根据pid查出项目id
	// proj, err := models.GetProj(pidNum)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// parentidpath := strings.Replace(strings.Replace(proj.ParentIdPath, "#$", "-", -1), "$", "", -1)
	// parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
	// patharray := strings.Split(parentidpath1, "-")
	// topprojectid, err := strconv.ParseInt(patharray[0], 10, 64)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// code := time.Now().Format("2006-01-02 15:04")
	// code = strings.Replace(code, "-", "", -1)
	// code = strings.Replace(code, " ", "", -1)
	// code = strings.Replace(code, ":", "", -1)
	// //根据项目id添加成果code, title, label, principal, content string, projectid int64
	// Id, err := models.AddProduct(code, title, "wx", user.Nickname, user.Id, pidNum, topprojectid)
	// if err != nil {
	// 	logs.Error(err)
	// }
}

// @Title get wx diaries list
// @Description get diaries by page
// @Param projectid query string true "The projectid of diary"
// @Param page query string  true "The page for diaries list"
// @Param limit query string  true "The limit of page for diaries list"
// @Param skey query string  false "The skey for diary"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getwxdiaries [get]
// 小程序取得日志列表，分页_珠三角设代用
func (c *DiaryController) GetWxdiaries() {
	// id := c.Ctx.Input.Param(":id")
	// id := web.AppConfig.String("wxdiaryprojectid") //"26159" //25002珠三角设代日记id26159
	// wxsite := web.AppConfig.String("wxreqeustsite")
	var ProjectId int64
	var err error
	projectid := c.GetString("projectid")
	if projectid != "" {
		ProjectId, err = strconv.ParseInt(projectid, 10, 64)
		if err != nil {
			logs.Error(err)
		}
	}
	limit := c.GetString("limit")
	if limit == "" {
		limit = "12"
	}
	// limit1, err := strconv.ParseInt(limit, 10, 64)
	limit1, err := strconv.Atoi(limit)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("page")
	// page1, err := strconv.ParseInt(page, 10, 64)
	page1, err := strconv.Atoi(page)
	if err != nil {
		logs.Error(err)
	}

	// var idNum int64
	// //id转成64为
	// idNum, err = strconv.ParseInt(id, 10, 64)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// var offset int64
	var offset int
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	// diaries, err := models.GetWxDiaries(idNum, limit1, offset)
	diaries, err := models.GetWxDiaries2(ProjectId, limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(diaries)

	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "diaries": diaries}
	c.ServeJSON()
}

// @Title get wx diaries list
// @Description get diaries by page
// @Param id path string  true "The id of diaries"
// @Param page query string  true "The page for diaries list"
// @Param limit query string  true "The limit of page for diaries list"
// @Param skey query string  false "The skey for diary"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getwxdiaries2/:id [get]
// 网页页面取得日志列表，返回page rows total
func (c *DiaryController) GetWxdiaries2() {
	id := c.Ctx.Input.Param(":id")
	// id := web.AppConfig.String("wxdiaryprojectid") //"26159" //25002珠三角设代日记id26159
	// wxsite := web.AppConfig.String("wxreqeustsite")
	// limit := "10"
	limit := c.GetString("limit")

	limit1, err := strconv.Atoi(limit)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("page")
	// page1, err := strconv.ParseInt(page, 10, 64)
	page1, err := strconv.Atoi(page)
	if err != nil {
		logs.Error(err)
	}

	var idNum int64
	//id转成64为
	idNum, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	// var offset int64
	var offset int
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	// diaries, err := models.GetWxDiaries(idNum, limit1, offset)
	diaries, err := models.GetWxDiaries2(idNum, limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	count, err := models.GetWxDiaryCount(idNum)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(diaries)
	c.Data["json"] = map[string]interface{}{"page": page, "rows": diaries, "total": count}
	c.ServeJSON()
}

// @Title get wx diary by diaryId
// @Description get diary by diaryid
// @Param id path string  true "The id of diary"
// @Param skey path string  true "The skey of user"
// @Success 200 {object} models.GetDiary
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /getwxdiary/:id [get]
// 根据id查看一篇微信文章
func (c *DiaryController) GetWxDiary() {
	var err error
	id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	wxsite, err := web.AppConfig.String("wxreqeustsite")
	if err != nil {
		logs.Error(err)
	}

	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(idNum)
	Diary, err := models.GetDiary(idNum)
	if err != nil {
		logs.Error(err)
	}

	content := strings.Replace(Diary.Content, "/attachment/", wxsite+"/attachment/", -1)
	type Duration int64
	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)
	// hours := 8
	// const lll = "2006-01-02 15:04"
	// diarytime := Diary.Updated.Add(time.Duration(hours) * time.Hour).Format(lll)
	wxDiary := &models.Diary{
		Id:        Diary.Id,
		Title:     Diary.Title,
		Diarydate: Diary.Diarydate,
		Content:   content, //Diary.Content,
		// LeassonType: 1,
		Views:   Diary.Views,
		Created: Diary.Created,
		// Updated:     diarytime,
	}
	c.Data["json"] = wxDiary
	c.ServeJSON()
}

// @Title post wx diary by diaryid
// @Description post diary by diaryid
// @Param id query string true "The id of diary"
// @Param title query string true "The title of diary"
// @Param content query string true "The content of diary"
// @Success 200 {object} models.AddDiary
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /updatewxdiary [post]
// 编辑设代日记id下微信小程序文章_珠三角设代plus用_editor方式
func (c *DiaryController) UpdateWxDiary() {
	// pid := web.AppConfig.String("wxcatalogid") //"26159"
	//hotqinsessionid携带过来后，用下面的方法获取用户登录存储在服务端的session
	openid := c.GetSession("openID")
	if openid == nil {
		return
	}

	id := c.GetString("id")
	title := c.GetString("title")
	content := c.GetString("content")
	//将content中的http://ip/去掉
	wxsite, err := web.AppConfig.String("wxreqeustsite")
	if err != nil {
		logs.Error(err)
	}
	content = strings.Replace(content, wxsite, "", -1)
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}

	//取得文章
	_, err = models.GetDiary(idNum)
	if err != nil {
		logs.Error(err)
	} else {
		//更新文章
		err = models.UpdateDiary(idNum, title, content)
		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"info": "ERR", "id": id}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "id": id}
			c.ServeJSON()
		}
	}
}

// @Title post wx diary by diaryId
// @Description post diary by catalogid
// @Param id query string  true "The id of diary"
// @Success 200 {object} models.Adddiary
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /deletewxdiary [post]
// 根据id删除_没删除文章中的图片
func (c *DiaryController) DeleteWxDiary() {
	var openID string
	openid := c.GetSession("openID")
	// beego.Info(openid.(string))
	if openid == nil {
		c.Data["json"] = "没有注册，未查到openid"
		c.ServeJSON()
	} else {
		openID = openid.(string)
		user, err := models.GetUserByOpenID(openID)
		if err != nil {
			logs.Error(err)
			c.Data["json"] = "未查到openid对应的用户"
			c.ServeJSON()
		} else {
			//判断是否具备admin角色
			role, err := models.GetRoleByRolename("admin")
			if err != nil {
				logs.Error(err)
			}
			uid := strconv.FormatInt(user.Id, 10)
			roleid := strconv.FormatInt(role.Id, 10)
			hasrole, err := e.HasRoleForUser(uid, "role_"+roleid)
			if err != nil {
				logs.Error(err)
			}
			if hasrole {
				id := c.GetString("id")
				//id转成64为
				idNum, err := strconv.ParseInt(id, 10, 64)
				if err != nil {
					logs.Error(err)
				}
				err = models.DeleteDiary(idNum)
				if err != nil {
					logs.Error(err)
					c.Data["json"] = "delete wrong"
					c.ServeJSON()
				} else {
					c.Data["json"] = "ok"
					c.ServeJSON()
				}
			} else {
				c.Data["json"] = "不是管理员"
				c.ServeJSON()
			}
		}
	}
}

type DiaryContent struct {
	Txt  string
	Html string
}

// @Title get wx diaries to doc
// @Description get diaries to doc
// @Param projectid query string true "The projectid of diary"
// @Param page query string  true "The page for diaries list"
// @Param limit query string  true "The limit of page for diaries list"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getwxdiaries [get]
// 将日志导出到word
func (c *DiaryController) HtmlToDoc() {
	// id := web.AppConfig.String("wxdiaryprojectid") //"26159" //25002珠三角设代日记id26159
	var ProjectId int64
	var err error
	projectid := c.GetString("projectid")
	if projectid != "" {
		ProjectId, err = strconv.ParseInt(projectid, 10, 64)
		if err != nil {
			logs.Error(err)
		}
	}
	// limit := "10"
	limit := c.GetString("limit")
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

	// diaries, err := models.GetWxDiaries(idNum, limit1, offset)
	diaries, err := models.GetWxDiaries2(ProjectId, limit1, offset)
	if err != nil {
		logs.Error(err)
	}

	doc := document.New()

	for _, v := range diaries {
		did := v.Diary.Id
		// wxsite := web.AppConfig.String("wxreqeustsite")

		Diary, err := models.GetDiary(did)
		if err != nil {
			logs.Error(err)
		}
		para := doc.AddParagraph()
		run := para.AddRun()
		para.SetStyle("Title")
		run.AddText(Diary.Title)

		para = doc.AddParagraph()
		para.SetStyle("Heading1")
		run = para.AddRun()
		run.AddText(Diary.Diarydate)

		//将一篇日志分段，通过<p标签
		slice1 := make([]DiaryContent, 0)

		var r io.Reader = strings.NewReader(string(Diary.Content))
		goquerydoc, err := goquery.NewDocumentFromReader(r)
		if err != nil {
			logs.Error(err)
		}

		goquerydoc.Find("p").Each(func(i int, s *goquery.Selection) {
			sel, _ := s.Html()
			bb := make([]DiaryContent, 1)
			bb[0].Html = sel
			txt := s.Text()
			bb[0].Txt = txt
			slice1 = append(slice1, bb...)
		})

		for _, w := range slice1 {
			//在每段里查找img标签
			// beego.Info(w)
			var r2 io.Reader = strings.NewReader(w.Html)
			goquerydoc2, err := goquery.NewDocumentFromReader(r2)
			if err != nil {
				logs.Error(err)
			}
			slice2 := make([]Img, 0)
			goquerydoc2.Find("img").Each(func(i int, s2 *goquery.Selection) {
				sel2, _ := s2.Attr("src")
				// beego.Info(sel2)
				aa := make([]Img, 1)
				sel3 := strings.Replace(sel2, "/attachment/", "attachment/", -1)
				aa[0].Src = sel3
				aa[0].Name = path.Base(sel2)
				slice2 = append(slice2, aa...)
			})

			para = doc.AddParagraph()
			para.Properties().SetFirstLineIndent(0.354331 * measurement.Inch)
			run = para.AddRun()
			run.AddText(w.Txt)

			if len(slice2) > 0 {
				for _, x := range slice2 {
					img1, err := common.ImageFromFile(x.Src)
					if err != nil {
						logs.Error("unable to create image: %s", err)
					}
					img1ref, err := doc.AddImage(img1)
					if err != nil {
						logs.Error("unable to add image to document: %s", err)
					}
					para = doc.AddParagraph()
					run = para.AddRun()

					inl, err := run.AddDrawingInline(img1ref)
					if err != nil {
						logs.Error("unable to add inline image: %s", err)
					}
					inl.SetSize(5.5*measurement.Inch, 5.5*measurement.Inch)
				}
			}
		}

		// wxDiary := &models.Diary{
		// 	Id:        Diary.Id,
		// 	Title:     Diary.Title,
		// 	Diarydate: Diary.Diarydate,
		// 	Content:   content,
		// 	Views:     Diary.Views,
		// 	Created:   Diary.Created,
		// }

		// img2data, err := ioutil.ReadFile("gophercolor.png")
		// if err != nil {
		// 	log.Fatalf("unable to read file: %s", err)
		// }
		// img2, err := common.ImageFromBytes(img2data)
		// if err != nil {
		// 	log.Fatalf("unable to create image: %s", err)
		// }
		// img2ref, err := doc.AddImage(img2)
		// if err != nil {
		// 	log.Fatalf("unable to add image to document: %s", err)
		// }

		// para := doc.AddParagraph()

		// anchored, err := para.AddRun().AddDrawingAnchored(img1ref)
		// if err != nil {
		// 	log.Fatalf("unable to add anchored image: %s", err)
		// }
		// anchored.SetName("Gopher")
		// anchored.SetSize(2*measurement.Inch, 2*measurement.Inch)
		// anchored.SetOrigin(wml.WdST_RelFromHPage, wml.WdST_RelFromVTopMargin)
		// anchored.SetHAlignment(wml.WdST_AlignHCenter)
		// anchored.SetYOffset(3 * measurement.Inch)
		// anchored.SetTextWrapSquare(wml.WdST_WrapTextBothSides)

	}
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + ".docx"
	doc.SaveToFile("static/" + newname)
	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "filename": newname}
	c.ServeJSON()
}

func createParaRun(doc *document.Document, s string) document.Run {
	para := doc.AddParagraph()
	run := para.AddRun()
	run.AddText(s)
	return run
}

// golang的log.Fatal()和panic()函数的区别

// 在讲两者区别之前我们先看一下os.Exit()函数的定义：

// func Exit(code int)

// Exit causes the current program to exit with the given status code.
// Conventionally, code zero indicates success, non-zero an error.
// The program terminates immediately; deferred functions are not run.
// 注意两点：

// 应用程序马上退出。
// defer函数不会执行。
// 再来看log.Fatal函数定义

// func Fatal(v ...interface{})

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
// 看源代码：go/src/log/log.go

// // Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
// func (l *Logger) Fatal(v ...interface{}) {
//     l.Output(2, fmt.Sprint(v...))
//     os.Exit(1)
// }
// 总结起来log.Fatal函数完成：

// 打印输出内容
// 退出应用程序
// defer函数不会执行
// 和os.Exit()相比多了第一步。

// 再来看内置函数panic()函数定义：

// // The panic built-in function stops normal execution of the current
// // goroutine. When a function F calls panic, normal execution of F stops
// // immediately. Any functions whose execution was deferred by F are run in
// // the usual way, and then F returns to its caller. To the caller G, the
// // invocation of F then behaves like a call to panic, terminating G's
// // execution and running any deferred functions. This continues until all
// // functions in the executing goroutine have stopped, in reverse order. At
// // that point, the program is terminated and the error condition is reported,
// // including the value of the argument to panic. This termination sequence
// // is called panicking and can be controlled by the built-in function
// // recover.
// func panic(v interface{})
// 注意几点：

// 函数立刻停止执行 (注意是函数本身，不是应用程序停止)
// defer函数被执行
// 返回给调用者(caller)
// 调用者函数假装也收到了一个panic函数，从而
// 4.1 立即停止执行当前函数
// 4.2 它defer函数被执行
// 4.3 返回给它的调用者(caller)
// ...(递归重复上述步骤，直到最上层函数)
// 应用程序停止。
// panic的行为
// 简单的总结panic()就有点类似java语言的exception的处理，因而panic的行为和java的exception处理行为就非常类似，行为结合catch，和final语句块的处理流程。

// 下面给几个例子：

// 例子1：log.Fatal

// package main

// import (
//     "log"
// )

// func foo() {
//     defer func () { log.Print("3333")} ()
//     log.Fatal("4444")
// }

// func main() {
//     log.Print("1111")
//     defer func () { log.Print("2222")} ()
//     foo()
//     log.Print("9999")
// }
// 运行结果：

// $ go build && ./main
// 2018/08/20 17:48:44 1111
// 2018/08/20 17:48:44 4444
// 可见defer函数的内容并没有被执行，程序在log.Fatal(...)处直接就退出了。

// 例子2：panic()函数

// package main

// import (
//     "log"
// )

// func foo() {
//     defer func () { log.Print("3333")} ()
//     panic("4444")
// }

// func main() {
//     log.Print("1111")
//     defer func () { log.Print("2222")} ()
//     foo()
//     log.Print("9999")
// }
// 运行结果：

// $ go build && ./main
// 2018/08/20 17:49:28 1111
// 2018/08/20 17:49:28 3333
// 2018/08/20 17:49:28 2222
// panic: 4444

// goroutine 1 [running]:
// main.foo()
//         /home/.../main.go:9 +0x55
// main.main()
//         /home/.../main.go:15 +0x82
// 可见所有的defer都被调用到了，函数根据父子调用关系所有的defer都被调用直到最上层。
// 当然如果其中某一层函数定义了recover()功能，那么panic会在那一层函数里面被截获，然后由recover()定义如何处理这个panic，是丢弃，还是向上再抛出。(是不是和exception的处理机制一模一样呢？)

// 作者：CodingCode
// 链接：https://www.jianshu.com/p/f85ecae6e7df
// 来源：简书
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
