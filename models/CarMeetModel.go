package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	// "strconv"
	// "strings"
	"time"
)

type CarCalendar struct {
	Id        int64     `json:"id",form:"-"`
	Title     string    `json:"title",form:"title;text;title:",valid:"MinSize(1);MaxSize(100)"` //orm:"unique",
	Content   string    `json:"content",orm:"sie(20)"`
	Starttime time.Time `json:"start",orm:"type(datetime)"`
	Endtime   time.Time `json:"end",orm:"type(datetime)"`
	Allday    bool      `json:"allDay",orm:"default(0)"`
	Color     string    `json:"color",orm:"null"`
	Public    bool      `default(true)`
	Ip        string    `json:"ip",orm:"null"`
	// Color     string    `json:"backgroundColor",orm:"null"`
	// BColor    string    `json:"borderColor",orm:"null"`
}

type MeetCalendar struct {
	Id        int64     `json:"id",form:"-"`
	Title     string    `json:"title",form:"title;text;title:",valid:"MinSize(1);MaxSize(100)"` //orm:"unique",
	Content   string    `json:"content",orm:"sie(20)"`
	Starttime time.Time `json:"start",orm:"type(datetime)"`
	Endtime   time.Time `json:"end",orm:"type(datetime)"`
	Allday    bool      `json:"allDay",orm:"default(0)"`
	Color     string    `json:"color",orm:"null"`
	Public    bool      `default(true)`
	Ip        string    `json:"ip",orm:"null"`
	// Color     string    `json:"backgroundColor",orm:"null"`
	// BColor    string    `json:"borderColor",orm:"null"`
}

func init() {
	orm.RegisterModel(new(CarCalendar), new(MeetCalendar)) //, new(Article)
}

//********汽车日历********
//添加
func AddCarCalendar(title, content, color, ip string, allday, public bool, start, end time.Time) (id int64, err error) {
	o := orm.NewOrm()
	calendar := &CarCalendar{
		Title:   title,
		Content: content,
		Color:   color,
		Ip:      ip,
		Allday:  allday,
		Public:  public,
		// BColor:    color,
		Starttime: start,
		Endtime:   end,
	}
	id, err = o.Insert(calendar)
	if err != nil {
		return id, err
	}
	// }
	return id, err
}

//取所有——要修改为支持时间段的，比如某个月份
func GetCarCalendar(start, end time.Time, public bool) (calendars []*CarCalendar, err error) {
	cond := orm.NewCondition()
	cond1 := cond.And("Starttime__gte", start).And("Starttime__lt", end) //这里全部用开始时间来判断
	o := orm.NewOrm()
	qs := o.QueryTable("CarCalendar")
	qs = qs.SetCond(cond1)

	// o := orm.NewOrm()
	calendars = make([]*CarCalendar, 0)
	// qs := o.QueryTable("CarCalendar")
	if public { //只取公开的
		_, err = qs.Filter("public", true).OrderBy("-Starttime").All(&calendars)
		if err != nil {
			return calendars, err
		}
	} else { //取全部
		_, err = qs.OrderBy("-Starttime").All(&calendars)
		if err != nil {
			return calendars, err
		}
	}
	return calendars, err
}

//修改
func UpdateCarCalendar(cid int64, title, content, color string, allday, public bool, start, end time.Time) error {
	o := orm.NewOrm()
	calendar := &CarCalendar{Id: cid}
	if o.Read(calendar) == nil {
		calendar.Title = title
		calendar.Content = content
		calendar.Color = color
		calendar.Allday = allday
		calendar.Public = public
		// calendar.BColor = color
		calendar.Starttime = start
		calendar.Endtime = end
		// calendar.Updated = time.Now()
		_, err := o.Update(calendar)
		if err != nil {
			return err
		}
	}
	return nil
}

//拖曳
func DropCarCalendar(cid int64, start, end time.Time) error {
	o := orm.NewOrm()
	calendar := &CarCalendar{Id: cid}
	if o.Read(calendar) == nil {
		calendar.Starttime = start
		calendar.Endtime = end
		_, err := o.Update(calendar)
		if err != nil {
			return err
		}
	}
	return nil
}

