package api

import (
	"magic/internal/db/application"
	"magic/internal/db/domain/entity"
	"magic/pkg/ctx"
	"magic/pkg/ginx"
)

type DbSqlExec struct {
	DbSqlExecApp application.DbSqlExec
}

func (d *DbSqlExec) DbSqlExecs(rc *ctx.ReqCtx) {
	g := rc.GinCtx
	m := &entity.DbSqlExec{DbId: uint64(ginx.QueryInt(g, "dbId", 0)),
		Db:    g.Query("db"),
		Table: g.Query("table"),
		Type:  int8(ginx.QueryInt(g, "type", 0)),
	}
	m.CreatorId = rc.LoginAccount.Id
	rc.ResData = d.DbSqlExecApp.GetPageList(m, ginx.GetPageParam(rc.GinCtx), new([]entity.DbSqlExec))
}
