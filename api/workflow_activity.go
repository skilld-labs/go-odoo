package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type WorkflowActivityService struct {
	client *Client
}

func NewWorkflowActivityService(c *Client) *WorkflowActivityService {
	return &WorkflowActivityService{client: c}
}

func (svc *WorkflowActivityService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.WorkflowActivityModel, name)
}

func (svc *WorkflowActivityService) GetByIds(ids []int64) (*types.WorkflowActivitys, error) {
	w := &types.WorkflowActivitys{}
	return w, svc.client.getByIds(types.WorkflowActivityModel, ids, w)
}

func (svc *WorkflowActivityService) GetByName(name string) (*types.WorkflowActivitys, error) {
	w := &types.WorkflowActivitys{}
	return w, svc.client.getByName(types.WorkflowActivityModel, name, w)
}

func (svc *WorkflowActivityService) GetByField(field string, value string) (*types.WorkflowActivitys, error) {
	w := &types.WorkflowActivitys{}
	return w, svc.client.getByField(types.WorkflowActivityModel, field, value, w)
}

func (svc *WorkflowActivityService) GetAll() (*types.WorkflowActivitys, error) {
	w := &types.WorkflowActivitys{}
	return w, svc.client.getAll(types.WorkflowActivityModel, w)
}

func (svc *WorkflowActivityService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.WorkflowActivityModel, fields, relations)
}

func (svc *WorkflowActivityService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.WorkflowActivityModel, ids, fields, relations)
}

func (svc *WorkflowActivityService) Delete(ids []int64) error {
	return svc.client.delete(types.WorkflowActivityModel, ids)
}
