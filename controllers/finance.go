package controllers

import (
	// "crypto/md5"
	// "encoding/hex"
	// "encoding/json"
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	// beego "github.com/beego/beego/v2/adapter"
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
	"log"

	"github.com/unidoc/unioffice/common"
	"github.com/unidoc/unioffice/document"
	"github.com/unidoc/unioffice/measurement"
	// "github.com/unidoc/unioffice/schema/soo/wml"
)

// CMSWXDIARY API
type FinanceController struct {
	web.Controller
}

//日志列表页
func (c *FinanceController) Get() {
	c.TplName = "wxfinance.tpl"
}

//财务登记页面——作废
func (c *FinanceController) GetWxFinance2() {
	c.TplName = "wxfinance.tpl"
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
	Finance, err := models.GetFinance(idNum)
	if err != nil {
		logs.Error(err)
	}

	content := strings.Replace(Finance.Content, "/attachment/", wxsite+"/attachment/", -1)
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
	// financetime := Finance.Updated.Add(time.Duration(hours) * time.Hour).Format(lll)
	wxFinance := &models.Finance{
		Id:          Finance.Id,
		Amount:      Finance.Amount,
		Financedate: Finance.Financedate,
		Content:     content, //Finance.Content,
		// LeassonType: 1,
		Views:   Finance.Views,
		Created: Finance.Created,
		// Updated:     financetime,
	}
	c.Data["Finance"] = wxFinance
}

// @Title post wx finance by catalogId
// @Description post finance by projectid
// @Param id path string  true "The projectid of finance"
// @Param amount query string true "The amount of finance"
// @Param radio query string true "The radio of finance"
// @Param radio2 query string true "The radio2 of finance"
// @Param financedate query string true "The financedate of finance"
// @Param financeactivity query string true "The financeactivity of finance"
// @Param content query string true "The content of finance"
// @Param skey query string false "The skey of user"
// @Success 200 {object} models.AddFinance
// @Failure 400 Invalid page supplied
// @Failure 404 Finance not found
// @router /addwxfinance/:id [post]
// 向设代日记id下添加微信小程序文章_珠三角设代plus用_
func (c *FinanceController) AddWxFinance() {
	var user models.User
	var err error
	openID := c.GetSession("openID")
	// beego.Info(openID)
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			logs.Error(err)
		} else {
			pid := c.Ctx.Input.Param(":id")
			pidNum, err := strconv.ParseInt(pid, 10, 64)
			if err != nil {
				logs.Error(err)
			}
			amount := c.GetString("amount")
			amountint, err := strconv.Atoi(amount)
			if err != nil {
				logs.Error(err)
			}
			radio := c.GetString("radio")
			if radio == "1" {
				amountint = 0 - amountint
			}
			var consider bool
			radio2 := c.GetString("radio2")
			if radio2 == "1" {
				consider = true
			}
			financedate := c.GetString("financedate")
			array := strings.Split(financedate, "-")
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
			financedate2 := year + "-" + month + "-" + day
			financeactivity := c.GetString("financeactivity")
			content := c.GetString("content")
			content = "<p style='font-size: 16px;'>分部：" + financeactivity + "；</p><p style='font-size: 16px;'>记录：" + user.Nickname + "；</p>" + content //<span style="font-size: 18px;">这个字体到底是多大才好看</span>

			aid, err := models.AddFinance(amountint, content, financedate2, pidNum, user.Id, consider)
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
		c.Data["json"] = map[string]interface{}{"status": 0, "info": "用户未登录", "id": 0}
		c.ServeJSON()
		// return
	}
}

