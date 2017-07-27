package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseActionRuleService struct {
	client *Client
}

func NewBaseActionRuleService(c *Client) *BaseActionRuleService {
	return &BaseActionRuleService{client: c}
}

func (svc *BaseActionRuleService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BaseActionRuleModel, name)
}

func (svc *BaseActionRuleService) GetByIds(ids []int) (*types.BaseActionRules, error) {
	b := &types.BaseActionRules{}
	return b, svc.client.getByIds(types.BaseActionRuleModel, ids, b)
}

func (svc *BaseActionRuleService) GetByName(name string) (*types.BaseActionRules, error) {
	b := &types.BaseActionRules{}
	return b, svc.client.getByName(types.BaseActionRuleModel, name, b)
}

func (svc *BaseActionRuleService) GetAll() (*types.BaseActionRules, error) {
	b := &types.BaseActionRules{}
	return b, svc.client.getAll(types.BaseActionRuleModel, b)
}

func (svc *BaseActionRuleService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BaseActionRuleModel, fields, relations)
}

func (svc *BaseActionRuleService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseActionRuleModel, ids, fields, relations)
}

func (svc *BaseActionRuleService) Delete(ids []int) error {
	return svc.client.delete(types.BaseActionRuleModel, ids)
}
