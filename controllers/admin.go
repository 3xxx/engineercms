package controllers

import (
	// "crypto/md5"
	// "encoding/hex"
	"encoding/json"
	"github.com/3xxx/engineercms/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

// CMSADMIN API
type AdminController struct {
	beego.Controller
}

//Catalog添加附件链接和设计说明、校审意见
type CatalogLinkCont struct {
	Id            int64     `json:"id"`
	ProjectNumber string    //项目编号
	ProjectName   string    //项目名称
	DesignStage   string    //阶段
	Section       string    //专业
	Tnumber       string    //成果编号
	Name          string    //成果名称
	Category      string    //成果类型
	Page          string    //成果计量单位
	Count         float64   //成果数量
	Drawn         string    //编制、绘制
	Designd       string    //设计
	Checked       string    //校核
	Examined      string    //审查
	Verified      string    //核定
	Approved      string    //批准
	Complex       float64   //难度系数
	Drawnratio    float64   //编制、绘制占比系数
	Designdratio  float64   //设计系数
	Checkedratio  float64   //校核系数
	Examinedratio float64   //审查系数
	Datestring    string    //保存字符型日期
	Date          time.Time `orm:"null;auto_now_add;type(datetime)"`
	Created       time.Time `orm:"auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"auto_now_add;type(datetime)"`
	Author        string    //上传者
	State         int
	Link          []models.CatalogLink
}

//附件链接表
type CatalogLinkEditable struct {
	Id        int64
	CatalogId int64
	Url       string `orm:"sie(500)"`
	Editable  bool
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now_add;type(datetime)"`
}

// @Title getAdminBlock
// @Description get admin page
// @Success 200 {object} success
// @Failure 400 Invalid page
// @Failure 404 page not found
// @router / [get]
func (c *AdminController) Get() {
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid

	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		return
	}
	c.TplName = "admin.tpl"
}

func (c *AdminController) Admin() {
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	if role == "1" {
		id := c.Ctx.Input.Param(":id")
		c.Data["Id"] = id
		switch id {
		case "010": //日历实物
			c.TplName = "admin_calendar.tpl"
		case "011": //基本设置
			c.TplName = "admin_base.tpl"
		case "012": //分级目录
			c.TplName = "admin_category.tpl"
		case "013": //搜索引擎
			c.TplName = "admin_spiderip.tpl"
		case "014": //升级数据库
			c.TplName = "admin_updatedatabase.tpl"
		case "021": //角色权限分配
			c.TplName = "admin_role.tpl"
			// c.TplName = "admin_systemrole.tpl"
		case "022": //用户-oo权限
			c.TplName = "admin_roleoo.tpl"
			// c.TplName = "admin_projectrole.tpl"
		case "023": //角色-oo权限
			c.TplName = "admin_useroo.tpl"
		case "024": //IP地址段
			c.TplName = "admin_ipsegment.tpl"
		case "030": //组织结构
			c.TplName = "admin_department.tpl"
		case "031": //用户-组织结构
			c.TplName = "admin_users.tpl"
		case "032": //角色-用户
			c.TplName = "admin_users.tpl"

		// case "033": //用户组
		// 	c.TplName = "admin_usergroup.tpl"
		case "041": //项目编辑
			c.TplName = "admin_projectstree.tpl"
		case "042": //同步IP async
			c.TplName = "admin_projectsynch.tpl"
		case "043": //项目权限
			c.TplName = "admin_projectsrole.tpl"
		case "044": //项目目录快捷编辑
			c.TplName = "admin_projectseditor.tpl"
		case "051": //merit基本信息
			c.TplName = "admin_meritbasic.tpl"
		case "052": //未提交成果清单
			c.TplName = "admin_meritlist.tpl"
		case "053": //预留
			c.TplName = "admin_merit.tpl"
		case "061": //系统信息日志
			c.TplName = "admin_infolog.tpl"
		case "062": //系统错误日志
			c.TplName = "admin_errlog.tpl"

		default:
			c.TplName = "admin_calendar.tpl"
		}
	}
}

//添加ip地址段
func (c *AdminController) AddIpsegment() {
	_, role, _, _, _ := checkprodRole(c.Ctx)

	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// pid := c.Ctx.Input.Param(":id")
	title := c.Input().Get("title")
	startip := c.Input().Get("startip")
	endip := c.Input().Get("endip")
	iprole := c.Input().Get("iprole")
	iproleNum, err := strconv.Atoi(iprole)
	if err != nil {
		beego.Error(err)
	}
	_, err = models.AddAdminIpsegment(title, startip, endip, iproleNum)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
	Createip()
}

//修改ip地址段
func (c *AdminController) UpdateIpsegment() {
	_, role, _, _, _ := checkprodRole(c.Ctx)

	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// pid := c.Ctx.Input.Param(":id")
	cid := c.Input().Get("cid")
	title := c.Input().Get("title")
	startip := c.Input().Get("startip")
	endip := c.Input().Get("endip")
	iprole := c.Input().Get("iprole")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	iproleNum, err := strconv.Atoi(iprole)
	if err != nil {
		beego.Error(err)
	}
	err = models.UpdateAdminIpsegment(cidNum, title, startip, endip, iproleNum)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
	Createip()
}

