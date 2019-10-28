package api

import (
	"github.com/morezig/go-odoo/types"
)

type CashBoxOutService struct {
	client *Client
}

func NewCashBoxOutService(c *Client) *CashBoxOutService {
	return &CashBoxOutService{client: c}
}

func (svc *CashBoxOutService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.CashBoxOutModel, name)
}

func (svc *CashBoxOutService) GetByIds(ids []int64) (*types.CashBoxOuts, error) {
	c := &types.CashBoxOuts{}
	return c, svc.client.getByIds(types.CashBoxOutModel, ids, c)
}

func (svc *CashBoxOutService) GetByName(name string) (*types.CashBoxOuts, error) {
	c := &types.CashBoxOuts{}
	return c, svc.client.getByName(types.CashBoxOutModel, name, c)
}

func (svc *CashBoxOutService) GetByField(field string, value string) (*types.CashBoxOuts, error) {
	c := &types.CashBoxOuts{}
	return c, svc.client.getByField(types.CashBoxOutModel, field, value, c)
}

func (svc *CashBoxOutService) GetAll() (*types.CashBoxOuts, error) {
	c := &types.CashBoxOuts{}
	return c, svc.client.getAll(types.CashBoxOutModel, c)
}

func (svc *CashBoxOutService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.CashBoxOutModel, fields, relations)
}

func (svc *CashBoxOutService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CashBoxOutModel, ids, fields, relations)
}

func (svc *CashBoxOutService) Delete(ids []int64) error {
	return svc.client.delete(types.CashBoxOutModel, ids)
}
