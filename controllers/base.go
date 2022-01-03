package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	jwt "github.com/golang-jwt/jwt"
	"strings"
)

// CMSToken API
type BaseController struct {
	web.Controller
}

// ParseToken parse JWT token in http header.
func (base *BaseController) ParseToken() (t *jwt.Token, err error) {
	authString := base.Ctx.Input.Header("Authorization")
	// web.Debug("AuthString:", authString)

	kv := strings.Split(authString, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		logs.Error("AuthString invalid:", authString)
		return nil, err
	}
	tokenString := kv[1]

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 必要的验证 RS256
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, err
			fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
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
			return token, err
			fmt.Errorf("Invalid issuer.")
		}
		// 我们的公钥,可以在<a href="https://manage.auth0.com/" target="_blank">https://manage.auth0.com/</a> 上下载到对应的封装好的json，里面包括了签名
		k5c := "xxxx"
		cert := "-----BEGIN CERTIFICATE-----\n" + k5c + "\n-----END CERTIFICATE-----"
		result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
		//result := []byte(cert) // 不是正确的 PUBKEY 格式 都会 报  key is of invalid type
		return result, nil
	})
	if err != nil {
		logs.Error("Parse token:", err)
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
		logs.Error("Token invalid:", tokenString)
		return nil, err
	}
	// beego.Debug("Token:", token)

	return token, nil
}

// func (c *BaseController) GetMy() {
// 	var fields []string
// 	var query = make(map[string]string)
// 	//get token
// 	token, e := c.ParseToken()
// 	if e != nil {
// 		beego.Debug("ParseToken error")
// 	}
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		beego.Debug("get ParseToken claims error")
// 	}
// 	beego.Debug("claims:", claims)
// 	var Email string = claims["Email"].(string)
// 	beego.Debug("Email:", Email)
// 	query["Uid"] = Email

// 	// fields: col1,col2,entity.col3
// 	if v := c.GetString("fields"); v != "" {
// 		fields = strings.Split(v, ",")
// 	}

// 	l, err := models.GetMyPayment(query, fields)
// 	if err != nil {
// 		c.Data["json"] = err.Error()
// 	} else {
// 		c.Data["json"] = l
// 	}
// 	c.ServeJSON()
// }

// if passwd == user.Password {

//             // 带权限创建令牌
//             claims := make(jwt.MapClaims)
//             claims["username"] = username
//             if username == "admin" {
//                 claims["admin"] = "true"
//             } else {
//                 claims["admin"] = "false"
//             }
//             claims["exp"] = time.Now().Add(time.Hour * 480).Unix() //20天有效期，过期需要重新登录获取token
//             token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

//             // 使用自定义字符串加密 and get the complete encoded token as a string
//             tokenString, err := token.SignedString([]byte("mykey"))
//             if err != nil {
//                 logs.Error("jwt.SignedString:", err)
//                 this.RetError(errSystem)
//                 return
//             }
//             this.Data["json"] = map[string]interface{}{"status": 200, "message": "login success ", "moreinfo": tokenString}
//         } else {
//             this.Data["json"] = map[string]interface{}{"status": 400, "message": "login failed ", "moreinfo": time.Now().Format("2006-01-02 15:04:05")}
//         }

// //**********
//         token, e := this.ParseToken()
//     if e != nil {
//         this.RetError(e)
//         return
//     }
//     claims, ok := token.Claims.(jwt.MapClaims)
//     if !ok {
//         this.RetError(errPermission)
//         return
//     }
// var user string = claims["username"].(string)
