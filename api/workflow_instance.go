package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type WorkflowInstanceService struct {
	client *Client
}

func NewWorkflowInstanceService(c *Client) *WorkflowInstanceService {
	return &WorkflowInstanceService{client: c}
}

func (svc *WorkflowInstanceService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.WorkflowInstanceModel, name)
}

func (svc *WorkflowInstanceService) GetByIds(ids []int) (*types.WorkflowInstances, error) {
	w := &types.WorkflowInstances{}
	return w, svc.client.getByIds(types.WorkflowInstanceModel, ids, w)
}

func (svc *WorkflowInstanceService) GetByName(name string) (*types.WorkflowInstances, error) {
	w := &types.WorkflowInstances{}
	return w, svc.client.getByName(types.WorkflowInstanceModel, name, w)
}

func (svc *WorkflowInstanceService) GetAll() (*types.WorkflowInstances, error) {
	w := &types.WorkflowInstances{}
	return w, svc.client.getAll(types.WorkflowInstanceModel, w)
}

func (svc *WorkflowInstanceService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.WorkflowInstanceModel, fields, relations)
}

func (svc *WorkflowInstanceService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.WorkflowInstanceModel, ids, fields, relations)
}

func (svc *WorkflowInstanceService) Delete(ids []int) error {
	return svc.client.delete(types.WorkflowInstanceModel, ids)
}