//删除ip
func (c *AdminController) DeleteIpsegment() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	ids := c.GetString("ids")
	array := strings.Split(ids, ",")
	for _, v := range array {
		// pid = strconv.FormatInt(v1, 10)
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		err = models.DeleteAdminIpsegment(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	}
	Createip()
}

//查询IP地址段
func (c *AdminController) Ipsegment() {
	ipsegments, err := models.GetAdminIpsegment()
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = ipsegments
	c.ServeJSON()
}

// 1 init函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等
// 2 每个包可以拥有多个init函数
// 3 包的每个源文件也可以拥有多个init函数
// 4 同一个包中多个init函数的执行顺序go语言没有明确的定义(说明)
// 5 不同包的init函数按照包导入的依赖关系决定该初始化函数的执行顺序
// 6 init函数不能被其他函数调用，而是在main函数执行之前，自动被调用
//读取iprole.txt文件，作为全局变量Iprolemaps，供调用访问者ip的权限用
var (
	Iprolemaps map[string]int
)

func init() {
	Iprolemaps = make(map[string]int)
	// f, err := os.OpenFile("./conf/iprole.txt", os.O_RDONLY, 0660)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "%s err read from %s : %s\n", err)
	// }
	// var scanner *bufio.Scanner
	// scanner = bufio.NewScanner(f)
	//从IP地址段数据表读取数据
	ipsegments, err := models.GetAdminIpsegment()
	if err != nil {
		beego.Error(err)
	}
	// for scanner.Scan() {
	//循环行
	argslice := make([]string, 0)
	for _, w := range ipsegments {
		// args := strings.Split(scanner.Text(), " ")
		//分割ip起始、终止和权限
		// maps := processFlag(args)
		// args := [3]string{v.StartIp, v.EndIp, strconv.Itoa(v.Iprole)}
		if w.EndIp != "" {
			argslice = append(argslice, w.StartIp, w.EndIp, strconv.Itoa(w.Iprole))
		} else {
			argslice = append(argslice, w.StartIp, strconv.Itoa(w.Iprole))
		}
		maps := processFlag(argslice)
		for i, v := range maps {
			Iprolemaps[i] = v
		}
		argslice = make([]string, 0)
	}
	// beego.Info(Iprolemaps)
	// }
	// f.Close()
}

func Createip() {
	Iprolemaps = make(map[string]int)
	// f, err := os.OpenFile("./conf/iprole.txt", os.O_RDONLY, 0660)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "%s err read from %s : %s\n", err)
	// }
	// var scanner *bufio.Scanner
	// scanner = bufio.NewScanner(f)
	//从IP地址段数据表读取数据
	ipsegments, err := models.GetAdminIpsegment()
	if err != nil {
		beego.Error(err)
	}
	// for scanner.Scan() {
	//循环行
	argslice := make([]string, 0)
	for _, w := range ipsegments {
		// args := strings.Split(scanner.Text(), " ")
		//分割ip起始、终止和权限
		// maps := processFlag(args)
		// args := [3]string{v.StartIp, v.EndIp, strconv.Itoa(v.Iprole)}
		if w.EndIp != "" {
			argslice = append(argslice, w.StartIp, w.EndIp, strconv.Itoa(w.Iprole))
		} else {
			argslice = append(argslice, w.StartIp, strconv.Itoa(w.Iprole))
		}
		maps := processFlag(argslice)
		for i, v := range maps {
			Iprolemaps[i] = v
		}
		argslice = make([]string, 0)
	}
	// beego.Info(Iprolemaps)
	// }
	// f.Close()
}

//取得访问者的权限
func Getiprole(ip string) (role int) {
	role, ok := Iprolemaps[ip]
	if ok {
		return role
	} else {
		return 5
	}
	//元素查找，这是通用的使用方法
	// v, ok := personDB["test1"]
	// if !ok {
	//     fmt.Println(" 没有找到信息")
	//     return
	// }
}

//获取下一个IP
func nextIp(ip string) string {
	ips := strings.Split(ip, ".")
	var i int
	for i = len(ips) - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(ips[i])
		if n >= 255 {
			//进位
			ips[i] = "1"
		} else {
			//+1
			n++
			ips[i] = strconv.Itoa(n)
			break
		}
	}
	if i == -1 {
		//全部IP段都进行了进位,说明此IP本身已超出范围
		return ""
	}
	ip = ""
	leng := len(ips)
	for i := 0; i < leng; i++ {
		if i == leng-1 {
			ip += ips[i]
		} else {
			ip += ips[i] + "."
		}
	}
	return ip
}

//生成IP地址列表
func processIp(startIp, endIp string) []string {
	var ips = make([]string, 0)
	for ; startIp != endIp; startIp = nextIp(startIp) {
		if startIp != "" {
			ips = append(ips, startIp)
		}
	}
	ips = append(ips, startIp)
	return ips
}

