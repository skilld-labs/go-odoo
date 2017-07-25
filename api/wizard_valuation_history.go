package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type WizardValuationHistoryService struct {
	client *Client
}

func NewWizardValuationHistoryService(c *Client) *WizardValuationHistoryService {
	return &WizardValuationHistoryService{client: c}
}

func (svc *WizardValuationHistoryService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.WizardValuationHistoryModel, name)
}

func (svc *WizardValuationHistoryService) GetByIds(ids []int) (*types.WizardValuationHistorys, error) {
	w := &types.WizardValuationHistorys{}
	return w, svc.client.getByIds(types.WizardValuationHistoryModel, ids, w)
}

func (svc *WizardValuationHistoryService) GetByName(name string) (*types.WizardValuationHistorys, error) {
	w := &types.WizardValuationHistorys{}
	return w, svc.client.getByName(types.WizardValuationHistoryModel, name, w)
}

func (svc *WizardValuationHistoryService) GetAll() (*types.WizardValuationHistorys, error) {
	w := &types.WizardValuationHistorys{}
	return w, svc.client.getAll(types.WizardValuationHistoryModel, w)
}

func (svc *WizardValuationHistoryService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.WizardValuationHistoryModel, fields, relations)
}

func (svc *WizardValuationHistoryService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.WizardValuationHistoryModel, ids, fields, relations)
}

func (svc *WizardValuationHistoryService) Delete(ids []int) error {
	return svc.client.delete(types.WizardValuationHistoryModel, ids)
}
