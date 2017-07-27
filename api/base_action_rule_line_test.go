package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseActionRuleLineTestService struct {
	client *Client
}

func NewBaseActionRuleLineTestService(c *Client) *BaseActionRuleLineTestService {
	return &BaseActionRuleLineTestService{client: c}
}

func (svc *BaseActionRuleLineTestService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BaseActionRuleLineTestModel, name)
}

func (svc *BaseActionRuleLineTestService) GetByIds(ids []int) (*types.BaseActionRuleLineTests, error) {
	b := &types.BaseActionRuleLineTests{}
	return b, svc.client.getByIds(types.BaseActionRuleLineTestModel, ids, b)
}

func (svc *BaseActionRuleLineTestService) GetByName(name string) (*types.BaseActionRuleLineTests, error) {
	b := &types.BaseActionRuleLineTests{}
	return b, svc.client.getByName(types.BaseActionRuleLineTestModel, name, b)
}

func (svc *BaseActionRuleLineTestService) GetAll() (*types.BaseActionRuleLineTests, error) {
	b := &types.BaseActionRuleLineTests{}
	return b, svc.client.getAll(types.BaseActionRuleLineTestModel, b)
}

func (svc *BaseActionRuleLineTestService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BaseActionRuleLineTestModel, fields, relations)
}

func (svc *BaseActionRuleLineTestService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseActionRuleLineTestModel, ids, fields, relations)
}

func (svc *BaseActionRuleLineTestService) Delete(ids []int) error {
	return svc.client.delete(types.BaseActionRuleLineTestModel, ids)
}
