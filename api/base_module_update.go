package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseModuleUpdateService struct {
	client *Client
}

func NewBaseModuleUpdateService(c *Client) *BaseModuleUpdateService {
	return &BaseModuleUpdateService{client: c}
}

func (svc *BaseModuleUpdateService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.BaseModuleUpdateModel, name)
}

func (svc *BaseModuleUpdateService) GetByIds(ids []int64) (*types.BaseModuleUpdates, error) {
	b := &types.BaseModuleUpdates{}
	return b, svc.client.getByIds(types.BaseModuleUpdateModel, ids, b)
}

func (svc *BaseModuleUpdateService) GetByName(name string) (*types.BaseModuleUpdates, error) {
	b := &types.BaseModuleUpdates{}
	return b, svc.client.getByName(types.BaseModuleUpdateModel, name, b)
}

func (svc *BaseModuleUpdateService) GetByField(field string, value string) (*types.BaseModuleUpdates, error) {
	b := &types.BaseModuleUpdates{}
	return b, svc.client.getByField(types.BaseModuleUpdateModel, field, value, b)
}

func (svc *BaseModuleUpdateService) GetAll() (*types.BaseModuleUpdates, error) {
	b := &types.BaseModuleUpdates{}
	return b, svc.client.getAll(types.BaseModuleUpdateModel, b)
}

func (svc *BaseModuleUpdateService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.BaseModuleUpdateModel, fields, relations)
}

func (svc *BaseModuleUpdateService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseModuleUpdateModel, ids, fields, relations)
}

func (svc *BaseModuleUpdateService) Delete(ids []int64) error {
	return svc.client.delete(types.BaseModuleUpdateModel, ids)
}
