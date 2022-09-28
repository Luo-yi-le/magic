package persistence

import (
	"magic/internal/machine/domain/entity"
	"magic/internal/machine/domain/repository"
	"magic/pkg/biz"
	"magic/pkg/model"
)

type machineScriptRepoImpl struct{}

func newMachineScriptRepo() repository.MachineScript {
	return new(machineScriptRepoImpl)
}

// 分页获取机器信息列表
func (m *machineScriptRepoImpl) GetPageList(condition *entity.MachineScript, pageParam *model.PageParam, toEntity interface{}, orderBy ...string) *model.PageResult {
	return model.GetPage(pageParam, condition, toEntity, orderBy...)
}

// 根据条件获取账号信息
func (m *machineScriptRepoImpl) GetMachineScript(condition *entity.MachineScript, cols ...string) error {
	return model.GetBy(condition, cols...)
}

// 根据id获取
func (m *machineScriptRepoImpl) GetById(id uint64, cols ...string) *entity.MachineScript {
	ms := new(entity.MachineScript)
	if err := model.GetById(ms, id, cols...); err != nil {
		return nil

	}
	return ms
}

// 根据id获取
func (m *machineScriptRepoImpl) Delete(id uint64) {
	biz.ErrIsNil(model.DeleteById(new(entity.MachineScript), id), "删除失败")
}

func (m *machineScriptRepoImpl) Create(entity *entity.MachineScript) {
	model.Insert(entity)
}

func (m *machineScriptRepoImpl) UpdateById(entity *entity.MachineScript) {
	model.UpdateById(entity)
}