//resize
func ResizeCarCalendar(cid int64, end time.Time) error {
	o := orm.NewOrm()
	calendar := &CarCalendar{Id: cid}
	if o.Read(calendar) == nil {
		// calendar.Starttime = start
		calendar.Endtime = end
		// calendar.Updated = time.Now()
		_, err := o.Update(calendar)
		if err != nil {
			return err
		}
	}
	return nil
}

//根据id查询事件
func GetCarCalendarbyid(id int64) (calendar CarCalendar, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("CarCalendar")
	err = qs.Filter("id", id).One(&calendar)
	if err != nil {
		return calendar, err
	}
	return calendar, err
}

//删除事件
func DeleteCarCalendar(cid int64) error {
	o := orm.NewOrm()
	calendar := &CarCalendar{Id: cid}
	if o.Read(calendar) == nil {
		_, err := o.Delete(calendar)
		if err != nil {
			return err
		}
	}
	return nil
}

//*******会议室日历
//添加
func AddMeetCalendar(title, content, color, ip string, allday, public bool, start, end time.Time) (id int64, err error) {
	o := orm.NewOrm()
	calendar := &MeetCalendar{
		Title:   title,
		Content: content,
		Color:   color,
		Ip:      ip,
		Allday:  allday,
		Public:  public,
		// BColor:    color,
		Starttime: start,
		Endtime:   end,
	}
	id, err = o.Insert(calendar)
	if err != nil {
		return id, err
	}
	// }
	return id, err
}

//取所有——要修改为支持时间段的，比如某个月份
func GetMeetCalendar(start, end time.Time, public bool) (calendars []*MeetCalendar, err error) {
	cond := orm.NewCondition()
	cond1 := cond.And("Starttime__gte", start).And("Starttime__lt", end) //这里全部用开始时间来判断
	o := orm.NewOrm()
	qs := o.QueryTable("MeetCalendar")
	qs = qs.SetCond(cond1)

	// o := orm.NewOrm()
	calendars = make([]*MeetCalendar, 0)
	// qs := o.QueryTable("MeetCalendar")
	if public { //只取公开的
		_, err = qs.Filter("public", true).OrderBy("-Starttime").All(&calendars)
		if err != nil {
			return calendars, err
		}
	} else { //取全部
		_, err = qs.OrderBy("-Starttime").All(&calendars)
		if err != nil {
			return calendars, err
		}
	}
	return calendars, err
}

//修改
func UpdateMeetCalendar(cid int64, title, content, color string, allday, public bool, start, end time.Time) error {
	o := orm.NewOrm()
	calendar := &MeetCalendar{Id: cid}
	if o.Read(calendar) == nil {
		calendar.Title = title
		calendar.Content = content
		calendar.Color = color
		calendar.Allday = allday
		calendar.Public = public
		// calendar.BColor = color
		calendar.Starttime = start
		calendar.Endtime = end
		// calendar.Updated = time.Now()
		_, err := o.Update(calendar)
		if err != nil {
			return err
		}
	}
	return nil
}

//拖曳
func DropMeetCalendar(cid int64, start, end time.Time) error {
	o := orm.NewOrm()
	calendar := &MeetCalendar{Id: cid}
	if o.Read(calendar) == nil {
		calendar.Starttime = start
		calendar.Endtime = end
		_, err := o.Update(calendar)
		if err != nil {
			return err
		}
	}
	return nil
}

//resize
func ResizeMeetCalendar(cid int64, end time.Time) error {
	o := orm.NewOrm()
	calendar := &MeetCalendar{Id: cid}
	if o.Read(calendar) == nil {
		// calendar.Starttime = start
		calendar.Endtime = end
		// calendar.Updated = time.Now()
		_, err := o.Update(calendar)
		if err != nil {
			return err
		}
	}
	return nil
}

//根据id查询事件
func GetMeetCalendarbyid(id int64) (calendar MeetCalendar, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("MeetCalendar")
	err = qs.Filter("id", id).One(&calendar)
	if err != nil {
		return calendar, err
	}
	return calendar, err
}

//删除事件
func DeleteMeetCalendar(cid int64) error {
	o := orm.NewOrm()
	calendar := &MeetCalendar{Id: cid}
	if o.Read(calendar) == nil {
		_, err := o.Delete(calendar)
		if err != nil {
			return err
		}
	}
	return nil
}
