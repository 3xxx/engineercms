package models

import (
	// "crypto/md5"
	// "encoding/hex"
	// "errors"
	"strconv"
	// "fmt"
	// "log"
	// "time"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
	// . "github.com/beego/admin/src/lib"
	"strings"
)

func UpdateDatabase() (err1, err2, err3, err4 error) {
	o := orm.NewOrm()
	o.Raw("PRAGMA synchronous = OFF; ", 0, 0, 0).Exec()

	//取得所有
	qs := o.QueryTable("Project")
	var proj []*Project
	_, err := qs.Limit(-1).All(&proj)
	//1.替换parentidpath
	for _, v := range proj {
		if !strings.Contains(v.ParentIdPath, "$") && v.ParentIdPath != "" {
			category := &Project{Id: v.Id}
			//patharray := strings.Split(parentidpath1, "-")
			patharray := "$" + strings.Replace(v.ParentIdPath, "-", "#$", -1) + "#"
			category.ParentIdPath = patharray
			_, err := o.Update(category, "ParentIdPath")
			if err != nil {
				return err
			}
		}
	}

	//2.product表添加projid字段
	var topprojectid int64
	qs1 := o.QueryTable("Product")
	var prod []*Product
	var proj Project
	_, err1 = qs1.Limit(-1).All(&prod)
	for _, v1 := range prod {
		//根据pid查出目录id
		proj, err2 = GetProj(v1.ProjectId)
		// if err != nil {
		// 	return err2
		// }
		if v1.TopProjectId == 0 {
			if proj.ParentIdPath != "" { //如果不是根目录
				parentidpath := strings.Replace(strings.Replace(proj.ParentIdPath, "#$", "-", -1), "$", "", -1)
				parentidpath1 := strings.Replace(parentidpath, "#", "", -1)
				patharray := strings.Split(parentidpath1, "-")
				topprojectid, err3 = strconv.ParseInt(patharray[0], 10, 64)
				// if err != nil {
				// 	return err
				// }
			} else {
				topprojectid = proj.Id
			}
			product := &Product{Id: v1.Id}
			product.TopProjectId = topprojectid
			_, err4 = o.Update(product, "TopProjectId")
			// if err != nil {
			// 	return err
			// }
		}
	}
	return err1, err2, err3, err4
}
