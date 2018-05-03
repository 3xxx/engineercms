package controllers

import (
	"encoding/json"
	"github.com/3xxx/engineercms/models"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	// "github.com/casbin/beego-orm-adapter"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
	// "mime/multipart"
	// "github.com/astaxie/beego/httplib"
	// "bytes"
	// "github.com/astaxie/beego/utils/pagination"
	// "io"
)

type OnlyController struct {
	beego.Controller
}

type Callback struct {
	Key         string    `json:"key"`
	Status      int       `json:"status"`
	Url         string    `json:"url"`
	Changesurl  string    `json:"changesurl"`
	History     history1  `json:"history"`
	Users       []string  `json:"users"`
	Actions     []action  `json:"actions"`
	Lastsave    time.Time `json:"lastsave"`
	Notmodified bool      `json:"notmodified"`
}

type action struct {
	Type   int   `json:"type"`
	Userid int64 `json:"userid"`
}

type history1 struct {
	ServerVersion string    `json:"serverVersion"`
	Changes       []change  `json:"changes"`
	Created       time.Time `json:"created"`
	Key           string    `json:"key"`
	// ChangesUrl    string    `json:"changesurl"`
	User    User1 `json:"user"`
	Version int   `json:"version"`
}

type change struct {
	Created string `json:"created"` //time.Time
	User    User1  `json:"user"`
}

type User1 struct {
	Id   string `json:"id"` //必须大写才能在tpl中显示{{.json}}
	Name string `json:"name"`
}

//构造changesurl结构
type changesurl struct {
	Version    int    `json:"version"`
	ChangesUrl string `json:"changesurl"`
}

// type Callback2 struct {
// 	Key         string    `json:"key"`
// 	Status      int       `json:"status"`
// 	Url         string    `json:"url"`
// 	Changesurl  string    `json:"changesurl"`
// 	History     history2  `json:"history"`
// 	Users       []string  `json:"users"`
// 	Actions     []action  `json:"actions"`
// 	Lastsave    time.Time `json:"lastsave"`
// 	Notmodified bool      `json:"notmodified"`
// }

// type history2 struct {
// 	ServerVersion string    `json:"serverVersion"`
// 	Changes       []change2 `json:"changes"`
// 	Created       time.Time `json:"created"`
// 	Key           string    `json:"key"`
// 	User          User1     `json:"user"`
// 	Version       int       `json:"version"`
// }

// type change2 struct {
// 	Created string `json:"created"` //time.Time
// 	User    User2  `json:"user"`
// }

// type User2 struct {
// 	Id   string `json:"id"` //必须大写才能在tpl中显示{{.json}}
// 	Name string `json:"name"`
// }

// type FileNode struct {
// 	Id        int64       `json:"id"`
// 	Title     string      `json:"text"`
// 	Code      string      `json:"code"` //分级目录代码
// 	FileNodes []*FileNode `json:"nodes"`
// }

type Onlyoffice1 struct {
	Id      int64
	Code    string
	Title   string
	Ext     string
	Created time.Time
	Updated time.Time
}

type OnlyLink struct {
	Id        int64
	Code      string
	Title     string
	Label     string
	End       time.Time
	Principal string
	Uid       int64
	Created   time.Time
	Updated   time.Time
	Docxlink  []DocxLink
	// Xlsxlink  []XlsxLink
	// Pptxlink  []PptxLink
}

type DocxLink struct {
	Id         int64
	Title      string
	Suffix     string
	Permission string
	// Link    string
	Created time.Time
	Updated time.Time
}

//权限表提交的table中json数据解析成struct
type Rolepermission struct {
	Id         int64
	Name       string `json:"name"`
	Rolenumber string
	Permission string `json:"role"`
}

// type XlsxLink struct {
// 	Id    int64
// 	Title string
// 	Created time.Time
// 	Updated time.Time
// }

// type PptxLink struct {
// 	Id    int64
// 	Title string
// 	Created time.Time
// 	Updated time.Time
// }

