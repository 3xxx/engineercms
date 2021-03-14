package controllers

import (
	// "encoding/json"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	m "github.com/3xxx/engineercms/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/casbin/beego-orm-adapter"
	"github.com/casbin/casbin"
	"github.com/casbin/xorm-adapter"
	_ "github.com/mattn/go-sqlite3"
	"path"
	"regexp"
	"strconv"
	"strings"
	// "engineercms/controllers/validator"
	// "github.com/astaxie/beego/context"
	// "github.com/asofdate/hauth/core/groupcache"
	// "github.com/asofdate/hauth/core/hrpc"
	// "github.com/asofdate/hauth/core/models"
	// "github.com/asofdate/hauth/utils"
	// "github.com/asofdate/hauth/utils/hret"
	// "github.com/asofdate/hauth/utils/i18n"
	// "github.com/asofdate/hauth/utils/jwt"
	// "github.com/asofdate/hauth/utils/logs"
	// "github.com/asofdate/hauth/utils/validator"
)

type RoleController struct {
	beego.Controller
}

type Userrole struct {
	Id         int64
	Rolename   string `json:"name"`
	Rolenumber string
	Status     string `json:"status"` //20201225这里是否要修改为role
	Level      string
}

type Tree struct {
	Id    int64  `json:"id"`
	Nodes []Tree `json:"nodes"`
}

// type CasbinRule struct {
// 	Id    int
// 	PType string
// 	V0    string
// 	V1    string
// 	V2    string
// 	V3    string
// 	V4    string
// 	V5    string
// }

// type RoleController struct {
// 	models models.RoleModel
// }

// var RoleCtl = &RoleController{
// 	models.RoleModel{},
// }

// 2019/5/27
var e *casbin.Enforcer

// var a *beegoormadapter.Adapter

