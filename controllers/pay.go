package controllers

import (
	"github.com/3xxx/engineercms/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	// beego "github.com/beego/beego/v2/adapter"
	"regexp"
	"strconv"
)

type PayController struct {
	web.Controller
}

// @Title get wx pay list
// @Description get pay by page
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getwxpay [get]
// 根据pay id获得一条记录
func (c *PayController) GetWxPay() {
	var user models.User
	var err error
	openID := c.GetSession("openID")
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			logs.Error(err)
		}
	} else {
		c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		c.ServeJSON()
		return
		// user.Id = 9
	}
	pay, err := models.GetPay(user.Id)
	if err != nil {
		logs.Error(err)
	}
	c.Data["json"] = pay
	c.ServeJSON()
}

// @Title get wx userpays list
// @Description get userpay by page
// @Param page query string true "The page for mymoney list"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getwxuserpays [get]
// 根据用户id获得所有收支记录
func (c *PayController) GetWxUserPays() {
	var user models.User
	var err error
	openID := c.GetSession("openID")
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			logs.Error(err)
		}
	} else {
		c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		c.ServeJSON()
		return
		// user.Id = 9
	}

	limit := "8"
	limit1, err := strconv.Atoi(limit)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("page")
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

	pays, err := models.GetUserPay(user.Id, limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	// c.Data["json"] = pays
	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "mymoney": pays}
	c.ServeJSON()
}

// @Title get wx userappreciations list
// @Description get userappreciation by page
// @Param page query string true "The page for myappreciation list"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getwxuserpayappreciations [get]
// 根据用户id查询所有花费记录，即查询用户支出记录
func (c *PayController) GetWxUserPayAppreciations() {
	var user models.User
	var err error
	openID := c.GetSession("openID")
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			logs.Error(err)
		}
	} else {
		c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		c.ServeJSON()
		return
		// user.Id = 9
	}

	limit := "8"
	limit1, err := strconv.Atoi(limit)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("page")
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

	appreciations, err := models.GetUserPayAppreciation(user.Id, limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	// c.Data["json"] = pays
	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "mymoney": appreciations}
	c.ServeJSON()
}

// @Title get wx userappreciations list
// @Description get userappreciation by page
// @Param page query string true "The page for myappreciation list"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getwxusergetappreciations [get]
// 根据用户id获得所有赞赏记录
func (c *PayController) GetWxUserGetAppreciations() {
	var user models.User
	var err error
	openID := c.GetSession("openID")
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			logs.Error(err)
		}
	} else {
		c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		c.ServeJSON()
		return
		// user.Id = 9
	}

	limit := "8"
	limit1, err := strconv.Atoi(limit)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("page")
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

	appreciations, err := models.GetUserGetAppreciation(user.Id, limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	// c.Data["json"] = pays
	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "mymoney": appreciations}
	c.ServeJSON()
}

// @Title get wx user money
// @Description get user money
// @Success 200 {object} models.GetUserMoney
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getwxusermoney [get]
// 根据用户id获得账户余额
func (c *PayController) GetWxUserMoney() {
	var user models.User
	var err error
	openID := c.GetSession("openID")
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			logs.Error(err)
		}
	} else {
		c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		c.ServeJSON()
		return
		// user.Id = 9
	}
	var money models.Money
	money, err = models.GetUserMoney(user.Id)
	if err != nil {
		logs.Error(err)
		//如果用户账户为空，则默认充值10000金币
		err := models.AddUserRecharge(user.Id, 10000)
		if err != nil {
			logs.Error(err)
		} else {
			money.Amount = 10000
			money.UserID = user.Id
			c.Data["json"] = money
		}
	} else {
		c.Data["json"] = money
	}
	//这里要返回用户userid
	c.ServeJSON()
}

