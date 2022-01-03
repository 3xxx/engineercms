package controllers

import (
	"encoding/json"
	"errors"
	"github.com/3xxx/engineercms/controllers/tool/result"
	"github.com/3xxx/engineercms/controllers/utils"
	"github.com/3xxx/engineercms/controllers/utils/ziptil"
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// CMSSHARE PRODUCT API
// type ShareController struct {
// 	beego.Controller
// }
const (
	//single file.
	SHARE_TYPE_FILE = "FILE"
	//directory
	SHARE_TYPE_DIRECTORY = "DIRECTORY"
	//mix things
	SHARE_TYPE_MIX = "MIX"
)

const (
	SHARE_MAX_NUM = 100
)

type WebResult struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ShareController struct {
	// BaseController
	web.Controller
	// shareDao  *ShareDao
	// bridgeDao *BridgeDao
	// matterDao *MatterDao
	// matterService *MatterService
	// shareService  *ShareService
}

// func (this *ShareController) Init() {
// 	this.BaseController.Init()
// 	b := core.CONTEXT.GetBean(this.shareDao)
// 	if b, ok := b.(*ShareDao); ok {
// 		this.shareDao = b
// 	}
// 	b = core.CONTEXT.GetBean(this.bridgeDao)
// 	if b, ok := b.(*BridgeDao); ok {
// 		this.bridgeDao = b
// 	}
// 	b = core.CONTEXT.GetBean(this.matterDao)
// 	if b, ok := b.(*MatterDao); ok {
// 		this.matterDao = b
// 	}
// 	b = core.CONTEXT.GetBean(this.matterService)
// 	if b, ok := b.(*MatterService); ok {
// 		this.matterService = b
// 	}
// 	b = core.CONTEXT.GetBean(this.shareService)
// 	if b, ok := b.(*ShareService); ok {
// 		this.shareService = b
// 	}
// }

// func (this *ShareController) RegisterRoutes() map[string]func(writer http.ResponseWriter, request *http.Request) {
// 	routeMap := make(map[string]func(writer http.ResponseWriter, request *http.Request))
// 	routeMap["/api/share/create"] = this.Wrap(this.Create, USER_ROLE_USER)
// 	routeMap["/api/share/delete"] = this.Wrap(this.Delete, USER_ROLE_USER)
// 	routeMap["/api/share/delete/batch"] = this.Wrap(this.DeleteBatch, USER_ROLE_USER)
// 	routeMap["/api/share/detail"] = this.Wrap(this.Detail, USER_ROLE_USER)
// 	routeMap["/api/share/page"] = this.Wrap(this.Page, USER_ROLE_USER)
// 	routeMap["/api/share/browse"] = this.Wrap(this.Browse, USER_ROLE_GUEST)
// 	routeMap["/api/share/zip"] = this.Wrap(this.Zip, USER_ROLE_GUEST)
// 	return routeMap
// }

// @Title post create a new share
// @Description post create a new share
// @Param matterUuids query string true "The matterUuids of share"
// @Param expireInfinity query bool true "The xpireInfinity of share"
// @Param expireTime query string true "The expireTime of share"
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /create [post]
func (c *ShareController) Create() {
	// matterUuids := request.FormValue("matterUuids")
	// expireInfinityStr := request.FormValue("expireInfinity")
	// expireTimeStr := request.FormValue("expireTime")
	matterUuids := c.GetString("matterUuids")
	expireInfinityStr := c.GetString("expireInfinityStr")
	expireTimeStr := c.GetString("expireTimeStr")
	if matterUuids == "" {
		// panic(result.BadRequest("matterUuids cannot be null"))
		logs.Error("matterUuids cannot be null")
	}
	var expireTime time.Time
	expireInfinity := false
	if expireInfinityStr == "true" {
		expireInfinity = true
		expireTime = time.Now()
	} else {
		if expireTimeStr == "" {
			// panic(result.BadRequest("time format error"))
			logs.Error("time format error")
		} else {
			expireTime = utils.ConvertDateTimeStringToTime(expireTimeStr)
		}

		if expireTime.Before(time.Now()) {
			// panic(result.BadRequest("expire time cannot before now"))
			logs.Error("expire time cannot before now")
		}
	}
	uuidArray := strings.Split(matterUuids, ",")

	if len(uuidArray) == 0 {
		// panic(result.BadRequest("share at least one file"))
		logs.Error("share at least one file")
	} else if len(uuidArray) > SHARE_MAX_NUM {
		// panic(result.BadRequestI18n(request, i18n.ShareNumExceedLimit, len(uuidArray), SHARE_MAX_NUM))
		logs.Error("ShareNumExceedLimit")
	}

	var name string
	shareType := SHARE_TYPE_MIX
	// user := this.checkUser(request)
	v := c.GetSession("uname") //用来获取存储在服务器端中的数据??。
	// beego.Info(v)                          //qin.xc
	var user models.User
	var err error
	if v != nil { //如果登录了
		user, err = models.GetUserByUsername(v.(string))
		if err != nil {
			logs.Error(err)
		}
	}

	// var puuid int64 //string
	// var matters []*Matter
	// var products []*models.Product
	products := make([]models.Product, 0)
	productslice := make([]models.Product, 1)
	for key, uuid := range uuidArray {
		// matter := this.matterDao.CheckByUuid(uuid)//原始文件是查找matter（file）是否属于登陆者
		// if matter.UserUuid != user.Uuid {
		// 	panic(result.UNAUTHORIZED)
		// }
		idNum, err := strconv.ParseInt(uuid, 10, 64)
		if err != nil {
			logs.Error(err)
		}
		product, err := models.GetProd(idNum)
		productslice[0] = product
		// matters = append(matters, matter)
		products = append(products, productslice...)
		if key == 0 {
			// puuid = matter.Puuid
			// puuid = product.Id
			name = product.Title
			// if matter.Dir {
			shareType = SHARE_TYPE_DIRECTORY
			// } else {
			// 	shareType = SHARE_TYPE_FILE
			// }
		} else {
			// if matter.Puuid != puuid {
			// 	panic(utils.Unauthorized("you can only share files in the same directory"))
			// }
		}
	}
	if len(products) > 1 {
		shareType = SHARE_TYPE_MIX
		name = products[0].Title + "," + products[1].Title + " ..."
	}
	share := &models.Share{
		Name:           name,
		ShareType:      shareType,
		UserId:         user.Id,
		Username:       user.Username,
		DownloadTimes:  0,
		Code:           utils.RandomString4(),
		ExpireInfinity: expireInfinity,
		ExpireTime:     expireTime,
	}
	// this.shareDao.Create(share)
	share, err = models.CreateShare(share)
	if err != nil {
		logs.Error(err)
	}
	for _, product := range products {
		bridge := &models.Bridge{
			ShareUuid: share.Uuid,
			ProductId: product.Id,
		}
		// this.bridgeDao.Create(bridge)
		_, err = models.CreateBridge(bridge)
		if err != nil {
			logs.Error(err)
		}
	}
	c.Data["json"] = map[string]interface{}{"code": "OK", "msg": "", "data": share}
	c.ServeJSON()
	// return c.Success(share)
}

// func (this *ShareController) Delete(writer http.ResponseWriter, request *http.Request) *result.WebResult {
// 	uuid := request.FormValue("uuid")
// 	if uuid == "" {
// 		panic(result.BadRequest("uuid cannot be null"))
// 	}
// 	share := this.shareDao.FindByUuid(uuid)
// 	if share != nil {
// 		this.bridgeDao.DeleteByShareUuid(share.Uuid)
// 		this.shareDao.Delete(share)
// 	}
// 	return this.Success(nil)
// }

// func (this *ShareController) DeleteBatch(writer http.ResponseWriter, request *http.Request) *result.WebResult {
// 	uuids := request.FormValue("uuids")
// 	if uuids == "" {
// 		panic(result.BadRequest("uuids cannot be null"))
// 	}
// 	uuidArray := strings.Split(uuids, ",")
// 	for _, uuid := range uuidArray {
// 		imageCache := this.shareDao.FindByUuid(uuid)
// 		user := this.checkUser(request)
// 		if imageCache.UserUuid != user.Uuid {
// 			panic(result.UNAUTHORIZED)
// 		}
// 		this.shareDao.Delete(imageCache)
// 	}
// 	return this.Success("OK")
// }

// @Title get a share
// @Description get a share
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /detail/:id [get]
func (c *ShareController) Detail() {
	uuid := c.Ctx.Input.Param(":id")
	if uuid == "" {
		panic(result.BadRequest("uuid cannot be null"))
	}
	share, err := models.CheckByUuidShare(uuid)
	if err != nil {
		logs.Error(err)
	}
	// v := c.GetSession("uname") //用来获取存储在服务器端中的数据??。
	// var user models.User
	// if v != nil { //如果登录了
	// 	user, err = models.GetUserByUsername(v.(string))
	// 	if err != nil {
	// 		logs.Error(err)
	// 	}
	// }

	// if share.UserId != user.Id {
	// 	c.Data["ShowCode"] = true
	// } else {
	// 	c.Data["ShowCode"] = false
	// }
	c.Data["Share"] = share
	// c.Data["Shareuuid"] = share.Uuid
	c.Data["Shareuuid"] = uuid

	c.TplName = "share/share.tpl"
}

// @Title get a share
// @Description get a share
// @Param shareUuid query string true "The shareUuid of share"
// @Param code query string false "The code of share"
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /browse [get]
func (c *ShareController) Browse() {
	// shareUuid := request.FormValue("shareUuid")
	// code := request.FormValue("code")
	//puuid can be "root"
	// puuid := request.FormValue("puuid")
	// rootUuid := request.FormValue("rootUuid")
	shareUuid := c.GetString("shareUuid")
	code := c.GetString("code")
	v := c.GetSession("uname") //用来获取存储在服务器端中的数据??。
	// beego.Info(v)                          //qin.xc
	var user models.User
	var err error
	if v != nil { //如果登录了
		user, err = models.GetUserByUsername(v.(string))
		if err != nil {
			logs.Error(err)
		}
	}
	// beego.Info(user)
	// if share.UserId != user.Id {
	// 	panic(result.UNAUTHORIZED)
	// }
	// user := this.findUser(request)
	share, err := CheckShare(c.Ctx, shareUuid, code, &user)
	if err != nil {
		logs.Error(err)
		// code: "NEED_SHARE_CODE"
		// data: null
		// msg: "提取码必填"
		c.Data["json"] = map[string]interface{}{"code": err, "msg": "提取码必填", "data": nil}
		c.ServeJSON()
		return
	}
	bridges, err := models.FindByShareUuid(share.Uuid)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"code": err, "msg": "bridges查询错误", "data": nil}
		c.ServeJSON()
		return
	}
	var products []*models.Product
	if len(bridges) != 0 {
		uuids := make([]int64, 0)
		for _, bridge := range bridges {
			uuids = append(uuids, bridge.ProductId)
		}
		// sortArray := []builder.OrderPair{
		// 	{
		// 		Key:   "dir",
		// 		Value: DIRECTION_DESC,
		// 	},
		// }
		products, err = models.FindByUuids(uuids)
	} else {
		c.Data["json"] = map[string]interface{}{"code": err, "msg": "未查到成果", "data": nil}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{"code": "OK", "msg": "", "data": products}
	c.ServeJSON()
}

