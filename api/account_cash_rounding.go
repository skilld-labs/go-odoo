package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountCashRoundingService struct {
	client *Client
}

func NewAccountCashRoundingService(c *Client) *AccountCashRoundingService {
	return &AccountCashRoundingService{client: c}
}

func (svc *AccountCashRoundingService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountCashRoundingModel, name)
}

func (svc *AccountCashRoundingService) GetByIds(ids []int) (*types.AccountCashRoundings, error) {
	a := &types.AccountCashRoundings{}
	return a, svc.client.getByIds(types.AccountCashRoundingModel, ids, a)
}

func (svc *AccountCashRoundingService) GetByName(name string) (*types.AccountCashRoundings, error) {
	a := &types.AccountCashRoundings{}
	return a, svc.client.getByName(types.AccountCashRoundingModel, name, a)
}

func (svc *AccountCashRoundingService) GetByField(field string, value string) (*types.AccountCashRoundings, error) {
	a := &types.AccountCashRoundings{}
	return a, svc.client.getByField(types.AccountCashRoundingModel, field, value, a)
}

func (svc *AccountCashRoundingService) GetAll() (*types.AccountCashRoundings, error) {
	a := &types.AccountCashRoundings{}
	return a, svc.client.getAll(types.AccountCashRoundingModel, a)
}

func (svc *AccountCashRoundingService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountCashRoundingModel, fields, relations)
}

func (svc *AccountCashRoundingService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountCashRoundingModel, ids, fields, relations)
}

func (svc *AccountCashRoundingService) Delete(ids []int) error {
	return svc.client.delete(types.AccountCashRoundingModel, ids)
}
