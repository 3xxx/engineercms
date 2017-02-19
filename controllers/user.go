package controllers

import (
	// m "github.com/beego/admin/src/models"
	// "github.com/astaxie/beego/orm"
	"crypto/md5"
	"encoding/hex"
	m "engineercms/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/tealeg/xlsx"
	"os"
	"strconv"
	"time"
)

type UserController struct {
	beego.Controller
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
	if !checkAccount(c.Ctx) {
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/login?url="+route, 302)
		// c.Redirect("/login", 302)
		return
	}
	//2.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		c.Data["Uname"] = v.(string)
	}
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err == nil {
	// 	c.Data["Uname"] = ck.Value
	// } else {
	// 	beego.Error(err)
	// }
	//2.取得客户端用户名
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// uname := ck.Value
	//3.取出用户的权限等级
	// category, err := models.GetCategory(id)
	// beego.Info(username)
	//4.取得客户端用户名
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// }
	// uname := ck.Value
	//5.取出用户的权限等级
	role, _ := checkRole(c.Ctx) //login里的
	// beego.Info(role)
	//6.进行逻辑分析：
	// rolename, _ := strconv.ParseInt(role, 10, 64)
	// if filetype != "pdf" && filetype != "jpg" && filetype != "diary" {
	if role > 1 { //&& uname != category.Author
		// port := strconv.Itoa(c.Ctx.Input.Port())//c.Ctx.Input.Site() + ":" + port +
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// }

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//2.取得客户端用户名
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
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

//如果不带id则取到所有用户
//如果带id，则取一个用户
func (c *UserController) User() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	c.Data["Ip"] = c.Ctx.Input.IP()
	// var categories []*models.AdminCategory
	if id == "" { //如果id为空，则查询类别
		users, err := m.GetUsers()
		if err != nil {
			beego.Error(err)
		}
		c.Data["json"] = &users
		c.ServeJSON()
	} else {
		//pid转成64为
		idNum, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		user := m.GetUserByUserId(idNum)
		if err != nil {
			beego.Error(err)
		}
		// var users1 []*m.User
		users := make([]*m.User, 1)
		users[0] = &user
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

//用户登录查看自己的资料
func (c *UserController) View() {
	// c.Data["IsCategory"] = true
	// c.TplName = "category.tpl"
	c.Data["IsLogin"] = checkAccount(c.Ctx)
	//2.取得客户端用户名
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		c.Data["Uname"] = v.(string)
	}
	// ck, err := c.Ctx.Request.Cookie("uname")
	// if err != nil {
	// 	beego.Error(err)
	// } else {
	// 	c.Data["Uname"] = ck.Value
	// }
	// userid, _ := c.GetInt64("Id")
	// id := c.Ctx.Input.Param("0")这里为何无效？？？？这个需要routers中设置AutoRouter
	// beego.Info(id)
	// userid, _ := strconv.ParseInt(id, 10, 64)

	userid, _ := strconv.ParseInt(c.Input().Get("useid"), 10, 64)
	user := m.GetUserByUserId(userid)
	c.Data["User"] = user
	c.TplName = "admin_user_view.tpl"
}

func (c *UserController) AddUser() {
	// u := m.User{}
	// if err := c.ParseForm(&u); err != nil {
	// 	beego.Error(err.Error)
	// 	return
	// }
	var user m.User
	user.Username = c.Input().Get("username")
	user.Nickname = c.Input().Get("nickname")

	md5Ctx := md5.New()
	md5Ctx.Write([]byte(c.Input().Get("password")))
	cipherStr := md5Ctx.Sum(nil)
	user.Password = hex.EncodeToString(cipherStr)
	user.Repassword = c.Input().Get("repassword")
	user.Email = c.Input().Get("email")
	user.Department = c.Input().Get("department")
	user.Secoffice = c.Input().Get("secoffice")
	user.Ip = c.Input().Get("ip")
	user.Port = c.Input().Get("port")
	statusint, err := strconv.Atoi(c.Input().Get("status"))
	if err != nil {
		beego.Error(err)
	}
	user.Status = statusint
	roleint, err := strconv.Atoi(c.Input().Get("role"))
	if err != nil {
		beego.Error(err)
	}
	user.Role = roleint
	id, err := m.SaveUser(user)
	if err == nil && id > 0 {
		// c.Rsp(true, "Success")
		// return
		c.Data["json"] = "ok"
		c.ServeJSON()
	} else {
		// c.Rsp(false, err.Error())
		beego.Error(err)
		// return
	}
}

// func (c *UserController) UpdateUser() {
// 	u := m.User{}
// 	if err := c.ParseForm(&u); err != nil {
// 		//handle error
// 		// c.Rsp(false, err.Error())
// 		beego.Error(err.Error)
// 		return
// 	}
// 	id, err := m.UpdateUser(&u)
// 	if err == nil && id > 0 {
// 		// c.Rsp(true, "Success")
// 		return
// 	} else {
// 		// c.Rsp(false, err.Error())
// 		beego.Error(err.Error)
// 		return
// 	}

// }
//在线修改保存某个字段
func (c *UserController) UpdateUser() {
	name := c.Input().Get("name")
	value := c.Input().Get("value")
	pk := c.Input().Get("pk")
	id, err := strconv.ParseInt(pk, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	err = m.UpdateUser(id, name, value)
	if err != nil {
		beego.Error(err)
	} else {
		data := "ok!"
		c.Ctx.WriteString(data)
	}

	logs := logs.NewLogger(1000)
	logs.SetLogger("file", `{"filename":"log/engineercms.log"}`)
	logs.EnableFuncCallDepth(true)
	logs.Info(c.Ctx.Input.IP() + " " + "修改保存设计记录" + pk)
	logs.Close()
}

//这个作废，用在线修改
// func (c *UserController) UpdateUser() {
// 	userid := c.Input().Get("userid")
// 	nickname := c.Input().Get("nickname")
// 	email := c.Input().Get("email")
// 	Pwd1 := c.Input().Get("password")
// 	if Pwd1 != "" {
// 		md5Ctx := md5.New()
// 		md5Ctx.Write([]byte(Pwd1))
// 		cipherStr := md5Ctx.Sum(nil)
// 		password := hex.EncodeToString(cipherStr)
// 		err := m.UpdateUser(userid, nickname, email, password) //这里修改
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 	} else {
// 		err := m.UpdateUser(userid, nickname, email, "") //这里修改
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 	}
// 	c.TplName = "user_view.tpl"
// }

func (c *UserController) DeleteUser() {
	Id, _ := c.GetInt64("userid")
	status, err := m.DelUserById(Id)
	if err == nil && status > 0 {
		// c.Rsp(true, "Success")
		c.Redirect("/user/index", 302)
		return
	} else {
		// c.Rsp(false, err.Error())
		beego.Error(err.Error)
		return
	}
}

//用户查看自己，修改密码等
func (c *UserController) GetUserByUsername() {
	_, role := checkprodRole(c.Ctx)
	if role == 1 {
		c.Data["IsAdmin"] = true
	} else if role <= 1 && role > 5 {
		c.Data["IsLogin"] = true
	} else {
		c.Data["IsAdmin"] = false
		c.Data["IsLogin"] = false
	}

	// c.Data["IsProject"] = true
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	// 	c.Data["IsCategory"] = true
	// c.TplName = "category.tpl"
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	//4.取得客户端用户名
	var uname string
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		uname = v.(string)
		c.Data["Uname"] = v.(string)
	}
	if uname == "" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	// uname := v.(string) //ck.Value
	// 4.取出用户的权限等级
	// role, _ := checkRole(c.Ctx) //login里的
	// 5.进行逻辑分析：
	// rolename, err := strconv.ParseInt(role, 10, 64)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// username := c.Input().Get("username")
	// user, err := m.GetUserByUsername(uname)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// beego.Info(user)
	// list, _, _ := m.GetRoleByUsername(uname)
	c.Data["User"] = uname
	// c.Data["Role"] = list
	c.TplName = "user_view.tpl"
}

//用户个人数据，填充table，以便编辑
func (c *UserController) Usermyself() {
	//4.取得客户端用户名
	var uname string
	sess, _ := globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer sess.SessionRelease(c.Ctx.ResponseWriter)
	v := sess.Get("uname")
	if v != nil {
		uname = v.(string)
		c.Data["Uname"] = v.(string)
	}
	if uname == "" {
		route := c.Ctx.Request.URL.String()
		c.Data["Url"] = route
		c.Redirect("/roleerr?url="+route, 302)
		// c.Redirect("/roleerr", 302)
		return
	}
	user, err := m.GetUserByUsername(uname)
	if err != nil {
		beego.Error(err)
	}
	users := make([]*m.User, 1)
	users[0] = &user
	c.Data["json"] = users
	c.ServeJSON()
}

//上传excel文件，导入到数据库
//引用来自category的查看成果类型里的成果
func (c *UserController) ImportUsers() {
	//获取上传的文件
	_, h, err := c.GetFile("usersexcel")
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(h.path)
	// var attachment string
	var path string
	// var filesize int64
	if h != nil {
		//保存附件
		path = ".\\attachment\\" + h.Filename  // 关闭上传的文件，不然的话会出现临时文件不能清除的情况
		err = c.SaveToFile("usersexcel", path) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
	}

	var user m.User
	//读出excel内容写入数据库
	xlFile, err := xlsx.OpenFile(path) //
	if err != nil {
		beego.Error(err)
	}
	for _, sheet := range xlFile.Sheets {
		for i, row := range sheet.Rows {
			if i != 0 { //忽略第一行标题
				// 这里要判断单元格列数，如果超过单元格使用范围的列数，则出错for j := 2; j < 7; j += 5 {
				j := 1
				if len(row.Cells) >= 2 { //总列数，从1开始
					user.Username, err = row.Cells[j].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 3 {
					user.Nickname, err = row.Cells[j+1].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 4 {
					Pwd1, err := row.Cells[j+2].String()
					if err != nil {
						beego.Error(err)
					}

					md5Ctx := md5.New()
					md5Ctx.Write([]byte(Pwd1))
					cipherStr := md5Ctx.Sum(nil)
					user.Password = hex.EncodeToString(cipherStr)
				}
				if len(row.Cells) >= 5 {
					user.Email, err = row.Cells[j+3].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 6 {
					user.Department, err = row.Cells[j+4].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 7 {
					user.Secoffice, err = row.Cells[j+5].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 8 {
					user.Ip, err = row.Cells[j+6].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 9 {
					user.Port, err = row.Cells[j+7].String()
					if err != nil {
						beego.Error(err)
					}
				}
				if len(row.Cells) >= 10 {
					status, err := row.Cells[j+8].String()
					if err != nil {
						beego.Error(err)
					}
					status1, err := strconv.Atoi(status)
					if err != nil {
						beego.Error(err)
					}
					user.Status = status1
				}
				if len(row.Cells) >= 11 {
					role, err := row.Cells[j+9].String()
					if err != nil {
						beego.Error(err)
					}
					role1, err := strconv.Atoi(role)
					if err != nil {
						beego.Error(err)
					}
					user.Role = role1
					user.Lastlogintime = time.Now()
					_, err = m.SaveUser(user) //如果姓名重复，也要返回uid
					if err != nil {
						beego.Error(err)
					}
				}
			}
		}
	}
	//删除附件
	err = os.Remove(path)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = "ok"
	c.ServeJSON()
}

func (this *UserController) Roleerr() {
	// url := this.Input().Get("url")
	url1 := this.Input().Get("url") //这里不支持这样的url，http://192.168.9.13/login?url=/topic/add?id=955&mid=3
	url2 := this.Input().Get("level")
	url3 := this.Input().Get("key")
	var url string
	if url2 == "" {
		url = url1
	} else {
		url = url1 + "&level=" + url2 + "&key=" + url3
	}
	this.Data["Url"] = url
	this.TplName = "role_err.tpl"
}
