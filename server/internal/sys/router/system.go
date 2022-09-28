package router

import (
	"magic/internal/sys/api"

	"github.com/gin-gonic/gin"
)

func InitSystemRouter(router *gin.RouterGroup) {
	s := &api.System{}
	sys := router.Group("sysmsg")

	{
		sys.GET("", s.ConnectWs)
	}
}
