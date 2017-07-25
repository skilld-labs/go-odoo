package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ProjectConfigSettingsService struct {
	client *Client
}

func NewProjectConfigSettingsService(c *Client) *ProjectConfigSettingsService {
	return &ProjectConfigSettingsService{client: c}
}

func (svc *ProjectConfigSettingsService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ProjectConfigSettingsModel, name)
}

func (svc *ProjectConfigSettingsService) GetByIds(ids []int) (*types.ProjectConfigSettingss, error) {
	p := &types.ProjectConfigSettingss{}
	return p, svc.client.getByIds(types.ProjectConfigSettingsModel, ids, p)
}

func (svc *ProjectConfigSettingsService) GetByName(name string) (*types.ProjectConfigSettingss, error) {
	p := &types.ProjectConfigSettingss{}
	return p, svc.client.getByName(types.ProjectConfigSettingsModel, name, p)
}

func (svc *ProjectConfigSettingsService) GetAll() (*types.ProjectConfigSettingss, error) {
	p := &types.ProjectConfigSettingss{}
	return p, svc.client.getAll(types.ProjectConfigSettingsModel, p)
}

func (svc *ProjectConfigSettingsService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ProjectConfigSettingsModel, fields, relations)
}

func (svc *ProjectConfigSettingsService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ProjectConfigSettingsModel, ids, fields, relations)
}

func (svc *ProjectConfigSettingsService) Delete(ids []int) error {
	return svc.client.delete(types.ProjectConfigSettingsModel, ids)
}
