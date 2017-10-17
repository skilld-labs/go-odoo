package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockMoveLineService struct {
	client *Client
}

func NewStockMoveLineService(c *Client) *StockMoveLineService {
	return &StockMoveLineService{client: c}
}

func (svc *StockMoveLineService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.StockMoveLineModel, name)
}

func (svc *StockMoveLineService) GetByIds(ids []int) (*types.StockMoveLines, error) {
	s := &types.StockMoveLines{}
	return s, svc.client.getByIds(types.StockMoveLineModel, ids, s)
}

func (svc *StockMoveLineService) GetByName(name string) (*types.StockMoveLines, error) {
	s := &types.StockMoveLines{}
	return s, svc.client.getByName(types.StockMoveLineModel, name, s)
}

func (svc *StockMoveLineService) GetByField(field string, value string) (*types.StockMoveLines, error) {
	s := &types.StockMoveLines{}
	return s, svc.client.getByField(types.StockMoveLineModel, field, value, s)
}

func (svc *StockMoveLineService) GetAll() (*types.StockMoveLines, error) {
	s := &types.StockMoveLines{}
	return s, svc.client.getAll(types.StockMoveLineModel, s)
}

func (svc *StockMoveLineService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.StockMoveLineModel, fields, relations)
}

func (svc *StockMoveLineService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockMoveLineModel, ids, fields, relations)
}

func (svc *StockMoveLineService) Delete(ids []int) error {
	return svc.client.delete(types.StockMoveLineModel, ids)
}
