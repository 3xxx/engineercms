package models

import (
	// "database/sql"
	// beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
	// _ "github.com/go-sql-driver/mysql"
	// "github.com/go-xorm/xorm"
	// _ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
	// "strconv"
	// "strings"
	// "fmt"
	// "log"
	// "os"
	"time"
)

// var engine *xorm.Engine

// -- ----------------------------
// -- Table structure for t_activity
// -- ----------------------------
// DROP TABLE IF EXISTS `t_activity`;
// CREATE TABLE `t_activity`  (
//   `F_ID` int(11) NOT NULL AUTO_INCREMENT,
//   `F_CreaterId` int(11) NOT NULL,
//   `F_Caption` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
//   `F_Desc` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
//   `F_Location` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
//   `F_Lat` float NULL DEFAULT NULL,
//   `F_Lng` float NULL DEFAULT NULL,
//   `F_StartDate` date NULL DEFAULT NULL,
//   `F_EndDate` date NULL DEFAULT NULL,
//   `F_IfFace` int(11) NULL DEFAULT NULL,
//   `F_IfLocation` int(11) NULL DEFAULT NULL,
//   `F_IfPhoto` int(11) NULL DEFAULT NULL,
//   PRIMARY KEY (`F_ID`) USING BTREE,
//   INDEX `F_CreaterId`(`F_CreaterId`) USING BTREE,
//   CONSTRAINT `t_activity_ibfk_1` FOREIGN KEY (`F_CreaterId`) REFERENCES `t_user` (`F_ID`) ON DELETE RESTRICT ON UPDATE RESTRICT
// ) ENGINE = InnoDB AUTO_INCREMENT = 40 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

// -- ----------------------------
// -- Table structure for t_apply
// -- ----------------------------
// DROP TABLE IF EXISTS `t_apply`;
// CREATE TABLE `t_apply`  (
//   `F_ID` int(11) NOT NULL AUTO_INCREMENT,
//   `F_ApplyerId` int(11) NOT NULL,
//   `F_ActivityId` int(11) NOT NULL,
//   `F_ApplyDate` date NULL DEFAULT NULL,
//   PRIMARY KEY (`F_ID`) USING BTREE,
//   INDEX `F_ApplyerId`(`F_ApplyerId`) USING BTREE,
//   INDEX `F_ActivityId`(`F_ActivityId`) USING BTREE,
//   CONSTRAINT `t_apply_ibfk_1` FOREIGN KEY (`F_ApplyerId`) REFERENCES `t_user` (`F_ID`) ON DELETE RESTRICT ON UPDATE RESTRICT,
//   CONSTRAINT `t_apply_ibfk_2` FOREIGN KEY (`F_ActivityId`) REFERENCES `t_activity` (`F_ID`) ON DELETE RESTRICT ON UPDATE RESTRICT
// ) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

// -- ----------------------------
// -- Table structure for t_check
// -- ----------------------------
// DROP TABLE IF EXISTS `t_check`;
// CREATE TABLE `t_check`  (
//   `F_ID` int(11) NOT NULL AUTO_INCREMENT,
//   `F_ActivityId` int(11) NULL DEFAULT NULL,
//   `F_UserID` int(11) NULL DEFAULT NULL,
//   `F_CheckDate` date NULL DEFAULT NULL,
//   `F_PhotoUrl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
//   PRIMARY KEY (`F_ID`) USING BTREE,
//   INDEX `F_ActivityId`(`F_ActivityId`) USING BTREE,
//   INDEX `F_UserID`(`F_UserID`) USING BTREE,
//   CONSTRAINT `t_check_ibfk_1` FOREIGN KEY (`F_ActivityId`) REFERENCES `t_activity` (`F_ID`) ON DELETE RESTRICT ON UPDATE RESTRICT,
//   CONSTRAINT `t_check_ibfk_2` FOREIGN KEY (`F_UserID`) REFERENCES `t_user` (`F_ID`) ON DELETE RESTRICT ON UPDATE RESTRICT
// ) ENGINE = InnoDB AUTO_INCREMENT = 30 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

// -- ----------------------------
// -- Table structure for t_user
// -- ----------------------------
// DROP TABLE IF EXISTS `t_user`;
// CREATE TABLE `t_user`  (
//   `F_ID` int(11) NOT NULL AUTO_INCREMENT,
//   `F_Name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
//   `F_PhotoUrl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
//   PRIMARY KEY (`F_ID`) USING BTREE
// ) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

