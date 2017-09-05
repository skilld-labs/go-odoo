package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockProductionLotService struct {
	client *Client
}

func NewStockProductionLotService(c *Client) *StockProductionLotService {
	return &StockProductionLotService{client: c}
}

func (svc *StockProductionLotService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.StockProductionLotModel, name)
}

func (svc *StockProductionLotService) GetByIds(ids []int) (*types.StockProductionLots, error) {
	s := &types.StockProductionLots{}
	return s, svc.client.getByIds(types.StockProductionLotModel, ids, s)
}

func (svc *StockProductionLotService) GetByName(name string) (*types.StockProductionLots, error) {
	s := &types.StockProductionLots{}
	return s, svc.client.getByName(types.StockProductionLotModel, name, s)
}

func (svc *StockProductionLotService) GetByField(field string, value string) (*types.StockProductionLots, error) {
	s := &types.StockProductionLots{}
	return s, svc.client.getByField(types.StockProductionLotModel, field, value, s)
}

func (svc *StockProductionLotService) GetAll() (*types.StockProductionLots, error) {
	s := &types.StockProductionLots{}
	return s, svc.client.getAll(types.StockProductionLotModel, s)
}

func (svc *StockProductionLotService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.StockProductionLotModel, fields, relations)
}

func (svc *StockProductionLotService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockProductionLotModel, ids, fields, relations)
}

func (svc *StockProductionLotService) Delete(ids []int) error {
	return svc.client.delete(types.StockProductionLotModel, ids)
}
