package api

import (
	"github.com/morezig/go-odoo/types"
)

type IrActionsClientService struct {
	client *Client
}

func NewIrActionsClientService(c *Client) *IrActionsClientService {
	return &IrActionsClientService{client: c}
}

func (svc *IrActionsClientService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrActionsClientModel, name)
}

func (svc *IrActionsClientService) GetByIds(ids []int64) (*types.IrActionsClients, error) {
	i := &types.IrActionsClients{}
	return i, svc.client.getByIds(types.IrActionsClientModel, ids, i)
}

func (svc *IrActionsClientService) GetByName(name string) (*types.IrActionsClients, error) {
	i := &types.IrActionsClients{}
	return i, svc.client.getByName(types.IrActionsClientModel, name, i)
}

func (svc *IrActionsClientService) GetByField(field string, value string) (*types.IrActionsClients, error) {
	i := &types.IrActionsClients{}
	return i, svc.client.getByField(types.IrActionsClientModel, field, value, i)
}

func (svc *IrActionsClientService) GetAll() (*types.IrActionsClients, error) {
	i := &types.IrActionsClients{}
	return i, svc.client.getAll(types.IrActionsClientModel, i)
}

func (svc *IrActionsClientService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrActionsClientModel, fields, relations)
}

func (svc *IrActionsClientService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrActionsClientModel, ids, fields, relations)
}

func (svc *IrActionsClientService) Delete(ids []int64) error {
	return svc.client.delete(types.IrActionsClientModel, ids)
}
