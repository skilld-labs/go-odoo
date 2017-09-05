package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountInvoiceRefundService struct {
	client *Client
}

func NewAccountInvoiceRefundService(c *Client) *AccountInvoiceRefundService {
	return &AccountInvoiceRefundService{client: c}
}

func (svc *AccountInvoiceRefundService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountInvoiceRefundModel, name)
}

func (svc *AccountInvoiceRefundService) GetByIds(ids []int) (*types.AccountInvoiceRefunds, error) {
	a := &types.AccountInvoiceRefunds{}
	return a, svc.client.getByIds(types.AccountInvoiceRefundModel, ids, a)
}

func (svc *AccountInvoiceRefundService) GetByName(name string) (*types.AccountInvoiceRefunds, error) {
	a := &types.AccountInvoiceRefunds{}
	return a, svc.client.getByName(types.AccountInvoiceRefundModel, name, a)
}

func (svc *AccountInvoiceRefundService) GetByField(field string, value string) (*types.AccountInvoiceRefunds, error) {
	a := &types.AccountInvoiceRefunds{}
	return a, svc.client.getByName(types.AccountInvoiceRefundModel, field, value, a)
}

func (svc *AccountInvoiceRefundService) GetAll() (*types.AccountInvoiceRefunds, error) {
	a := &types.AccountInvoiceRefunds{}
	return a, svc.client.getAll(types.AccountInvoiceRefundModel, a)
}

func (svc *AccountInvoiceRefundService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountInvoiceRefundModel, fields, relations)
}

func (svc *AccountInvoiceRefundService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountInvoiceRefundModel, ids, fields, relations)
}

func (svc *AccountInvoiceRefundService) Delete(ids []int) error {
	return svc.client.delete(types.AccountInvoiceRefundModel, ids)
}
