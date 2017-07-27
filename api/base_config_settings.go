package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseConfigSettingsService struct {
	client *Client
}

func NewBaseConfigSettingsService(c *Client) *BaseConfigSettingsService {
	return &BaseConfigSettingsService{client: c}
}

func (svc *BaseConfigSettingsService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BaseConfigSettingsModel, name)
}

func (svc *BaseConfigSettingsService) GetByIds(ids []int) (*types.BaseConfigSettingss, error) {
	b := &types.BaseConfigSettingss{}
	return b, svc.client.getByIds(types.BaseConfigSettingsModel, ids, b)
}

func (svc *BaseConfigSettingsService) GetByName(name string) (*types.BaseConfigSettingss, error) {
	b := &types.BaseConfigSettingss{}
	return b, svc.client.getByName(types.BaseConfigSettingsModel, name, b)
}

func (svc *BaseConfigSettingsService) GetAll() (*types.BaseConfigSettingss, error) {
	b := &types.BaseConfigSettingss{}
	return b, svc.client.getAll(types.BaseConfigSettingsModel, b)
}

func (svc *BaseConfigSettingsService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BaseConfigSettingsModel, fields, relations)
}

func (svc *BaseConfigSettingsService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseConfigSettingsModel, ids, fields, relations)
}

func (svc *BaseConfigSettingsService) Delete(ids []int) error {
	return svc.client.delete(types.BaseConfigSettingsModel, ids)
}
