package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ReportAccountReportAgedpartnerbalanceService struct {
	client *Client
}

func NewReportAccountReportAgedpartnerbalanceService(c *Client) *ReportAccountReportAgedpartnerbalanceService {
	return &ReportAccountReportAgedpartnerbalanceService{client: c}
}

func (svc *ReportAccountReportAgedpartnerbalanceService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ReportAccountReportAgedpartnerbalanceModel, name)
}

func (svc *ReportAccountReportAgedpartnerbalanceService) GetByIds(ids []int) (*types.ReportAccountReportAgedpartnerbalances, error) {
	r := &types.ReportAccountReportAgedpartnerbalances{}
	return r, svc.client.getByIds(types.ReportAccountReportAgedpartnerbalanceModel, ids, r)
}

func (svc *ReportAccountReportAgedpartnerbalanceService) GetByName(name string) (*types.ReportAccountReportAgedpartnerbalances, error) {
	r := &types.ReportAccountReportAgedpartnerbalances{}
	return r, svc.client.getByName(types.ReportAccountReportAgedpartnerbalanceModel, name, r)
}

func (svc *ReportAccountReportAgedpartnerbalanceService) GetByField(field string, value string) (*types.ReportAccountReportAgedpartnerbalances, error) {
	r := &types.ReportAccountReportAgedpartnerbalances{}
	return r, svc.client.getByField(types.ReportAccountReportAgedpartnerbalanceModel, field, value, r)
}

func (svc *ReportAccountReportAgedpartnerbalanceService) GetAll() (*types.ReportAccountReportAgedpartnerbalances, error) {
	r := &types.ReportAccountReportAgedpartnerbalances{}
	return r, svc.client.getAll(types.ReportAccountReportAgedpartnerbalanceModel, r)
}

func (svc *ReportAccountReportAgedpartnerbalanceService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ReportAccountReportAgedpartnerbalanceModel, fields, relations)
}

func (svc *ReportAccountReportAgedpartnerbalanceService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportAccountReportAgedpartnerbalanceModel, ids, fields, relations)
}

func (svc *ReportAccountReportAgedpartnerbalanceService) Delete(ids []int) error {
	return svc.client.delete(types.ReportAccountReportAgedpartnerbalanceModel, ids)
}
