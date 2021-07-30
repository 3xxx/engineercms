package controllers

import (
	"crypto/md5"
	"encoding/hex"
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	// "net/url"
	"github.com/3xxx/engineercms/controllers/utils"
	"github.com/3xxx/engineercms/models"
	"strconv"
	// "github.com/astaxie/beego/session"
	"encoding/json"
	// "github.com/casbin/beego-orm-adapter"
	// "github.com/casbin/casbin"
	"net/http"
	// "strings"
	"math/rand"
	"regexp"
	"time"
)

// CMSWX login API
type LoginController struct {
	beego.Controller
}

type SuccessController struct {
	beego.Controller
}

type ServiceValidateController struct {
	beego.Controller
}

// func (c *LoginController) Get() {
// 	isExit := c.Input().Get("exit") == "true"
// 	// secofficeshow?secid=1643&level=3&key=modify
// 	url1 := c.Input().Get("url") //这里不支持这样的url，http://192.168.9.13/login?url=/topic/add?id=955&mid=3
// 	url2 := c.Input().Get("level")
// 	url3 := c.Input().Get("key")
// 	var url string
// 	if url2 == "" {
// 		url = url1
// 	} else {
// 		url = url1 + "&level=" + url2 + "&key=" + url3
// 	}
// 	c.Data["Url"] = url
// 	if isExit {
// 		// c.Ctx.SetCookie("uname", "", -1, "/")
// 		// c.Ctx.SetCookie("pwd", "", -1, "/")
// 		// c.DelSession("gosessionid")
// 		// c.DelSession("uname")//这个不行
// 		// c.Destroy/Session()
// 		// c.Ctx.Input.CruSession.Delete("gosessionid")这句与上面一句重复
// 		// c.Ctx.Input.CruSession.Flush()
// 		// beego.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)
// 		v := c.GetSession("uname")
// 		// islogin := false
// 		if v != nil {
// 			//删除指定的session
// 			c.DelSession("uname")
// 			//销毁全部的session
// 			c.DestroySession()
// 		}
// 		// sess.Flush()//这个不灵
// 		c.Redirect("/", 302)
// 		return
// 	}
// 	c.TplName = "login.tpl"
// }

// func (c *LoginController) Loginerr() {
// 	url1 := c.Input().Get("url") //这里不支持这样的url，http://192.168.9.13/login?url=/topic/add?id=955&mid=3
// 	url2 := c.Input().Get("level")
// 	url3 := c.Input().Get("key")
// 	var url string
// 	if url2 == "" {
// 		url = url1
// 	} else {
// 		url = url1 + "&level=" + url2 + "&key=" + url3
// 	}
// 	// port := strconv.Itoa(c.Ctx.Input.Port())
// 	// url := c.Ctx.Input.Site() + ":" + port + c.Ctx.Request.URL.String()
// 	c.Data["Url"] = url
// 	// beego.Info(url)
// 	c.TplName = "loginerr.tpl"
// }

//登录页面
func (c *LoginController) Login() {
	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		beego.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		c.TplName = "mobile/mlogin.tpl"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "login.tpl"
	}
}

