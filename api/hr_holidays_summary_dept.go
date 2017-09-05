package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type HrHolidaysSummaryDeptService struct {
	client *Client
}

func NewHrHolidaysSummaryDeptService(c *Client) *HrHolidaysSummaryDeptService {
	return &HrHolidaysSummaryDeptService{client: c}
}

func (svc *HrHolidaysSummaryDeptService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.HrHolidaysSummaryDeptModel, name)
}

func (svc *HrHolidaysSummaryDeptService) GetByIds(ids []int) (*types.HrHolidaysSummaryDepts, error) {
	h := &types.HrHolidaysSummaryDepts{}
	return h, svc.client.getByIds(types.HrHolidaysSummaryDeptModel, ids, h)
}

func (svc *HrHolidaysSummaryDeptService) GetByName(name string) (*types.HrHolidaysSummaryDepts, error) {
	h := &types.HrHolidaysSummaryDepts{}
	return h, svc.client.getByName(types.HrHolidaysSummaryDeptModel, name, h)
}

func (svc *HrHolidaysSummaryDeptService) GetByField(field string, value string) (*types.HrHolidaysSummaryDepts, error) {
	h := &types.HrHolidaysSummaryDepts{}
	return h, svc.client.getByName(types.HrHolidaysSummaryDeptModel, field, value, h)
}

func (svc *HrHolidaysSummaryDeptService) GetAll() (*types.HrHolidaysSummaryDepts, error) {
	h := &types.HrHolidaysSummaryDepts{}
	return h, svc.client.getAll(types.HrHolidaysSummaryDeptModel, h)
}

func (svc *HrHolidaysSummaryDeptService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.HrHolidaysSummaryDeptModel, fields, relations)
}

func (svc *HrHolidaysSummaryDeptService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.HrHolidaysSummaryDeptModel, ids, fields, relations)
}

func (svc *HrHolidaysSummaryDeptService) Delete(ids []int) error {
	return svc.client.delete(types.HrHolidaysSummaryDeptModel, ids)
}