//文档管理页面
func (c *OnlyController) Get() {
	//取得客户端用户名
	v := c.GetSession("uname")
	if v != nil {
		uname := v.(string)
		user, err := models.GetUserByUsername(uname)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Uid"] = user.Id
	} else {
		c.Data["Uid"] = 0
	}
	username, role := checkprodRole(c.Ctx)
	roleint, err := strconv.Atoi(role)
	if err != nil {
		beego.Error(err)
	}
	if role == "1" {
		c.Data["IsAdmin"] = true
	} else if roleint > 1 && roleint < 5 {
		c.Data["IsLogin"] = true
	} else {
		c.Data["IsAdmin"] = false
		c.Data["IsLogin"] = false
	}
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	// beego.Info(c.Ctx.Input.UserAgent())
	u := c.Ctx.Input.UserAgent()
	// re := regexp.MustCompile("Trident")
	// loc := re.FindStringIndex(u)
	// loc[0] > 1
	c.Data["IsOnlyOffice"] = true
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		beego.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		c.TplName = "onlyoffice/docs.tpl"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "onlyoffice/docs.tpl"
	}
	// c.Data["Url"] = c.Ctx.Request.URL.String()
	// var u = navigator.userAgent, app = navigator.appVersion;
	//       return {
	//           trident: u.indexOf('Trident') > -1, //IE内核
	//           presto: u.indexOf('Presto') > -1, //opera内核
	//           webKit: u.indexOf('AppleWebKit') > -1, //苹果、谷歌内核
	//           gecko: u.indexOf('Gecko') > -1 && u.indexOf('KHTML') == -1,//火狐内核
	//           mobile: !!u.match(/AppleWebKit.*Mobile.*/), //是否为移动终端
	//           ios: !!u.match(/\(i[^;]+;( U;)? CPU.+Mac OS X/), //ios终端//两个感叹号的作用就在于，如果明确设置了变量的值（非null/undifined/0/”“等值),结果就会根据变量的实际值来返回，如果没有设置，结果就会返回false。
	//           android: u.indexOf('Android') > -1 || u.indexOf('Adr') > -1, //android终端
	//           iPhone: u.indexOf('iPhone') > -1 , //是否为iPhone或者QQHD浏览器
	//           iPad: u.indexOf('iPad') > -1, //是否iPad
	//           webApp: u.indexOf('Safari') == -1, //是否web应该程序，没有头部与底部
	//           //webApp: ("standalone" in window.navigator)&&window.navigator.standalone, //是否web应该程序，没有头部与底部
	//           weixin: u.indexOf('MicroMessenger') > -1, //是否微信 （2015-01-22新增）
	//           qq: u.match(/\sQQ/i) == " qq" //是否QQ

	// public static boolean  isMobileDevice(String requestHeader){
	/**
	 * android : 所有android设备
	 * mac os : iphone ipad
	 * windows phone:Nokia等windows系统的手机
	 */
	// String[] deviceArray = new String[]{"android","mac os","windows phone"};
	// if(requestHeader == null)
	//     return false;
	// requestHeader = requestHeader.toLowerCase();
	// for(int i=0;i<deviceArray.length;i++){
	//     if(requestHeader.indexOf(deviceArray[i])>0){
	//         return true;
	//     }
	// }
	// return false;

	// }

	//         Enumeration   typestr = request.getHeaderNames();
	// String s1 = request.getHeader("user-agent");
	// if(s1.contains("Android")) {
	// System.out.println("Android移动客户端");
	// } else if(s1.contains("iPhone")) {
	// System.out.println("iPhone移动客户端");
	// }  else if(s1.contains("iPad")) {
	// System.out.println("iPad客户端");
	// }  else {
	// System.out.println("其他客户端");
	// }
	// 	if (/(iPhone|iPad|iPod|iOS)/i.test(navigator.userAgent)) {
	//   	//alert(navigator.userAgent);
	//   	window.location.href ="iPhone.html";
	// } else if (/(Android)/i.test(navigator.userAgent)) {
	//     //alert(navigator.userAgent);
	//     window.location.href ="Android.html";
	// } else {
	//     window.location.href ="pc.html";
	// };
}

//提供给列表页的table中json数据
func (c *OnlyController) GetData() {
	//1.取得客户端用户名
	var uname, useridstring string
	var user models.User
	var err error
	v := c.GetSession("uname")
	// var role, userrole int
	if v != nil {
		uname = v.(string)
		// c.Data["Uname"] = v.(string)
		user, err = models.GetUserByUsername(uname)
		if err != nil {
			beego.Error(err)
		}
		c.Data["Uid"] = user.Id
		// userrole = user.Role
		useridstring = strconv.FormatInt(user.Id, 10)
	}
	var myRes, roleRes [][]string
	if useridstring != "" {
		myRes = e.GetPermissionsForUser(useridstring)
		// beego.Info(myRes)
	}
	myResall := e.GetPermissionsForUser("") //取出所有设置了权限的数据

	docs, err := models.GetDocs()
	if err != nil {
		beego.Error(err)
	}

	link := make([]OnlyLink, 0)
	Docxslice := make([]DocxLink, 0)
	for _, w := range docs {

		Attachments, err := models.GetOnlyAttachments(w.Id)
		if err != nil {
			beego.Error(err)
		}
		linkarr := make([]OnlyLink, 1)
		linkarr[0].Id = w.Id
		linkarr[0].Code = w.Code
		linkarr[0].Title = w.Title
		linkarr[0].Label = w.Label
		linkarr[0].End = w.End
		linkarr[0].Principal = w.Principal
		linkarr[0].Uid = w.Uid
		linkarr[0].Created = w.Created
		linkarr[0].Updated = w.Updated
		//docid——me——1
		for _, v := range Attachments {
			// beego.Info(v.FileName)
			// fileext := path.Ext(v.FileName)
			docxarr := make([]DocxLink, 1)
			docxarr[0].Permission = "1"
			//查询v.Id是否和myres的V1路径后面的id一致，如果一致，则取得V2（权限）
			//查询用户具有的权限
			// beego.Info(useridstring)
			if useridstring != "" { //如果是登录用户，则设置了权限的文档不能看
				// beego.Info(myRes)
				// myRes1 := e.GetPermissionsForUser("") //取出所有设置了权限的数据
				if w.Uid != user.Id { //如果不是作者本人
					for _, k := range myResall { //所有设置了权限的都不能看
						// beego.Info(k)
						if strconv.FormatInt(v.Id, 10) == path.Base(k[1]) {
							docxarr[0].Permission = "4"
						}
					}

					for _, k := range myRes { //如果与登录用户对应上，则赋予权限
						// beego.Info(k)
						if strconv.FormatInt(v.Id, 10) == path.Base(k[1]) {
							docxarr[0].Permission = k[2]
						}
					}
					roles := e.GetRolesForUser(useridstring) //取出用户的所有角色
					for _, w1 := range roles {               //2018.4.30修改这个bug，这里原先w改为w1
						roleRes = e.GetPermissionsForUser(w1) //取出角色的所有权限，改为w1
						for _, k := range roleRes {
							// beego.Info(k)
							if strconv.FormatInt(v.Id, 10) == path.Base(k[1]) {
								// docxarr[0].Permission = k[2]
								int1, err := strconv.Atoi(k[2])
								if err != nil {
									beego.Error(err)
								}
								int2, err := strconv.Atoi(docxarr[0].Permission)
								if err != nil {
									beego.Error(err)
								}
								if int1 < int2 {
									docxarr[0].Permission = k[2] //按最小值权限
								}
							}
						}
					}

				} //如果是用户自己的文档，则permission为1，默认
			} else { //如果用户没登录，则设置了权限的文档不能看
				for _, k := range myResall { //所有设置了权限的不能看
					// beego.Info(k)
					if strconv.FormatInt(v.Id, 10) == path.Base(k[1]) {
						// beego.Info(i)
						// beego.Info(strconv.FormatInt(v.Id, 10))
						// beego.Info(path.Base(k[1]))
						docxarr[0].Permission = "4"
						// beego.Info(strconv.FormatInt(v.Id, 10))
						// beego.Info(path.Base(k[1]))
					}
				}
			}

			docxarr[0].Id = v.Id
			docxarr[0].Title = v.FileName
			if path.Ext(v.FileName) == ".docx" || path.Ext(v.FileName) == ".DOCX" || path.Ext(v.FileName) == ".doc" || path.Ext(v.FileName) == ".DOC" {
				docxarr[0].Suffix = "docx"

			} else if path.Ext(v.FileName) == ".XLSX" || path.Ext(v.FileName) == ".xlsx" || path.Ext(v.FileName) == ".XLS" || path.Ext(v.FileName) == ".xls" {
				docxarr[0].Suffix = "xlsx"
				// xlsxarr := make([]XlsxLink, 1)
				// xlsxarr[0].Id = v.Id
				// xlsxarr[0].Title = v.FileName
				// Xlsxslice = append(Xlsxslice, xlsxarr...)
			} else if path.Ext(v.FileName) == ".pptx" || path.Ext(v.FileName) == ".PPTX" || path.Ext(v.FileName) == ".ppt" || path.Ext(v.FileName) == ".PPT" {
				docxarr[0].Suffix = "pptx"
				// pptxarr := make([]PptxLink, 1)
				// pptxarr[0].Id = v.Id
				// pptxarr[0].Title = v.FileName
				// Pptxslice = append(Pptxslice, pptxarr...)
			} else if path.Ext(v.FileName) == ".pdf" || path.Ext(v.FileName) == ".PDF" {
				docxarr[0].Suffix = "pdf"
			} else if path.Ext(v.FileName) == ".txt" || path.Ext(v.FileName) == ".TXT" {
				docxarr[0].Suffix = "txt"
			}
			Docxslice = append(Docxslice, docxarr...)
		}
		linkarr[0].Docxlink = Docxslice
		// linkarr[0].Xlsxlink = Xlsxslice
		// linkarr[0].Pptxlink = Pptxslice
		Docxslice = make([]DocxLink, 0) //再把slice置0
		// Xlsxslice = make([]XlsxLink, 0) //再把slice置0
		// Pptxslice = make([]PptxLink, 0)

		link = append(link, linkarr...)
	}
	c.Data["json"] = link //products
	c.ServeJSON()
}

