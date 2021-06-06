package models

import (
	// "fmt"
	// "github.com/astaxie/beego"
	// "gorm.io/gorm"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

//出差登记信息表
type Business struct {
	// gorm.Model
	ID               uint      `json:"id" gorm:"primary_key"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	UpdatedAt        time.Time
	DeletedAt        *time.Time
	ProjectID        int64 `json:"projectid" gorm:"column:project_id;foreignkey:ProjectId;"`
	UserID           int64 `gorm:"column:user_id"` // One-To-One (属于 - 本表的BillingAddressID作外键
	Location         string
	Lat              float64
	Lng              float64
	StartDate        time.Time
	EndDate          time.Time
	Projecttitle     string
	Drivername       string
	Subsidy          int
	Carfare          int
	Hotelfee         int
	BusinessUsers    []BusinessUser
	BusinessCheckins []BusinessCheckin
	// ArticleID     int64   `gorm:"column:article_id;foreignkey:ArticleId;"`
	Article  Article `gorm:"foreignkey:ArticleId"`
	Worktime float64
	Overtime int
}

//出差人员表
type BusinessUser struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	UserID     int64 // 外键 (属于), tag `index`是为该列创建索引
	BusinessID uint  `json:"businessid"` //这个对应business表中的ID
	// NickNames  NickName `gorm:"foreignkey:UserID"` //加不加这个references:UserID没所谓，奇怪
	NickNames User `gorm:"foreignkey:UserID"`
}

// 按道理，上面应该是`gorm:"foreignkey:ID;references:UserID"`，即主表businessuser中的USERID=
// =从表中的ID啊
// 这个是测试用的，实际用的是user
type NickName struct {
	gorm.Model
	NickName string
}

// 出差每天打卡，还是另外一个功能吧，不要放一起
type BusinessCheckin struct {
	gorm.Model
	BusinessID uint
	UserID     int64
	Location   string `json:"F_Location"`
	Lat        float64
	Lng        float64
	CheckTime  time.Time `gorm:"autoCreateTime"`
	SelectDate time.Time `gorm:"type:date"`
	PhotoUrl   string
	Users      User `gorm:"foreignkey:UserID"`
}

//用户充值记录
// type Recharge struct {
// 	gorm.Model
// 	UserID int64 `gorm:"column:user_id;foreignkey:UserId;"` // 外键 (属于), tag `index`是为该列创建索引
// 	Amount int   `gorm:"column:amount"`
// 	User   User  `gorm:"foreignkey:UserId"`
// }

func init() {
	_db.CreateTable(&Business{}, &BusinessUser{}, &NickName{}, &BusinessCheckin{}) //当第一个存在后，后面的不建立！！！bug
	_db.CreateTable(&BusinessUser{})
	_db.CreateTable(&NickName{})
	_db.CreateTable(&BusinessCheckin{})
}

// 登记出差活动
func CreateBusiness(business Business) (id uint, err error) {
	db := GetDB()
	// projectuser := ProjectUser{ProjectId: pid, UserId: uid}
	// 查询项目名称和时间段
	//判断是否有重名
	result := db.Where("projecttitle = ? AND start_date = ? AND end_date = ?", business.Projecttitle, business.StartDate, business.EndDate).First(&business)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result = db.Create(&business) // 通过数据的指针来创建
		return business.ID, result.Error
	} else {
		return 0, result.Error
	}
	// user.ID             // 返回插入数据的主键
	// result.Error        // 返回 error
	// result.RowsAffected // 返回插入记录的条数
}

// 更新出差活动
func UpdateBusiness(business Business) (err error) {
	db := GetDB()
	// projectuser := ProjectUser{ProjectId: pid, UserId: uid}
	// db.First(&business, id)
	result := db.Model(&business).Updates(business)
	return result.Error
	// user.ID             // 返回插入数据的主键
	// result.Error        // 返回 error
	// result.RowsAffected // 返回插入记录的条数
}

// 添加用户~出差关联表格
func CreateUserBusiness(businessuser BusinessUser) (id uint, err error) {
	db := GetDB()
	// projectuser := ProjectUser{ProjectId: pid, UserId: uid}
	result := db.Create(&businessuser) // 通过数据的指针来创建
	// user.ID             // 返回插入数据的主键
	// result.Error        // 返回 error
	// result.RowsAffected // 返回插入记录的条数
	return businessuser.ID, result.Error
}

// 删除用户~出差关联表格
func DeleteUserBusiness(userid int64, businessid uint) (err error) {
	db := GetDB()
	result := db.Where("user_id = ? AND business_id = ? ", userid, businessid).Delete(&BusinessUser{})
	// user.ID             // 返回插入数据的主键
	// result.Error        // 返回 error
	// result.RowsAffected // 返回插入记录的条数
	return result.Error
}

// 查出未过期的出差活动,与自己有关的出差活动
func GetAllBusiness(projectid, uid int64) (business []Business, err error) {
	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//join一定要select,其他不用select的话默认查询全部。
	// Preload("BusinessUsers.NickNames")——嵌套预加载！！
	// 加8个小时
	now := time.Now()
	h, _ := time.ParseDuration("1h")
	h1 := now.Add(8 * h)
	db := GetDB()
	err = db.Order("business.updated_at desc").
		Preload("BusinessUsers").
		// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
		Preload("BusinessUsers.NickNames").
		Joins("left JOIN business_user on business_user.business_id = business.id").
		Where("business_user.user_id = ?", uid).
		Where("business.project_id = ?", projectid).
		Where("business.start_date < ? AND business.end_date > ?", h1, h1).
		Find(&business).Error
	return business, err
}

// 查出出差活动
func GetBusinessById(businessid int64) (business Business, err error) {
	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//join一定要select,其他不用select的话默认查询全部。
	// Preload("BusinessUsers.NickNames")——嵌套预加载！！
	// 加8个小时
	db := GetDB()
	err = db.
		Preload("BusinessUsers").
		// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
		Preload("BusinessUsers.NickNames").
		Joins("left JOIN business_user on business_user.business_id = business.id").
		Where("business.id = ?", businessid).
		Find(&business).Error
	return business, err
}

// 查出未过期的出差活动——没有用到，试验用，保留
func GetAllBusiness2(projectid int64) (business []Business, err error) {
	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//join一定要select,其他不用select的话默认查询全部。
	// Preload("BusinessUsers.NickNames")——嵌套预加载！！
	db := GetDB()
	err = db.Order("business.updated_at desc").
		Preload("BusinessUsers").
		// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
		Preload("BusinessUsers.NickNames").
		Where("business.project_id = ?", projectid).
		Where("business.end_date > ?", time.Now()).
		Find(&business).Error
	return business, err
}

//打卡记录写入数据库
func BusinessCheck(businessid uint, userid int64, Lat, Lng float64, PhotoUrl, location string, SelectDate time.Time) (id uint, err error) {
	db := GetDB()
	//查询数据库中有无打卡
	// var businesscheckin BusinessCheckin
	businesscheckin := &BusinessCheckin{
		BusinessID: businessid,
		UserID:     userid,
		CheckTime:  time.Now(),
		PhotoUrl:   PhotoUrl,
		Lat:        Lat,
		Lng:        Lng,
		SelectDate: SelectDate,
		// PhotoUrl:PhotoUrl,
		Location: location,
	}
	//判断是否有重名
	err = db.Where("business_id = ? AND user_id = ? AND Select_date = ?", businessid, userid, SelectDate).FirstOrCreate(&businesscheckin).Error
	// err = o.QueryTable("business_checkin").Filter("ActivityId", ActivityId).Filter("UserId", UserId).Filter("SelectDate", SelectDate).One(&check1, "Id")
	// if err == orm.ErrNoRows {
	// 没有找到记录
	return businesscheckin.ID, err
}

// 按月查询打卡记录-businessid是unint，换成int64也行？
func GetBusinessCheck(businessid, UserId int64, SelectMonth1, SelectMonth2 time.Time) (check []*BusinessCheckin, err error) {
	db := GetDB()
	err = db.
		Where("business_id = ? AND user_id = ? AND select_date >= ? AND select_date <= ?", businessid, UserId, SelectMonth1, SelectMonth2).
		Find(&check).Error
	return check, err
}

//按月统计**************
func GetBusinessCheckUser(selectmonth1, selectmonth2 time.Time, limit, offset int) (business []Business, err error) {
	db := GetDB()
	err = db.Order("business.updated_at desc").
		Preload("BusinessCheckins", "select_date >= ? AND select_date <= ? ", selectmonth1, selectmonth2).
		Preload("BusinessCheckins.Users").
		// Joins("left JOIN business_checkin on business_checkin.business_id = business.id").
		// Where("business_checkin.select_date >= ? AND business_checkin.select_date <= ?", selectmonth1, selectmonth2).
		Where("business.start_date <= ? AND business.end_date >= ?", selectmonth2, selectmonth1).
		Find(&business).Error
	return business, err
}

// 根据businessid查询关联的users
func GetBusinessUsers(businessid uint) (users []*BusinessUser, err error) {
	db := GetDB()
	err = db.
		Preload("NickNames").
		Where("business_id = ? ", businessid).
		Find(&users).Error
	return users, err
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
