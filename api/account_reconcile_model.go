package api

import (
	"github.com/morezig/go-odoo/types"
)

type AccountReconcileModelService struct {
	client *Client
}

func NewAccountReconcileModelService(c *Client) *AccountReconcileModelService {
	return &AccountReconcileModelService{client: c}
}

func (svc *AccountReconcileModelService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountReconcileModelModel, name)
}

func (svc *AccountReconcileModelService) GetByIds(ids []int64) (*types.AccountReconcileModels, error) {
	a := &types.AccountReconcileModels{}
	return a, svc.client.getByIds(types.AccountReconcileModelModel, ids, a)
}

func (svc *AccountReconcileModelService) GetByName(name string) (*types.AccountReconcileModels, error) {
	a := &types.AccountReconcileModels{}
	return a, svc.client.getByName(types.AccountReconcileModelModel, name, a)
}

func (svc *AccountReconcileModelService) GetByField(field string, value string) (*types.AccountReconcileModels, error) {
	a := &types.AccountReconcileModels{}
	return a, svc.client.getByField(types.AccountReconcileModelModel, field, value, a)
}

func (svc *AccountReconcileModelService) GetAll() (*types.AccountReconcileModels, error) {
	a := &types.AccountReconcileModels{}
	return a, svc.client.getAll(types.AccountReconcileModelModel, a)
}

func (svc *AccountReconcileModelService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountReconcileModelModel, fields, relations)
}

func (svc *AccountReconcileModelService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountReconcileModelModel, ids, fields, relations)
}

func (svc *AccountReconcileModelService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountReconcileModelModel, ids)
}
