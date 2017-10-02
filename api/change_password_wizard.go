package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ChangePasswordWizardService struct {
	client *Client
}

func NewChangePasswordWizardService(c *Client) *ChangePasswordWizardService {
	return &ChangePasswordWizardService{client: c}
}

func (svc *ChangePasswordWizardService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ChangePasswordWizardModel, name)
}

func (svc *ChangePasswordWizardService) GetByIds(ids []int64) (*types.ChangePasswordWizards, error) {
	c := &types.ChangePasswordWizards{}
	return c, svc.client.getByIds(types.ChangePasswordWizardModel, ids, c)
}

func (svc *ChangePasswordWizardService) GetByName(name string) (*types.ChangePasswordWizards, error) {
	c := &types.ChangePasswordWizards{}
	return c, svc.client.getByName(types.ChangePasswordWizardModel, name, c)
}

func (svc *ChangePasswordWizardService) GetByField(field string, value string) (*types.ChangePasswordWizards, error) {
	c := &types.ChangePasswordWizards{}
	return c, svc.client.getByField(types.ChangePasswordWizardModel, field, value, c)
}

func (svc *ChangePasswordWizardService) GetAll() (*types.ChangePasswordWizards, error) {
	c := &types.ChangePasswordWizards{}
	return c, svc.client.getAll(types.ChangePasswordWizardModel, c)
}

func (svc *ChangePasswordWizardService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ChangePasswordWizardModel, fields, relations)
}

func (svc *ChangePasswordWizardService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ChangePasswordWizardModel, ids, fields, relations)
}

func (svc *ChangePasswordWizardService) Delete(ids []int64) error {
	return svc.client.delete(types.ChangePasswordWizardModel, ids)
}
