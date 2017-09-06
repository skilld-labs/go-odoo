package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrActionsActWindowCloseService struct {
	client *Client
}

func NewIrActionsActWindowCloseService(c *Client) *IrActionsActWindowCloseService {
	return &IrActionsActWindowCloseService{client: c}
}

func (svc *IrActionsActWindowCloseService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrActionsActWindowCloseModel, name)
}

func (svc *IrActionsActWindowCloseService) GetByIds(ids []int) (*types.IrActionsActWindowCloses, error) {
	i := &types.IrActionsActWindowCloses{}
	return i, svc.client.getByIds(types.IrActionsActWindowCloseModel, ids, i)
}

func (svc *IrActionsActWindowCloseService) GetByName(name string) (*types.IrActionsActWindowCloses, error) {
	i := &types.IrActionsActWindowCloses{}
	return i, svc.client.getByName(types.IrActionsActWindowCloseModel, name, i)
}

func (svc *IrActionsActWindowCloseService) GetByField(field string, value string) (*types.IrActionsActWindowCloses, error) {
	i := &types.IrActionsActWindowCloses{}
	return i, svc.client.getByField(types.IrActionsActWindowCloseModel, field, value, i)
}

func (svc *IrActionsActWindowCloseService) GetAll() (*types.IrActionsActWindowCloses, error) {
	i := &types.IrActionsActWindowCloses{}
	return i, svc.client.getAll(types.IrActionsActWindowCloseModel, i)
}

func (svc *IrActionsActWindowCloseService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrActionsActWindowCloseModel, fields, relations)
}

func (svc *IrActionsActWindowCloseService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrActionsActWindowCloseModel, ids, fields, relations)
}

func (svc *IrActionsActWindowCloseService) Delete(ids []int) error {
	return svc.client.delete(types.IrActionsActWindowCloseModel, ids)
}
