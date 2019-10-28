package api

import (
	"github.com/morezig/go-odoo/types"
)

type MailMassMailingContactService struct {
	client *Client
}

func NewMailMassMailingContactService(c *Client) *MailMassMailingContactService {
	return &MailMassMailingContactService{client: c}
}

func (svc *MailMassMailingContactService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.MailMassMailingContactModel, name)
}

func (svc *MailMassMailingContactService) GetByIds(ids []int64) (*types.MailMassMailingContacts, error) {
	m := &types.MailMassMailingContacts{}
	return m, svc.client.getByIds(types.MailMassMailingContactModel, ids, m)
}

func (svc *MailMassMailingContactService) GetByName(name string) (*types.MailMassMailingContacts, error) {
	m := &types.MailMassMailingContacts{}
	return m, svc.client.getByName(types.MailMassMailingContactModel, name, m)
}

func (svc *MailMassMailingContactService) GetByField(field string, value string) (*types.MailMassMailingContacts, error) {
	m := &types.MailMassMailingContacts{}
	return m, svc.client.getByField(types.MailMassMailingContactModel, field, value, m)
}

func (svc *MailMassMailingContactService) GetAll() (*types.MailMassMailingContacts, error) {
	m := &types.MailMassMailingContacts{}
	return m, svc.client.getAll(types.MailMassMailingContactModel, m)
}

func (svc *MailMassMailingContactService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.MailMassMailingContactModel, fields, relations)
}

func (svc *MailMassMailingContactService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailMassMailingContactModel, ids, fields, relations)
}

func (svc *MailMassMailingContactService) Delete(ids []int64) error {
	return svc.client.delete(types.MailMassMailingContactModel, ids)
}
