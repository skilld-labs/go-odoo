package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountFullReconcileService struct {
	client *Client
}

func NewAccountFullReconcileService(c *Client) *AccountFullReconcileService {
	return &AccountFullReconcileService{client: c}
}

func (svc *AccountFullReconcileService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountFullReconcileModel, name)
}

func (svc *AccountFullReconcileService) GetByIds(ids []int) (*types.AccountFullReconciles, error) {
	a := &types.AccountFullReconciles{}
	return a, svc.client.getByIds(types.AccountFullReconcileModel, ids, a)
}

func (svc *AccountFullReconcileService) GetByName(name string) (*types.AccountFullReconciles, error) {
	a := &types.AccountFullReconciles{}
	return a, svc.client.getByName(types.AccountFullReconcileModel, name, a)
}

func (svc *AccountFullReconcileService) GetByField(field string, value string) (*types.AccountFullReconciles, error) {
	a := &types.AccountFullReconciles{}
	return a, svc.client.getByField(types.AccountFullReconcileModel, field, value, a)
}

func (svc *AccountFullReconcileService) GetAll() (*types.AccountFullReconciles, error) {
	a := &types.AccountFullReconciles{}
	return a, svc.client.getAll(types.AccountFullReconcileModel, a)
}

func (svc *AccountFullReconcileService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountFullReconcileModel, fields, relations)
}

func (svc *AccountFullReconcileService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountFullReconcileModel, ids, fields, relations)
}

func (svc *AccountFullReconcileService) Delete(ids []int) error {
	return svc.client.delete(types.AccountFullReconcileModel, ids)
}
