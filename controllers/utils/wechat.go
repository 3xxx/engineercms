package utils

import (
	"encoding/json"
	// "encoding/xml"
	// "errors"
	"bytes"
	"encoding/base64"
	"fmt"
	// beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/cache"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	// "strconv"
	"context"
	"strings"
	"time"
)

type WxPushContent struct {
	Push_record_id   int    `json:"push_record_id"`
	Template_id      string `json:"template_id"`
	Page             string `json:"page"`
	Data             Data   `json:"data"`
	Emphasis_keyword string `json:"emphasis_keyword"`
}

type Data struct {
	Keyword1 WxKeyword `json:"keyword1"`
	Keyword2 WxKeyword `json:"keyword2"`
	Keyword3 WxKeyword `json:"keyword3"`
}

type WxKeyword struct {
	Value string `json:"value"`
}

var bm cache.Cache

const (
	//以下均为公众号管理后台设置项
	token          = "XXXXXXXX"
	appID          = "XXXXXXXXXX"
	encodingAESKey = "XXXXXXXXXXXXXXX"
)

var AesKey []byte

func EncodingAESKey2AESKey(encodingKey string) []byte {
	data, _ := base64.StdEncoding.DecodeString(encodingKey + "=")
	return data
}

func init() {
	var err error
	bm, err = cache.NewCache("memory", `{"interval":7200}`)
	if err != nil {
		logs.Error(err)
	}

	AesKey = EncodingAESKey2AESKey(encodingAESKey)
}

//验证token时效性
func GetAccessToken(app_version string) (accesstoken string, errcode float64, errmsg string, err error) {
	// category := BM.Get("category")
	//   if category == nil{
	//       return nil
	//   }
	//   var data []models.Category
	//   _ = json.Unmarshal(category.([]uint8), &data)
	//   return data
	// var accessToken cache.Cache
	resultToken, err := bm.Get(context.TODO(), "access_token")
	if err != nil {
		logs.Error(err)
	}
	if resultToken == nil {
		var APPID, SECRET string
		if app_version == "1" {
			APPID, err = web.AppConfig.String("wxAPPID")
			if err != nil {
				logs.Error(err)
			}
			SECRET, err = web.AppConfig.String("wxSECRET")
			if err != nil {
				logs.Error(err)
			}
		} else {
			appstring := "wxAPPID" + app_version
			APPID, err = web.AppConfig.String(appstring)
			if err != nil {
				logs.Error(err)
			}
			secretstring := "wxSECRET" + app_version
			SECRET, err = web.AppConfig.String(secretstring)
			if err != nil {
				logs.Error(err)
			}
		}
		// 重新获取，存入缓存，返回值
		accessToken, errcode, errmsg, err := requestToken(APPID, SECRET)
		if accessToken != "" {
			bm.Put(context.TODO(), "access_token", accessToken, 7200*time.Second)
			return accessToken, 0, "", nil
		} else {
			return "", errcode, errmsg, err
		}
		// timeoutDuration := 10000000 * time.Second
		// accessToken, err = cache.NewCache("memory", `{"interval":7200}`)
		// bm, err := bm.NewCache("memory", `{"interval":60}`)
	} else {
		// var data string
		// data = resultToken.(string)
		return resultToken.(string), 0, "", nil
	}
}

//请求获得accessToken
func requestToken(appid, secret string) (accessToken string, errcode float64, errmsg string, err error) {
	u, err := url.Parse("https://api.weixin.qq.com/cgi-bin/token")
	if err != nil {
		log.Fatal(err)
	}
	paras := &url.Values{}
	//设置请求参数
	paras.Set("appid", appid)
	paras.Set("secret", secret)
	paras.Set("grant_type", "client_credential")
	u.RawQuery = paras.Encode()
	resp, err := http.Get(u.String())
	//关闭资源
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", 0, "", err //errors.New("request token err :" + err.Error())
	}

	jMap := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&jMap)
	if err != nil {
		return "", 0, "", err //errors.New("request token response json parse err :" + err.Error())
	}
	if jMap["errcode"] == nil || jMap["errcode"].(float64) == 0 {
		accessToken, _ = jMap["access_token"].(string)
		return accessToken, 0, "", nil
	} else {
		//返回错误信息
		errcode := jMap["errcode"].(float64)
		errmsg := jMap["errmsg"].(string)
		// err = errors.New(errcode + ":" + errmsg)
		return "", errcode, errmsg, err
	}
}

