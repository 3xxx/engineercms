package controllers

import (
	"crypto/md5"
	"encoding/hex"
	// beego "github.com/beego/beego/v2/adapter"
	// "github.com/bitly/go-simplejson"
	"encoding/json"
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

// type Userselect struct { //
// 	Id   int64  `json:"id"`
// 	Name string `json:"text"`
// }

type RegistController struct {
	web.Controller
}

func (this *RegistController) Get() {
	u := this.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		this.TplName = "mobile/mregister.tpl"
	} else {
		// beego.Info("电脑端！")
		this.TplName = "regist.tpl"
	}
}

func (this *RegistController) RegistErr() {
	this.TplName = "registerr.tpl"
}

func (this *RegistController) CheckUname() {
	var user models.User //这里修改
	//fmt.Println(inputs)
	user.Username = this.GetString("uname")
	err := models.CheckUname(user) //这里修改
	if err == nil {
		this.Ctx.WriteString("false")
		// return false
	} else {
		this.Ctx.WriteString("true")
		// return true
	}
	// return
}

//提交注册名称
func (this *RegistController) Post() {
	var user models.User //这里修改
	// inputs := this.Input()
	// beego.Info(inputs)
	user.Username = this.GetString("uname")
	user.Email = this.GetString("email")
	user.Nickname = this.GetString("nickname")
	Pwd1 := this.GetString("pwd")

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(Pwd1))
	cipherStr := md5Ctx.Sum(nil)
	// fmt.Print(cipherStr)
	// fmt.Print("\n")
	// fmt.Print(hex.EncodeToString(cipherStr))
	user.Password = hex.EncodeToString(cipherStr)
	user.Lastlogintime = time.Now()
	user.Status = 1
	user.Role = "4"
	_, err2 := models.SaveUser(user) //这里修改

	// _, err = models.AddRoleUser(4, uid)
	u := this.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		if err2 == nil {
			// this.TplName = "success.tpl"
			// this.Data["json"] = map[string]interface{}{"islogin": 0}
			// this.ServeJSON()
			this.Redirect("/login", 301)
		} else {
			// fmt.Println(err)
			// this.TplName = "registerr.tpl"
			// this.Data["json"] = map[string]interface{}{"islogin": 1}
			// this.ServeJSON()
			this.Redirect("/regist", 301)
		}
	} else {
		// beego.Info("电脑端！")
		if err2 == nil {
			// this.TplName = "success.tpl"
			this.Redirect("/login", 301)
		} else {
			// fmt.Println(err)
			// this.TplName = "registerr.tpl"
			this.Redirect("/regist", 301)
		}
	}

}

