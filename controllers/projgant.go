//项目进度控制器，具体任务控制器另外做吧
package controllers

import (
	"encoding/json"
	// beego "github.com/beego/beego/v2/adapter"
	"github.com/tealeg/xlsx"
	// "github.com/beego/beego/v2/adapter/utils/pagination"
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"os"
	// "path"
	// "path/filepath"
	"strconv"
	"strings"
	"time"
)

type ProjGantController struct {
	web.Controller
}

type Gantt struct {
	Tasks            []Task           `json:"tasks"`
	Resources        []Resourcesvalue `json:"resources"`
	Roles            []Rolesvalue     `json:"roles"`
	SelectedRow      int64            `json:"selectedRow"`
	DeletedTaskIds   []int64          `json:"deletedTaskIds"`
	CanWrite         bool             `json:"canWrite"`
	CanWriteOnParent bool             `json:"canWriteOnParent"`
	Zoom             string           `json:"zoom"` //"w3"
}

type Task struct {
	Id               int64         `json:"id"`
	Status           string        `json:"status"`
	Level            int           `json:"level"`
	Code             string        `json:"code"`
	Name             string        `json:"name"`
	StartIsMilestone bool          `json:"startIsMilestone"`
	Start            int64         `json:"start"`
	EndIsMilestone   bool          `json:"endIsMilestone"`
	End              int64         `json:"end"`
	Duration         int           `json:"duration"`
	Progress         int           `json:"progress"`
	Depends          string        `json:"depends"`
	HasChild         bool          `json:"hasChild"`
	Description      string        `json:"description"`
	Relevance        int           `json:"relevance"`
	Type             string        `json:"type"`
	TypeId           string        `json:"typeId"`
	CanWrite         bool          `json:"canWrite"`
	Collapsed        bool          `json:"collapsed"`
	Assigs           []Assigsvalue `json:"assigs"`
	// 						"resourceId":"tmp_1",
	//             "id":"tmp_1345625008213",
	//             "roleId":"tmp_1",
	//             "effort":7200000

	// "resources":[
	//       {"id": "tmp_1",
	//        "name": "秦晓川"
	//        },
	//       {"id": "tmp_2", "name": "冯文涛"},
	//       {"id": "tmp_3", "name": "张武"},
	//       {"id": "tmp_4", "name": "陈小云"}
	//     ],
	//     "roles":[
	//       {"id": "tmp_1",
	//       "name": "项目负责人"
	//       },
	//       {"id": "tmp_2", "name": "专业负责人"},
	//       {"id": "tmp_3", "name": "专业负责人"},
	//       {"id": "tmp_4", "name": "审查"}
	//     ],
}

type Assigsvalue struct {
	ResourceId string `json:"resourceId"`
	Id         string `json:"id"`
	RoleId     string `json:"roleId"`
	Effort     int64  `json:"effort"`
}

type Resourcesvalue struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Rolesvalue struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

