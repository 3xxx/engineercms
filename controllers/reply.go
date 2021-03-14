package controllers

import (
	"encoding/json"
	"github.com/3xxx/engineercms/controllers/utils"
	"github.com/3xxx/engineercms/models"
	"github.com/astaxie/beego"
	"net/http"
	"os"
	"strconv"
)

type ReplyController struct {
	beego.Controller
}

// @Title post wx release by articleId
// @Description post release by articleId
// @Param content query string  true "The content of release"
// @Param app_version query string  false "the app_version of wx"
// @Param avatar query string  false "The avatar of release"
// @Param username query string  false "The username of release"
// @Param publish_time query string  false "The time of release"
// @Success 200 {object} models.AddTopicReply
// @Failure 400 Invalid page supplied
// @Failure 404 article not found
// @router /addwxrelease/:id [post]
//添加文章评论
func (c *ReplyController) AddWxRelease() {
	//content去验证
	app_version := c.Input().Get("app_version")
	accessToken, _, _, err := utils.GetAccessToken(app_version)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": err}
		c.ServeJSON()
	}
	// beego.Info(accessToken)
	content := c.Input().Get("content")
	errcode, errmsg, err := utils.MsgSecCheck(accessToken, content)
	if err != nil {
		beego.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": err}
		c.ServeJSON()
	} else if errcode != 87014 {
		id := c.Ctx.Input.Param(":id")
		//id转成64为
		idNum, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			beego.Error(err)
		}
		// content := c.Input().Get("content")
		avatar := c.Input().Get("avatar")
		username := c.Input().Get("username")
		created := c.Input().Get("publish_time")
		var openID string
		openid := c.GetSession("openID")
		if openid == nil {

		} else {
			openID = openid.(string)
		}

		releaseid, err := models.AddTopicReply(idNum, openID, content, avatar, username, created)
		if err != nil {
			beego.Error(err)
			c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": err}
		} else {
			c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "releaseid": releaseid}
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": errmsg}
		c.ServeJSON()
	}

}

// @Title delete wx release by releaseid
// @Description delete release by releaseid
// @Param id query string  true "The id of release"
// @Success 200 {object} models.DeleteTopicReply
// @Failure 400 Invalid page supplied
// @Failure 404 article not found
// @router /deletewxrelease/:id [post]
//删除文章评论
func (c *ReplyController) DeleteWxRelease() {
	id := c.Ctx.Input.Param(":id")
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	err = models.DeleteTopicReply(idNum)
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "releaseid": id}
		c.ServeJSON()
	}
}

// @Title post wx like by articleId
// @Description post like by articleId
// @Param id query string  true "The id of article"
// @Success 200 {object} models.AddTopicLike
// @Failure 400 Invalid page supplied
// @Failure 404 article not found
// @router /addwxlike/:id [post]
//添加文章点赞
func (c *ReplyController) AddWxLike() {
	// tid := c.Input().Get("tid") //tid是文章id
	id := c.Ctx.Input.Param(":id")
	//id转成64为
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		beego.Error(err)
	}
	liked := c.Input().Get("liked")
	//根据用户openid，获取用户名，如果没有openid，则用用户昵称
	JSCODE := c.Input().Get("code")
	// beego.Info(JSCODE)
	APPID := beego.AppConfig.String("wxAPPID2")
	SECRET := beego.AppConfig.String("wxSECRET2")
	app_version := c.Input().Get("app_version")
	if app_version == "3" {
		APPID = beego.AppConfig.String("wxAPPID3")
		SECRET = beego.AppConfig.String("wxSECRET3")
	}
	// APPID := "wx05b480781ac8f4c8"                //这里一定要修改啊
	// SECRET := "0ddfd84afc74f3bc5b5db373a4090dc5" //这里一定要修改啊
	requestUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=" + APPID + "&secret=" + SECRET + "&js_code=" + JSCODE + "&grant_type=authorization_code"
	resp, err := http.Get(requestUrl)
	if err != nil {
		beego.Error(err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		beego.Error(err)
		// return
	}
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		beego.Error(err)
		// return
	}
	// beego.Info(data)
	var openID string
	// var sessionKey string
	if _, ok := data["session_key"]; !ok {
		errcode := data["errcode"]
		errmsg := data["errmsg"].(string)
		// return
		c.Data["json"] = map[string]interface{}{"errNo": errcode, "msg": errmsg, "data": "session_key 不存在"}
		// c.ServeJSON()
	} else {
		// var unionId string
		openID = data["openid"].(string)
		// sessionKey = data["session_key"].(string)
		// unionId = data["unionid"].(string)
		// beego.Info(openID)
		// beego.Info(sessionKey)
	}
	if liked == "true" { //如果已经点赞了，则取消
		err = models.DeleteTopicLike(openID)
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = map[string]interface{}{"info": "SUCCESS"}
			c.ServeJSON()
		}
	} else if liked == "false" {
		likeid, err := models.AddTopicLike(idNum, openID)
		if err != nil {
			beego.Error(err)
		} else {
			c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "likeid": likeid}
			c.ServeJSON()
		}
	}
}

// @Title delete wx like by likeid
// @Description delete like by likeid
// @Param id query string  true "The id of like"
// @Success 200 {object} models.DeleteTopicLike
// @Failure 400 Invalid page supplied
// @Failure 404 article not found
// @router /deletewxlike/:id [post]
//删除文章点赞
// func (c *ReplyController) DeleteWxLike() {
// 	id := c.Ctx.Input.Param(":id")
// 	//id转成64为
// 	idNum, err := strconv.ParseInt(id, 10, 64)
// 	if err != nil {
// 		beego.Error(err)
// 	}
// 	err = models.DeleteTopicLike(idNum)
// 	if err != nil {
// 		beego.Error(err)
// 	} else {
// 		c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "likeid": id}
// 		c.ServeJSON()
// 	}
// }

//添加wiki评论
func (c *ReplyController) AddWiki() {
	tid := c.Input().Get("tid") //tid是文章id
	username, _, _, _, _ := checkprodRole(c.Ctx)
	err := models.AddWikiReply(tid, username, c.Input().Get("content"))
	if err != nil {
		beego.Error(err)
	} else {
		c.Data["json"] = tid
		c.ServeJSON()
	}
	op := c.Input().Get("op")
	switch op {
	case "b":
		c.Redirect("/topic/view_b/"+tid, 302)
	case "c":
		c.Redirect("/wiki/view/"+tid, 302)
	default:
		c.Redirect("/topic/view/"+tid, 302)
	}
}

//删除wiki评论
func (c *ReplyController) DeleteWiki() {
	if !checkAccount(c.Ctx) {
		return
	}
	tid := c.Input().Get("tid")
	err := models.DeleteWikiReply(c.Input().Get("rid"))
	if err != nil {
		beego.Error(err)
	}
	c.Redirect("/wiki/view/"+tid, 302)
}

// get file modified time
func FileMTime(file string) (int64, error) {
	f, e := os.Stat(file)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

// put string to file
func FilePutContent(file string, content string) (int, error) {
	fs, e := os.Create(file)
	if e != nil {
		return 0, e
	}
	defer fs.Close()
	return fs.WriteString(content)
}

// get string from text file
// func FileGetContent(file string) (string, error) {
// 	if !IsFile(file) {
// 		return "", os.ErrNotExist
// 	}
// 	b, e := ioutil.ReadFile(file)
// 	if e != nil {
// 		return "", e
// 	}
// 	return string(b), nil
// }
