package models

import (
	// "fmt"
	// beego "github.com/beego/beego/v2/adapter"
	// "github.com/jinzhu/gorm"
	// "time"
	"errors"
	"gorm.io/gorm"
	// "github.com/casbin/beego-orm-adapter"
)

// 用户-模板表
type ExcelTemple struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	ClassID int64  `gorm:"column:class_id;foreignkey:ProjectId;"`
	UserID  int64  `gorm:"column:user_id;` // 外键 (属于), tag `index`是为该列创建索引
	Number  string `json:"number" gorm:"column:number"`
	Title   string `json:"title" gorm:"column:title"`
	TitleB  string `json:"titleb" gorm:"column:title_b"` //简要名称，无日期，无版本号，无扩展名
	Path    string `json:"path" gorm:"column:path"`
	// TempUrl   string `json:"excelurl" gorm:"column:excel_url"`
	Status  bool   `json:"status"`
	Version string `json:"version"`
	User    User   `gorm:"foreignkey:Id;references:UserID;"`
	Article ExcelArticle
}

// 模板-输入表
type ExcelInputs struct {
	gorm.Model
	ExcelTempleID          uint
	InputAlias             string                 `json:"inputalias" gorm:"column:input_alias"`
	InputValue             string                 `json:"inputvalue" gorm:"column:input_value"` //默认值
	ResultType             string                 `json:"resulttype" gorm:"column:result_type"`
	Units                  string                 `json:"units" gorm:"column:units"`
	Comment                string                 `json:"comment"`
	RealMin                string                 `json:"realmin"`
	RealMax                string                 `json:"realmax"`
	ExcelHistoryInputValue ExcelHistoryInputValue `json:"historyinputvalue"`
	SelectValue            []Select2              `json:"selectvalue" gorm:"-"`
	TextArealValue         TextAreal              `json:"textarealvalue" gorm:"-"`
	// RealValue    float64 `json:"realvalue" gorm:"column:real_value"`
}

// 模板-输出表
type ExcelOutputs struct {
	gorm.Model
	ExcelTempleID           uint
	OutputAlias             string                  `json:"outputalias" gorm:"column:output_alias"`
	OutputValue             string                  `json:"outputvalue" gorm:"column:output_value"`
	ResultType              string                  `json:"resulttype" gorm:"column:result_type"`
	Units                   string                  `json:"units" gorm:"column:units"`
	Comment                 string                  `json:"comment"`
	Images                  []ExcelImage            `json:"images" gorm:"-"`
	ExcelHistoryOutputValue ExcelHistoryOutputValue `json:"historyoutputvalue"`
	// RealValue    float64 `json:"realvalue" gorm:"column:real_value"`
}

type ExcelImage struct {
	Url string
}

// 用户-历史计算记录表
type ExcelHistory struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	UserID        int64 `json:"userid" gorm:"column:user_id;"` // 外键 (属于), tag `index`是为该列创建索引
	ExcelTempleID uint  `json:"exceltempleid" gorm:"column:excel_temple_id"`
	User          User  `gorm:"foreignkey:Id;references:UserID;"`
	ExcelTemple   ExcelTemple
	PdfUrl        string `json:"pdfurl"`
	FaceImgUrl    string `json:"faceimgurl"`
	Description   string `json:"description"`
	// ExcelInputs  ExcelInputs //根据exceltempleid查询到excelinputs
	// ExcelOutputs ExcelOutputs
	// HistoryInputValue  []HistoryInputValue  `gorm:"foreignkey:ID"`
	// HistoryOutputValue []HistoryOutputValue `gorm:"foreignkey:ID"`
}

// 历史计算记录-输入参数记录表
type ExcelHistoryInputValue struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	ExcelHistoryID uint   `json:"excel_history_id" foreignkey:ExcelHistoryID;" gorm:"column:excel_history_id"`
	ExcelInputsID  uint   `json:"excelinputsid" gorm:"column:excel_inputs_id"`
	InputValue     string `json:"inputvalue" gorm:"column:input_value"`
	// ExcelInputs   ExcelInputs//invalid recursive type
}

// 历史计算记录-输出参数记录表
type ExcelHistoryOutputValue struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	ExcelHistoryID uint   `foreignkey:ExcelHistoryID;" gorm:"column:excel_history_id"` // 外键 (属于), tag `index`是为该列创建索引
	ExcelOutputsID uint   `json:"exceloutputsid" gorm:"column:excel_outputs_id"`
	OutputValue    string `json:"outputvalue" gorm:"column:output_value"`
	// ExcelOutputs   ExcelOutputs//invalid recursive type
}

