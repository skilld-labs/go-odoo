package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MassMailingConfigSettingsService struct {
	client *Client
}

func NewMassMailingConfigSettingsService(c *Client) *MassMailingConfigSettingsService {
	return &MassMailingConfigSettingsService{client: c}
}

func (svc *MassMailingConfigSettingsService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.MassMailingConfigSettingsModel, name)
}

func (svc *MassMailingConfigSettingsService) GetByIds(ids []int64) (*types.MassMailingConfigSettingss, error) {
	m := &types.MassMailingConfigSettingss{}
	return m, svc.client.getByIds(types.MassMailingConfigSettingsModel, ids, m)
}

func (svc *MassMailingConfigSettingsService) GetByName(name string) (*types.MassMailingConfigSettingss, error) {
	m := &types.MassMailingConfigSettingss{}
	return m, svc.client.getByName(types.MassMailingConfigSettingsModel, name, m)
}

func (svc *MassMailingConfigSettingsService) GetByField(field string, value string) (*types.MassMailingConfigSettingss, error) {
	m := &types.MassMailingConfigSettingss{}
	return m, svc.client.getByField(types.MassMailingConfigSettingsModel, field, value, m)
}

func (svc *MassMailingConfigSettingsService) GetAll() (*types.MassMailingConfigSettingss, error) {
	m := &types.MassMailingConfigSettingss{}
	return m, svc.client.getAll(types.MassMailingConfigSettingsModel, m)
}

func (svc *MassMailingConfigSettingsService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.MassMailingConfigSettingsModel, fields, relations)
}

func (svc *MassMailingConfigSettingsService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MassMailingConfigSettingsModel, ids, fields, relations)
}

func (svc *MassMailingConfigSettingsService) Delete(ids []int64) error {
	return svc.client.delete(types.MassMailingConfigSettingsModel, ids)
}