// check whether shareUuid and shareCode matches. check whether user can access.
func CheckShare(ctx *context.Context, shareUuid string, code string, user *models.User) (*models.Share, error) {
	share, err := models.CheckByUuidShare(shareUuid)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(user)
	// beego.Info(code)
	//if self, not need shareCode
	//因为user的模型查询设计问题，永远不会为nil，只会是&{0 0 0001--01-01 00:00:00 +0000
	if user == nil || user.Id != share.UserId {
		//if not login or not self's share, shareCode is required.
		if code == "" {
			// panic(result.CustomWebResultI18n(request, result.NEED_SHARE_CODE, i18n.ShareCodeRequired))
			logs.Error("NEED_SHARE_CODE")
			return nil, errors.New("NEED_SHARE_CODE")
		} else if share.Code != code {
			// panic(result.CustomWebResultI18n(request, result.SHARE_CODE_ERROR, i18n.ShareCodeError))
			logs.Error("SHARE_CODE_ERROR")
			return nil, errors.New("SHARE_CODE_ERROR")
		} else {
			if !share.ExpireInfinity {
				if share.ExpireTime.Before(time.Now()) {
					// panic(result.BadRequest("share expired"))
					logs.Error("share expired")
					return nil, errors.New("share expired")
				}
			}
		}
	}
	return share, err
}

