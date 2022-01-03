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
	// _ "github.com/mattn/go-sqlite3"
)

//const (
//	_DB_NAME        = "database/orm_test.db"
//	_SQLITE3_DRIVER = "sqlite3"
//)

type Wiki struct {
	Id                int64
	Uid               int64
	Title             string
	Content           string `orm:"sie(5000)"`
	Attachment        string
	Created           time.Time `orm:"auto_now_add;type(datetime)"`
	Updated           time.Time `orm:"auto_now;type(datetime)"`
	Views             int64
	Author            string
	ReplyTime         time.Time
	ReplyCount        int64
	ReplyLastUserName string
	// Attachments     []*Attachment `orm:"reverse(many)"` // fk 的反向关系
}

func init() {
	orm.RegisterModel(new(Wiki)) //, new(Article)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/orm_test.db", 10)
}

//一对一模式
func AddWikiOne(title, content, uname string) (id int64, err error) {
	o := orm.NewOrm()
	// var wiki Wiki //下面这个filter放在wiki=&Wiki{后面用返回one(wiki)则查询出错！
	wiki1 := &Wiki{
		Title:     title,
		Content:   content,
		Author:    uname,
		Created:   time.Now(),
		Updated:   time.Now(),
		ReplyTime: time.Now(), //这里不要的话写不了数据库，也不提示错误。
		// Attachment: attachment,
	}
	id, err = o.Insert(wiki1)
	if err != nil {
		return id, err
	}
	return id, err
}

//一对多模式
func AddWikiMany(title, uname, content, attachment string) (id int64, err error) {
	o := orm.NewOrm()
	// var wiki Wiki //下面这个filter放在wiki=&Wiki{后面用返回one(wiki)则查询出错！
	wiki1 := &Wiki{
		Title:      title,
		Content:    content,
		Attachment: attachment,
		Author:     uname,
		Created:    time.Now(),
		Updated:    time.Now(),
		ReplyTime:  time.Now(),
	}
	id, err = o.Insert(wiki1)
	if err != nil {
		return id, err //如果文章编号相同，则唯一性检查错误，返回id吗？
	}
	return id, err
}

func DeletWiki(tid string) error { //应该在controllers中显示警告
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	// Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
	// user := User{Name: "slene"}
	// err = o.Read(&user, "Name")
	wiki := Wiki{Id: tidNum}
	if o.Read(&wiki) == nil {
		_, err = o.Delete(&wiki)
		if err != nil {
			return err
		}
	}

	// attachment := Attachment{TopicId: tidNum}
	// if o.Read(&attachment, "WikiId") == nil {
	// oldCate = wiki.Category
	// _, err = o.Delete(&attachment)
	// if err != nil {
	// return err
	// }
	// }
	// _, err = o.Delete(&wiki) //这句为何重复？
	return err
}

//删除文章中的附件
func DeletAttachment(aid string) error { //应该显示警告
	aidNum, err := strconv.ParseInt(aid, 10, 64)
	if err != nil {
		return err
	}

	// var oldCate string
	o := orm.NewOrm()
	// Read 默认通过查询主键赋值，可以使用指定的字段进行查询：
	// user := User{Name: "slene"}
	// err = o.Read(&user, "Name")

	// topic := Topic{Id: tidNum}
	// if o.Read(&topic) == nil {
	// 	oldCate = topic.Category
	// 	_, err = o.Delete(&topic)
	// 	if err != nil {
	// 		return err
	// 	}
	// }

	attachment := Attachment{Id: aidNum}
	if o.Read(&attachment) == nil {
		// oldCate = topic.Category
		_, err = o.Delete(&attachment)
		if err != nil {
			return err
		}
	}

	// if len(oldCate) > 0 {
	// 	cate := new(Category)
	// 	qs := o.QueryTable("category")
	// 	err = qs.Filter("title", oldCate).One(cate)
	// 	if err == nil {
	// 		cate.TopicCount--
	// 		_, err = o.Update(cate)
	// 	}
	// }
	// _, err = o.Delete(&topic) //这句为何重复？
	return err
}

