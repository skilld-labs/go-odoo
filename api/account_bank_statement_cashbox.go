package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountBankStatementCashboxService struct {
	client *Client
}

func NewAccountBankStatementCashboxService(c *Client) *AccountBankStatementCashboxService {
	return &AccountBankStatementCashboxService{client: c}
}

func (svc *AccountBankStatementCashboxService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountBankStatementCashboxModel, name)
}

func (svc *AccountBankStatementCashboxService) GetByIds(ids []int64) (*types.AccountBankStatementCashboxs, error) {
	a := &types.AccountBankStatementCashboxs{}
	return a, svc.client.getByIds(types.AccountBankStatementCashboxModel, ids, a)
}

func (svc *AccountBankStatementCashboxService) GetByName(name string) (*types.AccountBankStatementCashboxs, error) {
	a := &types.AccountBankStatementCashboxs{}
	return a, svc.client.getByName(types.AccountBankStatementCashboxModel, name, a)
}

func (svc *AccountBankStatementCashboxService) GetByField(field string, value string) (*types.AccountBankStatementCashboxs, error) {
	a := &types.AccountBankStatementCashboxs{}
	return a, svc.client.getByField(types.AccountBankStatementCashboxModel, field, value, a)
}

func (svc *AccountBankStatementCashboxService) GetAll() (*types.AccountBankStatementCashboxs, error) {
	a := &types.AccountBankStatementCashboxs{}
	return a, svc.client.getAll(types.AccountBankStatementCashboxModel, a)
}

func (svc *AccountBankStatementCashboxService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountBankStatementCashboxModel, fields, relations)
}

func (svc *AccountBankStatementCashboxService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountBankStatementCashboxModel, ids, fields, relations)
}

func (svc *AccountBankStatementCashboxService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountBankStatementCashboxModel, ids)
}
