package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrQwebFieldDatetimeService struct {
	client *Client
}

func NewIrQwebFieldDatetimeService(c *Client) *IrQwebFieldDatetimeService {
	return &IrQwebFieldDatetimeService{client: c}
}

func (svc *IrQwebFieldDatetimeService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrQwebFieldDatetimeModel, name)
}

func (svc *IrQwebFieldDatetimeService) GetByIds(ids []int) (*types.IrQwebFieldDatetimes, error) {
	i := &types.IrQwebFieldDatetimes{}
	return i, svc.client.getByIds(types.IrQwebFieldDatetimeModel, ids, i)
}

func (svc *IrQwebFieldDatetimeService) GetByName(name string) (*types.IrQwebFieldDatetimes, error) {
	i := &types.IrQwebFieldDatetimes{}
	return i, svc.client.getByName(types.IrQwebFieldDatetimeModel, name, i)
}

func (svc *IrQwebFieldDatetimeService) GetAll() (*types.IrQwebFieldDatetimes, error) {
	i := &types.IrQwebFieldDatetimes{}
	return i, svc.client.getAll(types.IrQwebFieldDatetimeModel, i)
}

func (svc *IrQwebFieldDatetimeService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrQwebFieldDatetimeModel, fields, relations)
}

func (svc *IrQwebFieldDatetimeService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebFieldDatetimeModel, ids, fields, relations)
}

func (svc *IrQwebFieldDatetimeService) Delete(ids []int) error {
	return svc.client.delete(types.IrQwebFieldDatetimeModel, ids)
}
