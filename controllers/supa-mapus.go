package controllers

import (
	"encoding/json"

	"github.com/beego/beego/v2/client/httplib"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	// "strings"
	"time"
)

// CMSWX article API
type SupaMapusController struct {
	web.Controller
}

type MapResult struct {
	Area Areastruct   `json:"area"`
	Pois []Poisstruct `json:"pois"`
	// Statistics StatisticsStruct `json:"statistics"`
	// AllAdmins AllAdmins `json:"allAdmins"`
}

type Areastruct struct {
	Name   string `json:"name"`
	Bound  string `json:"bound"`
	Lonlat string `json:"lonlat"`
	// AdminCode int    `json:"adminCode"`
	// Level     int    `json:"level"`
}

type Poisstruct struct {
	Address string `json:"address"`
	PoiType string `json:"poiType"`
	Name    string `json:"name"`
	Lonlat  string `json:"lonlat"`
}

// type StatisticsStruct struct{
// 	PriorityCitys []PriorityCitysStruct
// 	AllAdmins []ChildAdmins allAdmins

// }

// type PriorityCitysStruct struct{
//  	Name string		`json:"name"`
//   lon string    `json:"lon"`
//   lat string    `json:"lat"`
//  }

// @Title get tianditu location
// @Description get tianditu location
// @Param keyword query string true "The keyword of location"
// @Success 200 {object} models.GetAttachbyId
// @Failure 400 Invalid page supplied
// @Failure 404 location not found
// @router /tianditusearch [get]
// 取得同步文章列表
func (c *SupaMapusController) Search() {
	keyword := c.GetString("keyword")
	logs.Info(keyword)
	//	通过如下接口可以设置请求的超时时间和数据读取时间
	tiandituurl := "https://api.tianditu.gov.cn/v2/search?postStr={'keyWord':'" + keyword + "','level':12,'mapBound':'-180,-90,180,90','queryType':'1','start':'0','count':'10'}&type=query&tk=5ea85acc4bf47ec49258e859b5908686"
	jsonstring, err := httplib.Get(tiandituurl).SetTimeout(100*time.Second, 30*time.Second).String() //.ToJSON(&productlink)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "info": err, "message": "查询错误"}
		c.ServeJSON()
		return
	}
	// logs.Info(jsonstring)
	var mapresult MapResult
	//json字符串解析到结构体，以便进行追加
	err = json.Unmarshal([]byte(jsonstring), &mapresult)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "info": err, "message": "查询结果解析错误"}
		c.ServeJSON()
		return
	}
	// logs.Info(mapresult.Pois)
	// lonlattext := strings.Split(mapresult.Area.Lonlat, ",")
	if len(mapresult.Pois) > 0 {
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "info": "SUCCESS", "pois": mapresult.Pois}
		c.ServeJSON()
	} else if mapresult.Area != (Areastruct{}) {
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "info": "SUCCESS", "area": mapresult.Area}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "info": "ERROR", "message": "无pois，也无area值"}
		c.ServeJSON()
	}
}

// , "areaname": mapresult.Area.Name, "lon": lonlattext[0], "lat": lonlattext[1]

// 获取访问凭据-请求
// func RequestAccessToken(corpid string, secret string) (cache_token AccessTokenCache, ok bool) {
// 	url := "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
// 	req := httplib.Get(url)
// 	req.Param("corpid", corpid)     // 企业ID
// 	req.Param("corpsecret", secret) // 应用的凭证密钥
// 	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: false})
// 	req.AddFilters(httpFilter)
// 	resp, err := req.Response()
// 	_ = resp
// 	var token AccessTokenCache
// 	if err != nil {
// 		logs.Error(err)
// 		return token, false
// 	}
// 	var atr AccessTokenResponse
// 	err = req.ToJSON(&atr)
// 	if err != nil {
// 		logs.Error(err)
// 		return token, false
// 	}
// 	token = AccessTokenCache{
// 		AccessToken: atr.AccessToken,
// 		ExpiresIn:   atr.ExpiresIn,
// 		UpdateTime:  time.Now(),
// 	}
// 	return token, true
// }
