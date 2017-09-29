package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ChangePasswordUserService struct {
	client *Client
}

func NewChangePasswordUserService(c *Client) *ChangePasswordUserService {
	return &ChangePasswordUserService{client: c}
}

func (svc *ChangePasswordUserService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ChangePasswordUserModel, name)
}

func (svc *ChangePasswordUserService) GetByIds(ids []int64) (*types.ChangePasswordUsers, error) {
	c := &types.ChangePasswordUsers{}
	return c, svc.client.getByIds(types.ChangePasswordUserModel, ids, c)
}

func (svc *ChangePasswordUserService) GetByName(name string) (*types.ChangePasswordUsers, error) {
	c := &types.ChangePasswordUsers{}
	return c, svc.client.getByName(types.ChangePasswordUserModel, name, c)
}

func (svc *ChangePasswordUserService) GetByField(field string, value string) (*types.ChangePasswordUsers, error) {
	c := &types.ChangePasswordUsers{}
	return c, svc.client.getByField(types.ChangePasswordUserModel, field, value, c)
}

func (svc *ChangePasswordUserService) GetAll() (*types.ChangePasswordUsers, error) {
	c := &types.ChangePasswordUsers{}
	return c, svc.client.getAll(types.ChangePasswordUserModel, c)
}

func (svc *ChangePasswordUserService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ChangePasswordUserModel, fields, relations)
}

func (svc *ChangePasswordUserService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ChangePasswordUserModel, ids, fields, relations)
}

func (svc *ChangePasswordUserService) Delete(ids []int64) error {
	return svc.client.delete(types.ChangePasswordUserModel, ids)
}
