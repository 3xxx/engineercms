package models

import (
	// "database/sql"
	// beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
	// _ "github.com/go-sql-driver/mysql"
	// "github.com/go-xorm/xorm"
	// _ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
	// "strconv"
	// "strings"
	// "fmt"
	// "log"
	// "os"
	"time"
)

//待办事件
type Todo struct {
	Id        int64  `json:"todoid"`
	ProjectId int64  `json:"projectid"`
	Name      string `json:"todoname"`
	UserId    int64
	Completed bool      `json:"completed"`
	Created   time.Time `orm:"auto_now_add;type(date)"`
	Updated   time.Time `orm:"auto_now_add;type(date)"`
}

//操作记录
type Todolog struct {
	Id      int64
	TodoId  int64
	UserId  int64
	Action  string    //'Finish' : 'Restart'
	Created time.Time `orm:"auto_now_add;type(date)"`
}

func init() {
	orm.RegisterModel(new(Todo), new(Todolog))
}

func TodoCreate(projectid int64, name string, userid int64) (id int64, err error) {
	o := orm.NewOrm()
	todo := &Todo{
		ProjectId: projectid,
		Name:      name,
		UserId:    userid,
		Completed: false,
		Created:   time.Now(),
	}
	id, err = o.Insert(todo)
	return id, err
}

type TodoUser struct {
	Todo `xorm:"extends"`
	User `xorm:"extends"`
}

func GetTodoUser(projectid int64, limit, offset int) ([]*TodoUser, error) {
	users := make([]*TodoUser, 0)
	return users, engine.Table("todo").Join("INNER", "user", "todo.user_id = user.id").Where("todo.project_id=?", projectid).Desc("todo.created").Limit(limit, offset).Find(&users)
}

//修改
func UpdateTodo(todoid int64) error {
	o := orm.NewOrm()
	todo := &Todo{Id: todoid}
	if o.Read(todo) == nil {
		todo.Completed = !todo.Completed
		todo.Updated = time.Now()
		// calendar.Updated = time.Now()
		_, err := o.Update(todo)
		if err != nil {
			return err
		}
	}
	return nil
}

//删除待办
// func DeleteTodo(todoid int64) (int64, error) {
// 	o := orm.NewOrm()
// 	status, err := o.Delete(&Todo{Id: todoid})
// 	return status, err
// }

//删除待办
func DeleteTodo(todoid int64) error {
	o := orm.NewOrm()
	todo := &Todo{Id: todoid}
	if o.Read(todo) == nil {
		_, err := o.Delete(todo)
		if err != nil {
			return err
		}
	}
	return nil
}

// func GetTodos() (todos []*Todo, err error) {
// 	o := orm.NewOrm()
// 	qs := o.QueryTable("Todo")
// 	_, err = qs.Filter("completed", false).All(&todos)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return todos, err
// }
