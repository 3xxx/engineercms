package controllers

import (
	"encoding/json"
	"encoding/xml"
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	// beego "github.com/beego/beego/v2/adapter"
	// "github.com/beego/beego/v2/adapter/orm"
	// "github.com/casbin/beego-orm-adapter"
	// "baliance.com/gooxml/document"
	// "github.com/unidoc/unioffice/common"
	"github.com/3xxx/engineercms/controllers/utils/ziptil"
	"github.com/unidoc/unioffice/document"
	// "github.com/unidoc/unioffice/measurement"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
	// "mime/multipart"
	"bytes"
	"github.com/beego/beego/v2/adapter/httplib"
	// "github.com/beego/beego/v2/adapter/utils/pagination"
	// "crypto/aes"
	// "crypto/cipher"
	// "io"
	// "github.com/beego/beego/v2/adapter/session"
)

var OnlyUsers []string

type OnlyController struct {
	web.Controller
}

type Callback struct {
	Key           string    `json:"key"`
	Status        int       `json:"status"`
	Url           string    `json:"url"`
	Changesurl    string    `json:"changesurl"`
	History       history1  `json:"history"`
	Users         []string  `json:"users"`
	Actions       []action  `json:"actions"`
	Lastsave      time.Time `json:"lastsave"`
	Notmodified   bool      `json:"notmodified"`
	Forcesavetype int       `json:"forcesavetype"`
}

type action struct {
	Type   int   `json:"type"`
	Userid int64 `json:"userid"`
}

