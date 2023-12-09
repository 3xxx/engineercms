package controllers

import (
	// m "github.com/beego/admin/src/models"
	// "github.com/beego/beego/v2/adapter/orm"
	"crypto/md5"
	"encoding/hex"
	m "github.com/3xxx/engineercms/models"
	// beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	// "github.com/beego/beego/v2/adapter/logs"
	"github.com/tealeg/xlsx"
	"html/template"
	"os"
	"strconv"
	"strings"
	"time"
)

type UserController struct {
	web.Controller
}

func (c *UserController) Index() {
	// page, _ := c.GetInt64("page")
	// page_size, _ := c.GetInt64("rows")
	// sort := c.GetString("sort")
	// order := c.GetString("order")
	// if len(order) > 0 {
	// 	if order == "desc" {
	// 		sort = "-" + sort
	// 	}
	// } else {
	// 	sort = "Id"
	// }
	// 	c.Data["IsCategory"] = true
	// c.TplName = "category.tpl"
	//1.首先判断是否注册
	// if !checkAccount(c.Ctx) {
	// 	// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/login?url="+route, 302)
	// 	// c.Redirect("/login", 302)
	// 	return
	// }
	//2.取得客户端用户名
	// sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	// defer sess.SessionRelease(c.Ctx.ResponseWriter)
	// v := sess.Get("uname")
	// if v != nil {
	// 	c.Data["Uname"] = v.(string)
	// }
	// v := c.GetSession("uname")
	// if v != nil {
	// 	c.Data["Uname"] = v.(string)
	// }
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err == nil {
	// 	c.Data["Uname"] = ck.Value
	// } else {
	// 	logs.Error(err)
	// }
	//2.取得客户端用户名
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	logs.Error(err)
	// }
	// uname := ck.Value
	//3.取出用户的权限等级
	// category, err := models.GetCategory(id)
	// beego.Info(username)
	//4.取得客户端用户名
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	logs.Error(err)
	// }
	// uname := ck.Value
	//5.取出用户的权限等级
	// role, _ := checkRole(c.Ctx) //login里的
	// beego.Info(role)
	//6.进行逻辑分析：
	// rolename, _ := strconv.ParseInt(role, 10, 64)
	// if filetype != "pdf" && filetype != "jpg" && filetype != "diary" {
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	if !isadmin { //&& uname != category.Author
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// }

	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	//2.取得客户端用户名
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	logs.Error(err)
	// } else {
	// 	c.Data["Uname"] = ck.Value
	// }
	users, count := m.Getuserlist(1, 2000, "Id")
	if c.IsAjax() {
		c.Data["json"] = &map[string]interface{}{"total": count, "rows": &users}
		c.ServeJSON()
		return
	} else {
		// tree := c.GetTree()
		// c.Data["tree"] = &tree
		c.Data["Users"] = &users
		c.TplName = "user.tpl"
		// if c.GetTemplatetype() != "easyui" {
		// c.Layout = c.GetTemplatetype() + "/public/layout.tpl"
		// }
		// c.TplName = c.GetTemplatetype() + "/rbac/user.tpl"
	}
}

// 后端分页的数据结构
type userTableserver struct {
	Rows  []*m.User `json:"rows"`
	Page  int64     `json:"page"`
	Total int64     `json:"total"` //string或int64都行！
}

// @Title get user
// @Description get user by userid
// @Param id path string false "The id of user"
// @Param role query string false "The role of user"
// @Param searchText query string false "The searchText of users"
// @Param pageNo query string true "The page for users list"
// @Param limit query string true "The limit of page for users list"
// @Success 200 {object} models.User
// @Failure 400 Invalid page supplied
// @Failure 404 user not found
// @router /user/:id [get]
// 如果不带id则取到所有用户
// 如果带id，则取一个用户
func (c *UserController) User() {
	_, _, _, isadmin, _ := CheckprodRole(c.Ctx)
	if !isadmin {
		c.Data["json"] = map[string]interface{}{"msg": "非管理员权限，无法查询用户信息！", "data": "请登录管理员。"}
		c.ServeJSON()
	}
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	c.Data["Ip"] = c.Ctx.Input.IP()
	// var categories []*models.AdminCategory
	if id == "0" { //如果id为空，则查询类别
		limit := c.GetString("limit")
		if limit == "" {
			limit = "5"
		}
		limit1, err := strconv.ParseInt(limit, 10, 64)
		if err != nil {
			logs.Error(err)
		}
		page := c.GetString("pageNo")
		page1, err := strconv.ParseInt(page, 10, 64)
		if err != nil {
			logs.Error(err)
		}

		var offset int64
		if page1 <= 1 {
			offset = 0
		} else {
			offset = (page1 - 1) * limit1
		}
		searchText := c.GetString("searchText")
		users, count, err := m.GetUsersPage(limit1, offset, "Username", searchText)
		if err != nil {
			logs.Error(err)
		}
		for _, w := range users {
			w.Password = "0123456789abcdefghijklmnopqrstuv"
		}
		//如果设置了role,用于onlyoffice/officeview的权限设置
		role := c.GetString("role")
		if role != "" {
			for _, v := range users {
				v.Role = role
			}
		}
		table := userTableserver{users, page1, count}
		// table := prodTableserver2{products, 1, 20}
		c.Data["json"] = table
		// c.Data["json"] = &users
		c.ServeJSON()
	} else {
		//pid转成64为
		idNum, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			logs.Error(err)
		}
		user := m.GetUserByUserId(idNum)
		if err != nil {
			logs.Error(err)
		}

		// var users1 []*m.User
		users := make([]*m.User, 1)
		users[0] = &user
		users[0].Password = "0123456789abcdefghijklmnopqrstuv"
		// users = append(users, &user...)
		c.Data["json"] = users //取到一个用户数据，不是数组，所以table无法显示
		c.ServeJSON()
	}
	// if c.IsAjax() {
	// 	c.Data["json"] = &map[string]interface{}{"total": count, "rows": &users}
	// 	c.ServeJSON()
	// 	return
	// } else {
	// }
}

// 用户登录查看自己的资料_不是这个，是GetUserByUsername
func (c *UserController) View() {
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	userid, _ := strconv.ParseInt(c.GetString("useid"), 10, 64)
	user := m.GetUserByUserId(userid)
	c.Data["User"] = user
	c.TplName = "admin_user_view.tpl"
}

// @Title add user
// @Description add user
// @Param username query string false "The name of user"
// @Param nickname query string false "The nickname of user"
// @Param password query string false "The password of user"
// @Param email query string false "The email of user"
// @Param department query string false "The department of user"
// @Param secoffice query string false "The secoffice of user"
// @Param ip query string false "The ip of user"
// @Param port query string false "The port of user"
// @Param status query string false "The status of user"
// @Param role query string false "The role of user"
// @Success 200 {object} models.User
// @Failure 400 Invalid page supplied
// @Failure 404 user not found
// @router /adduser [post]
// 添加用户
func (c *UserController) AddUser() {
	var user m.User
	user.Username = c.GetString("username")
	user.Nickname = c.GetString("nickname")

	Pwd1 := c.GetString("password")
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(Pwd1))
	cipherStr := md5Ctx.Sum(nil)
	user.Password = hex.EncodeToString(cipherStr)
	// user.Repassword = c.GetString("repassword")
	// Pwd1=c.GetString("password")
	// md5Ctx := md5.New()
	// md5Ctx.Write([]byte(Pwd1))
	// cipherStr := md5Ctx.Sum(nil)
	// user.Password = hex.EncodeToString(cipherStr)

	user.Email = c.GetString("email")
	user.Sex = c.GetString("sex")
	ispartymember := c.GetString("ispartymember")
	if ispartymember == "true" {
		user.IsPartyMember = true
	} else {
		user.IsPartyMember = false
	}
	user.Department = c.GetString("department")
	user.Secoffice = c.GetString("secoffice")
	user.Ip = c.GetString("ip")
	user.Port = c.GetString("port")
	statusint, err := strconv.Atoi(c.GetString("status"))
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "message": err}
		c.ServeJSON()
	}
	user.Status = statusint
	user.Role = c.GetString("role")
	id, err := m.SaveUser(user)
	if err == nil && id > 0 {
		// c.Rsp(true, "Success")
		// return
		c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "message": "ok", "data": id}
		c.ServeJSON()
	} else if err == nil && id == 0 {
		// c.Rsp(false, err.Error())
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "message": "用户已存在", "data": id}
		c.ServeJSON()
		// return
	} else {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "message": err, "data": id}
		c.ServeJSON()
	}
}

