package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type CrmLeadTagService struct {
	client *Client
}

func NewCrmLeadTagService(c *Client) *CrmLeadTagService {
	return &CrmLeadTagService{client: c}
}

func (svc *CrmLeadTagService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.CrmLeadTagModel, name)
}

func (svc *CrmLeadTagService) GetByIds(ids []int) (*types.CrmLeadTags, error) {
	c := &types.CrmLeadTags{}
	return c, svc.client.getByIds(types.CrmLeadTagModel, ids, c)
}

func (svc *CrmLeadTagService) GetByName(name string) (*types.CrmLeadTags, error) {
	c := &types.CrmLeadTags{}
	return c, svc.client.getByName(types.CrmLeadTagModel, name, c)
}

func (svc *CrmLeadTagService) GetByField(field string, value string) (*types.CrmLeadTags, error) {
	c := &types.CrmLeadTags{}
	return c, svc.client.getByField(types.CrmLeadTagModel, field, value, c)
}

func (svc *CrmLeadTagService) GetAll() (*types.CrmLeadTags, error) {
	c := &types.CrmLeadTags{}
	return c, svc.client.getAll(types.CrmLeadTagModel, c)
}

func (svc *CrmLeadTagService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.CrmLeadTagModel, fields, relations)
}

func (svc *CrmLeadTagService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CrmLeadTagModel, ids, fields, relations)
}

func (svc *CrmLeadTagService) Delete(ids []int) error {
	return svc.client.delete(types.CrmLeadTagModel, ids)
}
