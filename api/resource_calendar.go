package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ResourceCalendarService struct {
	client *Client
}

func NewResourceCalendarService(c *Client) *ResourceCalendarService {
	return &ResourceCalendarService{client: c}
}

func (svc *ResourceCalendarService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ResourceCalendarModel, name)
}

func (svc *ResourceCalendarService) GetByIds(ids []int) (*types.ResourceCalendars, error) {
	r := &types.ResourceCalendars{}
	return r, svc.client.getByIds(types.ResourceCalendarModel, ids, r)
}

func (svc *ResourceCalendarService) GetByName(name string) (*types.ResourceCalendars, error) {
	r := &types.ResourceCalendars{}
	return r, svc.client.getByName(types.ResourceCalendarModel, name, r)
}

func (svc *ResourceCalendarService) GetAll() (*types.ResourceCalendars, error) {
	r := &types.ResourceCalendars{}
	return r, svc.client.getAll(types.ResourceCalendarModel, r)
}

func (svc *ResourceCalendarService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ResourceCalendarModel, fields, relations)
}

func (svc *ResourceCalendarService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResourceCalendarModel, ids, fields, relations)
}

func (svc *ResourceCalendarService) Delete(ids []int) error {
	return svc.client.delete(types.ResourceCalendarModel, ids)
}
