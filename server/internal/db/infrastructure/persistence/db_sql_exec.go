package persistence

import (
	"magic/internal/db/domain/entity"
	"magic/internal/db/domain/repository"
	"magic/pkg/biz"
	"magic/pkg/model"
)

type dbSqlExecRepoImpl struct{}

func newDbSqlExecRepo() repository.DbSqlExec {
	return new(dbSqlExecRepoImpl)
}

func (d *dbSqlExecRepoImpl) Insert(dse *entity.DbSqlExec) {
	model.Insert(dse)
}

func (d *dbSqlExecRepoImpl) DeleteBy(condition *entity.DbSqlExec) {
	biz.ErrIsNil(model.DeleteByCondition(condition), "删除sql执行记录失败")
}

// 分页获取
func (d *dbSqlExecRepoImpl) GetPageList(condition *entity.DbSqlExec, pageParam *model.PageParam, toEntity interface{}, orderBy ...string) *model.PageResult {
	return model.GetPage(pageParam, condition, toEntity, orderBy...)
}
