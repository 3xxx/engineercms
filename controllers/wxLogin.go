package controllers

import (
	"github.com/3xxx/engineercms/controllers/utils"
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/go-pay/util"
	"github.com/golang-jwt/jwt/v4"
	// "github.com/gorilla/websocket"
	// "github.com/skip2/go-qrcode"
	"encoding/json"
	// "fmt"
	"context"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// 微信扫码
type WechatLoginController struct {
	web.Controller
}

// @Title get wx login
// @Description wx login
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 qrcode not found
// @router /wxlogin [get]
// 页面获取qrcode地址，前端显示二维码
func (c *WechatLoginController) WxLogin() {
	c.TplName = "wxlogin/wxlogin.tpl"
}

// @Title wx user web
// @Description wx login
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 qrcode not found
// @router /getwxuser [get]
// 页面获取qrcode地址，前端显示二维码
// func (c *WechatLoginController) GetWxUser() {
// 	c.TplName = "wxlogin/wxuser.tpl"
// }

// @Title get wx login qrcode
// @Description wx qrcode generate
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 qrcode not found
// @router /wxloginqrcode [get]
// 页面获取qrcode地址，前端显示二维码
func (c *WechatLoginController) WxLoginQrcode() {
	appstring := "wxAPPID7"
	APPID, err := web.AppConfig.String(appstring)
	if err != nil {
		logs.Error(err)
	}

	state := util.RandomString(10) //防止跨站请求伪造攻击 增加安全性
	path := "https://zsj.itdos.net/v1/wx/redirecturi"
	redirectURL := url.QueryEscape(path) //get userinfo,授权后重定向的回调链接地址， 请使用 urlEncode 对链接进行处理
	requestUrl := "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + APPID + "&redirect_uri=" + redirectURL + "&response_type=code&state=" + state + "&scope=snsapi_userinfo#wechat_redirect"

	// 测试写入响应头头，看看是否前端能带上。可以带上
	c.Ctx.Output.Header("Authorization", "wechat login is so nan a")
	// authString := c.Ctx.Input.Header("Authorization")
	// c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	// c.Header("Authorization", atoken)

	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "code": 200, "status": "", "errorMsg": "", "CodeUrl": requestUrl}
	c.ServeJSON()
	// 生成二维码
	// qrCode, err := qrcode.Encode(wechatLoginURL, qrcode.Medium, 256)
	// if err != nil {
	// 	// 错误处理
	// 	c.String(http.StatusInternalServerError, "Error generating QR code")
	// 	return
	// }
	// 将二维码图片作为响应返回给用户
	// c.Header("Content-Type", "image/png")
	// c.Writer.Write(qrCode)
}

type JWTCustomClaims struct {
	UserId int64 `json:"userid"`
	jwt.RegisteredClaims
}

