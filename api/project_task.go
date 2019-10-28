package api

import (
	"github.com/morezig/go-odoo/types"
)

type ProjectTaskService struct {
	client *Client
}

func NewProjectTaskService(c *Client) *ProjectTaskService {
	return &ProjectTaskService{client: c}
}

func (svc *ProjectTaskService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ProjectTaskModel, name)
}

func (svc *ProjectTaskService) GetByIds(ids []int64) (*types.ProjectTasks, error) {
	p := &types.ProjectTasks{}
	return p, svc.client.getByIds(types.ProjectTaskModel, ids, p)
}

func (svc *ProjectTaskService) GetByName(name string) (*types.ProjectTasks, error) {
	p := &types.ProjectTasks{}
	return p, svc.client.getByName(types.ProjectTaskModel, name, p)
}

func (svc *ProjectTaskService) GetByField(field string, value string) (*types.ProjectTasks, error) {
	p := &types.ProjectTasks{}
	return p, svc.client.getByField(types.ProjectTaskModel, field, value, p)
}

func (svc *ProjectTaskService) GetAll() (*types.ProjectTasks, error) {
	p := &types.ProjectTasks{}
	return p, svc.client.getAll(types.ProjectTaskModel, p)
}

func (svc *ProjectTaskService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ProjectTaskModel, fields, relations)
}

func (svc *ProjectTaskService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProjectTaskModel, ids, fields, relations)
}

func (svc *ProjectTaskService) Delete(ids []int64) error {
	return svc.client.delete(types.ProjectTaskModel, ids)
}
