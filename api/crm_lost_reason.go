package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type CrmLostReasonService struct {
	client *Client
}

func NewCrmLostReasonService(c *Client) *CrmLostReasonService {
	return &CrmLostReasonService{client: c}
}

func (svc *CrmLostReasonService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.CrmLostReasonModel, name)
}

func (svc *CrmLostReasonService) GetByIds(ids []int) (*types.CrmLostReasons, error) {
	c := &types.CrmLostReasons{}
	return c, svc.client.getByIds(types.CrmLostReasonModel, ids, c)
}

func (svc *CrmLostReasonService) GetByName(name string) (*types.CrmLostReasons, error) {
	c := &types.CrmLostReasons{}
	return c, svc.client.getByName(types.CrmLostReasonModel, name, c)
}

func (svc *CrmLostReasonService) GetByField(field string, value string) (*types.CrmLostReasons, error) {
	c := &types.CrmLostReasons{}
	return c, svc.client.getByName(types.CrmLostReasonModel, field, value, c)
}

func (svc *CrmLostReasonService) GetAll() (*types.CrmLostReasons, error) {
	c := &types.CrmLostReasons{}
	return c, svc.client.getAll(types.CrmLostReasonModel, c)
}

func (svc *CrmLostReasonService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.CrmLostReasonModel, fields, relations)
}

func (svc *CrmLostReasonService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CrmLostReasonModel, ids, fields, relations)
}

func (svc *CrmLostReasonService) Delete(ids []int) error {
	return svc.client.delete(types.CrmLostReasonModel, ids)
}
