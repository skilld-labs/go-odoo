package api

import (
	"github.com/morezig/go-odoo/types"
)

type IrActionsActionsService struct {
	client *Client
}

func NewIrActionsActionsService(c *Client) *IrActionsActionsService {
	return &IrActionsActionsService{client: c}
}

func (svc *IrActionsActionsService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrActionsActionsModel, name)
}

func (svc *IrActionsActionsService) GetByIds(ids []int64) (*types.IrActionsActionss, error) {
	i := &types.IrActionsActionss{}
	return i, svc.client.getByIds(types.IrActionsActionsModel, ids, i)
}

func (svc *IrActionsActionsService) GetByName(name string) (*types.IrActionsActionss, error) {
	i := &types.IrActionsActionss{}
	return i, svc.client.getByName(types.IrActionsActionsModel, name, i)
}

func (svc *IrActionsActionsService) GetByField(field string, value string) (*types.IrActionsActionss, error) {
	i := &types.IrActionsActionss{}
	return i, svc.client.getByField(types.IrActionsActionsModel, field, value, i)
}

func (svc *IrActionsActionsService) GetAll() (*types.IrActionsActionss, error) {
	i := &types.IrActionsActionss{}
	return i, svc.client.getAll(types.IrActionsActionsModel, i)
}

func (svc *IrActionsActionsService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrActionsActionsModel, fields, relations)
}

func (svc *IrActionsActionsService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrActionsActionsModel, ids, fields, relations)
}

func (svc *IrActionsActionsService) Delete(ids []int64) error {
	return svc.client.delete(types.IrActionsActionsModel, ids)
}
