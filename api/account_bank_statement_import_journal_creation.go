package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountBankStatementImportJournalCreationService struct {
	client *Client
}

func NewAccountBankStatementImportJournalCreationService(c *Client) *AccountBankStatementImportJournalCreationService {
	return &AccountBankStatementImportJournalCreationService{client: c}
}

func (svc *AccountBankStatementImportJournalCreationService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountBankStatementImportJournalCreationModel, name)
}

func (svc *AccountBankStatementImportJournalCreationService) GetByIds(ids []int) (*types.AccountBankStatementImportJournalCreations, error) {
	a := &types.AccountBankStatementImportJournalCreations{}
	return a, svc.client.getByIds(types.AccountBankStatementImportJournalCreationModel, ids, a)
}

func (svc *AccountBankStatementImportJournalCreationService) GetByName(name string) (*types.AccountBankStatementImportJournalCreations, error) {
	a := &types.AccountBankStatementImportJournalCreations{}
	return a, svc.client.getByName(types.AccountBankStatementImportJournalCreationModel, name, a)
}

func (svc *AccountBankStatementImportJournalCreationService) GetAll() (*types.AccountBankStatementImportJournalCreations, error) {
	a := &types.AccountBankStatementImportJournalCreations{}
	return a, svc.client.getAll(types.AccountBankStatementImportJournalCreationModel, a)
}

func (svc *AccountBankStatementImportJournalCreationService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountBankStatementImportJournalCreationModel, fields, relations)
}

func (svc *AccountBankStatementImportJournalCreationService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountBankStatementImportJournalCreationModel, ids, fields, relations)
}

func (svc *AccountBankStatementImportJournalCreationService) Delete(ids []int) error {
	return svc.client.delete(types.AccountBankStatementImportJournalCreationModel, ids)
}
