//提交给merit的成果记录表
package models

import (
	// beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
	// _ "github.com/mattn/go-sqlite3"
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

//附件链接表
type CatalogLink struct {
	Id        int64
	CatalogId int64
	Url       string    `orm:"sie(500)"`
	Created   time.Time `orm:"auto_now_add;type(datetime)"`
	Updated   time.Time `orm:"auto_now_add;type(datetime)"`
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
	EcmsIp     string    //用户engineercms的ip地址
	EcmsPort   string    //用户engineercms的端口
	Createtime time.Time `orm:"type(datetime);auto_now_add" `
	Updated    time.Time `orm:"type(datetime);auto_now_add" `
}

func init() {
	orm.RegisterModel(new(PostMerit), new(CatalogLink), new(MeritBasic)) //, new(Article)
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

//删除成果列表
func DeletePostMerit(cid int64) error {
	o := orm.NewOrm()
	product := &PostMerit{Id: cid}
	if o.Read(product) == nil {
		_, err := o.Delete(product)
		if err != nil {
			return err
		}
	}
	return nil
}

//添加成果列表
func AddPostMerit(catalog PostMerit) (cid int64, err error, news string) {
	// orm := orm.NewOrm()
	// fmt.Println(user)
	//判断重复性
	o := orm.NewOrm()
	var catalog1 PostMerit
	//保证成果的唯一性
	//出差必须在成果名称中写入自己的名字以示区别
	//Filter("Drawn", catalog.Drawn).Filter("Designd", catalog.Designd).Filter("Checked", catalog.Checked).
	err = o.QueryTable("PostMerit").Filter("ProjectNumber", catalog.ProjectNumber).Filter("ProjectName", catalog.ProjectName).Filter("Tnumber", catalog.Tnumber).Filter("Name", catalog.Name).One(&catalog1)
	if err == orm.ErrNoRows {
		cid, err1 := o.Insert(&catalog) //_, err = o.Insert(reply)
		if err1 != nil {
			return 0, err1, "insert err"
		} else {
			return cid, nil, "insert ok"
		}
		// fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		return 0, err, "找不到主键"
		//     fmt.Println("找不到主键")
	} else {
		return catalog1.Id, nil, "数据已存在"
	}
}

//用户修改一条成果的某个字段
func ModifyCatalog(cid int64, fieldname, value string) error {
	o := orm.NewOrm()
	var catalog PostMerit
	// catalog := &Catalog{Id: cid}
	err := o.QueryTable("PostMerit").Filter("Id", cid).One(&catalog)
	// err:=o.Read(catalog).One()
	if err == nil {
		// type Duration int64
		// const (
		// 	Nanosecond  Duration = 1
		// 	Microsecond          = 1000 * Nanosecond
		// 	Millisecond          = 1000 * Microsecond
		// 	Second               = 1000 * Millisecond
		// 	Minute               = 60 * Second
		// 	Hour                 = 60 * Minute
		// )
		// const lll = "2006-01-02"
		catalog.Updated = time.Now() //.Add(+time.Duration(hours) * time.Hour)
		switch fieldname {
		case "ProjectNumber":
			catalog.ProjectNumber = value
			_, err := o.Update(&catalog, "ProjectNumber", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "ProjectName":
			catalog.ProjectName = value
			_, err := o.Update(&catalog, "ProjectName", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "DesignStage":
			catalog.DesignStage = value
			_, err := o.Update(&catalog, "DesignStage", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Section":
			catalog.Section = value
			_, err := o.Update(&catalog, "Section", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Tnumber":
			catalog.Tnumber = value
			_, err := o.Update(&catalog, "Tnumber", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Name":
			catalog.Name = value
			_, err := o.Update(&catalog, "Name", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Category":
			catalog.Category = value
			_, err := o.Update(&catalog, "Category", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Count":
			//转成float64
			catalog.Count, err = strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			_, err := o.Update(&catalog, "Count", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Drawn":
			catalog.Drawn = value
			_, err := o.Update(&catalog, "Drawn", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Designd":
			catalog.Designd = value
			_, err := o.Update(&catalog, "Designd", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Checked":
			catalog.Checked = value
			_, err := o.Update(&catalog, "Checked", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Examined":
			catalog.Examined = value
			_, err := o.Update(&catalog, "Examined", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Drawnratio":
			catalog.Drawnratio, err = strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			_, err := o.Update(&catalog, "Drawnratio", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Designdratio":
			catalog.Designdratio, err = strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			_, err := o.Update(&catalog, "Designdratio", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Checkedratio":
			catalog.Checkedratio, err = strconv.ParseFloat(value, 64)
			if err != nil {
				return err
			}
			_, err := o.Update(&catalog, "Checkedratio", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Examinedratio":
			catalog.Examinedratio, err = strconv.ParseFloat(value, 64)
			_, err := o.Update(&catalog, "Examinedratio", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Complex":
			catalog.Complex, err = strconv.ParseFloat(value, 64)
			_, err := o.Update(&catalog, "Complex", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Datestring":
			const lll = "2006-01-02" //"2006-01-02 15:04:05" //12-19-2015 22:40:24
			catalog.Date, err = time.Parse(lll, value)
			if err != nil {
				return err
			}
			catalog.Datestring = value
			_, err := o.Update(&catalog, "Datestring", "Date", "Updated") //这里不能用&catalog
			if err != nil {
				return err
			} else {
				return nil
			}
		}
		// 指定多个字段
		// o.Update(&user, "Field1", "Field2", ...)这个试验没成功
	} else {
		return o.Read(&catalog)
	}
	return nil
}

//添加附件链接表
func AddCatalogLink(cid int64, link string) (id int64, err error) {
	o := orm.NewOrm()
	cataloglink := &CatalogLink{
		CatalogId: cid,
		Url:       link,
		Created:   time.Now(),
		Updated:   time.Now(),
	}
	id, err = o.Insert(cataloglink)
	if err != nil {
		return id, err //如果文章编号相同，则唯一性检查错误，返回id吗？
	}
	return id, err
}

//根据成果id查出附件链接表
func GetCatalogLinks(cid int64) (links []*CatalogLink, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("CatalogLink")
	_, err = qs.Filter("CatalogId", cid).All(&links)
	if err != nil {
		return nil, err
	}
	return links, err
}

//修改links
func ModifyCatalogLink(id, cid int64, fieldname, value string) error {
	o := orm.NewOrm()
	var cataloglink CatalogLink
	// catalog := &Catalog{Id: cid}
	err := o.QueryTable("CatalogLink").Filter("Id", id).One(&cataloglink)
	// err:=o.Read(catalog).One()
	if err == nil {
		cataloglink.Updated = time.Now() //.Add(+time.Duration(hours) * time.Hour)
		cataloglink.Url = value
		_, err := o.Update(&cataloglink, "Url", "Updated")
		if err != nil {
			return err
		} else {
			return nil
		}
	} else {
		cataloglink := &CatalogLink{
			CatalogId: cid,
			Url:       value,
			Created:   time.Now(),
			Updated:   time.Now(),
		}
		_, err = o.Insert(cataloglink)
		if err != nil {
			return err
		}
	}
	return nil
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

	if err == nil {
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
			_, err := o.Update(&merit, "Ip", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "Port":
			merit.Port = value
			_, err := o.Update(&merit, "Port", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "EcmsIp":
			merit.EcmsIp = value
			_, err := o.Update(&merit, "EcmsIp", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		case "EcmsPort":
			merit.EcmsPort = value
			_, err := o.Update(&merit, "EcmsPort", "Updated")
			if err != nil {
				return err
			} else {
				return nil
			}
		}
		// 指定多个字段
		// o.Update(&user, "Field1", "Field2", ...)这个试验没成功
	} else {
		switch fieldname {
		case "Username":
			merit.Username = value
		case "Nickname":
			merit.Nickname = value
		case "Password":
			//这里要加密
			md5Ctx := md5.New()
			md5Ctx.Write([]byte(value))
			cipherStr := md5Ctx.Sum(nil)
			merit.Password = hex.EncodeToString(cipherStr)
		case "Ip":
			merit.Ip = value
		case "Port":
			merit.Port = value
		case "EcmsIp":
			merit.EcmsIp = value
		case "EcmsPort":
			merit.EcmsPort = value
		}
		_, err := o.Insert(&merit)
		if err != nil {
			return err
		}
		// return o.Read(&merit)
	}
	return nil
}