// @Title get wx redirecturi
// @Description wx redirecturi
// @Param code query string true "The code of wechat callback"
// @Param state query string true "The state of wechat callback"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 pay not found
// @router /redirecturi [get]
// 微信服务端回调地址——必须线上调试才起作用！！
func (c *WechatLoginController) RedirectUri() {
	// 获取微信返回的授权码
	code := c.GetString("code")
	// logs.Info(code)
	state := c.GetString("state")
	c.Data["Code"] = code
	c.Data["State"] = state
	c.TplName = "wxlogin/wxloginsuccess.tpl"

	appstring := "wxAPPID7"
	APPID, err := web.AppConfig.String(appstring)
	if err != nil {
		logs.Error(err)
	}
	secretstring := "wxSECRET7"
	SECRET, err := web.AppConfig.String(secretstring)
	if err != nil {
		logs.Error(err)
	}
	requestUrl := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + APPID + "&secret=" + SECRET + "&code=" + code + "&grant_type=authorization_code"
	tokenResp, err := http.Get(requestUrl)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"data": "", "info": "ERROR", "status": "", "errorMsg": "获取wechat token失败"}
		c.ServeJSON()
	}
	// 解析响应中的access_token和openid
	var tokenData struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		OpenID       string `json:"openid"`
		Scope        string `json:"scope"`
	}
	if err = json.NewDecoder(tokenResp.Body).Decode(&tokenData); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"data": "", "info": "ERROR", "status": "", "errorMsg": "token解析失败"}
		c.ServeJSON()
	}
	// 拉取用户信息(需scope为 snsapi_userinfo)
	userInfoURL := "https://api.weixin.qq.com/sns/userinfo?access_token=" + tokenData.AccessToken + "&openid=" + tokenData.OpenID + "&lang=zh_CN"
	userInfoResp, err := http.Get(userInfoURL)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"data": "", "info": "ERROR", "status": "", "errorMsg": "获取用户信息失败"}
		c.ServeJSON()
	}
	defer userInfoResp.Body.Close()
	var userData struct {
		OpenID   string `json:"openid"`
		Nickname string `json:"nickname"`
		// "sex": 1,
		// 	"province":"PROVINCE",
		// 	"city":"CITY",
		// 	"country":"COUNTRY",
		// 	"headimgurl":"https://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
		// 	"privilege":[ "PRIVILEGE1" "PRIVILEGE2"     ],
		// 	"unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL"
	}
	if err = json.NewDecoder(userInfoResp.Body).Decode(&userData); err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"data": "", "info": "ERROR", "status": "", "errorMsg": "用户信息解析失败"}
		c.ServeJSON()
	}
	// logs.Info(userData.Nickname)
	// user, _, err := models.GetwxUserByNickname(userData.Nickname)
	// if err != nil {
	// 	logs.Error(err)
	// 	c.Data["json"] = map[string]interface{}{"errNo": 0, "msg": "未查到用户", "data": "这个Nickname的用户不存在，新注册该用户。", "openID": userData.Nickname}
	// 	c.ServeJSON()
	// }
	// logs.Info(userData.OpenID)
	// logs.Info(userData.Nickname)
	user, _, err := models.GetwxUserByOpenID(userData.OpenID, userData.Nickname)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"errNo": 0, "msg": "未查到用户", "data": "这个openID的用户不存在，新存入openid", "openID": userData.OpenID}
		c.ServeJSON()
	}
	// logs.Info(user)
	// logs.Info(userid)
	// 写入缓存。如果有值存在，则等待2秒，等GetWxUserLogin里的Bm.Get掉里面的值，即变为空后再写入。因为后者轮询的时间间隔是2s
	// Bm定义与utiles.wechat.go
	hasopenID, err := utils.Bm.IsExist(context.TODO(), "openID")
	if err != nil {
		logs.Error(err)
	}
	if hasopenID {
		time.Sleep(2 * time.Second)
	}
	err = utils.Bm.Put(context.TODO(), "openID", userData.OpenID, 360*time.Second)
	if err != nil {
		logs.Error(err)
	}
	hasnickname, err := utils.Bm.IsExist(context.TODO(), "nickname")
	if err != nil {
		logs.Error(err)
	}
	if hasnickname {
		time.Sleep(2 * time.Second)
	}
	err = utils.Bm.Put(context.TODO(), "nickname", userData.Nickname, 360*time.Second)
	if err != nil {
		logs.Error(err)
	}
	hasusername, err := utils.Bm.IsExist(context.TODO(), "uname")
	if err != nil {
		logs.Error(err)
	}
	if hasusername {
		time.Sleep(2 * time.Second)
	}
	err = utils.Bm.Put(context.TODO(), "uname", user.Username, 360*time.Second)
	if err != nil {
		logs.Error(err)
	}
	hasuserId, err := utils.Bm.IsExist(context.TODO(), "userid")
	if err != nil {
		logs.Error(err)
	}
	if hasuserId {
		time.Sleep(2 * time.Second)
	}
	userid_str := strconv.FormatInt(user.Id, 10)
	err = utils.Bm.Put(context.TODO(), "userid", userid_str, 360*time.Second)
	if err != nil {
		logs.Error(err)
	}
}

