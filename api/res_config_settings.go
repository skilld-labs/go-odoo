package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ResConfigSettingsService struct {
	client *Client
}

func NewResConfigSettingsService(c *Client) *ResConfigSettingsService {
	return &ResConfigSettingsService{client: c}
}

func (svc *ResConfigSettingsService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ResConfigSettingsModel, name)
}

func (svc *ResConfigSettingsService) GetByIds(ids []int) (*types.ResConfigSettingss, error) {
	r := &types.ResConfigSettingss{}
	return r, svc.client.getByIds(types.ResConfigSettingsModel, ids, r)
}

func (svc *ResConfigSettingsService) GetByName(name string) (*types.ResConfigSettingss, error) {
	r := &types.ResConfigSettingss{}
	return r, svc.client.getByName(types.ResConfigSettingsModel, name, r)
}

func (svc *ResConfigSettingsService) GetByField(field string, value string) (*types.ResConfigSettingss, error) {
	r := &types.ResConfigSettingss{}
	return r, svc.client.getByField(types.ResConfigSettingsModel, field, value, r)
}

func (svc *ResConfigSettingsService) GetAll() (*types.ResConfigSettingss, error) {
	r := &types.ResConfigSettingss{}
	return r, svc.client.getAll(types.ResConfigSettingsModel, r)
}

func (svc *ResConfigSettingsService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ResConfigSettingsModel, fields, relations)
}

func (svc *ResConfigSettingsService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResConfigSettingsModel, ids, fields, relations)
}

func (svc *ResConfigSettingsService) Delete(ids []int) error {
	return svc.client.delete(types.ResConfigSettingsModel, ids)
}