type MsgInput struct {
	Openid  string `json:"openid"`
	Scene   int    `json:"scene"`
	Version int    `json:"version"`
	Content string `json:"content"`
}

type MsgOut struct {
	Errcode int         `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
	Result  MsgResult   `json:"result"`
	Detail  []MsgDetail `json:"detail"`
}
type MsgResult struct {
	Suggest string `json:"suggest"`
	Label   int    `json:"label"`
}
type MsgDetail struct {
	Strategy string `json:"strategy"`
	Errcode  int    `json:"errcode"`
	Suggest  string `json:"suggest"`
	Label    int    `json:"label"`
	Prob     int    `json:"prob"`
}

// 文本敏感字符检测
func MsgSecCheck(scene, version int, access_token, openid, content string) (suggest string, label int, err error) {
	// access_token = data["access_token"].(string)
	// https://api.weixin.qq.com/wxa/msg_sec_check?access_token=ACCESS_TOKEN
	requestUrl := "https://api.weixin.qq.com/wxa/msg_sec_check?access_token=" + access_token
	// 对于 POST 请求，部分参数需以 QueryString 的形式写在 URL 中（一般只有 access_token，如有额外参数会在文档里的 URL 中体现），其他参数如无特殊说明均以 JSON 字符串格式写在 POST 请求的 body 中。
	// "content":content
	// content = "特3456书yuuo莞6543李zxcz蒜7782法fgnv级"
	// utf := strconv.QuoteToASCII(content)
	// content = utf[1 : len(utf)-1]

	msg := MsgInput{
		Openid:  openid,
		Scene:   scene,
		Version: version,
		Content: content,
	}
	// b, err := json.Marshal(map[string]string{
	// 	"openid":  openid,
	// 	"scene":   scene,
	// 	"version": version,
	// 	"content": content,
	// })
	b, err := json.Marshal(msg)
	if err != nil {
		logs.Error(err)
		return "json转换错误", 0, err
	}
	// logs.Info(string(b))

	resp, err := http.Post(requestUrl, "application/x-www-form-urlencoded", bytes.NewBuffer(b)) //注意，这里是post
	if err != nil {
		logs.Error(err)
		return "post请求错误", 0, err
		// return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		logs.Error(err)
		return "请求返回错误", 0, err
	}
	var msgout MsgOut
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.Error(err)
		return "请求返回错误", 0, err
	}
	err = json.Unmarshal([]byte(body), &msgout)
	if err != nil {
		logs.Error(err)
		return "解析至json错误", 0, err
	}
	// logs.Info(msgout)
	return msgout.Result.Suggest, msgout.Result.Label, nil
	// var data map[string]interface{}
	// err = json.NewDecoder(resp.Body).Decode(&data)
	// if err != nil {
	// 	logs.Error(err)
	// 	return 0, "解析body错误", err
	// } else {
	// 	errcode := data["errcode"].(float64)
	// 	logs.Info(errcode)
	// 	errmsg := data["errmsg"].(string)
	// 	// err = json.Unmarshal(data["result"], &MsgResult)
	// 	// result := data["result"].(string)
	// 	// detail := data["detail"].(string)
	// 	// logs.Info(detail)
	// 	return errcode, errmsg, nil
	// }
}

// 图片检测

// type ReqImgCheck struct {
// 	Media *multipart.FileHeader `json:"media"`
// }

//图片检测
// header设置为"Content-Type","application/octet-stream"
// func ImgCheck(media *multipart.FileHeader) error {
// 	req.Media = media
// 	url := fmt.Sprintf("https://api.weixin.qq.com/wxa/img_sec_check?access_token=%s", ak)
// 	bt, err := json.Marshal(req)
// 	if err != nil {

// 	}
// 	body, err := util.PostCurl(url, bt, util.JSONHeader)
// 	if err != nil {

// 	}
// }

func ImgSecCheck(bts []byte, accessToken string) (errcode float64, errmsg string, err error) {
	// beego.Info(accessToken)
	var bufReader bytes.Buffer
	//	"mime/multipart" 可以将上传文件封装
	mpWriter := multipart.NewWriter(&bufReader)
	//文件名无所谓
	fileName := "detect"
	//字段名必须为media
	writer, err := mpWriter.CreateFormFile("media", fileName)
	if err != nil {
		logs.Error(err)
		return 0, "建立media错误", err
	}
	reader := bytes.NewReader(bts)
	io.Copy(writer, reader)
	//关闭了才能把内容写入
	mpWriter.Close()

	client := http.DefaultClient
	destURL := "https://api.weixin.qq.com/wxa/img_sec_check?access_token=" + accessToken
	req, _ := http.NewRequest("POST", destURL, &bufReader)
	//从mpWriter中获取content-Type
	req.Header.Set("Content-Type", mpWriter.FormDataContentType())
	req.Header.Set("Content-Type", "application/octet-stream")
	// request.addHeader("Content-Type", "application/octet-stream");
	// InputStream inputStream = multipartFile.getInputStream();
	// byte[] byt = new byte[inputStream.available()];
	// inputStream.read(byt);
	// request.setEntity(new ByteArrayEntity(byt, ContentType.create("image/jpg")));
	// response = httpclient.execute(request);
	// HttpEntity httpEntity = response.getEntity();
	// String result = EntityUtils.toString(httpEntity, "UTF-8");// 转成string
	// JSONObject jsonObject = JSONObject.parseObject(result);
	// return jsonObject;

	resp, err := client.Do(req)
	if err != nil {
		logs.Error(err)
		return 0, "请求错误", err
	}
	defer resp.Body.Close()

	vs := make(map[string]interface{})
	result, _ := ioutil.ReadAll(resp.Body)
	if err != nil {
		// return false
		logs.Error(err)
		return 0, "读取返回值错误", err
	}
	err = json.Unmarshal(result, &vs)
	if err != nil {
		// return false
		logs.Error(err)
		return 0, "解析json错误", err
	}
	//errcode 存在，且为0，返回通过，87014有风险
	// errmsg "ok","risky content"
	if _, ok := vs["errcode"]; ok {
		// return true
		return vs["errcode"].(float64), vs["errmsg"].(string), nil
	} else {
		return 0, "无返回code", nil
	}
}

type phrase1 struct {
	Value string `json:"value"`
}

//注意下面是date，不是data
type date2 struct {
	Value string `json:"value"`
}
type thing3 struct {
	Value string `json:"value"`
}
type data struct {
	Phrase1 phrase1 `json:"phrase1"`
	Date2   date2   `json:"date2"`
	Thing3  thing3  `json:"thing3"`
}

type Message struct {
	Touser      string `json:"touser"`
	Template_id string `json:"template_id"`
	Page        string `json:"page"`
	Data        data   `json:"data"`
}

// 发送订阅消息
func SendMessage(access_token, openid, template_id string) (errcode float64, errmsg string, err error) {
	requestUrl := "https://api.weixin.qq.com/cgi-bin/message/subscribe/send?access_token=" + access_token
	// 对于 POST 请求，部分参数需以 QueryString 的形式写在 URL 中（一般只有 access_token，如有额外参数会在文档里的 URL 中体现），其他参数如无特殊说明均以 JSON 字符串格式写在 POST 请求的 body 中。
	// "content":content
	// 打卡类型
	// {{phrase1.DATA}}

	// 打卡时间
	// {{date2.DATA}}

	// 打卡地点
	// {{thing3.DATA}}

	//***********这个也行
	// b := `{
	// 	 	"touser":"opl4B5YvCVXaatKF6VGUSbohMlWQ",
	// 	 	"template_id":"c8Dped7TztDiMJ2bzHMu4G3nikn-mSOHmQEh_7aTlLo",
	// 	 	"page":"index",
	// 	 	"data":{
	// 	 		"phrase1":{"value":"设代打卡"},
	// 	 		"date2":{"value":"8:30"},
	// 	 		"thing3":{"value":"顺德、南沙、东莞、深圳"}
	// 	 		}
	// 	}`
	// resp, err := http.Post(requestUrl, "application/json;charset=utf-8", bytes.NewBuffer([]byte(b))) //注意，这里是post
	//*************

	// phrase1 := map[string]string{
	// 	"value": "设代打卡",
	// }
	// date2 := map[string]string{
	// 	"value": "8:30",
	// }
	// thing3 := map[string]string{
	// 	"value": "顺德、南沙、东莞、深圳",
	// }

	// c := map[string]string{
	// 	"touser":      opneid,
	// 	"template_id": template_id,
	// 	"page":        "index",
	// 	"data":        b,
	// }

	var phrase11 phrase1
	phrase11.Value = "设代打卡"
	var date21 date2
	date21.Value = "8:30"
	var thing31 thing3
	thing31.Value = "顺德、南沙、东莞、深圳"
	var data1 data
	data1.Phrase1 = phrase11
	data1.Date2 = date21
	data1.Thing3 = thing31
	var message Message
	message.Data = data1
	message.Template_id = template_id
	message.Page = "packageA/pages/search/search"
	message.Touser = openid
	wxmessage, err := json.Marshal(message)
	if err != nil {
		logs.Error(err)
		return 0, "json转换错误", err
	}
	// beego.Info(message)
	// fmt.Println("buf = ", string(wxmessage))

	// jsons,_:=json.Marshal(teamworkinfo)
	// result := string(wxmessage)
	// jsoninfo := strings.NewReader(result)——这个不是以body形式传递

	// resp, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(wxmessage))
	// req.Header.Add("appCode", "winner")
	// req.Header.Add("token", "466221e4-593d-4bb8-b41b-hcfedhf")//应用对接的token
	// resp.Header.Add("Content-Type", "application/json")

	//估计下面的"application/json;charset=utf-8"不影响结果吧，没有试
	resp, err := http.Post(requestUrl, "application/json;charset=utf-8", bytes.NewBuffer(wxmessage))
	if err != nil {
		logs.Error(err)
		return 0, "post请求错误", err
		// return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		logs.Error(err)
		return 0, "请求返回错误", err
	}
	var msgdata map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&msgdata)
	if err != nil {
		logs.Error(err)
		return 0, "解析body错误", err
	} else {
		errcode := msgdata["errcode"].(float64)
		// beego.Info(errcode)
		errmsg := msgdata["errmsg"].(string)
		return errcode, errmsg, nil
	}
}

