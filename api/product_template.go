package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ProductTemplateService struct {
	client *Client
}

func NewProductTemplateService(c *Client) *ProductTemplateService {
	return &ProductTemplateService{client: c}
}

func (svc *ProductTemplateService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ProductTemplateModel, name)
}

func (svc *ProductTemplateService) GetByIds(ids []int64) (*types.ProductTemplates, error) {
	p := &types.ProductTemplates{}
	return p, svc.client.getByIds(types.ProductTemplateModel, ids, p)
}

func (svc *ProductTemplateService) GetByName(name string) (*types.ProductTemplates, error) {
	p := &types.ProductTemplates{}
	return p, svc.client.getByName(types.ProductTemplateModel, name, p)
}

func (svc *ProductTemplateService) GetByField(field string, value string) (*types.ProductTemplates, error) {
	p := &types.ProductTemplates{}
	return p, svc.client.getByField(types.ProductTemplateModel, field, value, p)
}

func (svc *ProductTemplateService) GetAll() (*types.ProductTemplates, error) {
	p := &types.ProductTemplates{}
	return p, svc.client.getAll(types.ProductTemplateModel, p)
}

func (svc *ProductTemplateService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ProductTemplateModel, fields, relations)
}

func (svc *ProductTemplateService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProductTemplateModel, ids, fields, relations)
}

func (svc *ProductTemplateService) Delete(ids []int64) error {
	return svc.client.delete(types.ProductTemplateModel, ids)
}
