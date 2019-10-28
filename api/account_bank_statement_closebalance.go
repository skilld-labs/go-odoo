package api

import (
	"github.com/morezig/go-odoo/types"
)

type AccountBankStatementClosebalanceService struct {
	client *Client
}

func NewAccountBankStatementClosebalanceService(c *Client) *AccountBankStatementClosebalanceService {
	return &AccountBankStatementClosebalanceService{client: c}
}

func (svc *AccountBankStatementClosebalanceService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountBankStatementClosebalanceModel, name)
}

func (svc *AccountBankStatementClosebalanceService) GetByIds(ids []int64) (*types.AccountBankStatementClosebalances, error) {
	a := &types.AccountBankStatementClosebalances{}
	return a, svc.client.getByIds(types.AccountBankStatementClosebalanceModel, ids, a)
}

func (svc *AccountBankStatementClosebalanceService) GetByName(name string) (*types.AccountBankStatementClosebalances, error) {
	a := &types.AccountBankStatementClosebalances{}
	return a, svc.client.getByName(types.AccountBankStatementClosebalanceModel, name, a)
}

func (svc *AccountBankStatementClosebalanceService) GetByField(field string, value string) (*types.AccountBankStatementClosebalances, error) {
	a := &types.AccountBankStatementClosebalances{}
	return a, svc.client.getByField(types.AccountBankStatementClosebalanceModel, field, value, a)
}

func (svc *AccountBankStatementClosebalanceService) GetAll() (*types.AccountBankStatementClosebalances, error) {
	a := &types.AccountBankStatementClosebalances{}
	return a, svc.client.getAll(types.AccountBankStatementClosebalanceModel, a)
}

func (svc *AccountBankStatementClosebalanceService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountBankStatementClosebalanceModel, fields, relations)
}

func (svc *AccountBankStatementClosebalanceService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountBankStatementClosebalanceModel, ids, fields, relations)
}

func (svc *AccountBankStatementClosebalanceService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountBankStatementClosebalanceModel, ids)
}
