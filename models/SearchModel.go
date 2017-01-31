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
