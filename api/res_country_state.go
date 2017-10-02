package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ResCountryStateService struct {
	client *Client
}

func NewResCountryStateService(c *Client) *ResCountryStateService {
	return &ResCountryStateService{client: c}
}

func (svc *ResCountryStateService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ResCountryStateModel, name)
}

func (svc *ResCountryStateService) GetByIds(ids []int64) (*types.ResCountryStates, error) {
	r := &types.ResCountryStates{}
	return r, svc.client.getByIds(types.ResCountryStateModel, ids, r)
}

func (svc *ResCountryStateService) GetByName(name string) (*types.ResCountryStates, error) {
	r := &types.ResCountryStates{}
	return r, svc.client.getByName(types.ResCountryStateModel, name, r)
}

func (svc *ResCountryStateService) GetByField(field string, value string) (*types.ResCountryStates, error) {
	r := &types.ResCountryStates{}
	return r, svc.client.getByField(types.ResCountryStateModel, field, value, r)
}

func (svc *ResCountryStateService) GetAll() (*types.ResCountryStates, error) {
	r := &types.ResCountryStates{}
	return r, svc.client.getAll(types.ResCountryStateModel, r)
}

func (svc *ResCountryStateService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ResCountryStateModel, fields, relations)
}

func (svc *ResCountryStateService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResCountryStateModel, ids, fields, relations)
}

func (svc *ResCountryStateService) Delete(ids []int64) error {
	return svc.client.delete(types.ResCountryStateModel, ids)
}
