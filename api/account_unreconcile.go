package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountUnreconcileService struct {
	client *Client
}

func NewAccountUnreconcileService(c *Client) *AccountUnreconcileService {
	return &AccountUnreconcileService{client: c}
}

func (svc *AccountUnreconcileService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountUnreconcileModel, name)
}

func (svc *AccountUnreconcileService) GetByIds(ids []int64) (*types.AccountUnreconciles, error) {
	a := &types.AccountUnreconciles{}
	return a, svc.client.getByIds(types.AccountUnreconcileModel, ids, a)
}

func (svc *AccountUnreconcileService) GetByName(name string) (*types.AccountUnreconciles, error) {
	a := &types.AccountUnreconciles{}
	return a, svc.client.getByName(types.AccountUnreconcileModel, name, a)
}

func (svc *AccountUnreconcileService) GetByField(field string, value string) (*types.AccountUnreconciles, error) {
	a := &types.AccountUnreconciles{}
	return a, svc.client.getByField(types.AccountUnreconcileModel, field, value, a)
}

func (svc *AccountUnreconcileService) GetAll() (*types.AccountUnreconciles, error) {
	a := &types.AccountUnreconciles{}
	return a, svc.client.getAll(types.AccountUnreconcileModel, a)
}

func (svc *AccountUnreconcileService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountUnreconcileModel, fields, relations)
}

func (svc *AccountUnreconcileService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountUnreconcileModel, ids, fields, relations)
}

func (svc *AccountUnreconcileService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountUnreconcileModel, ids)
}
