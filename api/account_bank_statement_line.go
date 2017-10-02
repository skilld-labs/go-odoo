package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountBankStatementLineService struct {
	client *Client
}

func NewAccountBankStatementLineService(c *Client) *AccountBankStatementLineService {
	return &AccountBankStatementLineService{client: c}
}

func (svc *AccountBankStatementLineService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountBankStatementLineModel, name)
}

func (svc *AccountBankStatementLineService) GetByIds(ids []int64) (*types.AccountBankStatementLines, error) {
	a := &types.AccountBankStatementLines{}
	return a, svc.client.getByIds(types.AccountBankStatementLineModel, ids, a)
}

func (svc *AccountBankStatementLineService) GetByName(name string) (*types.AccountBankStatementLines, error) {
	a := &types.AccountBankStatementLines{}
	return a, svc.client.getByName(types.AccountBankStatementLineModel, name, a)
}

func (svc *AccountBankStatementLineService) GetByField(field string, value string) (*types.AccountBankStatementLines, error) {
	a := &types.AccountBankStatementLines{}
	return a, svc.client.getByField(types.AccountBankStatementLineModel, field, value, a)
}

func (svc *AccountBankStatementLineService) GetAll() (*types.AccountBankStatementLines, error) {
	a := &types.AccountBankStatementLines{}
	return a, svc.client.getAll(types.AccountBankStatementLineModel, a)
}

func (svc *AccountBankStatementLineService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountBankStatementLineModel, fields, relations)
}

func (svc *AccountBankStatementLineService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountBankStatementLineModel, ids, fields, relations)
}

func (svc *AccountBankStatementLineService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountBankStatementLineModel, ids)
}
