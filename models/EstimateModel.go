package models

import (
	// "fmt"
	// beego "github.com/beego/beego/v2/adapter"
	// "github.com/jinzhu/gorm"
	// "time"
	// "errors"
	"gorm.io/gorm"
	// "github.com/casbin/beego-orm-adapter"
)

// 项目表
type EstimateProject struct {
	gorm.Model
	Number  string `json:"number"`
	Name    string `json:"name"`
	Profile string `json:"profile"` // 概述
	Grade   string `json:"grade"`   // 工程等级
	Period  int    `json:"period"`  //  施工工期
	UserID  int64
	// User    User `json:"user" gorm:"foreignKey:UserID;references:Id;"` // 这个写法错误，所以无法建表
	User User `json:"user" gorm:"foreignkey:Id;references:UserID;"`
	// TotalInvestment    float64 `json:"totalinvestment"`
	// EstimateCategoryID int `json:"estimatecategoryid"`
	// EstimateCategory   EstimateCategory
	EstimateProjPhase []EstimateProjPhase `json:"estimateprojphase" gorm:"foreignkey:EstimateProjectID;references:ID;"` // 阶段
	// Project       Project     `json:"project" gorm:"foreignKey:Id;references:ProjectId;"`
}

// foreignKey指定在连接表中用作外键的当前模型的列名。
// references表示引用表中与连接表外键映射的列名。

// 项目阶段划分表
// type EstimatePhase struct {
// 	gorm.Model
// 	Name string `json:"name"`
// }

// 地区分类表
// type EstimateCategory struct {
// 	gorm.Model
// 	Province string `json:"province"`
// 	Category string `json:"category"`
// }

// 项目——阶段——价格水平
type EstimateProjPhase struct {
	gorm.Model
	EstimateProjectID uint    `json:"estimateprojectid"`
	PhaseName         string  `json:"phasename"`        // 项目阶段
	Information       string  `json:"information"`      // 价格水平年
	TotalInvestment   float64 `json:"totalinvestment"`  // 总投资
	StaticInvestment  float64 `json:"staticinvestment"` // 静态投资
}

// 作废！项目——阶段——工程专业部分professional建筑工程，机电设备及安装……
type EstimateProfessional struct {
	gorm.Model
	EstimateProjPhaseID uint `json:"estimateprojphaseid"`
	// ParentID            uint   `json:"parentid"`
	Number    string `json:"number"`
	Component string `json:"component"`
	// CostName  string  `json:"costname"`
	// Unit      string  `json:"unit"`
	// Quantity  float64 `json:"quantity"`
	// UnitPrice float64 `json:"unitprice"`
	Total float64 `json:"total"`
}

// 作废！工程专业部分——工程二级Secondary
type EstimateSecondary struct {
	gorm.Model
	EstimateProfessionalID uint    `json:"estimateprofessionalid"`
	Number                 string  `json:"number"`
	Component              string  `json:"component"`
	Total                  float64 `json:"total"`
}

// 作废！工程二级——工程三~六级
type EstimateTertiary struct {
	gorm.Model
	EstimateSecondaryID uint    `json:"estimatesecondaryid"`
	ParentID            uint    `json:"parentid"`
	Number              string  `json:"number"`
	Component           string  `json:"component"`
	Total               float64 `json:"total"`
}

// 建筑工程费用表Architectural
type EstimateCostArchi struct {
	gorm.Model
	EstimateProjPhaseID uint `json:"estimateprojphaseid"`
	ParentID            uint `json:"parentId"`
	// EstimateTertiaryID uint    `json:"estimatetertiaryid"`
	Number    string               `json:"number"`
	CostName  string               `json:"name"`
	Unit      string               `json:"unit"`
	Quantity  float64              `json:"quantity"`
	UnitPrice float64              `json:"unitprice"`
	Total     float64              `json:"total"`
	Children  []*EstimateCostArchi `json:"children" gorm:"-"`
}

