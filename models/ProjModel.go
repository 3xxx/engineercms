package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	// "strconv"
	// "strings"
	"time"
)

type Project struct {
	Id           int64     `form:"-"`
	Code         string    `orm:"null"`                                              //编号
	Title        string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Label        string    `orm:"null"`                                              //标签
	Principal    string    `orm:"null"`                                              //负责人id
	ParentId     int64     `orm:"null"`
	ParentIdPath string    `orm:"null"`
	Grade        int       `orm:"null"`
	Created      time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated      time.Time `orm:"index","auto_now_add;type(datetime)"`
}

// type Product struct {
// 	Id        int64
// 	Code      string    `orm:"null"`                                              //编号                                             //编号
// 	Title     string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
// 	Label     string    `orm:"null"`                                              //关键字                                           //标签
// 	Uid       int64     `orm:"null"`
// 	Principal string    `orm:"null"`      //提供人                                            //负责人id
// 	ProjectId int64     `orm:"null"`      //侧栏id
// 	Content   string    `orm:"sie(5000)"` //内容
// 	Created   time.Time `orm:"index","auto_now_add;type(datetime)"`
// 	Updated   time.Time `orm:"index","auto_now;type(datetime)"`
// 	Views     int64
// 	// ReplyTime         time.Time
// 	// ReplyCount        int64
// 	// ReplyLastUserName string
// 	// Attachments     []*Attachment `orm:"reverse(many)"` // fk 的反向关系
// }

// //附件,attachment 和 topic 是 ManyToOne 关系，也就是 ForeignKey 为 topic
// type Attachment struct {
// 	Id        int64
// 	FileName  string
// 	FileSize  string
// 	Downloads int64
// 	ProductId int64     //*Topic    `orm:"rel(fk)"`
// 	Created   time.Time `orm:"index","auto_now_add;type(datetime)"`
// 	Updated   time.Time `orm:"index","auto_now;type(datetime)"`
// }

func init() {
	orm.RegisterModel(new(Project)) //, new(Article)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/engineer.db", 10)
}

//添加项目
func AddProject(code, title, label, principal string, parentid int64, parentidpath string, grade int) (id int64, err error) {
	o := orm.NewOrm()
	// var project Project
	// if pid == "" {
	project := &Project{
		Code:         code,
		Title:        title,
		Label:        label,
		Principal:    principal,
		ParentId:     parentid,
		ParentIdPath: parentidpath,
		Grade:        grade,
		Created:      time.Now(),
		Updated:      time.Now(),
	}
	id, err = o.Insert(project)
	if err != nil {
		return 0, err
	}
	// } else {

	// }
	return id, nil
}

//修改——还没改
func UpdateProject(cid int64, title, code string, grade int) error {
	o := orm.NewOrm()
	//id转成64为
	// cidNum, err := strconv.ParseInt(cid, 10, 64)
	// if err != nil {
	// 	return err
	// }
	project := &Project{Id: cid}
	if o.Read(project) == nil {
		project.Title = title
		project.Code = code
		project.Grade = grade
		project.Updated = time.Now()
		_, err := o.Update(project)
		if err != nil {
			return err
		}
	}
	return nil
}

//取得所有项目
func GetProjects() (proj []*Project, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Project") //这个表名AchievementTopic需要用驼峰式，
	_, err = qs.Filter("parentid", 0).All(&proj)
	if err != nil {
		return proj, err
	}
	return proj, err
}

//根据id取得项目目录
func GetProj(id int64) (proj Project, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Project") //这个表名AchievementTopic需要用驼峰式，
	err = qs.Filter("id", id).One(&proj)
	if err != nil {
		return proj, err
	}
	return proj, err
}

//根据id取得项目文件夹_未使用
// func GetProjPath(id int64) (path string, err error) {
// 	o := orm.NewOrm()
// 	qs := o.QueryTable("Project") //这个表名AchievementTopic需要用驼峰式，
// 	err = qs.Filter("id", id).One(&proj)
// 	if err != nil {
// 		return proj, err
// 	}
// 	return proj, err
// }
//根据id查出所有子孙——这个不对，要用ParentIdPath
func GetProjectsbyPid(id int64) (projects []*Project, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Project")                         //这个表名AchievementTopic需要用驼峰式，
	_, err = qs.Filter("ParentIdPath", id).All(&projects) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	return projects, err
}

//根据id查出所有儿子
func GetProjSonbyId(id int64) (projects []*Project, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Project")                     //这个表名AchievementTopic需要用驼峰式，
	_, err = qs.Filter("parentid", id).All(&projects) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	return projects, err
}

//根据id查是否有下级
func Projhasson(id int64) bool {
	o := orm.NewOrm()
	// qs := o.QueryTable("Project")
	proj := Project{ParentId: id}
	err := o.Read(&proj, "ParentId")
	if err == orm.ErrNoRows {
		return false
	} else if err == orm.ErrMissPK {
		return false
	} else {
		return true
	}
}

//根据名字title查询到项目目录
func GetProjectTitle(title string) (cate Project, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Project")
	// var cate Project
	err = qs.Filter("title", title).One(&cate)
	// if pid != "" {
	// cate := Project{Title: title}这句无效
	// categories = make([]*Project, 0)
	// _, err = qs.Filter("parentid", cate.Id).All(&categories)
	// if err != nil {
	// 	return nil, err
	// }
	return cate, err
	// } else { //如果不给定父id（PID=0），则取所有一级
	// _, err = qs.Filter("parentid", 0).All(&categories)
	// if err != nil {
	// return nil, err
	// }
	// return categories, err
	// }
}

//添加成果到项目侧栏某个id下
// func AddProduct(code, title, label, principal, content string, projectid int64) (id int64, err error) {
// 	o := orm.NewOrm()
// 	// var project Project
// 	// if pid == "" {
// 	product := &Product{
// 		Code:      code,
// 		Title:     title,
// 		Label:     label,
// 		Principal: principal,
// 		ProjectId: projectid,
// 		Content:   content,
// 		Created:   time.Now(),
// 		Updated:   time.Now(),
// 	}
// 	id, err = o.Insert(product)
// 	if err != nil {
// 		return 0, err
// 	}
// 	// } else {

// 	// }
// 	return id, nil
// }

//根据侧栏id查出所有成果
// func GetProducts(id int64) (products []*Product, err error) {
// 	o := orm.NewOrm()
// 	qs := o.QueryTable("Product")                      //这个表名AchievementTopic需要用驼峰式，
// 	_, err = qs.Filter("projectid", id).All(&products) //而这个字段parentid为何又不用呢
// 	if err != nil {
// 		return nil, err
// 	}
// 	return products, err
// }
