package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailNotificationService struct {
	client *Client
}

func NewMailNotificationService(c *Client) *MailNotificationService {
	return &MailNotificationService{client: c}
}

func (svc *MailNotificationService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.MailNotificationModel, name)
}

func (svc *MailNotificationService) GetByIds(ids []int) (*types.MailNotifications, error) {
	m := &types.MailNotifications{}
	return m, svc.client.getByIds(types.MailNotificationModel, ids, m)
}

func (svc *MailNotificationService) GetByName(name string) (*types.MailNotifications, error) {
	m := &types.MailNotifications{}
	return m, svc.client.getByName(types.MailNotificationModel, name, m)
}

func (svc *MailNotificationService) GetByField(field string, value string) (*types.MailNotifications, error) {
	m := &types.MailNotifications{}
	return m, svc.client.getByField(types.MailNotificationModel, field, value, m)
}

func (svc *MailNotificationService) GetAll() (*types.MailNotifications, error) {
	m := &types.MailNotifications{}
	return m, svc.client.getAll(types.MailNotificationModel, m)
}

func (svc *MailNotificationService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.MailNotificationModel, fields, relations)
}

func (svc *MailNotificationService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailNotificationModel, ids, fields, relations)
}

func (svc *MailNotificationService) Delete(ids []int) error {
	return svc.client.delete(types.MailNotificationModel, ids)
}
