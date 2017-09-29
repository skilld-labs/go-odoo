package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type PortalWizardUserService struct {
	client *Client
}

func NewPortalWizardUserService(c *Client) *PortalWizardUserService {
	return &PortalWizardUserService{client: c}
}

func (svc *PortalWizardUserService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.PortalWizardUserModel, name)
}

func (svc *PortalWizardUserService) GetByIds(ids []int64) (*types.PortalWizardUsers, error) {
	p := &types.PortalWizardUsers{}
	return p, svc.client.getByIds(types.PortalWizardUserModel, ids, p)
}

func (svc *PortalWizardUserService) GetByName(name string) (*types.PortalWizardUsers, error) {
	p := &types.PortalWizardUsers{}
	return p, svc.client.getByName(types.PortalWizardUserModel, name, p)
}

func (svc *PortalWizardUserService) GetByField(field string, value string) (*types.PortalWizardUsers, error) {
	p := &types.PortalWizardUsers{}
	return p, svc.client.getByField(types.PortalWizardUserModel, field, value, p)
}

func (svc *PortalWizardUserService) GetAll() (*types.PortalWizardUsers, error) {
	p := &types.PortalWizardUsers{}
	return p, svc.client.getAll(types.PortalWizardUserModel, p)
}

func (svc *PortalWizardUserService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.PortalWizardUserModel, fields, relations)
}

func (svc *PortalWizardUserService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.PortalWizardUserModel, ids, fields, relations)
}

func (svc *PortalWizardUserService) Delete(ids []int64) error {
	return svc.client.delete(types.PortalWizardUserModel, ids)
}
