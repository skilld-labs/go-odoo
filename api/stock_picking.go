package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockPickingService struct {
	client *Client
}

func NewStockPickingService(c *Client) *StockPickingService {
	return &StockPickingService{client: c}
}

func (svc *StockPickingService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.StockPickingModel, name)
}

func (svc *StockPickingService) GetByIds(ids []int) (*types.StockPickings, error) {
	s := &types.StockPickings{}
	return s, svc.client.getByIds(types.StockPickingModel, ids, s)
}

func (svc *StockPickingService) GetByName(name string) (*types.StockPickings, error) {
	s := &types.StockPickings{}
	return s, svc.client.getByName(types.StockPickingModel, name, s)
}

func (svc *StockPickingService) GetByField(field string, value string) (*types.StockPickings, error) {
	s := &types.StockPickings{}
	return s, svc.client.getByName(types.StockPickingModel, field, value, s)
}

func (svc *StockPickingService) GetAll() (*types.StockPickings, error) {
	s := &types.StockPickings{}
	return s, svc.client.getAll(types.StockPickingModel, s)
}

func (svc *StockPickingService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.StockPickingModel, fields, relations)
}

func (svc *StockPickingService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockPickingModel, ids, fields, relations)
}

func (svc *StockPickingService) Delete(ids []int) error {
	return svc.client.delete(types.StockPickingModel, ids)
}
