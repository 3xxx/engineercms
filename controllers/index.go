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
	"strconv"
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
	Href     string `json:"href"`
}

type AchSecoffice struct { //专业室：水工、施工……
	Id         int64         `json:"Id"` //`form:"-"`
	Pid        int64         `form:"-"`
	Title      string        `json:"text"`
	Tags       [1]string     `json:"tags"` //显示员工数量，如果定义为数值[1]int，则无论如何都显示0，所以要做成字符
	Employee   []AchEmployee `json:"nodes"`
	Level      string        `json:"Level"`
	Href       string        `json:"href"`
	Selectable bool          `json:"selectable"` //这个不能要，虽然没赋值。否则点击node，没反应，即默认false？？
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
	var depcount int                           //部门人员数
	category1, err := models.GetAdminDepart(0) //得到所有分院（部门）
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
		category2, err := models.GetAdminDepart(category1[i1].Id) //得到所有科室
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(category2)
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
				// beego.Info(category2[i2].Title)
				//根据分院和科室查所有员工
				users, count, err := models.GetUsersbySec(category1[i1].Title, category2[i2].Title) //得到员工姓名
				if err != nil {
					beego.Error(err)
				}
				beego.Info(count)
				for i3, _ := range users {
					cc := make([]AchEmployee, 1)
					cc[0].Id = users[i3].Id
					cc[0].Level = "3"
					cc[0].Href = users[i3].Ip + ":" + users[i3].Port
					cc[0].Pid = category2[i2].Id
					cc[0].Nickname = users[i3].Nickname //名称
					// beego.Info(users[i3].Nickname)
					// cc[0].Selectable = false
					achemployee = append(achemployee, cc...)
				}
				bb[0].Tags[0] = strconv.Itoa(count)
				bb[0].Employee = achemployee
				bb[0].Selectable = false
				achemployee = make([]AchEmployee, 0) //再把slice置0
				achsecoffice = append(achsecoffice, bb...)
				depcount = depcount + count //部门人员数=科室人员数相加
			}
			// aa[0].Tags[0] = depcount
			// aa[0].Secoffice = achsecoffice
			// aa[0].Selectable = false               //点击展开，默认是true
			// achsecoffice = make([]AchSecoffice, 0) //再把slice置0
			// achdepart = append(achdepart, aa...)
		}
		//查出所有有这个部门但科室名为空的人员
		//根据分院查所有员工
		users, count, err := models.GetUsersbySecOnly(category1[i1].Title) //得到员工姓名
		if err != nil {
			beego.Error(err)
		}
		for i3, _ := range users {
			dd := make([]AchSecoffice, 1)
			dd[0].Id = users[i3].Id
			dd[0].Level = "3"
			dd[0].Href = users[i3].Ip + ":" + users[i3].Port
			dd[0].Pid = category1[i1].Id
			dd[0].Title = users[i3].Nickname //名称——关键，把人员当作科室名
			dd[0].Selectable = true
			achsecoffice = append(achsecoffice, dd...)
		}
		aa[0].Tags[0] = count + depcount
		// count = 0
		depcount = 0
		aa[0].Secoffice = achsecoffice
		aa[0].Selectable = false               //点击展开，默认是true
		achsecoffice = make([]AchSecoffice, 0) //再把slice置0
		achdepart = append(achdepart, aa...)
	}

	c.Data["json"] = achdepart
	c.TplName = "index.tpl"
	beego.Info(achdepart)
}

//上面那个是显示侧栏
//这个是显示右侧iframe框架
func (c *IndexController) GetUser() {
	carousels, err := models.GetAdminCarousel()
	if err != nil {
		beego.Error(err)
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

//上面那个是显示右侧页面
//这个是填充数据最新成果、项目、文章
func (c *IndexController) Product() {

}

func (c *IndexController) Calendar() {
	c.TplName = "index_calendar.tpl"
}
