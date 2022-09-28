package persistence

import (
	"magic/internal/sys/domain/entity"
	"magic/internal/sys/domain/repository"
	"magic/pkg/biz"
	"magic/pkg/model"
)

type msgRepoImpl struct{}

func newMsgRepo() repository.Msg {
	return new(msgRepoImpl)
}

func (m *msgRepoImpl) GetPageList(condition *entity.Msg, pageParam *model.PageParam, toEntity interface{}, orderBy ...string) *model.PageResult {
	return model.GetPage(pageParam, condition, toEntity)
}

func (m *msgRepoImpl) Insert(account *entity.Msg) {
	biz.ErrIsNil(model.Insert(account), "新增消息记录失败")
}
