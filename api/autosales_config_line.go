package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AutosalesConfigLineService struct {
	client *Client
}

func NewAutosalesConfigLineService(c *Client) *AutosalesConfigLineService {
	return &AutosalesConfigLineService{client: c}
}

func (svc *AutosalesConfigLineService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AutosalesConfigLineModel, name)
}

func (svc *AutosalesConfigLineService) GetByIds(ids []int) (*types.AutosalesConfigLines, error) {
	a := &types.AutosalesConfigLines{}
	return a, svc.client.getByIds(types.AutosalesConfigLineModel, ids, a)
}

func (svc *AutosalesConfigLineService) GetByName(name string) (*types.AutosalesConfigLines, error) {
	a := &types.AutosalesConfigLines{}
	return a, svc.client.getByName(types.AutosalesConfigLineModel, name, a)
}

func (svc *AutosalesConfigLineService) GetByField(field string, value string) (*types.AutosalesConfigLines, error) {
	a := &types.AutosalesConfigLines{}
	return a, svc.client.getByField(types.AutosalesConfigLineModel, field, value, a)
}

func (svc *AutosalesConfigLineService) GetAll() (*types.AutosalesConfigLines, error) {
	a := &types.AutosalesConfigLines{}
	return a, svc.client.getAll(types.AutosalesConfigLineModel, a)
}

func (svc *AutosalesConfigLineService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AutosalesConfigLineModel, fields, relations)
}

func (svc *AutosalesConfigLineService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AutosalesConfigLineModel, ids, fields, relations)
}

func (svc *AutosalesConfigLineService) Delete(ids []int) error {
	return svc.client.delete(types.AutosalesConfigLineModel, ids)
}
