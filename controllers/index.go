//在线成果登记
package controllers

import (
	// json "encoding/json"
	// "fmt"
	"github.com/astaxie/beego"
	// "github.com/tealeg/xlsx"
	// "github.com/bitly/go-simplejson"
	// "io/ioutil"
	// "github.com/astaxie/beego/logs"
	"engineercms/models"
	// "sort"
	// "strconv"
	// "strings"
	// "time"
)

type IndexController struct {
	beego.Controller
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
	Id       int64  `json:"Id"` //`form:"-"`
	Pid      int64  `form:"-"`
	Nickname string `json:"text"` //这个是侧栏显示的内容
	Level    string `json:"Level"`
}

type AchSecoffice struct { //专业室：水工、施工……
	Id       int64         `json:"Id"` //`form:"-"`
	Pid      int64         `form:"-"`
	Title    string        `json:"text"`
	Tags     [1]int        `json:"tags"` //显示员工数量
	Employee []AchEmployee `json:"nodes"`
	Level    string        `json:"Level"`
	// Href       string     `json:"href"` //点击科室，显示总体情况
	// Selectable bool       `json:"selectable"`这个不能要，虽然没赋值。否则点击node，没反应，即默认false？？
}

type AchDepart struct { //分院：施工预算、水工分院……
	Id int64 `json:"Id"` //`form:"-"`
	// Pid       int64          `form:"-"`
	Title     string         `json:"text"` //这个后面json仅仅对于encode解析有用
	Secoffice []AchSecoffice `json:"nodes"`
	Level     string         `json:"Level"`
	// Selectable bool       `json:"selectable"`
}

type Employee struct { //职员的分院和科室属性
	Id         int64  `form:"-"`
	Name       string `json:"Name"`
	Department string `json:"Department"` //分院
	Secoffice  string `json:"Keshi"`      //科室。当controller返回json给view的时候，必须用text作为字段
	Numbers    int    //分值
	Marks      int    //记录个数
}

//显示侧栏结构，科室里员工
func (c *IndexController) GetIndex() {
	role := checkprodRole(c.Ctx)
	if role == 1 {
		c.Data["IsAdmin"] = true
	} else if role <= 1 && role > 5 {
		c.Data["IsLogin"] = true
	} else {
		c.Data["IsAdmin"] = false
		c.Data["IsLogin"] = false
	}

	c.Data["IsIndex"] = true
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role

	achemployee := make([]AchEmployee, 0)
	achsecoffice := make([]AchSecoffice, 0)
	achdepart := make([]AchDepart, 0)
	//由uname取得user,获得user的分院名称
	// user, err := models.GetUserByUsername(uname)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// switch role {
	// case 1: //管理员登录显示的侧栏是全部的
	category1, err := models.GetAdminDepart(0) //得到多个分院
	if err != nil {
		beego.Error(err)
	}
	for i1, _ := range category1 {
		aa := make([]AchDepart, 1)
		aa[0].Id = category1[i1].Id
		aa[0].Level = "1"
		// aa[0].Pid = category[0].Id
		aa[0].Title = category1[i1].Title //分院名称
		// beego.Info(category1[i1].Title)
		category2, err := models.GetAdminDepart(category1[i1].Id) //得到多个科室
		if err != nil {
			beego.Error(err)
		}
		for i2, _ := range category2 {
			bb := make([]AchSecoffice, 1)
			bb[0].Id = category2[i2].Id
			bb[0].Level = "2"
			bb[0].Pid = category1[i1].Id
			bb[0].Title = category2[i2].Title //科室名称
			// beego.Info(category2[i2].Title)
			//根据分院和科室查所有员工
			users, count, err := models.GetUsersbySec(category1[i1].Title, category2[i2].Title) //得到员工姓名
			if err != nil {
				beego.Error(err)
			}
			for i3, _ := range users {
				cc := make([]AchEmployee, 1)
				cc[0].Id = users[i3].Id
				cc[0].Level = "3"
				cc[0].Pid = category2[i2].Id
				cc[0].Nickname = users[i3].Nickname //名称
				// beego.Info(users[i3].Nickname)
				// cc[0].Selectable = false
				achemployee = append(achemployee, cc...)
			}
			bb[0].Tags[0] = count
			bb[0].Employee = achemployee
			achemployee = make([]AchEmployee, 0) //再把slice置0
			achsecoffice = append(achsecoffice, bb...)
		}
		aa[0].Secoffice = achsecoffice
		achsecoffice = make([]AchSecoffice, 0) //再把slice置0
		achdepart = append(achdepart, aa...)
	}
	c.Data["json"] = achdepart
	c.TplName = "index.tpl"
}

//上面那个是显示侧栏
//这个是显示右侧iframe框架
func (c *IndexController) GetUser() {
	c.TplName = "index_user.tpl"
}

//上面那个是显示右侧页面
//这个是填充数据最新成果、项目、文章
func (c *IndexController) Product() {

}

func (c *IndexController) Calendar() {
	c.TplName = "index_calendar.tpl"
}
