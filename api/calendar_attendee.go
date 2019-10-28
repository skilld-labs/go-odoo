package api

import (
	"github.com/morezig/go-odoo/types"
)

type CalendarAttendeeService struct {
	client *Client
}

func NewCalendarAttendeeService(c *Client) *CalendarAttendeeService {
	return &CalendarAttendeeService{client: c}
}

func (svc *CalendarAttendeeService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.CalendarAttendeeModel, name)
}

func (svc *CalendarAttendeeService) GetByIds(ids []int64) (*types.CalendarAttendees, error) {
	c := &types.CalendarAttendees{}
	return c, svc.client.getByIds(types.CalendarAttendeeModel, ids, c)
}

func (svc *CalendarAttendeeService) GetByName(name string) (*types.CalendarAttendees, error) {
	c := &types.CalendarAttendees{}
	return c, svc.client.getByName(types.CalendarAttendeeModel, name, c)
}

func (svc *CalendarAttendeeService) GetByField(field string, value string) (*types.CalendarAttendees, error) {
	c := &types.CalendarAttendees{}
	return c, svc.client.getByField(types.CalendarAttendeeModel, field, value, c)
}

func (svc *CalendarAttendeeService) GetAll() (*types.CalendarAttendees, error) {
	c := &types.CalendarAttendees{}
	return c, svc.client.getAll(types.CalendarAttendeeModel, c)
}

func (svc *CalendarAttendeeService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.CalendarAttendeeModel, fields, relations)
}

func (svc *CalendarAttendeeService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CalendarAttendeeModel, ids, fields, relations)
}

func (svc *CalendarAttendeeService) Delete(ids []int64) error {
	return svc.client.delete(types.CalendarAttendeeModel, ids)
}
