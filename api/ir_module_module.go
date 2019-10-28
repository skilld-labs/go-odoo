package api

import (
	"github.com/morezig/go-odoo/types"
)

type IrModuleModuleService struct {
	client *Client
}

func NewIrModuleModuleService(c *Client) *IrModuleModuleService {
	return &IrModuleModuleService{client: c}
}

func (svc *IrModuleModuleService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrModuleModuleModel, name)
}

func (svc *IrModuleModuleService) GetByIds(ids []int64) (*types.IrModuleModules, error) {
	i := &types.IrModuleModules{}
	return i, svc.client.getByIds(types.IrModuleModuleModel, ids, i)
}

func (svc *IrModuleModuleService) GetByName(name string) (*types.IrModuleModules, error) {
	i := &types.IrModuleModules{}
	return i, svc.client.getByName(types.IrModuleModuleModel, name, i)
}

func (svc *IrModuleModuleService) GetByField(field string, value string) (*types.IrModuleModules, error) {
	i := &types.IrModuleModules{}
	return i, svc.client.getByField(types.IrModuleModuleModel, field, value, i)
}

func (svc *IrModuleModuleService) GetAll() (*types.IrModuleModules, error) {
	i := &types.IrModuleModules{}
	return i, svc.client.getAll(types.IrModuleModuleModel, i)
}

func (svc *IrModuleModuleService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrModuleModuleModel, fields, relations)
}

func (svc *IrModuleModuleService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrModuleModuleModel, ids, fields, relations)
}

func (svc *IrModuleModuleService) Delete(ids []int64) error {
	return svc.client.delete(types.IrModuleModuleModel, ids)
}
