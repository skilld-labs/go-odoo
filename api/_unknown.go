package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type UnknownService struct {
	client *Client
}

func NewUnknownService(c *Client) *UnknownService {
	return &UnknownService{client: c}
}

func (svc *UnknownService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.UnknownModel, name)
}

func (svc *UnknownService) GetByIds(ids []int) (*types.Unknowns, error) {
	u := &types.Unknowns{}
	return u, svc.client.getByIds(types.UnknownModel, ids, u)
}

func (svc *UnknownService) GetByName(name string) (*types.Unknowns, error) {
	u := &types.Unknowns{}
	return u, svc.client.getByName(types.UnknownModel, name, u)
}

func (svc *UnknownService) GetByField(field string, value string) (*types.Unknowns, error) {
	u := &types.Unknowns{}
	return u, svc.client.getByField(types.UnknownModel, field, value, u)
}

func (svc *UnknownService) GetAll() (*types.Unknowns, error) {
	u := &types.Unknowns{}
	return u, svc.client.getAll(types.UnknownModel, u)
}

func (svc *UnknownService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.UnknownModel, fields, relations)
}

func (svc *UnknownService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.UnknownModel, ids, fields, relations)
}

func (svc *UnknownService) Delete(ids []int) error {
	return svc.client.delete(types.UnknownModel, ids)
}
