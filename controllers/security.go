package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	// 	"crypto/aes"
	// 	"crypto/cipher"
	// 	"encoding/base64"
	// "encoding/json"
	// "errors"
	// 	"fmt"
	// "github.com/xlstudio/wxbizdatacrypt"
	// "regexp"
)

// DecodeWeAppUserInfo 解密微信小程序用户信息
func DecodeWeAppUserInfo(encryptedData string, sessionKey string, iv string) (string, error) {
	cipher, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		fmt.Println("encryptedData: ", encryptedData, "\n", err.Error())
		return "", err
	}

	key, keyErr := base64.StdEncoding.DecodeString(sessionKey)
	if keyErr != nil {
		fmt.Println("sessionKey: ", sessionKey, "\n", keyErr.Error())
		return "", keyErr
	}

	theIV, ivErr := base64.StdEncoding.DecodeString(iv)
	if ivErr != nil {
		fmt.Println("iv: ", iv, "\n", ivErr.Error())
		return "", ivErr
	}

	result, resultErr := AESDecrypt(cipher, key, theIV)
	if resultErr != nil {
		return "", resultErr
	}
	return string(result), nil
}

// AESDecrypt AES解密
func AESDecrypt(cipherBytes, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockModel := cipher.NewCBCDecrypter(block, iv)
	dst := make([]byte, len(cipherBytes))
	blockModel.CryptBlocks(dst, cipherBytes)
	dst = PKCS7UnPadding(dst, block.BlockSize())
	return dst, nil
}

// PKCS7UnPadding pkcs7填充方式
func PKCS7UnPadding(dst []byte, blockSize int) []byte {
	length := len(dst)
	unpadding := int(dst[length-1])
	return dst[:(length - unpadding)]
}

// package wechat

// import (
// 	"encoding/json"
// 	"errors"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"net/url"
// )
// 我们可以使用go的ticker来设置7000秒请求一次来保证token的持续可用
// func init() {
// 	appid := ""
// 	secret := ""
// 	freshTokenTicker := time.NewTicker(7000 * time.Second)
// 	//requestToken()
// 	go func() {
// 		select {
// 		case <-freshTokenTicker.C:
// 			{
// 				accessToken, err := requestToken(appid, secret)
// 				if err != nil {
// 					//TODO 错误处理
// 				}
// 				log.Printf("token refresh :%s", accessToken)
// 			}
// 		}
// 	}()
// }

// func requestToken(appid, secret string) (string, error) {
// 	u, err := url.Parse("https://api.weixin.qq.com/cgi-bin/token")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	paras := &url.Values{}
// 	//设置请求参数
// 	paras.Set("appid", appid)
// 	paras.Set("secret", secret)
// 	paras.Set("grant_type", "client_credential")
// 	u.RawQuery = paras.Encode()
// 	resp, err := http.Get(u.String())
// 	//关闭资源
// 	if resp != nil && resp.Body != nil {
// 		defer resp.Body.Close()
// 	}
// 	if err != nil {
// 		return "", errors.New("request token err :" + err.Error())
// 	}
// 	jMap := make(map[string]interface{})
// 	err = json.NewDecoder(resp.Body).Decode(&jMap)
// 	if err != nil {
// 		return "", errors.New("request token response json parse err :" + err.Error())
// 	}
// 	if jMap["errcode"] == nil || jMap["errcode"] == 0 {
// 		accessToken, _ := jMap["access_token"].(string)
// 		return accessToken, nil
// 	} else {
// 		//返回错误信息
// 		errcode := jMap["errcode"].(string)
// 		errmsg := jMap["errmsg"].(string)
// 		err = errors.New(errcode + ":" + errmsg)
// 		return "", err
// 	}
// }

// //返回的map可以替换为专门的结构体
// func WechatLogin(js_code, appid, secret string) (map[string]interface{}, error) {
// 	Code2SessURL := "https://api.weixin.qq.com/sns/jscode2session?appid={appid}&secret={secret}&js_code={code}&grant_type=authorization_code"
// 	Code2SessURL = strings.Replace(Code2SessURL, "{appid}", appid, -1)
// 	Code2SessURL = strings.Replace(Code2SessURL, "{secret}", secret, -1)
// 	Code2SessURL = strings.Replace(Code2SessURL, "{code}", js_code, -1)
// 	resp, err := http.Get(Code2SessURL)
// 	//关闭资源
// 	if resp != nil && resp.Body != nil {
// 		defer resp.Body.Close()
// 	}
// 	if err != nil {
// 		return nil, errors.New("WechatLogin request err :" + err.Error())
// 	}

// 	var jMap map[string]interface{}
// 	err = json.NewDecoder(resp.Body).Decode(&jMap)
// 	if err != nil {
// 		return nil, errors.New("request token response json parse err :" + err.Error())
// 	}
// 	if jMap["errcode"] == nil || jMap["errcode"] == 0 {
// 		return jMap, nil
// 	} else {
// 		//返回错误信息
// 		errcode := jMap["errcode"].(string)
// 		errmsg := jMap["errmsg"].(string)
// 		err = errors.New(errcode + ":" + errmsg)
// 		return nil, err
// 	}
// }

