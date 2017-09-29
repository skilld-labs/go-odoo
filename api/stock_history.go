package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockHistoryService struct {
	client *Client
}

func NewStockHistoryService(c *Client) *StockHistoryService {
	return &StockHistoryService{client: c}
}

func (svc *StockHistoryService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.StockHistoryModel, name)
}

func (svc *StockHistoryService) GetByIds(ids []int64) (*types.StockHistorys, error) {
	s := &types.StockHistorys{}
	return s, svc.client.getByIds(types.StockHistoryModel, ids, s)
}

func (svc *StockHistoryService) GetByName(name string) (*types.StockHistorys, error) {
	s := &types.StockHistorys{}
	return s, svc.client.getByName(types.StockHistoryModel, name, s)
}

func (svc *StockHistoryService) GetByField(field string, value string) (*types.StockHistorys, error) {
	s := &types.StockHistorys{}
	return s, svc.client.getByField(types.StockHistoryModel, field, value, s)
}

func (svc *StockHistoryService) GetAll() (*types.StockHistorys, error) {
	s := &types.StockHistorys{}
	return s, svc.client.getAll(types.StockHistoryModel, s)
}

func (svc *StockHistoryService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.StockHistoryModel, fields, relations)
}

func (svc *StockHistoryService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockHistoryModel, ids, fields, relations)
}

func (svc *StockHistoryService) Delete(ids []int64) error {
	return svc.client.delete(types.StockHistoryModel, ids)
}
