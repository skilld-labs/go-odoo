package api

import (
	"github.com/morezig/go-odoo/types"
)

type IrQwebFieldDatetimeService struct {
	client *Client
}

func NewIrQwebFieldDatetimeService(c *Client) *IrQwebFieldDatetimeService {
	return &IrQwebFieldDatetimeService{client: c}
}

func (svc *IrQwebFieldDatetimeService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrQwebFieldDatetimeModel, name)
}

func (svc *IrQwebFieldDatetimeService) GetByIds(ids []int64) (*types.IrQwebFieldDatetimes, error) {
	i := &types.IrQwebFieldDatetimes{}
	return i, svc.client.getByIds(types.IrQwebFieldDatetimeModel, ids, i)
}

func (svc *IrQwebFieldDatetimeService) GetByName(name string) (*types.IrQwebFieldDatetimes, error) {
	i := &types.IrQwebFieldDatetimes{}
	return i, svc.client.getByName(types.IrQwebFieldDatetimeModel, name, i)
}

func (svc *IrQwebFieldDatetimeService) GetByField(field string, value string) (*types.IrQwebFieldDatetimes, error) {
	i := &types.IrQwebFieldDatetimes{}
	return i, svc.client.getByField(types.IrQwebFieldDatetimeModel, field, value, i)
}

func (svc *IrQwebFieldDatetimeService) GetAll() (*types.IrQwebFieldDatetimes, error) {
	i := &types.IrQwebFieldDatetimes{}
	return i, svc.client.getAll(types.IrQwebFieldDatetimeModel, i)
}

func (svc *IrQwebFieldDatetimeService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrQwebFieldDatetimeModel, fields, relations)
}

func (svc *IrQwebFieldDatetimeService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebFieldDatetimeModel, ids, fields, relations)
}

func (svc *IrQwebFieldDatetimeService) Delete(ids []int64) error {
	return svc.client.delete(types.IrQwebFieldDatetimeModel, ids)
}
