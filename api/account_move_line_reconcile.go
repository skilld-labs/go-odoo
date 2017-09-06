package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountMoveLineReconcileService struct {
	client *Client
}

func NewAccountMoveLineReconcileService(c *Client) *AccountMoveLineReconcileService {
	return &AccountMoveLineReconcileService{client: c}
}

func (svc *AccountMoveLineReconcileService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountMoveLineReconcileModel, name)
}

func (svc *AccountMoveLineReconcileService) GetByIds(ids []int) (*types.AccountMoveLineReconciles, error) {
	a := &types.AccountMoveLineReconciles{}
	return a, svc.client.getByIds(types.AccountMoveLineReconcileModel, ids, a)
}

func (svc *AccountMoveLineReconcileService) GetByName(name string) (*types.AccountMoveLineReconciles, error) {
	a := &types.AccountMoveLineReconciles{}
	return a, svc.client.getByName(types.AccountMoveLineReconcileModel, name, a)
}

func (svc *AccountMoveLineReconcileService) GetByField(field string, value string) (*types.AccountMoveLineReconciles, error) {
	a := &types.AccountMoveLineReconciles{}
	return a, svc.client.getByField(types.AccountMoveLineReconcileModel, field, value, a)
}

func (svc *AccountMoveLineReconcileService) GetAll() (*types.AccountMoveLineReconciles, error) {
	a := &types.AccountMoveLineReconciles{}
	return a, svc.client.getAll(types.AccountMoveLineReconcileModel, a)
}

func (svc *AccountMoveLineReconcileService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountMoveLineReconcileModel, fields, relations)
}

func (svc *AccountMoveLineReconcileService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountMoveLineReconcileModel, ids, fields, relations)
}

func (svc *AccountMoveLineReconcileService) Delete(ids []int) error {
	return svc.client.delete(types.AccountMoveLineReconcileModel, ids)
}
