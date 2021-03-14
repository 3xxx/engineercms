package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego"
	// "github.com/bitly/go-simplejson"
	"encoding/json"
	"github.com/3xxx/engineercms/models"
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
	beego.Controller
}

func (this *RegistController) Get() {
	u := this.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		beego.Error(err)
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
	inputs := this.Input()
	//fmt.Println(inputs)
	user.Username = inputs.Get("uname")
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
	user.Username = this.Input().Get("uname")
	user.Email = this.Input().Get("email")
	user.Nickname = this.Input().Get("nickname")
	Pwd1 := this.Input().Get("pwd")

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
		beego.Error(err)
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
			this.TplName = "success.tpl"
		} else {
			// fmt.Println(err)
			this.TplName = "registerr.tpl"
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
//添加微信小程序珠三角设代阅览版注册
func (c *RegistController) WxRegist() {
	var user models.User
	var uid string
	var isAdmin bool
	user.Username = c.Input().Get("uname")
	Pwd1 := c.Input().Get("password") //注意这里用的是全称password，不是pwd
	// autoLogin := c.Input().Get("autoLogin") == "on"
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(Pwd1))
	cipherStr := md5Ctx.Sum(nil)
	user.Password = hex.EncodeToString(cipherStr)
	// beego.Info(user.Password)
	// beego.Info(user.Username)
	err := models.ValidateUser(user)
	if err == nil {
		JSCODE := c.Input().Get("code")
		// beego.Info(JSCODE)
		var APPID, SECRET string
		app_version := c.Input().Get("app_version")
		if app_version == "1" {
			APPID = beego.AppConfig.String("wxAPPID")
			SECRET = beego.AppConfig.String("wxSECRET")
		} else {
			appstring := "wxAPPID" + app_version
			APPID = beego.AppConfig.String(appstring)
			secretstring := "wxSECRET" + app_version
			SECRET = beego.AppConfig.String(secretstring)
		}
		requestUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=" + APPID + "&secret=" + SECRET + "&js_code=" + JSCODE + "&grant_type=authorization_code"
		resp, err := http.Get(requestUrl)
		if err != nil {
			beego.Error(err)
			// return
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			beego.Error(err)
		}
		var data map[string]interface{}
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(data)
		var openID string
		// var sessionKey string
		if _, ok := data["session_key"]; !ok {
			errcode := data["errcode"]
			// beego.Info(errcode)
			errmsg := data["errmsg"].(string)
			c.Data["json"] = map[string]interface{}{"errNo": errcode, "msg": errmsg, "data": "session_key 不存在"}
		} else {
			openID = data["openid"].(string)
			// beego.Info(openID)
			//将openid写入数据库
			user, err = models.GetUserByUsername(user.Username)
			if err != nil {
				beego.Error(err)
			} else {
				uid = strconv.FormatInt(user.Id, 10)
			}
			_, err = models.AddUserOpenID(user.Id, openID)
			if err != nil {
				beego.Error(err)
			}
			//根据userid取出user和avatorUrl
			useravatar, err := models.GetUserAvatorUrl(user.Id)
			if err != nil {
				beego.Error(err)
			}
			var photo string
			wxsite := beego.AppConfig.String("wxreqeustsite")
			if len(useravatar) != 0 {
				photo = wxsite + useravatar[0].UserAvatar.AvatarUrl
				// beego.Info(photo)
			}
			//根据userid取出appreciation赞赏码
			userappreciation, err := models.GetUserAppreciationUrl(user.Id)
			if err != nil {
				beego.Error(err)
			}
			var appreciationphoto string
			if len(userappreciation) != 0 {
				appreciationphoto = wxsite + userappreciation[0].UserAppreciation.AppreciationUrl
			}
			// roles, err := models.GetRolenameByUserId(user.Id)
			// if err != nil {
			// 	beego.Error(err)
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
				beego.Error(err)
			} else {
				roleid := strconv.FormatInt(role.Id, 10)
				isAdmin = e.HasRoleForUser(uid, "role_"+roleid)
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
				beego.Error(err)
				c.Data["json"] = map[string]interface{}{"info": "已经注册", "data": ""}
				c.ServeJSON()
			} else {
				c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "userId": uid, "isAdmin": isAdmin, "sessionId": sessionId, "photo": photo, "appreciationphoto": appreciationphoto}
				c.ServeJSON()
			}
		}
	} else {
		c.Data["json"] = map[string]interface{}{"info": "用户名或密码错误", "data": ""}
		c.ServeJSON()
	}
}

//post方法
func (this *RegistController) GetUname() {
	var user models.User //这里修改[]*models.User(uname string)
	inputs := this.Input()
	//fmt.Println(inputs)
	user.Username = inputs.Get("uname")
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
	inputs := this.Input()
	//fmt.Println(inputs)
	user.Username = inputs.Get("uname")
	// beego.Info(user.Username)
	uname1, err := models.GetUname(user) //这里修改
	//转换成json数据？
	// beego.Info(uname1[0].Username)
	// b, err := json.Marshal(uname1)
	if err != nil {
		beego.Error(err)
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
		beego.Error(err)
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
