package router

import (
	"magic/internal/sys/api"
	"magic/internal/sys/application"
	"magic/pkg/ctx"

	"github.com/gin-gonic/gin"
)

func InitSyslogRouter(router *gin.RouterGroup) {
	s := &api.Syslog{
		SyslogApp: application.GetSyslogApp(),
	}
	sys := router.Group("syslogs")
	{
		sys.GET("", func(c *gin.Context) {
			ctx.NewReqCtxWithGin(c).Handle(s.Syslogs)
		})
	}
}