// 机电设备费用表Electromechanical equipment
type EstimateCostElect struct {
	gorm.Model
	EstimateProjPhaseID uint `json:"estimateprojphaseid"`
	ParentID            uint `json:"parentId"`
	// EstimateTertiaryID uint    `json:"estimatetertiaryid"`
	Number                string               `json:"number"`
	CostName              string               `json:"name"`
	Unit                  string               `json:"unit"`
	Quantity              float64              `json:"quantity"`
	UnitPriceEquipment    float64              `json:"unitpriceequipment"`
	UnitPriceInstallation float64              `json:"unitpriceinstallation"`
	TotalEquipment        float64              `json:"totalequipment"`
	TotalInstallation     float64              `json:"totalinstallation"`
	Children              []*EstimateCostElect `json:"children" gorm:"-"`
}

// 金属结构费用表
type EstimateCostMetal struct {
	gorm.Model
	EstimateProjPhaseID uint `json:"estimateprojphaseid"`
	ParentID            uint `json:"parentId"`
	// EstimateTertiaryID uint    `json:"estimatetertiaryid"`
	Number                string               `json:"number"`
	CostName              string               `json:"name"`
	Unit                  string               `json:"unit"`
	Quantity              float64              `json:"quantity"`
	UnitPriceEquipment    float64              `json:"unitpriceequipment"`
	UnitPriceInstallation float64              `json:"unitpriceinstallation"`
	TotalEquipment        float64              `json:"totalequipment"`
	TotalInstallation     float64              `json:"totalinstallation"`
	Children              []*EstimateCostMetal `json:"children" gorm:"-"`
}

// 临时工程费用表Temporary
type EstimateCostTemp struct {
	gorm.Model
	EstimateProjPhaseID uint `json:"estimateprojphaseid"`
	ParentID            uint `json:"parentId"`
	// EstimateTertiaryID uint    `json:"estimatetertiaryid"`
	Number    string              `json:"number"`
	CostName  string              `json:"name"`
	Unit      string              `json:"unit"`
	Quantity  float64             `json:"quantity"`
	UnitPrice float64             `json:"unitprice"`
	Total     float64             `json:"total"`
	Children  []*EstimateCostTemp `json:"children" gorm:"-"`
}

//	type EstimatePriceInformation struct {
//		gorm.Model
//		EstimatePhaseID int    `json:"estimatephaseid"`
//		Information     string `json:"information"`
//	}
// &EstimatePhase{},

func init() {
	// _db.AutoMigrate(&EstimateProject{}, &EstimateProjPhase{}, &EstimateProfessional{}, &EstimateSecondary{}, &EstimateTertiary{}, &EstimateCost{})
}

// 项目信息写入数据库
func AddEstimateProject(number, name, profile, grade string, period int, userid int64) (id uint, err error) {
	db := _db //GetDB()
	//查询数据库中有无打卡
	// var businesscheckin BusinessCheckin
	estimateproject := &EstimateProject{
		Number:  number,
		Name:    name,
		Profile: profile,
		Grade:   grade,
		Period:  period,
		UserID:  userid,
	}
	//判断是否有重名
	err = db.Where("number = ? AND name = ?", number, name).FirstOrCreate(&estimateproject).Error
	// err = o.QueryTable("business_checkin").Filter("ActivityId", ActivityId).Filter("UserId", UserId).Filter("SelectDate", SelectDate).One(&check1, "Id")
	// if err == orm.ErrNoRows {
	// 没有找到记录
	return estimateproject.ID, err
}

// 修改项目信息
func UpdateEstProject(projectid uint, fieldname, value string) (err error) {
	//获取DB
	db := _db //GetDB()
	// 条件更新
	// var new_value bool
	// switch value {
	// case "true":
	// 	new_value = true
	// case "false":
	// 	new_value = false
	// }
	// 	case "comment":
	result := db.Model(&EstimateProject{}).Where("id = ?", projectid).Update(fieldname, value)
	return result.Error
}

// 项目阶段特性写入数据库
func AddEstimatePhase(projectid uint, phasename, information string, totalinvestment, staticinvestment float64) (id uint, err error) {
	db := _db //GetDB()
	//查询数据库中有无打卡
	// var businesscheckin BusinessCheckin
	estimatephase := &EstimateProjPhase{
		EstimateProjectID: projectid,
		PhaseName:         phasename,
		Information:       information,
		TotalInvestment:   totalinvestment,
		StaticInvestment:  staticinvestment,
	}
	//判断是否有重名
	err = db.Where("estimate_project_id = ? AND phase_name = ?", projectid, phasename).FirstOrCreate(&estimatephase).Error
	// err = o.QueryTable("business_checkin").Filter("ActivityId", ActivityId).Filter("UserId", UserId).Filter("SelectDate", SelectDate).One(&check1, "Id")
	// if err == orm.ErrNoRows {
	// 没有找到记录
	return estimatephase.ID, err
}

