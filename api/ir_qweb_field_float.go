package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrQwebFieldFloatService struct {
	client *Client
}

func NewIrQwebFieldFloatService(c *Client) *IrQwebFieldFloatService {
	return &IrQwebFieldFloatService{client: c}
}

func (svc *IrQwebFieldFloatService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrQwebFieldFloatModel, name)
}

func (svc *IrQwebFieldFloatService) GetByIds(ids []int64) (*types.IrQwebFieldFloats, error) {
	i := &types.IrQwebFieldFloats{}
	return i, svc.client.getByIds(types.IrQwebFieldFloatModel, ids, i)
}

func (svc *IrQwebFieldFloatService) GetByName(name string) (*types.IrQwebFieldFloats, error) {
	i := &types.IrQwebFieldFloats{}
	return i, svc.client.getByName(types.IrQwebFieldFloatModel, name, i)
}

func (svc *IrQwebFieldFloatService) GetByField(field string, value string) (*types.IrQwebFieldFloats, error) {
	i := &types.IrQwebFieldFloats{}
	return i, svc.client.getByField(types.IrQwebFieldFloatModel, field, value, i)
}

func (svc *IrQwebFieldFloatService) GetAll() (*types.IrQwebFieldFloats, error) {
	i := &types.IrQwebFieldFloats{}
	return i, svc.client.getAll(types.IrQwebFieldFloatModel, i)
}

func (svc *IrQwebFieldFloatService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrQwebFieldFloatModel, fields, relations)
}

func (svc *IrQwebFieldFloatService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebFieldFloatModel, ids, fields, relations)
}

func (svc *IrQwebFieldFloatService) Delete(ids []int64) error {
	return svc.client.delete(types.IrQwebFieldFloatModel, ids)
}
