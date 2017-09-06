package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountPaymentMethodService struct {
	client *Client
}

func NewAccountPaymentMethodService(c *Client) *AccountPaymentMethodService {
	return &AccountPaymentMethodService{client: c}
}

func (svc *AccountPaymentMethodService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountPaymentMethodModel, name)
}

func (svc *AccountPaymentMethodService) GetByIds(ids []int) (*types.AccountPaymentMethods, error) {
	a := &types.AccountPaymentMethods{}
	return a, svc.client.getByIds(types.AccountPaymentMethodModel, ids, a)
}

func (svc *AccountPaymentMethodService) GetByName(name string) (*types.AccountPaymentMethods, error) {
	a := &types.AccountPaymentMethods{}
	return a, svc.client.getByName(types.AccountPaymentMethodModel, name, a)
}

func (svc *AccountPaymentMethodService) GetByField(field string, value string) (*types.AccountPaymentMethods, error) {
	a := &types.AccountPaymentMethods{}
	return a, svc.client.getByField(types.AccountPaymentMethodModel, field, value, a)
}

func (svc *AccountPaymentMethodService) GetAll() (*types.AccountPaymentMethods, error) {
	a := &types.AccountPaymentMethods{}
	return a, svc.client.getAll(types.AccountPaymentMethodModel, a)
}

func (svc *AccountPaymentMethodService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountPaymentMethodModel, fields, relations)
}

func (svc *AccountPaymentMethodService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountPaymentMethodModel, ids, fields, relations)
}

func (svc *AccountPaymentMethodService) Delete(ids []int) error {
	return svc.client.delete(types.AccountPaymentMethodModel, ids)
}
