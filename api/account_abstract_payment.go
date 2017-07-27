package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountAbstractPaymentService struct {
	client *Client
}

func NewAccountAbstractPaymentService(c *Client) *AccountAbstractPaymentService {
	return &AccountAbstractPaymentService{client: c}
}

func (svc *AccountAbstractPaymentService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountAbstractPaymentModel, name)
}

func (svc *AccountAbstractPaymentService) GetByIds(ids []int) (*types.AccountAbstractPayments, error) {
	a := &types.AccountAbstractPayments{}
	return a, svc.client.getByIds(types.AccountAbstractPaymentModel, ids, a)
}

func (svc *AccountAbstractPaymentService) GetByName(name string) (*types.AccountAbstractPayments, error) {
	a := &types.AccountAbstractPayments{}
	return a, svc.client.getByName(types.AccountAbstractPaymentModel, name, a)
}

func (svc *AccountAbstractPaymentService) GetAll() (*types.AccountAbstractPayments, error) {
	a := &types.AccountAbstractPayments{}
	return a, svc.client.getAll(types.AccountAbstractPaymentModel, a)
}

func (svc *AccountAbstractPaymentService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountAbstractPaymentModel, fields, relations)
}

func (svc *AccountAbstractPaymentService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountAbstractPaymentModel, ids, fields, relations)
}

func (svc *AccountAbstractPaymentService) Delete(ids []int) error {
	return svc.client.delete(types.AccountAbstractPaymentModel, ids)
}
