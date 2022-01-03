// //定义一个工具包，用来管理gorm数据库连接池的初始化工作。
package models

// import (
// 	// "github.com/beego/beego/v2/adapter/orm"
// 	"fmt"
// 	beego "github.com/beego/beego/v2/adapter"
// 	"github.com/jinzhu/gorm"
// 	// _ "github.com/jinzhu/gorm/dialects/sqlite"
// )

// //定义全局的db对象，我们执行数据库操作主要通过他实现。
// var _db *gorm.DB

// //包初始化函数，golang特性，每个包初始化的时候会自动执行init函数，这里用来初始化gorm。
// func init() {
// 	// ...忽略dsn配置，请参考上面例子...
// 	//配置MySQL连接参数
// 	// username := "root"  //账号
// 	// password := "123456" //密码
// 	// host := "127.0.0.1" //数据库地址，可以是Ip或者域名
// 	// port := 3306 //数据库端口
// 	// Dbname := "tizi365" //数据库名
// 	// timeout := "10s" //连接超时，10秒

// 	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
// 	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

// 	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
// 	// db, err := gorm.Open("mysql", dsn)
// 	// if err != nil {
// 	// 	panic("连接数据库失败, error=" + err.Error())
// 	// }
// 	var err error
// 	var dns string
// 	db_type := beego.AppConfig.String("db_type")
// 	db_name := beego.AppConfig.String("db_name")
// 	db_path := beego.AppConfig.String("db_path")
// 	if db_path == "" {
// 		db_path = "./"
// 	}

// 	dns = fmt.Sprintf("%s%s.db", db_path, db_name)
// 	_db, err = gorm.Open(db_type, dns)
// 	// _db.LogMode(true)
// 	if err != nil {
// 		panic("连接数据库失败, error=" + err.Error())
// 	}
// 	// defer gdb.Close()
// 	//禁止表名复数形式
// 	_db.SingularTable(true)

// 	//延时关闭数据库连接
// 	// defer db.Close()

// 	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
// 	//    _db, err := gorm.Open("mysql", dsn)
// 	//    if err != nil {
// 	// 	panic("连接数据库失败, error=" + err.Error())
// 	// }
// 	// 开发的时候需要打开调试日志
// 	_db.LogMode(true)
// 	//设置数据库连接池参数
// 	_db.DB().SetMaxOpenConns(100) //设置数据库连接池最大连接数
// 	_db.DB().SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
// }

// //获取gorm db对象，其他包需要执行数据库查询的时候，只要通过tools.getDB()获取db对象即可。
// //不用担心协程并发使用同样的db对象会共用同一个连接，
// // db对象在调用他的方法的时候会从数据库连接池中获取新的连接
// // 注意：使用连接池技术后，千万不要使用完db后调用db.Close关闭数据库连接，
// // 这样会导致整个数据库连接池关闭，导致连接池没有可用的连接
// func GetDB() *gorm.DB {
// 	return _db
// }

// // func GetDB() *gorm.DB {
// // 	var err error
// // 	var dns string
// // 	db_type := beego.AppConfig.String("db_type")
// // 	db_name := beego.AppConfig.String("db_name")
// // 	db_path := beego.AppConfig.String("db_path")
// // 	if db_path == "" {
// // 		db_path = "./"
// // 	}

// // 	dns = fmt.Sprintf("%s%s.db", db_path, db_name)
// // 	gdb, err = gorm.Open(db_type, dns)
// // 	gdb.LogMode(true)
// // 	if err != nil {
// // 		panic("failed to connect database")
// // 	}
// // 	defer gdb.Close()
// // }
