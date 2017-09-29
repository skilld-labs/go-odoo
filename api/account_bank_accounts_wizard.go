package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountBankAccountsWizardService struct {
	client *Client
}

func NewAccountBankAccountsWizardService(c *Client) *AccountBankAccountsWizardService {
	return &AccountBankAccountsWizardService{client: c}
}

func (svc *AccountBankAccountsWizardService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountBankAccountsWizardModel, name)
}

func (svc *AccountBankAccountsWizardService) GetByIds(ids []int64) (*types.AccountBankAccountsWizards, error) {
	a := &types.AccountBankAccountsWizards{}
	return a, svc.client.getByIds(types.AccountBankAccountsWizardModel, ids, a)
}

func (svc *AccountBankAccountsWizardService) GetByName(name string) (*types.AccountBankAccountsWizards, error) {
	a := &types.AccountBankAccountsWizards{}
	return a, svc.client.getByName(types.AccountBankAccountsWizardModel, name, a)
}

func (svc *AccountBankAccountsWizardService) GetByField(field string, value string) (*types.AccountBankAccountsWizards, error) {
	a := &types.AccountBankAccountsWizards{}
	return a, svc.client.getByField(types.AccountBankAccountsWizardModel, field, value, a)
}

func (svc *AccountBankAccountsWizardService) GetAll() (*types.AccountBankAccountsWizards, error) {
	a := &types.AccountBankAccountsWizards{}
	return a, svc.client.getAll(types.AccountBankAccountsWizardModel, a)
}

func (svc *AccountBankAccountsWizardService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountBankAccountsWizardModel, fields, relations)
}

func (svc *AccountBankAccountsWizardService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountBankAccountsWizardModel, ids, fields, relations)
}

func (svc *AccountBankAccountsWizardService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountBankAccountsWizardModel, ids)
}
