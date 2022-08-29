package main

import (
	"errors"
	"fmt"
	"github.com/3xxx/engineercms/commands/daemon"
	"github.com/3xxx/engineercms/conf"
	"github.com/3xxx/engineercms/controllers"
	"github.com/3xxx/engineercms/models"
	_ "github.com/3xxx/engineercms/routers" //这个最大的坑！！！！！
	"github.com/beego/beego/v2/adapter/toolbox"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/i18n"
	"github.com/kardianos/service"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

// graceful错误——删除router下的另外2个文件，重新运行生成即可。
// 项目引用了vendor里面的库，里面库go文件在init函数里面有同样的命名，从而在加载引用的时候会被调用，从而在main运行的时候提示已经解析多次。
// 用goland查找默认是该项目下的文件，可以通过指定vendor目录进行查找。

// 2021-8-30发现建表有问题：
// 1.beego是首先运行controllers里的init()，然后运行models里的init()，这是不科学的。gorm就先运行models里的init()，建立表格
// 2.因为controllers里先查询表还是操作表啥的出错，eforce那个，总是提示没有casbin_rule表，出错，程序走不下去，不去models里执行init()里的建表代码
// 3.将controllers里错误注释掉，它就执行建表了。但是user表总是到它为止，将user表后面的注释，就可以继续自动建表了，也没找到什么地方不对，再恢复，好像也可以。
// 4.在eforce里加个判断，如果查询有admin角色，则说明建好了角色表，可以执行操作了，否则return，不执行后面的建立角色-权限表。

func main() {
	// beego.AddFuncMap("dict", dict)
	//web.AutoRender = false
	//beego.TemplateLeft = "<<<"
	//beego.TemplateRight = ">>>"
	//web.TemplateLeft = "<<<"
	//web.TemplateRight = ">>>"
	//自动建表
	orm.RunSyncdb("default", false, true)
	// orm.RunSyncdb("default", true, true)
	models.InsertUser()
	// models.InsertGroup()
	models.InsertRole()

	// err := orm.RunSyncdb("default", false, true) // mindoc
	// if err == nil {
	// 	initialization()
	// } else {
	// 	panic(err.Error())
	// }

	web.AddFuncMap("loadtimes", loadtimes)
	web.AddFuncMap("subsuffix", subsuffix)
	//默认关闭orm调试模式——仅仅针对beego的orm有效。
	ormDebug, err := web.AppConfig.String("ormDebug")
	if err != nil {
		logs.Error("获取ormDebug ->", err.Error())
	}
	if ormDebug == "true" {
		orm.Debug = true
	} else {
		orm.Debug = false
	}
	//创建附件目录ModePerm FileMode = 0777 // 覆盖所有Unix权限位（用于通过&获取类型位）
	os.Mkdir("attachment", os.ModePerm)
	//创建轮播图片存放目录
	os.Mkdir("attachment/carousel", os.ModePerm)

	// time1 := "0/" + time + " * * * * *"

	// time1 := "* 30 8 * * 1-5"
	time1, err := web.AppConfig.String("tasktime")
	if err != nil {
		logs.Error("获取tasktime ->", err.Error())
	}
	if time1 != "" {
		tk1 := toolbox.NewTask("tk1", time1, func() error { controllers.SendMessage(); return nil }) //func() error { fmt.Println("tk1"); return nil }
		toolbox.AddTask("tk1", tk1)
		toolbox.StartTask()
		defer toolbox.StopTask()
	}
	// 定时备份数据库
	time2, err := web.AppConfig.String("backupdatatime")
	if err != nil {
		logs.Error("获取tasktime ->", err.Error())
	}
	// logs.Info(time2)
	if time2 != "" {
		tk2 := toolbox.NewTask("tk2", time2, func() error { controllers.Postdata(); return nil }) //func() error { fmt.Println("tk1"); return nil }
		toolbox.AddTask("tk2", tk2)
		toolbox.StartTask()
		defer toolbox.StopTask()
	}
	// ********mindoc*********
	// if len(os.Args) >= 3 && os.Args[1] == "service" {
	// 	if os.Args[2] == "install" {
	//    daemon.Install()
	// 	} else if os.Args[2] == "remove" {
	// 		daemon.Uninstall()
	// 	} else if os.Args[2] == "restart" {
	// 		daemon.Restart()
	// 	}
	// }

	initialization()
	// commands.RegisterCache()
	// commands.RegisterLogger(conf.LogFile)
	// commands.RegisterCommand()

	d := daemon.NewDaemon()
	s, err := service.New(d, d.Config())
	if err != nil {
		fmt.Println("Create service error => ", err)
		os.Exit(1)
	}
	if err := s.Run(); err != nil {
		log.Fatal("启动程序失败 ->", err)
	}

	// ********mindoc*********
	// web.SetStaticPath("/down1", "download1")
	web.Run()
	// 开启pprof，监听请求,无效
	// go func() {
	// 	log.Println(http.ListenAndServe("127.0.0.1:6060", nil))
	// }()
}

//显示页面加载时间
func loadtimes(t time.Time) int {
	return int(time.Now().Sub(t).Nanoseconds() / 1e6)
}

//去除扩展名
func subsuffix(in string) string {
	fileSuffix := path.Ext(in)
	return strings.TrimSuffix(in, fileSuffix)
}

// GoLang 如何在网页显示当前环境的版本号
// func main() {
// 	server()
// }
func server() {
	http.HandleFunc("/version", version)
	http.ListenAndServe(":8080", nil)
}

func version(w http.ResponseWriter, r *http.Request) {
	out, err := exec.Command("go", "version").Output()
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, fmt.Sprintf("%s", out))
}

