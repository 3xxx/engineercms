package models

import (
	// "errors"
	// "strconv"
	// "fmt"
	// "log"
	// "time"
	// beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
	// "github.com/beego/beego/v2/adapter/validation"
	// . "github.com/beego/admin/src/lib"
)

// 搜索本地项目
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

// 搜索本地成果
func SearchProduct(key, searchtext string) (prod []*Product, err error) {
	cond := orm.NewCondition()
	cond1 := cond.Or("Code__contains", key).Or("Title__contains", key).Or("Label__contains", key).Or("Principal__contains", key)
	var cond4 *orm.Condition
	if searchtext != "" {
		cond3 := cond.Or("Code__contains", searchtext).Or("Title__contains", searchtext).Or("Label__contains", searchtext).Or("Principal__contains", searchtext)
		cond4 = cond.AndCond(cond1).AndCond(cond3)
	} else {
		cond4 = cond1
	}
	o := orm.NewOrm()
	qs := o.QueryTable("Product")
	qs = qs.SetCond(cond4)
	_, err = qs.Distinct().OrderBy("-created").All(&prod) //qs.Filter("Drawn", user.Nickname).All(&aa)
	if err != nil {
		return prod, err
	}
	return prod, err
}

// 搜索本地成果,分页
func SearchProductPage(limit, offset int64, key, searchtext string) (prod []*Product, err error) {
	cond := orm.NewCondition()
	cond1 := cond.Or("Code__contains", key).Or("Title__contains", key).Or("Label__contains", key).Or("Principal__contains", key)
	var cond4 *orm.Condition
	if searchtext != "" {
		cond3 := cond.Or("Code__contains", searchtext).Or("Title__contains", searchtext).Or("Label__contains", searchtext).Or("Principal__contains", searchtext)
		cond4 = cond.AndCond(cond1).AndCond(cond3)
	} else {
		cond4 = cond1
	}
	o := orm.NewOrm()
	qs := o.QueryTable("Product")
	qs = qs.SetCond(cond4)
	_, err = qs.Distinct().Limit(limit, offset).OrderBy("-created").All(&prod) //qs.Filter("Drawn", user.Nickname).All(&aa)
	if err != nil {
		return prod, err
	}
	return prod, err
	// db := GetDB()
	// if isadmin {
	// 	err = db.Order("cart.updated desc").Table("cart").Select("cart.id,cart.user_id,cart.product_id,cart.status,cart.updated,user.nickname as user_nickname, product.title as product_title, product.top_project_id as top_project_id,t1.title as project_title,t2.title as top_project_title").
	// 		Where("cart.status=?", status).
	// 		Joins("left JOIN user on user.id = cart.user_id").
	// 		Joins("left join product on product.id = cart.product_id").
	// 		Joins("left join project AS t1 on t1.id = product.project_id").
	// 		Joins("left join project AS t2 ON t2.id = product.top_project_id").
	// 		Limit(limit).Offset(offset).Scan(&usercarts).Error
	// }
}

// 搜索本地成果,分页
func SearchProductCount(key, searchtext string) (count int64, err error) {
	cond := orm.NewCondition()
	cond1 := cond.Or("Code__contains", key).Or("Title__contains", key).Or("Label__contains", key).Or("Principal__contains", key)
	var cond4 *orm.Condition
	if searchtext != "" {
		cond3 := cond.Or("Code__contains", searchtext).Or("Title__contains", searchtext).Or("Label__contains", searchtext).Or("Principal__contains", searchtext)
		cond4 = cond.AndCond(cond1).AndCond(cond3)
	} else {
		cond4 = cond1
	}
	o := orm.NewOrm()
	qs := o.QueryTable("Product")
	qs = qs.SetCond(cond4)
	count, err = qs.Distinct().Count()
	if err != nil {
		return count, err
	}
	return count, err
	// db := GetDB()
	// if isadmin {
	// 	err = db.Order("cart.updated desc").Table("cart").Select("cart.id,cart.user_id,cart.product_id,cart.status,cart.updated,user.nickname as user_nickname, product.title as product_title, product.top_project_id as top_project_id,t1.title as project_title,t2.title as top_project_title").
	// 		Where("cart.status=?", status).
	// 		Joins("left JOIN user on user.id = cart.user_id").
	// 		Joins("left join product on product.id = cart.product_id").
	// 		Joins("left join project AS t1 on t1.id = product.project_id").
	// 		Joins("left join project AS t2 ON t2.id = product.top_project_id").
	// 		Limit(limit).Offset(offset).Scan(&usercarts).Error
	// }
}

