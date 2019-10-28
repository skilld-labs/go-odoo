package api

import (
	"github.com/morezig/go-odoo/types"
)

type ResourceTestService struct {
	client *Client
}

func NewResourceTestService(c *Client) *ResourceTestService {
	return &ResourceTestService{client: c}
}

func (svc *ResourceTestService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ResourceTestModel, name)
}

func (svc *ResourceTestService) GetByIds(ids []int64) (*types.ResourceTests, error) {
	r := &types.ResourceTests{}
	return r, svc.client.getByIds(types.ResourceTestModel, ids, r)
}

func (svc *ResourceTestService) GetByName(name string) (*types.ResourceTests, error) {
	r := &types.ResourceTests{}
	return r, svc.client.getByName(types.ResourceTestModel, name, r)
}

func (svc *ResourceTestService) GetByField(field string, value string) (*types.ResourceTests, error) {
	r := &types.ResourceTests{}
	return r, svc.client.getByField(types.ResourceTestModel, field, value, r)
}

func (svc *ResourceTestService) GetAll() (*types.ResourceTests, error) {
	r := &types.ResourceTests{}
	return r, svc.client.getAll(types.ResourceTestModel, r)
}

func (svc *ResourceTestService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ResourceTestModel, fields, relations)
}

func (svc *ResourceTestService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResourceTestModel, ids, fields, relations)
}

func (svc *ResourceTestService) Delete(ids []int64) error {
	return svc.client.delete(types.ResourceTestModel, ids)
}
