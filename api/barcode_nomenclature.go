package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BarcodeNomenclatureService struct {
	client *Client
}

func NewBarcodeNomenclatureService(c *Client) *BarcodeNomenclatureService {
	return &BarcodeNomenclatureService{client: c}
}

func (svc *BarcodeNomenclatureService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BarcodeNomenclatureModel, name)
}

func (svc *BarcodeNomenclatureService) GetByIds(ids []int) (*types.BarcodeNomenclatures, error) {
	b := &types.BarcodeNomenclatures{}
	return b, svc.client.getByIds(types.BarcodeNomenclatureModel, ids, b)
}

func (svc *BarcodeNomenclatureService) GetByName(name string) (*types.BarcodeNomenclatures, error) {
	b := &types.BarcodeNomenclatures{}
	return b, svc.client.getByName(types.BarcodeNomenclatureModel, name, b)
}

func (svc *BarcodeNomenclatureService) GetByField(field string, value string) (*types.BarcodeNomenclatures, error) {
	b := &types.BarcodeNomenclatures{}
	return b, svc.client.getByField(types.BarcodeNomenclatureModel, field, value, b)
}

func (svc *BarcodeNomenclatureService) GetAll() (*types.BarcodeNomenclatures, error) {
	b := &types.BarcodeNomenclatures{}
	return b, svc.client.getAll(types.BarcodeNomenclatureModel, b)
}

func (svc *BarcodeNomenclatureService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BarcodeNomenclatureModel, fields, relations)
}

func (svc *BarcodeNomenclatureService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BarcodeNomenclatureModel, ids, fields, relations)
}

func (svc *BarcodeNomenclatureService) Delete(ids []int) error {
	return svc.client.delete(types.BarcodeNomenclatureModel, ids)
}
