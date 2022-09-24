package models

import (
	// "errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

//用户-模板表
type PhotoData struct {
	gorm.Model
	YearMonth    string
	YearMonthDay string
	Url          string `gorm:"unique"`
	Name         string
}

// 写入数据库
func AddPhotoData(photodatas []PhotoData) error {
	db := _db //GetDB()
	// err := db.Create(&photodatas).Error //sqlite不能超过999条
	err := db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(photodatas, 100).Error
	// err := db.CreateInBatches(photodatas, 100).Error
	// err = db.Where("user_id = ? AND temp_title = ?", userid, templetitle).FirstOrCreate(&usertemple).Error
	return err
}

// type PhotoResult struct {
// 	YearMonth    string
// 	YearMonthDay string
// 	Url          string
// }

// 查询最新的limit个按天分组的照片数据，按yearmonth分组
func GetPhotoData(limit, offset int) (results []PhotoData, err error) {
	db := _db
	err = db.Limit(limit).Order("created_at desc").Offset(offset).Table("photo_data").Group("year_month_day").Scan(&results).Error
	return results, err
}

// 按yearmonthday查询所有照片
func GetDayPhotoData(yearmonthday string, limit, offset int) (results []PhotoData, err error) {
	db := _db
	err = db.Limit(limit).Offset(offset).Table("photo_data").Where("year_month_day = ?", yearmonthday).Scan(&results).Error
	return results, err
}

// 查找前一天的照片，降序排列DATE_FORMAT(from_unixtime(art_time),'%Y-%m')
// from_unixtime(timestamp,'%Y-%m-%d %H:%i:%s')小写m是数字月份，大写M是英文月份
// to_date(,"yyyy-mm-dd")
// DATE_FORMAT(time,'%Y-%m-%d')
// date(created_at) > ?
func GetPrevPhotoData(createdat time.Time) (results []PhotoData, err error) {
	db := _db
	err = db.Limit(1).Offset(0).Order("created_at desc").Table("photo_data").Where("created_at < ?", createdat).Scan(&results).Error
	return results, err
}

// 查找后一天的照片,升序排列
func GetNextPhotoData(createdat time.Time) (results []PhotoData, err error) {
	db := _db
	err = db.Limit(1).Offset(0).Order("created_at asc").Table("photo_data").Where("created_at > ?", createdat).Scan(&results).Error
	return results, err
}
