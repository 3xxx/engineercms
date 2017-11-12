package controllers

import (
	"engineercms/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils/pagination"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
	// "github.com/astaxie/beego/httplib"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	username, role := checkprodRole(c.Ctx)
	if role == 1 {
		c.Data["IsAdmin"] = true
	} else if role > 1 && role < 5 {
		c.Data["IsLogin"] = true
	} else {
		c.Data["IsAdmin"] = false
		c.Data["IsLogin"] = false
	}
	c.Data["Username"] = username
	c.Data["IsIndex"] = true
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role

	// beego.Info(username)
	// beego.Info(role)
	if role == 1 {
		c.Data["IsAdmin"] = true
	} else if role > 1 && role < 5 {
		c.Data["IsLogin"] = true
	} else {
		c.Data["IsAdmin"] = false
		c.Data["IsLogin"] = false
	}
	c.Data["Username"] = username
	// c.Data["IsProjects"] = true
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role

	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		beego.Error(err)
	}
	if matched == true {
		// beego.Info("移动端~")
		// c.TplName = "mproject.tpl"
		c.Redirect("/project/25002", 301)
	} else {
		// beego.Info("电脑端！")
		achemployee := make([]AchEmployee, 0)
		achsecoffice := make([]AchSecoffice, 0)
		achdepart := make([]AchDepart, 0)
		//由uname取得user,获得user的分院名称
		// user, err := models.GetUserByUsername(uname)
		// if err != nil {
		// 	beego.Error(err)
		// }
		// switch role {
		// case 1: //管理员登录显示的侧栏是全部的
		var depcount int                           //部门人员数
		category1, err := models.GetAdminDepart(0) //得到所有分院（部门）
		if err != nil {
			beego.Error(err)
		}
		for i1, _ := range category1 {
			aa := make([]AchDepart, 1)
			aa[0].Id = category1[i1].Id
			aa[0].Level = "1"
			// aa[0].Pid = category[0].Id
			aa[0].Title = category1[i1].Title //分院名称
			// beego.Info(category1[i1].Title)
			category2, err := models.GetAdminDepart(category1[i1].Id) //得到所有科室
			if err != nil {
				beego.Error(err)
			}
			// beego.Info(category2)
			//如果返回科室为空，则直接取得员工
			//这个逻辑判断不完美，如果一个部门即有科室，又有人没有科室属性怎么办，直接挂在部门下的呢？
			//应该是反过来找出所有没有科室字段的人员，把他放在部门下
			if len(category2) > 0 {
				//如果这个部门有科室，则查出科室和科室里的人
				//如果这个部门下无科室
				//或者部门下有科室，但一些人只属于这个部门而无科室属性
				for i2, _ := range category2 {
					bb := make([]AchSecoffice, 1)
					bb[0].Id = category2[i2].Id
					bb[0].Level = "2"
					bb[0].Pid = category1[i1].Id
					bb[0].Title = category2[i2].Title //科室名称
					// beego.Info(category2[i2].Title)
					//根据分院和科室查所有员工
					users, count, err := models.GetUsersbySec(category1[i1].Title, category2[i2].Title) //得到员工姓名
					if err != nil {
						beego.Error(err)
					}
					// beego.Info(count)
					for i3, _ := range users {
						cc := make([]AchEmployee, 1)
						cc[0].Id = users[i3].Id
						cc[0].Level = "3"
						cc[0].Href = users[i3].Ip + ":" + users[i3].Port
						cc[0].Pid = category2[i2].Id
						cc[0].Nickname = users[i3].Nickname //名称
						// beego.Info(users[i3].Nickname)
						// cc[0].Selectable = false
						achemployee = append(achemployee, cc...)
					}
					bb[0].Tags[0] = strconv.Itoa(count)
					bb[0].Employee = achemployee
					bb[0].Selectable = false
					achemployee = make([]AchEmployee, 0) //再把slice置0
					achsecoffice = append(achsecoffice, bb...)
					depcount = depcount + count //部门人员数=科室人员数相加
				}
				// aa[0].Tags[0] = depcount
				// aa[0].Secoffice = achsecoffice
				// aa[0].Selectable = false               //点击展开，默认是true
				// achsecoffice = make([]AchSecoffice, 0) //再把slice置0
				// achdepart = append(achdepart, aa...)
			}
			//查出所有有这个部门但科室名为空的人员
			//根据分院查所有员工
			users, count, err := models.GetUsersbySecOnly(category1[i1].Title) //得到员工姓名
			if err != nil {
				beego.Error(err)
			}
			for i3, _ := range users {
				dd := make([]AchSecoffice, 1)
				dd[0].Id = users[i3].Id
				dd[0].Level = "3"
				dd[0].Href = users[i3].Ip + ":" + users[i3].Port
				dd[0].Pid = category1[i1].Id
				dd[0].Title = users[i3].Nickname //名称——关键，把人员当作科室名
				dd[0].Selectable = true
				achsecoffice = append(achsecoffice, dd...)
			}
			aa[0].Tags[0] = count + depcount
			// count = 0
			depcount = 0
			aa[0].Secoffice = achsecoffice
			aa[0].Selectable = false               //点击展开，默认是true
			achsecoffice = make([]AchSecoffice, 0) //再把slice置0
			achdepart = append(achdepart, aa...)
		}
		c.Data["json"] = achdepart
		c.TplName = "index.tpl"
		// c.TplName = "engineercms.tpl"

		// c.Redirect("/project/25002", 301)
	}
}

