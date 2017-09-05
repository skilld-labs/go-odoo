package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ReportAccountReportTrialbalanceService struct {
	client *Client
}

func NewReportAccountReportTrialbalanceService(c *Client) *ReportAccountReportTrialbalanceService {
	return &ReportAccountReportTrialbalanceService{client: c}
}

func (svc *ReportAccountReportTrialbalanceService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ReportAccountReportTrialbalanceModel, name)
}

func (svc *ReportAccountReportTrialbalanceService) GetByIds(ids []int) (*types.ReportAccountReportTrialbalances, error) {
	r := &types.ReportAccountReportTrialbalances{}
	return r, svc.client.getByIds(types.ReportAccountReportTrialbalanceModel, ids, r)
}

func (svc *ReportAccountReportTrialbalanceService) GetByName(name string) (*types.ReportAccountReportTrialbalances, error) {
	r := &types.ReportAccountReportTrialbalances{}
	return r, svc.client.getByName(types.ReportAccountReportTrialbalanceModel, name, r)
}

func (svc *ReportAccountReportTrialbalanceService) GetByField(field string, value string) (*types.ReportAccountReportTrialbalances, error) {
	r := &types.ReportAccountReportTrialbalances{}
	return r, svc.client.getByField(types.ReportAccountReportTrialbalanceModel, field, value, r)
}

func (svc *ReportAccountReportTrialbalanceService) GetAll() (*types.ReportAccountReportTrialbalances, error) {
	r := &types.ReportAccountReportTrialbalances{}
	return r, svc.client.getAll(types.ReportAccountReportTrialbalanceModel, r)
}

func (svc *ReportAccountReportTrialbalanceService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ReportAccountReportTrialbalanceModel, fields, relations)
}

func (svc *ReportAccountReportTrialbalanceService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportAccountReportTrialbalanceModel, ids, fields, relations)
}

func (svc *ReportAccountReportTrialbalanceService) Delete(ids []int) error {
	return svc.client.delete(types.ReportAccountReportTrialbalanceModel, ids)
}
