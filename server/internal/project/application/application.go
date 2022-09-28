package application

import (
	dbapp "magic/internal/db/application"
	machineapp "magic/internal/machine/application"
	mongoapp "magic/internal/mongo/application"
	"magic/internal/project/infrastructure/persistence"
	redisapp "magic/internal/redis/application"
)

var (
	projectApp Project = newProjectApp(
		persistence.GetProjectRepo(),
		persistence.GetProjectEnvRepo(),
		persistence.GetProjectMemberRepo(),
		machineapp.GetMachineApp(),
		redisapp.GetRedisApp(),
		dbapp.GetDbApp(),
		mongoapp.GetMongoApp(),
	)
)

func GetProjectApp() Project {
	return projectApp
}
