package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockBackorderConfirmationService struct {
	client *Client
}

func NewStockBackorderConfirmationService(c *Client) *StockBackorderConfirmationService {
	return &StockBackorderConfirmationService{client: c}
}

func (svc *StockBackorderConfirmationService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.StockBackorderConfirmationModel, name)
}

func (svc *StockBackorderConfirmationService) GetByIds(ids []int64) (*types.StockBackorderConfirmations, error) {
	s := &types.StockBackorderConfirmations{}
	return s, svc.client.getByIds(types.StockBackorderConfirmationModel, ids, s)
}

func (svc *StockBackorderConfirmationService) GetByName(name string) (*types.StockBackorderConfirmations, error) {
	s := &types.StockBackorderConfirmations{}
	return s, svc.client.getByName(types.StockBackorderConfirmationModel, name, s)
}

func (svc *StockBackorderConfirmationService) GetByField(field string, value string) (*types.StockBackorderConfirmations, error) {
	s := &types.StockBackorderConfirmations{}
	return s, svc.client.getByField(types.StockBackorderConfirmationModel, field, value, s)
}

func (svc *StockBackorderConfirmationService) GetAll() (*types.StockBackorderConfirmations, error) {
	s := &types.StockBackorderConfirmations{}
	return s, svc.client.getAll(types.StockBackorderConfirmationModel, s)
}

func (svc *StockBackorderConfirmationService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.StockBackorderConfirmationModel, fields, relations)
}

func (svc *StockBackorderConfirmationService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockBackorderConfirmationModel, ids, fields, relations)
}

func (svc *StockBackorderConfirmationService) Delete(ids []int64) error {
	return svc.client.delete(types.StockBackorderConfirmationModel, ids)
}
