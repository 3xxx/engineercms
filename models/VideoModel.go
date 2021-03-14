package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// _ "github.com/mattn/go-sqlite3"
	// "strconv"
	// "strings"
	"time"
)

// 视频表
type Video struct {
	Id int64
	// Code         string `orm:"null"`                                              //编号
	Name     string `form:"title;text;title:",valid:"MinSize(1);MaxSize(20)"` //orm:"unique",
	Url      string `orm:"null"`                                              //关键字
	CoverUrl string `orm:"null"`
	UserId   int64  `orm:"null"`
	// Principal    string `orm:"null"`       //提供人 负责人id
	ProjectId int64 `orm:"null"` //侧栏id
	// TopProjectId int64  `orm:"default(0)"` //项目id
	Content string    `orm:"sie(5000)"` //内容
	Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Video)) //, new(Article)
}

//查询返回的视频表
type UserVideo struct {
	// gorm.Model
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Url      string `json:"url"`
	CoverUrl string `json:"coverurl"`
	Content  string `json:"content"`
	// Created time.Time `orm:"auto_now_add;type(datetime)"`
	Updated time.Time `json:"updated"`
	// DeletedAt *time.Time
	UserId       int64  `json:"userid"`
	ProjectId    int64  `json:"projectid"`
	TopProjectId int64  `json:"topprojectid"`
	UserNickname string `json:"usernickname"`
	ProductTitle string `json:"producttitle"`
}

func CreateVideo(projectid, userid int64, content, name, url string) (id int64, err error) {
	o := orm.NewOrm()
	//查询数据库中有无
	// var video Video
	video := &Video{
		Name:      name,
		Url:       url,
		ProjectId: projectid,
		UserId:    userid,
		Content:   content,
		Created:   time.Now(),
		Updated:   time.Now(),
	}
	id, err = o.Insert(video)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func UpdateVideo(vid int64, coverurl string) error {
	db := GetDB()
	err := db.Table("video").Where("id = ?", vid).Update("cover_url", coverurl).Error
	return err
}

//查询所有视频
func GetUserVideo(pid int64, limit, offset int, searchText string) (uservideos []UserVideo, err error) {
	//获取DB Where("product.title LIKE ?", "%searchText%").不对
	db := GetDB()
	if searchText != "" {
		err = db.Order("video.updated desc").Table("video").Select("video.id,video.name,video.url,video.cover_url,video.content,video.user_id,video.project_id,video.updated,user.nickname as user_nickname, project.title as project_title").Where("project_id=?", pid).
			Joins("left JOIN user on user.id = video.user_id").
			Joins("left join project on project.id = video.project_id").
			Limit(limit).Offset(offset).Scan(&uservideos).Error
	} else { //和上面一样，没有做关键字检索！！！！！
		err = db.Order("video.updated desc").Table("video").Select("video.id,video.name,video.url,video.cover_url,video.content,video.user_id,video.project_id,video.updated,user.nickname as user_nickname, project.title as project_title").Where("project_id=?", pid).
			Joins("left JOIN user on user.id = video.user_id").
			Joins("left join project on project.id = video.project_id").
			Limit(limit).Offset(offset).Scan(&uservideos).Error
	}
	return uservideos, err
	// 多连接及参数
	// db.Joins("JOIN pays ON pays.user_id = users.id", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("user_id = ?", uid).Find(&pays)
}

//查询某个用户借阅记录总数
func GetUserVideoCount(pid int64, searchText string) (count int64, err error) {
	//获取DB
	db := GetDB()
	if searchText != "" {
		err = db.Table("video").Where("project_id=?", pid).
			Count(&count).Error
	} else {
		err = db.Table("video").Where("project_id=?", pid).
			Count(&count).Error
	}
	return count, err
}

//查询一个video
func GetVideobyId(id int64) (video Video, err error) {
	//获取DB
	db := GetDB()
	err = db.Where("id = ?", id).Find(&video).Error
	return video, err
}

// 删除
func Deletevideo(id int64) error {
	//获取DB
	db := GetDB()
	err := db.Where("id = ?", id).Delete(Video{}).Error
	return err
}