//取出分页的wiki
func ListWikisByOffsetAndLimit(set, postsPerPage int) ([]*Wiki, error) {
	o := orm.NewOrm()
	topics := make([]*Wiki, 0)
	qs := o.QueryTable("wiki")
	var err error
	_, err = qs.Limit(postsPerPage, set).OrderBy("-created").All(&topics)
	return topics, err
}

//缺少排序，由项目名称获取项目下所有成果，如果没有项目名称，则获取所有成果
func GetAllWikis(isDesc bool) ([]*Wiki, error) {
	o := orm.NewOrm()
	wikis := make([]*Wiki, 0)
	qs := o.QueryTable("wiki")
	var err error
	// _, err = qs.OrderBy("-created").All(&wikis)
	_, err = qs.OrderBy("-reply_time").All(&wikis)
	return wikis, err
}

//取出分页的文章
// func ListPostsByOffsetAndLimit(set, postsPerPage int) ([]*Wiki, error) {
// 	o := orm.NewOrm()
// 	wikis := make([]*Wiki, 0)
// 	qs := o.QueryTable("wiki")
// 	var err error
// 	_, err = qs.Limit(postsPerPage, set).OrderBy("-created").All(&wikis)
// 	return wikis, err
// }

func SearchWikis(tuming string, isDesc bool) ([]*Wiki, error) {
	o := orm.NewOrm()
	wikis := make([]*Wiki, 0)
	qs := o.QueryTable("wiki")
	var err error
	if isDesc {
		if len(tuming) > 0 {
			qs = qs.Filter("Title__contains", tuming) //这里取回
		}
		_, err = qs.OrderBy("-created").All(&wikis)
	} else {
		_, err = qs.Filter("Title__contains", tuming).OrderBy("-created").All(&wikis)
		//o.QueryTable("user").Filter("name", "slene").All(&users)
	}
	return wikis, err
}

//查看一个文章
func GetWiki(tid string) (*Wiki, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	wiki := new(Wiki)
	qs := o.QueryTable("wiki")
	err = qs.Filter("id", tidNum).One(wiki)
	if err != nil {
		return nil, err
	}
	wiki.Views++
	_, err = o.Update(wiki)

	return wiki, err
}

//由用户名取得文章
func Getwikisbyuname(uname string) ([]*Wiki, error) {
	o := orm.NewOrm()
	wikis := make([]*Wiki, 0)
	qs := o.QueryTable("wiki")
	var err error
	qs = qs.Filter("Author", uname)
	_, err = qs.OrderBy("-created").All(&wikis)
	// _, err := qs.All(&cates)
	return wikis, err
}

//只修改编号、名称和内容，不修改附件及分类树状目录
func ModifyWiki(tid, title, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	wiki := &Wiki{Id: tidNum}
	if o.Read(wiki) == nil {
		wiki.Title = title
		wiki.Content = content
		// wiki.Attachment = attachment
		wiki.Updated = time.Now()
		_, err = o.Update(wiki)
		if err != nil {
			return err
		}
	}
	//删除旧的附件
	// if len(oldAttach) > 0 {
	// 	os.Remove(path.Join("attachment", oldAttach))
	// }
	return err
}

//设代日记：由图片附件的id，存入图片的content
// func ModifyAttachment(aid, content string) error {
// 	aidNum, err := strconv.ParseInt(aid, 10, 64)
// 	if err != nil {
// 		return err
// 	}
// 	o := orm.NewOrm()
// 	attachment := &Attachment{Id: aidNum}
// 	if o.Read(attachment) == nil {
// 		attachment.Content = content
// 		attachment.Updated = time.Now()
// 		_, err = o.Update(attachment)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return err
// }

//删除文章中的附件
// func DeletAttachment(aid string) error { //应该显示警告
// 	aidNum, err := strconv.ParseInt(aid, 10, 64)
// 	if err != nil {
// 		return err
// 	}
// 	o := orm.NewOrm()
// 	attachment := Attachment{Id: aidNum}
// 	if o.Read(&attachment) == nil {
// 		// oldCate = wiki.Category
// 		_, err = o.Delete(&attachment)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	// _, err = o.Delete(&wiki) //这句为何重复？
// 	return err
// }
