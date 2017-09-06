package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type CrmStageService struct {
	client *Client
}

func NewCrmStageService(c *Client) *CrmStageService {
	return &CrmStageService{client: c}
}

func (svc *CrmStageService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.CrmStageModel, name)
}

func (svc *CrmStageService) GetByIds(ids []int) (*types.CrmStages, error) {
	c := &types.CrmStages{}
	return c, svc.client.getByIds(types.CrmStageModel, ids, c)
}

func (svc *CrmStageService) GetByName(name string) (*types.CrmStages, error) {
	c := &types.CrmStages{}
	return c, svc.client.getByName(types.CrmStageModel, name, c)
}

func (svc *CrmStageService) GetByField(field string, value string) (*types.CrmStages, error) {
	c := &types.CrmStages{}
	return c, svc.client.getByField(types.CrmStageModel, field, value, c)
}

func (svc *CrmStageService) GetAll() (*types.CrmStages, error) {
	c := &types.CrmStages{}
	return c, svc.client.getAll(types.CrmStageModel, c)
}

func (svc *CrmStageService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.CrmStageModel, fields, relations)
}

func (svc *CrmStageService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CrmStageModel, ids, fields, relations)
}

func (svc *CrmStageService) Delete(ids []int) error {
	return svc.client.delete(types.CrmStageModel, ids)
}
