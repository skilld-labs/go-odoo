package api

import (
	"github.com/morezig/go-odoo/types"
)

type CalendarEventTypeService struct {
	client *Client
}

func NewCalendarEventTypeService(c *Client) *CalendarEventTypeService {
	return &CalendarEventTypeService{client: c}
}

func (svc *CalendarEventTypeService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.CalendarEventTypeModel, name)
}

func (svc *CalendarEventTypeService) GetByIds(ids []int64) (*types.CalendarEventTypes, error) {
	c := &types.CalendarEventTypes{}
	return c, svc.client.getByIds(types.CalendarEventTypeModel, ids, c)
}

func (svc *CalendarEventTypeService) GetByName(name string) (*types.CalendarEventTypes, error) {
	c := &types.CalendarEventTypes{}
	return c, svc.client.getByName(types.CalendarEventTypeModel, name, c)
}

func (svc *CalendarEventTypeService) GetByField(field string, value string) (*types.CalendarEventTypes, error) {
	c := &types.CalendarEventTypes{}
	return c, svc.client.getByField(types.CalendarEventTypeModel, field, value, c)
}

func (svc *CalendarEventTypeService) GetAll() (*types.CalendarEventTypes, error) {
	c := &types.CalendarEventTypes{}
	return c, svc.client.getAll(types.CalendarEventTypeModel, c)
}

func (svc *CalendarEventTypeService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.CalendarEventTypeModel, fields, relations)
}

func (svc *CalendarEventTypeService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CalendarEventTypeModel, ids, fields, relations)
}

func (svc *CalendarEventTypeService) Delete(ids []int64) error {
	return svc.client.delete(types.CalendarEventTypeModel, ids)
}
