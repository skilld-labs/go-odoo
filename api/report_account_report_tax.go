package api

import (
	"github.com/morezig/go-odoo/types"
)

type ReportAccountReportTaxService struct {
	client *Client
}

func NewReportAccountReportTaxService(c *Client) *ReportAccountReportTaxService {
	return &ReportAccountReportTaxService{client: c}
}

func (svc *ReportAccountReportTaxService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ReportAccountReportTaxModel, name)
}

func (svc *ReportAccountReportTaxService) GetByIds(ids []int64) (*types.ReportAccountReportTaxs, error) {
	r := &types.ReportAccountReportTaxs{}
	return r, svc.client.getByIds(types.ReportAccountReportTaxModel, ids, r)
}

func (svc *ReportAccountReportTaxService) GetByName(name string) (*types.ReportAccountReportTaxs, error) {
	r := &types.ReportAccountReportTaxs{}
	return r, svc.client.getByName(types.ReportAccountReportTaxModel, name, r)
}

func (svc *ReportAccountReportTaxService) GetByField(field string, value string) (*types.ReportAccountReportTaxs, error) {
	r := &types.ReportAccountReportTaxs{}
	return r, svc.client.getByField(types.ReportAccountReportTaxModel, field, value, r)
}

func (svc *ReportAccountReportTaxService) GetAll() (*types.ReportAccountReportTaxs, error) {
	r := &types.ReportAccountReportTaxs{}
	return r, svc.client.getAll(types.ReportAccountReportTaxModel, r)
}

func (svc *ReportAccountReportTaxService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ReportAccountReportTaxModel, fields, relations)
}

func (svc *ReportAccountReportTaxService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportAccountReportTaxModel, ids, fields, relations)
}

func (svc *ReportAccountReportTaxService) Delete(ids []int64) error {
	return svc.client.delete(types.ReportAccountReportTaxModel, ids)
}
