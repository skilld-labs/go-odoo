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

func (svc *IrModelConstraintService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrModelConstraintModel, name)
}

func (svc *IrModelConstraintService) GetByIds(ids []int) (*types.IrModelConstraints, error) {
	i := &types.IrModelConstraints{}
	return i, svc.client.getByIds(types.IrModelConstraintModel, ids, i)
}

func (svc *IrModelConstraintService) GetByName(name string) (*types.IrModelConstraints, error) {
	i := &types.IrModelConstraints{}
	return i, svc.client.getByName(types.IrModelConstraintModel, name, i)
}

func (svc *IrModelConstraintService) GetAll() (*types.IrModelConstraints, error) {
	i := &types.IrModelConstraints{}
	return i, svc.client.getAll(types.IrModelConstraintModel, i)
}

func (svc *IrModelConstraintService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrModelConstraintModel, fields, relations)
}

func (svc *IrModelConstraintService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrModelConstraintModel, ids, fields, relations)
}

func (svc *IrModelConstraintService) Delete(ids []int) error {
	return svc.client.delete(types.IrModelConstraintModel, ids)
}
