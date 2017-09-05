package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type CashBoxInService struct {
	client *Client
}

func NewCashBoxInService(c *Client) *CashBoxInService {
	return &CashBoxInService{client: c}
}

func (svc *CashBoxInService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.CashBoxInModel, name)
}

func (svc *CashBoxInService) GetByIds(ids []int) (*types.CashBoxIns, error) {
	c := &types.CashBoxIns{}
	return c, svc.client.getByIds(types.CashBoxInModel, ids, c)
}

func (svc *CashBoxInService) GetByName(name string) (*types.CashBoxIns, error) {
	c := &types.CashBoxIns{}
	return c, svc.client.getByName(types.CashBoxInModel, name, c)
}

func (svc *CashBoxInService) GetByField(field string, value string) (*types.CashBoxIns, error) {
	c := &types.CashBoxIns{}
	return c, svc.client.getByName(types.CashBoxInModel, field, value, c)
}

func (svc *CashBoxInService) GetAll() (*types.CashBoxIns, error) {
	c := &types.CashBoxIns{}
	return c, svc.client.getAll(types.CashBoxInModel, c)
}

func (svc *CashBoxInService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.CashBoxInModel, fields, relations)
}

func (svc *CashBoxInService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CashBoxInModel, ids, fields, relations)
}

func (svc *CashBoxInService) Delete(ids []int) error {
	return svc.client.delete(types.CashBoxInModel, ids)
}
