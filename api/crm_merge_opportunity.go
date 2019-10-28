package api

import (
	"github.com/morezig/go-odoo/types"
)

type CrmMergeOpportunityService struct {
	client *Client
}

func NewCrmMergeOpportunityService(c *Client) *CrmMergeOpportunityService {
	return &CrmMergeOpportunityService{client: c}
}

func (svc *CrmMergeOpportunityService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.CrmMergeOpportunityModel, name)
}

func (svc *CrmMergeOpportunityService) GetByIds(ids []int64) (*types.CrmMergeOpportunitys, error) {
	c := &types.CrmMergeOpportunitys{}
	return c, svc.client.getByIds(types.CrmMergeOpportunityModel, ids, c)
}

func (svc *CrmMergeOpportunityService) GetByName(name string) (*types.CrmMergeOpportunitys, error) {
	c := &types.CrmMergeOpportunitys{}
	return c, svc.client.getByName(types.CrmMergeOpportunityModel, name, c)
}

func (svc *CrmMergeOpportunityService) GetByField(field string, value string) (*types.CrmMergeOpportunitys, error) {
	c := &types.CrmMergeOpportunitys{}
	return c, svc.client.getByField(types.CrmMergeOpportunityModel, field, value, c)
}

func (svc *CrmMergeOpportunityService) GetAll() (*types.CrmMergeOpportunitys, error) {
	c := &types.CrmMergeOpportunitys{}
	return c, svc.client.getAll(types.CrmMergeOpportunityModel, c)
}

func (svc *CrmMergeOpportunityService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.CrmMergeOpportunityModel, fields, relations)
}

func (svc *CrmMergeOpportunityService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CrmMergeOpportunityModel, ids, fields, relations)
}

func (svc *CrmMergeOpportunityService) Delete(ids []int64) error {
	return svc.client.delete(types.CrmMergeOpportunityModel, ids)
}
