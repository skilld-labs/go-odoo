package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseService struct {
	client *Client
}

func NewBaseService(c *Client) *BaseService {
	return &BaseService{client: c}
}

func (svc *BaseService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BaseModel, name)
}

func (svc *BaseService) GetByIds(ids []int) (*types.Bases, error) {
	b := &types.Bases{}
	return b, svc.client.getByIds(types.BaseModel, ids, b)
}

func (svc *BaseService) GetByName(name string) (*types.Bases, error) {
	b := &types.Bases{}
	return b, svc.client.getByName(types.BaseModel, name, b)
}

func (svc *BaseService) GetByField(field string, value string) (*types.Bases, error) {
	b := &types.Bases{}
	return b, svc.client.getByField(types.BaseModel, field, value, b)
}

func (svc *BaseService) GetAll() (*types.Bases, error) {
	b := &types.Bases{}
	return b, svc.client.getAll(types.BaseModel, b)
}

func (svc *BaseService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BaseModel, fields, relations)
}

func (svc *BaseService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseModel, ids, fields, relations)
}

func (svc *BaseService) Delete(ids []int) error {
	return svc.client.delete(types.BaseModel, ids)
}
