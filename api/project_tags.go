package api

import (
	"github.com/morezig/go-odoo/types"
)

type ProjectTagsService struct {
	client *Client
}

func NewProjectTagsService(c *Client) *ProjectTagsService {
	return &ProjectTagsService{client: c}
}

func (svc *ProjectTagsService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ProjectTagsModel, name)
}

func (svc *ProjectTagsService) GetByIds(ids []int64) (*types.ProjectTagss, error) {
	p := &types.ProjectTagss{}
	return p, svc.client.getByIds(types.ProjectTagsModel, ids, p)
}

func (svc *ProjectTagsService) GetByName(name string) (*types.ProjectTagss, error) {
	p := &types.ProjectTagss{}
	return p, svc.client.getByName(types.ProjectTagsModel, name, p)
}

func (svc *ProjectTagsService) GetByField(field string, value string) (*types.ProjectTagss, error) {
	p := &types.ProjectTagss{}
	return p, svc.client.getByField(types.ProjectTagsModel, field, value, p)
}

func (svc *ProjectTagsService) GetAll() (*types.ProjectTagss, error) {
	p := &types.ProjectTagss{}
	return p, svc.client.getAll(types.ProjectTagsModel, p)
}

func (svc *ProjectTagsService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ProjectTagsModel, fields, relations)
}

func (svc *ProjectTagsService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProjectTagsModel, ids, fields, relations)
}

func (svc *ProjectTagsService) Delete(ids []int64) error {
	return svc.client.delete(types.ProjectTagsModel, ids)
}