func init() {
	var dns string
	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	// 数据库名称
	db_name := beego.AppConfig.String("db_name")
	db_path := beego.AppConfig.String("db_path")
	db_sslmode := beego.AppConfig.String("db_sslmode")
	switch db_type {
	case "mysql":
		// dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/", db_user, db_pass, db_host, db_port, db_name)
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/", db_user, db_pass, db_host, db_port)
		// a = beegoormadapter.NewAdapter("mysql", dns)
		//因为beegoormadapter里需要手动调整mysql和sqlite，所以，如果是mysql，就用xorm
		// Initialize a Xorm adapter and use it in a Casbin enforcer:
		// The adapter will use the MySQL database named "casbin".
		// If it doesn't exist, the adapter will create it automatically.
		// a := xormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/") // Your driver and data source.
		a := xormadapter.NewAdapter("mysql", dns)
		e = casbin.NewEnforcer("conf/rbac_model.conf", a)
		break
	case "postgres":
		// orm.RegisterDriver("postgres", orm.DRPostgres)
		dns = fmt.Sprintf("dbname=%s host=%s  user=%s  password=%s  port=%s  sslmode=%s", db_name, db_host, db_user, db_pass, db_port, db_sslmode)
		a := beegoormadapter.NewAdapter("postgres", dns)
		e = casbin.NewEnforcer("conf/rbac_model.conf", a)
		break
	case "sqlite3":
		if db_path == "" {
			db_path = "./"
		}
		dns = fmt.Sprintf("%s%s.db", db_path, db_name)
		// 注册casbin
		a := beegoormadapter.NewAdapter("sqlite3", dns, true)
		e = casbin.NewEnforcer("conf/rbac_model.conf", a)
		break
	default:
		beego.Critical("Database driver is not allowed:", db_type)
	}

	// Initialize a Beego ORM adapter and use it in a Casbin enforcer:
	// The adapter will use the MySQL database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	// a := beegoormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/") // Your driver and data source.
	// a = beegoormadapter.NewAdapter("sqlite3", "database/engineer.db", true) // Your driver and data source.
	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := beegoormadapter.NewAdapter("mysql", "mysql_username:mysql_password@tcp(127.0.0.1:3306)/abc", true)
	// e := casbin.NewEnforcer("examples/rbac_model.conf", a)
	// e = casbin.NewEnforcer("conf/rbac_model.conf", a)
	// Load the policy from DB.
	e.LoadPolicy()
	// p, admin, /*, *
	// p, anonymous, /login, *
	// p, member, /logout, *
	// p, member, /member/*, *
	// e.AddPolicy("role_admin", "/*", "*", ".*")

	// e.AddPolicy("role_anonymous", "/login", "*", ".*")
	// e.AddPolicy("role_everyone", "/login", "*", ".*")
	// e.AddPolicy("role_member", "/logout", "*", ".*")
	// e.AddPolicy("role_member", "/member", "*", ".*")
	beego.Info("insert user ...")
	// u := new(User)
	var u m.User
	u.Username = "admin"
	u.Nickname = "Hotqin888"
	Pwd1 := "admin"
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(Pwd1))
	cipherStr := md5Ctx.Sum(nil)
	u.Password = hex.EncodeToString(cipherStr)
	// u.Password = Pwdhash("admin")
	u.Email = "504284@qq.com"
	u.Remark = "I'm admin"
	u.Status = 1
	u.Role = "1"
	id, err := m.SaveUser(u)
	// o = orm.NewOrm()
	// o.Insert(u)
	// fmt.Println("insert user end")
	if err == nil && id > 0 {
		beego.Info("insert user end")
	} else {
		beego.Info(err)
	}

	beego.Info("insert role ...")
	// r := new(Role)
	var r m.Role
	r.Rolename = "admin"
	r.Rolenumber = "1"
	r.Status = "0"

	id, err = m.SaveRole(r)
	if err == nil && id > 0 {
		beego.Info("insert role end")
	} else {
		beego.Error(err)
		//重新获取roleid
		role, err := m.GetRoleByRolename("admin")
		if err != nil {
			beego.Error(err)
		} else {
			id = role.Id
		}
	}
	user_admin, err := m.GetUserByUsername("admin")
	if err != nil {
		beego.Error(err)
	}
	//将用户admin加入到角色admin里
	e.AddGroupingPolicy(strconv.FormatInt(user_admin.Id, 10), "role_"+strconv.FormatInt(id, 10))
	//添加admin角色的权限/*
	e.AddPolicy("role_"+strconv.FormatInt(id, 10), "/*", "*", ".*")

	//匿名用户角色
	r.Rolename = "anonymous"
	r.Rolenumber = "5"
	r.Status = "0"
	// r.Remark = "I'm a admin role"
	// r.Status = 2
	// r.Title = "Admin role"
	// o.Insert(r)
	id, err = m.SaveRole(r)
	if err == nil && id > 0 {
		beego.Info("insert role end")
	} else {
		beego.Error(err)
		//重新获取roleid
		role, err := m.GetRoleByRolename("anonymous")
		if err != nil {
			beego.Error(err)
		} else {
			id = role.Id
		}
	}
	//添加anonymous角色的权限/*
	e.AddPolicy("role_"+strconv.FormatInt(id, 10), "/login", "*", ".*")

	r.Rolename = "everyone"
	r.Rolenumber = "5"
	r.Status = "0"
	// r.Remark = "I'm a admin role"
	// r.Status = 2
	// r.Title = "Admin role"
	// o.Insert(r)
	id, err = m.SaveRole(r)
	if err == nil && id > 0 {
		beego.Info("insert role end")
	} else {
		beego.Error(err)
	}

	r.Rolename = "isme"
	r.Rolenumber = "4"
	r.Status = "0"
	// r.Remark = "I'm a admin role"
	// r.Status = 2
	// r.Title = "Admin role"
	// o.Insert(r)
	id, err = m.SaveRole(r)
	if err == nil && id > 0 {
		beego.Info("insert role end")
	} else {
		beego.Error(err)
	}
}

