package controllers

import (
	// "crypto/md5"
	// "encoding/hex"
	"encoding/json"
	"github.com/3xxx/engineercms/controllers/utils"
	"github.com/3xxx/engineercms/models"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/httplib"
	// "github.com/astaxie/beego/logs"
	// "net"
	// "net/http"
	// "net/url"
	// "path"
	"fmt"
	"strconv"
	"strings"
	// "reflect"
	"sort"
	"time"
	// "github.com/xlstudio/wxbizdatacrypt"
	"crypto/sha1"
	// "encoding/xml"
	// "fmt"
	"io"
	// "io/ioutil"
	// "log"
)

// CMSCHECKIN API
type CheckController struct {
	beego.Controller
}

// @Title post checkin person
// @Description post person
// @Param username query string true "The userId for person"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /person [post]
func (c *CheckController) Person() {
	// var userId = -1
	//    var avatorUrl = ''
	//    var username = req.body.username
	//    // console.log(username);
	//    // var username = '李四';
	//    users.getUserByName(username,function (userInfos) {
	//        console.log(userInfos)
	//        if(userInfos.length === 0) {
	//            users.create(username,function (add_result) {
	//                userId = add_result.insertId
	//                var result = {
	//                    userId: userId
	//                    avatorUrl:avatorUrl
	//                }
	//                res.send(result)
	//            })
	//        }else {
	//            userId = userInfos[0].F_ID
	//            avatorUrl = userInfos[0].F_PhotoUrl
	//            var result = {
	//                userId: userId
	//                avatorUrl: avatorUrl
	//            }
	//            res.send(result)
	//        }

	//        console.log(result)

	//    })
	c.Data["json"] = map[string]interface{}{"userId": 1, "avatorUrl": "头像地址-person"}
	c.ServeJSON()
}

