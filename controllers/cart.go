package controllers

import (
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"strconv"
	"strings"
)

type CartResult struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type CartController struct {
	web.Controller
}

// @Title post create a new cart
// @Description post create a new cart
// @Param ids query string true "The ids of product"
// @Success 200 {object} models.AddArticle
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /createproductcart [post]
func (c *CartController) CreateProductCart() {
	ids := c.GetString("ids")
	if ids == "" {
		logs.Error("matterUuids cannot be null")
		c.Data["json"] = map[string]interface{}{"code": "ERROR", "msg": "matterUuids cannot be null"}
		c.ServeJSON()
		return
	}

	Array := strings.Split(ids, ",")

	if len(Array) == 0 {
		logs.Error("share at least one file")
		c.Data["json"] = map[string]interface{}{"code": "ERROR", "msg": "share at least one file"}
		c.ServeJSON()
		return
	} else if len(Array) > SHARE_MAX_NUM {
		logs.Error("ShareNumExceedLimit")
		c.Data["json"] = map[string]interface{}{"code": "ERROR", "msg": "ShareNumExceedLimit"}
		c.ServeJSON()
		return
	}

	v := c.GetSession("uname")
	var user models.User
	var err error
	if v != nil { //如果登录了
		user, err = models.GetUserByUsername(v.(string))
		if err != nil {
			logs.Error(err)
		}
	} else {
		c.Data["json"] = map[string]interface{}{"code": "ERROR", "msg": "用户未登录"}
		c.ServeJSON()
		return
	}
	// beego.Info(user)
	// var cart models.Cart
	products := make([]models.Product, 0)
	productslice := make([]models.Product, 1)
	for _, id := range Array {
		idNum, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			logs.Error(err)
		}
		product, err := models.GetProd(idNum)
		// beego.Info(product.Id)
		productslice[0] = product
		products = append(products, productslice...)
		// cart.ProductId = product.Id
		// cart.UserId = user.Id
		_, err = models.CreateCart(product.Id, user.Id)
		if err != nil {
			logs.Error(err)
		}
	}
	// share := &models.Share{
	// 	Name:     name,
	// 	UserId:   user.Id,
	// 	Username: user.Username,
	// }
	c.Data["json"] = map[string]interface{}{"code": "OK", "msg": "", "data": products}
	c.ServeJSON()
}

// @Title get usercarttpl
// @Description get usercarttpl
// @Success 200 {object} models.Create
// @Failure 400 Invalid page supplied
// @Failure 404 cart not found
// @router /getcart [get]
func (c *CartController) GetCart() {
	c.TplName = "cart/usercart.tpl"
	v := c.GetSession("uname")
	var user models.User
	var userid, roleid string
	var err error
	var islogin bool
	if v != nil { //如果登录了
		user, err = models.GetUserByUsername(v.(string))
		if err != nil {
			logs.Error(err)
		}
		//查询admin角色的id
		//重新获取roleid
		role, err := models.GetRoleByRolename("admin")
		if err != nil {
			logs.Error(err)
		}
		userid = strconv.FormatInt(user.Id, 10)
		roleid = strconv.FormatInt(role.Id, 10)
		islogin = true
	} else {
		c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		c.ServeJSON()
		islogin = false
		return
	}
	isadmin, err := e.HasRoleForUser(userid, "role_"+roleid)
	if err != nil {
		logs.Error(err)
	}
	c.Data["IsAdmin"] = isadmin
	c.Data["IsLogin"] = islogin
	c.Data["Username"] = user.Nickname
	site := c.Ctx.Input.Site()
	port := strconv.Itoa(c.Ctx.Input.Port())
	if port == "80" {
		c.Data["Site"] = site
	} else {
		c.Data["Site"] = site + ":" + port
	}
}

