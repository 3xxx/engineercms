package models

import (
	// "fmt"
	// "github.com/astaxie/beego"
	// "gorm.io/gorm"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

//定位组表
type Location struct {
	// gorm.Model
	ID                uint      `json:"id" gorm:"primary_key"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time
	DeletedAt         *time.Time
	ProjectID         int64  `json:"projectid" gorm:"column:project_id;foreignkey:ProjectId;"`
	UserID            int64  `gorm:"column:user_id"` // One-To-One (属于 - 本表的BillingAddressID作外键
	Title             string `json:"title"`
	Describe          string `json:"describe"`
	Sort              int
	LocationNavigates []LocationNavigate `json:"locationnavigates"`
}

// 定位数据表
type LocationNavigate struct {
	gorm.Model
	LocationID uint    `json:"locationid"`
	Title      string  `json:"title"`
	Label      string  `json:"label"`
	Location   string  `json:"F_Location"`
	Address    string  `json:"F_Address"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	Sort       int
}

func init() {
	_db.CreateTable(&Location{}, &LocationNavigate{})
	_db.CreateTable(&LocationNavigate{})
}

// 登记定位数据组
func CreateLocation(location Location) (id uint, err error) {
	db := GetDB()
	// projectuser := ProjectUser{ProjectId: pid, UserId: uid}
	// 查询项目名称和时间段
	//判断是否有重名
	result := db.Where("title = ? AND project_id = ? ", location.Title, location.ProjectID).First(&location)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result = db.Create(&location) // 通过数据的指针来创建
		return location.ID, result.Error
	} else {
		return 0, result.Error
	}
	// user.ID             // 返回插入数据的主键
	// result.Error        // 返回 error
	// result.RowsAffected // 返回插入记录的条数
}

//定位导航写入数据库
func CreateLocationNavigate(locationnavigate LocationNavigate) (id uint, err error) {
	db := GetDB()
	//判断是否有重名
	err = db.Where("location_id = ? AND location = ?", locationnavigate.LocationID, locationnavigate.Location).FirstOrCreate(&locationnavigate).Error
	// err = o.QueryTable("location_checkin").Filter("ActivityId", ActivityId).Filter("UserId", UserId).Filter("SelectDate", SelectDate).One(&check1, "Id")
	// if err == orm.ErrNoRows {
	// 没有找到记录
	return locationnavigate.ID, err
}

// 更新出差活动
func UpdateLocation(location Location) (err error) {
	db := GetDB()
	// projectuser := ProjectUser{ProjectId: pid, UserId: uid}
	// db.First(&location, id)
	result := db.Model(&location).Updates(location)
	return result.Error
	// user.ID             // 返回插入数据的主键
	// result.Error        // 返回 error
	// result.RowsAffected // 返回插入记录的条数
}

// 查出定位组及里面的定位导航
func GetAllLocation(projectid int64) (location []Location, err error) {
	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//join一定要select,其他不用select的话默认查询全部。
	// Preload("LocationUsers.NickNames")——嵌套预加载！！
	// 加8个小时
	db := GetDB()
	err = db.Order("location.sort desc").
		Preload("LocationNavigates").
		Where("location.project_id = ?", projectid).
		Find(&location).Error
	return location, err
}

// 查出定位导航
func GetLocationById(locationid int64) (location Location, err error) {
	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//join一定要select,其他不用select的话默认查询全部。
	// Preload("LocationUsers.NickNames")——嵌套预加载！！
	// 加8个小时
	db := GetDB()
	err = db.
		Preload("LocationUsers").
		// Preload("LocationUsers.NickNames", "id = ?", uid).//只预加载匹配的！
		Preload("LocationUsers.NickNames").
		Joins("left JOIN location_user on location_user.location_id = location.id").
		Where("location.id = ?", locationid).
		Find(&location).Error
	return location, err
}

// 查出未过期的出差活动——没有用到，试验用，保留
func GetAllLocation2(projectid int64) (location []Location, err error) {
	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//join一定要select,其他不用select的话默认查询全部。
	// Preload("LocationUsers.NickNames")——嵌套预加载！！
	db := GetDB()
	err = db.Order("location.updated_at desc").
		Preload("LocationUsers").
		// Preload("LocationUsers.NickNames", "id = ?", uid).//只预加载匹配的！
		Preload("LocationUsers.NickNames").
		Where("location.project_id = ?", projectid).
		Where("location.end_date > ?", time.Now()).
		Find(&location).Error
	return location, err
}

//按月统计**************
func GetLocationCheckUser(selectmonth1, selectmonth2 time.Time, limit, offset int) (location []Location, err error) {
	db := GetDB()
	err = db.Order("location.updated_at desc").
		Preload("LocationCheckins", "select_date >= ? AND select_date <= ? ", selectmonth1, selectmonth2).
		Preload("LocationCheckins.Users").
		// Joins("left JOIN location_checkin on location_checkin.location_id = location.id").
		// Where("location_checkin.select_date >= ? AND location_checkin.select_date <= ?", selectmonth1, selectmonth2).
		Where("location.start_date <= ? AND location.end_date >= ?", selectmonth2, selectmonth1).
		Find(&location).Error
	return location, err
}

