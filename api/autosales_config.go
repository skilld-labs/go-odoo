package api

import (
	"github.com/morezig/go-odoo/types"
)

type AutosalesConfigService struct {
	client *Client
}

func NewAutosalesConfigService(c *Client) *AutosalesConfigService {
	return &AutosalesConfigService{client: c}
}

func (svc *AutosalesConfigService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AutosalesConfigModel, name)
}

func (svc *AutosalesConfigService) GetByIds(ids []int64) (*types.AutosalesConfigs, error) {
	a := &types.AutosalesConfigs{}
	return a, svc.client.getByIds(types.AutosalesConfigModel, ids, a)
}

func (svc *AutosalesConfigService) GetByName(name string) (*types.AutosalesConfigs, error) {
	a := &types.AutosalesConfigs{}
	return a, svc.client.getByName(types.AutosalesConfigModel, name, a)
}

func (svc *AutosalesConfigService) GetByField(field string, value string) (*types.AutosalesConfigs, error) {
	a := &types.AutosalesConfigs{}
	return a, svc.client.getByField(types.AutosalesConfigModel, field, value, a)
}

func (svc *AutosalesConfigService) GetAll() (*types.AutosalesConfigs, error) {
	a := &types.AutosalesConfigs{}
	return a, svc.client.getAll(types.AutosalesConfigModel, a)
}

func (svc *AutosalesConfigService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AutosalesConfigModel, fields, relations)
}

func (svc *AutosalesConfigService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AutosalesConfigModel, ids, fields, relations)
}

func (svc *AutosalesConfigService) Delete(ids []int64) error {
	return svc.client.delete(types.AutosalesConfigModel, ids)
}
