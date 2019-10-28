package api

import (
	"github.com/morezig/go-odoo/types"
)

type MailActivityMixinService struct {
	client *Client
}

func NewMailActivityMixinService(c *Client) *MailActivityMixinService {
	return &MailActivityMixinService{client: c}
}

func (svc *MailActivityMixinService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.MailActivityMixinModel, name)
}

func (svc *MailActivityMixinService) GetByIds(ids []int64) (*types.MailActivityMixins, error) {
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

func (svc *MailActivityMixinService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.MailActivityMixinModel, fields, relations)
}

func (svc *MailActivityMixinService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailActivityMixinModel, ids, fields, relations)
}

func (svc *MailActivityMixinService) Delete(ids []int64) error {
	return svc.client.delete(types.MailActivityMixinModel, ids)
}