func (c *RoleController) Test() {
	// Check the permission.
	//请求的资源v1/v2/aaa.jpg
	//得到资源的扩展名Suffix-jpg，输入enforce中间那个(?i:pdf)不分大小写
	beego.Info(e.Enforce("alice", "/v1/v2/aaa.jpg", "write", "jpg"))
	beego.Info(e.Enforce("bob", "/v1/v2/aaa.PDF", "delete", "PDF"))
	beego.Info(e.Enforce("bob", "/v1/v2/aaa.jpg", "write", "jpg"))
	// beego.Info(e.Enforce("bob", "/v1/aaa.pdf", "read", "pdf"))
	// beego.Info(e.Enforce("bob", "/v1/v2/aaa.dwg", "read", "dwg"))
	// beego.Info(e.Enforce("bob", "/v1/v2/aaa.pdf", "write", "pdf"))
	beego.Info(e.Enforce("bob", "/v1/v2/aaa.ttt", "read", "ttt")) //任意扩展名
	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)
	// Save the policy back to DB.
	// e.SavePolicy()
}

// swagger:operation GET /v1/auth/role/page StaticFiles RoleController
//
// 角色管理页面
//
// 如果用户被授权访问角色管理页面,则系统返回角色管理页面内容,否则返回404错误
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
// func (RoleController) Page(ctx *context.Context) {
// 	ctx.Request.ParseForm()
// 	if !hrpc.BasicAuth(ctx.Request) {
// 		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
// 		return
// 	}

// 	rst, err := groupcache.GetStaticFile("AsofdateRolePage")
// 	if err != nil {
// 		hret.Error(ctx.ResponseWriter, 404, i18n.PageNotFound(ctx.Request))
// 		return
// 	}
// 	ctx.ResponseWriter.Write(rst)
// }

// swagger:operation GET /v1/auth/role/get RoleController RoleController
//
// 查询角色信息
//
// 查询指定域中的角色信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
//这个作废，用下面的从casbin中获取用户角色
// func (c *RoleController) Get() {
// 	id := c.Ctx.Input.Param(":id")
// 	c.Data["Id"] = id
// 	c.Data["Ip"] = c.Ctx.Input.IP()
// 	// if id == "" { //如果id为空，则查询
// 	roles, err := m.GetRoles()
// 	if err != nil {
// 		beego.Error(err)
// 	}

// 	if id != "" {
// 		//pid转成64为
// 		idNum, err := strconv.ParseInt(id, 10, 64)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		//查出用户的角色，处于勾选状态
// 		userroles, err := m.GetRoleByUserId(idNum)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		userrole := make([]Userrole, 0)
// 		var level string
// 		level = "2"
// 		for _, v1 := range roles {
// 			for _, v2 := range userroles {
// 				if v2.RoleId == v1.Id {
// 					level = "1"
// 				}
// 			}
// 			aa := make([]Userrole, 1)
// 			aa[0].Id = v1.Id
// 			aa[0].Rolename = v1.Rolename
// 			aa[0].Rolenumber = v1.Rolenumber
// 			aa[0].Level = level
// 			userrole = append(userrole, aa...)
// 			aa = make([]Userrole, 0)
// 			level = "2"
// 		}
// 		c.Data["json"] = userrole
// 		c.ServeJSON()
// 	}
// 	c.Data["json"] = roles
// 	c.ServeJSON()
// }

