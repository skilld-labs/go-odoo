package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type CalendarEventService struct {
	client *Client
}

func NewCalendarEventService(c *Client) *CalendarEventService {
	return &CalendarEventService{client: c}
}

func (svc *CalendarEventService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.CalendarEventModel, name)
}

func (svc *CalendarEventService) GetByIds(ids []int) (*types.CalendarEvents, error) {
	c := &types.CalendarEvents{}
	return c, svc.client.getByIds(types.CalendarEventModel, ids, c)
}

func (svc *CalendarEventService) GetByName(name string) (*types.CalendarEvents, error) {
	c := &types.CalendarEvents{}
	return c, svc.client.getByName(types.CalendarEventModel, name, c)
}

func (svc *CalendarEventService) GetByField(field string, value string) (*types.CalendarEvents, error) {
	c := &types.CalendarEvents{}
	return c, svc.client.getByName(types.CalendarEventModel, field, value, c)
}

func (svc *CalendarEventService) GetAll() (*types.CalendarEvents, error) {
	c := &types.CalendarEvents{}
	return c, svc.client.getAll(types.CalendarEventModel, c)
}

func (svc *CalendarEventService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.CalendarEventModel, fields, relations)
}

func (svc *CalendarEventService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CalendarEventModel, ids, fields, relations)
}

func (svc *CalendarEventService) Delete(ids []int) error {
	return svc.client.delete(types.CalendarEventModel, ids)
}
