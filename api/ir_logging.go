package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrLoggingService struct {
	client *Client
}

func NewIrLoggingService(c *Client) *IrLoggingService {
	return &IrLoggingService{client: c}
}

func (svc *IrLoggingService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrLoggingModel, name)
}

func (svc *IrLoggingService) GetByIds(ids []int) (*types.IrLoggings, error) {
	i := &types.IrLoggings{}
	return i, svc.client.getByIds(types.IrLoggingModel, ids, i)
}

func (svc *IrLoggingService) GetByName(name string) (*types.IrLoggings, error) {
	i := &types.IrLoggings{}
	return i, svc.client.getByName(types.IrLoggingModel, name, i)
}

func (svc *IrLoggingService) GetByField(field string, value string) (*types.IrLoggings, error) {
	i := &types.IrLoggings{}
	return i, svc.client.getByName(types.IrLoggingModel, field, value, i)
}

func (svc *IrLoggingService) GetAll() (*types.IrLoggings, error) {
	i := &types.IrLoggings{}
	return i, svc.client.getAll(types.IrLoggingModel, i)
}

func (svc *IrLoggingService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrLoggingModel, fields, relations)
}

func (svc *IrLoggingService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrLoggingModel, ids, fields, relations)
}

func (svc *IrLoggingService) Delete(ids []int) error {
	return svc.client.delete(types.IrLoggingModel, ids)
}
