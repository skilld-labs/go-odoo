package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type WorkflowWorkitemService struct {
	client *Client
}

func NewWorkflowWorkitemService(c *Client) *WorkflowWorkitemService {
	return &WorkflowWorkitemService{client: c}
}

func (svc *WorkflowWorkitemService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.WorkflowWorkitemModel, name)
}

func (svc *WorkflowWorkitemService) GetByIds(ids []int64) (*types.WorkflowWorkitems, error) {
	w := &types.WorkflowWorkitems{}
	return w, svc.client.getByIds(types.WorkflowWorkitemModel, ids, w)
}

func (svc *WorkflowWorkitemService) GetByName(name string) (*types.WorkflowWorkitems, error) {
	w := &types.WorkflowWorkitems{}
	return w, svc.client.getByName(types.WorkflowWorkitemModel, name, w)
}

func (svc *WorkflowWorkitemService) GetByField(field string, value string) (*types.WorkflowWorkitems, error) {
	w := &types.WorkflowWorkitems{}
	return w, svc.client.getByField(types.WorkflowWorkitemModel, field, value, w)
}

func (svc *WorkflowWorkitemService) GetAll() (*types.WorkflowWorkitems, error) {
	w := &types.WorkflowWorkitems{}
	return w, svc.client.getAll(types.WorkflowWorkitemModel, w)
}

func (svc *WorkflowWorkitemService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.WorkflowWorkitemModel, fields, relations)
}

func (svc *WorkflowWorkitemService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.WorkflowWorkitemModel, ids, fields, relations)
}

func (svc *WorkflowWorkitemService) Delete(ids []int64) error {
	return svc.client.delete(types.WorkflowWorkitemModel, ids)
}
