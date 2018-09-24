package main

import (
	_ "github.com/3xxx/engineercms/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/plugins/cors"
	"github.com/3xxx/engineercms/models"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"time"
)

func main() {
	// beego.AddFuncMap("dict", dict)
	beego.AddFuncMap("loadtimes", loadtimes)

	//开启orm调试模式
	orm.Debug = true
	//创建附件目录
	os.Mkdir("attachment", os.ModePerm)
	//创建轮播图片存放目录
	os.Mkdir("attachment//carousel", os.ModePerm)
	//自动建表
	orm.RunSyncdb("default", false, true)
	models.InsertUser()
	// insertGroup()
	// insertRole()
	beego.Run()
}

//显示页面加载时间
func loadtimes(t time.Time) int {
	return int(time.Now().Sub(t).Nanoseconds() / 1e6)
}

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
