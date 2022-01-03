package controllers

import (
	"bufio"
	"github.com/3xxx/engineercms/controllers/utils"
	"github.com/beego/beego/v2/server/web"
	"io"
	"os"
	"strings"
)

// CMSADMIN API
type AdminLogController struct {
	web.Controller
}

type Logstruct struct {
	Id       int64  `json:"id"`
	Datetime string `json:"datetime"`
	Tag      string `json:"tag"`
	File     string `json:"file"`
	User     string `json:"user"`
	Action   string `json:"action"`
	Url      string `json:"url"`
}

// @Title getAdminBlock
// @Description get log list
// @Success 200 {object} success
// @Failure 400 Invalid page
// @Failure 404 page not found
// @router /infolog [get]
func (c *AdminLogController) InfoLog() {
	username, _, _, _, _ := checkprodRole(c.Ctx)
	var logslice = make([]Logstruct, 0)
	var logs = make([]Logstruct, 1)
	fileName := "log/engineercms.info.log"
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		// fmt.Println("Open file error!", err)
		utils.FileLogs.Error(username + " openfile " + err.Error())
		return
	}
	defer file.Close()
	// stat, err := file.Stat()
	// if err != nil {
	// 	panic(err)
	// 	utils.FileLogs.Error(username + " panic " + err.Error())
	// }
	// var size = stat.Size()
	// fmt.Println("file size=", size)
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		// fmt.Println(line)
		array := strings.Split(string(line), " ")
		if len(array) > 6 {
			logs[0].Datetime = array[0] + " " + array[1]
			logs[0].Tag = array[2]
			logs[0].File = array[3]
			logs[0].User = array[4]
			logs[0].Action = array[5]
			logs[0].Url = array[6]
			logslice = append(logslice, logs...)
		}
		// for _, v := range array {
		// 	fmt.Println(v)
		// }
		//array 到 json str
		// arr := []string{"hello", "apple", "python", "golang", "base", "peach", "pear"}
		// lang, err := json.Marshal(array)
		// if err == nil {
		// 	fmt.Println(lang)
		// }
		if err != nil {
			if err == io.EOF {
				// fmt.Println("File read ok!")
				utils.FileLogs.Error(username + " File read ok! " + err.Error())
				break
			} else {
				// fmt.Println("Read file error!", err)
				utils.FileLogs.Error(username + " Read file error! " + err.Error())
				return
			}
		}
	}
	c.Data["json"] = logslice
	c.ServeJSON()
}

// @Title getAdminBlock
// @Description get log list
// @Success 200 {object} success
// @Failure 400 Invalid page
// @Failure 404 page not found
// @router /errlog [get]
func (c *AdminLogController) ErrLog() {
	username, _, _, _, _ := checkprodRole(c.Ctx)
	var logslice = make([]Logstruct, 0)
	var logs = make([]Logstruct, 1)
	fileName := "log/engineercms.error.log"
	file, err := os.OpenFile(fileName, os.O_RDWR, 0666)
	if err != nil {
		// fmt.Println("Open file error!", err)
		utils.FileLogs.Error(username + " openfile " + err.Error())
		return
	}
	defer file.Close()
	// stat, err := file.Stat()
	// if err != nil {
	// 	panic(err)
	// 	utils.FileLogs.Error(username + " panic " + err.Error())
	// }
	// var size = stat.Size()
	// fmt.Println("file size=", size)
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		// fmt.Println(line)
		array := strings.Split(string(line), " ")
		if len(array) > 6 {
			logs[0].Datetime = array[0] + " " + array[1]
			logs[0].Tag = array[2]
			logs[0].File = array[3]
			logs[0].User = array[4]
			logs[0].Action = array[5]
			logs[0].Url = array[6]
			logslice = append(logslice, logs...)
		}
		// for _, v := range array {
		// 	fmt.Println(v)
		// }
		//array 到 json str
		// arr := []string{"hello", "apple", "python", "golang", "base", "peach", "pear"}
		// lang, err := json.Marshal(array)
		// if err == nil {
		// 	fmt.Println(lang)
		// }
		if err != nil {
			if err == io.EOF {
				// fmt.Println("File read ok!")
				utils.FileLogs.Error(username + " File read ok! " + err.Error())
				break
			} else {
				// fmt.Println("Read file error!", err)
				utils.FileLogs.Error(username + " Read file error! " + err.Error())
				return
			}
		}
	}
	c.Data["json"] = logslice
	c.ServeJSON()
}

// golang 逐行读取文件
// 复制代码
// package main

// import (
//     "bufio"
//     "fmt"
//     "io"
//     "os"
// )

// func main() {
//     fi, err := os.Open("C:/Documents and Settings/xxx/Desktop/tax.txt")
//     if err != nil {
//         fmt.Printf("Error: %s\n", err)
//         return
//     }
//     defer fi.Close()
//     br := bufio.NewReader(fi)
//     for {
//         a, _, c := br.ReadLine()
//         if c == io.EOF {
//             break
//         }
//         fmt.Println(string(a))
//     }
// }
