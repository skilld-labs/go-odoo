package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountFiscalPositionAccountService struct {
	client *Client
}

func NewAccountFiscalPositionAccountService(c *Client) *AccountFiscalPositionAccountService {
	return &AccountFiscalPositionAccountService{client: c}
}

func (svc *AccountFiscalPositionAccountService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountFiscalPositionAccountModel, name)
}

func (svc *AccountFiscalPositionAccountService) GetByIds(ids []int) (*types.AccountFiscalPositionAccounts, error) {
	a := &types.AccountFiscalPositionAccounts{}
	return a, svc.client.getByIds(types.AccountFiscalPositionAccountModel, ids, a)
}

func (svc *AccountFiscalPositionAccountService) GetByName(name string) (*types.AccountFiscalPositionAccounts, error) {
	a := &types.AccountFiscalPositionAccounts{}
	return a, svc.client.getByName(types.AccountFiscalPositionAccountModel, name, a)
}

func (svc *AccountFiscalPositionAccountService) GetAll() (*types.AccountFiscalPositionAccounts, error) {
	a := &types.AccountFiscalPositionAccounts{}
	return a, svc.client.getAll(types.AccountFiscalPositionAccountModel, a)
}

func (svc *AccountFiscalPositionAccountService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountFiscalPositionAccountModel, fields, relations)
}

func (svc *AccountFiscalPositionAccountService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountFiscalPositionAccountModel, ids, fields, relations)
}

func (svc *AccountFiscalPositionAccountService) Delete(ids []int) error {
	return svc.client.delete(types.AccountFiscalPositionAccountModel, ids)
}
