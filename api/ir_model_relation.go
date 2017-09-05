package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrModelRelationService struct {
	client *Client
}

func NewIrModelRelationService(c *Client) *IrModelRelationService {
	return &IrModelRelationService{client: c}
}

func (svc *IrModelRelationService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrModelRelationModel, name)
}

func (svc *IrModelRelationService) GetByIds(ids []int) (*types.IrModelRelations, error) {
	i := &types.IrModelRelations{}
	return i, svc.client.getByIds(types.IrModelRelationModel, ids, i)
}

func (svc *IrModelRelationService) GetByName(name string) (*types.IrModelRelations, error) {
	i := &types.IrModelRelations{}
	return i, svc.client.getByName(types.IrModelRelationModel, name, i)
}

func (svc *IrModelRelationService) GetByField(field string, value string) (*types.IrModelRelations, error) {
	i := &types.IrModelRelations{}
	return i, svc.client.getByName(types.IrModelRelationModel, field, value, i)
}

func (svc *IrModelRelationService) GetAll() (*types.IrModelRelations, error) {
	i := &types.IrModelRelations{}
	return i, svc.client.getAll(types.IrModelRelationModel, i)
}

func (svc *IrModelRelationService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrModelRelationModel, fields, relations)
}

func (svc *IrModelRelationService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrModelRelationModel, ids, fields, relations)
}

func (svc *IrModelRelationService) Delete(ids []int) error {
	return svc.client.delete(types.IrModelRelationModel, ids)
}
