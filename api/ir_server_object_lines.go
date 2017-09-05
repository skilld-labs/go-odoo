package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrServerObjectLinesService struct {
	client *Client
}

func NewIrServerObjectLinesService(c *Client) *IrServerObjectLinesService {
	return &IrServerObjectLinesService{client: c}
}

func (svc *IrServerObjectLinesService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrServerObjectLinesModel, name)
}

func (svc *IrServerObjectLinesService) GetByIds(ids []int) (*types.IrServerObjectLiness, error) {
	i := &types.IrServerObjectLiness{}
	return i, svc.client.getByIds(types.IrServerObjectLinesModel, ids, i)
}

func (svc *IrServerObjectLinesService) GetByName(name string) (*types.IrServerObjectLiness, error) {
	i := &types.IrServerObjectLiness{}
	return i, svc.client.getByName(types.IrServerObjectLinesModel, name, i)
}

func (svc *IrServerObjectLinesService) GetByField(field string, value string) (*types.IrServerObjectLiness, error) {
	i := &types.IrServerObjectLiness{}
	return i, svc.client.getByField(types.IrServerObjectLinesModel, field, value, i)
}

func (svc *IrServerObjectLinesService) GetAll() (*types.IrServerObjectLiness, error) {
	i := &types.IrServerObjectLiness{}
	return i, svc.client.getAll(types.IrServerObjectLinesModel, i)
}

func (svc *IrServerObjectLinesService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrServerObjectLinesModel, fields, relations)
}

func (svc *IrServerObjectLinesService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrServerObjectLinesModel, ids, fields, relations)
}

func (svc *IrServerObjectLinesService) Delete(ids []int) error {
	return svc.client.delete(types.IrServerObjectLinesModel, ids)
}
