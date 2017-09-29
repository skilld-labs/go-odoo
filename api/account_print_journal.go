package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountPrintJournalService struct {
	client *Client
}

func NewAccountPrintJournalService(c *Client) *AccountPrintJournalService {
	return &AccountPrintJournalService{client: c}
}

func (svc *AccountPrintJournalService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountPrintJournalModel, name)
}

func (svc *AccountPrintJournalService) GetByIds(ids []int64) (*types.AccountPrintJournals, error) {
	a := &types.AccountPrintJournals{}
	return a, svc.client.getByIds(types.AccountPrintJournalModel, ids, a)
}

func (svc *AccountPrintJournalService) GetByName(name string) (*types.AccountPrintJournals, error) {
	a := &types.AccountPrintJournals{}
	return a, svc.client.getByName(types.AccountPrintJournalModel, name, a)
}

func (svc *AccountPrintJournalService) GetByField(field string, value string) (*types.AccountPrintJournals, error) {
	a := &types.AccountPrintJournals{}
	return a, svc.client.getByField(types.AccountPrintJournalModel, field, value, a)
}

func (svc *AccountPrintJournalService) GetAll() (*types.AccountPrintJournals, error) {
	a := &types.AccountPrintJournals{}
	return a, svc.client.getAll(types.AccountPrintJournalModel, a)
}

func (svc *AccountPrintJournalService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountPrintJournalModel, fields, relations)
}

func (svc *AccountPrintJournalService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountPrintJournalModel, ids, fields, relations)
}

func (svc *AccountPrintJournalService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountPrintJournalModel, ids)
}
