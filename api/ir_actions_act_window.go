package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrActionsActWindowService struct {
	client *Client
}

func NewIrActionsActWindowService(c *Client) *IrActionsActWindowService {
	return &IrActionsActWindowService{client: c}
}

func (svc *IrActionsActWindowService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrActionsActWindowModel, name)
}

func (svc *IrActionsActWindowService) GetByIds(ids []int64) (*types.IrActionsActWindows, error) {
	i := &types.IrActionsActWindows{}
	return i, svc.client.getByIds(types.IrActionsActWindowModel, ids, i)
}

func (svc *IrActionsActWindowService) GetByName(name string) (*types.IrActionsActWindows, error) {
	i := &types.IrActionsActWindows{}
	return i, svc.client.getByName(types.IrActionsActWindowModel, name, i)
}

func (svc *IrActionsActWindowService) GetByField(field string, value string) (*types.IrActionsActWindows, error) {
	i := &types.IrActionsActWindows{}
	return i, svc.client.getByField(types.IrActionsActWindowModel, field, value, i)
}

func (svc *IrActionsActWindowService) GetAll() (*types.IrActionsActWindows, error) {
	i := &types.IrActionsActWindows{}
	return i, svc.client.getAll(types.IrActionsActWindowModel, i)
}

func (svc *IrActionsActWindowService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrActionsActWindowModel, fields, relations)
}

func (svc *IrActionsActWindowService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrActionsActWindowModel, ids, fields, relations)
}

func (svc *IrActionsActWindowService) Delete(ids []int64) error {
	return svc.client.delete(types.IrActionsActWindowModel, ids)
}
