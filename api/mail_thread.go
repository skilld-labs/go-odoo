package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailThreadService struct {
	client *Client
}

func NewMailThreadService(c *Client) *MailThreadService {
	return &MailThreadService{client: c}
}

func (svc *MailThreadService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.MailThreadModel, name)
}

func (svc *MailThreadService) GetByIds(ids []int) (*types.MailThreads, error) {
	m := &types.MailThreads{}
	return m, svc.client.getByIds(types.MailThreadModel, ids, m)
}

func (svc *MailThreadService) GetByName(name string) (*types.MailThreads, error) {
	m := &types.MailThreads{}
	return m, svc.client.getByName(types.MailThreadModel, name, m)
}

func (svc *MailThreadService) GetByField(field string, value string) (*types.MailThreads, error) {
	m := &types.MailThreads{}
	return m, svc.client.getByField(types.MailThreadModel, field, value, m)
}

func (svc *MailThreadService) GetAll() (*types.MailThreads, error) {
	m := &types.MailThreads{}
	return m, svc.client.getAll(types.MailThreadModel, m)
}

func (svc *MailThreadService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.MailThreadModel, fields, relations)
}

func (svc *MailThreadService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailThreadModel, ids, fields, relations)
}

func (svc *MailThreadService) Delete(ids []int) error {
	return svc.client.delete(types.MailThreadModel, ids)
}
