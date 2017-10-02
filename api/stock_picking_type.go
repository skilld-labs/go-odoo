package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockPickingTypeService struct {
	client *Client
}

func NewStockPickingTypeService(c *Client) *StockPickingTypeService {
	return &StockPickingTypeService{client: c}
}

func (svc *StockPickingTypeService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.StockPickingTypeModel, name)
}

func (svc *StockPickingTypeService) GetByIds(ids []int64) (*types.StockPickingTypes, error) {
	s := &types.StockPickingTypes{}
	return s, svc.client.getByIds(types.StockPickingTypeModel, ids, s)
}

func (svc *StockPickingTypeService) GetByName(name string) (*types.StockPickingTypes, error) {
	s := &types.StockPickingTypes{}
	return s, svc.client.getByName(types.StockPickingTypeModel, name, s)
}

func (svc *StockPickingTypeService) GetByField(field string, value string) (*types.StockPickingTypes, error) {
	s := &types.StockPickingTypes{}
	return s, svc.client.getByField(types.StockPickingTypeModel, field, value, s)
}

func (svc *StockPickingTypeService) GetAll() (*types.StockPickingTypes, error) {
	s := &types.StockPickingTypes{}
	return s, svc.client.getAll(types.StockPickingTypeModel, s)
}

func (svc *StockPickingTypeService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.StockPickingTypeModel, fields, relations)
}

func (svc *StockPickingTypeService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockPickingTypeModel, ids, fields, relations)
}

func (svc *StockPickingTypeService) Delete(ids []int64) error {
	return svc.client.delete(types.StockPickingTypeModel, ids)
}