// 搜索某个项目里的成果：article的全文，待完善
func SearchProjProduct(pid, limit, offset int64, key, searchtext string) (count int64, prod []*Product, err error) {
	cond := orm.NewCondition()
	cond1 := cond.Or("ProjectId", pid)
	//查出所有子孙项目
	//取到所有子孙pid
	sonproj, _ := GetProjectsbyPid(pid)
	for _, v := range sonproj {
		cond1 = cond1.Or("ProjectId", v.Id)
	}
	cond2 := cond.Or("Code__contains", key).Or("Title__contains", key).Or("Label__contains", key).Or("Principal__contains", key)
	//(...or...or...)and()
	var cond4 *orm.Condition
	if searchtext != "" {
		cond3 := cond.Or("Code__contains", searchtext).Or("Title__contains", searchtext).Or("Label__contains", searchtext).Or("Principal__contains", searchtext)
		cond4 = cond.AndCond(cond1).AndCond(cond2).AndCond(cond3)
	} else {
		cond4 = cond.AndCond(cond1).AndCond(cond2)
	}
	o := orm.NewOrm()
	qs := o.QueryTable("Product")
	qs = qs.SetCond(cond4)
	// cond := orm.NewCondition()
	// cond1 := cond.Or("Code__contains", key).Or("Title__contains", key).Or("Label__contains", key).Or("Principal__contains", key)
	// cond2 := cond.AndCond(cond1).And("Content__contains", key)
	// o := orm.NewOrm()
	// qs := o.QueryTable("Product")
	// qs1 := qs.SetCond(cond2)
	_, err = qs.Limit(limit, offset).Distinct().OrderBy("-created").All(&prod)
	// _, err = qs.Filter("ProjectId", pid).Distinct().OrderBy("-created").All(&prod)
	if err != nil {
		return count, prod, err
	}
	count, err = qs.Distinct().Count()
	if err != nil {
		return count, prod, err
	}
	//取出所有成果
	// articls := make([]*Article, 0)
	// _, products, err := GetProjProducts(pid, 1)
	// qs2 := o.QueryTable("Article")
	// for _, v := range prod {
	// 	_, err = qs2.Filter("ProductId", v.Id).Filter("Content__contains", key).OrderBy("-created").All(&articls)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if len(articls) > 0 {
	// 		prod = append(prod, v)
	// 	}
	// 	articls = make([]*Article, 0)
	// }
	return count, prod, err
}

// 搜索某个项目里的成果，分页，
// 递归出所有子孙项目，这个厉害
// 这个条件不行，id1或id2或id3 and code 或key 或
func SearchProjProductPage(pid, limit, offset int64, key, searchtext string) (prod []*Product, err error) {
	cond := orm.NewCondition()
	cond1 := cond.Or("ProjectId", pid)
	//查出所有子孙项目
	//取到所有子孙pid
	sonproj, _ := GetProjectsbyPid(pid)

	for _, v := range sonproj {
		cond1 = cond1.Or("ProjectId", v.Id)
	}
	var cond4 *orm.Condition
	if key != "" {
		cond2 := cond.Or("Code__contains", key).Or("Title__contains", key).Or("Label__contains", key).Or("Principal__contains", key)
		// cond1 = cond1.AndCond(cond2)// 这里会出问题
		// cond1 = cond.AndCond(cond1).AndCond(cond2) // 这里是正确的写法

		if searchtext != "" {
			cond3 := cond.Or("Code__contains", searchtext).Or("Title__contains", searchtext).Or("Label__contains", searchtext).Or("Principal__contains", searchtext)
			cond4 = cond.AndCond(cond1).AndCond(cond2).AndCond(cond3)
		} else {
			cond4 = cond.AndCond(cond1).AndCond(cond2)
		}
		// cond2 := cond.AndCond(cond1).OrCond(cond.And("name", "slene"))
		//循环
		// cond2 := cond.Or("ProjectId", pid1).Or("ProjectId", pid2)……
		// ids := []int{1, 2, 3}
		// p.Raw("SELECT name FROM user WHERE id IN (?, ?, ?)", ids)

		// 假定表名test,列id是数值类型。
		// 用同一个字段的多个值作为条件来查询可以使用in或者or
		// select * from test where id in (1,2,3)

		// SELECT * FROM `employee`' AND ' : ' WHERE '
		// SELECT id, user_name FROM user WHERE id = ?
		// elect * from tbl_employee where id=#{id} and last_name like #{lastName} and email=#{email}
	}
	//(...or...or...)and()
	o := orm.NewOrm()
	qs := o.QueryTable("Product")
	qs = qs.SetCond(cond4)
	//循环这个id下所有项目？
	_, err = qs.Limit(limit, offset).Distinct().OrderBy("-created").All(&prod) //qs.Filter("Drawn", user.Nickname).All(&aa)
	if err != nil {
		return prod, err
	}
	return prod, err
}

// String name = request.getParameter("name");  //姓名
// String rank= request.getParameter("age");  //年龄
// String address= request.getParameter("address");  //地址
// String sql = "select * from  student where 1=1 ";
// if(name!=null && !name.equals("")){
//     sql += "t.name like '%"+name+"%'";
// }
// if(rank!=null && !rank.equals("")){
//     sql += "t.age like '%"+age+"%'";
// }
// if(address!=null && !address.equals("")){
//     sql += "t.address like '%"+address+"%'";
// }

//定义一个存储实际参数的容器
// List<String>list=new ArrayList<String>();
// Stringsql = "select * from product where 1=1";
// if(condition.getPname()!=null&&condition.getPname().trim().equals("")){
//        sql+="and pname  like ? ";
//        list.add("%"+condition.getPname().trim()+"%");
// }
// if(condition.getIsHot()!=null&&condition.getIsHot().trim().equals("")){
// sql+="and is_hot=?  ";
// list.add(condition.getIsHot().trim());
// }
// if(condition.getCid()!=null&&condition.getCid().trim().equals("")){
//        sql+="and cid=? ";
//        list.add("%"+condition.getCid().trim()+"%");
//        }
// List<Product>productList = runner.query(sql,new BeanListHandler<Product>(Product.class),list.toArray());
// return  productList;

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