// 作废！一级专业部分（建筑工程）写入数据库
func AddEstimateProfessional(estimateprojphaseid uint, component string, total float64) (id uint, err error) {
	db := _db //GetDB()
	//查询数据库中有无打卡
	// var businesscheckin BusinessCheckin
	estimateprofessional := &EstimateProfessional{
		EstimateProjPhaseID: estimateprojphaseid,
		Component:           component,
		Total:               total,
	}
	//判断是否有重名
	err = db.Where("estimate_proj_phase_id = ? AND component = ?", estimateprojphaseid, component).FirstOrCreate(&estimateprofessional).Error
	// err = o.QueryTable("business_checkin").Filter("ActivityId", ActivityId).Filter("UserId", UserId).Filter("SelectDate", SelectDate).One(&check1, "Id")
	// if err == orm.ErrNoRows {
	// 没有找到记录
	return estimateprofessional.ID, err
}

// 作废！二级工程部分（一）写入数据库
func AddEstimateSecondary(estimateprofessionalid uint, number, component string, total float64) (id uint, err error) {
	db := _db //GetDB()
	//查询数据库中有无打卡
	// var businesscheckin BusinessCheckin
	estimatesecondary := &EstimateSecondary{
		EstimateProfessionalID: estimateprofessionalid,
		Number:                 number,
		Component:              component,
		Total:                  total,
	}
	//判断是否有重名
	err = db.Where("estimate_professional_id = ? AND component = ?", estimateprofessionalid, component).FirstOrCreate(&estimatesecondary).Error
	// err = o.QueryTable("business_checkin").Filter("ActivityId", ActivityId).Filter("UserId", UserId).Filter("SelectDate", SelectDate).One(&check1, "Id")
	// if err == orm.ErrNoRows {
	// 没有找到记录
	return estimatesecondary.ID, err
}

// 作废！三~五级工程部分（ (一) 1 (1) ）写入数据库
func AddEstimateTertiary(estimatesecondaryid, parentid uint, number, component string, total float64) (id uint, err error) {
	db := _db //GetDB()
	//查询数据库中有无打卡
	// var businesscheckin BusinessCheckin
	estimatetertiary := &EstimateTertiary{
		EstimateSecondaryID: estimatesecondaryid,
		ParentID:            parentid,
		Number:              number,
		Component:           component,
		Total:               total,
	}
	//判断是否有重名
	err = db.Where("estimate_secondary_id = ? AND component = ?", estimatesecondaryid, component).FirstOrCreate(&estimatetertiary).Error
	// err = o.QueryTable("business_checkin").Filter("ActivityId", ActivityId).Filter("UserId", UserId).Filter("SelectDate", SelectDate).One(&check1, "Id")
	// if err == orm.ErrNoRows {
	// 没有找到记录
	return estimatetertiary.ID, err
}

// cost写入数据库——建筑工程
func AddEstimateCostArchi(estimateProjPhaseID, parentID uint, number, costname, unit string, quantity, unitprice, total float64) (id uint, err error) {
	db := _db //GetDB()
	//查询数据库中有无打卡
	var estimatecost *EstimateCostArchi
	if estimateProjPhaseID == 0 && quantity == 0 && unitprice == 0 { // 2345级
		estimatecost = &EstimateCostArchi{
			ParentID: parentID,
			Number:   number,
			CostName: costname,
			Total:    total,
		}
	} else if parentID == 0 && quantity == 0 && unitprice == 0 { // 一级
		estimatecost = &EstimateCostArchi{
			EstimateProjPhaseID: estimateProjPhaseID,
			ParentID:            parentID,
			Number:              number,
			CostName:            costname,
			Total:               total,
		}
	} else if estimateProjPhaseID == 0 { // 工程量表
		estimatecost = &EstimateCostArchi{
			// EstimateProjPhaseID: estimateProjPhaseID,
			ParentID:  parentID,
			Number:    number,
			CostName:  costname,
			Unit:      unit,
			Quantity:  quantity,
			UnitPrice: unitprice,
			Total:     total,
		}
	}
	//判断是否有重名
	err = db.Where("estimate_proj_phase_id = ? AND cost_name = ? AND quantity=?", estimateProjPhaseID, costname, quantity).FirstOrCreate(&estimatecost).Error
	// err = o.QueryTable("business_checkin").Filter("ActivityId", ActivityId).Filter("UserId", UserId).Filter("SelectDate", SelectDate).One(&check1, "Id")
	// if err == orm.ErrNoRows {
	// 没有找到记录
	return estimatecost.ID, err
}