// @Title post wx regist
// @Description post wx regist
// @Param uname query string  true "The username of ueser"
// @Param password query string  true "The password of account"
// @Param code query string  true "The code of wx"
// @Param app_version query string  true "The app_version of wx"
// @Success 200 {object} models.SaveUser
// @Failure 400 Invalid page supplied
// @Failure 404 user not found
// @router /wxregist [post]
// 微信小程序根据用户输入的用户名和密码后，将openid存到用户名对应的数据表中，
// 方便下次自动根据openid匹配用户后自动登录wxlogin
// 这个不是真正的注册。另见WxRegion()
func (c *RegistController) WxRegist() {
	var user models.User
	var uid string
	var isAdmin bool
	user.Username = c.GetString("uname")
	Pwd1 := c.GetString("password") //注意这里用的是全称password，不是pwd
	// autoLogin := c.GetString("autoLogin") == "on"
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(Pwd1))
	cipherStr := md5Ctx.Sum(nil)
	user.Password = hex.EncodeToString(cipherStr)
	logs.Info(user.Password)
	logs.Info(user.Username)
	err := models.ValidateUser(user)
	if err == nil {
		JSCODE := c.GetString("code")
		// beego.Info(JSCODE)
		var APPID, SECRET string
		app_version := c.GetString("app_version")
		if app_version == "1" {
			APPID, err = web.AppConfig.String("wxAPPID")
			if err != nil {
				logs.Error(err)
			}
			SECRET, err = web.AppConfig.String("wxSECRET")
			if err != nil {
				logs.Error(err)
			}
		} else {
			appstring := "wxAPPID" + app_version
			APPID, err = web.AppConfig.String(appstring)
			if err != nil {
				logs.Error(err)
			}
			secretstring := "wxSECRET" + app_version
			SECRET, err = web.AppConfig.String(secretstring)
			if err != nil {
				logs.Error(err)
			}
		}
		requestUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=" + APPID + "&secret=" + SECRET + "&js_code=" + JSCODE + "&grant_type=authorization_code"
		resp, err := http.Get(requestUrl)
		if err != nil {
			logs.Error(err)
			// return
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			logs.Error(err)
		}
		var data map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			logs.Error(err)
		}
		// beego.Info(data)
		var openID string
		// var sessionKey string
		if _, ok := data["session_key"]; !ok {
			errcode := data["errcode"]
			// beego.Info(errcode)
			errmsg := data["errmsg"].(string)
			c.Data["json"] = map[string]interface{}{"errNo": errcode, "msg": errmsg, "data": "session_key 不存在"}
			c.ServeJSON()
		} else {
			openID = data["openid"].(string)
			// beego.Info(openID)
			//将openid写入数据库
			user, err = models.GetUserByUsername(user.Username)
			if err != nil {
				logs.Error(err)
			} else {
				uid = strconv.FormatInt(user.Id, 10)
			}
			_, err = models.AddUserOpenID(user.Id, openID)
			if err != nil {
				logs.Error(err)
			}
			//根据userid取出user和avatorUrl
			useravatar, err := models.GetUserAvatorUrl(user.Id)
			if err != nil {
				logs.Error(err)
			}
			var photo string
			wxsite, err := web.AppConfig.String("wxreqeustsite")
			if err != nil {
				logs.Error(err)
			}
			if len(useravatar) != 0 {
				photo = wxsite + useravatar[0].UserAvatar.AvatarUrl
				// beego.Info(photo)
			}
			//根据userid取出appreciation赞赏码
			userappreciation, err := models.GetUserAppreciationUrl(user.Id)
			if err != nil {
				logs.Error(err)
			}
			var appreciationphoto string
			if len(userappreciation) != 0 {
				appreciationphoto = wxsite + userappreciation[0].UserAppreciation.AppreciationUrl
			}
			// roles, err := models.GetRolenameByUserId(user.Id)
			// if err != nil {
			// 	logs.Error(err)
			// }
			// var isAdmin bool
			// for _, v := range roles {
			// 	if v.Rolename == "admin" {
			// 		isAdmin = true
			// 	}
			// }
			//bug:下面这句是取得角色admin，如果用户未设置角色admin，则出错20191224。
			//因此，必须有系统默认设置admin角色才行。
			role, err := models.GetRoleByRolename("admin")
			if err != nil {
				logs.Error(err)
			} else {
				roleid := strconv.FormatInt(role.Id, 10)
				isAdmin, err = e.HasRoleForUser(uid, "role_"+roleid)
				if err != nil {
					logs.Error(err)
				}
			}

			//用户登录后，存储openid在服务端的session里，下次用户通过hotqinsessionid来取得openid
			c.SetSession("openID", openID)
			// c.SetSession("sessionKey", sessionKey) //这个没用
			// https://blog.csdn.net/yelin042/article/details/71773636
			// c.SetSession("skey", skey) //这个没用
			//为何有这个hotqinsessionid?
			//用户小程序register后，只是存入服务器数据库中的openid和用户名对应
			//用户小程序login的时候，即这里，将openid存入session
			//下次用户请求携带hotqinsessionid即可取到session-openid了。
			sessionId := c.Ctx.Input.Cookie("hotqinsessionid") //这一步什么意思

			if err != nil {
				logs.Error(err)
				c.Data["json"] = map[string]interface{}{"info": "已经注册", "data": "已经注册"}
				c.ServeJSON()
			} else {
				c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "userId": uid, "userName": user.Username, "userNickname": user.Nickname, "isAdmin": isAdmin, "sessionId": sessionId, "photo": photo, "appreciationphoto": appreciationphoto}
				c.ServeJSON()
			}
		}
	} else {
		c.Data["json"] = map[string]interface{}{"info": "用户名或密码错误", "data": ""}
		c.ServeJSON()
	}
}

