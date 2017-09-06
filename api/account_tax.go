package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountTaxService struct {
	client *Client
}

func NewAccountTaxService(c *Client) *AccountTaxService {
	return &AccountTaxService{client: c}
}

func (svc *AccountTaxService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountTaxModel, name)
}

func (svc *AccountTaxService) GetByIds(ids []int) (*types.AccountTaxs, error) {
	a := &types.AccountTaxs{}
	return a, svc.client.getByIds(types.AccountTaxModel, ids, a)
}

func (svc *AccountTaxService) GetByName(name string) (*types.AccountTaxs, error) {
	a := &types.AccountTaxs{}
	return a, svc.client.getByName(types.AccountTaxModel, name, a)
}

func (svc *AccountTaxService) GetByField(field string, value string) (*types.AccountTaxs, error) {
	a := &types.AccountTaxs{}
	return a, svc.client.getByField(types.AccountTaxModel, field, value, a)
}

func (svc *AccountTaxService) GetAll() (*types.AccountTaxs, error) {
	a := &types.AccountTaxs{}
	return a, svc.client.getAll(types.AccountTaxModel, a)
}

func (svc *AccountTaxService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountTaxModel, fields, relations)
}

func (svc *AccountTaxService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountTaxModel, ids, fields, relations)
}

func (svc *AccountTaxService) Delete(ids []int) error {
	return svc.client.delete(types.AccountTaxModel, ids)
}