func testmain() {

	mstr := map[string]string{}

	mstr["key"] = "chatengine"
	mstr["conn"] = ":6379"
	mstr["dbNum"] = "0"

	bytes, _ := json.Marshal(mstr)

	//从redis缓存中拿数据拿数据
	cache_conn, err := cache.NewCache("redis", string(bytes))
	if err != nil {
		fmt.Println(err)
	}

	timeoutDuration := 10000000 * time.Second

	err = cache_conn.Put(context.TODO(), "wilson1231111", "xu", timeoutDuration*time.Second)
	if err != nil {
		fmt.Println("数据读取出错，错误为：", err)
	} else {
		fmt.Println("redis 读取正常")
	}
	cache_areardata, err := cache_conn.Get(context.TODO(), "area")
	if err != nil {
		fmt.Println("数据读取出错，错误为：", err)
	}
	if areaData := cache_areardata; areaData != nil {
		// beego.Info("get data from cache===========")
		//resp["data"] = areaData
		fmt.Println("从redis中读取出的数据为：", areaData)
	} else {
		fmt.Println("需要读取MySQL")
	}

	fmt.Println(strings.Index("i am a good man", "haha"))
}

// func checkSignature() {
// 	//"r"为*http.Request
// 	var r *http.Request
// 	r.ParseForm()

