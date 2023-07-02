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
type UserTemple struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	ClassID int64  `gorm:"column:class_id;foreignkey:ProjectId;"`
	UserID  int64  `gorm:"column:user_id;` // 外键 (属于), tag `index`是为该列创建索引
	Number  string `json:"number" gorm:"column:number"`
	// TempTitle  string `json:"temptitle" gorm:"column:temp_title"`
	// TempTitleB string `json:"temptitleb" gorm:"column:temp_title_b"` //简要名称，无日期，无版本号，无扩展名
	// TempPath   string `json:"temppath" gorm:"column:temp_path"`
	TempTitle  string `json:"title" gorm:"column:temp_title"`
	TempTitleB string `json:"titleb" gorm:"column:temp_title_b"` //简要名称，无日期，无版本号，无扩展名
	TempPath   string `json:"path" gorm:"column:temp_path"`
	// TempUrl   string `json:"tempurl" gorm:"column:temp_url"`
	Status  bool   `json:"status"`
	Version string `json:"version"`
	User    User   `gorm:"foreignkey:Id;references:UserID;"`
	// MathArticle MathArticle
	Article MathArticle
}

// 模板-输入表
type TempleInputs struct {
	gorm.Model
	UserTempleID      uint
	InputAlias        string            `json:"inputalias" gorm:"column:input_alias"`
	InputValue        string            `json:"inputvalue" gorm:"column:input_value"` //默认值
	ResultType        string            `json:"resulttype" gorm:"column:result_type"`
	Units             string            `json:"units" gorm:"column:units"`
	Comment           string            `json:"comment"`
	RealMin           string            `json:"realmin"`
	RealMax           string            `json:"realmax"`
	HistoryInputValue HistoryInputValue `json:"historyinputvalue"`    //带条件的预加载,条件是UserHistoryID
	SelectValue       []Select2         `json:"selectvalue" gorm:"-"` // hasMany
	TextArealValue    TextAreal         `json:"textarealvalue" gorm:"-"`
	// RealValue    float64 `json:"realvalue" gorm:"column:real_value"`
}

type Select2 struct {
	// gorm.Model
	ID    string `json:"id"` // 这个在前端table中显示，必须是string
	Text  string `json:"text"`
	Value string `json:"value"`
}

type TextAreal struct {
	Value string `json:"value"`
}

// type Select struct {
// 	Value []string
// }

// 模板-输出表
type TempleOutputs struct {
	gorm.Model
	UserTempleID       uint
	OutputAlias        string             `json:"outputalias" gorm:"column:output_alias"`
	OutputValue        string             `json:"outputvalue" gorm:"column:output_value"`
	ResultType         string             `json:"resulttype" gorm:"column:result_type"`
	Units              string             `json:"units" gorm:"column:units"`
	Comment            string             `json:"comment"`
	HistoryOutputValue HistoryOutputValue `json:"historyoutputvalue"`
	// RealValue    float64 `json:"realvalue" gorm:"column:real_value"`
}

// 用户-历史计算记录表
type UserHistory struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	UserID       int64 `json:"userid" gorm:"column:user_id;"` // 外键 (属于), tag `index`是为该列创建索引
	UserTempleID uint  `json:"tempid" gorm:"column:user_temple_id"`
	User         User  `gorm:"foreignkey:Id;references:UserID;"`
	UserTemple   UserTemple
	PdfUrl       string `json:"pdfurl"`
	FaceImgUrl   string `json:"faceimgurl"`
	Description  string `json:"description"`
	// TempleInputs  TempleInputs //根据usertempleid查询到templeinputs
	// TempleOutputs TempleOutputs
	// HistoryInputValue  []HistoryInputValue  `gorm:"foreignkey:ID"`
	// HistoryOutputValue []HistoryOutputValue `gorm:"foreignkey:ID"`
}

// 历史计算记录-输入参数记录表
type HistoryInputValue struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	UserHistoryID  uint   `json:"user_history_id" foreignkey:UserHistoryID;" gorm:"column:user_history_id"`
	TempleInputsID uint   `json:"templeinputsid" gorm:"column:temple_inputs_id"`
	InputValue     string `json:"inputvalue" gorm:"column:input_value"`
	// TempleInputs   TempleInputs//invalid recursive type
}

// 历史计算记录-输出参数记录表
type HistoryOutputValue struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	UserHistoryID   uint   `foreignkey:UserHistoryID;" gorm:"column:user_history_id"` // 外键 (属于), tag `index`是为该列创建索引
	TempleOutputsID uint   `json:"templeoutputsid" gorm:"column:temple_outputs_id"`
	OutputValue     string `json:"outputvalue" gorm:"column:output_value"`
	// TempleOutputs   TempleOutputs//invalid recursive type
}

// 用户自定义计算书封面
type UserCalFace struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	UserID    int64 `gorm:"column:user_id;foreignkey:UserId;"` // 外键 (属于), tag `index`是为该列创建索引
	CalFaceID uint  `json:"calfaceid" gorm:"column:cal_face_id"`
}

type CalFace struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	Tittle  string `gorm:"column:title;"` // 外键 (属于), tag `index`是为该列创建索引
	Content string `json:"content" gorm:"column:content"`
}

// 用户添加计算书说明（前说明和后分析）explain
type UserCalExplain struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	UserID       int64 `gorm:"column:user_id;foreignkey:UserId;"` // 外键 (属于), tag `index`是为该列创建索引
	CalExplainID uint  `json:"calexplainid" gorm:"column:cal_explain_id"`
}

type CalExplain struct {
	gorm.Model
	// ID     int    `gorm:"primary_key"`
	Tittle      string `gorm:"column:title"` // 外键 (属于), tag `index`是为该列创建索引
	Content     string `json:"content" gorm:"column:content"`
	FrontOrBack bool   `json:"frontorback" gorm:"column:front_or_back"`
}

// mathcad的文章
type MathArticle struct {
	gorm.Model
	UserTempleID uint
	Title        string `json:"title" gorm:"column:title;size:20"`
	Subtext      string `json:"subtext" gorm:"column:subtext;size:20"`
	Content      string `json:"html" gorm:"column:content;size:5000"`
	// UserTemple   UserTemple
}
