package models

import (
	// "crypto/md5"
	// "encoding/hex"
	// "errors"
	"fmt"
	"log"
	// "strconv"
	"time"
	// beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
	// "github.com/beego/beego/v2/adapter/validation"
	// . "github.com/beego/admin/src/lib"
	// "github.com/casbin/beego-orm-adapter"
	// "github.com/casbin/casbin"
	// "github.com/3xxx/engineercms/controllers"
	// "github.com/casbin/casbin"
)

// var e *casbin.Enforcer

//角色表-20201225修改json status为role
type Role struct {
	Id         int64
	Rolename   string `json:"name",orm:"unique"` //这个拼音的简写
	Rolenumber string
	Status     string    `json:"role",orm:"default('0');size(2)"` //,form:"Status",valid:"Range('0','1','2','3','4')"`
	Createtime time.Time `orm:"type(datetime);auto_now_add" `
	Updated    time.Time `orm:"type(datetime);auto_now_add" `
}

type UserRole struct {
	Id     int64
	UserId int64
	RoleId int64 `xorm:"extends"`
}

func init() {
	orm.RegisterModel(new(Role), new(UserRole))
	// _db.CreateTable(&Role{})
	// InsertRole()
}

//添加权限
func SaveRole(role Role) (rid int64, err error) {
	o := orm.NewOrm()
	var role1 Role
	//判断是否有重名
	err = o.QueryTable("role").Filter("rolename", role.Rolename).One(&role1, "Id")
	if err == orm.ErrNoRows { //Filter("tnumber", tnumber).One(topic, "Id")==nil则无法建立
		// 没有找到记录
		rid, err = o.Insert(&role)
		if err != nil {
			return rid, err
		}
	} //else { //应该进行更新操作
	// user1 := &User{Id: user1.Id}
	// 	user1.Username = user.Username
	// 	user1.Nickname = user.Nickname
	// 	user1.Password = user.Password
	// 	user1.Repassword = user.Repassword
	// 	user1.Email = user.Email
	// 	user1.Department = user.Department
	// 	user1.Secoffice = user.Secoffice
	// 	// user1.Remark = user.Remark
	// 	user1.Ip = user.Ip
	// 	user1.Status = user.Status
	// 	user1.Lastlogintime = user.Lastlogintime
	// 	user1.Createtime = time.Now()
	// 	user1.Role = user.Role
	// 	_, err = o.Update(&user1)
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// 	uid = user1.Id
	// }
	return rid, err
}

//取出所有角色
func GetRoles() (roles []*Role, err error) {
	o := orm.NewOrm()
	role := new(Role)
	qs := o.QueryTable(role)
	// var offset int64
	// if page <= 1 {
	// 	offset = 0
	// } else {
	// 	offset = (page - 1) * page_size
	// }
	_, err = qs.All(&roles)
	if err != nil {
		return nil, err
	}
	// count, _ = qs.Count()
	return roles, err
}

//由rolename取得角色
func GetRoleByRolename(name string) (role Role, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Role")
	err = qs.Filter("Rolename", name).One(&role)
	if err != nil {
		return role, err
	}
	return role, err
}

//取到一个角色数据，不是数组，所以table无法显示
func GetRoleByRoleId(roleid int64) (role Role) {
	role = Role{Id: roleid}
	o := orm.NewOrm()
	o.Read(&role) //这里是默认主键查询。=(&user,"Id")
	return role
}

//由用户id取得所拥有的角色
func GetRoleByUserId(id int64) (roles []*UserRole, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("UserRole")
	_, err = qs.Filter("UserId", id).All(&roles)
	if err != nil {
		return nil, err
	}
	return roles, err
}

//由用户id取得所拥有的角色rolename
func GetRolenameByUserId(id int64) (roles []*Role, err error) {
	engine.Table("user_role").Join("INNER", "role", "role.id = user_role.role_id").Where("user_role.user_id = ?", id).Find(&roles)
	return roles, err
}

