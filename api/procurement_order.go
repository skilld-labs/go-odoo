package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ProcurementOrderService struct {
	client *Client
}

func NewProcurementOrderService(c *Client) *ProcurementOrderService {
	return &ProcurementOrderService{client: c}
}

func (svc *ProcurementOrderService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ProcurementOrderModel, name)
}

func (svc *ProcurementOrderService) GetByIds(ids []int) (*types.ProcurementOrders, error) {
	p := &types.ProcurementOrders{}
	return p, svc.client.getByIds(types.ProcurementOrderModel, ids, p)
}

func (svc *ProcurementOrderService) GetByName(name string) (*types.ProcurementOrders, error) {
	p := &types.ProcurementOrders{}
	return p, svc.client.getByName(types.ProcurementOrderModel, name, p)
}

func (svc *ProcurementOrderService) GetByField(field string, value string) (*types.ProcurementOrders, error) {
	p := &types.ProcurementOrders{}
	return p, svc.client.getByField(types.ProcurementOrderModel, field, value, p)
}

func (svc *ProcurementOrderService) GetAll() (*types.ProcurementOrders, error) {
	p := &types.ProcurementOrders{}
	return p, svc.client.getAll(types.ProcurementOrderModel, p)
}

func (svc *ProcurementOrderService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ProcurementOrderModel, fields, relations)
}

func (svc *ProcurementOrderService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProcurementOrderModel, ids, fields, relations)
}

func (svc *ProcurementOrderService) Delete(ids []int) error {
	return svc.client.delete(types.ProcurementOrderModel, ids)
}
