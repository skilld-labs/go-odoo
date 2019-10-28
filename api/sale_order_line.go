package api

import (
	"github.com/morezig/go-odoo/types"
)

type SaleOrderLineService struct {
	client *Client
}

func NewSaleOrderLineService(c *Client) *SaleOrderLineService {
	return &SaleOrderLineService{client: c}
}

func (svc *SaleOrderLineService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.SaleOrderLineModel, name)
}

func (svc *SaleOrderLineService) GetByIds(ids []int64) (*types.SaleOrderLines, error) {
	s := &types.SaleOrderLines{}
	return s, svc.client.getByIds(types.SaleOrderLineModel, ids, s)
}

func (svc *SaleOrderLineService) GetByName(name string) (*types.SaleOrderLines, error) {
	s := &types.SaleOrderLines{}
	return s, svc.client.getByName(types.SaleOrderLineModel, name, s)
}

func (svc *SaleOrderLineService) GetByField(field string, value string) (*types.SaleOrderLines, error) {
	s := &types.SaleOrderLines{}
	return s, svc.client.getByField(types.SaleOrderLineModel, field, value, s)
}

func (svc *SaleOrderLineService) GetAll() (*types.SaleOrderLines, error) {
	s := &types.SaleOrderLines{}
	return s, svc.client.getAll(types.SaleOrderLineModel, s)
}

func (svc *SaleOrderLineService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.SaleOrderLineModel, fields, relations)
}

func (svc *SaleOrderLineService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.SaleOrderLineModel, ids, fields, relations)
}

func (svc *SaleOrderLineService) Delete(ids []int64) error {
	return svc.client.delete(types.SaleOrderLineModel, ids)
}