// @Title post wx region
// @Description post wx region
// @Param uname query string  true "The username of user"
// @Param password query string  true "The password of account"
// @Param code query string  true "The code of wx"
// @Param app_version query string  true "The app_version of wx"
// @Success 200 {object} models.SaveUser
// @Failure 400 Invalid page supplied
// @Failure 404 user not found
// @router /wxregion [post]
// 将用户名和密码存入数据表，openid也存入数据表
func (c *RegistController) WxRegion() {
	var user models.User
	// var uid string
	// var isAdmin bool
	isAdmin := false
	user.Username = c.GetString("uname")
	Pwd1 := c.GetString("password") //注意这里用的是全称password，不是pwd
	// autoLogin := c.GetString("autoLogin") == "on"
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(Pwd1))
	cipherStr := md5Ctx.Sum(nil)
	user.Password = hex.EncodeToString(cipherStr)
	// beego.Info(user.Password)
	// beego.Info(user.Username)
	err := models.CheckUname(user)
	if err != nil {
		logs.Error(err)
	} else { //用户存在
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "用户名已存在！"}
		c.ServeJSON()
		return
	}
	uid, err := models.SaveUser(user)
	if err == nil {
		// JSCODE := c.GetString("code")
		openID := c.GetString("openID")
		// beego.Info(JSCODE)
		// var APPID, SECRET string
		// app_version := c.GetString("app_version")
		// if app_version == "1" {
		// 	APPID = web.AppConfig.String("wxAPPID")
		// 	SECRET = web.AppConfig.String("wxSECRET")
		// } else {
		// 	appstring := "wxAPPID" + app_version
		// 	APPID = web.AppConfig.String(appstring)
		// 	secretstring := "wxSECRET" + app_version
		// 	SECRET = web.AppConfig.String(secretstring)
		// }
		// requestUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=" + APPID + "&secret=" + SECRET + "&js_code=" + JSCODE + "&grant_type=authorization_code"
		// resp, err := http.Get(requestUrl)
		// beego.Info(resp)
		// if err != nil {
		// 	logs.Error(err)
		// 	// return
		// }
		// defer resp.Body.Close()
		// if resp.StatusCode != 200 {
		// 	logs.Error(err)
		// }
		// var data map[string]interface{}
		// err = json.NewDecoder(resp.Body).Decode(&data)
		// if err != nil {
		// 	logs.Error(err)
		// }
		// beego.Info(data)
		// // var sessionKey string
		// if _, ok := data["session_key"]; !ok {
		// 	errcode := data["errcode"]
		// 	beego.Info(errcode)
		// 	errmsg := data["errmsg"].(string)
		// 	c.Data["json"] = map[string]interface{}{"info": "ERROR", "errNo": errcode, "msg": errmsg, "data": "session_key 不存在"}
		// 	c.ServeJSON()
		// } else {
		// openID = data["openid"].(string)
		// beego.Info(openID)
		//将openid写入数据库
		// user, err = models.GetUserByUsername(user.Username)
		// if err != nil {
		// 	logs.Error(err)
		// } else {
		// 	uid = strconv.FormatInt(user.Id, 10)
		// }
		_, err = models.AddUserOpenID(uid, openID)
		if err != nil {
			logs.Error(err)
		}
		//根据userid取出user和avatorUrl
		useravatar, err := models.GetUserAvatorUrl(user.Id)
		if err != nil {
			logs.Error(err)
		}
		var photo string
		wxsite, err := web.AppConfig.String("wxreqeustsite")
		if err != nil {
			logs.Error(err)
		}
		if len(useravatar) != 0 {
			photo = wxsite + useravatar[0].UserAvatar.AvatarUrl
			// beego.Info(photo)
		}
		//根据userid取出appreciation赞赏码
		userappreciation, err := models.GetUserAppreciationUrl(user.Id)
		if err != nil {
			logs.Error(err)
		}
		var appreciationphoto string
		if len(userappreciation) != 0 {
			appreciationphoto = wxsite + userappreciation[0].UserAppreciation.AppreciationUrl
		}
		// roles, err := models.GetRolenameByUserId(user.Id)
		// if err != nil {
		// 	logs.Error(err)
		// }
		// var isAdmin bool
		// for _, v := range roles {
		// 	if v.Rolename == "admin" {
		// 		isAdmin = true
		// 	}
		// }
		//bug:下面这句是取得角色admin，如果用户未设置角色admin，则出错20191224。
		//因此，必须有系统默认设置admin角色才行。
		// role, err := models.GetRoleByRolename("admin")
		// if err != nil {
		// 	logs.Error(err)
		// } else {
		// 	roleid := strconv.FormatInt(role.Id, 10)
		// 	isAdmin = e.HasRoleForUser(uid, "role_"+roleid)
		// }

		//用户登录后，存储openid在服务端的session里，下次用户通过hotqinsessionid来取得openid
		c.SetSession("openID", openID)
		// c.SetSession("sessionKey", sessionKey) //这个没用
		// https://blog.csdn.net/yelin042/article/details/71773636
		// c.SetSession("skey", skey) //这个没用
		//为何有这个hotqinsessionid?
		//用户小程序register后，只是存入服务器数据库中的openid和用户名对应
		//用户小程序login的时候，即这里，将openid存入session
		//下次用户请求携带hotqinsessionid即可取到session-openid（key字段）了。
		//这一步是取得session名称为hotqinsessionid的session，前端请求携带这个后，服务端可用session中的key字段（如uname、openID）取得值
		sessionId := c.Ctx.Input.Cookie("hotqinsessionid")

		if err != nil {
			logs.Error(err)
			c.Data["json"] = map[string]interface{}{"info": "已经注册", "data": "已经注册"}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "userId": uid, "isAdmin": isAdmin, "sessionId": sessionId, "photo": photo, "appreciationphoto": appreciationphoto}
			c.ServeJSON()
		}
		// }
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "保存用户出错！"}
		c.ServeJSON()
	}
}

