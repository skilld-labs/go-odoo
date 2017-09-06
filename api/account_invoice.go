package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountInvoiceService struct {
	client *Client
}

func NewAccountInvoiceService(c *Client) *AccountInvoiceService {
	return &AccountInvoiceService{client: c}
}

func (svc *AccountInvoiceService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountInvoiceModel, name)
}

func (svc *AccountInvoiceService) GetByIds(ids []int) (*types.AccountInvoices, error) {
	a := &types.AccountInvoices{}
	return a, svc.client.getByIds(types.AccountInvoiceModel, ids, a)
}

func (svc *AccountInvoiceService) GetByName(name string) (*types.AccountInvoices, error) {
	a := &types.AccountInvoices{}
	return a, svc.client.getByName(types.AccountInvoiceModel, name, a)
}

func (svc *AccountInvoiceService) GetByField(field string, value string) (*types.AccountInvoices, error) {
	a := &types.AccountInvoices{}
	return a, svc.client.getByField(types.AccountInvoiceModel, field, value, a)
}

func (svc *AccountInvoiceService) GetAll() (*types.AccountInvoices, error) {
	a := &types.AccountInvoices{}
	return a, svc.client.getAll(types.AccountInvoiceModel, a)
}

func (svc *AccountInvoiceService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountInvoiceModel, fields, relations)
}

func (svc *AccountInvoiceService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountInvoiceModel, ids, fields, relations)
}

func (svc *AccountInvoiceService) Delete(ids []int) error {
	return svc.client.delete(types.AccountInvoiceModel, ids)
}
