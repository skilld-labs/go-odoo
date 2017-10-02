package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrModelConstraintService struct {
	client *Client
}

func NewIrModelConstraintService(c *Client) *IrModelConstraintService {
	return &IrModelConstraintService{client: c}
}

func (svc *IrModelConstraintService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrModelConstraintModel, name)
}

func (svc *IrModelConstraintService) GetByIds(ids []int64) (*types.IrModelConstraints, error) {
	i := &types.IrModelConstraints{}
	return i, svc.client.getByIds(types.IrModelConstraintModel, ids, i)
}

func (svc *IrModelConstraintService) GetByName(name string) (*types.IrModelConstraints, error) {
	i := &types.IrModelConstraints{}
	return i, svc.client.getByName(types.IrModelConstraintModel, name, i)
}

func (svc *IrModelConstraintService) GetByField(field string, value string) (*types.IrModelConstraints, error) {
	i := &types.IrModelConstraints{}
	return i, svc.client.getByField(types.IrModelConstraintModel, field, value, i)
}

func (svc *IrModelConstraintService) GetAll() (*types.IrModelConstraints, error) {
	i := &types.IrModelConstraints{}
	return i, svc.client.getAll(types.IrModelConstraintModel, i)
}

func (svc *IrModelConstraintService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrModelConstraintModel, fields, relations)
}

func (svc *IrModelConstraintService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrModelConstraintModel, ids, fields, relations)
}

func (svc *IrModelConstraintService) Delete(ids []int64) error {
	return svc.client.delete(types.IrModelConstraintModel, ids)
}