func (c *RoleController) Get() {
	id := c.Ctx.Input.Param(":id")
	c.Data["Id"] = id
	c.Data["Ip"] = c.Ctx.Input.IP()
	// if id == "" { //如果id为空，则查询
	roles, err := m.GetRoles()
	if err != nil {
		beego.Error(err)
	}
	//如果设置了role,用于onlyoffice的权限设置
	role := c.Input().Get("role")
	if role != "" {
		// roleint, err := strconv.Atoi(role)
		// if err != nil {
		// 	beego.Error(err)
		// }
		for _, v := range roles {
			v.Status = role
		}
	}

	if id != "" { //如果选中了用户，则显示用户所具有的角色
		//pid转成64为
		// idNum, err := strconv.ParseInt(id, 10, 64)
		// if err != nil {
		// 	beego.Error(err)
		// }
		//查出用户的角色，处于勾选状态，来自casbin\rbac_api.go
		userroles := e.GetRolesForUser(id)
		userrole := make([]Userrole, 0)
		var level string
		level = "2"
		for _, v1 := range roles {
			for _, v2 := range userroles {
				ridNum, err := strconv.ParseInt(strings.Replace(v2, "role_", "", -1), 10, 64)
				if err != nil {
					beego.Error(err)
				}
				if ridNum == v1.Id {
					level = "1" //if (row.Level === "1") checked: true
				}
			}
			aa := make([]Userrole, 1)
			aa[0].Id = v1.Id
			aa[0].Rolename = v1.Rolename
			aa[0].Rolenumber = v1.Rolenumber
			aa[0].Level = level
			aa[0].Status = v1.Status
			userrole = append(userrole, aa...)
			aa = make([]Userrole, 0)
			level = "2"
		}
		c.Data["json"] = userrole //用户所具有的角色，勾选
		c.ServeJSON()
	} else {
		c.Data["json"] = roles //角色列表
		c.ServeJSON()
	}
}

// swagger:operation POST /v1/auth/role/post RoleController RoleController
//
// 新增角色信息
//
// 在某个指定的域中,新增角色信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success

//添加角色
func (c *RoleController) Post() {
	// u := m.Role{}
	// if err := c.ParseForm(&u); err != nil {
	// 	beego.Error(err.Error)
	// 	return
	// }
	var role m.Role
	role.Rolename = c.Input().Get("rolename")
	role.Rolenumber = c.Input().Get("rolenumber")

	// statusint, err := strconv.Atoi(c.Input().Get("status"))
	// if err != nil {
	// 	beego.Error(err)
	// }
	role.Status = c.Input().Get("status")

	id, err := m.SaveRole(role)
	if err == nil && id > 0 {
		// c.Rsp(true, "Success")
		// return
		c.Data["json"] = "ok"
		c.ServeJSON()
	} else {
		// c.Rsp(false, err.Error())
		beego.Error(err)
		c.Data["json"] = "wrong"
		c.ServeJSON()
		// return
	}
}

//向userid里添加权限——这个作废，用casbin的
// func (c *RoleController) UserRole() {
// 	uid := c.GetString("uid") //secofficeid
// 	//id转成64位
// 	uidNum, err := strconv.ParseInt(uid, 10, 64)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	//取出所有uidnum的role
// 	userroles, err := m.GetRoleByUserId(uidNum)
// 	if err != nil {
// 		beego.Error(err)
// 	}

// 	ids := c.GetString("ids") //roleid
// 	// beego.Info(ids)
// 	array := strings.Split(ids, ",")
// 	// beego.Info(array)
// 	bool := false
// 	for _, v1 := range array {
// 		// pid = strconv.FormatInt(v1, 10)
// 		//id转成64位
// 		idNum, err := strconv.ParseInt(v1, 10, 64)
// 		if err != nil {
// 			beego.Error(err)
// 		}
// 		for _, v2 := range userroles {
// 			//没有找到则插入
// 			if v2.RoleId == idNum {
// 				bool = true
// 			}
// 		}
// 		if bool == false {
// 			//存入数据库
// 			err = m.AddUserRole(uidNum, idNum)
// 			if err != nil {
// 				beego.Error(err)
// 			}
// 			// beego.Info(uidNum)
// 			// beego.Info(idNum)
// 		}
// 		bool = false
// 	}

// 	for _, v3 := range userroles {
// 		for _, v4 := range array {
// 			//id转成64位
// 			idNum, err := strconv.ParseInt(v4, 10, 64)
// 			if err != nil {
// 				beego.Error(err)
// 			}
// 			//没有找到则删除
// 			if v3.RoleId == idNum {
// 				bool = true
// 			}
// 		}

