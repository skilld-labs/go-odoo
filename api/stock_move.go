package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockMoveService struct {
	client *Client
}

func NewStockMoveService(c *Client) *StockMoveService {
	return &StockMoveService{client: c}
}

func (svc *StockMoveService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.StockMoveModel, name)
}

func (svc *StockMoveService) GetByIds(ids []int64) (*types.StockMoves, error) {
	s := &types.StockMoves{}
	return s, svc.client.getByIds(types.StockMoveModel, ids, s)
}

func (svc *StockMoveService) GetByName(name string) (*types.StockMoves, error) {
	s := &types.StockMoves{}
	return s, svc.client.getByName(types.StockMoveModel, name, s)
}

func (svc *StockMoveService) GetByField(field string, value string) (*types.StockMoves, error) {
	s := &types.StockMoves{}
	return s, svc.client.getByField(types.StockMoveModel, field, value, s)
}

func (svc *StockMoveService) GetAll() (*types.StockMoves, error) {
	s := &types.StockMoves{}
	return s, svc.client.getAll(types.StockMoveModel, s)
}

func (svc *StockMoveService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.StockMoveModel, fields, relations)
}

func (svc *StockMoveService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockMoveModel, ids, fields, relations)
}

func (svc *StockMoveService) Delete(ids []int64) error {
	return svc.client.delete(types.StockMoveModel, ids)
}
