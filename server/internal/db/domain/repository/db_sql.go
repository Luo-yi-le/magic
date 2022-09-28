package repository

import "magic/internal/db/domain/entity"

type DbSql interface {
	DeleteBy(condition *entity.DbSql)
}
