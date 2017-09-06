package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountFiscalPositionAccountTemplateService struct {
	client *Client
}

func NewAccountFiscalPositionAccountTemplateService(c *Client) *AccountFiscalPositionAccountTemplateService {
	return &AccountFiscalPositionAccountTemplateService{client: c}
}

func (svc *AccountFiscalPositionAccountTemplateService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountFiscalPositionAccountTemplateModel, name)
}

func (svc *AccountFiscalPositionAccountTemplateService) GetByIds(ids []int) (*types.AccountFiscalPositionAccountTemplates, error) {
	a := &types.AccountFiscalPositionAccountTemplates{}
	return a, svc.client.getByIds(types.AccountFiscalPositionAccountTemplateModel, ids, a)
}

func (svc *AccountFiscalPositionAccountTemplateService) GetByName(name string) (*types.AccountFiscalPositionAccountTemplates, error) {
	a := &types.AccountFiscalPositionAccountTemplates{}
	return a, svc.client.getByName(types.AccountFiscalPositionAccountTemplateModel, name, a)
}

func (svc *AccountFiscalPositionAccountTemplateService) GetByField(field string, value string) (*types.AccountFiscalPositionAccountTemplates, error) {
	a := &types.AccountFiscalPositionAccountTemplates{}
	return a, svc.client.getByField(types.AccountFiscalPositionAccountTemplateModel, field, value, a)
}

func (svc *AccountFiscalPositionAccountTemplateService) GetAll() (*types.AccountFiscalPositionAccountTemplates, error) {
	a := &types.AccountFiscalPositionAccountTemplates{}
	return a, svc.client.getAll(types.AccountFiscalPositionAccountTemplateModel, a)
}

func (svc *AccountFiscalPositionAccountTemplateService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountFiscalPositionAccountTemplateModel, fields, relations)
}

func (svc *AccountFiscalPositionAccountTemplateService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountFiscalPositionAccountTemplateModel, ids, fields, relations)
}

func (svc *AccountFiscalPositionAccountTemplateService) Delete(ids []int) error {
	return svc.client.delete(types.AccountFiscalPositionAccountTemplateModel, ids)
}