//port代替权限role
func processFlag(arg []string) (maps map[string]int) {
	//开始IP,结束IP
	var startIp, endIp string
	//端口
	var ports []int = make([]int, 0)
	index := 0
	startIp = arg[index]
	si := net.ParseIP(startIp)
	if si == nil {
		//开始IP不合法
		// fmt.Println("'startIp' Setting error")
		beego.Error("开始IP不合法")
		return nil
	}
	index++
	endIp = arg[index]
	ei := net.ParseIP(endIp)
	if ei == nil {
		//未指定结束IP,即只扫描一个IP
		endIp = startIp
	} else {
		index++
	}

	tmpPort := arg[index]
	if strings.Index(tmpPort, "-") != -1 {
		//连续端口
		tmpPorts := strings.Split(tmpPort, "-")
		var startPort, endPort int
		var err error
		startPort, err = strconv.Atoi(tmpPorts[0])
		if err != nil || startPort < 1 || startPort > 65535 {
			//开始端口不合法
			return nil
		}
		if len(tmpPorts) >= 2 {
			//指定结束端口
			endPort, err = strconv.Atoi(tmpPorts[1])
			if err != nil || endPort < 1 || endPort > 65535 || endPort < startPort {
				//结束端口不合法
				// fmt.Println("'endPort' Setting error")
				beego.Error("'endPort' Setting error")
				return nil
			}
		} else {
			//未指定结束端口
			endPort = 65535
		}
		for i := 0; startPort+i <= endPort; i++ {
			ports = append(ports, startPort+i)
		}
	} else {
		//一个或多个端口
		ps := strings.Split(tmpPort, ",")
		for i := 0; i < len(ps); i++ {
			p, err := strconv.Atoi(ps[i])
			if err != nil {
				//端口不合法
				// fmt.Println("'port' Setting error")
				beego.Error("'port' Setting error")
				return nil
			}
			ports = append(ports, p)
		}
	}

	//生成扫描地址列表
	ips := processIp(startIp, endIp)
	il := len(ips)
	m1 := make(map[string]int)
	for i := 0; i < il; i++ {
		pl := len(ports)
		for j := 0; j < pl; j++ {
			//			ipAddrs <- ips[i] + ":" + strconv.Itoa(ports[j])
			//			ipAddrs := ips[i] + ":" + strconv.Itoa(ports[j])
			m1[ips[i]] = ports[j]
		}
	}
	//	fmt.Print(slice1)
	return m1
	//	close(ipAddrs)
}

// @Title Get Category list
// @Description Get Category list by some info
// @Success 200 {object} models.GetAdminCategory
// @Param   id     path   string false       "category id"
// @router /category/:id [get]
//根据数字id或空查询分类，如果有pid，则查询下级，如果pid为空，则查询类别
func (c *AdminController) Category() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	c.Data["Ip"] = c.Ctx.Input.IP()
	// var categories []*models.AdminCategory
	var err error
	if id == "" { //如果id为空，则查询类别
		id = "0"
	}
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	categories, err := models.GetAdminCategory(idNum)
	if err != nil {
		beego.Error(err)
	}

	c.Data["json"] = categories
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

