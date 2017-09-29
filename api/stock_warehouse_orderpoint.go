package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockWarehouseOrderpointService struct {
	client *Client
}

func NewStockWarehouseOrderpointService(c *Client) *StockWarehouseOrderpointService {
	return &StockWarehouseOrderpointService{client: c}
}

func (svc *StockWarehouseOrderpointService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.StockWarehouseOrderpointModel, name)
}

func (svc *StockWarehouseOrderpointService) GetByIds(ids []int64) (*types.StockWarehouseOrderpoints, error) {
	s := &types.StockWarehouseOrderpoints{}
	return s, svc.client.getByIds(types.StockWarehouseOrderpointModel, ids, s)
}

func (svc *StockWarehouseOrderpointService) GetByName(name string) (*types.StockWarehouseOrderpoints, error) {
	s := &types.StockWarehouseOrderpoints{}
	return s, svc.client.getByName(types.StockWarehouseOrderpointModel, name, s)
}

func (svc *StockWarehouseOrderpointService) GetByField(field string, value string) (*types.StockWarehouseOrderpoints, error) {
	s := &types.StockWarehouseOrderpoints{}
	return s, svc.client.getByField(types.StockWarehouseOrderpointModel, field, value, s)
}

func (svc *StockWarehouseOrderpointService) GetAll() (*types.StockWarehouseOrderpoints, error) {
	s := &types.StockWarehouseOrderpoints{}
	return s, svc.client.getAll(types.StockWarehouseOrderpointModel, s)
}

func (svc *StockWarehouseOrderpointService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.StockWarehouseOrderpointModel, fields, relations)
}

func (svc *StockWarehouseOrderpointService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockWarehouseOrderpointModel, ids, fields, relations)
}

func (svc *StockWarehouseOrderpointService) Delete(ids []int64) error {
	return svc.client.delete(types.StockWarehouseOrderpointModel, ids)
}
