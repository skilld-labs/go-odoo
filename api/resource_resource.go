package api

import (
	"github.com/morezig/go-odoo/types"
)

type ResourceResourceService struct {
	client *Client
}

func NewResourceResourceService(c *Client) *ResourceResourceService {
	return &ResourceResourceService{client: c}
}

func (svc *ResourceResourceService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ResourceResourceModel, name)
}

func (svc *ResourceResourceService) GetByIds(ids []int64) (*types.ResourceResources, error) {
	r := &types.ResourceResources{}
	return r, svc.client.getByIds(types.ResourceResourceModel, ids, r)
}

func (svc *ResourceResourceService) GetByName(name string) (*types.ResourceResources, error) {
	r := &types.ResourceResources{}
	return r, svc.client.getByName(types.ResourceResourceModel, name, r)
}

func (svc *ResourceResourceService) GetByField(field string, value string) (*types.ResourceResources, error) {
	r := &types.ResourceResources{}
	return r, svc.client.getByField(types.ResourceResourceModel, field, value, r)
}

func (svc *ResourceResourceService) GetAll() (*types.ResourceResources, error) {
	r := &types.ResourceResources{}
	return r, svc.client.getAll(types.ResourceResourceModel, r)
}

func (svc *ResourceResourceService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ResourceResourceModel, fields, relations)
}

func (svc *ResourceResourceService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResourceResourceModel, ids, fields, relations)
}

func (svc *ResourceResourceService) Delete(ids []int64) error {
	return svc.client.delete(types.ResourceResourceModel, ids)
}
