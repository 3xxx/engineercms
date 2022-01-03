package controllers

import (
	// beego "github.com/beego/beego/v2/adapter"
	// "github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"net/http/pprof"
)

type ProfController struct {
	web.Controller
}

func (this *ProfController) Get() {
	switch this.Ctx.Input.Param(":app") {
	default:
		pprof.Index(this.Ctx.ResponseWriter, this.Ctx.Request)
	case "":
		pprof.Index(this.Ctx.ResponseWriter, this.Ctx.Request)
	case "cmdline":
		pprof.Cmdline(this.Ctx.ResponseWriter, this.Ctx.Request)
	case "profile":
		pprof.Profile(this.Ctx.ResponseWriter, this.Ctx.Request)
	case "symbol":
		pprof.Symbol(this.Ctx.ResponseWriter, this.Ctx.Request)
	}
	this.Ctx.ResponseWriter.WriteHeader(200)
}