//取得changesurl
// func (c *OnlyController) ChangesUrl() {
// 	version := c.Input().Get("version")

// 	versionint, err := strconv.Atoi(version)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	attachmentid := c.Input().Get("attachmentid")
// 	idNum, err := strconv.ParseInt(attachmentid, 10, 64)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	changesurl, err := models.GetOnlyChangesUrl(idNum, versionint)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	beego.Info(changesurl.ChangesUrl)
// 	c.Data["json"] = changesurl.ChangesUrl
// 	c.ServeJSON()
// }

//协作页面的显示
//补充权限判断
//补充token
func (c *OnlyController) OnlyOffice() {
	id := c.Ctx.Input.Param(":id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据附件id取得附件的prodid，路径
	onlyattachment, err := models.GetOnlyAttachbyId(idNum)
	if err != nil {
		beego.Error(err)
	}

	//docid——uid——me
	doc, err := models.Getdocbyid(onlyattachment.DocId)
	if err != nil {
		beego.Error(err)
	}

	var uname, useridstring, Permission string
	var myRes, roleRes [][]string
	v := c.GetSession("uname")
	// var role, userrole int
	myResall := e.GetPermissionsForUser("") //取出所有设置了权限的数据
	if v != nil {
		uname = v.(string)
		c.Data["Uname"] = v.(string)
		user, err := models.GetUserByUsername(uname)
		if err != nil {
			beego.Error(err)
		}
		useridstring = strconv.FormatInt(user.Id, 10)
		myRes = e.GetPermissionsForUser(useridstring)
		if doc.Uid == user.Id {
			Permission = "1"
		} else { //如果是登录用户，则设置了权限的文档不能看
			Permission = "1"
			for _, k := range myResall {
				if strconv.FormatInt(onlyattachment.Id, 10) == path.Base(k[1]) {
					Permission = "4"
				}
			}
			for _, k := range myRes {
				if strconv.FormatInt(onlyattachment.Id, 10) == path.Base(k[1]) {
					Permission = k[2]
				}
			}

			roles := e.GetRolesForUser(useridstring) //取出用户的所有角色
			for _, w1 := range roles {               //2018.4.30修改这个bug，这里原先w改为w1
				roleRes = e.GetPermissionsForUser(w1) //取出角色的所有权限，改为w1
				for _, k := range roleRes {
					// beego.Info(k)
					if id == path.Base(k[1]) {
						// docxarr[0].Permission = k[2]
						int1, err := strconv.Atoi(k[2])
						if err != nil {
							beego.Error(err)
						}
						int2, err := strconv.Atoi(Permission)
						if err != nil {
							beego.Error(err)
						}
						if int1 < int2 {
							Permission = k[2] //按最小值权限
						}
					}
				}
			}
		}
		c.Data["Uid"] = user.Id
		// userrole = user.Role
	} else { //如果用户没登录，则设置了权限的文档不能看
		Permission = "1"
		for _, k := range myResall { //所有设置了权限的不能看
			if strconv.FormatInt(onlyattachment.Id, 10) == path.Base(k[1]) {
				Permission = "4"
			}
		}
		c.Data["Uname"] = c.Ctx.Input.IP()
		c.Data["Uid"] = 0
	}

	// var myRes [][]string
	// if useridstring != "" {
	// 	myRes = e.GetPermissionsForUser(useridstring)
	// }
	// myResall := e.GetPermissionsForUser("") //取出所有设置了权限的数据
	// Permission = "1"
	// if useridstring != "" {
	// 	for _, k := range myResall {
	// 		if strconv.FormatInt(onlyattachment.Id, 10) == path.Base(k[1]) {
	// 			Permission = "4"
	// 		}
	// 	}
	// 	for _, k := range myRes {
	// 		if strconv.FormatInt(onlyattachment.Id, 10) == path.Base(k[1]) {
	// 			Permission = k[2]
	// 		}
	// 	}
	// } else { //如果用户没登录，则设置了权限的文档不能看
	// 	for _, k := range myResall { //所有设置了权限的不能看
	// 		if strconv.FormatInt(onlyattachment.Id, 10) == path.Base(k[1]) {
	// 			Permission = "4"
	// 		}
	// 	}
	// }
	// beego.Info(Permission)
	// In case edit is set to "false" and review is set to "true",
	// the document will be available in review mode only.
	if Permission == "1" {
		c.Data["Mode"] = "edit"
		c.Data["Edit"] = true
		c.Data["Review"] = true
	} else if Permission == "2" {
		c.Data["Mode"] = "edit"
		c.Data["Edit"] = false
		c.Data["Review"] = true
	} else if Permission == "3" {
		c.Data["Mode"] = "view"
		c.Data["Edit"] = false
		c.Data["Review"] = false
	} else if Permission == "4" {
		route := c.Ctx.Request.URL.String()
		c.Redirect("/roleerr?url="+route, 302)
		return
	}

	c.Data["Doc"] = onlyattachment
	c.Data["attachid"] = idNum
	c.Data["Key"] = strconv.FormatInt(onlyattachment.Updated.UnixNano(), 10)

	//构造[]history
	history, err := models.GetOnlyHistory(onlyattachment.Id)
	if err != nil {
		beego.Error(err)
	}

	onlyhistory := make([]history1, 0)
	onlychanges := make([]change, 0)
	onlychangesurl := make([]changesurl, 0)
	for _, v := range history {
		aa := make([]history1, 1)
		cc := make([]changesurl, 1)
		// aa[0].Created = v.Created
		aa[0].Key = v.HistoryKey
		aa[0].User.Id = strconv.FormatInt(v.UserId, 10) //

		if v.UserId != 0 {
			user := models.GetUserByUserId(v.UserId)
			aa[0].User.Name = user.Nickname
		}
		aa[0].Version = v.Version

		//取得changes
		changes, err := models.GetOnlyChanges(v.HistoryKey)
		if err != nil {
			beego.Error(err)
		}
		for _, v1 := range changes {
			bb := make([]change, 1)
			bb[0].Created = v1.Created
			bb[0].User.Id = v1.UserId
			bb[0].User.Name = v1.UserName
			onlychanges = append(onlychanges, bb...)
		}
		aa[0].Changes = onlychanges
		// aa[0].ChangesUrl = v.ChangesUrl
		aa[0].Created = v.Created

		cc[0].Version = v.Version
		cc[0].ChangesUrl = v.ChangesUrl
		onlychanges = make([]change, 0)
		onlyhistory = append(onlyhistory, aa...)
		onlychangesurl = append(onlychangesurl, cc...)
	}

	c.Data["onlyhistory"] = onlyhistory
	c.Data["changesurl"] = onlychangesurl

	historyversion, err := models.GetOnlyHistoryVersion(onlyattachment.Id)
	if err != nil {
		beego.Error(err)
	}
	var first int
	for _, v := range historyversion {
		if first < v.Version {
			first = v.Version
		}
	}
	c.Data["currentversion"] = first
	// beego.Info(first)

	if path.Ext(onlyattachment.FileName) == ".docx" || path.Ext(onlyattachment.FileName) == ".DOCX" {
		c.Data["fileType"] = "docx"
		c.Data["documentType"] = "text"
	} else if path.Ext(onlyattachment.FileName) == ".XLSX" || path.Ext(onlyattachment.FileName) == ".xlsx" {
		c.Data["fileType"] = "xlsx"
		c.Data["documentType"] = "spreadsheet"
	} else if path.Ext(onlyattachment.FileName) == ".pptx" || path.Ext(onlyattachment.FileName) == ".PPTX" {
		c.Data["fileType"] = "pptx"
		c.Data["documentType"] = "presentation"
	} else if path.Ext(onlyattachment.FileName) == ".doc" || path.Ext(onlyattachment.FileName) == ".DOC" {
		c.Data["fileType"] = "doc"
		c.Data["documentType"] = "text"
	} else if path.Ext(onlyattachment.FileName) == ".txt" || path.Ext(onlyattachment.FileName) == ".TXT" {
		c.Data["fileType"] = "txt"
		c.Data["documentType"] = "text"
	} else if path.Ext(onlyattachment.FileName) == ".XLS" || path.Ext(onlyattachment.FileName) == ".xls" {
		c.Data["fileType"] = "xls"
		c.Data["documentType"] = "spreadsheet"
	} else if path.Ext(onlyattachment.FileName) == ".csv" || path.Ext(onlyattachment.FileName) == ".CSV" {
		c.Data["fileType"] = "csv"
		c.Data["documentType"] = "spreadsheet"
	} else if path.Ext(onlyattachment.FileName) == ".ppt" || path.Ext(onlyattachment.FileName) == ".PPT" {
		c.Data["fileType"] = "ppt"
		c.Data["documentType"] = "presentation"
	} else if path.Ext(onlyattachment.FileName) == ".pdf" || path.Ext(onlyattachment.FileName) == ".PDF" {
		c.Data["fileType"] = "pdf"
		c.Data["documentType"] = "text"
		c.Data["Mode"] = "view"
	}

	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		beego.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		// c.TplName = "onlyoffice/onlyoffice.tpl"
		c.Data["Type"] = "mobile"
	} else {
		// beego.Info("电脑端！")
		// c.TplName = "onlyoffice/onlyoffice.tpl"
		c.Data["Type"] = "desktop"
	}
	c.TplName = "onlyoffice/onlyoffice.tpl"
}

