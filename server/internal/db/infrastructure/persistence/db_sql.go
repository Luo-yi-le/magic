package persistence

import (
	"magic/internal/db/domain/entity"
	"magic/internal/db/domain/repository"
	"magic/pkg/biz"
	"magic/pkg/model"
)

type dbSqlRepoImpl struct{}

func newDbSqlRepo() repository.DbSql {
	return new(dbSqlRepoImpl)
}

// 分页获取数据库信息列表
func (d *dbSqlRepoImpl) DeleteBy(condition *entity.DbSql) {
	biz.ErrIsNil(model.DeleteByCondition(condition), "删除sql失败")
}
