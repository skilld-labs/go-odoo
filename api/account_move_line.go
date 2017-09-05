package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountMoveLineService struct {
	client *Client
}

func NewAccountMoveLineService(c *Client) *AccountMoveLineService {
	return &AccountMoveLineService{client: c}
}

func (svc *AccountMoveLineService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountMoveLineModel, name)
}

func (svc *AccountMoveLineService) GetByIds(ids []int) (*types.AccountMoveLines, error) {
	a := &types.AccountMoveLines{}
	return a, svc.client.getByIds(types.AccountMoveLineModel, ids, a)
}

func (svc *AccountMoveLineService) GetByName(name string) (*types.AccountMoveLines, error) {
	a := &types.AccountMoveLines{}
	return a, svc.client.getByName(types.AccountMoveLineModel, name, a)
}

func (svc *AccountMoveLineService) GetByField(field string, value string) (*types.AccountMoveLines, error) {
	a := &types.AccountMoveLines{}
	return a, svc.client.getByName(types.AccountMoveLineModel, field, value, a)
}

func (svc *AccountMoveLineService) GetAll() (*types.AccountMoveLines, error) {
	a := &types.AccountMoveLines{}
	return a, svc.client.getAll(types.AccountMoveLineModel, a)
}

func (svc *AccountMoveLineService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountMoveLineModel, fields, relations)
}

func (svc *AccountMoveLineService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountMoveLineModel, ids, fields, relations)
}

func (svc *AccountMoveLineService) Delete(ids []int) error {
	return svc.client.delete(types.AccountMoveLineModel, ids)
}
