package models

import (
	"errors"
	"github.com/3xxx/engineercms/conf"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type TeamMember struct {
	TeamMemberId int `orm:"column(team_member_id);pk;auto;unique;" json:"team_member_id"`
	TeamId       int `orm:"column(team_id);type(int)" json:"team_id"`
	MemberId     int `orm:"column(member_id);type(int)" json:"member_id"`
	// RoleId 角色：0 创始人(创始人不能被移除) / 1 管理员/2 编辑者/3 观察者
	RoleId   conf.BookRole `orm:"column(role_id);type(int)" json:"role_id"`
	RoleName string        `orm:"-" json:"role_name"`
	Account  string        `orm:"-" json:"account"`
	RealName string        `orm:"-" json:"real_name"`
	Avatar   string        `orm:"-" json:"avatar"`
}

// TableName 获取对应数据库表名.
func (m *TeamMember) TableName() string {
	return "team_member"
}
func (m *TeamMember) TableNameWithPrefix() string {
	return conf.GetDatabasePrefix() + m.TableName()
}

// TableEngine 获取数据使用的引擎.
func (m *TeamMember) TableEngine() string {
	return "INNODB"
}

// 联合唯一键
func (m *TeamMember) TableUnique() [][]string {
	return [][]string{{"team_id", "member_id"}}
}

func (m *TeamMember) QueryTable() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m.TableNameWithPrefix())
}

func NewTeamMember() *TeamMember {
	return &TeamMember{}
}

func (m *TeamMember) First(id int, cols ...string) (*TeamMember, error) {
	if id <= 0 {
		return nil, errors.New("参数错误")
	}
	o := orm.NewOrm()

	err := o.QueryTable(m.TableNameWithPrefix()).Filter("team_member_id", id).One(m, cols...)

	if err != nil && err != orm.ErrNoRows {
		beego.Error("查询团队成员错误 ->", err)
	}

	return m.Include(), err
}

func (m *TeamMember) ChangeRoleId(teamId int, memberId int, roleId conf.BookRole) (member *TeamMember, err error) {

	if teamId <= 0 || memberId <= 0 || roleId <= 0 || roleId > conf.BookObserver {
		return nil, ErrInvalidParameter
	}
	o := orm.NewOrm()

	err = o.QueryTable(m.TableNameWithPrefix()).Filter("team_id", teamId).Filter("member_id", memberId).OrderBy("-team_member_id").One(m)

	if err != nil {
		beego.Error("查询团队用户时失败 ->", err)
		return m, err
	}
	m.RoleId = roleId

	err = m.Save("role_id")

	if err == nil {
		m.Include()
	}
	return m, err
}

//查询团队中指定的用户.
func (m *TeamMember) FindFirst(teamId, memberId int) (*TeamMember, error) {
	if teamId <= 0 || memberId <= 0 {
		return nil, ErrInvalidParameter
	}
	o := orm.NewOrm()
	err := o.QueryTable(m.TableNameWithPrefix()).Filter("team_id", teamId).Filter("member_id", memberId).One(m)

	if err != nil {
		beego.Error("查询团队用户失败 ->", err)
		return nil, err
	}
	return m.Include(), nil
}

//更新或插入团队用户.
func (m *TeamMember) Save(cols ...string) (err error) {

	if m.TeamId <= 0 {
		return errors.New("团队不能为空")
	}
	if m.MemberId <= 0 {
		return errors.New("用户不能为空")
	}

	o := orm.NewOrm()

	if !o.QueryTable(NewTeam().TableNameWithPrefix()).Filter("team_id", m.TeamId).Exist() {
		return errors.New("团队不存在")
	}
	if !o.QueryTable(NewMember()).Filter("member_id", m.MemberId).Filter("status", 0).Exist() {
		return errors.New("用户不存在或已禁用")
	}

	if m.TeamMemberId <= 0 {
		if o.QueryTable(m.TableNameWithPrefix()).Filter("team_id", m.TeamId).Filter("member_id", m.MemberId).Exist() {
			return errors.New("团队中已存在该用户")
		}
		_, err = o.Insert(m)
	} else {
		_, err = o.Update(m, cols...)
	}
	if err != nil {
		beego.Error("在保存团队时出错 ->", err)
	}
	return
}