type history1 struct {
	ServerVersion string    `json:"serverVersion"`
	Changes       []change  `json:"changes"`
	Created       time.Time `json:"created"`
	ChangesUrl    string    `json:"changesurl"`
	FileUrl       string    `json:"fileurl"`
	Key           string    `json:"key"`
	// ChangesUrl    string
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

// 构造changesurl结构
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
	Uname     string
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

// 权限表提交的table中json数据解析成struct
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

// 文档管理页面
func (c *OnlyController) Get() {
	//取得客户端用户名
	// v := c.GetSession("uname")
	// if v != nil {
	// 	uname := v.(string)
	// 	user, err := models.GetUserByUsername(uname)
	// 	if err != nil {
	// 		logs.Error(err)
	// 	}
	// 	c.Data["Uid"] = user.Id
	// } else {
	// 	c.Data["Uid"] = 0
	// }
	// username, role := checkprodRole(c.Ctx)
	// roleint, err := strconv.Atoi(role)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// if role == "1" {
	// 	c.Data["IsAdmin"] = true
	// } else if roleint > 1 && roleint < 5 {
	// 	c.Data["IsLogin"] = true
	// } else {
	// 	c.Data["IsAdmin"] = false
	// 	c.Data["IsLogin"] = false
	// }
	// c.Data["Username"] = username
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid

	u := c.Ctx.Input.UserAgent()
	// re := regexp.MustCompile("Trident")
	// loc := re.FindStringIndex(u)
	// loc[0] > 1
	c.Data["IsOnlyOffice"] = true
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
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

// @Title get only documents list
// @Description get documents by page & limit
// @Param page query string  true "The page for document list"
// @Param limit query string  true "The limit for document list"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getonlydocs [get]
//vue文档列表数据
// func (c *ArticleController) GetOnlyDocs() {
// 	var offset, limit1, page1 int
// 	var err error
// 	limit := c.GetString("limit")
// 	if limit == "" {
// 		limit1 = 0
// 	} else {
// 		limit1, err = strconv.Atoi(limit)
// 		if err != nil {
// 			logs.Error(err)
// 		}
// 	}
// 	page := c.GetString("page")
// 	if page == "" {
// 		limit1 = 0
// 		page1 = 1
// 	} else {
// 		page1, err = strconv.Atoi(page)
// 		if err != nil {
// 			logs.Error(err)
// 		}
// 	}

// 	if page1 <= 1 {
// 		offset = 0
// 	} else {
// 		offset = (page1 - 1) * limit1
// 	}

// 	username, role, uid := Authorizer(c.Ctx)

// 	//这里用jion，取得uname和attachment和permission
// 	docs, err := models.GetDocList(offset, limit1)
// 	if err != nil {
// 		logs.Error(err)
// 	}

// 	//1.anonymous，首先查permission表，
// 	//所有均为4
// 	//permission表中针对anonymous的权限，以它为准
// 	//设置了权限的文档，anonymous则为4（下面这条包含了）
// 	//permission表中没有设置权限的文档，开放

// 	//2.isme
// 	//如果设置了isme权限，以它为准
// 	//如果没有，则为1

// 	//3.everyone
// 	//如果设置了，则以它为准
// 	//如果没有设置，则没有

// 	//4.全部文档AllDocs
// 	for _, v := range Attachments {
// 		// beego.Info(v.FileName)
// 		// fileext := path.Ext(v.FileName)
// 		docxarr := make([]DocxLink, 1)
// 		docxarr[0].Permission = "1"
// 		//查询v.Id是否和myres的V1路径后面的id一致，如果一致，则取得V2（权限）
// 		//查询用户具有的权限
// 		// beego.Info(useridstring)
// 		if useridstring != "0" { //如果是登录用户，则设置了权限的文档不能看
// 			// beego.Info(myRes)
// 			// myRes1 := e.GetPermissionsForUser("") //取出所有设置了权限的数据
// 			if w.Uid != uid { //如果不是作者本人
// 				for _, k := range myResall { //所有设置了权限的都不能看
// 					// beego.Info(k)
// 					if strconv.FormatInt(v.Id, 10) == path.Base(k[1]) {
// 						docxarr[0].Permission = "4"
// 					}
// 				}

// 				for _, k := range myRes { //如果与登录用户对应上，则赋予权限
// 					// beego.Info(k)
// 					if strconv.FormatInt(v.Id, 10) == path.Base(k[1]) {
// 						docxarr[0].Permission = k[2]
// 					}
// 				}
// 				roles := e.GetRolesForUser(useridstring) //取出用户的所有角色
// 				for _, w1 := range roles {               //2018.4.30修改这个bug，这里原先w改为w1
// 					roleRes = e.GetPermissionsForUser(w1) //取出角色的所有权限，改为w1
// 					for _, k := range roleRes {
// 						// beego.Info(k)
// 						if strconv.FormatInt(v.Id, 10) == path.Base(k[1]) {
// 							// docxarr[0].Permission = k[2]
// 							int1, err := strconv.Atoi(k[2])
// 							if err != nil {
// 								logs.Error(err)
// 							}
// 							int2, err := strconv.Atoi(docxarr[0].Permission)
// 							if err != nil {
// 								logs.Error(err)
// 							}
// 							if int1 < int2 {
// 								docxarr[0].Permission = k[2] //按最小值权限
// 							}
// 							//补充everyone
// 						}
// 					}
// 				}
// 			} //如果是用户自己的文档，则permission为1，默认
// 		} else { //如果用户没登录，则设置了权限的文档不能看
// 			for _, k := range myResall { //所有设置了权限的不能看
// 				// beego.Info(k)
// 				if strconv.FormatInt(v.Id, 10) == path.Base(k[1]) {
// 					// beego.Info(i)
// 					// beego.Info(strconv.FormatInt(v.Id, 10))
// 					// beego.Info(path.Base(k[1]))
// 					docxarr[0].Permission = "4"
// 					// beego.Info(strconv.FormatInt(v.Id, 10))
// 					// beego.Info(path.Base(k[1]))
// 				}
// 				//如果有everyone权限，则按everyone的
// 			}
// 		}

// 		docxarr[0].Id = v.Id
// 		docxarr[0].Title = v.FileName
// 		if path.Ext(v.FileName) == ".docx" || path.Ext(v.FileName) == ".DOCX" || path.Ext(v.FileName) == ".doc" || path.Ext(v.FileName) == ".DOC" {
// 			docxarr[0].Suffix = "docx"
// 		} else if path.Ext(v.FileName) == ".wps" || path.Ext(v.FileName) == ".WPS" {
// 			docxarr[0].Suffix = "docx"
// 		} else if path.Ext(v.FileName) == ".XLSX" || path.Ext(v.FileName) == ".xlsx" || path.Ext(v.FileName) == ".XLS" || path.Ext(v.FileName) == ".xls" {
// 			docxarr[0].Suffix = "xlsx"
// 			// xlsxarr := make([]XlsxLink, 1)
// 			// xlsxarr[0].Id = v.Id
// 			// xlsxarr[0].Title = v.FileName
// 			// Xlsxslice = append(Xlsxslice, xlsxarr...)
// 		} else if path.Ext(v.FileName) == ".ET" || path.Ext(v.FileName) == ".et" {
// 			docxarr[0].Suffix = "xlsx"
// 		} else if path.Ext(v.FileName) == ".pptx" || path.Ext(v.FileName) == ".PPTX" || path.Ext(v.FileName) == ".ppt" || path.Ext(v.FileName) == ".PPT" {
// 			docxarr[0].Suffix = "pptx"
// 			// pptxarr := make([]PptxLink, 1)
// 			// pptxarr[0].Id = v.Id
// 			// pptxarr[0].Title = v.FileName
// 			// Pptxslice = append(Pptxslice, pptxarr...)
// 		} else if path.Ext(v.FileName) == ".DPS" || path.Ext(v.FileName) == ".dps" {
// 			docxarr[0].Suffix = "pptx"
// 		} else if path.Ext(v.FileName) == ".pdf" || path.Ext(v.FileName) == ".PDF" {
// 			docxarr[0].Suffix = "pdf"
// 		} else if path.Ext(v.FileName) == ".txt" || path.Ext(v.FileName) == ".TXT" {
// 			docxarr[0].Suffix = "txt"
// 		}
// 		Docxslice = append(Docxslice, docxarr...)
// 	}
// 	linkarr[0].Docxlink = Docxslice
// 	// linkarr[0].Xlsxlink = Xlsxslice
// 	// linkarr[0].Pptxlink = Pptxslice
// 	Docxslice = make([]DocxLink, 0) //再把slice置0
// 	// Xlsxslice = make([]XlsxLink, 0) //再把slice置0
// 	// Pptxslice = make([]PptxLink, 0)
// 	//无权限的不显示
// 	//如果permission=4，则不赋值给link
// 	if linkarr[0].Docxlink[0].Permission != "4" {
// 		link = append(link, linkarr...)
// 	}

// 	c.Data["json"] = link //products
// 	c.ServeJSON()
// }

// 提供给列表页的table中json数据
func (c *OnlyController) GetData() {
	//1.取得客户端用户名
	// var uname, useridstring string
	// var user models.User
	var err error
	// v := c.GetSession("uname")
	// var role, userrole int
	// if v != nil {
	// 	uname = v.(string)
	// 	// c.Data["Uname"] = v.(string)
	// 	user, err = models.GetUserByUsername(uname)
	// 	if err != nil {
	// 		logs.Error(err)
	// 	}
	// 	c.Data["Uid"] = user.Id
	// 	// userrole = user.Role
	// 	useridstring = strconv.FormatInt(user.Id, 10)
	// }
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	useridstring := strconv.FormatInt(uid, 10)
	//增加admin，everyone，isme

	var myRes, roleRes [][]string
	if useridstring != "0" {
		myRes = e.GetPermissionsForUser(useridstring)
		// beego.Info(myRes)
	}
	myResall := e.GetPermissionsForUser("") //取出所有设置了权限的数据

	docs, err := models.GetDocs()
	if err != nil {
		logs.Error(err)
	}

	link := make([]OnlyLink, 0)
	Docxslice := make([]DocxLink, 0)
	for _, w := range docs {
		linkarr := make([]OnlyLink, 1)
		linkarr[0].Id = w.Id
		linkarr[0].Code = w.Code
		linkarr[0].Title = w.Title
		linkarr[0].Label = w.Label
		linkarr[0].End = w.End
		linkarr[0].Principal = w.Principal
		linkarr[0].Uid = w.Uid
		user := models.GetUserByUserId(w.Uid)
		linkarr[0].Uname = user.Nickname
		linkarr[0].Created = w.Created
		linkarr[0].Updated = w.Updated

		Attachments, err := models.GetOnlyAttachments(w.Id)
		if err != nil {
			logs.Error(err)
		}
		//docid——me——1
		for _, v := range Attachments {
			// beego.Info(v.FileName)
			// fileext := path.Ext(v.FileName)
			docxarr := make([]DocxLink, 1)
			docxarr[0].Permission = "1"
			//查询v.Id是否和myres的V1路径后面的id一致，如果一致，则取得V2（权限）
			//查询用户具有的权限
			// beego.Info(useridstring)
			if useridstring != "0" { //如果是登录用户，则设置了权限的文档不能看
				// beego.Info(myRes)
				// myRes1 := e.GetPermissionsForUser("") //取出所有设置了权限的数据
				if w.Uid != uid { //如果不是作者本人
					for _, k := range myResall { //所有设置了权限的都不能看
						// beego.Info(k)
						if strconv.FormatInt(v.Id, 10) == path.Base(k[1]) && !isadmin {
							docxarr[0].Permission = "4"
							// docxarr[0].Permission = "3"
						}
					}

					for _, k := range myRes { //如果与登录用户对应上，则赋予权限
						// beego.Info(k)
						if strconv.FormatInt(v.Id, 10) == path.Base(k[1]) {
							docxarr[0].Permission = k[2]
						}
					}
					roles, err := e.GetRolesForUser(useridstring) //取出用户的所有角色
					if err != nil {
						logs.Error(err)
					}
					for _, w1 := range roles { //2018.4.30修改这个bug，这里原先w改为w1
						roleRes = e.GetPermissionsForUser(w1) //取出角色的所有权限，改为w1
						for _, k := range roleRes {
							// beego.Info(k)
							if strconv.FormatInt(v.Id, 10) == path.Base(k[1]) {
								// docxarr[0].Permission = k[2]
								int1, err := strconv.Atoi(k[2])
								if err != nil {
									logs.Error(err)
								}
								int2, err := strconv.Atoi(docxarr[0].Permission)
								if err != nil {
									logs.Error(err)
								}
								if int1 < int2 {
									docxarr[0].Permission = k[2] //按最小值权限
								}
								//补充everyone
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
						// docxarr[0].Permission = "3"
						docxarr[0].Permission = "4"
						// beego.Info(strconv.FormatInt(v.Id, 10))
						// beego.Info(path.Base(k[1]))
					}
					//如果有everyone权限，则按everyone的
				}
			}

			docxarr[0].Id = v.Id
			docxarr[0].Title = v.FileName
			if path.Ext(v.FileName) == ".docx" || path.Ext(v.FileName) == ".DOCX" || path.Ext(v.FileName) == ".doc" || path.Ext(v.FileName) == ".DOC" {
				docxarr[0].Suffix = "docx"
			} else if path.Ext(v.FileName) == ".wps" || path.Ext(v.FileName) == ".WPS" {
				docxarr[0].Suffix = "docx"
			} else if path.Ext(v.FileName) == ".XLSX" || path.Ext(v.FileName) == ".xlsx" || path.Ext(v.FileName) == ".XLS" || path.Ext(v.FileName) == ".xls" {
				docxarr[0].Suffix = "xlsx"
				// xlsxarr := make([]XlsxLink, 1)
				// xlsxarr[0].Id = v.Id
				// xlsxarr[0].Title = v.FileName
				// Xlsxslice = append(Xlsxslice, xlsxarr...)
			} else if path.Ext(v.FileName) == ".ET" || path.Ext(v.FileName) == ".et" {
				docxarr[0].Suffix = "xlsx"
			} else if path.Ext(v.FileName) == ".pptx" || path.Ext(v.FileName) == ".PPTX" || path.Ext(v.FileName) == ".ppt" || path.Ext(v.FileName) == ".PPT" {
				docxarr[0].Suffix = "pptx"
				// pptxarr := make([]PptxLink, 1)
				// pptxarr[0].Id = v.Id
				// pptxarr[0].Title = v.FileName
				// Pptxslice = append(Pptxslice, pptxarr...)
			} else if path.Ext(v.FileName) == ".DPS" || path.Ext(v.FileName) == ".dps" {
				docxarr[0].Suffix = "pptx"
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
		//无权限的不显示
		//如果permission！=4或者是管理，则赋值给link
		if linkarr[0].Docxlink[0].Permission != "4" || isadmin {
			link = append(link, linkarr...)
		}
	}
	c.Data["json"] = link //products
	c.ServeJSON()
}

// WPS文字：wps,wpt,doc,dot,rtf
// WPS演示：dps,dpt,ppt,pot,pps
// WPS表格：et,ett,xls,xlt
//取得changesurl
// func (c *OnlyController) ChangesUrl() {
// 	version := c.GetString("version")

// 	versionint, err := strconv.Atoi(version)
// 	if err != nil {
// 		logs.Error(err)
// 	}
// 	attachmentid := c.GetString("attachmentid")
// 	idNum, err := strconv.ParseInt(attachmentid, 10, 64)
// 	if err != nil {
// 		logs.Error(err)
// 	}
// 	changesurl, err := models.GetOnlyChangesUrl(idNum, versionint)
// 	if err != nil {
// 		logs.Error(err)
// 	}
// 	beego.Info(changesurl.ChangesUrl)
// 	c.Data["json"] = changesurl.ChangesUrl
// 	c.ServeJSON()
// }

// @Title get onlyoffce
// @Description get onlyoffice
// @Param id query string true "The id of office"
// @Success 200 {object} models.Ollyoffice
// @Failure 400 Invalid page supplied
// @Failure 404 office not found
// @router /onlyoffice [get]
// 协作页面的显示——路由是/onlyoffce/onlyoffice/123
// 补充权限判断
// 补充token
func (c *OnlyController) OnlyOffice() {
	id := c.Ctx.Input.Param(":id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	//根据附件id取得附件的prodid，路径
	onlyattachment, err := models.GetOnlyAttachbyId(idNum)
	if err != nil {
		logs.Error(err)
	}

	//docid——uid——me
	doc, err := models.Getdocbyid(onlyattachment.DocId)
	if err != nil {
		logs.Error(err)
	}

	var useridstring, Permission string
	var myRes, roleRes [][]string

	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	// c.Data["Uid"] = uid
	useridstring = strconv.FormatInt(uid, 10)
	c.Data["Uid"] = useridstring
	//增加admin，everyone，isme

	var usersessionid string //客户端sesssionid
	if islogin {
		usersessionid = c.Ctx.Input.Cookie("hotqinsessionid")
		//服务端sessionid怎么取出
		// v := c.GetSession("uname")
		// beego.Info(v.(string))
	}
	// v := c.GetSession("uname")
	// var role, userrole int
	myResall := e.GetPermissionsForUser("") //取出所有设置了权限的数据
	if uid != 0 {                           //无论是登录还是ip查出了用户id
		// uname = v.(string)
		// c.Data["Uname"] = v.(string)
		// user, err := models.GetUserByUsername(uname)
		// if err != nil {
		// 	logs.Error(err)
		// }
		// useridstring = strconv.FormatInt(user.Id, 10)
		myRes = e.GetPermissionsForUser(useridstring)
		if doc.Uid == uid || isadmin { //isme or isadmin
			Permission = "1"
		} else { //如果是登录用户，则设置了权限的文档根据权限查看
			Permission = "1"
			for _, k := range myResall {
				if strconv.FormatInt(onlyattachment.Id, 10) == path.Base(k[1]) {
					Permission = "4"
					// Permission = "3"
				}
			}
			for _, k := range myRes {
				if strconv.FormatInt(onlyattachment.Id, 10) == path.Base(k[1]) {
					Permission = k[2]
				}
			}

			roles, err := e.GetRolesForUser(useridstring) //取出用户的所有角色
			if err != nil {
				logs.Error(err)
			}
			for _, w1 := range roles { //2018.4.30修改这个bug，这里原先w改为w1
				roleRes = e.GetPermissionsForUser(w1) //取出角色的所有权限，改为w1
				for _, k := range roleRes {
					// beego.Info(k)
					if id == path.Base(k[1]) {
						// docxarr[0].Permission = k[2]
						int1, err := strconv.Atoi(k[2])
						if err != nil {
							logs.Error(err)
						}
						int2, err := strconv.Atoi(Permission)
						if err != nil {
							logs.Error(err)
						}
						if int1 < int2 {
							Permission = k[2] //按最小值权限
						}
						//补充everyone权限，如果登录用户权限大于everyone，则用小的
					}
				}
			}
		}
		// c.Data["Uid"] = user.Id
		// userrole = user.Role
	} else { //如果用户没登录，则设置了权限的文档不能看
		Permission = "1"
		for _, k := range myResall { //所有设置了权限的不能看
			if strconv.FormatInt(onlyattachment.Id, 10) == path.Base(k[1]) {
				// Permission = "3"
				Permission = "4"
			}
			//如果设置了everyone用户权限，则按everyone的权限
		}
		// c.Data["Uname"] = c.Ctx.Input.IP()
		// c.Data["Uid"] = 0
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
		c.Data["Comment"] = true
		c.Data["Download"] = true
		c.Data["Print"] = true
	} else if Permission == "2" {
		c.Data["Mode"] = "edit"
		c.Data["Edit"] = false
		c.Data["Review"] = true
		c.Data["Comment"] = true
		c.Data["Download"] = false
		c.Data["Print"] = false
	} else if Permission == "3" {
		c.Data["Mode"] = "view"
		c.Data["Edit"] = false
		c.Data["Review"] = false
		c.Data["Comment"] = false
		c.Data["Download"] = false
		c.Data["Print"] = false
	} else if Permission == "4" {
		// route := c.Ctx.Request.URL.String()
		// c.Redirect("/roleerr?url="+route, 302)
		c.TplName = "onlyoffice/login.tpl"
		return
	}

	c.Data["Doc"] = onlyattachment
	c.Data["Sessionid"] = usersessionid
	c.Data["attachid"] = idNum
	c.Data["Key"] = strconv.FormatInt(onlyattachment.Updated.UnixNano(), 10)

	//构造[]history
	history, err := models.GetOnlyHistory(onlyattachment.Id)
	if err != nil {
		logs.Error(err)
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
		aa[0].ServerVersion = v.ServerVersion
		aa[0].Version = v.Version
		aa[0].FileUrl = v.FileUrl
		aa[0].ChangesUrl = v.ChangesUrl
		//取得changes
		changes, err := models.GetOnlyChanges(v.HistoryKey)
		if err != nil {
			logs.Error(err)
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
		logs.Error(err)
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
		c.Data["documentType"] = "text" //word
	} else if path.Ext(onlyattachment.FileName) == ".wps" || path.Ext(onlyattachment.FileName) == ".WPS" {
		c.Data["fileType"] = "docx"
		c.Data["documentType"] = "text"
	} else if path.Ext(onlyattachment.FileName) == ".XLSX" || path.Ext(onlyattachment.FileName) == ".xlsx" {
		c.Data["fileType"] = "xlsx"
		c.Data["documentType"] = "spreadsheet" //cell
	} else if path.Ext(onlyattachment.FileName) == ".ET" || path.Ext(onlyattachment.FileName) == ".et" {
		c.Data["fileType"] = "xlsx"
		c.Data["documentType"] = "spreadsheet"
	} else if path.Ext(onlyattachment.FileName) == ".pptx" || path.Ext(onlyattachment.FileName) == ".PPTX" {
		c.Data["fileType"] = "pptx"
		c.Data["documentType"] = "presentation" //slide
	} else if path.Ext(onlyattachment.FileName) == ".dps" || path.Ext(onlyattachment.FileName) == ".DPS" {
		c.Data["fileType"] = "pptx"
		c.Data["documentType"] = "presentation"
	} else if path.Ext(onlyattachment.FileName) == ".doc" || path.Ext(onlyattachment.FileName) == ".DOC" {
		c.Data["fileType"] = "doc"
		c.Data["documentType"] = "text"
	} else if path.Ext(onlyattachment.FileName) == ".txt" || path.Ext(onlyattachment.FileName) == ".TXT" {
		c.Data["fileType"] = "txt"
		c.Data["documentType"] = "text" //word
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
		c.Data["documentType"] = "text" //word
		c.Data["Mode"] = "view"
	}

	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
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
	onlyofficeapi_url, err := web.AppConfig.String("onlyofficeapi_url")
	if err != nil {
		logs.Error(err)
	}
	c.Data["Onlyofficeapi_url"] = onlyofficeapi_url
	engineercmsapi_url, err := web.AppConfig.String("engineercmsapi_url")
	if err != nil {
		logs.Error(err)
	}
	c.Data["Engineercmsapi_url"] = engineercmsapi_url

	c.TplName = "onlyoffice/onlyoffice.tpl"
}

// 协作页面的保存和回调
// 关闭浏览器标签后获取最新文档保存到文件夹
func (c *OnlyController) UrltoCallback() {
	var actionuserid int64
	// pk1 := c.Ctx.Input.RequestBody
	id := c.GetString("id")
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}

	//根据附件id取得附件的prodid，路径
	onlyattachment, err := models.GetOnlyAttachbyId(idNum)
	if err != nil {
		logs.Error(err)
	}

	var callback Callback
	json.Unmarshal(c.Ctx.Input.RequestBody, &callback)
	// beego.Info(string(c.Ctx.Input.RequestBody))
	// beego.Info(callback.Status)
	// beego.Info(callback.Forcesavetype)
	if callback.Status == 1 || callback.Status == 4 {
		//•	1 - document is being edited,
		//•	4 - document is closed with no changes,
		// beego.Info(callback.Users)
		OnlyUsers = callback.Users
		c.Data["json"] = map[string]interface{}{"error": 0}
		c.ServeJSON()
	} else if callback.Status == 2 && callback.Notmodified == false {
		//•	2 - document is ready for saving
		resp, err := http.Get(callback.Url) //Changesurl
		if err != nil {
			logs.Error(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logs.Error(err)
		}
		defer resp.Body.Close()
		if err != nil {
			logs.Error(err)
		}
		//1.
		//2.将document server中的文件下载下来存入，名称就是编号v1
		//3.再将变化文件changesurl文件存下来，名称为changesv1.zip
		//1.改名保存
		FileSuffix := path.Ext(onlyattachment.FileName) //只留下后缀名
		filenameOnly := strings.TrimSuffix(onlyattachment.FileName, FileSuffix)

		var first int
		//写入历史版本
		historyversion, err := models.GetOnlyHistoryVersion(onlyattachment.Id)
		if err != nil {
			logs.Error(err)
		}
		for _, v := range historyversion {
			if first < v.Version {
				first = v.Version
			}
		}
		//去掉版本号：要么用正则，要么用history中的版本号进行替换？风险有点大
		filenameOnly = strings.Replace(filenameOnly, "v"+strconv.Itoa(first), "", -1)
		vnumber := strconv.Itoa(first + 1)
		// file := "./attachment/onlyoffice/" + onlyattachment.FileName                          //源文件路径
		// err = os.Rename(file, "./attachment/onlyoffice/"+filenameOnly+"v"+vnumber+FileSuffix) //重命名 C:\\log\\2013.log 文件为install.txt
		// if err != nil {
		// 	logs.Error(err)
		// }

		// f, err := os.OpenFile("./attachment/onlyoffice/"+onlyattachment.FileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		f, err := os.Create("./attachment/onlyoffice/" + filenameOnly + "v" + vnumber + FileSuffix)
		if err != nil {
			logs.Error(err)
		}
		defer f.Close()
		_, err = f.Write(body) //这里直接用resp.Body如何？
		// _, err = f.WriteString(str)
		// _, err = io.Copy(body, f)
		if err != nil {
			logs.Error(err)
		} else {
			//更新文档更新时间
			err = models.UpdateDocTime(onlyattachment.DocId)
			if err != nil {
				logs.Error(err)
			}
		}

		//更新附件的时间和changesurl
		err = models.UpdateOnlyAttachment(idNum, filenameOnly+"v"+vnumber+FileSuffix)
		if err != nil {
			logs.Error(err)
		}

		//写入历史版本数据
		array := strings.Split(callback.Changesurl, "&")
		Expires1 := strings.Split(array[1], "=")
		Expires := Expires1[1]
		Expirestime, err := strconv.ParseInt(Expires, 10, 64)
		if err != nil {
			logs.Error(err)
		}
		//获取本地location
		// toBeCharge := "2015-01-01 00:00:00"
		//待转化为时间戳的字符串 注意 这里的小时和分钟还有秒必须写 因为是跟着模板走的 修改模板的话也可以不写
		// timeLayout := "2006-01-02T15:04:05.999Z"
		//转化所需模板
		// loc, _ := time.LoadLocation("Local") //重要：获取时区
		// theTime, _ := time.ParseInLocation(timeLayout, toBeCharge, loc) //使用模板在对应时区转化为time.time类型
		// sr := theTime.Unix()
		//转化为时间戳 类型是int64
		//打印输出时间戳 1420041600

		//changes文档保存存下来
		respchanges, err := http.Get(callback.Changesurl) //Changesurl
		if err != nil {
			logs.Error(err)
		}
		bodychanges, err := ioutil.ReadAll(respchanges.Body)
		if err != nil {
			logs.Error(err)
		}
		defer respchanges.Body.Close()
		if err != nil {
			logs.Error(err)
		}
		//建立目录，并返回作为父级目录
		err = os.MkdirAll("./attachment/onlyoffice/changes/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
		if err != nil {
			logs.Error(err)
		}
		fchanges, err := os.Create("./attachment/onlyoffice/changes/" + filenameOnly + "v" + vnumber + "changes.zip")
		if err != nil {
			logs.Error(err)
		}
		defer fchanges.Close()
		_, err = fchanges.Write(bodychanges)
		if err != nil {
			logs.Error(err)
		}
		//写入历史数据库
		nowdocurl := "/attachment/onlyoffice/" + filenameOnly + "v" + vnumber + FileSuffix
		changeszipurl := "/attachment/onlyoffice/changes/" + filenameOnly + "v" + vnumber + "changes.zip"
		//时间戳转日期
		dataTimeStr := time.Unix(Expirestime, 0) //.Format(timeLayout) //设置时间戳 使用模板格式化为日期字符串
		// t, _ := time.Parse(timeLayout, callback.Lastsave)
		// beego.Info(callback.Lastsave)
		// beego.Info(callback.Users[0])
		if len(callback.Actions) == 0 {
			actionuserid = 0
		} else {
			actionuserid = callback.Actions[0].Userid
		}
		// _, err1, err2 := models.AddOnlyHistory(onlyattachment.Id, actionuserid, callback.History.ServerVersion, first+1, callback.Key, callback.Url, callback.Changesurl, dataTimeStr, callback.Lastsave)
		_, err1, err2 := models.AddOnlyHistory(onlyattachment.Id, actionuserid, callback.History.ServerVersion, first+1, callback.Key, nowdocurl, changeszipurl, dataTimeStr, callback.Lastsave)
		if err1 != nil {
			logs.Error(err1)
		}
		if err2 != nil {
			logs.Error(err2)
		}
		//写入changes数据库
		for _, v := range callback.History.Changes {
			_, err1, err2 = models.AddOnlyChanges(callback.Key, v.User.Id, v.User.Name, v.Created)
			if err1 != nil {
				logs.Error(err1)
			}
			if err2 != nil {
				logs.Error(err2)
			}
		}

		c.Data["json"] = map[string]interface{}{"error": 0}
		c.ServeJSON()
	} else if callback.Status == 6 && callback.Forcesavetype == 1 {
		//•	6 - document is being edited, but the current document state is saved,
		// 0 - the force saving request is performed to the command service,
		// 1 - the force saving request is performed each time the saving is done (e.g. the Save button is clicked), which is only available when the forcesave option is set to true.
		// 2 - the force saving request is performed by timer with the settings from the server config.
		resp, err := http.Get(callback.Url)
		if err != nil {
			logs.Error(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logs.Error(err)
		}
		defer resp.Body.Close()
		if err != nil {
			logs.Error(err)
		}
		// f, err := os.OpenFile("./attachment/onlyoffice/"+onlyattachment.FileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
		//强制保存——不好用，前端不要设置成强制保存！！
		f, err := os.Create("./attachment/onlyoffice/" + onlyattachment.FileName)
		if err != nil {
			logs.Error(err)
		}
		defer f.Close()
		_, err = f.Write(body) //这里直接用resp.Body如何？
		if err != nil {
			logs.Error(err)
		}
		// else {
		//更新文档更新时间_不能更新时间，会造成key失效
		// err = models.UpdateDocTime(onlyattachment.DocId)
		// if err != nil {
		// 	logs.Error(err)
		// }
		// }
		c.Data["json"] = map[string]interface{}{"error": 0}
		c.ServeJSON()
	} else if callback.Status == 3 || callback.Status == 7 {
		//•	3 - document saving error has occurred.
		//•	7 - error has occurred while force saving the document.
		//更新附件的时间和changesurl
		//不更新可以吗？此时有人没有关闭浏览器，有人重新打开文档，
		//用新的key在服务器上编辑文档了！！！
		err = models.UpdateOnlyAttachment(idNum, onlyattachment.FileName)
		if err != nil {
			logs.Error(err)
		}
		c.Data["json"] = map[string]interface{}{"error": 0}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"error": 0}
		c.ServeJSON()
	}
}

// 批量添加一对一模式
// 要避免同名覆盖的严重bug！！！！
func (c *OnlyController) AddOnlyAttachment() {
	//取得客户端用户名
	// v := c.GetSession("uname")
	// var user models.User
	// var err error
	// if v != nil {
	// 	uname := v.(string)
	// 	user, err = models.GetUserByUsername(uname)
	// 	if err != nil {
	// 		logs.Error(err)
	// 	}
	// }
	_, _, uid, _, _ := checkprodRole(c.Ctx)
	var filepath, DiskDirectory, Url string
	err := os.MkdirAll("./attachment/onlyoffice/", 0777) //..代表本当前exe文件目录的上级，.表示当前目录，没有.表示盘的根目录
	if err != nil {
		logs.Error(err)
	}
	DiskDirectory = "./attachment/onlyoffice/"
	Url = "/attachment/onlyoffice/"

	//获取上传的文件
	_, h, err := c.GetFile("file")
	if err != nil {
		logs.Error(err)
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
		prodlabel := c.GetString("prodlabel")
		prodprincipal := c.GetString("prodprincipal")
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
		inputdate := c.GetString("proddate")
		// beego.Info(inputdate)
		var t1, end time.Time
		// var convdate1, convdate2 string
		const lll = "2006-01-02"
		if len(inputdate) > 9 { //如果是datepick获取的时间，则不用加8小时
			t1, err = time.Parse(lll, inputdate) //这里t1要是用t1:=就不是前面那个t1了
			if err != nil {
				logs.Error(err)
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
				logs.Error(err)
			}
			end = date
		}
		// beego.Info(end)
		prodId, err := models.AddDoc(code, title, prodlabel, prodprincipal, end, uid)
		if err != nil {
			logs.Error(err)
		}
		//改名，替换文件名中的#和斜杠
		title = strings.Replace(title, "#", "号", -1)
		title = strings.Replace(title, "/", "-", -1)

		// FileSuffix := path.Ext(h.Filename)
		// filepath = DiskDirectory + "/" + code + title + FileSuffix
		// attachmentname := code + title + FileSuffix
		filepath = DiskDirectory + "/" + h.Filename
		//把成果id作为附件的parentid，把附件的名称等信息存入附件数据库
		//如果附件名称相同，则不能上传，数据库添加
		attachmentname := h.Filename

		_, _, err2 := models.AddOnlyAttachment(attachmentname, 0, 0, prodId)
		if err2 != nil {
			logs.Error(err2)
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
						logs.Error(err)
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

// 协作页面下载的文档，采用绝对路径型式
func (c *OnlyController) DownloadDoc() {
	// v := c.GetSession("uname")
	// if v != nil {
	// uname := v.(string)
	// beego.Info(uname)
	// 	c.Data["Uname"] = v.(string)
	// 	user, err := models.GetUserByUsername(uname)
	// 	if err != nil {
	// 		logs.Error(err)
	// 	}
	// 	useridstring = strconv.FormatInt(user.Id, 10)
	// }
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	//4.取得客户端用户名
	// var uname, useridstring string
	// v := c.GetSession("uname")
	// if v != nil {
	// 	uname = v.(string)
	// 	c.Data["Uname"] = v.(string)
	// 	user, err := models.GetUserByUsername(uname)
	// 	if err != nil {
	// 		logs.Error(err)
	// 	}
	// 	useridstring = strconv.FormatInt(user.Id, 10)
	// }
	_, _, uid, isadmin, _ := checkprodRole(c.Ctx)
	// docid := c.Ctx.Input.Param(":id")
	// beego.Info(docid)
	//pid转成64为
	// idNum, err := strconv.ParseInt(docid, 10, 64)
	// if err != nil {
	// 	logs.Error(err)
	// }
	//根据成果id取得所有附件
	// attachments, err := models.GetOnlyAttachments(idNum)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// filePath := "attachment/onlyoffice/" + attachments[0].FileName

	filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:]) //attachment/onlyoffice/id
	if err != nil {
		logs.Error(err)
	}
	//由附件名取得附件id
	var downloadfile models.OnlyAttachment
	if strings.Contains(filePath, "?hotqinsessionid=") {
		filePathtemp := strings.Split(filePath, "?")
		filePath = filePathtemp[0]
		// beego.Info(filePath)
	}
	filename := filepath.Base(filePath)

	fileext := path.Ext(filename)
	matched, err := regexp.MatchString("\\.*[m|M][c|C][d|D]", fileext)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(matched)
	if matched {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "不能下载mcd文件!"}
		c.ServeJSON()
		return
	}
	matched, err = regexp.MatchString("\\.*[f|F][c|C][s|S][t|T][d|D]", fileext)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(matched)
	if matched {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "不能下载fcstd文件!"}
		c.ServeJSON()
		return
	}

	downloadfile, err = models.GetOnlyAttachbyName(filename)
	if err != nil {
		logs.Error(err)
	}

	//1.管理员或者没有设置权限的文档直接可以下载。
	police := e.GetFilteredPolicy(1, "/onlyoffice/"+strconv.FormatInt(downloadfile.Id, 10))
	// beego.Info(police)
	if isadmin || len(police) == 0 {
		// c.Ctx.Output.Download(filePath) //这个能保证下载文件名称正确
		http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)
		return
	}

	//2.取得用户权限
	police = e.GetFilteredPolicy(0, strconv.FormatInt(uid, 10), "/onlyoffice/"+strconv.FormatInt(downloadfile.Id, 10))
	// beego.Info(police)
	for _, v2 := range police {
		// beego.Info(v2)
		v2int, err := strconv.ParseInt(v2[2], 10, 64)
		if err != nil {
			logs.Error(err)
		}
		if v2int <= 3 {
			// canidown = true
			// c.Ctx.Output.Download(filePath) //这个能保证下载文件名称正确
			http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)
			return
		}
	}

	//3.取得用户角色——取得角色的权限
	userroles, err := e.GetRolesForUser(strconv.FormatInt(uid, 10))
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(userroles)
	// userrole := make([]Userrole, 0)
	// var canidown bool
	for _, v1 := range userroles {
		police := e.GetFilteredPolicy(0, v1, "/onlyoffice/"+strconv.FormatInt(downloadfile.Id, 10))
		// beego.Info(police)
		for _, v2 := range police {
			// beego.Info(v2)
			v2int, err := strconv.ParseInt(v2[2], 10, 64)
			if err != nil {
				logs.Error(err)
			}
			if v2int <= 3 {
				// canidown = true
				// c.Ctx.Output.Download(filePath) //这个能保证下载文件名称正确
				http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)
				return
			}
		}
	}
	//1.url处理中文字符路径，[1:]截掉路径前面的/斜杠
	// filePath1 := path.Base(c.Ctx.Request.RequestURI)
	// beego.Info(filePath1)
	// c.Data["json"] = "无权限下载"//这个导致文档被覆盖，不可取
	// c.ServeJSON()
}

// @Title post download onlydoc
// @Description post download onlydoc by id
// @Param id query string true "The id of onlydoc"
// @Param url query string true "The url of onlyofficeserver"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /downloadonlydoc [post]
// 文档管理页面下载最新的文档——先检查是否有人打开文档
func (c *OnlyController) DownloadOnlyDoc() {
	id := c.GetString("id")
	// key := c.GetString("key")
	// beego.Info(id)
	url := c.GetString("url")
	// array := strings.Split(ids, ",")
	// for _, v := range array {
	//pid转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"errNo": 0, "info": "ERR", "data": "字符转int64错误", "msg": "字符转int64错误"}
		c.ServeJSON()
		return
	}
	//根据附件id取得附件的prodid，路径
	onlyattachment, err := models.GetOnlyAttachbyId(idNum)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"errNo": 1, "info": "ERR", "data": "查询onlyattachment错误", "msg": "查询onlyattachment错误"}
		c.ServeJSON()
		return
	}
	key := strconv.FormatInt(onlyattachment.Updated.UnixNano(), 10)

	// For the interaction with the document command service the POST requests are used.
	// The request parameters are entered in JSON format in the request body.
	requestUrl := url + "/coauthoring/CommandService.ashx"
	b, err := json.Marshal(map[string]string{
		"c":        "forcesave",
		"key":      key,
		"userdata": "sample userdata",
	})
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"errNo": 2, "info": "ERR", "data": "json转换错误", "msg": "json转换错误"}
		c.ServeJSON()
		return
	}
	resp, err := http.Post(requestUrl, "application/x-www-form-urlencoded", bytes.NewBuffer(b)) //注意，这里是post
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"errNo": 3, "info": "ERR", "data": "post请求错误", "msg": "post请求错误"}
		c.ServeJSON()
		return
	}
	// beego.Info(resp)
	// beego.Info(resp.Body)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"errNo": 4, "info": "ERR", "data": "请求返回错误", "msg": "请求返回错误"}
		c.ServeJSON()
		return
	}
	var msgdata map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&msgdata)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"errNo": 5, "info": "ERR", "data": "解码错误", "msg": "解码错误"}
		c.ServeJSON()
		return
	}
	// beego.Info(msgdata)
	if msgdata["error"].(float64) == 0 {
		c.Data["json"] = map[string]interface{}{"errNo": 6, "info": "SUCCESS", "data": "forcesave成功", "msg": "forcesave成功"}
		c.ServeJSON()
		return
	} else if msgdata["error"].(float64) == 4 {
		c.Data["json"] = map[string]interface{}{"errNo": 7, "info": "ERR", "data": "文档无修改", "msg": "文档无修改"}
		c.ServeJSON()
		return
	}
}

