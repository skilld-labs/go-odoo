package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockConfigSettingsService struct {
	client *Client
}

func NewStockConfigSettingsService(c *Client) *StockConfigSettingsService {
	return &StockConfigSettingsService{client: c}
}

func (svc *StockConfigSettingsService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.StockConfigSettingsModel, name)
}

func (svc *StockConfigSettingsService) GetByIds(ids []int) (*types.StockConfigSettingss, error) {
	s := &types.StockConfigSettingss{}
	return s, svc.client.getByIds(types.StockConfigSettingsModel, ids, s)
}

func (svc *StockConfigSettingsService) GetByName(name string) (*types.StockConfigSettingss, error) {
	s := &types.StockConfigSettingss{}
	return s, svc.client.getByName(types.StockConfigSettingsModel, name, s)
}

func (svc *StockConfigSettingsService) GetAll() (*types.StockConfigSettingss, error) {
	s := &types.StockConfigSettingss{}
	return s, svc.client.getAll(types.StockConfigSettingsModel, s)
}

func (svc *StockConfigSettingsService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.StockConfigSettingsModel, fields, relations)
}

func (svc *StockConfigSettingsService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockConfigSettingsModel, ids, fields, relations)
}

func (svc *StockConfigSettingsService) Delete(ids []int) error {
	return svc.client.delete(types.StockConfigSettingsModel, ids)
}