// 	timestamp := strings.Join(r.Form["timestamp"], "")
// 	nonce := strings.Join(r.Form["nonce"], "")
// 	signature := strings.Join(r.Form["signature"], "")
// 	encryptType := strings.Join(r.Form["encrypt_type"], "")
// 	msgSignature := strings.Join(r.Form["msg_signature"], "")

// 	if !wechat.ValidateUrl(timestamp, nonce, signature) {
// 		wmm.log.AppendObj(nil, "Wechat Message Service: this http request is not from Wechat platform!")
// 		return
// 	}

// 	//微信安全模式更改/首次添加url，需要验证，将参数原样写会即可
// 	var es string
// 	if e = req.ParseGet("echostr", &es); e != nil {

// 	}
// 	if es != "" {
// 		//todo 将参数值es原先写回即可
// 		return
// 	}

// 	if r.Method == "POST" {
// 		if encryptType == "aes" {
// 			encryptRequestBody := wechat.ParseEncryptRequestBody(r)
// 			// Validate mstBody signature
// 			if !wechat.ValidateMsg(timestamp, nonce, encryptRequestBody.Encrypt, msgSignature) {
// 				return
// 			}

// 			// Decode base64
// 			cipherData, err := base64.StdEncoding.DecodeString(encryptRequestBody.Encrypt)
// 			if err != nil {
// 				return
// 			}

