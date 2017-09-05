package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrQwebService struct {
	client *Client
}

func NewIrQwebService(c *Client) *IrQwebService {
	return &IrQwebService{client: c}
}

func (svc *IrQwebService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrQwebModel, name)
}

func (svc *IrQwebService) GetByIds(ids []int) (*types.IrQwebs, error) {
	i := &types.IrQwebs{}
	return i, svc.client.getByIds(types.IrQwebModel, ids, i)
}

func (svc *IrQwebService) GetByName(name string) (*types.IrQwebs, error) {
	i := &types.IrQwebs{}
	return i, svc.client.getByName(types.IrQwebModel, name, i)
}

func (svc *IrQwebService) GetByField(field string, value string) (*types.IrQwebs, error) {
	i := &types.IrQwebs{}
	return i, svc.client.getByName(types.IrQwebModel, field, value, i)
}

func (svc *IrQwebService) GetAll() (*types.IrQwebs, error) {
	i := &types.IrQwebs{}
	return i, svc.client.getAll(types.IrQwebModel, i)
}

func (svc *IrQwebService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrQwebModel, fields, relations)
}

func (svc *IrQwebService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebModel, ids, fields, relations)
}

func (svc *IrQwebService) Delete(ids []int) error {
	return svc.client.delete(types.IrQwebModel, ids)
}
