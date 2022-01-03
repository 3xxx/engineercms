package utils

import (
	// "crypto/md5"
	"errors"
	"fmt"
	// "github.com/3xxx/go-sso/models"
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	"encoding/base64"
	"encoding/json"
	// "github.com/dgrijalva/jwt-go"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	KEY                    string = "JWT-ARY-STARK"
	DEFAULT_EXPIRE_SECONDS int    = 3000 //600 // default 10 minutes
	//SecretKey = "welcome to wangshubo's blog"
	SecretKey = "I have login"
	secret    = "I have login"
)

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct {
	Data string `json:"data"`
}

type Token struct {
	Token string `json:"token"`
}

type User struct {
	Id       string `json:"id"`
	Name     string `json:"json"`
	Password string `json:"password"`
}

// JWT -- json web token
// HEADER PAYLOAD SIGNATURE
// This struct is the PAYLOAD
type MyCustomClaims struct {
	User
	jwt.StandardClaims
}

//使用这个生成token
func CreateToken(userName string) (tokenString string, err error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	//添加令牌关键信息
	Tokenexp, err := strconv.Atoi(beego.AppConfig.String("Tokenexp"))
	if err != nil {
		return tokenString, err
	}
	//添加令牌期限
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(Tokenexp)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["userName"] = userName
	token.Claims = claims
	tokenString, err = token.SignedString([]byte(beego.AppConfig.String("TokenSecrets")))
	if err != nil {
		fmt.Println("generate json web token failed !! error :", err)
		return tokenString, err
	}
	return tokenString, err
}

// 校验token是否有效 返回参数
func CheckToken(tokenString string) (userName string, err error) {
	// tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"
	secret := beego.AppConfig.String("TokenSecrets")
	// secret = "AllYourBase"
	// userName := ""
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})
	if err != nil {
		beego.Error(err)
		return "", err
	}
	// token.Valid里已经包含了过期判断
	if token != nil && token.Valid {
		// fmt.Println("You look nice today")
		claims, _ := token.Claims.(jwt.MapClaims)
		userName = claims["userName"].(string)
	}

	// Token from another example.  This token is expired
	// token, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	//  return []byte("secret"), nil
	// })
	// if token.Valid {
	//  fmt.Println("You look nice today")
	// } else if ve, ok := err.(*jwt.ValidationError); ok {
	//  if ve.Errors&jwt.ValidationErrorMalformed != 0 {
	//      fmt.Println("That's not even a token")
	//  } else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
	//      // Token is either expired or not active yet
	//      fmt.Println("Timing is everything")
	//  } else {
	//      fmt.Println("Couldn't handle this token:", err)
	//  }
	// } else {
	//  fmt.Println("Couldn't handle this token:", err)
	// }
	return userName, err
}

// 校验token是否有效 返回参数
func LubanCheckToken(tokenString string) (userId, userName, userNickname string, err error) {
	// tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIiLCJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.HE7fK0xOQwFEr4WDgRWj4teRPZ6i3GLwD5YCm6Pwu_c"
	mySignKey := beego.AppConfig.String("LubanTokenSecrets") //坑：密钥必须要长，要达到这个位数，26个英文字母是不行的。！！！！
	// beego.Info(secret)
	// mySignKey := "whatthefuck123weishenmebuneng123" //密钥，同java代码，
	// mySignKeyBytes, err := base64.URLEncoding.DecodeString(mySignKey) //需要用和加密时同样的方式转化成对应的字节数组
	mySignKeyBytes, err := base64.RawStdEncoding.DecodeString(mySignKey)
	if err != nil {
		fmt.Println("base64 decodeString failed.", err)
		return "", "", "", err
	}

	// 原文链接：https://blog.csdn.net/qiang527052/article/details/80748700

	// secret = "AllYourBase"
	// userName := ""
	// tokenString = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJkb3VibGVzIn0.cBDFVLuQZV2B7J76kk17LE2hmni_3RbzTBzIH_OsriE"
	// tokenString = "eyJhbGciOiJIUzI1NiJ9.eyJqdGkiOiI4OCIsInN1YiI6IuWwj-eZvSIsImlhdCI6MTYzMzAxNjU0OSwiY29tcGFueUlkIjoiOTk5OTk5OSIsImNvbXBhbnlOYW1lIjoi6bKB54-tIn0.W9p52lAzo9EcNSb_Uwf7MK9Lw-vNvG_8HHI-piMkNpY"
	tokenString = SubString(tokenString, 4, len(tokenString))
	beego.Info(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		// return []byte(mySignKey), nil
		return mySignKeyBytes, nil
	})
	if err != nil {
		beego.Error(err)
		return "", "", "", err
	}
	beego.Info(token)
	// token.Valid里已经包含了过期判断
	if token != nil { //&& token.Valid
		beego.Info("You look nice today")
		claims, _ := token.Claims.(jwt.MapClaims)
		userId = claims["userId"].(string)
		userName = claims["userName"].(string)
		userNickname = claims["nickname"].(string)
		// companyName = claims["companyName"].(string)
	}
	return userId, userName, userNickname, err
}

