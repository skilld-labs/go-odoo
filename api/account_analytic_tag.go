package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountAnalyticTagService struct {
	client *Client
}

func NewAccountAnalyticTagService(c *Client) *AccountAnalyticTagService {
	return &AccountAnalyticTagService{client: c}
}

func (svc *AccountAnalyticTagService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountAnalyticTagModel, name)
}

func (svc *AccountAnalyticTagService) GetByIds(ids []int) (*types.AccountAnalyticTags, error) {
	a := &types.AccountAnalyticTags{}
	return a, svc.client.getByIds(types.AccountAnalyticTagModel, ids, a)
}

func (svc *AccountAnalyticTagService) GetByName(name string) (*types.AccountAnalyticTags, error) {
	a := &types.AccountAnalyticTags{}
	return a, svc.client.getByName(types.AccountAnalyticTagModel, name, a)
}

func (svc *AccountAnalyticTagService) GetByField(field string, value string) (*types.AccountAnalyticTags, error) {
	a := &types.AccountAnalyticTags{}
	return a, svc.client.getByField(types.AccountAnalyticTagModel, field, value, a)
}

func (svc *AccountAnalyticTagService) GetAll() (*types.AccountAnalyticTags, error) {
	a := &types.AccountAnalyticTags{}
	return a, svc.client.getAll(types.AccountAnalyticTagModel, a)
}

func (svc *AccountAnalyticTagService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountAnalyticTagModel, fields, relations)
}

func (svc *AccountAnalyticTagService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountAnalyticTagModel, ids, fields, relations)
}

func (svc *AccountAnalyticTagService) Delete(ids []int) error {
	return svc.client.delete(types.AccountAnalyticTagModel, ids)
}