//初始化数据，来自commands install.go
func initialization() {
	err := models.NewOption().Init()
	if err != nil {
		panic(err.Error())
	}
	// command.go里已经注册过了
	lang, _ := web.AppConfig.String("default_lang")
	// err = i18n.SetMessage(lang, "conf/lang/"+lang+".ini")
	// if err != nil {
	// 	panic(fmt.Errorf("initialize locale error: %s", err))
	// }

	member, err := models.NewMember().FindByFieldFirst("account", "admin")
	if errors.Is(err, orm.ErrNoRows) {

		// create admin user
		logs.Info("creating admin user")
		member.Account = "admin"
		member.Avatar = conf.URLForWithCdnImage("/static/mindoc/images/headimgurl.jpg")
		member.Password = "123456"
		member.AuthMethod = "local"
		member.Role = conf.MemberSuperRole
		member.Email = "admin@iminho.me"

		if err := member.Add(); err != nil {
			panic("Member.Add => " + err.Error())
		}

		// create demo book
		logs.Info("creating demo book")
		book := models.NewBook()

		book.MemberId = member.MemberId
		book.BookName = i18n.Tr(lang, "init.default_proj_name") //"MinDoc演示项目"
		book.Status = 0
		book.ItemId = 1
		book.Description = i18n.Tr(lang, "init.default_proj_desc") //"这是一个MinDoc演示项目，该项目是由系统初始化时自动创建。"
		book.CommentCount = 0
		book.PrivatelyOwned = 0
		book.CommentStatus = "open" //"closed"
		book.Identify = "mindoc"
		book.DocCount = 0
		book.CommentCount = 0
		book.Version = time.Now().Unix()
		book.Cover = conf.GetDefaultCover()
		book.Editor = "markdown"
		book.Theme = "default"

		if err := book.Insert(lang); err != nil {
			panic("初始化项目失败 -> " + err.Error())
		}
	} else if err != nil {
		panic(fmt.Errorf("occur errors when initialize: %s", err))
	}

	if !models.NewItemsets().Exist(1) {
		item := models.NewItemsets()
		item.ItemName = i18n.Tr(lang, "init.default_proj_space") //"默认项目空间"
		item.MemberId = 1
		if err := item.Save(); err != nil {
			panic("初始化项目空间失败 -> " + err.Error())
		}
	}
}

// *******Golang 各种类型的默认值*********
// bool //默认值为false
// string //默认值为空字符串
// int int8 int16 int32 int64 //默认值为0
// uint uint8 uint16 uint32 uint64 uintptr //默认值为0
// byte // uint8 的别名
// rune // int32 的别名
// float32 float64 //默认值为0
// complex64 complex128 //默认值为0
// 在Go语言中，布尔类型的"0"（初始值）为false，数值类型的"0"为0，
// 字符串类型的"0"为空字符串""，而指针/切片/映射/通道/函数和接口的"0"即为nil
// 当你声明一个结构体变量并未初始化时，该结构体所有的域(Field)都为"0""nil"（初始值）

