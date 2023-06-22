package controllers

import (
	// json "encoding/json"
	// "fmt"
	// beego "github.com/beego/beego/v2/adapter"
	// "github.com/tealeg/xlsx"
	// "github.com/bitly/go-simplejson"
	// "io/ioutil"
	// "github.com/beego/beego/v2/adapter/logs"
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	// "sort"
	"strconv"
	// "strings"
	"regexp"
	"time"
)

type IndexController struct {
	web.Controller
}

type Userselect struct { //
	Id   int64  //`json:"id"`
	Ad   string `json:"id"`
	Name string `json:"text"`
}

type Select1 struct {
	Title string `json:"title"`
}

type AchEmployee struct { //员工姓名
	Id             int64  `json:"Id"` //`form:"-"`
	Pid            int64  `form:"-"`
	Nickname       string `json:"text"` //这个是侧栏显示的内容
	Level          string `json:"Level"`
	Href           string `json:"href"`
	Icon           string `json:"icon"`
	Image          string `json:"image"`
	Color          string `json:"color"`
	BackColor      string `json:"backColor"`
	IconColor      string `json:"iconColor"`
	IconBackground string `json:"iconBackground"`
}

type AchSecoffice struct { //专业室：水工、施工……
	Id             int64         `json:"Id"` //`form:"-"`
	Pid            int64         `form:"-"`
	Title          string        `json:"text"`
	Tags           [1]string     `json:"tags"` //显示员工数量，如果定义为数值[1]int，则无论如何都显示0，所以要做成字符
	Employee       []AchEmployee `json:"nodes"`
	Level          string        `json:"Level"`
	Href           string        `json:"href"`
	Icon           string        `json:"icon"`
	Image          string        `json:"image"`
	Color          string        `json:"color"`
	BackColor      string        `json:"backColor"`
	IconColor      string        `json:"iconColor"`
	IconBackground string        `json:"iconBackground"`
	Selectable     bool          `json:"selectable"` //这个不能要，虽然没赋值。否则点击node，没反应，即默认false？？
}

type AchDepart struct { //分院：施工预算、水工分院……
	Id int64 `json:"Id"` //`form:"-"`
	// Pid       int64          `form:"-"`
	Title     string         `json:"text"` //这个后面json仅仅对于encode解析有用
	Secoffice []AchSecoffice `json:"nodes"`
	// Employee  []AchEmployee  `json:"nodes"`不能有2个nodes啊，如果部门下有人，只能用secoffice代替员工
	Level      string `json:"Level"`
	Tags       [1]int `json:"tags"` //显示员工数量
	Selectable bool   `json:"selectable"`
}

type Employee struct { //职员的分院和科室属性
	Id         int64  `form:"-"`
	Name       string `json:"Name"`
	Department string `json:"Department"` //分院
	Secoffice  string `json:"Keshi"`      //科室。当controller返回json给view的时候，必须用text作为字段
	Numbers    int    //分值
	Marks      int    //记录个数
}

func (c *IndexController) Cms() {
	c.TplName = "index.html"
}

