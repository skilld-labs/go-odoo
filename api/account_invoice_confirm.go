package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountInvoiceConfirmService struct {
	client *Client
}

func NewAccountInvoiceConfirmService(c *Client) *AccountInvoiceConfirmService {
	return &AccountInvoiceConfirmService{client: c}
}

func (svc *AccountInvoiceConfirmService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountInvoiceConfirmModel, name)
}

func (svc *AccountInvoiceConfirmService) GetByIds(ids []int64) (*types.AccountInvoiceConfirms, error) {
	a := &types.AccountInvoiceConfirms{}
	return a, svc.client.getByIds(types.AccountInvoiceConfirmModel, ids, a)
}

func (svc *AccountInvoiceConfirmService) GetByName(name string) (*types.AccountInvoiceConfirms, error) {
	a := &types.AccountInvoiceConfirms{}
	return a, svc.client.getByName(types.AccountInvoiceConfirmModel, name, a)
}

func (svc *AccountInvoiceConfirmService) GetByField(field string, value string) (*types.AccountInvoiceConfirms, error) {
	a := &types.AccountInvoiceConfirms{}
	return a, svc.client.getByField(types.AccountInvoiceConfirmModel, field, value, a)
}

func (svc *AccountInvoiceConfirmService) GetAll() (*types.AccountInvoiceConfirms, error) {
	a := &types.AccountInvoiceConfirms{}
	return a, svc.client.getAll(types.AccountInvoiceConfirmModel, a)
}

func (svc *AccountInvoiceConfirmService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountInvoiceConfirmModel, fields, relations)
}

func (svc *AccountInvoiceConfirmService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountInvoiceConfirmModel, ids, fields, relations)
}

func (svc *AccountInvoiceConfirmService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountInvoiceConfirmModel, ids)
}