// @Title get wx finance list
// @Description get finance by page
// @Param id path string  true "The projectid of finance"
// @Param page query string  true "The page for finance list"
// @Param limit query string  true "The limit of page for finance list"
// @Param skey query string  false "The skey for finance"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getwxfinancelist/:id [get]
//小程序取得日志列表，分页,通用
func (c *FinanceController) GetWxFinanceList() {
	id := c.Ctx.Input.Param(":id")
	// id := web.AppConfig.String("wxfinanceprojectid") //"26159" //25002珠三角设代日记id26159
	// wxsite := web.AppConfig.String("wxreqeustsite")
	limit := c.GetString("limit")
	if limit == "" {
		limit = "12"
	}
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

	// finance, err := models.GetWxFinance(idNum, limit1, offset)
	finance, err := models.GetWxFinance2(idNum, limit1, offset)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "finance": finance}
		c.ServeJSON()
	}
}

// @Title get wx finance list
// @Description get finance by page
// @Param id path string  true "The projectid of finance"
// @Param page query string  true "The page for finance list"
// @Param limit query string  true "The limit of page for finance list"
// @Param skey query string  false "The skey for finance"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getwxfinance2/:id [get]
//网页页面取得日志列表，返回page rows total
func (c *FinanceController) GetWxfinance2() {
	id := c.Ctx.Input.Param(":id")
	// id := web.AppConfig.String("wxfinanceprojectid") //"26159" //25002珠三角设代日记id26159
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

	// finance, err := models.GetWxFinance(idNum, limit1, offset)
	finance, err := models.GetWxFinance2(idNum, limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	count, err := models.GetWxFinanceCount(idNum)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"page": 1, "rows": finance, "total": 0}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"page": page, "rows": finance, "total": count}
		c.ServeJSON()
	}
}

type WxFinance struct {
	Id int64 `json:"id",form:"-"`
	// Title       string    `orm:"sie(20)"`
	Financedate string    `orm:"sie(20)"`
	Content     string    `json:"html",orm:"sie(5000)"`
	ProjectId   int64     `orm:"null"`
	UserId      int64     `orm:"null"`
	Amount      int       `json:"amount"`
	Consider    bool      `json:"consider"`
	IsArticleMe bool      `json:"isArticleMe"`
	Views       int64     `orm:"default(0)"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now_add;type(datetime)"`
}

// @Title get wx finance by financeId
// @Description get finance by financeid
// @Param id path string  true "The id of finance"
// @Param skey path string  true "The skey of user"
// @Success 200 {object} models.GetFinance
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /getwxfinance/:id [get]
//根据id查看一篇微信文章
func (c *FinanceController) GetWxFinance() {
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
	Finance, err := models.GetFinance(idNum)
	if err != nil {
		logs.Error(err)
	}

	content := strings.Replace(Finance.Content, "/attachment/", wxsite+"/attachment/", -1)
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
	// financetime := Finance.Updated.Add(time.Duration(hours) * time.Hour).Format(lll)
	// var openID string
	var user models.User
	var isArticleMe bool
	var wxFinance *WxFinance
	openid := c.GetSession("openID")
	if openid != nil {
		// openID = openid.(string)
		user, err = models.GetUserByOpenID(openid.(string))
		if err != nil {
			logs.Error(err)
		}
	}
	if Finance.UserId == user.Id { //20191122修改
		isArticleMe = true
	}
	wxFinance = &WxFinance{
		Id:          Finance.Id,
		Amount:      Finance.Amount,
		Financedate: Finance.Financedate,
		Content:     content, //Finance.Content,
		// LeassonType: 1,
		Consider:    Finance.Consider,
		IsArticleMe: isArticleMe,
		Views:       Finance.Views,
		Created:     Finance.Created,
		// Updated:     financetime,
	}
	c.Data["json"] = wxFinance
	c.ServeJSON()
}

