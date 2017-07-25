package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type WorkflowTriggersService struct {
	client *Client
}

func NewWorkflowTriggersService(c *Client) *WorkflowTriggersService {
	return &WorkflowTriggersService{client: c}
}

func (svc *WorkflowTriggersService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.WorkflowTriggersModel, name)
}

func (svc *WorkflowTriggersService) GetByIds(ids []int) (*types.WorkflowTriggerss, error) {
	w := &types.WorkflowTriggerss{}
	return w, svc.client.getByIds(types.WorkflowTriggersModel, ids, w)
}

func (svc *WorkflowTriggersService) GetByName(name string) (*types.WorkflowTriggerss, error) {
	w := &types.WorkflowTriggerss{}
	return w, svc.client.getByName(types.WorkflowTriggersModel, name, w)
}

func (svc *WorkflowTriggersService) GetAll() (*types.WorkflowTriggerss, error) {
	w := &types.WorkflowTriggerss{}
	return w, svc.client.getAll(types.WorkflowTriggersModel, w)
}

func (svc *WorkflowTriggersService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.WorkflowTriggersModel, fields, relations)
}

func (svc *WorkflowTriggersService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.WorkflowTriggersModel, ids, fields, relations)
}

func (svc *WorkflowTriggersService) Delete(ids []int) error {
	return svc.client.delete(types.WorkflowTriggersModel, ids)
}
