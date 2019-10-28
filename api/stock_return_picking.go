package api

import (
	"github.com/morezig/go-odoo/types"
)

type StockReturnPickingService struct {
	client *Client
}

func NewStockReturnPickingService(c *Client) *StockReturnPickingService {
	return &StockReturnPickingService{client: c}
}

func (svc *StockReturnPickingService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.StockReturnPickingModel, name)
}

func (svc *StockReturnPickingService) GetByIds(ids []int64) (*types.StockReturnPickings, error) {
	s := &types.StockReturnPickings{}
	return s, svc.client.getByIds(types.StockReturnPickingModel, ids, s)
}

func (svc *StockReturnPickingService) GetByName(name string) (*types.StockReturnPickings, error) {
	s := &types.StockReturnPickings{}
	return s, svc.client.getByName(types.StockReturnPickingModel, name, s)
}

func (svc *StockReturnPickingService) GetByField(field string, value string) (*types.StockReturnPickings, error) {
	s := &types.StockReturnPickings{}
	return s, svc.client.getByField(types.StockReturnPickingModel, field, value, s)
}

func (svc *StockReturnPickingService) GetAll() (*types.StockReturnPickings, error) {
	s := &types.StockReturnPickings{}
	return s, svc.client.getAll(types.StockReturnPickingModel, s)
}

func (svc *StockReturnPickingService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.StockReturnPickingModel, fields, relations)
}

func (svc *StockReturnPickingService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockReturnPickingModel, ids, fields, relations)
}

func (svc *StockReturnPickingService) Delete(ids []int64) error {
	return svc.client.delete(types.StockReturnPickingModel, ids)
}
