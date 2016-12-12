package controllers

import (
	// "encoding/json"
	"engineercms/models"
	"github.com/astaxie/beego"
	"os"
	"strconv"
	"strings"
)

type ProjController struct {
	beego.Controller
}

type Pidstruct struct {
	ParentId     int64
	ParentIdPath string
}

//项目列表页面
func (c *ProjController) Get() {
	c.Data["IsCategory"] = true
	c.TplName = "projects.tpl"
	//取得项目类别，给添加项目模态框选项用
	var slice1 []string
	categories, err := models.GetAdminCategory(0)
	if err != nil {
		beego.Error(err)
	}
	for _, v := range categories {
		aa := make([]string, 1)
		aa[0] = v.Title //名称
		// cc[0].Selectable = false
		slice1 = append(slice1, aa...)
	}
	c.Data["Select2"] = slice1
}

//提供给项目列表页的table中json数据，扩展后按标签显示
func (c *ProjController) GetProjects() {
	id := c.Ctx.Input.Param(":id")
	if id == "" {
		//显示全部
		projects, err := models.GetProjects()
		if err != nil {
			beego.Error(err)
		}
		c.Data["json"] = projects
		c.ServeJSON()
	} else {
		//根据标签查询
	}
}

//根据ip查看项目，查出项目目录
func (c *ProjController) GetProject() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	// var categories []*models.ProjCategory
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
	// categories, err := models.GetProjectsbyPid(idNum)
	// if err != nil {
	// 	beego.Error(err)
	// }
	//算出最大级数
	// grade := make([]int, 0)
	// for _, v := range categories {
	// 	grade = append(grade, v.Grade)
	// }
	// height := intmax(grade[0], grade[1:]...)
	//递归生成目录json
	root := FileNode{category.Id, category.Title, []*FileNode{}}
	walk(category.Id, &root)
	// beego.Info(root)
	// data, _ := json.Marshal(root)
	c.Data["json"] = root //data
	// c.ServeJSON()
	c.Data["Category"] = category
	c.TplName = "project.tpl"
}

//根据项目侧栏id查看这个id下的成果，不含子目录中的成果
//任何一级目录下都可以放成果
func (c *ProjController) GetProjProd() {
	id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	c.Data["Id"] = id
	// var categories []*models.ProjCategory
	// var err error
	//id转成64为
	// idNum, err := strconv.ParseInt(id, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	//取项目本身
	// category, err := models.GetProj(idNum)
	// if err != nil {
	// 	beego.Error(err)
	// }
	//取项目所有子孙
	// categories, err := models.GetProjectsbyPid(idNum)
	// if err != nil {
	// 	beego.Error(err)
	// }
	//算出最大级数
	// grade := make([]int, 0)
	// for _, v := range categories {
	// 	grade = append(grade, v.Grade)
	// }
	// height := intmax(grade[0], grade[1:]...)

	// c.Data["json"] = root
	// c.ServeJSON()
	c.TplName = "project_products.tpl"
}

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

//添加项目和项目目录、文件夹
func (c *ProjController) AddProject() {
	// rows := c.Input().Get("rows2[0][0]")
	// beego.Info(rows)
	projcode := c.Input().Get("code")
	projname := c.Input().Get("name")
	projlabe := c.Input().Get("label")
	principal := c.Input().Get("principal")
	//先保存项目名称到数据库，parentid为0，返回id作为下面二级三级四级……的parentid
	//然后递归保存二级三级……到数据库
	//最后递归生成硬盘目录
	Id, err := models.AddProject(projcode, projname, projlabe, principal, 0, "", 1)
	if err != nil {
		beego.Error(err)
	}
	//根据id查出分级目录的名称、代码和层数
	//如果不建立下级怎样？页面中允许不选择任何
	ids := c.GetString("ids")
	array := strings.Split(ids, ",")
	// nodes := make([]models.AdminCategory, 0)

	// for i1, v1 := range array {
	// 	category, err := models.GetAdminCategorybyId(v1)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// 	aa := make([]models.AdminCategory, 1)
	// 	aa[0].Title = category.Title
	// 	aa[0].Code = category.Code
	// 	aa[0].Grade = category.Grade
	// 	nodes = append(nodes, aa...)
	// }
	nodes := make([]*models.AdminCategory, 0) //这样完美解决像上面那样借助aa[0]那样
	grade := make([]int, 0)
	for _, v2 := range array {
		// pid = strconv.FormatInt(v1, 10)
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
	idarr := make([]Pidstruct, 1)
	idarr[0].ParentId = Id
	idarr[0].ParentIdPath = "" //strconv.FormatInt(Id, 10)
	write(idarr, nodes, 2, height)
	//递归创建文件夹
	patharr := make([]Pathstruct, 1)
	patharr[0].ParentPath = ".\\attachment\\" + projcode + projname
	create(patharr, nodes, 2, height)
	c.Data["json"] = "ok"
	c.ServeJSON()
}

//还没改，应该是updateproj
func (c *ProjController) UpdateCategory() {
	// pid := c.Ctx.Input.Param(":id")
	cid := c.Input().Get("cid")
	title := c.Input().Get("title")
	code := c.Input().Get("code")
	grade := c.Input().Get("grade")
	//cid转成64为
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
func write(pid []Pidstruct, nodes []*models.AdminCategory, igrade, height int) (cid []Pidstruct) {
	for _, v := range pid {
		for _, v1 := range nodes {
			if v1.Grade == igrade {
				title := v1.Title
				parentid := v.ParentId
				var parentidpath string
				if v.ParentIdPath != "" {
					parentidpath = v.ParentIdPath + "-" + strconv.FormatInt(v.ParentId, 10)
				} else {
					parentidpath = strconv.FormatInt(v.ParentId, 10)
				}
				grade := igrade
				Id, err := models.AddProject("", title, "", "", parentid, parentidpath, grade)
				if err != nil {
					beego.Error(err)
				}
				var cid1 Pidstruct
				cid1.ParentId = Id
				cid1.ParentIdPath = parentidpath
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

type FileNode struct {
	Id        int64       `json:"id"`
	Title     string      `json:"text"`
	FileNodes []*FileNode `json:"nodes"`
}

//递归构造项目树状目录
func walk(id int64, node *FileNode) {
	//列出当前id下子节点，不要列出孙节点……
	files, err := models.GetProjSonbyId(id)
	if err != nil {
		beego.Error(err)
	}
	// 遍历目录
	for _, proj := range files {
		id := proj.Id
		title := proj.Title
		// 将当前名和id作为子节点添加到目录下
		child := FileNode{id, title, []*FileNode{}}
		node.FileNodes = append(node.FileNodes, &child)
		// 如果遍历的当前节点下还有节点，则进入该节点进行递归
		if models.Projhasson(proj.Id) {
			walk(proj.Id, &child)
		}
	}
	return
}

type Pathstruct struct {
	ParentPath string
}

//递归建立文件夹
func create(path []Pathstruct, nodes []*models.AdminCategory, igrade, height int) (cpath []Pathstruct) {
	for _, v := range path {
		for _, v1 := range nodes {
			if v1.Grade == igrade {
				title := v1.Title
				parentpath := v.ParentPath
				//建立目录，并返回作为父级目录
				err := os.MkdirAll(parentpath+"\\"+title, 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
				if err != nil {
					beego.Error(err)
				}

				var cpath1 Pathstruct
				cpath1.ParentPath = parentpath + "\\" + title
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
