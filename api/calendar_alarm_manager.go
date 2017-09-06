package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type CalendarAlarmManagerService struct {
	client *Client
}

func NewCalendarAlarmManagerService(c *Client) *CalendarAlarmManagerService {
	return &CalendarAlarmManagerService{client: c}
}

func (svc *CalendarAlarmManagerService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.CalendarAlarmManagerModel, name)
}

func (svc *CalendarAlarmManagerService) GetByIds(ids []int) (*types.CalendarAlarmManagers, error) {
	c := &types.CalendarAlarmManagers{}
	return c, svc.client.getByIds(types.CalendarAlarmManagerModel, ids, c)
}

func (svc *CalendarAlarmManagerService) GetByName(name string) (*types.CalendarAlarmManagers, error) {
	c := &types.CalendarAlarmManagers{}
	return c, svc.client.getByName(types.CalendarAlarmManagerModel, name, c)
}

func (svc *CalendarAlarmManagerService) GetByField(field string, value string) (*types.CalendarAlarmManagers, error) {
	c := &types.CalendarAlarmManagers{}
	return c, svc.client.getByField(types.CalendarAlarmManagerModel, field, value, c)
}

func (svc *CalendarAlarmManagerService) GetAll() (*types.CalendarAlarmManagers, error) {
	c := &types.CalendarAlarmManagers{}
	return c, svc.client.getAll(types.CalendarAlarmManagerModel, c)
}

func (svc *CalendarAlarmManagerService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.CalendarAlarmManagerModel, fields, relations)
}

func (svc *CalendarAlarmManagerService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CalendarAlarmManagerModel, ids, fields, relations)
}

func (svc *CalendarAlarmManagerService) Delete(ids []int) error {
	return svc.client.delete(types.CalendarAlarmManagerModel, ids)
}
