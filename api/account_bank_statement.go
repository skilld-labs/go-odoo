package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountBankStatementService struct {
	client *Client
}

func NewAccountBankStatementService(c *Client) *AccountBankStatementService {
	return &AccountBankStatementService{client: c}
}

func (svc *AccountBankStatementService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountBankStatementModel, name)
}

func (svc *AccountBankStatementService) GetByIds(ids []int64) (*types.AccountBankStatements, error) {
	a := &types.AccountBankStatements{}
	return a, svc.client.getByIds(types.AccountBankStatementModel, ids, a)
}

func (svc *AccountBankStatementService) GetByName(name string) (*types.AccountBankStatements, error) {
	a := &types.AccountBankStatements{}
	return a, svc.client.getByName(types.AccountBankStatementModel, name, a)
}

func (svc *AccountBankStatementService) GetByField(field string, value string) (*types.AccountBankStatements, error) {
	a := &types.AccountBankStatements{}
	return a, svc.client.getByField(types.AccountBankStatementModel, field, value, a)
}

func (svc *AccountBankStatementService) GetAll() (*types.AccountBankStatements, error) {
	a := &types.AccountBankStatements{}
	return a, svc.client.getAll(types.AccountBankStatementModel, a)
}

func (svc *AccountBankStatementService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountBankStatementModel, fields, relations)
}

func (svc *AccountBankStatementService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountBankStatementModel, ids, fields, relations)
}

func (svc *AccountBankStatementService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountBankStatementModel, ids)
}
