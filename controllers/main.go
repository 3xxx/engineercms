package controllers

import (
	// "encoding/json"
	"github.com/3xxx/engineercms/models"
	"github.com/astaxie/beego"
	// "io"
	// "encoding/json"
	// "net/http"
	// "github.com/astaxie/beego/httplib"
	// "bytes"
	// "io/ioutil"
	// "mime/multipart"
	// "os"
	"regexp"
	"strings"
	// "time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	var id string
	navid1 := beego.AppConfig.String("navigationid1")
	navid2 := beego.AppConfig.String("navigationid2")
	navid3 := beego.AppConfig.String("navigationid3")
	navid4 := beego.AppConfig.String("navigationid4")
	navid5 := beego.AppConfig.String("navigationid5")
	navid6 := beego.AppConfig.String("navigationid6")
	navid7 := beego.AppConfig.String("navigationid7")
	navid8 := beego.AppConfig.String("navigationid8")
	navid9 := beego.AppConfig.String("navigationid9")
	index := beego.AppConfig.String("defaultindex")
	// beego.Info(index)
	switch index {
	case "IsNav1":
		id = navid1
		c.Redirect("/project/"+id, 301)
	case "IsNav2":
		id = navid2
		// beego.Info(id)
		c.Redirect("/project/"+id, 301)
	case "IsNav3":
		id = navid3
		c.Redirect("/project/"+id, 301)
	case "IsNav4":
		id = navid4
		c.Redirect("/project/"+id, 301)
	case "IsNav5":
		id = navid5
		c.Redirect("/project/"+id, 301)
	case "IsNav6":
		id = navid6
		c.Redirect("/project/"+id, 301)
	case "IsNav7":
		id = navid7
		c.Redirect("/project/"+id, 301)
	case "IsNav8":
		id = navid8
		c.Redirect("/project/"+id, 301)
	case "IsNav9":
		id = navid9
		c.Redirect("/project/"+id, 301)
	case "IsProject":
		c.Redirect("/project", 301)
	case "IsOnlyOffice":
		c.Redirect("/onlyoffice", 301)
	case "IsDesignGant", "IsConstructGant":
		c.Redirect("/projectgant", 301)
	case "IsLogin":
		c.Redirect("/login", 301)
		//登录后默认跳转……
	case "IsIndex":
		c.Redirect("/index", 301)
	default:
		c.Redirect("/index", 301)
		// c.Redirect("/cms", 301)
		// c.TplName = "index.tpl"
		// c.TplName = "index.html"
		// beego.Info("cmsdefault")
		// c.Redirect("/cms", 301)
	}
	// c.TplName = "index.tpl"
	// c.TplName = "engineercms.tpl"
	// c.Data["IsIndex"] = true
	// username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	// c.Data["Username"] = username
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	// c.Data["IsAdmin"] = isadmin
	// c.Data["IsLogin"] = islogin
	// c.Data["Uid"] = uid
	// c.Data["PageStartTime"] = time.Now()
	// u := c.Ctx.Input.UserAgent()
	// matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// if matched == true {
	// 	// beego.Info("移动端~")
	// 	c.Redirect("/project/", 301)
	// } else {
	// 	// beego.Info("电脑端！")
	// 	achemployee := make([]AchEmployee, 0)
	// 	achsecoffice := make([]AchSecoffice, 0)
	// 	achdepart := make([]AchDepart, 0)
	// 	// case 1: //管理员登录显示的侧栏是全部的
	// 	var depcount int                           //部门人员数
	// 	category1, err := models.GetAdminDepart(0) //得到所有分院（部门）
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// 	for i1, _ := range category1 {
	// 		aa := make([]AchDepart, 1)
	// 		aa[0].Id = category1[i1].Id
	// 		aa[0].Level = "1"
	// 		// aa[0].Pid = category[0].Id
	// 		aa[0].Title = category1[i1].Title //分院名称
	// 		// beego.Info(category1[i1].Title)
	// 		category2, err := models.GetAdminDepart(category1[i1].Id) //得到所有科室
	// 		if err != nil {
	// 			beego.Error(err)
	// 		}
	// 		// beego.Info(category2)
	// 		//如果返回科室为空，则直接取得员工
	// 		//这个逻辑判断不完美，如果一个部门即有科室，又有人没有科室属性怎么办，直接挂在部门下的呢？
	// 		//应该是反过来找出所有没有科室字段的人员，把他放在部门下
	// 		if len(category2) > 0 {
	// 			//如果这个部门有科室，则查出科室和科室里的人
	// 			//如果这个部门下无科室
	// 			//或者部门下有科室，但一些人只属于这个部门而无科室属性
	// 			for i2, _ := range category2 {
	// 				bb := make([]AchSecoffice, 1)
	// 				bb[0].Id = category2[i2].Id
	// 				bb[0].Level = "2"
	// 				bb[0].Pid = category1[i1].Id
	// 				bb[0].Title = category2[i2].Title //科室名称
	// 				// beego.Info(category2[i2].Title)
	// 				//根据分院和科室查所有员工
	// 				users, count, err := models.GetUsersbySec(category1[i1].Title, category2[i2].Title) //得到员工姓名
	// 				if err != nil {
	// 					beego.Error(err)
	// 				}
	// 				// beego.Info(count)
	// 				for i3, _ := range users {
	// 					cc := make([]AchEmployee, 1)
	// 					cc[0].Id = users[i3].Id
	// 					cc[0].Level = "3"
	// 					cc[0].Href = users[i3].Ip + ":" + users[i3].Port
	// 					cc[0].Pid = category2[i2].Id
	// 					cc[0].Nickname = users[i3].Nickname //名称
	// 					// beego.Info(users[i3].Nickname)
	// 					// cc[0].Selectable = false
	// 					achemployee = append(achemployee, cc...)
	// 				}
	// 				bb[0].Tags[0] = strconv.Itoa(count)
	// 				bb[0].Employee = achemployee
	// 				bb[0].Selectable = false
	// 				achemployee = make([]AchEmployee, 0) //再把slice置0
	// 				achsecoffice = append(achsecoffice, bb...)
	// 				depcount = depcount + count //部门人员数=科室人员数相加
	// 			}
	// 			// aa[0].Tags[0] = depcount
	// 			// aa[0].Secoffice = achsecoffice
	// 			// aa[0].Selectable = false               //点击展开，默认是true
	// 			// achsecoffice = make([]AchSecoffice, 0) //再把slice置0
	// 			// achdepart = append(achdepart, aa...)
	// 		}
	// 		//查出所有有这个部门但科室名为空的人员
	// 		//根据分院查所有员工
	// 		users, count, err := models.GetUsersbySecOnly(category1[i1].Title) //得到员工姓名
	// 		if err != nil {
	// 			beego.Error(err)
	// 		}
	// 		for i3, _ := range users {
	// 			dd := make([]AchSecoffice, 1)
	// 			dd[0].Id = users[i3].Id
	// 			dd[0].Level = "3"
	// 			dd[0].Href = users[i3].Ip + ":" + users[i3].Port
	// 			dd[0].Pid = category1[i1].Id
	// 			dd[0].Title = users[i3].Nickname //名称——关键，把人员当作科室名
	// 			dd[0].Selectable = true
	// 			achsecoffice = append(achsecoffice, dd...)
	// 		}
	// 		aa[0].Tags[0] = count + depcount
	// 		// count = 0
	// 		depcount = 0
	// 		aa[0].Secoffice = achsecoffice
	// 		aa[0].Selectable = false               //点击展开，默认是true
	// 		achsecoffice = make([]AchSecoffice, 0) //再把slice置0
	// 		achdepart = append(achdepart, aa...)
	// 	}
	// 	c.Data["json"] = achdepart
}

