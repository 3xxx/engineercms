// estimate
package controllers

import (
	"fmt"
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/xuri/excelize/v2"
	"html/template"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

// CMSexcel API
type EstimateController struct {
	web.Controller
}

// 后端分页的项目列表结构
type estimateProjectsTable struct {
	Rows  []models.EstimateProject `json:"rows"`
	Page  int                      `json:"page"`
	Total int                      `json:"total"` //string或int64都行！
}

// 后端分页的工程费用表结构
//
//	type EstimateCostTable struct {
//		Rows  []models.EstimateCost `json:"rows"`
//		Page  int                   `json:"page"`
//		Total int                   `json:"total"` //string或int64都行！
//	}
type EstimateCostArchiTable struct {
	Code  int                         `json:"code"`
	Count int                         `json:"count"`
	Data  []*models.EstimateCostArchi `json:"data"`
}

type EstimateCostElectTable struct {
	Code  int                         `json:"code"`
	Count int                         `json:"count"`
	Data  []*models.EstimateCostElect `json:"data"`
}

type EstimateCostMetalTable struct {
	Code  int                         `json:"code"`
	Count int                         `json:"count"`
	Data  []*models.EstimateCostMetal `json:"data"`
}

type EstimateCostTempTable struct {
	Code  int                        `json:"code"`
	Count int                        `json:"count"`
	Data  []*models.EstimateCostTemp `json:"data"`
}

// type CostNode struct {
// 	ID                  uint        `json:"id"`
// 	EstimateProjPhaseID uint        `json:"estimateprojphaseid"`
// 	ParentID            uint        `json:"parentid"`
// 	Number              string      `json:"number"`
// 	CostName            string      `json:"costname"`
// 	Unit                string      `json:"unit"`
// 	Quantity            float64     `json:"quantity"`
// 	UnitPrice           float64     `json:"unitprice"`
// 	Total               float64     `json:"total"`
// 	Children            []*CostNode `json:"children"`
// }

func getEsArchiTreeRecursive(list []*models.EstimateCostArchi, parentId uint) []*models.EstimateCostArchi {
	res := make([]*models.EstimateCostArchi, 0)
	for _, v := range list {
		if v.ParentID == parentId {
			v.Children = getEsArchiTreeRecursive(list, v.ID)
			res = append(res, v)
		}
	}
	return res
}

func getEsElectTreeRecursive(list []*models.EstimateCostElect, parentId uint) []*models.EstimateCostElect {
	res := make([]*models.EstimateCostElect, 0)
	for _, v := range list {
		if v.ParentID == parentId {
			v.Children = getEsElectTreeRecursive(list, v.ID)
			res = append(res, v)
		}
	}
	return res
}

func getEsMetalTreeRecursive(list []*models.EstimateCostMetal, parentId uint) []*models.EstimateCostMetal {
	res := make([]*models.EstimateCostMetal, 0)
	for _, v := range list {
		if v.ParentID == parentId {
			v.Children = getEsMetalTreeRecursive(list, v.ID)
			res = append(res, v)
		}
	}
	return res
}

func getEsTempTreeRecursive(list []*models.EstimateCostTemp, parentId uint) []*models.EstimateCostTemp {
	res := make([]*models.EstimateCostTemp, 0)
	for _, v := range list {
		if v.ParentID == parentId {
			v.Children = getEsTempTreeRecursive(list, v.ID)
			res = append(res, v)
		}
	}
	return res
}

// func Tree(node []*TreeList, pid int) []*TreeList {
//   res := make([]*TreeList, 0)
//   for _, v := range node {
//     if v.Pid == pid {
//       v.Children = Tree(node, v.Id)
//       res = append(res, v)
//     }
//   }
//   return res
// }

// @Title upload excel
// @Description get excel
// @Success 200 {object} models.GetexcelPage
// @Failure 400 Invalid page supplied
// @Failure 404 Page not found
// @router /uploadexcel [get]
// 上传页面
func (c *EstimateController) UploadExcel() {
	username, role, _, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	if !islogin {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "用户未登录！", "title": "", "original": ""}
		c.ServeJSON()
		return
	}
	c.TplName = "estimate/uploadexcel.tpl"
}

