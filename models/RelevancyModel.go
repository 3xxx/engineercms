package models

import (
	"github.com/astaxie/beego/orm"
	// _ "github.com/mattn/go-sqlite3"
	// "strconv"
	// "fmt"
	// "strings"
	"time"
)

type Relevancy struct {
	Id        int64     `form:"-"`
	ProductId int64     `orm:"null"` //编号
	Relevancy string    `orm:"null"` //标签
	Created   time.Time `orm:"null","auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"null","auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Relevancy)) //, new(Article)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/engineer.db", 10)
}

//添加项目
func AddRelevancy(prodid int64, relevancy string) (id int64, err error) {
	o := orm.NewOrm()
	//关闭写同步
	o.Raw("PRAGMA synchronous = OFF; ", 0, 0, 0).Exec()
	// var project Project
	// if pid == "" {
	relev := &Relevancy{
		ProductId: prodid,
		Relevancy: relevancy,
		Created:   time.Now(),
		Updated:   time.Now(),
	}
	id, err = o.Insert(relev)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//根据成果id取得关联文件
//根据成果编号，取得关联文件编号
func GetRelevancy(prodid int64) (relevancies []*Relevancy, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Relevancy")
	// relevancies1 := make([]*Relevancy, 0)
	_, err = qs.Filter("ProductId", prodid).All(&relevancies)
	if err != nil {
		return nil, err
	}
	return relevancies, err
}

func GetAllRelevancies() (relevancies []*Relevancy, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Relevancy")
	// relevancies1 := make([]*Relevancy, 0)
	_, err = qs.All(&relevancies)
	if err != nil {
		return nil, err
	}
	return relevancies, err
}