// cost写入数据库——机电设备安装工程
func AddEstimateCostElect(estimateProjPhaseID, parentID uint, number, costname, unit string, quantity, unitpriceEquipment, unitpriceInstallation, totalEquipment, totalInstallation float64) (id uint, err error) {
	db := _db //GetDB()
	//查询数据库中有无打卡
	var estimatecost *EstimateCostElect
	if estimateProjPhaseID == 0 && quantity == 0 && unitpriceEquipment == 0 { // 2345级
		estimatecost = &EstimateCostElect{
			ParentID:          parentID,
			Number:            number,
			CostName:          costname,
			TotalEquipment:    totalEquipment,
			TotalInstallation: totalInstallation,
		}
	} else if parentID == 0 && quantity == 0 && unitpriceEquipment == 0 { // 一级
		estimatecost = &EstimateCostElect{
			EstimateProjPhaseID: estimateProjPhaseID,
			ParentID:            parentID,
			Number:              number,
			CostName:            costname,
			TotalEquipment:      totalEquipment,
			TotalInstallation:   totalInstallation,
		}
	} else if estimateProjPhaseID == 0 { // 工程量表
		estimatecost = &EstimateCostElect{
			// EstimateProjPhaseID: estimateProjPhaseID,
			ParentID:              parentID,
			Number:                number,
			CostName:              costname,
			Unit:                  unit,
			Quantity:              quantity,
			UnitPriceEquipment:    unitpriceEquipment,
			UnitPriceInstallation: unitpriceInstallation,
			TotalEquipment:        totalEquipment,
			TotalInstallation:     totalInstallation,
		}
	}
	//判断是否有重名
	err = db.Where("estimate_proj_phase_id = ? AND cost_name = ? AND quantity=?", estimateProjPhaseID, costname, quantity).FirstOrCreate(&estimatecost).Error
	// err = o.QueryTable("business_checkin").Filter("ActivityId", ActivityId).Filter("UserId", UserId).Filter("SelectDate", SelectDate).One(&check1, "Id")
	// if err == orm.ErrNoRows {
	// 没有找到记录
	return estimatecost.ID, err
}

// cost写入数据库——金属结构安装工程
func AddEstimateCostMetal(estimateProjPhaseID, parentID uint, number, costname, unit string, quantity, unitpriceEquipment, unitpriceInstallation, totalEquipment, totalInstallation float64) (id uint, err error) {
	db := _db //GetDB()
	//查询数据库中有无打卡
	var estimatecost *EstimateCostMetal
	if estimateProjPhaseID == 0 && quantity == 0 && unitpriceEquipment == 0 { // 2345级
		estimatecost = &EstimateCostMetal{
			ParentID:          parentID,
			Number:            number,
			CostName:          costname,
			TotalEquipment:    totalEquipment,
			TotalInstallation: totalInstallation,
		}
	} else if parentID == 0 && quantity == 0 && unitpriceEquipment == 0 { // 一级
		estimatecost = &EstimateCostMetal{
			EstimateProjPhaseID: estimateProjPhaseID,
			ParentID:            parentID,
			Number:              number,
			CostName:            costname,
			TotalEquipment:      totalEquipment,
			TotalInstallation:   totalInstallation,
		}
	} else if estimateProjPhaseID == 0 { // 工程量表
		estimatecost = &EstimateCostMetal{
			// EstimateProjPhaseID: estimateProjPhaseID,
			ParentID:              parentID,
			Number:                number,
			CostName:              costname,
			Unit:                  unit,
			Quantity:              quantity,
			UnitPriceEquipment:    unitpriceEquipment,
			UnitPriceInstallation: unitpriceInstallation,
			TotalEquipment:        totalEquipment,
			TotalInstallation:     totalInstallation,
		}
	}
	//判断是否有重名
	err = db.Where("estimate_proj_phase_id = ? AND cost_name = ? AND quantity=?", estimateProjPhaseID, costname, quantity).FirstOrCreate(&estimatecost).Error
	// err = o.QueryTable("business_checkin").Filter("ActivityId", ActivityId).Filter("UserId", UserId).Filter("SelectDate", SelectDate).One(&check1, "Id")
	// if err == orm.ErrNoRows {
	// 没有找到记录
	return estimatecost.ID, err
}

