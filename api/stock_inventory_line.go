package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockInventoryLineService struct {
	client *Client
}

func NewStockInventoryLineService(c *Client) *StockInventoryLineService {
	return &StockInventoryLineService{client: c}
}

func (svc *StockInventoryLineService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.StockInventoryLineModel, name)
}

func (svc *StockInventoryLineService) GetByIds(ids []int64) (*types.StockInventoryLines, error) {
	s := &types.StockInventoryLines{}
	return s, svc.client.getByIds(types.StockInventoryLineModel, ids, s)
}

func (svc *StockInventoryLineService) GetByName(name string) (*types.StockInventoryLines, error) {
	s := &types.StockInventoryLines{}
	return s, svc.client.getByName(types.StockInventoryLineModel, name, s)
}

func (svc *StockInventoryLineService) GetByField(field string, value string) (*types.StockInventoryLines, error) {
	s := &types.StockInventoryLines{}
	return s, svc.client.getByField(types.StockInventoryLineModel, field, value, s)
}

func (svc *StockInventoryLineService) GetAll() (*types.StockInventoryLines, error) {
	s := &types.StockInventoryLines{}
	return s, svc.client.getAll(types.StockInventoryLineModel, s)
}

func (svc *StockInventoryLineService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.StockInventoryLineModel, fields, relations)
}

func (svc *StockInventoryLineService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockInventoryLineModel, ids, fields, relations)
}

func (svc *StockInventoryLineService) Delete(ids []int64) error {
	return svc.client.delete(types.StockInventoryLineModel, ids)
}
