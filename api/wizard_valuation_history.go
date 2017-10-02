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

func (svc *WizardValuationHistoryService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.WizardValuationHistoryModel, name)
}

func (svc *WizardValuationHistoryService) GetByIds(ids []int64) (*types.WizardValuationHistorys, error) {
	w := &types.WizardValuationHistorys{}
	return w, svc.client.getByIds(types.WizardValuationHistoryModel, ids, w)
}

func (svc *WizardValuationHistoryService) GetByName(name string) (*types.WizardValuationHistorys, error) {
	w := &types.WizardValuationHistorys{}
	return w, svc.client.getByName(types.WizardValuationHistoryModel, name, w)
}

func (svc *WizardValuationHistoryService) GetByField(field string, value string) (*types.WizardValuationHistorys, error) {
	w := &types.WizardValuationHistorys{}
	return w, svc.client.getByField(types.WizardValuationHistoryModel, field, value, w)
}

func (svc *WizardValuationHistoryService) GetAll() (*types.WizardValuationHistorys, error) {
	w := &types.WizardValuationHistorys{}
	return w, svc.client.getAll(types.WizardValuationHistoryModel, w)
}

func (svc *WizardValuationHistoryService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.WizardValuationHistoryModel, fields, relations)
}

func (svc *WizardValuationHistoryService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.WizardValuationHistoryModel, ids, fields, relations)
}

func (svc *WizardValuationHistoryService) Delete(ids []int64) error {
	return svc.client.delete(types.WizardValuationHistoryModel, ids)
}
