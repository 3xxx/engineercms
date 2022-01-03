package controllers

import (
	// "bytes"
	// "encoding/json"
	// "fmt"
	// "github.com/360EntSecGroup-Skylar/excelize"
	"github.com/astaxie/beego"
	"github.com/tealeg/xlsx"
	// "github.com/pborman/uuid"
	// "image/png"
	// "io"
	// "log"
	// "net/http"
	"os"
	"path"
	// "hydrocms/models"
	// "encoding/base64"
	"github.com/3xxx/engineercms/models"
	// "io/ioutil"
	// "regexp"
	"strconv"
	// "strings"
	"time"
)

// CMSWX froala API
type FileinputController struct {
	beego.Controller
}

type Fileinput struct {
	InitialPreview       []string        `json:"initialPreview"`
	InitialPreviewConfig []PreviewConfig `json:"initialPreviewConfig"`
	// InitialPreviewDownloadUrl []string        `json:"initialPreviewDownloadUrl"`
}

// type Preview struct {
// 	Url  string
// }

type PreviewConfig struct {
	Caption     string `json:"caption"`
	Size        int64  `json:"size"`
	Url         string `json:"url"`
	DownloadUrl string `json:"downloadUrl"`
	Key         string `json:"key"`
}

// @Title post bootstrapfileinput
// @Description post file by BootstrapFileInput
// @Success 200 {object} SUCCESS
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /bootstrapfileinput [post]
// 上传excel文件，格式：第一行，序号-名称-编号，都打上格子
func (c *FileinputController) BootstrapFileInput() {
	//获取上传的文件
	_, h, err := c.GetFile("input-ke-2[]")
	if err != nil {
		beego.Error(err)
	}
	fileSuffix := path.Ext(h.Filename)
	// random_name
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) + fileSuffix // + "_" + filename
	year, month, _ := time.Now().Date()
	err = os.MkdirAll("./attachment/legislation/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		beego.Error(err)
	}
	var filesize int64

	if h != nil {
		//保存附件
		filepath := "./attachment/legislation/" + strconv.Itoa(year) + month.String() + "/" + newname
		Url := "/attachment/legislation/" + strconv.Itoa(year) + month.String() + "/"
		err = c.SaveToFile("input-ke-2[]", filepath) //.Join("attachment", attachment)) //存文件    WaterMark(path)    //给文件加水印
		if err != nil {
			beego.Error(err)
		}
		filesize, _ = FileSize(filepath)
		filesize = filesize / 1000.0

		// 解析excel
		if fileSuffix == ".XLSX" || fileSuffix == ".xlsx" || fileSuffix == ".XLS" || fileSuffix == ".xls" {
			// xlsx, err := excelize.OpenFile(filepath)
			xlsx, err := xlsx.OpenFile(filepath)
			if err != nil {
				beego.Error(err)
				return
			}
			for _, sheet := range xlsx.Sheets {
				for i, row := range sheet.Rows {
					if i != 0 {
						// 这里要判断单元格列数，如果超过单元格使用范围的列数，则出错for j := 2; j < 7; j += 5 {
						j := 1
						standardtitle := row.Cells[j].String()
						if err != nil {
							beego.Error(err)
						}
						//2、根据名称搜索标准版本库，取得名称和版本号
						library, err := models.SearchLiabraryName(standardtitle)
						// beego.Info(library)
						if err != nil {
							beego.Error(err)
						}
						var LibraryNumber string
						if len(library) != 0 { //library != nil这样不行，空数组不是nil
							// beego.Info(library)
							//3、构造struct
							for j, w := range library {
								// beego.Info(w)
								if j == 0 {
									LibraryNumber = w.Category + " " + w.Number + "-" + w.Year //规范有效版本库中的完整编号
								} else {
									//数据库中出现了category和number前面都多了一个空格，执行日期出错：2011\-m\-10
									LibraryNumber = LibraryNumber + "," + w.Category + " " + w.Number + "-" + w.Year //规范有效版本库中的完整编号
								}
							}
							row.Cells[j+1].Value = LibraryNumber
						}

					}
				}
			}
			err = xlsx.Save(filepath)
			if err != nil {
				beego.Error(err)
			}
			// for index, name := range xlsx.GetSheetMap() {
			// 	fmt.Println(index, name)
			// 	// Get all the rows in the Sheet1.
			// 	rows, err := xlsx.GetRows(name)
			// 	for _, row := range rows {
			// 		for _, colCell := range row {
			// 			fmt.Print(colCell, "\t")
			// 		}
			// 		fmt.Println()
			// 	}
			// }
		}
		var fileinput Fileinput
		fileinput.InitialPreview = []string{Url + newname}
		// fileinput.InitialPreviewDownloadUrl = []string{Url + newname}
		// var config PreviewConfig
		config := make([]PreviewConfig, 1)
		config[0].Caption = newname
		config[0].DownloadUrl = Url + newname
		config[0].Size = filesize
		config[0].Key = Url + newname
		config[0].Url = Url + newname
		fileinput.InitialPreviewConfig = config
		c.Data["json"] = fileinput
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "link": "", "title": "", "original": ""}
		c.ServeJSON()
	}
}