// @Title post checkin activity
// @Description post person
// @Param CreaterId query string true "The CreaterId for activity"
// @Param projectid query string true "The projectid of activity"
// @Param Caption query string true "The Caption for activity"
// @Param Desc query string true "The Desc for activity"
// @Param Location query string true "The Location for activity"
// @Param Lat query string true "The Lat for activity"
// @Param Lng query string true "The Lng for activity"
// @Param SatrtDate query string true "The SatrtDate for activity"
// @Param EndDate query string true "The EndDate for activity"
// @Param IfFace query string true "The IfFace for activity"
// @Param IfPhoto query string true "The IfPhoto for activity"
// @Param IfLocation query string true "The IfLocation for activity"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /activity/create [post]
// 新增打卡活动
func (c *CheckController) Create() {
	/**
	 * /activity/create
	 * POST
	 * 创建新的签到活动
	 */
	// router.post('/create', function (req, res, next) {
	//     var F_CreaterId = req.body.createrId;
	//     var F_Caption = req.body.activity_name;
	//     var F_Desc = req.body.activity_desc;
	//     var F_Location = req.body.location;
	//     var F_Lat = req.body.lat;
	//     var F_Lng = req.body.lng;
	//     var F_SatrtDate = req.body.startDate;
	//     var F_EndDate = req.body.endDate;
	//     var F_IfFace = req.body.ifFace;
	//     var F_IfPhoto = req.body.ifPhoto;
	//     var F_IfLocation = req.body.ifLocation;
	//     var act = {
	//         F_CreaterId: F_CreaterId,
	//         F_Caption: F_Caption,
	//         F_Desc: F_Desc,
	//         F_Location: F_Location,
	//         F_Lat: F_Lat,
	//         F_Lng: F_Lng,
	//         F_StartDate: F_SatrtDate,
	//         F_EndDate: F_EndDate,
	//         F_IfFace: F_IfFace,
	//         F_IfPhoto: F_IfPhoto,
	//         F_IfLocation: F_IfLocation
	//     };
	//     // console.log(act);
	//     activity.create(act,function (act_res) {
	//         console.log(act_res);
	//         apply.apply(F_CreaterId,act_res.insertId,function (apply_res) {
	//             var insert_res = {
	//                 code:1
	//             };
	//             res.send(insert_res);
	//         })
	//     });
	// });
	// 	 'createrId
	// 'activity_name
	// 'activity_desc
	// 'location': va
	// 'lat': app.dat
	// 'lng': app.dat
	// 'startDate': v
	// 'endDate': val
	// 'ifFace': ifFa
	// 'ifPhoto': ifP
	// 'ifLocation':
	// 	 activity_desc:
	// "南沙管理部设代成立！"
	// activity_name:"大南沙设代总部"
	// createrId:1
	// endDate:"2022-12-31"
	// ifFace:0
	// ifLocation:1
	// ifPhoto:0
	// lat:23.12247
	// lng:113.36148
	// location:"东璟花园"
	// startDate:"2019-05-01"
	createrid := c.Input().Get("createrId")
	// beego.Info(createrid)
	//pid转成64为
	CreaterId, err := strconv.ParseInt(createrid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	projectid := c.Input().Get("projectid")
	var ProjectId int64
	if projectid != "" {
		ProjectId, err = strconv.ParseInt(projectid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}
	Caption := c.Input().Get("activity_name")
	Desc := c.Input().Get("activity_desc")
	Location := c.Input().Get("location")
	lat := c.Input().Get("lat")
	//string到float64
	Lat, err := strconv.ParseFloat(lat, 64)

	lng := c.Input().Get("lng")
	Lng, err := strconv.ParseFloat(lng, 64)

	startdate := c.Input().Get("startDate")
	datearray := strings.Split(startdate, "-")
	year := datearray[0]
	month := datearray[1]
	if len(month) == 1 {
		month = "0" + month
	}
	day := datearray[2]
	if len(day) == 1 {
		day = "0" + day
	}
	startdate = year + "-" + month + "-" + day
	const base_format = "2006-01-02"
	StartDate, err := time.Parse(base_format, startdate)
	if err != nil {
		beego.Error(err)
	}
	enddate := c.Input().Get("endDate")
	datearray = strings.Split(enddate, "-")
	year = datearray[0]
	month = datearray[1]
	if len(month) == 1 {
		month = "0" + month
	}
	day = datearray[2]
	if len(day) == 1 {
		day = "0" + day
	}
	enddate = year + "-" + month + "-" + day
	// beego.Info(enddate)
	EndDate, err := time.Parse(base_format, enddate)
	if err != nil {
		beego.Error(err)
	}
	ifface := c.Input().Get("ifFace")
	IfFace, err := strconv.ParseBool(ifface) // string 转bool
	ifphoto := c.Input().Get("ifPhoto")
	IfPhoto, err := strconv.ParseBool(ifphoto) // string 转bool
	iflocation := c.Input().Get("ifLocation")
	IfLocation, err := strconv.ParseBool(iflocation) // string 转bool
	_, err = models.CheckCreate(CreaterId, ProjectId, Caption, Desc, Location, Lat, Lng, StartDate, EndDate, IfFace, IfPhoto, IfLocation)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = "ok"
		c.ServeJSON()
	}
}

type GetAll struct {
	Completed  []models.Activity `json:"completed"`
	Processing []models.Activity `json:"processing"`
}

type detail struct {
	ID        string `json:"F_ID"`
	Caption   string `json:"F_Caption"`
	StartDate string `json:"F_StartDate"`
	EndDate   string `json:"F_EndDate"`
	Location  string `json:"F_Location"`
}

// @Title post checkin person
// @Description post person
// @Param projectid query string true "The projectid of activity"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /activity/getall [post]
//取出所有活动
func (c *CheckController) Getall() {
	// var user_id = req.body.userId;
	//   var now = moment();
	//   var completed = [];
	//   var processing = [];
	//   console.log(now);
	//   activity.getall(user_id,function (activitys) {
	//       for(var i = 0; i < activitys.length; i ++){
	//           activitys[i].F_StartDate = moment(activitys[i].F_StartDate).format('YYYY-MM-DD');
	//           activitys[i].F_EndDate = moment(activitys[i].F_EndDate).format('YYYY-MM-DD');
	//           //<0说明在endDate在现在之前---已经结束了
	//           if(moment(activitys[i].F_EndDate).diff(now)<0){
	//               completed.push(activitys[i]);
	//           }else {
	//               processing.push(activitys[i]);
	//           }
	//       }
	//       var res_getall = {
	//           'completed':completed,
	//           'processing':processing
	//       };
	//       res.send(res_getall);
	//   })
	//   getall: function (user_id,callback) {
	//       var sql = 'SELECT * FROM v_apply_activity WHERE F_ApplyerId=? ';
	//       var params = [user_id];
	//       db.connection.query(sql,params,function (err,res) {
	//           if(err){
	//               console.log(err);
	//               return;
	//           }
	//           callback(res);
	//       })
	//   },
	var ProjectId int64
	var err error
	projectid := c.Input().Get("projectid")
	if projectid != "" {
		ProjectId, err = strconv.ParseInt(projectid, 10, 64)
		if err != nil {
			beego.Error(err)
		}
	}

	//查出未过期的
	activities, err := models.GetAll(ProjectId)
	if err != nil {
		beego.Error(err)
	}
	var getall GetAll
	// local, _ := time.LoadLocation("Local")
	// starttime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-10-23 15:01:01", local)
	// endtime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2017-10-24 15:05:01", local)
	// now := time.Now()
	// fmt.Println(now)
	// fmt.Println(endtime, endtime.After(now))
	// fmt.Println(starttime, starttime.Before(now))
	activityslice := make([]models.Activity, 0)
	activityslice2 := make([]models.Activity, 0)
	for _, v := range activities {
		// endtime, err := time.ParseInLocation("2006-01-02 15:04:05", v.EndDate, time.Local)
		// if err != nil {
		// 	beego.Error(err)
		// }
		detail2 := make([]models.Activity, 1)
		detail2[0].Id = v.Id
		detail2[0].Caption = v.Caption
		detail2[0].StartDate = v.StartDate
		detail2[0].EndDate = v.EndDate
		detail2[0].Location = v.Location
		if v.EndDate.After(time.Now()) {
			activityslice = append(activityslice, detail2...)
		} else { //查出过期的
			activityslice2 = append(activityslice2, detail2...)
		}
	}

	getall.Processing = activityslice
	getall.Completed = activityslice2

	c.Data["json"] = getall //map[string]interface{}{"userId": 1, "avatorUrl": "Filename"}
	c.ServeJSON()
}

// @Title post checkin person
// @Description post person
// @Param username query string true "The userId for person"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /activity/like [post]
//搜索
func (c *CheckController) Like() {
	//  * 模糊查询所有的活动
	//  */
	// router.post('/like',function (req,res,next) {
	//     // var re = {
	//     //     code:1
	//     // };
	//     // res.send(re);
	//     var str = req.body.str;
	//     // var str = '火锅';
	//     var now = moment();
	//     var completed = [];
	//     var processing = [];
	//     activity.getallLike(str,function (activitys) {
	//         for(var i = 0; i < activitys.length; i ++){
	//             activitys[i].F_StartDate = moment(activitys[i].F_StartDate).format('YYYY-MM-DD');
	//             activitys[i].F_EndDate = moment(activitys[i].F_EndDate).format('YYYY-MM-DD');
	//             //<0说明在endDate在现在之前---已经结束了
	//             if(moment(activitys[i].F_EndDate).diff(now)<0){
	//                 completed.push(activitys[i]);
	//             }else {
	//                 processing.push(activitys[i]);
	//             }
	//         }
	//         var res_getall = {
	//             'completed':completed,
	//             'processing':processing
	//         };
	//         res.send(res_getall);
	//     })
	// });
	// getallLike:function (str,callback) {
	//         var sql = 'SELECT * FROM T_activity WHERE F_Caption LIKE ? ';
	//         str = '%' + str + '%';
	//         var params = [str];
	//         db.connection.query(sql,params,function (err,res) {
	//             if(err){
	//                 console.log(err);
	//                 return;
	//             }
	//             callback(res);
	//         })
	//     },
	Caption := c.Input().Get("str")
	activities, err := models.CheckLike(Caption)
	if err != nil {
		beego.Error(err)
	}
	var getall GetAll
	activityslice := make([]models.Activity, 0)
	activityslice2 := make([]models.Activity, 0)
	for _, v := range activities {
		detail2 := make([]models.Activity, 1)
		detail2[0].Id = v.Id
		detail2[0].Caption = v.Caption
		detail2[0].StartDate = v.StartDate
		detail2[0].EndDate = v.EndDate
		detail2[0].Location = v.Location
		if v.EndDate.After(time.Now()) {
			activityslice = append(activityslice, detail2...)
		} else { //查出过期的
			activityslice2 = append(activityslice2, detail2...)
		}
	}

	getall.Processing = activityslice
	getall.Completed = activityslice2

	c.Data["json"] = getall //map[string]interface{}{"userId": 1, "avatorUrl": "Filename"}
	c.ServeJSON()
}

type Details struct {
	Dates []datess `json:"dates"`
}

type datess struct {
	F_CheckDate string
}

// @Title post checkin person
// @Description post person
// @Param username query string true "The userId for person"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /check/details [post]
//活动详情
func (c *CheckController) Details() {
	// 	/router.get('/details',function (req,res,next) {
	//     // var activity_id = req.body.id;
	//      var activity_id = 31;
	//     check.getDetailById(activity_id,function (details) {
	//         check.getDateById(activity_id,function (dates) {
	//             for(var i = 0; i < details.length; i ++) {
	//                 details[i].F_CheckDate = moment(details[i].F_CheckDate).format('YYYY-MM-DD');
	//             }
	//             for(var i = 0; i < dates.length; i ++) {
	//                 dates[i].F_CheckDate = moment(dates[i].F_CheckDate).format('YYYY-MM-DD');
	//             }
	//             var response = {
	//                 'dateNum':dates.length,
	//                 'dates':dates,
	//                 'details':details
	//             }
	//             res.send(response);

	//             // res.send(details);
	//             // res.send(details)
	//         })

	//     })
	// });

	// /**
	//  * /check/details
	//  * activity_id
	//  */
	// router.post('/details',function (req,res,next) {
	//    var activity_id = req.body.id;
	//    //  var activity_id = 25;
	//    check.getDetailById(activity_id,function (details) {
	//        check.getDateById(activity_id,function (dates) {
	//            for(var i = 0; i < details.length; i ++) {
	//                details[i].F_CheckDate = moment(details[i].F_CheckDate).format('YYYY-MM-DD');
	//            }
	//            for(var i = 0; i < dates.length; i ++) {
	//                dates[i].F_CheckDate = moment(dates[i].F_CheckDate).format('YYYY-MM-DD');
	//            }
	//            var response = {
	//                'dateNum':dates.length,
	//                'dates':dates,
	//                'details':details
	//            }
	//            res.send(response);

	//            // res.send(details);
	//            // res.send(details)
	//        })

	//    })
	// });

	//     getDetailById: function getDetailById(activity_id,callback) {
	//         var sql = 'SELECT * FROM v_check_detail WHERE F_ActivityId = ?';
	//         var params = [activity_id];
	//         db.connection.query(sql,params,function (err,res) {
	//             if(err){
	//                 console.log(err);
	//                 return;
	//             }
	//             // console.log(res);
	//             callback(res);
	//         })
	//     },

	datess := make([]datess, 1)
	datess[0].F_CheckDate = "checkdate-details"
	var details Details
	details.Dates = datess
	c.Data["json"] = details //map[string]interface{}{"userId": 1, "avatorUrl": "Filename"}
	c.ServeJSON()
}

// @Title post checkin person
// @Description post person
// @Param username query string true "The userId for person"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /check/compare [post]
func (c *CheckController) Compare() {
	c.Data["json"] = map[string]interface{}{"confidence": 80, "avatorUrl": "Filename"}
	c.ServeJSON()
}

// @Title post checkin person
// @Description post person
// @Param userId query string true "The userId for person"
// @Param activityId query string true "The activityId of activity"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /check [post]
// 打卡记录写入数据库
func (c *CheckController) Check() {
	/**
	 * /check
	 * userid activityid photourl
	 */
	// router.post('/',function (req,res,next) {
	//    var userId = req.body.userId;
	//    var activityId = req.body.activityId;
	//    var photoUrl = req.body.photoUrl;
	//    check.check(userId,activityId,photoUrl,function (insert_res) {
	//        console.log(insert_res);
	//        var result = {
	//            code:1
	//        };
	//        res.send(result);
	//    })
	// });

	// 	check: function (user_id,activity_id,photo_url,callback) {
	//        var sql = 'INSERT INTO T_check(F_ID,F_ActivityId,F_UserID,F_CheckDate,F_PhotoUrl)VALUES(0,?,?,?,?)';
	//        var now = new Date();
	//        var params = [activity_id, user_id, now, photo_url];
	//        console.log(params);
	//        db.connection.query(sql,params,function (err,result) {
	//            if(err){
	//                console.log(err.message);
	//                return;
	//            }
	//            console.log("INSERT CHECK SUCCESS",result.insertId);
	//            callback(result);

	//        });
	//    },

	userid := c.Input().Get("userId")
	//pid转成64为
	UserId, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	activityid := c.Input().Get("activityId")
	ActivityId, err := strconv.ParseInt(activityid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	lat := c.Input().Get("lat")
	//string到float64
	Lat, err := strconv.ParseFloat(lat, 64)
	lng := c.Input().Get("lng")
	Lng, err := strconv.ParseFloat(lng, 64)

	// startdate := c.Input().Get("startDate")
	// beego.Info(startdate)
	const base_format = "2006-01-02"
	// StartDate, err := time.Parse(base_format, startdate)
	// if err != nil {
	// 	beego.Error(err)
	// }
	PhotoUrl := c.Input().Get("photoUrl")
	//根据userid取出user和avatorUrl
	useravatar, err := models.GetUserAvatorUrl(UserId)
	if err != nil {
		beego.Error(err)
	}
	var photo string
	if len(useravatar) != 0 {
		wxsite := beego.AppConfig.String("wxreqeustsite")
		photo = wxsite + useravatar[0].UserAvatar.AvatarUrl
	}
	year := c.Input().Get("year")
	month := c.Input().Get("month")
	if len(month) == 1 {
		month = "0" + month
	}
	day := c.Input().Get("day")
	if len(day) == 1 {
		day = "0" + day
	}
	SelectDate, err := time.Parse(base_format, year+"-"+month+"-"+day)
	if err != nil {
		beego.Error(err)
	}
	_, err = models.CheckCheck(ActivityId, UserId, Lat, Lng, PhotoUrl, SelectDate)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"code": 2, "message": PhotoUrl}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"code": 1, "avatorUrl": photo}
		c.ServeJSON()
	}
}

