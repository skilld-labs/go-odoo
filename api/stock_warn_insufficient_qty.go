package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockWarnInsufficientQtyService struct {
	client *Client
}

func NewStockWarnInsufficientQtyService(c *Client) *StockWarnInsufficientQtyService {
	return &StockWarnInsufficientQtyService{client: c}
}

func (svc *StockWarnInsufficientQtyService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.StockWarnInsufficientQtyModel, name)
}

func (svc *StockWarnInsufficientQtyService) GetByIds(ids []int) (*types.StockWarnInsufficientQtys, error) {
	s := &types.StockWarnInsufficientQtys{}
	return s, svc.client.getByIds(types.StockWarnInsufficientQtyModel, ids, s)
}

func (svc *StockWarnInsufficientQtyService) GetByName(name string) (*types.StockWarnInsufficientQtys, error) {
	s := &types.StockWarnInsufficientQtys{}
	return s, svc.client.getByName(types.StockWarnInsufficientQtyModel, name, s)
}

func (svc *StockWarnInsufficientQtyService) GetByField(field string, value string) (*types.StockWarnInsufficientQtys, error) {
	s := &types.StockWarnInsufficientQtys{}
	return s, svc.client.getByField(types.StockWarnInsufficientQtyModel, field, value, s)
}

func (svc *StockWarnInsufficientQtyService) GetAll() (*types.StockWarnInsufficientQtys, error) {
	s := &types.StockWarnInsufficientQtys{}
	return s, svc.client.getAll(types.StockWarnInsufficientQtyModel, s)
}

func (svc *StockWarnInsufficientQtyService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.StockWarnInsufficientQtyModel, fields, relations)
}

func (svc *StockWarnInsufficientQtyService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockWarnInsufficientQtyModel, ids, fields, relations)
}

func (svc *StockWarnInsufficientQtyService) Delete(ids []int) error {
	return svc.client.delete(types.StockWarnInsufficientQtyModel, ids)
}