// @Title get wx login get user info
// @Description wx login get user info
// @Param code query string true "The code of wechat callback"
// @Param state query string true "The state of wechat callback"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 pay not found
// @router /getwxuserinfo [get]
// 微信服务端回调地址——作废
func (c *WechatLoginController) GetWxUserInfo() {
	appstring := "wxAPPID7"
	APPID, err := web.AppConfig.String(appstring)
	if err != nil {
		logs.Error(err)
	}
	secretstring := "wxSECRET7"
	SECRET, err := web.AppConfig.String(secretstring)
	if err != nil {
		logs.Error(err)
	}
	// 获取微信返回的授权码
	code := c.GetString("code")
	// logs.Info(code)
	// 向微信服务器发送请求，获取access_token和openid
	// https://api.weixin.qq.com/sns/oauth2/access_token?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code
	requestUrl := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + APPID + "&secret=" + SECRET + "&code=" + code + "&grant_type=authorization_code"
	// logs.Info(requestUrl)
	tokenResp, err := http.Get(requestUrl)
	if err != nil {
		logs.Error(err)
		// fmt.Println(err)
		// resp := &ResponseData{
		// 	Data:    nil,
		// 	Message: "error,获取token失败",
		// 	Code:    CodeServerBusy,
		// }
		c.Data["json"] = map[string]interface{}{"data": "", "info": "ERROR", "status": "", "errorMsg": "获取wechat token失败"}
		c.ServeJSON()
		// c.JSON(http.StatusBadRequest, resp)
		// return
	}
	// defer tokenResp.Body.Close()
	// logs.Info(tokenResp)
	// logs.Info(tokenResp.openid)
	// logs.Info(tokenResp.Body)
	// 解析响应中的access_token和openid
	var tokenData struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		OpenID       string `json:"openid"`
		Scope        string `json:"scope"`
	}
	if err = json.NewDecoder(tokenResp.Body).Decode(&tokenData); err != nil {
		c.Data["json"] = map[string]interface{}{"data": "", "info": "ERROR", "status": "", "errorMsg": "token解析失败"}
		c.ServeJSON()
		// resp := &ResponseData{
		// 	Data:    nil,
		// 	Message: "error,获取token失败",
		// 	Code:    CodeServerBusy,
		// }
		// c.JSON(http.StatusBadRequest, resp)
		// return
	}
	// logs.Info(tokenData.OpenID)
	// logs.Info(tokenData.AccessToken)
	userInfoURL := "https://api.weixin.qq.com/sns/userinfo?access_token=" + tokenData.AccessToken + "&openid=" + tokenData.OpenID + "&lang=zh_CN"
	// userInfoURL := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s", tokenData.AccessToken, tokenData.OpenID)
	userInfoResp, err := http.Get(userInfoURL)
	if err != nil {
		// 错误处理
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"data": "", "info": "ERROR", "status": "", "errorMsg": "获取用户信息失败"}
		c.ServeJSON()
		// zap.L().Error("获取失败")
		// return
	}
	defer userInfoResp.Body.Close()

	//------------------------------------
	var userData struct {
		OpenID   string `json:"openid"`
		Nickname string `json:"nickname"`
	}
	if err = json.NewDecoder(userInfoResp.Body).Decode(&userData); err != nil {
		// 错误处理
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"data": "", "info": "ERROR", "status": "", "errorMsg": "用户信息解析失败"}
		c.ServeJSON()
		// zap.L().Error("获取用户信息失败")
		// return
	}
	//用户的名字
	// var user models.User
	// nickname := userData.Nickname
	// logs.Info(userData.Nickname)
	// user, _, err := models.GetwxUserByNickname(userData.Nickname)
	// if err != nil {
	// 	logs.Error(err)
	// 	c.Data["json"] = map[string]interface{}{"errNo": 0, "msg": "未查到用户", "data": "这个Nickname的用户不存在，新注册该用户。", "openID": userData.Nickname}
	// 	c.ServeJSON()
	// }
	user, _, err := models.GetwxUserByOpenID(userData.OpenID, userData.Nickname)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"errNo": 0, "msg": "未查到用户", "data": "这个openID的用户不存在，新存入openid", "openID": userData.OpenID}
		c.ServeJSON()
	}

	// if err = mysql.DB.Where("user_name=?", nickname).First(&user).Error; err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		user.UserName = nickname
	// 		user.UserID, _ = snowflake.GetID()
	// 		user.Identity = "普通用户"
	// 	} else {
	// 		logs.Error(err)
	// 		c.Data["json"] = map[string]interface{}{"data": "", "info": "ERROR", "status": "", "errorMsg": "查询用户数据表失败"}
	// 		c.ServeJSON()
	// 		// zap.L().Error("验证登录信息过程中出错")
	// 		// ResponseError(c, CodeServerBusy)
	// 		// return
	// 	}
	// }
	//添加jwt验证
	// 生成 access token 在 access token 中需要包含我们自定义的字段，比如用户 ID
	mySigningKey := []byte("hotqin888")
	mc := JWTCustomClaims{
		user.Id,
		jwt.RegisteredClaims{
			// ExpiresAt 是一个时间戳，代表 access token 的过期时间
			ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Add(time.Duration(120)*time.Minute).Unix(), 0)), // time.Now().Add(time.Duration(120) * time.Minute).Unix(),
			// 签发人
			Issuer: user.Username,
		},
	}

	// 生成 access token
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, mc).SignedString(mySigningKey)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"data": "", "info": "ERROR", "status": "", "errorMsg": "生成JWT令牌失败"}
		c.ServeJSON()
		// log.Printf("generate access token failed: %v \n", err)
		// return "", "", err
	}

	// // 生成 refresh token
	// // refresh token 只需要包含标准的声明，不需要包含自定义的声明
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		// ExpiresAt 是一个时间戳，代表 refresh token 的过期时间
		ExpiresAt: time.Now().Add(time.Duration(120) * time.Minute).Unix(),
		// 签发人
		Issuer: user.Username,
	}).SignedString(mySigningKey)

	c.Ctx.Output.Header("Authorization", accessToken)
	// authString := c.Ctx.Input.Header("Authorization")
	// c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	// c.Header("Authorization", atoken)

	//用户登录后，存储openid在服务端的session里，下次用户通过hotqinsessionid来取得openID、sessionKey等。
	// c.SetSession("openID", userData.OpenID)
	// c.SetSession("uname", user.Username) // 用户第一扫码，这个不存在
	// c.SetSession("nickname", userData.Nickname)
	//发送成功响应
	c.Data["json"] = map[string]interface{}{"data": "登录成功", "code": 200, "info": "SUCCESS", "status": "登录成功", "errorMsg": "", "nickname": userData.Nickname, "openid": userData.OpenID, "AccessToken": accessToken, "RefreshToken": refreshToken}
	c.ServeJSON()
	// c.TplName = "wxlogin/wxloginsuccess.tpl"

	// c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	// c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	// c.Data["json"] = map[string]interface{}{"code":"200","msg":"請求成功","data":"33333"}
	// c.ServeJSON()

	// c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	// c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	// s := map[string]interface{}{
	// 	"code":200,
	// 	"msg":"请求成功",
	// 	"data":[]map[string]interface{}{map[string]interface{}{"name":"小明","old":"7"},map[string]interface{}{"name":"小紅","old":"10"}},
	// }
	// c.Data["json"]=s
	// c.ServeJSON()
	// 	附：检验授权凭证（access_token）是否有效
	// 请求方法

	// http：GET（请使用https协议）：

	// https://api.weixin.qq.com/sns/auth?access_token=ACCESS_TOKEN&openid=OPENID
}