//协作页面的保存和回调
//关闭浏览器标签后获取最新文档保存到文件夹
func (c *OnlyController) UrltoCallback() {
	// pk1 := c.Ctx.Input.RequestBody
	id := c.Input().Get("id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据附件id取得附件的prodid，路径
	onlyattachment, err := models.GetOnlyAttachbyId(idNum)
	if err != nil {
		beego.Error(err)
	}

	var callback Callback
	json.Unmarshal(c.Ctx.Input.RequestBody, &callback)
	// beego.Info(string(c.Ctx.Input.RequestBody))

	if callback.Status == 1 || callback.Status == 4 {
		c.Data["json"] = map[string]interface{}{"error": 0}
		c.ServeJSON()
	} else if callback.Status == 2 {
		resp, err := http.Get(callback.Url)
		if err != nil {
			beego.Error(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			beego.Error(err)
		}
		defer resp.Body.Close()
		if err != nil {
			beego.Error(err)
		}

		// f, err := os.OpenFile(".\\attachment\\onlyoffice\\"+onlyattachment.FileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		f, err := os.Create(".\\attachment\\onlyoffice\\" + onlyattachment.FileName)
		if err != nil {
			beego.Error(err)
		}
		defer f.Close()
		_, err = f.Write(body)
		// _, err = f.WriteString(str)
		// _, err = io.Copy(body, f)
		if err != nil {
			beego.Error(err)
		} else {
			//更新附件的时间和changesurl
			err = models.UpdateOnlyAttachment(idNum)
			if err != nil {
				beego.Error(err)
			}

			//写入历史版本数据
			array := strings.Split(callback.Changesurl, "&")
			Expires1 := strings.Split(array[1], "=")
			Expires := Expires1[1]
			Expirestime, err := strconv.ParseInt(Expires, 10, 64)
			if err != nil {
				beego.Error(err)
			}
			//获取本地location
			// toBeCharge := "2015-01-01 00:00:00"                             //待转化为时间戳的字符串 注意 这里的小时和分钟还要秒必须写 因为是跟着模板走的 修改模板的话也可以不写
			// timeLayout := "2006-01-02T15:04:05.999Z" //转化所需模板

			// loc, _ := time.LoadLocation("Local") //重要：获取时区
			// theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
			// sr := theTime.Unix()                                            //转化为时间戳 类型是int64                                               //打印输出时间戳 1420041600

			//时间戳转日期
			dataTimeStr := time.Unix(Expirestime, 0) //.Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
			// t, _ := time.Parse(timeLayout, callback.Lastsave)
			// beego.Info(callback.Lastsave)
			//写入历史版本
			historyversion, err := models.GetOnlyHistoryVersion(onlyattachment.Id)
			if err != nil {
				beego.Error(err)
			}
			var first int
			for _, v := range historyversion {
				if first < v.Version {
					first = v.Version
				}
			}
			// beego.Info(callback.Users[0])
			_, err1, err2 := models.AddOnlyHistory(onlyattachment.Id, callback.Actions[0].Userid, first+1, callback.Key, callback.Changesurl, dataTimeStr, callback.Lastsave)
			if err1 != nil {
				beego.Error(err1)
			}
			if err2 != nil {
				beego.Error(err2)
			}
			//写入changes
			for _, v := range callback.History.Changes {
				_, err1, err2 = models.AddOnlyChanges(callback.Key, v.User.Id, v.User.Name, v.Created)
				if err1 != nil {
					beego.Error(err1)
				}
				if err2 != nil {
					beego.Error(err2)
				}
			}
			//更新文档更新时间
			err = models.UpdateDocTime(onlyattachment.DocId)
			if err != nil {
				beego.Error(err)
			}
		}
		c.Data["json"] = map[string]interface{}{"error": 0}
		c.ServeJSON()
	} else if callback.Status == 3 {
		//更新附件的时间和changesurl
		err = models.UpdateOnlyAttachment(idNum)
		if err != nil {
			beego.Error(err)
		}
		c.Data["json"] = map[string]interface{}{"error": 0}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"error": 0}
		c.ServeJSON()
	}
}

//批量添加一对一模式
//要避免同名覆盖的严重bug！！！！
func (c *OnlyController) AddOnlyAttachment() {
	//取得客户端用户名
	v := c.GetSession("uname")
	var user models.User
	var err error
	if v != nil {
		uname := v.(string)
		user, err = models.GetUserByUsername(uname)
		if err != nil {
			beego.Error(err)
		}
	}

	var filepath, DiskDirectory, Url string

	err = os.MkdirAll(".\\attachment\\onlyoffice\\", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		beego.Error(err)
	}
	DiskDirectory = ".\\attachment\\onlyoffice\\"
	Url = "/attachment/onlyoffice/"

	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
	}
	if h != nil {
		//保存附件
		//将附件的编号和名称写入数据库
		_, code, title, _, _, _, _ := Record(h.Filename)
		// filename1, filename2 := SubStrings(attachment)
		//当2个文件都取不到filename1的时候，数据库里的tnumber的唯一性检查出错。
		if code == "" {
			code = title //如果编号为空，则用文件名代替，否则多个编号为空导致存入数据库唯一性检查错误
		}
		//存入成果数据库
		//如果编号重复，则不写入，只返回Id值。
		//根据id添加成果code, title, label, principal, content string, projectid int64
		prodlabel := c.Input().Get("prodlabel")
		prodprincipal := c.Input().Get("prodprincipal")
		// type Duration int64
		// const (
		// 	Nanosecond  Duration = 1
		// 	Microsecond          = 1000 * Nanosecond
		// 	Millisecond          = 1000 * Microsecond
		// 	Second               = 1000 * Millisecond
		// 	Minute               = 60 * Second
		// 	Hour                 = 60 * Minute
		// )
		// hours := 8
		inputdate := c.Input().Get("proddate")
		// beego.Info(inputdate)
		var t1, end time.Time
		// var convdate1, convdate2 string
		const lll = "2006-01-02"
		if len(inputdate) > 9 { //如果是datepick获取的时间，则不用加8小时
			t1, err = time.Parse(lll, inputdate) //这里t1要是用t1:=就不是前面那个t1了
			if err != nil {
				beego.Error(err)
			}
			// convdate := t1.Format(lll)
			// catalog.Datestring = convdate
			end = t1
			// t1 = printtime.Add(+time.Duration(hours) * time.Hour)
		} else { //如果取系统时间，则需要加8小时
			date := time.Now()
			convdate := date.Format(lll)
			// catalog.Datestring = convdate
			date, err = time.Parse(lll, convdate)
			if err != nil {
				beego.Error(err)
			}
			end = date
		}
		// beego.Info(end)
		prodId, err := models.AddDoc(code, title, prodlabel, prodprincipal, end, user.Id)
		if err != nil {
			beego.Error(err)
		}
		//改名，替换文件名中的#和斜杠
		title = strings.Replace(title, "#", "号", -1)
		title = strings.Replace(title, "/", "-", -1)
		FileSuffix := path.Ext(h.Filename)

		filepath = DiskDirectory + "\\" + code + title + FileSuffix
		//把成果id作为附件的parentid，把附件的名称等信息存入附件数据库
		//如果附件名称相同，则不能上传，数据库添加
		attachmentname := code + title + FileSuffix

		_, _, err2 := models.AddOnlyAttachment(attachmentname, 0, 0, prodId)
		if err2 != nil {
			beego.Error(err2)
		} else {
			//存入文件夹
			//判断文件是否存在，如果不存在就写入
			if _, err := os.Stat(filepath); err != nil {
				// beego.Info(err)
				if os.IsNotExist(err) {
					// beego.Info(err)
					//return false
					err = c.SaveToFile("file", filepath) //存文件
					if err != nil {
						beego.Error(err)
					}
					c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "title": h.Filename, "original": h.Filename, "url": Url + "/" + h.Filename}
					c.ServeJSON()
				}
			} else {
				// beego.Info(err)
				c.Data["json"] = map[string]interface{}{"state": "error"}
				c.ServeJSON()
			}

		}
	}
}

