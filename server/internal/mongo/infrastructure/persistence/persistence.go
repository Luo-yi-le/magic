package persistence

import (
	"magic/internal/mongo/domain/repository"
)

var (
	mongoRepo repository.Mongo = newMongoRepo()
)

func GetMongoRepo() repository.Mongo {
	return mongoRepo
}
