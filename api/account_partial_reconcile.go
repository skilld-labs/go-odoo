package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountPartialReconcileService struct {
	client *Client
}

func NewAccountPartialReconcileService(c *Client) *AccountPartialReconcileService {
	return &AccountPartialReconcileService{client: c}
}

func (svc *AccountPartialReconcileService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountPartialReconcileModel, name)
}

func (svc *AccountPartialReconcileService) GetByIds(ids []int) (*types.AccountPartialReconciles, error) {
	a := &types.AccountPartialReconciles{}
	return a, svc.client.getByIds(types.AccountPartialReconcileModel, ids, a)
}

func (svc *AccountPartialReconcileService) GetByName(name string) (*types.AccountPartialReconciles, error) {
	a := &types.AccountPartialReconciles{}
	return a, svc.client.getByName(types.AccountPartialReconcileModel, name, a)
}

func (svc *AccountPartialReconcileService) GetByField(field string, value string) (*types.AccountPartialReconciles, error) {
	a := &types.AccountPartialReconciles{}
	return a, svc.client.getByField(types.AccountPartialReconcileModel, field, value, a)
}

func (svc *AccountPartialReconcileService) GetAll() (*types.AccountPartialReconciles, error) {
	a := &types.AccountPartialReconciles{}
	return a, svc.client.getAll(types.AccountPartialReconcileModel, a)
}

func (svc *AccountPartialReconcileService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountPartialReconcileModel, fields, relations)
}

func (svc *AccountPartialReconcileService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountPartialReconcileModel, ids, fields, relations)
}

func (svc *AccountPartialReconcileService) Delete(ids []int) error {
	return svc.client.delete(types.AccountPartialReconcileModel, ids)
}