// @Title Get Category by title
// @Description Get Category list by title info
// @Success 200 {object} models.GetAdminCategory
// @Param   title   query   string  false       "title of search"
// @router /categorytitle [get]
//根据名称title查询分级表
func (c *AdminController) CategoryTitle() {
	// title := c.Ctx.Input.Param(":id")
	title := c.Input().Get("title")
	// beego.Info(title)
	categories, err := models.GetAdminCategoryTitle(title)
	// beego.Info(categories)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = categories
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

// @Title Post Category by pid title code grade
// @Description Get Category list by title info
// @Success 200 {object} models.AddAdminCategory
// @Param   pid   query   string  false       "parentid of category"
// @Param   title   query   string  false       "title of category"
// @Param   code   query   string  false       "code of category"
// @Param   grade   query   string  false       "grade of category"
// @router /category/addcategory [post]
//添加
func (c *AdminController) AddCategory() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// pid := c.Ctx.Input.Param(":id")
	pid := c.Input().Get("pid")
	title := c.Input().Get("title")
	code := c.Input().Get("code")
	grade := c.Input().Get("grade")
	//pid转成64为
	var pidNum int64
	var err error
	if pid != "" {
		pidNum, err = strconv.ParseInt(pid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	} else {
		pidNum = 0
	}
	gradeNum, err := strconv.Atoi(grade)
	if err != nil {
		beego.Error(err)
	}
	_, err = models.AddAdminCategory(pidNum, title, code, gradeNum)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

//修改
func (c *AdminController) UpdateCategory() {
	_, role, _, _, _ := checkprodRole(c.Ctx)

	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// pid := c.Ctx.Input.Param(":id")
	cid := c.Input().Get("cid")
	title := c.Input().Get("title")
	code := c.Input().Get("code")
	grade := c.Input().Get("grade")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	gradeNum, err := strconv.Atoi(grade)
	if err != nil {
		beego.Error(err)
	}
	err = models.UpdateAdminCategory(cidNum, title, code, gradeNum)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

//删除，如果有下级，一起删除
func (c *AdminController) DeleteCategory() {
	_, role, _, _, _ := checkprodRole(c.Ctx)

	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	ids := c.GetString("ids")
	array := strings.Split(ids, ",")
	for _, v := range array {
		// pid = strconv.FormatInt(v1, 10)
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//查询下级，即分级
		categories, err := models.GetAdminCategory(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			for _, v1 := range categories {
				err = models.DeleteAdminCategory(v1.Id)
				if err != nil {
					beego.Error(err)
				}
			}
		}
		err = models.DeleteAdminCategory(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	}
}

//********************日历开始**************
//添加日历
func (c *AdminController) AddCalendar() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	start := c.Input().Get("start")
	end := c.Input().Get("end")
	color := c.Input().Get("color")
	allday1 := c.Input().Get("allday")
	var allday bool
	if allday1 == "true" {
		allday = true
	} else {
		allday = false
	}
	public1 := c.Input().Get("public")
	var public bool
	if public1 == "true" {
		public = true
	} else {
		public = false
	}
	const lll = "2006-01-02 15:04"
	var starttime, endtime time.Time
	var err error
	if start != "" {
		starttime, err = time.Parse(lll, start)
		// beego.Info(start)
		// beego.Info(starttime)
		if err != nil {
			beego.Error(err)
		}
	} else {
		starttime = time.Now()
	}
	if end != "" {
		endtime, err = time.Parse(lll, end)
		if err != nil {
			beego.Error(err)
		}
	} else {
		endtime = starttime
	}
	_, err = models.AddAdminCalendar(title, content, color, allday, public, starttime, endtime)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = title
		c.ServeJSON()
	}
}

//返回日历json数据
//如果是管理员，则显示全部，非管理员，显示公开
func (c *AdminController) Calendar() {
	start := c.Input().Get("start")
	end := c.Input().Get("end")
	const lll = "2006-01-02"
	startdate, err := time.Parse(lll, start)
	if err != nil {
		beego.Error(err)
	}
	enddate, err := time.Parse(lll, end)
	if err != nil {
		beego.Error(err)
	}
	var calendars []*models.AdminCalendar
	// _, role := checkprodRole(c.Ctx)
	_, role, _, _, _ := checkprodRole(c.Ctx)
	// c.Data["Username"] = username
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	// c.Data["IsAdmin"] = isadmin
	// c.Data["IsLogin"] = islogin
	// c.Data["Uid"] = uid
	if role == "1" {
		calendars, err = models.GetAdminCalendar(startdate, enddate, false)
		if err != nil {
			beego.Error(err)
		}
	} else {
		calendars, err = models.GetAdminCalendar(startdate, enddate, true)
		if err != nil {
			beego.Error(err)
		}
	}
	c.Data["json"] = calendars
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

//修改
func (c *AdminController) UpdateCalendar() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	cid := c.Input().Get("cid")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	start := c.Input().Get("start")
	end := c.Input().Get("end")
	color := c.Input().Get("color")
	allday1 := c.Input().Get("allday")
	var allday bool
	if allday1 == "true" {
		allday = true
	} else {
		allday = false
	}
	public1 := c.Input().Get("public")
	var public bool
	if public1 == "true" {
		public = true
	} else {
		public = false
	}
	const lll = "2006-01-02 15:04"
	var starttime, endtime time.Time
	if start != "" {
		starttime, err = time.Parse(lll, start)
		// beego.Info(start)
		// beego.Info(starttime)
		if err != nil {
			beego.Error(err)
		}
	} else {
		starttime = time.Now()
	}
	if end != "" {
		endtime, err = time.Parse(lll, end)
		if err != nil {
			beego.Error(err)
		}
	} else {
		endtime = starttime
	}
	err = models.UpdateAdminCalendar(cidNum, title, content, color, allday, public, starttime, endtime)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = title
		c.ServeJSON()
	}
	// pid := c.Ctx.Input.Param(":id")
	//
	// title := c.Input().Get("title")
	// code := c.Input().Get("code")
	// grade := c.Input().Get("grade")
	// //pid转成64为
	// cidNum, err := strconv.ParseInt(cid, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// gradeNum, err := strconv.Atoi(grade)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// err = models.UpdateAdminCategory(cidNum, title, code, gradeNum)
	// if err != nil {
	// 	beego.Error(err)
	// } else {
	// 	c.Data["json"] = "ok"
	// 	c.ServeJSON()
	// }
}

//拖曳
func (c *AdminController) DropCalendar() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	id := c.Input().Get("id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	delta := c.Input().Get("delta")
	daltaint, err := strconv.Atoi(delta)
	if err != nil {
		beego.Error(err)
	}
	calendar, err := models.GetAdminCalendarbyid(idNum)
	if err != nil {
		beego.Error(err)
	}
	t1 := calendar.Starttime.AddDate(0, 0, daltaint)
	t2 := calendar.Endtime.AddDate(0, 0, daltaint)
	err = models.DropAdminCalendar(idNum, t1, t2)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = calendar.Title
		c.ServeJSON()
	}
}

//resize
func (c *AdminController) ResizeCalendar() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	id := c.Input().Get("id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	delta := c.Input().Get("delta")
	delta = delta + "h"
	deltahour, err := time.ParseDuration(delta)
	if err != nil {
		beego.Error(err)
	}
	// starttime.Add(-time.Duration(hours) * time.Hour)
	calendar, err := models.GetAdminCalendarbyid(idNum)
	if err != nil {
		beego.Error(err)
	}
	// t1 := calendar.Starttime.Add(deltahour)
	t2 := calendar.Endtime.Add(deltahour)
	err = models.ResizeAdminCalendar(idNum, t2)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = calendar.Title
		c.ServeJSON()
	}
}

//删除，如果有下级，一起删除
func (c *AdminController) DeleteCalendar() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	cid := c.Input().Get("cid")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	err = models.DeleteAdminCalendar(cidNum)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

