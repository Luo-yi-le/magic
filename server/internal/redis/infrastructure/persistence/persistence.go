package persistence

import "magic/internal/redis/domain/repository"

var (
	redisRepo repository.Redis = newRedisRepo()
)

func GetRedisRepo() repository.Redis {
	return redisRepo
}