// @Title post bootstrapfileinput
// @Description post file by BootstrapFileInput
// @Param id path string true "The id of project"
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 page not found
// @router /uploadexcelestimate [post]
// 上传excel模板文件，解析模板文件的输入输出存到数据库中——作废！！
func (c *EstimateController) UploadExcelEstimate() {
	// 取得用户名
	username, _, uid, _, _ := checkprodRole(c.Ctx)
	//获取上传的文件
	_, h, err := c.GetFile("input-ke-2[]")
	// beego.Info(h.Filename)自动将英文括号改成了_下划线
	if err != nil {
		logs.Error(err)
	}
	// 解析txt
	fileext := path.Ext(h.Filename)
	// beego.Info(fileext)
	matched, err := regexp.MatchString(".*[x|X][l|L][s|S]", fileext) //xlsx和xlsm
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "正则匹配错误！"}
		c.ServeJSON()
		return
	}
	if !matched {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "文件类型错误！"}
		c.ServeJSON()
		return
	}

	// random_name
	// newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	// year, month, _ := time.Now().Date()
	err = os.MkdirAll("./attachment/estimate/"+username+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		logs.Error(err)
	}

	var filesize int64
	// var excelinputalias, excelinputvalue, exceloutputalias, exceloutputvalue, inputcomment, outputcomment string
	// var k int
	// var isinputcell, isoutputcell bool

	if h != nil {
		//保存附件——要防止重名覆盖！！！！先判断是否存在！！！
		// filepath := "./attachment/excel/" + username + "/" + h.Filename
		filepath := "./attachment/estimate/" + username + "/" + h.Filename
		Url := "/attachment/estimate/" + username + "/"
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
		filesize, _ = FileSize(h.Filename)
		filesize = filesize / 1000.0

		matched, err := regexp.MatchString("\\.[x|X][l|L][s|S]*", fileext)
		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "非excel文件！"}
			c.ServeJSON()
			return
		}
		if matched { //如果是temple文件，则解析，如果是pdf文件，则不解析
			var estimateProjID, estimatePhaseID uint
			f, err := excelize.OpenFile(filepath)
			if err != nil {
				logs.Error(err)
				c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "打开excel表格文件错误！"}
				c.ServeJSON()
				return
			}

			sheetlist := f.GetSheetList()
			for i, sheetname := range sheetlist {
				fmt.Print(sheetname, "\n")
				if strings.Contains(sheetname, "配置") { // sheetname == "配置"
					var number, name, profile, grade, phasename, information string
					var periodint int
					var totalinvestmentfloat, staticinvestmentfloat float64

					rows, err := f.GetRows(sheetlist[i])
					if err != nil {
						fmt.Println(err)
						return
					}
					for _, row := range rows {
						for j, colCell := range row {
							if strings.Contains(string(colCell), "项目编号") {
								fmt.Print(colCell, "\t")
								fmt.Print(row[j+1], "\n")
								number = row[j+1]
							} else if strings.Contains(string(colCell), "项目名称") {
								fmt.Print(colCell, "\t")
								fmt.Print(row[j+1], "\n")
								name = row[j+1]
							} else if strings.Contains(string(colCell), "项目简介") {
								fmt.Print(colCell, "\t")
								fmt.Print(row[j+1], "\n")
								profile = row[j+1]
							} else if strings.Contains(string(colCell), "工程等级") {
								fmt.Print(colCell, "\t")
								fmt.Print(row[j+1], "\n")
								grade = row[j+1]
							} else if strings.Contains(string(colCell), "施工总工期") {
								fmt.Print(colCell, "\t")
								fmt.Print(row[j+1], "\n")
								period := row[j+1]
								periodint, err = strconv.Atoi(period)
								if err != nil {
									logs.Error(err)
								}
							} else if strings.Contains(string(colCell), "阶段") {
								fmt.Print(colCell, "\t")
								fmt.Print(row[j+1], "\n")
								phasename = row[j+1]
							} else if strings.Contains(string(colCell), "价格水平") {
								fmt.Print(colCell, "\t")
								fmt.Print(row[j+1], "\n")
								information = row[j+1]
							} else if strings.Contains(string(colCell), "总投资") {
								fmt.Print(colCell, "\t")
								fmt.Print(row[j+1], "\n")
								totalinvestment := row[j+1]
								totalinvestmentfloat, err = strconv.ParseFloat(totalinvestment, 64)
								if err != nil {
									logs.Error(err)
								}
							} else if strings.Contains(string(colCell), "静态投资") {
								fmt.Print(colCell, "\t")
								fmt.Print(row[j+1], "\n")
								staticinvestment := row[j+1]
								staticinvestmentfloat, err = strconv.ParseFloat(staticinvestment, 64)
								if err != nil {
									logs.Error(err)
								}
							}
							// 地类
							// 阶段
							// 价格水平
						}
					}
					// 写入项目信息
					estimateProjID, err = models.AddEstimateProject(number, name, profile, grade, periodint, uid)
					if err != nil {
						logs.Error(err)
					}
					// 写入阶段信息
					estimatePhaseID, err = models.AddEstimatePhase(estimateProjID, phasename, information, totalinvestmentfloat, staticinvestmentfloat)
					if err != nil {
						logs.Error(err)
					}
				} else if strings.Contains(sheetname, "建筑") { // sheetname == "建筑"
					// var cost_name, unit string
					// var quantityfloat, unit_pricefloat, totalfloat float64,professionalID, secondaryID, tertiaryID, fourthID, fifthID,
					var parentID_first, parentID_second, parentID_third, parentID_fourth, parentID uint

					rows, err := f.GetRows(sheetlist[i])
					if err != nil {
						fmt.Println(err)
						return
					}
					for i, row := range rows {
						if i > 3 { // len(row[i])
							// for j, colCell := range row {
							// logs.Info(row)
							// for _, colCell := range row {
							// 	logs.Info(colCell)
							// }
							//  rex := regexp.MustCompile("([\u4e00-\u9fa5]{1,2})").MatchString(string(colCell)) // 括号中1~2个中文字符
							//  rex := regexp.MustCompile(`[\u4e00-\u9fa5]{1,2}`).MatchString(string(colCell)) // 中文1~2个字符
							//  rex := regexp.MustCompile(`\(\d\)`).MatchString(string(colCell)) // 括号中数字
							//  rex := regexp.MustCompile(`\d`).MatchString(string(colCell)) // 数字
							//  str := `"老干a爹a干老"` //要查找的字符串
							//  re := regexp.MustCompile("[\u4e00-\u9fa5]{1,}")
							//  "^[\u4E00-\u9FA5]{1,2}$"
							//  re := regexp.MustCompile(`^[\d\|[A-Z].*\]`)
							//  regexp.MustCompile("[").MatchString(string(r))
							// https://blog.csdn.net/qq_20432379/article/details/86506432 尽量用双引号！！！不然就报/u错误
							var partOne bool
							// logs.Info(partOne) // false
							number := row[0]
							costname := row[1]
							costname4 := SubString(costname, 0, 4)
							// logs.Info(costname4)
							//设定一个含有中文的字符串
							rex := regexp.MustCompile("^第.部分$")
							if rex.MatchString(string(number)) || rex.MatchString(string(costname4)) {
								partOne = true
							}

							// if strings.Contains(string(row[1]), "第一部分：建筑工程") && row[2] == "" && row[3] == "" {
							if partOne && row[2] == "" && row[3] == "" {
								// 一级
								// number := row[0]
								// costname := row[1]
								total := row[5]
								totalfloat, err := strconv.ParseFloat(total, 64)
								if err != nil {
									logs.Error(err)
								}
								// professionalID, err = models.AddEstimateProfessional(estimatePhaseID, professional_component, totalfloat)
								costID, err := models.AddEstimateCostArchi(estimatePhaseID, 0, number, costname, "", 0, 0, totalfloat)
								if err != nil {
									logs.Error(err)
								}
								parentID_first = costID
								parentID = costID
								// break \u4e00\u4e8c\u4e09\u56db\u4e94\u516d\u4e03\u516b\u4e5d\u5341
							} else if regexp.MustCompile(`^[一-龟]{1,2}$`).MatchString(string(row[0])) && row[2] == "" && row[3] == "" && row[4] == "" {
								// 二级
								// number := row[0]
								// costname := row[1]
								total := row[5]
								totalfloat, err := strconv.ParseFloat(total, 64)
								if err != nil {
									logs.Error(err)
								}
								// secondaryID, err = models.AddEstimateSecondary(professionalID, number, secondary_component, totalfloat)
								costID, err := models.AddEstimateCostArchi(0, parentID_first, number, costname, "", 0, 0, totalfloat)
								if err != nil {
									logs.Error(err)
								}
								parentID_second = costID
								parentID = costID
								// break
							} else if regexp.MustCompile(`^\([一-龟]{1,2}\)$`).MatchString(string(row[0])) && row[2] == "" && row[3] == "" && row[4] == "" {
								// 三级
								// number := row[0]
								// costname := row[1]
								total := row[5]
								totalfloat, err := strconv.ParseFloat(total, 64)
								if err != nil {
									logs.Error(err)
								}
								// tertiaryID, err = models.AddEstimateTertiary(secondaryID, parentID, number, tertiary_component, totalfloat)
								costID, err := models.AddEstimateCostArchi(0, parentID_second, number, costname, "", 0, 0, totalfloat)
								if err != nil {
									logs.Error(err)
								}
								parentID_third = costID
								parentID = costID
								// break
							} else if regexp.MustCompile(`^[1-9]\d*$`).MatchString(string(row[0])) && row[2] == "" && row[3] == "" && row[4] == "" {
								// 四级
								// number := row[0]
								// costname := row[1]
								total := row[5]
								totalfloat, err := strconv.ParseFloat(total, 64)
								if err != nil {
									logs.Error(err)
								}
								// fourthID, err = models.AddEstimateTertiary(tertiaryID, parentID, number, tertiary_component, totalfloat)
								costID, err := models.AddEstimateCostArchi(0, parentID_third, number, costname, "", 0, 0, totalfloat)
								if err != nil {
									logs.Error(err)
								}
								parentID_fourth = costID
								parentID = costID
								// break
							} else if regexp.MustCompile(`^\([1-9]\d*\)$`).MatchString(string(row[0])) && row[2] == "" && row[3] == "" && row[4] == "" {
								// 五级
								// number := row[0]
								// costname := row[1]
								total := row[5]
								totalfloat, err := strconv.ParseFloat(total, 64)
								if err != nil {
									logs.Error(err)
								}
								// fifthID, err = models.AddEstimateTertiary(fourthID, parentID, number, tertiary_component, totalfloat)
								costID, err := models.AddEstimateCostArchi(0, parentID_fourth, number, costname, "", 0, 0, totalfloat)
								if err != nil {
									logs.Error(err)
								}
								// parentID_fifth = costID
								parentID = costID
								// break
							} else if row[1] != "" && row[2] != "" && row[3] != "" && row[4] != "" && row[5] != "" {
								// number := row[0]
								// costname := row[1]
								unit := row[2]
								quantity := row[3]
								quantityfloat, err := strconv.ParseFloat(quantity, 64)
								if err != nil {
									logs.Error(err)
								}
								unitprice := row[4]
								unitpricefloat, err := strconv.ParseFloat(unitprice, 64)
								if err != nil {
									logs.Error(err)
								}
								total := row[5]
								totalfloat, err := strconv.ParseFloat(total, 64)
								if err != nil {
									logs.Error(err)
								}
								costID, err := models.AddEstimateCostArchi(0, parentID, number, costname, unit, quantityfloat, unitpricefloat, totalfloat)
								if err != nil {
									logs.Error(err)
								}
								logs.Info(costID)
							}
							// }
						}
					}
				} else if strings.Contains(sheetname, "机电") {
					var parentID_first, parentID_second, parentID_third, parentID_fourth, parentID uint
					rows, err := f.GetRows(sheetlist[i])
					if err != nil {
						logs.Error(err)
						return
					}
					for i, row := range rows {
						if i > 3 { // len(row[i])
							var partOne bool
							number := row[0]
							costname := row[1]
							costname4 := SubString(costname, 0, 4)
							//设定一个含有中文的字符串
							rex := regexp.MustCompile("^第.部分$")
							if rex.MatchString(string(number)) || rex.MatchString(string(costname4)) {
								partOne = true
							}

							if partOne && row[2] == "" && row[3] == "" {
								// 一级
								totalEquipment := row[6]
								totalEquipmentfloat, err := strconv.ParseFloat(totalEquipment, 64)
								if err != nil {
									logs.Error(err)
								}
								totalInstallation := row[7]
								totalInstallationfloat, err := strconv.ParseFloat(totalInstallation, 64)
								if err != nil {
									logs.Error(err)
								}
								costID, err := models.AddEstimateCostElect(estimatePhaseID, 0, number, costname, "", 0, 0, 0, totalEquipmentfloat, totalInstallationfloat)
								if err != nil {
									logs.Error(err)
								}
								parentID_first = costID
								parentID = costID
							} else if regexp.MustCompile(`^[一-龟]{1,2}$`).MatchString(string(row[0])) && row[2] == "" && row[3] == "" && row[4] == "" {
								// 二级
								totalEquipment := row[6]
								totalEquipmentfloat, err := strconv.ParseFloat(totalEquipment, 64)
								if err != nil {
									logs.Error(err)
								}
								totalInstallation := row[7]
								totalInstallationfloat, err := strconv.ParseFloat(totalInstallation, 64)
								if err != nil {
									logs.Error(err)
								}
								costID, err := models.AddEstimateCostElect(0, parentID_first, number, costname, "", 0, 0, 0, totalEquipmentfloat, totalInstallationfloat)
								if err != nil {
									logs.Error(err)
								}
								parentID_second = costID
								parentID = costID
							} else if regexp.MustCompile(`^\([一-龟]{1,2}\)$`).MatchString(string(row[0])) && row[2] == "" && row[3] == "" && row[4] == "" {
								// 三级
								totalEquipment := row[6]
								totalEquipmentfloat, err := strconv.ParseFloat(totalEquipment, 64)
								if err != nil {
									logs.Error(err)
								}
								totalInstallation := row[7]
								totalInstallationfloat, err := strconv.ParseFloat(totalInstallation, 64)
								if err != nil {
									logs.Error(err)
								}
								costID, err := models.AddEstimateCostElect(0, parentID_second, number, costname, "", 0, 0, 0, totalEquipmentfloat, totalInstallationfloat)
								if err != nil {
									logs.Error(err)
								}
								parentID_third = costID
								parentID = costID
							} else if regexp.MustCompile(`^[1-9]\d*$`).MatchString(string(row[0])) && row[2] == "" && row[3] == "" && row[4] == "" {
								// 四级
								totalEquipment := row[6]
								totalEquipmentfloat, err := strconv.ParseFloat(totalEquipment, 64)
								if err != nil {
									logs.Error(err)
								}
								totalInstallation := row[7]
								totalInstallationfloat, err := strconv.ParseFloat(totalInstallation, 64)
								if err != nil {
									logs.Error(err)
								}
								costID, err := models.AddEstimateCostElect(0, parentID_third, number, costname, "", 0, 0, 0, totalEquipmentfloat, totalInstallationfloat)
								if err != nil {
									logs.Error(err)
								}
								parentID_fourth = costID
								parentID = costID
							} else if regexp.MustCompile(`^\([1-9]\d*\)$`).MatchString(string(row[0])) && row[2] == "" && row[3] == "" && row[4] == "" {
								// 五级
								totalEquipment := row[6]
								totalEquipmentfloat, err := strconv.ParseFloat(totalEquipment, 64)
								if err != nil {
									logs.Error(err)
								}
								totalInstallation := row[7]
								totalInstallationfloat, err := strconv.ParseFloat(totalInstallation, 64)
								if err != nil {
									logs.Error(err)
								}
								costID, err := models.AddEstimateCostElect(0, parentID_fourth, number, costname, "", 0, 0, 0, totalEquipmentfloat, totalInstallationfloat)
								if err != nil {
									logs.Error(err)
								}
								parentID = costID
							} else if row[1] != "" && row[2] != "" && row[3] != "" && row[4] != "" && row[5] != "" {
								unit := row[2]
								quantity := row[3]
								quantityfloat, err := strconv.ParseFloat(quantity, 64)
								if err != nil {
									logs.Error(err)
								}
								unitpriceEquipment := row[4]
								unitpriceEquipmentfloat, err := strconv.ParseFloat(unitpriceEquipment, 64)
								if err != nil {
									logs.Error(err)
								}
								unitpriceInstallation := row[5]
								unitpriceInstallationfloat, err := strconv.ParseFloat(unitpriceInstallation, 64)
								if err != nil {
									logs.Error(err)
								}
								totalEquipment := row[6]
								totalEquipmentfloat, err := strconv.ParseFloat(totalEquipment, 64)
								if err != nil {
									logs.Error(err)
								}
								totalInstallation := row[7]
								totalInstallationfloat, err := strconv.ParseFloat(totalInstallation, 64)
								if err != nil {
									logs.Error(err)
								}
								costID, err := models.AddEstimateCostElect(0, parentID, number, costname, unit, quantityfloat, unitpriceEquipmentfloat, unitpriceInstallationfloat, totalEquipmentfloat, totalInstallationfloat)
								if err != nil {
									logs.Error(err)
								}
								logs.Info(costID)
							}
						}
					}
				} else if strings.Contains(sheetname, "金结") ||
					strings.Contains(sheetname, "金属") {
					var parentID_first, parentID_second, parentID_third, parentID_fourth, parentID uint

					rows, err := f.GetRows(sheetlist[i])
					if err != nil {
						logs.Error(err)
						return
					}
					for i, row := range rows {
						if i > 3 { // len(row[i])
							var partOne bool
							number := row[0]
							costname := row[1]
							costname4 := SubString(costname, 0, 4)
							//设定一个含有中文的字符串
							rex := regexp.MustCompile("^第.部分$")
							if rex.MatchString(string(number)) || rex.MatchString(string(costname4)) {
								partOne = true
							}

							if partOne && row[2] == "" && row[3] == "" {
								// 一级
								totalEquipment := row[6]
								totalEquipmentfloat, err := strconv.ParseFloat(totalEquipment, 64)
								if err != nil {
									logs.Error(err)
								}
								totalInstallation := row[7]
								totalInstallationfloat, err := strconv.ParseFloat(totalInstallation, 64)
								if err != nil {
									logs.Error(err)
								}
								costID, err := models.AddEstimateCostMetal(estimatePhaseID, 0, number, costname, "", 0, 0, 0, totalEquipmentfloat, totalInstallationfloat)
								if err != nil {
									logs.Error(err)
								}
								parentID_first = costID
								parentID = costID
							} else if regexp.MustCompile(`^[一-龟]{1,2}$`).MatchString(string(row[0])) && row[2] == "" && row[3] == "" && row[4] == "" {
								// 二级
								totalEquipment := row[6]
								totalEquipmentfloat, err := strconv.ParseFloat(totalEquipment, 64)
								if err != nil {
									logs.Error(err)
								}
								totalInstallation := row[7]
								totalInstallationfloat, err := strconv.ParseFloat(totalInstallation, 64)
								if err != nil {
									logs.Error(err)
								}
								costID, err := models.AddEstimateCostMetal(0, parentID_first, number, costname, "", 0, 0, 0, totalEquipmentfloat, totalInstallationfloat)
								if err != nil {
									logs.Error(err)
								}
								parentID_second = costID
								parentID = costID
							} else if regexp.MustCompile(`^\([一-龟]{1,2}\)$`).MatchString(string(row[0])) && row[2] == "" && row[3] == "" && row[4] == "" {
								// 三级
								totalEquipment := row[6]
								totalEquipmentfloat, err := strconv.ParseFloat(totalEquipment, 64)
								if err != nil {
									logs.Error(err)
								}
								totalInstallation := row[7]
								totalInstallationfloat, err := strconv.ParseFloat(totalInstallation, 64)
								if err != nil {
									logs.Error(err)
								}
								costID, err := models.AddEstimateCostMetal(0, parentID_second, number, costname, "", 0, 0, 0, totalEquipmentfloat, totalInstallationfloat)
								if err != nil {
									logs.Error(err)
								}
								parentID_third = costID
								parentID = costID
							} else if regexp.MustCompile(`^[1-9]\d*$`).MatchString(string(row[0])) && row[2] == "" && row[3] == "" && row[4] == "" {
								// 四级
								totalEquipment := row[6]
								totalEquipmentfloat, err := strconv.ParseFloat(totalEquipment, 64)
								if err != nil {
									logs.Error(err)
								}
								totalInstallation := row[7]
								totalInstallationfloat, err := strconv.ParseFloat(totalInstallation, 64)
								if err != nil {
									logs.Error(err)
								}
								costID, err := models.AddEstimateCostMetal(0, parentID_third, number, costname, "", 0, 0, 0, totalEquipmentfloat, totalInstallationfloat)
								if err != nil {
									logs.Error(err)
								}
								parentID_fourth = costID
								parentID = costID
							} else if regexp.MustCompile(`^\([1-9]\d*\)$`).MatchString(string(row[0])) && row[2] == "" && row[3] == "" && row[4] == "" {
								// 五级
								totalEquipment := row[6]
								totalEquipmentfloat, err := strconv.ParseFloat(totalEquipment, 64)
								if err != nil {
									logs.Error(err)
								}
								totalInstallation := row[7]
								totalInstallationfloat, err := strconv.ParseFloat(totalInstallation, 64)
								if err != nil {
									logs.Error(err)
								}
								costID, err := models.AddEstimateCostMetal(0, parentID_fourth, number, costname, "", 0, 0, 0, totalEquipmentfloat, totalInstallationfloat)
								if err != nil {
									logs.Error(err)
								}
								parentID = costID
							} else if row[1] != "" && row[2] != "" && row[3] != "" && row[4] != "" && row[5] != "" {
								unit := row[2]
								quantity := row[3]
								quantityfloat, err := strconv.ParseFloat(quantity, 64)
								if err != nil {
									logs.Error(err)
								}
								unitpriceEquipment := row[4]
								unitpriceEquipmentfloat, err := strconv.ParseFloat(unitpriceEquipment, 64)
								if err != nil {
									logs.Error(err)
								}
								unitpriceInstallation := row[5]
								unitpriceInstallationfloat, err := strconv.ParseFloat(unitpriceInstallation, 64)
								if err != nil {
									logs.Error(err)
								}
								totalEquipment := row[6]
								totalEquipmentfloat, err := strconv.ParseFloat(totalEquipment, 64)
								if err != nil {
									logs.Error(err)
								}
								totalInstallation := row[7]
								totalInstallationfloat, err := strconv.ParseFloat(totalInstallation, 64)
								if err != nil {
									logs.Error(err)
								}
								costID, err := models.AddEstimateCostMetal(0, parentID, number, costname, unit, quantityfloat, unitpriceEquipmentfloat, unitpriceInstallationfloat, totalEquipmentfloat, totalInstallationfloat)
								if err != nil {
									logs.Error(err)
								}
								logs.Info(costID)
							}
						}
					}
				} else if strings.Contains(sheetname, "施工") ||
					strings.Contains(sheetname, "临时") {
					var parentID_first, parentID_second, parentID_third, parentID_fourth, parentID uint

					rows, err := f.GetRows(sheetlist[i])
					if err != nil {
						fmt.Println(err)
						return
					}
					for i, row := range rows {
						if i > 3 { // len(row[i])
							var partOne bool
							number := row[0]
							costname := row[1]
							costname4 := SubString(costname, 0, 4)
							//设定一个含有中文的字符串
							rex := regexp.MustCompile("^第.部分$")
							if rex.MatchString(string(number)) || rex.MatchString(string(costname4)) {
								partOne = true
							}

							if partOne && row[2] == "" && row[3] == "" {
								// 一级
								total := row[5]
								totalfloat, err := strconv.ParseFloat(total, 64)
								if err != nil {
									logs.Error(err)
								}
								costID, err := models.AddEstimateCostTemp(estimatePhaseID, 0, number, costname, "", 0, 0, totalfloat)
								if err != nil {
									logs.Error(err)
								}
								parentID_first = costID
								parentID = costID
							} else if regexp.MustCompile(`^[一-龟]{1,2}$`).MatchString(string(row[0])) && row[2] == "" && row[3] == "" && row[4] == "" {
								// 二级
								total := row[5]
								totalfloat, err := strconv.ParseFloat(total, 64)
								if err != nil {
									logs.Error(err)
								}
								costID, err := models.AddEstimateCostTemp(0, parentID_first, number, costname, "", 0, 0, totalfloat)
								if err != nil {
									logs.Error(err)
								}
								parentID_second = costID
								parentID = costID
							} else if regexp.MustCompile(`^\([一-龟]{1,2}\)$`).MatchString(string(row[0])) && row[2] == "" && row[3] == "" && row[4] == "" {
								// 三级
								total := row[5]
								totalfloat, err := strconv.ParseFloat(total, 64)
								if err != nil {
									logs.Error(err)
								}
								costID, err := models.AddEstimateCostTemp(0, parentID_second, number, costname, "", 0, 0, totalfloat)
								if err != nil {
									logs.Error(err)
								}
								parentID_third = costID
								parentID = costID
							} else if regexp.MustCompile(`^[1-9]\d*$`).MatchString(string(row[0])) && row[2] == "" && row[3] == "" && row[4] == "" {
								// 四级
								total := row[5]
								totalfloat, err := strconv.ParseFloat(total, 64)
								if err != nil {
									logs.Error(err)
								}
								costID, err := models.AddEstimateCostTemp(0, parentID_third, number, costname, "", 0, 0, totalfloat)
								if err != nil {
									logs.Error(err)
								}
								parentID_fourth = costID
								parentID = costID
							} else if regexp.MustCompile(`^\([1-9]\d*\)$`).MatchString(string(row[0])) && row[2] == "" && row[3] == "" && row[4] == "" {
								// 五级
								total := row[5]
								totalfloat, err := strconv.ParseFloat(total, 64)
								if err != nil {
									logs.Error(err)
								}
								costID, err := models.AddEstimateCostTemp(0, parentID_fourth, number, costname, "", 0, 0, totalfloat)
								if err != nil {
									logs.Error(err)
								}
								parentID = costID
							} else if row[1] != "" && row[2] != "" && row[3] != "" && row[4] != "" && row[5] != "" {
								unit := row[2]
								quantity := row[3]
								quantityfloat, err := strconv.ParseFloat(quantity, 64)
								if err != nil {
									logs.Error(err)
								}
								unitprice := row[4]
								unitpricefloat, err := strconv.ParseFloat(unitprice, 64)
								if err != nil {
									logs.Error(err)
								}
								total := row[5]
								totalfloat, err := strconv.ParseFloat(total, 64)
								if err != nil {
									logs.Error(err)
								}
								costID, err := models.AddEstimateCostTemp(0, parentID, number, costname, unit, quantityfloat, unitpricefloat, totalfloat)
								if err != nil {
									logs.Error(err)
								}
								logs.Info(costID)
							}
						}
					}
				}
			}
		}
		var fileinput Fileinput
		fileinput.InitialPreview = []string{Url + h.Filename}
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
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "文件上传为空", "filelink": ""}
		c.ServeJSON()
	}
}