func (c *AdminController) SearchCalendar() {
	title := c.Input().Get("title")
	const lll = "2006-01-02"

	var calendars []*models.AdminCalendar
	var err error
	// _, role := checkprodRole(c.Ctx)
	_, role, _, _, _ := checkprodRole(c.Ctx)
	// c.Data["Username"] = username
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	// c.Data["IsAdmin"] = isadmin
	// c.Data["IsLogin"] = islogin
	// c.Data["Uid"] = uid
	if role == "1" {
		calendars, err = models.SearchAdminCalendar(title, false)
		if err != nil {
			beego.Error(err)
		}
	} else {
		calendars, err = models.SearchAdminCalendar(title, true)
		if err != nil {
			beego.Error(err)
		}
	}
	c.Data["json"] = calendars
	c.ServeJSON()
}

//******编辑项目同步ip**********
//根据项目id查询ip
func (c *AdminController) SynchIp() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	c.Data["Ip"] = c.Ctx.Input.IP()
	// var categories []*models.AdminCategory
	var err error
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	synchips, err := models.GetAdminSynchIp(idNum)
	if err != nil {
		beego.Error(err)
	}

	c.Data["json"] = synchips
	c.ServeJSON()
}

//添加
func (c *AdminController) AddsynchIp() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// pid := c.Ctx.Input.Param(":id")
	pid := c.Input().Get("pid")
	username := c.Input().Get("username")
	ip := c.Input().Get("ip")
	port := c.Input().Get("port")
	//pid转成64为
	var pidNum int64
	var err error
	// if pid != "" {
	pidNum, err = strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	// } else {
	// 	pidNum = 0
	// }
	// gradeNum, err := strconv.Atoi(grade)
	// if err != nil {
	// 	beego.Error(err)
	// }
	_, err = models.AddAdminSynchIp(pidNum, username, ip, port)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

//修改
func (c *AdminController) UpdatesynchIp() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// pid := c.Ctx.Input.Param(":id")
	cid := c.Input().Get("cid")
	username := c.Input().Get("username")
	ip := c.Input().Get("ip")
	port := c.Input().Get("port")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	// gradeNum, err := strconv.Atoi(grade)
	// if err != nil {
	// 	beego.Error(err)
	// }
	err = models.UpdateAdminSynchIp(cidNum, username, ip, port)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

//删除
func (c *AdminController) DeletesynchIp() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	ids := c.GetString("ids")
	array := strings.Split(ids, ",")
	for _, v := range array {
		// pid = strconv.FormatInt(v1, 10)
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}

		err = models.DeleteAdminSynchIp(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	}
}

//******后台部门结构********
//根据数字id或空查询分类，如果有pid，则查询下级，如果pid为空，则查询类别
func (c *AdminController) Department() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	c.Data["Ip"] = c.Ctx.Input.IP()
	// var categories []*models.AdminDepartment
	var err error
	if id == "" { //如果id为空，则查询类别
		id = "0"
	}
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	categories, err := models.GetAdminDepart(idNum)
	if err != nil {
		beego.Error(err)
	}

	c.Data["json"] = categories
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

//根据名称title查询分级表
func (c *AdminController) DepartmentTitle() {
	// title := c.Ctx.Input.Param(":id")
	title := c.Input().Get("title")
	// beego.Info(title)
	categories, err := models.GetAdminDepartTitle(title)
	// beego.Info(categories)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = categories
	c.ServeJSON()
	// c.TplName = "admin_category.tpl"
}