// 文档管理页面下载文档2
func (c *OnlyController) Download() {
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	//4.取得客户端用户名
	_, _, uid, isadmin, _ := checkprodRole(c.Ctx)
	docid := c.Ctx.Input.Param(":id")
	// beego.Info(docid)
	//pid转成64为
	idNum, err := strconv.ParseInt(docid, 10, 64)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"errNo": 0, "info": "ERR", "data": "字符转int64错误", "msg": "字符转int64错误"}
		c.ServeJSON()
		return
	}
	//根据成果id取得所有附件
	attachments, err := models.GetOnlyAttachments(idNum)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"errNo": 1, "info": "ERR", "data": "查询onlyattachment错误", "msg": "查询onlyattachment错误"}
		c.ServeJSON()
		return
	}

	fileext := path.Ext(attachments[0].FileName)
	matched, err := regexp.MatchString("\\.*[m|M][c|C][d|D]", fileext)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(matched)
	if matched {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "不能下载mcd文件!"}
		c.ServeJSON()
		return
	}

	filePath := "attachment/onlyoffice/" + attachments[0].FileName
	//管理员或者没有设置权限的文档直接可以下载。
	police := e.GetFilteredPolicy(1, "/onlyoffice/"+docid)
	if isadmin || len(police) == 0 {
		c.Ctx.Output.Download(filePath) //这个能保证下载文件名称正确
		return
	}
	// beego.Info(police[0][0])

	userroles, err := e.GetRolesForUser(strconv.FormatInt(uid, 10))
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(userroles)
	// userrole := make([]Userrole, 0)
	// var canidown bool

	for _, v1 := range userroles {
		police := e.GetFilteredPolicy(0, v1, "/onlyoffice/"+docid)
		// beego.Info(police)
		for _, v2 := range police {
			// beego.Info(v2)
			v2int, err := strconv.ParseInt(v2[2], 10, 64)
			if err != nil {
				logs.Error(err)
				c.Data["json"] = map[string]interface{}{"errNo": 2, "info": "ERR", "data": "字符转int64错误", "msg": "字符转int64错误"}
				c.ServeJSON()
				return
			}
			if v2int <= 3 {
				// canidown = true
				c.Ctx.Output.Download(filePath) //这个能保证下载文件名称正确
				return
			}
		}
	}
	c.Data["json"] = map[string]interface{}{"errNo": 3, "info": "ERR", "data": "无权下载", "msg": "无权下载"}
	c.ServeJSON()
	// var paths []beegoormadapter.CasbinRule
	// o := orm.NewOrm()
	// qs := o.QueryTable("casbin_rule")
	// _, err := qs.Filter("PType", "p").Filter("v0", "role_"+roleid).Filter("v1", "/onlyoffice"+docid).All(&paths)
	// if err != nil {
	// 	logs.Error(err)
	// }
	// beego.Info(paths)

	// var projids []string
	// for _, v1 := range paths {
	// 	beego.Info(v1.V2)
	// projid := strings.Replace(v1.V1, "/*", "", -1)
	// projids = append(projids, path.Base(projid))
	// }

	// if e.Enforce(useridstring, "attachment/onlyoffice/" +docid, c.Ctx.Request.Method, fileext) || isadmin {

	// beego.Info(filePath)
	//1.url处理中文字符路径，[1:]截掉路径前面的/斜杠
	// filePath := path.Base(ctx.Request.RequestURI)
	// filePath, err := url.QueryUnescape(c.Ctx.Request.RequestURI[1:]) //  attachment/SL2016测试添加成果/A/FB/1/Your First Meteor Application.pdf
	// if err != nil {
	// 	logs.Error(err)
	// }
	// fileext := path.Ext(filePath)
	//根据路由path.Dir——再转成数组strings.Split——查出项目id——加上名称——查出下级id
	// beego.Info(path.Dir(filePath))
	// filepath1 := path.Dir(filePath)
	// array := strings.Split(filepath1, "/")
	// beego.Info(strings.Split(filepath1, "/"))
	// http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, filePath)//这个下载文件名不对

	// }
	// *******解密****************
	// aesKey := "0123456789123456"
	// ciphertext, err := ioutil.ReadFile("attachment/onlyoffice/MyXLSXFile1.xlsx")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// // Key
	// key := []byte(aesKey)

	// // Create the AES cipher
	// block, err := aes.NewCipher(key)
	// if err != nil {
	// 	panic(err)
	// }

	// // Before even testing the decryption,
	// // if the text is too small, then it is incorrect
	// if len(ciphertext) < aes.BlockSize {
	// 	panic("Text is too short")
	// }

	// // Get the 16 byte IV
	// iv := ciphertext[:aes.BlockSize]

	// // Remove the IV from the ciphertext
	// ciphertext = ciphertext[aes.BlockSize:]

	// // Return a decrypted stream
	// stream := cipher.NewCFBDecrypter(block, iv)
	// // Decrypt bytes from ciphertext
	// stream.XORKeyStream(ciphertext, ciphertext)
	// // create a new file for saving the encrypted data.
	// // f, err := os.Create("11.ts") //生成一个新的文件
	// // if err != nil {
	// // 	panic(err.Error())
	// // }
	// // defer f.Close()
	// // _, err = io.Copy(f, bytes.NewReader(ciphertext))
	// // if err != nil {
	// // 	logs.Error(err)
	// // }
	// io.Copy(c.Ctx.ResponseWriter, bytes.NewReader(ciphertext))
}