type CheckDate struct {
	Year  string `json:"year"`
	Month string `json:"month"`
	Day   string `json:"day"`
}

// @Title get checkin check
// @Description get check
// @Param userId query string true "The userId for check"
// @Param activityId query string true "The activityid for check"
// @Param year query string true "The year for check"
// @Param month query string true "The month for check"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /checkgetcheck [get]
// 取得当月的打卡记录
func (c *CheckController) CheckGetCheck() {
	userid := c.Input().Get("userId")
	//pid转成64为
	UserId, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	activityid := c.Input().Get("activityId")
	ActivityId, err := strconv.ParseInt(activityid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	const base_format = "2006-01-02"
	year := c.Input().Get("year")
	month := c.Input().Get("month")
	if len(month) == 1 {
		month = "0" + month
	}
	SelectMonth1, err := time.Parse(base_format, year+"-"+month+"-01")
	if err != nil {
		beego.Error(err)
	}
	SelectMonth2 := SelectMonth1.AddDate(0, 1, 0)

	data, err := models.CheckGetCheck(ActivityId, UserId, SelectMonth1, SelectMonth2)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"code": 2, "data": nil}
		c.ServeJSON()
	} else {
		// beego.Info(data)
		checkdataslice := make([]CheckDate, 0)
		for _, v := range data {
			checkdata := make([]CheckDate, 1)
			checkdata[0].Year = v.SelectDate.Format("2006")
			checkdata[0].Month = v.SelectDate.Format("01")
			checkdata[0].Day = v.SelectDate.Format("02")
			checkdataslice = append(checkdataslice, checkdata...)
		}
		c.Data["json"] = map[string]interface{}{"code": 1, "data": checkdataslice}
		c.ServeJSON()
	}
}

