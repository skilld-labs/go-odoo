package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ResUsersService struct {
	client *Client
}

func NewResUsersService(c *Client) *ResUsersService {
	return &ResUsersService{client: c}
}

func (svc *ResUsersService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ResUsersModel, name)
}

func (svc *ResUsersService) GetByIds(ids []int) (*types.ResUserss, error) {
	r := &types.ResUserss{}
	return r, svc.client.getByIds(types.ResUsersModel, ids, r)
}

func (svc *ResUsersService) GetByName(name string) (*types.ResUserss, error) {
	r := &types.ResUserss{}
	return r, svc.client.getByName(types.ResUsersModel, name, r)
}

func (svc *ResUsersService) GetByField(field string, value string) (*types.ResUserss, error) {
	r := &types.ResUserss{}
	return r, svc.client.getByField(types.ResUsersModel, field, value, r)
}

func (svc *ResUsersService) GetAll() (*types.ResUserss, error) {
	r := &types.ResUserss{}
	return r, svc.client.getAll(types.ResUsersModel, r)
}

func (svc *ResUsersService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ResUsersModel, fields, relations)
}

func (svc *ResUsersService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResUsersModel, ids, fields, relations)
}

func (svc *ResUsersService) Delete(ids []int) error {
	return svc.client.delete(types.ResUsersModel, ids)
}
