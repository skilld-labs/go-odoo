package api

import (
	"github.com/morezig/go-odoo/types"
)

type IrQwebFieldQwebService struct {
	client *Client
}

func NewIrQwebFieldQwebService(c *Client) *IrQwebFieldQwebService {
	return &IrQwebFieldQwebService{client: c}
}

func (svc *IrQwebFieldQwebService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrQwebFieldQwebModel, name)
}

func (svc *IrQwebFieldQwebService) GetByIds(ids []int64) (*types.IrQwebFieldQwebs, error) {
	i := &types.IrQwebFieldQwebs{}
	return i, svc.client.getByIds(types.IrQwebFieldQwebModel, ids, i)
}

func (svc *IrQwebFieldQwebService) GetByName(name string) (*types.IrQwebFieldQwebs, error) {
	i := &types.IrQwebFieldQwebs{}
	return i, svc.client.getByName(types.IrQwebFieldQwebModel, name, i)
}

func (svc *IrQwebFieldQwebService) GetByField(field string, value string) (*types.IrQwebFieldQwebs, error) {
	i := &types.IrQwebFieldQwebs{}
	return i, svc.client.getByField(types.IrQwebFieldQwebModel, field, value, i)
}

func (svc *IrQwebFieldQwebService) GetAll() (*types.IrQwebFieldQwebs, error) {
	i := &types.IrQwebFieldQwebs{}
	return i, svc.client.getAll(types.IrQwebFieldQwebModel, i)
}

func (svc *IrQwebFieldQwebService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrQwebFieldQwebModel, fields, relations)
}

func (svc *IrQwebFieldQwebService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebFieldQwebModel, ids, fields, relations)
}

func (svc *IrQwebFieldQwebService) Delete(ids []int64) error {
	return svc.client.delete(types.IrQwebFieldQwebModel, ids)
}
