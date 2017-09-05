package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type WizardIrModelMenuCreateService struct {
	client *Client
}

func NewWizardIrModelMenuCreateService(c *Client) *WizardIrModelMenuCreateService {
	return &WizardIrModelMenuCreateService{client: c}
}

func (svc *WizardIrModelMenuCreateService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.WizardIrModelMenuCreateModel, name)
}

func (svc *WizardIrModelMenuCreateService) GetByIds(ids []int) (*types.WizardIrModelMenuCreates, error) {
	w := &types.WizardIrModelMenuCreates{}
	return w, svc.client.getByIds(types.WizardIrModelMenuCreateModel, ids, w)
}

func (svc *WizardIrModelMenuCreateService) GetByName(name string) (*types.WizardIrModelMenuCreates, error) {
	w := &types.WizardIrModelMenuCreates{}
	return w, svc.client.getByName(types.WizardIrModelMenuCreateModel, name, w)
}

func (svc *WizardIrModelMenuCreateService) GetByField(field string, value string) (*types.WizardIrModelMenuCreates, error) {
	w := &types.WizardIrModelMenuCreates{}
	return w, svc.client.getByField(types.WizardIrModelMenuCreateModel, field, value, w)
}

func (svc *WizardIrModelMenuCreateService) GetAll() (*types.WizardIrModelMenuCreates, error) {
	w := &types.WizardIrModelMenuCreates{}
	return w, svc.client.getAll(types.WizardIrModelMenuCreateModel, w)
}

func (svc *WizardIrModelMenuCreateService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.WizardIrModelMenuCreateModel, fields, relations)
}

func (svc *WizardIrModelMenuCreateService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.WizardIrModelMenuCreateModel, ids, fields, relations)
}

func (svc *WizardIrModelMenuCreateService) Delete(ids []int) error {
	return svc.client.delete(types.WizardIrModelMenuCreateModel, ids)
}
