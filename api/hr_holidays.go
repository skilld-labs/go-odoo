package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type HrHolidaysService struct {
	client *Client
}

func NewHrHolidaysService(c *Client) *HrHolidaysService {
	return &HrHolidaysService{client: c}
}

func (svc *HrHolidaysService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.HrHolidaysModel, name)
}

func (svc *HrHolidaysService) GetByIds(ids []int64) (*types.HrHolidayss, error) {
	h := &types.HrHolidayss{}
	return h, svc.client.getByIds(types.HrHolidaysModel, ids, h)
}

func (svc *HrHolidaysService) GetByName(name string) (*types.HrHolidayss, error) {
	h := &types.HrHolidayss{}
	return h, svc.client.getByName(types.HrHolidaysModel, name, h)
}

func (svc *HrHolidaysService) GetByField(field string, value string) (*types.HrHolidayss, error) {
	h := &types.HrHolidayss{}
	return h, svc.client.getByField(types.HrHolidaysModel, field, value, h)
}

func (svc *HrHolidaysService) GetAll() (*types.HrHolidayss, error) {
	h := &types.HrHolidayss{}
	return h, svc.client.getAll(types.HrHolidaysModel, h)
}

func (svc *HrHolidaysService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.HrHolidaysModel, fields, relations)
}

func (svc *HrHolidaysService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.HrHolidaysModel, ids, fields, relations)
}

func (svc *HrHolidaysService) Delete(ids []int64) error {
	return svc.client.delete(types.HrHolidaysModel, ids)
}
