package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrFiltersService struct {
	client *Client
}

func NewIrFiltersService(c *Client) *IrFiltersService {
	return &IrFiltersService{client: c}
}

func (svc *IrFiltersService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrFiltersModel, name)
}

func (svc *IrFiltersService) GetByIds(ids []int) (*types.IrFilterss, error) {
	i := &types.IrFilterss{}
	return i, svc.client.getByIds(types.IrFiltersModel, ids, i)
}

func (svc *IrFiltersService) GetByName(name string) (*types.IrFilterss, error) {
	i := &types.IrFilterss{}
	return i, svc.client.getByName(types.IrFiltersModel, name, i)
}

func (svc *IrFiltersService) GetByField(field string, value string) (*types.IrFilterss, error) {
	i := &types.IrFilterss{}
	return i, svc.client.getByField(types.IrFiltersModel, field, value, i)
}

func (svc *IrFiltersService) GetAll() (*types.IrFilterss, error) {
	i := &types.IrFilterss{}
	return i, svc.client.getAll(types.IrFiltersModel, i)
}

func (svc *IrFiltersService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrFiltersModel, fields, relations)
}

func (svc *IrFiltersService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrFiltersModel, ids, fields, relations)
}

func (svc *IrFiltersService) Delete(ids []int) error {
	return svc.client.delete(types.IrFiltersModel, ids)
}