// @Title post wx finance by financeid
// @Description post finance by financeid
// @Param id query string true "The id of finance"
// @Param amount query string true "The amount of finance"
// @Param radio query string true "The radio of finance"
// @Param radio2 query string true "The radio2 of finance"
// @Param financedate query string true "The financedate of finance"
// @Param content query string true "The content of finance"
// @Success 200 {object} models.AddFinance
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /updatewxfinance [post]
//编辑设代日记id下微信小程序文章_珠三角设代plus用_editor方式
func (c *FinanceController) UpdateWxFinance() {
	// pid := web.AppConfig.String("wxcatalogid") //"26159"
	//hotqinsessionid携带过来后，用下面的方法获取用户登录存储在服务端的session
	openid := c.GetSession("openID")
	if openid == nil {
		c.Data["json"] = map[string]interface{}{"info": "ERR", "id": 0, "data": "openid为空"}
		c.ServeJSON()
		// return
	}

	id := c.GetString("id")
	amount := c.GetString("amount")
	amountint, err := strconv.Atoi(amount)
	if err != nil {
		logs.Error(err)
	}
	radio := c.GetString("radio")
	if radio == "1" {
		amountint = 0 - amountint
	}
	var consider bool
	radio2 := c.GetString("radio2")
	if radio2 == "1" {
		consider = true
	}
	content := c.GetString("content")

	financedate := c.GetString("financedate")
	array := strings.Split(financedate, "-")
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
	financedate2 := year + "-" + month + "-" + day

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
	_, err = models.GetFinance(idNum)
	if err != nil {
		logs.Error(err)
	} else {
		//更新文章
		err = models.UpdateFinance(idNum, amountint, content, financedate2, consider)
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

// @Title post wx finance by financeId
// @Description post finance by catalogid
// @Param id query string  true "The id of finance"
// @Success 200 {object} models.Addfinance
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /deletewxfinance [post]
//根据id删除_没删除文章中的图片
func (c *FinanceController) DeleteWxFinance() {
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
				err = models.DeleteFinance(idNum)
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

type FinanceContent struct {
	Txt  string
	Html string
}

// 下面这个没更改
func (c *FinanceController) HtmlToDoc() {
	id, err := web.AppConfig.String("wxfinanceprojectid") //"26159" //25002珠三角设代日记id26159
	if err != nil {
		logs.Error(err)
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

	var idNum int64
	//id转成64为
	idNum, err = strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}

	var offset int
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	// finance, err := models.GetWxFinance(idNum, limit1, offset)
	finance, err := models.GetWxFinance2(idNum, limit1, offset)
	if err != nil {
		logs.Error(err)
	}

	doc := document.New()

	for _, v := range finance {
		did := v.Finance.Id
		// wxsite := web.AppConfig.String("wxreqeustsite")

		Finance, err := models.GetFinance(did)
		if err != nil {
			logs.Error(err)
		}
		para := doc.AddParagraph()
		run := para.AddRun()
		para.SetStyle("Title")
		run.AddText(strconv.Itoa(Finance.Amount))

		para = doc.AddParagraph()
		para.SetStyle("Heading1")
		run = para.AddRun()
		run.AddText(Finance.Financedate)

		//将一篇日志分段，通过<p标签
		slice1 := make([]FinanceContent, 0)

		var r io.Reader = strings.NewReader(string(Finance.Content))
		goquerydoc, err := goquery.NewDocumentFromReader(r)
		if err != nil {
			logs.Error(err)
		}

		goquerydoc.Find("p").Each(func(i int, s *goquery.Selection) {
			sel, _ := s.Html()
			bb := make([]FinanceContent, 1)
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
						log.Fatalf("unable to create image: %s", err)
					}
					img1ref, err := doc.AddImage(img1)
					if err != nil {
						log.Fatalf("unable to add image to document: %s", err)
					}
					para = doc.AddParagraph()
					run = para.AddRun()

					inl, err := run.AddDrawingInline(img1ref)
					if err != nil {
						log.Fatalf("unable to add inline image: %s", err)
					}
					inl.SetSize(5.5*measurement.Inch, 5.5*measurement.Inch)
				}
			}
		}

		// wxFinance := &models.Finance{
		// 	Id:        Finance.Id,
		// 	Title:     Finance.Title,
		// 	Financedate: Finance.Financedate,
		// 	Content:   content,
		// 	Views:     Finance.Views,
		// 	Created:   Finance.Created,
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
