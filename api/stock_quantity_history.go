package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockQuantityHistoryService struct {
	client *Client
}

func NewStockQuantityHistoryService(c *Client) *StockQuantityHistoryService {
	return &StockQuantityHistoryService{client: c}
}

func (svc *StockQuantityHistoryService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.StockQuantityHistoryModel, name)
}

func (svc *StockQuantityHistoryService) GetByIds(ids []int) (*types.StockQuantityHistorys, error) {
	s := &types.StockQuantityHistorys{}
	return s, svc.client.getByIds(types.StockQuantityHistoryModel, ids, s)
}

func (svc *StockQuantityHistoryService) GetByName(name string) (*types.StockQuantityHistorys, error) {
	s := &types.StockQuantityHistorys{}
	return s, svc.client.getByName(types.StockQuantityHistoryModel, name, s)
}

func (svc *StockQuantityHistoryService) GetByField(field string, value string) (*types.StockQuantityHistorys, error) {
	s := &types.StockQuantityHistorys{}
	return s, svc.client.getByField(types.StockQuantityHistoryModel, field, value, s)
}

func (svc *StockQuantityHistoryService) GetAll() (*types.StockQuantityHistorys, error) {
	s := &types.StockQuantityHistorys{}
	return s, svc.client.getAll(types.StockQuantityHistoryModel, s)
}

func (svc *StockQuantityHistoryService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.StockQuantityHistoryModel, fields, relations)
}

func (svc *StockQuantityHistoryService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockQuantityHistoryModel, ids, fields, relations)
}

func (svc *StockQuantityHistoryService) Delete(ids []int) error {
	return svc.client.delete(types.StockQuantityHistoryModel, ids)
}
