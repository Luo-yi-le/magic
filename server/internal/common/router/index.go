package router

import (
	"magic/internal/common/api"
	dbapp "magic/internal/db/application"
	machineapp "magic/internal/machine/application"
	projectapp "magic/internal/project/application"
	redisapp "magic/internal/redis/application"
	"magic/pkg/ctx"

	"github.com/gin-gonic/gin"
)

func InitIndexRouter(router *gin.RouterGroup) {
	index := router.Group("common/index")
	i := &api.Index{
		ProjectApp: projectapp.GetProjectApp(),
		MachineApp: machineapp.GetMachineApp(),
		DbApp:      dbapp.GetDbApp(),
		RedisApp:   redisapp.GetRedisApp(),
	}
	{
		// 首页基本信息统计
		index.GET("count", func(g *gin.Context) {
			ctx.NewReqCtxWithGin(g).
				Handle(i.Count)
		})
	}
}
