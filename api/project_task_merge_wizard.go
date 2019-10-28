package api

import (
	"github.com/morezig/go-odoo/types"
)

type ProjectTaskMergeWizardService struct {
	client *Client
}

func NewProjectTaskMergeWizardService(c *Client) *ProjectTaskMergeWizardService {
	return &ProjectTaskMergeWizardService{client: c}
}

func (svc *ProjectTaskMergeWizardService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ProjectTaskMergeWizardModel, name)
}

func (svc *ProjectTaskMergeWizardService) GetByIds(ids []int64) (*types.ProjectTaskMergeWizards, error) {
	p := &types.ProjectTaskMergeWizards{}
	return p, svc.client.getByIds(types.ProjectTaskMergeWizardModel, ids, p)
}

func (svc *ProjectTaskMergeWizardService) GetByName(name string) (*types.ProjectTaskMergeWizards, error) {
	p := &types.ProjectTaskMergeWizards{}
	return p, svc.client.getByName(types.ProjectTaskMergeWizardModel, name, p)
}

func (svc *ProjectTaskMergeWizardService) GetByField(field string, value string) (*types.ProjectTaskMergeWizards, error) {
	p := &types.ProjectTaskMergeWizards{}
	return p, svc.client.getByField(types.ProjectTaskMergeWizardModel, field, value, p)
}

func (svc *ProjectTaskMergeWizardService) GetAll() (*types.ProjectTaskMergeWizards, error) {
	p := &types.ProjectTaskMergeWizards{}
	return p, svc.client.getAll(types.ProjectTaskMergeWizardModel, p)
}

func (svc *ProjectTaskMergeWizardService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ProjectTaskMergeWizardModel, fields, relations)
}

func (svc *ProjectTaskMergeWizardService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProjectTaskMergeWizardModel, ids, fields, relations)
}

func (svc *ProjectTaskMergeWizardService) Delete(ids []int64) error {
	return svc.client.delete(types.ProjectTaskMergeWizardModel, ids)
}