// 			// AES Decrypt
// 			plainData, err := wechat.AesDecrypt(cipherData, wechat.AesKey)
// 			if err != nil {
// 				return
// 			}

// 			//封装struct
// 			textRequestBody, err := wechat.ParseEncryptTextRequestBody(plainData)
// 			if err != nil {
// 				return
// 			}

// 			var url string

// 			tp := textRequestBody.MsgType
// 			if tp == "text" && textRequestBody.Content == "【收到不支持的消息类型，暂无法显示】" {
// 				//安全模式下向用户回复消息也需要加密
// 				respBody, e := wechat.MakeEncryptResponseBody(textRequestBody.ToUserName, textRequestBody.FromUserName, "一些回复给用户的消息", nonce, timestamp)
// 				if e != nil {
// 					return e
// 				}
// 				//此处return NewSimpleError是一个对返回值处理的封装，返回xml格式消息，并不是返回错误
// 				return service.NewSimpleError(service.SERVER_WRITE_XML, string(respBody))

// 			}
// 			if tp == "event" {
// 				//某个类型的消息暂时后台不作处理，也需要向微信服务器做出响应
// 				return service.NewSimpleError(service.SERVER_WRITE_TEXT, "success")
// 			}
// 		}
// 		return service.NewSimpleError(service.SERVER_WRITE_TEXT, "success")
// 	}
// }

// 附：微信消息加解密工具包（GoLang）

// package wechat
//微信消息加解密工具包

// func init() {
//     AesKey = EncodingAESKey2AESKey(encodingAESKey)
// }

// type TextRequestBody struct {
// 	XMLName      xml.Name `xml:"xml"`
// 	ToUserName   string
// 	FromUserName string
// 	CreateTime   time.Duration
// 	MsgType      string
// 	Url          string
// 	PicUrl       string
// 	MediaId      string
// 	ThumbMediaId string
// 	Content      string
// 	MsgId        int
// 	Location_X   string
// 	Location_Y   string
// 	Label        string
// }

// type TextResponseBody struct {
// 	XMLName      xml.Name `xml:"xml"`
// 	ToUserName   CDATAText
// 	FromUserName CDATAText
// 	CreateTime   string
// 	MsgType      CDATAText
// 	Content      CDATAText
// }

// type EncryptRequestBody struct {
// 	XMLName    xml.Name `xml:"xml"`
// 	ToUserName string
// 	Encrypt    string
// }

// type EncryptResponseBody struct {
// 	XMLName      xml.Name `xml:"xml"`
// 	Encrypt      CDATAText
// 	MsgSignature CDATAText
// 	TimeStamp    string
// 	Nonce        CDATAText
// }

// type EncryptResponseBody1 struct {
// 	XMLName      xml.Name `xml:"xml"`
// 	Encrypt      string
// 	MsgSignature string
// 	TimeStamp    string
// 	Nonce        string
// }

// type CDATAText struct {
// 	Text string `xml:",innerxml"`
// }

// func MakeSignature(timestamp, nonce string) string {
// 	sl := []string{token, timestamp, nonce}
// 	sort.Strings(sl)
// 	s := sha1.New()
// 	io.WriteString(s, strings.Join(sl, ""))
// 	return fmt.Sprintf("%x", s.Sum(nil))
// }

// func MakeMsgSignature(timestamp, nonce, msg_encrypt string) string {
// 	sl := []string{token, timestamp, nonce, msg_encrypt}
// 	sort.Strings(sl)
// 	s := sha1.New()
// 	io.WriteString(s, strings.Join(sl, ""))
// 	return fmt.Sprintf("%x", s.Sum(nil))
// }

// func ValidateUrl(timestamp, nonce, signatureIn string) bool {
// 	signatureGen := MakeSignature(timestamp, nonce)
// 	if signatureGen != signatureIn {
// 		return false
// 	}
// 	return true
// }

// func ValidateMsg(timestamp, nonce, msgEncrypt, msgSignatureIn string) bool {
// 	msgSignatureGen := MakeMsgSignature(timestamp, nonce, msgEncrypt)
// 	if msgSignatureGen != msgSignatureIn {
// 		return false
// 	}
// 	return true
// }

// func ParseEncryptRequestBody(r *http.Request) *EncryptRequestBody {
// 	body, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		return nil
// 	}
// 	//  mlog.AppendObj(nil, "Wechat Message Service: RequestBody--", body)
// 	requestBody := &EncryptRequestBody{}
// 	xml.Unmarshal(body, requestBody)
// 	return requestBody

