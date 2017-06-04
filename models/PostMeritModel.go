//提交给merit的成果记录表
package models

import (
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	// "strings"
	"crypto/md5"
	"encoding/hex"
	"time"
)

type PostMerit struct {
	Id            int64     `json:"id"`
	ProjectNumber string    //项目编号
	ProjectName   string    //项目名称
	DesignStage   string    //阶段
	Section       string    //专业
	Tnumber       string    //成果编号
	Name          string    //成果名称
	Category      string    //成果类型
	Page          string    //成果计量单位
	Count         float64   //成果数量
	Drawn         string    //编制、绘制
	Designd       string    //设计
	Checked       string    //校核
	Examined      string    //审查
	Verified      string    //核定
	Approved      string    //批准
	Complex       float64   //难度系数
	Drawnratio    float64   //编制、绘制占比系数
	Designdratio  float64   //设计系数
	Checkedratio  float64   //校核系数
	Examinedratio float64   //审查系数
	Datestring    string    //保存字符型日期
	Date          time.Time `orm:"null;auto_now_add;type(datetime)"`
	Created       time.Time `orm:"auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"auto_now_add;type(datetime)"`
	Author        string    //上传者
	State         int       //1编写状态，未提交；2编写者提交，
}

//用户表
type MeritBasic struct {
	Id         int64  `PK`
	Username   string `orm:"unique"` //这个拼音的简写
	Nickname   string //中文名，注意这里，很多都要查询中文名才行`orm:"unique;size(32)" form:"Nickname" valid:"Required;MaxSize(20);MinSize(2)"`
	Password   string
	Repassword string `orm:"-" form:"Repassword" valid:"Required" form:"-"`
	Ip         string //ip地址
	Port       string
	Createtime time.Time `orm:"type(datetime);auto_now_add" `
	Updated    time.Time `orm:"type(datetime);auto_now_add" `
}

func init() {
	orm.RegisterModel(new(PostMerit), new(MeritBasic)) //, new(Article)
}

//修改成果信息
// func UpdatePostMerit(cid int64, code, title, label, principal string) error {
// 	o := orm.NewOrm()
// 	product := &Product{Id: cid}
// 	if o.Read(product) == nil {
// 		product.Code = code
// 		product.Title = title
// 		product.Label = label
// 		product.Principal = principal
// 		product.Updated = time.Now()
// 		_, err := o.Update(product)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

//删除_把附件也一并删除（在controllers中实现吧）
func DeletePostMerit(cid int64) error {
	o := orm.NewOrm()
	product := &Product{Id: cid}
	if o.Read(product) == nil {
		_, err := o.Delete(product)
		if err != nil {
			return err
		}
	}
	return nil
}

//
func AddPostMerit(catalog PostMerit) (cid int64, err error, news string) {
	// orm := orm.NewOrm()
	// fmt.Println(user)
	//判断重复性
	o := orm.NewOrm()
	var catalog1 PostMerit
	//保证成果的唯一性
	//出差必须在成果名称中写入自己的名字以示区别
	//Filter("Drawn", catalog.Drawn).Filter("Designd", catalog.Designd).Filter("Checked", catalog.Checked).
	err = o.QueryTable("PostMerit").Filter("Tnumber", catalog.Tnumber).Filter("Name", catalog.Name).Filter("Category", catalog.Category).One(&catalog1)
	if err == orm.ErrNoRows {
		cid, err = o.Insert(&catalog) //_, err = o.Insert(reply)
		if err != nil {
			return 0, err, "insert err"
		} else {
			return cid, nil, "insert ok"
		}
		// fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		return 0, err, "找不到主键"
		//     fmt.Println("找不到主键")
	} else {
		return 0, nil, "数据已存在"
	}
}

