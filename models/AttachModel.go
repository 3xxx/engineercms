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
	err = o.QueryTable("Attachment").Filter("productid", productid).Filter("filename", filename).One(&attach)
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

//修改_这个意义不大，对于附件的修改，一般是追加或删除
func UpdateAttachment(cid int64, filename string, filesize, downloads int64) (err error) {
	o := orm.NewOrm()
	attachment := &Attachment{Id: cid}
	if o.Read(attachment) == nil {
		attachment.FileName = filename
		attachment.FileSize = filesize
		attachment.Downloads = downloads
		attachment.Updated = time.Now()
		_, err := o.Update(attachment)
		if err != nil {
			return err
		}
	}
	return err
}

//删除
func DeleteAttachment(cid int64) error {
	o := orm.NewOrm()
	attachment := &Attachment{Id: cid}
	if o.Read(attachment) == nil {
		_, err := o.Delete(attachment)
		if err != nil {
			return err
		}
	}
	return nil
}

//根据成果id取得所有附件
func GetAttachments(id int64) (attachments []*Attachment, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Attachment")
	_, err = qs.Filter("Productid", id).All(&attachments)
	if err != nil {
		return nil, err
	}
	return attachments, err
}

//根据名字title查询到附件
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

//根据附件id查询附件
func GetAttachbyId(Id int64) (attach Attachment, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Attachment")
	err = qs.Filter("id", Id).One(&attach)
	return attach, err
}