// func Create(name string) (file *File, err error)
// func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
//这里的flag int是用下面的这些方式，O_CREAT
// const (
//     O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
//     O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
//     O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
//     O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
//     O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
//     O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
//     O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
//     O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
// )
//打开 以只读,文件不存在将创建 方式打开  要存放的路径资源
// f, f_err := os.OpenFile(file_save, os.O_WRONLY|os.O_CREATE, 0666)
// 		if f_err != nil {
// 					fmt.Fprintf(w, "file open fail:%s", f_err)
// 		}		//文件 copy
// 		_, copy_err := io.Copy(f, file)
// 		if copy_err != nil {
// 				fmt.Fprintf(w, "file copy fail:%s", copy_err)
// 		}		//关闭对应打开的文件
// defer f.Close()
// defer file.Close()

//获得文件名最快的代码实现方式比较
// func main() {
//     filename := "/root/Desktop/模特唐嘉灵2.jpg"
//     //方案一
//     startTime1 := time.Now()
//     for i := 0; i < 1000000; i++ {
//         _ = filepath.Base(filename)
//     }
//     endTime1 := time.Now()

// 当每次读取块的大小小于4KB，建议使用bufio.NewReader(f), 大于4KB用bufio.NewReaderSize(f,缓存大小)
// 要读Reader, 图方便用ioutil.ReadAll()
// 一次性读取文件，使用ioutil.ReadFile()
// func Ioutil(name string) {
//     if contents,err := ioutil.ReadFile(name);err == nil {
//         //因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
//         result := strings.Replace(string(contents),"\n","",1)
//         fmt.Println("Use ioutil.ReadFile to read a file:",result)
//         }
//     }
// func main() {
//     b, err := ioutil.ReadFile("test.log")
//     if err != nil {
//         fmt.Print(err)
//     }
//     fmt.Println(b)
//     str := string(b)
//     fmt.Println(str)
// }
// 读文件方式一：利用ioutil.ReadFile直接从文件读取到[]byte中
// func Read0()  (string){
//     f, err := ioutil.ReadFile("file/test")
//     if err != nil {
//         fmt.Println("read fail", err)
//     }
//     return string(f)
// }

// func read3(path string)string{
//     fi,err := os.Open(path)
//     if err != nil{panic(err)}
//     defer fi.Close()
//     fd,err := ioutil.ReadAll(fi)
//     // fmt.Println(string(fd))
//     return string(fd)
// }

//将客户端说的一句话记录在【以他的名字命名的文件里】
// func writeMsgToLog(msg string, client Client) {
// 	//打开文件
// 	file, e := os.OpenFile(
// 		"D:/BJBlockChain1801/demos/W4/day1/01ChatRoomII/logs/"+client.name+".log",
// 		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
// 		0644)
// 	SHandleError(e, "os.OpenFile")
// 	defer file.Close()

// 	//追加这句话
// 	logMsg := fmt.Sprintln(time.Now().Format("2006-01-02 15:04:05"), msg)
// 	file.Write([]byte(logMsg))
// }

//map for  Http Content-Type  Http 文件类型对应的content-Type
// var HttpContentType = map[string]string{
//   ".avi": "video/avi",
//   ".mp3": "   audio/mp3",
//   ".mp4": "video/mp4",
//   ".wmv": "   video/x-ms-wmv",
//   ".asf":  "video/x-ms-asf",
//   ".rm":   "application/vnd.rn-realmedia",
//   ".rmvb": "application/vnd.rn-realmedia-vbr",
//   ".mov":  "video/quicktime",
//   ".m4v":  "video/mp4",
//   ".flv":  "video/x-flv",
//   ".jpg":  "image/jpeg",
//   ".png":  "image/png",
// }
// //根据文件路径读取返回流文件 参数url
// func PubResFileStreamGetService(c *gin.Context) {
// filePath := c.Query("url")
// //获取文件名称带后缀
// fileNameWithSuffix := path.Base(filePath)
// //获取文件的后缀
// fileType := path.Ext(fileNameWithSuffix)
// //获取文件类型对应的http ContentType 类型
// fileContentType := HttpContentType[fileType]
// if common.IsEmpty(fileContentType) {
//   c.String(http.StatusNotFound, "file http contentType not found")
//   return
// }
// c.Header("Content-Type", fileContentType)
// c.File(filePath)
// }

// beego文件流
// step1:在beego项目中添加:
// copyrequestbody = true
// 然后在控制器中添加代码：
// req:=this.Ctx.Input.RequestBody
// data:=string(req)