// @Title post wx userpay by articleid
// @Description post userpay by articleid
// @Param articleid query string  true "The id of article"
// @Param amount query string  true "The amout of pays"
// @Success 200 {object} models.AddUserPays
// @Failure 400 Invalid page supplied
// @Failure 404 pas not found
// @router /addwxuserpays [post]
// 用户id打赏一个文章id
func (c *PayController) AddWxUserPays() {
	var user models.User
	var err error
	openID := c.GetSession("openID")
	if openID != nil {
		user, err = models.GetUserByOpenID(openID.(string))
		if err != nil {
			logs.Error(err)
		}
	} else {
		c.Data["json"] = map[string]interface{}{"info": "用户未登录", "id": 0}
		c.ServeJSON()
		return
		// user.Id = 9
	}
	articleid := c.GetString("articleid")
	//id转成64为
	aidNum, err := strconv.ParseInt(articleid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	amount := c.GetString("amount")
	amountint, err := strconv.Atoi(amount)
	if err != nil {
		logs.Error(err)
	}
	//如果作者账户里没有记录，则充值
	// var money models.Money
	// money, err = models.GetUserMoney(product.Uid)
	// if err != nil {
	// 	logs.Error(err)
	// 	//如果用户账户为空，则默认充值10000金币
	// 	err := models.AddUserRecharge(product.Uid, 10000)
	// 	if err != nil {
	// 		logs.Error(err)
	// 	} else {
	// 		money.Amount = 10000
	// 		money.UserID = user.Id
	// 		c.Data["json"] = money
	// 	}
	// }
	err = models.AddUserPay(aidNum, user.Id, amountint)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"info": "写入数据错误", "id": 1}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"info": "SUCCESS"}
		c.ServeJSON()
	}
}

// ***********网页端***********

// @Title get pay list
// @Description get pay by page
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getpay [get]
// 根据userid获得花费和收入记录——作废
func (c *PayController) GetPay() {
	_, _, uid, _, _ := checkprodRole(c.Ctx)
	pay, err := models.GetPay(uid)
	if err != nil {
		logs.Error(err)
	}
	c.Data["json"] = pay
	c.ServeJSON()
}

// @Title get pay list
// @Description get pay by page
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getuserpay [get]
// 显示用户收支记录
func (c *PayController) GetUserPay() {
	_, _, uid, isadmin, islogin := checkprodRole(c.Ctx)
	if !islogin {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "用户未登录！", "title": "", "original": ""}
		c.ServeJSON()
		return
	}
	user := models.GetUserByUserId(uid)
	c.Data["UserNickname"] = user.Nickname
	c.Data["IsAdmin"] = isadmin
	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		c.TplName = "pay/pay.tpl"
	} else {
		c.TplName = "pay/pay.tpl"
	}
}

//后端分页的模板列表数据结构
type MathPayTable struct {
	Rows  []*models.PayMath `json:"rows"`
	Page  int64             `json:"page"`
	Total int64             `json:"total"` //string或int64都行！
}

// @Title get userpays list
// @Description get userpay by page
// @Param page query string true "The page for mymoney list"
// @Param page query string false "The page of projproducts"
// @Param limit query string false "The size of page"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getuserpaylist [get]
// 根据用户id获得所有收支记录
func (c *PayController) GetUserPaylist() {
	_, _, uid, _, islogin := checkprodRole(c.Ctx)
	if !islogin {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "用户未登录！", "title": "", "original": ""}
		c.ServeJSON()
		return
	}
	limit := "8"
	limit1, err := strconv.Atoi(limit)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("page")

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

	pays, err := models.GetUserPayMath(uid, limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	// 要做成分页的！！！
	table := MathPayTable{pays, 1, 10}
	// c.Data["json"] = pays
	c.Data["json"] = table
	c.ServeJSON()
}

// @Title get userappreciations list
// @Description get userappreciation by page
// @Param page query string true "The page for myappreciation list"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getuserpayappreciations [get]
// 根据用户id查询所有花费记录，即查询用户支出记录
func (c *PayController) GetUserPayAppreciations() {
	_, _, uid, _, _ := checkprodRole(c.Ctx)
	limit := "8"
	limit1, err := strconv.Atoi(limit)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("page")
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

	appreciations, err := models.GetUserPayMathAppreciation(uid, limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	// c.Data["json"] = pays
	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "mymoney": appreciations}
	c.ServeJSON()
}

