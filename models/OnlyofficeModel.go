package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// _ "github.com/mattn/go-sqlite3"
	// "strconv"
	// "strings"
	"time"
)

type OnlyOffice struct {
	Id        int64     `form:"-"`
	Code      string    `orm:"null"`
	Title     string    `orm:"null"`
	Label     string    `orm:"null"`
	End       time.Time `orm:"null;type(datetime)"`
	Principal string    `orm:"null"`
	// Ext     string    `orm:"null"`
	Uid     int64     `orm:"null"`
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

//附件
type OnlyAttachment struct {
	Id        int64
	FileName  string
	FileSize  int64
	Downloads int64
	DocId     int64 //*Topic `orm:"rel(fk)"`
	// Changesurl string    `orm:"null"` //文件修改记录
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

//历史版本
type OnlyHistory struct {
	Id            int64
	AttachId      int64
	UserId        int64
	ServerVersion string
	Version       int
	FileUrl       string
	ChangesUrl    string    //`orm:"null"`
	HistoryKey    string    `orm:"sie(19)"`
	Expires       time.Time `orm:"type(datetime)"`
	Created       time.Time `orm:"type(datetime)"`
}

//修改情况
type OnlyChanges struct {
	Id         int64
	HistoryKey string `orm:"sie(19)"`
	UserId     string `orm:"sie(10)"`
	UserName   string `orm:"sie(20)"`
	Created    string `orm:"sie(19)"`
}

func init() {
	orm.RegisterModel(new(OnlyOffice), new(OnlyAttachment), new(OnlyHistory), new(OnlyChanges))
}

//取得所有项目
func GetDocs() (docs []*OnlyOffice, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("OnlyOffice")
	_, err = qs.All(&docs)
	if err != nil {
		return docs, err
	}
	return docs, err
}

//添加成果到项目侧栏某个id下
//如果这个侧栏id下的这个成果编号已经存在，则返回id
////应该用ReadOrCreate尝试从数据库读取，不存在的话就创建一个
func AddDoc(code, title, label, principal string, end time.Time, uid int64) (id int64, err error) {
	o := orm.NewOrm()
	var prod OnlyOffice
	err = o.QueryTable("OnlyOffice").Filter("code", code).One(&prod)
	if err == orm.ErrNoRows { // 没有找到记录
		product := &OnlyOffice{
			Code:      code,
			Title:     title,
			Label:     label,
			Principal: principal,
			End:       end,
			Uid:       uid,
			Created:   time.Now(),
			Updated:   time.Now(),
		}
		id, err = o.Insert(product)
		if err != nil {
			return 0, err
		}
	} else if err == orm.ErrMultiRows {
		return 0, err
	} else if err == nil {
		return prod.Id, err
	}
	return id, err
}

//根据docid获得docuid
func Getdocbyid(id int64) (doc OnlyOffice, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("OnlyOffice")
	err = qs.Filter("id", id).One(&doc)
	return doc, err
}

//添加附件到成果id下
//如果附件名称已经存在，则不再追加写入数据库
//应该用ReadOrCreate尝试从数据库读取，不存在的话就创建一个
func AddOnlyAttachment(filename string, filesize, downloads, productid int64) (id int64, err1, err2 error) {
	o := orm.NewOrm()
	var attach OnlyAttachment
	err1 = o.QueryTable("OnlyAttachment").Filter("DocId", productid).Filter("filename", filename).One(&attach)
	if err1 == orm.ErrNoRows { // 没有找到记录
		attachment := &OnlyAttachment{
			FileName:  filename,
			FileSize:  filesize,
			Downloads: downloads,
			DocId:     productid,
			Created:   time.Now(),
			Updated:   time.Now(),
		}
		id, err2 = o.Insert(attachment)
		if err2 != nil {
			return 0, err1, err2
		}
	} else if err1 == orm.ErrMultiRows {
		return 0, err1, err2
	} else if err1 == nil {
		return attach.Id, err1, err2
	}
	return id, err1, err2
}

//根据成果id取得所有附件
func GetOnlyAttachments(id int64) (attachments []*OnlyAttachment, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("OnlyAttachment")
	_, err = qs.Filter("DocId", id).All(&attachments)
	if err != nil {
		return nil, err
	}
	return attachments, err
}

//根据附件id查询附件
func GetOnlyAttachbyId(Id int64) (attach OnlyAttachment, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("OnlyAttachment")
	err = qs.Filter("id", Id).One(&attach)
	return attach, err
}

//修改附件的日期和changesurl修改记录地址
func UpdateOnlyAttachment(cid int64) (err error) {
	o := orm.NewOrm()
	attachment := &OnlyAttachment{Id: cid}
	if o.Read(attachment) == nil {
		attachment.Updated = time.Now()
		// if changesurl != "" {
		// 	attachment.Changesurl = changesurl
		// 	_, err = o.Update(attachment, "Updated", "Changesurl")
		// } else {
		_, err = o.Update(attachment, "Updated")
		// }
		if err != nil {
			return err
		}
	}
	return err
}

//修改成果信息
func UpdateDoc(cid int64, code, title, label, principal string, end time.Time) error {
	o := orm.NewOrm()
	product := &OnlyOffice{Id: cid}
	if o.Read(product) == nil {
		product.Code = code
		product.Title = title
		product.Label = label
		product.Principal = principal
		product.End = end
		product.Updated = time.Now()
		_, err := o.Update(product)
		if err != nil {
			return err
		}
	}
	return nil
}

//修改成果时间信息
func UpdateDocTime(cid int64) error {
	o := orm.NewOrm()
	product := &OnlyOffice{Id: cid}
	if o.Read(product) == nil {
		product.Updated = time.Now()
		_, err := o.Update(product, "Updated")
		if err != nil {
			return err
		}
	}
	return nil
}

//删除_把附件也一并删除（在controllers中实现吧）
func DeleteDoc(cid int64) error {
	o := orm.NewOrm()
	product := &OnlyOffice{Id: cid}
	if o.Read(product) == nil {
		_, err := o.Delete(product)
		if err != nil {
			return err
		}
	}
	return nil
}

//删除附件
func DeleteOnlyAttachment(cid int64) error {
	o := orm.NewOrm()
	attachment := &OnlyAttachment{Id: cid}
	if o.Read(attachment) == nil {
		_, err := o.Delete(attachment)
		if err != nil {
			return err
		}
	}
	return nil
}

//添加历史版本
func AddOnlyHistory(onlyattachmentid, uid int64, serverversion string, version int, key, fileurl, changesurl string, expires, created time.Time) (id int64, err1, err2 error) {
	o := orm.NewOrm()
	var history OnlyHistory
	err1 = o.QueryTable("OnlyHistory").Filter("HistoryKey", key).One(&history)
	if err1 == orm.ErrNoRows { // 没有找到记录
		onlyhistory := &OnlyHistory{
			AttachId:      onlyattachmentid,
			UserId:        uid,
			ServerVersion: serverversion,
			Version:       version,
			HistoryKey:    key,
			FileUrl:       fileurl,
			ChangesUrl:    changesurl,
			Expires:       expires,
			Created:       created,
		}
		id, err2 = o.Insert(onlyhistory)
		if err2 != nil {
			return 0, err1, err2
		}
	} else if err1 == orm.ErrMultiRows {
		return 0, err1, err2
	} else if err1 == nil {
		return history.Id, err1, err2
	}
	return id, err1, err2
}

//写入附件文档的修改记录地址
// func UpdateChangesUrl(onlyattachmentid int64, changesurl string) error {
// 	o := orm.NewOrm()
// 	onlyattachment := &OnlyAttachment{Id: onlyattachmentid}
// 	if o.Read(onlyattachment) == nil {
// 		onlyattachment.Changesurl = changesurl
// 		_, err := o.Update(onlyattachment, "Changesurl")
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

//根据附件id获取历史版本信息
func GetOnlyHistory(onlyattachmentid int64) (onlyhistories []*OnlyHistory, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("OnlyHistory")
	_, err = qs.Filter("AttachId", onlyattachmentid).All(&onlyhistories)
	if err != nil {
		return nil, err
	}
	return onlyhistories, err
}

//根据附件id和version获取历史changeurl
// func GetOnlyChangesUrl(onlyattachmentid int64, version int) (onlyhistory *OnlyHistory, err error) {
// 	o := orm.NewOrm()
// 	qs := o.QueryTable("OnlyHistory")
// 	err = qs.Filter("AttachId", onlyattachmentid).Filter("Version", version).One(&onlyhistory)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return onlyhistory, err
// }

//获取附件id历史版本号
func GetOnlyHistoryVersion(onlyattachmentid int64) (onlyhistories []OnlyHistory, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("OnlyHistory")
	_, err = qs.Filter("AttachId", onlyattachmentid).All(&onlyhistories, "Version")
	if err != nil {
		return nil, err
	}
	return onlyhistories, err
}

//添加历史版本
func AddOnlyChanges(key, uid, uname, created string) (id int64, err1, err2 error) {
	o := orm.NewOrm()
	var changes OnlyChanges
	err1 = o.QueryTable("OnlyChanges").Filter("HistoryKey", key).One(&changes)
	if err1 == orm.ErrNoRows { // 没有找到记录
		onlychanges := &OnlyChanges{
			UserId:     uid,
			UserName:   uname,
			HistoryKey: key,
			Created:    created,
		}
		id, err2 = o.Insert(onlychanges)
		if err2 != nil {
			return 0, err1, err2
		}
	} else if err1 == orm.ErrMultiRows {
		return 0, err1, err2
	} else if err1 == nil {
		return changes.Id, err1, err2
	}
	return id, err1, err2
}

//根据附件历史版本key获取历史版本修改信息
func GetOnlyChanges(key string) (onlychanges []*OnlyChanges, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("OnlyChanges")
	_, err = qs.Filter("HistoryKey", key).All(&onlychanges)
	if err != nil {
		return nil, err
	}
	return onlychanges, err
}
