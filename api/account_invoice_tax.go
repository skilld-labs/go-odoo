package api

import (
	"github.com/morezig/go-odoo/types"
)

type AccountInvoiceTaxService struct {
	client *Client
}

func NewAccountInvoiceTaxService(c *Client) *AccountInvoiceTaxService {
	return &AccountInvoiceTaxService{client: c}
}

func (svc *AccountInvoiceTaxService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountInvoiceTaxModel, name)
}

func (svc *AccountInvoiceTaxService) GetByIds(ids []int64) (*types.AccountInvoiceTaxs, error) {
	a := &types.AccountInvoiceTaxs{}
	return a, svc.client.getByIds(types.AccountInvoiceTaxModel, ids, a)
}

func (svc *AccountInvoiceTaxService) GetByName(name string) (*types.AccountInvoiceTaxs, error) {
	a := &types.AccountInvoiceTaxs{}
	return a, svc.client.getByName(types.AccountInvoiceTaxModel, name, a)
}

func (svc *AccountInvoiceTaxService) GetByField(field string, value string) (*types.AccountInvoiceTaxs, error) {
	a := &types.AccountInvoiceTaxs{}
	return a, svc.client.getByField(types.AccountInvoiceTaxModel, field, value, a)
}

func (svc *AccountInvoiceTaxService) GetAll() (*types.AccountInvoiceTaxs, error) {
	a := &types.AccountInvoiceTaxs{}
	return a, svc.client.getAll(types.AccountInvoiceTaxModel, a)
}

func (svc *AccountInvoiceTaxService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountInvoiceTaxModel, fields, relations)
}

func (svc *AccountInvoiceTaxService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountInvoiceTaxModel, ids, fields, relations)
}

func (svc *AccountInvoiceTaxService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountInvoiceTaxModel, ids)
}