// 显示侧栏结构，科室里员工
func (c *IndexController) GetIndex() {
	c.Data["IsIndex"] = true
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	c.Data["PageStartTime"] = time.Now()
	achemployee := make([]AchEmployee, 0)
	achsecoffice := make([]AchSecoffice, 0)
	achdepart := make([]AchDepart, 0)
	//由uname取得user,获得user的分院名称
	// user, err := models.GetUserByUsername(uname)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// switch role {
	// case 1: //管理员登录显示的侧栏是全部的
	var depcount int                           //部门人员数
	category1, err := models.GetAdminDepart(0) //得到所有分院（部门）
	if err != nil {
		logs.Error(err)
	}
	for i1, _ := range category1 {
		aa := make([]AchDepart, 1)
		aa[0].Id = category1[i1].Id
		aa[0].Level = "1"
		aa[0].Title = category1[i1].Title                         //分院名称
		category2, err := models.GetAdminDepart(category1[i1].Id) //得到所有科室
		if err != nil {
			logs.Error(err)
		}
		//如果返回科室为空，则直接取得员工
		//这个逻辑判断不完美，如果一个部门即有科室，又有人没有科室属性怎么办，直接挂在部门下的呢？
		//应该是反过来找出所有没有科室字段的人员，把他放在部门下
		if len(category2) > 0 {
			//如果这个部门有科室，则查出科室和科室里的人
			//如果这个部门下无科室
			//或者部门下有科室，但一些人只属于这个部门而无科室属性
			for i2, _ := range category2 {
				bb := make([]AchSecoffice, 1)
				bb[0].Id = category2[i2].Id
				bb[0].Level = "2"
				bb[0].Pid = category1[i1].Id
				bb[0].Title = category2[i2].Title //科室名称
				//根据分院和科室查所有员工
				users, count, err := models.GetUsersbySec(category1[i1].Title, category2[i2].Title) //得到员工姓名
				if err != nil {
					logs.Error(err)
				}
				for i3, _ := range users {
					cc := make([]AchEmployee, 1)
					cc[0].Icon = "glyphicon glyphicon-user" //-flag
					// cc[0].Image = "/static/img/go.jpg"
					// cc[0].Color = "#0969DA" //文字颜色HOTPINKFF69B4_黑色000000
					if users[i3].IsPartyMember == true {
						cc[0].BackColor = "#FFF0F5" // 横条整体背景色
					} else {
						cc[0].BackColor = "#FFF" // 横条整体背景色
					}
					if users[i3].Sex == "女" {
						cc[0].IconColor = "#54aeff"
						cc[0].Color = "#54aeff"
					} else {
						cc[0].IconColor = "#0969DA" //浅蓝色——白色FFFFFF
						cc[0].Color = "#0969DA"
					}
					// cc[0].IconBackground = "#90EE90" //头像颜色 浅绿
					cc[0].Id = users[i3].Id
					cc[0].Level = "3"
					cc[0].Href = users[i3].Ip + ":" + users[i3].Port
					cc[0].Pid = category2[i2].Id
					cc[0].Nickname = users[i3].Nickname //名称
					achemployee = append(achemployee, cc...)
				}
				bb[0].Tags[0] = strconv.Itoa(count)
				bb[0].Employee = achemployee
				bb[0].Selectable = false
				achemployee = make([]AchEmployee, 0) //再把slice置0
				achsecoffice = append(achsecoffice, bb...)
				depcount = depcount + count //部门人员数=科室人员数相加
			}
		}
		//查出所有有这个部门但科室名为空的人员
		//根据分院查所有员工
		users, count, err := models.GetUsersbySecOnly(category1[i1].Title) //得到员工姓名
		if err != nil {
			logs.Error(err)
		}
		for i3, _ := range users {
			dd := make([]AchSecoffice, 1)
			dd[0].Icon = "glyphicon glyphicon-user" //-flag
			if users[i3].IsPartyMember == true {
				dd[0].BackColor = "#FFF0F5" // 横条整体背景色
			} else {
				dd[0].BackColor = "#FFF" // 横条整体背景色
			}
			if users[i3].Sex == "女" {
				dd[0].IconColor = "#54aeff"
				dd[0].Color = "#54aeff"
			} else {
				dd[0].IconColor = "#0969DA" //浅蓝色——白色FFFFFF
				dd[0].Color = "#0969DA"
			}
			dd[0].Id = users[i3].Id
			dd[0].Level = "3"
			dd[0].Href = users[i3].Ip + ":" + users[i3].Port
			dd[0].Pid = category1[i1].Id
			dd[0].Title = users[i3].Nickname //名称——关键，把人员当作科室名
			dd[0].Selectable = true
			achsecoffice = append(achsecoffice, dd...)
		}
		aa[0].Tags[0] = count + depcount
		depcount = 0
		aa[0].Secoffice = achsecoffice
		aa[0].Selectable = false               //点击展开，默认是true
		achsecoffice = make([]AchSecoffice, 0) //再把slice置0
		achdepart = append(achdepart, aa...)
	}
	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		c.TplName = "mobile/mhome.tpl"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "index.tpl"
	}
	c.Data["json"] = achdepart
	// c.TplName = "index.tpl"
}

// 上面那个是显示侧栏
// 这个是显示右侧iframe框架
func (c *IndexController) GetUser() {
	carousels, err := models.GetAdminCarousel()
	if err != nil {
		logs.Error(err)
	}
	mySlice := make([]*models.AdminCarousel, 10)
	if len(carousels) >= 10 {
		mySlice = carousels[:10]
	} else {
		mySlice = carousels[:len(carousels)]
		//
		// for i := 0; i < 10; i++ {
		// 	sum += i
		// }
		// for i, v := range carousels {
		// }
	}
	c.Data["Carousel"] = mySlice
	c.TplName = "index_user.tpl"
}

