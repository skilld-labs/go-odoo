package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BarcodeRuleService struct {
	client *Client
}

func NewBarcodeRuleService(c *Client) *BarcodeRuleService {
	return &BarcodeRuleService{client: c}
}

func (svc *BarcodeRuleService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BarcodeRuleModel, name)
}

func (svc *BarcodeRuleService) GetByIds(ids []int) (*types.BarcodeRules, error) {
	b := &types.BarcodeRules{}
	return b, svc.client.getByIds(types.BarcodeRuleModel, ids, b)
}

func (svc *BarcodeRuleService) GetByName(name string) (*types.BarcodeRules, error) {
	b := &types.BarcodeRules{}
	return b, svc.client.getByName(types.BarcodeRuleModel, name, b)
}

func (svc *BarcodeRuleService) GetAll() (*types.BarcodeRules, error) {
	b := &types.BarcodeRules{}
	return b, svc.client.getAll(types.BarcodeRuleModel, b)
}

func (svc *BarcodeRuleService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BarcodeRuleModel, fields, relations)
}

func (svc *BarcodeRuleService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BarcodeRuleModel, ids, fields, relations)
}

func (svc *BarcodeRuleService) Delete(ids []int) error {
	return svc.client.delete(types.BarcodeRuleModel, ids)
}
