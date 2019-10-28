package api

import (
	"github.com/morezig/go-odoo/types"
)

type IrActionsTodoService struct {
	client *Client
}

func NewIrActionsTodoService(c *Client) *IrActionsTodoService {
	return &IrActionsTodoService{client: c}
}

func (svc *IrActionsTodoService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrActionsTodoModel, name)
}

func (svc *IrActionsTodoService) GetByIds(ids []int64) (*types.IrActionsTodos, error) {
	i := &types.IrActionsTodos{}
	return i, svc.client.getByIds(types.IrActionsTodoModel, ids, i)
}

func (svc *IrActionsTodoService) GetByName(name string) (*types.IrActionsTodos, error) {
	i := &types.IrActionsTodos{}
	return i, svc.client.getByName(types.IrActionsTodoModel, name, i)
}

func (svc *IrActionsTodoService) GetByField(field string, value string) (*types.IrActionsTodos, error) {
	i := &types.IrActionsTodos{}
	return i, svc.client.getByField(types.IrActionsTodoModel, field, value, i)
}

func (svc *IrActionsTodoService) GetAll() (*types.IrActionsTodos, error) {
	i := &types.IrActionsTodos{}
	return i, svc.client.getAll(types.IrActionsTodoModel, i)
}

func (svc *IrActionsTodoService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrActionsTodoModel, fields, relations)
}

func (svc *IrActionsTodoService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrActionsTodoModel, ids, fields, relations)
}

func (svc *IrActionsTodoService) Delete(ids []int64) error {
	return svc.client.delete(types.IrActionsTodoModel, ids)
}
