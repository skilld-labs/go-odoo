package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountBankStatementImportService struct {
	client *Client
}

func NewAccountBankStatementImportService(c *Client) *AccountBankStatementImportService {
	return &AccountBankStatementImportService{client: c}
}

func (svc *AccountBankStatementImportService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountBankStatementImportModel, name)
}

func (svc *AccountBankStatementImportService) GetByIds(ids []int) (*types.AccountBankStatementImports, error) {
	a := &types.AccountBankStatementImports{}
	return a, svc.client.getByIds(types.AccountBankStatementImportModel, ids, a)
}

func (svc *AccountBankStatementImportService) GetByName(name string) (*types.AccountBankStatementImports, error) {
	a := &types.AccountBankStatementImports{}
	return a, svc.client.getByName(types.AccountBankStatementImportModel, name, a)
}

func (svc *AccountBankStatementImportService) GetByField(field string, value string) (*types.AccountBankStatementImports, error) {
	a := &types.AccountBankStatementImports{}
	return a, svc.client.getByName(types.AccountBankStatementImportModel, field, value, a)
}

func (svc *AccountBankStatementImportService) GetAll() (*types.AccountBankStatementImports, error) {
	a := &types.AccountBankStatementImports{}
	return a, svc.client.getAll(types.AccountBankStatementImportModel, a)
}

func (svc *AccountBankStatementImportService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountBankStatementImportModel, fields, relations)
}

func (svc *AccountBankStatementImportService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountBankStatementImportModel, ids, fields, relations)
}

func (svc *AccountBankStatementImportService) Delete(ids []int) error {
	return svc.client.delete(types.AccountBankStatementImportModel, ids)
}
