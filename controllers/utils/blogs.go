package utils

import "github.com/beego/beego/v2/core/logs"

// var ConsoleLogs *logs.BeeLogger
var FileLogs *logs.BeeLogger

func init() {
	FileLogs = logs.NewLogger()
	FileLogs.EnableFuncCallDepth(true)
	FileLogs.SetLogger("multifile", `{"filename":"log/engineercms.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"separate":["emergency", "alert", "critical", "error", "warning", "notice", "info"]}`)
	// ConsoleLogs = logs.NewLogger(1000)
	// ConsoleLogs.SetLogger("console")
	// FileLogs = logs.NewLogger(1000)
	// FileLogs.SetLogger("file", `{"filename":”logs/test.log"}`)
}
