package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrQwebFieldDateService struct {
	client *Client
}

func NewIrQwebFieldDateService(c *Client) *IrQwebFieldDateService {
	return &IrQwebFieldDateService{client: c}
}

func (svc *IrQwebFieldDateService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrQwebFieldDateModel, name)
}

func (svc *IrQwebFieldDateService) GetByIds(ids []int) (*types.IrQwebFieldDates, error) {
	i := &types.IrQwebFieldDates{}
	return i, svc.client.getByIds(types.IrQwebFieldDateModel, ids, i)
}

func (svc *IrQwebFieldDateService) GetByName(name string) (*types.IrQwebFieldDates, error) {
	i := &types.IrQwebFieldDates{}
	return i, svc.client.getByName(types.IrQwebFieldDateModel, name, i)
}

func (svc *IrQwebFieldDateService) GetByField(field string, value string) (*types.IrQwebFieldDates, error) {
	i := &types.IrQwebFieldDates{}
	return i, svc.client.getByField(types.IrQwebFieldDateModel, field, value, i)
}

func (svc *IrQwebFieldDateService) GetAll() (*types.IrQwebFieldDates, error) {
	i := &types.IrQwebFieldDates{}
	return i, svc.client.getAll(types.IrQwebFieldDateModel, i)
}

func (svc *IrQwebFieldDateService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrQwebFieldDateModel, fields, relations)
}

func (svc *IrQwebFieldDateService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebFieldDateModel, ids, fields, relations)
}

func (svc *IrQwebFieldDateService) Delete(ids []int) error {
	return svc.client.delete(types.IrQwebFieldDateModel, ids)
}