// -- ----------------------------
// -- View structure for v_activity_info
// -- ----------------------------
// DROP VIEW IF EXISTS `v_activity_info`;
// CREATE ALGORITHM = UNDEFINED DEFINER = `root`@`localhost` SQL SECURITY DEFINER VIEW `v_activity_info` AS select `t_activity`.`F_ID` AS `F_ActivityId`,`t_user`.`F_Name` AS `F_CreaterName`,`t_activity`.`F_Caption` AS `F_Caption`,`t_activity`.`F_Desc` AS `F_Desc`,`t_activity`.`F_Location` AS `F_Location`,`t_activity`.`F_Lat` AS `F_Lat`,`t_activity`.`F_Lng` AS `F_Lng`,`t_activity`.`F_StartDate` AS `F_StartDate`,`t_activity`.`F_EndDate` AS `F_EndDate`,`t_activity`.`F_IfFace` AS `F_IfFace`,`t_activity`.`F_IfLocation` AS `F_IfLocation`,`t_activity`.`F_IfPhoto` AS `F_IfPhoto` from (`t_activity` join `t_user` on((`t_activity`.`F_CreaterId` = `t_user`.`F_ID`)));

// -- ----------------------------
// -- View structure for v_apply_activity
// -- ----------------------------
// DROP VIEW IF EXISTS `v_apply_activity`;
// CREATE ALGORITHM = UNDEFINED DEFINER = `root`@`localhost` SQL SECURITY DEFINER VIEW `v_apply_activity` AS select `t_activity`.`F_ID` AS `F_ID`,`t_activity`.`F_CreaterId` AS `F_CreaterId`,`t_activity`.`F_Caption` AS `F_Caption`,`t_activity`.`F_Desc` AS `F_Desc`,`t_activity`.`F_Lat` AS `F_Lat`,`t_activity`.`F_Lng` AS `F_Lng`,`t_activity`.`F_StartDate` AS `F_StartDate`,`t_activity`.`F_EndDate` AS `F_EndDate`,`t_activity`.`F_IfFace` AS `F_IfFace`,`t_activity`.`F_IfLocation` AS `F_IfLocation`,`t_activity`.`F_IfPhoto` AS `F_IfPhoto`,`t_apply`.`F_ApplyDate` AS `F_ApplyDate`,`t_apply`.`F_ApplyerId` AS `F_ApplyerId`,`t_activity`.`F_Location` AS `F_Location` from (`t_apply` join `t_activity` on((`t_apply`.`F_ActivityId` = `t_activity`.`F_ID`)));

// -- ----------------------------
// -- View structure for v_apply_info
// -- ----------------------------
// DROP VIEW IF EXISTS `v_apply_info`;
// CREATE ALGORITHM = UNDEFINED DEFINER = `root`@`localhost` SQL SECURITY DEFINER VIEW `v_apply_info` AS select `t_activity`.`F_ID` AS `F_ID`,count(0) AS `F_ApplyNum`,`t_activity`.`F_Caption` AS `F_Caption`,`t_apply`.`F_ApplyerId` AS `F_ApplyerId` from (`t_activity` join `t_apply` on((`t_apply`.`F_ActivityId` = `t_activity`.`F_ID`))) group by `t_activity`.`F_ID`;

// -- ----------------------------
// -- View structure for v_check_detail
// -- ----------------------------
// DROP VIEW IF EXISTS `v_check_detail`;
// CREATE ALGORITHM = UNDEFINED DEFINER = `root`@`localhost` SQL SECURITY DEFINER VIEW `v_check_detail` AS select `t_check`.`F_ActivityId` AS `F_ActivityId`,`t_check`.`F_UserID` AS `F_UserId`,`t_check`.`F_PhotoUrl` AS `F_CheckPhotoUrl`,`t_user`.`F_Name` AS `F_UserName`,`t_check`.`F_CheckDate` AS `F_CheckDate` from (`t_user` join `t_check` on((`t_check`.`F_UserID` = `t_user`.`F_ID`)));

// SET FOREIGN_KEY_CHECKS = 1;

type Activity struct {
	Id         int64  `json:"F_ID"`
	CreaterId  int64  `json:"userid"`
	ProjectId  int64  `json:"projectid"`
	Caption    string `json:"F_Caption"`
	Desc       string
	Location   string `json:"F_Location"`
	Lat        float64
	Lng        float64
	StartDate  time.Time `json:"F_StartDate",orm:"auto_now_add;type(date)"`
	EndDate    time.Time `json:"F_EndDate",orm:"auto_now_add;type(date)"`
	IfFace     bool
	IfLocation bool
	IfPhoto    bool
}