// 用户自定义计算书封面
type ExcelCalFace struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	UserID    int64 `gorm:"column:user_id;foreignkey:UserId;"` // 外键 (属于), tag `index`是为该列创建索引
	CalFaceID uint  `json:"calfaceid" gorm:"column:cal_face_id"`
}

// type ExcelCalFace struct {
// 	gorm.Model
// 	// ID     int    `gorm:"primary_key"`
// 	Tittle  string `gorm:"column:title;"` // 外键 (属于), tag `index`是为该列创建索引
// 	Content string `json:"content" gorm:"column:content"`
// }

// 用户添加计算书说明（前说明和后分析）explain
type ExcelCalExplain struct {
	gorm.Model
	UserID       int64 `gorm:"column:user_id;foreignkey:UserId;"` // 外键 (属于), tag `index`是为该列创建索引
	CalExplainID uint  `json:"calexplainid" gorm:"column:cal_explain_id"`
}

// type ExcelCalExplain struct {
// 	gorm.Model
// 	Tittle      string `gorm:"column:title"` // 外键 (属于), tag `index`是为该列创建索引
// 	Content     string `json:"content" gorm:"column:content"`
// 	FrontOrBack bool   `json:"frontorback" gorm:"column:front_or_back"`
// }

// excel的文章
type ExcelArticle struct {
	gorm.Model
	ExcelTempleID uint
	Title         string `json:"title" gorm:"column:title;size:20"`
	Subtext       string `json:"subtext" gorm:"column:subtext;size:20"`
	Content       string `json:"html" gorm:"column:content;size:5000"`
	// ExcelTemple   ExcelTemple
}

func init() {
	// _db.AutoMigrate(&ExcelTemple{}, &ExcelInputs{}, &ExcelOutputs{}, &ExcelHistory{}, &ExcelHistoryInputValue{}, &ExcelHistoryOutputValue{}, &ExcelArticle{})
	// _db.CreateTable(&ExcelTemple{})
	// _db.CreateTable(&ExcelInputs{})
	// _db.CreateTable(&ExcelOutputs{})

	// _db.CreateTable(&ExcelHistory{})
	// _db.CreateTable(&ExcelHistoryInputValue{})
	// _db.CreateTable(&ExcelHistoryOutputValue{})
	// _db.CreateTable(&ExcelArticle{})
}

// 用户模板写入数据库
func AddExcelTemple(classid, userid int64, number, title, titleb, path, version string) (id uint, err error) {
	db := _db //GetDB()
	//查询数据库中有无打卡
	// var businesscheckin BusinessCheckin
	exceltemple := &ExcelTemple{
		ClassID: classid,
		UserID:  userid,
		Number:  number,
		Title:   title,
		TitleB:  titleb,
		Path:    path,
		// TempUrl:   excelurl,
		Version: version,
		Status:  true,
	}
	//判断是否有重名
	err = db.Where("user_id = ? AND title = ?", userid, title).FirstOrCreate(&exceltemple).Error
	// err = o.QueryTable("business_checkin").Filter("ActivityId", ActivityId).Filter("UserId", UserId).Filter("SelectDate", SelectDate).One(&check1, "Id")
	// if err == orm.ErrNoRows {
	// 没有找到记录
	return exceltemple.ID, err
}

// 修改输入参数的信息
func UpdateExcel(excelid uint, fieldname, value string) (err error) {
	//获取DB
	db := _db //GetDB()
	// 条件更新
	var new_value bool
	switch value {
	case "true":
		new_value = true
	case "false":
		new_value = false
	}
	// 	case "comment":
	result := db.Model(&ExcelTemple{}).Where("id = ?", excelid).Update(fieldname, new_value)
	return result.Error
}

// 解析模板的输入参数写入数据库
func AddExcelInputAlias(exceltempleid uint, inputvalue, inputalias, resulttype, units string) (id uint, err error) {
	//获取DB
	db := _db //GetDB()
	//保证id正确
	var excelinputs ExcelInputs
	// result := db.Where("excel_temple_id = ? AND input_alias=?", exceltempleid, inputalias).First(&excelinputs)
	// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 	excelinputs = ExcelInputs{ExcelTempleID: exceltempleid, InputAlias: inputalias, InputValue: inputvalue, ResultType: resulttype, Units: units}
	// 	// result = db.Create(&ExcelInputs{ExcelTempleID: exceltempleid, InputAlias: inputalias, InputValue: inputvalue, ResultType: resulttype, Units: units}) // 通过数据的指针来创建
	// 	result = db.Create(&excelinputs)
	// 	return excelinputs.ID, result.Error
	// } else {
	// 	return 0, result.Error
	// }
	excelinputs = ExcelInputs{ExcelTempleID: exceltempleid, InputAlias: inputalias, InputValue: inputvalue, ResultType: resulttype, Units: units}
	result := db.Create(&excelinputs) // 通过数据的指针来创建
	// user.ID             // 返回插入数据的主键
	// result.Error        // 返回 error
	// result.RowsAffected // 返回插入记录的条数
	return excelinputs.ID, result.Error

	// if err = tx.Create(&Pay{UserID: uid, User2ID: product.Uid, ArticleID: articleid, Amount: newamount}).Error; err != nil {
	// 	return err
	// }
}

