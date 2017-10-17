package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountGroupService struct {
	client *Client
}

func NewAccountGroupService(c *Client) *AccountGroupService {
	return &AccountGroupService{client: c}
}

func (svc *AccountGroupService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountGroupModel, name)
}

func (svc *AccountGroupService) GetByIds(ids []int) (*types.AccountGroups, error) {
	a := &types.AccountGroups{}
	return a, svc.client.getByIds(types.AccountGroupModel, ids, a)
}

func (svc *AccountGroupService) GetByName(name string) (*types.AccountGroups, error) {
	a := &types.AccountGroups{}
	return a, svc.client.getByName(types.AccountGroupModel, name, a)
}

func (svc *AccountGroupService) GetByField(field string, value string) (*types.AccountGroups, error) {
	a := &types.AccountGroups{}
	return a, svc.client.getByField(types.AccountGroupModel, field, value, a)
}

func (svc *AccountGroupService) GetAll() (*types.AccountGroups, error) {
	a := &types.AccountGroups{}
	return a, svc.client.getAll(types.AccountGroupModel, a)
}

func (svc *AccountGroupService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountGroupModel, fields, relations)
}

func (svc *AccountGroupService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountGroupModel, ids, fields, relations)
}

func (svc *AccountGroupService) Delete(ids []int) error {
	return svc.client.delete(types.AccountGroupModel, ids)
}
