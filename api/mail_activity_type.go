package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailActivityTypeService struct {
	client *Client
}

func NewMailActivityTypeService(c *Client) *MailActivityTypeService {
	return &MailActivityTypeService{client: c}
}

func (svc *MailActivityTypeService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.MailActivityTypeModel, name)
}

func (svc *MailActivityTypeService) GetByIds(ids []int) (*types.MailActivityTypes, error) {
	m := &types.MailActivityTypes{}
	return m, svc.client.getByIds(types.MailActivityTypeModel, ids, m)
}

func (svc *MailActivityTypeService) GetByName(name string) (*types.MailActivityTypes, error) {
	m := &types.MailActivityTypes{}
	return m, svc.client.getByName(types.MailActivityTypeModel, name, m)
}

func (svc *MailActivityTypeService) GetByField(field string, value string) (*types.MailActivityTypes, error) {
	m := &types.MailActivityTypes{}
	return m, svc.client.getByField(types.MailActivityTypeModel, field, value, m)
}

func (svc *MailActivityTypeService) GetAll() (*types.MailActivityTypes, error) {
	m := &types.MailActivityTypes{}
	return m, svc.client.getAll(types.MailActivityTypeModel, m)
}

func (svc *MailActivityTypeService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.MailActivityTypeModel, fields, relations)
}

func (svc *MailActivityTypeService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailActivityTypeModel, ids, fields, relations)
}

func (svc *MailActivityTypeService) Delete(ids []int) error {
	return svc.client.delete(types.MailActivityTypeModel, ids)
}
