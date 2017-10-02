package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountMoveService struct {
	client *Client
}

func NewAccountMoveService(c *Client) *AccountMoveService {
	return &AccountMoveService{client: c}
}

func (svc *AccountMoveService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountMoveModel, name)
}

func (svc *AccountMoveService) GetByIds(ids []int64) (*types.AccountMoves, error) {
	a := &types.AccountMoves{}
	return a, svc.client.getByIds(types.AccountMoveModel, ids, a)
}

func (svc *AccountMoveService) GetByName(name string) (*types.AccountMoves, error) {
	a := &types.AccountMoves{}
	return a, svc.client.getByName(types.AccountMoveModel, name, a)
}

func (svc *AccountMoveService) GetByField(field string, value string) (*types.AccountMoves, error) {
	a := &types.AccountMoves{}
	return a, svc.client.getByField(types.AccountMoveModel, field, value, a)
}

func (svc *AccountMoveService) GetAll() (*types.AccountMoves, error) {
	a := &types.AccountMoves{}
	return a, svc.client.getAll(types.AccountMoveModel, a)
}

func (svc *AccountMoveService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountMoveModel, fields, relations)
}

func (svc *AccountMoveService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountMoveModel, ids, fields, relations)
}

func (svc *AccountMoveService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountMoveModel, ids)
}
