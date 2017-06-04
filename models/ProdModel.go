package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	// "strings"
	"time"
)

type Product struct {
	Id        int64
	Code      string    `orm:"null"`                                              //编号                                          //编号                                             //编号
	Title     string    `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Label     string    `orm:"null"`                                              //关键字                                           //标签
	Uid       int64     `orm:"null"`
	Principal string    `orm:"null"`      //提供人                                            //负责人id
	ProjectId int64     `orm:"null"`      //侧栏id
	Content   string    `orm:"sie(5000)"` //内容
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now;type(datetime)"`
	Views     int64
	// ReplyTime         time.Time
	// ReplyCount        int64
	// ReplyLastUserName string
	// Attachments     []*Attachment `orm:"reverse(many)"` // fk 的反向关系
}

func init() {
	orm.RegisterModel(new(Product)) //, new(Article)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/engineer.db", 10)
}

//修改成果信息
func UpdateProduct(cid int64, code, title, label, principal string) error {
	o := orm.NewOrm()
	product := &Product{Id: cid}
	if o.Read(product) == nil {
		product.Code = code
		product.Title = title
		product.Label = label
		product.Principal = principal
		product.Updated = time.Now()
		_, err := o.Update(product)
		if err != nil {
			return err
		}
	}
	return nil
}

//删除_把附件也一并删除（在controllers中实现吧）
func DeleteProduct(cid int64) error {
	o := orm.NewOrm()
	product := &Product{Id: cid}
	if o.Read(product) == nil {
		_, err := o.Delete(product)
		if err != nil {
			return err
		}
	}
	return nil
}

//添加成果到项目侧栏某个id下
//如果这个侧栏id下的这个成果编号已经存在，则返回id
////应该用ReadOrCreate尝试从数据库读取，不存在的话就创建一个
func AddProduct(code, title, label, principal, content string, Projectid int64) (id int64, err error) {
	o := orm.NewOrm()
	// err := o.QueryTable("user").Filter("name", "slene").One(&user)
	// if err == orm.ErrMultiRows {
	// 	// 多条的时候报错——单条呢？没有错误，错误为nil
	// 	fmt.Printf("Returned Multi Rows Not One")
	// }
	// if err == orm.ErrNoRows {
	// 	// 没有找到记录
	// 	fmt.Printf("Not row found")
	// }
	var prod Product
	err = o.QueryTable("Product").Filter("Projectid", Projectid).Filter("code", code).One(&prod)
	if err == orm.ErrNoRows { // 没有找到记录
		product := &Product{
			Code:      code,
			Title:     title,
			Label:     label,
			Principal: principal,
			ProjectId: Projectid,
			Content:   content,
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
	// var Product Product
	// if pid == "" {

	// } else {

	// }
	return id, err
}

//根据侧栏id查出所有成果
func GetProducts(id int64) (products []*Product, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Product")
	_, err = qs.Filter("ProjectId", id).OrderBy("-created").All(&products)
	if err != nil {
		return nil, err
	}
	return products, err
}

//根据项目id查出所有成果
func GetProjProducts(id int64) (products []*Product, err error) {
	idstring := strconv.FormatInt(id, 10)
	projects := make([]*Project, 0)

	cond := orm.NewCondition()
	cond1 := cond.Or("Id", id).Or("ParentIdPath__contains", idstring)
	o := orm.NewOrm()
	//先查出所有项目parent id path中包含id的数据
	qs := o.QueryTable("Project")
	qs = qs.SetCond(cond1)

	_, err = qs.All(&projects)
	if err != nil {
		return nil, err
	}
	//循环数据的id，查出成果product
	qs1 := o.QueryTable("Product")
	products1 := make([]*Product, 0)
	for _, v := range projects {
		_, err = qs1.Filter("ProjectId", v.Id).OrderBy("-created").All(&products1)
		if err != nil {
			return nil, err
		}
		products = append(products, products1...)
		products1 = make([]*Product, 0)
	}
	return products, err
}

//根据成果id取得成果
func GetProd(id int64) (prod Product, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Product") //这个表名AchievementTopic需要用驼峰式，
	err = qs.Filter("id", id).One(&prod)
	if err != nil {
		return prod, err
	}
	return prod, err
}
