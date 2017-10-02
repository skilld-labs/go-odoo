package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailMailService struct {
	client *Client
}

func NewMailMailService(c *Client) *MailMailService {
	return &MailMailService{client: c}
}

func (svc *MailMailService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.MailMailModel, name)
}

func (svc *MailMailService) GetByIds(ids []int64) (*types.MailMails, error) {
	m := &types.MailMails{}
	return m, svc.client.getByIds(types.MailMailModel, ids, m)
}

func (svc *MailMailService) GetByName(name string) (*types.MailMails, error) {
	m := &types.MailMails{}
	return m, svc.client.getByName(types.MailMailModel, name, m)
}

func (svc *MailMailService) GetByField(field string, value string) (*types.MailMails, error) {
	m := &types.MailMails{}
	return m, svc.client.getByField(types.MailMailModel, field, value, m)
}

func (svc *MailMailService) GetAll() (*types.MailMails, error) {
	m := &types.MailMails{}
	return m, svc.client.getAll(types.MailMailModel, m)
}

func (svc *MailMailService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.MailMailModel, fields, relations)
}

func (svc *MailMailService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailMailModel, ids, fields, relations)
}

func (svc *MailMailService) Delete(ids []int64) error {
	return svc.client.delete(types.MailMailModel, ids)
}
