package models

import (
	// "fmt"
	// beego "github.com/beego/beego/v2/adapter"
	// "github.com/jinzhu/gorm"
	// "time"
	// "errors"
	"gorm.io/gorm"
)

// 用户-模板表
type FreecadModel struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	// ClassID int64  `gorm:"column:class_id;foreignkey:ProjectId;"`
	UserID int64  `gorm:"column:user_id;` // 外键 (属于), tag `index`是为该列创建索引
	Number string `json:"number" gorm:"column:number"`
	// TempTitle  string `json:"temptitle" gorm:"column:temp_title"`
	// TempTitleB string `json:"temptitleb" gorm:"column:temp_title_b"` //简要名称，无日期，无版本号，无扩展名
	// TempPath   string `json:"temppath" gorm:"column:temp_path"`
	Title  string `json:"title" gorm:"column:title"`
	TitleB string `json:"titleb" gorm:"column:title_b"` //简要名称，无日期，无版本号，无扩展名
	Path   string `json:"path" gorm:"column:path"`
	// TempUrl   string `json:"tempurl" gorm:"column:temp_url"`
	Status  bool   `json:"status"`
	Version string `json:"version"`
	User    User   `gorm:"foreignkey:Id;references:UserID;"`
}

// 用户模板写入数据库
func AddFreecadModel(userid int64, number, title, titleb, path, version string) (id uint, err error) {
	db := _db //GetDB()
	//查询数据库中有无打卡
	// var businesscheckin BusinessCheckin
	freecadmodel := &FreecadModel{
		UserID: userid,
		Number: number,
		Title:  title,
		TitleB: titleb,
		Path:   path,
		// TempUrl:   templeurl,
		Version: version,
		Status:  true,
	}
	//判断是否有重名
	err = db.Where("user_id = ? AND title = ?", userid, title).FirstOrCreate(&freecadmodel).Error
	// err = o.QueryTable("business_checkin").Filter("ActivityId", ActivityId).Filter("UserId", UserId).Filter("SelectDate", SelectDate).One(&check1, "Id")
	// if err == orm.ErrNoRows {
	// 没有找到记录
	return freecadmodel.ID, err
}

// 查出所有模板
func GetFreecadModels(limit, offset int, searchtext string) (freecadmodels []FreecadModel, err error) {
	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//join一定要select,其他不用select的话默认查询全部。
	// Preload("BusinessUsers.NickNames")——嵌套预加载！！
	// number desc
	db := _db //GetDB()
	if searchtext != "" {
		err = db.Order("number").
			Preload("User").
			// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
			// Preload("User.Nickname").
			Where("title_b like ?", "%"+searchtext+"%").
			Limit(limit).Offset(offset).
			Find(&freecadmodels).Error
	} else {
		err = db.Order("number").
			Preload("User").
			// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
			// Preload("User.Nickname").
			// Where("class_id = ?", classid).
			Limit(limit).Offset(offset).
			Find(&freecadmodels).Error
	}
	return freecadmodels, err
}

// 查出所有模板数量
func GetFreecadModelCount(classid int64, searchtext string) (count int64, err error) {
	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//join一定要select,其他不用select的话默认查询全部。
	// Preload("BusinessUsers.NickNames")——嵌套预加载！！
	db := _db //GetDB()
	if searchtext != "" {
		err = db.Model(&FreecadModel{}).
			// Preload("User").
			// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
			// Preload("User.Nickname").
			Where("title_b like ?", "%"+searchtext+"%").
			Count(&count).Error
	} else {
		err = db.Model(&FreecadModel{}).
			// Preload("User").
			// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
			// Preload("User.Nickname").
			// Where("class_id = ?", classid).
			Count(&count).Error
	}
	return count, err
}

// 查出某个模板
func GetFreecadModel(id uint) (freecadmodel FreecadModel, err error) {
	db := _db //GetDB()
	err = db.
		// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
		Preload("User").
		Where("id = ?", id).
		Find(&freecadmodel).Error
	return freecadmodel, err
}