// @Title post wx user
// @Description post user
// @Param uname query string true "The username of user"
// @Param nickname query string true "The nickname of user"
// @Param password query string true "The Password of user"
// @Param email query string false "The email of user"
// @Param department query string false "The department of user"
// @Param secoffice query string false "The secoffice of user"
// @Success 200 {object} models.AddUser
// @Failure 400 Invalid page supplied
// @Failure 404 user not found
// @router /addwxuser [post]
// 小程序添加用户
func (c *UserController) AddWxUser() {
	var user m.User
	var err error
	openID := c.GetSession("openID")
	if openID != nil {
		user2, err := m.GetUserByOpenID(openID.(string))
		if err != nil {
			logs.Error(err)
		}
		//判断是否具备admin角色
		role, err := m.GetRoleByRolename("admin")
		if err != nil {
			logs.Error(err)
		}
		uid := strconv.FormatInt(user2.Id, 10)
		roleid := strconv.FormatInt(role.Id, 10)
		isAdmin, err := e.HasRoleForUser(uid, "role_"+roleid)
		if err != nil {
			logs.Error(err)
		}
		if !isAdmin {
			c.Data["json"] = map[string]interface{}{"info": "非管理员", "id": 0}
			c.ServeJSON()
			return
		}
	} else {
		c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		c.ServeJSON()
		return
		// user.Id = 9
	}
	user.Username = c.GetString("uname")
	user.Nickname = c.GetString("nickname")
	// beego.Info(user.Username)
	Pwd1 := c.GetString("password")
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(Pwd1))
	cipherStr := md5Ctx.Sum(nil)
	user.Password = hex.EncodeToString(cipherStr)
	// user.Email = c.GetString("email")
	// user.Department = c.GetString("department")
	// user.Secoffice = c.GetString("secoffice")
	// user.Ip = c.GetString("ip")
	// user.Port = c.GetString("port")
	// statusint, err := strconv.Atoi(c.GetString("status"))
	// if err != nil {
	// 	logs.Error(err)
	// }
	user.Status = 1
	user.Role = "4"
	id, err := m.SaveUser(user)
	if err == nil && id > 0 {
		// c.Rsp(true, "Success")
		// return
		c.Data["json"] = map[string]interface{}{"info": "success", "id": id}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"info": "添加失败", "id": 0}
		c.ServeJSON()
	}
}

