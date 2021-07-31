//project只能是运行ip权限下操作，即只判断iprole，不提供远程操作
package controllers

import (
	// "encoding/json"
	// "github.com/3xxx/engineercms/controllers/utils"
	"github.com/3xxx/engineercms/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/casbin/beego-orm-adapter"
	// "github.com/casbin/casbin"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ProjController struct {
	beego.Controller
}

//成果页导航条
type Navbartruct struct {
	Id    int64
	Title string
}

type Project1 struct {
	Id        int64 //`json:"id"`_微信小程序的项目选择里要记得修改
	Code      string
	Title     string
	Label     string
	Principal string
	Number    int64
	Created   time.Time
	Updated   time.Time
}

//后端分页的数据结构
type Tableserver struct {
	Rows  []Project1 `json:"rows"`
	Page  int64      `json:"page"`
	Total int64      `json:"total"` //string或int64都行！
}

//项目列表页面
func (c *ProjController) Get() {
	// username, role := checkprodRole(c.Ctx)
	// roleint, err := strconv.Atoi(role)
	// if err != nil {
	// 	beego.Error(err)
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
	c.Data["IsProject"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["IsProjects"] = true
	// beego.Info(c.Ctx.Input.IP())
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		beego.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		c.TplName = "mobile/mprojects.tpl"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "projects.tpl"
	}

	//取得项目类别，给添加项目模态框选项用
	var slice1 []string
	categories, err := models.GetAdminCategory(0)
	if err != nil {
		beego.Error(err)
	}
	for _, v := range categories {
		// aa := make([]string, 1)
		// aa[0] = v.Title //名称
		// cc[0].Selectable = false
		// slice1 = append(slice1, aa...)当aa为slice的时候要...,
		slice1 = append(slice1, v.Title) //当v.title为值的时候不用...
	}
	c.Data["Select2"] = slice1
}