// func DecodeWeAppUserInfotest() {
// 	appID := "wx4f4bc4dec97d474b"
// 	sessionKey := "tiihtNczf5v6AKRyjwEUhQ=="
// 	encryptedData := "CiyLU1Aw2KjvrjMdj8YKliAjtP4gsMZMQmRzooG2xrDcvSnxIMXFufNstNGTyaGS9uT5geRa0W4oTOb1WT7fJlAC+oNPdbB+3hVbJSRgv+4lGOETKUQz6OYStslQ142dNCuabNPGBzlooOmB231qMM85d2/fV6ChevvXvQP8Hkue1poOFtnEtpyxVLW1zAo6/1Xx1COxFvrc2d7UL/lmHInNlxuacJXwu0fjpXfz/YqYzBIBzD6WUfTIF9GRHpOn/Hz7saL8xz+W//FRAUid1OksQaQx4CMs8LOddcQhULW4ucetDf96JcR3g0gfRK4PC7E/r7Z6xNrXd2UIeorGj5Ef7b1pJAYB6Y5anaHqZ9J6nKEBvB4DnNLIVWSgARns/8wR2SiRS7MNACwTyrGvt9ts8p12PKFdlqYTopNHR1Vf7XjfhQlVsAJdNiKdYmYVoKlaRv85IfVunYzO0IKXsyl7JCUjCpoG20f0a04COwfneQAGGwd5oa+T8yO5hzuyDb/XcxxmK01EpqOyuxINew=="
// 	iv := "r7BXXKkLb8qrSNn05n0qiA=="

// 	pc := wxbizdatacrypt.WxBizDataCrypt{AppID: appID, SessionKey: sessionKey}
// 	result, err := pc.Decrypt(encryptedData, iv, true) //第三个参数解释： 需要返回 JSON 数据类型时 使用 true, 需要返回 map 数据类型时 使用 false
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(result)
// 	}
// }

// package wxbizdatacrypt

// import (
// 	"crypto/aes"
// 	"crypto/cipher"
// 	"encoding/base64"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"regexp"
// )

// var errorCode = map[string]int{
// 	"IllegalAppID":      -41000,
// 	"IllegalAesKey":     -41001,
// 	"IllegalIV":         -41002,
// 	"IllegalBuffer":     -41003,
// 	"DecodeBase64Error": -41004,
// 	"DecodeJsonError":   -41005,
// }

// // WxBizDataCrypt represents an active WxBizDataCrypt object
// type WxBizDataCrypt struct {
// 	AppID      string
// 	SessionKey string
// }

// type showError struct {
// 	errorCode int
// 	errorMsg  error
// }

// func (e showError) Error() string {
// 	return fmt.Sprintf("{code: %v, error: \"%v\"}", e.errorCode, e.errorMsg)
// }

// // Decrypt Weixin APP's AES Data
// // If isJSON is true, Decrypt return JSON type.
// // If isJSON is false, Decrypt return map type.
// func (wxCrypt *WxBizDataCrypt) Decrypt(encryptedData string, iv string, isJSON bool) (interface{}, error) {
// 	if len(wxCrypt.SessionKey) != 24 {
// 		return nil, showError{errorCode["IllegalAesKey"], errors.New("sessionKey length is error")}
// 	}
// 	aesKey, err := base64.StdEncoding.DecodeString(wxCrypt.SessionKey)
// 	if err != nil {
// 		return nil, showError{errorCode["DecodeBase64Error"], err}
// 	}

// 	if len(iv) != 24 {
// 		return nil, showError{errorCode["IllegalIV"], errors.New("iv length is error")}
// 	}
// 	aesIV, err := base64.StdEncoding.DecodeString(iv)
// 	if err != nil {
// 		return nil, showError{errorCode["DecodeBase64Error"], err}
// 	}

// 	aesCipherText, err := base64.StdEncoding.DecodeString(encryptedData)
// 	if err != nil {
// 		return nil, showError{errorCode["DecodeBase64Error"], err}
// 	}
// 	aesPlantText := make([]byte, len(aesCipherText))

// 	aesBlock, err := aes.NewCipher(aesKey)
// 	if err != nil {
// 		return nil, showError{errorCode["IllegalBuffer"], err}
// 	}

// 	mode := cipher.NewCBCDecrypter(aesBlock, aesIV)
// 	mode.CryptBlocks(aesPlantText, aesCipherText)
// 	aesPlantText = PKCS7UnPadding(aesPlantText)

// 	var decrypted map[string]interface{}

// 	re := regexp.MustCompile(`[^\{]*(\{.*\})[^\}]*`)
// 	aesPlantText = []byte(re.ReplaceAllString(string(aesPlantText), "$1"))

// 	err = json.Unmarshal(aesPlantText, &decrypted)
// 	if err != nil {
// 		return nil, showError{errorCode["DecodeJsonError"], err}
// 	}

// 	if decrypted["watermark"].(map[string]interface{})["appid"] != wxCrypt.AppID {
// 		return nil, showError{errorCode["IllegalAppID"], errors.New("appID is not match")}
// 	}

// 	if isJSON == true {
// 		return string(aesPlantText), nil
// 	}

// 	return decrypted, nil
// }

// // PKCS7UnPadding return unpadding []Byte plantText
// func PKCS7UnPadding(plantText []byte) []byte {
// 	length := len(plantText)
// 	unPadding := int(plantText[length-1])
// 	return plantText[:(length - unPadding)]
// }
