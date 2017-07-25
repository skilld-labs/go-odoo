package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type WorkflowService struct {
	client *Client
}

func NewWorkflowService(c *Client) *WorkflowService {
	return &WorkflowService{client: c}
}

func (svc *WorkflowService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.WorkflowModel, name)
}

func (svc *WorkflowService) GetByIds(ids []int) (*types.Workflows, error) {
	w := &types.Workflows{}
	return w, svc.client.getByIds(types.WorkflowModel, ids, w)
}

func (svc *WorkflowService) GetByName(name string) (*types.Workflows, error) {
	w := &types.Workflows{}
	return w, svc.client.getByName(types.WorkflowModel, name, w)
}

func (svc *WorkflowService) GetAll() (*types.Workflows, error) {
	w := &types.Workflows{}
	return w, svc.client.getAll(types.WorkflowModel, w)
}

func (svc *WorkflowService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.WorkflowModel, fields, relations)
}

func (svc *WorkflowService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.WorkflowModel, ids, fields, relations)
}

func (svc *WorkflowService) Delete(ids []int) error {
	return svc.client.delete(types.WorkflowModel, ids)
}
