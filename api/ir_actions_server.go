package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrActionsServerService struct {
	client *Client
}

func NewIrActionsServerService(c *Client) *IrActionsServerService {
	return &IrActionsServerService{client: c}
}

func (svc *IrActionsServerService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrActionsServerModel, name)
}

func (svc *IrActionsServerService) GetByIds(ids []int) (*types.IrActionsServers, error) {
	i := &types.IrActionsServers{}
	return i, svc.client.getByIds(types.IrActionsServerModel, ids, i)
}

func (svc *IrActionsServerService) GetByName(name string) (*types.IrActionsServers, error) {
	i := &types.IrActionsServers{}
	return i, svc.client.getByName(types.IrActionsServerModel, name, i)
}

func (svc *IrActionsServerService) GetAll() (*types.IrActionsServers, error) {
	i := &types.IrActionsServers{}
	return i, svc.client.getAll(types.IrActionsServerModel, i)
}

func (svc *IrActionsServerService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrActionsServerModel, fields, relations)
}

func (svc *IrActionsServerService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrActionsServerModel, ids, fields, relations)
}

func (svc *IrActionsServerService) Delete(ids []int) error {
	return svc.client.delete(types.IrActionsServerModel, ids)
}