type sharerequest struct {
	Id        int64  `json:"id"`
	ShareUuid string `json:"shareUuid"`
	Code      string `json:"code"`
}

// @Title download a share
// @Description download a share
// @Param shareUuid query string true "The shareUuid of share"
// @Param code query string false "The code of share"
// @Param id query string true "The id of attachment"
// @Success 200 {object} models.Attachment
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /download [post]
func (c *ShareController) Download() {
	content := c.Ctx.Input.RequestBody
	var ob sharerequest
	err := json.Unmarshal(content, &ob)
	// beego.Info(ob.Id)
	// beego.Info(ob.ShareUuid)

	shareUuid := ob.ShareUuid  //c.GetString("shareUuid")
	code := ob.Code            //c.GetString("code")
	id := ob.Id                //c.GetString("id")
	v := c.GetSession("uname") //用来获取存储在服务器端中的数据??。
	var user models.User
	// var err error
	if v != nil { //如果登录了
		user, err = models.GetUserByUsername(v.(string))
		if err != nil {
			logs.Error(err)
		}
	}
	// beego.Info(&user)
	// if &user == nil {
	// 	beego.Info("user为nil")
	// }
	_, err = CheckShare(c.Ctx, shareUuid, code, &user)
	if err != nil {
		logs.Error(err)
		// code: "NEED_SHARE_CODE"
		// data: null
		// msg: "提取码必填"
		c.Data["json"] = map[string]interface{}{"code": err, "msg": "提取码必填", "data": nil}
		c.ServeJSON()
		return
	}
	//根据id查出attachment进行下载
	//pid转成64为
	// idNum, err := strconv.ParseInt(id, 10, 64)
	// if err != nil {
	// 	logs.Error(err)
	// 	utils.FileLogs.Error(c.Ctx.Input.IP() + " 转换id " + err.Error())
	// }
	product, err := models.GetProd(id)
	if err != nil {
		logs.Error(err)
	}
	//查出attachment id
	attachments, err := models.GetAttachments(product.Id)
	if err != nil {
		logs.Error(err)
	}

	//由proj id取得url
	fileurl, _, err := GetUrlPath(product.ProjectId)
	if err != nil {
		logs.Error(err)
	}
	//这里只下载第一个文件哦！！
	if len(attachments) > 0 {
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

		c.Ctx.Output.Download(fileurl + "/" + attachments[0].FileName)
	} else {
		c.Data["json"] = map[string]interface{}{"code": err, "msg": "", "data": product}
		c.ServeJSON()
	}
}

