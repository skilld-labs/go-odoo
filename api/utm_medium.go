package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type UtmMediumService struct {
	client *Client
}

func NewUtmMediumService(c *Client) *UtmMediumService {
	return &UtmMediumService{client: c}
}

func (svc *UtmMediumService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.UtmMediumModel, name)
}

func (svc *UtmMediumService) GetByIds(ids []int) (*types.UtmMediums, error) {
	u := &types.UtmMediums{}
	return u, svc.client.getByIds(types.UtmMediumModel, ids, u)
}

func (svc *UtmMediumService) GetByName(name string) (*types.UtmMediums, error) {
	u := &types.UtmMediums{}
	return u, svc.client.getByName(types.UtmMediumModel, name, u)
}

func (svc *UtmMediumService) GetByField(field string, value string) (*types.UtmMediums, error) {
	u := &types.UtmMediums{}
	return u, svc.client.getByField(types.UtmMediumModel, field, value, u)
}

func (svc *UtmMediumService) GetAll() (*types.UtmMediums, error) {
	u := &types.UtmMediums{}
	return u, svc.client.getAll(types.UtmMediumModel, u)
}

func (svc *UtmMediumService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.UtmMediumModel, fields, relations)
}

func (svc *UtmMediumService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.UtmMediumModel, ids, fields, relations)
}

func (svc *UtmMediumService) Delete(ids []int) error {
	return svc.client.delete(types.UtmMediumModel, ids)
}
