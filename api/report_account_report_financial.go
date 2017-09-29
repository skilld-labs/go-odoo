package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ReportAccountReportFinancialService struct {
	client *Client
}

func NewReportAccountReportFinancialService(c *Client) *ReportAccountReportFinancialService {
	return &ReportAccountReportFinancialService{client: c}
}

func (svc *ReportAccountReportFinancialService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ReportAccountReportFinancialModel, name)
}

func (svc *ReportAccountReportFinancialService) GetByIds(ids []int64) (*types.ReportAccountReportFinancials, error) {
	r := &types.ReportAccountReportFinancials{}
	return r, svc.client.getByIds(types.ReportAccountReportFinancialModel, ids, r)
}

func (svc *ReportAccountReportFinancialService) GetByName(name string) (*types.ReportAccountReportFinancials, error) {
	r := &types.ReportAccountReportFinancials{}
	return r, svc.client.getByName(types.ReportAccountReportFinancialModel, name, r)
}

func (svc *ReportAccountReportFinancialService) GetByField(field string, value string) (*types.ReportAccountReportFinancials, error) {
	r := &types.ReportAccountReportFinancials{}
	return r, svc.client.getByField(types.ReportAccountReportFinancialModel, field, value, r)
}

func (svc *ReportAccountReportFinancialService) GetAll() (*types.ReportAccountReportFinancials, error) {
	r := &types.ReportAccountReportFinancials{}
	return r, svc.client.getAll(types.ReportAccountReportFinancialModel, r)
}

func (svc *ReportAccountReportFinancialService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ReportAccountReportFinancialModel, fields, relations)
}

func (svc *ReportAccountReportFinancialService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportAccountReportFinancialModel, ids, fields, relations)
}

func (svc *ReportAccountReportFinancialService) Delete(ids []int64) error {
	return svc.client.delete(types.ReportAccountReportFinancialModel, ids)
}
