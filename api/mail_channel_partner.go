package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailChannelPartnerService struct {
	client *Client
}

func NewMailChannelPartnerService(c *Client) *MailChannelPartnerService {
	return &MailChannelPartnerService{client: c}
}

func (svc *MailChannelPartnerService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.MailChannelPartnerModel, name)
}

func (svc *MailChannelPartnerService) GetByIds(ids []int) (*types.MailChannelPartners, error) {
	m := &types.MailChannelPartners{}
	return m, svc.client.getByIds(types.MailChannelPartnerModel, ids, m)
}

func (svc *MailChannelPartnerService) GetByName(name string) (*types.MailChannelPartners, error) {
	m := &types.MailChannelPartners{}
	return m, svc.client.getByName(types.MailChannelPartnerModel, name, m)
}

func (svc *MailChannelPartnerService) GetByField(field string, value string) (*types.MailChannelPartners, error) {
	m := &types.MailChannelPartners{}
	return m, svc.client.getByField(types.MailChannelPartnerModel, field, value, m)
}

func (svc *MailChannelPartnerService) GetAll() (*types.MailChannelPartners, error) {
	m := &types.MailChannelPartners{}
	return m, svc.client.getAll(types.MailChannelPartnerModel, m)
}

func (svc *MailChannelPartnerService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.MailChannelPartnerModel, fields, relations)
}

func (svc *MailChannelPartnerService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailChannelPartnerModel, ids, fields, relations)
}

func (svc *MailChannelPartnerService) Delete(ids []int) error {
	return svc.client.delete(types.MailChannelPartnerModel, ids)
}
