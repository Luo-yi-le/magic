package api

import (
	"magic/internal/sys/application"
	"magic/internal/sys/domain/entity"
	"magic/pkg/ctx"
	"magic/pkg/ginx"
)

type Syslog struct {
	SyslogApp application.Syslog
}

func (r *Syslog) Syslogs(rc *ctx.ReqCtx) {
	g := rc.GinCtx
	condition := &entity.Syslog{
		Type:      int8(ginx.QueryInt(g, "type", 0)),
		CreatorId: uint64(ginx.QueryInt(g, "creatorId", 0)),
	}
	rc.ResData = r.SyslogApp.GetPageList(condition, ginx.GetPageParam(g), new([]entity.Syslog), "create_time DESC")
}
