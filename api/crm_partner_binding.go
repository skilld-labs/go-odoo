package api

import (
	"github.com/morezig/go-odoo/types"
)

type CrmPartnerBindingService struct {
	client *Client
}

func NewCrmPartnerBindingService(c *Client) *CrmPartnerBindingService {
	return &CrmPartnerBindingService{client: c}
}

func (svc *CrmPartnerBindingService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.CrmPartnerBindingModel, name)
}

func (svc *CrmPartnerBindingService) GetByIds(ids []int64) (*types.CrmPartnerBindings, error) {
	c := &types.CrmPartnerBindings{}
	return c, svc.client.getByIds(types.CrmPartnerBindingModel, ids, c)
}

func (svc *CrmPartnerBindingService) GetByName(name string) (*types.CrmPartnerBindings, error) {
	c := &types.CrmPartnerBindings{}
	return c, svc.client.getByName(types.CrmPartnerBindingModel, name, c)
}

func (svc *CrmPartnerBindingService) GetByField(field string, value string) (*types.CrmPartnerBindings, error) {
	c := &types.CrmPartnerBindings{}
	return c, svc.client.getByField(types.CrmPartnerBindingModel, field, value, c)
}

func (svc *CrmPartnerBindingService) GetAll() (*types.CrmPartnerBindings, error) {
	c := &types.CrmPartnerBindings{}
	return c, svc.client.getAll(types.CrmPartnerBindingModel, c)
}

func (svc *CrmPartnerBindingService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.CrmPartnerBindingModel, fields, relations)
}

func (svc *CrmPartnerBindingService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CrmPartnerBindingModel, ids, fields, relations)
}

func (svc *CrmPartnerBindingService) Delete(ids []int64) error {
	return svc.client.delete(types.CrmPartnerBindingModel, ids)
}
