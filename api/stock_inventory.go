package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockInventoryService struct {
	client *Client
}

func NewStockInventoryService(c *Client) *StockInventoryService {
	return &StockInventoryService{client: c}
}

func (svc *StockInventoryService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.StockInventoryModel, name)
}

func (svc *StockInventoryService) GetByIds(ids []int) (*types.StockInventorys, error) {
	s := &types.StockInventorys{}
	return s, svc.client.getByIds(types.StockInventoryModel, ids, s)
}

func (svc *StockInventoryService) GetByName(name string) (*types.StockInventorys, error) {
	s := &types.StockInventorys{}
	return s, svc.client.getByName(types.StockInventoryModel, name, s)
}

func (svc *StockInventoryService) GetByField(field string, value string) (*types.StockInventorys, error) {
	s := &types.StockInventorys{}
	return s, svc.client.getByField(types.StockInventoryModel, field, value, s)
}

func (svc *StockInventoryService) GetAll() (*types.StockInventorys, error) {
	s := &types.StockInventorys{}
	return s, svc.client.getAll(types.StockInventoryModel, s)
}

func (svc *StockInventoryService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.StockInventoryModel, fields, relations)
}

func (svc *StockInventoryService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockInventoryModel, ids, fields, relations)
}

func (svc *StockInventoryService) Delete(ids []int) error {
	return svc.client.delete(types.StockInventoryModel, ids)
}
