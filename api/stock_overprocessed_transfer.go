package api

import (
	"github.com/morezig/go-odoo/types"
)

type StockOverprocessedTransferService struct {
	client *Client
}

func NewStockOverprocessedTransferService(c *Client) *StockOverprocessedTransferService {
	return &StockOverprocessedTransferService{client: c}
}

func (svc *StockOverprocessedTransferService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.StockOverprocessedTransferModel, name)
}

func (svc *StockOverprocessedTransferService) GetByIds(ids []int64) (*types.StockOverprocessedTransfers, error) {
	s := &types.StockOverprocessedTransfers{}
	return s, svc.client.getByIds(types.StockOverprocessedTransferModel, ids, s)
}

func (svc *StockOverprocessedTransferService) GetByName(name string) (*types.StockOverprocessedTransfers, error) {
	s := &types.StockOverprocessedTransfers{}
	return s, svc.client.getByName(types.StockOverprocessedTransferModel, name, s)
}

func (svc *StockOverprocessedTransferService) GetByField(field string, value string) (*types.StockOverprocessedTransfers, error) {
	s := &types.StockOverprocessedTransfers{}
	return s, svc.client.getByField(types.StockOverprocessedTransferModel, field, value, s)
}

func (svc *StockOverprocessedTransferService) GetAll() (*types.StockOverprocessedTransfers, error) {
	s := &types.StockOverprocessedTransfers{}
	return s, svc.client.getAll(types.StockOverprocessedTransferModel, s)
}

func (svc *StockOverprocessedTransferService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.StockOverprocessedTransferModel, fields, relations)
}

func (svc *StockOverprocessedTransferService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockOverprocessedTransferModel, ids, fields, relations)
}

func (svc *StockOverprocessedTransferService) Delete(ids []int64) error {
	return svc.client.delete(types.StockOverprocessedTransferModel, ids)
}
