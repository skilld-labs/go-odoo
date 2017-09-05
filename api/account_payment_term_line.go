package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountPaymentTermLineService struct {
	client *Client
}

func NewAccountPaymentTermLineService(c *Client) *AccountPaymentTermLineService {
	return &AccountPaymentTermLineService{client: c}
}

func (svc *AccountPaymentTermLineService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountPaymentTermLineModel, name)
}

func (svc *AccountPaymentTermLineService) GetByIds(ids []int) (*types.AccountPaymentTermLines, error) {
	a := &types.AccountPaymentTermLines{}
	return a, svc.client.getByIds(types.AccountPaymentTermLineModel, ids, a)
}

func (svc *AccountPaymentTermLineService) GetByName(name string) (*types.AccountPaymentTermLines, error) {
	a := &types.AccountPaymentTermLines{}
	return a, svc.client.getByName(types.AccountPaymentTermLineModel, name, a)
}

func (svc *AccountPaymentTermLineService) GetByField(field string, value string) (*types.AccountPaymentTermLines, error) {
	a := &types.AccountPaymentTermLines{}
	return a, svc.client.getByName(types.AccountPaymentTermLineModel, field, value, a)
}

func (svc *AccountPaymentTermLineService) GetAll() (*types.AccountPaymentTermLines, error) {
	a := &types.AccountPaymentTermLines{}
	return a, svc.client.getAll(types.AccountPaymentTermLineModel, a)
}

func (svc *AccountPaymentTermLineService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountPaymentTermLineModel, fields, relations)
}

func (svc *AccountPaymentTermLineService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountPaymentTermLineModel, ids, fields, relations)
}

func (svc *AccountPaymentTermLineService) Delete(ids []int) error {
	return svc.client.delete(types.AccountPaymentTermLineModel, ids)
}