//文档
func (c *MainController) Getecmsdoc() {
	// username, role := checkprodRole(c.Ctx)
	// roleint, err := strconv.Atoi(role)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// // beego.Info(username)
	// // beego.Info(role)
	// if role == "1" {
	// 	c.Data["IsAdmin"] = true
	// } else if roleint > 1 && roleint < 5 {
	// 	c.Data["IsLogin"] = true
	// } else {
	// 	c.Data["IsAdmin"] = false
	// 	c.Data["IsLogin"] = false
	// }
	// c.Data["Username"] = username
	// // c.Data["IsProjects"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	c.TplName = "ecmsdoc.tpl"
}

func (c *MainController) Getmeritmsdoc() {
	// username, role := checkprodRole(c.Ctx)
	// roleint, err := strconv.Atoi(role)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// // beego.Info(username)
	// // beego.Info(role)
	// if role == "1" {
	// 	c.Data["IsAdmin"] = true
	// } else if roleint > 1 && roleint < 5 {
	// 	c.Data["IsLogin"] = true
	// } else {
	// 	c.Data["IsAdmin"] = false
	// 	c.Data["IsLogin"] = false
	// }
	// c.Data["Username"] = username
	// // c.Data["IsProjects"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	c.TplName = "meritmsdoc.tpl"
}

