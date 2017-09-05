package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrUiViewCustomService struct {
	client *Client
}

func NewIrUiViewCustomService(c *Client) *IrUiViewCustomService {
	return &IrUiViewCustomService{client: c}
}

func (svc *IrUiViewCustomService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrUiViewCustomModel, name)
}

func (svc *IrUiViewCustomService) GetByIds(ids []int) (*types.IrUiViewCustoms, error) {
	i := &types.IrUiViewCustoms{}
	return i, svc.client.getByIds(types.IrUiViewCustomModel, ids, i)
}

func (svc *IrUiViewCustomService) GetByName(name string) (*types.IrUiViewCustoms, error) {
	i := &types.IrUiViewCustoms{}
	return i, svc.client.getByName(types.IrUiViewCustomModel, name, i)
}

func (svc *IrUiViewCustomService) GetByField(field string, value string) (*types.IrUiViewCustoms, error) {
	i := &types.IrUiViewCustoms{}
	return i, svc.client.getByName(types.IrUiViewCustomModel, field, value, i)
}

func (svc *IrUiViewCustomService) GetAll() (*types.IrUiViewCustoms, error) {
	i := &types.IrUiViewCustoms{}
	return i, svc.client.getAll(types.IrUiViewCustomModel, i)
}

func (svc *IrUiViewCustomService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrUiViewCustomModel, fields, relations)
}

func (svc *IrUiViewCustomService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrUiViewCustomModel, ids, fields, relations)
}

func (svc *IrUiViewCustomService) Delete(ids []int) error {
	return svc.client.delete(types.IrUiViewCustomModel, ids)
}