// activity/actInfo
// .infos[0].F_Caption
// <view>发起人： {{activity_info.infos[0].F_CreaterName}}</view>
// <view>活动开始时间： {{activity_info.infos[0].F_StartDate}}</view>
// <view>活动结束时间： {{activity_info.infos[0].F_EndDate}}</view>
// <view>活动地点： {{activity_info.infos[0].F_Location}}</view>
// <view>简介： {{activity_info.infos[0].F_Desc}}</view>

// 已有{{activity_info.nums}}人报名

// lng = res.data.infos[0].F_Lng;
// lat = res.data.infos[0].F_Lat;
// face = res.data.infos[0].F_IfFace;
// location = res.data.infos[0].F_IfLocation;

type ActInfos struct {
	Nums  int     `json:"nums"`
	Infos []infos `json:"infos"`
}

type infos struct {
	F_Caption     string
	F_CreaterName string
	F_StartDate   string
	F_EndDate     string
	F_Location    string
	F_Desc        string
	F_Lng         string
	F_Lat         string
	F_IfFace      string
	F_IfLocation  string
}

// @Title post checkin person
// @Description post person
// @Param username query string true "The userId for person"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /activity/actInfo [post]
func (c *CheckController) ActInfo() {
	infos := make([]infos, 1)
	// infos[0].F_CheckDate = "111"
	infos[0].F_Caption = "Caption-actinfo"
	infos[0].F_CreaterName = "创建者"
	infos[0].F_StartDate = "开始时间"
	infos[0].F_EndDate = "结束时间"
	infos[0].F_Location = "地点"
	infos[0].F_Desc = "描述"
	infos[0].F_Lng = "113.26436"
	infos[0].F_Lat = "23.12908"
	infos[0].F_IfFace = "ifface"
	infos[0].F_IfLocation = "iflocation"
	var actinfos ActInfos
	actinfos.Infos = infos
	actinfos.Nums = 2
	c.Data["json"] = actinfos
	c.ServeJSON()
}

