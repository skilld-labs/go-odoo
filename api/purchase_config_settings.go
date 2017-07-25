package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type PurchaseConfigSettingsService struct {
	client *Client
}

func NewPurchaseConfigSettingsService(c *Client) *PurchaseConfigSettingsService {
	return &PurchaseConfigSettingsService{client: c}
}

func (svc *PurchaseConfigSettingsService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.PurchaseConfigSettingsModel, name)
}

func (svc *PurchaseConfigSettingsService) GetByIds(ids []int) (*types.PurchaseConfigSettingss, error) {
	p := &types.PurchaseConfigSettingss{}
	return p, svc.client.getByIds(types.PurchaseConfigSettingsModel, ids, p)
}

func (svc *PurchaseConfigSettingsService) GetByName(name string) (*types.PurchaseConfigSettingss, error) {
	p := &types.PurchaseConfigSettingss{}
	return p, svc.client.getByName(types.PurchaseConfigSettingsModel, name, p)
}

func (svc *PurchaseConfigSettingsService) GetAll() (*types.PurchaseConfigSettingss, error) {
	p := &types.PurchaseConfigSettingss{}
	return p, svc.client.getAll(types.PurchaseConfigSettingsModel, p)
}

func (svc *PurchaseConfigSettingsService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.PurchaseConfigSettingsModel, fields, relations)
}

func (svc *PurchaseConfigSettingsService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.PurchaseConfigSettingsModel, ids, fields, relations)
}

func (svc *PurchaseConfigSettingsService) Delete(ids []int) error {
	return svc.client.delete(types.PurchaseConfigSettingsModel, ids)
}
