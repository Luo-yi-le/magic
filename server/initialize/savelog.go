package initialize

import (
	sysapp "magic/internal/sys/application"
	"magic/pkg/ctx"
)

func InitSaveLogFunc() ctx.SaveLogFunc {
	return sysapp.GetSyslogApp().SaveFromReq
}
