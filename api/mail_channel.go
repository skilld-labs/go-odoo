package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type MailChannelService struct {
	client *Client
}

func NewMailChannelService(c *Client) *MailChannelService {
	return &MailChannelService{client: c}
}

func (svc *MailChannelService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.MailChannelModel, name)
}

func (svc *MailChannelService) GetByIds(ids []int) (*types.MailChannels, error) {
	m := &types.MailChannels{}
	return m, svc.client.getByIds(types.MailChannelModel, ids, m)
}

func (svc *MailChannelService) GetByName(name string) (*types.MailChannels, error) {
	m := &types.MailChannels{}
	return m, svc.client.getByName(types.MailChannelModel, name, m)
}

func (svc *MailChannelService) GetAll() (*types.MailChannels, error) {
	m := &types.MailChannels{}
	return m, svc.client.getAll(types.MailChannelModel, m)
}

func (svc *MailChannelService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.MailChannelModel, fields, relations)
}

func (svc *MailChannelService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailChannelModel, ids, fields, relations)
}

func (svc *MailChannelService) Delete(ids []int) error {
	return svc.client.delete(types.MailChannelModel, ids)
}
