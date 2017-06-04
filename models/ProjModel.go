package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	// "strings"
	"fmt"
	"time"
)

type Project struct {
	Id              int64     `form:"-"`
	Code            string    `orm:"null"`                                              //编号
	Title           string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Label           string    `orm:"null"`                                              //标签
	Principal       string    `orm:"null"`                                              //负责人id
	ParentId        int64     `orm:"null"`
	ParentIdPath    string    `orm:"null"`
	ParentTitlePath string    `orm:"null"`
	Grade           int       `orm:"null"`
	Created         time.Time `orm:"null","auto_now_add;type(datetime)"`
	Updated         time.Time `orm:"null","auto_now_add;type(datetime)"`
}

type Pidstruct struct {
	ParentId        int64
	ParentTitle     string
	ParentIdPath    string
	ParentTitlePath string
}

type ProjCalendar struct {
	Id          int64     `json:"id",form:"-"`
	ProjectId   int64     `json:"projectid"`
	Title       string    `json:"title",form:"title;text;title:",valid:"MinSize(1);MaxSize(100)"` //orm:"unique",
	Content     string    `json:"content",orm:"sie(20)"`
	Starttime   time.Time `json:"start",orm:"type(datetime)"`
	Endtime     time.Time `json:"end",orm:"null;type(datetime)"`
	Allday      bool      `json:"allDay",orm:"default(true)"`
	Memorabilia bool      `json:"memorabilia",orm:"default(false)"` //是否属于大事记
	Image       string    `json:"image",orm:"null"`                 //图片链接地址
	Color       string    `json:"color",orm:"null"`
	Public      bool      `default(true)` //是否公开
	// Color     string    `json:"backgroundColor",orm:"null"`
	// BColor    string    `json:"borderColor",orm:"null"`
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
	orm.RegisterModel(new(Project), new(ProjCalendar)) //, new(Article)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/engineer.db", 10)
}

