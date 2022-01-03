package controllers

import (
	"github.com/astaxie/beego"
	"net/http/pprof"
)

type ProfController struct {
	beego.Controller
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