// func AuthenticateUserForLogin(loginname, password string) (*models.User, error) {
// 	if len(password) == 0 || len(loginname) == 0 {
// 		return nil, errors.New("Error: 用户名或密码为空")
// 	}
// 	//密码md5加密
// 	data := []byte(password)
// 	has := md5.Sum(data)
// 	password = fmt.Sprintf("%x", has) //将[]byte转成16进制
// 	var user *models.User
// 	o := orm.NewOrm()
// 	ids := []string{loginname, password}
// 	err := o.Raw("SELECT user_id,login_name FROM user where login_name= ? and user_password= ? ", ids).QueryRow(&user)
// 	if err != nil {
// 		return nil, errors.New("Error: 未找到该用户")
// 	} else {
// 		return user, nil
// 	}
// }

// update expireAt and return a new token
func RefreshToken(tokenString string) (string, error) {
	// first get previous token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(KEY), nil
		})
	claims, ok := token.Claims.(*MyCustomClaims)
	if !ok || !token.Valid {
		return "", err
	}
	mySigningKey := []byte(KEY)
	expireAt := time.Now().Add(time.Second * time.Duration(DEFAULT_EXPIRE_SECONDS)).Unix()
	newClaims := MyCustomClaims{
		claims.User,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    claims.User.Name,
			IssuedAt:  time.Now().Unix(),
		},
	}
	// generate new token with new claims
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	tokenStr, err := newToken.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("generate new fresh json web token failed !! error :", err)
		return "", err
	}
	return tokenStr, err
}

func ValidateToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(KEY), nil
		})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		fmt.Printf("%v %v", claims.User, claims.StandardClaims.ExpiresAt)
		fmt.Println("token will be expired at ", time.Unix(claims.StandardClaims.ExpiresAt, 0))
	} else {
		fmt.Println("validate tokenString failed !!!", err)
		return err
	}
	return nil
}

//验证Token
// func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
// 	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
// 		func(token *jwt.Token) (interface{}, error) {
// 			return []byte(SecretKey), nil
// 		})

// 	if err == nil {
// 		if token.Valid {
// 			next(w, r)
// 		} else {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			fmt.Fprint(w, "Token is not valid")
// 		}
// 	} else {
// 		w.WriteHeader(http.StatusUnauthorized)
// 		fmt.Fprint(w, "Unauthorized access to this resource")
// 	}
// }

// token验证
func VerifyToken(tokenString string) bool {
	segments := strings.Split(tokenString, ".")
	res, _ := jwt.DecodeSegment(segments[1])
	var mapResult map[string]int
	//使用 json.Unmarshal(data []byte, v interface{})进行转换,返回 error 信息
	if err := json.Unmarshal([]byte(res), &mapResult); err != nil {
		return false
	}
	userId := mapResult["user_id"]
	beego.Debug(userId)
	// model, err := m.FindLastByUserId(userId)
	// if err != nil {
	// 	return false
	// }
	// secret := model.Secret
	//验证token合法性
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return false
	}
	//更新过期时间为30分钟后
	// model.ExpiredAt = time.Now().In(utils.Tz).Add(60 * 30 * 1e9)
	// beego.Debug(model.ExpiredAt)
	// m.UpdateExpired(model)
	return true
}

func GenerateToken(expiredSeconds int) (tokenString string) {
	if expiredSeconds == 0 {
		expiredSeconds = DEFAULT_EXPIRE_SECONDS
	}
	// Create the Claims
	mySigningKey := []byte(KEY)
	expireAt := time.Now().Add(time.Second * time.Duration(expiredSeconds)).Unix()
	fmt.Println("token will be expired at ", time.Unix(expireAt, 0))
	// pass parameter to this func or not
	user := User{"007", "Kev", "password"}
	claims := MyCustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    user.Name,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("generate json web token failed !! error :", err)
	}
	return tokenStr
}

