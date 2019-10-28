package api

import (
	"github.com/morezig/go-odoo/types"
)

type StockLocationRouteService struct {
	client *Client
}

func NewStockLocationRouteService(c *Client) *StockLocationRouteService {
	return &StockLocationRouteService{client: c}
}

func (svc *StockLocationRouteService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.StockLocationRouteModel, name)
}

func (svc *StockLocationRouteService) GetByIds(ids []int64) (*types.StockLocationRoutes, error) {
	s := &types.StockLocationRoutes{}
	return s, svc.client.getByIds(types.StockLocationRouteModel, ids, s)
}

func (svc *StockLocationRouteService) GetByName(name string) (*types.StockLocationRoutes, error) {
	s := &types.StockLocationRoutes{}
	return s, svc.client.getByName(types.StockLocationRouteModel, name, s)
}

func (svc *StockLocationRouteService) GetByField(field string, value string) (*types.StockLocationRoutes, error) {
	s := &types.StockLocationRoutes{}
	return s, svc.client.getByField(types.StockLocationRouteModel, field, value, s)
}

func (svc *StockLocationRouteService) GetAll() (*types.StockLocationRoutes, error) {
	s := &types.StockLocationRoutes{}
	return s, svc.client.getAll(types.StockLocationRouteModel, s)
}

func (svc *StockLocationRouteService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.StockLocationRouteModel, fields, relations)
}

func (svc *StockLocationRouteService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockLocationRouteModel, ids, fields, relations)
}

func (svc *StockLocationRouteService) Delete(ids []int64) error {
	return svc.client.delete(types.StockLocationRouteModel, ids)
}
