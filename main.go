package main

import (
	"os"

	_ "github.com/3xxx/engineercms/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//开启orm调试模式
	orm.Debug = true
	//创建附件目录
	os.Mkdir("attachment", os.ModePerm)
	//创建轮播图片存放目录
	os.Mkdir("attachment//carousel", os.ModePerm)
	//自动建表
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
