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

func (svc *PortalWizardService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.PortalWizardModel, name)
}

func (svc *PortalWizardService) GetByIds(ids []int) (*types.PortalWizards, error) {
	p := &types.PortalWizards{}
	return p, svc.client.getByIds(types.PortalWizardModel, ids, p)
}

func (svc *PortalWizardService) GetByName(name string) (*types.PortalWizards, error) {
	p := &types.PortalWizards{}
	return p, svc.client.getByName(types.PortalWizardModel, name, p)
}

func (svc *PortalWizardService) GetByField(field string, value string) (*types.PortalWizards, error) {
	p := &types.PortalWizards{}
	return p, svc.client.getByName(types.PortalWizardModel, field, value, p)
}

func (svc *PortalWizardService) GetAll() (*types.PortalWizards, error) {
	p := &types.PortalWizards{}
	return p, svc.client.getAll(types.PortalWizardModel, p)
}

func (svc *PortalWizardService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.PortalWizardModel, fields, relations)
}

func (svc *PortalWizardService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.PortalWizardModel, ids, fields, relations)
}

func (svc *PortalWizardService) Delete(ids []int) error {
	return svc.client.delete(types.PortalWizardModel, ids)
}
