package api

import (
	"github.com/morezig/go-odoo/types"
)

type AccountInvoiceLineService struct {
	client *Client
}

func NewAccountInvoiceLineService(c *Client) *AccountInvoiceLineService {
	return &AccountInvoiceLineService{client: c}
}

func (svc *AccountInvoiceLineService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountInvoiceLineModel, name)
}

func (svc *AccountInvoiceLineService) GetByIds(ids []int64) (*types.AccountInvoiceLines, error) {
	a := &types.AccountInvoiceLines{}
	return a, svc.client.getByIds(types.AccountInvoiceLineModel, ids, a)
}

func (svc *AccountInvoiceLineService) GetByName(name string) (*types.AccountInvoiceLines, error) {
	a := &types.AccountInvoiceLines{}
	return a, svc.client.getByName(types.AccountInvoiceLineModel, name, a)
}

func (svc *AccountInvoiceLineService) GetByField(field string, value string) (*types.AccountInvoiceLines, error) {
	a := &types.AccountInvoiceLines{}
	return a, svc.client.getByField(types.AccountInvoiceLineModel, field, value, a)
}

func (svc *AccountInvoiceLineService) GetAll() (*types.AccountInvoiceLines, error) {
	a := &types.AccountInvoiceLines{}
	return a, svc.client.getAll(types.AccountInvoiceLineModel, a)
}

func (svc *AccountInvoiceLineService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountInvoiceLineModel, fields, relations)
}

func (svc *AccountInvoiceLineService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountInvoiceLineModel, ids, fields, relations)
}

func (svc *AccountInvoiceLineService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountInvoiceLineModel, ids)
}
