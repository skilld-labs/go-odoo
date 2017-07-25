package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountPaymentTermService struct {
	client *Client
}

func NewAccountPaymentTermService(c *Client) *AccountPaymentTermService {
	return &AccountPaymentTermService{client: c}
}

func (svc *AccountPaymentTermService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountPaymentTermModel, name)
}

func (svc *AccountPaymentTermService) GetByIds(ids []int) (*types.AccountPaymentTerms, error) {
	a := &types.AccountPaymentTerms{}
	return a, svc.client.getByIds(types.AccountPaymentTermModel, ids, a)
}

func (svc *AccountPaymentTermService) GetByName(name string) (*types.AccountPaymentTerms, error) {
	a := &types.AccountPaymentTerms{}
	return a, svc.client.getByName(types.AccountPaymentTermModel, name, a)
}

func (svc *AccountPaymentTermService) GetAll() (*types.AccountPaymentTerms, error) {
	a := &types.AccountPaymentTerms{}
	return a, svc.client.getAll(types.AccountPaymentTermModel, a)
}

func (svc *AccountPaymentTermService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountPaymentTermModel, fields, relations)
}

func (svc *AccountPaymentTermService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountPaymentTermModel, ids, fields, relations)
}

func (svc *AccountPaymentTermService) Delete(ids []int) error {
	return svc.client.delete(types.AccountPaymentTermModel, ids)
}
