package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockLocationService struct {
	client *Client
}

func NewStockLocationService(c *Client) *StockLocationService {
	return &StockLocationService{client: c}
}

func (svc *StockLocationService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.StockLocationModel, name)
}

func (svc *StockLocationService) GetByIds(ids []int) (*types.StockLocations, error) {
	s := &types.StockLocations{}
	return s, svc.client.getByIds(types.StockLocationModel, ids, s)
}

func (svc *StockLocationService) GetByName(name string) (*types.StockLocations, error) {
	s := &types.StockLocations{}
	return s, svc.client.getByName(types.StockLocationModel, name, s)
}

func (svc *StockLocationService) GetAll() (*types.StockLocations, error) {
	s := &types.StockLocations{}
	return s, svc.client.getAll(types.StockLocationModel, s)
}

func (svc *StockLocationService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.StockLocationModel, fields, relations)
}

func (svc *StockLocationService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockLocationModel, ids, fields, relations)
}

func (svc *StockLocationService) Delete(ids []int) error {
	return svc.client.delete(types.StockLocationModel, ids)
}