//报名活动
type Apply struct {
	Id         int64
	ApplyerId  int64
	ActivityId int64
	ApplyDate  time.Time `orm:"auto_now_add;type(date)"`
}

//打卡记录_作废
// type Check struct {
// 	Id         int64
// 	ActivityId int64
// 	UserId     int64
// 	CheckDate  time.Time `orm:"auto_now_add;type(date)"`
// 	PhotoUrl   string
// 	Lat        float64
// 	Lng        float64
// 	SelectDate time.Time `orm:"auto_now_add;type(date)"`
// }

//打卡记录
type Checkin struct {
	Id         int64
	ActivityId int64
	UserId     int64
	CheckDate  time.Time `orm:"auto_now_add;type(date)"`
	PhotoUrl   string
	Lat        float64
	Lng        float64
	SelectDate time.Time `orm:"auto_now_add;type(date)"`
}

type OpenidTmplId struct {
	Id     int64
	OpenID string
	TmplId string
}

func init() {
	orm.RegisterModel(new(Activity), new(Apply), new(Checkin), new(OpenidTmplId))
}

// activity
// create: function create (a,callback) {
//     var sql = 'INSERT INTO T_activity(F_ID,F_CreaterId,F_Caption,F_Desc,F_Location,F_Lat,F_Lng,F_StartDate' +
//         ',F_EndDate,F_IfFace,F_IfPhoto,F_IfLocation)' +
//         ' VALUES(0,?,?,?,?,?,?,?,?,?,?,?)';
//     var params=[a.F_CreaterId,a.F_Caption,a.F_Desc,a.F_Location,a.F_Lat,a.F_Lng,a.F_StartDate,a.F_EndDate
//         ,a.F_IfFace,a.F_IfPhoto,a.F_IfLocation];
//     db.connection.query(sql,params,function (err,res) {
//         if(err) {
//             console.log('INSERT ACTIVITY:', err.message);
//             return;
//         }
//         callback(res);
//         console.log('INSERT ACTIVITY SUCCESS');
//     });
// },

func CheckCreate(CreaterId, projectid int64, Caption, Desc, Location string, Lat, Lng float64, StartDate, EndDate time.Time, IfFace, IfPhoto, IfLocation bool) (id int64, err error) {
	o := orm.NewOrm()
	// var category AdminCategory
	// if pid == "" {
	// 	category := &AdminCategory{
	// 		ParentId: 0,
	// 		Title:    title,
	// 		Code:     code,
	// 		Grade:    grade,
	// 		Created:  time.Now(),
	// 		Updated:  time.Now(),
	// 	}
	// 	id, err = o.Insert(category)
	// 	if err != nil {
	// 		return 0, err
	// 	}
	// } else {

	activity := &Activity{
		CreaterId:  CreaterId,
		ProjectId:  projectid,
		Caption:    Caption,
		Desc:       Desc,
		Location:   Location,
		Lat:        Lat,
		Lng:        Lng,
		StartDate:  StartDate,
		EndDate:    EndDate,
		IfFace:     IfFace,
		IfLocation: IfLocation,
		IfPhoto:    IfPhoto,
	}
	id, err = o.Insert(activity)
	// if err != nil {
	return id, err
	// }
	// return id, nil
}

// getall: function (user_id,callback) {
//     var sql = 'SELECT * FROM v_apply_activity WHERE F_ApplyerId=? ';
//     var params = [user_id];
//     db.connection.query(sql,params,function (err,res) {
//         if(err){
//             console.log(err);
//             return;
//         }
//         callback(res);
//     })
// },
func GetAll(projectid int64) (activity []*Activity, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Activity")
	if projectid != 0 {
		_, err = qs.Filter("ProjectId", projectid).All(&activity)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = qs.All(&activity)
		if err != nil {
			return nil, err
		}
	}
	return activity, err
}

// getallLike:function (str,callback) {
//     var sql = 'SELECT * FROM T_activity WHERE F_Caption LIKE ? ';
//     str = '%' + str + '%';
//     var params = [str];
//     db.connection.query(sql,params,function (err,res) {
//         if(err){
//             console.log(err);
//             return;
//         }
//         callback(res);
//     })
// },
func CheckLike(name string) (activity []*Activity, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Activity")
	_, err = qs.Filter("Caption__contains", name).All(&activity)
	if err != nil {
		return nil, err
	}
	return activity, err
}

