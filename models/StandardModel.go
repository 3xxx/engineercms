package models

import (
	//"database/sql"
	//"github.com/astaxie/beedb"
	//_ "github.com/ziutek/mymysql/godrv"
	//"time"
	// "fmt"
	// "os"
	// "path"
	// "github.com/astaxie/beego"
	"strconv"
	// "strings"
	"time"
	//"github.com/Unknwon/com
	// "errors"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
	_ "github.com/mattn/go-sqlite3"
)

//const (
//	_DB_NAME        = "database/orm_test.db"
//	_SQLITE3_DRIVER = "sqlite3"
//)

type Standard struct {
	Id       int64
	Number   string //`orm:"unique"`
	Title    string
	Uid      int64
	Category string
	Content  string `orm:"sie(5000)"`
	Route    string
	// AttachmentId int64
	// Attachments     []*Attachment `orm:"reverse(many)"` // fk 的反向关系
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	Views   int64
}

type Library struct {
	Id       int64
	Number   string //规范的编号`orm:"unique"`
	Title    string
	LiNumber string    //完整编号，含年份
	Category string    //行业分类
	Year     string    //编号里的年份
	Execute  string    //执行时间
	Content  string    `orm:"sie(5000)"`
	Created  time.Time `orm:"auto_now_add;type(datetime)"`
	Updated  time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Standard), new(Library)) //, new(Attachment), new(Article)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/standardms.db", 10)
}

//标准存入数据库
func SaveStandard(standard Standard) (sid int64, err error) {
	o := orm.NewOrm()
	//判断是否有重名
	// var spider Spider //下面这个filter放在topic=&Topic{后面用返回one(topic)则查询出错！
	//只有编号和主机都不同才写入。
	err = o.QueryTable("standard").Filter("number", standard.Number).Filter("title", standard.Title).One(&standard, "Id")
	// err = o.QueryTable("topic").Filter("categoryid", cid).Filter("tnumber", tnumber).One(&topic, "Id")
	if err == orm.ErrNoRows { //Filter("tnumber", tnumber).One(topic, "Id")==nil则无法建立
		// 没有找到记录
		// spider1 := &Spider{
		// 	Number:   number,
		// 	Name:     name,
		// 	Link:     link,
		// 	UserName: username,
		// 	UserIp:   userip,
		// 	Created:  time.Now(),
		// 	Updated:  time.Now(),
		// }
		sid, err = o.Insert(&standard)
		if err != nil {
			return 0, err //如果文章编号相同，则唯一性检查错误，返回id吗？
		}
	} else { //进行更新
		// _, err = o.Update(cate)
		sid, err = o.Update(&standard)
	}
	return sid, err
	// 原来的代码orm := orm.NewOrm()
	// // fmt.Println(user)
	// uid, err = orm.Insert(&user) //_, err = o.Insert(reply)
	// return uid, err
}

//有效版本库存入数据库
func SaveLibrary(library Library) (lid int64, err error) {
	o := orm.NewOrm()
	//关闭写同步
	o.Raw("PRAGMA synchronous = OFF; ", 0, 0, 0).Exec()
	//判断是否有重名
	// var spider Spider //下面这个filter放在topic=&Topic{后面用返回one(topic)则查询出错！
	//只有编号和主机都不同才写入。
	var library1 Library //数据库中已有的数据
	err = o.QueryTable("library").Filter("number", library.Number).Filter("title", library.Title).Filter("category", library.Category).One(&library1)
	// err = o.QueryTable("topic").Filter("categoryid", cid).Filter("tnumber", tnumber).One(&topic, "Id")
	if err == orm.ErrNoRows { //Filter("tnumber", tnumber).One(topic, "Id")==nil则无法建立
		// 没有找到记录
		// spider1 := &Spider{
		// 	Number:   number,
		// 	Name:     name,
		// 	Link:     link,
		// 	UserName: username,
		// 	UserIp:   userip,
		// 	Created:  time.Now(),
		// 	Updated:  time.Now(),
		// }
		lid, err = o.Insert(&library)
		if err != nil {
			return 0, err //如果文章编号相同，则唯一性检查错误，返回id吗？
		}
	} else { //如果有记录，则进行判断年份，还要判断是否有值
		//判断年份，如果是新则进行更新操作
		//如果年份旧，就判断有无值，没有值的才更新
		// 年份string转成int
		if library1.Year != "" { //如果数据库中的year有值，则进行判断
			Year, err := strconv.Atoi(library.Year)
			if err != nil {
				return 0, err
			}
			Year1, err := strconv.Atoi(library1.Year)
			if err != nil {
				return 0, err
			}
			if Year >= Year1 { //如果比数据库中的新，进行更新操作
				// library1.LiNumber = library.LiNumber //完整编号，含年份
				library1.Year = library.Year
				library1.Execute = library.Execute
				library1.Updated = time.Now()
				lid, err = o.Update(&library1)
			}
		} else { //如果数据库中的year没有值，则直接进行update
			// library1.LiNumber = library.LiNumber //完整编号，含年份
			library1.Year = library.Year
			library1.Execute = library.Execute
			library1.Updated = time.Now()
			lid, err = o.Update(&library1)
		}
	}
	return lid, err
	// 原来的代码orm := orm.NewOrm()
	// // fmt.Println(user)
	// uid, err = orm.Insert(&user) //_, err = o.Insert(reply)
	// return uid, err
}

