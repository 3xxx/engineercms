package models

import (
	// "crypto/md5"
	// "encoding/hex"
	// "errors"
	// "strconv"
	// "fmt"
	// "log"
	"time"
	// "github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
	// . "github.com/beego/admin/src/lib"
	// "github.com/casbin/beego-orm-adapter"
	// "github.com/casbin/casbin"
)

//用户表
type Role struct {
	Id         int64  `PK`
	Rolename   string `orm:"unique"` //这个拼音的简写
	Rolenumber string
	Status     int       `orm:"default(0)" form:"Status" valid:"Range(1,2)"`
	Createtime time.Time `orm:"type(datetime);auto_now_add" `
	Updated    time.Time `orm:"type(datetime);auto_now_add" `
}

type UserRole struct {
	Id     int64 `PK`
	UserId int64
	RoleId int64
}

func init() {
	orm.RegisterModel(new(Role), new(UserRole))
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

//取到一个用户数据，不是数组，所以table无法显示
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
	_, err := o.Insert(userrole)
	if err != nil {
		return err
	}
	return nil
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
	if err != nil {
		return err
	}
	// }
	return err
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