// @Title post checkin person
// @Description post person
// @Param username query string true "The userId for person"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /activity/haveApply [post]
func (c *CheckController) HaveApply() {
	c.Data["json"] = map[string]interface{}{"code": true}
	c.ServeJSON()
}

type Detailsinfo struct {
	F_CheckPhotoUrl []string
	F_UserName      string
	Details         []details `json:"details"`
	// ActivityId      string `json:"activity_id"`
	// Date            string `json:"date"`
}

type details struct {
	F_UserName      string
	F_CheckPhotoUrl string
}

// @Title post checkin person
// @Description post person
// @Param username query string true "The userId for person"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /details/date [post]
func (c *CheckController) Date() {
	/**
	 * 查找某个活动某天的签到情况
	 */
	// router.post('/details/date',function (req,res,next) {
	//    var activity_id = req.body.activityId;
	//    var check_date = req.body.checkDate;
	//    // var activity_id = 31;
	//    // var check_date = "2018-01-17";
	//    console.log(activity_id,check_date);
	//    check.getDetailByIdAndDate(activity_id,check_date,function (details) {
	//        for(var i = 0; i < details.length; i ++) {
	//            details[i].F_CheckDate = moment(details[i].F_CheckDate).format('YYYY-MM-DD');
	//        }
	//        var response = {
	//            'details':details
	//        }
	//        console.log(check_date,details);
	//        res.send(response);
	//    })
	// });
	//    getDetailByIdAndDate: function getDetailByIdAndDate(activity_id,date,callback) {
	//        var sql = 'SELECT * FROM v_check_detail WHERE F_ActivityId = ? AND F_CheckDate = ?';
	//        var params = [activity_id,date];
	//        db.connection.query(sql,params,function (err,res) {
	//            if(err){
	//                console.log(err);
	//                return;
	//            }
	//            // console.log(res);
	//            callback(res);
	//        })
	//    },
	// list := []string{
	// 		{},
	// 		{},
	// 	}
	// 	catchimage := Catchimage{"SUCCESS", list}
	f_checkphotourl := []string{"http://img", "http://img2"}
	var detailsinfo Detailsinfo
	details1 := make([]details, 1)

	details1[0].F_UserName = "F_UserName"
	details1[0].F_CheckPhotoUrl = "http://details.img"
	detailsinfo.Details = details1
	detailsinfo.F_CheckPhotoUrl = f_checkphotourl
	// detailsinfo.ActivityId = "112-date"
	// detailsinfo.Date = "日期-Date"
	// detailsinfo.Details=
	c.Data["json"] = detailsinfo
	c.ServeJSON()
}

// @Title post checkin person
// @Description post person
// @Param username query string true "The userId for person"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /activity/apply [post]
//报名某个活动
func (c *CheckController) Apply() {
	c.Data["json"] = map[string]interface{}{"userId": 1, "avatorUrl": "Filename"}
	c.ServeJSON()
}

