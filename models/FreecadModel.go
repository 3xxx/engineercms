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
	FaceImgUrl   string `json:"faceimgurl"`
	DxfSvgUrl    string `json:"dxfsvgurl"`
	RenderImgUrl string `json:"renderimgurl"`
	Status       bool   `json:"status"`
	Version      string `json:"version"`
	Description  string `json:"description"`
	Indicators   string `json:"indicators"`
	User         User   `gorm:"foreignkey:Id;references:UserID;"`
}

// 参数-输入表
type FreecadInputs struct {
	gorm.Model
	FreecadModelID      uint
	Number              int                 `json:"number"`
	Name                string              `json:"name"`
	Symbol              string              `json:"symbol"`
	Unit                string              `json:"unit" gorm:"column:unit"`
	InputValue          float64             `json:"inputvalue" gorm:"column:input_value"` //默认值
	Remark              string              `json:"remark"`
	RealMin             string              `json:"realmin"`
	RealMax             string              `json:"realmax"`
	FCHistoryInputValue FCHistoryInputValue `json:"fchistoryinputvalue"`  //带条件的预加载,条件是UserHistoryID
	SelectValue         []Select2           `json:"selectvalue" gorm:"-"` // hasMany
	TextArealValue      TextAreal           `json:"textarealvalue" gorm:"-"`
	// RealValue    float64 `json:"realvalue" gorm:"column:real_value"`
}

// 用户-历史计算记录表
type FCUserHistory struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	UserID         int64 `json:"userid" gorm:"column:user_id;"` // 外键 (属于), tag `index`是为该列创建索引
	FreecadModelID uint  `json:"tempid" gorm:"column:freecad_model_id"`
	User           User  `gorm:"foreignkey:Id;references:UserID;"`
	FreecadModel   FreecadModel
	GlbUrl         string `json:"glburl"`
	FaceImgUrl     string `json:"faceimgurl"`
	Description    string `json:"description"`
	// TempleInputs  TempleInputs //根据usertempleid查询到templeinputs
	// TempleOutputs TempleOutputs
	// HistoryInputValue  []HistoryInputValue  `gorm:"foreignkey:ID"`
	// HistoryOutputValue []HistoryOutputValue `gorm:"foreignkey:ID"`
}

// 历史计算记录-输入参数记录表
type FCHistoryInputValue struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	FCUserHistoryID uint   `json:"fcuserhistoryid" foreignkey:FCUserHistoryID;" gorm:"column:fc_user_history_id"`
	FreecadInputsID uint   `json:"freecadinputsid" gorm:"column:freecad_inputs_id"`
	InputValue      string `json:"inputvalue" gorm:"column:input_value"`
	// TempleInputs   TempleInputs//invalid recursive type
}

// 来自passmathcadmodel
// type Select2 struct {
// 	// gorm.Model
// 	ID    string `json:"id"` // 这个在前端table中显示，必须是string
// 	Text  string `json:"text"`
// 	Value string `json:"value"`
// }

// type TextAreal struct {
// 	Value string `json:"value"`
// }

// 参数化三维模型写入数据库
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
	// err = db.Where("user_id = ? AND title = ?", userid, title).FirstOrCreate(&freecadmodel).Error
	err = db.Where(FreecadModel{Number: number, TitleB: titleb, Version: version}).Assign(FreecadModel{Title: title, Path: path}).FirstOrCreate(&freecadmodel).Error
	return freecadmodel.ID, err
	// err = o.QueryTable("business_checkin").Filter("ActivityId", ActivityId).Filter("UserId", UserId).Filter("SelectDate", SelectDate).One(&check1, "Id")
	// if err == orm.ErrNoRows {
	// 没有找到记录
	return freecadmodel.ID, err
}

// 模型封面写入数据库，如果数据已经存在，则更新封面
func AddFreecadModelFace(userid int64, number, titleb, faceimgurl, version string) (id uint, err error) {
	db := _db //GetDB()
	//查询数据库中有无打卡
	// var businesscheckin BusinessCheckin
	freecadmodel := &FreecadModel{
		UserID:     userid,
		Number:     number,
		TitleB:     titleb,
		FaceImgUrl: faceimgurl,
		// TempUrl:   templeurl,
		Version: version,
		Status:  true,
	}
	//使用了gorm的assign方法，没有找到，则建立（where里的字段加上assign里的字段？）；如果找到，则更新assign里的字段
	// err = db.Where(FreecadModel{Number: number, TitleB: titleb, Version: version}).Assign(FreecadModel{Number: number, TitleB: titleb, Version: version, FaceImgUrl: faceimgurl}).FirstOrCreate(&freecadmodel).Error
	err = db.Where(FreecadModel{Number: number, TitleB: titleb, Version: version}).Assign(FreecadModel{FaceImgUrl: faceimgurl}).FirstOrCreate(&freecadmodel).Error
	return freecadmodel.ID, err

	// Initialize and save new record with `Assign` attributes if not found
	// db.Where(User{Name: "non_existing"}).Assign(User{Age: 20}).FirstOrCreate(&user)
	// SQL: SELECT * FROM users WHERE name = 'non_existing';
	// SQL: INSERT INTO "users" (name, age) VALUES ("non_existing", 20);
	// user -> User{ID: 112, Name: "non_existing", Age: 20}

	// Update found record with `Assign` attributes
	// db.Where(User{Name: "jinzhu"}).Assign(User{Age: 20}).FirstOrCreate(&user)
	// SQL: SELECT * FROM users WHERE name = 'jinzhu';
	// SQL: UPDATE users SET age=20 WHERE id = 111;
	// user -> User{ID: 111, Name: "Jinzhu", Age: 20}
}

