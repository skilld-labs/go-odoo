package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrQwebFieldIntegerService struct {
	client *Client
}

func NewIrQwebFieldIntegerService(c *Client) *IrQwebFieldIntegerService {
	return &IrQwebFieldIntegerService{client: c}
}

func (svc *IrQwebFieldIntegerService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrQwebFieldIntegerModel, name)
}

func (svc *IrQwebFieldIntegerService) GetByIds(ids []int) (*types.IrQwebFieldIntegers, error) {
	i := &types.IrQwebFieldIntegers{}
	return i, svc.client.getByIds(types.IrQwebFieldIntegerModel, ids, i)
}

func (svc *IrQwebFieldIntegerService) GetByName(name string) (*types.IrQwebFieldIntegers, error) {
	i := &types.IrQwebFieldIntegers{}
	return i, svc.client.getByName(types.IrQwebFieldIntegerModel, name, i)
}

func (svc *IrQwebFieldIntegerService) GetByField(field string, value string) (*types.IrQwebFieldIntegers, error) {
	i := &types.IrQwebFieldIntegers{}
	return i, svc.client.getByField(types.IrQwebFieldIntegerModel, field, value, i)
}

func (svc *IrQwebFieldIntegerService) GetAll() (*types.IrQwebFieldIntegers, error) {
	i := &types.IrQwebFieldIntegers{}
	return i, svc.client.getAll(types.IrQwebFieldIntegerModel, i)
}

func (svc *IrQwebFieldIntegerService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrQwebFieldIntegerModel, fields, relations)
}

func (svc *IrQwebFieldIntegerService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebFieldIntegerModel, ids, fields, relations)
}

func (svc *IrQwebFieldIntegerService) Delete(ids []int) error {
	return svc.client.delete(types.IrQwebFieldIntegerModel, ids)
}
