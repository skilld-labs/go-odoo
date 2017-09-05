package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountAgedTrialBalanceService struct {
	client *Client
}

func NewAccountAgedTrialBalanceService(c *Client) *AccountAgedTrialBalanceService {
	return &AccountAgedTrialBalanceService{client: c}
}

func (svc *AccountAgedTrialBalanceService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountAgedTrialBalanceModel, name)
}

func (svc *AccountAgedTrialBalanceService) GetByIds(ids []int) (*types.AccountAgedTrialBalances, error) {
	a := &types.AccountAgedTrialBalances{}
	return a, svc.client.getByIds(types.AccountAgedTrialBalanceModel, ids, a)
}

func (svc *AccountAgedTrialBalanceService) GetByName(name string) (*types.AccountAgedTrialBalances, error) {
	a := &types.AccountAgedTrialBalances{}
	return a, svc.client.getByName(types.AccountAgedTrialBalanceModel, name, a)
}

func (svc *AccountAgedTrialBalanceService) GetByField(field string, value string) (*types.AccountAgedTrialBalances, error) {
	a := &types.AccountAgedTrialBalances{}
	return a, svc.client.getByField(types.AccountAgedTrialBalanceModel, field, value, a)
}

func (svc *AccountAgedTrialBalanceService) GetAll() (*types.AccountAgedTrialBalances, error) {
	a := &types.AccountAgedTrialBalances{}
	return a, svc.client.getAll(types.AccountAgedTrialBalanceModel, a)
}

func (svc *AccountAgedTrialBalanceService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountAgedTrialBalanceModel, fields, relations)
}

func (svc *AccountAgedTrialBalanceService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountAgedTrialBalanceModel, ids, fields, relations)
}

func (svc *AccountAgedTrialBalanceService) Delete(ids []int) error {
	return svc.client.delete(types.AccountAgedTrialBalanceModel, ids)
}