func (c *MainController) Gethydrowsdoc() {
	// username, role := checkprodRole(c.Ctx)
	// roleint, err := strconv.Atoi(role)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// // beego.Info(username)
	// // beego.Info(role)
	// if role == "1" {
	// 	c.Data["IsAdmin"] = true
	// } else if roleint > 1 && roleint < 5 {
	// 	c.Data["IsLogin"] = true
	// } else {
	// 	c.Data["IsAdmin"] = false
	// 	c.Data["IsLogin"] = false
	// }
	// c.Data["Username"] = username
	// // c.Data["IsProjects"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	c.TplName = "hydrowsdoc.tpl"
}

//api
func (c *MainController) Getecmsapi() {
	// username, role := checkprodRole(c.Ctx)
	// roleint, err := strconv.Atoi(role)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// // beego.Info(username)
	// // beego.Info(role)
	// if role == "1" {
	// 	c.Data["IsAdmin"] = true
	// } else if roleint > 1 && roleint < 5 {
	// 	c.Data["IsLogin"] = true
	// } else {
	// 	c.Data["IsAdmin"] = false
	// 	c.Data["IsLogin"] = false
	// }
	// c.Data["Username"] = username
	// // c.Data["IsProjects"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	c.TplName = "ecmsapi.tpl"
}

func (c *MainController) Getmeritmsapi() {
	// username, role := checkprodRole(c.Ctx)
	// roleint, err := strconv.Atoi(role)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// // beego.Info(username)
	// // beego.Info(role)
	// if role == "1" {
	// 	c.Data["IsAdmin"] = true
	// } else if roleint > 1 && roleint < 5 {
	// 	c.Data["IsLogin"] = true
	// } else {
	// 	c.Data["IsAdmin"] = false
	// 	c.Data["IsLogin"] = false
	// }
	// c.Data["Username"] = username
	// // c.Data["IsProjects"] = true
	// c.Data["Ip"] = c.Ctx.Input.IP()
	// c.Data["role"] = role
	username, role, uid, isadmin, islogin := checkprodRole(c.Ctx)
	c.Data["Username"] = username
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Uid"] = uid
	c.TplName = "meritmsapi.tpl"
}

// @Title get usermanage
// @Description get usermanage
// @Success 200 {object} models.GetAttachbyId
// @Failure 400 Invalid page supplied
// @Failure 404 pdf not found
// @router /usermanage [get]
func (c *MainController) UserManage() {
	// beego.Info(c.Ctx.Input.UserAgent())
	u := c.Ctx.Input.UserAgent()
	// re := regexp.MustCompile("Trident")
	// loc := re.FindStringIndex(u)
	// loc[0] > 1
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		beego.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		c.TplName = "userManagement.html"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "userManagement.html"
	}
}

// @Title get photoswipe
// @Description get photoswipe
// @Success 200 {object} models.GetAttachbyId
// @Failure 400 Invalid page supplied
// @Failure 404 pdf not found
// @router /photoswipe [get]
func (c *MainController) PhotoSwipe() {
	// beego.Info(c.Ctx.Input.UserAgent())
	u := c.Ctx.Input.UserAgent()
	// re := regexp.MustCompile("Trident")
	// loc := re.FindStringIndex(u)
	// loc[0] > 1
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		beego.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		c.TplName = "photoswipe.html"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "photoswipe.html"
	}
}

func (c *MainController) Test() {
	// beego.Info(c.Ctx.Input.UserAgent())
	u := c.Ctx.Input.UserAgent()
	// re := regexp.MustCompile("Trident")
	// loc := re.FindStringIndex(u)
	// loc[0] > 1
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		beego.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		c.TplName = "test1.tpl"
	} else {
		// beego.Info("电脑端！")
		c.TplName = "test.tpl"
	}
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

func (c *MainController) Slide() {
	c.TplName = "slide.tpl"
}