// 修改输入参数的备注信息
func UpdateExcelInputs(excelinputid uint, fieldname, value string) (err error) {
	//获取DB
	db := _db //GetDB()
	// 条件更新
	switch fieldname {
	case "realmin":
		fieldname = "real_min"
	case "realmax":
		fieldname = "real_max"
	}
	// 	case "comment":
	result := db.Model(&ExcelInputs{}).Where("id = ?", excelinputid).Update(fieldname, value)
	return result.Error
}

// 修改输出参数的备注信息
func UpdateExcelOutputs(exceloutputid uint, comment string) (err error) {
	//获取DB
	db := _db //GetDB()
	// 条件更新
	result := db.Model(&ExcelOutputs{}).Where("id = ?", exceloutputid).Update("comment", comment)
	return result.Error
}

// 解析模板的输出参数写入数据库
func AddExcelOutputAlias(exceltempleid uint, outputvalue, outputalias, resulttype, units string) (id uint, err error) {
	//获取DB
	db := _db //GetDB()
	//保证id正确
	var exceloutputs ExcelOutputs
	// result := db.Where("excel_temple_id = ? AND output_alias=?", exceltempleid, outputalias).First(&exceloutputs)
	// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 	exceloutputs = ExcelOutputs{ExcelTempleID: exceltempleid, OutputAlias: outputalias, OutputValue: outputvalue, ResultType: resulttype, Units: units}
	// 	result = db.Create(&exceloutputs)
	// 	return exceloutputs.ID, result.Error
	// } else {
	// 	return 0, result.Error
	// }
	exceloutputs = ExcelOutputs{ExcelTempleID: exceltempleid, OutputAlias: outputalias, OutputValue: outputvalue, ResultType: resulttype, Units: units}
	result := db.Create(&exceloutputs) // 通过数据的指针来创建
	return exceloutputs.ID, result.Error
}

// 查出所有模板
func GetExcelTemples(classid int64, limit, offset int, searchtext string) (exceltemples []ExcelTemple, err error) {
	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//join一定要select,其他不用select的话默认查询全部。
	// Preload("BusinessUsers.NickNames")——嵌套预加载！！
	db := _db //GetDB()
	if searchtext != "" {
		err = db.Order("updated_at desc").
			Preload("User").
			Preload("Article").
			// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
			// Preload("User.Nickname").
			Where("class_id = ? AND title_b like ?", classid, "%"+searchtext+"%").
			Limit(limit).Offset(offset).
			Find(&exceltemples).Error
	} else {
		err = db.Order("updated_at desc").
			Preload("User").
			Preload("Article").
			// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
			// Preload("User.Nickname").
			Where("class_id = ?", classid).
			Limit(limit).Offset(offset).
			Find(&exceltemples).Error
	}
	return exceltemples, err
}

// 查出所有模板数量
func GetExcelTempleCount(classid int64, searchtext string) (count int64, err error) {
	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//join一定要select,其他不用select的话默认查询全部。
	// Preload("BusinessUsers.NickNames")——嵌套预加载！！
	db := _db //GetDB()
	if searchtext != "" {
		err = db.Model(&ExcelTemple{}).
			// Preload("User").
			// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
			// Preload("User.Nickname").
			Where("class_id = ? AND title_b like ?", classid, "%"+searchtext+"%").
			Count(&count).Error
	} else {
		err = db.Model(&ExcelTemple{}).
			// Preload("User").
			// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
			// Preload("User.Nickname").
			Where("class_id = ?", classid).
			Count(&count).Error
	}

	return count, err
}

// 查出某个模板
func GetExcelTemple(excelid uint) (exceltemple ExcelTemple, err error) {
	db := _db //GetDB()
	err = db.
		// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
		Preload("User").
		Where("id = ?", excelid).
		Find(&exceltemple).Error
	return exceltemple, err
}

