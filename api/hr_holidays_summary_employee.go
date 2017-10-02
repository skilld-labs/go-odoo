package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type HrHolidaysSummaryEmployeeService struct {
	client *Client
}

func NewHrHolidaysSummaryEmployeeService(c *Client) *HrHolidaysSummaryEmployeeService {
	return &HrHolidaysSummaryEmployeeService{client: c}
}

func (svc *HrHolidaysSummaryEmployeeService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.HrHolidaysSummaryEmployeeModel, name)
}

func (svc *HrHolidaysSummaryEmployeeService) GetByIds(ids []int64) (*types.HrHolidaysSummaryEmployees, error) {
	h := &types.HrHolidaysSummaryEmployees{}
	return h, svc.client.getByIds(types.HrHolidaysSummaryEmployeeModel, ids, h)
}

func (svc *HrHolidaysSummaryEmployeeService) GetByName(name string) (*types.HrHolidaysSummaryEmployees, error) {
	h := &types.HrHolidaysSummaryEmployees{}
	return h, svc.client.getByName(types.HrHolidaysSummaryEmployeeModel, name, h)
}

func (svc *HrHolidaysSummaryEmployeeService) GetByField(field string, value string) (*types.HrHolidaysSummaryEmployees, error) {
	h := &types.HrHolidaysSummaryEmployees{}
	return h, svc.client.getByField(types.HrHolidaysSummaryEmployeeModel, field, value, h)
}

func (svc *HrHolidaysSummaryEmployeeService) GetAll() (*types.HrHolidaysSummaryEmployees, error) {
	h := &types.HrHolidaysSummaryEmployees{}
	return h, svc.client.getAll(types.HrHolidaysSummaryEmployeeModel, h)
}

func (svc *HrHolidaysSummaryEmployeeService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.HrHolidaysSummaryEmployeeModel, fields, relations)
}

func (svc *HrHolidaysSummaryEmployeeService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.HrHolidaysSummaryEmployeeModel, ids, fields, relations)
}

func (svc *HrHolidaysSummaryEmployeeService) Delete(ids []int64) error {
	return svc.client.delete(types.HrHolidaysSummaryEmployeeModel, ids)
}