type downloadzip struct {
	Ids string `json:"ids"`
}

// @Title download zip
// @Description download zip
// @Param ids query string true "The ids of onlydoc"
// @Success 200 {object} models.Onlyoffice
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /downloadzip [post]
func (c *OnlyController) DownloadZip() {
	// ids1 := c.GetString("ids")
	// logs.Info(ids1)
	content := c.Ctx.Input.RequestBody
	var ob downloadzip
	err := json.Unmarshal(content, &ob)
	// logs.Info(ob.Ids)
	array := strings.Split(ob.Ids, ",")
	// logs.Info(array)
	//建立文件夹
	_, err = os.Stat("./temp/onlyoffice")
	if err != nil {
		if os.IsNotExist(err) {
			// return false
			err := os.MkdirAll("./temp/onlyoffice", 0777)
			if err != nil {
				logs.Error(err)
			}
		}
	}

	for _, v := range array {
		// logs.Info(v)
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			logs.Error(err)
		}
		//根据成果id取得所有附件
		attachments, err := models.GetOnlyAttachments(idNum)
		if err != nil {
			logs.Error(err)
		}
		for _, w := range attachments {
			//取得附件的成果id
			attach, err := models.GetOnlyAttachbyId(w.Id)
			if err != nil {
				logs.Error(err)
			}
			path := "./attachment/onlyoffice/" + attach.FileName
			_, err = CopyFile("./temp/onlyoffice/"+attach.FileName, path)
			//targetfile,sourcefile
			if err != nil {
				logs.Error(err)
			}
		}
	}

	//压缩
	ziptil.Zip("./temp/onlyoffice", "./temp/onlyoffice.zip")
	// ajax无法下载zip文件
	c.Ctx.Output.Download("./temp/onlyoffice.zip")
	err = os.Remove("./temp/onlyoffice.zip")
	if err != nil {
		logs.Error(err)
	}
	err = RemoveContents("./temp/onlyoffice/")
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
	// c.Data["json"] = map[string]interface{}{"code": "OK", "msg": "", "data": ids}
	// c.ServeJSON()
}

