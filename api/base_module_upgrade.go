package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseModuleUpgradeService struct {
	client *Client
}

func NewBaseModuleUpgradeService(c *Client) *BaseModuleUpgradeService {
	return &BaseModuleUpgradeService{client: c}
}

func (svc *BaseModuleUpgradeService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.BaseModuleUpgradeModel, name)
}

func (svc *BaseModuleUpgradeService) GetByIds(ids []int64) (*types.BaseModuleUpgrades, error) {
	b := &types.BaseModuleUpgrades{}
	return b, svc.client.getByIds(types.BaseModuleUpgradeModel, ids, b)
}

func (svc *BaseModuleUpgradeService) GetByName(name string) (*types.BaseModuleUpgrades, error) {
	b := &types.BaseModuleUpgrades{}
	return b, svc.client.getByName(types.BaseModuleUpgradeModel, name, b)
}

func (svc *BaseModuleUpgradeService) GetByField(field string, value string) (*types.BaseModuleUpgrades, error) {
	b := &types.BaseModuleUpgrades{}
	return b, svc.client.getByField(types.BaseModuleUpgradeModel, field, value, b)
}

func (svc *BaseModuleUpgradeService) GetAll() (*types.BaseModuleUpgrades, error) {
	b := &types.BaseModuleUpgrades{}
	return b, svc.client.getAll(types.BaseModuleUpgradeModel, b)
}

func (svc *BaseModuleUpgradeService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.BaseModuleUpgradeModel, fields, relations)
}

func (svc *BaseModuleUpgradeService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseModuleUpgradeModel, ids, fields, relations)
}

func (svc *BaseModuleUpgradeService) Delete(ids []int64) error {
	return svc.client.delete(types.BaseModuleUpgradeModel, ids)
}
