package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ImLivechatChannelService struct {
	client *Client
}

func NewImLivechatChannelService(c *Client) *ImLivechatChannelService {
	return &ImLivechatChannelService{client: c}
}

func (svc *ImLivechatChannelService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ImLivechatChannelModel, name)
}

func (svc *ImLivechatChannelService) GetByIds(ids []int64) (*types.ImLivechatChannels, error) {
	i := &types.ImLivechatChannels{}
	return i, svc.client.getByIds(types.ImLivechatChannelModel, ids, i)
}

func (svc *ImLivechatChannelService) GetByName(name string) (*types.ImLivechatChannels, error) {
	i := &types.ImLivechatChannels{}
	return i, svc.client.getByName(types.ImLivechatChannelModel, name, i)
}

func (svc *ImLivechatChannelService) GetByField(field string, value string) (*types.ImLivechatChannels, error) {
	i := &types.ImLivechatChannels{}
	return i, svc.client.getByField(types.ImLivechatChannelModel, field, value, i)
}

func (svc *ImLivechatChannelService) GetAll() (*types.ImLivechatChannels, error) {
	i := &types.ImLivechatChannels{}
	return i, svc.client.getAll(types.ImLivechatChannelModel, i)
}

func (svc *ImLivechatChannelService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ImLivechatChannelModel, fields, relations)
}

func (svc *ImLivechatChannelService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ImLivechatChannelModel, ids, fields, relations)
}

func (svc *ImLivechatChannelService) Delete(ids []int64) error {
	return svc.client.delete(types.ImLivechatChannelModel, ids)
}