//协作页面下载的文档，采用绝对路径型式
func (c *OnlyController) DownloadDoc() {
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	//4.取得客户端用户名
	// var uname, useridstring string
	// v := c.GetSession("uname")
	// if v != nil {
	// 	uname = v.(string)
	// 	c.Data["Uname"] = v.(string)
	// 	user, err := models.GetUserByUsername(uname)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// 	useridstring = strconv.FormatInt(user.Id, 10)
	// }

	//1.url处理中文字符路径，[1:]截掉路径前面的/斜杠
	// filePath := path.Base(ctx.Request.RequestURI)
	filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:]) //  attachment/SL2016测试添加成果/A/FB/1/Your First Meteor Application.pdf
	if err != nil {
		beego.Error(err)
	}

	// fileext := path.Ext(filePath)
	//根据路由path.Dir——再转成数组strings.Split——查出项目id——加上名称——查出下级id
	// beego.Info(path.Dir(filePath))
	// filepath1 := path.Dir(filePath)
	// array := strings.Split(filepath1, "/")
	// beego.Info(strings.Split(filepath1, "/"))
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)
}

//文档管理页面下载文档
func (c *OnlyController) Download() {
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	//4.取得客户端用户名
	// var uname, useridstring string
	// v := c.GetSession("uname")
	// if v != nil {
	// 	uname = v.(string)
	// 	c.Data["Uname"] = v.(string)
	// 	user, err := models.GetUserByUsername(uname)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// 	useridstring = strconv.FormatInt(user.Id, 10)
	// }
	docid := c.Ctx.Input.Param(":id")
	//pid转成64为
	idNum, err := strconv.ParseInt(docid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据成果id取得所有附件
	attachments, err := models.GetOnlyAttachments(idNum)
	if err != nil {
		beego.Error(err)
	}

	filePath := "attachment/onlyoffice/" + attachments[0].FileName
	//1.url处理中文字符路径，[1:]截掉路径前面的/斜杠
	// filePath := path.Base(ctx.Request.RequestURI)
	// filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:]) //  attachment/SL2016测试添加成果/A/FB/1/Your First Meteor Application.pdf
	// if err != nil {
	// 	beego.Error(err)
	// }

	// fileext := path.Ext(filePath)
	//根据路由path.Dir——再转成数组strings.Split——查出项目id——加上名称——查出下级id
	// beego.Info(path.Dir(filePath))
	// filepath1 := path.Dir(filePath)
	// array := strings.Split(filepath1, "/")
	// beego.Info(strings.Split(filepath1, "/"))
	http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)
}

