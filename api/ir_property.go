package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrPropertyService struct {
	client *Client
}

func NewIrPropertyService(c *Client) *IrPropertyService {
	return &IrPropertyService{client: c}
}

func (svc *IrPropertyService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrPropertyModel, name)
}

func (svc *IrPropertyService) GetByIds(ids []int64) (*types.IrPropertys, error) {
	i := &types.IrPropertys{}
	return i, svc.client.getByIds(types.IrPropertyModel, ids, i)
}

func (svc *IrPropertyService) GetByName(name string) (*types.IrPropertys, error) {
	i := &types.IrPropertys{}
	return i, svc.client.getByName(types.IrPropertyModel, name, i)
}

func (svc *IrPropertyService) GetByField(field string, value string) (*types.IrPropertys, error) {
	i := &types.IrPropertys{}
	return i, svc.client.getByField(types.IrPropertyModel, field, value, i)
}

func (svc *IrPropertyService) GetAll() (*types.IrPropertys, error) {
	i := &types.IrPropertys{}
	return i, svc.client.getAll(types.IrPropertyModel, i)
}

func (svc *IrPropertyService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrPropertyModel, fields, relations)
}

func (svc *IrPropertyService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrPropertyModel, ids, fields, relations)
}

func (svc *IrPropertyService) Delete(ids []int64) error {
	return svc.client.delete(types.IrPropertyModel, ids)
}