// @Title update user
// @Description add user
// @Param pk query string false "The pk of user"
// @Param name query string false "The name of user"
// @Param value query string false "The value of user"
// @Success 200 {object} models.User
// @Failure 400 Invalid page supplied
// @Failure 404 user not found
// @router /updateestproject [post]
// 在线修改保存某个字段
func (c *EstimateController) UpdateEstProject() {
	//进行权限判断isme or isadmin
	_, _, _, isadmin, _ := checkprodRole(c.Ctx)
	id := c.GetString("id") //这个其实就是userid
	//id转成uint为
	idint, err := strconv.Atoi(id)
	if err != nil {
		logs.Error(err)
	}
	projid := uint(idint)

	if isadmin { //|| uid == id
		name := c.GetString("name")
		value := c.GetString("value")
		value = template.HTMLEscapeString(value) //过滤xss攻击
		logs.Info(value)
		err = models.UpdateEstProject(projid, name, value)
		if err != nil {
			logs.Error(err)
			data := "写入数据错误!"
			c.Ctx.WriteString(data)
		} else {
			logs.Info("ok")
			data := "ok!"
			c.Ctx.WriteString(data)
		}
	} else {
		data := "权限不足！"
		c.Ctx.WriteString(data)
	}
}

// @Title getestimateprojects
// @Description get estimateprojects
// @Success 200 {object} models.GetEstimateProjects
// @Failure 400 Invalid page supplied
// @Failure 404 Page not found
// @router /getestimateprojects [get]
// 项目列表页面
func (c *EstimateController) GetEstimateProject() {
	username, role, _, isadmin, islogin := checkprodRole(c.Ctx)
	if !islogin {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "用户未登录！", "title": "", "original": ""}
		c.ServeJSON()
		return
	}
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["IsEstimate"] = true
	// c.TplName = "estimate/getEstimateCost.tpl"
	c.TplName = "estimate/estimateProjects.tpl"
}