// 编辑成果信息
func (c *OnlyController) UpdateDoc() {
	id := c.GetString("pid")
	code := c.GetString("code")
	title := c.GetString("title")
	label := c.GetString("label")
	principal := c.GetString("principal")
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	inputdate := c.GetString("proddate")
	var t1, end time.Time
	const lll = "2006-01-02"
	if len(inputdate) > 9 { //如果是datepick获取的时间，则不用加8小时
		t1, err = time.Parse(lll, inputdate) //这里t1要是用t1:=就不是前面那个t1了
		if err != nil {
			logs.Error(err)
		}
		end = t1
		// t1 = printtime.Add(+time.Duration(hours) * time.Hour)
	} else { //如果取系统时间，则需要加8小时
		date := time.Now()
		convdate := date.Format(lll)
		date, err = time.Parse(lll, convdate)
		if err != nil {
			logs.Error(err)
		}
		end = date
	}
	//根据id添加成果
	err = models.UpdateDoc(idNum, code, title, label, principal, end)
	if err != nil {
		logs.Error(err)
	}
	c.Data["json"] = "ok"
	c.ServeJSON()
}

// 删除成果，包含成果里的附件。删除附件用attachment中的
func (c *OnlyController) DeleteDoc() {
	ids := c.GetString("ids")
	array := strings.Split(ids, ",")
	for _, v := range array {
		//id转成64位
		idNum, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			logs.Error(err)
		}
		//循环删除成果
		//根据成果id取得所有附件
		attachments, err := models.GetOnlyAttachments(idNum)
		if err != nil {
			logs.Error(err)
		}
		for _, w := range attachments {
			//取得附件的成果id
			attach, err := models.GetOnlyAttachbyId(w.Id)
			if err != nil {
				logs.Error(err)
			}
			path := "./attachment/onlyoffice/" + attach.FileName
			//删除附件
			err = os.Remove(path)
			if err != nil {
				logs.Error(err)
			}
			//删除附件数据表
			err = models.DeleteOnlyAttachment(w.Id)
			if err != nil {
				logs.Error(err)
			}
		}
		err = models.DeleteDoc(idNum) //删除成果数据表
		if err != nil {
			logs.Error(err)
		} else {
			c.Data["json"] = "ok"
			c.ServeJSON()
		}
	}
}