// 	var matters []*Matter
// 	if len(bridges) != 0 {
// 		uuids := make([]string, 0)
// 		for _, bridge := range bridges {
// 			uuids = append(uuids, bridge.MatterUuid)
// 		}
// 		sortArray := []builder.OrderPair{
// 			{
// 				Key:   "dir",
// 				Value: DIRECTION_DESC,
// 			},
// 		}
// 		matters = this.matterDao.FindByUuids(uuids, sortArray)
// 		share.Matters = matters
// 	}
// 	return this.Success(share)
// }

// check whether shareUuid and shareCode matches. check whether user can access.
// func (this *ShareService) CheckShare(request *http.Request, shareUuid string, code string, user *User) *Share {
// 	share := this.shareDao.CheckByUuid(shareUuid)
// 	//if self, not need shareCode
// 	if user == nil || user.Uuid != share.UserUuid {
// 		//if not login or not self's share, shareCode is required.
// 		if code == "" {
// 			panic(result.CustomWebResultI18n(request, result.NEED_SHARE_CODE, i18n.ShareCodeRequired))
// 		} else if share.Code != code {
// 			panic(result.CustomWebResultI18n(request, result.SHARE_CODE_ERROR, i18n.ShareCodeError))
// 		} else {
// 			if !share.ExpireInfinity {
// 				if share.ExpireTime.Before(time.Now()) {
// 					panic(result.BadRequest("share expired"))
// 				}
// 			}
// 		}
// 	}
// 	return share
// }

// func (this *ShareController) Page(writer http.ResponseWriter, request *http.Request) *result.WebResult {
// 	pageStr := request.FormValue("page")
// 	pageSizeStr := request.FormValue("pageSize")
// 	orderCreateTime := request.FormValue("orderCreateTime")
// 	user := this.checkUser(request)
// 	var page int
// 	if pageStr != "" {
// 		page, _ = strconv.Atoi(pageStr)
// 	}
// 	pageSize := 200
// 	if pageSizeStr != "" {
// 		tmp, err := strconv.Atoi(pageSizeStr)
// 		if err == nil {
// 			pageSize = tmp
// 		}
// 	}
// 	sortArray := []builder.OrderPair{
// 		{
// 			Key:   "create_time",
// 			Value: orderCreateTime,
// 		},
// 	}
// 	pager := this.shareDao.Page(page, pageSize, user.Uuid, sortArray)
// 	return this.Success(pager)
// }

