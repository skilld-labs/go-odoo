package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BasePartnerMergeAutomaticWizardService struct {
	client *Client
}

func NewBasePartnerMergeAutomaticWizardService(c *Client) *BasePartnerMergeAutomaticWizardService {
	return &BasePartnerMergeAutomaticWizardService{client: c}
}

func (svc *BasePartnerMergeAutomaticWizardService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BasePartnerMergeAutomaticWizardModel, name)
}

func (svc *BasePartnerMergeAutomaticWizardService) GetByIds(ids []int) (*types.BasePartnerMergeAutomaticWizards, error) {
	b := &types.BasePartnerMergeAutomaticWizards{}
	return b, svc.client.getByIds(types.BasePartnerMergeAutomaticWizardModel, ids, b)
}

func (svc *BasePartnerMergeAutomaticWizardService) GetByName(name string) (*types.BasePartnerMergeAutomaticWizards, error) {
	b := &types.BasePartnerMergeAutomaticWizards{}
	return b, svc.client.getByName(types.BasePartnerMergeAutomaticWizardModel, name, b)
}

func (svc *BasePartnerMergeAutomaticWizardService) GetByField(field string, value string) (*types.BasePartnerMergeAutomaticWizards, error) {
	b := &types.BasePartnerMergeAutomaticWizards{}
	return b, svc.client.getByName(types.BasePartnerMergeAutomaticWizardModel, field, value, b)
}

func (svc *BasePartnerMergeAutomaticWizardService) GetAll() (*types.BasePartnerMergeAutomaticWizards, error) {
	b := &types.BasePartnerMergeAutomaticWizards{}
	return b, svc.client.getAll(types.BasePartnerMergeAutomaticWizardModel, b)
}

func (svc *BasePartnerMergeAutomaticWizardService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BasePartnerMergeAutomaticWizardModel, fields, relations)
}

func (svc *BasePartnerMergeAutomaticWizardService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BasePartnerMergeAutomaticWizardModel, ids, fields, relations)
}

func (svc *BasePartnerMergeAutomaticWizardService) Delete(ids []int) error {
	return svc.client.delete(types.BasePartnerMergeAutomaticWizardModel, ids)
}