//login页面输入用户名和密码后登陆提交
func (c *LoginController) Post() {
	// uname := c.Input().Get("uname")
	// url := c.Input().Get("returnUrl")
	url1 := c.Input().Get("url") //这里不支持这样的url，http://192.168.9.13/login?url=/topic/add?id=955&mid=3
	url2 := c.Input().Get("level")
	url3 := c.Input().Get("key")
	var url string
	if url2 == "" && url1 != "" {
		url = url1
	} else if url2 != "" {
		url = url1 + "&level=" + url2 + "&key=" + url3
	} else {
		url = c.Input().Get("referrer")
	}
	var user models.User
	user.Username = c.Input().Get("uname")
	Pwd1 := c.Input().Get("pwd")
	// autoLogin := c.Input().Get("autoLogin") == "on"
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(Pwd1))
	cipherStr := md5Ctx.Sum(nil)

	user.Password = hex.EncodeToString(cipherStr)
	err := models.ValidateUser(user)
	if err == nil {
		c.SetSession("uname", user.Username)
		c.SetSession("pwd", user.Password)
		utils.FileLogs.Info(user.Username + " " + "login" + " 成功")
		User, err := models.GetUserByUsername(user.Username)
		if err != nil {
			beego.Error(err)
			utils.FileLogs.Error(user.Username + " 查询用户 " + err.Error())
		}
		if User.Ip == "" {
			err = models.UpdateUser(User.Id, "Ip", c.Ctx.Input.IP())
			if err != nil {
				beego.Error(err)
				utils.FileLogs.Error(user.Username + " 添加用户ip " + err.Error())
			}
		} else {
			//更新user表的lastlogintime
			err = models.UpdateUserlastlogintime(user.Username)
			if err != nil {
				beego.Error(err)
				utils.FileLogs.Error(user.Username + " 更新用户登录时间 " + err.Error())
			}
		}
		if url != "" {
			c.Redirect(url, 302)
		} else {
			var id string
			index := beego.AppConfig.String("redirect")
			navid1 := beego.AppConfig.String("navigationid1")
			navid2 := beego.AppConfig.String("navigationid2")
			navid3 := beego.AppConfig.String("navigationid3")
			navid4 := beego.AppConfig.String("navigationid4")
			navid5 := beego.AppConfig.String("navigationid5")
			navid6 := beego.AppConfig.String("navigationid6")
			navid7 := beego.AppConfig.String("navigationid7")
			navid8 := beego.AppConfig.String("navigationid8")
			navid9 := beego.AppConfig.String("navigationid9")
			// beego.Info(index)
			switch index {
			case "":
				c.Redirect("/index", 302)
			case "IsNav1":
				id = navid1
				c.Redirect("/project/"+id, 302)
			case "IsNav2":
				id = navid2
				// beego.Info(id)
				c.Redirect("/project/"+id, 302)
			case "IsNav3":
				id = navid3
				c.Redirect("/project/"+id, 302)
			case "IsNav4":
				id = navid4
				c.Redirect("/project/"+id, 302)
			case "IsNav5":
				id = navid5
				c.Redirect("/project/"+id, 302)
			case "IsNav6":
				id = navid6
				c.Redirect("/project/"+id, 302)
			case "IsNav7":
				id = navid7
				c.Redirect("/project/"+id, 302)
			case "IsNav8":
				id = navid8
				c.Redirect("/project/"+id, 302)
			case "IsNav9":
				id = navid9
				c.Redirect("/project/"+id, 302)
			case "IsProject":
				c.Redirect("/project", 302)
			case "IsOnlyOffice":
				c.Redirect("/onlyoffice", 302)
			case "IsDesignGant", "IsConstructGant":
				c.Redirect("/projectgant", 302)
			default:
				c.Redirect("/index", 302)
			}
		}
	} else {
		c.Redirect("/loginerr?url="+url, 302)
	}
	return
}

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(GetRandomSalt()))
}

// return len=8  salt
func GetRandomSalt() []byte {
	return GetRandomString(8)
}

//生成随机字符串
func GetRandomString(length int) []byte {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return result
}

// @Title post user login...
// @Description post login..
// @Param uname query string  true "The name of user"
// @Param pwd query string  true "The password of user"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /loginpost [post]
//login弹框输入用户名和密码后登陆提交//微信用户注册登录用register
func (c *LoginController) LoginPost() {
	var user models.User
	user.Username = c.Input().Get("uname")
	// beego.Info(user.Username)
	// uname := c.GetString("uname")
	// Pwd1 := c.GetString("pwd")
	Pwd1 := c.Input().Get("pwd")
	// beego.Info(Pwd1)
	// autoLogin := c.Input().Get("autoLogin") == "on"
	islogin := 0
	// maxAge := 0
	// if autoLogin {
	// 	maxAge = 1<<31 - 1
	// }
	// c.Ctx.SetCookie("uname", uname, maxAge, "/")
	// c.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(Pwd1))
	cipherStr := md5Ctx.Sum(nil)
	user.Password = hex.EncodeToString(cipherStr)
	// beego.Info(user.Password)
	// beego.Info(islogin)
	err := models.ValidateUser(user)
	if err == nil {
		c.SetSession("uname", user.Username)
		c.SetSession("pwd", user.Password) //这个没用
		utils.FileLogs.Info(user.Username + " " + "login" + " 成功")
		User, err := models.GetUserByUsername(user.Username)
		if err != nil {
			beego.Error(err)
			utils.FileLogs.Error(user.Username + " 查询用户 " + err.Error())
		}
		if User.Ip == "" {
			err = models.UpdateUser(User.Id, "Ip", c.Ctx.Input.IP())
			if err != nil {
				beego.Error(err)
				utils.FileLogs.Error(user.Username + " 添加用户ip " + err.Error())
			}
		} else {
			//更新user表的lastlogintime
			err = models.UpdateUserlastlogintime(user.Username)
			if err != nil {
				beego.Error(err)
				utils.FileLogs.Error(user.Username + " 更新用户登录时间 " + err.Error())
			}
		}
		// beego.Info(islogin)
	} else {
		islogin = 1
	}
	// beego.Info(islogin)
	// if name == "admin" && pwd == "123456" {
	// 	c.SetSession("loginuser", "adminuser")
	// 	fmt.Println("当前的session:")
	// 	fmt.Println(c.CruSession)
	c.Data["json"] = map[string]interface{}{"islogin": islogin}
	c.ServeJSON()
}

