package main

import (
	_ "engineercms/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//开启orm调试模式
	orm.Debug = true
	//自动建表
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
