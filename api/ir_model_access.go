package api

import (
	"github.com/morezig/go-odoo/types"
)

type IrModelAccessService struct {
	client *Client
}

func NewIrModelAccessService(c *Client) *IrModelAccessService {
	return &IrModelAccessService{client: c}
}

func (svc *IrModelAccessService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrModelAccessModel, name)
}

func (svc *IrModelAccessService) GetByIds(ids []int64) (*types.IrModelAccesss, error) {
	i := &types.IrModelAccesss{}
	return i, svc.client.getByIds(types.IrModelAccessModel, ids, i)
}

func (svc *IrModelAccessService) GetByName(name string) (*types.IrModelAccesss, error) {
	i := &types.IrModelAccesss{}
	return i, svc.client.getByName(types.IrModelAccessModel, name, i)
}

func (svc *IrModelAccessService) GetByField(field string, value string) (*types.IrModelAccesss, error) {
	i := &types.IrModelAccesss{}
	return i, svc.client.getByField(types.IrModelAccessModel, field, value, i)
}

func (svc *IrModelAccessService) GetAll() (*types.IrModelAccesss, error) {
	i := &types.IrModelAccesss{}
	return i, svc.client.getAll(types.IrModelAccessModel, i)
}

func (svc *IrModelAccessService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrModelAccessModel, fields, relations)
}

func (svc *IrModelAccessService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrModelAccessModel, ids, fields, relations)
}

func (svc *IrModelAccessService) Delete(ids []int64) error {
	return svc.client.delete(types.IrModelAccessModel, ids)
}
