package application

import "magic/internal/redis/infrastructure/persistence"

var (
	redisApp Redis = newRedisApp(persistence.GetRedisRepo())
)

func GetRedisApp() Redis {
	return redisApp
}