//编辑成果信息
func (c *OnlyController) UpdateDoc() {
	id := c.Input().Get("pid")
	code := c.Input().Get("code")
	title := c.Input().Get("title")
	label := c.Input().Get("label")
	principal := c.Input().Get("principal")

	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	inputdate := c.Input().Get("proddate")
	var t1, end time.Time
	const lll = "2006-01-02"
	if len(inputdate) > 9 { //如果是datepick获取的时间，则不用加8小时
		t1, err = time.Parse(lll, inputdate) //这里t1要是用t1:=就不是前面那个t1了
		if err != nil {
			beego.Error(err)
		}
		end = t1
		// t1 = printtime.Add(+time.Duration(hours) * time.Hour)
	} else { //如果取系统时间，则需要加8小时
		date := time.Now()
		convdate := date.Format(lll)
		date, err = time.Parse(lll, convdate)
		if err != nil {
			beego.Error(err)
		}
		end = date
	}
	//根据id添加成果
	err = models.UpdateDoc(idNum, code, title, label, principal, end)
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = "ok"
	c.ServeJSON()
}

//删除成果，包含成果里的附件。删除附件用attachment中的
func (c *OnlyController) DeleteDoc() {
	ids := c.GetString("ids")
	array := strings.Split(ids, ",")
	for _, v := range array {
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		//循环删除成果
		//根据成果id取得所有附件
		attachments, err := models.GetOnlyAttachments(idNum)
		if err != nil {
			beego.Error(err)
		}
		for _, w := range attachments {
			//取得附件的成果id
			attach, err := models.GetOnlyAttachbyId(w.Id)
			if err != nil {
				beego.Error(err)
			}

			path := ".\\attachment\\onlyoffice\\" + attach.FileName
			//删除附件
			err = os.Remove(path)
			if err != nil {
				beego.Error(err)
			}
			//删除附件数据表
			err = models.DeleteOnlyAttachment(w.Id)
			if err != nil {
				beego.Error(err)
			}
		}

		err = models.DeleteDoc(idNum) //删除成果数据表
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	}
}