//添加项目
func AddProject(code, title, label, principal string, parentid int64, parentidpath, parenttitlepath string, grade int) (id int64, err error) {
	o := orm.NewOrm()
	//关闭写同步
	// o.Raw("PRAGMA synchronous = OFF; ", 0, 0, 0).Exec()
	// var project Project
	// if pid == "" {
	project := &Project{
		Code:            code,
		Title:           title,
		Label:           label,
		Principal:       principal,
		ParentId:        parentid,
		ParentIdPath:    parentidpath,
		ParentTitlePath: parenttitlepath,
		Grade:           grade,
		Created:         time.Now(),
		Updated:         time.Now(),
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
func UpdateProject(cid int64, code, title, label, principal string) error {
	o := orm.NewOrm()
	project := &Project{Id: cid}
	if o.Read(project) == nil {
		project.Code = code
		project.Title = title
		project.Label = label
		project.Principal = principal
		project.Updated = time.Now()
		_, err := o.Update(project)
		if err != nil {
			return err
		}
	}
	return nil
}

//删除
func DeleteProject(id int64) error {
	o := orm.NewOrm()
	project := &Project{Id: id}
	if o.Read(project) == nil {
		_, err := o.Delete(project)
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
//根据id查出所有子孙，用ParentIdPath
func GetProjectsbyPid(id int64) (projects []*Project, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Project")
	_, err = qs.Filter("ParentIdPath__contains", id).All(&projects)
	if err != nil {
		return nil, err
	}
	return projects, err
}

//根据id查出所有儿子
func GetProjSonbyId(id int64) (projects []*Project, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Project")
	_, err = qs.Filter("parentid", id).All(&projects)
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

//根据目录id取得第一级项目
//用parentidpath的第一个数字就行了。

//根据parenttitlepath和title取得proj目录
func GetProjbyParenttitlepath(parenttitlepath, title string) (proj Project, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Project") //这个表名AchievementTopic需要用驼峰式，
	err = qs.Filter("ParentTitlePath", parenttitlepath).Filter("title", title).One(&proj)
	if err != nil {
		return proj, err
	}
	return proj, err
}

//递归将目录写入数据库
func Insertproj(pid []Pidstruct, nodes []*AdminCategory, igrade, height int) (cid []Pidstruct) {
	o := orm.NewOrm() //实例化数据库操作对象
	// o.Using("default")
	//关闭写同步
	o.Raw("PRAGMA synchronous = OFF; ", 0, 0, 0).Exec()
	// var project models.Project
	var Id int64
	for _, v := range pid {
		for _, v1 := range nodes {
			if v1.Grade == igrade {
				title := v1.Title
				code := v1.Code
				parentid := v.ParentId

				var parentidpath string
				var parenttitlepath string
				if v.ParentIdPath != "" {
					parentidpath = v.ParentIdPath + "-" + strconv.FormatInt(v.ParentId, 10)
					parenttitlepath = v.ParentTitlePath + "-" + v.ParentTitle
				} else {
					parentidpath = strconv.FormatInt(v.ParentId, 10)
					parenttitlepath = v.ParentTitle
				}

				grade := igrade
				//通过事务方式来进行数据插入
				// err := o.Begin()
				const lll = "2006-01-02 15:04:05.000"
				date := time.Now().Format(lll)
				sql := fmt.Sprintf("insert into Project (Code, Title, Label, Principal, Parent_id, Parent_id_path, Parent_title_path, Grade,Created,Updated)"+
					" values('%s','%s','%s','%s',%d,'%s','%s',%d,'%s','%s')", code, title, "", "", parentid, parentidpath, parenttitlepath, grade, date, date)
				res, err := o.Raw(sql).Exec()
				if err != nil {
					o.Rollback()
					// beego.Info("插入t_studentInfo表出错,事务回滚")
				} else {
					// o.Commit()
					// beego.Info("插入t_studenInfo表成功,事务提交")
					// num, _ = res.RowsAffected()
					Id, _ = res.LastInsertId()
				}
				// Id, err := models.AddProject(code, title, "", "", parentid, parentidpath, parenttitlepath, grade)
				// if err != nil {
				// 	beego.Error(err)
				// }
				var cid1 Pidstruct
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
		Insertproj(cid, nodes, igrade, height)
	}
	return
}

//************项目日历
//********日历********
//添加
func AddProjCalendar(pid int64, title, content, color, imgurl string, allday, public, memorabilia bool, start, end time.Time) (id int64, err error) {
	o := orm.NewOrm()
	calendar := &ProjCalendar{
		ProjectId:   pid,
		Title:       title,
		Content:     content,
		Color:       color,
		Allday:      allday,
		Public:      public,
		Memorabilia: memorabilia,
		Image:       imgurl,
		// BColor:    color,
		Starttime: start,
		Endtime:   end,
	}
	id, err = o.Insert(calendar)
	if err != nil {
		return id, err
	}
	// }
	return id, err
}

//取所有——要修改为支持时间段的，比如某个月份
func GetProjCalendar(pid int64, start, end time.Time, public bool) (calendars []*ProjCalendar, err error) {
	cond := orm.NewCondition()
	cond1 := cond.And("Starttime__gte", start).And("Starttime__lt", end) //这里全部用开始时间来判断
	o := orm.NewOrm()
	qs := o.QueryTable("ProjCalendar")
	qs = qs.SetCond(cond1)

	// o := orm.NewOrm()
	calendars = make([]*ProjCalendar, 0)
	// qs := o.QueryTable("ProjCalendar")
	if public { //只取公开的
		_, err = qs.Filter("ProjectId", pid).Filter("public", true).OrderBy("-Starttime").All(&calendars)
		if err != nil {
			return calendars, err
		}
	} else { //取全部
		_, err = qs.Filter("ProjectId", pid).OrderBy("-Starttime").All(&calendars)
		if err != nil {
			return calendars, err
		}
	}
	return calendars, err
}

//取出所有日历日程
func GetAllProjCalendar(pid int64, public bool) (calendars []*ProjCalendar, err error) {
	// cond := orm.NewCondition()
	// cond1 := cond.And("Starttime__gte", start).And("Starttime__lt", end) //这里全部用开始时间来判断
	o := orm.NewOrm()
	qs := o.QueryTable("ProjCalendar")
	// qs = qs.SetCond(cond1)

	// o := orm.NewOrm()
	calendars = make([]*ProjCalendar, 0)
	// qs := o.QueryTable("ProjCalendar")
	if public { //只取公开的
		_, err = qs.Filter("ProjectId", pid).Filter("public", true).Filter("memorabilia", true).OrderBy("-Starttime").All(&calendars)
		if err != nil {
			return calendars, err
		}
	} else { //取全部
		_, err = qs.Filter("ProjectId", pid).Filter("memorabilia", true).OrderBy("-Starttime").All(&calendars)
		if err != nil {
			return calendars, err
		}
	}
	return calendars, err
}

//取出分页的日历
func ListPostsByOffsetAndLimit(pid int64, set, postsPerPage int, public bool) ([]*ProjCalendar, error) {
	o := orm.NewOrm()
	calendars := make([]*ProjCalendar, 0)
	qs := o.QueryTable("ProjCalendar")
	var err error
	if public { //只取公开的
		_, err = qs.Filter("ProjectId", pid).Filter("public", true).Filter("memorabilia", true).Limit(postsPerPage, set).OrderBy("-Starttime").All(&calendars)
		return calendars, err
	} else { //取全部
		_, err = qs.Filter("ProjectId", pid).Filter("memorabilia", true).Limit(postsPerPage, set).OrderBy("-Starttime").All(&calendars)
		return calendars, err
	}
}

//修改
func UpdateProjCalendar(cid int64, title, content, color, url string, allday, public, memorabilia bool, start, end time.Time) error {
	o := orm.NewOrm()
	calendar := &ProjCalendar{Id: cid}
	if o.Read(calendar) == nil {
		calendar.Title = title
		calendar.Content = content
		calendar.Color = color
		calendar.Image = url
		calendar.Allday = allday
		calendar.Public = public
		calendar.Memorabilia = memorabilia
		// calendar.BColor = color
		calendar.Starttime = start
		calendar.Endtime = end
		// calendar.Updated = time.Now()
		_, err := o.Update(calendar)
		if err != nil {
			return err
		}
	}
	return nil
}

//拖曳
func DropProjCalendar(cid int64, start, end time.Time) error {
	o := orm.NewOrm()
	calendar := &ProjCalendar{Id: cid}
	if o.Read(calendar) == nil {
		calendar.Starttime = start
		calendar.Endtime = end
		_, err := o.Update(calendar)
		if err != nil {
			return err
		}
	}
	return nil
}

//resize
func ResizeProjCalendar(cid int64, end time.Time) error {
	o := orm.NewOrm()
	calendar := &ProjCalendar{Id: cid}
	if o.Read(calendar) == nil {
		// calendar.Starttime = start
		calendar.Endtime = end
		// calendar.Updated = time.Now()
		_, err := o.Update(calendar)
		if err != nil {
			return err
		}
	}
	return nil
}

//根据id查询事件
func GetProjCalendarbyid(id int64) (calendar ProjCalendar, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("ProjCalendar")
	err = qs.Filter("id", id).One(&calendar)
	if err != nil {
		return calendar, err
	}
	return calendar, err
}

//删除事件
func DeleteProjCalendar(cid int64) error {
	o := orm.NewOrm()
	calendar := &ProjCalendar{Id: cid}
	if o.Read(calendar) == nil {
		_, err := o.Delete(calendar)
		if err != nil {
			return err
		}
	}
	return nil
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
