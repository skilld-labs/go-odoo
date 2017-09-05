package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrQwebFieldDurationService struct {
	client *Client
}

func NewIrQwebFieldDurationService(c *Client) *IrQwebFieldDurationService {
	return &IrQwebFieldDurationService{client: c}
}

func (svc *IrQwebFieldDurationService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrQwebFieldDurationModel, name)
}

func (svc *IrQwebFieldDurationService) GetByIds(ids []int) (*types.IrQwebFieldDurations, error) {
	i := &types.IrQwebFieldDurations{}
	return i, svc.client.getByIds(types.IrQwebFieldDurationModel, ids, i)
}

func (svc *IrQwebFieldDurationService) GetByName(name string) (*types.IrQwebFieldDurations, error) {
	i := &types.IrQwebFieldDurations{}
	return i, svc.client.getByName(types.IrQwebFieldDurationModel, name, i)
}

func (svc *IrQwebFieldDurationService) GetByField(field string, value string) (*types.IrQwebFieldDurations, error) {
	i := &types.IrQwebFieldDurations{}
	return i, svc.client.getByName(types.IrQwebFieldDurationModel, field, value, i)
}

func (svc *IrQwebFieldDurationService) GetAll() (*types.IrQwebFieldDurations, error) {
	i := &types.IrQwebFieldDurations{}
	return i, svc.client.getAll(types.IrQwebFieldDurationModel, i)
}

func (svc *IrQwebFieldDurationService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrQwebFieldDurationModel, fields, relations)
}

func (svc *IrQwebFieldDurationService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebFieldDurationModel, ids, fields, relations)
}

func (svc *IrQwebFieldDurationService) Delete(ids []int) error {
	return svc.client.delete(types.IrQwebFieldDurationModel, ids)
}
