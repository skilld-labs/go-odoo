package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrFieldsConverterService struct {
	client *Client
}

func NewIrFieldsConverterService(c *Client) *IrFieldsConverterService {
	return &IrFieldsConverterService{client: c}
}

func (svc *IrFieldsConverterService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrFieldsConverterModel, name)
}

func (svc *IrFieldsConverterService) GetByIds(ids []int) (*types.IrFieldsConverters, error) {
	i := &types.IrFieldsConverters{}
	return i, svc.client.getByIds(types.IrFieldsConverterModel, ids, i)
}

func (svc *IrFieldsConverterService) GetByName(name string) (*types.IrFieldsConverters, error) {
	i := &types.IrFieldsConverters{}
	return i, svc.client.getByName(types.IrFieldsConverterModel, name, i)
}

func (svc *IrFieldsConverterService) GetByField(field string, value string) (*types.IrFieldsConverters, error) {
	i := &types.IrFieldsConverters{}
	return i, svc.client.getByName(types.IrFieldsConverterModel, field, value, i)
}

func (svc *IrFieldsConverterService) GetAll() (*types.IrFieldsConverters, error) {
	i := &types.IrFieldsConverters{}
	return i, svc.client.getAll(types.IrFieldsConverterModel, i)
}

func (svc *IrFieldsConverterService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrFieldsConverterModel, fields, relations)
}

func (svc *IrFieldsConverterService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrFieldsConverterModel, ids, fields, relations)
}

func (svc *IrFieldsConverterService) Delete(ids []int) error {
	return svc.client.delete(types.IrFieldsConverterModel, ids)
}
