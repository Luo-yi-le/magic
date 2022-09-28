package repository

import (
	"magic/internal/sys/domain/entity"
	"magic/pkg/model"
)

type Syslog interface {
	GetPageList(condition *entity.Syslog, pageParam *model.PageParam, toEntity interface{}, orderBy ...string) *model.PageResult

	Insert(log *entity.Syslog)
}
