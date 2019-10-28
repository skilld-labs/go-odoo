package api

import (
	"github.com/morezig/go-odoo/types"
)

type IrActionsActWindowCloseService struct {
	client *Client
}

func NewIrActionsActWindowCloseService(c *Client) *IrActionsActWindowCloseService {
	return &IrActionsActWindowCloseService{client: c}
}

func (svc *IrActionsActWindowCloseService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrActionsActWindowCloseModel, name)
}

func (svc *IrActionsActWindowCloseService) GetByIds(ids []int64) (*types.IrActionsActWindowCloses, error) {
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

func (svc *IrActionsActWindowCloseService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrActionsActWindowCloseModel, fields, relations)
}

func (svc *IrActionsActWindowCloseService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrActionsActWindowCloseModel, ids, fields, relations)
}

func (svc *IrActionsActWindowCloseService) Delete(ids []int64) error {
	return svc.client.delete(types.IrActionsActWindowCloseModel, ids)
}