// @Title getEstimateProjectsData
// @Description get estimateProjects Data
// @Param page query string false "The page of estimateProjectsData"
// @Param limit query string false "The size of page"
// @Success 200 {object} models.GetEstimateProjectsData
// @Failure 400 Invalid page supplied
// @Failure 404 Page not found
// @router /getestimateprojectsdata [get]
// 查询数据
func (c *EstimateController) GetEstimateProjectsData() {
	_, _, uid, _, islogin := checkprodRole(c.Ctx)
	if !islogin {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "用户未登录！", "title": "", "original": ""}
		c.ServeJSON()
		return
	}

	c.Data["UserId"] = uid

	limit := c.GetString("limit")
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

	estimateProjects, err := models.GetEstimateProjects(limit1, offset)
	if err != nil {
		logs.Error(err)
	}

	count, err := models.GetEstimateProjectsCount()
	if err != nil {
		logs.Error(err)
	}
	// 要做成分页的！！！
	table := estimateProjectsTable{estimateProjects, page1, int(count)}
	c.Data["json"] = table
	c.ServeJSON()
}

// @Title getestimatecostarchi
// @Description get costarchi tpl
// @Param id path string  true "The id of projectphase"
// @Success 200 {object} models.GetEstimateCostArchi
// @Failure 400 Invalid page supplied
// @Failure 404 Page not found
// @router /getestimatecostarchi/:id [get]
// 建筑工程费用数据查询页面
func (c *EstimateController) GetEstimateCostArchi() {
	username, role, _, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["IsEstimate"] = true
	phaseID := c.Ctx.Input.Param(":id")
	c.Data["PhaseID"] = phaseID
	// c.TplName = "estimate/getEstimateCost.tpl"
	c.TplName = "estimate/estimateCostArchi.tpl"
}