// 获取上传过来的文件
// func (c *MainController) GetPostdata() {
// 	const lll = "2006-01-02"
// 	// convdate := time.Now().Format(lll)
// 	// t1, err := time.Parse(lll, convdate) //这里t1要是用t1:=就不是前面那个t1了
// 	// if err != nil {
// 	// 		beego.Error(err)
// 	// 	}
// 	date := time.Now()
// 	convdate := string(date.Format(lll))

// 	f, _, err := c.GetFile("uploadfile")
// 	// beego.Info(h) //这里 filename 是路径，所以不能以filename作为保存的文件名。坑！！
// 	defer f.Close()
// 	getpostdatafilepath := beego.AppConfig.String("getpostdatafilepath") //"./static/upload/"
// 	getpostdatafilename := beego.AppConfig.String("getpostdatafilename")
// 	if err != nil {
// 		beego.Error(err)
// 	} else {
// 		c.SaveToFile("uploadfile", getpostdatafilepath+convdate+getpostdatafilename) // ".data"保存位置在 static/upload, 没有文件夹要先创建
// 		c.Ctx.WriteString("ok")
// 	}
// }

// @Title post Data(file)
// @Description post Data(file)
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /sendmessage [post]
// 检验signature v1/checkin/sendmessage
// 上传文件、备份文件
// func (c *MainController) Postdata() {
// 	postdataurl := beego.AppConfig.String("postdataurl")
// 	postdatausername := beego.AppConfig.String("postdatausername")
// 	postdatapassword := beego.AppConfig.String("postdatapassword")
// 	postdatafilepath := beego.AppConfig.String("postdatafilepath")
// 	b := httplib.Post(postdataurl) //"./database/meritms.db"
// 	b.Param("username", postdatausername)
// 	b.Param("password", postdatapassword)
// 	b.PostFile("uploadfile", postdatafilepath) //./static/
// 	// b.PostFile("uploadfile2", "httplib.txt")PostFile 第一个参数是 form 表单的字段名,第二个是需要发送的文件名或者文件路径
// 	str, err := b.String()
// 	if err != nil {
// 		beego.Error(str)
// 	}
// }

// 上传文件、备份文件
// func Postdata() {
// 	postdataurl := beego.AppConfig.String("postdataurl")
// 	postdatausername := beego.AppConfig.String("postdatausername")
// 	postdatapassword := beego.AppConfig.String("postdatapassword")
// 	postdatafilepath := beego.AppConfig.String("postdatafilepath")
// 	b := httplib.Post(postdataurl) //"./database/meritms.db"
// 	b.Param("username", postdatausername)
// 	b.Param("password", postdatapassword)
// 	b.PostFile("uploadfile", postdatafilepath) //./static/
// 	// b.PostFile("uploadfile2", "httplib.txt")PostFile 第一个参数是 form 表单的字段名,第二个是需要发送的文件名或者文件路径
// 	str, err := b.String()
// 	if err != nil {
// 		beego.Error(str)
// 	}
// }

func (c *MainController) Register() {
	// flash := beego.NewFlash()
	token := c.Input().Get("token")
	//是否重复提交
	if c.IsSubmitAgain(token) {
		c.Redirect("/registerpage", 302)
		return
	}
}

func (c *MainController) IsSubmitAgain(token string) bool {
	cotoken := c.Ctx.GetCookie("token")
	if token == "" || len(token) == 0 || token != cotoken || strings.Compare(cotoken, token) != 0 {
		return true
	}
	return false
}

//升级数据库
func (c *MainController) UpdateDatabase() {
	beego.Info("ok")
	err1, err2, err3, err4, err5, err6, err7 := models.UpdateDatabase()
	if err1 != nil {
		beego.Error(err1)
	}
	if err2 != nil {
		beego.Error(err2)
	}
	if err3 != nil {
		beego.Error(err3)
	}
	if err4 != nil {
		beego.Error(err4)
	}
	if err5 != nil {
		beego.Error(err5)
	}
	if err6 != nil {
		beego.Error(err6)
	}
	if err7 != nil {
		beego.Error(err7)
	}
	c.Data["json"] = "ok"
	c.ServeJSON()
}

//删除数据表和字段测试
func (c *MainController) ModifyDatabase() {
	err := models.ModifyDatabase()
	if err != nil {
		beego.Error(err)
	}
	c.Data["json"] = err
	c.ServeJSON()
}

func (c *MainController) Autodesk() {
	c.TplName = "autodeskview.tpl"
}
