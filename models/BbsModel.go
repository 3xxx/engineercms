package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

// var engine *xorm.Engine

//公告记录
type Bbs struct {
	Id         int64
	UserId     int64
	BbsDate    time.Time `orm:"auto_now_add;type(date)"`
	Desc       string
	SelectDate time.Time `orm:"auto_now_add;type(date)"`
}

func init() {
	orm.RegisterModel(new(Bbs))
}

//打卡记录写入数据库
func BbsBbs(UserId int64, Desc string, SelectDate time.Time) (id int64, err error) {
	o := orm.NewOrm()
	//查询数据库中有无打卡
	var bbs1 Bbs
	//判断是否有重名
	err = o.QueryTable("bbs").Filter("SelectDate", SelectDate).One(&bbs1, "Id")
	if err == orm.ErrNoRows {
		// 没有找到记录
		bbs := &Bbs{
			UserId:     UserId,
			BbsDate:    time.Now(),
			Desc:       Desc,
			SelectDate: SelectDate,
		}
		id, err = o.Insert(bbs)
		if err != nil {
			return id, err
		}
	} else if err == nil { //如果存在记录，或多于一条，则进行更新
		bbs := &Bbs{Id: bbs1.Id}
		bbs.BbsDate = time.Now()
		bbs.Desc = Desc
		_, err = o.Update(bbs, "BbsDate", "Desc")
		if err != nil {
			return id, err
		}
	}
	return id, nil
}

// 按月查询打卡记录
func BbsGetBbs(SelectMonth1, SelectMonth2 time.Time) (bbs []*Bbs, err error) {
	cond := orm.NewCondition()
	cond1 := cond.And("SelectDate__gte", SelectMonth1).And("SelectDate__lte", SelectMonth2)
	o := orm.NewOrm()
	qs := o.QueryTable("bbs")
	qs = qs.SetCond(cond1)
	_, err = qs.All(&bbs)
	if err != nil {
		return nil, err
	}
	return bbs, nil
}

// 查询选定日期记录
func GetBbs(SelectDate time.Time) (bbs Bbs, err error) {
	o := orm.NewOrm()
	//判断是否有重名
	err = o.QueryTable("bbs").Filter("SelectDate", SelectDate).One(&bbs)
	if err != nil {
		return bbs, err
	} else {
		return bbs, err
	}
}
