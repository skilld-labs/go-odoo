package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrValuesService struct {
	client *Client
}

func NewIrValuesService(c *Client) *IrValuesService {
	return &IrValuesService{client: c}
}

func (svc *IrValuesService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrValuesModel, name)
}

func (svc *IrValuesService) GetByIds(ids []int) (*types.IrValuess, error) {
	i := &types.IrValuess{}
	return i, svc.client.getByIds(types.IrValuesModel, ids, i)
}

func (svc *IrValuesService) GetByName(name string) (*types.IrValuess, error) {
	i := &types.IrValuess{}
	return i, svc.client.getByName(types.IrValuesModel, name, i)
}

func (svc *IrValuesService) GetByField(field string, value string) (*types.IrValuess, error) {
	i := &types.IrValuess{}
	return i, svc.client.getByField(types.IrValuesModel, field, value, i)
}

func (svc *IrValuesService) GetAll() (*types.IrValuess, error) {
	i := &types.IrValuess{}
	return i, svc.client.getAll(types.IrValuesModel, i)
}

func (svc *IrValuesService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrValuesModel, fields, relations)
}

func (svc *IrValuesService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrValuesModel, ids, fields, relations)
}

func (svc *IrValuesService) Delete(ids []int) error {
	return svc.client.delete(types.IrValuesModel, ids)
}
