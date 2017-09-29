package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountConfigSettingsService struct {
	client *Client
}

func NewAccountConfigSettingsService(c *Client) *AccountConfigSettingsService {
	return &AccountConfigSettingsService{client: c}
}

func (svc *AccountConfigSettingsService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountConfigSettingsModel, name)
}

func (svc *AccountConfigSettingsService) GetByIds(ids []int64) (*types.AccountConfigSettingss, error) {
	a := &types.AccountConfigSettingss{}
	return a, svc.client.getByIds(types.AccountConfigSettingsModel, ids, a)
}

func (svc *AccountConfigSettingsService) GetByName(name string) (*types.AccountConfigSettingss, error) {
	a := &types.AccountConfigSettingss{}
	return a, svc.client.getByName(types.AccountConfigSettingsModel, name, a)
}

func (svc *AccountConfigSettingsService) GetByField(field string, value string) (*types.AccountConfigSettingss, error) {
	a := &types.AccountConfigSettingss{}
	return a, svc.client.getByField(types.AccountConfigSettingsModel, field, value, a)
}

func (svc *AccountConfigSettingsService) GetAll() (*types.AccountConfigSettingss, error) {
	a := &types.AccountConfigSettingss{}
	return a, svc.client.getAll(types.AccountConfigSettingsModel, a)
}

func (svc *AccountConfigSettingsService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountConfigSettingsModel, fields, relations)
}

func (svc *AccountConfigSettingsService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountConfigSettingsModel, ids, fields, relations)
}

func (svc *AccountConfigSettingsService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountConfigSettingsModel, ids)
}
