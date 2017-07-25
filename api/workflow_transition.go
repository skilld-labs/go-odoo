package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type WorkflowTransitionService struct {
	client *Client
}

func NewWorkflowTransitionService(c *Client) *WorkflowTransitionService {
	return &WorkflowTransitionService{client: c}
}

func (svc *WorkflowTransitionService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.WorkflowTransitionModel, name)
}

func (svc *WorkflowTransitionService) GetByIds(ids []int) (*types.WorkflowTransitions, error) {
	w := &types.WorkflowTransitions{}
	return w, svc.client.getByIds(types.WorkflowTransitionModel, ids, w)
}

func (svc *WorkflowTransitionService) GetByName(name string) (*types.WorkflowTransitions, error) {
	w := &types.WorkflowTransitions{}
	return w, svc.client.getByName(types.WorkflowTransitionModel, name, w)
}

func (svc *WorkflowTransitionService) GetAll() (*types.WorkflowTransitions, error) {
	w := &types.WorkflowTransitions{}
	return w, svc.client.getAll(types.WorkflowTransitionModel, w)
}

func (svc *WorkflowTransitionService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.WorkflowTransitionModel, fields, relations)
}

func (svc *WorkflowTransitionService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.WorkflowTransitionModel, ids, fields, relations)
}

func (svc *WorkflowTransitionService) Delete(ids []int) error {
	return svc.client.delete(types.WorkflowTransitionModel, ids)
}
