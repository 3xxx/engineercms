package controllers

import (
	"crypto/md5"
	"encoding/hex"
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	// "net/url"
	"github.com/3xxx/engineercms/models"
	"strconv"
	// "github.com/astaxie/beego/session"
	"encoding/json"
	"net/http"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	isExit := c.Input().Get("exit") == "true"
	// secofficeshow?secid=1643&level=3&key=modify
	url1 := c.Input().Get("url") //这里不支持这样的url，http://192.168.9.13/login?url=/topic/add?id=955&mid=3
	url2 := c.Input().Get("level")
	url3 := c.Input().Get("key")
	// beego.Info(url1)
	// beego.Info(url2)
	// beego.Info(url3)
	var url string
	if url2 == "" {
		url = url1
	} else {
		url = url1 + "&level=" + url2 + "&key=" + url3
	}

	c.Data["Url"] = url
	// beego.Info(isExit)
	// logout user
	// func LogoutUser(ctx *context.Context) {
	// 	DeleteRememberCookie(ctx)
	// 	ctx.Input.CruSession.Delete("auth_user_id")
	// 	ctx.Input.CruSession.Flush()
	// 	beego.GlobalSessions.SessionDestroy(ctx.ResponseWriter, ctx.Request)
	// }

	if isExit {
		// c.Ctx.SetCookie("uname", "", -1, "/")
		// c.Ctx.SetCookie("pwd", "", -1, "/")
		// c.DelSession("gosessionid")
		// c.DelSession("uname")//这个不行
		// c.Destroy/Session()
		// c.Ctx.Input.CruSession.Delete("gosessionid")这句与上面一句重复
		// c.Ctx.Input.CruSession.Flush()
		// beego.GlobalSessions.SessionDestroy(c.Ctx.ResponseWriter, c.Ctx.Request)
		v := c.GetSession("uname")
		// islogin := false
		if v != nil {
			//删除指定的session
			c.DelSession("uname")
			//销毁全部的session
			c.DestroySession()
			// islogin = true

			//beego.Info("当前的session:")
			//beego.Info(c.CruSession)
		}
		// sess.Flush()//这个不灵
		c.Redirect("/", 301)
		return
	}

	//	c.Data["Website"] = "My Website"
	//	c.Data["Email"] = "your.email.address@example.com"
	//	c.Data["EmailName"] = "Your Name"
	//	c.Data["Id"] = c.Ctx.Input.Param(":id")
	c.TplName = "login.tpl"
}

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
	// beego.Info(url)
	c.TplName = "loginerr.tpl"
}

//输入用户名和密码后登陆提交
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
	// beego.Info(url)
	//（4）获取当前的请求会话，并返回当前请求会话的对象
	// sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	// defer sess.SessionRelease(c.Ctx.ResponseWriter)
	//（5）根据当前请求对象，设置一个session
	// sess.Set("mySession", "qq504284")
	// c.Data["Website"] = "广东省水利电力勘测设计研究院■☆●施工预算分院"
	//（6）从session中读取值
	// c.Data["Email"] = sess.Get("mySession")

	// beego.Info(url)
	// pwd := c.Input().Get("pwd")
	// autoLogin := c.Input().Get("autoLogin") == "on"

	// if beego.AppConfig.String("uname") == uname &&
	// 	beego.AppConfig.String("pwd") == pwd {
	// 	maxAge := 0
	// 	if autoLogin {
	// 		maxAge = 1<<31 - 1
	// 	}
	// 	c.Ctx.SetCookie("uname", uname, maxAge, "/")
	// 	c.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	// 	c.Redirect("/", 301)
	// } else {
	// 	c.Redirect("/login", 302)
	// }
	// return
	var user models.User
	user.Username = c.Input().Get("uname")
	Pwd1 := c.Input().Get("pwd")
	// autoLogin := c.Input().Get("autoLogin") == "on"
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(Pwd1))
	cipherStr := md5Ctx.Sum(nil)
	// fmt.Print(cipherStr)
	// fmt.Print("\n")
	// fmt.Print(hex.EncodeToString(cipherStr))

	user.Password = hex.EncodeToString(cipherStr)
	err := models.ValidateUser(user)
	if err == nil {
		// if beego.AppConfig.String("uname") == uname &&
		// 	beego.AppConfig.String("pwd") == pwd {
		// maxAge := 0
		// if autoLogin {
		// 	maxAge = 1<<31 - 1
		// }
		// c.Ctx.SetCookie("uname", user.Username, maxAge, "/")
		c.SetSession("uname", user.Username)
		c.SetSession("pwd", user.Password)
		// beego.Info(sess.Get("uname"))
		// c.Ctx.SetCookie("pwd", user.Password, maxAge, "/")
		User, err := models.GetUserByUsername(user.Username)
		if err != nil {
			beego.Error(err)
		}
		if User.Ip == "" {
			err = models.UpdateUser(User.Id, "Ip", c.Ctx.Input.IP())
			if err != nil {
				beego.Error(err)
			}
		} else {
			//更新user表的lastlogintime
			err = models.UpdateUserlastlogintime(user.Username)
			if err != nil {
				beego.Error(err)
			}
		}
		if url != "" {
			c.Redirect(url, 301)
			// c.Redirect("/onlyoffice", 301)
			// beego.Info(url)
		} else {
			c.Redirect("/", 301)
		}
	} else {
		// port := strconv.Itoa(c.Ctx.Input.Port())
		// route := c.Ctx.Input.Site() + ":" + port + c.Ctx.Input.URL()
		// c.Data["Url"] = route
		// c.Redirect("/login?url="+route, 302)
		c.Redirect("/loginerr?url="+url, 302)
	}
	return
	// sess := index.StartSession()
	// var user models.User
	// inputs := index.Input()
	//fmt.Println(inputs)
	// user.Username = c.Input().Get("uname")
	// Pwd1 := c.Input().Get("pwd")

	// md5Ctx := md5.New()
	// md5Ctx.Write([]byte(Pwd1))
	// cipherStr := md5Ctx.Sum(nil)
	// fmt.Print(cipherStr)
	// fmt.Print("\n")
	// fmt.Print(hex.EncodeToString(cipherStr))

	// user.Pwd = hex.EncodeToString(cipherStr)
	// err := models.ValidateUser(user)
	// if err == nil {
	// 	sess.Set("username", user.Username)
	// 	fmt.Println("username:", sess.Get("username"))
	// 	index.TplName = "success.tpl"
	// } else {
	// 	fmt.Println(err)
	// 	index.TplName = "error.tpl"
	// }
}

