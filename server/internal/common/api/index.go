package api

import (
	dbapp "magic/internal/db/application"
	dbentity "magic/internal/db/domain/entity"
	machineentity "magic/internal/machine/domain/entity"
	projectentity "magic/internal/project/domain/entity"
	redisentity "magic/internal/redis/domain/entity"

	machineapp "magic/internal/machine/application"
	projectapp "magic/internal/project/application"
	redisapp "magic/internal/redis/application"
	"magic/pkg/ctx"
)

type Index struct {
	ProjectApp projectapp.Project
	MachineApp machineapp.Machine
	DbApp      dbapp.Db
	RedisApp   redisapp.Redis
}

func (i *Index) Count(rc *ctx.ReqCtx) {
	rc.ResData = map[string]interface{}{
		"projectNum": i.ProjectApp.Count(new(projectentity.Project)),
		"machineNum": i.MachineApp.Count(new(machineentity.Machine)),
		"dbNum":      i.DbApp.Count(new(dbentity.Db)),
		"redisNum":   i.RedisApp.Count(new(redisentity.Redis)),
	}
}
