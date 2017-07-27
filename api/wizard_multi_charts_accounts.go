package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type WizardMultiChartsAccountsService struct {
	client *Client
}

func NewWizardMultiChartsAccountsService(c *Client) *WizardMultiChartsAccountsService {
	return &WizardMultiChartsAccountsService{client: c}
}

func (svc *WizardMultiChartsAccountsService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.WizardMultiChartsAccountsModel, name)
}

func (svc *WizardMultiChartsAccountsService) GetByIds(ids []int) (*types.WizardMultiChartsAccountss, error) {
	w := &types.WizardMultiChartsAccountss{}
	return w, svc.client.getByIds(types.WizardMultiChartsAccountsModel, ids, w)
}

func (svc *WizardMultiChartsAccountsService) GetByName(name string) (*types.WizardMultiChartsAccountss, error) {
	w := &types.WizardMultiChartsAccountss{}
	return w, svc.client.getByName(types.WizardMultiChartsAccountsModel, name, w)
}

func (svc *WizardMultiChartsAccountsService) GetAll() (*types.WizardMultiChartsAccountss, error) {
	w := &types.WizardMultiChartsAccountss{}
	return w, svc.client.getAll(types.WizardMultiChartsAccountsModel, w)
}

func (svc *WizardMultiChartsAccountsService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.WizardMultiChartsAccountsModel, fields, relations)
}

func (svc *WizardMultiChartsAccountsService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.WizardMultiChartsAccountsModel, ids, fields, relations)
}

func (svc *WizardMultiChartsAccountsService) Delete(ids []int) error {
	return svc.client.delete(types.WizardMultiChartsAccountsModel, ids)
}
