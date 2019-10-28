package api

import (
	"github.com/morezig/go-odoo/types"
)

type AccountFiscalPositionTaxTemplateService struct {
	client *Client
}

func NewAccountFiscalPositionTaxTemplateService(c *Client) *AccountFiscalPositionTaxTemplateService {
	return &AccountFiscalPositionTaxTemplateService{client: c}
}

func (svc *AccountFiscalPositionTaxTemplateService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountFiscalPositionTaxTemplateModel, name)
}

func (svc *AccountFiscalPositionTaxTemplateService) GetByIds(ids []int64) (*types.AccountFiscalPositionTaxTemplates, error) {
	a := &types.AccountFiscalPositionTaxTemplates{}
	return a, svc.client.getByIds(types.AccountFiscalPositionTaxTemplateModel, ids, a)
}

func (svc *AccountFiscalPositionTaxTemplateService) GetByName(name string) (*types.AccountFiscalPositionTaxTemplates, error) {
	a := &types.AccountFiscalPositionTaxTemplates{}
	return a, svc.client.getByName(types.AccountFiscalPositionTaxTemplateModel, name, a)
}

func (svc *AccountFiscalPositionTaxTemplateService) GetByField(field string, value string) (*types.AccountFiscalPositionTaxTemplates, error) {
	a := &types.AccountFiscalPositionTaxTemplates{}
	return a, svc.client.getByField(types.AccountFiscalPositionTaxTemplateModel, field, value, a)
}

func (svc *AccountFiscalPositionTaxTemplateService) GetAll() (*types.AccountFiscalPositionTaxTemplates, error) {
	a := &types.AccountFiscalPositionTaxTemplates{}
	return a, svc.client.getAll(types.AccountFiscalPositionTaxTemplateModel, a)
}

func (svc *AccountFiscalPositionTaxTemplateService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountFiscalPositionTaxTemplateModel, fields, relations)
}

func (svc *AccountFiscalPositionTaxTemplateService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountFiscalPositionTaxTemplateModel, ids, fields, relations)
}

func (svc *AccountFiscalPositionTaxTemplateService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountFiscalPositionTaxTemplateModel, ids)
}
