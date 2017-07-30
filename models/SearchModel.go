package models

import (
	// "errors"
	// "strconv"
	// "fmt"
	// "log"
	// "time"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
	// . "github.com/beego/admin/src/lib"
)

//搜索本地项目
func SearchProject(key string) (proj []*Project, err error) {
	cond := orm.NewCondition()
	cond1 := cond.Or("Code__contains", key).Or("Title__contains", key).Or("Label__contains", key).Or("Principal__contains", key)

	o := orm.NewOrm()
	qs := o.QueryTable("Project")
	qs = qs.SetCond(cond1)
	_, err = qs.Distinct().OrderBy("-created").All(&proj) //qs.Filter("Drawn", user.Nickname).All(&aa)
	if err != nil {
		return proj, err
	}
	return proj, err
}

//搜索本地成果
func SearchProduct(key string) (prod []*Product, err error) {
	cond := orm.NewCondition()
	cond1 := cond.Or("Code__contains", key).Or("Title__contains", key).Or("Label__contains", key).Or("Principal__contains", key).Or("Content__contains", key)
	o := orm.NewOrm()
	qs := o.QueryTable("Product")
	qs = qs.SetCond(cond1)
	_, err = qs.Distinct().OrderBy("-created").All(&prod) //qs.Filter("Drawn", user.Nickname).All(&aa)
	if err != nil {
		return prod, err
	}
	return prod, err
}

//搜索某个项目里的成果：article的全文，待完善
func SearchProjProduct(pid int64, key string) (prod []*Product, err error) {
	cond := orm.NewCondition()
	cond1 := cond.Or("Code__contains", key).Or("Title__contains", key).Or("Label__contains", key).Or("Principal__contains", key)
	// cond2 := cond.Or("Content__contains", key)

	o := orm.NewOrm()
	qs := o.QueryTable("Product")
	qs1 := qs.SetCond(cond1)
	_, err = qs1.Filter("ProjectId", pid).Distinct().OrderBy("-created").All(&prod) //qs.Filter("Drawn", user.Nickname).All(&aa)
	if err != nil {
		return prod, err
	}
	//取出所有成果
	articls := make([]*Article, 0)
	products, err := GetProjProducts(pid)

	qs2 := o.QueryTable("Article")
	// qs3 := qs2.SetCond(cond2)

	for _, v := range products {
		_, err = qs2.Filter("ProductId", v.Id).Filter("Content__contains", key).OrderBy("-created").All(&articls)
		if err != nil {
			return nil, err
		}
		if len(articls) > 0 {
			prod = append(prod, v)
		}
		articls = make([]*Article, 0)
	}

	return prod, err
}

//设计院首页全局搜索——未修改
// func Searchspidertopics(title string, isDesc bool) ([]*Spidertopic, []*Spidercategory, error) {
// 	o := orm.NewOrm()
// 	spidertopics := make([]*Spidertopic, 0)
// 	spidercategories := make([]*Spidercategory, 0)
// 	// spidercategories := make([]*Spidercategory, 0)
// 	qs := o.QueryTable("spidertopic")
// 	var err error
// 	if isDesc {
// 		if len(title) > 0 {
// 			qs = qs.Filter("Name__contains", title) //这里取回
// 		}
// 		_, err = qs.OrderBy("-created").All(&spidertopics)
// 	} else {
// 		_, err = qs.Filter("Name__contains", title).OrderBy("-created").All(&spidertopics)
// 		//o.QueryTable("user").Filter("name", "slene").All(&users)
// 	}
// 	qs1 := o.QueryTable("spidercategory")
// 	if isDesc {
// 		if len(title) > 0 {
// 			qs1 = qs1.Filter("Name__contains", title) //这里取回
// 		}
// 		_, err = qs1.OrderBy("-created").All(&spidercategories)
// 	} else {
// 		_, err = qs1.Filter("Name__contains", title).OrderBy("-created").All(&spidercategories)
// 		//o.QueryTable("user").Filter("name", "slene").All(&users)
// 	}

// 	return spidertopics, spidercategories, err
// }