//项目列表页面
func (c *ProjGantController) Get() {
	// username, role := checkprodRole(c.Ctx)
	// roleint, err := strconv.Atoi(role)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// if role == "1" {
	// 	c.Data["IsAdmin"] = true
	// } else if roleint > 1 && roleint < 5 {
	// 	c.Data["IsLogin"] = true
	// } else {
	// 	c.Data["IsAdmin"] = false
	// 	c.Data["IsLogin"] = false
	// }
	// c.Data["Username"] = username
	c.Data["IsProjectgant"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	c.TplName = "projects_gant.tpl"

	projgants, err := models.GetProjGants()
	if err != nil {
		logs.Error(err)
	}
	//一次性取出所有没有关闭的项目
	//A将leve=0级项目结束时间与当前时间对比，在当前时间前的，放到后面A2；
	//在当前时间后的，按先后顺序排列A1
	//A1循环leve=0级
	//——循环所有leve=1级，parentid为当前A1的，则append进A1，B1
	//————循环所有leve=2级，parentid为当前B1的，C1
	//循环A2，同上
	A1 := make([]*models.ProjGant, 0)
	A2 := make([]*models.ProjGant, 0)
	B1 := make([]*models.ProjGant, 0)
	C1 := make([]*models.ProjGant, 0)
	for i, v := range projgants {
		//所有0级项目分为当前时间前和后
		if v.End.After(time.Now()) && v.Level == 0 {
			A1 = append(A1, projgants[i])
		} else if v.End.Before(time.Now()) && v.Level == 0 {
			A2 = append(A2, projgants[i])
		} else if v.Level == 1 {
			B1 = append(B1, projgants[i])
		} else if v.Level == 2 {
			C1 = append(C1, projgants[i])
		}
	}

	gantt := new(Gantt)
	task := make([]Task, 0)
	assigsvalue := make([]Assigsvalue, 0)
	resourcesvalue := make([]Resourcesvalue, 0)
	rolesvalue := make([]Rolesvalue, 0)

	for _, v1 := range A1 {
		aa := make([]Task, 1)
		aa[0].Id = v1.Id
		aa[0].Status = v1.Status
		aa[0].Level = v1.Level
		aa[0].Code = v1.Code
		aa[0].Name = v1.Name
		aa[0].StartIsMilestone = v1.StartIsMilestone
		aa[0].EndIsMilestone = v1.EndIsMilestone
		aa[0].Duration = v1.Duration
		aa[0].Progress = v1.Progress
		aa[0].Depends = v1.Depends
		aa[0].HasChild = v1.HasChild
		aa[0].Description = v1.Description
		aa[0].Start = (v1.Start).UnixNano() / 1e6
		aa[0].End = (v1.End).UnixNano() / 1e6
		aa[0].CanWrite = true
		aa[0].Collapsed = false

		bb := make([]Assigsvalue, 1)
		// //查出一个task的资源，循环
		bb[0].Id = "tmp_1345625008213"
		bb[0].ResourceId = "tmp_1"
		bb[0].RoleId = "tmp_1"
		bb[0].Effort = 7200000
		assigsvalue = append(assigsvalue, bb...)
		aa[0].Assigs = assigsvalue
		assigsvalue = make([]Assigsvalue, 0)
		task = append(task, aa...)

		for _, v2 := range B1 {
			if v2.ParentId == v1.Id {
				aa := make([]Task, 1)
				aa[0].Id = v2.Id
				aa[0].Status = v2.Status
				aa[0].Level = v2.Level
				aa[0].Code = v2.Code
				aa[0].Name = v2.Name
				aa[0].StartIsMilestone = v2.StartIsMilestone
				aa[0].EndIsMilestone = v2.EndIsMilestone
				aa[0].Duration = v2.Duration
				aa[0].Progress = v2.Progress
				aa[0].Depends = v2.Depends
				aa[0].HasChild = v2.HasChild
				aa[0].Description = v2.Description
				aa[0].Start = (v2.Start).UnixNano() / 1e6
				aa[0].End = (v2.End).UnixNano() / 1e6
				aa[0].CanWrite = true
				aa[0].Collapsed = false

				bb := make([]Assigsvalue, 1)
				// //查出一个task的资源，循环
				bb[0].Id = "tmp_1345625008213"
				bb[0].ResourceId = "tmp_1"
				bb[0].RoleId = "tmp_1"
				bb[0].Effort = 7200000
				assigsvalue = append(assigsvalue, bb...)
				aa[0].Assigs = assigsvalue
				assigsvalue = make([]Assigsvalue, 0)
				task = append(task, aa...)
				for _, v3 := range C1 {
					if v3.ParentId == v2.Id {
						aa := make([]Task, 1)
						aa[0].Id = v3.Id
						aa[0].Status = v3.Status
						aa[0].Level = v3.Level
						aa[0].Code = v3.Code
						aa[0].Name = v3.Name
						aa[0].StartIsMilestone = v3.StartIsMilestone
						aa[0].EndIsMilestone = v3.EndIsMilestone
						aa[0].Duration = v3.Duration
						aa[0].Progress = v3.Progress
						aa[0].Depends = v3.Depends
						aa[0].HasChild = v3.HasChild
						aa[0].Description = v3.Description
						aa[0].Start = (v3.Start).UnixNano() / 1e6
						aa[0].End = (v3.End).UnixNano() / 1e6
						aa[0].CanWrite = true
						aa[0].Collapsed = false

						bb := make([]Assigsvalue, 1)
						// //查出一个task的资源，循环
						bb[0].Id = "tmp_1345625008213"
						bb[0].ResourceId = "tmp_1"
						bb[0].RoleId = "tmp_1"
						bb[0].Effort = 7200000
						assigsvalue = append(assigsvalue, bb...)
						aa[0].Assigs = assigsvalue
						assigsvalue = make([]Assigsvalue, 0)
						task = append(task, aa...)
					}
				}
			}
		}
	}

	// for _, v1 := range projgants {
	// 	aa := make([]Task, 1)
	// 	aa[0].Id = v1.Id
	// 	aa[0].Status = v1.Status
	// 	aa[0].Level = v1.Level
	// 	aa[0].Code = v1.Code
	// 	aa[0].Name = v1.Name
	// 	aa[0].StartIsMilestone = v1.StartIsMilestone
	// 	aa[0].EndIsMilestone = v1.EndIsMilestone
	// 	aa[0].Duration = v1.Duration
	// 	aa[0].Progress = v1.Progress
	// 	aa[0].Depends = v1.Depends
	// 	aa[0].HasChild = v1.HasChild
	// 	aa[0].Description = v1.Description
	// 	aa[0].Start = (v1.Start).UnixNano() / 1e6
	// 	aa[0].End = (v1.End).UnixNano() / 1e6
	// 	aa[0].CanWrite = true
	// 	aa[0].Collapsed = false
	// 	bb := make([]Assigsvalue, 1)
	// 	// //查出一个task的资源，循环
	// 	bb[0].Id = "tmp_1345625008213"
	// 	bb[0].ResourceId = "tmp_1"
	// 	bb[0].RoleId = "tmp_1"
	// 	bb[0].Effort = 7200000
	// 	assigsvalue = append(assigsvalue, bb...)
	// 	aa[0].Assigs = assigsvalue
	// 	task = append(task, aa...)
	// }
	cc := make([]Resourcesvalue, 1)
	cc[0].Id = "tmp_1"
	cc[0].Name = "秦晓川"
	resourcesvalue = append(resourcesvalue, cc...)

	dd := make([]Rolesvalue, 1)
	dd[0].Id = "tmp_1"
	dd[0].Name = "项目负责人"
	rolesvalue = append(rolesvalue, dd...)
	gantt.Tasks = task
	gantt.Resources = resourcesvalue
	gantt.Roles = rolesvalue
	gantt.SelectedRow = 0
	// gantt.DeletedTaskIds=[0,1]
	gantt.CanWrite = true
	gantt.CanWriteOnParent = true
	gantt.Zoom = "w3"

	c.Data["Gantt"] = gantt
	// c.ServeJSON()
}

//提供给项目列表页的table中json数据，扩展后按标签显示
// func (c *ProjGantController) GetProjGants() {

// }

//根据id查看项目，查出项目目录
func (c *ProjGantController) GetProjectGant() {
	// username, role := checkprodRole(c.Ctx)
	// if role == 1 {
	// 	c.Data["IsAdmin"] = true
	// } else if role > 1 && role < 5 {
	// 	c.Data["IsLogin"] = true
	// } else {
	// 	c.Data["IsAdmin"] = false
	// 	c.Data["IsLogin"] = false
	// }
	// c.Data["Username"] = username
	// c.Data["IsProject"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role

	// id := c.Ctx.Input.Param(":id")
	// c.Data["Id"] = id
	// // var categories []*models.ProjCategory
	// var err error
	// //id转成64为
	// idNum, err := strconv.ParseInt(id, 10, 64)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// //取项目本身
	// category, err := models.GetProj(idNum)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// //取项目所有子孙
	// categories, err := models.GetProjectsbyPid(idNum)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// //根据id取出下级
	// cates := getsons(idNum, categories)
	// //算出最大级数
	// // grade := make([]int, 0)
	// // for _, v := range categories {
	// // 	grade = append(grade, v.Grade)
	// // }
	// // height := intmax(grade[0], grade[1:]...)
	// //递归生成目录json
	// root := FileNode{category.Id, category.Title, "", []*FileNode{}}
	// // walk(category.Id, &root)
	// maketreejson(cates, categories, &root)
	// // beego.Info(root)
	// // data, _ := json.Marshal(root)
	// c.Data["json"] = root //data
	// // c.ServeJSON()
	// c.Data["Category"] = category
	// c.TplName = "cms/project.tpl"
}

//根据项目侧栏id查看这个id下的成果，不含子目录中的成果
//任何一级目录下都可以放成果
//这个作废——以product中的GetProducts
func (c *ProjGantController) GetProjGant() {
	id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	c.Data["Id"] = id
	// var categories []*models.ProjCategory
	// var err error
	//id转成64为
	// idNum, err := strconv.ParseInt(id, 10, 64)
	// if err != nil {
	// 	logs.Error(err)
	// }
	//取项目本身
	// category, err := models.GetProj(idNum)
	// if err != nil {
	// 	logs.Error(err)
	// }
	//取项目所有子孙
	// categories, err := models.GetProjectsbyPid(idNum)
	// if err != nil {
	// 	logs.Error(err)
	// }
	//算出最大级数
	// grade := make([]int, 0)
	// for _, v := range categories {
	// 	grade = append(grade, v.Grade)
	// }
	// height := intmax(grade[0], grade[1:]...)

	// c.Data["json"] = root
	// c.ServeJSON()
	c.TplName = "cms/project_products.tpl"
}

//添加项目和项目目录、文件夹
func (c *ProjGantController) AddProjGant() {
	// iprole := Getiprole(c.Ctx.Input.IP())
	// if iprole != 1 {
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/roleerr?url="+route, 302)
	// 	// c.Redirect("/roleerr", 302)
	// 	return
	// }
	// rows := c.GetString("rows2[0][0]")

	// status						:= c.GetString("status					")
	// level						:= c.GetString("level						")
	// code 						:= c.GetString("code 						")
	// name							:= c.GetString("name						")
	// startismilestone:= c.GetString("startismilestone")
	// start 						:= c.GetString("start 					")
	// endismilestone		:= c.GetString("endismilestone	")
	// end							:= c.GetString("end							")
	// duration					:= c.GetString("duration				")
	// progress					:= c.GetString("progress				")
	// depends					:= c.GetString("depends					")
	// haschild					:= c.GetString("haschild				")
	// description			:= c.GetString("description			")

	var Ganttstruct Gantt
	tt := []byte(c.GetString("prj"))
	json.Unmarshal(tt, &Ganttstruct)
	// beego.Info(Ganttstruct.Tasks[0].Id)
	var parentid int64
	for i, v := range Ganttstruct.Tasks {
		if v.Level != 0 {
			for j := i; j >= 0; j-- {
				if Ganttstruct.Tasks[j].Level == v.Level-1 {
					//这个id错误。如果是还没建立的，则需呀用后面的插入数据库后返回的id来作为下一个的parentid
					//根据编号和名称查询数据库中的id？
					parentid = Ganttstruct.Tasks[j].Id
					break //跳出循环
				}
			}
		} else {
			parentid = 0
		}
		// beego.Info(parentid)
		// beego.Info(v.Id)
		id := v.Id
		status := v.Status
		level := v.Level
		// levelint, err := strconv.Atoi(level)
		// if err != nil {
		// 	logs.Error(err)
		// }
		code := v.Code
		name := v.Name
		startismilestone := v.StartIsMilestone
		// var startismilestonebool, endismilestonebool, haschildbool bool
		// if startismilestone == "true" {
		// 	startismilestonebool = true
		// } else {
		// 	startismilestonebool = false
		// }1396994400.000
		start := time.Unix(v.Start/1e3, 0)
		// (v1.Start).UnixNano() / 1e6
		endismilestone := v.EndIsMilestone
		// if endismilestone == "true" {
		// 	endismilestonebool = true
		// } else {
		// 	endismilestonebool = false
		// }
		end := time.Unix(v.End/1e3, 0)
		duration := v.Duration
		// durationint, err := strconv.Atoi(duration)
		// if err != nil {
		// 	logs.Error(err)
		// }
		progress := v.Progress
		// progressint, err := strconv.Atoi(progress)
		// if err != nil {
		// 	logs.Error(err)
		// }
		depends := v.Depends
		haschild := v.HasChild
		// if haschild == "true" {
		// 	haschildbool = true
		// } else {
		// 	haschildbool = false
		// }
		description := v.Description

		// type Duration int64
		// const (
		// 	Nanosecond  Duration = 1
		// 	Microsecond          = 1000 * Nanosecond
		// 	Millisecond          = 1000 * Microsecond
		// 	Second               = 1000 * Millisecond
		// 	Minute               = 60 * Second
		// 	Hour                 = 60 * Minute
		// )
		// hours := 0
		// var t1, t2 time.Time
		// const lll = "2006-01-02"
		// starttime, _ := time.Parse(lll, start)
		// endtime, _ := time.Parse(lll, end)
		// t1 = starttime.Add(-time.Duration(hours) * time.Hour)
		// beego.Info(t1)：2016-08-19 00:00:00 +0000 UTC
		// t2 = endtime.Add(-time.Duration(hours) * time.Hour)
		// beego.Info(t2)
		_, err := models.AddProjGant(id, parentid, status, code, name, depends, description, level, duration, progress, start, end, startismilestone, endismilestone, haschild)
		if err != nil {
			logs.Error(err)
		}

	}
	c.Data["json"] = "ok"
	c.ServeJSON()
}

//导入甘特数据
//上传excel文件，导入到数据库
func (c *ProjGantController) ImportProjGant() {
	// 获取上传的文件
	_, h, err := c.GetFile("gantsexcel")
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(h.path)
	// var attachment string
	var path string
	const lll = "2006-01-02"
	var convdate string
	var id, secid int64
	var duration, progress int
	var code, name, description string
	var date, t1 time.Time
	// var filesize int64
	if h != nil {
		//保存附件
		path = "./attachment/" + h.Filename    // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("gantsexcel", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			logs.Error(err)
			c.Data["json"] = "err保存文件失败！"
			c.ServeJSON()
		} else {
			//读出excel内容写入数据库
			xlFile, err := xlsx.OpenFile(path) //
			if err != nil {
				logs.Error(err)
			}
			for _, sheet := range xlFile.Sheets {
				for i, row := range sheet.Rows {
					if i != 0 { //忽略第一行标题
						// 这里要判断单元格列数，如果超过单元格使用范围的列数，则出错for j := 2; j < 7; j += 5 {
						j := 1
						//读取编号
						if len(row.Cells) >= 2 { //总列数，从1开始
							code = row.Cells[j].String()
							if err != nil {
								logs.Error(err)
							}
						}
						//读取开始时间
						if len(row.Cells) >= 7 {
							if row.Cells[j+5].Value != "" {
								Starttime, err := row.Cells[j+5].Float()
								if err != nil {
									logs.Error(err)
								} else {
									date = xlsx.TimeFromExcelTime(Starttime, false)
								}
							} else {
								date = time.Now()
							}
							convdate = date.Format(lll)

							date, err = time.Parse(lll, convdate)
							if err != nil {
								logs.Error(err)
							}
							t1 = date
						}
						//读取时间跨度
						if len(row.Cells) >= 8 {
							duration, err = row.Cells[j+6].Int()
							if err != nil {
								logs.Error(err)
							}
						}
						//读取进度
						if len(row.Cells) >= 9 {
							progress, err = row.Cells[j+7].Int()
							if err != nil {
								logs.Error(err)
							}
						}
						//读取描述
						if len(row.Cells) >= 10 {
							description = row.Cells[j+8].String()
							if err != nil {
								logs.Error(err)
							}
						}
						//读取项目名称
						if len(row.Cells) >= 3 {
							name = row.Cells[j+1].String()
							if err != nil {
								logs.Error(err)
							}
							//查询编号和名称，如果存在，则返回id作为parentid
							gantname, err := models.GetProjGantName(code, name)
							if err != nil {
								logs.Error(err) //没找到记录，新建
								id, err = models.AddProjGant(0, 0, "STATUS_ACTIVE", code, name, "", "", 0, duration, progress, t1, t1, false, false, true)
								if err != nil {
									logs.Error(err)
								}
							} else { //找到，则作为parentid
								id = gantname.Id
							}
						}
						//读取阶段
						if len(row.Cells) >= 4 {
							designstage := row.Cells[j+2].String()
							if err != nil {
								logs.Error(err)
							}
							//查询名称和parentid，如果存在，则返回id作为parentid
							gantparent, err := models.GetProjGantParent(designstage, id)
							if err != nil {
								logs.Error(err) //没找到记录，新建
								secid, err = models.AddProjGant(0, id, "STATUS_ACTIVE", "", designstage, "", "", 1, duration, progress, t1, t1, false, false, true)
								if err != nil {
									logs.Error(err)
								}
							} else { //找到，则作为parentid
								secid = gantparent.Id
							}
						}

						//读取专业
						if len(row.Cells) >= 5 {
							section := row.Cells[j+3].String()
							if err != nil {
								logs.Error(err)
							}
							//查询名称和parentid，如果存在，则返回id作为parentid
							_, err = models.GetProjGantParent(section, secid)
							if err != nil {
								logs.Error(err) //没找到记录，新建
								_, err := models.AddProjGant(0, secid, "STATUS_ACTIVE", "", section, "", description, 2, duration, progress, t1, t1, false, false, false)
								if err != nil {
									logs.Error(err)
								}
							}
						}
						//读取任务
						// if len(row.Cells) >= 6 {
						// 	task, err = row.Cells[j+4].String()
						// 	if err != nil {
						// 		logs.Error(err)
						// 	}
						// 	//查询名称和parentid，如果存在，则返回id作为parentid
						// 	gantparent, err := models.GetProjGantParent(task, id3)
						// 	if err != nil {
						// 		logs.Error(err) //没找到记录，新建
						// 		_, err := models.AddProjGant2(id3, "", task, start, duration, progress)
						// 		if err != nil {
						// 			logs.Error(err)
						// 		}
						// 	} else { //找到，则作为parentid
						// 		id = gantparent.Id
						// 	}
						// }

						// if len(row.Cells) >= 10 {
						// 	if row.Cells[j+8].Value != "" {
						// 		endtime2, err := row.Cells[j+8].Float()
						// 		if err != nil {
						// 			logs.Error(err)
						// 		} else {
						// 			date = xlsx.TimeFromExcelTime(endtime2, false)
						// 		}
						// 	} else {
						// 		date = time.Now()
						// 	}
						// 	convdate = date.Format(lll)

						// 	date, err = time.Parse(lll, convdate)
						// 	if err != nil {
						// 		logs.Error(err)
						// 	}
						// 	t1 = date
						// }
						// if len(row.Cells) >= 11 {
						// 	if row.Cells[j+9].Value != "" {
						// 		endtime2, err := row.Cells[j+9].Float()
						// 		if err != nil {
						// 			logs.Error(err)
						// 		} else {
						// 			date = xlsx.TimeFromExcelTime(endtime2, false)
						// 		}
						// 	} else {
						// 		date = time.Now()
						// 	}
						// 	convdate = date.Format(lll)

						// 	date, err = time.Parse(lll, convdate)
						// 	if err != nil {
						// 		logs.Error(err)
						// 	}
						// 	t2 = date
						// }
						// _, err := models.AddProjGant(code, title, designstage, section, label, desc, customclass, dataobj, t1, t2)
						// if err != nil {
						// 	logs.Error(err)
						// }
					}
				}
			}
			//删除附件
			err = os.Remove(path)
			if err != nil {
				logs.Error(err)
			}
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = "err文件为空！"
		c.ServeJSON()
	}
}

//修改项目名称、负责人等，
func (c *ProjGantController) UpdateProjGant() {
	iprole := Getiprole(c.Ctx.Input.IP())
	if iprole != 1 {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		return
	}
	var err error
	projcode := c.GetString("code")
	projname := c.GetString("name")
	projlabe := c.GetString("label")
	principal := c.GetString("principal")
	pid := c.GetString("pid")
	//id转成64位
	idNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	err = models.UpdateProject(idNum, projcode, projname, projlabe, principal)
	if err != nil {
		logs.Error(err)
	}

	if err != nil {
		c.Data["json"] = "no"
		c.ServeJSON()
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
	// c.Data["json"] = "ok"
	// c.ServeJSON()
}

//根据id删除proj
//后台删除目录，
func (c *ProjGantController) DeleteProjGant() {
	iprole := Getiprole(c.Ctx.Input.IP())
	if iprole != 1 {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// var err error
	//查所有子孙项目，循环删除
	ids := c.GetString("ids")
	// beego.Info(ids)
	array := strings.Split(ids, ",")
	//循环项目id
	for _, v := range array {
		//id转成64位
		projid, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			logs.Error(err)
		}
		//根据项目id取得所有子孙id
		projs, err := models.GetProjectsbyPid(projid)
		if err != nil {
			logs.Error(err)
		}
		//循环子孙项目
		for _, w := range projs {
			//取得子孙项目的成果列表
			//根据项目id取得所有成果
			products, err := models.GetProducts(w.Id)
			if err != nil {
				logs.Error(err)
			}
			for _, x := range products {
				//删除子孙成果表
				//循环删除成果
				//根据成果id取得所有附件
				attachments, err := models.GetAttachments(x.Id)
				if err != nil {
					logs.Error(err)
				}
				//删除附件表
				for _, y := range attachments {
					//删除附件数据表
					err = models.DeleteAttachment(y.Id)
					if err != nil {
						logs.Error(err)
					}
				}

				//删除子孙文章表
				//取得成果id下所有文章
				articles, err := models.GetArticles(x.Id)
				if err != nil {
					logs.Error(err)
				}
				//删除文章表
				for _, z := range articles {
					//删除文章数据表
					err = models.DeleteArticle(z.Id)
					if err != nil {
						logs.Error(err)
					}
				}
				//删除成果表自身
				err = models.DeleteProduct(x.Id) //删除成果数据表
				if err != nil {
					logs.Error(err)
				}
			}
			//删除子孙proj数据表
			err = models.DeleteProject(w.Id)
			if err != nil {
				logs.Error(err)
			}
			//删除子孙文章图片文件夹（下面已经全部删除了）
		}
		//根据proj的id——这个放deleteproject前面，否则项目数据表删除了就取不到路径了
		_, DiskDirectory, err := GetUrlPath(projid)
		if err != nil {
			logs.Error(err)
		} else {
			// beego.Info(DiskDirectory)
			path := DiskDirectory
			//直接删除这个文件夹，remove删除文件
			err = os.RemoveAll(path)
			if err != nil {
				logs.Error(err)
			}
			//删除项目自身数据表
			err = models.DeleteProject(projid)
			if err != nil {
				logs.Error(err)
			}
		}
	}
	// if err != nil {
	// 	c.Data["json"] = "no"
	// 	c.ServeJSON()
	// } else {
	c.Data["json"] = "ok"
	c.ServeJSON()
	// }
}

//关闭项目进度
func (c *ProjGantController) CloseProjGant() {
	iprole := Getiprole(c.Ctx.Input.IP())
	if iprole != 1 {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// var err error
	//查所有子孙项目，循环删除
	ids := c.GetString("ids")
	// beego.Info(ids)
	array := strings.Split(ids, ",")
	//循环项目id
	for _, v := range array {
		//id转成64位
		projid, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			logs.Error(err)
		}
		//根据项目id取得所有子孙id
		projs, err := models.GetProjectsbyPid(projid)
		if err != nil {
			logs.Error(err)
		}
		//循环子孙项目
		for _, w := range projs {
			//取得子孙项目的成果列表
			//根据项目id取得所有成果
			products, err := models.GetProducts(w.Id)
			if err != nil {
				logs.Error(err)
			}
			for _, x := range products {
				//删除子孙成果表
				//循环删除成果
				//根据成果id取得所有附件
				attachments, err := models.GetAttachments(x.Id)
				if err != nil {
					logs.Error(err)
				}
				//删除附件表
				for _, y := range attachments {
					//删除附件数据表
					err = models.DeleteAttachment(y.Id)
					if err != nil {
						logs.Error(err)
					}
				}

				//删除子孙文章表
				//取得成果id下所有文章
				articles, err := models.GetArticles(x.Id)
				if err != nil {
					logs.Error(err)
				}
				//删除文章表
				for _, z := range articles {
					//删除文章数据表
					err = models.DeleteArticle(z.Id)
					if err != nil {
						logs.Error(err)
					}
				}
				//删除成果表自身
				err = models.DeleteProduct(x.Id) //删除成果数据表
				if err != nil {
					logs.Error(err)
				}
			}
			//删除子孙proj数据表
			err = models.DeleteProject(w.Id)
			if err != nil {
				logs.Error(err)
			}
			//删除子孙文章图片文件夹（下面已经全部删除了）
		}
		//根据proj的id——这个放deleteproject前面，否则项目数据表删除了就取不到路径了
		_, DiskDirectory, err := GetUrlPath(projid)
		if err != nil {
			logs.Error(err)
		} else {
			// beego.Info(DiskDirectory)
			path := DiskDirectory
			//直接删除这个文件夹，remove删除文件
			err = os.RemoveAll(path)
			if err != nil {
				logs.Error(err)
			}
			//删除项目自身数据表
			err = models.DeleteProject(projid)
			if err != nil {
				logs.Error(err)
			}
		}
	}
	// if err != nil {
	// 	c.Data["json"] = "no"
	// 	c.ServeJSON()
	// } else {
	c.Data["json"] = "ok"
	c.ServeJSON()
	// }
}
