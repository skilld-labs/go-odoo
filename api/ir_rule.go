package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrRuleService struct {
	client *Client
}

func NewIrRuleService(c *Client) *IrRuleService {
	return &IrRuleService{client: c}
}

func (svc *IrRuleService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrRuleModel, name)
}

func (svc *IrRuleService) GetByIds(ids []int) (*types.IrRules, error) {
	i := &types.IrRules{}
	return i, svc.client.getByIds(types.IrRuleModel, ids, i)
}

func (svc *IrRuleService) GetByName(name string) (*types.IrRules, error) {
	i := &types.IrRules{}
	return i, svc.client.getByName(types.IrRuleModel, name, i)
}

func (svc *IrRuleService) GetByField(field string, value string) (*types.IrRules, error) {
	i := &types.IrRules{}
	return i, svc.client.getByField(types.IrRuleModel, field, value, i)
}

func (svc *IrRuleService) GetAll() (*types.IrRules, error) {
	i := &types.IrRules{}
	return i, svc.client.getAll(types.IrRuleModel, i)
}

func (svc *IrRuleService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrRuleModel, fields, relations)
}

func (svc *IrRuleService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrRuleModel, ids, fields, relations)
}

func (svc *IrRuleService) Delete(ids []int) error {
	return svc.client.delete(types.IrRuleModel, ids)
}