// @Title get wx user login bool
// @Description wx login user login bool
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 pay not found
// @router /getwxuserlogin [get]
// 二维码页面轮询
func (c *WechatLoginController) GetWxUserLogin() {
	// Get方法：查询缓存里的一条数据，根据key查询
	openID, err := utils.Bm.Get(context.TODO(), "openID")
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"data": "登录失败", "code": 400, "info": "ERROR", "status": "登录失败，无openid", "errorMsg": "login error", "islogin": false}
		c.ServeJSON()
		return
	}
	c.SetSession("openID", openID.(string))
	// 清除缓存
	utils.Bm.Delete(context.TODO(), "openID")
	nickName, err := utils.Bm.Get(context.TODO(), "nickname")
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"data": "登录失败", "code": 400, "info": "ERROR", "status": "登录失败，无nickename", "errorMsg": "login error", "islogin": false}
		c.ServeJSON()
		return
	}
	c.SetSession("nickname", nickName.(string))
	// 清除缓存
	utils.Bm.Delete(context.TODO(), "nickname")
	userName, err := utils.Bm.Get(context.TODO(), "uname")
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"data": "登录失败", "code": 400, "info": "ERROR", "status": "登录失败，无username", "errorMsg": "login error", "islogin": false}
		c.ServeJSON()
		return
	}
	c.SetSession("uname", userName.(string))
	// 清除缓存
	utils.Bm.Delete(context.TODO(), "uname")
	userId, err := utils.Bm.Get(context.TODO(), "userid")
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"data": "登录失败", "code": 400, "info": "ERROR", "status": "登录失败，无userid", "errorMsg": "login error", "islogin": false}
		c.ServeJSON()
		return
	}
	c.SetSession("userid", userId.(string))
	// 清除缓存
	utils.Bm.Delete(context.TODO(), "userid")
	// user, err := models.GetUserByOpenID(openID.(string))
	// if err != nil {
	// 	logs.Error(err)
	// 	c.Data["json"] = map[string]interface{}{"data": "查询用户openid失败", "code": 400, "info": "ERROR", "status": "查询用户openid失败", "errorMsg": "query user openid error", "islogin": false}
	// 	c.ServeJSON()
	// }

	// _, _, _, _, islogin := checkprodRole(c.Ctx)
	// // v := ctx.Input.Session("uname")
	// openID := c.GetSession("openID")
	// nickName := c.GetSession("nickname")
	c.Data["json"] = map[string]interface{}{"data": "登录成功", "code": 200, "info": "SUCCESS", "status": "登录成功", "errorMsg": "", "islogin": true, "nickname": nickName.(string), "openid": openID.(string), "username": userName.(string), "userid": userId.(string)}
	c.ServeJSON()
}