// func (c *UserController) UpdateUser() {
// 	u := m.User{}
// 	if err := c.ParseForm(&u); err != nil {
// 		//handle error
// 		// c.Rsp(false, err.Error())
// 		logs.Error(err.Error)
// 		return
// 	}
// 	id, err := m.UpdateUser(&u)
// 	if err == nil && id > 0 {
// 		// c.Rsp(true, "Success")
// 		return
// 	} else {
// 		// c.Rsp(false, err.Error())
// 		logs.Error(err.Error)
// 		return
// 	}
// }

// @Title post wx userpassword by uid
// @Description post user password by uid
// @Param uid query string  true "The id of user"
// @Param oldpass query string  true "The oldPassword of user"
// @Param newpass query string  true "The newpassword of user"
// @Param id path string  true "The id of wx"
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /updatewxuser [post]
// 小程序修改用户密码
func (c *UserController) UpdateWxUser() {
	oldpass := c.GetString("oldpass")
	uid := c.GetString("uid")
	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	//取出用户旧密码
	user := m.GetUserByUserId(id)
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(oldpass))
	cipherStr := md5Ctx.Sum(nil)
	Password := hex.EncodeToString(cipherStr)
	if user.Password == Password {
		newpass := c.GetString("newpass")
		err = m.UpdateUser(id, "Password", newpass)
		if err != nil {
			logs.Error(err)
			c.Data["json"] = "wrong"
			c.ServeJSON()
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = "旧密码不正确"
		c.ServeJSON()
	}
}