// 查出模板输入参数
func GetExcelInputs(excelid uint) (excelinputs []ExcelInputs, err error) {
	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//join一定要select,其他不用select的话默认查询全部。
	// Preload("BusinessUsers.NickNames")——嵌套预加载！！
	db := _db //GetDB()
	err = db.
		// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
		// Preload("User.Nickname").
		Where("excel_temple_id = ?", excelid).
		Find(&excelinputs).Error
	return excelinputs, err
}

// 查出模板输出参数
func GetExcelOutputs(excelid uint) (exceloutputs []ExcelOutputs, err error) {
	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//join一定要select,其他不用select的话默认查询全部。
	// Preload("BusinessUsers.NickNames")——嵌套预加载！！
	db := _db //GetDB()
	err = db.
		// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
		// Preload("User.Nickname").
		Where("excel_temple_id = ?", excelid).
		Find(&exceloutputs).Error
	return exceloutputs, err
}

// 用户计算历史写入
func CreateExcelHistory(userid int64, exceltempleid uint, pdfurl, faceimgurl, description string) (id uint, err error) {
	//获取DB
	db := _db //GetDB()
	//保证id正确
	excelhistory := ExcelHistory{UserID: userid, ExcelTempleID: exceltempleid, PdfUrl: pdfurl, FaceImgUrl: faceimgurl, Description: description}
	result := db.Create(&excelhistory) // 通过数据的指针来创建
	return excelhistory.ID, result.Error
}

// 用户计输入算数据写入, excelinputsid
func AddExcelHistoryInputValue(excelhistoryid uint, excelinputs []ExcelInputs) (id uint, err error) {
	db := _db //GetDB()
	var excelhistoryinputvalue ExcelHistoryInputValue
	// result := map[string]interface{}{}
	var result *gorm.DB
	for _, v := range excelinputs {
		// for i := 0; i < len(excelinputs); i++ {
		result = db.Where("excel_history_id = ? AND excel_inputs_id = ? ", excelhistoryid, v.ID).First(&excelhistoryinputvalue)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			excelhistoryinputvalue = ExcelHistoryInputValue{ExcelHistoryID: excelhistoryid, ExcelInputsID: v.ID, InputValue: v.InputValue}
			result = db.Create(&excelhistoryinputvalue)
			// 开始这里有return，导致只算一次就返回了，百思不得其解！！不应该这里放return
		}
	}
	return excelhistoryinputvalue.ID, result.Error
}

// 用户计算输出数据写入, exceloutputsid
func AddExcelHistoryOutputValue(excelhistoryid uint, exceloutputs []ExcelOutputs) (id uint, err error) {
	db := _db //GetDB()
	var excelhistoryoutputvalue ExcelHistoryOutputValue
	// result := map[string]interface{}{}
	var result *gorm.DB
	for _, v := range exceloutputs {
		// for i := 0; i < len(excelinputs); i++ {
		result = db.Where("excel_history_id = ? AND excel_outputs_id = ? ", excelhistoryid, v.ID).First(&excelhistoryoutputvalue)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			excelhistoryoutputvalue = ExcelHistoryOutputValue{ExcelHistoryID: excelhistoryid, ExcelOutputsID: v.ID, OutputValue: v.OutputValue}
			result = db.Create(&excelhistoryoutputvalue)
			// 开始这里有return，导致只算一次就返回了，百思不得其解！！不应该这里放return
		}
	}
	return excelhistoryoutputvalue.ID, result.Error
}

// 查询用户计算历史列表
func GetExcelHistory(userid int64, exceltempleid uint, limit, offset int) (excelhistory []ExcelHistory, err error) {
	db := _db //GetDB()
	if exceltempleid == 0 {
		err = db.Order("excel_history.updated_at desc").
			Preload("ExcelTemple").
			Preload("User", "id = ?", userid).
			Where("user_id = ?", userid).
			Limit(limit).Offset(offset).
			Find(&excelhistory).Error
	} else {
		err = db.Order("excel_history.updated_at desc").
			Preload("ExcelTemple", "id = ?", exceltempleid).
			Preload("User", "id = ?", userid).
			Where("user_id = ? AND excel_temple_id=?", userid, exceltempleid).
			Limit(limit).Offset(offset).
			Find(&excelhistory).Error
	}
	return excelhistory, err
}

