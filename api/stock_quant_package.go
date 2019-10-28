package api

import (
	"github.com/morezig/go-odoo/types"
)

type StockQuantPackageService struct {
	client *Client
}

func NewStockQuantPackageService(c *Client) *StockQuantPackageService {
	return &StockQuantPackageService{client: c}
}

func (svc *StockQuantPackageService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.StockQuantPackageModel, name)
}

func (svc *StockQuantPackageService) GetByIds(ids []int64) (*types.StockQuantPackages, error) {
	s := &types.StockQuantPackages{}
	return s, svc.client.getByIds(types.StockQuantPackageModel, ids, s)
}

func (svc *StockQuantPackageService) GetByName(name string) (*types.StockQuantPackages, error) {
	s := &types.StockQuantPackages{}
	return s, svc.client.getByName(types.StockQuantPackageModel, name, s)
}

func (svc *StockQuantPackageService) GetByField(field string, value string) (*types.StockQuantPackages, error) {
	s := &types.StockQuantPackages{}
	return s, svc.client.getByField(types.StockQuantPackageModel, field, value, s)
}

func (svc *StockQuantPackageService) GetAll() (*types.StockQuantPackages, error) {
	s := &types.StockQuantPackages{}
	return s, svc.client.getAll(types.StockQuantPackageModel, s)
}

func (svc *StockQuantPackageService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.StockQuantPackageModel, fields, relations)
}

func (svc *StockQuantPackageService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockQuantPackageModel, ids, fields, relations)
}

func (svc *StockQuantPackageService) Delete(ids []int64) error {
	return svc.client.delete(types.StockQuantPackageModel, ids)
}