//添加
func (c *AdminController) AddDepartment() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// pid := c.Ctx.Input.Param(":id")
	pid := c.Input().Get("pid")
	title := c.Input().Get("title")
	code := c.Input().Get("code")
	//pid转成64为
	var pidNum int64
	var err error
	if pid != "" {
		pidNum, err = strconv.ParseInt(pid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	} else {
		pidNum = 0
	}

	_, err = models.AddAdminDepart(pidNum, title, code)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

//修改
func (c *AdminController) UpdateDepartment() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// pid := c.Ctx.Input.Param(":id")
	cid := c.Input().Get("cid")
	title := c.Input().Get("title")
	code := c.Input().Get("code")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	err = models.UpdateAdminDepart(cidNum, title, code)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

//删除，如果有下级，一起删除
func (c *AdminController) DeleteDepartment() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	ids := c.GetString("ids")
	array := strings.Split(ids, ",")
	for _, v := range array {
		// pid = strconv.FormatInt(v1, 10)
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//查询下级，即分级
		categories, err := models.GetAdminDepart(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			for _, v1 := range categories {
				err = models.DeleteAdminDepart(v1.Id)
				if err != nil {
					beego.Error(err)
				}
			}
		}
		err = models.DeleteAdminDepart(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	}
}

//批量上传首页轮播图片
func (c *AdminController) AddCarousel() {
	// _, role := checkprodRole(c.Ctx)
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role == "1" {
		//获取上传的文件
		_, h, err := c.GetFile("file")
		if err != nil {
			beego.Error(err)
		}
		// var attachment string
		// var filesize int64
		if h != nil {
			//保存附件
			// attachment = h.Filename
			// beego.Info(attachment)
			path := "./attachment/carousel/" // + h.Filename
			url := "/attachment/carousel"    //+ h.Filename
			//存入成果数据库
			//如果编号重复，则不写入，值返回Id值。
			//根据id添加成果code, title, label, principal, content string, projectid int64
			_, err := models.AddAdminCarousel(h.Filename, url)
			if err != nil {
				beego.Error(err)
			} else {
				//存入文件夹
				err = c.SaveToFile("file", path+h.Filename) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
				if err != nil {
					beego.Error(err)
				}
				c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "title": h.Filename, "original": h.Filename, "url": url + "/" + h.Filename}
				c.ServeJSON()
			}
		}
	} else {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
}

//查询所有轮播图片
func (c *AdminController) Carousel() {
	carousels, err := models.GetAdminCarousel()
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = carousels
	c.ServeJSON()
}

//删除选中的轮播图片
func (c *AdminController) DeleteCarousel() {
	ids := c.GetString("ids")
	array := strings.Split(ids, ",")
	for _, v := range array {
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		status, err := models.DelAdminCarouselById(idNum)
		if err == nil && status > 0 {
			c.Data["json"] = "ok"
			c.ServeJSON()
		} else if err != nil {
			beego.Error(err)
			c.Data["json"] = "出错"
			c.ServeJSON()
		}
	}
}

//merit基本信息*************************************
//IP，用户名，姓名，密码
func (c *AdminController) MeritBasic() {
	meritbasic, err := models.GetMeritBasic()
	if err != nil {
		beego.Error(err)
	}
	//取到一个数据，不是数组，所以table无法显示
	//如果数据为空，则构造一个空数据给table，方便修改
	merits := make([]*models.MeritBasic, 1)
	merits[0] = &meritbasic
	c.Data["json"] = &merits
	c.ServeJSON()
}

//在线修改保存某个字段
func (c *AdminController) UpdateMeritBasic() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	name := c.Input().Get("name")
	value := c.Input().Get("value")
	pk := c.Input().Get("pk")
	id, err := strconv.ParseInt(pk, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	err = models.UpdateMeritBasic(id, name, value)
	if err != nil {
		beego.Error(err)
	} else {
		data := "ok!"
		c.Ctx.WriteString(data)
	}
}

//取得成果给table
//成果清单
//未提交status=0和已提交status=1
func (c *AdminController) GetPostMerit() {
	id := c.Ctx.Input.Param(":id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		beego.Error(err)
	}
	var postmerits []*models.PostMerit
	if idint == 0 {
		//取得这个id下的所有merittopic
		postmerits, err = models.GetPostMerits(0)
		if err != nil {
			beego.Error(err)
		}
	} else if idint == 1 {
		postmerits, err = models.GetPostMerits(1)
		if err != nil {
			beego.Error(err)
		}
	}
	link := make([]CatalogLinkCont, 0)
	Attachslice := make([]models.CatalogLink, 0)
	linkarr := make([]CatalogLinkCont, 1)
	attacharr := make([]models.CatalogLink, 1)

	//这里循环，添加附件链接和设计说，校审意见
	for _, w := range postmerits {
		linkarr[0].Id = w.Id
		linkarr[0].ProjectNumber = w.ProjectNumber
		linkarr[0].ProjectName = w.ProjectName
		linkarr[0].DesignStage = w.DesignStage
		linkarr[0].Section = w.Section
		linkarr[0].Tnumber = w.Tnumber
		linkarr[0].Name = w.Name
		linkarr[0].Category = w.Category
		linkarr[0].Page = w.Page
		linkarr[0].Count = w.Count
		linkarr[0].Drawn = w.Drawn
		linkarr[0].Designd = w.Designd
		linkarr[0].Checked = w.Checked
		linkarr[0].Examined = w.Examined
		linkarr[0].Verified = w.Verified
		linkarr[0].Approved = w.Approved
		linkarr[0].Complex = w.Complex
		linkarr[0].Drawnratio = w.Drawnratio
		linkarr[0].Designdratio = w.Designdratio
		linkarr[0].Checkedratio = w.Checkedratio
		linkarr[0].Examinedratio = w.Examinedratio
		linkarr[0].Datestring = w.Datestring
		linkarr[0].Date = w.Date
		linkarr[0].Created = w.Created
		linkarr[0].Updated = w.Updated
		linkarr[0].Author = w.Author
		linkarr[0].State = w.State
		links, err := models.GetCatalogLinks(w.Id)
		if err != nil {
			beego.Error(err)
		}
		for _, v := range links {
			attacharr[0].Url = v.Url
			// beego.Info(v.Url)
			Attachslice = append(Attachslice, attacharr...)
		}
		linkarr[0].Link = Attachslice
		Attachslice = make([]models.CatalogLink, 0)
		link = append(link, linkarr...)
	}
	c.Data["json"] = link //postmerits //products
	c.ServeJSON()
}

//在线修改保存某个字段
func (c *AdminController) ModifyCatalog() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	name := c.Input().Get("name")
	value := c.Input().Get("value")
	pk := c.Input().Get("pk")

	ids := c.GetString("ids")
	if ids != "" { //修改选中。问题，选中的是其他几个，修改当前这个没有选中，则不修改，会不会很奇怪？
		array := strings.Split(ids, ",")
		for _, v := range array {
			// pid = strconv.FormatInt(v1, 10)
			if v != pk { //避免与下面的id重复
				//id转成64位
				idNum, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					beego.Error(err)
				}
				err = models.ModifyCatalog(idNum, name, value)
				if err != nil {
					beego.Error(err)
				} else {
					// data := value
					// c.Ctx.WriteString(data)
				}
			}
		}
	}

	id, err := strconv.ParseInt(pk, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//无论如何都修改当前，重复修改了
	err = models.ModifyCatalog(id, name, value)
	if err != nil {
		beego.Error(err)
	} else {
		data := value
		c.Ctx.WriteString(data)
	}

	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/meritlog.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "修改保存设计记录" + pk)
	logs.Close()
}

