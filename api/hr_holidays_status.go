package api

import (
	"github.com/morezig/go-odoo/types"
)

type HrHolidaysStatusService struct {
	client *Client
}

func NewHrHolidaysStatusService(c *Client) *HrHolidaysStatusService {
	return &HrHolidaysStatusService{client: c}
}

func (svc *HrHolidaysStatusService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.HrHolidaysStatusModel, name)
}

func (svc *HrHolidaysStatusService) GetByIds(ids []int64) (*types.HrHolidaysStatuss, error) {
	h := &types.HrHolidaysStatuss{}
	return h, svc.client.getByIds(types.HrHolidaysStatusModel, ids, h)
}

func (svc *HrHolidaysStatusService) GetByName(name string) (*types.HrHolidaysStatuss, error) {
	h := &types.HrHolidaysStatuss{}
	return h, svc.client.getByName(types.HrHolidaysStatusModel, name, h)
}

func (svc *HrHolidaysStatusService) GetByField(field string, value string) (*types.HrHolidaysStatuss, error) {
	h := &types.HrHolidaysStatuss{}
	return h, svc.client.getByField(types.HrHolidaysStatusModel, field, value, h)
}

func (svc *HrHolidaysStatusService) GetAll() (*types.HrHolidaysStatuss, error) {
	h := &types.HrHolidaysStatuss{}
	return h, svc.client.getAll(types.HrHolidaysStatusModel, h)
}

func (svc *HrHolidaysStatusService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.HrHolidaysStatusModel, fields, relations)
}

func (svc *HrHolidaysStatusService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.HrHolidaysStatusModel, ids, fields, relations)
}

func (svc *HrHolidaysStatusService) Delete(ids []int64) error {
	return svc.client.delete(types.HrHolidaysStatusModel, ids)
}