func UpdateRole(role Role) (err error) {
	o := orm.NewOrm()
	role1 := &Role{Id: role.Id}
	if o.Read(role1) == nil {
		role1.Rolename = role.Rolename
		role1.Rolenumber = role.Rolenumber
		role1.Status = role.Status
		role1.Updated = time.Now()
		_, err := o.Update(role1)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteRole(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Role{Id: Id})
	return status, err
}

//将userid和roleid存入对应数据库
//如果存在，
func AddUserRole(uid, rid int64) error {
	//重复性检查
	o := orm.NewOrm()
	userrole := &UserRole{
		UserId: uid,
		RoleId: rid,
	}

	// 三个返回参数依次为：是否新创建的，对象 Id 值，错误
	//这里用这个出错！！
	// created, id, err := o.ReadOrCreate(&userrole, "UserId", "RoleId")
	// if created {
	// 	fmt.Println("New Insert an object. Id:", id)
	// } else {
	// 	fmt.Println("Get an object. Id:", id)
	// }

	//判断是否有重名
	err := o.QueryTable("UserRole").Filter("UserId", uid).Filter("RoleId", rid).One(userrole, "Id")
	if err == orm.ErrNoRows { //Filter("tnumber", tnumber).One(topic, "Id")==nil则无法建立
		// 没有找到记录
		_, err := o.Insert(userrole)
		if err != nil {
			return err
		}
	} //else {

	return err
}

func DeleteUserRole(uid, rid int64) error {
	o := orm.NewOrm()
	var role UserRole
	qs := o.QueryTable("UserRole")
	_, err := qs.Filter("UserId", uid).Filter("RoleId", rid).All(&role)
	if err != nil {
		return err
	}
	// if o.Read(&merit) == nil {
	_, err = o.Delete(&role) //删除用户角色
	// if err != nil {
	// 	return err
	// }
	// }
	return err
}

func InsertRole() error {
	fmt.Println("insert role ...")
	// r := new(Role)
	var r Role
	r.Rolename = "admin"
	r.Rolenumber = "1"
	r.Status = "1"
	// r.Remark = "I'm a admin role"
	// r.Status = 2
	// r.Title = "Admin role"
	// o.Insert(r)
	id, err := SaveRole(r)
	if err == nil && id > 0 {
		fmt.Println("insert role end")
	} else {
		log.Println(err)
		role, err := GetRoleByRolename(r.Rolename)
		if err != nil {
			log.Println(err)
		}
		id = role.Id
	}

	user, err := GetUserByUsername("admin")
	if err != nil {
		log.Println(err)
	}

	// 将用户admin加入到角色admin里

	// e.AddGroupingPolicy(strconv.FormatInt(user.Id, 10), "role_"+strconv.FormatInt(id, 10))

	//将用户admin加入到角色admin里
	err = AddUserRole(user.Id, id)
	if err != nil {
		log.Println(err)
	}
	//匿名用户角色
	r.Rolename = "anonymous"
	r.Rolenumber = "5"
	r.Status = "1"
	// r.Remark = "I'm a admin role"
	// r.Status = 2
	// r.Title = "Admin role"
	// o.Insert(r)
	id, err = SaveRole(r)
	if err == nil && id > 0 {
		fmt.Println("insert role end")
	} else {
		log.Println(err)
	}

	r.Rolename = "everyone"
	r.Rolenumber = "5"
	r.Status = "1"
	// r.Remark = "I'm a admin role"
	// r.Status = 2
	// r.Title = "Admin role"
	// o.Insert(r)
	id, err = SaveRole(r)
	if err == nil && id > 0 {
		fmt.Println("insert role end")
	} else {
		log.Println(err)
	}

	r.Rolename = "isme"
	r.Rolenumber = "4"
	r.Status = "1"
	// r.Remark = "I'm a admin role"
	// r.Status = 2
	// r.Title = "Admin role"
	// o.Insert(r)
	id, err = SaveRole(r)
	if err == nil && id > 0 {
		fmt.Println("insert role end")
	} else {
		log.Println(err)
	}
	return err
	// fmt.Println("insert role end")
}

//由角色id、action和项目id，取得所有的路径
// func GetPermissions(roleid, projectid, action) (paths []*CasbinRule, err error) {
// 	o := orm.NewOrm()
// 	qs := o.QueryTable("casbin_rule")
// 	_, err = qs.Filter("PType", "p").Filter("v0", roleid).Filter("v1__contains", projectid).Filter("v2", action).All(&paths)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return roles, err
// }