func GetExcelHistoryCount(userid int64, exceltempleid uint) (count int64, err error) {
	db := _db //GetDB()
	if exceltempleid == 0 {
		err = db.Model(&ExcelHistory{}).
			Preload("ExcelTemple").
			Preload("User", "id = ?", userid).
			Where("user_id = ?", userid).
			Count(&count).Error
	} else {
		err = db.Model(&ExcelHistory{}).
			Preload("ExcelTemple", "id = ?", exceltempleid).
			Preload("User", "id = ?", userid).
			Where("user_id = ? AND excel_temple_id=?", userid, exceltempleid).
			Count(&count).Error
	}
	return count, err
}

// 查询历史计算数值_作废！！
func GetExcelHistoryValue(excelhistoryid uint) (excelhistory ExcelHistory, err error) {
	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//join一定要select,其他不用select的话默认查询全部。
	// Preload("BusinessUsers.NickNames")——嵌套预加载！！
	db := _db //GetDB()
	err = db.
		Preload("User").
		Preload("ExcelInputs").
		Where("id = ?", excelhistoryid).
		Joins("left join history_input_value AS t1 on t1.excel_history_id = excelhistoryid").
		Joins("left join history_input_value AS t2 ON t2.excel_inputs_id = excel_inputs.id").
		Joins("left join history_output_value AS t3 on t3.excel_history_id = excelhistoryid").
		Joins("left join history_output_value AS t4 ON t4.excel_outputs_id = excel_outputs.id").
		Find(&excelhistory).Error
	return excelhistory, err
}

// 根据historyid查询历史计算
func GetHistoryExcel(historyid uint) (excelhistory ExcelHistory, err error) {
	db := _db //GetDB()
	err = db.
		Where("id = ?", historyid).
		Find(&excelhistory).Error
	return excelhistory, err
}

// 查出模板历史输入参数
func GetExcelHistoryInputs(exceltempleid, excelhistoryid uint) (excelinputs []ExcelInputs, err error) {
	db := _db //GetDB()
	err = db.
		Preload("ExcelHistoryInputValue", "excel_history_id = ?", excelhistoryid).
		Where("excel_temple_id = ?", exceltempleid).
		// Joins("left join excel_history AS t1 on t1.ExcelTempleID = HistoryInputValue.").
		// Joins("left join history_input_value AS t2 ON t2.excel_inputs_id = excel_inputs.id").
		Find(&excelinputs).Error
	return excelinputs, err
}

// 查出模板历史输出参数
func GetExcelHistoryOutputs(exceltempleid, excelhistoryid uint) (exceloutputs []ExcelOutputs, err error) {
	db := _db //GetDB()
	err = db.
		Preload("ExcelHistoryOutputValue", "excel_history_id = ?", excelhistoryid).
		Where("excel_temple_id = ?", exceltempleid).
		// Joins("left join excel_history AS t1 on t1.ExcelTempleID = HistoryInputValue.").
		// Joins("left join history_input_value AS t2 ON t2.excel_inputs_id = excel_inputs.id").
		Find(&exceloutputs).Error
	return exceloutputs, err
}

// 删除模板
func DeleteExcel(excelid uint) error {
	db := _db //GetDB()
	result := db.Where("id = ?", excelid).Delete(&ExcelTemple{})
	// user.ID             // 返回插入数据的主键
	// result.Error        // 返回 error
	// result.RowsAffected // 返回插入记录的条数
	return result.Error
}

// 添加文章作为用户模板的附件
func AddExcelArticle(exceltempleid uint, title, subtext, content string) (id uint, err1, err2 error) {
	db := _db //GetDB()
	var result2 *gorm.DB
	//保证id正确
	var excelarticle ExcelArticle
	result1 := db.Where("excel_temple_id = ?", exceltempleid).First(&excelarticle)
	if errors.Is(result1.Error, gorm.ErrRecordNotFound) {
		result2 = db.Create(&ExcelArticle{ExcelTempleID: exceltempleid, Title: title, Subtext: subtext, Content: content}) // 通过数据的指针来创建
		return excelarticle.ID, result1.Error, result2.Error                                                               //这样无法返回ID，参见CreateExcelHistory
	} else {
		return excelarticle.ID, result1.Error, nil
	}
}

// 根据excelarticle取得文章
func GetExcelArticle(excelarticleid uint) (excelarticle ExcelArticle, err error) {
	db := _db //GetDB()
	err = db.
		Where("id = ?", excelarticleid).
		Find(&excelarticle).Error
	return excelarticle, err
}

// 编辑文章
func UpdateExcelArticle(excelarticleid uint, title, subtext, content string) (err error) {
	//获取DB
	db := _db //GetDB()
	// 条件更新
	result := db.Model(ExcelArticle{}).Where("id = ?", excelarticleid).Updates(ExcelArticle{Title: title, Subtext: subtext, Content: content})
	// db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
	return result.Error
}
