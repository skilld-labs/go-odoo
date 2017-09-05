package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailMessageSubtypeService struct {
	client *Client
}

func NewMailMessageSubtypeService(c *Client) *MailMessageSubtypeService {
	return &MailMessageSubtypeService{client: c}
}

func (svc *MailMessageSubtypeService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.MailMessageSubtypeModel, name)
}

func (svc *MailMessageSubtypeService) GetByIds(ids []int) (*types.MailMessageSubtypes, error) {
	m := &types.MailMessageSubtypes{}
	return m, svc.client.getByIds(types.MailMessageSubtypeModel, ids, m)
}

func (svc *MailMessageSubtypeService) GetByName(name string) (*types.MailMessageSubtypes, error) {
	m := &types.MailMessageSubtypes{}
	return m, svc.client.getByName(types.MailMessageSubtypeModel, name, m)
}

func (svc *MailMessageSubtypeService) GetByField(field string, value string) (*types.MailMessageSubtypes, error) {
	m := &types.MailMessageSubtypes{}
	return m, svc.client.getByField(types.MailMessageSubtypeModel, field, value, m)
}

func (svc *MailMessageSubtypeService) GetAll() (*types.MailMessageSubtypes, error) {
	m := &types.MailMessageSubtypes{}
	return m, svc.client.getAll(types.MailMessageSubtypeModel, m)
}

func (svc *MailMessageSubtypeService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.MailMessageSubtypeModel, fields, relations)
}

func (svc *MailMessageSubtypeService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailMessageSubtypeModel, ids, fields, relations)
}

func (svc *MailMessageSubtypeService) Delete(ids []int) error {
	return svc.client.delete(types.MailMessageSubtypeModel, ids)
}
