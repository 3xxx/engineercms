package controllers

// import (
// 	"encoding/json"

// 	"github.com/asofdate/hauth/core/groupcache"
// 	"github.com/asofdate/hauth/core/hrpc"
// 	"github.com/asofdate/hauth/core/models"
// 	"github.com/beego/beego/v2/adapter/context"
// 	"github.com/asofdate/hauth/utils/hret"
// 	"github.com/asofdate/hauth/utils/i18n"
// 	"github.com/asofdate/hauth/utils/jwt"
// 	"github.com/asofdate/hauth/utils/logs"
// )

// type userRolesController struct {
// 	models *models.UserRolesModel
// }

// var UserRolesCtl = &userRolesController{
// 	models: new(models.UserRolesModel),
// }

// // swagger:operation GET /v1/auth/batch/page StaticFiles domainShareControll
// //
// // If the request is successful,
// // will return authorization page information to the client
// //
// // The system will check user permissions.
// // So,you must first login system,and then you can send the request.
// //
// // If the user is authorized to visit, the return authorization information page
// //
// // ---
// // produces:
// // - application/json
// // - application/xml
// // - text/xml
// // - text/html
// // responses:
// //   '200':
// //     description: request success.
// //   '404':
// //     description: page not found.
// func (this *userRolesController) Page(ctx *context.Context) {
// 	if !hrpc.BasicAuth(ctx.Request) {
// 		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
// 		return
// 	}

// 	// According to the key get the value from the groupCache system
// 	rst, err := groupcache.GetStaticFile("AuthorityPage")
// 	if err != nil {
// 		hret.Error(ctx.ResponseWriter, 404, i18n.Get(ctx.Request, "as_of_date_page_not_exist"))
// 		return
// 	}

// 	ctx.ResponseWriter.Write(rst)
// }

// // swagger:operation GET /v1/auth/user/roles/get userRolesController userRolesController
// //
// // 通过user_id用户账号，来查询这个用户拥有的角色信息
// //
// // 查询角色信息
// //
// // ---
// // produces:
// // - application/json
// // - application/xml
// // - text/xml
// // - text/html
// // parameters:
// // - name: domain_id
// //   in: query
// //   description: domain code number
// //   required: true
// //   type: string
// //   format:
// // responses:
// //   '200':
// //     description: success
// func (this userRolesController) GetRolesByUserId(ctx *context.Context) {
// 	ctx.Request.ParseForm()
// 	user_id := ctx.Request.FormValue("user_id")

// 	rst, err := this.models.GetRolesByUser(user_id)
// 	if err != nil {
// 		logs.Error(err)
// 		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_user_role_query"), err)
// 		return
// 	}
// 	hret.Json(ctx.ResponseWriter, rst)
// }

// // swagger:operation GET /v1/auth/user/roles/other userRolesController userRolesController
// //
// // 通过user_id账号，查询这个用户能够访问，但是又没有获取到的角色信息
// //
// // 查询用户没有获取的角色
// //
// // ---
// // produces:
// // - application/json
// // - application/xml
// // - text/xml
// // - text/html
// // parameters:
// // - name: domain_id
// //   in: query
// //   description: domain code number
// //   required: true
// //   type: string
// //   format:
// // responses:
// //   '200':
// //     description: all domain information
// func (this userRolesController) GetOtherRoles(ctx *context.Context) {
// 	ctx.Request.ParseForm()
// 	user_id := ctx.Request.FormValue("user_id")

// 	if user_id == "" {
// 		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_user_role_no_user"))
// 		return
// 	}

// 	rst, err := this.models.GetOtherRoles(user_id)
// 	if err != nil {
// 		logs.Error(err)
// 		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_user_role_un_auth"), err)
// 		return
// 	}
// 	hret.Json(ctx.ResponseWriter, rst)
// }

