package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailTrackingValueService struct {
	client *Client
}

func NewMailTrackingValueService(c *Client) *MailTrackingValueService {
	return &MailTrackingValueService{client: c}
}

func (svc *MailTrackingValueService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.MailTrackingValueModel, name)
}

func (svc *MailTrackingValueService) GetByIds(ids []int64) (*types.MailTrackingValues, error) {
	m := &types.MailTrackingValues{}
	return m, svc.client.getByIds(types.MailTrackingValueModel, ids, m)
}

func (svc *MailTrackingValueService) GetByName(name string) (*types.MailTrackingValues, error) {
	m := &types.MailTrackingValues{}
	return m, svc.client.getByName(types.MailTrackingValueModel, name, m)
}

func (svc *MailTrackingValueService) GetByField(field string, value string) (*types.MailTrackingValues, error) {
	m := &types.MailTrackingValues{}
	return m, svc.client.getByField(types.MailTrackingValueModel, field, value, m)
}

func (svc *MailTrackingValueService) GetAll() (*types.MailTrackingValues, error) {
	m := &types.MailTrackingValues{}
	return m, svc.client.getAll(types.MailTrackingValueModel, m)
}

func (svc *MailTrackingValueService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.MailTrackingValueModel, fields, relations)
}

func (svc *MailTrackingValueService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailTrackingValueModel, ids, fields, relations)
}

func (svc *MailTrackingValueService) Delete(ids []int64) error {
	return svc.client.delete(types.MailTrackingValueModel, ids)
}
