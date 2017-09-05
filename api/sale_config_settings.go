package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type SaleConfigSettingsService struct {
	client *Client
}

func NewSaleConfigSettingsService(c *Client) *SaleConfigSettingsService {
	return &SaleConfigSettingsService{client: c}
}

func (svc *SaleConfigSettingsService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.SaleConfigSettingsModel, name)
}

func (svc *SaleConfigSettingsService) GetByIds(ids []int) (*types.SaleConfigSettingss, error) {
	s := &types.SaleConfigSettingss{}
	return s, svc.client.getByIds(types.SaleConfigSettingsModel, ids, s)
}

func (svc *SaleConfigSettingsService) GetByName(name string) (*types.SaleConfigSettingss, error) {
	s := &types.SaleConfigSettingss{}
	return s, svc.client.getByName(types.SaleConfigSettingsModel, name, s)
}

func (svc *SaleConfigSettingsService) GetByField(field string, value string) (*types.SaleConfigSettingss, error) {
	s := &types.SaleConfigSettingss{}
	return s, svc.client.getByName(types.SaleConfigSettingsModel, field, value, s)
}

func (svc *SaleConfigSettingsService) GetAll() (*types.SaleConfigSettingss, error) {
	s := &types.SaleConfigSettingss{}
	return s, svc.client.getAll(types.SaleConfigSettingsModel, s)
}

func (svc *SaleConfigSettingsService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.SaleConfigSettingsModel, fields, relations)
}

func (svc *SaleConfigSettingsService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.SaleConfigSettingsModel, ids, fields, relations)
}

func (svc *SaleConfigSettingsService) Delete(ids []int) error {
	return svc.client.delete(types.SaleConfigSettingsModel, ids)
}