// func (this *ShareController) CheckShare(writer http.ResponseWriter, request *http.Request) *Share {

// 	shareUuid := request.FormValue("shareUuid")
// 	code := request.FormValue("code")
// 	user := this.findUser(request)

// 	return this.shareService.CheckShare(request, shareUuid, code, user)
// }

// func (this *ShareController) Browse(writer http.ResponseWriter, request *http.Request) *result.WebResult {
// 	shareUuid := request.FormValue("shareUuid")
// 	code := request.FormValue("code")
// 	//puuid can be "root"
// 	puuid := request.FormValue("puuid")
// 	rootUuid := request.FormValue("rootUuid")
// 	user := this.findUser(request)
// 	share := this.shareService.CheckShare(request, shareUuid, code, user)
// 	bridges := this.bridgeDao.FindByShareUuid(share.Uuid)

// 	if puuid == MATTER_ROOT {
// 		var matters []*Matter
// 		if len(bridges) != 0 {
// 			uuids := make([]string, 0)
// 			for _, bridge := range bridges {
// 				uuids = append(uuids, bridge.MatterUuid)
// 			}
// 			sortArray := []builder.OrderPair{
// 				{
// 					Key:   "dir",
// 					Value: DIRECTION_DESC,
// 				},
// 			}
// 			matters = this.matterDao.FindByUuids(uuids, sortArray)
// 			share.Matters = matters
// 		}
// 	} else {
// 		//if root. No need to validate.
// 		if puuid == rootUuid {
// 			dirMatter := this.matterDao.CheckByUuid(puuid)
// 			share.DirMatter = dirMatter
// 		} else {
// 			dirMatter := this.matterService.Detail(request, puuid)
// 			//check whether shareRootMatter is being sharing
// 			shareRootMatter := this.matterDao.CheckByUuid(rootUuid)
// 			if !shareRootMatter.Dir {
// 				panic(result.BadRequestI18n(request, i18n.MatterDestinationMustDirectory))
// 			}
// 			this.bridgeDao.CheckByShareUuidAndMatterUuid(share.Uuid, shareRootMatter.Uuid)
// 			//stop at rootUuid
// 			find := false
// 			parentMatter := dirMatter.Parent
// 			for parentMatter != nil {
// 				if parentMatter.Uuid == rootUuid {
// 					parentMatter.Parent = nil
// 					find = true
// 					break
// 				}
// 				parentMatter = parentMatter.Parent
// 			}
// 			if !find {
// 				panic(result.BadRequest("rootUuid is not the root of share."))
// 			}
// 			share.DirMatter = dirMatter
// 		}
// 	}
// 	return this.Success(share)
// }

