package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockScrapService struct {
	client *Client
}

func NewStockScrapService(c *Client) *StockScrapService {
	return &StockScrapService{client: c}
}

func (svc *StockScrapService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.StockScrapModel, name)
}

func (svc *StockScrapService) GetByIds(ids []int64) (*types.StockScraps, error) {
	s := &types.StockScraps{}
	return s, svc.client.getByIds(types.StockScrapModel, ids, s)
}

func (svc *StockScrapService) GetByName(name string) (*types.StockScraps, error) {
	s := &types.StockScraps{}
	return s, svc.client.getByName(types.StockScrapModel, name, s)
}

func (svc *StockScrapService) GetByField(field string, value string) (*types.StockScraps, error) {
	s := &types.StockScraps{}
	return s, svc.client.getByField(types.StockScrapModel, field, value, s)
}

func (svc *StockScrapService) GetAll() (*types.StockScraps, error) {
	s := &types.StockScraps{}
	return s, svc.client.getAll(types.StockScrapModel, s)
}

func (svc *StockScrapService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.StockScrapModel, fields, relations)
}

func (svc *StockScrapService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockScrapModel, ids, fields, relations)
}

func (svc *StockScrapService) Delete(ids []int64) error {
	return svc.client.delete(types.StockScrapModel, ids)
}
