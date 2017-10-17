package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrDefaultService struct {
	client *Client
}

func NewIrDefaultService(c *Client) *IrDefaultService {
	return &IrDefaultService{client: c}
}

func (svc *IrDefaultService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrDefaultModel, name)
}

func (svc *IrDefaultService) GetByIds(ids []int) (*types.IrDefaults, error) {
	i := &types.IrDefaults{}
	return i, svc.client.getByIds(types.IrDefaultModel, ids, i)
}

func (svc *IrDefaultService) GetByName(name string) (*types.IrDefaults, error) {
	i := &types.IrDefaults{}
	return i, svc.client.getByName(types.IrDefaultModel, name, i)
}

func (svc *IrDefaultService) GetByField(field string, value string) (*types.IrDefaults, error) {
	i := &types.IrDefaults{}
	return i, svc.client.getByField(types.IrDefaultModel, field, value, i)
}

func (svc *IrDefaultService) GetAll() (*types.IrDefaults, error) {
	i := &types.IrDefaults{}
	return i, svc.client.getAll(types.IrDefaultModel, i)
}

func (svc *IrDefaultService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrDefaultModel, fields, relations)
}

func (svc *IrDefaultService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrDefaultModel, ids, fields, relations)
}

func (svc *IrDefaultService) Delete(ids []int) error {
	return svc.client.delete(types.IrDefaultModel, ids)
}