// @Title post Addpermission...
// @Description post Addpermission..
// @Param ids query string true "The id of document"
// @Success 200 {object} models.GetOnlyoffice
// @Failure 400 Invalid page supplied
// @Failure 404 data not found
// @router /addpermission [post]
// onlyoffice权限管理
// 添加用户和角色的权限
// 先删除这个文档id下所有permission，再添加新的。
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
		logs.Error(err)
	}
	//根据成果id取得所有附件——这里只取第一个
	attachments, err := models.GetOnlyAttachments(idNum)
	if err != nil {
		logs.Error(err)
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
	// 	logs.Error(err)
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
	// 	logs.Error(err)
	// }
	// _, err = o.Delete(&paths)
	// if err != nil {
	// 	logs.Error(err)
	// }
	//再添加permission
	for _, v1 := range rolepermission {
		// beego.Info(v1.Id)
		if v1.Rolenumber != "" { //存储角色id
			success, err = e.AddPolicy("role_"+strconv.FormatInt(v1.Id, 10), "/onlyoffice/"+strconv.FormatInt(attachments[0].Id, 10), v1.Permission, suf)
			if err != nil {
				logs.Error(err)
			}
		} else { //存储用户id
			success, err = e.AddPolicy(strconv.FormatInt(v1.Id, 10), "/onlyoffice/"+strconv.FormatInt(attachments[0].Id, 10), v1.Permission, suf)
			if err != nil {
				logs.Error(err)
			}
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

// 查询一个文档，哪些用户和角色拥有什么样的权限
// 用casbin的内置方法，不应该用查询数据库方法
func (c *OnlyController) Getpermission() {
	docid := c.GetString("docid")
	//id转成64位
	idNum, err := strconv.ParseInt(docid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	//根据成果id取得所有附件
	attachments, err := models.GetOnlyAttachments(idNum)
	if err != nil {
		logs.Error(err)
	}
	// var users []beegoormadapter.CasbinRule
	rolepermission := make([]Rolepermission, 0)
	for _, w := range attachments {
		// o := orm.NewOrm()
		// qs := o.QueryTable("casbin_rule")
		// _, err = qs.Filter("PType", "p").Filter("v1", "/onlyoffice/"+strconv.FormatInt(w.Id, 10)).All(&users)
		// if err != nil {
		// 	logs.Error(err)
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
					logs.Error(err)
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
					logs.Error(err)
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

//用户新建模板
//上传文档分类：word，excel和ppt

// 文档结构数据
type DocNode struct {
	Id       int    `json:"id"`
	Heading  string `json:"text"`
	Level    int    `json:"level"` //分级
	ParentId int
}

// 树状目录数据——如何定位到word的位置呢
type WordTree struct {
	Id        int         `json:"id"`
	Heading   string      `json:"text"`
	Level     int         `json:"level"` //分级目录，这里其实没什么用了
	WordTrees []*WordTree `json:"nodes"`
}

// 生成word文档的文档结构图
func (c *OnlyController) GetTree() {
	doc, err := document.Open("./attachment/toc.docx")
	if err != nil {
		// log.Fatalf("error opening document: %s", err)
		logs.Error(err)
	}
	var docnode []DocNode
	var id int
	id = 1
	for _, para := range doc.Paragraphs() {
		//if para.Style() == "1" || para.Style() == "Heading1" {
		if para.Style() != "" {
			var text1 string
			for _, run := range para.Runs() {
				text1 = text1 + run.Text()
			}
			aa := make([]DocNode, 1)
			aa[0].Id = id
			aa[0].Heading = text1
			level, err := strconv.Atoi(strings.Replace(para.Style(), "Heading", "", -1))
			if err != nil {
				logs.Error(err)
			}
			aa[0].Level = level
			//循环赋给parentid
			var ispass bool
			ispass = false
			if len(docnode) > 0 {
				for i := len(docnode); i > 0; i-- {
					if level > docnode[i-1].Level {
						aa[0].ParentId = docnode[i-1].Id
						ispass = true
						break
					}
				}
			}
			if !ispass {
				aa[0].ParentId = 0
			}
			//aa := DocNode{id, text1, level, parentid}
			docnode = append(docnode, aa...)
			id++
		}
	}
	//先构造第0级的树状数据结构
	root := WordTree{0, "文档结构", 0, []*WordTree{}}
	//下面node是那些level=1的
	var node []DocNode
	for _, k := range docnode {
		if k.Level == 1 {
			node = append(node, k)
		}
	}
	//递归生成目录json
	makedoctree(node, docnode, &root)
	// beego.Info(root.WordTrees[0])//指针是这样显示的！！！！！！！
	c.Data["json"] = root
	c.TplName = "doctree.tpl"
}

// 递归生成树状结构数据
func makedoctree(node, nodes []DocNode, tree *WordTree) {
	// 遍历第一层
	for _, v := range node {
		id := v.Id
		heading := v.Heading
		level := v.Level
		// 将当前名、层级和id作为子节点添加到目录下
		child := WordTree{id, heading, level, []*WordTree{}}
		tree.WordTrees = append(tree.WordTrees, &child)
		slice := getnodesons(id, nodes)
		//fmt.Println(slice)
		// 如果遍历的当前节点下还有节点，则进入该节点进行递归
		if len(slice) > 0 {
			makedoctree(slice, nodes, &child)
		}
	}
	return
}

// 取得这个id的下级（儿子）目录
func getnodesons(idNum int, nodes []DocNode) (slice []DocNode) {
	for _, k := range nodes {
		if k.ParentId == idNum {
			slice = append(slice, k)
		}
	}
	return slice
}

// 文档格式转换
type Conversionsend struct {
	Async      bool   `json:"async"`
	Filetype   string `json:"filetype"`
	Key        string `json:"key"`
	Outputtype string `json:"async"`
	Thumbnail  Nail   `json:"thumbnail"`
	Title      string `json:"title"`
	Url        string `json:"url"`
}

type Nail struct {
	Aspect int  `json:"aspect"`
	First  bool `json:"first"`
	Height int  `json:"height"`
	Width  int  `json:"width"`
}
type Conversionresponse struct {
	EndConvert bool   `json:"endconvert"`
	FileUrl    string `json:"fileurl"`
	Percent    int    `json:"percent"`
}

// @Title post conversion doc
// @Description post doc to onlyoffice conversion
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /conversion [post]
func (c *OnlyController) Conversion() {
	// {
	// 	"async": false,
	//    "filetype": "docx",
	//    "key": "Khirz6zTPdfd7",
	//    "outputtype": "png",
	//    "thumbnail": {
	//        "aspect": 0,
	//        "first": true,
	//        "height": 150,
	//        "width": 100
	//    },
	//    "title": "Example Document Title.docx",
	//    "url": "https://example.com/url-to-example-document.docx"
	// }
	// {
	//    "endConvert": true,
	//    "fileUrl": "https://documentserver/ResourceService.ashx?filename=output.doc",
	//    "percent": 100
	// }
	var nail Nail
	nail.Aspect = 0
	nail.First = true
	nail.Height = 850
	nail.Width = 600
	var conversionsend Conversionsend
	conversionsend.Async = false
	conversionsend.Filetype = "docx"
	conversionsend.Key = "Khirz6zTPdfd7"
	conversionsend.Outputtype = "pdf"
	conversionsend.Thumbnail = nail
	conversionsend.Title = "Example Document Title.docx"
	conversionsend.Url = "http://192.168.99.1/attachment/onlyoffice/111历史版本试验v4.docx"

	req := httplib.Post("http://192.168.99.100:9000/convertservice.ashx")
	// req.Header("contentType", "application/json")
	req.Header("Content-Type", "application/json")
	// 	bt,err:=ioutil.ReadFile("hello.txt")
	// if err!=nil{
	//     log.Fatal("read file err:",err)
	// }
	// beego.Info(conversionsend)
	b, err := json.Marshal(conversionsend)
	req.Body(string(b))
	// beego.Info(string(b))
	var conversionresponse Conversionresponse

	jsonstring, err := req.String()
	if err != nil {
		logs.Error(err)
	} else {
		//json字符串解析到结构体，以便进行追加
		// beego.Info(jsonstring)
		// err = json.Unmarshal([]byte(jsonstring), &conversionresponse)
		err = xml.Unmarshal([]byte(jsonstring), &conversionresponse)
		// 	fmt.Println(s)
		if err != nil {
			logs.Error(err)
		}

		resp, err := http.Get(conversionresponse.FileUrl)
		if err != nil {
			logs.Error(err)
		}
		// beego.Info(resp)
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logs.Error(err)
		}
		defer resp.Body.Close()
		// if err != nil {
		// 	logs.Error(err)
		// }
		f, err := os.Create("./attachment/onlyoffice/" + "Example Document Title.pdf")
		if err != nil {
			logs.Error(err)
		}
		defer f.Close()
		_, err = f.Write(body) //这里直接用resp.Body如何？
		// _, err = f.WriteString(str)
		// _, err = io.Copy(body, f)
		if err != nil {
			logs.Error(err)
		}

		// http.ServeFile(c.Ctx.ResponseWriter, c.Ctx.Request, "//attachment/onlyoffice/Example Document Title.docx")
		filePath := "attachment/onlyoffice/Example Document Title.pdf"
		c.Ctx.Output.Download(filePath) //这个能保证下载文件名称正确
		c.Data["json"] = conversionresponse
		c.ServeJSON()
	}
}

// https://192.168.101.24:9000/coauthoring/CommandService.ashx
// {
//     "c": "info",
//     "key": "Khirz6zTPdfd7"
// }
// {
//     "key": "Khirz6zTPdfd7",
//     "status": 1,
//     "users": ["6d5a81d0", "78e1e841"]
// }

type CommandSend struct {
	C        string     `json:"c"`
	Key      string     `json:"key"`
	Users    []string   `json:"users"`
	UserData string     `json:"userdata"`
	Meta     metastring `json:"meta"`
}

type metastring struct {
	Title    string `json:"title"`
	Type     string `json:"type"`
	Presence string `json:"presence"`
}

type CommandResponse struct {
	Key    string   `json:"key"`
	Error  int      `json:"error"`
	Users  []string `json:"users"`
	Status int      `json:"status"`
}

type OnlyUsersTable struct {
	UserNickname string `json:"usernickname"`
	UserId       string `json:"userid"`
	Key          string `json:"key"`
}

// @Title post CommandService info doc
// @Description post doc to onlyoffice CommandService
// @Param id path string true "The id of doc"
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 article not found
// @router /commandinfo/:id [get]
// 查询正在编辑的用户列表
func (c *OnlyController) CommandInfo() {
	id := c.Ctx.Input.Param(":id")
	if id == "" {
		return
	}
	//id转成64为
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

	var commandsend CommandSend
	commandsend.C = "info"
	// beego.Info(onlyattachment.Updated)
	commandsend.Key = strconv.FormatInt(onlyattachment.Updated.UnixNano(), 10) //"1640401834797572400"
	// beego.Info(commandsend.Key)
	onlyofficeapi_url, err := web.AppConfig.String("onlyofficeapi_url")
	if err != nil {
		logs.Error(err)
	}
	req := httplib.Post(onlyofficeapi_url + "/coauthoring/CommandService.ashx")
	req.Header("Content-Type", "application/json")
	req.Header("accept", "*/*")
	b, err := json.Marshal(commandsend)
	req.Body(string(b))
	// beego.Info(string(b))
	var commandresponse CommandResponse

	jsonstring, err := req.String()
	if err != nil {
		// logs.Error(err)
		c.Data["json"] = map[string]interface{}{"errNo": 0, "info": "ERR", "data": "请求错误", "msg": "CommandService无返回"}
		c.ServeJSON()
	} else {
		//json字符串解析到结构体，以便进行追加
		// beego.Info(jsonstring)
		err = json.Unmarshal([]byte(jsonstring), &commandresponse)
		// err = xml.Unmarshal([]byte(jsonstring), &commandresponse)
		if err != nil {
			logs.Error(err)
		}
		// beego.Info(commandresponse)
		onlyuserstableslice := make([]OnlyUsersTable, 0)
		// beego.Info(OnlyUsers)
		for _, v := range OnlyUsers {
			onlyuserstable := make([]OnlyUsersTable, 1)
			// 根据id查询用户昵称
			useridint64, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				logs.Error(err)
				return
			}
			user := models.GetUserByUserId(useridint64)
			onlyuserstable[0].UserNickname = user.Nickname
			onlyuserstable[0].UserId = strconv.FormatInt(user.Id, 10)
			onlyuserstable[0].Key = commandsend.Key
			onlyuserstableslice = append(onlyuserstableslice, onlyuserstable...)
		}
		c.Data["json"] = onlyuserstableslice //commandresponse
		c.ServeJSON()
	}
}

// @Title get CommandService drop doc
// @Description get doc to onlyoffice CommandService
// @Param id path string true "The id of user"
// @Param key query string  true "The key of doc"
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 article not found
// @router /commanddrop/:id [get]
// 强制断开指定用户连接
func (c *OnlyController) CommandDrop() {
	id := c.Ctx.Input.Param(":id")
	if id == "" {
		return
	}
	key := c.GetString("key")
	var commandsend CommandSend
	commandsend.C = "drop"
	commandsend.Key = key
	userid := []string{id}
	commandsend.Users = userid
	onlyofficeapi_url, err := web.AppConfig.String("onlyofficeapi_url")
	if err != nil {
		logs.Error(err)
	}
	req := httplib.Post(onlyofficeapi_url + "/coauthoring/CommandService.ashx")
	req.Header("Content-Type", "application/json")
	req.Header("accept", "*/*")
	b, err := json.Marshal(commandsend)
	req.Body(string(b))
	// beego.Info(string(b))
	var commandresponse CommandResponse

	jsonstring, err := req.String()
	if err != nil {
		// logs.Error(err)
		c.Data["json"] = map[string]interface{}{"errNo": 0, "info": "ERR", "data": "请求错误", "msg": "CommandService无返回"}
		c.ServeJSON()
	} else {
		//json字符串解析到结构体，以便进行追加
		// beego.Info(jsonstring)
		err = json.Unmarshal([]byte(jsonstring), &commandresponse)
		// err = xml.Unmarshal([]byte(jsonstring), &commandresponse)
		if err != nil {
			logs.Error(err)
		}
		// beego.Info(commandresponse)
		c.Data["json"] = commandresponse
		c.ServeJSON()
	}
}

// @Title get CommandService builder doc
// @Description get doc to onlyoffice CommandService
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 article not found
// @router /commandbuilder [get]
// 新建文档——必须有build服务器
func (c *OnlyController) CommandBuilder() {
	// id := c.Ctx.Input.Param(":id")
	// if id == "" {
	// 	return
	// }
	// key := c.GetString("key")
	var commandsend CommandSend
	commandsend.C = "drop"
	// commandsend.Key = key
	// userid := []string{id}
	// commandsend.Users = userid
	req := httplib.Post("http://192.168.101.24:9000/docbuilder")
	req.Header("Content-Type", "application/json")
	req.Header("accept", "*/*")
	b, err := json.Marshal(commandsend)
	req.Body(string(b))
	// beego.Info(string(b))
	var commandresponse CommandResponse

	jsonstring, err := req.String()
	if err != nil {
		// logs.Error(err)
		c.Data["json"] = map[string]interface{}{"errNo": 0, "info": "ERR", "data": "请求错误", "msg": "CommandService无返回"}
		c.ServeJSON()
	} else {
		//json字符串解析到结构体，以便进行追加
		// beego.Info(jsonstring)
		err = json.Unmarshal([]byte(jsonstring), &commandresponse)
		// err = xml.Unmarshal([]byte(jsonstring), &commandresponse)
		if err != nil {
			logs.Error(err)
		}
		// beego.Info(commandresponse)
		c.Data["json"] = commandresponse
		c.ServeJSON()
	}
}

// 4.func Unmarshal(data []byte, v interface{}) error
// 将一个 xml 反序列化为对象 结构体中的字段必须是公有的,即大写字母开头的。
// 如果要解析的 xml 是小的,可以 使用 tag 来指定 Struct 的字段与 xml 标记的对应关系

// package main

// import (
// 	"encoding/xml"
// 	"fmt"
// )

// type Student struct {
// 	XMLName xml.Name `xml:"student"`
// 	Name    string   `xml:"name"`
// 	Age     int      `xml:"age"`
// }

// func main() {
// 	str := `<?xml version="1.0" encoding="utf-8"?>
//            <student>
// <name>张三</name> <age>19</age> </student>`
// 	var s Student
// 	xml.Unmarshal([]byte(str), &s)
// 	fmt.Println(s)
// }

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
// 	logs.Error(err)
// }
//打开文件句柄操作
// fh, err := os.Open("tt")
// if err != nil {
// 	logs.Error(err)
// }
// defer fh.Close()

//iocopy
// _, err = io.Copy(fileWriter, fh)
// if err != nil {
// 	logs.Error(err)
// }

// contentType := bodyWriter.FormDataContentType()
// bodyWriter.Close()

// resp1, err := http.Post("http://192.168.99.1/onlyoffice/post?id="+id, contentType, bodyBuf)
// if err != nil {
// 	logs.Error(err)
// }
// defer resp1.Body.Close()

// err = c.SaveToFile("tt", "/attachment/wiki/2018February/1.doc")
// if err != nil {
// 	logs.Error(err)
// }

// resp_body, err := ioutil.ReadAll(resp1.Body)
// if err != nil {
// 	logs.Error(err)
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
// 	id := c.GetString("id")
// 	//pid转成64为
// 	idNum, err := strconv.ParseInt(id, 10, 64)
// 	if err != nil {
// 		logs.Error(err)
// 	}

// 	//获取上传的文件
// 	_, h, err := c.GetFile("uploadfile")
// 	if err != nil {
// 		logs.Error(err)
// 	}
// 	if h != nil {
// 		//存入文件夹
// 		err = c.SaveToFile("uploadfile", "./attachment/onlyoffice/"+h.Filename)
// 		if err != nil {
// 			logs.Error(err)
// 		} else {
// 			err = models.UpdateOnlyAttachment(idNum)
// 			if err != nil {
// 				logs.Error(err)
// 			}
// 		}
// 	}
// }