// https://user.3d66.com/login/Wx_Official_Account/checkLogin
// code: 1
// data: []
// msg: "待扫码"
// status: 500

// TOKEN 假设您在Go代码中定义了一个名为TOKEN的常量，用于存储您的令牌值
// const TOKEN = "111"

// 配置公众号的token
// func CheckSignature(c *gin.Context) {
// 	// 获取查询参数中的签名、时间戳和随机数
// 	signature := c.Query("signature")
// 	timestamp := c.Query("timestamp")
// 	nonce := c.Query("nonce")
// 	echostr := c.Query("echostr")
// 	// 创建包含令牌、时间戳和随机数的字符串切片
// 	tmpArr := []string{TOKEN, timestamp, nonce}
// 	// 对切片进行字典排序
// 	sort.Strings(tmpArr)
// 	// 将排序后的元素拼接成单个字符串
// 	tmpStr := ""
// 	for _, v := range tmpArr {
// 		tmpStr += v
// 	}
// 	// 对字符串进行SHA-1哈希计算
// 	tmpHash := sha1.New()
// 	tmpHash.Write([]byte(tmpStr))
// 	tmpStr = fmt.Sprintf("%x", tmpHash.Sum(nil))
// 	fmt.Println(tmpStr)
// 	fmt.Println(signature)
// 	// 将计算得到的签名与请求中提供的签名进行比较，并根据结果发送相应的响应
// 	if tmpStr == signature {
// 		c.String(200, echostr)
// 		redis.Client.Set(context.Background(), "library:token", tmpStr, 7*24*time.Hour)
// 	} else {
// 		c.String(403, "签名验证失败 "+timestamp)
// 	}
// }

