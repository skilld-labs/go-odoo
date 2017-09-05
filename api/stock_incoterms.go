package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockIncotermsService struct {
	client *Client
}

func NewStockIncotermsService(c *Client) *StockIncotermsService {
	return &StockIncotermsService{client: c}
}

func (svc *StockIncotermsService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.StockIncotermsModel, name)
}

func (svc *StockIncotermsService) GetByIds(ids []int) (*types.StockIncotermss, error) {
	s := &types.StockIncotermss{}
	return s, svc.client.getByIds(types.StockIncotermsModel, ids, s)
}

func (svc *StockIncotermsService) GetByName(name string) (*types.StockIncotermss, error) {
	s := &types.StockIncotermss{}
	return s, svc.client.getByName(types.StockIncotermsModel, name, s)
}

func (svc *StockIncotermsService) GetByField(field string, value string) (*types.StockIncotermss, error) {
	s := &types.StockIncotermss{}
	return s, svc.client.getByField(types.StockIncotermsModel, field, value, s)
}

func (svc *StockIncotermsService) GetAll() (*types.StockIncotermss, error) {
	s := &types.StockIncotermss{}
	return s, svc.client.getAll(types.StockIncotermsModel, s)
}

func (svc *StockIncotermsService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.StockIncotermsModel, fields, relations)
}

func (svc *StockIncotermsService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockIncotermsModel, ids, fields, relations)
}

func (svc *StockIncotermsService) Delete(ids []int) error {
	return svc.client.delete(types.StockIncotermsModel, ids)
}