// 更新模型描述和指标
func UpdateFreecadModel(id uint, description, indicators string) (err error) {
	//获取DB
	db := _db //GetDB()
	// 条件更新
	result := db.Model(FreecadModel{}).Where("id = ?", id).Updates(FreecadModel{Description: description, Indicators: indicators})
	// db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
	return result.Error
}

// 删除模型
func DeleteFreeCAD(id uint) error {
	db := _db //GetDB()
	result := db.Where("id = ?", id).Delete(&FreecadModel{})
	// user.ID             // 返回插入数据的主键
	// result.Error        // 返回 error
	// result.RowsAffected // 返回插入记录的条数
	return result.Error
}

// 查出所有参数化模型
func GetFreecadModels(limit, offset int, searchtext string) (freecadmodels []FreecadModel, err error) {
	// 坑：preload里不是对应的表的名字，而是主表中字段名字！！！
	//join一定要select,其他不用select的话默认查询全部。
	// Preload("BusinessUsers.NickNames")——嵌套预加载！！
	// number desc
	db := _db //GetDB()
	if searchtext != "" {
		err = db.Order("ID").
			Preload("User").
			// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
			// Preload("User.Nickname").
			Where("title_b like ?", "%"+searchtext+"%").
			Limit(limit).Offset(offset).
			Find(&freecadmodels).Error
	} else {
		err = db.Order("ID").
			Preload("User").
			// Preload("BusinessUsers.NickNames", "id = ?", uid).//只预加载匹配的！
			// Preload("User.Nickname").
			// Where("class_id = ?", classid).
			Limit(limit).Offset(offset).
			Find(&freecadmodels).Error
	}
	return freecadmodels, err
}

// 查出所有参数化模型数量
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

// 解析模板的输入参数写入数据库,存入FC参数表格
func CreateFCModelInput(fcid uint, freecadinputs FreecadInputs) (id uint, err error) {
	//获取DB
	db := _db //GetDB()
	//保证id正确
	var fcinputs FreecadInputs
	// result := db.Where("excel_temple_id = ? AND input_alias=?", exceltempleid, inputalias).First(&excelinputs)
	// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 	excelinputs = ExcelInputs{ExcelTempleID: exceltempleid, InputAlias: inputalias, InputValue: inputvalue, ResultType: resulttype, Units: units}
	// 	// result = db.Create(&ExcelInputs{ExcelTempleID: exceltempleid, InputAlias: inputalias, InputValue: inputvalue, ResultType: resulttype, Units: units}) // 通过数据的指针来创建
	// 	result = db.Create(&excelinputs)
	// 	return excelinputs.ID, result.Error
	// } else {
	// 	return 0, result.Error
	// }
	fcinputs = FreecadInputs{
		FreecadModelID: fcid,
		Number:         freecadinputs.Number,
		Symbol:         freecadinputs.Symbol,
		Remark:         freecadinputs.Remark,
		Name:           freecadinputs.Name,
		InputValue:     freecadinputs.InputValue,
		Unit:           freecadinputs.Unit,
	}
	result := db.Create(&fcinputs) // 通过数据的指针来创建
	// user.ID             // 返回插入数据的主键
	// result.Error        // 返回 error
	// result.RowsAffected // 返回插入记录的条数
	return fcinputs.ID, result.Error
	// user.ID             // 返回插入数据的主键
	// result.Error        // 返回 error
	// result.RowsAffected // 返回插入记录的条数
}

// 查出FC参数表格
func GetFCModelInput(fcid uint) (freecadinputs []FreecadInputs, err error) {
	//获取DB
	db := _db //GetDB()
	err = db.
		Where("freecad_model_id = ?", fcid).
		Find(&freecadinputs).Error
	return freecadinputs, err
}
