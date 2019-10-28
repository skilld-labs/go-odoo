package api

import (
	"github.com/morezig/go-odoo/types"
)

type AccountAccountTemplateService struct {
	client *Client
}

func NewAccountAccountTemplateService(c *Client) *AccountAccountTemplateService {
	return &AccountAccountTemplateService{client: c}
}

func (svc *AccountAccountTemplateService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountAccountTemplateModel, name)
}

func (svc *AccountAccountTemplateService) GetByIds(ids []int64) (*types.AccountAccountTemplates, error) {
	a := &types.AccountAccountTemplates{}
	return a, svc.client.getByIds(types.AccountAccountTemplateModel, ids, a)
}

func (svc *AccountAccountTemplateService) GetByName(name string) (*types.AccountAccountTemplates, error) {
	a := &types.AccountAccountTemplates{}
	return a, svc.client.getByName(types.AccountAccountTemplateModel, name, a)
}

func (svc *AccountAccountTemplateService) GetByField(field string, value string) (*types.AccountAccountTemplates, error) {
	a := &types.AccountAccountTemplates{}
	return a, svc.client.getByField(types.AccountAccountTemplateModel, field, value, a)
}

func (svc *AccountAccountTemplateService) GetAll() (*types.AccountAccountTemplates, error) {
	a := &types.AccountAccountTemplates{}
	return a, svc.client.getAll(types.AccountAccountTemplateModel, a)
}

func (svc *AccountAccountTemplateService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountAccountTemplateModel, fields, relations)
}

func (svc *AccountAccountTemplateService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountAccountTemplateModel, ids, fields, relations)
}

func (svc *AccountAccountTemplateService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountAccountTemplateModel, ids)
}