//退出登录
func (c *LoginController) Logout() {
	v := c.GetSession("uname")
	if v != nil {
		//删除指定的session
		c.DelSession("uname")
		c.DelSession("uid") //删除mindoc的用户登录信息
		//销毁全部的session
		// c.DestroySession()
	}
	islogin := false
	c.Data["json"] = map[string]interface{}{"islogin": islogin}
	c.ServeJSON()
}

//作废20180915
func (c *LoginController) Loginerr() {
	url1 := c.Input().Get("url") //这里不支持这样的url，http://192.168.9.13/login?url=/topic/add?id=955&mid=3
	url2 := c.Input().Get("level")
	url3 := c.Input().Get("key")
	var url string
	if url2 == "" {
		url = url1
	} else {
		url = url1 + "&level=" + url2 + "&key=" + url3
	}
	// port := strconv.Itoa(c.Ctx.Input.Port())
	// url := c.Ctx.Input.Site() + ":" + port + c.Ctx.Request.URL.String()
	c.Data["Url"] = url
	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		beego.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		c.TplName = "mobile/mloginerr.tpl"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "loginerr.tpl"
	}

}

// @Title post wx login
// @Description post wx login
// @Param id path string  true "The id of wx"
// @Param code path string  true "The jscode of wxuser"
// @Success 200 {object} success
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /wxlogin/:id [get]
//微信小程序访问微信服务器获取用户信息
func (c *LoginController) WxLogin() {
	id := c.Ctx.Input.Param(":id")
	JSCODE := c.Input().Get("code")
	// beego.Info(JSCODE)
	var APPID, SECRET string
	if id == "1" {
		APPID = "wx7f77b90a1a891d93"
		SECRET = "f58ca4f28cbb52ccd805d66118060449"
	} else {
		appstring := "wxAPPID" + id
		APPID = beego.AppConfig.String(appstring)
		secretstring := "wxSECRET" + id
		SECRET = beego.AppConfig.String(secretstring)
		// beego.Info(APPID)
	}
	//这里用security.go里的方法
	requestUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=" + APPID + "&secret=" + SECRET + "&js_code=" + JSCODE + "&grant_type=authorization_code"
	resp, err := http.Get(requestUrl)
	if err != nil {
		beego.Error(err)
		// return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		beego.Error(err)
		// return
	}
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		beego.Error(err)
		// return
	}
	// beego.Info(data)
	if _, ok := data["session_key"]; !ok {
		errcode := data["errcode"]
		errmsg := data["errmsg"].(string)
		// return
		c.Data["json"] = map[string]interface{}{"errNo": errcode, "msg": errmsg, "data": "session_key 不存在"}
		c.ServeJSON()
	} else {
		var openID string
		var sessionKey string
		// var unionId string
		openID = data["openid"].(string)
		sessionKey = data["session_key"].(string)

		md5Ctx := md5.New()
		md5Ctx.Write([]byte(sessionKey))
		skey := md5Ctx.Sum(nil)

		user, err := models.GetUserByOpenID(openID)
		if err != nil {
			beego.Error(err)
			c.Data["json"] = map[string]interface{}{"errNo": 0, "msg": "未查到用户", "data": "这个openID的用户不存在"}
			c.ServeJSON()
		} else {
			//根据userid取出user和avatorUrl
			useravatar, err := models.GetUserAvatorUrl(user.Id)
			if err != nil {
				beego.Error(err)
			}
			wxsite := beego.AppConfig.String("wxreqeustsite")
			var photo string
			if len(useravatar) != 0 {
				photo = wxsite + useravatar[0].UserAvatar.AvatarUrl
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

			var isAdmin bool
			// // beego.Info(roles)
			// for _, v := range roles {
			// 	// beego.Info(v.Rolename)
			// 	if v.Rolename == "admin" {
			// 		isAdmin = true
			// 	}
			// }
			//判断是否具备admin角色
			role, err := models.GetRoleByRolename("admin")
			if err != nil {
				beego.Error(err)
			}
			uid := strconv.FormatInt(user.Id, 10)
			roleid := strconv.FormatInt(role.Id, 10)
			isAdmin = e.HasRoleForUser(uid, "role_"+roleid)
			// useridstring := strconv.FormatInt(user.Id, 10)
			//用户登录后，存储openid在服务端的session里，下次用户通过hotqinsessionid来取得openid
			c.SetSession("openID", openID)
			c.SetSession("sessionKey", sessionKey) //这个没用
			// https://blog.csdn.net/yelin042/article/details/71773636
			c.SetSession("skey", skey) //这个没用
			//为何有这个hotqinsessionid?
			//用户小程序register后，只是存入服务器数据库中的openid和用户名对应
			//用户小程序login的时候，即这里，将openid存入session
			//下次用户请求携带hotqinsessionid即可取到session-openid了。
			sessionId := c.Ctx.Input.Cookie(beego.AppConfig.String("SessionName")) //这一步什么意思
			c.Data["json"] = map[string]interface{}{"errNo": 1, "msg": "success", "userId": uid, "isAdmin": isAdmin, "sessionId": sessionId, "photo": photo, "appreciationphoto": appreciationphoto}
			c.ServeJSON()
		}
		// beego.Info(useridstring)
		// unionId = data["unionid"].(string)
		// beego.Info(openID)
		// beego.Info(sessionKey)
		// beego.Info(unionId)
		//如果数据库存在记录，则存入session？
		//上传文档的时候，检查session？
	}
}

