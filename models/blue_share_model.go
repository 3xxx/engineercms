package models

import (
	"github.com/nu7hatch/gouuid"
	"time"
)

// type Article struct {
// 	Id        int64     `json:"id",form:"-"`
// 	Subtext   string    `orm:"sie(20)"`
// 	Content   string    `json:"html",orm:"sie(5000)"`
// 	ProductId int64     `orm:"null"`
// 	Views     int64     `orm:"default(0)"`
// 	Created   time.Time `orm:"auto_now_add;type(datetime)"`
// 	Updated   time.Time `orm:"auto_now_add;type(datetime)"`
// }

const (
	//root matter's uuid
	MATTER_ROOT = "root"
	//cache directory name.
	MATTER_CACHE = "cache"
	//zip file temp directory.
	MATTER_ZIP             = "zip"
	MATTER_NAME_MAX_LENGTH = 200
	MATTER_NAME_MAX_DEPTH  = 32
	//matter name pattern
	MATTER_NAME_PATTERN = `[\\/:*?"<>|]`
)

type Share struct {
	// Base
	// gorm.Model
	Id        uint   `json:"id" gorm:"primary_key"`
	Uuid      string `json:"uuid" gorm:"type:char(36);index:idx_uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt *time.Time
	Name      string `json:"name" gorm:"type:varchar(255)"`
	ShareType string `json:"shareType" gorm:"type:varchar(45)"`
	Username  string `json:"username" gorm:"type:varchar(45)"`
	// UserUuid       string    `json:"userUuid" gorm:"type:char(36)"`
	UserId         int64     `json:"userid" gorm:"type:bigint(20) not null;default:0"`
	DownloadTimes  int64     `json:"downloadTimes" gorm:"type:bigint(20) not null;default:0"`
	Code           string    `json:"code" gorm:"type:varchar(45) not null"`
	ExpireInfinity bool      `json:"expireInfinity" gorm:"type:tinyint(1) not null;default:0"`
	ExpireTime     time.Time `json:"expireTime" gorm:"type:timestamp not null;default:'2018-01-01 00:00:00'"`
	// DirMatter      *Matter   `json:"dirMatter" gorm:"-"`
	// Matters        []*Matter `json:"matters" gorm:"-"`
}

type Matter struct {
	// Base
	Puuid    string    `json:"puuid" gorm:"type:char(36);index:idx_puuid"`
	UserUuid string    `json:"userUuid" gorm:"type:char(36);index:idx_uu"`
	Username string    `json:"username" gorm:"type:varchar(45) not null"`
	Dir      bool      `json:"dir" gorm:"type:tinyint(1) not null;default:0"`
	Name     string    `json:"name" gorm:"type:varchar(255) not null"`
	Md5      string    `json:"md5" gorm:"type:varchar(45)"`
	Size     int64     `json:"size" gorm:"type:bigint(20) not null;default:0"`
	Privacy  bool      `json:"privacy" gorm:"type:tinyint(1) not null;default:0"`
	Path     string    `json:"path" gorm:"type:varchar(1024)"`
	Times    int64     `json:"times" gorm:"type:bigint(20) not null;default:0"`
	Parent   *Matter   `json:"parent" gorm:"-"`
	Children []*Matter `json:"-" gorm:"-"`
}

//放到adminmodel里去了
// func init() {
// orm.RegisterModel(new(Share))
// orm.RegisterModelWithPrefix("share_", new(Share))
// }

func CreateShare(share *Share) (*Share, error) {
	timeUUID, _ := uuid.NewV4()
	share.Uuid = string(timeUUID.String())
	share.CreatedAt = time.Now()
	share.UpdatedAt = time.Now()
	// share.Sort = time.Now().UnixNano() / 1e6
	// db := core.CONTEXT.GetDB().Create(share)
	// this.PanicError(db.Error)
	// o := orm.NewOrm()
	// id, err = o.Insert(share)
	db := _db //GetDB()
	db = db.Table("share_share").Create(&share)

	return share, db.Error
}

//find by uuid. if not found return nil.
func FindByUuidShare(uuid string) (*Share, error) {
	// db := GetDB()
	// err = db.Where("user_id = ?", uid).First(&money).Error
	// return money, err
	var entity = &Share{}
	db := _db //GetDB()
	// db := core.CONTEXT.GetDB().Where("uuid = ?", uuid).First(entity)
	db = db.Table("share_share").Where("uuid = ?", uuid).First(entity)
	// db.Where("name = ?", "jinzhu").First(&user)
	// if db.Error != nil {
	// 	if db.Error.Error() == result.DB_ERROR_NOT_FOUND {
	// 		return nil, db.Error
	// 	} else {
	// 		panic(db.Error)
	// 	}
	// }
	return entity, db.Error
}

//find by uuid. if not found panic NotFound error
func CheckByUuidShare(uuid string) (*Share, error) {
	entity, err := FindByUuidShare(uuid)
	// if entity == nil {
	// 	panic(result.NotFound("not found record with uuid = %s", uuid))
	// }
	return entity, err
}

// func PageShare(page int, pageSize int, userUuid string, sortArray []OrderPair) *Pager {
// 	var wp = &builder.WherePair{}
// 	if userUuid != "" {
// 		wp = wp.And(&builder.WherePair{Query: "user_uuid = ?", Args: []interface{}{userUuid}})
// 	}
// 	var conditionDB *gorm.DB
// 	conditionDB = core.CONTEXT.GetDB().Model(&Share{}).Where(wp.Query, wp.Args...)
// 	count := 0
// 	db := conditionDB.Count(&count)
// 	this.PanicError(db.Error)
// 	var shares []*Share
// 	db = conditionDB.Order(this.GetSortString(sortArray)).Offset(page * pageSize).Limit(pageSize).Find(&shares)
// 	this.PanicError(db.Error)
// 	pager := NewPager(page, pageSize, count, shares)
// 	return pager
// }

// func (this *ShareDao) Create(share *Share) *Share {

// 	timeUUID, _ := uuid.NewV4()
// 	share.Uuid = string(timeUUID.String())
// 	share.CreateTime = time.Now()
// 	share.UpdateTime = time.Now()
// 	share.Sort = time.Now().UnixNano() / 1e6
// 	db := core.CONTEXT.GetDB().Create(share)
// 	this.PanicError(db.Error)

// 	return share
// }

// func SaveShare(share *Share) *Share {
// 	share.UpdateTime = time.Now()
// 	db := core.CONTEXT.GetDB().Save(share)
// 	this.PanicError(db.Error)
// 	return share
// }

// func DeleteShare(share *Share) {
// 	db := core.CONTEXT.GetDB().Delete(&share)
// 	this.PanicError(db.Error)
// }

//System cleanup.
// func CleanupShare() {
// 	this.logger.Info("[ShareDao] clean up. Delete all Share")
// 	db := core.CONTEXT.GetDB().Where("uuid is not null").Delete(Share{})
// 	this.PanicError(db.Error)
// }

func FindByUuids(uuids []int64) ([]*Product, error) {
	var products []*Product

	// if sortArray == nil || len(sortArray) == 0 {
	// 	return ""
	// }
	// str := ""
	// for _, pair := range sortArray {
	// 	if pair.Value == DIRECTION_DESC || pair.Value == DIRECTION_ASC {
	// 		if str != "" {
	// 			str = str + ","
	// 		}
	// 		str = str + " " + pair.Key + " " + pair.Value
	// 	}
	// }
	// return str
	// GetDB()
	db := _db.Where(uuids).Find(&products)
	// this.PanicError(db.Error)
	return products, db.Error
}
