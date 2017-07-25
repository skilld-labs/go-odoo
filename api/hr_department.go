package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type HrDepartmentService struct {
	client *Client
}

func NewHrDepartmentService(c *Client) *HrDepartmentService {
	return &HrDepartmentService{client: c}
}

func (svc *HrDepartmentService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.HrDepartmentModel, name)
}

func (svc *HrDepartmentService) GetByIds(ids []int) (*types.HrDepartments, error) {
	h := &types.HrDepartments{}
	return h, svc.client.getByIds(types.HrDepartmentModel, ids, h)
}

func (svc *HrDepartmentService) GetByName(name string) (*types.HrDepartments, error) {
	h := &types.HrDepartments{}
	return h, svc.client.getByName(types.HrDepartmentModel, name, h)
}

func (svc *HrDepartmentService) GetAll() (*types.HrDepartments, error) {
	h := &types.HrDepartments{}
	return h, svc.client.getAll(types.HrDepartmentModel, h)
}

func (svc *HrDepartmentService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.HrDepartmentModel, fields, relations)
}

func (svc *HrDepartmentService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.HrDepartmentModel, ids, fields, relations)
}

func (svc *HrDepartmentService) Delete(ids []int) error {
	return svc.client.delete(types.HrDepartmentModel, ids)
}
