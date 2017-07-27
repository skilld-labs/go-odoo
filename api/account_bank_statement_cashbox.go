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

func (svc *AccountBankStatementCashboxService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountBankStatementCashboxModel, name)
}

func (svc *AccountBankStatementCashboxService) GetByIds(ids []int) (*types.AccountBankStatementCashboxs, error) {
	a := &types.AccountBankStatementCashboxs{}
	return a, svc.client.getByIds(types.AccountBankStatementCashboxModel, ids, a)
}

func (svc *AccountBankStatementCashboxService) GetByName(name string) (*types.AccountBankStatementCashboxs, error) {
	a := &types.AccountBankStatementCashboxs{}
	return a, svc.client.getByName(types.AccountBankStatementCashboxModel, name, a)
}

func (svc *AccountBankStatementCashboxService) GetAll() (*types.AccountBankStatementCashboxs, error) {
	a := &types.AccountBankStatementCashboxs{}
	return a, svc.client.getAll(types.AccountBankStatementCashboxModel, a)
}

func (svc *AccountBankStatementCashboxService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountBankStatementCashboxModel, fields, relations)
}

func (svc *AccountBankStatementCashboxService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountBankStatementCashboxModel, ids, fields, relations)
}

func (svc *AccountBankStatementCashboxService) Delete(ids []int) error {
	return svc.client.delete(types.AccountBankStatementCashboxModel, ids)
}
