package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailActivityMixinService struct {
	client *Client
}

func NewMailActivityMixinService(c *Client) *MailActivityMixinService {
	return &MailActivityMixinService{client: c}
}

func (svc *MailActivityMixinService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.MailActivityMixinModel, name)
}

func (svc *MailActivityMixinService) GetByIds(ids []int) (*types.MailActivityMixins, error) {
	m := &types.MailActivityMixins{}
	return m, svc.client.getByIds(types.MailActivityMixinModel, ids, m)
}

func (svc *MailActivityMixinService) GetByName(name string) (*types.MailActivityMixins, error) {
	m := &types.MailActivityMixins{}
	return m, svc.client.getByName(types.MailActivityMixinModel, name, m)
}

func (svc *MailActivityMixinService) GetByField(field string, value string) (*types.MailActivityMixins, error) {
	m := &types.MailActivityMixins{}
	return m, svc.client.getByField(types.MailActivityMixinModel, field, value, m)
}

func (svc *MailActivityMixinService) GetAll() (*types.MailActivityMixins, error) {
	m := &types.MailActivityMixins{}
	return m, svc.client.getAll(types.MailActivityMixinModel, m)
}

func (svc *MailActivityMixinService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.MailActivityMixinModel, fields, relations)
}

func (svc *MailActivityMixinService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailActivityMixinModel, ids, fields, relations)
}

func (svc *MailActivityMixinService) Delete(ids []int) error {
	return svc.client.delete(types.MailActivityMixinModel, ids)
}
