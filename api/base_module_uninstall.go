package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseModuleUninstallService struct {
	client *Client
}

func NewBaseModuleUninstallService(c *Client) *BaseModuleUninstallService {
	return &BaseModuleUninstallService{client: c}
}

func (svc *BaseModuleUninstallService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BaseModuleUninstallModel, name)
}

func (svc *BaseModuleUninstallService) GetByIds(ids []int) (*types.BaseModuleUninstalls, error) {
	b := &types.BaseModuleUninstalls{}
	return b, svc.client.getByIds(types.BaseModuleUninstallModel, ids, b)
}

func (svc *BaseModuleUninstallService) GetByName(name string) (*types.BaseModuleUninstalls, error) {
	b := &types.BaseModuleUninstalls{}
	return b, svc.client.getByName(types.BaseModuleUninstallModel, name, b)
}

func (svc *BaseModuleUninstallService) GetByField(field string, value string) (*types.BaseModuleUninstalls, error) {
	b := &types.BaseModuleUninstalls{}
	return b, svc.client.getByField(types.BaseModuleUninstallModel, field, value, b)
}

func (svc *BaseModuleUninstallService) GetAll() (*types.BaseModuleUninstalls, error) {
	b := &types.BaseModuleUninstalls{}
	return b, svc.client.getAll(types.BaseModuleUninstallModel, b)
}

func (svc *BaseModuleUninstallService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BaseModuleUninstallModel, fields, relations)
}

func (svc *BaseModuleUninstallService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseModuleUninstallModel, ids, fields, relations)
}

func (svc *BaseModuleUninstallService) Delete(ids []int) error {
	return svc.client.delete(types.BaseModuleUninstallModel, ids)
}
