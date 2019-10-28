package api

import (
	"github.com/morezig/go-odoo/types"
)

type AccountAnalyticAccountService struct {
	client *Client
}

func NewAccountAnalyticAccountService(c *Client) *AccountAnalyticAccountService {
	return &AccountAnalyticAccountService{client: c}
}

func (svc *AccountAnalyticAccountService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountAnalyticAccountModel, name)
}

func (svc *AccountAnalyticAccountService) GetByIds(ids []int64) (*types.AccountAnalyticAccounts, error) {
	a := &types.AccountAnalyticAccounts{}
	return a, svc.client.getByIds(types.AccountAnalyticAccountModel, ids, a)
}

func (svc *AccountAnalyticAccountService) GetByName(name string) (*types.AccountAnalyticAccounts, error) {
	a := &types.AccountAnalyticAccounts{}
	return a, svc.client.getByName(types.AccountAnalyticAccountModel, name, a)
}

func (svc *AccountAnalyticAccountService) GetByField(field string, value string) (*types.AccountAnalyticAccounts, error) {
	a := &types.AccountAnalyticAccounts{}
	return a, svc.client.getByField(types.AccountAnalyticAccountModel, field, value, a)
}

func (svc *AccountAnalyticAccountService) GetAll() (*types.AccountAnalyticAccounts, error) {
	a := &types.AccountAnalyticAccounts{}
	return a, svc.client.getAll(types.AccountAnalyticAccountModel, a)
}

func (svc *AccountAnalyticAccountService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountAnalyticAccountModel, fields, relations)
}

func (svc *AccountAnalyticAccountService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountAnalyticAccountModel, ids, fields, relations)
}

func (svc *AccountAnalyticAccountService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountAnalyticAccountModel, ids)
}