// 生成jwt token
// func GenerateToken2(id int64) (string, string, error) {
// claims := make(jwt.MapClaims)
// claims["user_id"] = id
// claims["exp"] = time.Now().In(utils.Tz).Add(time.Hour * 24).Unix()
// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// secret := utils.GetRandomString(32)
// beego.Warning(secret)
// 使用自定义字符串加密 and get the complete encoded token as a string
// 	tokenString, err := token.SignedString([]byte(secret))
// 	if err != nil {
// 		beego.Error("token生成失败", err)
// 		return "", "", errors.New("token生成失败")
// 	}
// 	return tokenString, secret, nil
// }

//创建
// func Create(secret string, userId int64, ip string) (int64, error) {
// expiredAt := time.Now().Add(60 * 30 * 1e9)
// beego.Debug(expiredAt)
// model := models.UserSecret{Secret: secret, Ip: ip, UserId: userId, ExpiredAt: expiredAt}
// return m.Create(&model)
// }

// ParseToken parse JWT token in http header.
func ParseToken(authString string) (t *jwt.Token, err error) {
	// authString := base.Ctx.Input.Header("Authorization")
	beego.Debug("AuthString:", authString)

	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		beego.Error("AuthString invalid:", authString)
		return nil, err
	}
	tokenString := kv[1]

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 必要的验证 RS256
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		//// 可选项验证  'aud' claim
		//aud := "https://api.cn.atomintl.com"
		//checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
		//if !checkAud {
		//  return token, errors.New("Invalid audience.")
		//}
		// 必要的验证 'iss' claim
		iss := "https://atomintl.auth0.com/"
		checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		if !checkIss {
			return token, errors.New("Invalid issuer.")
		}
		// 我们的公钥,可以在<a href="https://manage.auth0.com/" target="_blank">https://manage.auth0.com/</a> 上下载到对应的封装好的json，里面包括了签名
		k5c := "xxxx"
		cert := "-----BEGIN CERTIFICATE-----\n" + k5c + "\n-----END CERTIFICATE-----"
		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		//result := []byte(cert) // 不是正确的 PUBKEY 格式 都会 报  key is of invalid type
		return result, nil
	})
	if err != nil {
		beego.Error("Parse token:", err)
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That's not even a token
				return nil, err
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return nil, err
			} else {
				// Couldn't handle this token
				return nil, err
			}
		} else {
			// Couldn't handle this token
			return nil, err
		}
	}
	if !token.Valid {
		beego.Error("Token invalid:", tokenString)
		return token, err
		// beego.Debug("Token:", token)
		// return token, nil
	}
	return token, nil
}

func CheckAuth(authString string) bool {
	// authString := this.Ctx.Input.Header("Authorization")
	beego.Debug("AuthString:", authString)
	kv := strings.Split(authString, " ")
	if authString == "" || len(kv) != 2 || kv[0] != "Bearer" {
		beego.Debug("no auth")
		return false
	}
	// us := new(services.UserSecretService)
	// return us.VerifyToken(kv[1])
	return true
}

// func GetMy() {
// 	var fields []string
// 	var query = make(map[string]string)
// 	//get token
// 	token, e := ParseToken()
// 	if e != nil {
// 		// c.RetError(e)
// 		beego.Debug("ParseToken error")
// 		return
// 	}
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		beego.Debug("get ParseToken claims error")
// 		//return
// 	}
// 	beego.Debug("claims:", claims)
// 	var Email string = claims["Email"].(string)
// 	beego.Debug("Email:", Email)
// 	query["Uid"] = Email

// 	// fields: col1,col2,entity.col3
// 	// if v := c.GetString("fields"); v != "" {
// 	// 	fields = strings.Split(v, ",")
// 	// }

// 	// l, err := models.GetMyPayment(query, fields)
// 	// if err != nil {
// 	// 	c.Data["json"] = err.Error()
// 	// } else {
// 	// 	c.Data["json"] = l
// 	// }
// 	// c.ServeJSON()
// }