// 		if bool == false {
// 			//删除数据库
// 			err = m.DeleteUserRole(uidNum, v3.RoleId)
// 			if err != nil {
// 				beego.Error(err)
// 			}
// 		}
// 		bool = false
// 	}
// 	if err != nil {
// 		beego.Error(err)
// 	} else {
// 		c.Data["json"] = "ok"
// 		c.ServeJSON()
// 	}
// }

//AddPolicy(sec string, ptype string, rule []string)
//添加用户角色
//先删除用户所有角色
func (c *RoleController) UserRole() {
	//要支持批量分配角色，循环用户id
	uid := c.GetString("uid") //secofficeid
	//先删除用户的权限
	e.DeleteRolesForUser(uid) //数据库没有删掉！
	//删除数据库中角色中的用户
	// o := orm.NewOrm()
	// qs := o.QueryTable("casbin_rule")
	// _, err := qs.Filter("PType", "g").Filter("v0", uid).Delete()
	// if err != nil {
	// 	beego.Error(err)
	// }
	//再添加，如果没有选择，相当于删除了全部角色
	ids := c.GetString("ids") //roleid
	if ids != "" {
		array := strings.Split(ids, ",")
		// var rule []string
		for _, v1 := range array {
			// rule = append(rule, uid, v1)
			// e.AddPolicy(uid, v1)
			e.AddGroupingPolicy(uid, "role_"+v1) //management_api.go
			//应该用AddRoleForUser()//rbac_api.go
			// rule = make([]string, 0)
		}
		// a.SavePolicy(e.GetModel())//autosave默认是true
		// 	[{0 p 12 1    } {0 g 8 1    } {0 g 7 1
		//    } {0 g 7 2    } {0 g 5 1    } {0 g 5 2    }]
		// lines := [7][4]string{{"0", "p", "100", "1"}, {"0", "p", "101", "1"}}
		// _, err := a.o.InsertMulti(len(lines), lines)
		// return err
	}
	c.Data["json"] = "ok"
	c.ServeJSON()
}

