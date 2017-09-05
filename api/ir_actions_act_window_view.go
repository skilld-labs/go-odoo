package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrActionsActWindowViewService struct {
	client *Client
}

func NewIrActionsActWindowViewService(c *Client) *IrActionsActWindowViewService {
	return &IrActionsActWindowViewService{client: c}
}

func (svc *IrActionsActWindowViewService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrActionsActWindowViewModel, name)
}

func (svc *IrActionsActWindowViewService) GetByIds(ids []int) (*types.IrActionsActWindowViews, error) {
	i := &types.IrActionsActWindowViews{}
	return i, svc.client.getByIds(types.IrActionsActWindowViewModel, ids, i)
}

func (svc *IrActionsActWindowViewService) GetByName(name string) (*types.IrActionsActWindowViews, error) {
	i := &types.IrActionsActWindowViews{}
	return i, svc.client.getByName(types.IrActionsActWindowViewModel, name, i)
}

func (svc *IrActionsActWindowViewService) GetByField(field string, value string) (*types.IrActionsActWindowViews, error) {
	i := &types.IrActionsActWindowViews{}
	return i, svc.client.getByName(types.IrActionsActWindowViewModel, field, value, i)
}

func (svc *IrActionsActWindowViewService) GetAll() (*types.IrActionsActWindowViews, error) {
	i := &types.IrActionsActWindowViews{}
	return i, svc.client.getAll(types.IrActionsActWindowViewModel, i)
}

func (svc *IrActionsActWindowViewService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrActionsActWindowViewModel, fields, relations)
}

func (svc *IrActionsActWindowViewService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrActionsActWindowViewModel, ids, fields, relations)
}

func (svc *IrActionsActWindowViewService) Delete(ids []int) error {
	return svc.client.delete(types.IrActionsActWindowViewModel, ids)
}