// cost写入数据库——临时工程
func AddEstimateCostTemp(estimateProjPhaseID, parentID uint, number, costname, unit string, quantity, unitprice, total float64) (id uint, err error) {
	db := _db //GetDB()
	//查询数据库中有无打卡
	var estimatecost *EstimateCostTemp
	if estimateProjPhaseID == 0 && quantity == 0 && unitprice == 0 { // 2345级
		estimatecost = &EstimateCostTemp{
			ParentID: parentID,
			Number:   number,
			CostName: costname,
			Total:    total,
		}
	} else if parentID == 0 && quantity == 0 && unitprice == 0 { // 一级
		estimatecost = &EstimateCostTemp{
			EstimateProjPhaseID: estimateProjPhaseID,
			ParentID:            parentID,
			Number:              number,
			CostName:            costname,
			Total:               total,
		}
	} else if estimateProjPhaseID == 0 { // 工程量表
		estimatecost = &EstimateCostTemp{
			// EstimateProjPhaseID: estimateProjPhaseID,
			ParentID:  parentID,
			Number:    number,
			CostName:  costname,
			Unit:      unit,
			Quantity:  quantity,
			UnitPrice: unitprice,
			Total:     total,
		}
	}
	//判断是否有重名
	err = db.Where("estimate_proj_phase_id = ? AND cost_name = ? AND quantity=?", estimateProjPhaseID, costname, quantity).FirstOrCreate(&estimatecost).Error
	// err = o.QueryTable("business_checkin").Filter("ActivityId", ActivityId).Filter("UserId", UserId).Filter("SelectDate", SelectDate).One(&check1, "Id")
	// if err == orm.ErrNoRows {
	// 没有找到记录
	return estimatecost.ID, err
}

// 查询所有项目
func GetEstimateProjects(limit, offset int) (estimateProjects []EstimateProject, err error) {
	db := _db
	err = db.Order("updated_at desc").
		Preload("User").
		Preload("EstimateProjPhase").
		Limit(limit).Offset(offset).
		Find(&estimateProjects).Error
	return estimateProjects, err
}

// 查询所有项目
func GetEstimateProjectsCount() (count int64, err error) {
	db := _db
	err = db.Order("updated_at desc").
		Count(&count).Error
	return count, err
}

// 查询建筑工程投资
func GetEstimateCostArchi(estimateProjPhaseID uint, limit, offset int) (estimateCost []*EstimateCostArchi, err error) {
	db := _db
	// err = db.Raw("with RECURSIVE temp(id,parent_id,number,cost_name) as ( select id,parent_id,number,cost_name from estimate_cost where estimate_proj_phase_id = ? union all select t.* from estimate_cost t,temp where temp.id = t.parent_id ) select * from temp", estimateProjPhaseID).Scan(&estimateCost).Error
	// err = db.Raw("with RECURSIVE temp(id,estimate_proj_phase_id,parent_id,number,cost_name) as ( select id,estimate_proj_phase_id,parent_id,number,cost_name from estimate_cost where estimate_proj_phase_id = ? union all select estimate_cost.* from estimate_cost JOIN temp ON estimate_cost.parent_id = temp.id ) select * from temp", estimateProjPhaseID).Scan(&estimateCost).Error
	var q string
	q = `
		with recursive
    cao as (
      select * from estimate_cost_archi where estimate_proj_phase_id = ?
      union all
      select estimate_cost_archi.* from cao join estimate_cost_archi on cao.id = estimate_cost_archi.parent_id
    )
		select * from cao;
		`
	err = db.Raw(q, estimateProjPhaseID).Scan(&estimateCost).Error
	return estimateCost, err
}