//打卡记录写入数据库
func CheckCheck(ActivityId, UserId int64, Lat, Lng float64, PhotoUrl string, SelectDate time.Time) (id int64, err error) {
	o := orm.NewOrm()
	//查询数据库中有无打卡
	var check1 Checkin
	//判断是否有重名
	err = o.QueryTable("checkin").Filter("ActivityId", ActivityId).Filter("UserId", UserId).Filter("SelectDate", SelectDate).One(&check1, "Id")
	if err == orm.ErrNoRows {
		// 没有找到记录
		check := &Checkin{
			ActivityId: ActivityId,
			UserId:     UserId,
			CheckDate:  time.Now(),
			PhotoUrl:   PhotoUrl,
			Lat:        Lat,
			Lng:        Lng,
			SelectDate: SelectDate,
		}
		id, err = o.Insert(check)
		if err != nil {
			return id, err
		}
	}
	return id, nil
}

// type checkgetcheck struct{
// 	Id int64
// 	Uid int64
// 	activityid int64
// }

// 按月查询打卡记录
func CheckGetCheck(ActivityId, UserId int64, SelectMonth1, SelectMonth2 time.Time) (check []*Checkin, err error) {
	cond := orm.NewCondition()
	cond1 := cond.And("SelectDate__gte", SelectMonth1).And("SelectDate__lte", SelectMonth2)
	o := orm.NewOrm()
	qs := o.QueryTable("checkin")
	qs = qs.SetCond(cond1)
	_, err = qs.Filter("ActivityId", ActivityId).Filter("UserId", UserId).All(&check)
	if err != nil {
		return nil, err
	}
	return check, nil
}

//按月统计**************这个没用
//查询check里的用户Distinct("check.user_id").Limit(limit, offset)
func GetCheckUser(selectmonth1, selectmonth2 time.Time, activityid int64) ([]*Checkin, error) {
	checks := make([]*Checkin, 0) //分配的空间是0，容量是8个字段
	return checks, engine.Where("select_date >= ? AND select_date <=? AND activity_id", selectmonth1, selectmonth2, activityid).Distinct("user_id", "select_date").Find(&checks)
	// return users, nil Select("check.* as name")这个什么鬼语法？
}

// GetPkgInfosByIDs returns a list of package info by given IDs.
// func GetPkgInfosByIDs(ids []int64) ([]*PkgInfo, error) {
// 	if len(ids) == 0 {
// 		return []*PkgInfo{}, nil
// 	}
// 	pkgInfos := make([]*PkgInfo, 0, len(ids))
// 	return pkgInfos, x.Where("id > 0").In("id", base.Int64sToStrings(ids)).Find(&pkgInfos)
// }
type CheckUser struct {
	Checkin `xorm:"extends"`
	User    `xorm:"extends"`
}

//按月统计**************
func GetCheckUser2(selectmonth1, selectmonth2 time.Time, activityid int64, limit, offset int) ([]*CheckUser, error) {
	// var users = make([]User, 0)
	users := make([]*CheckUser, 0)
	// err = engine.Distinct("user_id").AllCols().Find(&checks)
	// err := engine.Sql("select * from userinfo, userdetail where userinfo.detail_id = userdetail.id").Find(&users)
	// return users, engine.Sql("SELECT DISTINCT checkin.user_id FROM checkin INNER JOIN user ON checkin.user_id = user.id LIMIT ? OFFSET ?", limit, offset).Find(&users)
	return users, engine.Table("checkin").Join("INNER", "user", "checkin.user_id = user.id").Where("checkin.select_date >= ? AND checkin.select_date <= ? AND checkin.activity_id = ?", selectmonth1, selectmonth2, activityid).Distinct("checkin.user_id", "user.nickname").Limit(limit, offset).Find(&users)
	// return users, engine.Sql("SELECT * from checkin").Find(&users)
	//distinct和limit不能和sql组合
}

//存储用户订阅消息模板id和用户openid
func AddWxSubscribeMessage(openid, tmplid string) (id int64, err error) {
	o := orm.NewOrm()
	//查询数据库中有无打卡
	var opentmpid OpenidTmplId
	//判断是否有重名
	err = o.QueryTable("OpenidTmplId").Filter("OpenId", openid).Filter("TmplId", tmplid).One(&opentmpid, "Id")
	if err == orm.ErrNoRows {
		// 没有找到记录
		opentmpid := &OpenidTmplId{
			OpenID: openid,
			TmplId: tmplid,
		}
		id, err = o.Insert(opentmpid)
		if err != nil {
			return id, err
		}
	}
	return id, nil
}

