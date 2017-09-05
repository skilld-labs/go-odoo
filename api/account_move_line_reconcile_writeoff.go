package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountMoveLineReconcileWriteoffService struct {
	client *Client
}

func NewAccountMoveLineReconcileWriteoffService(c *Client) *AccountMoveLineReconcileWriteoffService {
	return &AccountMoveLineReconcileWriteoffService{client: c}
}

func (svc *AccountMoveLineReconcileWriteoffService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountMoveLineReconcileWriteoffModel, name)
}

func (svc *AccountMoveLineReconcileWriteoffService) GetByIds(ids []int) (*types.AccountMoveLineReconcileWriteoffs, error) {
	a := &types.AccountMoveLineReconcileWriteoffs{}
	return a, svc.client.getByIds(types.AccountMoveLineReconcileWriteoffModel, ids, a)
}

func (svc *AccountMoveLineReconcileWriteoffService) GetByName(name string) (*types.AccountMoveLineReconcileWriteoffs, error) {
	a := &types.AccountMoveLineReconcileWriteoffs{}
	return a, svc.client.getByName(types.AccountMoveLineReconcileWriteoffModel, name, a)
}

func (svc *AccountMoveLineReconcileWriteoffService) GetByField(field string, value string) (*types.AccountMoveLineReconcileWriteoffs, error) {
	a := &types.AccountMoveLineReconcileWriteoffs{}
	return a, svc.client.getByField(types.AccountMoveLineReconcileWriteoffModel, field, value, a)
}

func (svc *AccountMoveLineReconcileWriteoffService) GetAll() (*types.AccountMoveLineReconcileWriteoffs, error) {
	a := &types.AccountMoveLineReconcileWriteoffs{}
	return a, svc.client.getAll(types.AccountMoveLineReconcileWriteoffModel, a)
}

func (svc *AccountMoveLineReconcileWriteoffService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountMoveLineReconcileWriteoffModel, fields, relations)
}

func (svc *AccountMoveLineReconcileWriteoffService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountMoveLineReconcileWriteoffModel, ids, fields, relations)
}

func (svc *AccountMoveLineReconcileWriteoffService) Delete(ids []int) error {
	return svc.client.delete(types.AccountMoveLineReconcileWriteoffModel, ids)
}
