package controllers

// import (
// 	"html/template"

// 	"github.com/beego/beego/v2/adapter/context"
// 	"github.com/asofdate/hauth/core/hrpc"
// 	"github.com/asofdate/hauth/core/models"
// 	"github.com/asofdate/hauth/utils/hret"
// 	"github.com/asofdate/hauth/utils/i18n"
// 	"github.com/asofdate/hauth/utils/logs"
// 	"github.com/asofdate/hauth/utils/validator"
// )

// type roleAndResourceController struct {
// 	model        *models.RoleModel
// 	resRoleModel *models.RoleAndResourceModel
// 	resModel     *models.ResourceModel
// }

// var RoleAndResourceCtl = &roleAndResourceController{
// 	new(models.RoleModel),
// 	new(models.RoleAndResourceModel),
// 	new(models.ResourceModel),
// }

// // swagger:operation GET /v1/auth/role/resource/details StaticFiles domainShareControll
// //
// // 角色菜单资源配置管理页面
// //
// // 如果用户被授权,将会返回指定角色资源管理页面.
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
// func (this roleAndResourceController) ResourcePage(ctx *context.Context) {
// 	ctx.Request.ParseForm()
// 	if !hrpc.BasicAuth(ctx.Request) {
// 		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
// 		return
// 	}

// 	var role_id = ctx.Request.FormValue("role_id")

// 	rst, err := this.model.GetRow(role_id)

// 	if err != nil {
// 		logs.Error(err)
// 		hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_role_resource_query"))
// 		return
// 	}
// 	file, _ := template.ParseFiles("./views/hauth/res_role_rel_page.tpl")

// 	file.Execute(ctx.ResponseWriter, rst)
// }

// // swagger:operation GET /v1/auth/role/resource/get roleAndResourceController roleAndResourceController
// //
// // 查询角色指定的拥有的菜单资源和没有拥有的菜单资源
// //
// // type_id = 0 表示查询角色拥有的菜单资源, type_id = 1 表示查询角色没有获取的菜单资源
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
// func (this roleAndResourceController) GetResource(ctx *context.Context) {
// 	ctx.Request.ParseForm()
// 	if !hrpc.BasicAuth(ctx.Request) {
// 		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
// 		return
// 	}

// 	role_id := ctx.Request.FormValue("role_id")
// 	type_id := ctx.Request.FormValue("type_id")

// 	if type_id == "0" {
// 		// 查询角色已经获取到的资源信息
// 		rst, err := this.resRoleModel.Get(role_id)
// 		if err != nil {
// 			logs.Error(err)
// 			hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_role_get_resource"))
// 			return
// 		}
// 		hret.Json(ctx.ResponseWriter, rst)
// 	} else if type_id == "1" {
// 		// 查询角色没有获取到的资源信息
// 		rst, err := this.resRoleModel.UnGetted(role_id)
// 		if err != nil {
// 			logs.Error(err)
// 			hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_role_unget_resource"))
// 			return
// 		}
// 		hret.Json(ctx.ResponseWriter, rst)
// 	}
// }

// // swagger:operation POST /v1/auth/role/resource/rights roleAndResourceController roleAndResourceController
// //
// // 授予角色菜单资源或删除角色菜单资源
// //
// // type_id = 0 表示移除某个指定角色的菜单资源.
// //
// // type_id = 1 表示给某个指定角色增加菜单资源.
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
// func (this roleAndResourceController) HandleResource(ctx *context.Context) {
// 	ctx.Request.ParseForm()
// 	if !hrpc.BasicAuth(ctx.Request) {
// 		hret.Error(ctx.ResponseWriter, 403, i18n.NoAuth(ctx.Request))
// 		return
// 	}

// 	res_id := ctx.Request.FormValue("res_id")
// 	role_id := ctx.Request.FormValue("role_id")
// 	type_id := ctx.Request.FormValue("type_id")

// 	if !validator.IsWord(res_id) {
// 		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_resource_res_id"))
// 		return
// 	}

// 	if !validator.IsWord(role_id) {
// 		hret.Error(ctx.ResponseWriter, 421, i18n.Get(ctx.Request, "error_role_id_format"))
// 		return
// 	}

// 	// 撤销权限操作
// 	if type_id == "0" {
// 		err := this.resRoleModel.Delete(role_id, res_id)
// 		if err != nil {
// 			logs.Error(err)
// 			hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_role_delete_failed"))
// 			return
// 		} else {
// 			hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
// 			return
// 		}
// 	} else {
// 		//授权操作
// 		err := this.resRoleModel.Post(role_id, res_id)
// 		if err != nil {
// 			logs.Error(err)
// 			hret.Error(ctx.ResponseWriter, 419, i18n.Get(ctx.Request, "error_role_delete_failed"))
// 			return
// 		} else {
// 			hret.Success(ctx.ResponseWriter, i18n.Success(ctx.Request))
// 			return
// 		}
// 	}
// }
