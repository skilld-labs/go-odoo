package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailMassMailingTestService struct {
	client *Client
}

func NewMailMassMailingTestService(c *Client) *MailMassMailingTestService {
	return &MailMassMailingTestService{client: c}
}

func (svc *MailMassMailingTestService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.MailMassMailingTestModel, name)
}

func (svc *MailMassMailingTestService) GetByIds(ids []int64) (*types.MailMassMailingTests, error) {
	m := &types.MailMassMailingTests{}
	return m, svc.client.getByIds(types.MailMassMailingTestModel, ids, m)
}

func (svc *MailMassMailingTestService) GetByName(name string) (*types.MailMassMailingTests, error) {
	m := &types.MailMassMailingTests{}
	return m, svc.client.getByName(types.MailMassMailingTestModel, name, m)
}

func (svc *MailMassMailingTestService) GetByField(field string, value string) (*types.MailMassMailingTests, error) {
	m := &types.MailMassMailingTests{}
	return m, svc.client.getByField(types.MailMassMailingTestModel, field, value, m)
}

func (svc *MailMassMailingTestService) GetAll() (*types.MailMassMailingTests, error) {
	m := &types.MailMassMailingTests{}
	return m, svc.client.getAll(types.MailMassMailingTestModel, m)
}

func (svc *MailMassMailingTestService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.MailMassMailingTestModel, fields, relations)
}

func (svc *MailMassMailingTestService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailMassMailingTestModel, ids, fields, relations)
}

func (svc *MailMassMailingTestService) Delete(ids []int64) error {
	return svc.client.delete(types.MailMassMailingTestModel, ids)
}
