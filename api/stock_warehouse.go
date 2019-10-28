package api

import (
	"github.com/morezig/go-odoo/types"
)

type StockWarehouseService struct {
	client *Client
}

func NewStockWarehouseService(c *Client) *StockWarehouseService {
	return &StockWarehouseService{client: c}
}

func (svc *StockWarehouseService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.StockWarehouseModel, name)
}

func (svc *StockWarehouseService) GetByIds(ids []int64) (*types.StockWarehouses, error) {
	s := &types.StockWarehouses{}
	return s, svc.client.getByIds(types.StockWarehouseModel, ids, s)
}

func (svc *StockWarehouseService) GetByName(name string) (*types.StockWarehouses, error) {
	s := &types.StockWarehouses{}
	return s, svc.client.getByName(types.StockWarehouseModel, name, s)
}

func (svc *StockWarehouseService) GetByField(field string, value string) (*types.StockWarehouses, error) {
	s := &types.StockWarehouses{}
	return s, svc.client.getByField(types.StockWarehouseModel, field, value, s)
}

func (svc *StockWarehouseService) GetAll() (*types.StockWarehouses, error) {
	s := &types.StockWarehouses{}
	return s, svc.client.getAll(types.StockWarehouseModel, s)
}

func (svc *StockWarehouseService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.StockWarehouseModel, fields, relations)
}

func (svc *StockWarehouseService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockWarehouseModel, ids, fields, relations)
}

func (svc *StockWarehouseService) Delete(ids []int64) error {
	return svc.client.delete(types.StockWarehouseModel, ids)
}
