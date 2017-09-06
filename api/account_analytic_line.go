package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountAnalyticLineService struct {
	client *Client
}

func NewAccountAnalyticLineService(c *Client) *AccountAnalyticLineService {
	return &AccountAnalyticLineService{client: c}
}

func (svc *AccountAnalyticLineService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountAnalyticLineModel, name)
}

func (svc *AccountAnalyticLineService) GetByIds(ids []int) (*types.AccountAnalyticLines, error) {
	a := &types.AccountAnalyticLines{}
	return a, svc.client.getByIds(types.AccountAnalyticLineModel, ids, a)
}

func (svc *AccountAnalyticLineService) GetByName(name string) (*types.AccountAnalyticLines, error) {
	a := &types.AccountAnalyticLines{}
	return a, svc.client.getByName(types.AccountAnalyticLineModel, name, a)
}

func (svc *AccountAnalyticLineService) GetByField(field string, value string) (*types.AccountAnalyticLines, error) {
	a := &types.AccountAnalyticLines{}
	return a, svc.client.getByField(types.AccountAnalyticLineModel, field, value, a)
}

func (svc *AccountAnalyticLineService) GetAll() (*types.AccountAnalyticLines, error) {
	a := &types.AccountAnalyticLines{}
	return a, svc.client.getAll(types.AccountAnalyticLineModel, a)
}

func (svc *AccountAnalyticLineService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountAnalyticLineModel, fields, relations)
}

func (svc *AccountAnalyticLineService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountAnalyticLineModel, ids, fields, relations)
}

func (svc *AccountAnalyticLineService) Delete(ids []int) error {
	return svc.client.delete(types.AccountAnalyticLineModel, ids)
}
