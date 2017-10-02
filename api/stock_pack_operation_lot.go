package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockPackOperationLotService struct {
	client *Client
}

func NewStockPackOperationLotService(c *Client) *StockPackOperationLotService {
	return &StockPackOperationLotService{client: c}
}

func (svc *StockPackOperationLotService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.StockPackOperationLotModel, name)
}

func (svc *StockPackOperationLotService) GetByIds(ids []int64) (*types.StockPackOperationLots, error) {
	s := &types.StockPackOperationLots{}
	return s, svc.client.getByIds(types.StockPackOperationLotModel, ids, s)
}

func (svc *StockPackOperationLotService) GetByName(name string) (*types.StockPackOperationLots, error) {
	s := &types.StockPackOperationLots{}
	return s, svc.client.getByName(types.StockPackOperationLotModel, name, s)
}

func (svc *StockPackOperationLotService) GetByField(field string, value string) (*types.StockPackOperationLots, error) {
	s := &types.StockPackOperationLots{}
	return s, svc.client.getByField(types.StockPackOperationLotModel, field, value, s)
}

func (svc *StockPackOperationLotService) GetAll() (*types.StockPackOperationLots, error) {
	s := &types.StockPackOperationLots{}
	return s, svc.client.getAll(types.StockPackOperationLotModel, s)
}

func (svc *StockPackOperationLotService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.StockPackOperationLotModel, fields, relations)
}

func (svc *StockPackOperationLotService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockPackOperationLotModel, ids, fields, relations)
}

func (svc *StockPackOperationLotService) Delete(ids []int64) error {
	return svc.client.delete(types.StockPackOperationLotModel, ids)
}
