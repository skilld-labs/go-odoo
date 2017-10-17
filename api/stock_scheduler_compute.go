package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockSchedulerComputeService struct {
	client *Client
}

func NewStockSchedulerComputeService(c *Client) *StockSchedulerComputeService {
	return &StockSchedulerComputeService{client: c}
}

func (svc *StockSchedulerComputeService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.StockSchedulerComputeModel, name)
}

func (svc *StockSchedulerComputeService) GetByIds(ids []int) (*types.StockSchedulerComputes, error) {
	s := &types.StockSchedulerComputes{}
	return s, svc.client.getByIds(types.StockSchedulerComputeModel, ids, s)
}

func (svc *StockSchedulerComputeService) GetByName(name string) (*types.StockSchedulerComputes, error) {
	s := &types.StockSchedulerComputes{}
	return s, svc.client.getByName(types.StockSchedulerComputeModel, name, s)
}

func (svc *StockSchedulerComputeService) GetByField(field string, value string) (*types.StockSchedulerComputes, error) {
	s := &types.StockSchedulerComputes{}
	return s, svc.client.getByField(types.StockSchedulerComputeModel, field, value, s)
}

func (svc *StockSchedulerComputeService) GetAll() (*types.StockSchedulerComputes, error) {
	s := &types.StockSchedulerComputes{}
	return s, svc.client.getAll(types.StockSchedulerComputeModel, s)
}

func (svc *StockSchedulerComputeService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.StockSchedulerComputeModel, fields, relations)
}

func (svc *StockSchedulerComputeService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockSchedulerComputeModel, ids, fields, relations)
}

func (svc *StockSchedulerComputeService) Delete(ids []int) error {
	return svc.client.delete(types.StockSchedulerComputeModel, ids)
}
