package api

import (
	"github.com/morezig/go-odoo/types"
)

type MailMassMailingService struct {
	client *Client
}

func NewMailMassMailingService(c *Client) *MailMassMailingService {
	return &MailMassMailingService{client: c}
}

func (svc *MailMassMailingService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.MailMassMailingModel, name)
}

func (svc *MailMassMailingService) GetByIds(ids []int64) (*types.MailMassMailings, error) {
	m := &types.MailMassMailings{}
	return m, svc.client.getByIds(types.MailMassMailingModel, ids, m)
}

func (svc *MailMassMailingService) GetByName(name string) (*types.MailMassMailings, error) {
	m := &types.MailMassMailings{}
	return m, svc.client.getByName(types.MailMassMailingModel, name, m)
}

func (svc *MailMassMailingService) GetByField(field string, value string) (*types.MailMassMailings, error) {
	m := &types.MailMassMailings{}
	return m, svc.client.getByField(types.MailMassMailingModel, field, value, m)
}

func (svc *MailMassMailingService) GetAll() (*types.MailMassMailings, error) {
	m := &types.MailMassMailings{}
	return m, svc.client.getAll(types.MailMassMailingModel, m)
}

func (svc *MailMassMailingService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.MailMassMailingModel, fields, relations)
}

func (svc *MailMassMailingService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailMassMailingModel, ids, fields, relations)
}

func (svc *MailMassMailingService) Delete(ids []int64) error {
	return svc.client.delete(types.MailMassMailingModel, ids)
}
