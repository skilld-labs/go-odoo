package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountInvoiceCancelService struct {
	client *Client
}

func NewAccountInvoiceCancelService(c *Client) *AccountInvoiceCancelService {
	return &AccountInvoiceCancelService{client: c}
}

func (svc *AccountInvoiceCancelService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountInvoiceCancelModel, name)
}

func (svc *AccountInvoiceCancelService) GetByIds(ids []int) (*types.AccountInvoiceCancels, error) {
	a := &types.AccountInvoiceCancels{}
	return a, svc.client.getByIds(types.AccountInvoiceCancelModel, ids, a)
}

func (svc *AccountInvoiceCancelService) GetByName(name string) (*types.AccountInvoiceCancels, error) {
	a := &types.AccountInvoiceCancels{}
	return a, svc.client.getByName(types.AccountInvoiceCancelModel, name, a)
}

func (svc *AccountInvoiceCancelService) GetByField(field string, value string) (*types.AccountInvoiceCancels, error) {
	a := &types.AccountInvoiceCancels{}
	return a, svc.client.getByField(types.AccountInvoiceCancelModel, field, value, a)
}

func (svc *AccountInvoiceCancelService) GetAll() (*types.AccountInvoiceCancels, error) {
	a := &types.AccountInvoiceCancels{}
	return a, svc.client.getAll(types.AccountInvoiceCancelModel, a)
}

func (svc *AccountInvoiceCancelService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountInvoiceCancelModel, fields, relations)
}

func (svc *AccountInvoiceCancelService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountInvoiceCancelModel, ids, fields, relations)
}

func (svc *AccountInvoiceCancelService) Delete(ids []int) error {
	return svc.client.delete(types.AccountInvoiceCancelModel, ids)
}
