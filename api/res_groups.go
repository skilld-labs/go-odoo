package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ResGroupsService struct {
	client *Client
}

func NewResGroupsService(c *Client) *ResGroupsService {
	return &ResGroupsService{client: c}
}

func (svc *ResGroupsService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ResGroupsModel, name)
}

func (svc *ResGroupsService) GetByIds(ids []int) (*types.ResGroupss, error) {
	r := &types.ResGroupss{}
	return r, svc.client.getByIds(types.ResGroupsModel, ids, r)
}

func (svc *ResGroupsService) GetByName(name string) (*types.ResGroupss, error) {
	r := &types.ResGroupss{}
	return r, svc.client.getByName(types.ResGroupsModel, name, r)
}

func (svc *ResGroupsService) GetByField(field string, value string) (*types.ResGroupss, error) {
	r := &types.ResGroupss{}
	return r, svc.client.getByField(types.ResGroupsModel, field, value, r)
}

func (svc *ResGroupsService) GetAll() (*types.ResGroupss, error) {
	r := &types.ResGroupss{}
	return r, svc.client.getAll(types.ResGroupsModel, r)
}

func (svc *ResGroupsService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ResGroupsModel, fields, relations)
}

func (svc *ResGroupsService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResGroupsModel, ids, fields, relations)
}

func (svc *ResGroupsService) Delete(ids []int) error {
	return svc.client.delete(types.ResGroupsModel, ids)
}
