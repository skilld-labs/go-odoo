package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrUiMenuService struct {
	client *Client
}

func NewIrUiMenuService(c *Client) *IrUiMenuService {
	return &IrUiMenuService{client: c}
}

func (svc *IrUiMenuService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrUiMenuModel, name)
}

func (svc *IrUiMenuService) GetByIds(ids []int64) (*types.IrUiMenus, error) {
	i := &types.IrUiMenus{}
	return i, svc.client.getByIds(types.IrUiMenuModel, ids, i)
}

func (svc *IrUiMenuService) GetByName(name string) (*types.IrUiMenus, error) {
	i := &types.IrUiMenus{}
	return i, svc.client.getByName(types.IrUiMenuModel, name, i)
}

func (svc *IrUiMenuService) GetByField(field string, value string) (*types.IrUiMenus, error) {
	i := &types.IrUiMenus{}
	return i, svc.client.getByField(types.IrUiMenuModel, field, value, i)
}

func (svc *IrUiMenuService) GetAll() (*types.IrUiMenus, error) {
	i := &types.IrUiMenus{}
	return i, svc.client.getAll(types.IrUiMenuModel, i)
}

func (svc *IrUiMenuService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrUiMenuModel, fields, relations)
}

func (svc *IrUiMenuService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrUiMenuModel, ids, fields, relations)
}

func (svc *IrUiMenuService) Delete(ids []int64) error {
	return svc.client.delete(types.IrUiMenuModel, ids)
}
