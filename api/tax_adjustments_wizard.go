package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type TaxAdjustmentsWizardService struct {
	client *Client
}

func NewTaxAdjustmentsWizardService(c *Client) *TaxAdjustmentsWizardService {
	return &TaxAdjustmentsWizardService{client: c}
}

func (svc *TaxAdjustmentsWizardService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.TaxAdjustmentsWizardModel, name)
}

func (svc *TaxAdjustmentsWizardService) GetByIds(ids []int64) (*types.TaxAdjustmentsWizards, error) {
	t := &types.TaxAdjustmentsWizards{}
	return t, svc.client.getByIds(types.TaxAdjustmentsWizardModel, ids, t)
}

func (svc *TaxAdjustmentsWizardService) GetByName(name string) (*types.TaxAdjustmentsWizards, error) {
	t := &types.TaxAdjustmentsWizards{}
	return t, svc.client.getByName(types.TaxAdjustmentsWizardModel, name, t)
}

func (svc *TaxAdjustmentsWizardService) GetByField(field string, value string) (*types.TaxAdjustmentsWizards, error) {
	t := &types.TaxAdjustmentsWizards{}
	return t, svc.client.getByField(types.TaxAdjustmentsWizardModel, field, value, t)
}

func (svc *TaxAdjustmentsWizardService) GetAll() (*types.TaxAdjustmentsWizards, error) {
	t := &types.TaxAdjustmentsWizards{}
	return t, svc.client.getAll(types.TaxAdjustmentsWizardModel, t)
}

func (svc *TaxAdjustmentsWizardService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.TaxAdjustmentsWizardModel, fields, relations)
}

func (svc *TaxAdjustmentsWizardService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.TaxAdjustmentsWizardModel, ids, fields, relations)
}

func (svc *TaxAdjustmentsWizardService) Delete(ids []int64) error {
	return svc.client.delete(types.TaxAdjustmentsWizardModel, ids)
}
