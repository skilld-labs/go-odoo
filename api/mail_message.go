package api

import (
	"github.com/morezig/go-odoo/types"
)

type MailMessageService struct {
	client *Client
}

func NewMailMessageService(c *Client) *MailMessageService {
	return &MailMessageService{client: c}
}

func (svc *MailMessageService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.MailMessageModel, name)
}

func (svc *MailMessageService) GetByIds(ids []int64) (*types.MailMessages, error) {
	m := &types.MailMessages{}
	return m, svc.client.getByIds(types.MailMessageModel, ids, m)
}

func (svc *MailMessageService) GetByName(name string) (*types.MailMessages, error) {
	m := &types.MailMessages{}
	return m, svc.client.getByName(types.MailMessageModel, name, m)
}

func (svc *MailMessageService) GetByField(field string, value string) (*types.MailMessages, error) {
	m := &types.MailMessages{}
	return m, svc.client.getByField(types.MailMessageModel, field, value, m)
}

func (svc *MailMessageService) GetAll() (*types.MailMessages, error) {
	m := &types.MailMessages{}
	return m, svc.client.getAll(types.MailMessageModel, m)
}

func (svc *MailMessageService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.MailMessageModel, fields, relations)
}

func (svc *MailMessageService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailMessageModel, ids, fields, relations)
}

func (svc *MailMessageService) Delete(ids []int64) error {
	return svc.client.delete(types.MailMessageModel, ids)
}
