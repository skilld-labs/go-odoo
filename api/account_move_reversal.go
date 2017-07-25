package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountMoveReversalService struct {
	client *Client
}

func NewAccountMoveReversalService(c *Client) *AccountMoveReversalService {
	return &AccountMoveReversalService{client: c}
}

func (svc *AccountMoveReversalService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountMoveReversalModel, name)
}

func (svc *AccountMoveReversalService) GetByIds(ids []int) (*types.AccountMoveReversals, error) {
	a := &types.AccountMoveReversals{}
	return a, svc.client.getByIds(types.AccountMoveReversalModel, ids, a)
}

func (svc *AccountMoveReversalService) GetByName(name string) (*types.AccountMoveReversals, error) {
	a := &types.AccountMoveReversals{}
	return a, svc.client.getByName(types.AccountMoveReversalModel, name, a)
}

func (svc *AccountMoveReversalService) GetAll() (*types.AccountMoveReversals, error) {
	a := &types.AccountMoveReversals{}
	return a, svc.client.getAll(types.AccountMoveReversalModel, a)
}

func (svc *AccountMoveReversalService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountMoveReversalModel, fields, relations)
}

func (svc *AccountMoveReversalService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountMoveReversalModel, ids, fields, relations)
}

func (svc *AccountMoveReversalService) Delete(ids []int) error {
	return svc.client.delete(types.AccountMoveReversalModel, ids)
}
