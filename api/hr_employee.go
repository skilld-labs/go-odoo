package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type HrEmployeeService struct {
	client *Client
}

func NewHrEmployeeService(c *Client) *HrEmployeeService {
	return &HrEmployeeService{client: c}
}

func (svc *HrEmployeeService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.HrEmployeeModel, name)
}

func (svc *HrEmployeeService) GetByIds(ids []int) (*types.HrEmployees, error) {
	h := &types.HrEmployees{}
	return h, svc.client.getByIds(types.HrEmployeeModel, ids, h)
}

func (svc *HrEmployeeService) GetByName(name string) (*types.HrEmployees, error) {
	h := &types.HrEmployees{}
	return h, svc.client.getByName(types.HrEmployeeModel, name, h)
}

func (svc *HrEmployeeService) GetAll() (*types.HrEmployees, error) {
	h := &types.HrEmployees{}
	return h, svc.client.getAll(types.HrEmployeeModel, h)
}

func (svc *HrEmployeeService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.HrEmployeeModel, fields, relations)
}

func (svc *HrEmployeeService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.HrEmployeeModel, ids, fields, relations)
}

func (svc *HrEmployeeService) Delete(ids []int) error {
	return svc.client.delete(types.HrEmployeeModel, ids)
}