// }

// func ParseTextRequestBody(r *http.Request) *TextRequestBody {
// 	body, err := ioutil.ReadAll(r.Body)
// 	r.Body.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 		return nil
// 	}
// 	requestBody := &TextRequestBody{}
// 	xml.Unmarshal(body, requestBody)
// 	return requestBody
// }

// func Value2CDATA(v string) CDATAText {
// 	//return CDATAText{[]byte("<![CDATA[" + v + "]]>")}
// 	return CDATAText{"<![CDATA[" + v + "]]>"}
// }

// func MakeTextResponseBody(fromUserName, toUserName, content string) ([]byte, error) {
// 	textResponseBody := &TextResponseBody{}
// 	textResponseBody.FromUserName = Value2CDATA(fromUserName)
// 	textResponseBody.ToUserName = Value2CDATA(toUserName)
// 	textResponseBody.MsgType = Value2CDATA("text")
// 	textResponseBody.Content = Value2CDATA(content)
// 	textResponseBody.CreateTime = strconv.Itoa(int(time.Duration(time.Now().Unix())))
// 	return xml.MarshalIndent(textResponseBody, " ", "  ")
// }
// func MakeEncryptResponseBody(fromUserName, toUserName, content, nonce, timestamp string) ([]byte, error) {
// 	encryptBody := &EncryptResponseBody{}

// 	encryptXmlData, _ := MakeEncryptXmlData(fromUserName, toUserName, timestamp, content)
// 	encryptBody.Encrypt = Value2CDATA(encryptXmlData)
// 	encryptBody.MsgSignature = Value2CDATA(MakeMsgSignature(timestamp, nonce, encryptXmlData))
// 	encryptBody.TimeStamp = timestamp
// 	encryptBody.Nonce = Value2CDATA(nonce)

// 	return xml.MarshalIndent(encryptBody, " ", "  ")
// }

// func MakeEncryptXmlData(fromUserName, toUserName, timestamp, content string) (string, error) {
// 	textResponseBody := &TextResponseBody{}
// 	textResponseBody.FromUserName = Value2CDATA(fromUserName)
// 	textResponseBody.ToUserName = Value2CDATA(toUserName)
// 	textResponseBody.MsgType = Value2CDATA("text")
// 	textResponseBody.Content = Value2CDATA(content)
// 	textResponseBody.CreateTime = timestamp
// 	body, err := xml.MarshalIndent(textResponseBody, " ", "  ")
// 	if err != nil {
// 		return "", errors.New("xml marshal error")
// 	}

// 	buf := new(bytes.Buffer)
// 	err = binary.Write(buf, binary.BigEndian, int32(len(body)))
// 	if err != nil {
// 		mlog.AppendObj(err, "Binary write err:", err)
// 	}
// 	bodyLength := buf.Bytes()

// 	randomBytes := []byte("abcdefghijklmnop")

// 	plainData := bytes.Join([][]byte{randomBytes, bodyLength, body, []byte(appID)}, nil)
// 	cipherData, err := AesEncrypt(plainData, AesKey)
// 	if err != nil {
// 		return "", errors.New("AesEncrypt error")
// 	}
// 	return base64.StdEncoding.EncodeToString(cipherData), nil
// }

// PadLength calculates padding length, from github.com/vgorin/cryptogo
// func PadLength(slice_length, blocksize int) (padlen int) {
// 	padlen = blocksize - slice_length%blocksize
// 	if padlen == 0 {
// 		padlen = blocksize
// 	}
// 	return padlen
// }

//from github.com/vgorin/cryptogo
// func PKCS7Pad(message []byte, blocksize int) (padded []byte) {
// 	// block size must be bigger or equal 2
// 	if blocksize < 1<<1 {
// 		panic("block size is too small (minimum is 2 bytes)")
// 	}
// 	// block size up to 255 requires 1 byte padding
// 	if blocksize < 1<<8 {
// 		// calculate padding length
// 		padlen := PadLength(len(message), blocksize)

// 		// define PKCS7 padding block
// 		padding := bytes.Repeat([]byte{byte(padlen)}, padlen)

// 		// apply padding
// 		padded = append(message, padding...)
// 		return padded
// 	}
// 	// block size bigger or equal 256 is not currently supported
// 	panic("unsupported block size")
// }