//onlyoffice权限管理
//添加用户和角色的权限
//先删除这个文档id下所有permission，再添加新的。
func (c *OnlyController) Addpermission() {
	// roleids := c.GetString("roleids")
	// rolearray := strings.Split(roleids, ",")
	// userids := c.GetString("userids")
	// userarray := strings.Split(userids, ",")
	ids := c.GetString("ids")
	// beego.Info(ids)
	tt := []byte(ids)
	var rolepermission []Rolepermission
	json.Unmarshal(tt, &rolepermission)
	// beego.Info(rolepermission)
	docid := c.GetString("docid")
	//id转成64位
	idNum, err := strconv.ParseInt(docid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	//根据成果id取得所有附件——这里只取第一个
	attachments, err := models.GetOnlyAttachments(idNum)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(docid)
	// action := "get"
	// var suf string
	suf := ".*"
	var success bool

	//先删除这个文档所有的permission
	// var paths []beegoormadapter.CasbinRule
	// o := orm.NewOrm()
	// qs := o.QueryTable("casbin_rule")
	// _, err = qs.Filter("v1", "/onlyoffice/"+strconv.FormatInt(attachments[0].Id, 10)).All(&paths)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// for _, v := range paths {
	e.RemoveFilteredPolicy(1, "/onlyoffice/"+strconv.FormatInt(attachments[0].Id, 10))
	// success = e.RemovePolicy(v.V0, "/onlyoffice/"+strconv.FormatInt(attachments[0].Id, 10))
	// beego.Info(success)
	// success = e.RemoveGroupingPolicy(v.V0, "/onlyoffice/"+strconv.FormatInt(attachments[0].Id, 10))
	// beego.Info(success)
	// }
	// var paths []beegoormadapter.CasbinRule
	//因为上面的代码无法删除数据库
	// o := orm.NewOrm()
	// qs := o.QueryTable("casbin_rule")
	// _, err = qs.Filter("v1", "/onlyoffice/"+strconv.FormatInt(attachments[0].Id, 10)).Delete()
	// if err != nil {
	// 	beego.Error(err)
	// }
	// _, err = o.Delete(&paths)
	// if err != nil {
	// 	beego.Error(err)
	// }
	//再添加permission
	for _, v1 := range rolepermission {
		// beego.Info(v1.Id)
		if v1.Rolenumber != "" { //存储角色id
			success = e.AddPolicy("role_"+strconv.FormatInt(v1.Id, 10), "/onlyoffice/"+strconv.FormatInt(attachments[0].Id, 10), v1.Permission, suf)
		} else { //存储用户id
			success = e.AddPolicy(strconv.FormatInt(v1.Id, 10), "/onlyoffice/"+strconv.FormatInt(attachments[0].Id, 10), v1.Permission, suf)
		}
		//这里应该用AddPermissionForUser()，来自casbin\rbac_api.go
	}
	if success == true {
		c.Data["json"] = "ok"
	} else {
		c.Data["json"] = "wrong"
	}
	c.ServeJSON()
}

//查询一个文档，哪些用户和角色拥有什么样的权限
//用casbin的内置方法，不应该用查询数据库方法
func (c *OnlyController) Getpermission() {
	docid := c.GetString("docid")
	//id转成64位
	idNum, err := strconv.ParseInt(docid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//根据成果id取得所有附件
	attachments, err := models.GetOnlyAttachments(idNum)
	if err != nil {
		beego.Error(err)
	}
	// var users []beegoormadapter.CasbinRule
	rolepermission := make([]Rolepermission, 0)
	for _, w := range attachments {
		// o := orm.NewOrm()
		// qs := o.QueryTable("casbin_rule")
		// _, err = qs.Filter("PType", "p").Filter("v1", "/onlyoffice/"+strconv.FormatInt(w.Id, 10)).All(&users)
		// if err != nil {
		// 	beego.Error(err)
		// }
		users := e.GetFilteredPolicy(1, "/onlyoffice/"+strconv.FormatInt(w.Id, 10))
		// beego.Info(users)
		for _, v := range users {
			rolepermission1 := make([]Rolepermission, 1)
			if strings.Contains(v[0], "role_") { //是角色
				// beego.Info(v.V0)
				roleid := strings.Replace(v[0], "role_", "", -1)
				//id转成64位
				roleidNum, err := strconv.ParseInt(roleid, 10, 64)
				if err != nil {
					beego.Error(err)
				}
				// beego.Info(roleidNum)
				role := models.GetRoleByRoleId(roleidNum)
				// beego.Info(role)
				rolepermission1[0].Id = roleidNum
				rolepermission1[0].Name = role.Rolename
				rolepermission1[0].Rolenumber = role.Rolenumber
				rolepermission1[0].Permission = v[2]
				// rolepermission = append(rolepermission, rolepermission1...)
			} else { //是用户
				// rolepermission1 := make([]Rolepermission, 1)
				//id转成64位
				uidNum, err := strconv.ParseInt(v[0], 10, 64)
				if err != nil {
					beego.Error(err)
				}
				user := models.GetUserByUserId(uidNum)
				rolepermission1[0].Id = uidNum
				rolepermission1[0].Name = user.Nickname
				rolepermission1[0].Permission = v[2]
			}
			rolepermission = append(rolepermission, rolepermission1...)
			// rolepermission1 = make([]Rolepermission, 0)
		}
		// myRes := e.GetPermissionsForUser(roleid)
	}
	c.Data["json"] = rolepermission
	c.ServeJSON()
}

// 原来golang下载文件这么简单
// 分为三步
// 第一步 用http.get 获取到res.Body 这个流
// 第二步 用os.Creat创建文件并取到文件
// 第三步 io.Copy把得到的res.Body拷贝到文件的流里面

// import (
//     "fmt"
//     "github.com/PuerkitoBio/goquery"
//     "io"
//     "net/http"
//     "os"
// )

// func main() {
//     x, _ := goquery.NewDocument("http://www.fengyun5.com/Sibao/600/1.html")
//     urls, _ := x.Find("#content img").Attr("src")
//     res, _ := http.Get(urls)
//     file, _ := os.Create("xxx.jpg")
//     io.Copy(file, res.Body)
//     fmt.Println("下载完成！")
// }

// func downloadFile(fileFullPath string, res *restful.Response) {
//     file, err := os.Open(fileFullPath)

//     if err != nil {
//         res.WriteEntity(_dto.ErrorDto{Err: err})
//         return
//     }

//     defer file.Close()
//     fileName := path.Base(fileFullPath)
//     fileName = url.QueryEscape(fileName) // 防止中文乱码
//     res.AddHeader("Content-Type", "application/octet-stream")
//     res.AddHeader("content-disposition", "attachment; filename=\""+fileName+"\"")
//     _, error := io.Copy(res.ResponseWriter, file)
//     if error != nil {
//         res.WriteErrorString(http.StatusInternalServerError, err.Error())
//         return
//     }
// }

//判断某个文件是否存在 package main

// import (
//     "fmt"
//     "os"
// )

// func main() {
//     originalPath := "test.txt"
//     result := Exists(originalPath)
//     fmt.Println(result)
// }

// func Exists(name string) bool {
//     if _, err := os.Stat(name); err != nil {
//         if os.IsNotExist(err) {
//             return false
//         }
//     }
//     return true
// }

//****模拟表单上传附件********************
//***最开始先httpget到onlyoffice的output.doc，
//再读出来
//再新建一个文件
//把读出来的写入新建的文件
//把新建的文件模拟表单上传
//接受上传的文件，保存到attachment对应的onlyoffice文件夹中
// req := httplib.Get("http://192.168.99.100:9000/cache/files/E7FAFC9C22A8_5080/output.docx/output.docx?md5=-fOxyXjPpAyxS_QZdt5jBw==&expires=1520306521&disposition=attachment&ooname=output.docx")
// str, err := req.Bytes() //req.String()
// 打开根目录文件f, err := os.OpenFile("tt", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm) //可读写，追加的方式打开（或创建文件）
// f, err := os.Create("")

// 模拟表单上传bodyBuf := &bytes.Buffer{}
// bodyWriter := multipart.NewWriter(bodyBuf)
//关键的一步操作
// fileWriter, err := bodyWriter.CreateFormFile("uploadfile", onlyattachment.FileName)
// if err != nil {
// 	beego.Error(err)
// }
//打开文件句柄操作
// fh, err := os.Open("tt")
// if err != nil {
// 	beego.Error(err)
// }
// defer fh.Close()

//iocopy
// _, err = io.Copy(fileWriter, fh)
// if err != nil {
// 	beego.Error(err)
// }

// contentType := bodyWriter.FormDataContentType()
// bodyWriter.Close()

// resp1, err := http.Post("http://192.168.99.1/onlyoffice/post?id="+id, contentType, bodyBuf)
// if err != nil {
// 	beego.Error(err)
// }
// defer resp1.Body.Close()

// err = c.SaveToFile("tt", "\\attachment\\wiki\\2018February\\1.doc")
// if err != nil {
// 	beego.Error(err)
// }

// resp_body, err := ioutil.ReadAll(resp1.Body)
// if err != nil {
// 	beego.Error(err)
// }
// beego.Info(resp1.Status)
// beego.Info(string(resp_body))
// fmt.Println(resp1.Status)
// fmt.Println(string(resp_body))
// return nil
//******************************

// // 创建表单文件
//    // CreateFormFile 用来创建表单，第一个参数是字段名，第二个参数是文件名
//    buf := new(bytes.Buffer)
//    writer := multipart.NewWriter(buf)
//    formFile, err := writer.CreateFormFile("uploadfile", "test.jpg")
//    if err != nil {
//        log.Fatalf("Create form file failed: %s\n", err)
//    }

//    // 从文件读取数据，写入表单
//    srcFile, err := os.Open("test.jpg")
//    if err != nil {
//        log.Fatalf("%Open source file failed: s\n", err)
//    }
//    defer srcFile.Close()
//    _, err = io.Copy(formFile, srcFile)
//    if err != nil {
//        log.Fatalf("Write to form file falied: %s\n", err)
//    }
//    // 发送表单
//    contentType := writer.FormDataContentType()
//    writer.Close() // 发送之前必须调用Close()以写入结尾行
//    _, err = http.Post("http://localhost:9090/upload", contentType, buf)
//    if err != nil {
//        log.Fatalf("Post failed: %s\n", err)
//    }

//——关闭浏览器最后一个标签，存储文件——取消这种方法
// func (c *OnlyController) PostOnlyoffice() {
// 	id := c.Input().Get("id")
// 	//pid转成64为
// 	idNum, err := strconv.ParseInt(id, 10, 64)
// 	if err != nil {
// 		beego.Error(err)
// 	}

// 	//获取上传的文件
// 	_, h, err := c.GetFile("uploadfile")
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	if h != nil {
// 		//存入文件夹
// 		err = c.SaveToFile("uploadfile", ".\\attachment\\onlyoffice\\"+h.Filename)
// 		if err != nil {
// 			beego.Error(err)
// 		} else {
// 			err = models.UpdateOnlyAttachment(idNum)
// 			if err != nil {
// 				beego.Error(err)
// 			}
// 		}
// 	}
// }