// @Title get wx haslogin
// @Description get wx usersession
// @Success 200 {object} success
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /wxhassession [get]
//微信小程序根据小程序自身存储的session，检查ecms里的openid session是否有效
func (c *LoginController) WxHasSession() {
	// var openID string
	openid := c.GetSession("openID")
	// beego.Info(openid)
	if openid == nil {
		c.Data["json"] = map[string]interface{}{"errNo": 0, "msg": "ecms中session失效"}
		c.ServeJSON()
	} else {
		// openID = openid.(string)
		c.Data["json"] = map[string]interface{}{"errNo": 1, "msg": "ecms中session有效"}
		c.ServeJSON()
	}
}

// [login.go:224] 0716J5410OfFLF1Daw610f6a4106J54e
// [login.go:247] map[session_key:3NaIB1t/AOjCQKitWx1fr
// Q== openid:oRgfy5MQlRRxyyNrENpZWnhniO-I]
// 2018/09/09 18:57:04.791 [C] [asm_amd64.s:509] the request url is  /wx/wxlogin
// 2018/09/09 18:57:04.807 [C] [asm_amd64.s:509] Handler crashed with error interfa
// ce conversion: interface {} is nil, not string
// 2018/09/09 18:57:04.807 [C] [asm_amd64.s:509] D:/gowork/src/github.com/3xxx/engi
// neercms/controllers/login.go:260
//判断用户是否登录
func checkAccount(ctx *context.Context) bool {
	var user models.User
	//（4）获取当前的请求会话，并返回当前请求会话的对象
	//但是我还是建议大家采用 SetSession、GetSession、DelSession 三个方法来操作，避免自己在操作的过程中资源没释放的问题
	// sess, _ := globalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	// defer sess.SessionRelease(ctx.ResponseWriter)
	v := ctx.Input.CruSession.Get("uname")
	if v == nil {
		return false
		//     this.SetSession("asta", int(1))
		//     this.Data["num"] = 0
	} else {
		//     this.SetSession("asta", v.(int)+1)
		//     this.Data["num"] = v.(int)
		user.Username = v.(string)
		v = ctx.Input.CruSession.Get("pwd")
		user.Password = v.(string) //ck.Value
		err := models.ValidateUser(user)
		if err == nil {
			return true
		} else {
			return false
		}
	}
}

func checkRole(ctx *context.Context) (role string, err error) { //这里返回用户的role
	v := ctx.Input.CruSession.Get("uname")
	var user models.User
	user.Username = v.(string) //ck.Value
	user, err = models.GetUserByUsername(user.Username)
	if err != nil {
		beego.Error(err)
	}
	return user.Role, err
}