// @Title getEstimateCostArchiData
// @Description get costArchi
// @Param id path string  true "The id of estimatephase"
// @Param page query string false "The page of estimatecostarchi"
// @Param limit query string false "The size of page"
// @Success 200 {object} models.GetEstimateCostArchi
// @Failure 400 Invalid page supplied
// @Failure 404 Page not found
// @router /getestimatecostarchidata/:id [get]
// 建筑工程费用数据查询
func (c *EstimateController) GetEstimateCostArchiData() {
	_, _, uid, _, islogin := checkprodRole(c.Ctx)
	if !islogin {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "用户未登录！", "title": "", "original": ""}
		c.ServeJSON()
		return
	}

	id := c.Ctx.Input.Param(":id")
	var iduint uint
	if id != "" {
		//id转成uint为
		idint, err := strconv.Atoi(id)
		if err != nil {
			logs.Error(err)
		}
		iduint = uint(idint)
	}

	c.Data["UserId"] = uid

	limit := c.GetString("limit")
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

	estimatecost, err := models.GetEstimateCostArchi(iduint, limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	// logs.Info(estimatecost)

	count, err := models.GetEstimateCostArchiCount(iduint)
	if err != nil {
		logs.Error(err)
	}
	// logs.Info(count)
	// 要做成分页的！！！
	res := getEsArchiTreeRecursive(estimatecost, 0)
	logs.Info(res)
	table := EstimateCostArchiTable{0, int(count), res}
	// logs.Info(table)
	c.Data["json"] = table
	// c.Data["json"] = estimatecost
	c.ServeJSON()
}

