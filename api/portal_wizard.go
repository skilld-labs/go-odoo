package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type PortalWizardService struct {
	client *Client
}

func NewPortalWizardService(c *Client) *PortalWizardService {
	return &PortalWizardService{client: c}
}

func (svc *PortalWizardService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.PortalWizardModel, name)
}

func (svc *PortalWizardService) GetByIds(ids []int64) (*types.PortalWizards, error) {
	p := &types.PortalWizards{}
	return p, svc.client.getByIds(types.PortalWizardModel, ids, p)
}

func (svc *PortalWizardService) GetByName(name string) (*types.PortalWizards, error) {
	p := &types.PortalWizards{}
	return p, svc.client.getByName(types.PortalWizardModel, name, p)
}

func (svc *PortalWizardService) GetByField(field string, value string) (*types.PortalWizards, error) {
	p := &types.PortalWizards{}
	return p, svc.client.getByField(types.PortalWizardModel, field, value, p)
}

func (svc *PortalWizardService) GetAll() (*types.PortalWizards, error) {
	p := &types.PortalWizards{}
	return p, svc.client.getAll(types.PortalWizardModel, p)
}

func (svc *PortalWizardService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.PortalWizardModel, fields, relations)
}

func (svc *PortalWizardService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.PortalWizardModel, ids, fields, relations)
}

func (svc *PortalWizardService) Delete(ids []int64) error {
	return svc.client.delete(types.PortalWizardModel, ids)
}
