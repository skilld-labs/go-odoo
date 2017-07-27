package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseActionRuleLeadTestService struct {
	client *Client
}

func NewBaseActionRuleLeadTestService(c *Client) *BaseActionRuleLeadTestService {
	return &BaseActionRuleLeadTestService{client: c}
}

func (svc *BaseActionRuleLeadTestService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BaseActionRuleLeadTestModel, name)
}

func (svc *BaseActionRuleLeadTestService) GetByIds(ids []int) (*types.BaseActionRuleLeadTests, error) {
	b := &types.BaseActionRuleLeadTests{}
	return b, svc.client.getByIds(types.BaseActionRuleLeadTestModel, ids, b)
}

func (svc *BaseActionRuleLeadTestService) GetByName(name string) (*types.BaseActionRuleLeadTests, error) {
	b := &types.BaseActionRuleLeadTests{}
	return b, svc.client.getByName(types.BaseActionRuleLeadTestModel, name, b)
}

func (svc *BaseActionRuleLeadTestService) GetAll() (*types.BaseActionRuleLeadTests, error) {
	b := &types.BaseActionRuleLeadTests{}
	return b, svc.client.getAll(types.BaseActionRuleLeadTestModel, b)
}

func (svc *BaseActionRuleLeadTestService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BaseActionRuleLeadTestModel, fields, relations)
}

func (svc *BaseActionRuleLeadTestService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseActionRuleLeadTestModel, ids, fields, relations)
}

func (svc *BaseActionRuleLeadTestService) Delete(ids []int) error {
	return svc.client.delete(types.BaseActionRuleLeadTestModel, ids)
}