// type CheckUser struct {
// 	Checkin `xorm:"extends"`
// 	User    `xorm:"extends"`
// }

//按月统计**************
// func GetCheckUser2(selectmonth1, selectmonth2 time.Time, activityid int64, limit, offset int) ([]*CheckUser, error) {
// var users = make([]User, 0)
// users := make([]*CheckUser, 0)
// err = engine.Distinct("user_id").AllCols().Find(&checks)
// err := engine.Sql("select * from userinfo, userdetail where userinfo.detail_id = userdetail.id").Find(&users)
// return users, engine.Sql("SELECT DISTINCT checkin.user_id FROM checkin INNER JOIN user ON checkin.user_id = user.id LIMIT ? OFFSET ?", limit, offset).Find(&users)
// return users, engine.Table("checkin").Join("INNER", "user", "checkin.user_id = user.id").Where("checkin.select_date >= ? AND checkin.select_date <= ? AND checkin.activity_id = ?", selectmonth1, selectmonth2, activityid).Distinct("checkin.user_id", "user.nickname").Limit(limit, offset).Find(&users)
// return users, engine.Sql("SELECT * from checkin").Find(&users)
//distinct和limit不能和sql组合
// }

// 注释：Has Many一对多的外键、引用
// 1.默认外键是 从表中的字段为 主表模型的类型（type）加上其 主键（ID） 生成 ，如：从表card中的UserID
// 2.可以改变外键`gorm:"foreignKey:UserName"`
// 3.可以改变引用references:MemberNumber
// 4.用preload来查询关联，preload中的名字必须是主表中的字段名，不是从表名
// 5.不必是gorm建立的表才能这样用，beego orm建立的表也可以用
// 6.嵌套预加载中的foreignkey似乎反了？

// User 有多张 CreditCard，UserID 是外键
// type User struct {——主表
//   gorm.Model
//   CreditCards []CreditCard
// }

// type CreditCard struct {——从表
//   gorm.Model
//   Number string
//   UserID uint——这个是默认外键（主表名+ID），对应User主表中的ID，gorm.Model意味着ID和created等
// }

// type User struct {
//   gorm.Model
//   MemberNumber string
//   CreditCards  []CreditCard `gorm:"foreignKey:UserNumber;references:MemberNumber"`
// }

// type CreditCard struct {
//   gorm.Model
//   Number     string
//   UserNumber string——外键，这个值等于User表中的MemberNumber时，则查询到
// }

// 奶奶的，为啥和官网例子相反？？
// 文章
// type Topics struct {
// 	Id         int        `gorm:"primary_key"`
// 	Title      string     `gorm:"not null"`
// 	UserId     int        `gorm:"not null"`
// 	CategoryId int        `gorm:"not null"`
// 	Category   Categories `gorm:"foreignkey:CategoryId"`//文章所属分类外键
// 	User       Users      `gorm:"foreignkey:UserId"`//文章所属用户外键
// }

// 用户
// type Users struct {
// 	Id   int    `gorm:"primary_key"`
// 	Name string `gorm:"not null"`
// }

// 分类
// type Categories struct {
// 	Id   int    `gorm:"primary_key"`
// 	Name string `gorm:"not null"`
// }

// result := db.Where("id=?", id).Preload("Category").Preload("User").First(&topic)

// gorm的错误类型
// var (
// 	// ErrRecordNotFound record not found error
// 	ErrRecordNotFound = errors.New("record not found")
// 	// ErrInvalidTransaction invalid transaction when you are trying to `Commit` or `Rollback`
// 	ErrInvalidTransaction = errors.New("no valid transaction")
// 	// ErrNotImplemented not implemented
// 	ErrNotImplemented = errors.New("not implemented")
// 	// ErrMissingWhereClause missing where clause
// 	ErrMissingWhereClause = errors.New("WHERE conditions required")
// 	// ErrUnsupportedRelation unsupported relations
// 	ErrUnsupportedRelation = errors.New("unsupported relations")
// 	// ErrPrimaryKeyRequired primary keys required
// 	ErrPrimaryKeyRequired = errors.New("primary key required")
// 	// ErrModelValueRequired model value required
// 	ErrModelValueRequired = errors.New("model value required")
// 	// ErrInvalidData unsupported data
// 	ErrInvalidData = errors.New("unsupported data")
// 	// ErrUnsupportedDriver unsupported driver
// 	ErrUnsupportedDriver = errors.New("unsupported driver")
// 	// ErrRegistered registered
// 	ErrRegistered = errors.New("registered")
// 	// ErrInvalidField invalid field
// 	ErrInvalidField = errors.New("invalid field")
// 	// ErrEmptySlice empty slice found
// 	ErrEmptySlice = errors.New("empty slice found")
// 	// ErrDryRunModeUnsupported dry run mode unsupported
// 	ErrDryRunModeUnsupported = errors.New("dry run mode unsupported")
// )
