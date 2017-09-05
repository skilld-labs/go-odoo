package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type SaleOrderService struct {
	client *Client
}

func NewSaleOrderService(c *Client) *SaleOrderService {
	return &SaleOrderService{client: c}
}

func (svc *SaleOrderService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.SaleOrderModel, name)
}

func (svc *SaleOrderService) GetByIds(ids []int) (*types.SaleOrders, error) {
	s := &types.SaleOrders{}
	return s, svc.client.getByIds(types.SaleOrderModel, ids, s)
}

func (svc *SaleOrderService) GetByName(name string) (*types.SaleOrders, error) {
	s := &types.SaleOrders{}
	return s, svc.client.getByName(types.SaleOrderModel, name, s)
}

func (svc *SaleOrderService) GetByField(field string, value string) (*types.SaleOrders, error) {
	s := &types.SaleOrders{}
	return s, svc.client.getByName(types.SaleOrderModel, field, value, s)
}

func (svc *SaleOrderService) GetAll() (*types.SaleOrders, error) {
	s := &types.SaleOrders{}
	return s, svc.client.getAll(types.SaleOrderModel, s)
}

func (svc *SaleOrderService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.SaleOrderModel, fields, relations)
}

func (svc *SaleOrderService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.SaleOrderModel, ids, fields, relations)
}

func (svc *SaleOrderService) Delete(ids []int) error {
	return svc.client.delete(types.SaleOrderModel, ids)
}
