package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	// "strconv"
	// "strings"
	"time"
)

type Article struct {
	Id        int64     `form:"-"`
	Subtext   string    `orm:"sie(20)"`
	Content   string    `orm:"sie(5000)"`
	ProductId int64     `orm:"null"`
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Article)) //, new(Article)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/engineer.db", 10)
}

//添加文章作为成果的附件
func AddArticle(subtext, content string, productid int64) (id int64, err error) {
	o := orm.NewOrm()
	Article := &Article{
		Subtext:   subtext,
		Content:   content,
		ProductId: productid,
		Created:   time.Now(),
		Updated:   time.Now(),
	}
	id, err = o.Insert(Article)
	if err != nil {
		return 0, err
	}
	return id, nil
}

//修改
func UpdateArticle(id int64, subtext, content string) error {
	o := orm.NewOrm()
	article := &Article{Id: id}
	if o.Read(article) == nil {
		article.Subtext = subtext
		article.Content = content
		article.Updated = time.Now()
		_, err := o.Update(article)
		if err != nil {
			return err
		}
	}
	return nil
}

//删除
func DeleteArticle(id int64) error {
	o := orm.NewOrm()
	article := &Article{Id: id}
	if o.Read(article) == nil {
		_, err := o.Delete(article)
		if err != nil {
			return err
		}
	}
	return nil
}

//取得所有项目
// func GetArticles() (Artic []*Article, err error) {
// 	o := orm.NewOrm()
// 	qs := o.QueryTable("Article") //这个表名AchievementTopic需要用驼峰式，
// 	_, err = qs.Filter("parentid", 0).All(&Artic)
// 	if err != nil {
// 		return Artic, err
// 	}
// 	return Artic, err
// }
//根据成果id取得所有文章
func GetArticles(pid int64) (Articles []*Article, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Article")
	_, err = qs.Filter("Productid", pid).All(&Articles)
	if err != nil {
		return nil, err
	}
	return Articles, err
}

//根据文章id取得文章
func GetArticle(id int64) (Artic Article, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Article") //这个表名AchievementTopic需要用驼峰式，
	err = qs.Filter("id", id).One(&Artic)
	if err != nil {
		return Artic, err
	}
	return Artic, err
}
