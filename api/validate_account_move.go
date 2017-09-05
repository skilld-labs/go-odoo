package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ValidateAccountMoveService struct {
	client *Client
}

func NewValidateAccountMoveService(c *Client) *ValidateAccountMoveService {
	return &ValidateAccountMoveService{client: c}
}

func (svc *ValidateAccountMoveService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ValidateAccountMoveModel, name)
}

func (svc *ValidateAccountMoveService) GetByIds(ids []int) (*types.ValidateAccountMoves, error) {
	v := &types.ValidateAccountMoves{}
	return v, svc.client.getByIds(types.ValidateAccountMoveModel, ids, v)
}

func (svc *ValidateAccountMoveService) GetByName(name string) (*types.ValidateAccountMoves, error) {
	v := &types.ValidateAccountMoves{}
	return v, svc.client.getByName(types.ValidateAccountMoveModel, name, v)
}

func (svc *ValidateAccountMoveService) GetByField(field string, value string) (*types.ValidateAccountMoves, error) {
	v := &types.ValidateAccountMoves{}
	return v, svc.client.getByName(types.ValidateAccountMoveModel, field, value, v)
}

func (svc *ValidateAccountMoveService) GetAll() (*types.ValidateAccountMoves, error) {
	v := &types.ValidateAccountMoves{}
	return v, svc.client.getAll(types.ValidateAccountMoveModel, v)
}

func (svc *ValidateAccountMoveService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ValidateAccountMoveModel, fields, relations)
}

func (svc *ValidateAccountMoveService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ValidateAccountMoveModel, ids, fields, relations)
}

func (svc *ValidateAccountMoveService) Delete(ids []int) error {
	return svc.client.delete(types.ValidateAccountMoveModel, ids)
}