// type Session struct {
// 	Session int
// }
// type Login struct {
// 	UserName string
// 	Password string
// }

func Authorizer(ctx *context.Context) (uname, role string, uid int64) {
	v := ctx.Input.CruSession.Get("uname") //用来获取存储在服务器端中的数据??。
	// beego.Info(v)                          //qin.xc
	var user models.User
	var err error
	if v != nil { //如果登录了
		uname = v.(string)
		user, err = models.GetUserByUsername(uname)
		if err != nil {
			beego.Error(err)
		} else {
			uid = user.Id
			//获取用户角色，用户角色也要存入token么？
			role = user.Role
		}
	} else { //如果没登录
		role = "anonymous"
	}
	return uname, role, uid
}

//用户登录，则role是1则是admin，其余没有意义
//ip区段，casbin中表示，比如9楼ip区段作为用户，赋予了角色，这个角色具有访问项目目录权限
func checkprodRole(ctx *context.Context) (uname, role string, uid int64, isadmin, islogin bool) {
	v := ctx.Input.CruSession.Get("uname") //用来获取存储在服务器端中的数据??。
	var userid, roleid, userrole string
	var user models.User
	var err error
	var iprole int
	if v != nil { //如果登录了
		islogin = true
		uname = v.(string)
		user, err = models.GetUserByUsername(uname)
		if err != nil {
			beego.Error(err)
		} else {
			//查询admin角色的id
			//重新获取roleid
			role, err := models.GetRoleByRolename("admin")
			if err != nil {
				beego.Error(err)
			}
			userid = strconv.FormatInt(user.Id, 10)
			roleid = strconv.FormatInt(role.Id, 10)
			isadmin = e.HasRoleForUser(userid, "role_"+roleid)

			uid = user.Id
			if user.Role == "0" {
				// isadmin = false
				userrole = "4"
				// } else if user.Role == "1" {
				// isadmin = true
				// userrole = user.Role
			} else {
				// isadmin = false
				userrole = user.Role
			}
		}
	} else { //如果没登录,查询ip对应的用户
		islogin = false
		isadmin = false
		uid = 0
		uname = ctx.Input.IP()
		// beego.Info(uname)
		user, err = models.GetUserByIp(uname)
		if err != nil { //如果查不到，则用户名就是ip，role再根据ip地址段权限查询
			// beego.Error(err)
			iprole = Getiprole(ctx.Input.IP()) //查不到，则是5——这个应该取消，采用casbin里的ip区段
			userrole = strconv.Itoa(iprole)
		} else { //如果查到，则role和用户名
			//查询admin角色的id
			//重新获取roleid
			role, err := models.GetRoleByRolename("admin")
			if err != nil {
				beego.Error(err)
			}
			userid = strconv.FormatInt(user.Id, 10)
			roleid = strconv.FormatInt(role.Id, 10)
			isadmin = e.HasRoleForUser(userid, "role_"+roleid)

			// if user.Role == "1" {
			// 	isadmin = true
			// }
			uid = user.Id
			userrole = user.Role
			uname = user.Username
			islogin = true
		}
	}
	return uname, userrole, uid, isadmin, islogin
}

func CheckprodRole(ctx *context.Context) (uname, role string, uid int64, isadmin, islogin bool) {
	v := ctx.Input.CruSession.Get("uname") //用来获取存储在服务器端中的数据??。
	var userid, roleid, userrole string
	var user models.User
	var err error
	var iprole int
	if v != nil { //如果登录了
		islogin = true
		uname = v.(string)
		user, err = models.GetUserByUsername(uname)
		if err != nil {
			beego.Error(err)
		} else {
			//查询admin角色的id
			//重新获取roleid
			role, err := models.GetRoleByRolename("admin")
			if err != nil {
				beego.Error(err)
			}
			userid = strconv.FormatInt(user.Id, 10)
			roleid = strconv.FormatInt(role.Id, 10)
			isadmin = e.HasRoleForUser(userid, "role_"+roleid)

			uid = user.Id
			if user.Role == "0" {
				// isadmin = false
				userrole = "4"
				// } else if user.Role == "1" {
				// isadmin = true
				// userrole = user.Role
			} else {
				// isadmin = false
				userrole = user.Role
			}
		}
	} else { //如果没登录,查询ip对应的用户
		islogin = false
		isadmin = false
		uid = 0
		uname = ctx.Input.IP()
		// beego.Info(uname)
		user, err = models.GetUserByIp(uname)
		if err != nil { //如果查不到，则用户名就是ip，role再根据ip地址段权限查询
			// beego.Error(err)
			iprole = Getiprole(ctx.Input.IP()) //查不到，则是5——这个应该取消，采用casbin里的ip区段
			userrole = strconv.Itoa(iprole)
		} else { //如果查到，则role和用户名
			//查询admin角色的id
			//重新获取roleid
			role, err := models.GetRoleByRolename("admin")
			if err != nil {
				beego.Error(err)
			}
			userid = strconv.FormatInt(user.Id, 10)
			roleid = strconv.FormatInt(role.Id, 10)
			isadmin = e.HasRoleForUser(userid, "role_"+roleid)

			// if user.Role == "1" {
			// 	isadmin = true
			// }
			uid = user.Id
			userrole = user.Role
			uname = user.Username
			islogin = true
		}
	}
	return uname, userrole, uid, isadmin, islogin
}

