package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ImLivechatChannelRuleService struct {
	client *Client
}

func NewImLivechatChannelRuleService(c *Client) *ImLivechatChannelRuleService {
	return &ImLivechatChannelRuleService{client: c}
}

func (svc *ImLivechatChannelRuleService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ImLivechatChannelRuleModel, name)
}

func (svc *ImLivechatChannelRuleService) GetByIds(ids []int) (*types.ImLivechatChannelRules, error) {
	i := &types.ImLivechatChannelRules{}
	return i, svc.client.getByIds(types.ImLivechatChannelRuleModel, ids, i)
}

func (svc *ImLivechatChannelRuleService) GetByName(name string) (*types.ImLivechatChannelRules, error) {
	i := &types.ImLivechatChannelRules{}
	return i, svc.client.getByName(types.ImLivechatChannelRuleModel, name, i)
}

func (svc *ImLivechatChannelRuleService) GetByField(field string, value string) (*types.ImLivechatChannelRules, error) {
	i := &types.ImLivechatChannelRules{}
	return i, svc.client.getByField(types.ImLivechatChannelRuleModel, field, value, i)
}

func (svc *ImLivechatChannelRuleService) GetAll() (*types.ImLivechatChannelRules, error) {
	i := &types.ImLivechatChannelRules{}
	return i, svc.client.getAll(types.ImLivechatChannelRuleModel, i)
}

func (svc *ImLivechatChannelRuleService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ImLivechatChannelRuleModel, fields, relations)
}

func (svc *ImLivechatChannelRuleService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ImLivechatChannelRuleModel, ids, fields, relations)
}

func (svc *ImLivechatChannelRuleService) Delete(ids []int) error {
	return svc.client.delete(types.ImLivechatChannelRuleModel, ids)
}