// @Title get wx user
// @Description wx user
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 pay not found
// @router /getwxuser [get]
// 回调地址
// func (c *WechatLoginController) GetWxUser() {
// 	userToken := c.GetString("Authorization")
// 	fmt.Println("222")
// 	fmt.Println(userToken)
// 	// 获取所有参数
// 	rqu := c.Ctx.Request.Header
// 	fmt.Println(rqu)
// 	rqu2 := c.Ctx.Request.Header["Authorization"]
// 	fmt.Println(rqu2)
// 	rqu3 := c.Ctx.Request.Header["User-Agent"]
// 	fmt.Println(rqu3)
// 	// 获取用户的usertoken
// 	rquri := c.Ctx.Request.Header["Usertoken"]
// 	fmt.Println(rquri)
// 	t2 := c.Ctx.Request.Response
// 	fmt.Println(t2)
// 	// 获取微信返回的授权码
// 	// code := c.GetString("code")
// }

// 获取网页授权登陆二维码
// func GetAuthQrUrl(ctx *gin.Context) {
// 	// 生成ticket
// 	ticket := wxgo.GenerateRandomTicket(20)
// 	// 生成授权地址
// 	redirect_url := "http://your_ip/wechat/accessusercode" // 微信授权后重定向地址,用于接收用户code
// 	scope := "snsapi_base"                                 //授权权限
// 	oauthUrl := wechat.Wx.GetOauth2CodeUrl(redirect_url, scope, ticket)
// 	// 将授权地址生成QR码
// 	savePath := "./resource/image"
// 	err := wxgo.GenerateQrCode(oauthUrl, savePath, ticket)
// 	if err != nil {
// 		log.Fatal(err.Error())
// 		return
// 	}
// 	qrUrl := fmt.Sprintf("http://your_ip/static/image/%s.png", ticket)
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"ticket": ticket,
// 		"qrUrl":  qrUrl,
// 	})
// }

// func GenerateQrCode(url string, savedir string, fname string) error {
// 	// "github.com/skip2/go-qrcode" 包
// 	qrcode, err := qrcode.New(url, qrcode.Highest)
// 	if err != nil {
// 		return err
// 	}
// 	qrcode.DisableBorder = true
// 	//保存成文件
// 	savepath := fmt.Sprintf("%s/%s.png", savedir, fname)
// 	err = qrcode.WriteFile(256, savepath)
// 	return err
// }

// 网页授权接收code
// func AccessUserCode(ctx *gin.Context) {
// 	code := ctx.Query("code")
// 	ticke := ctx.Query("state")
// 	// 用code获取用户access token
// 	uat, _ := wechat.Wx.GetUserATReq(code)
// 	// 用access token获取用户信息
// 	uInfo, _ := wechat.Wx.GetUserInfoReq(uat)
// 	log.Printf(uInfo.NickName, uInfo.Headimgurl, uInfo.City)
// 	// 将ticket和openid放入redis
// 	redis.RedisClient.Set(ticke, uInfo.OpenId, 60*time.Second)
// 	// 重定向到成功/失败页面
// 	ctx.Redirect(http.StatusTemporaryRedirect, "http://your_ip/static/html/loginsucceed.html")
// }

// 六、WebSocket实现扫码登录
// 为了实现PC端扫码登录，我们可以使用WebSocket技术。以下是一个简单的示例：

// func main() {
//     http.HandleFunc("/ws", websocketHandler)
//     http.ListenAndServe(":7777", nil)
// }

// import (
//     "github.com/gorilla/websocket"
//     "net/http"
// )

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// func websocketHandler(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		return
// 	}
// 	defer conn.Close()

// 	for {
// 		_, message, err := conn.ReadMessage()
// 		if err != nil {
// 			break
// 		}

// 		// 处理消息，例如验证token
// 		token := string(message)
// 		if token == "有效的token" {
// 			conn.WriteMessage(websocket.TextMessage, []byte("登录成功"))
// 		}
// 	}
// }
