package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrQwebFieldFloatTimeService struct {
	client *Client
}

func NewIrQwebFieldFloatTimeService(c *Client) *IrQwebFieldFloatTimeService {
	return &IrQwebFieldFloatTimeService{client: c}
}

func (svc *IrQwebFieldFloatTimeService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrQwebFieldFloatTimeModel, name)
}

func (svc *IrQwebFieldFloatTimeService) GetByIds(ids []int) (*types.IrQwebFieldFloatTimes, error) {
	i := &types.IrQwebFieldFloatTimes{}
	return i, svc.client.getByIds(types.IrQwebFieldFloatTimeModel, ids, i)
}

func (svc *IrQwebFieldFloatTimeService) GetByName(name string) (*types.IrQwebFieldFloatTimes, error) {
	i := &types.IrQwebFieldFloatTimes{}
	return i, svc.client.getByName(types.IrQwebFieldFloatTimeModel, name, i)
}

func (svc *IrQwebFieldFloatTimeService) GetByField(field string, value string) (*types.IrQwebFieldFloatTimes, error) {
	i := &types.IrQwebFieldFloatTimes{}
	return i, svc.client.getByField(types.IrQwebFieldFloatTimeModel, field, value, i)
}

func (svc *IrQwebFieldFloatTimeService) GetAll() (*types.IrQwebFieldFloatTimes, error) {
	i := &types.IrQwebFieldFloatTimes{}
	return i, svc.client.getAll(types.IrQwebFieldFloatTimeModel, i)
}

func (svc *IrQwebFieldFloatTimeService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrQwebFieldFloatTimeModel, fields, relations)
}

func (svc *IrQwebFieldFloatTimeService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebFieldFloatTimeModel, ids, fields, relations)
}

func (svc *IrQwebFieldFloatTimeService) Delete(ids []int) error {
	return svc.client.delete(types.IrQwebFieldFloatTimeModel, ids)
}