// @Title get userappreciations list
// @Description get userappreciation by page
// @Param page query string true "The page for myappreciation list"
// @Success 200 {object} models.GetProductsPage
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getusergetappreciations [get]
// 根据用户id获得所有赞赏记录
func (c *PayController) GetUserGetAppreciations() {
	_, _, uid, _, _ := checkprodRole(c.Ctx)

	limit := "8"
	limit1, err := strconv.Atoi(limit)
	if err != nil {
		logs.Error(err)
	}
	page := c.GetString("page")
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

	appreciations, err := models.GetUserGetMathAppreciation(uid, limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	// c.Data["json"] = pays
	c.Data["json"] = map[string]interface{}{"info": "SUCCESS", "mymoney": appreciations}
	c.ServeJSON()
}

// @Title get user money
// @Description get user money
// @Success 200 {object} models.GetUserMoney
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getusermoney [get]
// 根据用户id获得账户余额
func (c *PayController) GetUserMoney() {
	_, _, uid, _, islogin := checkprodRole(c.Ctx)
	if !islogin {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "用户未登录！", "title": "", "original": ""}
		c.ServeJSON()
		return
	}
	var money models.Money
	money, err := models.GetUserMoney(uid)
	if err != nil {
		logs.Error(err)
		//如果用户账户为空，则默认充值10000金币
		err := models.AddUserRecharge(uid, 100)
		if err != nil {
			logs.Error(err)
		} else {
			money.Amount = 100
			money.UserID = uid
			c.Data["json"] = money
		}
	} else {
		c.Data["json"] = money
	}
	//这里要返回用户userid
	c.ServeJSON()
}

// @Title post userpay by mathtempleid
// @Description post userpay by mathtempid
// @Param usertempleid query string  true "The id of mathtemple"
// @Param amount query string  true "The amout of pays"
// @Success 200 {object} models.Pay
// @Failure 400 Invalid page supplied
// @Failure 404 pas not found
// @router /adduserpays [post]
// 添加用户id打赏一个mathtemple记录
func (c *PayController) AddUserPays() {
	_, _, uid, _, islogin := checkprodRole(c.Ctx)
	if !islogin {
		c.Data["json"] = map[string]interface{}{"info": "ERROR", "state": "ERROR", "data": "用户未登录！", "title": "", "original": ""}
		c.ServeJSON()
		return
	}
	var templeid uint
	usertempleid := c.GetString("usertempleid")
	//id转成uint为
	usertempleidint, err := strconv.Atoi(usertempleid)
	if err != nil {
		logs.Error(err)
	}
	templeid = uint(usertempleidint)

	amount := c.GetString("amount")
	amountint, err := strconv.Atoi(amount)
	if err != nil {
		logs.Error(err)
	}
	//如果作者账户里没有记录，则充值
	// var money models.Money
	// money, err = models.GetUserMoney(product.Uid)
	// if err != nil {
	// 	logs.Error(err)
	// 	//如果用户账户为空，则默认充值10000金币
	// 	err := models.AddUserRecharge(product.Uid, 10000)
	// 	if err != nil {
	// 		logs.Error(err)
	// 	} else {
	// 		money.Amount = 10000
	// 		money.UserID = user.Id
	// 		c.Data["json"] = money
	// 	}
	// }
	err = models.AddUserPayMath(templeid, uid, amountint)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "info": "ERROR", "data": "写入数据错误!", "id": 1}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{"info": "SUCCESS"}
		c.ServeJSON()
	}
}

// @Title get user recharge
// @Description get user recharge
// @Success 200 {object} models.Recharge
// @Failure 400 Invalid page supplied
// @Failure 404 recharge not found
// @router /applyrecharge [get]
// 用户申请充值页面
func (c *PayController) ApplyRecharge() {
	_, _, uid, _, islogin := checkprodRole(c.Ctx)
	if !islogin {
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "info": "ERROR", "data": "未登录~", "title": ""}
		//这里要返回用户userid
		c.ServeJSON()
		return
	}
	c.Data["Uid"] = uid
	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		c.TplName = "pay/applyrecharge.tpl"
	} else {
		c.TplName = "pay/applyrecharge.tpl"
	}
}

// @Title post user recharge
// @Description get user recharge
// @Param id path string  true "The id of user"
// @Param amount query string true "The amount of money"
// @Success 200 {object} models.Recharge
// @Failure 400 Invalid page supplied
// @Failure 404 recharge not found
// @router /addapplyrecharge/:id [post]
// 用户添加充值申请方法-发送数据
func (c *PayController) AddApplyRecharge() {
	userid := c.Ctx.Input.Param(":id")
	//id转成64为
	useridNum, err := strconv.ParseInt(userid, 10, 64)
	if err != nil {
		logs.Error(err)
	}
	amount := c.GetString("amount")
	amountint, err := strconv.Atoi(amount)
	if err != nil {
		logs.Error(err)
	}

	// 身份证
	// 联系方式

	err = models.AddApplyRecharge(useridNum, amountint)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "info": "ERROR", "data": "写入数据库出错！"}
	} else {
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "info": "SUCCESS", "data": "发送申请成功~"}
	}
	c.ServeJSON()
}