// @Title getestimatecostelect
// @Description get costelect tpl
// @Param id path string  true "The id of projectphase"
// @Success 200 {object} models.GetEstimateCostElect
// @Failure 400 Invalid page supplied
// @Failure 404 Page not found
// @router /getestimatecostelect/:id [get]
// 机电设备安装工程费用数据查询页面
func (c *EstimateController) GetEstimateCostElect() {
	username, role, _, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["IsEstimate"] = true
	phaseID := c.Ctx.Input.Param(":id")
	c.Data["PhaseID"] = phaseID
	// c.TplName = "estimate/getEstimateCost.tpl"
	c.TplName = "estimate/estimateCostElect.tpl"
}

// @Title getEstimateCostElectData
// @Description get costArchi
// @Param id path string  true "The id of estimatephase"
// @Param page query string false "The page of estimatecostelect"
// @Param limit query string false "The size of page"
// @Success 200 {object} models.GetEstimateCostElect
// @Failure 400 Invalid page supplied
// @Failure 404 Page not found
// @router /getestimatecostelectdata/:id [get]
// 建筑工程费用数据查询
func (c *EstimateController) GetEstimateCostElectData() {
	_, _, uid, _, islogin := checkprodRole(c.Ctx)
	if !islogin {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "用户未登录！", "title": "", "original": ""}
		c.ServeJSON()
		return
	}

	id := c.Ctx.Input.Param(":id")
	var iduint uint
	if id != "" {
		//id转成uint为
		idint, err := strconv.Atoi(id)
		if err != nil {
			logs.Error(err)
		}
		iduint = uint(idint)
	}

	c.Data["UserId"] = uid

	limit := c.GetString("limit")
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

	estimatecost, err := models.GetEstimateCostElect(iduint, limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	// logs.Info(estimatecost)

	count, err := models.GetEstimateCostElectCount(iduint)
	if err != nil {
		logs.Error(err)
	}
	// logs.Info(count)
	// 要做成分页的！！！
	res := getEsElectTreeRecursive(estimatecost, 0)
	logs.Info(res)
	table := EstimateCostElectTable{0, int(count), res}
	// logs.Info(table)
	c.Data["json"] = table
	// c.Data["json"] = estimatecost
	c.ServeJSON()
}

