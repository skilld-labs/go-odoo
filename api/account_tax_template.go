package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountTaxTemplateService struct {
	client *Client
}

func NewAccountTaxTemplateService(c *Client) *AccountTaxTemplateService {
	return &AccountTaxTemplateService{client: c}
}

func (svc *AccountTaxTemplateService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountTaxTemplateModel, name)
}

func (svc *AccountTaxTemplateService) GetByIds(ids []int64) (*types.AccountTaxTemplates, error) {
	a := &types.AccountTaxTemplates{}
	return a, svc.client.getByIds(types.AccountTaxTemplateModel, ids, a)
}

func (svc *AccountTaxTemplateService) GetByName(name string) (*types.AccountTaxTemplates, error) {
	a := &types.AccountTaxTemplates{}
	return a, svc.client.getByName(types.AccountTaxTemplateModel, name, a)
}

func (svc *AccountTaxTemplateService) GetByField(field string, value string) (*types.AccountTaxTemplates, error) {
	a := &types.AccountTaxTemplates{}
	return a, svc.client.getByField(types.AccountTaxTemplateModel, field, value, a)
}

func (svc *AccountTaxTemplateService) GetAll() (*types.AccountTaxTemplates, error) {
	a := &types.AccountTaxTemplates{}
	return a, svc.client.getAll(types.AccountTaxTemplateModel, a)
}

func (svc *AccountTaxTemplateService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountTaxTemplateModel, fields, relations)
}

func (svc *AccountTaxTemplateService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountTaxTemplateModel, ids, fields, relations)
}

func (svc *AccountTaxTemplateService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountTaxTemplateModel, ids)
}