// 上面那个是显示右侧页面
// 这个是填充数据最新成果、项目、文章
func (c *IndexController) Product() {

}

func (c *IndexController) Calendar() {
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
	c.Data["IsCalendar"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	c.TplName = "index_calendar.tpl"
}

// *******汽车
// 显示页面
func (c *IndexController) GetCarCalendar() {
	// username, role := checkprodRole(c.Ctx)
	// roleint, err := strconv.Atoi(role)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// if role == "1" {
	// 	c.Data["IsAdmin"] = true
	// } else if roleint > 1 && roleint < 5 {
	// 	c.Data["IsLogin"] = true
	// } else {
	// 	c.Data["IsAdmin"] = false
	// 	c.Data["IsLogin"] = false
	// }
	// c.Data["Username"] = username
	c.Data["IsCarCalendar"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	c.TplName = "car_calendar.tpl"
}

// 添加日历
func (c *IndexController) AddCarCalendar() {
	// username, _ := checkprodRole(c.Ctx)
	ip := c.Ctx.Input.IP()
	title := c.GetString("title")
	content := c.GetString("content")
	start := c.GetString("start")
	end := c.GetString("end")
	color := c.GetString("color")
	allday1 := c.GetString("allday")
	var allday bool
	if allday1 == "true" {
		allday = true
	} else {
		allday = false
	}
	// public1 := c.GetString("public")
	var public bool
	// if public1 == "true" {
	public = true
	// } else {
	// 	public = false
	// }
	const lll = "2006-01-02 15:04"
	starttime, err := time.Parse(lll, start)
	// beego.Info(start)
	// beego.Info(starttime)
	if err != nil {
		logs.Error(err)
	}
	endtime, err := time.Parse(lll, end)
	if err != nil {
		logs.Error(err)
	}
	_, err = models.AddCarCalendar(title, content, color, ip, allday, public, starttime, endtime)
	if err != nil {
		logs.Error(err)
	} else {
		c.Data["json"] = title
		c.ServeJSON()
	}
}

// 返回日历json数据
// 如果是管理员，则显示全部，非管理员，显示公开
func (c *IndexController) CarCalendar() {
	start := c.GetString("start")
	end := c.GetString("end")
	const lll = "2006-01-02"
	startdate, err := time.Parse(lll, start)
	if err != nil {
		logs.Error(err)
	}
	enddate, err := time.Parse(lll, end)
	if err != nil {
		logs.Error(err)
	}
	var calendars []*models.CarCalendar
	// _, role := checkprodRole(c.Ctx)
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role == "1" {
		calendars, err = models.GetCarCalendar(startdate, enddate, false)
		if err != nil {
			logs.Error(err)
		}
	} else {
		calendars, err = models.GetCarCalendar(startdate, enddate, true)
		if err != nil {
			logs.Error(err)
		}
	}
	c.Data["json"] = calendars
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

// 修改
func (c *IndexController) UpdateCarCalendar() {
	cid := c.GetString("cid")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	title := c.GetString("title")
	content := c.GetString("content")
	start := c.GetString("start")
	end := c.GetString("end")
	color := c.GetString("color")
	allday1 := c.GetString("allday")
	var allday bool
	if allday1 == "true" {
		allday = true
	} else {
		allday = false
	}
	// public1 := c.GetString("public")
	var public bool
	// if public1 == "true" {
	public = true
	// } else {
	// 	public = false
	// }
	const lll = "2006-01-02 15:04"
	starttime, err := time.Parse(lll, start)
	// beego.Info(start)
	// beego.Info(starttime)
	if err != nil {
		logs.Error(err)
	}
	endtime, err := time.Parse(lll, end)
	if err != nil {
		logs.Error(err)
	}
	err = models.UpdateCarCalendar(cidNum, title, content, color, allday, public, starttime, endtime)
	if err != nil {
		logs.Error(err)
	} else {
		c.Data["json"] = title
		c.ServeJSON()
	}
	// pid := c.Ctx.Input.Param(":id")
	//
	// title := c.GetString("title")
	// code := c.GetString("code")
	// grade := c.GetString("grade")
	// //pid转成64为
	// cidNum, err := strconv.ParseInt(cid, 10, 64)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// gradeNum, err := strconv.Atoi(grade)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// err = models.UpdateAdminCategory(cidNum, title, code, gradeNum)
	// if err != nil {
	// 	logs.Error(err)
	// } else {
	// 	c.Data["json"] = "ok"
	// 	c.ServeJSON()
	// }
}

// 拖曳
func (c *IndexController) DropCarCalendar() {
	id := c.GetString("id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	delta := c.GetString("delta")
	daltaint, err := strconv.Atoi(delta)
	if err != nil {
		logs.Error(err)
	}
	calendar, err := models.GetCarCalendarbyid(idNum)
	if err != nil {
		logs.Error(err)
	}
	t1 := calendar.Starttime.AddDate(0, 0, daltaint)
	t2 := calendar.Endtime.AddDate(0, 0, daltaint)
	err = models.DropCarCalendar(idNum, t1, t2)
	if err != nil {
		logs.Error(err)
	} else {
		c.Data["json"] = calendar.Title
		c.ServeJSON()
	}
}

// resize
func (c *IndexController) ResizeCarCalendar() {
	id := c.GetString("id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	delta := c.GetString("delta")
	delta = delta + "h"
	deltahour, err := time.ParseDuration(delta)
	if err != nil {
		logs.Error(err)
	}
	// starttime.Add(-time.Duration(hours) * time.Hour)
	calendar, err := models.GetCarCalendarbyid(idNum)
	if err != nil {
		logs.Error(err)
	}
	// t1 := calendar.Starttime.Add(deltahour)
	t2 := calendar.Endtime.Add(deltahour)
	err = models.ResizeCarCalendar(idNum, t2)
	if err != nil {
		logs.Error(err)
	} else {
		c.Data["json"] = calendar.Title
		c.ServeJSON()
	}
}

// 删除，如果有下级，一起删除
func (c *IndexController) DeleteCarCalendar() {
	cid := c.GetString("cid")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		logs.Error(err)
	}

	err = models.DeleteCarCalendar(cidNum)
	if err != nil {
		logs.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

// *****会议室
// 显示页面
func (c *IndexController) MeetingroomCalendar() {
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
	c.Data["IsMeetingroomCalendar"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	c.TplName = "meetingroom_calendar.tpl"
}

// 添加日历
func (c *IndexController) AddMeetCalendar() {
	ip := c.Ctx.Input.IP()
	title := c.GetString("title")
	content := c.GetString("content")
	start := c.GetString("start")
	end := c.GetString("end")
	color := c.GetString("color")
	allday1 := c.GetString("allday")
	var allday bool
	if allday1 == "true" {
		allday = true
	} else {
		allday = false
	}
	// public1 := c.GetString("public")
	var public bool
	// if public1 == "true" {
	public = true
	// } else {
	// 	public = false
	// }
	const lll = "2006-01-02 15:04"
	starttime, err := time.Parse(lll, start)
	// beego.Info(start)
	// beego.Info(starttime)
	if err != nil {
		logs.Error(err)
	}
	endtime, err := time.Parse(lll, end)
	if err != nil {
		logs.Error(err)
	}
	_, err = models.AddMeetCalendar(title, content, color, ip, allday, public, starttime, endtime)
	if err != nil {
		logs.Error(err)
	} else {
		c.Data["json"] = title
		c.ServeJSON()
	}
}

// 返回日历json数据
// 如果是管理员，则显示全部，非管理员，显示公开
func (c *IndexController) MeetCalendar() {
	start := c.GetString("start")
	end := c.GetString("end")
	const lll = "2006-01-02"
	startdate, err := time.Parse(lll, start)
	if err != nil {
		logs.Error(err)
	}
	enddate, err := time.Parse(lll, end)
	if err != nil {
		logs.Error(err)
	}
	var calendars []*models.MeetCalendar
	// _, role := checkprodRole(c.Ctx)
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role == "1" {
		calendars, err = models.GetMeetCalendar(startdate, enddate, false)
		if err != nil {
			logs.Error(err)
		}
	} else {
		calendars, err = models.GetMeetCalendar(startdate, enddate, true)
		if err != nil {
			logs.Error(err)
		}
	}
	c.Data["json"] = calendars
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

// 修改
func (c *IndexController) UpdateMeetCalendar() {
	cid := c.GetString("cid")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	title := c.GetString("title")
	content := c.GetString("content")
	start := c.GetString("start")
	end := c.GetString("end")
	color := c.GetString("color")
	allday1 := c.GetString("allday")
	var allday bool
	if allday1 == "true" {
		allday = true
	} else {
		allday = false
	}
	// public1 := c.GetString("public")
	var public bool
	// if public1 == "true" {
	public = true
	// } else {
	// 	public = false
	// }
	const lll = "2006-01-02 15:04"
	starttime, err := time.Parse(lll, start)
	// beego.Info(start)
	// beego.Info(starttime)
	if err != nil {
		logs.Error(err)
	}
	endtime, err := time.Parse(lll, end)
	if err != nil {
		logs.Error(err)
	}
	err = models.UpdateMeetCalendar(cidNum, title, content, color, allday, public, starttime, endtime)
	if err != nil {
		logs.Error(err)
	} else {
		c.Data["json"] = title
		c.ServeJSON()
	}
	// pid := c.Ctx.Input.Param(":id")
	//
	// title := c.GetString("title")
	// code := c.GetString("code")
	// grade := c.GetString("grade")
	// //pid转成64为
	// cidNum, err := strconv.ParseInt(cid, 10, 64)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// gradeNum, err := strconv.Atoi(grade)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// err = models.UpdateAdminCategory(cidNum, title, code, gradeNum)
	// if err != nil {
	// 	logs.Error(err)
	// } else {
	// 	c.Data["json"] = "ok"
	// 	c.ServeJSON()
	// }
}

// 拖曳
func (c *IndexController) DropMeetCalendar() {
	id := c.GetString("id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	delta := c.GetString("delta")
	daltaint, err := strconv.Atoi(delta)
	if err != nil {
		logs.Error(err)
	}
	calendar, err := models.GetMeetCalendarbyid(idNum)
	if err != nil {
		logs.Error(err)
	}
	t1 := calendar.Starttime.AddDate(0, 0, daltaint)
	t2 := calendar.Endtime.AddDate(0, 0, daltaint)
	err = models.DropMeetCalendar(idNum, t1, t2)
	if err != nil {
		logs.Error(err)
	} else {
		c.Data["json"] = calendar.Title
		c.ServeJSON()
	}
}

// resize
func (c *IndexController) ResizeMeetCalendar() {
	id := c.GetString("id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	delta := c.GetString("delta")
	delta = delta + "h"
	deltahour, err := time.ParseDuration(delta)
	if err != nil {
		logs.Error(err)
	}
	// starttime.Add(-time.Duration(hours) * time.Hour)
	calendar, err := models.GetMeetCalendarbyid(idNum)
	if err != nil {
		logs.Error(err)
	}
	// t1 := calendar.Starttime.Add(deltahour)
	t2 := calendar.Endtime.Add(deltahour)
	err = models.ResizeMeetCalendar(idNum, t2)
	if err != nil {
		logs.Error(err)
	} else {
		c.Data["json"] = calendar.Title
		c.ServeJSON()
	}
}

// 删除，如果有下级，一起删除
func (c *IndexController) DeleteMeetCalendar() {
	cid := c.GetString("cid")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		logs.Error(err)
	}

	err = models.DeleteMeetCalendar(cidNum)
	if err != nil {
		logs.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

func (c *IndexController) SearchCalendar() {
	title := c.GetString("title")
	const lll = "2006-01-02"

	var calendars []*models.MeetCalendar
	var err error
	// _, role := checkprodRole(c.Ctx)
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role == "1" {
		calendars, err = models.SearchMeetCalendar(title, false)
		if err != nil {
			logs.Error(err)
		}
	} else {
		calendars, err = models.SearchMeetCalendar(title, true)
		if err != nil {
			logs.Error(err)
		}
	}
	c.Data["json"] = calendars
	c.ServeJSON()
}

// *****订餐
// 显示页面
func (c *IndexController) GetOrderCalendar() {
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
	c.Data["IsOrderCalendar"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	c.TplName = "order_calendar.tpl"
}

// *****考勤
// 显示页面
func (c *IndexController) GetAttendanceCalendar() {
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
	c.Data["IsAttendanceCalendar"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	c.TplName = "attendance_calendar.tpl"
}
