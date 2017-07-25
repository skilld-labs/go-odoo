package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type CrmActivityService struct {
	client *Client
}

func NewCrmActivityService(c *Client) *CrmActivityService {
	return &CrmActivityService{client: c}
}

func (svc *CrmActivityService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.CrmActivityModel, name)
}

func (svc *CrmActivityService) GetByIds(ids []int) (*types.CrmActivitys, error) {
	c := &types.CrmActivitys{}
	return c, svc.client.getByIds(types.CrmActivityModel, ids, c)
}

func (svc *CrmActivityService) GetByName(name string) (*types.CrmActivitys, error) {
	c := &types.CrmActivitys{}
	return c, svc.client.getByName(types.CrmActivityModel, name, c)
}

func (svc *CrmActivityService) GetAll() (*types.CrmActivitys, error) {
	c := &types.CrmActivitys{}
	return c, svc.client.getAll(types.CrmActivityModel, c)
}

func (svc *CrmActivityService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.CrmActivityModel, fields, relations)
}

func (svc *CrmActivityService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CrmActivityModel, ids, fields, relations)
}

func (svc *CrmActivityService) Delete(ids []int) error {
	return svc.client.delete(types.CrmActivityModel, ids)
}