// @Title update user
// @Description add user
// @Param pk query string false "The pk of user"
// @Param name query string false "The name of user"
// @Param value query string false "The value of user"
// @Success 200 {object} models.User
// @Failure 400 Invalid page supplied
// @Failure 404 user not found
// @router /updateuser [post]
// 在线修改保存某个字段
func (c *UserController) UpdateUser() {
	//进行权限判断isme or isadmin
	_, _, uid, isadmin, _ := checkprodRole(c.Ctx)
	pk := c.GetString("pk") //这个其实就是userid
	id, err := strconv.ParseInt(pk, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	if isadmin || uid == id {
		name := c.GetString("name")
		value := c.GetString("value")
		logs.Info(value)
		value = template.HTMLEscapeString(value) //过滤xss攻击
		logs.Info(value)
		err = m.UpdateUser(id, name, value)
		if err != nil {
			logs.Error(err)
			data := "写入数据错误!"
			c.Ctx.WriteString(data)
		} else {
			logs.Info("ok")
			data := "ok!"
			c.Ctx.WriteString(data)
		}
	} else {
		data := "权限不足！"
		c.Ctx.WriteString(data)
	}
}

//这个作废，用在线修改
// func (c *UserController) UpdateUser() {
// 	userid := c.GetString("userid")
// 	nickname := c.GetString("nickname")
// 	email := c.GetString("email")
// 	Pwd1 := c.GetString("password")
// 	if Pwd1 != "" {
// 		md5Ctx := md5.New()
// 		md5Ctx.Write([]byte(Pwd1))
// 		cipherStr := md5Ctx.Sum(nil)
// 		password := hex.EncodeToString(cipherStr)
// 		err := m.UpdateUser(userid, nickname, email, password) //这里修改
// 		if err != nil {
// 			logs.Error(err)
// 		}
// 	} else {
// 		err := m.UpdateUser(userid, nickname, email, "") //这里修改
// 		if err != nil {
// 			logs.Error(err)
// 		}
// 	}
// 	c.TplName = "user_view.tpl"
// }

// func (c *UserController) DeleteUser() {
// 	Id, _ := c.GetInt64("userid")
// 	status, err := m.DelUserById(Id)
// 	if err == nil && status > 0 {
// 		// c.Rsp(true, "Success")
// 		c.Redirect("/user/index", 302)
// 		return
// 	} else {
// 		// c.Rsp(false, err.Error())
// 		logs.Error(err.Error)
// 		return
// 	}
// }

// @Title delete user
// @Description delete user
// @Param ids query string false "The ids of user"
// @Success 200 {object} models.User
// @Failure 400 Invalid page supplied
// @Failure 404 user not found
// @router /deleteuser [post]
// 删除用户
func (c *UserController) DeleteUser() {
	ids := c.GetString("ids")
	array := strings.Split(ids, ",")
	for _, v := range array {
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			logs.Error(err)
		}
		status, err := m.DelUserById(idNum)
		if err == nil && status > 0 {
			c.Data["json"] = "ok"
			c.ServeJSON()
		} else if err != nil {
			logs.Error(err)
		}
	}
}

// @Title get user
// @Description get user
// @Success 200 {object} models.User
// @Failure 400 Invalid page supplied
// @Failure 404 user not found
// @router /getuserbyusername [get]
// 用户查看自己，修改密码等
func (c *UserController) GetUserByUsername() {
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	if islogin != true {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		return
	}
	c.TplName = "user_view.tpl"
}

// @Title get usermyself
// @Description get usermyself
// @Success 200 {object} models.User
// @Failure 400 Invalid page supplied
// @Failure 404 user not found
// @router /usermyself [get]
// 用户个人数据，填充table，以便编辑
func (c *UserController) Usermyself() {
	// 	_, role := checkprodRole(c.Ctx)
	// 	roleint, err := strconv.Atoi(role)
	// 	if err != nil {
	// 		logs.Error(err)
	// 	}
	// 	if role == "1" {
	// 		c.Data["IsAdmin"] = true
	// 	} else if roleint > 1 && roleint < 5 {
	// 		c.Data["IsLogin"] = true
	// 	} else {
	// 		c.Data["IsAdmin"] = false
	// 		c.Data["IsLogin"] = false
	// 	}
	// 	// c.Data["Username"] = username
	// 	// c.Data["IsIndex"] = true
	// 	// c.Data["Ip"] = c.Ctx.Input.IP()
	// 	c.Data["role"] = role
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	//4.取得客户端用户名
	// var uname string
	// sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	// defer sess.SessionRelease(c.Ctx.ResponseWriter)
	// v := c.GetSession("uname")
	// if v != nil {
	// 	uname = v.(string)
	// }
	// if uname == "" {
	// 	route := c.Ctx.Request.URL.String()
	// 	c.Data["Url"] = route
	// 	c.Redirect("/roleerr?url="+route, 302)
	// 	// c.Redirect("/roleerr", 302)
	// 	return
	// }
	user, err := m.GetUserByUsername(username)
	if err != nil {
		logs.Error(err)
	}
	users := make([]*m.User, 1)
	users[0] = &user
	users[0].Password = "0123456789abcdefghijklmnopqrstuv"
	c.Data["json"] = users
	c.ServeJSON()
}