//post方法
func (this *RegistController) GetUname() {
	var user models.User //这里修改[]*models.User(uname string)
	//fmt.Println(inputs)
	user.Username = this.GetString("uname")
	// beego.Info(user.Username)
	uname1, err := models.GetUname(user) //这里修改
	//转换成json数据？
	// beego.Info(uname1[0].Username)
	// b, err := json.Marshal(uname1)
	if err == nil {
		// this.Ctx.WriteString(string(b))
		this.Data["json"] = uname1 //string(b)
		this.ServeJSON()
	}
	// 	this.Ctx.WriteString(uname1[1].Username)
	// 	// return uname1[0].Username
	// }
	// return uname1[0].Username
}

//get方法，用于x-editable的select2方法
func (this *RegistController) GetUname1() {
	var user models.User //这里修改[]*models.User(uname string)
	//fmt.Println(inputs)
	user.Username = this.GetString("uname")
	// beego.Info(user.Username)
	uname1, err := models.GetUname(user) //这里修改
	//转换成json数据？
	// beego.Info(uname1[0].Username)
	// b, err := json.Marshal(uname1)
	if err != nil {
		logs.Error(err)
	}
	slice1 := make([]Userselect, 0)
	for _, v := range uname1 {
		aa := make([]Userselect, 1)
		aa[0].Id = v.Id //这里用for i1,v1,然后用v1.Id一样的意思
		// aa[0].Ad = v.Id
		aa[0].Name = v.Username
		slice1 = append(slice1, aa...)
	}
	// b, err := json.Marshal(slice1) //不需要转成json格式
	// beego.Info(string(b))
	// fmt.Println(string(b))
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(uname1) //[0xc08214b880 0xc08214b960 0xc08214ba40
	// beego.Info(slice1) //[{1471  admin} {1475  cai.wc} {1476  zeng.cw}
	// this.Data["Userselect"] = slice1
	this.Data["json"] = slice1 //string(b)
	this.ServeJSON()
	// this.TplName = "loginerr.html"
	// 	this.Ctx.WriteString(uname1[1].Username)
	// 	// return uname1[0].Username
	// }
	// return uname1[0].Username
	//结果如下：
	// [
	//  {
	//    "Id": 1471,
	//    "id": "",
	//    "text": "admin"
	//  },
	//  {
	//    "Id": 1475,
	//    "id": "",
	//    "text": "cai.wc"
	//  },
	//  {
	//    "Id": 1476,
	//    "id": "",
	//    "text": "zeng.cw"
	//  },
}
