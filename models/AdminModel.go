package models

import (
	"database/sql"
	"encoding/gob"
	// "github.com/3xxx/engineercms/commands"
	"github.com/3xxx/engineercms/conf"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	// "gorm.io/driver/sqlite"
	// "gorm.io/gorm"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
	// "strconv"
	// "strings"
	"fmt"
	"log"
	"os"
	"time"
)

var engine *xorm.Engine

//定义全局的db对象，我们执行数据库操作主要通过他实现。
var _db *gorm.DB

// var gdb *gorm.DB

type AdminCategory struct {
	Id       int64     `form:"-"`
	ParentId int64     `orm:"null"`
	Title    string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Code     string    `orm:"null"`
	Grade    int       `orm:"null"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}

type AdminIpsegment struct {
	Id      int64     `form:"-"`
	Title   string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	StartIp string    //`orm:"not null"`
	EndIp   string    `orm:"null"`
	Iprole  int       `orm:"null"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

type AdminCalendar struct {
	Id        int64     `json:"id",form:"-"`
	Title     string    `json:"title",form:"title;text;title:",valid:"MinSize(1);MaxSize(100)"` //orm:"unique",
	Content   string    `json:"content",orm:"sie(20)"`
	Starttime time.Time `json:"start",orm:"type(datetime)"`
	Endtime   time.Time `json:"end",orm:"null;type(datetime)"`
	Allday    bool      `json:"allDay",orm:"default(0)"`
	Color     string    `json:"color",orm:"null"`
	Public    bool      `default(true)`
	// Color     string    `json:"backgroundColor",orm:"null"`
	// BColor    string    `json:"borderColor",orm:"null"`
}

//项目同步ip表
type AdminSynchIp struct {
	Id       int64     `form:"-"`
	ParentId int64     `orm:"null"`
	UserName string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	SynchIp  string    //`orm:"not null"`
	Port     string    `orm:"default(80)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}

//科室结构
type AdminDepartment struct {
	Id       int64     `form:"-"`
	ParentId int64     `orm:"null"`
	Title    string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Code     string    `orm:"null"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now_add;type(datetime)"`
}

//首页轮播图片
type AdminCarousel struct {
	Id      int64     `json:"id",form:"-"`
	Title   string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Url     string    `orm:"null"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now_add;type(datetime)"`
}

// `id` int(11) NOT NULL AUTO_INCREMENT,
//   `title` varchar(100) NOT NULL,
//   `starttime` int(11) NOT NULL,
//   `endtime` int(11) DEFAULT NULL,
//   `allday` tinyint(1) NOT NULL DEFAULT '0',
//   `color` varchar(20) DEFAULT NULL,

func init() {
	orm.RegisterModel(new(AdminCategory), new(AdminIpsegment), new(AdminCalendar), new(AdminSynchIp), new(AdminDepartment), new(AdminCarousel)) //, new(Article)
	orm.RegisterModelWithPrefix(conf.GetDatabasePrefix(),
		new(Member),
		new(Book),
		new(Relationship),
		new(Option),
		new(Document),
		new(MindocAttachment),
		new(Logger),
		new(MemberToken),
		new(DocumentHistory),
		new(Migration),
		new(Label),
		new(Blog),
		new(Template),
		new(Team),
		new(TeamMember),
		new(TeamRelationship),
		new(Itemsets),
	)
	orm.RegisterModelWithPrefix("share_", new(Bridge), new(Share))
	//gorm设置默认表名前缀
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return "prefix_" + defaultTableName
	// }
	// //gorm自动生成表
	// db.AutoMigrate(&Product{}, &Email{})
	// db.CreateTable(&User{})
	gob.Register(Blog{})
	gob.Register(Document{})
	gob.Register(Template{})

	var dns string
	var err error
	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	db_path := beego.AppConfig.String("db_path")
	db_sslmode := beego.AppConfig.String("db_sslmode")
	switch db_type {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DRMySQL)
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db_user, db_pass, db_host, db_port, db_name)
		// 注册xorm
		// var err error
		// engine, err = xorm.NewEngine(db_type, dns)
		// if err != nil {
		// 	log.Println(err)
		// }
		break
	case "postgres":
		orm.RegisterDriver("postgres", orm.DRPostgres)
		dns = fmt.Sprintf("dbname=%s host=%s  user=%s  password=%s  port=%s  sslmode=%s", db_name, db_host, db_user, db_pass, db_port, db_sslmode)
	case "sqlite3":
		orm.RegisterDriver("sqlite", orm.DRSqlite)
		if db_path == "" {
			db_path = "./"
		}
		dns = fmt.Sprintf("%s%s.db", db_path, db_name)
		break
	default:
		beego.Critical("Database driver is not allowed:", db_type)
	}
	orm.RegisterDataBase("default", db_type, dns, 10)

	// 注册xorm
	// var err error
	engine, err = xorm.NewEngine(db_type, dns)
	if err != nil {
		log.Println(err)
	}

	// 注册gorm
	_db, err = gorm.Open(db_type, dns)
	// _db, err := gorm.Open(sqlite.Open(dns), &gorm.Config{})
	// defer _db.Close()//20200803这个不能打开。
	// _db.LogMode(true)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/engineer.db", 10)

	// 初始化gorm-20201008
	// 	var err error
	// var dns string
	// db_type := beego.AppConfig.String("db_type")
	// db_name := beego.AppConfig.String("db_name")
	// db_path := beego.AppConfig.String("db_path")
	// if db_path == "" {
	// 	db_path = "./"
	// }

	// defer gdb.Close()
	//禁止表名复数形式
	_db.SingularTable(true)
	// 开发的时候需要打开调试日志
	// _db.LogMode(true)
	//设置数据库连接池参数
	_db.DB().SetMaxOpenConns(100) //设置数据库连接池最大连接数
	_db.DB().SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