// // swagger:operation POST /v1/auth/user/roles/auth userRolesController userRolesController
// //
// // 给指定的用户授予角色
// //
// // 给指定的用户授予角色
// //
// // ---
// // produces:
// // - application/json
// // - application/xml
// // - text/xml
// // - text/html
// // parameters:
// // - name: domain_id
// //   in: query
// //   description: domain code number
// //   required: true
// //   type: string
// //   format:
// // responses:
// //   '200':
// //     description: success
// func (this userRolesController) Auth(ctx *context.Context) {
// 	ctx.Request.ParseForm()
// 	if !hrpc.BasicAuth(ctx.Request) {
// 		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
// 		return
// 	}
// 	var rst []models.UserRolesModel
// 	err := json.Unmarshal([]byte(ctx.Request.FormValue("JSON")), &rst)
// 	if err != nil {
// 		logs.Error(err)
// 		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_unmarsh_json"), err)
// 		return
// 	}

// 	for _, val := range rst {
// 		domain_id, err := hrpc.GetDomainId(val.User_id)
// 		if err != nil {
// 			logs.Error(err)
// 			hret.Error(ctx.ResponseWriter, 403, i18n.Get(ctx.Request, "error_user_role_no_auth"))
// 			return
// 		}

// 		if !hrpc.DomainAuth(ctx.Request, domain_id, "w") {
// 			hret.Error(ctx.ResponseWriter, 403, i18n.WriteDomain(ctx.Request, domain_id))
// 			return
// 		}
// 	}

// 	cok, _ := ctx.Request.Cookie("Authorization")
// 	jclaim, err := jwt.ParseJwt(cok.Value)
// 	if err != nil {
// 		logs.Error(err)
// 		hret.Error(ctx.ResponseWriter, 403, i18n.Disconnect(ctx.Request))
// 		return
// 	}

// 	msg, err := this.models.Auth(rst, jclaim.UserId)
// 	if err != nil {
// 		logs.Error(err)
// 		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, msg), err)
// 		return
// 	}
// 	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
// }

// // swagger:operation POST /v1/auth/user/roles/revoke userRolesController userRolesController
// //
// // Delete user has been granted the roles
// //
// // The system will check user permissions.
// // So,you must first login system,and then you can send the request.
// //
// // If the user is authorized to visit, the system will delete the roles that client request specified.
// //
// // ---
// // produces:
// // - application/json
// // - application/xml
// // - text/xml
// // - text/html
// // parameters:
// // - name: user_id
// //   in: query
// //   description: Removed the role of the user
// //   required: true
// //   type: string
// //   format:
// // - name: role_id
// //   in: query
// //   description: The role of ready to delete
// //   required: true
// //   type: string
// //   format:
// // responses:
// //   '200':
// //     description: success
// func (this userRolesController) Revoke(ctx *context.Context) {
// 	ctx.Request.ParseForm()
// 	if !hrpc.BasicAuth(ctx.Request) {
// 		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
// 		return
// 	}

// 	form := ctx.Request.FormValue("JSON")
// 	var rst []models.UserRolesModel

// 	err := json.Unmarshal([]byte(form), &rst)
// 	if err != nil {
// 		logs.Error("解析json格式数据失败，请联系管理员")
// 		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_unmarsh_json"))
// 		return
// 	}

// 	for _, val := range rst {
// 		domain_id, err := hrpc.GetDomainId(val.User_id)
// 		if err != nil {
// 			logs.Error(err)
// 			hret.Error(ctx.ResponseWriter, 403, i18n.Get(ctx.Request, "error_user_role_no_auth"))
// 			return
// 		}

// 		if !hrpc.DomainAuth(ctx.Request, domain_id, "w") {
// 			hret.Error(ctx.ResponseWriter, 403, i18n.WriteDomain(ctx.Request, domain_id))
// 			return
// 		}
// 	}

// 	msg, err := this.models.Revoke(rst)
// 	if err != nil {
// 		logs.Error(err)
// 		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, msg), err)
// 		return
// 	}
// 	hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
// }

// func init() {
// 	// Registered in the static page to the groupCache system
// 	groupcache.RegisterStaticFile("AuthorityPage", "./views/hauth/sys_batch_page.tpl")
// }