// @Title download a share
// @Description download a share
// @Param shareUuid query string true "The shareUuid of share"
// @Param code query string false "The code of share"
// @Success 200 {object} models.Attachment
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /downloadzip [post]
func (c *ShareController) DownloadZip() {
	// shareUuid := request.FormValue("shareUuid")
	// code := request.FormValue("code")
	// user := this.findUser(request)
	content := c.Ctx.Input.RequestBody
	var ob sharerequest
	err := json.Unmarshal(content, &ob)
	// beego.Info(ob.Id)
	// beego.Info(ob.ShareUuid)

	shareUuid := ob.ShareUuid //c.GetString("shareUuid")
	code := ob.Code           //c.GetString("code")
	// id := ob.Id
	// shareUuid := c.GetString("shareUuid")
	// code := c.GetString("code")
	v := c.GetSession("uname") //用来获取存储在服务器端中的数据??。
	var user models.User
	// var err error
	if v != nil { //如果登录了
		user, err = models.GetUserByUsername(v.(string))
		if err != nil {
			logs.Error(err)
		}
	}

	share, err := CheckShare(c.Ctx, shareUuid, code, &user)
	if err != nil {
		logs.Error(err)
		// code: "NEED_SHARE_CODE"
		// data: null
		// msg: "提取码必填"
		c.Data["json"] = map[string]interface{}{"code": err, "msg": "提取码必填", "data": nil}
		c.ServeJSON()
		return
	}
	bridges, err := models.FindByShareUuid(share.Uuid)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"code": err, "msg": "bridges查询错误", "data": nil}
		c.ServeJSON()
		return
	}
	var prod models.Product
	if len(bridges) != 0 {
		//建立文件夹
		_, err := os.Stat("./temp/engineercms")
		if err != nil {
			if os.IsNotExist(err) {
				// return false
				err := os.MkdirAll("./temp/engineercms", 0777)
				if err != nil {
					logs.Error(err)
				}
			}

		}

		uuids := make([]int64, 0)
		for _, bridge := range bridges {
			uuids = append(uuids, bridge.ProductId)

			prod, err = models.GetProd(bridge.ProductId)
			if err != nil {
				logs.Error(err)
			}
			_, DiskDirectory, err := GetUrlPath(prod.ProjectId)
			if err != nil {
				logs.Error(err)
			}
			attachments, err := models.GetAttachments(bridge.ProductId)
			for _, v := range attachments {
				fileext := path.Ext(v.FileName)
				matched, err := regexp.MatchString("\\.*[m|M][c|C][d|D]", fileext)
				if err != nil {
					logs.Error(err)
				}
				if matched {
					c.Data["json"] = map[string]interface{}{"info": "ERROR", "data": "不能下载mcd文件!"}
					c.ServeJSON()
					return
				}

				// pathLink = DiskDirectory + "/" + v.FileName
				_, err = CopyFile("./temp/engineercms/"+v.FileName, DiskDirectory+"/"+v.FileName)
				//targetfile,sourcefile
				if err != nil {
					// fmt.Println(err.Error())
					logs.Error(err)
				}
				// fmt.Println(w)
			}
		}
	}
	// products, err = models.FindByUuids(uuids)
	//压缩
	ziptil.Zip("./temp/engineercms", "./temp/"+prod.Title+".zip")
	// TarGz("./temp/engineercms/", "./temp/"+prod.Title+".tar.gz") //压缩
	// fileName := path.Base("./temp/" + prod.Title + ".tar.gz")
	// fileName = url.QueryEscape(fileName) // 防止中文乱码
	// c.Ctx.Output.Header("Content-Type", "application/octet-stream")
	// c.Ctx.Output.Header("content-disposition", "attachment; filename=\""+fileName+"\"")
	// c.Ctx.Output.Header("Content-Disposition", "attachment; filename = \""+fileName+"\"")

	// c.Ctx.ResponseWriter.Header().Set("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	// var name = prod.Title + ".tar.gz"
	// query := url.Values{}
	// query.Add("filename", prod.Title+".tar.gz")
	// filename := query.Encode()
	// beego.Info(fileName)
	// headers.add("Content-Disposition","attachment;filename*=UTF-8''"+URLEncoder.encode("中国","UTF-8")+".txt");
	// c.Ctx.ResponseWriter.Header().Set("Content-Disposition", "attachment; filename* = UTF-8''"+filename)
	// beego的download已经处理过了乱码
	c.Ctx.Output.Download("./temp/" + prod.Title + ".zip") //prod.Title + ".tar.gz"
	// let name = urlencode("文件.txt", "utf-8");
	// res.setHeader("Content-Disposition", "attachment; filename* = UTF-8''"+name);

	// c.Data["json"] = map[string]interface{}{"code": "OK", "msg": ""}
	// c.ServeJSON()
	//delete the temp zip file.
	err = os.Remove("./temp/" + prod.Title + ".zip")
	if err != nil {
		logs.Error(err)
	}
	err = RemoveContents("./temp/engineercms/")
	if err != nil {
		logs.Error(err)
		os.Exit(1)
	}
}

//删除文件夹下所有文件
func RemoveContents(dir string) error {
	d, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names, err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(dir, name))
		if err != nil {
			return err
		}
	}
	return nil
}

//Copyfile
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

//try to delete empty dir. true: delete an empty dir, false: delete nothing.
func DeleteEmptyDir(dirPath string) bool {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		logs.Error(err)
	}
	if len(dir) == 0 {
		//empty dir
		err = os.Remove(dirPath)
		if err != nil {
			logs.Error(err)
		}
		return true
	}
	return false
}

//Download specified matters. matters must have the same puuid.
// func (this *MatterService) DownloadZip(
// 	writer http.ResponseWriter,
// 	request *http.Request,
// 	matters []*Matter) {

// 	if matters == nil || len(matters) == 0 {
// 		panic(result.BadRequest("matters cannot be nil."))
// 	}
// 	userUuid := matters[0].UserUuid
// 	puuid := matters[0].Puuid

