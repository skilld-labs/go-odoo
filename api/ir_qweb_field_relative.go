package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrQwebFieldRelativeService struct {
	client *Client
}

func NewIrQwebFieldRelativeService(c *Client) *IrQwebFieldRelativeService {
	return &IrQwebFieldRelativeService{client: c}
}

func (svc *IrQwebFieldRelativeService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrQwebFieldRelativeModel, name)
}

func (svc *IrQwebFieldRelativeService) GetByIds(ids []int) (*types.IrQwebFieldRelatives, error) {
	i := &types.IrQwebFieldRelatives{}
	return i, svc.client.getByIds(types.IrQwebFieldRelativeModel, ids, i)
}

func (svc *IrQwebFieldRelativeService) GetByName(name string) (*types.IrQwebFieldRelatives, error) {
	i := &types.IrQwebFieldRelatives{}
	return i, svc.client.getByName(types.IrQwebFieldRelativeModel, name, i)
}

func (svc *IrQwebFieldRelativeService) GetAll() (*types.IrQwebFieldRelatives, error) {
	i := &types.IrQwebFieldRelatives{}
	return i, svc.client.getAll(types.IrQwebFieldRelativeModel, i)
}

func (svc *IrQwebFieldRelativeService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrQwebFieldRelativeModel, fields, relations)
}

func (svc *IrQwebFieldRelativeService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebFieldRelativeModel, ids, fields, relations)
}

func (svc *IrQwebFieldRelativeService) Delete(ids []int) error {
	return svc.client.delete(types.IrQwebFieldRelativeModel, ids)
}
