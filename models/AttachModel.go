package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	// "strconv"
	// "strings"
	"time"
)

//附件,attachment 和 topic 是 ManyToOne 关系，也就是 ForeignKey 为 topic
type Attachment struct {
	Id        int64
	FileName  string
	FileSize  int64
	Downloads int64
	ProductId int64     //*Topic    `orm:"rel(fk)"`
	Created   time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"index","auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Attachment)) //, new(Article)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/engineer.db", 10)
}

//添加附件到成果id下
//如果附件名称已经存在，则不再追加写入数据库
//应该用ReadOrCreate尝试从数据库读取，不存在的话就创建一个

func AddAttachment(filename string, filesize, downloads, productid int64) (id int64, err error) {
	o := orm.NewOrm()
	var attach Attachment
	err = o.QueryTable("Attachment").Filter("filename", filename).One(&attach)
	if err == orm.ErrNoRows { // 没有找到记录
		attachment := &Attachment{
			FileName:  filename,
			FileSize:  filesize,
			Downloads: downloads,
			ProductId: productid,
			Created:   time.Now(),
			Updated:   time.Now(),
		}
		id, err = o.Insert(attachment)
		if err != nil {
			return 0, err
		}
	} else if err == orm.ErrMultiRows {
		return 0, err
	} else if err == nil {
		return attach.Id, err
	}
	return id, err
}

//修改——还没改
func UpdateAttachment(cid int64, title, code string, grade int) (err error) {
	o := orm.NewOrm()
	//id转成64为
	// cidNum, err := strconv.ParseInt(cid, 10, 64)
	// if err != nil {
	// 	return err
	// }
	Attachment := &Attachment{Id: cid}
	if o.Read(Attachment) == nil {
		// Attachment.Title = title
		// Attachment.Code = code
		// Attachment.Grade = grade
		Attachment.Updated = time.Now()
		_, err := o.Update(Attachment)
		if err != nil {
			return err
		}
	}
	return err
}

//根据id取得所有附件
func GetAttachments(id int64) (attachments []*Attachment, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Attachment")
	_, err = qs.Filter("Productid", id).All(&attachments)
	if err != nil {
		return nil, err
	}
	return attachments, err
}

//根据名字title查询到项目目录
func GetAttachmentTitle(title string) (cate Attachment, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Attachment")
	// var cate Attachment
	err = qs.Filter("title", title).One(&cate)
	// if pid != "" {
	// cate := Attachment{Title: title}这句无效
	// categories = make([]*Attachment, 0)
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
