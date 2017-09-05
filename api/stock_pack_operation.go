package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockPackOperationService struct {
	client *Client
}

func NewStockPackOperationService(c *Client) *StockPackOperationService {
	return &StockPackOperationService{client: c}
}

func (svc *StockPackOperationService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.StockPackOperationModel, name)
}

func (svc *StockPackOperationService) GetByIds(ids []int) (*types.StockPackOperations, error) {
	s := &types.StockPackOperations{}
	return s, svc.client.getByIds(types.StockPackOperationModel, ids, s)
}

func (svc *StockPackOperationService) GetByName(name string) (*types.StockPackOperations, error) {
	s := &types.StockPackOperations{}
	return s, svc.client.getByName(types.StockPackOperationModel, name, s)
}

func (svc *StockPackOperationService) GetByField(field string, value string) (*types.StockPackOperations, error) {
	s := &types.StockPackOperations{}
	return s, svc.client.getByName(types.StockPackOperationModel, field, value, s)
}

func (svc *StockPackOperationService) GetAll() (*types.StockPackOperations, error) {
	s := &types.StockPackOperations{}
	return s, svc.client.getAll(types.StockPackOperationModel, s)
}

func (svc *StockPackOperationService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.StockPackOperationModel, fields, relations)
}

func (svc *StockPackOperationService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockPackOperationModel, ids, fields, relations)
}

func (svc *StockPackOperationService) Delete(ids []int) error {
	return svc.client.delete(types.StockPackOperationModel, ids)
}
