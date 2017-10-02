package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailFollowersService struct {
	client *Client
}

func NewMailFollowersService(c *Client) *MailFollowersService {
	return &MailFollowersService{client: c}
}

func (svc *MailFollowersService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.MailFollowersModel, name)
}

func (svc *MailFollowersService) GetByIds(ids []int64) (*types.MailFollowerss, error) {
	m := &types.MailFollowerss{}
	return m, svc.client.getByIds(types.MailFollowersModel, ids, m)
}

func (svc *MailFollowersService) GetByName(name string) (*types.MailFollowerss, error) {
	m := &types.MailFollowerss{}
	return m, svc.client.getByName(types.MailFollowersModel, name, m)
}

func (svc *MailFollowersService) GetByField(field string, value string) (*types.MailFollowerss, error) {
	m := &types.MailFollowerss{}
	return m, svc.client.getByField(types.MailFollowersModel, field, value, m)
}

func (svc *MailFollowersService) GetAll() (*types.MailFollowerss, error) {
	m := &types.MailFollowerss{}
	return m, svc.client.getAll(types.MailFollowersModel, m)
}

func (svc *MailFollowersService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.MailFollowersModel, fields, relations)
}

func (svc *MailFollowersService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailFollowersModel, ids, fields, relations)
}

func (svc *MailFollowersService) Delete(ids []int64) error {
	return svc.client.delete(types.MailFollowersModel, ids)
}
