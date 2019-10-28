package api

import (
	"github.com/morezig/go-odoo/types"
)

type ResourceCalendarAttendanceService struct {
	client *Client
}

func NewResourceCalendarAttendanceService(c *Client) *ResourceCalendarAttendanceService {
	return &ResourceCalendarAttendanceService{client: c}
}

func (svc *ResourceCalendarAttendanceService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ResourceCalendarAttendanceModel, name)
}

func (svc *ResourceCalendarAttendanceService) GetByIds(ids []int64) (*types.ResourceCalendarAttendances, error) {
	r := &types.ResourceCalendarAttendances{}
	return r, svc.client.getByIds(types.ResourceCalendarAttendanceModel, ids, r)
}

func (svc *ResourceCalendarAttendanceService) GetByName(name string) (*types.ResourceCalendarAttendances, error) {
	r := &types.ResourceCalendarAttendances{}
	return r, svc.client.getByName(types.ResourceCalendarAttendanceModel, name, r)
}

func (svc *ResourceCalendarAttendanceService) GetByField(field string, value string) (*types.ResourceCalendarAttendances, error) {
	r := &types.ResourceCalendarAttendances{}
	return r, svc.client.getByField(types.ResourceCalendarAttendanceModel, field, value, r)
}

func (svc *ResourceCalendarAttendanceService) GetAll() (*types.ResourceCalendarAttendances, error) {
	r := &types.ResourceCalendarAttendances{}
	return r, svc.client.getAll(types.ResourceCalendarAttendanceModel, r)
}

func (svc *ResourceCalendarAttendanceService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ResourceCalendarAttendanceModel, fields, relations)
}

func (svc *ResourceCalendarAttendanceService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ResourceCalendarAttendanceModel, ids, fields, relations)
}

func (svc *ResourceCalendarAttendanceService) Delete(ids []int64) error {
	return svc.client.delete(types.ResourceCalendarAttendanceModel, ids)
}
