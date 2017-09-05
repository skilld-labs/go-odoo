package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ResCountryGroupService struct {
	client *Client
}

func NewResCountryGroupService(c *Client) *ResCountryGroupService {
	return &ResCountryGroupService{client: c}
}

func (svc *ResCountryGroupService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ResCountryGroupModel, name)
}

func (svc *ResCountryGroupService) GetByIds(ids []int) (*types.ResCountryGroups, error) {
	r := &types.ResCountryGroups{}
	return r, svc.client.getByIds(types.ResCountryGroupModel, ids, r)
}

func (svc *ResCountryGroupService) GetByName(name string) (*types.ResCountryGroups, error) {
	r := &types.ResCountryGroups{}
	return r, svc.client.getByName(types.ResCountryGroupModel, name, r)
}

func (svc *ResCountryGroupService) GetByField(field string, value string) (*types.ResCountryGroups, error) {
	r := &types.ResCountryGroups{}
	return r, svc.client.getByField(types.ResCountryGroupModel, field, value, r)
}

func (svc *ResCountryGroupService) GetAll() (*types.ResCountryGroups, error) {
	r := &types.ResCountryGroups{}
	return r, svc.client.getAll(types.ResCountryGroupModel, r)
}

func (svc *ResCountryGroupService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ResCountryGroupModel, fields, relations)
}

func (svc *ResCountryGroupService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResCountryGroupModel, ids, fields, relations)
}

func (svc *ResCountryGroupService) Delete(ids []int) error {
	return svc.client.delete(types.ResCountryGroupModel, ids)
}