// @Title get usercartlist
// @Description get usercartlist
// @Param status query string true "The status for usercart list"
// @Param searchText query string false "The searchText of usercart"
// @Param pageNo query string true "The page for usercart list"
// @Param limit query string true "The limit of page for usercart list"
// @Success 200 {object} models.Create
// @Failure 400 Invalid page supplied
// @Failure 404 cart not found
// @router /getapprovalcart [get]
//根据用户id获得借阅记录，如果是admin角色，则查询全部
func (c *CartController) GetApprovalCart() {
	v := c.GetSession("uname")
	var user models.User
	var userid, roleid string
	var err error
	if v != nil { //如果登录了
		user, err = models.GetUserByUsername(v.(string))
		if err != nil {
			logs.Error(err)
		}
		//查询admin角色的id
		//重新获取roleid
		role, err := models.GetRoleByRolename("admin")
		if err != nil {
			logs.Error(err)
		}
		userid = strconv.FormatInt(user.Id, 10)
		roleid = strconv.FormatInt(role.Id, 10)

	} else {
		c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		c.ServeJSON()
		return
	}
	status := c.GetString("status")
	status1, err := strconv.Atoi(status)
	if err != nil {
		logs.Error(err)
	}
	searchText := c.GetString("searchText")
	limit := c.GetString("limit")
	if limit == "" {
		limit = "15"
	}
	limit1, err := strconv.Atoi(limit)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("pageNo")
	page1, err := strconv.Atoi(page)
	if err != nil {
		logs.Error(err)
	}
	var offset int
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}
	isadmin, err := e.HasRoleForUser(userid, "role_"+roleid)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(isadmin)
	carts, err := models.GetApprovalCart(user.Id, limit1, offset, status1, searchText, isadmin)
	if err != nil {
		logs.Error(err)
	}
	count, err := models.GetApprovalCartCount(user.Id, status1, searchText, isadmin)
	if err != nil {
		logs.Error(err)
	}
	// c.Data["json"] = carts
	c.Data["json"] = map[string]interface{}{"page": page1, "total": count, "rows": carts}

	// c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "mycarts": carts}
	c.ServeJSON()
}

// @Title get usercartlist
// @Description get usercartlist
// @Param status query string true "The status for usercart list"
// @Param searchText query string false "The searchText of usercart"
// @Param pageNo query string true "The page for usercart list"
// @Param limit query string true "The limit of page for usercart list"
// @Success 200 {object} models.Create
// @Failure 400 Invalid page supplied
// @Failure 404 cart not found
// @router /getapplycart [get]
//根据用户id获得借阅记录，如果是admin角色，则查询全部
func (c *CartController) GetApplyCart() {
	v := c.GetSession("uname")
	var user models.User
	var userid, roleid string
	var err error
	if v != nil { //如果登录了
		user, err = models.GetUserByUsername(v.(string))
		if err != nil {
			logs.Error(err)
		}
		//查询admin角色的id
		//重新获取roleid
		role, err := models.GetRoleByRolename("admin")
		if err != nil {
			logs.Error(err)
		}
		userid = strconv.FormatInt(user.Id, 10)
		roleid = strconv.FormatInt(role.Id, 10)

	} else {
		c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		c.ServeJSON()
		return
	}
	status := c.GetString("status")
	status1, err := strconv.Atoi(status)
	if err != nil {
		logs.Error(err)
	}
	searchText := c.GetString("searchText")
	limit := c.GetString("limit")
	if limit == "" {
		limit = "15"
	}
	limit1, err := strconv.Atoi(limit)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("pageNo")
	page1, err := strconv.Atoi(page)
	if err != nil {
		logs.Error(err)
	}
	var offset int
	if page1 <= 1 {
		offset = 0
	} else {
		offset = (page1 - 1) * limit1
	}
	isadmin, err := e.HasRoleForUser(userid, "role_"+roleid)
	if err != nil {
		logs.Error(err)
	}
	// beego.Info(isadmin)
	carts, err := models.GetApplyCart(user.Id, limit1, offset, status1, searchText, isadmin)
	if err != nil {
		logs.Error(err)
	}
	count, err := models.GetApplyCartCount(user.Id, status1, searchText, isadmin)
	if err != nil {
		logs.Error(err)
	}
	// c.Data["json"] = carts
	c.Data["json"] = map[string]interface{}{"page": page1, "total": count, "rows": carts}

	// c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "mycarts": carts}
	c.ServeJSON()
}

// @Title get usercartlist
// @Description get usercartlist
// @Param status query string true "The status for usercart list"
// @Param searchText query string false "The searchText of usercart"
// @Param pageNo query string true "The page for usercart list"
// @Param limit query string true "The limit of page for usercart list"
// @Success 200 {object} models.Create
// @Failure 400 Invalid page supplied
// @Failure 404 cart not found
// @router /gethistorycart [get]
//根据用户id获得借阅记录，如果是admin角色，则查询全部
// func (c *CartController) GetHistoryCart() {
// 	v := c.GetSession("uname")
// 	var user models.User
// 	var userid, roleid string
// 	var err error
// 	if v != nil {
// 		user, err = models.GetUserByUsername(v.(string))
// 		if err != nil {
// 			logs.Error(err)
// 		}
// 		role, err := models.GetRoleByRolename("admin")
// 		if err != nil {
// 			logs.Error(err)
// 		}
// 		userid = strconv.FormatInt(user.Id, 10)
// 		roleid = strconv.FormatInt(role.Id, 10)

