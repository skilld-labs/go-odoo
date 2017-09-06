package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrModuleModuleDependencyService struct {
	client *Client
}

func NewIrModuleModuleDependencyService(c *Client) *IrModuleModuleDependencyService {
	return &IrModuleModuleDependencyService{client: c}
}

func (svc *IrModuleModuleDependencyService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrModuleModuleDependencyModel, name)
}

func (svc *IrModuleModuleDependencyService) GetByIds(ids []int) (*types.IrModuleModuleDependencys, error) {
	i := &types.IrModuleModuleDependencys{}
	return i, svc.client.getByIds(types.IrModuleModuleDependencyModel, ids, i)
}

func (svc *IrModuleModuleDependencyService) GetByName(name string) (*types.IrModuleModuleDependencys, error) {
	i := &types.IrModuleModuleDependencys{}
	return i, svc.client.getByName(types.IrModuleModuleDependencyModel, name, i)
}

func (svc *IrModuleModuleDependencyService) GetByField(field string, value string) (*types.IrModuleModuleDependencys, error) {
	i := &types.IrModuleModuleDependencys{}
	return i, svc.client.getByField(types.IrModuleModuleDependencyModel, field, value, i)
}

func (svc *IrModuleModuleDependencyService) GetAll() (*types.IrModuleModuleDependencys, error) {
	i := &types.IrModuleModuleDependencys{}
	return i, svc.client.getAll(types.IrModuleModuleDependencyModel, i)
}

func (svc *IrModuleModuleDependencyService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrModuleModuleDependencyModel, fields, relations)
}

func (svc *IrModuleModuleDependencyService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrModuleModuleDependencyModel, ids, fields, relations)
}

func (svc *IrModuleModuleDependencyService) Delete(ids []int) error {
	return svc.client.delete(types.IrModuleModuleDependencyModel, ids)
}
