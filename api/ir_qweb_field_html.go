package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrQwebFieldHtmlService struct {
	client *Client
}

func NewIrQwebFieldHtmlService(c *Client) *IrQwebFieldHtmlService {
	return &IrQwebFieldHtmlService{client: c}
}

func (svc *IrQwebFieldHtmlService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrQwebFieldHtmlModel, name)
}

func (svc *IrQwebFieldHtmlService) GetByIds(ids []int) (*types.IrQwebFieldHtmls, error) {
	i := &types.IrQwebFieldHtmls{}
	return i, svc.client.getByIds(types.IrQwebFieldHtmlModel, ids, i)
}

func (svc *IrQwebFieldHtmlService) GetByName(name string) (*types.IrQwebFieldHtmls, error) {
	i := &types.IrQwebFieldHtmls{}
	return i, svc.client.getByName(types.IrQwebFieldHtmlModel, name, i)
}

func (svc *IrQwebFieldHtmlService) GetByField(field string, value string) (*types.IrQwebFieldHtmls, error) {
	i := &types.IrQwebFieldHtmls{}
	return i, svc.client.getByName(types.IrQwebFieldHtmlModel, field, value, i)
}

func (svc *IrQwebFieldHtmlService) GetAll() (*types.IrQwebFieldHtmls, error) {
	i := &types.IrQwebFieldHtmls{}
	return i, svc.client.getAll(types.IrQwebFieldHtmlModel, i)
}

func (svc *IrQwebFieldHtmlService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrQwebFieldHtmlModel, fields, relations)
}

func (svc *IrQwebFieldHtmlService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrQwebFieldHtmlModel, ids, fields, relations)
}

func (svc *IrQwebFieldHtmlService) Delete(ids []int) error {
	return svc.client.delete(types.IrQwebFieldHtmlModel, ids)
}
