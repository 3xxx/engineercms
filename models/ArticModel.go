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
	Content   string    `orm:"sie(5000)"` //标签
	ProductId int64     `orm:"null"`
	Created   time.Time `orm:"index","auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"index","auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Article)) //, new(Article)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/engineer.db", 10)
}

//添加项目
func AddArticle(code, title, label, principal string, parentid int64, parentidpath string, grade int) (id int64, err error) {
	o := orm.NewOrm()
	// var Article Article
	// if pid == "" {
	Article := &Article{
		// Code:         code,
		// Title:        title,
		// Label:        label,
		// Principal:    principal,
		// ParentId:     parentid,
		// ParentIdPath: parentidpath,
		// Grade:        grade,
		Created: time.Now(),
		Updated: time.Now(),
	}
	id, err = o.Insert(Article)
	if err != nil {
		return 0, err
	}
	// } else {

	// }
	return id, nil
}

//修改——还没改
func UpdateArticle(cid int64, title, code string, grade int) error {
	o := orm.NewOrm()
	//id转成64为
	// cidNum, err := strconv.ParseInt(cid, 10, 64)
	// if err != nil {
	// 	return err
	// }
	Article := &Article{Id: cid}
	if o.Read(Article) == nil {
		// Article.Title = title
		// Article.Code = code
		// Article.Grade = grade
		Article.Updated = time.Now()
		_, err := o.Update(Article)
		if err != nil {
			return err
		}
	}
	return nil
}

//取得所有项目
func GetArticles() (Artic []*Article, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Article") //这个表名AchievementTopic需要用驼峰式，
	_, err = qs.Filter("parentid", 0).All(&Artic)
	if err != nil {
		return Artic, err
	}
	return Artic, err
}

//根据id取得项目目录
func GetArtic(id int64) (Artic Article, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Article") //这个表名AchievementTopic需要用驼峰式，
	err = qs.Filter("id", id).One(&Artic)
	if err != nil {
		return Artic, err
	}
	return Artic, err
}

//根据id查出所有子孙——这个不对，要用ParentIdPath
func GetArticlesbyPid(id int64) (Articles []*Article, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Article")                         //这个表名AchievementTopic需要用驼峰式，
	_, err = qs.Filter("ParentIdPath", id).All(&Articles) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	return Articles, err
}

//根据id查出所有儿子
func GetArticSonbyId(id int64) (Articles []*Article, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Article")                     //这个表名AchievementTopic需要用驼峰式，
	_, err = qs.Filter("parentid", id).All(&Articles) //而这个字段parentid为何又不用呢
	if err != nil {
		return nil, err
	}
	return Articles, err
}

//根据id查是否有下级
func Artichasson(id int64) bool {
	o := orm.NewOrm()
	// qs := o.QueryTable("Article")
	Artic := Article{ProductId: id}
	err := o.Read(&Artic, "ProductId")
	if err == orm.ErrNoRows {
		return false
	} else if err == orm.ErrMissPK {
		return false
	} else {
		return true
	}
}

//根据名字title查询到项目目录
func GetArticleTitle(title string) (cate Article, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Article")
	// var cate Article
	err = qs.Filter("title", title).One(&cate)
	// if pid != "" {
	// cate := Article{Title: title}这句无效
	// categories = make([]*Article, 0)
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