//服务端生成token，并放入到response的header
/**
JWT由三部份组成：
* Header:头部 （对应：Header）
* Claims:声明 (对应：Payload)
* Signature:签名 (对应：Signature)
*/
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var u *User = new(User)
	var user UserCredentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in request")
		return
	}

	//验证是身份：若用户是someone，则生成token
	if strings.ToLower(user.Username) != "someone" {
		if user.Password != "p@ssword" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("Error logging in")
			fmt.Fprint(w, "Invalid credentials")
			return
		}
	}
	//1。生成token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	//2。添加令牌关键信息
	//添加令牌期限
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["id"] = u.Id
	claims["userName"] = u.Name
	claims["password"] = u.Password
	token.Claims = claims

	fmt.Println(claims)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error extracting the key")
		// fatal(err)
	}
	//获取令牌
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		// fatal(err)
	}
	//2。将生成的token放入到header
	response := Token{tokenString}
	JsonResponse(response, w)
}

// return this result to client then all later request should have header "Authorization: Bearer <token> "
func getHeaderTokenValue(tokenString string) string {
	//Authorization: Bearer <token>
	return fmt.Sprintf("Bearer %s", tokenString)
}

//如果不够length，返回全部长度范围
func SubString(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)
	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length
	if end > lth {
		end = lth
	}
	// 返回子串
	return string(rs[begin:end])
}

// 下一步我们给beego添加一个过滤器，假定：login接口返回给client 一个token，后续所有的请求都会带上这个token：
// var FilterToken = func(ctx *context.Context) {
// 	logs.Info("current router path is ", ctx.Request.RequestURI)
// 	if ctx.Request.RequestURI != "/login" && ctx.Input.Header("Authorization") == "" {
// 		logs.Error("without token, unauthorized !!")
// 		ctx.ResponseWriter.WriteHeader(401)
// 		ctx.ResponseWriter.Write([]byte("no permission"))
// 	}
// 	if ctx.Request.RequestURI != "/login" && ctx.Input.Header("Authorization") != "" {
// 		token := ctx.Input.Header("Authorization")
// 		token = strings.Split(token, "")[1]
// 		logs.Info("curernttoken: ", token)
// 		// validate token
// 		// invoke ValidateToken in utils/token
// 		// invalid or expired todo res 401
// 	}
// }

// 从request获取tokenstring
// 将tokenstring转化为未解密的token对象
// 将未解密的token对象解密得到解密后的token对象
// 从解密后的token对象里取参数
// 一、获取解密后的token
// 该函数根据request，获得tokenstring，并转为未解密token对象，解密后得到解密token对象

// import github.com/dgrijalva/jwt-go/request
// request.ParseFromRequest(req *http.Request, extractor Extractor, keyFunc jwt.Keyfunc)
// req即为http请求
// extractor 是一个实现了Extractor接口的对象，该接口需要实现的函数是ExtractToken(*http.Request) (string, error)，用于从http请求中提取tokenstring
// keyFunc是一个函数，需要接受一个“未解密的token”，并返回Secretkey的字节和错误信息
// func GetToken(r *http.Request) (token *jwt.Token, err error) { //由request获取token
// 	t := T{}
// 	// t是已经实现extract接口的对象，对request进行处理得到tokenString并生成为解密的token
// 	// request.ParseFromRequest的第三个参数是一个keyFunc，具体的直接看源代码
// 	// 该keyFunc参数需要接受一个“未解密的token”，并返回Secretkey的字节和错误信息
// 	// keyFunc被调用并传入未解密的token参数，返回解密好的token和可能出现的错误
// 	// 若解密是正确的，那么返回的token.valid = true
// 	return request.ParseFromRequest(r, t,
// 		func(token *jwt.Token) (interface{}, error) {
// 			return []byte(Secretkey), nil
// 		})
// }

// 二、(获得payload的信息)从token对象里获得参数(key)对应的值

// func GetIdFromClaims(key string, claims jwt.Claims) string {
// 	v := reflect.ValueOf(claims)
// 	if v.Kind() == reflect.Map {
// 		for _, k := range v.MapKeys() {
// 			value := v.MapIndex(k)

// 			if fmt.Sprintf("%s", k.Interface()) == key {
// 				return fmt.Sprintf("%v", value.Interface())
// 			}
// 		}
// 	}
// 	return ""
// }

// 示例 ：GetIdFromClaims("username", token.claims) 其中token是已经解密的token

// package main

// import (
//     "log"
//     "net/http"
//     "github.com/codegangsta/negroni"
//     "encoding/json"
//     "fmt"
//     "strings"
//     "github.com/dgrijalva/jwt-go"
//     "time"
//     "github.com/dgrijalva/jwt-go/request"
// )

