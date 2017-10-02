package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BarcodesBarcodeEventsMixinService struct {
	client *Client
}

func NewBarcodesBarcodeEventsMixinService(c *Client) *BarcodesBarcodeEventsMixinService {
	return &BarcodesBarcodeEventsMixinService{client: c}
}

func (svc *BarcodesBarcodeEventsMixinService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.BarcodesBarcodeEventsMixinModel, name)
}

func (svc *BarcodesBarcodeEventsMixinService) GetByIds(ids []int64) (*types.BarcodesBarcodeEventsMixins, error) {
	b := &types.BarcodesBarcodeEventsMixins{}
	return b, svc.client.getByIds(types.BarcodesBarcodeEventsMixinModel, ids, b)
}

func (svc *BarcodesBarcodeEventsMixinService) GetByName(name string) (*types.BarcodesBarcodeEventsMixins, error) {
	b := &types.BarcodesBarcodeEventsMixins{}
	return b, svc.client.getByName(types.BarcodesBarcodeEventsMixinModel, name, b)
}

func (svc *BarcodesBarcodeEventsMixinService) GetByField(field string, value string) (*types.BarcodesBarcodeEventsMixins, error) {
	b := &types.BarcodesBarcodeEventsMixins{}
	return b, svc.client.getByField(types.BarcodesBarcodeEventsMixinModel, field, value, b)
}

func (svc *BarcodesBarcodeEventsMixinService) GetAll() (*types.BarcodesBarcodeEventsMixins, error) {
	b := &types.BarcodesBarcodeEventsMixins{}
	return b, svc.client.getAll(types.BarcodesBarcodeEventsMixinModel, b)
}

func (svc *BarcodesBarcodeEventsMixinService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.BarcodesBarcodeEventsMixinModel, fields, relations)
}

func (svc *BarcodesBarcodeEventsMixinService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BarcodesBarcodeEventsMixinModel, ids, fields, relations)
}

func (svc *BarcodesBarcodeEventsMixinService) Delete(ids []int64) error {
	return svc.client.delete(types.BarcodesBarcodeEventsMixinModel, ids)
}