// 	for _, m := range matters {
// 		if m.UserUuid != userUuid {
// 			panic(result.BadRequest("userUuid not same"))
// 		} else if m.Puuid != puuid {
// 			panic(result.BadRequest("puuid not same"))
// 		}
// 	}

// 	preference := this.preferenceService.Fetch()

// 	//count the num of files will be downloaded.
// 	var count int64 = 0
// 	for _, matter := range matters {
// 		count = count + this.matterDao.CountByUserUuidAndPath(matter.UserUuid, matter.Path)
// 	}

// 	if preference.DownloadDirMaxNum >= 0 {
// 		if count > preference.DownloadDirMaxNum {
// 			panic(result.BadRequestI18n(request, i18n.MatterSelectNumExceedLimit, count, preference.DownloadDirMaxNum))
// 		}
// 	}

// 	//count the size of files will be downloaded.
// 	var sumSize int64 = 0
// 	for _, matter := range matters {
// 		sumSize = sumSize + this.matterDao.SumSizeByUserUuidAndPath(matter.UserUuid, matter.Path)
// 	}

// 	if preference.DownloadDirMaxSize >= 0 {
// 		if sumSize > preference.DownloadDirMaxSize {
// 			panic(result.BadRequestI18n(request, i18n.MatterSelectSizeExceedLimit, util.HumanFileSize(sumSize), util.HumanFileSize(preference.DownloadDirMaxSize)))
// 		}
// 	}

// 	//prepare the temp zip dir
// 	destZipDirPath := fmt.Sprintf("%s/%d", GetUserZipRootDir(matters[0].Username), time.Now().UnixNano()/1e6)
// 	util.MakeDirAll(destZipDirPath)

// 	destZipName := fmt.Sprintf("%s.zip", matters[0].Name)
// 	if len(matters) > 1 || !matters[0].Dir {
// 		destZipName = "archive.zip"
// 	}

// 	destZipPath := fmt.Sprintf("%s/%s", destZipDirPath, destZipName)

// 	this.zipMatters(request, matters, destZipPath)

// 	download.DownloadFile(writer, request, destZipPath, destZipName, true)

// 	//delete the temp zip file.
// 	err := os.Remove(destZipPath)
// 	if err != nil {
// 		this.logger.Error("error while deleting zip file %s", err.Error())
// 	}
// 	util.DeleteEmptyDir(destZipDirPath)

// }

//zip matters.
//必须用这个，将目标文件直接copy到压缩文件中
// func zipMatters(matters []*Matter, destPath string) {
// 	if util.PathExists(destPath) {
// 		// panic(result.BadRequest("%s exists", destPath))
// 	}
// 	//matters must have the same puuid.
// 	if matters == nil || len(matters) == 0 {
// 		// panic(result.BadRequest("matters cannot be nil."))
// 	}
// 	// userUuid := matters[0].UserUuid
// 	// puuid := matters[0].Puuid
// 	baseDirPath := util.GetDirOfPath(matters[0].AbsolutePath()) + "/"
// 	//wrap children for every matter.
// 	for _, m := range matters {
// 		this.WrapChildrenDetail(request, m)
// 	}
// 	// create temp zip file.
// 	fileWriter, err := os.Create(destPath)
// 	this.PanicError(err)
// 	defer func() {
// 		err := fileWriter.Close()
// 		this.PanicError(err)
// 	}()
// 	zipWriter := zip.NewWriter(fileWriter)
// 	defer func() {
// 		err := zipWriter.Close()
// 		this.PanicError(err)
// 	}()
// 	//DFS algorithm
// 	var walkFunc func(matter *Matter)
// 	walkFunc = func(matter *Matter) {
// 		path := matter.AbsolutePath()
// 		fileInfo, err := os.Stat(path)
// 		this.PanicError(err)
// 		// Create file info.
// 		fileHeader, err := zip.FileInfoHeader(fileInfo)
// 		this.PanicError(err)
// 		// Trim the baseDirPath
// 		fileHeader.Name = strings.TrimPrefix(path, baseDirPath)
// 		// directory has prefix /
// 		if matter.Dir {
// 			fileHeader.Name += "/"
// 		}
// 		writer, err := zipWriter.CreateHeader(fileHeader)
// 		this.PanicError(err)
// 		// only regular file has things to write.
// 		if fileHeader.Mode().IsRegular() {
// 			fileToBeZip, err := os.Open(path)
// 			defer func() {
// 				err = fileToBeZip.Close()
// 				this.PanicError(err)
// 			}()
// 			this.PanicError(err)
// 			_, err = io.Copy(writer, fileToBeZip)
// 			this.PanicError(err)
// 		}
// 		//dfs.
// 		for _, m := range matter.Children {
// 			walkFunc(m)
// 		}
// 	}
// 	for _, m := range matters {
// 		walkFunc(m)
// 	}
// }

