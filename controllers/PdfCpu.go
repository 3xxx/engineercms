package controllers

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	// "encoding/json"
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	// beego "github.com/beego/beego/v2/adapter"
	// "github.com/beego/beego/v2/adapter/httplib"
	// "github.com/beego/beego/v2/adapter/logs"
	// "github.com/bitly/go-simplejson"
	"io/ioutil"
	// "net"
	// "net/http"
	// "net/url"
	"os"
	// "path"
	"strconv"
	// "strings"
	"encoding/base64"
	"time"
)

// CMSADMIN API
type PdfCpuController struct {
	web.Controller
}

// func AddWatermarksFile(inFile, outFile string, selectedPages []string, wm *pdf.Watermark, conf *pdf.Configuration) (err error)

// @Title post signature to pdf
// @Description post signature to pdf
// @Param id path string true "The id of pdf"
// @Param pageNumber query string false "The pageNumber of pdf"
// @Param numPages query string true "The numPages of pdf"
// @Param offsetdx query string false "The offsetdx of signature"
// @Param offsetdy query string false "The offsetdy of signature"
// @Param scale query string false "The scale of signature"
// @Param image query string true "The base64 image of signature"
// @Success 200 {object} models.AddSignature
// @Failure 400 Invalid page supplied
// @Failure 404 pdf not found
// @router /addwatermarks/:id [post]
func (c *PdfCpuController) AddWatermarks() {
	// openID := c.GetSession("openID")
	// beego.Info(openID)
	// if openID != nil {
	// user, err := models.GetUserByOpenID(openID.(string))
	// if err != nil {
	// 	logs.Error(err)
	// } else {
	// 	useridstring := strconv.FormatInt(user.Id, 10)
	// 判断用户是否具有权限。
	id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
		return
	}

	//根据附件id取得附件的prodid，路径
	onlyattachment, err := models.GetOnlyAttachbyId(idNum)
	if err != nil {
		logs.Error(err)
	}

	//docid——uid——me
	// doc, err := models.Getdocbyid(onlyattachment.DocId)
	// if err != nil {
	// 	logs.Error(err)
	// }

	// fileext := path.Ext(doc.FileName)
	filepathname := "./attachment/onlyoffice/" + onlyattachment.FileName
	// filepathname = "./attachment/onlyoffice/bootstrap.pdf"
	//根据附件id取得附件的prodid，路径
	// attachment, err := models.GetAttachbyId(idNum)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// product, err := models.GetProd(attachment.ProductId)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// //根据projid取出路径
	// proj, err := models.GetProj(product.ProjectId)
	// if err != nil {
	// 	logs.Error(err)
	// 	utils.FileLogs.Error(err.Error())
	// }
	// var projurl string
	// if proj.ParentIdPath == "" || proj.ParentIdPath == "$#" {
	// 	projurl = "/" + strconv.FormatInt(proj.Id, 10) + "/"
	// } else {
	// 	projurl = "/" + strings.Replace(strings.Replace(proj.ParentIdPath, "#", "/", -1), "$", "", -1) + strconv.FormatInt(proj.Id, 10) + "/"
	// }
	//由proj id取得url
	// fileurl, _, err := GetUrlPath(product.ProjectId)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// fileext := path.Ext(attachment.FileName)
	// if e.Enforce(useridstring, projurl, c.Ctx.Request.Method, fileext) {
	// 	c.Ctx.Output.Download(fileurl + "/" + attachment.FileName)
	// }
	// 	}
	// } else {
	// 	c.Data["json"] = "未查到openID"
	// 	c.ServeJSON()
	// }
	pageNumber := c.GetString("pageNumber")
	selectedPages := make([]string, 1)
	selectedPages[0] = pageNumber
	imagebase64 := c.GetString("image")
	image, err := base64.StdEncoding.DecodeString(imagebase64) //成图片文件并把文件写入到buffer
	if err != nil {
		logs.Error(err)
	}
	// fileSuffix := path.Ext(h.Filename)
	// random_name
	newname := strconv.FormatInt(time.Now().UnixNano(), 10) //+ fileSuffix // + "_" + filename
	// err = ioutil.WriteFile(path1+newname+".jpg", ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
	// if err != nil {
	// 	logs.Error(err)
	// }
	year, month, _ := time.Now().Date()
	err = os.MkdirAll("./attachment/onlyoffice/signature/"+strconv.Itoa(year)+month.String()+"/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		logs.Error(err)
	}
	imagefilenamepath := "./attachment/onlyoffice/signature/" + strconv.Itoa(year) + month.String() + "/" + newname + ".png"

	err = ioutil.WriteFile(imagefilenamepath, image, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
	if err != nil {
		logs.Error(err)
	}

	conf := pdfcpu.NewDefaultConfiguration()
	conf.ValidationMode = pdfcpu.ValidationNone
	// Add a "Demo" watermark to all pages of in.pdf along the diagonal running from lower left to upper right.
	// onTop := false
	// wm, err := pdfcpu.ParseTextWatermarkDetails("Demo", "", onTop)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// err = api.AddWatermarksFile(filepathname, "", nil, wm, nil)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// Stamp all odd pages of in.pdf in red "Confidential" in 48 point Courier using a rotation angle of 45 degrees.
	// onTop = true
	// wm, err = pdfcpu.ParseTextWatermarkDetails("Confidential", "font:Courier, c: 1 0 0, rot:45, s:1 abs, points:48", onTop)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// err = api.AddWatermarksFile(filepathname, "", []string{"odd"}, wm, nil)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// Add image stamps to in.pdf using absolute scaling and a negative rotation of 90 degrees.
	// func ParseImageWatermarkDetails(fileName, desc string, onTop bool) (*Watermark, error)"pos: br, off: -30 40, s:1.0 a, rot:0"
	var desc string
	scale := c.GetString("scale")
	if scale == "" {
		scale = "1.0"
	}
	offsetdx := c.GetString("offsetdx")
	if offsetdx == "" {
		offsetdx = "-30"
	}
	offsetdy := c.GetString("offsetdy")
	if offsetdy == "" {
		offsetdy = "40"
	}
	desc = "pos: br, off: " + offsetdx + " " + offsetdy + ", s:" + scale + " a, rot:0"
	onTop := true
	wm, err := pdfcpu.ParseImageWatermarkDetails(imagefilenamepath, desc, onTop, 1)
	if err != nil {
		logs.Error(err)
	}
	err = api.AddWatermarksFile(filepathname, "", selectedPages, wm, conf)
	if err != nil {
		logs.Error(err)
	}

	// Add a PDF stamp to all pages of in.pdf using the 2nd page of stamp.pdf, use absolute scaling of 0.5
	// and rotate along the 2nd diagonal running from upper left to lower right corner.
	// onTop = true
	// wm, _ = pdfcpu.ParsePDFWatermarkDetails("stamp.pdf:2", "s:.5 a, d:2", onTop)
	// err = AddWatermarksFile(filepathname, "", nil, wm, nil)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERR"}
		c.ServeJSON()
	} else {
		// c.Data["json"] = id
		c.Data["json"] = map[string]interface{}{"info": "SUCCESS"}
		c.ServeJSON()
	}
}