/**
说明：
客户端通过在request对象header里添加token参数，发送到服务端。
服务端再拿出token进行比对。
token的第一次产生是发生在login检查账户存在并且正确之后，为该用户赋予一块令牌（加密字符串）
并将token放入response的header里。客户端登陆成功后，从response里取出token。并在以后操作request请求。
都保持在header里添加该段令牌，令牌有效期失效后，只有重新login，才能获取新的令牌。
*/

// func StartServer() {

// 	http.HandleFunc("/login", LoginHandler)

// 	http.Handle("/resource", negroni.New(
// 		negroni.HandlerFunc(ValidateTokenMiddleware),
// 		negroni.Wrap(http.HandlerFunc(ProtectedHandler)),
// 	))

// 	log.Println("Now listening...")
// 	http.ListenAndServe(":8080", nil)
// }

// func main() {
// 	StartServer()
// }

// func ProtectedHandler(w http.ResponseWriter, r *http.Request) {

// 	response := Response{"Gained access to protected resource !"}
// 	JsonResponse(response, w)

// }

func JsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

// import (
//     "fmt"

//     "time"

//     "github.com/dgrijalva/jwt-go"
// )

var (
	SIGN_NAME_SCERET = "aweQurt178BNI"
)

// func main() {
//     fmt.Println("Hello World!")

//     tokenString, err := createJwt()
//     if err != nil {
//         fmt.Println(err.Error())
//         return
//     }

//     fmt.Println(tokenString)

//     claims := parseJwt(tokenString)
//     fmt.Println(claims)

// }

//验证
//在调用Parse时，会进行加密验证，同时如果提供了exp，会进行过期验证；
//如果提供了iat，会进行发行时间验证;如果提供了nbf，会进行发行时间验证．

//创建 tokenString
func createJwt() (string, error) {
	//  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//      "foo": "bar",
	//      "nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),

	//  })

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["foo"] = "bar"
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(SIGN_NAME_SCERET))
	return tokenString, err
}

//解析tokenString
func parseJwt(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(SIGN_NAME_SCERET), nil
	})

	var claims jwt.MapClaims
	var ok bool

	if claims, ok = token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		fmt.Println(err)
	}

	return claims
}

// golang解析java token代码
// 在golang中解析java创建的token，需要注意token的签名加密字符串。在java中，
// 使用了方法TextCodec.BASE64.decode(base64EncodedSecretKey)。所以在token中，
// 也要进行相应的转换，将key的转化格式改成
// mySignKeyBytes, err := base64.URLEncoding.DecodeString(“abcdksdjkf123”)。

func ParseAuthToken() {
	mySignKey := "abcdksdjkf123"                                      //密钥，同java代码
	mySignKeyBytes, err := base64.URLEncoding.DecodeString(mySignKey) //需要用和加密时同样的方式转化成对应的字节数组
	if err != nil {
		fmt.Println("base64 decodeString failed.", err)
		return
	}
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDQ3NDk0NDIsImZvbyI6ImJhciIsImlhdCI6MTU0NDc1MzA0Mn0.t1NwZOUJP4Vj3L4YAiHletRNlc8vEtOMrjRAiyKl8aA"
	parseAuth, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return mySignKeyBytes, nil
	})
	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return
	}
	fmt.Println(parseAuth.Claims)
}

// 1. aud 标识token的接收者.
// 2. exp 过期时间.通常与Unix UTC时间做对比过期后token无效
// 3. jti 是自定义的id号
// 4. iat 签名发行时间.
// 5. iss 是签名的发行者.
// 6. nbf 这条token信息生效时间.这个值可以不设置,但是设定后,一定要大于当前Unix UTC,否则token将会延迟生效.
// 7. sub 签名面向的用户

// $.ajax({
//  url: ‘http://192.168.1.221:8080/v1/login/doLogin’,
//  type: ‘Post’,
//  dataType: ‘json’,
//  headers: {
//      Token:“wesdf2346dfioerjl2423mFADFOEJFds23SDFF”, //所有请求都带token 登陆就会得到token
//  },
//  data: {UserName: “retertsssssssssssssss”,UserId: ‘ertert22222222222’},
//  success: function(res){
//      console.log(res)
//  },
//  error: function(e) {
//      console.log(“errr==”)
//  }
// });
