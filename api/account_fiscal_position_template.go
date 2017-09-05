package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountFiscalPositionTemplateService struct {
	client *Client
}

func NewAccountFiscalPositionTemplateService(c *Client) *AccountFiscalPositionTemplateService {
	return &AccountFiscalPositionTemplateService{client: c}
}

func (svc *AccountFiscalPositionTemplateService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountFiscalPositionTemplateModel, name)
}

func (svc *AccountFiscalPositionTemplateService) GetByIds(ids []int) (*types.AccountFiscalPositionTemplates, error) {
	a := &types.AccountFiscalPositionTemplates{}
	return a, svc.client.getByIds(types.AccountFiscalPositionTemplateModel, ids, a)
}

func (svc *AccountFiscalPositionTemplateService) GetByName(name string) (*types.AccountFiscalPositionTemplates, error) {
	a := &types.AccountFiscalPositionTemplates{}
	return a, svc.client.getByName(types.AccountFiscalPositionTemplateModel, name, a)
}

func (svc *AccountFiscalPositionTemplateService) GetByField(field string, value string) (*types.AccountFiscalPositionTemplates, error) {
	a := &types.AccountFiscalPositionTemplates{}
	return a, svc.client.getByField(types.AccountFiscalPositionTemplateModel, field, value, a)
}

func (svc *AccountFiscalPositionTemplateService) GetAll() (*types.AccountFiscalPositionTemplates, error) {
	a := &types.AccountFiscalPositionTemplates{}
	return a, svc.client.getAll(types.AccountFiscalPositionTemplateModel, a)
}

func (svc *AccountFiscalPositionTemplateService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountFiscalPositionTemplateModel, fields, relations)
}

func (svc *AccountFiscalPositionTemplateService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountFiscalPositionTemplateModel, ids, fields, relations)
}

func (svc *AccountFiscalPositionTemplateService) Delete(ids []int) error {
	return svc.client.delete(types.AccountFiscalPositionTemplateModel, ids)
}
