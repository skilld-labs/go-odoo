package api

import (
	"github.com/morezig/go-odoo/types"
)

type CrmLeadLostService struct {
	client *Client
}

func NewCrmLeadLostService(c *Client) *CrmLeadLostService {
	return &CrmLeadLostService{client: c}
}

func (svc *CrmLeadLostService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.CrmLeadLostModel, name)
}

func (svc *CrmLeadLostService) GetByIds(ids []int64) (*types.CrmLeadLosts, error) {
	c := &types.CrmLeadLosts{}
	return c, svc.client.getByIds(types.CrmLeadLostModel, ids, c)
}

func (svc *CrmLeadLostService) GetByName(name string) (*types.CrmLeadLosts, error) {
	c := &types.CrmLeadLosts{}
	return c, svc.client.getByName(types.CrmLeadLostModel, name, c)
}

func (svc *CrmLeadLostService) GetByField(field string, value string) (*types.CrmLeadLosts, error) {
	c := &types.CrmLeadLosts{}
	return c, svc.client.getByField(types.CrmLeadLostModel, field, value, c)
}

func (svc *CrmLeadLostService) GetAll() (*types.CrmLeadLosts, error) {
	c := &types.CrmLeadLosts{}
	return c, svc.client.getAll(types.CrmLeadLostModel, c)
}

func (svc *CrmLeadLostService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.CrmLeadLostModel, fields, relations)
}

func (svc *CrmLeadLostService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CrmLeadLostModel, ids, fields, relations)
}

func (svc *CrmLeadLostService) Delete(ids []int64) error {
	return svc.client.delete(types.CrmLeadLostModel, ids)
}