//给角色赋项目目录的权限
//先删除角色对于这个项目的所有权限
func (c *RoleController) RolePermission() {
	var success bool
	var nodeidint int
	var projurl, action, suf1, suf string
	var err error
	roleids := c.GetString("roleids")
	rolearray := strings.Split(roleids, ",")
	// beego.Info(rolearray)
	permissionids := c.GetString("permissionids")
	permissionarray := strings.Split(permissionids, ",")
	switch permissionarray[0] {
	case "添加成果":
		action = "POST"
	case "编辑成果":
		action = "PUT"
	case "删除成果":
		action = "DELETE"
	case "读取成果":
		action = "GET"
	}
	// beego.Info(permissionarray)
	sufids := c.GetString("sufids")
	sufarray := strings.Split(sufids, ",")
	switch sufids {
	case "任意":
		suf = ".*"
	case "":
		suf = "(?i:PDF)"
	case "PDF":
		suf = "(?i:PDF)"
	}
	treeids := c.GetString("treeids") //项目目录id，25001,25002
	treearray := strings.Split(treeids, ",")
	// beego.Info(treearray)
	treenodeids := c.GetString("treenodeids") //项目目录的nodeid 0.0.0-0.0.1-0.1.0-0.1.0
	treenodearray := strings.Split(treenodeids, ",")
	projectid := c.GetString("projid")
	// beego.Info(treenodearray)
	// treeids := c.GetString("tree")
	//json字符串解析到结构体，以便进行追加
	// var tree []Tree
	// err := json.Unmarshal([]byte(treeids), &tree)
	// if err != nil {
	// 	beego.Error(err)
	// }

	//取出项目目录的顶级
	var nodesid, nodesids []string
	// beego.Info(len(treenodearray))
	if len(treenodearray) > 1 {
		nodesids, err = highest(treenodearray, nodesid, 0)
		if err != nil {
			beego.Error(err)
		}
	} else {
		nodesids = []string{"0"} //append(nodesids, "0")
	}
	// beego.Info(nodesids)

	//删除这些角色、项目id、权限的全部权限
	for _, v1 := range rolearray {
		// var paths []beegoormadapter.CasbinRule
		o := orm.NewOrm()
		qs := o.QueryTable("casbin_rule")
		if action == "GET" {
			_, err := qs.Filter("PType", "p").Filter("v0", "role_"+v1).Filter("v1__contains", "/"+projectid+"/").Filter("v2", action).Filter("v3", suf).Delete()
			if err != nil {
				beego.Error(err)
			}
		} else {
			_, err := qs.Filter("PType", "p").Filter("v0", "role_"+v1).Filter("v1__contains", "/"+projectid+"/").Filter("v2", action).Delete()
			if err != nil {
				beego.Error(err)
			}
		}
	}

	e.LoadPolicy() //重载权限
	// e.RemoveFilteredPolicy(1, "/onlyoffice/"+strconv.FormatInt(attachments[0].Id, 10))
	if treeids != "" {
		for _, v1 := range rolearray {
			for _, v2 := range permissionarray {
				//定义读取、添加、修改、删除
				switch v2 {
				case "添加成果":
					action = "POST"
					suf = ".*"
				case "编辑成果":
					action = "PUT"
					suf = ".*"
				case "删除成果":
					action = "DELETE"
					suf = ".*"
				case "读取成果":
					action = "GET"
					for i, v4 := range sufarray {
						if v4 == "任意" {
							suf = ".*"
							break
						} else if v4 == "" { //用户没展开则读取不到table4的select
							suf = "(?i:PDF)"
							break
						} else {
							suf1 = "(?i:" + v4 + ")"
							if i == 0 {
								suf = suf1
							} else {
								suf = suf + "," + suf1
							}
						}
					}
				}

				for _, v3 := range nodesids {
					nodeidint, err = strconv.Atoi(v3)
					if err != nil {
						beego.Error(err)
					}
					//id转成64位
					pidNum, err := strconv.ParseInt(treearray[nodeidint], 10, 64)
					if err != nil {
						beego.Error(err)
					}

					//根据projid取出路径
					proj, err := m.GetProj(pidNum)
					if err != nil {
						beego.Error(err)
					}
					if proj.ParentIdPath == "" || proj.ParentIdPath == "$#" {
						projurl = "/" + strconv.FormatInt(proj.Id, 10) + "/*"
					} else {
						// projurl = "/" + strings.Replace(proj.ParentIdPath, "-", "/", -1) + "/" + treearray[nodeidint] + "/*"
						projurl = "/" + strings.Replace(strings.Replace(proj.ParentIdPath, "#", "/", -1), "$", "", -1) + strconv.FormatInt(proj.Id, 10) + "/*"
					}
					// beego.Info(v1)
					// beego.Info(projurl)
					// beego.Info(action)
					// beego.Info(suf)
					sufarray := strings.Split(suf, ",")
					for _, v5 := range sufarray {
						success = e.AddPolicy("role_"+v1, projurl, action, v5) //来自casbin\management_api.go
						//这里应该用AddPermissionForUser()，来自casbin\rbac_api.go
					}
				}
			}
		}
	} else {
		success = true
	}
	// e.LoadPolicy() //重载权限

	if success == true {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "wrong"
	}
	c.ServeJSON()
}

//迭代查出最高级的树状目录
//nodesid是数组的序号
//nodeid是节点号：0.0.1   0.0.1.0
func highest(nodeid, nodesid []string, i int) (nodesid1 []string, err error) {
	if i == 0 {
		nodesid = append(nodesid, "0")
	}
	var i1 int
	for i1 = i; i1 < len(nodeid)-1; i1++ {
		matched, err := regexp.MatchString("(?i:"+nodeid[i]+")", nodeid[i1+1])
		// fmt.Println(matched)
		if err != nil {
			beego.Error(err)
		}
		if !matched {
			i = i1 + 1
			nodesid = append(nodesid, strconv.Itoa(i1+1))
			break
		} else {
			if i == len(nodeid)-2 {
				return nodesid, err
			}
		}
	}
	if i1 < len(nodeid)-1 {
		nodesid, err = highest(nodeid, nodesid, i)
	}
	return nodesid, err
}

