package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseModuleConfigurationService struct {
	client *Client
}

func NewBaseModuleConfigurationService(c *Client) *BaseModuleConfigurationService {
	return &BaseModuleConfigurationService{client: c}
}

func (svc *BaseModuleConfigurationService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.BaseModuleConfigurationModel, name)
}

func (svc *BaseModuleConfigurationService) GetByIds(ids []int64) (*types.BaseModuleConfigurations, error) {
	b := &types.BaseModuleConfigurations{}
	return b, svc.client.getByIds(types.BaseModuleConfigurationModel, ids, b)
}

func (svc *BaseModuleConfigurationService) GetByName(name string) (*types.BaseModuleConfigurations, error) {
	b := &types.BaseModuleConfigurations{}
	return b, svc.client.getByName(types.BaseModuleConfigurationModel, name, b)
}

func (svc *BaseModuleConfigurationService) GetByField(field string, value string) (*types.BaseModuleConfigurations, error) {
	b := &types.BaseModuleConfigurations{}
	return b, svc.client.getByField(types.BaseModuleConfigurationModel, field, value, b)
}

func (svc *BaseModuleConfigurationService) GetAll() (*types.BaseModuleConfigurations, error) {
	b := &types.BaseModuleConfigurations{}
	return b, svc.client.getAll(types.BaseModuleConfigurationModel, b)
}

func (svc *BaseModuleConfigurationService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.BaseModuleConfigurationModel, fields, relations)
}

func (svc *BaseModuleConfigurationService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseModuleConfigurationModel, ids, fields, relations)
}

func (svc *BaseModuleConfigurationService) Delete(ids []int64) error {
	return svc.client.delete(types.BaseModuleConfigurationModel, ids)
}
