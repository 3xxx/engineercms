package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// _ "github.com/mattn/go-sqlite3"
	// "strconv"
	// "strings"
	"time"
)

type Article struct {
	Id        int64     `json:"id",form:"-"`
	Subtext   string    `orm:"size(20)"`
	Content   string    `json:"html",orm:"size(5000)"`
	ProductId int64     `orm:"null"`
	Views     int64     `orm:"default(0)"`
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
//根据成果id取得所有文章——只返回id和prodid，因为返回content太慢了，没必要吧20171007
func GetArticles(pid int64) (Articles []*Article, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Article")
	_, err = qs.Filter("Productid", pid).All(&Articles, "Id", "ProductId", "Created", "Updated")
	if err != nil {
		return nil, err
	}
	return Articles, err
}

//微信小程序，根据成果id取得所有文章——返回id和prodid，content……
func GetWxArticles(pid int64) (Articles []*Article, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Article")
	_, err = qs.Filter("Productid", pid).All(&Articles)
	if err != nil {
		return nil, err
	}
	return Articles, err
}

//根据文章id取得文章
func GetArticle(id int64) (Artic *Article, err error) {
	o := orm.NewOrm()
	article := new(Article)
	qs := o.QueryTable("Article") //这个表名AchievementTopic需要用驼峰式，
	err = qs.Filter("id", id).One(article)
	if err != nil {
		return Artic, err
	}

	article.Views++
	_, err = o.Update(article)
	if err != nil {
		return article, err
	}

	return article, err
}

//取出用户文章数目
type Result struct {
	Usernickname string `json:"name"`
	Productid    int64
	Total        int64 `json:"value"`
}

func GetWxUserArticles(pid int64) (results []*Result, err error) {
	db := GetDB()
	// 这个可行db.Table("article").Select("product_id as productid, count(*) as total").Group("product_id").Scan(&results)
	// db.Table("article").Select("product_id as productid, count(*) as total,user.nickname as usernickname").Group("product_id").
	// 	Joins("left JOIN product on product.id = article.product_id").
	// 	Joins("left JOIN user on user.id = product.uid").
	// 	Scan(&results)
	err = db.Order("total desc").Table("article").Select("product_id as productid, count(*) as total,user.nickname as usernickname").
		Joins("left JOIN product on product.id = article.product_id").
		Joins("left JOIN user on user.id = product.uid").Group("product.uid").
		Joins("left JOIN project on project.id = product.project_id").Where("project.id=?", pid).
		Scan(&results).Error
	return results, err
}
