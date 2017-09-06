package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ReportAbstractReportService struct {
	client *Client
}

func NewReportAbstractReportService(c *Client) *ReportAbstractReportService {
	return &ReportAbstractReportService{client: c}
}

func (svc *ReportAbstractReportService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ReportAbstractReportModel, name)
}

func (svc *ReportAbstractReportService) GetByIds(ids []int) (*types.ReportAbstractReports, error) {
	r := &types.ReportAbstractReports{}
	return r, svc.client.getByIds(types.ReportAbstractReportModel, ids, r)
}

func (svc *ReportAbstractReportService) GetByName(name string) (*types.ReportAbstractReports, error) {
	r := &types.ReportAbstractReports{}
	return r, svc.client.getByName(types.ReportAbstractReportModel, name, r)
}

func (svc *ReportAbstractReportService) GetByField(field string, value string) (*types.ReportAbstractReports, error) {
	r := &types.ReportAbstractReports{}
	return r, svc.client.getByField(types.ReportAbstractReportModel, field, value, r)
}

func (svc *ReportAbstractReportService) GetAll() (*types.ReportAbstractReports, error) {
	r := &types.ReportAbstractReports{}
	return r, svc.client.getAll(types.ReportAbstractReportModel, r)
}

func (svc *ReportAbstractReportService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ReportAbstractReportModel, fields, relations)
}

func (svc *ReportAbstractReportService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportAbstractReportModel, ids, fields, relations)
}

func (svc *ReportAbstractReportService) Delete(ids []int) error {
	return svc.client.delete(types.ReportAbstractReportModel, ids)
}
