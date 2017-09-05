package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type HrJobService struct {
	client *Client
}

func NewHrJobService(c *Client) *HrJobService {
	return &HrJobService{client: c}
}

func (svc *HrJobService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.HrJobModel, name)
}

func (svc *HrJobService) GetByIds(ids []int) (*types.HrJobs, error) {
	h := &types.HrJobs{}
	return h, svc.client.getByIds(types.HrJobModel, ids, h)
}

func (svc *HrJobService) GetByName(name string) (*types.HrJobs, error) {
	h := &types.HrJobs{}
	return h, svc.client.getByName(types.HrJobModel, name, h)
}

func (svc *HrJobService) GetByField(field string, value string) (*types.HrJobs, error) {
	h := &types.HrJobs{}
	return h, svc.client.getByName(types.HrJobModel, field, value, h)
}

func (svc *HrJobService) GetAll() (*types.HrJobs, error) {
	h := &types.HrJobs{}
	return h, svc.client.getAll(types.HrJobModel, h)
}

func (svc *HrJobService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.HrJobModel, fields, relations)
}

func (svc *HrJobService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.HrJobModel, ids, fields, relations)
}

func (svc *HrJobService) Delete(ids []int) error {
	return svc.client.delete(types.HrJobModel, ids)
}
