package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockQuantService struct {
	client *Client
}

func NewStockQuantService(c *Client) *StockQuantService {
	return &StockQuantService{client: c}
}

func (svc *StockQuantService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.StockQuantModel, name)
}

func (svc *StockQuantService) GetByIds(ids []int) (*types.StockQuants, error) {
	s := &types.StockQuants{}
	return s, svc.client.getByIds(types.StockQuantModel, ids, s)
}

func (svc *StockQuantService) GetByName(name string) (*types.StockQuants, error) {
	s := &types.StockQuants{}
	return s, svc.client.getByName(types.StockQuantModel, name, s)
}

func (svc *StockQuantService) GetByField(field string, value string) (*types.StockQuants, error) {
	s := &types.StockQuants{}
	return s, svc.client.getByField(types.StockQuantModel, field, value, s)
}

func (svc *StockQuantService) GetAll() (*types.StockQuants, error) {
	s := &types.StockQuants{}
	return s, svc.client.getAll(types.StockQuantModel, s)
}

func (svc *StockQuantService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.StockQuantModel, fields, relations)
}

func (svc *StockQuantService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockQuantModel, ids, fields, relations)
}

func (svc *StockQuantService) Delete(ids []int) error {
	return svc.client.delete(types.StockQuantModel, ids)
}
