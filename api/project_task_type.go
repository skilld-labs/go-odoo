package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ProjectTaskTypeService struct {
	client *Client
}

func NewProjectTaskTypeService(c *Client) *ProjectTaskTypeService {
	return &ProjectTaskTypeService{client: c}
}

func (svc *ProjectTaskTypeService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ProjectTaskTypeModel, name)
}

func (svc *ProjectTaskTypeService) GetByIds(ids []int) (*types.ProjectTaskTypes, error) {
	p := &types.ProjectTaskTypes{}
	return p, svc.client.getByIds(types.ProjectTaskTypeModel, ids, p)
}

func (svc *ProjectTaskTypeService) GetByName(name string) (*types.ProjectTaskTypes, error) {
	p := &types.ProjectTaskTypes{}
	return p, svc.client.getByName(types.ProjectTaskTypeModel, name, p)
}

func (svc *ProjectTaskTypeService) GetByField(field string, value string) (*types.ProjectTaskTypes, error) {
	p := &types.ProjectTaskTypes{}
	return p, svc.client.getByField(types.ProjectTaskTypeModel, field, value, p)
}

func (svc *ProjectTaskTypeService) GetAll() (*types.ProjectTaskTypes, error) {
	p := &types.ProjectTaskTypes{}
	return p, svc.client.getAll(types.ProjectTaskTypeModel, p)
}

func (svc *ProjectTaskTypeService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ProjectTaskTypeModel, fields, relations)
}

func (svc *ProjectTaskTypeService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProjectTaskTypeModel, ids, fields, relations)
}

func (svc *ProjectTaskTypeService) Delete(ids []int) error {
	return svc.client.delete(types.ProjectTaskTypeModel, ids)
}
