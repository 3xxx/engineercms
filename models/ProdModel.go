package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// _ "github.com/mattn/go-sqlite3"
	// "strconv"
	// "strings"
	"time"
)

type Product struct {
	Id           int64
	Code         string `orm:"null"`                                              //编号                                          //编号                                             //编号
	Title        string `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Label        string `orm:"null"`                                              //关键字                                           //标签
	Uid          int64  `orm:"null"`
	Principal    string `orm:"null"`       //提供人                                            //负责人id
	ProjectId    int64  `orm:"null"`       //侧栏id
	TopProjectId int64  `orm:"default(0)"` //项目id
	// Content      string    `orm:"sie(5000)"` //内容
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
	// Views        int64
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

//修改成果时间信息
func UpdateProductTime(cid int64) error {
	o := orm.NewOrm()
	product := &Product{Id: cid}
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
func AddProduct(code, title, label, principal string, uid, Projectid, TopProjectId int64) (id int64, err error) {
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
			Code:         code,
			Title:        title,
			Label:        label,
			Uid:          uid,
			Principal:    principal,
			ProjectId:    Projectid,
			TopProjectId: TopProjectId,
			// Content:      content,
			Created: time.Now(),
			Updated: time.Now(),
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

//根据侧栏id查出所有成果——按编号排序
func GetProducts(id int64) (products []*Product, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Product")
	_, err = qs.Filter("ProjectId", id).OrderBy("-code").Limit(-1).All(&products)
	if err != nil {
		return nil, err
	}
	return products, err
}

//根据侧栏id分页查出所有成果——按编号排序
func GetProductsPage(id, limit, offset, uid int64, searchText string) (products []*Product, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Product")
	if searchText != "" {
		cond := orm.NewCondition()
		cond1 := cond.Or("Code__contains", searchText).Or("Title__contains", searchText).Or("Label__contains", searchText).Or("Principal__contains", searchText)
		cond2 := cond.AndCond(cond1).And("ProjectId", id)
		qs = qs.SetCond(cond2)
		_, err = qs.Limit(limit, offset).OrderBy("-created").All(&products)
	} else if uid == 0 {
		_, err = qs.Filter("ProjectId", id).Limit(limit, offset).OrderBy("-created").All(&products)
	} else if uid != 0 {
		_, err = qs.Filter("ProjectId", id).Filter("Uid", uid).Limit(limit, offset).OrderBy("-created").All(&products)
	}
	return products, err
}

//取得侧栏id下成果总数
func GetProductsCount(id int64, searchText string) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Product")
	if searchText != "" {
		cond := orm.NewCondition()
		cond1 := cond.Or("Code__contains", searchText).Or("Title__contains", searchText).Or("Label__contains", searchText).Or("Principal__contains", searchText)
		cond2 := cond.AndCond(cond1).And("ProjectId", id)
		qs = qs.SetCond(cond2)
		count, err = qs.Limit(-1).Count()
	} else {
		count, err = qs.Filter("ProjectId", id).Count()
	}
	return count, err
}

//根据项目顶级id查出所有成果
//这个不优雅，应该用循环子孙id，查询字段在id数组中，参考SearchProjProductPage
//直接把所有成果都查出来。getallproduct
func GetProjProducts(id int64, number int) (count int64, products []*Product, err error) {
	// idstring := strconv.FormatInt(id, 10)
	// projects := make([]*Project, 0)
	// cond := orm.NewCondition()
	// cond1 := cond.Or("Id", id).Or("ParentIdPath__contains", idstring+"-").Or("ParentId", id)
	o := orm.NewOrm()
	//先查出所有项目parent id path中包含id的数据
	// qs := o.QueryTable("Project")
	// qs = qs.SetCond(cond1)
	// _, err = qs.Filter("ParentIdPath__contains", "$"+idstring+"#").Limit(-1).All(&projects)
	// if err != nil {
	// 	return nil, err
	// }
	// _, err = qs.Limit(-1).All(&projects)
	// if err != nil {
	// 	return nil, err
	// }
	//循环数据的id，查出成果product
	qs1 := o.QueryTable("Product")
	// products1 := make([]*Product, 0)
	// for _, v := range projects {
	// _, err = qs1.Filter("ProjectId", v.Id).OrderBy("-created").Limit(-1).All(&products1) //, "ProjectId"
	if number == 1 { //查出所有成果列表
		_, err = qs1.Filter("TopProjectId", id).OrderBy("-created").Limit(-1).All(&products) //,
	} else if number == 2 { //查出所有成果属于目录的关系
		_, err = qs1.Filter("TopProjectId", id).OrderBy("-created").Limit(-1).All(&products, "Id", "ProjectId") //,
	} else if number == 3 {
		count, err = qs1.Filter("TopProjectId", id).Limit(-1).Count()
	}
	if err != nil {
		return 0, nil, err
	}
	// 	products = append(products, products1...)
	// 	products1 = make([]*Product, 0)
	// }
	return count, products, err
}

//直接把所有成果都查出来。——全表，避免使用
func GetAllProducts() (products []*Product, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Product")
	_, err = qs.Limit(-1).All(&products, "Id", "ProjectId")
	if err != nil {
		return nil, err
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