//find the current user from request.
// SessionName
// ctx.Input.Session("uname")
// web.AppConfig.String("wxcatalogid")
// func (this *ShareController) findUser(request *http.Request) *User {
// 	//try to find from SessionCache.
// 	// sessionId := util.GetSessionUuidFromRequest(request, core.COOKIE_AUTH_KEY)
// 	sessionId := util.GetSessionUuidFromRequest(request, web.AppConfig.String("SessionName"))
// 	if sessionId == "" {
// 		return nil
// 	}
// 	cacheItem, err := core.CONTEXT.GetSessionCache().Value(sessionId)
// 	if err != nil {
// 		this.logger.Warn("error while get from session cache. sessionId = %s, error = %v", sessionId, err)
// 		return nil
// 	}
// 	if cacheItem == nil || cacheItem.Data() == nil {
// 		this.logger.Warn("cache item doesn't exist with sessionId = %s", sessionId)
// 		return nil
// 	}
// 	if value, ok := cacheItem.Data().(*User); ok {
// 		return value
// 	} else {
// 		this.logger.Error("cache item not store the *User")
// 	}
// 	return nil
// }

//find current error. If not found, panic the LOGIN error.
// func (this *ShareController) checkUser(request *http.Request) *User {
// 	if this.findUser(request) == nil {
// 		// panic(result.LOGIN)
// 		return nil
// 	} else {
// 		return this.findUser(request)
// 	}
// }

// 来自share_service

// func (this *ShareService) Detail(uuid string) *Share {
// 	share := this.shareDao.CheckByUuid(uuid)
// 	return share
// }

// check whether shareUuid and shareCode matches. check whether user can access.
// func (this *ShareService) CheckShare(request *http.Request, shareUuid string, code string, user *User) *Share {
// 	share := this.shareDao.CheckByUuid(shareUuid)
// 	//if self, not need shareCode
// 	if user == nil || user.Uuid != share.UserUuid {
// 		//if not login or not self's share, shareCode is required.
// 		if code == "" {
// 			panic(result.CustomWebResultI18n(request, result.NEED_SHARE_CODE, i18n.ShareCodeRequired))
// 		} else if share.Code != code {
// 			panic(result.CustomWebResultI18n(request, result.SHARE_CODE_ERROR, i18n.ShareCodeError))
// 		} else {
// 			if !share.ExpireInfinity {
// 				if share.ExpireTime.Before(time.Now()) {
// 					panic(result.BadRequest("share expired"))
// 				}
// 			}
// 		}
// 	}
// 	return share
// }

//check whether a user can access a matter. shareRootUuid is matter's parent(or parent's parent and so on)
// func (this *ShareService) ValidateMatter(request *http.Request, shareUuid string, code string, user *User, shareRootUuid string, matter *Matter) {
// 	if matter == nil {
// 		panic(result.Unauthorized("matter cannot be nil"))
// 	}
// 	//if self's matter, ok.
// 	if user != nil && matter.UserUuid == user.Uuid {
// 		return
// 	}
// 	if shareUuid == "" || code == "" || shareRootUuid == "" {
// 		panic(result.Unauthorized("shareUuid,code,shareRootUuid cannot be null"))
// 	}
// 	share := this.CheckShare(request, shareUuid, code, user)
// 	//if shareRootUuid is root. Bridge must has record.
// 	if shareRootUuid == MATTER_ROOT {
// 		this.bridgeDao.CheckByShareUuidAndMatterUuid(share.Uuid, matter.Uuid)
// 	} else {
// 		//check whether shareRootMatter is being sharing
// 		shareRootMatter := this.matterDao.CheckByUuid(shareRootUuid)
// 		this.bridgeDao.CheckByShareUuidAndMatterUuid(share.Uuid, shareRootMatter.Uuid)
// 		// shareRootMatter is ancestor of matter.
// 		child := strings.HasPrefix(matter.Path, shareRootMatter.Path)
// 		if !child {
// 			panic(result.BadRequest("%s is not %s's children", matter.Uuid, shareRootUuid))
// 		}
// 	}
// }