//用户修改一个用户的某个字段
func UpdatePostMerit(id int64, fieldname, value string) error {
	o := orm.NewOrm()
	var merit PostMerit
	// user := &User{Id: cid}
	err := o.QueryTable("PostMerit").Filter("Id", id).One(&merit)
	// err:=o.Read(user).One()
	if err == nil {
		type Duration int64
		const (
			Nanosecond  Duration = 1
			Microsecond          = 1000 * Nanosecond
			Millisecond          = 1000 * Microsecond
			Second               = 1000 * Millisecond
			Minute               = 60 * Second
			Hour                 = 60 * Minute
		)
		// hours := 8

		const lll = "2006-01-02"
		merit.Updated = time.Now() //.Add(+time.Duration(hours) * time.Hour)
		switch fieldname {
		case "name":
			merit.Name = value
			_, err := o.Update(&merit, "Name", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		// case "Choose":
		// 	merit.Choose = value
		// 	_, err := o.Update(&merit, "Choose", "Updated")
		// 	if err != nil {
		// 		return err
		// 	} else {
		// 		return nil
		// 	}
		// case "Content":
		// 	merit.Content = value
		// 	_, err := o.Update(&merit, "Content", "Updated")
		// 	if err != nil {
		// 		return err
		// 	} else {
		// 		return nil
		// 	}
		case "State":
			merit.State, err = strconv.Atoi(value)
			if err != nil {
				return err
			}
			_, err := o.Update(&merit, "State", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		}
		// 指定多个字段
		// o.Update(&user, "Field1", "Field2", ...)这个试验没成功
	}
	return err
}

//根据侧栏id查出所有成果
// func GetPostMerits(id int64) (products []*Product, err error) {
// 	o := orm.NewOrm()
// 	qs := o.QueryTable("Product")
// 	_, err = qs.Filter("ProjectId", id).OrderBy("-created").All(&products)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return products, err
// }

//查出所有成果
func GetPostMerits(status int) (postmerit []*PostMerit, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("PostMerit") //这个表名AchievementTopic需要用驼峰式，
	_, err = qs.Filter("state", status).All(&postmerit)
	if err != nil {
		return postmerit, err
	}
	return postmerit, err
}

//根据成果id取得成果
func GetPostMerit(id int64) (postmerit PostMerit, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("PostMerit") //这个表名AchievementTopic需要用驼峰式，
	err = qs.Filter("id", id).One(&postmerit)
	if err != nil {
		return postmerit, err
	}
	return postmerit, err
}

//查出merit基本信息
//取到一个数据，不是数组，所以table无法显示
func GetMeritBasic() (meritbasic MeritBasic, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("MeritBasic") //这个表名AchievementTopic需要用驼峰式，
	err = qs.One(&meritbasic)
	if err != nil {
		return meritbasic, err
	}
	return meritbasic, err
}

//用户修改一个用户的某个字段
func UpdateMeritBasic(cid int64, fieldname, value string) error {
	o := orm.NewOrm()
	var merit MeritBasic
	// user := &User{Id: cid}
	err := o.QueryTable("MeritBasic").Filter("Id", cid).One(&merit)
	// err:=o.Read(user).One()
	if err == nil {
		type Duration int64
		const (
			Nanosecond  Duration = 1
			Microsecond          = 1000 * Nanosecond
			Millisecond          = 1000 * Microsecond
			Second               = 1000 * Millisecond
			Minute               = 60 * Second
			Hour                 = 60 * Minute
		)
		// hours := 8
		const lll = "2006-01-02"
		merit.Updated = time.Now() //.Add(+time.Duration(hours) * time.Hour)
		switch fieldname {
		case "Username":
			merit.Username = value
			_, err := o.Update(&merit, "Username", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Nickname":
			merit.Nickname = value
			_, err := o.Update(&merit, "Nickname", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Password":
			//这里要加密
			md5Ctx := md5.New()
			md5Ctx.Write([]byte(value))
			cipherStr := md5Ctx.Sum(nil)
			merit.Password = hex.EncodeToString(cipherStr)
			_, err := o.Update(&merit, "Password", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Ip":
			merit.Ip = value
			_, err := o.Update(&merit, "Ip", "Updated") //这里不能用&user
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Port":
			merit.Port = value
			_, err := o.Update(&merit, "Port", "Updated") //这里不能用&user
			if err != nil {
				return err
			} else {
				return nil
			}
		}
		// 指定多个字段
		// o.Update(&user, "Field1", "Field2", ...)这个试验没成功
	} else {
		return o.Read(&merit)
	}
	return nil
}
