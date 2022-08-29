package models

import (
	"github.com/beego/beego/v2/client/orm"
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
	_, err = qs.Filter("product_id", pid).All(&Articles, "Id", "ProductId", "Created", "Updated")
	if err != nil {
		return nil, err
	}
	return Articles, err
}

//微信小程序，根据成果id取得所有文章——返回id和prodid，content……
func GetWxArticles(pid int64) (Articles []*Article, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Article")
	_, err = qs.Filter("product_id", pid).All(&Articles)
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
	// db := GetDB()
	// 这个可行db.Table("article").Select("product_id as productid, count(*) as total").Group("product_id").Scan(&results)
	// db.Table("article").Select("product_id as productid, count(*) as total,user.nickname as usernickname").Group("product_id").
	// 	Joins("left JOIN product on product.id = article.product_id").
	// 	Joins("left JOIN user on user.id = product.uid").
	// 	Scan(&results)
	err = _db.Order("total desc").Table("article").Select("product_id as productid, count(*) as total,user.nickname as usernickname").
		Joins("left JOIN product on product.id = article.product_id").
		Joins("left JOIN user on user.id = product.uid").Group("product.uid").
		Joins("left JOIN project on project.id = product.project_id").Where("project.id=?", pid).
		Scan(&results).Error
	return results, err
}

//取出某个项目id下的文章列表
// type WxArticleResult struct {
// 	Usernickname string `json:"name"`
// 	ProductId    int64
// 	Subtext      string
// 	Content      string
// 	Created      time.Time
// 	Id           int64
// 	Total        int64 `json:"value"`
// }

// 由名字模糊搜索,productid——projectid
// Select("product_id as product_id,article.id as id,article.subtext as subtext,article.content as content,article.created as created").
// 不加上面这句，就会导致Id一直是projectid，因为是join，所以id会冲突，以最后查询出的id当做ariticle这个结构体的id了。
func SearchArticles(pid int64, limit, offset int, key string, isDesc bool) (articles []*Article, err error) {
	db := _db //GetDB()
	err = db.Order("created desc").Table("article").
		Select("product_id as product_id,article.id as id,article.subtext as subtext,article.content as content,article.created as created").
		Joins("left JOIN product on product.id = article.product_id").
		Joins("left JOIN project on project.id = product.project_id").
		Where("project.id=? AND article.subtext like ?", pid, "%"+key+"%").
		Limit(limit).Offset(offset).Scan(&articles).Error
	return articles, err

	// projects := db.Where("name like ?", "%"+this.Name+"%").Find(&result)

	// cond := orm.NewCondition()
	// cond1 := cond.Or("Subtext__contains", key).Or("Content__contains", key)
	// o := orm.NewOrm()
	// qs := o.QueryTable("Article")
	// qs = qs.SetCond(cond1)
	// if isDesc {
	// 	_, err = qs.Distinct().Limit(limit, offset).OrderBy("-created").All(&articles)
	// 	if err != nil {
	// 		return articles, err
	// 	}
	// } else {
	// 	_, err = qs.Distinct().Limit(limit, offset).All(&articles)
	// 	if err != nil {
	// 		return articles, err
	// 	}
	// }
	// return articles, err
}
