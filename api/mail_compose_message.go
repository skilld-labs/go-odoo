package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailComposeMessageService struct {
	client *Client
}

func NewMailComposeMessageService(c *Client) *MailComposeMessageService {
	return &MailComposeMessageService{client: c}
}

func (svc *MailComposeMessageService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.MailComposeMessageModel, name)
}

func (svc *MailComposeMessageService) GetByIds(ids []int64) (*types.MailComposeMessages, error) {
	m := &types.MailComposeMessages{}
	return m, svc.client.getByIds(types.MailComposeMessageModel, ids, m)
}

func (svc *MailComposeMessageService) GetByName(name string) (*types.MailComposeMessages, error) {
	m := &types.MailComposeMessages{}
	return m, svc.client.getByName(types.MailComposeMessageModel, name, m)
}

func (svc *MailComposeMessageService) GetByField(field string, value string) (*types.MailComposeMessages, error) {
	m := &types.MailComposeMessages{}
	return m, svc.client.getByField(types.MailComposeMessageModel, field, value, m)
}

func (svc *MailComposeMessageService) GetAll() (*types.MailComposeMessages, error) {
	m := &types.MailComposeMessages{}
	return m, svc.client.getAll(types.MailComposeMessageModel, m)
}

func (svc *MailComposeMessageService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.MailComposeMessageModel, fields, relations)
}

func (svc *MailComposeMessageService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailComposeMessageModel, ids, fields, relations)
}

func (svc *MailComposeMessageService) Delete(ids []int64) error {
	return svc.client.delete(types.MailComposeMessageModel, ids)
}