// func AesEncrypt(plainData []byte, aesKey []byte) ([]byte, error) {
// 	k := len(aesKey)
// 	if len(plainData)%k != 0 {
// 		plainData = PKCS7Pad(plainData, k)
// 	}

// 	block, err := aes.NewCipher(aesKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	iv := make([]byte, aes.BlockSize)
// 	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
// 		return nil, err
// 	}

// 	cipherData := make([]byte, len(plainData))
// 	blockMode := cipher.NewCBCEncrypter(block, iv)
// 	blockMode.CryptBlocks(cipherData, plainData)

// 	return cipherData, nil
// }

// func AesDecrypt(cipherData []byte, aesKey []byte) ([]byte, error) {
// 	k := len(aesKey) //PKCS#7
// 	if len(cipherData)%k != 0 {
// 		return nil, errors.New("crypto/cipher: ciphertext size is not multiple of aes key length")
// 	}

// 	block, err := aes.NewCipher(aesKey)
// 	if err != nil {
// 		return nil, err
// 	}

// 	iv := make([]byte, aes.BlockSize)
// 	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
// 		return nil, err
// 	}
// 	blockMode := cipher.NewCBCDecrypter(block, iv)
// 	plainData := make([]byte, len(cipherData))
// 	blockMode.CryptBlocks(plainData, cipherData)
// 	return plainData, nil
// }

// func ValidateAppId(id []byte) bool {
// 	if string(id) == appID {
// 		return true
// 	}
// 	return false
// }

// func ParseEncryptTextRequestBody(plainText []byte) (*TextRequestBody, error) {

// 	// Read length
// 	buf := bytes.NewBuffer(plainText[16:20])
// 	var length int32
// 	binary.Read(buf, binary.BigEndian, &length)

// 	// appID validation
// 	appIDstart := 20 + length
// 	id := plainText[appIDstart : int(appIDstart)+len(appID)]
// 	if !ValidateAppId(id) {
// 		mlog.AppendObj(nil, "Wechat Message Service: appid is invalid!")
// 		return nil, errors.New("Appid is invalid")
// 	}
// 	mlog.AppendObj(nil, "Wechat Message Service: appid validation is ok!")

// 	textRequestBody := &TextRequestBody{}
// 	xml.Unmarshal(plainText[20:20+length], textRequestBody)
// 	return textRequestBody, nil
// }

// func ParseEncryptResponse(responseEncryptTextBody []byte) {
// 	textResponseBody := &EncryptResponseBody1{}
// 	xml.Unmarshal(responseEncryptTextBody, textResponseBody)

// 	if !ValidateMsg(textResponseBody.TimeStamp, textResponseBody.Nonce, textResponseBody.Encrypt, textResponseBody.MsgSignature) {
// 		mlog.AppendInfo("msg signature is invalid")
// 		return
// 	}

// 	cipherData, err := base64.StdEncoding.DecodeString(textResponseBody.Encrypt)
// 	if err != nil {
// 		mlog.AppendObj(err, "Wechat Message Service: Decode base64 error")
// 		return
// 	}

// 	plainText, err := AesDecrypt(cipherData, AesKey)
// 	if err != nil {
// 		mlog.AppendInfo(err)
// 		return
// 	}

// 	mlog.AppendInfo(string(plainText))
// }

// func DecryptWechatAppletUser(encryptedData string, session_key string, iv string) ([]byte, error) {
// 	ciphertext, _ := base64.StdEncoding.DecodeString(encryptedData)
// 	key, _ := base64.StdEncoding.DecodeString(session_key)
// 	keyBytes := []byte(key)
// 	block, err := aes.NewCipher(keyBytes) //选择加密算法
// 	if err != nil {
// 		return nil, err
// 	}
// 	iv_b, _ := base64.StdEncoding.DecodeString(iv)
// 	blockModel := cipher.NewCBCDecrypter(block, iv_b)
// 	plantText := make([]byte, len(ciphertext))
// 	blockModel.CryptBlocks(plantText, ciphertext)
// 	plantText = PKCS7UnPadding(plantText, block.BlockSize())
// 	return plantText, nil
// }

// func PKCS7UnPadding(plantText []byte, blockSize int) []byte {
// 	length := len(plantText)
// 	unpadding := int(plantText[length-1])
// 	return plantText[:(length - unpadding)]
// }
