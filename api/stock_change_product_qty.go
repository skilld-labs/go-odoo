package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockChangeProductQtyService struct {
	client *Client
}

func NewStockChangeProductQtyService(c *Client) *StockChangeProductQtyService {
	return &StockChangeProductQtyService{client: c}
}

func (svc *StockChangeProductQtyService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.StockChangeProductQtyModel, name)
}

func (svc *StockChangeProductQtyService) GetByIds(ids []int) (*types.StockChangeProductQtys, error) {
	s := &types.StockChangeProductQtys{}
	return s, svc.client.getByIds(types.StockChangeProductQtyModel, ids, s)
}

func (svc *StockChangeProductQtyService) GetByName(name string) (*types.StockChangeProductQtys, error) {
	s := &types.StockChangeProductQtys{}
	return s, svc.client.getByName(types.StockChangeProductQtyModel, name, s)
}

func (svc *StockChangeProductQtyService) GetByField(field string, value string) (*types.StockChangeProductQtys, error) {
	s := &types.StockChangeProductQtys{}
	return s, svc.client.getByName(types.StockChangeProductQtyModel, field, value, s)
}

func (svc *StockChangeProductQtyService) GetAll() (*types.StockChangeProductQtys, error) {
	s := &types.StockChangeProductQtys{}
	return s, svc.client.getAll(types.StockChangeProductQtyModel, s)
}

func (svc *StockChangeProductQtyService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.StockChangeProductQtyModel, fields, relations)
}

func (svc *StockChangeProductQtyService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockChangeProductQtyModel, ids, fields, relations)
}

func (svc *StockChangeProductQtyService) Delete(ids []int) error {
	return svc.client.delete(types.StockChangeProductQtyModel, ids)
}
