package api

import (
	"github.com/morezig/go-odoo/types"
)

type MailChannelPartnerService struct {
	client *Client
}

func NewMailChannelPartnerService(c *Client) *MailChannelPartnerService {
	return &MailChannelPartnerService{client: c}
}

func (svc *MailChannelPartnerService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.MailChannelPartnerModel, name)
}

func (svc *MailChannelPartnerService) GetByIds(ids []int64) (*types.MailChannelPartners, error) {
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

func (svc *MailChannelPartnerService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.MailChannelPartnerModel, fields, relations)
}

func (svc *MailChannelPartnerService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailChannelPartnerModel, ids, fields, relations)
}

func (svc *MailChannelPartnerService) Delete(ids []int64) error {
	return svc.client.delete(types.MailChannelPartnerModel, ids)
}
