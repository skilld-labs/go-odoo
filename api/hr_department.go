package api

import (
	"github.com/morezig/go-odoo/types"
)

type HrDepartmentService struct {
	client *Client
}

func NewHrDepartmentService(c *Client) *HrDepartmentService {
	return &HrDepartmentService{client: c}
}

func (svc *HrDepartmentService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.HrDepartmentModel, name)
}

func (svc *HrDepartmentService) GetByIds(ids []int64) (*types.HrDepartments, error) {
	h := &types.HrDepartments{}
	return h, svc.client.getByIds(types.HrDepartmentModel, ids, h)
}

func (svc *HrDepartmentService) GetByName(name string) (*types.HrDepartments, error) {
	h := &types.HrDepartments{}
	return h, svc.client.getByName(types.HrDepartmentModel, name, h)
}

func (svc *HrDepartmentService) GetByField(field string, value string) (*types.HrDepartments, error) {
	h := &types.HrDepartments{}
	return h, svc.client.getByField(types.HrDepartmentModel, field, value, h)
}

func (svc *HrDepartmentService) GetAll() (*types.HrDepartments, error) {
	h := &types.HrDepartments{}
	return h, svc.client.getAll(types.HrDepartmentModel, h)
}

func (svc *HrDepartmentService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.HrDepartmentModel, fields, relations)
}

func (svc *HrDepartmentService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.HrDepartmentModel, ids, fields, relations)
}

func (svc *HrDepartmentService) Delete(ids []int64) error {
	return svc.client.delete(types.HrDepartmentModel, ids)
}
