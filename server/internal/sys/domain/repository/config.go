package repository

import (
	"magic/internal/sys/domain/entity"
	"magic/pkg/model"
)

type Config interface {
	GetPageList(condition *entity.Config, pageParam *model.PageParam, toEntity interface{}, orderBy ...string) *model.PageResult

	Insert(config *entity.Config)

	Update(config *entity.Config)

	GetConfig(config *entity.Config, cols ...string) error

	GetByCondition(condition *entity.Config, cols ...string) error
}