// getDateById: function getDateById(activity_id,callback) {
//     var sql = 'SELECT DISTINCT F_CheckDate FROM v_check_detail WHERE F_ActivityId = ?';
//     var params = [activity_id];
//     db.connection.query(sql,params,function (err,res) {
//         if(err){
//             console.log(err);
//             return;
//         }
//         console.log(res);
//         callback(res);
//     })
// },
// getCheckByDate: function getCheckByDate(check_date,callback) {
//     var sql = 'SELECT * FROM v_check_detail WHERE F_CheckDate = ?'
//     var params = [check_date];
//     db.connection.query(sql,params,function (err,res) {
//         if(err){
//             console.log(err);
//             return;
//         }
//         console.log(res);
//         callback(res);
//     })
// },
//********************统计**************
// @Title get mothcheckin
// @Description get monthcheck
// @Param projectid query string true "The projectid of check"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /monthchecksum [get]
func (c *CheckController) MonthCheckSum() {
	c.TplName = "check/check.tpl"
	c.Data["IsMonthCheck"] = true
	c.Data["ProjectId"] = c.Input().Get("projectid")
}

//后端分页的数据结构
type monthCheckTable struct {
	Rows  []map[int]interface{} `json:"rows"`
	Page  int64                 `json:"page"`
	Total int64                 `json:"total"` //string或int64都行！
}

// @Title get mothcheckin
// @Description get monthcheck
// @Param page query string false "The page of check"
// @Param limit query string false "The size of check"
// @Param activityId query string true "The activityid for check"
// @Param year query string true "The year for check"
// @Param month query string true "The month for check"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /monthcheck [get]
// 月度考勤统计
func (c *CheckController) MonthCheck() {
	var offset, limit1, page1 int
	var err error
	limit := c.Input().Get("limit")
	if limit == "" {
		limit1 = 10
	} else {
		limit1, err = strconv.Atoi(limit)
		if err != nil {
			beego.Error(err)
		}
	}
	page := c.Input().Get("page")
	if page == "" {
		// limit1 = 10
		page1 = 1
	} else {
		page1, err = strconv.Atoi(page)
		if err != nil {
			beego.Error(err)
		}
	}

	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}

	//当月天数
	const base_format = "2006-01-02"
	year := c.Input().Get("year")
	month := c.Input().Get("month")
	if len(month) == 1 {
		month = "0" + month
	}
	SelectMonth1, err := time.Parse(base_format, year+"-"+month+"-01")
	if err != nil {
		beego.Error(err)
	}
	SelectMonth2 := SelectMonth1.AddDate(0, 1, 0)
	//建立一个动态月日数组
	// beego.Info(SelectMonth2.Sub(SelectMonth1))
	days := SelectMonth2.Sub(SelectMonth1) / 24
	dayss := days.Hours()
	daysss := strconv.FormatFloat(dayss, 'f', -1, 64) //float64转string
	dayssss, err := strconv.Atoi(daysss)              //string转int
	// beego.Info(reflect.TypeOf(dayss)) //float64
	// beego.Info(dayssss)
	activityid := c.Input().Get("activityId")
	ActivityId, err := strconv.ParseInt(activityid, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	//查出check里所有用户
	// users, err := models.GetCheckUser(SelectMonth1, SelectMonth2, ActivityId)
	// if err != nil {
	// 	beego.Error(err)
	// }
	// for _, v := range users {
	// beego.Info(v.UserId)
	// beego.Info(v.SelectDate)
	// }
	users2, err := models.GetCheckUser2(SelectMonth1, SelectMonth2, ActivityId, limit1, offset)
	if err != nil {
		beego.Error(err)
	}
	// for _, v := range users2 {
	// beego.Info(v.User.Nickname)
	// beego.Info(v.Checkin.UserId)
	// }
	s := []map[int]interface{}{}
	// m := map[string]interface{}{"name": "John", "age": 10}
	// s = append(s, m)
	// b, err := json.Marshal(s)
	// if err != nil {
	// 	fmt.Println("json.Marshal failed:", err)
	// 	return
	// }
	//活动
	for _, v := range users2 {
		//这里直接用map[int]应该最好吧，可以排序，不用转string来回转？string能排序么
		var checkmap = make(map[int]interface{}, dayssss+2)
		// beego.Info(v.Checkin.UserId)
		for i := 0; i <= dayssss+1; i++ {
			// key = strconv.Itoa(i)
			if i == 0 {
				// key = "name"
				checkmap[0] = v.User.Nickname
			} else {
				// if len(key) == 1 {
				// 	key = "0" + key
				// 	checkmap[key] = ""
				checkmap[i] = ""
			}
		}
		//查出一个用户这个月的打卡记录
		data, err := models.CheckGetCheck(ActivityId, v.Checkin.UserId, SelectMonth1, SelectMonth2)
		// beego.Info(data)
		if err != nil {
			beego.Error(err)
		} else {
			for _, w := range data {
				day := w.SelectDate.Format("02")
				// beego.Info(day)
				dayint, err := strconv.Atoi(day)
				if err != nil {
					beego.Error(err)
				}
				checkmap[dayint] = "1"
			}
		}
		//排序map，按key
		var keys []int
		for k := range checkmap {
			// kint, err := strconv.Atoi(k)
			// if err != nil {
			// 	beego.Error(err)
			// }
			keys = append(keys, k)
		}
		sort.Ints(keys)
		// sort.Strings(keys)
		// To perform the opertion you want
		var checkmap2 = make(map[int]interface{}, dayssss+2)
		for _, k := range keys {
			// key = strconv.Itoa(k)
			// if key == "0" {
			// 	key = "name"
			// 	checkmap2[key] = v.UserId
			// } else if len(key) == 1 {
			// 	key = "0" + key
			// 	checkmap2[key] = checkmap[key]
			// } else {
			// 	checkmap2[key] = checkmap[key]
			// }
			// kstring := strconv.Itoa(k)
			checkmap2[k] = checkmap[k]
			// fmt.Println("Key:", k, "Value:", checkmap[k])
		}
		s = append(s, checkmap2)
	}
	b, err := json.Marshal(s)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"code": 2, "data": nil}
		c.ServeJSON()
	} else {
		// beego.Info(string(b))
		// c.Data["json"] = string(b)
		c.Ctx.WriteString(string(b))
		// c.Data["json"] = map[string]interface{}{"rows": s, "page": 1, "total": 10}
		// c.ServeJSON()
	}
}