// with recursive
//
//	cao as (
//	    select * from family where name = '曹操'
//	    union all
//	    select family.* from cao join family on cao.id = family.father
//	)
//
// select * from cao where emperor_no is not null;
// 查询机电设备安装工程投资
func GetEstimateCostElect(estimateProjPhaseID uint, limit, offset int) (estimateCost []*EstimateCostElect, err error) {
	db := _db
	// err = db.Raw("with RECURSIVE temp(id,parent_id,number,cost_name) as ( select id,parent_id,number,cost_name from estimate_cost where estimate_proj_phase_id = ? union all select t.* from estimate_cost t,temp where temp.id = t.parent_id ) select * from temp", estimateProjPhaseID).Scan(&estimateCost).Error
	// err = db.Raw("with RECURSIVE temp(id,estimate_proj_phase_id,parent_id,number,cost_name) as ( select id,estimate_proj_phase_id,parent_id,number,cost_name from estimate_cost where estimate_proj_phase_id = ? union all select estimate_cost.* from estimate_cost JOIN temp ON estimate_cost.parent_id = temp.id ) select * from temp", estimateProjPhaseID).Scan(&estimateCost).Error
	var q string
	q = `
		with recursive
    cao as (
      select * from estimate_cost_elect where estimate_proj_phase_id = ?
      union all
      select estimate_cost_elect.* from cao join estimate_cost_elect on cao.id = estimate_cost_elect.parent_id
    )
		select * from cao;
		`
	err = db.Raw(q, estimateProjPhaseID).Scan(&estimateCost).Error
	return estimateCost, err
}

// 查询金属结构安装工程投资
func GetEstimateCostMetal(estimateProjPhaseID uint, limit, offset int) (estimateCost []*EstimateCostMetal, err error) {
	db := _db
	// err = db.Raw("with RECURSIVE temp(id,parent_id,number,cost_name) as ( select id,parent_id,number,cost_name from estimate_cost where estimate_proj_phase_id = ? union all select t.* from estimate_cost t,temp where temp.id = t.parent_id ) select * from temp", estimateProjPhaseID).Scan(&estimateCost).Error
	// err = db.Raw("with RECURSIVE temp(id,estimate_proj_phase_id,parent_id,number,cost_name) as ( select id,estimate_proj_phase_id,parent_id,number,cost_name from estimate_cost where estimate_proj_phase_id = ? union all select estimate_cost.* from estimate_cost JOIN temp ON estimate_cost.parent_id = temp.id ) select * from temp", estimateProjPhaseID).Scan(&estimateCost).Error
	var q string
	q = `
		with recursive
    cao as (
      select * from estimate_cost_metal where estimate_proj_phase_id = ?
      union all
      select estimate_cost_metal.* from cao join estimate_cost_metal on cao.id = estimate_cost_metal.parent_id
    )
		select * from cao;
		`
	err = db.Raw(q, estimateProjPhaseID).Scan(&estimateCost).Error
	return estimateCost, err
}

// 查询临时工程投资
func GetEstimateCostTemp(estimateProjPhaseID uint, limit, offset int) (estimateCost []*EstimateCostTemp, err error) {
	db := _db
	// err = db.Raw("with RECURSIVE temp(id,parent_id,number,cost_name) as ( select id,parent_id,number,cost_name from estimate_cost where estimate_proj_phase_id = ? union all select t.* from estimate_cost t,temp where temp.id = t.parent_id ) select * from temp", estimateProjPhaseID).Scan(&estimateCost).Error
	// err = db.Raw("with RECURSIVE temp(id,estimate_proj_phase_id,parent_id,number,cost_name) as ( select id,estimate_proj_phase_id,parent_id,number,cost_name from estimate_cost where estimate_proj_phase_id = ? union all select estimate_cost.* from estimate_cost JOIN temp ON estimate_cost.parent_id = temp.id ) select * from temp", estimateProjPhaseID).Scan(&estimateCost).Error
	var q string
	q = `
		with recursive
    cao as (
      select * from estimate_cost_temp where estimate_proj_phase_id = ?
      union all
      select estimate_cost_temp.* from cao join estimate_cost_temp on cao.id = estimate_cost_temp.parent_id
    )
		select * from cao;
		`
	err = db.Raw(q, estimateProjPhaseID).Scan(&estimateCost).Error
	return estimateCost, err
}

