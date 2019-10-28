package api

import (
	"github.com/morezig/go-odoo/types"
)

type CrmLostReasonService struct {
	client *Client
}

func NewCrmLostReasonService(c *Client) *CrmLostReasonService {
	return &CrmLostReasonService{client: c}
}

func (svc *CrmLostReasonService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.CrmLostReasonModel, name)
}

func (svc *CrmLostReasonService) GetByIds(ids []int64) (*types.CrmLostReasons, error) {
	c := &types.CrmLostReasons{}
	return c, svc.client.getByIds(types.CrmLostReasonModel, ids, c)
}

func (svc *CrmLostReasonService) GetByName(name string) (*types.CrmLostReasons, error) {
	c := &types.CrmLostReasons{}
	return c, svc.client.getByName(types.CrmLostReasonModel, name, c)
}

func (svc *CrmLostReasonService) GetByField(field string, value string) (*types.CrmLostReasons, error) {
	c := &types.CrmLostReasons{}
	return c, svc.client.getByField(types.CrmLostReasonModel, field, value, c)
}

func (svc *CrmLostReasonService) GetAll() (*types.CrmLostReasons, error) {
	c := &types.CrmLostReasons{}
	return c, svc.client.getAll(types.CrmLostReasonModel, c)
}

func (svc *CrmLostReasonService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.CrmLostReasonModel, fields, relations)
}

func (svc *CrmLostReasonService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CrmLostReasonModel, ids, fields, relations)
}

func (svc *CrmLostReasonService) Delete(ids []int64) error {
	return svc.client.delete(types.CrmLostReasonModel, ids)
}