// @Title import users
// @Description import users
// @Success 200 {object} models.User
// @Failure 400 Invalid page supplied
// @Failure 404 user not found
// @router /importusers [post]
// 上传excel文件，导入到数据库
// 引用来自category的查看成果类型里的成果
func (c *UserController) ImportUsers() {
	//获取上传的文件
	_, h, err := c.GetFile("usersexcel")
	if err != nil {
		logs.Error(err)
	}
	var path string

	// var filesize int64
	if h != nil {
		//保存附件
		path = "./attachment/" + h.Filename    // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("usersexcel", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			logs.Error(err)
			c.Data["json"] = "err保存文件失败"
			c.ServeJSON()
		} else {
			var user m.User
			//读出excel内容写入数据库
			xlFile, err := xlsx.OpenFile(path) //
			if err != nil {
				logs.Error(err)
			}
			for _, sheet := range xlFile.Sheets {
				for i, row := range sheet.Rows {
					if i != 0 { //忽略第一行标题
						// 这里要判断单元格列数，如果超过单元格使用范围的列数，则出错for j := 2; j < 7; j += 5 {
						j := 1
						if len(row.Cells) >= 2 { //总列数，从1开始
							user.Username = row.Cells[j].String()
							if err != nil {
								logs.Error(err)
							}
						}
						if len(row.Cells) >= 3 {
							user.Nickname = row.Cells[j+1].String()
							if err != nil {
								logs.Error(err)
							}
						}
						if len(row.Cells) >= 4 {
							Pwd1 := row.Cells[j+2].String()
							if err != nil {
								logs.Error(err)
							}

							md5Ctx := md5.New()
							md5Ctx.Write([]byte(Pwd1))
							cipherStr := md5Ctx.Sum(nil)
							user.Password = hex.EncodeToString(cipherStr)
						}
						if len(row.Cells) >= 5 {
							user.Email = row.Cells[j+3].String()
							if err != nil {
								logs.Error(err)
							}
						}
						if len(row.Cells) >= 6 {
							user.Sex = row.Cells[j+4].String()
							if err != nil {
								logs.Error(err)
							}
						}
						if len(row.Cells) >= 7 {
							ispartymember := row.Cells[j+5].String()
							if ispartymember == "是" {
								user.IsPartyMember = true
							} else {
								user.IsPartyMember = false
							}
							if err != nil {
								logs.Error(err)
							}
						}
						if len(row.Cells) >= 8 {
							user.Department = row.Cells[j+6].String()
							if err != nil {
								logs.Error(err)
							}
						}
						if len(row.Cells) >= 9 {
							user.Secoffice = row.Cells[j+7].String()
							if err != nil {
								logs.Error(err)
							}
						}
						if len(row.Cells) >= 10 {
							user.Ip = row.Cells[j+8].String()
							if err != nil {
								logs.Error(err)
							}
						}
						if len(row.Cells) >= 11 {
							user.Port = row.Cells[j+9].String()
							if err != nil {
								logs.Error(err)
							}
						}
						if len(row.Cells) >= 12 {
							status := row.Cells[j+10].String()
							if err != nil {
								logs.Error(err)
							}
							status1, err := strconv.Atoi(status)
							if err != nil {
								logs.Error(err)
							}
							user.Status = status1
						}
						if len(row.Cells) >= 13 {
							role := row.Cells[j+11].String()
							if err != nil {
								logs.Error(err)
							}

							user.Role = role
							user.Lastlogintime = time.Now()
							_, err = m.SaveUser(user) //如果姓名重复，也要返回uid
							if err != nil {
								logs.Error(err)
							}
						}
					}
				}
			}
			//删除附件
			err = os.Remove(path)
			if err != nil {
				logs.Error(err)
			}
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = "err文件为空！"
		c.ServeJSON()
	}
}

func (this *UserController) Roleerr() {
	// url := this.GetString("url")
	url1 := this.GetString("url") //这里不支持这样的url，http://192.168.9.13/login?url=/topic/add?id=955&mid=3
	url2 := this.GetString("level")
	url3 := this.GetString("key")
	var url string
	if url2 == "" {
		url = url1
	} else {
		url = url1 + "&level=" + url2 + "&key=" + url3
	}
	this.Data["Url"] = url
	this.TplName = "role_err.tpl"
}