// @Title getestimatecostmetal
// @Description get costmetal tpl
// @Param id path string  true "The id of projectphase"
// @Success 200 {object} models.GetEstimateCostMetal
// @Failure 400 Invalid page supplied
// @Failure 404 Page not found
// @router /getestimatecostmetal/:id [get]
// 金属结构安装工程费用数据查询页面
func (c *EstimateController) GetEstimateCostMetal() {
	username, role, _, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["IsEstimate"] = true
	phaseID := c.Ctx.Input.Param(":id")
	c.Data["PhaseID"] = phaseID
	// c.TplName = "estimate/getEstimateCost.tpl"
	c.TplName = "estimate/estimateCostMetal.tpl"
}

// @Title getEstimateCostMetalData
// @Description get costArchi
// @Param id path string  true "The id of estimatephase"
// @Param page query string false "The page of estimatecostmetal"
// @Param limit query string false "The size of page"
// @Success 200 {object} models.GetEstimateCostMetal
// @Failure 400 Invalid page supplied
// @Failure 404 Page not found
// @router /getestimatecostmetaldata/:id [get]
// 建筑工程费用数据查询
func (c *EstimateController) GetEstimateCostMetalData() {
	_, _, uid, _, islogin := checkprodRole(c.Ctx)
	if !islogin {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "用户未登录！", "title": "", "original": ""}
		c.ServeJSON()
		return
	}

	id := c.Ctx.Input.Param(":id")
	var iduint uint
	if id != "" {
		//id转成uint为
		idint, err := strconv.Atoi(id)
		if err != nil {
			logs.Error(err)
		}
		iduint = uint(idint)
	}

	c.Data["UserId"] = uid

	limit := c.GetString("limit")
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

	estimatecost, err := models.GetEstimateCostMetal(iduint, limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	// logs.Info(estimatecost)

	count, err := models.GetEstimateCostMetalCount(iduint)
	if err != nil {
		logs.Error(err)
	}
	// logs.Info(count)
	// 要做成分页的！！！
	res := getEsMetalTreeRecursive(estimatecost, 0)
	logs.Info(res)
	table := EstimateCostMetalTable{0, int(count), res}
	// logs.Info(table)
	c.Data["json"] = table
	// c.Data["json"] = estimatecost
	c.ServeJSON()
}