// 管理员查看充值申请页面

// 查看申请页面的数据

// 同意申请，修改数额

// @Title get user recharge
// @Description get user recharge
// @Success 200 {object} models.Recharge
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getapplyrecharge [get]
// 管理员查看所有用户的充值申请页面
func (c *PayController) GetApplyRecharge() {
	_, _, _, isadmin, _ := checkprodRole(c.Ctx)
	if !isadmin {
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "info": "ERROR", "data": "不是管理员~", "title": ""}
		//这里要返回用户userid
		c.ServeJSON()
		return
	}
	u := c.Ctx.Input.UserAgent()
	matched, err := regexp.MatchString("AppleWebKit.*Mobile.*", u)
	if err != nil {
		logs.Error(err)
	}
	if matched == true {
		c.TplName = "pay/recharge.tpl"
	} else {
		c.TplName = "pay/recharge.tpl"
	}
}

//后端分页的用户申请充值数据结构
type ApplyRechargeTable struct {
	Rows  []models.Recharge `json:"rows"`
	Page  int               `json:"page"`
	Total int               `json:"total"` //string或int64都行！
}

// @Title get user rechargedata
// @Description get user rechargedata
// @Param page query string false "The page of recharges"
// @Param limit query string false "The size of recharges"
// @Success 200 {object} models.Recharge
// @Failure 400 Invalid page supplied
// @Failure 404 articls not found
// @router /getapplyrechargedata [get]
// 管理员查看所有用户的充值申请数据
func (c *PayController) GetApplyRechargeData() {
	_, _, _, isadmin, _ := checkprodRole(c.Ctx)
	if !isadmin {
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "info": "ERROR", "data": "不是管理员~", "title": ""}
		//这里要返回用户userid
		c.ServeJSON()
		return
	}
	limit := c.GetString("limit")
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

	recharges, err := models.GetApplyRecharge(limit1, offset)
	if err != nil {
		logs.Error(err)
	}
	// 要做成分页的！！！
	table := ApplyRechargeTable{recharges, page1, 100}
	// c.Data["json"] = pays
	c.Data["json"] = table
	c.ServeJSON()
}

// @Title post users recharge
// @Description get users recharge
// @Param id path string  true "The id of recharge"
// @Param amount query string true "The amount of money"
// @Success 200 {object} models.Recharge
// @Failure 400 Invalid page supplied
// @Failure 404 recharge not found
// @router /adduserrecharge/:id [post]
// 管理员同意后发送充值数据-修改数值
func (c *PayController) AddUserRecharge() {
	id := c.Ctx.Input.Param(":id")
	var rechargeid uint
	if id != "" {
		//id转成uint为
		idint, err := strconv.Atoi(id)
		if err != nil {
			logs.Error(err)
		}
		rechargeid = uint(idint)
	}
	amount := c.GetString("amount")
	amountint, err := strconv.Atoi(amount)
	if err != nil {
		logs.Error(err)
	}
	// var money models.Money
	//如果用户账户为空，则默认充值100金币

	err = models.UpdateRecharge(rechargeid, amountint)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = map[string]interface{}{"state": "ERROR", "info": "ERROR", "data": "写入数据库失败！"}
	} else {
		c.Data["json"] = map[string]interface{}{"state": "SUCCESS", "info": "SUCCESS", "data": "成功~"}
	}
	c.ServeJSON()
}

// //事务的提交以及回滚
// func gun(c *gin.Context) {
// 	//创建session
// 	session := models.X.NewSession()
// 	defer session.Close()
// 	//创建事务
// 	err := session.Begin()
// 	if err != nil {
// 		c.JSON(200, gin.H{"err": err})
// 		return
// 	}
// 	//操作事务,失败并回滚(模拟购物车结算情景)
// 	car_id := c.Query("car_id")
// 	if car_id == "" {
// 		c.JSON(200, gin.H{"msg": "car_id1不得为空!", "car_id": car_id})
// 		return
// 	}

