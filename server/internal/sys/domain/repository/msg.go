package repository

import (
	"magic/internal/sys/domain/entity"
	"magic/pkg/model"
)

type Msg interface {
	GetPageList(condition *entity.Msg, pageParam *model.PageParam, toEntity interface{}, orderBy ...string) *model.PageResult

	Insert(msg *entity.Msg)
}
