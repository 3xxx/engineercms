package models

import (
	// "database/sql"
	// beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"

	// "github.com/go-xorm/xorm"

	// "strconv"
	// "strings"
	// "fmt"
	// "log"
	// "os"
	"time"
)

type Diary struct {
	Id        int64     `json:"id",form:"-"`
	Title     string    `orm:"sie(20)"`
	Diarydate string    `orm:"sie(20)"`
	Content   string    `json:"html",orm:"sie(5000)"`
	ProjectId int64     `orm:"null"`
	UserId    int64     `orm:"null"`
	Views     int64     `orm:"default(0)"`
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Diary)) //, new(Diary)
	// orm.RegisterDriver("sqlite", orm.DRSqlite)
	// orm.RegisterDataBase("default", "sqlite3", "database/engineer.db", 10)
}

//添加日记
func AddDiary(title, content, diarydate string, projectid, uid int64) (id int64, err error) {
	o := orm.NewOrm()
	//查询数据库中有无打卡
	var diary Diary
	//判断是否有重名
	err = o.QueryTable("diary").Filter("Diarydate", diarydate).One(&diary)
	if err == orm.ErrNoRows {
		// 没有找到记录
		diary := &Diary{
			Title:     title,
			Diarydate: diarydate,
			Content:   content,
			ProjectId: projectid,
			UserId:    uid,
			Created:   time.Now(),
			Updated:   time.Now(),
		}
		id, err = o.Insert(diary)
		if err != nil {
			return 0, err
		}
		return id, nil
	} else if err == nil { //如果存在记录，或多于一条，则进行更新
		diary2 := &Diary{Id: diary.Id}
		// fmt.Printf(diary.Content)
		diary2.Updated = time.Now()
		diary2.Content = diary.Content + content
		_, err = o.Update(diary2, "Content", "Updated")
		if err != nil {
			return diary.Id, err
		}
		return diary.Id, nil
	}
	return id, nil
}

//微信小程序，根据projectid取得所有设代日志
func GetWxDiaries(pid, limit, offset int64) (diaries []*Diary, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Diary")
	_, err = qs.Filter("Projectid", pid).Limit(limit, offset).OrderBy("-Diarydate").All(&diaries)
	if err != nil {
		return nil, err
	}
	return diaries, err
}

//获取日志总数
func GetWxDiaryCount(pid int64) (count int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Diary")
	count, err = qs.Filter("Projectid", pid).Limit(-1).Count()
	if err != nil {
		return 0, err
	}
	return count, err
}

type DiaryUser struct {
	Diary `xorm:"extends"`
	User  `xorm:"extends"`
}

//微信小程序，根据projectid取得所有设代日志
func GetWxDiaries2(pid int64, limit, offset int) ([]*DiaryUser, error) {
	diaries := make([]*DiaryUser, 0)
	return diaries, engine.Table("diary").Join("INNER", "user", "diary.user_id = user.id").Where("diary.project_id = ?", pid).Desc("diarydate").Limit(limit, offset).Find(&diaries)
}

//根据id取得日志
func GetDiary(id int64) (Diary1 *Diary, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Diary") //这个表名AchievementTopic需要用驼峰式，
	diary := new(Diary)
	err = qs.Filter("id", id).One(diary)
	// err = qs.Filter("Id", id).One(diary)
	// if err != nil {
	// 	return diary, err
	// }
	// 	if err == orm.ErrMultiRows {
	//     // 多条的时候报错
	//     fmt.Printf("Returned Multi Rows Not One")
	// }
	// if err == orm.ErrNoRows {
	//     // 没有找到记录
	//     fmt.Printf("Not row found")
	// }

	diary.Views++
	_, err = o.Update(diary)
	if err != nil {
		return diary, err
	}

	return diary, err
}

//修改
func UpdateDiary(id int64, title, content string) error {
	o := orm.NewOrm()
	diary := &Diary{Id: id}
	if o.Read(diary) == nil {
		diary.Title = title
		diary.Content = content
		diary.Updated = time.Now()
		_, err := o.Update(diary, "Title", "Content", "Updated")
		if err != nil {
			return err
		}
	}
	return nil
}

//删除
func DeleteDiary(id int64) error {
	o := orm.NewOrm()
	diary := &Diary{Id: id}
	if o.Read(diary) == nil {
		_, err := o.Delete(diary)
		if err != nil {
			return err
		}
	}
	return nil
}