// 	//查找购物车中的商品id
// 	ids, _ := strconv.ParseInt(car_id, 10, 64)
// 	car := &models.Car{Car_id: ids}
// 	models.X.Get(car)

// 	/**
// 	 *	goods表库存减去销量
// 	 */
// 	//查询商品
// 	goods := &models.Goods{Goods_id: car.Goods_id}
// 	models.X.Get(goods)
// 	//更新库存
// 	good := models.Goods{}
// 	good.Stock = goods.Stock - car.Num
// 	rel4, err4 := session.ID(car.Goods_id).Update(good)
// 	if rel4 == 0 || err4 != nil {
// 		session.Rollback()
// 		c.JSON(200, gin.H{"err4": err4, "rel4": rel4, "carid": car.Goods_id, "goodsid": goods.Goods_id, "Stock": good.Stock})
// 		return
// 	}

// 	/**
// 	 *	用户扣费
// 	 */
// 	//查询用户
// 	user := &models.User{User_id: car.User_id}
// 	models.X.Get(user)
// 	//更新价格
// 	user_up := models.User{}
// 	user_up.Balance = user.Balance - car.Total_price
// 	rel1, err1 := session.ID(car.User_id).Update(user_up)
// 	if err1 != nil || rel1 == 0 {
// 		session.Rollback()
// 		c.JSON(200, gin.H{"err1": err1, "rel1": rel1})
// 		return
// 	}

// 	/**
// 	 *	删除用户的购物车信息
// 	 */
// 	rel2, err2 := session.Delete(car)
// 	if err2 != nil || rel2 == 0 {
// 		session.Rollback()
// 		c.JSON(200, gin.H{"err2": err2, "rel2": rel2})
// 		return
// 	}
// 	if user_up.Balance <= 0 {
// 		session.Rollback()
// 		c.JSON(200, gin.H{"msg": "余额不足"})
// 		return
// 	}

// 	err3 := session.Commit()
// 	if err3 != nil {
// 		c.JSON(200, gin.H{"err3": err3})
// 		return
// 	}
// 	c.JSON(200, gin.H{"msg": "用户扣费成功"})
// }

// func update_goods(c *gin.Context) {
// 	id := c.Query("id")
// 	if id == "" {
// 		c.JSON(200, gin.H{"msg": "id1不得为空!", "id": id})
// 		return
// 	}
// 	//string转换int64
// 	ids, err := strconv.ParseInt(id, 10, 64)

// 	goods_name := c.Query("goods_name")
// 	if goods_name == "" {
// 		c.JSON(200, gin.H{"msg": "goods_name不得为空!", "goods_name": goods_name})
// 		return
// 	}

// 	price := c.Query("price")
// 	if price == "" {
// 		c.JSON(200, gin.H{"msg": "price不得为空!", "price": price})
// 		return
// 	}
// 	prices, _ := strconv.ParseFloat(price, 64)

// 	stock := c.Query("stock")
// 	if stock == "" {
// 		c.JSON(200, gin.H{"msg": "stock不得为空!", "stock": stock})
// 		return
// 	}
// 	stocks, _ := strconv.ParseInt(stock, 10, 64)

// 	//修改
// 	goods := models.Goods{}
// 	goods.Stock = stocks
// 	goods.Goods_name = goods_name
// 	goods.Price = prices
// 	rel, err := models.X.ID(ids).Update(goods)
// 	if rel == 0 || err != nil {
// 		c.JSON(200, gin.H{"msg": "修改失败", "err": err, "stocks": stocks, "goods_name": goods_name, "prices": prices, "id": id})
// 	} else {
// 		c.JSON(200, gin.H{"msg": "修改成功"})
// 	}
// }

// func shiwu(c *gin.Context) {
// 	session := models.X.NewSession()
// 	defer session.Close()

// 	err := session.Begin()
// 	user1 := models.User{Name: "xiaoxiao1", Balance: 100}
// 	_, err = session.Insert(&user1)
// 	if err != nil {
// 		return
// 	}

// 	session.Rollback()
// 	data := make(map[string]interface{})
// 	data["msg"] = "错误"
// 	c.JSON(200, session)
// 	c.JSON(200, data)
// 	return

// 	//提交
// 	err = session.Commit()
// 	if err != nil {
// 		return
// 	}
// }