//文档
func (c *MainController) Getecmsdoc() {
	username, role := checkprodRole(c.Ctx)
	// beego.Info(username)
	// beego.Info(role)
	if role == 1 {
		c.Data["IsAdmin"] = true
	} else if role > 1 && role < 5 {
		c.Data["IsLogin"] = true
	} else {
		c.Data["IsAdmin"] = false
		c.Data["IsLogin"] = false
	}
	c.Data["Username"] = username
	// c.Data["IsProjects"] = true
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.TplName = "ecmsdoc.tpl"
}

func (c *MainController) Getmeritmsdoc() {
	username, role := checkprodRole(c.Ctx)
	// beego.Info(username)
	// beego.Info(role)
	if role == 1 {
		c.Data["IsAdmin"] = true
	} else if role > 1 && role < 5 {
		c.Data["IsLogin"] = true
	} else {
		c.Data["IsAdmin"] = false
		c.Data["IsLogin"] = false
	}
	c.Data["Username"] = username
	// c.Data["IsProjects"] = true
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.TplName = "meritmsdoc.tpl"
}

func (c *MainController) Gethydrowsdoc() {
	username, role := checkprodRole(c.Ctx)
	// beego.Info(username)
	// beego.Info(role)
	if role == 1 {
		c.Data["IsAdmin"] = true
	} else if role > 1 && role < 5 {
		c.Data["IsLogin"] = true
	} else {
		c.Data["IsAdmin"] = false
		c.Data["IsLogin"] = false
	}
	c.Data["Username"] = username
	// c.Data["IsProjects"] = true
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.TplName = "hydrowsdoc.tpl"
}

//api
func (c *MainController) Getecmsapi() {
	username, role := checkprodRole(c.Ctx)
	// beego.Info(username)
	// beego.Info(role)
	if role == 1 {
		c.Data["IsAdmin"] = true
	} else if role > 1 && role < 5 {
		c.Data["IsLogin"] = true
	} else {
		c.Data["IsAdmin"] = false
		c.Data["IsLogin"] = false
	}
	c.Data["Username"] = username
	// c.Data["IsProjects"] = true
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.TplName = "ecmsapi.tpl"
}