//获取gorm db对象，其他包需要执行数据库查询的时候，只要通过tools.getDB()获取db对象即可。
//不用担心协程并发使用同样的db对象会共用同一个连接，
// db对象在调用他的方法的时候会从数据库连接池中获取新的连接
// 注意：使用连接池技术后，千万不要使用完db后调用db.Close关闭数据库连接，
// 这样会导致整个数据库连接池关闭，导致连接池没有可用的连接
func GetDB() *gorm.DB {
	return _db
}

//创建数据库_来自github.com/beego/admin——这个仅作为参考用，没有使用
func createdb() {
	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	db_path := beego.AppConfig.String("db_path")
	db_sslmode := beego.AppConfig.String("db_sslmode")

	var dns string
	var sqlstring string
	switch db_type {
	case "mysql":
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", db_user, db_pass, db_host, db_port)
		sqlstring = fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", db_name)
		break
	case "postgres":
		dns = fmt.Sprintf("host=%s  user=%s  password=%s  port=%s  sslmode=%s", db_host, db_user, db_pass, db_port, db_sslmode)
		sqlstring = fmt.Sprintf("CREATE DATABASE %s", db_name)
		break
	case "sqlite3":
		if db_path == "" {
			db_path = "./"
		}
		dns = fmt.Sprintf("%s%s.db", db_path, db_name)
		os.Remove(dns)
		sqlstring = "create table init (n varchar(32));drop table init;"
		break
	default:
		beego.Critical("Database driver is not allowed:", db_type)
	}

	// gdb, err := gorm.Open(db_type, dns)
	// gdb.LogMode(true)
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// defer gdb.Close()

	db, err := sql.Open(db_type, dns)
	if err != nil {
		panic(err.Error())
	}
	r, err := db.Exec(sqlstring)
	if err != nil {
		log.Println(err)
		log.Println(r)
	} else {
		log.Println("Database ", db_name, " created")
	}
	defer db.Close()
}