// 查询建筑工程的投资数量
func GetEstimateCostArchiCount(estimateProjPhaseID uint) (count int64, err error) {
	db := _db
	// err = db.Raw("with RECURSIVE temp(id,estimate_proj_phase_id,parent_id,number,cost_name) as ( select id,estimate_proj_phase_id,parent_id,number,cost_name from estimate_cost where estimate_proj_phase_id = ? union all select estimate_cost.* from estimate_cost JOIN temp ON estimate_cost.parent_id = temp.id ) select * from temp", estimateProjPhaseID).Count(&count).Error
	var q string
	q = `
		with recursive
    cao as (
      select * from estimate_cost_archi where estimate_proj_phase_id = ?
      union all
      select estimate_cost_archi.* from cao join estimate_cost_archi on cao.id = estimate_cost_archi.parent_id
    )
		select id from cao;
		`
	err = db.Raw(q, estimateProjPhaseID).Count(&count).Error
	return count, err
}

// 查询机电设备安装工程的投资数量
func GetEstimateCostElectCount(estimateProjPhaseID uint) (count int64, err error) {
	db := _db
	// err = db.Raw("with RECURSIVE temp(id,estimate_proj_phase_id,parent_id,number,cost_name) as ( select id,estimate_proj_phase_id,parent_id,number,cost_name from estimate_cost where estimate_proj_phase_id = ? union all select estimate_cost.* from estimate_cost JOIN temp ON estimate_cost.parent_id = temp.id ) select * from temp", estimateProjPhaseID).Count(&count).Error
	var q string
	q = `
		with recursive
    cao as (
      select * from estimate_cost_elect where estimate_proj_phase_id = ?
      union all
      select estimate_cost_elect.* from cao join estimate_cost_elect on cao.id = estimate_cost_elect.parent_id
    )
		select id from cao;
		`
	err = db.Raw(q, estimateProjPhaseID).Count(&count).Error
	return count, err
}

// 查询金属结构安装工程的投资数量
func GetEstimateCostMetalCount(estimateProjPhaseID uint) (count int64, err error) {
	db := _db
	// err = db.Raw("with RECURSIVE temp(id,estimate_proj_phase_id,parent_id,number,cost_name) as ( select id,estimate_proj_phase_id,parent_id,number,cost_name from estimate_cost where estimate_proj_phase_id = ? union all select estimate_cost.* from estimate_cost JOIN temp ON estimate_cost.parent_id = temp.id ) select * from temp", estimateProjPhaseID).Count(&count).Error
	var q string
	q = `
		with recursive
    cao as (
      select * from estimate_cost_metal where estimate_proj_phase_id = ?
      union all
      select estimate_cost_metal.* from cao join estimate_cost_metal on cao.id = estimate_cost_metal.parent_id
    )
		select id from cao;
		`
	err = db.Raw(q, estimateProjPhaseID).Count(&count).Error
	return count, err
}

// 查询临时工程的投资数量
func GetEstimateCostTempCount(estimateProjPhaseID uint) (count int64, err error) {
	db := _db
	// err = db.Raw("with RECURSIVE temp(id,estimate_proj_phase_id,parent_id,number,cost_name) as ( select id,estimate_proj_phase_id,parent_id,number,cost_name from estimate_cost where estimate_proj_phase_id = ? union all select estimate_cost.* from estimate_cost JOIN temp ON estimate_cost.parent_id = temp.id ) select * from temp", estimateProjPhaseID).Count(&count).Error
	var q string
	q = `
		with recursive
    cao as (
      select * from estimate_cost_temp where estimate_proj_phase_id = ?
      union all
      select estimate_cost_temp.* from cao join estimate_cost_temp on cao.id = estimate_cost_temp.parent_id
    )
		select id from cao;
		`
	err = db.Raw(q, estimateProjPhaseID).Count(&count).Error
	return count, err
}
