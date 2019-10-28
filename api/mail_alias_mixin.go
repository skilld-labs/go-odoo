package api

import (
	"github.com/morezig/go-odoo/types"
)

type MailAliasMixinService struct {
	client *Client
}

func NewMailAliasMixinService(c *Client) *MailAliasMixinService {
	return &MailAliasMixinService{client: c}
}

func (svc *MailAliasMixinService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.MailAliasMixinModel, name)
}

func (svc *MailAliasMixinService) GetByIds(ids []int64) (*types.MailAliasMixins, error) {
	m := &types.MailAliasMixins{}
	return m, svc.client.getByIds(types.MailAliasMixinModel, ids, m)
}

func (svc *MailAliasMixinService) GetByName(name string) (*types.MailAliasMixins, error) {
	m := &types.MailAliasMixins{}
	return m, svc.client.getByName(types.MailAliasMixinModel, name, m)
}

func (svc *MailAliasMixinService) GetByField(field string, value string) (*types.MailAliasMixins, error) {
	m := &types.MailAliasMixins{}
	return m, svc.client.getByField(types.MailAliasMixinModel, field, value, m)
}

func (svc *MailAliasMixinService) GetAll() (*types.MailAliasMixins, error) {
	m := &types.MailAliasMixins{}
	return m, svc.client.getAll(types.MailAliasMixinModel, m)
}

func (svc *MailAliasMixinService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.MailAliasMixinModel, fields, relations)
}

func (svc *MailAliasMixinService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailAliasMixinModel, ids, fields, relations)
}

func (svc *MailAliasMixinService) Delete(ids []int64) error {
	return svc.client.delete(types.MailAliasMixinModel, ids)
}
