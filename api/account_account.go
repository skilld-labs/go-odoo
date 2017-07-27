package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountAccountService struct {
	client *Client
}

func NewAccountAccountService(c *Client) *AccountAccountService {
	return &AccountAccountService{client: c}
}

func (svc *AccountAccountService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountAccountModel, name)
}

func (svc *AccountAccountService) GetByIds(ids []int) (*types.AccountAccounts, error) {
	a := &types.AccountAccounts{}
	return a, svc.client.getByIds(types.AccountAccountModel, ids, a)
}

func (svc *AccountAccountService) GetByName(name string) (*types.AccountAccounts, error) {
	a := &types.AccountAccounts{}
	return a, svc.client.getByName(types.AccountAccountModel, name, a)
}

func (svc *AccountAccountService) GetAll() (*types.AccountAccounts, error) {
	a := &types.AccountAccounts{}
	return a, svc.client.getAll(types.AccountAccountModel, a)
}

func (svc *AccountAccountService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountAccountModel, fields, relations)
}

func (svc *AccountAccountService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountAccountModel, ids, fields, relations)
}

func (svc *AccountAccountService) Delete(ids []int) error {
	return svc.client.delete(types.AccountAccountModel, ids)
}