// @Title getestimatecostTemp
// @Description get costtemp tpl
// @Param id path string  true "The id of projectphase"
// @Success 200 {object} models.GetEstimateCostTemp
// @Failure 400 Invalid page supplied
// @Failure 404 Page not found
// @router /getestimatecosttemp/:id [get]
// 临时工程费用数据查询页面
func (c *EstimateController) GetEstimateCostTemp() {
	username, role, _, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["IsEstimate"] = true
	phaseID := c.Ctx.Input.Param(":id")
	c.Data["PhaseID"] = phaseID
	// c.TplName = "estimate/getEstimateCost.tpl"
	c.TplName = "estimate/estimateCostTemp.tpl"
}

// @Title getEstimateCostTempData
// @Description get costTemp
// @Param id path string  true "The id of estimatephase"
// @Param page query string false "The page of estimatecosttemp"
// @Param limit query string false "The size of page"
// @Success 200 {object} models.GetEstimateCostTemp
// @Failure 400 Invalid page supplied
// @Failure 404 Page not found
// @router /getestimatecosttempdata/:id [get]
// 建筑工程费用数据查询
func (c *EstimateController) GetEstimateCostTempData() {
	_, _, uid, _, islogin := checkprodRole(c.Ctx)
	if !islogin {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "用户未登录！", "title": "", "original": ""}
		c.ServeJSON()
		return
	}

	id := c.Ctx.Input.Param(":id")
	var iduint uint
	if id != "" {
		//id转成uint为
		idint, err := strconv.Atoi(id)
		if err != nil {
			logs.Error(err)
		}
		iduint = uint(idint)
	}

	c.Data["UserId"] = uid

	limit := c.GetString("limit")
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

	estimatecost, err := models.GetEstimateCostTemp(iduint, limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	// logs.Info(estimatecost)

	count, err := models.GetEstimateCostTempCount(iduint)
	if err != nil {
		logs.Error(err)
	}
	// logs.Info(count)
	// 要做成分页的！！！
	res := getEsTempTreeRecursive(estimatecost, 0)
	logs.Info(res)
	table := EstimateCostTempTable{0, int(count), res}
	// logs.Info(table)
	c.Data["json"] = table
	// c.Data["json"] = estimatecost
	c.ServeJSON()
}
