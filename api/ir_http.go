package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrHttpService struct {
	client *Client
}

func NewIrHttpService(c *Client) *IrHttpService {
	return &IrHttpService{client: c}
}

func (svc *IrHttpService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrHttpModel, name)
}

func (svc *IrHttpService) GetByIds(ids []int64) (*types.IrHttps, error) {
	i := &types.IrHttps{}
	return i, svc.client.getByIds(types.IrHttpModel, ids, i)
}

func (svc *IrHttpService) GetByName(name string) (*types.IrHttps, error) {
	i := &types.IrHttps{}
	return i, svc.client.getByName(types.IrHttpModel, name, i)
}

func (svc *IrHttpService) GetByField(field string, value string) (*types.IrHttps, error) {
	i := &types.IrHttps{}
	return i, svc.client.getByField(types.IrHttpModel, field, value, i)
}

func (svc *IrHttpService) GetAll() (*types.IrHttps, error) {
	i := &types.IrHttps{}
	return i, svc.client.getAll(types.IrHttpModel, i)
}

func (svc *IrHttpService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrHttpModel, fields, relations)
}

func (svc *IrHttpService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrHttpModel, ids, fields, relations)
}

func (svc *IrHttpService) Delete(ids []int64) error {
	return svc.client.delete(types.IrHttpModel, ids)
}
