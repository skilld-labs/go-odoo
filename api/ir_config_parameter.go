package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrConfigParameterService struct {
	client *Client
}

func NewIrConfigParameterService(c *Client) *IrConfigParameterService {
	return &IrConfigParameterService{client: c}
}

func (svc *IrConfigParameterService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrConfigParameterModel, name)
}

func (svc *IrConfigParameterService) GetByIds(ids []int) (*types.IrConfigParameters, error) {
	i := &types.IrConfigParameters{}
	return i, svc.client.getByIds(types.IrConfigParameterModel, ids, i)
}

func (svc *IrConfigParameterService) GetByName(name string) (*types.IrConfigParameters, error) {
	i := &types.IrConfigParameters{}
	return i, svc.client.getByName(types.IrConfigParameterModel, name, i)
}

func (svc *IrConfigParameterService) GetByField(field string, value string) (*types.IrConfigParameters, error) {
	i := &types.IrConfigParameters{}
	return i, svc.client.getByField(types.IrConfigParameterModel, field, value, i)
}

func (svc *IrConfigParameterService) GetAll() (*types.IrConfigParameters, error) {
	i := &types.IrConfigParameters{}
	return i, svc.client.getAll(types.IrConfigParameterModel, i)
}

func (svc *IrConfigParameterService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrConfigParameterModel, fields, relations)
}

func (svc *IrConfigParameterService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrConfigParameterModel, ids, fields, relations)
}

func (svc *IrConfigParameterService) Delete(ids []int) error {
	return svc.client.delete(types.IrConfigParameterModel, ids)
}
