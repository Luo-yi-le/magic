package repository

import (
	"magic/internal/sys/domain/entity"
	"magic/pkg/model"
)

type Account interface {
	// 根据条件获取账号信息
	GetAccount(condition *entity.Account, cols ...string) error

	GetPageList(condition *entity.Account, pageParam *model.PageParam, toEntity interface{}, orderBy ...string) *model.PageResult

	Insert(account *entity.Account)

	Update(account *entity.Account)
}
