package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountChartTemplateService struct {
	client *Client
}

func NewAccountChartTemplateService(c *Client) *AccountChartTemplateService {
	return &AccountChartTemplateService{client: c}
}

func (svc *AccountChartTemplateService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountChartTemplateModel, name)
}

func (svc *AccountChartTemplateService) GetByIds(ids []int) (*types.AccountChartTemplates, error) {
	a := &types.AccountChartTemplates{}
	return a, svc.client.getByIds(types.AccountChartTemplateModel, ids, a)
}

func (svc *AccountChartTemplateService) GetByName(name string) (*types.AccountChartTemplates, error) {
	a := &types.AccountChartTemplates{}
	return a, svc.client.getByName(types.AccountChartTemplateModel, name, a)
}

func (svc *AccountChartTemplateService) GetByField(field string, value string) (*types.AccountChartTemplates, error) {
	a := &types.AccountChartTemplates{}
	return a, svc.client.getByName(types.AccountChartTemplateModel, field, value, a)
}

func (svc *AccountChartTemplateService) GetAll() (*types.AccountChartTemplates, error) {
	a := &types.AccountChartTemplates{}
	return a, svc.client.getAll(types.AccountChartTemplateModel, a)
}

func (svc *AccountChartTemplateService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountChartTemplateModel, fields, relations)
}

func (svc *AccountChartTemplateService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountChartTemplateModel, ids, fields, relations)
}

func (svc *AccountChartTemplateService) Delete(ids []int) error {
	return svc.client.delete(types.AccountChartTemplateModel, ids)
}
