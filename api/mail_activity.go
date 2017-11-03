package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailActivityService struct {
	client *Client
}

func NewMailActivityService(c *Client) *MailActivityService {
	return &MailActivityService{client: c}
}

func (svc *MailActivityService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.MailActivityModel, name)
}

func (svc *MailActivityService) GetByIds(ids []int64) (*types.MailActivitys, error) {
	m := &types.MailActivitys{}
	return m, svc.client.getByIds(types.MailActivityModel, ids, m)
}

func (svc *MailActivityService) GetByName(name string) (*types.MailActivitys, error) {
	m := &types.MailActivitys{}
	return m, svc.client.getByName(types.MailActivityModel, name, m)
}

func (svc *MailActivityService) GetByField(field string, value string) (*types.MailActivitys, error) {
	m := &types.MailActivitys{}
	return m, svc.client.getByField(types.MailActivityModel, field, value, m)
}

func (svc *MailActivityService) GetAll() (*types.MailActivitys, error) {
	m := &types.MailActivitys{}
	return m, svc.client.getAll(types.MailActivityModel, m)
}

func (svc *MailActivityService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.MailActivityModel, fields, relations)
}

func (svc *MailActivityService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailActivityModel, ids, fields, relations)
}

func (svc *MailActivityService) Delete(ids []int64) error {
	return svc.client.delete(types.MailActivityModel, ids)
}
