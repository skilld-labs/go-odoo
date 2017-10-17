package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrModuleModuleExclusionService struct {
	client *Client
}

func NewIrModuleModuleExclusionService(c *Client) *IrModuleModuleExclusionService {
	return &IrModuleModuleExclusionService{client: c}
}

func (svc *IrModuleModuleExclusionService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrModuleModuleExclusionModel, name)
}

func (svc *IrModuleModuleExclusionService) GetByIds(ids []int) (*types.IrModuleModuleExclusions, error) {
	i := &types.IrModuleModuleExclusions{}
	return i, svc.client.getByIds(types.IrModuleModuleExclusionModel, ids, i)
}

func (svc *IrModuleModuleExclusionService) GetByName(name string) (*types.IrModuleModuleExclusions, error) {
	i := &types.IrModuleModuleExclusions{}
	return i, svc.client.getByName(types.IrModuleModuleExclusionModel, name, i)
}

func (svc *IrModuleModuleExclusionService) GetByField(field string, value string) (*types.IrModuleModuleExclusions, error) {
	i := &types.IrModuleModuleExclusions{}
	return i, svc.client.getByField(types.IrModuleModuleExclusionModel, field, value, i)
}

func (svc *IrModuleModuleExclusionService) GetAll() (*types.IrModuleModuleExclusions, error) {
	i := &types.IrModuleModuleExclusions{}
	return i, svc.client.getAll(types.IrModuleModuleExclusionModel, i)
}

func (svc *IrModuleModuleExclusionService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrModuleModuleExclusionModel, fields, relations)
}

func (svc *IrModuleModuleExclusionService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrModuleModuleExclusionModel, ids, fields, relations)
}

func (svc *IrModuleModuleExclusionService) Delete(ids []int) error {
	return svc.client.delete(types.IrModuleModuleExclusionModel, ids)
}