// @Title get only pdf
// @Description get only pdf
// @Param id path string true "The id of pdf"
// @Success 200 {object} models.GetOnlyPdf
// @Failure 400 Invalid page supplied
// @Failure 404 pdf not found
// @router /onlypdf/:id [get]
func (c *PdfCpuController) OnlyPdf() {

	id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	// id := c.GetString("id")
	// beego.Info(id)
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}

	//根据附件id取得附件的prodid，路径
	onlyattachment, err := models.GetOnlyAttachbyId(idNum)
	if err != nil {
		logs.Error(err)
	}
	// fileext := path.Ext(doc.FileName)
	filepathname := "/attachment/onlyoffice/" + onlyattachment.FileName
	// beego.Info(filepathname)
	// PdfLink := Url + "/" + Attachments[pNum-1].FileName
	// beego.Info(PdfLink)
	c.Data["DocId"] = id
	c.Data["PdfLink"] = filepathname
	c.TplName = "web/onlyviewer.html"

	// logs := logs.NewLogger(1000)
	// logs.SetLogger("file", `{"filename":"log/test.log"}`)
	// logs.EnableFuncCallDepth(true)
	// logs.Info(c.Ctx.Input.IP() + " " + "ListAllTopic")
	// logs.Close()
	// count, _ := models.M("logoperation").Alias(`op`).Field(`count(op.id) as count`).Where(where).Count()
	// if count > 0 {
	// 	pagesize := 10
	// 	p := tools.NewPaginator(this.Ctx.Request, pagesize, count)
	// 	log, _ := models.M("logoperation").Alias(`op`).Where(where).Limit(strconv.Itoa(p.Offset()), strconv.Itoa(pagesize)).Order(`op.id desc`).Select()
	// 	this.Data["data"] = log
	// 	this.Data["paginator"] = p
	// }
}

type Signature struct {
	BmpData string `json:"bmpdata"`
	Ret     string `json:"ret"`
}

// type BmpData1 struct {
// 	Action      string `json:"action"`
// 	Bmpbitcount string `json:"bmpbitcount"`
// 	Bbmpdata    string `json:"bmpdata"`
// 	Bmpdatalen  string `json:"bmpdatalen"`
// 	Bmpheight   string `json:"bmpheight"`
// 	Bmpwidth    string `json:"bmpwidth"`
// 	Ret         string `json:"ret"`
// 	Variable1   string `json:"variable1"`
// 	Variable2   string `json:"variable2"`
// 	Variable3   string `json:"variable3"`
// }

// type Ret1 struct {
// 	StartWith string `json:"startWith"`
// }

// @Title get test
// @Description get test
// @Param id path string true "The id of test"
// @Success 200 {object} models.Gettest
// @Failure 400 Invalid page supplied
// @Failure 404 pdf not found
// @router /test/:id [get]
func (c *PdfCpuController) Test() {
	id := c.Ctx.Input.Param(":id")
	logs.Info(id)
	// ret1 := Ret1{
	// 	StartWith: "00",
	// }
	signature := &Signature{
		BmpData: "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",
		Ret:     "0078",
	}

	c.Data["json"] = signature
	c.ServeJSON()
}
