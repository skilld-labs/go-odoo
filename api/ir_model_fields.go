package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrModelFieldsService struct {
	client *Client
}

func NewIrModelFieldsService(c *Client) *IrModelFieldsService {
	return &IrModelFieldsService{client: c}
}

func (svc *IrModelFieldsService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrModelFieldsModel, name)
}

func (svc *IrModelFieldsService) GetByIds(ids []int64) (*types.IrModelFieldss, error) {
	i := &types.IrModelFieldss{}
	return i, svc.client.getByIds(types.IrModelFieldsModel, ids, i)
}

func (svc *IrModelFieldsService) GetByName(name string) (*types.IrModelFieldss, error) {
	i := &types.IrModelFieldss{}
	return i, svc.client.getByName(types.IrModelFieldsModel, name, i)
}

func (svc *IrModelFieldsService) GetByField(field string, value string) (*types.IrModelFieldss, error) {
	i := &types.IrModelFieldss{}
	return i, svc.client.getByField(types.IrModelFieldsModel, field, value, i)
}

func (svc *IrModelFieldsService) GetAll() (*types.IrModelFieldss, error) {
	i := &types.IrModelFieldss{}
	return i, svc.client.getAll(types.IrModelFieldsModel, i)
}

func (svc *IrModelFieldsService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrModelFieldsModel, fields, relations)
}

func (svc *IrModelFieldsService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrModelFieldsModel, ids, fields, relations)
}

func (svc *IrModelFieldsService) Delete(ids []int64) error {
	return svc.client.delete(types.IrModelFieldsModel, ids)
}