// @Title get cms projectlist...
// @Description get projectlist..
// @Param projectid query string false "The id of project"
// @Param pageNo query string false "The page of projectlist"
// @Param limit query string false "The size of page"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /getprojects [get]
// 分页提供给项目列表页的table中json数据
// http://127.0.0.1/v1/project/getprojects?limit=15&pageNo=1
// 微信小程序里要改id等为小写
func (c *ProjController) GetProjects() {
	// id := c.Ctx.Input.Param(":id")
	id := c.Input().Get("projectid")
	var err error
	var offset, limit1, page1 int64
	// var err error
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 15
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("pageNo")
	if page == "" {
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	searchText := c.Input().Get("searchText")
	projects1 := make([]Project1, 0)
	if id == "" {
		//显示全部
		// var offset int64
		// if page1 <= 1 {
		// 	offset = 0
		// } else {
		// 	offset = (page1 - 1) * limit1
		// }
		// beego.Info(offset)
		// beego.Info(limit1)
		projects, err := models.GetProjectsPage(limit1, offset, searchText)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(projects)
		//记录开始时间
		// start := time.Now()

		//取得每个项目的成果数量
		// projects1 := make([]Project1, 0) //这里不能加*号
		for _, v := range projects {
			aa := make([]Project1, 1)
			aa[0].Id = v.Id
			aa[0].Code = v.Code
			aa[0].Title = v.Title
			aa[0].Label = v.Label
			aa[0].Principal = v.Principal
			//取得项目所有成果——速度太慢
			//修改为一次性取到所有成果，然后循环赋值给aa
			//取项目所有子孙
			//效率太低
			// categories, err := models.GetProjectsbyPid(v.Id)
			// if err != nil {
			// 	beego.Error(err)
			// }
			//根据项目id取得项目下所有成果数量
			count, _, err := models.GetProjProducts(v.Id, 3)
			if err != nil {
				beego.Error(err)
			}
			// var count int
			// for _, m := range products {
			// 	if v.Id == m.ProjectId {
			// 		count = count + 1
			// 	}
			// }
			// beego.Info(count)
			// slice := getsons(v.Id, categories)
			// beego.Info(slice)
			// 如果遍历的当前节点下还有子节点，则进入该子节点进行递归
			// if len(slice) > 0 {
			// 	getprodcount(slice, categories, products, &count)
			// }
			aa[0].Number = count //len(products)
			aa[0].Created = v.Created
			aa[0].Updated = v.Updated
			projects1 = append(projects1, aa...)
		}
		count, err := models.GetProjectsCount(searchText)
		if err != nil {
			beego.Error(err)
		}
		table := Tableserver{projects1, page1, count}

		c.Data["json"] = table
		c.ServeJSON()
		//记录结束时间差
		// elapsed := time.Since(start)
		// beego.Info(elapsed)
	} else {
		idNum, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//根据id查询下级
		projects, err := models.GetProjSonbyId(idNum)
		if err != nil {
			beego.Error(err)
		}
		//取得每个项目的成果数量
		// projects1 := make([]Project1, 0) //这里不能加*号
		for _, v := range projects {
			aa := make([]Project1, 1)
			aa[0].Id = v.Id
			aa[0].Code = v.Code
			aa[0].Title = v.Title
			aa[0].Label = v.Label
			aa[0].Principal = v.Principal

			aa[0].Created = v.Created
			aa[0].Updated = v.Updated
			projects1 = append(projects1, aa...)
		}
		// count, err := models.GetProjectsCount(searchText)
		// if err != nil {
		// 	beego.Error(err)
		// }
		count := int64(len(projects))
		table := Tableserver{projects1, page1, count}

		c.Data["json"] = table
		c.ServeJSON()
	}
}

// @Title get wx projectlist...
// @Description get projectlist..
// @Param projectid query string false "The id of project"
// @Param pageNo query string false "The page of projectlist"
// @Param limit query string false "The size of page"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /getwxprojects [get]
// 取出所有项目列表，table中json数据
//http://127.0.0.1/v1/project/getwxprojects
func (c *ProjController) GetWxProjects() {
	id := c.Input().Get("projectid")
	var err error
	var offset, limit1, page1 int64
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 150
	} else {
		limit1, err = strconv.ParseInt(limit, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("pageNo")
	if page == "" {
		page1 = 1
	} else {
		page1, err = strconv.ParseInt(page, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	searchText := c.Input().Get("searchText")
	projects1 := make([]Project1, 0)
	if id == "" {
		//显示全部
		projects, err := models.GetProjectsPage(limit1, offset, searchText)
		if err != nil {
			beego.Error(err)
		}
		c.Data["json"] = projects
		c.ServeJSON()
	} else {
		idNum, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//根据id查询下级
		projects, err := models.GetProjSonbyId(idNum)
		if err != nil {
			beego.Error(err)
		}
		//取得每个项目的成果数量
		// projects1 := make([]Project1, 0) //这里不能加*号
		for _, v := range projects {
			aa := make([]Project1, 1)
			aa[0].Id = v.Id
			aa[0].Code = v.Code
			aa[0].Title = v.Title
			aa[0].Label = v.Label
			aa[0].Principal = v.Principal

			aa[0].Created = v.Created
			aa[0].Updated = v.Updated
			projects1 = append(projects1, aa...)
		}
		// count, err := models.GetProjectsCount(searchText)
		// if err != nil {
		// 	beego.Error(err)
		// }
		count := int64(len(projects))
		table := Tableserver{projects1, page1, count}

		c.Data["json"] = table
		c.ServeJSON()
	}
}

//分页提供给项目列表页的table中json数据
//根据id查看项目，查出项目当前级和下一级目录
//点击第二级后，用下面的懒加载目录
func (c *ProjController) GetProject() {
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	id := c.Ctx.Input.Param(":id")

	navid1 := beego.AppConfig.String("navigationid1")
	navid2 := beego.AppConfig.String("navigationid2")
	navid3 := beego.AppConfig.String("navigationid3")
	navid4 := beego.AppConfig.String("navigationid4")
	navid5 := beego.AppConfig.String("navigationid5")
	navid6 := beego.AppConfig.String("navigationid6")
	navid7 := beego.AppConfig.String("navigationid7")
	navid8 := beego.AppConfig.String("navigationid8")
	navid9 := beego.AppConfig.String("navigationid9")

	switch id {
	case navid1:
		c.Data["IsNav1"] = true
	case navid2:
		c.Data["IsNav2"] = true
	case navid3:
		c.Data["IsNav3"] = true
	case navid4:
		c.Data["IsNav4"] = true
	case navid5:
		c.Data["IsNav5"] = true
	case navid6:
		c.Data["IsNav6"] = true
	case navid7:
		c.Data["IsNav7"] = true
	case navid8:
		c.Data["IsNav8"] = true
	case navid9:
		c.Data["IsNav9"] = true
	default:
		c.Data["IsProject"] = true
	}
	c.Data["Id"] = id
	// var categories []*models.ProjCategory
	// var err error
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//取项目本身
	category, err := models.GetProj(idNum)
	if err != nil {
		beego.Error(err)
	}

	//记录开始时间
	// start := time.Now()
	//取项目所有子孙
	categories, err := models.GetProjectsbyPid(idNum)
	if err != nil {
		beego.Error(err)
	}
	//记录结束时间差
	// elapsed := time.Since(start)
	// beego.Info(elapsed)
	//根据项目顶级id取得项目下所有成果
	var topprojectid int64
	if category.ParentId != 0 { //如果不是根目录
		parentidpath := strings.Replace(strings.Replace(category.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		topprojectid, err = strconv.ParseInt(patharray[0], 10, 64)
		if err != nil {
			beego.Error(err)
		}
	} else {
		topprojectid = category.Id
	}
	// 取出这个项目下所有成果！！
	_, products, err := models.GetProjProducts(topprojectid, 2)
	if err != nil {
		beego.Error(err)
	}
	//记录结束时间差
	// elapsed = time.Since(start)
	// beego.Info(elapsed)
	//一次性查出所有成果
	//或者存储成果数据的时候存上项目id，相当于加了个索引
	// products, err := models.GetAllProducts()
	// if err != nil {
	// 	beego.Error(err)
	// }
	//根据id取出下级
	cates := getsons(idNum, categories)
	//算出最大级数
	// grade := make([]int, 0)
	// for _, v := range categories {
	// 	grade = append(grade, v.Grade)
	// }
	// height := intmax(grade[0], grade[1:]...)
	var count int
	//取得这个项目目录下的成果数量
	productcount, err := models.GetProducts(idNum)
	if err != nil {
		beego.Error(err)
	}
	count = len(productcount)
	// beego.Info(count)

	for _, proj := range cates {
		id := proj.Id
		for _, m := range products {
			if id == m.ProjectId {
				count = count + 1
			}
		}
		// beego.Info(count)
		slice := getsons(id, categories)
		// 如果遍历的当前节点下还有节点，则进入该节点进行递归
		if len(slice) > 0 {
			getprodcount(slice, categories, products, &count)
		}
	}
	// beego.Info(count)
	var tags [1]string
	tags[0] = strconv.Itoa(count)
	//递归生成目录json
	// root := FileNode1{category.Id, category.Title, "", count, true, []*FileNode1{}}
	root := FileNode1{category.Id, category.Title, "", tags, false, []*FileNode1{}}
	// walk(category.Id, &root)
	// maketreejson1(cates, categories, products, &root)
	maketreejson2(cates, categories, products, &root)
	//记录结束时间差
	// elapsed = time.Since(start)
	// beego.Info(elapsed)
	// beego.Info(root)
	// data, _ := json.Marshal(root)
	c.Data["json"] = root //data
	// c.ServeJSON()
	c.Data["Category"] = category

	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		beego.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		c.TplName = "mobile/mproject.tpl"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "project.tpl"
	}
}

// @Title get cms projecttree...
// @Description get projecttree..
// @Param id path string  true "The id of projecttree"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /getprojecttree/:id [get]
//根据id查看项目，查出项目当前级和下一级目录
//点击第二级后，用下面的懒加载目录
func (c *ProjController) GetProjectTree() {
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	id := c.Ctx.Input.Param(":id")

	navid1 := beego.AppConfig.String("navigationid1")
	navid2 := beego.AppConfig.String("navigationid2")
	navid3 := beego.AppConfig.String("navigationid3")
	navid4 := beego.AppConfig.String("navigationid4")
	navid5 := beego.AppConfig.String("navigationid5")
	navid6 := beego.AppConfig.String("navigationid6")
	navid7 := beego.AppConfig.String("navigationid7")
	navid8 := beego.AppConfig.String("navigationid8")
	navid9 := beego.AppConfig.String("navigationid9")

	switch id {
	case navid1:
		c.Data["IsNav1"] = true
	case navid2:
		c.Data["IsNav2"] = true
	case navid3:
		c.Data["IsNav3"] = true
	case navid4:
		c.Data["IsNav4"] = true
	case navid5:
		c.Data["IsNav5"] = true
	case navid6:
		c.Data["IsNav6"] = true
	case navid7:
		c.Data["IsNav7"] = true
	case navid8:
		c.Data["IsNav8"] = true
	case navid9:
		c.Data["IsNav9"] = true
	default:
		c.Data["IsProject"] = true
	}
	c.Data["Id"] = id
	// var categories []*models.ProjCategory
	// var err error
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//取项目本身
	category, err := models.GetProj(idNum)
	if err != nil {
		beego.Error(err)
	}

	//记录开始时间
	// start := time.Now()
	//取项目所有子孙
	categories, err := models.GetProjectsbyPid(idNum)
	if err != nil {
		beego.Error(err)
	}
	//记录结束时间差
	// elapsed := time.Since(start)
	// beego.Info(elapsed)
	//根据项目顶级id取得项目下所有成果
	var topprojectid int64
	if category.ParentId != 0 { //如果不是根目录
		parentidpath := strings.Replace(strings.Replace(category.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		topprojectid, err = strconv.ParseInt(patharray[0], 10, 64)
		if err != nil {
			beego.Error(err)
		}
	} else {
		topprojectid = category.Id
	}
	_, products, err := models.GetProjProducts(topprojectid, 2)
	if err != nil {
		beego.Error(err)
	}
	//记录结束时间差
	//根据id取出下级
	cates := getsons(idNum, categories)
	//算出最大级数
	// grade := make([]int, 0)
	// for _, v := range categories {
	// 	grade = append(grade, v.Grade)
	// }
	// height := intmax(grade[0], grade[1:]...)
	var count int
	//取得这个项目目录下的成果数量
	productcount, err := models.GetProducts(idNum)
	if err != nil {
		beego.Error(err)
	}
	count = len(productcount)
	for _, proj := range cates {
		id := proj.Id
		for _, m := range products {
			if id == m.ProjectId {
				count = count + 1
			}
		}
		slice := getsons(id, categories)
		// 如果遍历的当前节点下还有节点，则进入该节点进行递归
		if len(slice) > 0 {
			getprodcount(slice, categories, products, &count)
		}
	}
	var tags [1]string
	tags[0] = strconv.Itoa(count)
	//递归生成目录json
	root := EleProjTree{category.Id, category.Title, "", tags, false, []*EleProjTree{}}
	makeeletreejson(cates, categories, products, &root)
	root1 := make([]EleProjTree, 1)
	root1[0] = root
	//记录结束时间差
	c.Data["json"] = root1
	c.ServeJSON()
}

//根据id懒加载项目下级目录——上面那个是显示第一级和第二级目录
func (c *ProjController) GetProjCate() {
	// id := c.Ctx.Input.Param(":id")
	id := c.Input().Get("id")
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	//记录开始时间
	start := time.Now()
	//取所有儿子
	// cates, err := models.GetProjSonbyId(idNum)
	// if err != nil {
	// 	beego.Error(err)
	// }

	//取项目所有子孙
	categories, err := models.GetProjectsbyPid(idNum)
	if err != nil {
		beego.Error(err)
	}

	//根据项目id取得项目下所有成果
	//这里的id必须是项目根目录的id
	//取项目本身
	category, err := models.GetProj(idNum)
	if err != nil {
		beego.Error(err)
	}
	// if proj.ParentIdPath != "" { //如果不是根目录
	parentidpath := strings.Replace(strings.Replace(category.ParentIdPath, "#$", "-", -1), "$", "", -1)
	parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
	patharray := strings.Split(parentidpath1, "-")
	topprojectid, err := strconv.ParseInt(patharray[0], 10, 64)
	if err != nil {
		beego.Error(err)
	}

	// } else {
	// 	topprojectid = proj.Id
	// }
	_, products, err := models.GetProjProducts(topprojectid, 2)
	if err != nil {
		beego.Error(err)
	}
	//记录结束时间差
	elapsed := time.Since(start)
	beego.Info(elapsed)
	//一次性查出所有成果
	//或者存储成果数据的时候存上项目id，相当于加了个索引
	// products, err := models.GetAllProducts()
	// if err != nil {
	// 	beego.Error(err)
	// }
	//根据id取出下级
	cates := getsons(idNum, categories)
	//算出最大级数
	// grade := make([]int, 0)
	// for _, v := range categories {
	// 	grade = append(grade, v.Grade)
	// }
	// height := intmax(grade[0], grade[1:]...)

	//取得项目类别，给添加项目模态框选项用
	var slice []FileNode2
	// var lazyload bool
	// var tags [1]string
	for _, v := range cates {
		var count int
		aa := make([]FileNode2, 1)
		aa[0].Id = v.Id
		aa[0].Title = v.Title //名称
		//是否有儿子_太慢>1S
		// if models.Projhasson(v.Id) {
		// 	aa[0].LazyLoad = true
		// } else {
		// 	aa[0].LazyLoad = false
		// }
		aa[0].LazyLoad = true

		for _, m := range products {
			if v.Id == m.ProjectId {
				count = count + 1
			}
		}
		slice2 := getsons(v.Id, categories)
		// 如果遍历的当前节点下还有节点，则进入该节点进行递归
		if len(slice2) > 0 {
			getprodcount(slice2, categories, products, &count)
		}
		// beego.Info(&count)
		var tags [1]string
		tags[0] = strconv.Itoa(count)
		aa[0].Tags = tags
		// 	cc[0].Selectable = false
		// 	// slice1 = append(slice1, aa...)当aa为slice的时候要...,
		slice = append(slice, aa...) //当v.title为值的时候不用...
	}

	// for _, proj := range cates {
	// 	id := proj.Id
	// 	title := proj.Title
	// 	// code := proj.Code
	// 	//是否有儿子_太慢
	// 	// if models.Projhasson(proj.Id) {
	// 	// 	lazyload = true
	// 	// } else {
	// 	// 	lazyload = false
	// 	// }
	// 	lazyload = true
	// 	// var count int
	// 	// for _, m := range products {
	// 	// 	if id == m.ProjectId {
	// 	// 		count = count + 1
	// 	// 	}
	// 	// }
	// 	// tags[0] = strconv.Itoa(count)
	// 	// 将当前名和id作为子节点添加到目录下
	// 	child := FileNode2{id, title, code, tags, lazyload}
	// 	slice = append(slice, child)
	// }
	//记录结束时间差696ms__不查询儿子则110ms
	// elapsed = time.Since(start) //97MS
	// beego.Info(elapsed)
	//beego.Info(patharray[0])
	//beego.Info(topprojectid)

	c.Data["json"] = slice //data
	c.ServeJSON()
}

//点击项目名称，根据id查看项目下所有成果
//这个只是页面。表格内的数据填充用product controllers里的getprojproducts方法
func (c *ProjController) GetProjProducts() {
	c.Data["IsProject"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = userrole
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid

	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	// var categories []*models.ProjCategory
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//取项目本身
	category, err := models.GetProj(idNum)
	if err != nil {
		beego.Error(err)
	}

	c.Data["Category"] = category
	c.TplName = "project_allproducts.tpl"
}

// @Title get user project editor tree
// @Description get user project editor tree
// @Param pid query string true "The id of project"
// @Success 200 {object} models.GetProjectPage
// @Failure 400 Invalid page supplied
// @Failure 404 project not found
// @router /userprojecteditortree [get]
//用户跳转到自己编辑项目目录页面
func (c *ProjController) UserpProjectEditorTree() {
	_, _, uid, isadmin, isLogin := checkprodRole(c.Ctx)
	if !isLogin {
		// route := c.Ctx.Request.URL.String()
		// c.Data["Url"] = route
		// c.Redirect("/roleerr?url="+route, 302)
		c.Data["json"] = "未登陆"
		c.ServeJSON()
		return
	}
	pid := c.Input().Get("pid")
	//id转成64位
	idNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	projectuser, err := models.GetProjectUser(idNum)
	if err != nil {
		beego.Error(err)
	}
	if projectuser.Id == uid || isadmin {
		c.Data["Id"] = pid
		c.TplName = "user_projecteditortree.tpl"
	} else {
		c.Data["json"] = "非管理员，也非本人"
		c.ServeJSON()
	}
}

// @Title get user project editor tree
// @Description get user project editor tree
// @Param pid query string true "The id of project"
// @Success 200 {object} models.GetProjectPage
// @Failure 400 Invalid page supplied
// @Failure 404 project not found
// @router /userprojectpermission [get]
//用户跳转到自己编辑项目目录页面
func (c *ProjController) UserProjectPermission() {
	_, _, uid, isadmin, isLogin := checkprodRole(c.Ctx)
	if !isLogin {
		// route := c.Ctx.Request.URL.String()
		// c.Data["Url"] = route
		// c.Redirect("/roleerr?url="+route, 302)
		c.Data["json"] = "未登陆"
		c.ServeJSON()
		return
	}
	pid := c.Input().Get("pid")
	//id转成64位
	idNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	projectuser, err := models.GetProjectUser(idNum)
	if err != nil {
		beego.Error(err)
	}
	if projectuser.Id == uid || isadmin {
		c.Data["Id"] = pid
		c.Data["IsAdmin"] = isadmin
		c.TplName = "user_projectpermission.tpl"
	} else {
		c.Data["json"] = "非管理员，也非本人"
		c.ServeJSON()
	}
}

//后台根据id查出项目目录，以便进行编辑
func (c *ProjController) GetProjectCate() {
	id := c.Ctx.Input.Param(":id")
	// id := c.Input().Get("id")
	var err error
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//取项目本身
	category, err := models.GetProj(idNum)
	if err != nil {
		beego.Error(err)
	}
	//取项目所有子孙
	categories, err := models.GetProjectsbyPid(idNum)
	if err != nil {
		beego.Error(err)
	}
	//根据id取出下级
	cates := getsons(idNum, categories)
	//递归生成目录json
	root := FileNode{category.Id, category.Title, category.Code, []*FileNode{}}
	// walk(category.Id, &root)
	maketreejson(cates, categories, &root)

	c.Data["json"] = root //data
	c.ServeJSON()
}

//后台添加项目id的子节点
func (c *ProjController) AddProjectCate() {
	_, _, uid, isadmin, _ := checkprodRole(c.Ctx)
	id := c.Input().Get("id")
	//id转成64位
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	//取项目本身
	category, err := models.GetProj(idNum)
	if err != nil {
		beego.Error(err)
	}
	//根据项目顶级id取得项目下所有成果
	var topprojectid int64
	if category.ParentId != 0 { //如果不是根目录
		parentidpath := strings.Replace(strings.Replace(category.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		topprojectid, err = strconv.ParseInt(patharray[0], 10, 64)
		if err != nil {
			beego.Error(err)
		}
	} else {
		topprojectid = category.Id
	}

	projectuser, err := models.GetProjectUser(topprojectid)
	if err != nil {
		beego.Error(err)
	}

	if projectuser.Id == uid || isadmin {
		//取节点id——
		category, err := models.GetProj(idNum)
		if err != nil {
			beego.Error(err)
		}
		//直接添加子节点
		title := c.Input().Get("name")
		code := c.Input().Get("code")
		parentid := category.Id
		var parentidpath, parenttitlepath string
		// if category.ParentIdPath != "" {
		// 	parentidpath = category.ParentIdPath + "-" + strconv.FormatInt(category.Id, 10)
		// 	parenttitlepath = category.ParentTitlePath + "-" + category.Title

		// } else {
		// 	parentidpath = strconv.FormatInt(category.Id, 10)
		// 	parenttitlepath = category.Title
		// }
		if category.ParentIdPath != "" {
			parentidpath = category.ParentIdPath + "$" + strconv.FormatInt(category.Id, 10) + "#"
			// parenttitlepath = category.ParentTitlePath + "#$" + category.Title + "#"
			parenttitlepath = category.ParentTitlePath + "-" + category.Title
		} else {
			parentidpath = "$" + strconv.FormatInt(category.Id, 10) + "#"
			// parenttitlepath = "$" + category.Title + "#"
			parenttitlepath = category.Title
		}
		grade := category.Grade + 1
		Id, err := models.AddProject(code, title, "", "", parentid, parentidpath, parenttitlepath, grade)
		if err != nil {
			beego.Error(err)
		}
		//添加文件夹
		//根据proj的id——这个放deleteproject前面，否则项目数据表删除了就取不到路径了
		_, DiskDirectory, err := GetUrlPath(idNum)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(DiskDirectory)
		parentpath := DiskDirectory
		// beego.Info(newpath)
		//建立目录，并返回作为父级目录
		err = os.MkdirAll(parentpath+"/"+title, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
		if err != nil {
			beego.Error(err)
		}
		c.Data["json"] = map[string]interface{}{"data": "ok", "id": Id}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"data": "非管理员、非本人"}
		c.ServeJSON()
	}
}

//后台修改项目目录节点名称
func (c *ProjController) UpdateProjectCate() {
	_, _, uid, isadmin, _ := checkprodRole(c.Ctx)
	id := c.Input().Get("id")
	//id转成64位
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//取项目本身
	category, err := models.GetProj(idNum)
	if err != nil {
		beego.Error(err)
	}
	//根据项目顶级id取得项目下所有成果
	var topprojectid int64
	if category.ParentId != 0 { //如果不是根目录
		parentidpath := strings.Replace(strings.Replace(category.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		topprojectid, err = strconv.ParseInt(patharray[0], 10, 64)
		if err != nil {
			beego.Error(err)
		}
	} else {
		topprojectid = category.Id
	}

	projectuser, err := models.GetProjectUser(topprojectid)
	if err != nil {
		beego.Error(err)
	}
	if projectuser.Id == uid || isadmin {
		code := c.Input().Get("code")
		title := c.Input().Get("name")
		//根据proj的id——这个放deleteproject前面，否则项目数据表删除了就取不到路径了
		_, DiskDirectory, err := GetUrlPath(idNum)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(DiskDirectory)
		path1 := DiskDirectory
		newpath1 := filepath.Dir(DiskDirectory)
		// beego.Info(newpath1)
		newpath := newpath1 + "/" + title
		// beego.Info(newpath)
		err = os.Rename(path1, newpath)
		if err != nil {
			beego.Error(err)
		}
		err = models.UpdateProject(idNum, code, title, "", "")
		if err != nil {
			beego.Error(err)
		}
		c.Data["json"] = map[string]interface{}{"data": "ok", "id": id} //id //data
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"data": "非管理员，非本人"}
		c.ServeJSON()
	}
	// c.Data["IsProjects"] = true
}

//根据项目侧栏id查看这个id下的成果，不含子目录中的成果
//任何一级目录下都可以放成果
//这个作废——以product中的GetProjProd()
func (c *ProjController) GetProjProd() {
	id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	c.Data["Id"] = id
	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		beego.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		c.TplName = "mobile/mproject_products.tpl"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "project_products.tpl"
	}
}

//取得某个侧栏id下的导航条
func (c *ProjController) GetProjNav() {
	id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	c.Data["Id"] = id
	// var categories []*models.ProjCategory
	// var err error
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//取项目本身
	proj, err := models.GetProj(idNum)
	if err != nil {
		beego.Error(err)
	}
	//将项目id路径转为名称路径

	//根据proj的parentIdpath
	navslice := make([]Navbartruct, 0)
	nav := make([]Navbartruct, 1)
	var parentidpath, parentidpath1 string
	if proj.ParentIdPath != "" { //如果不是根目录
		// patharray := strings.Split(proj.ParentIdPath, "-")
		parentidpath = strings.Replace(strings.Replace(proj.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 = strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		for _, v := range patharray {
			//pid转成64为
			idNum1, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				beego.Error(err)
			}
			proj1, err := models.GetProj(idNum1)
			if err != nil {
				beego.Error(err)
			}
			if proj1.ParentId != 0 { //如果是项目名称，则不要
				nav[0].Id = proj1.Id
				nav[0].Title = proj1.Title
				navslice = append(navslice, nav...)
			}
		}
		nav[0].Id = proj.Id
		nav[0].Title = proj.Title
		navslice = append(navslice, nav...) //加上自身名称
	}
	c.Data["json"] = navslice
	c.ServeJSON()
}

//添加项目和项目目录、文件夹
func (c *ProjController) AddProject() {
	c.Data["IsProjects"] = true
	_, _, uid, _, isLogin := checkprodRole(c.Ctx)
	if !isLogin {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		return
	}
	// rows := c.Input().Get("rows2[0][0]")
	// beego.Info(rows)
	projcode := c.Input().Get("code")
	projname := c.Input().Get("name")
	projlabel := c.Input().Get("label")
	principal := c.Input().Get("principal")
	//先保存项目名称到数据库，parentid为0，返回id作为下面二级三级四级……的parentid
	//然后递归保存二级三级……到数据库
	//最后递归生成硬盘目录
	Id, err := models.AddProject(projcode, projname, projlabel, principal, 0, "", "", 1)
	if err != nil {
		beego.Error(err)
	}
	_, err = models.AddProjectUser(Id, uid)
	if err != nil {
		beego.Error(err)
	}
	_, err = models.AddProjectLabel(Id, projlabel)
	if err != nil {
		beego.Error(err)
	}
	//根据id查出分级目录的名称、代码和层数
	//如果不建立下级怎样？页面中允许不选择任何
	ids := c.GetString("ids")
	array := strings.Split(ids, ",")

	nodes := make([]*models.AdminCategory, 0) //这样完美解决像上面那样借助aa[0]那样
	grade := make([]int, 0)
	for _, v2 := range array {
		//id转成64位
		idNum, err := strconv.ParseInt(v2, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		category, err := models.GetAdminCategorybyId(idNum)
		if err != nil {
			beego.Error(err)
		}
		nodes = append(nodes, category...)
		grade = append(grade, category[0].Grade)
	}
	//找出最大级数——*****如果用户没有选择下级怎样
	//******中间有空级怎么办？？——递归的时候注意
	// height := intmax(nodes[0].Grade, nodes[1:].Grade...)这句不行
	height := intmax(grade[0], grade[1:]...)
	// beego.Info(height)
	//可以递归了
	idarr := make([]models.Pidstruct, 1)
	idarr[0].ParentId = Id
	idarr[0].ParentTitle = projcode + projname
	idarr[0].ParentIdPath = "" //strconv.FormatInt(Id, 10)
	idarr[0].ParentTitlePath = ""
	// write(idarr, nodes, 2, height)
	models.Insertproj(idarr, nodes, 2, height)
	//递归创建文件夹
	patharr := make([]Pathstruct, 1)
	patharr[0].ParentPath = "./attachment/" + projcode + projname
	create(patharr, nodes, 2, height)
	c.Data["json"] = "ok"
	c.ServeJSON()
}

//根据项目模板添加项目
func (c *ProjController) AddProjTemplet() {
	_, _, uid, _, isLogin := checkprodRole(c.Ctx)
	if !isLogin {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}

	projcode := c.Input().Get("code")
	projname := c.Input().Get("name")
	projlabel := c.Input().Get("label")
	principal := c.Input().Get("principal")
	projid := c.Input().Get("projid")
	ispermission := c.Input().Get("ispermission")
	//code=sl123&name=dada&label=&principal=&projid=25001&ispermission=true
	//先保存项目名称到数据库，parentid为0，返回id作为下面的parentid进行递归
	//根据项目模板id，取出项目目录的json结构
	//然后递归生成硬盘目录
	//然后递归写入数据库
	//根据模板项目id，取出权限数据——将目录id路径转成目录名称路径——查出新项目对应的目录id——写入权限
	//id转成64为
	idNum, err := strconv.ParseInt(projid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//取项目本身
	category, err := models.GetProj(idNum)
	if err != nil {
		beego.Error(err)
	}
	//取项目所有子孙
	categories, err := models.GetProjectsbyPid(idNum)
	if err != nil {
		beego.Error(err)
	}
	//根据id取出下级
	cates := getsons(idNum, categories)
	//递归生成目录json
	root := models.FileNode{category.Id, category.Title, category.Code, 1, []*models.FileNode{}}
	// walk(category.Id, &root)
	maketreejsontemplet(cates, categories, &root)
	//先建立第一层项目编号和名称
	Id, err := models.AddProject(projcode, projname, projlabel, principal, 0, "", "", 1)
	if err != nil {
		beego.Error(err)
	}
	_, err = models.AddProjectUser(Id, uid)
	if err != nil {
		beego.Error(err)
	}
	_, err = models.AddProjectLabel(Id, projlabel)
	if err != nil {
		beego.Error(err)
	}
	//在递归写入数据库
	models.Insertprojtemplet(Id, "$"+strconv.FormatInt(Id, 10)+"#", projcode+projname, root.FileNodes)
	//递归创建文件夹
	// patharr := make([]Pathstruct, 1)
	//先建立第一层文件夹
	pathstring := "./attachment/" + projcode + projname
	//在递归建立下层文件夹
	createtemplet(pathstring, root.FileNodes)

	//权限继承
	var success bool
	var casbinv1 string
	if ispermission == "true" {
		//读取权限里包含projid的
		var paths []beegoormadapter.CasbinRule
		o := orm.NewOrm()
		qs := o.QueryTable("casbin_rule")
		_, err := qs.Filter("PType", "p").Filter("v1__contains", "/"+projid+"/").All(&paths)
		if err != nil {
			beego.Error(err)
		}
		//根据最后的/id/*查出proj的parenttitlepath，修改titlepath的项目编号和项目名称
		//据此再查出新项目对应的parentidpath和id
		//末端加上/id/*存入casbin
		for _, v := range paths {
			array := strings.Split(v.V1, "/")
			// lenth := len(array)
			// if len(array) <= 3 { //根目录
			// id := strings.Replace(strings.Replace(v.V1, "/*", "", -1), "/", "", -1)
			// } else {
			id := array[len(array)-2]
			// parentidpath := strings.Replace(v.V1, "/"+array[len(array)-2]+"/*", "/", -1)
			// }
			//id转成64为
			idNum, err = strconv.ParseInt(id, 10, 64)
			if err != nil {
				beego.Error(err)
			}
			oldproj, err := models.GetProj(idNum)
			if err != nil {
				beego.Error(err)
			}
			if len(array) <= 3 { //根目录
				newproj, err := models.GetProjectCodeTitle(projcode, projname)
				if err != nil {
					beego.Error(err)
				}
				casbinv1 = "/" + strconv.FormatInt(newproj.Id, 10) + "/*"
			} else {
				array1 := strings.Split(oldproj.ParentTitlePath, "-")
				newprojparenttitlepath := strings.Replace(oldproj.ParentTitlePath, array1[0], projcode+projname, -1)
				newproj, err := models.GetProjbyParenttitlepath(newprojparenttitlepath, oldproj.Title)
				if err != nil {
					beego.Error(err)
				}
				casbinv1 = strings.Replace(strings.Replace(strings.Replace(newproj.ParentIdPath, "#$", "/", -1), "$", "/", -1), "#", "/", -1) + strconv.FormatInt(newproj.Id, 10) + "/*"
			}
			success = e.AddPermissionForUser(v.V0, casbinv1, v.V2, v.V3)
			//这里应该用AddPermissionForUser()，来自casbin\rbac_api.go
		}
	}
	if success == true {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "wrong"
	}
	c.ServeJSON()
}

// @Title post wx quickaddproject...
// @Description post quickaddproject..
// @Param projectcode query string true "The projectcode of project"
// @Param projecttitle query string true "The projecttitle of project"
// @Param tempprojid query string true "The tempprojid of project"
// @Param istemppermission query string false "The permission of project"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /quickaddwxproject [post]
//根据项目模板添加项目
func (c *ProjController) QuickAddWxProjTemplet() {
	//content去验证
	// app_version := c.Input().Get("app_version")
	// accessToken, _, _, err := utils.GetAccessToken(app_version)
	// if err != nil {
	// 	beego.Error(err)
	// 	c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": err}
	// 	c.ServeJSON()
	// }
	var user models.User
	var err error
	openID := c.GetSession("openID")
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			beego.Error(err)
		}
	} else {
		// c.Data["json"] = map[string]interface{}{"data": "WRONG", "info": "用户未登录"}
		// c.ServeJSON()
		// return
		user.Id = 9
	}

	projcode := c.Input().Get("projectcode")
	projname := c.Input().Get("projecttitle")
	tempprojid := c.Input().Get("tempprojid")
	istemppermission := c.Input().Get("istemppermission")
	// beego.Info(istemppermission)
	if istemppermission == "" {
		istemppermission = "true"
	}
	//code=sl123&name=dada&label=&principal=&projid=25001&ispermission=true
	//先保存项目名称到数据库，parentid为0，返回id作为下面的parentid进行递归
	//根据项目模板id，取出项目目录的json结构
	//然后递归生成硬盘目录
	//然后递归写入数据库
	//根据模板项目id，取出权限数据——将目录id路径转成目录名称路径——查出新项目对应的目录id——写入权限
	//id转成64为
	idNum, err := strconv.ParseInt(tempprojid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//取项目本身
	category, err := models.GetProj(idNum)
	if err != nil {
		beego.Error(err)
	}
	//取项目所有子孙
	categories, err := models.GetProjectsbyPid(idNum)
	if err != nil {
		beego.Error(err)
	}
	//根据id取出下级
	cates := getsons(idNum, categories)
	//递归生成目录json
	root := models.FileNode{category.Id, category.Title, category.Code, 1, []*models.FileNode{}}
	// walk(category.Id, &root)
	maketreejsontemplet(cates, categories, &root)
	//先建立第一层项目编号和名称
	Id, err := models.AddProject(projcode, projname, "projlabel", "principal", 0, "", "", 1)
	if err != nil {
		beego.Error(err)
	}
	_, err = models.AddProjectUser(Id, user.Id)
	if err != nil {
		beego.Error(err)
	}
	_, err = models.AddProjectLabel(Id, "projlabel")
	if err != nil {
		beego.Error(err)
	}
	//在递归写入数据库
	lastid := models.Insertprojtemplet(Id, "$"+strconv.FormatInt(Id, 10)+"#", projcode+projname, root.FileNodes)
	//递归创建文件夹
	// patharr := make([]Pathstruct, 1)
	//先建立第一层文件夹
	pathstring := "./attachment/" + projcode + projname
	//在递归建立下层文件夹
	createtemplet(pathstring, root.FileNodes)

	projectjson := `{"articleid": "` + strconv.FormatInt(lastid, 10) + `", "articleprojid": "` + strconv.FormatInt(lastid, 10) + `",
	 "collapse": [], "diaryprojid": "` + strconv.FormatInt(Id, 10) + `", "financeprojid": "` + strconv.FormatInt(Id, 10) + `",
  "projectcode": "` + projcode + `", "projectid": "` + strconv.FormatInt(Id, 10) + `", "projecttitle": "` + projname + `", "text": ""}`
	f, err := os.Create("./conf/" + strconv.FormatInt(Id, 10) + ".json")
	if err != nil {
		beego.Error(err)
	}
	defer f.Close()
	// _, err = f.Write(body) //这里直接用resp.Body如何？
	// _, err = f.Write(c.Ctx.Input.RequestBody)
	_, err = f.WriteString(projectjson)
	// _, err = io.Copy(body, f)
	if err != nil {
		beego.Error(err)
	}
	//权限继承
	var success bool
	var casbinv1 string
	if istemppermission == "true" {
		//读取权限里包含projid的
		var paths []beegoormadapter.CasbinRule
		o := orm.NewOrm()
		qs := o.QueryTable("casbin_rule")
		_, err := qs.Filter("PType", "p").Filter("v1__contains", "/"+tempprojid+"/").All(&paths)
		if err != nil {
			beego.Error(err)
		}
		//根据最后的/id/*查出proj的parenttitlepath，修改titlepath的项目编号和项目名称
		//据此再查出新项目对应的parentidpath和id
		//末端加上/id/*存入casbin
		for _, v := range paths {
			array := strings.Split(v.V1, "/")
			// lenth := len(array)
			// if len(array) <= 3 { //根目录
			// id := strings.Replace(strings.Replace(v.V1, "/*", "", -1), "/", "", -1)
			// } else {
			id := array[len(array)-2]
			// parentidpath := strings.Replace(v.V1, "/"+array[len(array)-2]+"/*", "/", -1)
			// }
			//id转成64为
			idNum, err = strconv.ParseInt(id, 10, 64)
			if err != nil {
				beego.Error(err)
			}
			oldproj, err := models.GetProj(idNum)
			if err != nil {
				beego.Error(err)
			}
			if len(array) <= 3 { //根目录
				newproj, err := models.GetProjectCodeTitle(projcode, projname)
				if err != nil {
					beego.Error(err)
				}
				casbinv1 = "/" + strconv.FormatInt(newproj.Id, 10) + "/*"
			} else {
				array1 := strings.Split(oldproj.ParentTitlePath, "-")
				newprojparenttitlepath := strings.Replace(oldproj.ParentTitlePath, array1[0], projcode+projname, -1)
				newproj, err := models.GetProjbyParenttitlepath(newprojparenttitlepath, oldproj.Title)
				if err != nil {
					beego.Error(err)
				}
				casbinv1 = strings.Replace(strings.Replace(strings.Replace(newproj.ParentIdPath, "#$", "/", -1), "$", "/", -1), "#", "/", -1) + strconv.FormatInt(newproj.Id, 10) + "/*"
			}
			success = e.AddPermissionForUser(v.V0, casbinv1, v.V2, v.V3)
			//这里应该用AddPermissionForUser()，来自casbin\rbac_api.go
		}
	}
	if success == true {
		c.Data["json"] = map[string]interface{}{"data": "OK", "info": "SUCCESS"}
	} else {
		c.Data["json"] = map[string]interface{}{"data": "WRONG", "info": "用户未登录"}
	}
	c.ServeJSON()
}

// @Title get project user permission
// @Description get project user permission
// @Param pid query string true "The id of project"
// @Success 200 {object} models.GetProjectPage
// @Failure 400 Invalid page supplied
// @Failure 404 project not found
// @router /projectuserrole [get]
//判断登录用户是管理员还是isme
func (c *ProjController) ProjectUserRole() {
	_, _, uid, _, _ := checkprodRole(c.Ctx)
	pid := c.Input().Get("pid")
	//id转成64位
	idNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	projectuser, err := models.GetProjectUser(idNum)
	if err != nil {
		beego.Error(err)
	}
	if projectuser.Id == uid {
		beego.Info(projectuser.Id)
		c.Data["json"] = map[string]interface{}{"userrole": "isme"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"userrole": "isnotme"}
		c.ServeJSON()
	}
}

// 管理员或本人修改项目名称、负责人等，
// 没有修改硬盘目录，需要手动修改！！
func (c *ProjController) UpdateProject() {
	_, _, uid, isadmin, isLogin := checkprodRole(c.Ctx)
	if !isLogin {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		c.Data["json"] = "未登陆"
		c.ServeJSON()
		return
	}
	pid := c.GetString("pid")
	//id转成64位
	idNum, err := strconv.ParseInt(pid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	projectuser, err := models.GetProjectUser(idNum)
	if err != nil {
		beego.Error(err)
	}

	if projectuser.Id != uid && !isadmin {
		beego.Info(projectuser.Id)
		c.Data["json"] = "非管理员，也非本人"
		c.ServeJSON()
		return
	}
	c.Data["IsProjects"] = true

	projcode := c.Input().Get("code")
	projname := c.Input().Get("name")
	projlabe := c.Input().Get("label")
	principal := c.Input().Get("principal")

	err = models.UpdateProject(idNum, projcode, projname, projlabe, principal)
	if err != nil {
		beego.Error(err)
	}
	//没有修改硬盘目录，需要手动修改！！

	if err != nil {
		c.Data["json"] = "修改出错-写入数据库出错！"
		c.ServeJSON()
	} else {
		c.Data["json"] = "修改成功！"
		c.ServeJSON()
	}
}

//后台删除项目目录节点——这个用删除项目代替了。
//删除多节点
//删除多节点的子节点
// 作废了，用DeleteProject代替了！！！！
func (c *ProjController) DeleteProjectCate() {
	_, _, uid, isadmin, _ := checkprodRole(c.Ctx)
	id := c.Input().Get("id")
	beego.Info(id)
	//id转成64位
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//取项目本身
	category, err := models.GetProj(idNum)
	if err != nil {
		beego.Error(err)
	}
	//根据项目顶级id取得项目下所有成果
	var topprojectid int64
	if category.ParentId != 0 { //如果不是根目录
		parentidpath := strings.Replace(strings.Replace(category.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		topprojectid, err = strconv.ParseInt(patharray[0], 10, 64)
		if err != nil {
			beego.Error(err)
		}
	} else {
		topprojectid = category.Id
	}

	projectuser, err := models.GetProjectUser(topprojectid)
	if err != nil {
		beego.Error(err)
	}
	if projectuser.Id == uid || isadmin {
		ids := c.GetString("ids")
		array := strings.Split(ids, ",")
		beego.Info(array)
		for _, v2 := range array {
			//id转成64为
			idNum, err := strconv.ParseInt(v2, 10, 64)
			if err != nil {
				beego.Error(err)
			}
			//取出所有下级
			cates, err := models.GetProjectsbyPid(idNum)
			if err != nil {
				beego.Error(err)
			}
			for _, v1 := range cates { //删除下级目录
				err = models.DeleteProject(v1.Id)
				if err != nil {
					beego.Error(err)
				}
			}
			//根据proj的id——这个放deleteproject前面，否则项目数据表删除了就取不到路径了
			_, DiskDirectory, err := GetUrlPath(idNum)
			if err != nil {
				beego.Error(err)
			} else if DiskDirectory != "" {
				// beego.Info(DiskDirectory)
				// path := DiskDirectory
				// //直接删除这个文件夹，remove删除文件
				// err = os.RemoveAll(path)
				// if err != nil {
				// 	beego.Error(err)
				// }
				//删除目录本身
				err = models.DeleteProject(idNum)
				if err != nil {
					beego.Error(err)
				}
			}
		}
		c.Data["json"] = map[string]interface{}{"data": "ok"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"data": "非管理员、非本人"}
		c.ServeJSON()
	}
}

//根据id删除proj
//后台删除目录，代替DeleteProjectCate
func (c *ProjController) DeleteProject() {
	_, _, uid, isadmin, _ := checkprodRole(c.Ctx)
	ids := c.Input().Get("ids")
	// beego.Info(ids)
	array := strings.Split(ids, ",")
	//id转成64位
	idNum, err := strconv.ParseInt(array[0], 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//取项目本身
	category, err := models.GetProj(idNum)
	if err != nil {
		beego.Error(err)
	}
	//根据项目顶级id取得项目下所有成果
	var topprojectid int64
	if category.ParentId != 0 { //如果不是根目录
		parentidpath := strings.Replace(strings.Replace(category.ParentIdPath, "#$", "-", -1), "$", "", -1)
		parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
		patharray := strings.Split(parentidpath1, "-")
		topprojectid, err = strconv.ParseInt(patharray[0], 10, 64)
		if err != nil {
			beego.Error(err)
		}
	} else {
		topprojectid = category.Id
	}

	projectuser, err := models.GetProjectUser(topprojectid)
	if err != nil {
		beego.Error(err)
	}
	if projectuser.Id == uid || isadmin {
		//查所有子孙项目，循环删除
		// ids := c.GetString("ids")
		// beego.Info(ids)
		// array := strings.Split(ids, ",")
		//循环项目id
		for _, v := range array {
			//id转成64位
			projid, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				beego.Error(err)
			}
			//根据项目id取得所有子孙id
			projs, err := models.GetProjectsbyPid(projid)
			if err != nil {
				beego.Error(err)
			}
			//循环子孙项目
			for _, w := range projs {
				//取得子孙项目的成果列表
				//根据项目id取得所有成果
				products, err := models.GetProducts(w.Id)
				if err != nil {
					beego.Error(err)
				}
				for _, x := range products {
					//删除子孙成果表
					//循环删除成果
					//根据成果id取得所有附件
					attachments, err := models.GetAttachments(x.Id)
					if err != nil {
						beego.Error(err)
					}
					//删除附件表
					for _, y := range attachments {
						//删除附件数据表
						err = models.DeleteAttachment(y.Id)
						if err != nil {
							beego.Error(err)
						}
					}

					//删除子孙文章表
					//取得成果id下所有文章
					articles, err := models.GetArticles(x.Id)
					if err != nil {
						beego.Error(err)
					}
					//删除文章表
					for _, z := range articles {
						//删除文章数据表
						err = models.DeleteArticle(z.Id)
						if err != nil {
							beego.Error(err)
						}
					}
					//删除成果表自身
					err = models.DeleteProduct(x.Id) //删除成果数据表
					if err != nil {
						beego.Error(err)
					}
				}
				//删除子孙proj数据表
				err = models.DeleteProject(w.Id)
				if err != nil {
					beego.Error(err)
				}
				//删除子孙文章图片文件夹（下面已经全部删除了）
			}
			//根据proj的id——这个放deleteproject前面，否则项目数据表删除了就取不到路径了
			_, DiskDirectory, err := GetUrlPath(projid)
			if err != nil {
				beego.Error(err)
			} else if DiskDirectory != "" {
				// beego.Info(DiskDirectory)
				// path := DiskDirectory
				// //直接删除这个文件夹，remove删除文件
				// err = os.RemoveAll(path)
				// if err != nil {
				// 	beego.Error(err)
				// }
				//20181008删除一个“空”项目，导致attachment文件夹下所有附件都删除的悲惨事件
				//所以，不再提供删除文件夹功能
				//只修改文件夹名称??

				//删除项目自身数据表
				err = models.DeleteProject(projid)
				if err != nil {
					beego.Error(err)
				}
			}
		}
		// if err != nil {
		// 	c.Data["json"] = "no"
		// 	c.ServeJSON()
		// } else {
		c.Data["json"] = map[string]interface{}{"data": "ok"}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"data": "非管理员、非本人"}
		c.ServeJSON()
	}
}

//*******项目日历*****
//添加日历
func (c *ProjController) AddCalendar() {
	var starttime, endtime time.Time
	projectid := c.Ctx.Input.Param(":id")
	pid, err := strconv.ParseInt(projectid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(pid)
	title := c.Input().Get("title")
	content := c.Input().Get("content")
	start := c.Input().Get("start")
	end := c.Input().Get("end")
	color := c.Input().Get("color")
	url := c.Input().Get("url") //"/" +
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
	memorabilia1 := c.Input().Get("memorabilia")
	var memorabilia bool
	if memorabilia1 == "true" {
		memorabilia = true
	} else {
		memorabilia = false
	}
	const lll = "2006-01-02 15:04"
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

	_, err = models.AddProjCalendar(pid, title, content, color, url, allday, public, memorabilia, starttime, endtime)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = title
		c.ServeJSON()
	}
}

func (c *ProjController) GetCalendar() {
	c.Data["ProjectId"] = c.Ctx.Input.Param(":id")
	// beego.Info(c.Ctx.Input.Param(":id"))
	c.TplName = "project_calendar.tpl"
}

//返回日历json数据
//如果是管理员，则显示全部，非管理员，显示公开
func (c *ProjController) Calendar() {
	projectid := c.Ctx.Input.Param(":id")
	pid, err := strconv.ParseInt(projectid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
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
	var calendars []*models.ProjCalendar
	// _, role := checkprodRole(c.Ctx)
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role == "1" {
		calendars, err = models.GetProjCalendar(pid, startdate, enddate, false)
		if err != nil {
			beego.Error(err)
		}
	} else {
		calendars, err = models.GetProjCalendar(pid, startdate, enddate, true)
		if err != nil {
			beego.Error(err)
		}
	}
	c.Data["json"] = calendars
	c.ServeJSON()
	// c.TplName = "Proj_category.tpl"
}

//修改
func (c *ProjController) UpdateCalendar() {
	var starttime, endtime time.Time
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
	url := c.Input().Get("url") //"/" +
	memorabilia1 := c.Input().Get("memorabilia")
	var memorabilia bool
	if memorabilia1 == "true" {
		memorabilia = true
	} else {
		memorabilia = false
	}
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
	err = models.UpdateProjCalendar(cidNum, title, content, color, url, allday, public, memorabilia, starttime, endtime)
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
	// err = models.UpdateProjCategory(cidNum, title, code, gradeNum)
	// if err != nil {
	// 	beego.Error(err)
	// } else {
	// 	c.Data["json"] = "ok"
	// 	c.ServeJSON()
	// }
}

//拖曳
func (c *ProjController) DropCalendar() {
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
	calendar, err := models.GetProjCalendarbyid(idNum)
	if err != nil {
		beego.Error(err)
	}
	t1 := calendar.Starttime.AddDate(0, 0, daltaint)
	t2 := calendar.Endtime.AddDate(0, 0, daltaint)
	err = models.DropProjCalendar(idNum, t1, t2)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = calendar.Title
		c.ServeJSON()
	}
}

//resize
func (c *ProjController) ResizeCalendar() {
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
	calendar, err := models.GetProjCalendarbyid(idNum)
	if err != nil {
		beego.Error(err)
	}
	// t1 := calendar.Starttime.Add(deltahour)
	t2 := calendar.Endtime.Add(deltahour)
	err = models.ResizeProjCalendar(idNum, t2)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = calendar.Title
		c.ServeJSON()
	}
}

//删除，如果有下级，一起删除
func (c *ProjController) DeleteCalendar() {
	cid := c.Input().Get("cid")
	//pid转成64为
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	err = models.DeleteProjCalendar(cidNum)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

//****项目时间轴——大事记
type List struct {
	Name string `json:"name"`
}

type Listimage struct {
	Id        int64    `json:"id"`
	UserNo    string   `json:"userNo"`
	DiagTime  string   `json:"diagTime"`
	DiagDoc   string   `json:"diagDoc"`
	Feature   string   `json:"feature"`
	MatchList string   `json:"matchList"`
	Result    string   `json:"result"`
	Desc      string   `json:"desc"`
	Images    []string `json:"images"`
	Ctime     string   `json:"ctime"`
	Utime     string   `json:"utime"`
}

//项目时间轴
func (c *ProjController) ProjectTimeline() {
	projectid := c.Ctx.Input.Param(":id")
	pid, err := strconv.ParseInt(projectid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	var calendars []*models.ProjCalendar
	// var err error
	// _, role := checkprodRole(c.Ctx)
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role == "1" { //显示公开和私有的大事记
		calendars, err = models.GetAllProjCalendar(pid, false)
		if err != nil {
			beego.Error(err)
		}
	} else { //只显示公开的大事记
		calendars, err = models.GetAllProjCalendar(pid, true)
		if err != nil {
			beego.Error(err)
		}
	}
	count := len(calendars)
	// beego.Info(count)
	// count1 := strconv.Itoa(count)
	// count2, err := strconv.ParseInt(count1, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	project, err := models.GetProj(pid)
	c.Data["ProjectId"] = c.Ctx.Input.Param(":id")
	c.Data["ProjectTile"] = project.Title
	c.Data["Count"] = count
	c.TplName = "project_timeline.tpl"
	// c.Data["json"] = map[string]interface{}{
	// 	"id":    2,
	// 	"name":  "111",
	// 	"price": "demo.jpg",
	// }
}

//要分页
func (c *ProjController) Timeline() {
	// page := c.Input().Get("p")
	// pagenum, err := strconv.Atoi(page)
	// if err != nil {
	// 	beego.Error(err)
	// }
	projectid := c.Ctx.Input.Param(":id")
	pid, err := strconv.ParseInt(projectid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	var calendars []*models.ProjCalendar
	// var err error
	// _, role := checkprodRole(c.Ctx)
	_, role, _, _, _ := checkprodRole(c.Ctx)
	if role == "1" { //显示公开和私有的大事记
		calendars, err = models.GetAllProjCalendar(pid, false)
		if err != nil {
			beego.Error(err)
		}
	} else { //只显示公开的大事记
		calendars, err = models.GetAllProjCalendar(pid, true)
		if err != nil {
			beego.Error(err)
		}
	}
	count := len(calendars)
	// beego.Info(count)
	count1 := strconv.Itoa(count)
	count2, err := strconv.ParseInt(count1, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	// sets this.Data["paginator"] with the current offset (from the url query param)
	postsPerPage := 2
	paginator := pagination.SetPaginator(c.Ctx, postsPerPage, count2)
	// beego.Info(c.Ctx)
	// beego.Info(paginator.Offset()) //0
	// p := pagination.NewPaginator(c.Ctx.Request, 10, 9)
	// beego.Info(p.Offset())   0
	// fetch the next 5 posts
	if role == "1" { //显示公开和私有的大事记
		calendars, err = models.ListPostsByOffsetAndLimit(pid, paginator.Offset(), postsPerPage, true)
		if err != nil {
			beego.Error(err)
		}
	} else { //显示公开的大事记
		calendars, err = models.ListPostsByOffsetAndLimit(pid, paginator.Offset(), postsPerPage, false)
		if err != nil {
			beego.Error(err)
		}
	}

	// start := "2016-11-01" //c.Input().Get("start")
	// end := "2017-04-10"   //c.Input().Get("end")

	// const lll = "2006-01-02"
	// startdate, err := time.Parse(lll, start)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// enddate, err := time.Parse(lll, end)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// var calendars []*models.ProjCalendar
	// _, role := checkprodRole(c.Ctx)
	// if role == 1 { //显示公开和私有的大事记
	// 	calendars, err = models.GetProjCalendar(false, true)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// } else { //显示公开的大事记
	// 	calendars, err = models.GetProjCalendar(true, true)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// }
	c.Data["json"] = calendars
	c.ServeJSON()
}

//应该将日历改为froala，那么这个就可以淘汰了。
func (c *ProjController) UploadImage() {
	id := c.Input().Get("pid")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//由proj id取得url
	Url, DiskDirectory, err := GetUrlPath(idNum)
	if err != nil {
		beego.Error(err)
	}
	//保存上传的图片
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}

	var filesize int64
	fileSuffix := path.Ext(h.Filename)
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix
	year, month, _ := time.Now().Date()

	err = os.MkdirAll(DiskDirectory+"/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		beego.Error(err)
	}
	path1 := DiskDirectory + "/" + strconv.Itoa(year) + month.String() + "/" + newname //h.Filename
	Url1 := Url + "/" + strconv.Itoa(year) + month.String() + "/"
	err = c.SaveToFile("file", path1) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
	if err != nil {
		beego.Error(err)
	}
	filesize, _ = FileSize(path1)
	filesize = filesize / 1000.0
	c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "url": Url1 + newname, "title": h.Filename, "original": h.Filename}
	c.ServeJSON()
}

//求出[]int最大值
func intmax(first int, args ...int) int {
	for _, v := range args {
		if first < v {
			first = v
		}
	}
	return first
}

//递归将目录写入数据库
func write(pid []models.Pidstruct, nodes []*models.AdminCategory, igrade, height int) (cid []models.Pidstruct) {
	for _, v := range pid {
		for _, v1 := range nodes {
			if v1.Grade == igrade {
				title := v1.Title
				code := v1.Code
				parentid := v.ParentId

				var parentidpath string
				var parenttitlepath string
				// if v.ParentIdPath != "" {
				// 	parentidpath = v.ParentIdPath + "-" + strconv.FormatInt(v.ParentId, 10)
				// 	parenttitlepath = v.ParentTitlePath + "-" + v.ParentTitle
				// } else {
				// 	parentidpath = strconv.FormatInt(v.ParentId, 10)
				// 	parenttitlepath = v.ParentTitle
				// }
				if v.ParentIdPath != "" {
					parentidpath = v.ParentIdPath + "$" + strconv.FormatInt(v.ParentId, 10) + "#"
					// parenttitlepath = v.ParentTitlePath + "#$" + v.ParentTitle + "#"
					parenttitlepath = v.ParentTitlePath + "-" + v.ParentTitle
				} else {
					parentidpath = "$" + strconv.FormatInt(v.ParentId, 10) + "#"
					// parenttitlepath = "$" + v.ParentTitle + "#"
					parenttitlepath = v.ParentTitle
				}

				grade := igrade
				Id, err := models.AddProject(code, title, "", "", parentid, parentidpath, parenttitlepath, grade)
				if err != nil {
					beego.Error(err)
				}
				var cid1 models.Pidstruct
				cid1.ParentId = Id
				cid1.ParentTitle = title
				cid1.ParentIdPath = parentidpath
				cid1.ParentTitlePath = parenttitlepath
				cid = append(cid, cid1) //每次要清0吗？
			}
		}
	}
	igrade = igrade + 1
	if igrade <= height {
		write(cid, nodes, igrade, height)
	}
	return
}

//树状目录数据——带成果数量和懒加载
type FileNode2 struct {
	Id    int64  `json:"id"`
	Title string `json:"text"`
	// Code     string    `json:"code"` //分级目录代码
	Tags     [1]string `json:"tags"` //显示员工数量，如果定义为数值[1]int，则无论如何都显示0，所以要做成字符
	LazyLoad bool      `json:"lazyLoad"`
}

//树状目录数据——带成果数量
type FileNode1 struct {
	Id        int64        `json:"id"`
	Title     string       `json:"text"`
	Code      string       `json:"code"` //分级目录代码
	Tags      [1]string    `json:"tags"` //显示员工数量，如果定义为数值[1]int，则无论如何都显示0，所以要做成字符
	LazyLoad  bool         `json:"lazyLoad"`
	FileNodes []*FileNode1 `json:"nodes"`
}

//vue.js-project树状目录数据——带成果数量
type EleProjTree struct {
	Id       int64          `json:"id"`
	Label    string         `json:"label"`
	Code     string         `json:"code"` //分级目录代码
	Tags     [1]string      `json:"tags"` //显示员工数量，如果定义为数值[1]int，则无论如何都显示0，所以要做成字符
	Lazy     bool           `json:"lazy"`
	Children []*EleProjTree `json:"children"`
}

//树状目录数据
type FileNode struct {
	Id        int64       `json:"id"`
	Title     string      `json:"text"`
	Code      string      `json:"code"` //分级目录代码
	FileNodes []*FileNode `json:"nodes"`
}

//递归构造项目树状目录——反复查询数据库，速度太慢，淘汰
// func walk(id int64, node *FileNode) {
// 	//列出当前id下子节点，不要列出孙节点……
// 	files, err := models.GetProjSonbyId(id)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	// 遍历目录
// 	for _, proj := range files {
// 		id := proj.Id
// 		title := proj.Title
// 		code := proj.Code
// 		// 将当前名和id作为子节点添加到目录下
// 		child := FileNode{id, title, code, []*FileNode{}}
// 		node.FileNodes = append(node.FileNodes, &child)
// 		// 如果遍历的当前节点下还有节点，则进入该节点进行递归
// 		if models.Projhasson(proj.Id) {
// 			walk(proj.Id, &child)
// 		}
// 	}
// 	return
// }

//递归构造项目树状目录_带成果数量_懒加载只显示一层
func maketreejson3(cates, categories []*models.Project, products []*models.Product, node *FileNode2) {
	// 遍历目录
	for _, proj := range cates {
		id := proj.Id
		title := proj.Title
		// code := proj.Code
		var count int
		for _, m := range products {
			if id == m.ProjectId {
				count = count + 1
			}
		}
		// beego.Info(count)
		slice := getsons(id, categories)
		// 如果遍历的当前节点下还有节点，则进入该节点进行递归
		if len(slice) > 0 {
			getprodcount(slice, categories, products, &count)
		}
		// beego.Info(&count)
		var tags [1]string
		tags[0] = strconv.Itoa(count)
		// 将当前名和id作为子节点添加到目录下
		child := FileNode2{id, title, tags, true}
		node = &child
	}
	return
}

//递归构造项目树状目录_带成果数量_只显示项目层和下面第一层
func maketreejson2(cates, categories []*models.Project, products []*models.Product, node *FileNode1) {
	// 遍历目录
	for _, proj := range cates {
		id := proj.Id
		title := proj.Title
		code := proj.Code
		var count int
		for _, m := range products {
			if id == m.ProjectId {
				count = count + 1
			}
		}
		// beego.Info(count)
		slice := getsons(id, categories)
		// 如果遍历的当前节点下还有节点，则进入该节点进行递归
		if len(slice) > 0 {
			getprodcount(slice, categories, products, &count)
		}
		// beego.Info(&count)
		var tags [1]string
		tags[0] = strconv.Itoa(count)
		// 将当前名和id作为子节点添加到目录下
		child := FileNode1{id, title, code, tags, true, []*FileNode1{}}
		node.FileNodes = append(node.FileNodes, &child)

		// 如果遍历的当前节点下还有节点，则进入该节点进行递归
		// if len(slice) > 0 {
		// 	maketreejson2(slice, categories, products, &child)
		// }
	}
	return
}

//递归构造项目树状目录_带成果数量_只显示项目层和下面第一层
func makeeletreejson(cates, categories []*models.Project, products []*models.Product, node *EleProjTree) {
	// 遍历目录
	for _, proj := range cates {
		id := proj.Id
		title := proj.Title
		code := proj.Code
		var count int
		for _, m := range products {
			if id == m.ProjectId {
				count = count + 1
			}
		}
		// beego.Info(count)
		slice := getsons(id, categories)
		// 如果遍历的当前节点下还有节点，则进入该节点进行递归
		if len(slice) > 0 {
			getprodcount(slice, categories, products, &count)
		}
		// beego.Info(&count)
		var tags [1]string
		tags[0] = strconv.Itoa(count)
		// 将当前名和id作为子节点添加到目录下
		child := EleProjTree{id, title, code, tags, true, []*EleProjTree{}}
		node.Children = append(node.Children, &child)

		// 如果遍历的当前节点下还有节点，则进入该节点进行递归
		// if len(slice) > 0 {
		// 	maketreejson2(slice, categories, products, &child)
		// }
	}
	return
}

//递归构造项目树状目录_带成果数量
func maketreejson1(cates, categories []*models.Project, products []*models.Product, node *FileNode1) {
	// 遍历目录
	for _, proj := range cates {
		id := proj.Id
		title := proj.Title
		code := proj.Code
		var count int
		for _, m := range products {
			if id == m.ProjectId {
				count = count + 1
			}
		}
		// beego.Info(count)
		slice := getsons(id, categories)
		// 如果遍历的当前节点下还有节点，则进入该节点进行递归
		if len(slice) > 0 {
			getprodcount(slice, categories, products, &count)
		}
		// beego.Info(&count)
		var tags [1]string
		tags[0] = strconv.Itoa(count)
		// 将当前名和id作为子节点添加到目录下
		child := FileNode1{id, title, code, tags, true, []*FileNode1{}}
		node.FileNodes = append(node.FileNodes, &child)

		// 如果遍历的当前节点下还有节点，则进入该节点进行递归
		if len(slice) > 0 {
			maketreejson1(slice, categories, products, &child)
		}
	}
	return
}

//递归构造项目树状目录_不带标签，用于后台目录编辑使用
func maketreejson(cates, categories []*models.Project, node *FileNode) {
	// 遍历目录
	for _, proj := range cates {
		id := proj.Id
		title := proj.Title
		code := proj.Code
		// 将当前名和id作为子节点添加到目录下
		child := FileNode{id, title, code, []*FileNode{}}
		node.FileNodes = append(node.FileNodes, &child)
		slice := getsons(id, categories)
		// 如果遍历的当前节点下还有节点，则进入该节点进行递归
		if len(slice) > 0 {
			maketreejson(slice, categories, &child)
		}
	}
	return
}

//递归构造项目树状目录_不带标签，用于项目模板生成项目使用
func maketreejsontemplet(cates, categories []*models.Project, node *models.FileNode) {
	// 遍历目录
	for _, proj := range cates {
		id := proj.Id
		title := proj.Title
		code := proj.Code
		grade := proj.Grade
		// 将当前名和id作为子节点添加到目录下
		child := models.FileNode{id, title, code, grade, []*models.FileNode{}}
		node.FileNodes = append(node.FileNodes, &child)
		slice := getsons(id, categories)
		// 如果遍历的当前节点下还有节点，则进入该节点进行递归
		if len(slice) > 0 {
			maketreejsontemplet(slice, categories, &child)
		}
	}
	return
}

//取得树状目录下的成果数量
func getprodcount(cates, categories []*models.Project, products []*models.Product, count *int) {
	for _, k := range cates {
		for _, m := range products {
			if k.Id == m.ProjectId {
				*count = *count + 1
				// beego.Info(count)
			}
		}
		slice := getsons(k.Id, categories)
		// 如果遍历的当前节点下还有节点，则进入该节点进行递归
		if len(slice) > 0 {
			getprodcount(slice, categories, products, count)
		}
	}
	return
}

//取得数组的下级目录
func getsons(idNum int64, categories []*models.Project) (slice []*models.Project) {
	// slice := make([]*models.Project, 0)
	for _, k := range categories {
		if k.ParentId == idNum {
			slice = append(slice, k)
		}
	}
	return slice
}

type Pathstruct struct {
	ParentPath string
}

//根据分级目录递归建立文件夹
func create(path []Pathstruct, nodes []*models.AdminCategory, igrade, height int) (cpath []Pathstruct) {
	for _, v := range path {
		for _, v1 := range nodes {
			if v1.Grade == igrade {
				title := v1.Title
				parentpath := v.ParentPath
				//建立目录，并返回作为父级目录
				err := os.MkdirAll(parentpath+"/"+title, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
				if err != nil {
					beego.Error(err)
				}

				var cpath1 Pathstruct
				cpath1.ParentPath = parentpath + "/" + title
				cpath = append(cpath, cpath1) //每次要清0吗？
			}
		}
	}
	igrade = igrade + 1
	if igrade <= height {
		create(cpath, nodes, igrade, height)
	}
	return
}

//根据项目模板递归建立文件夹
func createtemplet(parentpath string, nodes []*models.FileNode) {
	for _, v1 := range nodes {
		//建立目录，并返回作为父级目录
		err := os.MkdirAll(parentpath+"/"+v1.Title, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
		if err != nil {
			beego.Error(err)
		}
		ParentPath := parentpath + "/" + v1.Title
		if len(v1.FileNodes) > 0 {
			nodes1 := v1.FileNodes
			createtemplet(ParentPath, nodes1)
		}
	}
	return
}

// {
//   "id": 33,
//   "text": "test3",
//   "nodes": [
//     {
//       "id": 34,
//       "text": "项目建议书",
//       "nodes": [
//         {
//           "id": 36,
//           "text": "综合",
//           "nodes": [
//             {
//               "id": 40,
//               "text": "设计大纲",
//               "nodes": []
//             },
//             {
//               "id": 41,
//               "text": "计算书",
//               "nodes": []
//             }
//           ]
//         },
// type FileNode struct {
// 	Name      string      `json:"name"`
// 	Path      string      `json:"path"`
// 	FileNodes []*FileNode `json:"children"`
// }

// func walkback(path string, info os.FileInfo, node *FileNode) {
// 	// 列出当前目录下的所有目录、文件
// 	files := listFiles(path)
// 	// 遍历这些文件
// 	for _, filename := range files {
// 		// 拼接全路径
// 		fpath := filepath.Join(path, filename)
// 		// 构造文件结构
// 		fio, _ := os.Lstat(fpath)
// 		// 将当前文件作为子节点添加到目录下
// 		child := FileNode{filename, fpath, []*FileNode{}}
// 		node.FileNodes = append(node.FileNodes, &child)
// 		// 如果遍历的当前文件是个目录，则进入该目录进行递归
// 		if fio.IsDir() {
// 			walk(fpath, fio, &child)
// 		}
// 	}
// 	return
// }

// c.TplName = "Proj_category.tpl"
// func (c *ProjController) Timeline() {
// 	imagelist1 := []string{"/static/img/1.jpg", "/static/img/2.jpg", "/static/img/3.jpg"}
// 	imagelist2 := []string{"/static/img/4.jpg", "/static/img/5.jpg", "/static/img/6.jpg"}
// 	imagelist3 := []string{"/static/img/7.jpg", "/static/img/8.jpg", "/static/img/9.jpg"}
// 	imagelist4 := []string{"/static/img/10.jpg", "/static/img/11.jpg", "/static/img/12.jpg"}
// 	imagelist5 := []string{"/static/img/13.jpg", "/static/img/14.jpg", "/static/img/15.jpg"}
// 	imagelist6 := []string{"/static/img/16.jpg", "/static/img/17.jpg", "/static/img/18.jpg"}

// 	listimage1 := Listimage{
// 		1,
// 		"uer0001",
// 		"2017/03/18",
// 		"秦晓川",
// 		"通过图像识别获得眼像特征",
// 		"知识库自动获取的饼子",
// 		"根据病症信息分析结果",
// 		"\n\t对综合揭露进行\n\t\t\t 行详细描述",
// 		imagelist1,
// 		"2017-03-18",
// 		"",
// 	}
// 	listimage2 := Listimage{
// 		2,
// 		"uer0002",
// 		"2017/03/14",
// 		"秦晓川2",
// 		"识别技术更新",
// 		"来自库",
// 		"分析结果",
// 		"\n\t对综合\n\t\t\t 详细描述",
// 		imagelist2,
// 		"2017-03-13",
// 		"",
// 	}
// 	listimage3 := Listimage{
// 		3,
// 		"uer0003",
// 		"2017/03/10",
// 		"秦晓川3",
// 		"特征",
// 		"自动获取",
// 		"根据结果",
// 		"\n\t进行\n\t\t\t 详细描述",
// 		imagelist3,
// 		"2017-03-10",
// 		"",
// 	}
// 	listimage4 := Listimage{
// 		4,
// 		"uer0004",
// 		"2017/03/02",
// 		"秦晓川4",
// 		"通过特征",
// 		"知识库",
// 		"分析结果",
// 		"\n\t综合揭露\n\t\t\t 描述",
// 		imagelist4,
// 		"2014-07-13",
// 		"",
// 	}
// 	listimage5 := Listimage{
// 		5,
// 		"uer0005",
// 		"2016/07/14",
// 		"秦晓川5",
// 		"通过图像识别获得眼像特征",
// 		"知识库自动获取的饼子",
// 		"根据病症信息分析结果",
// 		"\n\t对综合揭露进行\n\t\t\t 行详细描述",
// 		imagelist5,
// 		"2014-07-13",
// 		"",
// 	}
// 	listimage6 := Listimage{
// 		6,
// 		"uer0006",
// 		"2015/07/14",
// 		"秦晓川6",
// 		"眼像特征",
// 		"获取",
// 		"信息结果",
// 		"\n\t揭露进行\n\t\t\t 详细描述",
// 		imagelist6,
// 		"2014-07-13",
// 		"",
// 	}
// 	listimage := []Listimage{listimage1, listimage2, listimage3, listimage4, listimage5, listimage6}
// 	c.Data["json"] = listimage
// 	// c.Data["json"] = catalogs
// 	c.ServeJSON()
// }

//取得某个侧栏id下的成果给table
// func (c *ProjController) GetProducts() {
// 	id := c.Ctx.Input.Param(":id")
// 	beego.Info(id)
// 	c.Data["Id"] = id
// 	var idNum int64
// 	var err error
// 	if id != "" {
// 		//id转成64为
// 		idNum, err = strconv.ParseInt(id, 10, 64)
// 		if err != nil {
// 			beego.Error(err)
// 		}

// 	} else {

// 	}
// 	//根据id取得所有成果
// 	products, err := models.GetProducts(idNum)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	c.Data["json"] = products
// 	c.ServeJSON()
// 	// c.Data["json"] = root
// 	// c.ServeJSON()
// }

// //向某个侧栏id下添加成果
// func (c *ProjController) AddProduct() {
// 	id := c.Ctx.Input.Param(":id")
// 	pid := c.Input().Get("pid")
// 	code := c.Input().Get("code")
// 	title := c.Input().Get("title")
// 	label := c.Input().Get("label")
// 	principal := c.Input().Get("principal")
// 	content := c.Input().Get("content")
// 	// beego.Info(id)
// 	c.Data["Id"] = id
// 	//id转成64为
// 	idNum, err := strconv.ParseInt(pid, 10, 64)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	//根据id添加成果code, title, label, principal, content string, projectid int64
// 	_, err = models.AddProduct(code, title, label, principal, content, idNum)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	c.Data["json"] = "ok"
// 	c.ServeJSON()
// 	// c.Data["json"] = root
// 	// c.ServeJSON()
// }
