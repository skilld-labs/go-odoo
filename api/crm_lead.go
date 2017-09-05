package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type CrmLeadService struct {
	client *Client
}

func NewCrmLeadService(c *Client) *CrmLeadService {
	return &CrmLeadService{client: c}
}

func (svc *CrmLeadService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.CrmLeadModel, name)
}

func (svc *CrmLeadService) GetByIds(ids []int) (*types.CrmLeads, error) {
	c := &types.CrmLeads{}
	return c, svc.client.getByIds(types.CrmLeadModel, ids, c)
}

func (svc *CrmLeadService) GetByName(name string) (*types.CrmLeads, error) {
	c := &types.CrmLeads{}
	return c, svc.client.getByName(types.CrmLeadModel, name, c)
}

func (svc *CrmLeadService) GetByField(field string, value string) (*types.CrmLeads, error) {
	c := &types.CrmLeads{}
	return c, svc.client.getByName(types.CrmLeadModel, field, value, c)
}

func (svc *CrmLeadService) GetAll() (*types.CrmLeads, error) {
	c := &types.CrmLeads{}
	return c, svc.client.getAll(types.CrmLeadModel, c)
}

func (svc *CrmLeadService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.CrmLeadModel, fields, relations)
}

func (svc *CrmLeadService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CrmLeadModel, ids, fields, relations)
}

func (svc *CrmLeadService) Delete(ids []int) error {
	return svc.client.delete(types.CrmLeadModel, ids)
}
