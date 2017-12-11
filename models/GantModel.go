package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	// "strconv"
	// "strings"
	// "fmt"
	"time"
)

// id": -1,
// "status": "STATUS_
// "level": 0,
// "code": "SL2017",
// "name": "珠三角水
// "startIsMilestone":
// "start": 139699440
// "endIsMilestone":
// "end": 13995863999
// "duration": 20,
// "progress": 80,
// "progressByWorklog
// "depends": "",
// "hasChild": true,
// "description": "17

// "relevance": 0,
// "type": "",
// "typeId": "",

// "canWrite": true,
// "collapsed": false

// "assigs":[
// 						"resourceId":"tmp_1",
//             "id":"tmp_1345625008213",
//             "roleId":"tmp_1",
//             "effort":7200000

// "resources":[
//       {"id": "tmp_1",
//        "name": "秦晓川"
//        },
//       {"id": "tmp_2", "name": "冯文涛"},
//       {"id": "tmp_3", "name": "张武"},
//       {"id": "tmp_4", "name": "陈小云"}
//     ],
//     "roles":[
//       {"id": "tmp_1",
//       "name": "项目负责人"
//       },
//       {"id": "tmp_2", "name": "专业负责人"},
//       {"id": "tmp_3", "name": "专业负责人"},
//       {"id": "tmp_4", "name": "审查"}
//     ],

type ProjGant struct {
	Id               int64  `form:"-"`
	ParentId         int64  `orm:"null"`
	Status           string `orm:"null"` //STATUS_ACTIVE, STATUS_DONE, STATUS_FAILED, STATUS_SUSPENDED, STATUS_UNDEFINED
	Level            int    `orm:"null"`
	Code             string `orm:"null"` //编号
	Name             string `orm:"null"` //项目名称
	StartIsMilestone bool
	Start            time.Time `orm:"type(datetime)"`
	EndIsMilestone   bool
	End              time.Time `orm:"null;type(datetime)"`
	Duration         int       `orm:"null"`
	Progress         int       `orm:"null"`
	Depends          string    `orm:"null"`
	HasChild         bool      `orm:"null"`
	Description      string    `orm:"null"`
	Show             bool
	Created          time.Time `orm:"null","auto_now_add;type(datetime)"`
	Updated          time.Time `orm:"null","auto_now_add;type(datetime)"`
	// ProgressByWorklog string    `orm:"null"`
}

func init() {
	orm.RegisterModel(new(ProjGant))
	// orm.RegisterModel(new(AdminIpsegment))
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/meritms.db", 10)
}

//添加项目进度
func AddProjGant(id1, parentid int64, status, code, name, depends, description string, level, duration, progress int, start, end time.Time, startismilestone, endismilestone, haschild bool) (id int64, err error) {
	o := orm.NewOrm()
	gantt := &ProjGant{Id: id1}
	err = o.Read(gantt)
	if err == orm.ErrNoRows {
		//关闭写同步
		// o.Raw("PRAGMA synchronous = OFF; ", 0, 0, 0).Exec()
		// var ProjGant ProjGant
		// if pid == "" {
		projgant := &ProjGant{
			ParentId:         parentid,
			Status:           status,
			Level:            level,
			Code:             code,
			Name:             name,
			StartIsMilestone: startismilestone,
			Start:            start,
			EndIsMilestone:   endismilestone,
			End:              end,
			Duration:         duration,
			Progress:         progress,
			Depends:          depends,
			HasChild:         haschild,
			Description:      description,
			Show:             true,
			Created:          time.Now(),
			Updated:          time.Now(),
		}
		id, err = o.Insert(projgant)
		if err != nil {
			return id, err
		}
	} else if err == orm.ErrMissPK {
		return id, err
	} else {
		return id, nil
	}
	return id, err
}

//修改——还没改
func UpdateProjGant(cid int64, code, title, label string) error {
	// o := orm.NewOrm()
	// project := &ProjGant{Id: cid}
	// if o.Read(project) == nil {
	// 	project.Code = code
	// 	project.Title = title
	// 	project.Label = label
	// 	project.Updated = time.Now()
	// 	_, err := o.Update(project)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

//删除
func DeleteProjGant(id int64) error {
	o := orm.NewOrm()
	projectgant := &ProjGant{Id: id}
	if o.Read(projectgant) == nil {
		_, err := o.Delete(projectgant)
		if err != nil {
			return err
		}
	}
	return nil
}

//关闭
func CloseProjGant(id int64) error {
	o := orm.NewOrm()
	projectgant := &ProjGant{Id: id}
	if o.Read(projectgant) == nil {
		_, err := o.Delete(projectgant)
		if err != nil {
			return err
		}
	}
	return nil
}

//取得所有项目进度，按结束时间早到晚排列。
func GetProjGants() (projgant []*ProjGant, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("ProjGant") //这个表名AchievementTopic需要用驼峰式，
	_, err = qs.Filter("show", true).OrderBy("End").All(&projgant)
	if err != nil {
		return projgant, err
	}
	return projgant, err
}

//根据id取得项目进度
func GetProjGant(id int64) (projgant ProjGant, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("ProjGant") //这个表名AchievementTopic需要用驼峰式，
	err = qs.Filter("id", id).One(&projgant)
	if err != nil {
		return projgant, err
	}
	return projgant, err
}

//根据编号code和名字name查询到项目进度
func GetProjGantName(code, name string) (projgant ProjGant, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("ProjGant")
	// var cate ProjGant
	err = qs.Filter("code", code).Filter("name", name).One(&projgant)
	// if pid != "" {
	// cate := ProjGant{Title: title}这句无效
	// categories = make([]*ProjGant, 0)
	// _, err = qs.Filter("parentid", cate.Id).All(&categories)
	// if err != nil {
	// 	return nil, err
	// }
	return projgant, err
	// } else { //如果不给定父id（PID=0），则取所有一级
	// _, err = qs.Filter("parentid", 0).All(&categories)
	// if err != nil {
	// return nil, err
	// }
	// return categories, err
	// }
}

//根据名字name和parentid查询到项目进度
func GetProjGantParent(name string, parentid int64) (projgant ProjGant, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("ProjGant")
	err = qs.Filter("name", name).Filter("parentid", parentid).One(&projgant)
	return projgant, err
}
