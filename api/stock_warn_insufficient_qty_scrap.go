package api

import (
	"github.com/morezig/go-odoo/types"
)

type StockWarnInsufficientQtyScrapService struct {
	client *Client
}

func NewStockWarnInsufficientQtyScrapService(c *Client) *StockWarnInsufficientQtyScrapService {
	return &StockWarnInsufficientQtyScrapService{client: c}
}

func (svc *StockWarnInsufficientQtyScrapService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.StockWarnInsufficientQtyScrapModel, name)
}

func (svc *StockWarnInsufficientQtyScrapService) GetByIds(ids []int64) (*types.StockWarnInsufficientQtyScraps, error) {
	s := &types.StockWarnInsufficientQtyScraps{}
	return s, svc.client.getByIds(types.StockWarnInsufficientQtyScrapModel, ids, s)
}

func (svc *StockWarnInsufficientQtyScrapService) GetByName(name string) (*types.StockWarnInsufficientQtyScraps, error) {
	s := &types.StockWarnInsufficientQtyScraps{}
	return s, svc.client.getByName(types.StockWarnInsufficientQtyScrapModel, name, s)
}

func (svc *StockWarnInsufficientQtyScrapService) GetByField(field string, value string) (*types.StockWarnInsufficientQtyScraps, error) {
	s := &types.StockWarnInsufficientQtyScraps{}
	return s, svc.client.getByField(types.StockWarnInsufficientQtyScrapModel, field, value, s)
}

func (svc *StockWarnInsufficientQtyScrapService) GetAll() (*types.StockWarnInsufficientQtyScraps, error) {
	s := &types.StockWarnInsufficientQtyScraps{}
	return s, svc.client.getAll(types.StockWarnInsufficientQtyScrapModel, s)
}

func (svc *StockWarnInsufficientQtyScrapService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.StockWarnInsufficientQtyScrapModel, fields, relations)
}

func (svc *StockWarnInsufficientQtyScrapService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockWarnInsufficientQtyScrapModel, ids, fields, relations)
}

func (svc *StockWarnInsufficientQtyScrapService) Delete(ids []int64) error {
	return svc.client.delete(types.StockWarnInsufficientQtyScrapModel, ids)
}
