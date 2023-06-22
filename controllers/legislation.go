package controllers

import (
	// "bufio"
	// beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	// "io"
	"github.com/3xxx/engineercms/models"
	"regexp"
	"strconv"
	"strings"
)

type LegislationController struct {
	web.Controller
}

type Legislationmore struct {
	Id            int64
	Number        string //`orm:"unique"`
	Title         string //原法规名称
	LibraryNumber string //规范有效版本库中的编号
	LibraryTitle  string
	Execute       string //执行时间
	Color         string
}

func (c *LegislationController) Index() {
	_, _, _, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["IsLegislation"] = true
	c.TplName = "legislation.tpl"

	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/test.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP())
	logs.Close()
}

// 搜索规范或者图集的名称或编号
func (c *LegislationController) Checklist() { //checklist用的是post方法
	logs2 := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/test.log"}`)
	logs.EnableFuncCallDepth(true)

	name := c.GetString("name")
	// beego.Info(name)
	array := strings.Split(name, "\n")
	aa := make([]Legislationmore, len(array))
	for i, v := range array {
		// beego.Info(v)
		//1、正则取到名称
		if v != "" { //空行的处理
			reg := regexp.MustCompile(`[《].*[》]`) //(`^\\<.*\\>`)
			text2 := reg.FindAllString(v, -1)
			// beego.Info(text2)
			if text2 != nil { //无书名号的处理。因为text2是数组，所以要用nil进行判断，而不能用""。
				text3 := SubString(text2[0], 1, len([]rune(text2[0]))-2)
				//2、根据名称搜索标准版本库，取得名称和版本号
				library, err := models.SearchLiabraryName(text3)
				if err != nil {
					logs.Error(err)
				}
				// logs.Info(len(library))
				text4 := strconv.Itoa(i + 1)
				Id1, err := strconv.ParseInt(text4, 10, 64)
				if err != nil {
					logs.Error(err)
				}
				aa[i].Id = Id1

				if len(library) != 0 { //如果找到，进行比对 library != nil这样不行，空数组不是nil
					logs.Info(library)
					//3、构造struct
					for j, w := range library {
						// logs.Info(w)
						if j == 0 {
							aa[i].LibraryNumber = w.Category + " " + w.Number + "-" + w.Year //规范有效版本库中的完整编号
							aa[i].LibraryTitle = w.Title
							aa[i].Execute = w.Execute //执行日期
						} else {
							aa[i].LibraryNumber = aa[i].LibraryNumber + "," + w.Category + " " + w.Number + "-" + w.Year //规范有效版本库中的完整编号
							aa[i].LibraryTitle = w.Title
							aa[i].Execute = aa[i].Execute + "," + w.Execute //执行日期
						}
					}
					// logs.Info(aa[i].LibraryNumber)
					// 比较LibraryNumber是否一致
					reg2 := regexp.MustCompile(`[（].*[）]`) //(`^\\<.*\\>`)
					LibyNum := reg2.FindAllString(v, -1)
					// logs.Info(LibyNum)
					if LibyNum != nil {
						LibyNum2 := SubString(LibyNum[0], 1, len([]rune(LibyNum[0]))-2)
						if aa[i].LibraryNumber != LibyNum2 {
							aa[i].Color = "red"
						}
					} else {
						aa[i].Color = "green"
					}
				} else { // 没找到
					// beego.Info(library)
					// aa[i].Number = library.Number //`orm:"unique"`
					// aa[i].Title = text3
					reg2 := regexp.MustCompile(`[（].*[）]`) //(`^\\<.*\\>`)
					LibyNum := reg2.FindAllString(v, -1)
					logs.Info(LibyNum)
					if LibyNum != nil {
						aa[i].LibraryNumber = SubString(LibyNum[0], 1, len([]rune(LibyNum[0]))-2)
					}
					// aa[i].LibraryNumber = "No LibraryNumber Match Find!"

					aa[i].LibraryTitle = text3
					aa[i].Execute = ""
					aa[i].Color = "blue"
					// logs.Info(c.Ctx.Input.IP() + " " + "No LibraryNumber:" + text3)
					// beego.Info(aa[i])
				}
			}
		}
	}
	c.Data["IsLegislation"] = true
	c.TplName = "legislation.tpl"
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	// uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	// c.Data["Uname"] = uname
	//逐行读取
	// br := bufio.NewReader(name)
	// for {
	// 	a, _, c := br.ReadLine()
	// 	if c == io.EOF {
	// 		break
	// 	}
	// 	beego.Info(string(a))
	// }

	// bfRd := bufio.NewReader(f)
	// for {
	// 	line, err := bfRd.ReadBytes('\n')
	// 	hookfn(line)    //放在错误处理前面，即使发生错误，也会处理已经读取到的数据。
	// 	if err != nil { //遇到任何错误立即返回，并忽略 EOF 错误信息
	// 		if err == io.EOF {
	// 			return nil
	// 		}
	// 		return err
	// 	}
	// }

	// buf := bufio.NewReader(f)
	// for {
	// 	line, err := buf.ReadString('\n')
	// 	line = strings.TrimSpace(line)
	// 	handler(line)
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			return nil
	// 		}
	// 		return err
	// 	}
	// }

	//由categoryid查categoryname
	// aa := make([]Legislationmore, len(Results1))
	// //由legislationnumber查librarynumber
	// for i, v := range Results1 {
	// 	//由userid查username
	// 	user := models.GetUserByUserId(v.Uid)
	// 	//由分类和编号查有效版本库中的编号
	// 	library, err := models.SearchLiabraryNumber(name, "Number")
	// 	if err != nil {
	// 		logs.Error(err.Error)
	// 	}
	// 	aa[i].Id = v.Id
	// 	aa[i].Number = v.Number //`orm:"unique"`
	// 	aa[i].Title = v.Title
	// 	aa[i].Uname = user.Username //换成用户名
	// 	aa[i].Route = v.Route
	// 	aa[i].Created = v.Created
	// 	aa[i].Updated = v.Updated
	// 	aa[i].Views = v.Views
	// 	if library != nil {
	// 		aa[i].LibraryNumber = library.Number //规范有效版本库中的编号
	// 		aa[i].LibraryTitle = library.Title
	// 		aa[i].LiNumber = library.LiNumber //完整编号
	// 	} else {
	// 		aa[i].LiNumber = "No LibraryNumber Match Find!"
	// 		aa[i].LibraryTitle = ""
	// 		aa[i].LibraryNumber = ""
	// 	}
	// }
	c.Data["json"] = aa //这里必须要是c.Data["json"]，其他c.Data["Data"]不行
	c.ServeJSON()

	// logs.Info(c.Ctx.Input.IP() + " " + "SearchLegislationsName:" + name)
	logs2.Close()
}

// 上传文档供解析-替换（增加）标准号
func (c *LegislationController) FileInput() { //
	c.Data["IsLegislation_upfile"] = true //
	c.TplName = "legislation_upfile.tpl"

	logs2 := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/test.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP())
	logs2.Close()
}