//删除一个团队用户.
func (m *TeamMember) Delete(id int) (err error) {

	if id <= 0 {
		return ErrInvalidParameter
	}
	_, err = orm.NewOrm().QueryTable(m.TableNameWithPrefix()).Filter("team_member_id", id).Delete()

	if err != nil {
		beego.Error("删除团队用户时出错 ->", err)
	}
	return
}

//分页查询团队用户.
func (m *TeamMember) FindToPager(teamId, pageIndex, pageSize int) (list []*TeamMember, totalCount int, err error) {
	if teamId <= 0 {
		err = ErrInvalidParameter
		return
	}
	offset := (pageIndex - 1) * pageSize

	o := orm.NewOrm()

	_, err = o.QueryTable(m.TableNameWithPrefix()).Filter("team_id", teamId).Offset(offset).Limit(pageSize).All(&list)

	if err != nil {
		if err != orm.ErrNoRows {
			beego.Error("查询团队成员失败 ->", err)
		}
		return
	}
	c, err := o.QueryTable(m.TableNameWithPrefix()).Filter("team_id", teamId).Count()

	if err != nil {
		return
	}
	totalCount = int(c)

	//将来优化
	for _, item := range list {
		item.Include()
	}
	return
}

//查询关联数据.
func (m *TeamMember) Include() *TeamMember {

	if member, err := NewMember().Find(m.MemberId, "account", "real_name", "avatar"); err == nil {
		m.Account = member.Account
		m.RealName = member.RealName
		m.Avatar = member.Avatar
	}
	if m.RoleId == 0 {
		m.RoleName = "创始人"
	} else if m.RoleId == 1 {
		m.RoleName = "管理员"
	} else if m.RoleId == 2 {
		m.RoleName = "编辑者"
	} else if m.RoleId == 3 {
		m.RoleName = "观察者"
	}
	return m
}

//查询未加入团队的用户。
func (m *TeamMember) FindNotJoinMemberByAccount(teamId int, account string, limit int) (*SelectMemberResult, error) {
	if teamId <= 0 {
		return nil, ErrInvalidParameter
	}
	o := orm.NewOrm()

	sql := `select member.member_id,member.account,team.team_member_id
from md_members as member 
  left join md_team_member as team on team.team_id = ? and member.member_id = team.member_id
  where member.account like ? AND team_member_id IS NULL
  order by member.member_id desc 
limit ?;`

	members := make([]*Member, 0)

	_, err := o.Raw(sql, teamId, "%"+account+"%", limit).QueryRows(&members)

	if err != nil {
		beego.Error("查询团队用户时出错 ->", err)
		return nil, err
	}

	result := SelectMemberResult{}
	items := make([]KeyValueItem, 0)

	for _, member := range members {
		item := KeyValueItem{}
		item.Id = member.MemberId
		item.Text = member.Account
		items = append(items, item)
	}
	result.Result = items

	return &result, err
}

func (m *TeamMember) FindByBookIdAndMemberId(bookId, memberId int) (*TeamMember, error) {
	if bookId <= 0 || memberId <= 0 {
		return nil, ErrInvalidParameter
	}
	//一个用户可能在多个团队中，且一个项目可能有多个团队参与。因此需要查询用户最大权限。
	sql := `select *
from md_team_member as team
where team.team_id in (select rel.team_id from md_team_relationship as rel where rel.book_id = ?) 
and team.member_id = ? order by team.role_id asc limit 1;`

	o := orm.NewOrm()

	err := o.Raw(sql, bookId, memberId).QueryRow(m)

	if err != nil {
		beego.Error("查询用户项目所在团队失败 ->bookId=", bookId, " memberId=", memberId, err)
		return nil, err
	}
	return m, nil
}