func (c *MainController) Getmeritmsapi() {
	username, role := checkprodRole(c.Ctx)
	// beego.Info(username)
	// beego.Info(role)
	if role == 1 {
		c.Data["IsAdmin"] = true
	} else if role > 1 && role < 5 {
		c.Data["IsLogin"] = true
	} else {
		c.Data["IsAdmin"] = false
		c.Data["IsLogin"] = false
	}
	c.Data["Username"] = username
	// c.Data["IsProjects"] = true
	c.Data["Ip"] = c.Ctx.Input.IP()
	c.Data["role"] = role
	c.TplName = "meritmsapi.tpl"
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
		beego.Info("移动端~")
		c.TplName = "test1.tpl"
	} else {
		beego.Info("电脑端！")
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

func (c *MainController) Postdata() {
	const lll = "2006-01-02"
	// convdate := time.Now().Format(lll)
	// t1, err := time.Parse(lll, convdate) //这里t1要是用t1:=就不是前面那个t1了
	// if err != nil {
	// 		beego.Error(err)
	// 	}
	date := time.Now()
	convdate := string(date.Format(lll))

	f, _, err := c.GetFile("uploadfile")
	// beego.Info(h) //这里 filename是路径，所以不能以filename作为保存的文件名。坑！！
	defer f.Close()

	if err != nil {
		beego.Error(err)
	} else {
		c.SaveToFile("uploadfile", "./static/upload/"+convdate+".data") // 保存位置在 static/upload, 没有文件夹要先创建
		c.Ctx.WriteString("ok")
	}
}

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

func (c *MainController) Pdf() {
	// c.TplName = "web/viewer.html"
	// c.Data["IsLogin"] = checkAccount(c.Ctx)
	// uname, _, _ := checkRoleread(c.Ctx) //login里的
	// rolename, _ = strconv.Atoi(role)
	// c.Data["Uname"] = uname
	// Directory:github.com/astaxie/beego/context     Pakage in Source:context
	// func (input *BeegoInput) IP() string {}
	//c是TopicController,TopicController是beego.controller，即beego.controller.ctx.input.ip

	// beego.Info(c.Ctx.Input.IP())

	//取得附件的id——成果的id——目录的id——查询目录下所有pdf返回数量
	// id := c.Input().Get("id")
	// if id != "" {
	// 	//id转成64为
	// 	idNum, err = strconv.ParseInt(id, 10, 64)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// }
	//取得某个目录下所有pdf
	// id := c.Ctx.Input.Param(":id")
	// beego.Info(id)
	var Url string
	var pNum int64
	id := c.Input().Get("id")
	// beego.Info(id)
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	p := c.Input().Get("p")
	// beego.Info(p)

	var prods []*models.Product
	var projid int64
	if p == "" { //id是附件或成果id
		attach, err := models.GetAttachbyId(idNum)
		if err != nil {
			beego.Error(err)
		}
		//由成果id（后台传过来的行id）取得成果
		prod, err := models.GetProd(attach.ProductId)
		if err != nil {
			beego.Error(err)
		}
		//由侧栏id取得所有成果和所有成果的附件pdf
		prods, err = models.GetProducts(prod.ProjectId)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(prod.ProjectId)
		//由proj id取得url
		// Url, _, err = GetUrlPath(prod.ProjectId)
		// if err != nil {
		// 	beego.Error(err)
		// }
		// beego.Info(Url)
		projid = prod.ProjectId
	} else { //id是侧栏目录id
		//由侧栏id取得所有成果和所有成果的附件pdf
		prods, err = models.GetProducts(idNum)
		if err != nil {
			beego.Error(err)
		}
		//由proj id取得url
		Url, _, err = GetUrlPath(idNum)
		if err != nil {
			beego.Error(err)
		}
		// beego.Info(Url)
		// projid = idNum
		//id转成64为
		pNum, err = strconv.ParseInt(p, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	// c.Data["AttachmentId"] = id

	// var idNum int64
	// var err error

	// if id != "" {

	// var count int64
	Attachments := make([]*models.Attachment, 0)
	for _, v := range prods {
		//根据成果id取得所有附件数量
		// count1, err := models.GetAttachmentsCount(v.Id)
		// if err != nil {
		// 	beego.Error(err)
		// }
		// count = count + count1
		Attachments1, err := models.GetAttachments(v.Id) //只返回Id?没意义，全部返回
		if err != nil {
			beego.Error(err)
		}
		for _, v := range Attachments1 {
			if path.Ext(v.FileName) == ".pdf" || path.Ext(v.FileName) == ".PDF" {
				Attachments = append(Attachments, Attachments1...)
			}
		}
	}

	//对成果进行循环
	//赋予url
	//如果是一个成果，直接给url;如果大于1个，则是数组:这个在前端实现
	// http.ServeFile(ctx.ResponseWriter, ctx.Request, filePath)
	// link := make([]AttachmentLink, 0)
	// for _, v := range Attachments {
	// 	if path.Ext(v.FileName) == ".pdf" || path.Ext(v.FileName) == ".PDF" {
	// 		linkarr := make([]AttachmentLink, 1)
	// 		linkarr[0].Id = v.Id
	// 		linkarr[0].Title = v.FileName
	// 		linkarr[0].FileSize = v.FileSize
	// 		linkarr[0].Downloads = v.Downloads
	// 		linkarr[0].Created = v.Created
	// 		linkarr[0].Updated = v.Updated
	// 		linkarr[0].Link = Url + "/" + v.FileName
	// 		link = append(link, linkarr...)
	// 	}
	// }
	// c.Data["json"] = link

	// pdfs, err := models.GetAllPdfs(idNum, false)
	// if err != nil {
	// 	beego.Error(err)
	// }
	count := len(Attachments)
	count1 := strconv.Itoa(count)
	count2, err := strconv.ParseInt(count1, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	// cnt, err := o.QueryTable("user").Count()
	// if err != nil {
	// 	beego.Error(err)
	// }

	// sets this.Data["paginator"] with the current offset (from the url query param)
	postsPerPage := 1
	paginator := pagination.SetPaginator(c.Ctx, postsPerPage, count2)
	// fmt.Println(*c.Ctx)
	// beego.Info(c.Ctx)
	// beego.Info(paginator.Offset())   0
	// p := pagination.NewPaginator(c.Ctx.Request, 10, 9)
	// beego.Info(p.Offset())   0
	// fetch the next 20 posts
	//根据paginator.Offset()序列，取得附件
	// Attachments := make([]*models.Attachment, 0)
	// for _, v := range prod {
	// 	//根据成果id取得所有附件id
	// 	Attachments1, err := models.GetAttachments(v.Id) //只返回Id?没意义，全部返回
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// 	Attachments = append(Attachments, Attachments1...)
	// }

	//查询当前附件id所在位置
	// var p1, PdfLink string
	// for i, v := range Attachments {
	// 	if v.Id == idNum {
	// 		p1 = strconv.Itoa(i)
	// 		PdfLink = Url + "/" + v.FileName
	// 		break
	// 	}
	// }

	// pdfs, err = models.ListPostsByOffsetAndLimit(paginator.Offset(), postsPerPage)
	// if err != nil {
	// 	beego.Error(err)
	// }
	//这里根据上面取得的分页topics，取出这页的成果对应的所有项目
	// slice1 := make([]models.Category, 0)
	// for _, v := range topics {
	// 	tid := strconv.FormatInt(v.Id, 10)
	// 	category, err := models.GetTopicProj(tid)
	// 	if err != nil {
	// 		beego.Error(err)
	// 	}
	// beego.Info(category.Title)
	// 	aa := make([]models.Category, 1)
	// 	aa[0].Author = category.Title//这句出错，不知何故runtime error: invalid memory address or nil pointer dereference
	// 	slice1 = append(slice1, aa...)
	// }
	// c.Data["Categories"] = slice1

	c.Data["paginator"] = paginator

	if p == "" {
		var p1 string
		for i, v := range Attachments {
			if v.Id == idNum {
				p1 = strconv.Itoa(i + 1)
				// PdfLink = Url + "/" + v.FileName
				break
			}
		}
		c.Redirect("/pdf?p="+p1+"&id="+strconv.FormatInt(projid, 10), 302)
	} else {
		PdfLink := Url + "/" + Attachments[pNum-1].FileName
		// beego.Info(PdfLink)
		c.Data["PdfLink"] = PdfLink
		c.TplName = "web/viewer.html"
	}
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