// @Title get user login...
// @Description get login..
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /islogin [get]
//login弹框输入用户名和密码后登陆提交
func (c *LoginController) Islogin() {
	var islogin, isadmin bool
	var uname string
	var uid int64
	v := c.GetSession("uname")
	// v := c.Ctx.CruSession.Get("uname") //用来获取存储在服务器端中的数据??。
	var userrole string
	var user models.User
	var err error
	var iprole int
	if v != nil { //如果登录了
		islogin = true
		uname = v.(string)
		user, err = models.GetUserByUsername(uname)
		if err != nil {
			beego.Error(err)
		} else {
			uid = user.Id
			if user.Role == "0" {
				isadmin = false
				userrole = "4"
			} else if user.Role == "1" {
				isadmin = true
				userrole = user.Role
			} else {
				isadmin = false
				userrole = user.Role
			}
		}
	} else { //如果没登录,查询ip对应的用户
		islogin = false
		isadmin = false
		uid = 0
		uname = c.Ctx.Input.IP()
		// beego.Info(uname)
		user, err = models.GetUserByIp(uname)
		if err != nil { //如果查不到，则用户名就是ip，role再根据ip地址段权限查询
			// beego.Error(err)
			iprole = Getiprole(c.Ctx.Input.IP()) //查不到，则是5——这个应该取消，采用casbin里的ip区段
			userrole = strconv.Itoa(iprole)
		} else { //如果查到，则role和用户名
			if user.Role == "1" {
				isadmin = true
			}
			uid = user.Id
			userrole = user.Role
			uname = user.Username
			islogin = true
		}
	}
	c.Data["json"] = map[string]interface{}{"uname": uname, "role": userrole, "uid": uid, "islogin": islogin, "isadmin": isadmin}
	c.ServeJSON()
}

