package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountReconcileModelTemplateService struct {
	client *Client
}

func NewAccountReconcileModelTemplateService(c *Client) *AccountReconcileModelTemplateService {
	return &AccountReconcileModelTemplateService{client: c}
}

func (svc *AccountReconcileModelTemplateService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountReconcileModelTemplateModel, name)
}

func (svc *AccountReconcileModelTemplateService) GetByIds(ids []int) (*types.AccountReconcileModelTemplates, error) {
	a := &types.AccountReconcileModelTemplates{}
	return a, svc.client.getByIds(types.AccountReconcileModelTemplateModel, ids, a)
}

func (svc *AccountReconcileModelTemplateService) GetByName(name string) (*types.AccountReconcileModelTemplates, error) {
	a := &types.AccountReconcileModelTemplates{}
	return a, svc.client.getByName(types.AccountReconcileModelTemplateModel, name, a)
}

func (svc *AccountReconcileModelTemplateService) GetByField(field string, value string) (*types.AccountReconcileModelTemplates, error) {
	a := &types.AccountReconcileModelTemplates{}
	return a, svc.client.getByField(types.AccountReconcileModelTemplateModel, field, value, a)
}

func (svc *AccountReconcileModelTemplateService) GetAll() (*types.AccountReconcileModelTemplates, error) {
	a := &types.AccountReconcileModelTemplates{}
	return a, svc.client.getAll(types.AccountReconcileModelTemplateModel, a)
}

func (svc *AccountReconcileModelTemplateService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountReconcileModelTemplateModel, fields, relations)
}

func (svc *AccountReconcileModelTemplateService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountReconcileModelTemplateModel, ids, fields, relations)
}

func (svc *AccountReconcileModelTemplateService) Delete(ids []int) error {
	return svc.client.delete(types.AccountReconcileModelTemplateModel, ids)
}