// UserId Id
// getInfo:function (activityId,callback) {
//     var sql = 'SELECT * FROM v_activity_info WHERE F_ActivityId = ?';
//     var params = [activityId];
//     db.connection.query(sql,params,function (err,res) {
//         if(err){
//             console.log(err.message);
//             return;
//         }
//         callback(res);

//     })
// },
// haveApplyed:function (activity_id,user_id,callback) {
//     var sql = 'SELECT * FROM t_apply WHERE F_ApplyerId = ? AND F_ActivityId = ?';
//     var params = [user_id,activity_id];
//     db.connection.query(sql,params,function (err,res) {
//         if(err){
//             console.log(err);
//             return;
//         }
//         console.log(res);
//         callback(res);
//     })
// }

// apply
/**
 * TODO
 * 用户对某个签到活动进行报名
 * @param user
 */
// apply: function apply(user_id,activity_id,callback) {
//     var sql = 'INSERT INTO t_apply (F_ID,F_ApplyerId,F_ActivityId,F_ApplyDate)VALUES(0,?,?,?)';
//     var now = moment().format('YYYY-MM-DD');
//     var params = [user_id,activity_id,now];
//     db.connection.query(sql,params,function (err,res) {
//         var re = true;
//         if(err){
//             console.log(err.message);
//             return;
//         }
//         callback(re);
//     })
// },

// getApplyNumByActId: function applyNum(activity_id,callback) {
//     var sql = 'SELECT COUNT(*) AS NUMS FROM t_apply WHERE F_ActivityId = ? GROUP BY F_ActivityId';
//     var params = [activity_id];
//     db.connection.query(sql,params,function (err,res) {
//         if(err){
//             console.log(err);
//             return;
//         }
//         // console.log(res[0].NUMS);
//         if(res[0]==null){
//             callback(0)
//         }else{
//             callback(res[0].NUMS);
//         }
//     });
// }

//check

// check: function (user_id,activity_id,photo_url,callback) {
//     var sql = 'INSERT INTO T_check(F_ID,F_ActivityId,F_UserID,F_CheckDate,F_PhotoUrl)VALUES(0,?,?,?,?)';
//     var now = new Date();
//     var params = [activity_id, user_id, now, photo_url];
//     console.log(params);
//     db.connection.query(sql,params,function (err,result) {
//         if(err){
//             console.log(err.message);
//             return;
//         }
//         console.log("INSERT CHECK SUCCESS",result.insertId);
//         callback(result);

//     });
// },

// getDetailById: function getDetailById(activity_id,callback) {
//     var sql = 'SELECT * FROM v_check_detail WHERE F_ActivityId = ?';
//     var params = [activity_id];
//     db.connection.query(sql,params,function (err,res) {
//         if(err){
//             console.log(err);
//             return;
//         }
//         // console.log(res);
//         callback(res);
//     })
// },

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
// getDetailByIdAndDate: function getDetailByIdAndDate(activity_id,date,callback) {
//     var sql = 'SELECT * FROM v_check_detail WHERE F_ActivityId = ? AND F_CheckDate = ?';
//     var params = [activity_id,date];
//     db.connection.query(sql,params,function (err,res) {
//         if(err){
//             console.log(err);
//             return;
//         }
//         callback(res);
//     })
// },

//person
/**
 * 第一次使用的时候创建用户，通过检索是否在数据库中有记录
 * @param user
 */
// create: function create(user,callback) {
//     var sql = 'INSERT INTO T_user(F_ID,F_Name)VALUES(0,?)';
//     var params = [user];
//     db.connection.query(sql,params,function (err,result) {
//         if(err){
//             console.log(err.message);
//             return;
//         }
//         console.log("INSERT SUCCESS",result.insertId);
//         callback(result);

//     });
// },
/**
 * 通过用户名查找用户
 * @param name
 */
// getUserByName: function getUserByName(name,callback) {
//     var sql = 'SELECT * FROM T_user WHERE F_Name=?';
//     var params = [name];
//     db.connection.query(sql,params,function (err,result) {
//        if(err){
//            console.log(err.message);
//            return;
//        }
//        callback(result);
//     });
// },

/**
 * 传入用户名,并将用户的信息更新
 * @param name
 */
// updateUserById: function updateUserByName(user_id,avator_url,callback) {
//     var sql = 'UPDATE T_user SET F_PhotoUrl = ? WHERE F_ID = ?';
//     var params = [avator_url,user_id];
//     db.connection.query(sql,params,function (err,result) {
//         if(err){
//             console.log(err.message);
//             return;
//         }
//         callback(result);
//     })
// }
