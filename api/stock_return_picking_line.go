package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockReturnPickingLineService struct {
	client *Client
}

func NewStockReturnPickingLineService(c *Client) *StockReturnPickingLineService {
	return &StockReturnPickingLineService{client: c}
}

func (svc *StockReturnPickingLineService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.StockReturnPickingLineModel, name)
}

func (svc *StockReturnPickingLineService) GetByIds(ids []int64) (*types.StockReturnPickingLines, error) {
	s := &types.StockReturnPickingLines{}
	return s, svc.client.getByIds(types.StockReturnPickingLineModel, ids, s)
}

func (svc *StockReturnPickingLineService) GetByName(name string) (*types.StockReturnPickingLines, error) {
	s := &types.StockReturnPickingLines{}
	return s, svc.client.getByName(types.StockReturnPickingLineModel, name, s)
}

func (svc *StockReturnPickingLineService) GetByField(field string, value string) (*types.StockReturnPickingLines, error) {
	s := &types.StockReturnPickingLines{}
	return s, svc.client.getByField(types.StockReturnPickingLineModel, field, value, s)
}

func (svc *StockReturnPickingLineService) GetAll() (*types.StockReturnPickingLines, error) {
	s := &types.StockReturnPickingLines{}
	return s, svc.client.getAll(types.StockReturnPickingLineModel, s)
}

func (svc *StockReturnPickingLineService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.StockReturnPickingLineModel, fields, relations)
}

func (svc *StockReturnPickingLineService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockReturnPickingLineModel, ids, fields, relations)
}

func (svc *StockReturnPickingLineService) Delete(ids []int64) error {
	return svc.client.delete(types.StockReturnPickingLineModel, ids)
}
