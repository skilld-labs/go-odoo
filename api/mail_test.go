package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailTestService struct {
	client *Client
}

func NewMailTestService(c *Client) *MailTestService {
	return &MailTestService{client: c}
}

func (svc *MailTestService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.MailTestModel, name)
}

func (svc *MailTestService) GetByIds(ids []int) (*types.MailTests, error) {
	m := &types.MailTests{}
	return m, svc.client.getByIds(types.MailTestModel, ids, m)
}

func (svc *MailTestService) GetByName(name string) (*types.MailTests, error) {
	m := &types.MailTests{}
	return m, svc.client.getByName(types.MailTestModel, name, m)
}

func (svc *MailTestService) GetByField(field string, value string) (*types.MailTests, error) {
	m := &types.MailTests{}
	return m, svc.client.getByField(types.MailTestModel, field, value, m)
}

func (svc *MailTestService) GetAll() (*types.MailTests, error) {
	m := &types.MailTests{}
	return m, svc.client.getAll(types.MailTestModel, m)
}

func (svc *MailTestService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.MailTestModel, fields, relations)
}

func (svc *MailTestService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailTestModel, ids, fields, relations)
}

func (svc *MailTestService) Delete(ids []int) error {
	return svc.client.delete(types.MailTestModel, ids)
}