// @Title get Signature
// @Description get Signature
// @Param signature query string true "The signature of wx"
// @Param timestamp query string true "The timestamp of wx"
// @Param nonce query string true "The nonce for wx"
// @Param echostr query string true "The echostr for wx"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /checksignature [get]
// 检验signature v1/checkin/checksignature
func (c *CheckController) CheckSignature() {
	// r.ParseForm()
	token := "3xxx"
	// EncodingAESKey := "MeiIgTIBnoGVjRMZkgrXlx3t7olpSKeieK331rdMSqf"
	// var signature string = strings.Join(r.Form["signature"], "")
	// var timestamp string = strings.Join(r.Form["timestamp"], "")
	// var nonce string = strings.Join(r.Form["nonce"], "")
	// var echostr string = strings.Join(r.Form["echostr"], "")
	signature := c.Input().Get("signature")
	timestamp := c.Input().Get("timestamp")
	// beego.Info(timestamp)
	nonce := c.Input().Get("nonce")
	// beego.Info(nonce)
	echostr := c.Input().Get("echostr")
	// beego.Info(echostr)

	tmps := []string{token, timestamp, nonce}
	sort.Strings(tmps)
	tmpStr := tmps[0] + tmps[1] + tmps[2]
	tmp := str2sha1(tmpStr)
	if tmp == signature {
		// beego.Info(tmp)
		// beego.Info(signature)
		// fmt.Fprintf(w, echostr)
		c.Ctx.WriteString(echostr)
		// c.Data["json"] = echostr
		// c.ServeJSON()
	} else {
		// beego.Info(tmp)
		// beego.Info(signature)
		c.Data["json"] = map[string]interface{}{"tmp": tmp, "signature": signature}
		c.ServeJSON()
	}
}

// @Title post subscribemessage
// @Description post subscribemessage
// @Param tmplIds query string true "The tmplIds of wx"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /subscribemessage [post]
// 检验signature v1/checkin/subscribemessage
// 存储用户订阅subscribemessage
func (c *CheckController) SubscribeMessage() {
	openID := c.GetSession("openID")
	if openID == nil {
		beego.Info("openid为空")
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "openid为空"}
		c.ServeJSON()
	} else {
		tmplIds := c.Input().Get("tmplIds")
		//存储openid对应订阅消息模板id
		id, err := models.AddWxSubscribeMessage(openID.(string), tmplIds)
		if err != nil {
			beego.Error(err)
			c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": 0}
			c.ServeJSON()
		} else {
			c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "data": id}
			c.ServeJSON()
		}
	}
}

// @Title post Message
// @Description post Message
// @Param app_version query string true "The app_version of wx"
// @Param template_id query string true "The template_id of Message"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /sendmessage [post]
// 检验signature v1/checkin/sendmessage
// 点击发送订阅消息
func (c *CheckController) SendMessage() {
	// POST https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=ACCESS_TOKEN
	app_version := c.Input().Get("app_version")
	accessToken, _, _, err := utils.GetAccessToken(app_version)

	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": err}
		c.ServeJSON()
	}
	// template_id := c.Input().Get("template_id")
	template_id := "c8Dped7TztDiMJ2bzHMu4G3nikn-mSOHmQEh_7aTlLo"
	// openid := "opl4B5YvCVXaatKF6VGUSbohMlWQ" //"opl4B5auOWiSxXo5-RB3NTitcHxk"
	//查出所有openid
	openids, err := models.GetOpenIDs()
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "openid为空！"}
		c.ServeJSON()
	} else {
		for _, i := range openids {
			errcode, errmsg, err := utils.SendMessage(accessToken, i.OpenID, template_id)
			if err != nil {
				beego.Error(err)
				c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": err}
				c.ServeJSON()
			} else if errcode == 40003 {
				c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "touser字段openid为空或者不正确", "errmsg": errmsg}
				c.ServeJSON()
			} else if errcode == 40037 {
				c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "订阅模板id为空不正确", "errmsg": errmsg}
				c.ServeJSON()
			} else if errcode == 43101 {
				c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "用户拒绝接受消息，如果用户之前曾经订阅过，则表示用户取消了订阅关系", "errmsg": errmsg}
				c.ServeJSON()
			} else if errcode == 47003 {
				c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "模板参数不准确，可能为空或者不满足规则，errmsg会提示具体是哪个字段出错", "errmsg": errmsg}
				c.ServeJSON()
			} else if errcode == 41030 {
				c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "page路径不正确，需要保证在现网版本小程序中存在，与app.json保持一致", "errmsg": errmsg}
				c.ServeJSON()
			} else {
				c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "data": "成功！"}
				c.ServeJSON()
			}
		}
	}
	// 40003	touser字段openid为空或者不正确
	// 40037	订阅模板id为空不正确
	// 43101	用户拒绝接受消息，如果用户之前曾经订阅过，则表示用户取消了订阅关系
	// 47003	模板参数不准确，可能为空或者不满足规则，errmsg会提示具体是哪个字段出错
	// 41030	page路径不正确，需要保证在现网版本小程序中存在，与app.json保持一致
	// 打卡类型
	// {{phrase1.DATA}}
	// 打卡时间
	// {{date2.DATA}}
	// 打卡地点
	// {{thing3.DATA}}
}