//由名字模糊搜索
func SearchStandardsName(name string, isDesc bool) ([]*Standard, error) {
	o := orm.NewOrm()
	Standards := make([]*Standard, 0)
	qs := o.QueryTable("Standard")
	var err error
	if isDesc {
		if len(name) > 0 {
			qs = qs.Filter("Title__contains", name) //这里取回
		}
		_, err = qs.OrderBy("-created").All(&Standards)
	} else {
		_, err = qs.Filter("Title__contains", name).OrderBy("-created").All(&Standards)
		//o.QueryTable("user").Filter("name", "slene").All(&users)
	}
	return Standards, err
}

//由名字模糊搜索
func SearchStandardsNamePage(limit, offset int64, name string, isDesc bool) ([]*Standard, error) {
	o := orm.NewOrm()
	Standards := make([]*Standard, 0)
	qs := o.QueryTable("Standard")
	var err error
	if isDesc {
		if len(name) > 0 {
			qs = qs.Filter("Title__contains", name) //这里取回
		}
		_, err = qs.Limit(limit, offset).OrderBy("-created").All(&Standards)
	} else {
		_, err = qs.Limit(limit, offset).Filter("Title__contains", name).OrderBy("-created").All(&Standards)
		//o.QueryTable("user").Filter("name", "slene").All(&users)
	}
	return Standards, err
}

//由编号模糊搜索
func SearchStandardsNumber(number string, isDesc bool) ([]*Standard, error) {
	o := orm.NewOrm()
	Standards := make([]*Standard, 0)
	qs := o.QueryTable("Standard")
	var err error
	if isDesc {
		if len(number) > 0 {
			qs = qs.Filter("Number__contains", number) //这里取回
		}
		_, err = qs.OrderBy("-created").All(&Standards)
	} else {
		_, err = qs.Filter("Number__contains", number).OrderBy("-created").All(&Standards)
		//o.QueryTable("user").Filter("name", "slene").All(&users)
	}
	return Standards, err
}

//由编号模糊搜索
func SearchStandardsNumberPage(limit, offset int64, number string, isDesc bool) ([]*Standard, error) {
	o := orm.NewOrm()
	Standards := make([]*Standard, 0)
	qs := o.QueryTable("Standard")
	var err error
	if isDesc {
		if len(number) > 0 {
			qs = qs.Filter("Number__contains", number) //这里取回
		}
		_, err = qs.Limit(limit, offset).OrderBy("-created").All(&Standards)
	} else {
		_, err = qs.Limit(limit, offset).Filter("Number__contains", number).OrderBy("-created").All(&Standards)
		//o.QueryTable("user").Filter("name", "slene").All(&users)
	}
	return Standards, err
}

//由分类SL和编号搜索有效版本库
func SearchLiabraryNumber(Category, Number string) (*Library, error) {
	o := orm.NewOrm()
	library := new(Library)
	qs := o.QueryTable("library")
	err := qs.Filter("category", Category).Filter("number", Number).One(library)
	if err != nil {
		return nil, err
	}
	return library, err
}

func GetAllStandards() ([]*Standard, error) {
	o := orm.NewOrm()
	standards := make([]*Standard, 0)
	qs := o.QueryTable("standard")
	var err error
	//这里进行过滤，parentid为空的才显示
	// qs = qs.Filter("ParentId", 0)
	_, err = qs.OrderBy("-created").All(&standards)
	// _, err := qs.All(&cates)
	return standards, err
}

//根据id取得规范
func GetStandard(id int64) (standard Standard, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Standard")
	err = qs.Filter("id", id).One(&standard)
	if err != nil {
		return standard, err
	}
	return standard, err
}

func UpdateStandard(id int64, number, title, route string) error {
	o := orm.NewOrm()
	standard := &Standard{Id: id}
	if o.Read(standard) == nil {
		standard.Number = number
		standard.Title = title
		standard.Updated = time.Now()
		standard.Route = route
		_, err := o.Update(standard)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteStandard(id int64) error {
	var err error
	o := orm.NewOrm()
	standard := Standard{Id: id}
	if o.Read(&standard) == nil {
		_, err = o.Delete(&standard)
		if err != nil {
			return err
		}
	}
	return err
}

//由法规名称精确搜索有效版本库
func SearchLiabraryName(Name string) ([]*Library, error) {
	o := orm.NewOrm()
	// library := new(Library)
	var libraries []*Library
	qs := o.QueryTable("library")
	_, err := qs.Filter("title", Name).All(&libraries) //(library)这样也可以。多条时会出错
	// if err == orm.ErrMultiRows {
	// 	// 多条的时候报错
	// 	// fmt.Printf("Returned Multi Rows Not One")
	// 	return nil, err
	// }
	// if err == orm.ErrNoRows {
	// 	// 没有找到记录
	// 	// fmt.Printf("Not row found")
	// 	return nil, err
	// }
	if err != nil {
		return nil, err
	}
	return libraries, err
}

func GetAllValids() ([]*Library, error) {
	o := orm.NewOrm()
	librarys := make([]*Library, 0)
	qs := o.QueryTable("library")
	var err error
	//这里进行过滤，parentid为空的才显示
	// qs = qs.Filter("ParentId", 0)
	_, err = qs.OrderBy("-created").Limit(10000).All(&librarys)
	// _, err := qs.All(&cates)
	return librarys, err
}

func DeleteValid(id int64) error {
	var err error
	o := orm.NewOrm()
	// Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
	// user := User{Name: "slene"}
	// err = o.Read(&user, "Name")
	valid := Library{Id: id}
	if o.Read(&valid) == nil {
		_, err = o.Delete(&valid)
		if err != nil {
			return err
		}
	}
	return err
}

//删除_把附件也一并删除（在controllers中实现吧）
// func DeleteProduct(cid int64) error {
// 	o := orm.NewOrm()
// 	product := &Product{Id: cid}
// 	if o.Read(product) == nil {
// 		_, err := o.Delete(product)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