//*********SSO单点登录JWT token
// @Title get sso login
// @Description get sso logintpl
// @Param service query string  false "The service of login"
// @Success 200 {object} success
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /ssologin [get]
func (c *LoginController) SsoLogin() {
	c.Data["service"] = c.GetString("service")
	// beego.Info(c.GetString("service"))
	// token_head := c.Ctx.GetCookie("TOKEN")
	//http://sso.dxy.cn/login?service=https://www.a.cn
	authString := c.Ctx.GetCookie("TOKEN") //
	// beego.Info(authString)
	// authString = c.Ctx.Input.Header("Authorization")
	// beego.Info(authString)

	// authString = c.Ctx.Request.Header.Get("Authorization")
	// beego.Info(authString)
	// beego.Debug("AuthString:", authString)

	// kv := strings.Split(authString, " ")
	// if len(kv) != 2 || kv[0] != "Bearer" {
	// beego.Error("AuthString invalid:", authString)
	// return
	// }
	// tokenString := kv[1]
	// var authorizeAttributes = attributeList.OfType<TestAuthorizeAttribute>().ToList();
	// var claims = context.HttpContext.User.Claims;
	// // 从claims取出用户相关信息，到数据库中取得用户具备的权限码，与当前Controller或Action标识的权限码做比较
	// var userPermissions = "User_Edit";
	// if (!authorizeAttributes.Any(s => s.Permission.Equals(userPermissions)))
	// {
	//     context.Result = new JsonResult("没有权限");
	// }
	// return;

	//校验用户名密码（对Session匹配，或数据库数据匹配）
	//  private AuthorizeState ValidateTicket(string encryptToken,string role)
	//  {
	//    if (encryptToken == null)
	//    {
	//        return AuthorizeState.TokenErro;
	//    }
	//    //解密Ticket
	//    var strTicket = FormsAuthentication.Decrypt(encryptToken).UserData;
	//    //从Ticket里面获取用户名和密码
	//    var index = strTicket.IndexOf("&");
	//    string userName = strTicket.Substring(0, index);
	//    string password = strTicket.Substring(index + 1);
	//    ArrayList arrayList = new ArrayList(role.Split(','));
	//    var roleName = HttpContext.Current.Session["Role"].ToString();
	//    var name = HttpContext.Current.Session["UserName"].ToString();
	//    //取得session，不通过说明用户退出，或者session已经过期
	//    if (arrayList.Contains(roleName) && name == userName)  //获取对应控制器 对应方法的访问角色权限 如果包含说明符合访问 否则返回权限错误
	//    {
	//        return AuthorizeState.ValidateSuucss;
	//    }
	//    else
	//    {
	//        return AuthorizeState.UserValidateErro;
	//    }
	// }

	//  当我们去访问这个方法时。他会先进行身份验证。进入MVCAuthorize中。
	//  这里你可以扩展开来。 比如我临时需要对这个方法在多开放些角色 ，可以直接在action 带上
	//    [MVCAuthorize(Roles = "VIP9,ActiveUser")]
	//    public ActionResult Chat()
	//    {
	//      var account = HttpContext.User.Identity.Name;
	//      ModelDBContext db = new ModelDBContext();
	//      //ViewBag.UserName =db.User.Where(x=>x.Email == account).FirstOrDefault().FullName;
	//      return View();
	//    }

	c.Data["service"] = c.GetString("service")
	username, err := utils.CheckToken(authString)
	if err != nil {
		beego.Error(err)
		c.TplName = "sso/index.tpl"
	} else if username != "" {
		c.TplName = "sso/success.tpl"
	} else {
		c.TplName = "sso/index.tpl"
	}
}

// @Title post sso login
// @Description post sso login data
// @Param login_name query string  true "The name of user"
// @Param password query string  true "The password of user"
// @Param service query string  false "The service of location.href"
// @Success 200 {object} success
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /ssologinpost [post]
func (c *LoginController) SsoLoginPost() {
	//获取service
	login_name := c.GetString("login_name")
	password := c.GetString("password")
	service := c.GetString("service")
	//returnurl := this.GetString("returnurl")
	// 验证用户名密码
	if len(password) == 0 || len(login_name) == 0 {
		c.Data["json"] = map[string]interface{}{"success": -1, "message": "用户名或密码为空"}
		c.ServeJSON()
		return
	}
	//密码md5加密
	var user models.User
	// user.Username = c.Input().Get("uname")
	// Pwd1 := c.Input().Get("pwd")
	// autoLogin := c.Input().Get("autoLogin") == "on"
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	cipherStr := md5Ctx.Sum(nil)

	user.Password = hex.EncodeToString(cipherStr)
	user.Username = login_name
	err := models.ValidateUser(user)
	if err != nil {
		// return nil, errors.New("Error: 未找到该用户")
		c.Data["json"] = map[string]interface{}{"success": -1, "message": "未找到该用户", "service": service}
		c.ServeJSON()
		return
	} else {
		//更新user表的lastlogintime
		err = models.UpdateUserlastlogintime(user.Username)
		if err != nil {
			beego.Error(err)
			utils.FileLogs.Error(user.Username + " 更新用户登录时间 " + err.Error())
		} else {
			tokenString, err := utils.CreateToken(login_name)
			if err != nil {
				beego.Error(err)
			}
			//请求中head中设置token

			// c.Ctx.Output.Header("TOKEN", tokenString)
			// c.Ctx.Output.Header("Content-Type", "application/json; charset=utf-8")

			//下面是写在response head中，没什么用处
			c.Ctx.Output.Header("Authorization", tokenString)

			// response := utils.Token{tokenString}
			// json, err := json.Marshal(response)
			// if err != nil {
			// 	http.Error(c.Ctx.ResponseWriter, err.Error(), http.StatusInternalServerError)
			// 	return
			// }
			// c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
			// c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
			// c.Ctx.ResponseWriter.Header().Set("Authorization", tokenString)
			// c.Ctx.ResponseWriter.Write(json)
			// beego.Info(tokenString)
			//设置cookie 名称,值,时间,路径
			c.Ctx.SetCookie("TOKEN", tokenString, "3600", "/")
			// this.Ctx.SetCookie("USERID", user.Id, "3600", "/")    "user": user,
			c.Data["json"] = map[string]interface{}{"success": 0, "msg": "登录成功", "service": service}
			c.ServeJSON()
		}
	}
}

