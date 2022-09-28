package persistence

import (
	"magic/internal/project/domain/entity"
	"magic/internal/project/domain/repository"
	"magic/pkg/biz"
	"magic/pkg/model"
)

type projectMemberRepoImpl struct{}

func newProjectMemberRepo() repository.ProjectMemeber {
	return new(projectMemberRepoImpl)
}

func (p *projectMemberRepoImpl) ListMemeber(condition *entity.ProjectMember, toEntity interface{}, orderBy ...string) {
	model.ListByOrder(condition, toEntity, orderBy...)
}

func (p *projectMemberRepoImpl) Save(pm *entity.ProjectMember) {
	biz.ErrIsNilAppendErr(model.Insert(pm), "保存项目成员失败：%s")
}

func (p *projectMemberRepoImpl) GetPageList(condition *entity.ProjectMember, pageParam *model.PageParam, toEntity interface{}, orderBy ...string) *model.PageResult {
	return model.GetPage(pageParam, condition, toEntity, orderBy...)
}

func (p *projectMemberRepoImpl) DeleteByPidMid(projectId, accountId uint64) {
	model.DeleteByCondition(&entity.ProjectMember{ProjectId: projectId, AccountId: accountId})
}

func (p *projectMemberRepoImpl) DeleteMems(projectId uint64) {
	model.DeleteByCondition(&entity.ProjectMember{ProjectId: projectId})
}

func (p *projectMemberRepoImpl) IsExist(projectId, accountId uint64) bool {
	return model.CountBy(&entity.ProjectMember{ProjectId: projectId, AccountId: accountId}) > 0
}
