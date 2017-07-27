package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailTemplateService struct {
	client *Client
}

func NewMailTemplateService(c *Client) *MailTemplateService {
	return &MailTemplateService{client: c}
}

func (svc *MailTemplateService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.MailTemplateModel, name)
}

func (svc *MailTemplateService) GetByIds(ids []int) (*types.MailTemplates, error) {
	m := &types.MailTemplates{}
	return m, svc.client.getByIds(types.MailTemplateModel, ids, m)
}

func (svc *MailTemplateService) GetByName(name string) (*types.MailTemplates, error) {
	m := &types.MailTemplates{}
	return m, svc.client.getByName(types.MailTemplateModel, name, m)
}

func (svc *MailTemplateService) GetAll() (*types.MailTemplates, error) {
	m := &types.MailTemplates{}
	return m, svc.client.getAll(types.MailTemplateModel, m)
}

func (svc *MailTemplateService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.MailTemplateModel, fields, relations)
}

func (svc *MailTemplateService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailTemplateModel, ids, fields, relations)
}

func (svc *MailTemplateService) Delete(ids []int) error {
	return svc.client.delete(types.MailTemplateModel, ids)
}