// 	} else {
// 		c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
// 		c.ServeJSON()
// 		return
// 	}
// 	status := c.GetString("status")
// 	status1, err := strconv.Atoi(status)
// 	if err != nil {
// 		logs.Error(err)
// 	}
// 	searchText := c.GetString("searchText")
// 	limit := c.GetString("limit")
// 	if limit == "" {
// 		limit = "15"
// 	}
// 	limit1, err := strconv.Atoi(limit)
// 	if err != nil {
// 		logs.Error(err)
// 	}
// 	page := c.GetString("pageNo")
// 	page1, err := strconv.Atoi(page)
// 	if err != nil {
// 		logs.Error(err)
// 	}
// 	var offset int
// 	if page1 <= 1 {
// 		offset = 0
// 	} else {
// 		offset = (page1 - 1) * limit1
// 	}
// 	isadmin := e.HasRoleForUser(userid, "role_"+roleid)
// 	beego.Info(isadmin)
// 	carts, err := models.GetHistoryCart(user.Id, limit1, offset, status1, searchText, isadmin)
// 	if err != nil {
// 		logs.Error(err)
// 	}
// 	count, err := models.GetHistoryCartCount(user.Id, status1, searchText, isadmin)
// 	if err != nil {
// 		logs.Error(err)
// 	}
// 	c.Data["json"] = map[string]interface{}{"page": page1, "total": count, "rows": carts}
// 	c.ServeJSON()
// }

// @Title post delete usercart
// @Description post delete a usercart
// @Param ids query string true "The ids of usercats"
// @Success 200 {object} models.DeleteUserCart
// @Failure 400 Invalid page supplied
// @Failure 404 articl not found
// @router /deleteusercart [post]
func (c *CartController) DeleteUserCart() {
	ids := c.GetString("ids")
	if ids == "" {
		logs.Error("matterUuids cannot be null")
	}

	Array := strings.Split(ids, ",")
	// var ob []models.PostMerit传结构体json格式
	// json.Unmarshal(c.Ctx.Input.RequestBody, &ob)

	if len(Array) == 0 {
		logs.Error("share at least one file")
	} else if len(Array) > SHARE_MAX_NUM {
		logs.Error("ShareNumExceedLimit")
	}

	v := c.GetSession("uname")
	var user models.User
	var userid, roleid string
	var err error
	if v != nil { //如果登录了
		user, err = models.GetUserByUsername(v.(string))
		if err != nil {
			logs.Error(err)
		}
		//查询admin角色的id
		//重新获取roleid
		role, err := models.GetRoleByRolename("admin")
		if err != nil {
			logs.Error(err)
		}
		userid = strconv.FormatInt(user.Id, 10)
		roleid = strconv.FormatInt(role.Id, 10)
	} else {
		c.Data["json"] = map[string]interface{}{"code": "ERROR", "info": "用户未登录", "msg": "未登陆"}
		return
	}
	isadmin, err := e.HasRoleForUser(userid, "role_"+roleid)
	if err != nil {
		logs.Error(err)
	}
	for _, id := range Array {
		idNum, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			logs.Error(err)
		}
		usercart, err := models.GetUserCartbyId(idNum)
		if err != nil {
			logs.Error(err)
		}
		if isadmin || usercart.UserId == user.Id {
			err = models.DeleteUserCart(idNum)
			if err != nil {
				logs.Error(err)
			}
		}
	}
	c.Data["json"] = map[string]interface{}{"code": "OK", "info": "OK", "msg": "OK", "data": ids}
	c.ServeJSON()
}

// @Title post update carts
// @Description post update carts
// @Param ids query string true "The ids of approvalcats"
// @Success 200 {object} models.Update
// @Failure 400 Invalid page supplied
// @Failure 404 cart not found
// @router /updateapprovalcart [post]
func (c *CartController) UpdateApprovalCart() {
	ids := c.GetString("ids")
	if ids == "" {
		logs.Error("matterUuids cannot be null")
	}

	Array := strings.Split(ids, ",")
	// var ob []models.PostMerit传结构体json格式
	// json.Unmarshal(c.Ctx.Input.RequestBody, &ob)

	if len(Array) == 0 {
		logs.Error("share at least one file")
	}
	for _, id := range Array {
		idNum, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			logs.Error(err)
		}
		err = models.UpdateApprovalCart(idNum)
		if err != nil {
			logs.Error(err)
		}
	}
	c.Data["json"] = map[string]interface{}{"code": "OK", "info": "OK", "msg": "OK", "data": ids}
	c.ServeJSON()
}