//列表显示成果附件
func (c *AdminController) CatalogAttachment() {
	id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	c.Data["Id"] = id
	var idNum int64
	var err error
	// var Url string
	if id != "" {
		//id转成64为
		idNum, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	//由id取得成果状态
	catalog, err := models.GetPostMerit(idNum)
	if err != nil {
		beego.Error(err)
	}
	//根据成果id取得所有附件
	links, err := models.GetCatalogLinks(idNum)
	if err != nil {
		beego.Error(err)
	}

	Attachslice := make([]CatalogLinkEditable, 0)
	attacharr := make([]CatalogLinkEditable, 1)
	if len(links) > 0 {
		for _, v := range links {
			attacharr[0].Id = v.Id
			// linkarr[0].Title = v.FileName
			attacharr[0].Url = v.Url
			attacharr[0].CatalogId = idNum
			if catalog.State == 0 {
				attacharr[0].Editable = true
			} else {
				attacharr[0].Editable = false
			}
			// beego.Info(v.Url)
			Attachslice = append(Attachslice, attacharr...)
		}
		if catalog.State == 0 {
			attacharr[0].Url = "http://"
			attacharr[0].Id = 0
			attacharr[0].CatalogId = idNum
			attacharr[0].Editable = true
			Attachslice = append(Attachslice, attacharr...)
		}
	} else {
		if catalog.State == 0 {
			attacharr[0].Created = time.Now()
			attacharr[0].Updated = time.Now()
			attacharr[0].Editable = true
			attacharr[0].Url = "http://"
			attacharr[0].CatalogId = idNum
			Attachslice = attacharr
		}
	}

	c.Data["json"] = Attachslice
	c.ServeJSON()
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
}

//修改link
func (c *AdminController) ModifyLink() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	name := c.Input().Get("name")
	value := c.Input().Get("value")
	pk := c.Input().Get("pk")

	id, err := strconv.ParseInt(pk, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	cid := c.Input().Get("cid") //成果id
	cidnum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//无论如何都修改当前，重复修改了
	err = models.ModifyCatalogLink(id, cidnum, name, value)
	if err != nil {
		beego.Error(err)
	} else {
		data := value
		c.Ctx.WriteString(data)
	}
}

//提交meritlist给merit，这个是关键代码
func (c *AdminController) SendMeritlist() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// req1 := httplib.Post("http://beego.me/")
	// req1.Param("username","astaxie")
	// req1.Param("password","123456")
	//1——将state从0变为1
	// req.Response())
	pk1 := c.Ctx.Input.RequestBody
	// beego.Info(pk1)
	var ob []models.PostMerit
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	// beego.Info(c.Ctx.Input.RequestBody	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	// pk := c.Input().Get("id")
	// beego.Info(ob.Id)
	// cid, err := strconv.ParseInt(pk, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	for _, v := range ob {
		err := models.UpdatePostMerit(v.Id, "State", "1")
		if err != nil {
			beego.Error(err)
		} else {
			data := "ok!"
			c.Ctx.WriteString(data)
		}
	}
	//2——将meritlist提交给merit服务器
	meritbasic, err := models.GetMeritBasic()
	if err != nil {
		beego.Error(err)
	}
	// postmerit, err := models.GetPostMerit(ob.Id)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// postmerit.Author = meritbasic.Username
	//编码JSON
	// body, err := json.Marshal(postmerit)
	// if err != nil {
	// 	beego.Error(err)
	// }
	req := httplib.Post("http://" + meritbasic.Ip + ":" + meritbasic.Port + "/getecmspost?ecmsip=" + meritbasic.EcmsIp + "&ecmsport=" + meritbasic.EcmsPort)
	// req.Param("username", meritbasic.Username)
	// req.Param("password", meritbasic.Password)
	// req.Param("ecmsip", meritbasic.EcmsIp)
	// beego.Info(meritbasic.EcmsIp)
	// req.Param("ecmsport", meritbasic.EcmsPort)
	// req.Body(body)
	req.Body(pk1) //传body不能带param，太奇怪了。
	_, err = req.String()
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(str)
}

//删除meritlist
func (c *AdminController) DeleteMeritlist() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// pk1 := c.Ctx.Input.RequestBody
	var ob []models.PostMerit
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	for _, v := range ob {
		err := models.DeletePostMerit(v.Id)
		if err != nil {
			beego.Error(err)
		} else {
			data := "ok!"
			c.Ctx.WriteString(data)
		}
	}
}

//回退meritlist已提交给未提交
func (c *AdminController) DownMeritlist() {
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role != "1" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// pk1 := c.Ctx.Input.RequestBody
	var ob []models.PostMerit
	json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	for _, v := range ob {
		err := models.UpdatePostMerit(v.Id, "State", "0")
		if err != nil {
			beego.Error(err)
		} else {
			data := "ok!"
			c.Ctx.WriteString(data)
		}
	}
}

func (c *AdminController) Testdown() {
	filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:]) //  attachment/SL2016测试添加成果/A/FB/1/Your First Meteor Application.pdf
	if err != nil {
		beego.Error(err)
	}
	filename := path.Base(filePath)
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "static/download/"+filename)
}

