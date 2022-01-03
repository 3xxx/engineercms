package models

import (
	"github.com/3xxx/engineercms/controllers/tool/result"
	"github.com/nu7hatch/gouuid"
	"time"
)

/**
 * the link table for Share and Matter.
 */
type Bridge struct {
	// Base
	// gorm.Model
	Id        uint   `json:"id" gorm:"primary_key"`
	Uuid      string `json:"uuid" gorm:"type:char(36);index:idx_uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt *time.Time
	ShareUuid string `json:"shareUuid" gorm:"type:char(36)"`
	ProductId int64  `json:"productid" gorm:"type:bigint(20) not null;default:0"`
	// MatterUuid string `json:"matterUuid" gorm:"type:char(36)"`
}

type OrderPair struct {
	Key   string
	Value string
}

type WherePair struct {
	Query string
	Args  []interface{}
}

// func (this *WherePair) And(where *WherePair) *WherePair {
// 	if this.Query == "" {
// 		return where
// 	} else {
// 		return &WherePair{Query: this.Query + " AND " + where.Query, Args: append(this.Args, where.Args...)}
// 	}
// }

// func (this *WherePair) Or(where *WherePair) *WherePair {
// 	if this.Query == "" {
// 		return where
// 	} else {
// 		return &WherePair{Query: this.Query + " OR " + where.Query, Args: append(this.Args, where.Args...)}
// 	}
// }

//放到adminmodel里去了
// func init() {
// orm.RegisterModel(new(Bridge))
// orm.RegisterModelWithPrefix("share_", new(Bridge))
// }

//pager
type Pager struct {
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	TotalItems int         `json:"totalItems"`
	TotalPages int         `json:"totalPages"`
	Data       interface{} `json:"data"`
}

// func NewPagerBridge(page int, pageSize int, totalItems int, data interface{}) *Pager {
// 	return &Pager{
// 		Page:       page,
// 		PageSize:   pageSize,
// 		TotalItems: totalItems,
// 		TotalPages: int(math.Ceil(float64(totalItems) / float64(pageSize))),
// 		Data:       data,
// 	}
// }

//find by uuid. if not found return nil.
// func FindByUuidBridge(uuid string) *Bridge {
// 	var bridge = &Bridge{}
// 	db := core.CONTEXT.GetDB().Where("uuid = ?", uuid).First(bridge)
// 	if db.Error != nil {
// 		if db.Error.Error() == result.DB_ERROR_NOT_FOUND {
// 			return nil
// 		} else {
// 			panic(db.Error)
// 		}
// 	}
// 	return bridge
// }

//find by uuid. if not found panic NotFound error
// func CheckByUuidBridge(uuid string) *Bridge {
// 	entity := this.FindByUuid(uuid)
// 	if entity == nil {
// 		panic(result.NotFound("not found record with uuid = %s", uuid))
// 	}
// 	return entity
// }

//find by shareUuid and matterUuid. if not found panic NotFound error.
// func CheckByShareUuidAndMatterUuid(shareUuid string, matterUuid string) *Bridge {
// 	var bridge = &Bridge{}
// 	db := core.CONTEXT.GetDB().Where("share_uuid = ? AND matter_uuid = ?", shareUuid, matterUuid).First(bridge)
// 	if db.Error != nil {
// 		if db.Error.Error() == result.DB_ERROR_NOT_FOUND {
// 			panic(result.NotFound("not found record with shareUuid = %s and matterUuid = %s", shareUuid, matterUuid))
// 		} else {
// 			panic(db.Error)
// 		}
// 	}
// 	return bridge
// }

//get pager
// func PageBridge(page int, pageSize int, shareUuid string, sortArray []OrderPair) *Pager {
// 	var wp = &builder.WherePair{}
// 	if shareUuid != "" {
// 		wp = wp.And(&builder.WherePair{Query: "share_uuid = ?", Args: []interface{}{shareUuid}})
// 	}
// 	var conditionDB *gorm.DB
// 	conditionDB = core.CONTEXT.GetDB().Model(&Bridge{}).Where(wp.Query, wp.Args...)
// 	count := 0
// 	db := conditionDB.Count(&count)
// 	this.PanicError(db.Error)
// 	var bridges []*Bridge
// 	db = conditionDB.Order(this.GetSortString(sortArray)).Offset(page * pageSize).Limit(pageSize).Find(&bridges)
// 	this.PanicError(db.Error)
// 	pager := NewPager(page, pageSize, count, bridges)
// 	return pager
// }

func CreateBridge(bridge *Bridge) (*Bridge, error) {
	db := _db //GetDB()
	timeUUID, _ := uuid.NewV4()
	bridge.Uuid = string(timeUUID.String())
	bridge.CreatedAt = time.Now()
	bridge.UpdatedAt = time.Now()
	// bridge.Sort = time.Now().UnixNano() / 1e6
	db = db.Table("share_bridge").Create(&bridge)
	// db := core.CONTEXT.GetDB().Create(bridge)
	// this.PanicError(db.Error)
	return bridge, db.Error
}

// func SaveBridge(bridge *Bridge) *Bridge {
// 	bridge.UpdateTime = time.Now()
// 	db := core.CONTEXT.GetDB().Save(bridge)
// 	this.PanicError(db.Error)
// 	return bridge
// }

// func DeleteBridge(bridge *Bridge) {
// 	db := core.CONTEXT.GetDB().Delete(&bridge)
// 	this.PanicError(db.Error)
// }

// func DeleteByMatterUuid(matterUuid string) {
// 	var wp = &builder.WherePair{}
// 	wp = wp.And(&builder.WherePair{Query: "matter_uuid = ?", Args: []interface{}{matterUuid}})
// 	db := core.CONTEXT.GetDB().Where(wp.Query, wp.Args).Delete(Bridge{})
// 	this.PanicError(db.Error)
// }

// func DeleteByShareUuid(shareUuid string) {
// 	var wp = &builder.WherePair{}
// 	wp = wp.And(&builder.WherePair{Query: "share_uuid = ?", Args: []interface{}{shareUuid}})
// 	db := core.CONTEXT.GetDB().Where(wp.Query, wp.Args).Delete(Bridge{})
// 	this.PanicError(db.Error)
// }

func FindByShareUuid(shareUuid string) ([]*Bridge, error) {
	if shareUuid == "" {
		panic(result.BadRequest("shareUuid cannot be nil"))
	}
	var bridges []*Bridge
	// db := GetDB()
	db := _db.Table("share_bridge").
		Where("share_uuid = ?", shareUuid).
		Find(&bridges)
	// this.PanicError(db.Error)
	return bridges, db.Error
}

// func CleanupBridge() {
// 	this.logger.Info("[BridgeDao] cleanup: delete all bridge records.")
// 	db := core.CONTEXT.GetDB().Where("uuid is not null").Delete(Bridge{})
// 	this.PanicError(db.Error)
// }

// func (this *Bridge) TableName() string {
// 	return core.TABLE_PREFIX + "bridge"
// }
