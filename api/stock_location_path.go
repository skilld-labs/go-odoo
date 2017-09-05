package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockLocationPathService struct {
	client *Client
}

func NewStockLocationPathService(c *Client) *StockLocationPathService {
	return &StockLocationPathService{client: c}
}

func (svc *StockLocationPathService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.StockLocationPathModel, name)
}

func (svc *StockLocationPathService) GetByIds(ids []int) (*types.StockLocationPaths, error) {
	s := &types.StockLocationPaths{}
	return s, svc.client.getByIds(types.StockLocationPathModel, ids, s)
}

func (svc *StockLocationPathService) GetByName(name string) (*types.StockLocationPaths, error) {
	s := &types.StockLocationPaths{}
	return s, svc.client.getByName(types.StockLocationPathModel, name, s)
}

func (svc *StockLocationPathService) GetByField(field string, value string) (*types.StockLocationPaths, error) {
	s := &types.StockLocationPaths{}
	return s, svc.client.getByName(types.StockLocationPathModel, field, value, s)
}

func (svc *StockLocationPathService) GetAll() (*types.StockLocationPaths, error) {
	s := &types.StockLocationPaths{}
	return s, svc.client.getAll(types.StockLocationPathModel, s)
}

func (svc *StockLocationPathService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.StockLocationPathModel, fields, relations)
}

func (svc *StockLocationPathService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockLocationPathModel, ids, fields, relations)
}

func (svc *StockLocationPathService) Delete(ids []int) error {
	return svc.client.delete(types.StockLocationPathModel, ids)
}