func (c *SuccessController) Get() {
	c.TplName = "Success.html"
}

func (c *ServiceValidateController) Get() {
	ticket := c.GetString("ticket")
	userName, err := utils.CheckToken(ticket)
	if err != nil {
		beego.Error(err)
	}
	if userName != "" {
		c.Data["json"] = map[string]interface{}{"success": 0, "msg": "登录成功", "user": userName}
	} else {
		c.Data["json"] = map[string]interface{}{"success": -1, "msg": "用户未登录"}
	}
	c.ServeJSON()
}

//********

// func Authorizer1(e *casbin.Enforcer, users models.User) func(next http.Handler) http.Handler {
// 	role, err := session.GetString(r, "role")
// 	if err != nil {
// 		writeError(http.StatusInternalServerError, "ERROR", w, err)
// 		return
// 	}
// 	if role == "" {
// 		role = "anonymous"
// 	}
// 	// if it's a member, check if the user still exists
// 	if role == "member" {
// 		uid, err := session.GetInt(r, "userID")
// 		if err != nil {
// 			writeError(http.StatusInternalServerError, "ERROR", w, err)
// 			return
// 		}
// 		exists := users.Exists(uid)
// 		if !exists {
// 			writeError(http.StatusForbidden, "FORBIDDEN", w, errors.New("user does not exist"))
// 			return
// 		}
// 	}
// 	// casbin rule enforcing
// 	res, err := e.EnforceSafe(role, r.URL.Path, r.Method)
// 	if err != nil {
// 		writeError(http.StatusInternalServerError, "ERROR", w, err)
// 		return
// 	}
// 	if res {
// 		next.ServeHTTP(w, r)
// 	} else {
// 		writeError(http.StatusForbidden, "FORBIDDEN", w, errors.New("unauthorized"))
// 		return
// 	}
// }

// func checkRole(ctx *context.Context) (roles []*models.Role, err error) {
// 	ck, err := ctx.Request.Cookie("uname")
// 	if err != nil {
// 		return roles, err
// 	}
// 	var user models.User
// 	user.Username = ck.Value

// 	roles, _ = models.GetRoleByUsername(user.Username)
// 	if err == nil {
// 		return roles, err
// 	} else {
// 		return roles, err
// 	}
// }

// func GetRoleByUserId(userid int64) (roles []*Role, count int64) { //*Topic, []*Attachment, error
// 	roles = make([]*Role, 0)
// 	o := orm.NewOrm()
// 	// role := new(Role)
// 	count, _ = o.QueryTable("role").Filter("Users__User__Id", userid).All(&roles)
// 	return roles, count
// 	// 通过 post title 查询这个 post 有哪些 tag
// 	// var tags []*Tag
// 	// num, err := dORM.QueryTable("tag").Filter("Posts__Post__Title", "Introduce Beego ORM").All(&tags)
// }

// func (this *UserController) HandleLogin() {
// 	//1.获取数据
// 	userName := this.GetString("username")
// 	pwd := this.GetString("pwd")
// 	check := this.GetString("check")
// 	if userName == "" || pwd == "" {
// 		this.Data["errmsg"] = "用户名或密码不能为空,请重新登陆！"
// 		this.TplName = "login.html"
// 		return
// 	}
// 	//2.查询数据
// 	o := orm.NewOrm()
// 	user := models.User{Name: userName}
// 	err := o.Read(&user, "Name")
// 	if err != nil {
// 		this.Data["errmsg"] = "用户名或密码错误,请重新登陆！"
// 		this.TplName = "login.html"
// 		return
// 	}
// 	if user.PassWord != pwd {
// 		this.Data["errmsg"] = "用户名或密码错误,请重新登陆！"
// 		this.TplName = "login.html"
// 		return
// 	}
// 	if user.Active != true {
// 		this.Data["errmsg"] = "该用户没有激活，请先激活！"
// 		this.TplName = "login.html"
// 		return
// 	}
// 	if check == "on" {
// 		this.Ctx.SetCookie("username", userName, time.Second*3600)
// 	} else {
// 		this.Ctx.SetCookie("username", userName, -1)
// 	}
// 	this.SetSession("userName", userName)
// 	this.Redirect("/", 302)
// }
