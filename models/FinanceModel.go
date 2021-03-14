package models

import (
	// "database/sql"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	// "github.com/go-xorm/xorm"

	// "strconv"
	// "strings"
	// "fmt"
	// "log"
	// "os"
	"time"
)

type Finance struct {
	Id int64 `json:"id",form:"-"`
	// Title       string    `orm:"sie(20)"`
	Financedate string    `orm:"sie(20)"`
	Content     string    `json:"html",orm:"sie(5000)"`
	ProjectId   int64     `orm:"null"`
	UserId      int64     `orm:"null"`
	Amount      int       `json:"amount"`
	Consider    bool      `json:"consider"`
	Views       int64     `orm:"default(0)"`
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Finance)) //, new(Finance)
}

//添加财务记录
func AddFinance(amount int, content, financedate string, projectid, uid int64, consider bool) (id int64, err error) {
	o := orm.NewOrm()
	//查询数据库中有无打卡
	// var finance Finance
	//判断是否有重名
	// err = o.QueryTable("finance").One(&finance)
	// if err == orm.ErrNoRows {
	// 没有找到记录
	finance := &Finance{
		Amount:      amount,
		Financedate: financedate,
		Content:     content,
		ProjectId:   projectid,
		UserId:      uid,
		Consider:    consider,
		Created:     time.Now(),
		Updated:     time.Now(),
	}
	id, err = o.Insert(finance)
	if err != nil {
		return 0, err
	}
	// return id, nil
	// } else if err == nil { //如果存在记录，或多于一条，则进行更新
	// 	finance2 := &Finance{Id: finance.Id}
	// 	// fmt.Printf(finance.Content)
	// 	finance2.Updated = time.Now()
	// 	finance2.Content = finance.Content + content
	// 	_, err = o.Update(finance2, "Content", "Updated")
	// 	if err != nil {
	// 		return finance.Id, err
	// 	}
	// 	return finance.Id, nil
	// }
	return id, nil
}

//微信小程序，根据projectid取得所有设代日志
func GetWxFinance(pid, limit, offset int64) (finance []*Finance, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Finance")
	_, err = qs.Filter("Projectid", pid).Limit(limit, offset).OrderBy("-Financedate").All(&finance)
	if err != nil {
		return nil, err
	}
	return finance, err
}

//获取日志总数
func GetWxFinanceCount(pid int64) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Finance")
	count, err = qs.Filter("Projectid", pid).Limit(-1).Count()
	if err != nil {
		return 0, err
	}
	return count, err
}

type FinanceUser struct {
	Finance `xorm:"extends"`
	User    `xorm:"extends"`
}

//微信小程序，根据projectid取得所有设代日志
func GetWxFinance2(pid int64, limit, offset int) ([]*FinanceUser, error) {
	finance := make([]*FinanceUser, 0)
	return finance, engine.Table("finance").Join("INNER", "user", "finance.user_id = user.id").Where("finance.project_id = ?", pid).Desc("financedate").Limit(limit, offset).Find(&finance)
}

//根据id取得日志
func GetFinance(id int64) (Finance1 *Finance, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Finance") //这个表名AchievementTopic需要用驼峰式，
	finance := new(Finance)
	err = qs.Filter("id", id).One(finance)
	// err = qs.Filter("Id", id).One(finance)
	// if err != nil {
	// 	return finance, err
	// }
	// 	if err == orm.ErrMultiRows {
	//     // 多条的时候报错
	//     fmt.Printf("Returned Multi Rows Not One")
	// }
	// if err == orm.ErrNoRows {
	//     // 没有找到记录
	//     fmt.Printf("Not row found")
	// }

	finance.Views++
	_, err = o.Update(finance)
	if err != nil {
		return finance, err
	}

	return finance, err
}

//修改
func UpdateFinance(id int64, amount int, content, financedate string, consider bool) error {
	o := orm.NewOrm()
	finance := &Finance{Id: id}
	if o.Read(finance) == nil {
		finance.Amount = amount
		finance.Content = content
		finance.Consider = consider
		finance.Financedate = financedate
		finance.Updated = time.Now()
		_, err := o.Update(finance, "Amount", "Content", "Consider", "Financedate", "Updated")
		if err != nil {
			return err
		}
	}
	return nil
}

//删除
func DeleteFinance(id int64) error {
	o := orm.NewOrm()
	finance := &Finance{Id: id}
	if o.Read(finance) == nil {
		_, err := o.Delete(finance)
		if err != nil {
			return err
		}
	}
	return nil
}