//查询角色所具有的权限对应的项目目录
func (c *RoleController) GetRolePermission() {
	roleid := c.GetString("roleid") //角色id
	action := c.GetString("action")
	projectid := c.GetString("projectid")
	sufids := c.GetString("sufids") //扩展名
	// beego.Info(sufids)
	var suf string
	switch sufids {
	case "任意":
		suf = ".*"
	case "":
		suf = "(?i:PDF)"
	case "PDF":
		suf = "(?i:PDF)"
	}
	// beego.Info(suf)
	// beego.Info(roleid)
	// beego.Info(action)
	// beego.Info(projectid)

	// myRes := e.GetPermissionsForUser(roleid)
	// beego.Info(myRes)

	// 	2018/01/03 21:42:15 [I] [roleControllers.go:543] [[1 /25001/* POST .*] [1 /25001
	// /* PUT .*] [1 /25001/* DELETE .*] [1 /25001/* GET .*] [1 /25001/25003/* GET (?i:
	// PDF)] [1 /25001/25002/25013/* GET (?i:PDF)] [1 /25001/25002/25012/* GET (?i:PDF)
	// ] [1 /25001/25002/25011/* GET (?i:PDF)] [1 /25001/* GET (?i:PDF)] [1 /25001/2500
	// 4/* POST .*]]
	// Permissions, err := models.GetPermissions(roleid,projectid,action)
	// if err != nil {
	// 	beego.Error(err)
	// }
	var paths []beegoormadapter.CasbinRule
	o := orm.NewOrm()
	qs := o.QueryTable("casbin_rule")
	if action == "GET" || action == "" {
		_, err := qs.Filter("PType", "p").Filter("v0", "role_"+roleid).Filter("v1__contains", "/"+projectid+"/").Filter("v2", "GET").Filter("v3", suf).All(&paths)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(paths)
	} else {
		_, err := qs.Filter("PType", "p").Filter("v0", "role_"+roleid).Filter("v1__contains", "/"+projectid+"/").Filter("v2", action).All(&paths)
		if err != nil {
			beego.Error(err)
		}
	}
	// beego.Info(paths)
	var projids []string
	for _, v1 := range paths {
		projid := strings.Replace(v1.V1, "/*", "", -1)
		projids = append(projids, path.Base(projid))
	}
	// beego.Info(projids)
	c.Data["json"] = projids
	c.ServeJSON()
}

// swagger:operation POST /v1/auth/role/delete RoleController RoleController
//
// 删除角色信息
//
// 删除某个指定域中的角色信息
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (c *RoleController) Delete() {
	roleids := c.GetString("ids")
	rolearray := strings.Split(roleids, ",")
	for _, v1 := range rolearray {
		// rid, _ := c.GetInt64("roleid")
		//id转成64位
		idNum, err := strconv.ParseInt(v1, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		_, err = m.DeleteRole(idNum)
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	}
}

// swagger:operation PUT /v1/auth/role/put RoleController RoleController
//
// 更新角色信息
//
// 更新某个域中的角色信息,角色编码不能更新
//
// ---
// produces:
// - application/json
// - application/xml
// - text/xml
// - text/html
// parameters:
// - name: domain_id
//   in: query
//   description: domain code number
//   required: true
//   type: string
//   format:
// responses:
//   '200':
//     description: success
func (c *RoleController) Update() {
	var role m.Role
	roleid := c.Input().Get("roleid")
	idNum, err := strconv.ParseInt(roleid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	role.Id = idNum
	role.Rolename = c.Input().Get("rolename")
	role.Rolenumber = c.Input().Get("rolenumber")
	// statusint, err := strconv.Atoi(c.Input().Get("status"))
	// if err != nil {
	// 	beego.Error(err)
	// }
	role.Status = c.Input().Get("status")
	err = m.UpdateRole(role)
	if err == nil {
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
