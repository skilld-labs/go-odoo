package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailMassMailingListService struct {
	client *Client
}

func NewMailMassMailingListService(c *Client) *MailMassMailingListService {
	return &MailMassMailingListService{client: c}
}

func (svc *MailMassMailingListService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.MailMassMailingListModel, name)
}

func (svc *MailMassMailingListService) GetByIds(ids []int) (*types.MailMassMailingLists, error) {
	m := &types.MailMassMailingLists{}
	return m, svc.client.getByIds(types.MailMassMailingListModel, ids, m)
}

func (svc *MailMassMailingListService) GetByName(name string) (*types.MailMassMailingLists, error) {
	m := &types.MailMassMailingLists{}
	return m, svc.client.getByName(types.MailMassMailingListModel, name, m)
}

func (svc *MailMassMailingListService) GetByField(field string, value string) (*types.MailMassMailingLists, error) {
	m := &types.MailMassMailingLists{}
	return m, svc.client.getByField(types.MailMassMailingListModel, field, value, m)
}

func (svc *MailMassMailingListService) GetAll() (*types.MailMassMailingLists, error) {
	m := &types.MailMassMailingLists{}
	return m, svc.client.getAll(types.MailMassMailingListModel, m)
}

func (svc *MailMassMailingListService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.MailMassMailingListModel, fields, relations)
}

func (svc *MailMassMailingListService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailMassMailingListModel, ids, fields, relations)
}

func (svc *MailMassMailingListService) Delete(ids []int) error {
	return svc.client.delete(types.MailMassMailingListModel, ids)
}
