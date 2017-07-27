package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailMassMailingTagService struct {
	client *Client
}

func NewMailMassMailingTagService(c *Client) *MailMassMailingTagService {
	return &MailMassMailingTagService{client: c}
}

func (svc *MailMassMailingTagService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.MailMassMailingTagModel, name)
}

func (svc *MailMassMailingTagService) GetByIds(ids []int) (*types.MailMassMailingTags, error) {
	m := &types.MailMassMailingTags{}
	return m, svc.client.getByIds(types.MailMassMailingTagModel, ids, m)
}

func (svc *MailMassMailingTagService) GetByName(name string) (*types.MailMassMailingTags, error) {
	m := &types.MailMassMailingTags{}
	return m, svc.client.getByName(types.MailMassMailingTagModel, name, m)
}

func (svc *MailMassMailingTagService) GetAll() (*types.MailMassMailingTags, error) {
	m := &types.MailMassMailingTags{}
	return m, svc.client.getAll(types.MailMassMailingTagModel, m)
}

func (svc *MailMassMailingTagService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.MailMassMailingTagModel, fields, relations)
}

func (svc *MailMassMailingTagService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailMassMailingTagModel, ids, fields, relations)
}

func (svc *MailMassMailingTagService) Delete(ids []int) error {
	return svc.client.delete(types.MailMassMailingTagModel, ids)
}
