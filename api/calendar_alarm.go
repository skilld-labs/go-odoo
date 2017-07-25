package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type CalendarAlarmService struct {
	client *Client
}

func NewCalendarAlarmService(c *Client) *CalendarAlarmService {
	return &CalendarAlarmService{client: c}
}

func (svc *CalendarAlarmService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.CalendarAlarmModel, name)
}

func (svc *CalendarAlarmService) GetByIds(ids []int) (*types.CalendarAlarms, error) {
	c := &types.CalendarAlarms{}
	return c, svc.client.getByIds(types.CalendarAlarmModel, ids, c)
}

func (svc *CalendarAlarmService) GetByName(name string) (*types.CalendarAlarms, error) {
	c := &types.CalendarAlarms{}
	return c, svc.client.getByName(types.CalendarAlarmModel, name, c)
}

func (svc *CalendarAlarmService) GetAll() (*types.CalendarAlarms, error) {
	c := &types.CalendarAlarms{}
	return c, svc.client.getAll(types.CalendarAlarmModel, c)
}

func (svc *CalendarAlarmService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.CalendarAlarmModel, fields, relations)
}

func (svc *CalendarAlarmService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CalendarAlarmModel, ids, fields, relations)
}

func (svc *CalendarAlarmService) Delete(ids []int) error {
	return svc.client.delete(types.CalendarAlarmModel, ids)
}