// @Title get wx projectconfig by projectid
// @Description get wx projectconfig by projectid
// @Param projectid query string true "The id of project"
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /jsoneditor [get]
//给jsoneditor返回json数据
func (c *AdminController) Jsoneditor() {
	id := c.Input().Get("projectid")
	c.TplName = "jsoneditor.tpl"
	c.Data["ProjectId"] = id
}

// @Title get wx projectconfig by projectid
// @Description get wx projectconfig by projectid
// @Param projectid query string true "The id of project"
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /getwxprojectconfig [get]
//给jsoneditor返回json数据
func (c *AdminController) GetWxProjectConfig() {
	id := c.Input().Get("projectid")
	contents, _ := ioutil.ReadFile("./conf/" + id + ".json")
	js, err := simplejson.NewJson([]byte(contents))
	if err != nil {
		panic("json format error")
	}
	c.Data["json"] = js
	c.ServeJSON()
}

// type Projectconfig struct {
// 	Text     string      `json:"text"`
// 	Collapse []Collapse2 `json:"collapse"`
// }

// type Collapse2 struct {
// 	Name  string  `json:"name"`
// 	Icon  string  `json:"icon"`
// 	Value string  `json:"value"`
// 	Title string  `json:"title"`
// 	Cell  []Cell2 `json:"cell"`
// }

// type Cell2 struct {
// 	Icon  string `json:"icon"`
// 	Title string `json:"title"`
// 	Url   string `json:"url"`
// 	Value string `json:"value"`
// 	Label string `json:"label"`
// }

// @Title update wx projectconfig by projectid
// @Description put wx projectconfig by projectid
// @Param projectid query string true "The id of project"
// @Param projectconfig query string true "The json of projectconfig"
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /putwxprojectconfig [post]
//更新json文件
func (c *AdminController) PutWxProjectConfig() {
	_, _, _, isadmin, _ := checkprodRole(c.Ctx)
	// beego.Info()
	// beego.Info(c.Input().Get("projectconfig"))
	// projectconfig := c.Ctx.Input.RequestBody
	// beego.Info(projectconfig)
	// var ob Projectconfig
	// json.Unmarshal(c.Ctx.Input.RequestBody, &ob)
	// beego.Info(ob)

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// defer resp.Body.Close()
	// if err != nil {
	// 	beego.Error(err)
	// }
	// f, err := os.OpenFile("./attachment/onlyoffice/"+onlyattachment.FileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if isadmin {
		id := c.Input().Get("projectid")
		f, err := os.Create("./conf/" + id + ".json")
		if err != nil {
			beego.Error(err)
		}
		defer f.Close()
		// _, err = f.Write(body) //这里直接用resp.Body如何？
		_, err = f.Write(c.Ctx.Input.RequestBody)
		// _, err = f.WriteString(str)
		// _, err = io.Copy(body, f)
		if err != nil {
			beego.Error(err)
		}
		c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "json": id + ".json"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"info": "非admin!"}
		c.ServeJSON()
	}
}

//导入json数据
func (c *AdminController) ImportJson() {
	//获取上传的文件
	_, h, err := c.GetFile("json")
	if err != nil {
		beego.Error(err)
	}
	var path string
	if h != nil {
		//保存附件
		path = "./config/" + h.Filename
		// f.Close()                                             // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("json", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
	}
	contents, _ := ioutil.ReadFile(path)
	js, err := simplejson.NewJson([]byte(contents))
	if err != nil {
		panic("json format error")
	}
	c.Data["json"] = js
	c.ServeJSON()
}

// 根据conf目录下的json.json文件初始化价值结构
func (c *AdminController) InitJson() {
	contents, _ := ioutil.ReadFile("./conf/json.json")
	// var r List6
	// err := json.Unmarshal([]byte(contents), &r)
	// if err != nil {
	// 	fmt.Printf("err was %v", err)
	// }
	js, err := simplejson.NewJson([]byte(contents))
	if err != nil {
		panic("json format error")
	}
	c.Data["json"] = js
	c.ServeJSON()
	//1.获取省水利院
	// text, err := js.Get("text").String()
	// Id, err := models.AddCategory(0, text, "", "", "")
	// if err != nil {
	// 	beego.Error(err)
	// }
}

// func NewJsonStruct() *JsonStruct {
// 	return &JsonStruct{}
// }

// func (self *JsonStruct) Load(filename string, v interface{}) {
// 	data, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		return
// 	}
// 	datajson := []byte(data)
// 	err = json.Unmarshal(datajson, v)
// 	if err != nil {
// 		return
// 	}
// }

// include_once('connect.php');//连接数据库
//   $sql = "select * from calendar";
//    $query = mysql_query($sql);
//     while($row=mysql_fetch_array($query)){
//     $allday = $row['allday'];
//     $is_allday = $allday==1?true:false;
//     $data[] = array(
//     'id' => $row['id'],//事件id
//     'title' => $row['title'],//事件标题
//     'start' => date('Y-m-d H:i',$row['starttime']),//事件开始时间
//      'end' => date('Y-m-d H:i',$row['endtime']),//结束时间
//     'allDay' => $is_allday, //是否为全天事件
//     'color' => $row['color'] //事件的背景色
//      );
//    }
//    echo json_encode($data);

//删除日历
