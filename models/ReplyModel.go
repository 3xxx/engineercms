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

type Topic struct {
	Id                int64
	Uid               int64
	Title             string
	Tnumber           string //`orm:"unique"`
	Category          string
	CategoryId        int64
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

//文章评论
type Commenttopic struct {
	Id      int64
	Tid     int64
	Name    string
	Content string `orm:"size(1000)"`
	Created time.Time
}

//wiki评论
type Commentwiki struct {
	Id      int64
	Tid     int64
	Name    string
	Content string `orm:"size(1000)"`
	Created time.Time
}

func init() {
	orm.RegisterModel(new(Commenttopic), new(Commentwiki))
}

func DeleteTopicReply(rid string) error {
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()

	var tidNum int64
	reply := &Commenttopic{Id: ridNum}
	if o.Read(reply) == nil {
		tidNum = reply.Tid
		_, err = o.Delete(reply)
		if err != nil {
			return err
		}
	}
	replies := make([]*Commenttopic, 0) //slice,将所有文章取出来
	qs := o.QueryTable("commenttopic")
	_, err = qs.Filter("tid", tidNum).OrderBy("-created").All(&replies)
	if err != nil {
		return err
	}
	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil { //如果错误为空
		if len(replies) != 0 { //如果回复不为空，则……
			topic.ReplyTime = replies[0].Created
			topic.ReplyLastUserName = replies[0].Name
			topic.ReplyCount = int64(len(replies))
			_, err = o.Update(topic)
		}
	}
	return err
}

func AddTopicReply(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	reply := &Commenttopic{
		Tid:     tidNum,
		Name:    nickname,
		Content: content,
		Created: time.Now(),
	}
	o := orm.NewOrm()
	_, err = o.Insert(reply)
	if err != nil {
		return err
	}

	topic := &Topic{Id: tidNum}
	if o.Read(topic) == nil {
		topic.ReplyTime = time.Now()
		topic.ReplyLastUserName = nickname
		topic.ReplyCount++
		_, err = o.Update(topic)
	}
	return err
}

func GetAllTopicReplies(tid string) (replies []*Commenttopic, err error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	replies = make([]*Commenttopic, 0)

	o := orm.NewOrm()
	qs := o.QueryTable("commenttopic")
	_, err = qs.Filter("tid", tidNum).All(&replies)
	return replies, err
}

func DeleteWikiReply(rid string) error {
	ridNum, err := strconv.ParseInt(rid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()

	var tidNum int64
	reply := &Commentwiki{Id: ridNum}
	if o.Read(reply) == nil {
		tidNum = reply.Tid
		_, err = o.Delete(reply)
		if err != nil {
			return err
		}
	}
	replies := make([]*Commentwiki, 0) //slice,将所有文章取出来
	qs := o.QueryTable("commentwiki")
	_, err = qs.Filter("tid", tidNum).OrderBy("-created").All(&replies)
	if err != nil {
		return err
	}
	wiki := &Wiki{Id: tidNum}
	if o.Read(wiki) == nil { //如果错误为空
		if len(replies) != 0 { //如果回复不为空，则……
			wiki.ReplyTime = replies[0].Created
			wiki.ReplyLastUserName = replies[0].Name
			wiki.ReplyCount = int64(len(replies))
			_, err = o.Update(wiki)
		}
	}
	return err
}

func AddWikiReply(tid, nickname, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	reply := &Commentwiki{
		Tid:     tidNum,
		Name:    nickname,
		Content: content,
		Created: time.Now(),
	}
	o := orm.NewOrm()
	_, err = o.Insert(reply)
	if err != nil {
		return err
	}
	topic := &Wiki{Id: tidNum}
	if o.Read(topic) == nil {
		topic.ReplyTime = time.Now()
		topic.ReplyLastUserName = nickname
		topic.ReplyCount++
		_, err = o.Update(topic)
	}
	return err
}

func GetAllWikiReplies(tid int64) (replies []*Commentwiki, err error) {
	// tidNum, err := strconv.ParseInt(tid, 10, 64)
	// if err != nil {
	// 	return nil, err
	// }
	replies = make([]*Commentwiki, 0)

	o := orm.NewOrm()
	qs := o.QueryTable("commentwiki")
	_, err = qs.Filter("tid", tid).All(&replies)
	return replies, err
}