//数据库连接_这个仅作为参考，和上面重复
func Connect() {
	var dns string
	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	db_path := beego.AppConfig.String("db_path")
	db_sslmode := beego.AppConfig.String("db_sslmode")
	switch db_type {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DRMySQL)
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_pass, db_host, db_port, db_name, 30)
		// set default database
		// orm.RegisterDataBase("default", "mysql", "username:password@tcp(127.0.0.1:3306)/db_name?charset=utf8", 30)
		break
	case "postgres":
		orm.RegisterDriver("postgres", orm.DRPostgres)
		dns = fmt.Sprintf("dbname=%s host=%s  user=%s  password=%s  port=%s  sslmode=%s", db_name, db_host, db_user, db_pass, db_port, db_sslmode)
	case "sqlite3":
		orm.RegisterDriver("sqlite3", orm.DRSqlite)
		if db_path == "" {
			db_path = "./"
		}
		dns = fmt.Sprintf("%s%s.db", db_path, db_name)
		break
	default:
		beego.Critical("Database driver is not allowed:", db_type)
	}
	orm.RegisterDataBase("default", db_type, dns)
}

//添加
func AddAdminCategory(pid int64, title, code string, grade int) (id int64, err error) {
	o := orm.NewOrm()
	// var category AdminCategory
	// if pid == "" {
	// 	category := &AdminCategory{
	// 		ParentId: 0,
	// 		Title:    title,
	// 		Code:     code,
	// 		Grade:    grade,
	// 		Created:  time.Now(),
	// 		Updated:  time.Now(),
	// 	}
	// 	id, err = o.Insert(category)
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// } else {
	//pid转成64为
	// pidNum, err := strconv.ParseInt(pid, 10, 64)
	// if err != nil {
	// 	return 0, err
	// }
	category := &AdminCategory{
		ParentId: pid,
		Title:    title,
		Code:     code,
		Grade:    grade,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
	id, err = o.Insert(category)
	if err != nil {
		return 0, err
	}
	// }
	return id, nil
}

//修改
func UpdateAdminCategory(cid int64, title, code string, grade int) error {
	o := orm.NewOrm()
	//id转成64为
	// cidNum, err := strconv.ParseInt(cid, 10, 64)
	// if err != nil {
	// 	return err
	// }
	category := &AdminCategory{Id: cid}
	if o.Read(category) == nil {
		category.Title = title
		category.Code = code
		category.Grade = grade
		category.Updated = time.Now()
		_, err := o.Update(category)
		if err != nil {
			return err
		}
	}
	return nil
}

//删除
func DeleteAdminCategory(cid int64) error {
	o := orm.NewOrm()
	category := &AdminCategory{Id: cid}
	if o.Read(category) == nil {
		_, err := o.Delete(category)
		if err != nil {
			return err
		}
	}
	// catalog := Catalog{Id: cidNum}
	// if o.Read(&catalog) == nil {
	// 	_, err = o.Delete(&catalog)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

//根据父级id取得所有
//如果父级id为空，则取所有一级category
func GetAdminCategory(pid int64) (categories []*AdminCategory, err error) {
	o := orm.NewOrm()
	categories = make([]*AdminCategory, 0)
	qs := o.QueryTable("AdminCategory")                   //这个表名AchievementTopic需要用驼峰式，
	_, err = qs.Filter("parent_id", pid).All(&categories) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	return categories, err
}

//根据类别名字title查询所有下级分级category
func GetAdminCategoryTitle(title string) (categories []*AdminCategory, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminCategory")
	var cate AdminCategory
	err = qs.Filter("title", title).One(&cate)
	// if pid != "" {
	// cate := AdminCategory{Title: title}这句无效
	categories = make([]*AdminCategory, 0)
	_, err = qs.Filter("parentid", cate.Id).All(&categories)
	if err != nil {
		return nil, err
	}
	return categories, err
	// } else { //如果不给定父id（PID=0），则取所有一级
	// _, err = qs.Filter("parentid", 0).All(&categories)
	// if err != nil {
	// return nil, err
	// }
	// return categories, err
	// }
}

//根据id查分级
func GetAdminCategorybyId(id int64) (category []*AdminCategory, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminCategory")

	err = qs.Filter("id", id).One(&category)
	if err != nil {
		return nil, err
	}
	return category, err
}

//添加ip地址段
func AddAdminIpsegment(title, startip, endip string, iprole int) (id int64, err error) {
	o := orm.NewOrm()
	ipsegment := &AdminIpsegment{
		Title:   title,
		StartIp: startip,
		EndIp:   endip,
		Iprole:  iprole,
		Created: time.Now(),
		Updated: time.Now(),
	}
	id, err = o.Insert(ipsegment)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//修改Ip地址段
func UpdateAdminIpsegment(cid int64, title, startip, endip string, iprole int) error {
	o := orm.NewOrm()
	ipsegment := &AdminIpsegment{Id: cid}
	if o.Read(ipsegment) == nil {
		ipsegment.Title = title
		ipsegment.StartIp = startip
		ipsegment.EndIp = endip
		ipsegment.Iprole = iprole
		ipsegment.Updated = time.Now()
		_, err := o.Update(ipsegment)
		if err != nil {
			return err
		}
	}
	return nil
}

//删除
func DeleteAdminIpsegment(cid int64) error {
	o := orm.NewOrm()
	ipsegment := &AdminIpsegment{Id: cid}
	if o.Read(ipsegment) == nil {
		_, err := o.Delete(ipsegment)
		if err != nil {
			return err
		}
	}
	return nil
}

//查询所有Ip地址段
func GetAdminIpsegment() (ipsegments []*AdminIpsegment, err error) {
	o := orm.NewOrm()
	// ipsegments = make([]*AdminIpsegment, 0)

	qs := o.QueryTable("AdminIpsegment") //这个表名AchievementTopic需要用驼峰式，
	// if pid != "" {                      //如果给定父id则进行过滤
	//pid转成64为
	// pidNum, err := strconv.ParseInt(pid, 10, 64)
	// if err != nil {
	// 	return nil, err
	// }
	_, err = qs.All(&ipsegments)
	if err != nil {
		return nil, err
	}

	return ipsegments, err
	// } else { //如果不给定父id（PID=0），则取所有一级
	// _, err = qs.Filter("parentid", 0).All(&categories)
	// if err != nil {
	// 	return nil, err
	// }
	// return categories, err
	// }
}

//********日历********
//添加
func AddAdminCalendar(title, content, color string, allday, public bool, start, end time.Time) (id int64, err error) {
	o := orm.NewOrm()
	calendar := &AdminCalendar{
		Title:   title,
		Content: content,
		Color:   color,
		Allday:  allday,
		Public:  public,
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
func GetAdminCalendar(start, end time.Time, public bool) (calendars []*AdminCalendar, err error) {
	cond := orm.NewCondition()
	cond1 := cond.And("Starttime__gte", start).And("Starttime__lt", end) //这里全部用开始时间来判断

	o := orm.NewOrm()
	qs := o.QueryTable("AdminCalendar")
	qs = qs.SetCond(cond1)

	// o := orm.NewOrm()
	calendars = make([]*AdminCalendar, 0)
	// qs := o.QueryTable("AdminCalendar")
	if public { //只取公开的
		_, err = qs.Filter("public", true).OrderBy("-Starttime").All(&calendars)
		if err != nil {
			return calendars, err
		}
	} else { //取全部
		_, err = qs.OrderBy("-Starttime").All(&calendars)
		if err != nil {
			return calendars, err
		}
	}
	return calendars, err
}

//取所有——要修改为支持时间段的，比如某个月份
//未测试！！！20181117修改filter为cond2
func SearchAdminCalendar(title string, public bool) (calendars []*AdminCalendar, err error) {
	cond := orm.NewCondition()
	cond1 := cond.Or("title__contains", title).Or("content__contains", title)
	cond2 := cond.And("title", title).And("public", true)
	cond3 := cond.AndCond(cond1).AndCond(cond2)
	o := orm.NewOrm()
	qs := o.QueryTable("AdminCalendar")
	qs = qs.SetCond(cond3)

	calendars = make([]*AdminCalendar, 0)
	// qs := o.QueryTable("AdminCalendar")
	if public { //只取公开的
		_, err = qs.OrderBy("-Starttime").Distinct().All(&calendars)
		if err != nil {
			return calendars, err
		}
	} else { //取全部
		_, err = qs.OrderBy("-Starttime").Distinct().All(&calendars)
		if err != nil {
			return calendars, err
		}
	}
	return calendars, err
}

//修改
func UpdateAdminCalendar(cid int64, title, content, color string, allday, public bool, start, end time.Time) error {
	o := orm.NewOrm()
	calendar := &AdminCalendar{Id: cid}
	if o.Read(calendar) == nil {
		calendar.Title = title
		calendar.Content = content
		calendar.Color = color
		calendar.Allday = allday
		calendar.Public = public
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
func DropAdminCalendar(cid int64, start, end time.Time) error {
	o := orm.NewOrm()
	calendar := &AdminCalendar{Id: cid}
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
func ResizeAdminCalendar(cid int64, end time.Time) error {
	o := orm.NewOrm()
	calendar := &AdminCalendar{Id: cid}
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
func GetAdminCalendarbyid(id int64) (calendar AdminCalendar, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminCalendar")
	err = qs.Filter("id", id).One(&calendar)
	if err != nil {
		return calendar, err
	}
	return calendar, err
}

//删除事件
func DeleteAdminCalendar(cid int64) error {
	o := orm.NewOrm()
	calendar := &AdminCalendar{Id: cid}
	if o.Read(calendar) == nil {
		_, err := o.Delete(calendar)
		if err != nil {
			return err
		}
	}
	return nil
}

//****项目同步ip****
//添加AdminSynchIp
func AddAdminSynchIp(pid int64, username, synchip, port string) (id int64, err error) {
	o := orm.NewOrm()
	adminsynchip := &AdminSynchIp{
		ParentId: pid,
		UserName: username,
		SynchIp:  synchip,
		Port:     port,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
	id, err = o.Insert(adminsynchip)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//修改
func UpdateAdminSynchIp(cid int64, username, synchip, port string) error {
	o := orm.NewOrm()
	adminsynchip := &AdminSynchIp{Id: cid}
	if o.Read(adminsynchip) == nil {
		adminsynchip.UserName = username
		adminsynchip.SynchIp = synchip
		adminsynchip.Port = port
		adminsynchip.Updated = time.Now()
		_, err := o.Update(adminsynchip)
		if err != nil {
			return err
		}
	}
	return nil
}

//删除
func DeleteAdminSynchIp(cid int64) error {
	o := orm.NewOrm()
	synchip := &AdminSynchIp{Id: cid}
	if o.Read(synchip) == nil {
		_, err := o.Delete(synchip)
		if err != nil {
			return err
		}
	}
	return nil
}

//根据父级id取得所有AdminSynchIp
func GetAdminSynchIp(pid int64) (synchips []*AdminSynchIp, err error) {
	o := orm.NewOrm()
	synchips = make([]*AdminSynchIp, 0)
	qs := o.QueryTable("AdminSynchIp")                 //这个表名AchievementTopic需要用驼峰式，
	_, err = qs.Filter("parentid", pid).All(&synchips) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	return synchips, err
}

//***科室结构********

//添加部门
func AddAdminDepart(pid int64, title, code string) (id int64, err error) {
	o := orm.NewOrm()
	depart := &AdminDepartment{
		ParentId: pid,
		Title:    title,
		Code:     code,
		Created:  time.Now(),
		Updated:  time.Now(),
	}
	id, err = o.Insert(depart)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//修改
func UpdateAdminDepart(cid int64, title, code string) error {
	o := orm.NewOrm()
	category := &AdminDepartment{Id: cid}
	if o.Read(category) == nil {
		category.Title = title
		category.Code = code
		category.Updated = time.Now()
		_, err := o.Update(category)
		if err != nil {
			return err
		}
	}
	return nil
}

//删除
func DeleteAdminDepart(cid int64) error {
	o := orm.NewOrm()
	category := &AdminDepartment{Id: cid}
	if o.Read(category) == nil {
		_, err := o.Delete(category)
		if err != nil {
			return err
		}
	}
	return nil
}

//根据部门id取得所有科室
//如果父级id为0，则取所有部门
func GetAdminDepart(pid int64) (departs []*AdminDepartment, err error) {
	o := orm.NewOrm()
	departs = make([]*AdminDepartment, 0)
	qs := o.QueryTable("AdminDepartment")
	_, err = qs.Filter("parentid", pid).All(&departs)
	if err != nil {
		return nil, err
	}
	return departs, err
}

//根据父级id取得所有
//如果父级id为空，则取所有一级category
// func GetAdminCategory(pid int64) (categories []*AdminDepartment, err error) {
// 	o := orm.NewOrm()
// 	categories = make([]*AdminDepartment, 0)
// 	qs := o.QueryTable("AdminDepartment")
// 	_, err = qs.Filter("parentid", pid).All(&categories)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return categories, err
// }
//根据部门名字title返回自身
func GetAdminDepartName(title string) (AdminDepartment, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminDepartment")
	var cate AdminDepartment
	err := qs.Filter("title", title).One(&cate)
	if err != nil {
		return cate, err
	}
	return cate, err
}

//根据部门名字title查询所有下级科室category
func GetAdminDepartTitle(title string) (categories []*AdminDepartment, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminDepartment")
	var cate AdminDepartment
	err = qs.Filter("title", title).One(&cate)
	categories = make([]*AdminDepartment, 0)
	_, err = qs.Filter("parentid", cate.Id).All(&categories)
	if err != nil {
		return nil, err
	}
	return categories, err
}

//根据id查科室
func GetAdminDepartbyId(id int64) (category AdminDepartment, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("AdminDepartment")
	err = qs.Filter("id", id).One(&category)
	if err != nil {
		return category, err
	}
	return category, err
}

//由分院id和科室 名称取得科室
func GetAdminDepartbyidtitle(id int64, title string) (*AdminDepartment, error) {
	o := orm.NewOrm()
	// cate := &Category{Id: id}
	category := new(AdminDepartment)
	qs := o.QueryTable("AdminDepartment")
	err := qs.Filter("parentid", id).Filter("title", title).One(category)
	if err != nil {
		return nil, err
	}
	return category, err
}

//添加轮播动画
func AddAdminCarousel(title, url string) (id int64, err error) {
	o := orm.NewOrm()
	carousel := &AdminCarousel{
		Title: title,
		// Path:    code,
		Url:     url,
		Created: time.Now(),
		Updated: time.Now(),
	}
	id, err = o.Insert(carousel)
	if err != nil {
		return 0, err
	}
	// }
	return id, nil
}

func GetAdminCarousel() (carousels []*AdminCarousel, err error) {
	o := orm.NewOrm()
	carousels = make([]*AdminCarousel, 0)
	qs := o.QueryTable("AdminCarousel")
	_, err = qs.OrderBy("-created").All(&carousels)
	if err != nil {
		return nil, err
	}
	return carousels, err
}

func DelAdminCarouselById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&AdminCarousel{Id: Id})
	return status, err
}

//添加提交给merit的ip地址和端口号，接口名称
