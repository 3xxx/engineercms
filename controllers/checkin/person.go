package checkin

import (
	beego "github.com/beego/beego/v2"
)

// CMSCHECKIN API
type PersonController struct {
	beego.Controller
}

func (c *PersonController) Person() {
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
	c.Data["json"] = "ok"
	c.ServeJSON()
}