//微信小程序访问微信服务器获取用户信息
func (c *LoginController) WxLogin() {
	JSCODE := c.Input().Get("code")
	APPID := "wx7f77b90a1a891d93"
	SECRET := "f58ca4f28cbb52ccd805d66118060449"
	requestUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=" + APPID + "&secret=" + SECRET + "&js_code=" + JSCODE + "&grant_type=authorization_code"
	resp, err := http.Get(requestUrl)
	if err != nil {
		beego.Error(err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		beego.Error(err)
		return
	}

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		beego.Error(err)
		return
	}

	if _, ok := data["session_key"]; !ok {
		beego.Info("session_key 不存在")
		beego.Info(data)
		beego.Error(err)
		return
	}

	var openID string
	var sessionKey string
	var unionId string
	openID = data["openid"].(string)
	sessionKey = data["session_key"].(string)
	unionId = data["unionid"].(string)
	beego.Info(openID)
	beego.Info(sessionKey)
	beego.Info(unionId)

	//如果数据库存在记录，则存入session？
	//上传文档的时候，检查session？
	c.SetSession("uname", openID)
	c.SetSession("pwd", sessionKey)
	c.Data["json"] = map[string]interface{}{"errNo": 0, "msg": "success", "data": "3rd_session"}
	c.ServeJSON()
	// ctx.JSON(iris.StatusOK, iris.Map{
	// 	"errNo": model.ErrorCode.SUCCESS,
	// 	"msg":   "success",
	// 	"data":  resData,
	// })
	// session := ctx.Session()
	// session.Set("weAppOpenID", openID)
	// session.Set("weAppSessionKey", sessionKey)

	// client := http.Client{}                                        //  结构化这个你想要的东西，因为返回的东西不一定都是你想要的，把你想要的建立一个struct来接收
	// formResponse, err := client.PostForm(requestUrl, url.Values{}) // 递交请求
	// if err != nil {
	// 	panic(err)
	// }
	// defer formResponse.Body.Close()
	// body, err := ioutil.ReadAll(formResponse.Body) // 将内容读取出来。
	// type BaseInfo struct {
	// 	OpenId string `json:"openid”`
	// 	Token  string `json:"access_token”`
	// }

	// var user models.User
	// user.Username = c.Input().Get("uname")
	// err = models.ValidateUser(user)
	// if err == nil {
	// 	c.SetSession("uname", user.Username)
	// 	c.SetSession("pwd", user.Password)
	// 	User, err := models.GetUserByUsername(user.Username)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// }
}

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
	// this.TplName = "index.tpl"

	// ck, err := ctx.Request.Cookie("uname")
	// if err != nil {
	// 	return false
	// }

	//ck.Value

	// ck, err = ctx.Request.Cookie("pwd")
	// if err != nil {
	// 	return false
	// }

	// return beego.AppConfig.String("uname") == uname &&
	// 	beego.AppConfig.String("pwd") == pwd
}

func checkRole(ctx *context.Context) (role string, err error) { //这里返回用户的role
	//（4）获取当前的请求会话，并返回当前请求会话的对象
	// sess, _ := globalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	// defer sess.SessionRelease(ctx.ResponseWriter)
	v := ctx.Input.CruSession.Get("uname")
	// ck, err := ctx.Request.Cookie("uname")
	// if err != nil {
	// 	return "", err
	// }
	var user models.User
	user.Username = v.(string) //ck.Value
	user, err = models.GetUserByUsername(user.Username)
	if err != nil {
		beego.Error(err)
	}
	return user.Role, err
	// var roles []*models.Role
	// roles, _, err = models.GetRoleByUsername(user.Username)
	// if err == nil {
	// 	return roles[0].Title, err //这里修改Name改为title就对了
	// } else {
	// 	return "", err
	// }
}

// type Session struct {
// 	Session int
// }
// type Login struct {
// 	UserName string
// 	Password string
// }

//用户登录，则role是1则是admin，其余没有意义
//ip区段，casbin中表示，比如9楼ip区段作为用户，赋予了角色，这个角色具有访问项目目录权限
func checkprodRole(ctx *context.Context) (uname, role string, uid int64, isadmin, islogin bool) {
	v := ctx.Input.CruSession.Get("uname")
	// beego.Info(v)
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
		uname = ctx.Input.IP()
		// beego.Info(uname)
		user, err = models.GetUserByIp(uname)
		if err != nil { //如果查不到，则用户名就是ip，role再根据ip地址段权限查询
			beego.Error(err)
			iprole = Getiprole(ctx.Input.IP()) //查不到，则是5——这个应该取消，采用casbin里的ip区段
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
	// beego.Info(iprole)
	//如果用户登录，则以登录权限为优先，即使给这个用户设置的ip权限
	// roleint, err := strconv.Atoi(userrole)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// if iprole <= roleint {
	// 	role = strconv.Itoa(iprole)
	// } else {
	// 	role = userrole
	// }
	return uname, userrole, uid, isadmin, islogin
}

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
