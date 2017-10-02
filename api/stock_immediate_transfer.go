package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockImmediateTransferService struct {
	client *Client
}

func NewStockImmediateTransferService(c *Client) *StockImmediateTransferService {
	return &StockImmediateTransferService{client: c}
}

func (svc *StockImmediateTransferService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.StockImmediateTransferModel, name)
}

func (svc *StockImmediateTransferService) GetByIds(ids []int64) (*types.StockImmediateTransfers, error) {
	s := &types.StockImmediateTransfers{}
	return s, svc.client.getByIds(types.StockImmediateTransferModel, ids, s)
}

func (svc *StockImmediateTransferService) GetByName(name string) (*types.StockImmediateTransfers, error) {
	s := &types.StockImmediateTransfers{}
	return s, svc.client.getByName(types.StockImmediateTransferModel, name, s)
}

func (svc *StockImmediateTransferService) GetByField(field string, value string) (*types.StockImmediateTransfers, error) {
	s := &types.StockImmediateTransfers{}
	return s, svc.client.getByField(types.StockImmediateTransferModel, field, value, s)
}

func (svc *StockImmediateTransferService) GetAll() (*types.StockImmediateTransfers, error) {
	s := &types.StockImmediateTransfers{}
	return s, svc.client.getAll(types.StockImmediateTransferModel, s)
}

func (svc *StockImmediateTransferService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.StockImmediateTransferModel, fields, relations)
}

func (svc *StockImmediateTransferService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockImmediateTransferModel, ids, fields, relations)
}

func (svc *StockImmediateTransferService) Delete(ids []int64) error {
	return svc.client.delete(types.StockImmediateTransferModel, ids)
}
