package api

import (
	"github.com/morezig/go-odoo/types"
)

type IrQwebFieldService struct {
	client *Client
}

func NewIrQwebFieldService(c *Client) *IrQwebFieldService {
	return &IrQwebFieldService{client: c}
}

func (svc *IrQwebFieldService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrQwebFieldModel, name)
}

func (svc *IrQwebFieldService) GetByIds(ids []int64) (*types.IrQwebFields, error) {
	i := &types.IrQwebFields{}
	return i, svc.client.getByIds(types.IrQwebFieldModel, ids, i)
}

func (svc *IrQwebFieldService) GetByName(name string) (*types.IrQwebFields, error) {
	i := &types.IrQwebFields{}
	return i, svc.client.getByName(types.IrQwebFieldModel, name, i)
}

func (svc *IrQwebFieldService) GetByField(field string, value string) (*types.IrQwebFields, error) {
	i := &types.IrQwebFields{}
	return i, svc.client.getByField(types.IrQwebFieldModel, field, value, i)
}

func (svc *IrQwebFieldService) GetAll() (*types.IrQwebFields, error) {
	i := &types.IrQwebFields{}
	return i, svc.client.getAll(types.IrQwebFieldModel, i)
}

func (svc *IrQwebFieldService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrQwebFieldModel, fields, relations)
}

func (svc *IrQwebFieldService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebFieldModel, ids, fields, relations)
}

func (svc *IrQwebFieldService) Delete(ids []int64) error {
	return svc.client.delete(types.IrQwebFieldModel, ids)
}