func SendMessage() {
	app_version := "4"
	accessToken, _, _, err := utils.GetAccessToken(app_version)
	if err != nil {
		beego.Error(err)
	}
	// beego.Info(accessToken)
	// template_id := c.Input().Get("template_id")
	template_id := "c8Dped7TztDiMJ2bzHMu4G3nikn-mSOHmQEh_7aTlLo"
	// openid := "opl4B5YvCVXaatKF6VGUSbohMlWQ" //"opl4B5auOWiSxXo5-RB3NTitcHxk"
	//查出所有openid
	openids, err := models.GetOpenIDs()
	if err != nil {
		beego.Error(err)
	} else {
		for _, i := range openids {
			errcode, errmsg, err := utils.SendMessage(accessToken, i.OpenID, template_id)
			// beego.Info(errcode)
			// beego.Info(errmsg)
			if err != nil {
				beego.Error(err)
			} else if errcode == 40003 {
				beego.Error("touser字段openid为空或者不正确" + errmsg)
			} else if errcode == 40037 {
				beego.Error("订阅模板id为空不正确" + errmsg)
			} else if errcode == 43101 {
				beego.Error("用户拒绝接受消息，如果用户之前曾经订阅过，则表示用户取消了订阅关系" + errmsg)
			} else if errcode == 47003 {
				beego.Error("模板参数不准确，可能为空或者不满足规则，errmsg会提示具体是哪个字段出错" + errmsg)
			} else if errcode == 41030 {
				beego.Error("page路径不正确，需要保证在现网版本小程序中存在，与app" + errmsg)
			} else {
				beego.Error("发送成功" + errmsg)
				// beego.Info(errmsg)
			}
		}
	}
}

// {
// 	"touser":"opl4B5YvCVXaatKF6VGUSbohMlWQ",
// "template_id":"c8Dped7TztDiMJ2bzHMu4G3nikn-mSOHmQEh_7aTlLo",
// "page":"index",
// "data":{
// 	"phrase1":{"value":"设代打卡"},
// 	"date2":{"value":"8:30"},
// 	"thing3":{"value":"顺德、南沙、东莞、深圳"}
// 	}
// }
// {
//   "touser": "OPENID",
//   "template_id": "TEMPLATE_ID",
//   "page": "index",
//   "miniprogram_state":"developer",
//   "lang":"zh_CN",
//   "data": {
//       "number01": {
//           "value": "339208499"
//       },
//       "date01": {
//           "value": "2015年01月05日"
//       },
//       "site01": {
//           "value": "TIT创意园"
//       } ,
//       "site02": {
//           "value": "广州市新港中路397号"
//       }
//   }
// }

// type Request struct {
// 	ToUserName   string
// 	FromUserName string
// 	CreateTime   time.Duration
// 	MsgType      string
// 	Content      string
// 	MsgId        int
// }

// type Response struct {
// 	ToUserName   string `xml:"xml>ToUserName"`
// 	FromUserName string `xml:"xml>FromUserName"`
// 	CreateTime   string `xml:"xml>CreateTime"`
// 	MsgType      string `xml:"xml>MsgType"`
// 	Content      string `xml:"xml>Content"`
// 	MsgId        int    `xml:"xml>MsgId"`
// }

func str2sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

// func action(w http.ResponseWriter, r *http.Request) {
// 	postedMsg, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	r.Body.Close()
// 	v := Request{}
// 	xml.Unmarshal(postedMsg, &v)
// 	if v.MsgType == "text" {
// 		v := Request{v.ToUserName, v.FromUserName, v.CreateTime, v.MsgType, v.Content, v.MsgId}
// 		output, err := xml.MarshalIndent(v, " ", " ")
// 		if err != nil {
// 			fmt.Printf("error:%v\n", err)
// 		}
// 		fmt.Fprintf(w, string(output))
// 	} else if v.MsgType == "event" {
// 		Content := `"欢迎关注
//                                 我的微信"`
// 		v := Request{v.ToUserName, v.FromUserName, v.CreateTime, v.MsgType, Content, v.MsgId}
// 		output, err := xml.MarshalIndent(v, " ", " ")
// 		if err != nil {
// 			fmt.Printf("error:%v\n", err)
// 		}
// 		fmt.Fprintf(w, string(output))
// 	}
// }
