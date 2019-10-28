package api

import (
	"github.com/morezig/go-odoo/types"
)

type MailChannelService struct {
	client *Client
}

func NewMailChannelService(c *Client) *MailChannelService {
	return &MailChannelService{client: c}
}

func (svc *MailChannelService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.MailChannelModel, name)
}

func (svc *MailChannelService) GetByIds(ids []int64) (*types.MailChannels, error) {
	m := &types.MailChannels{}
	return m, svc.client.getByIds(types.MailChannelModel, ids, m)
}

func (svc *MailChannelService) GetByName(name string) (*types.MailChannels, error) {
	m := &types.MailChannels{}
	return m, svc.client.getByName(types.MailChannelModel, name, m)
}

func (svc *MailChannelService) GetByField(field string, value string) (*types.MailChannels, error) {
	m := &types.MailChannels{}
	return m, svc.client.getByField(types.MailChannelModel, field, value, m)
}

func (svc *MailChannelService) GetAll() (*types.MailChannels, error) {
	m := &types.MailChannels{}
	return m, svc.client.getAll(types.MailChannelModel, m)
}

func (svc *MailChannelService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.MailChannelModel, fields, relations)
}

func (svc *MailChannelService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailChannelModel, ids, fields, relations)
}

func (svc *MailChannelService) Delete(ids []int64) error {
	return svc.client.delete(types.MailChannelModel, ids)
}
