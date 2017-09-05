package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type PurchaseOrderLineService struct {
	client *Client
}

func NewPurchaseOrderLineService(c *Client) *PurchaseOrderLineService {
	return &PurchaseOrderLineService{client: c}
}

func (svc *PurchaseOrderLineService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.PurchaseOrderLineModel, name)
}

func (svc *PurchaseOrderLineService) GetByIds(ids []int) (*types.PurchaseOrderLines, error) {
	p := &types.PurchaseOrderLines{}
	return p, svc.client.getByIds(types.PurchaseOrderLineModel, ids, p)
}

func (svc *PurchaseOrderLineService) GetByName(name string) (*types.PurchaseOrderLines, error) {
	p := &types.PurchaseOrderLines{}
	return p, svc.client.getByName(types.PurchaseOrderLineModel, name, p)
}

func (svc *PurchaseOrderLineService) GetByField(field string, value string) (*types.PurchaseOrderLines, error) {
	p := &types.PurchaseOrderLines{}
	return p, svc.client.getByName(types.PurchaseOrderLineModel, field, value, p)
}

func (svc *PurchaseOrderLineService) GetAll() (*types.PurchaseOrderLines, error) {
	p := &types.PurchaseOrderLines{}
	return p, svc.client.getAll(types.PurchaseOrderLineModel, p)
}

func (svc *PurchaseOrderLineService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.PurchaseOrderLineModel, fields, relations)
}

func (svc *PurchaseOrderLineService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.PurchaseOrderLineModel, ids, fields, relations)
}

func (svc *PurchaseOrderLineService) Delete(ids []int) error {
	return svc.client.delete(types.PurchaseOrderLineModel, ids)
}
