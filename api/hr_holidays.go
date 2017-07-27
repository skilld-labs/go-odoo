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

func (svc *HrHolidaysService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.HrHolidaysModel, name)
}

func (svc *HrHolidaysService) GetByIds(ids []int) (*types.HrHolidayss, error) {
	h := &types.HrHolidayss{}
	return h, svc.client.getByIds(types.HrHolidaysModel, ids, h)
}

func (svc *HrHolidaysService) GetByName(name string) (*types.HrHolidayss, error) {
	h := &types.HrHolidayss{}
	return h, svc.client.getByName(types.HrHolidaysModel, name, h)
}

func (svc *HrHolidaysService) GetAll() (*types.HrHolidayss, error) {
	h := &types.HrHolidayss{}
	return h, svc.client.getAll(types.HrHolidaysModel, h)
}

func (svc *HrHolidaysService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.HrHolidaysModel, fields, relations)
}

func (svc *HrHolidaysService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.HrHolidaysModel, ids, fields, relations)
}

func (svc *HrHolidaysService) Delete(ids []int) error {
	return svc.client.delete(types.HrHolidaysModel, ids)
}
