package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ReportAccountReportOverdueService struct {
	client *Client
}

func NewReportAccountReportOverdueService(c *Client) *ReportAccountReportOverdueService {
	return &ReportAccountReportOverdueService{client: c}
}

func (svc *ReportAccountReportOverdueService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ReportAccountReportOverdueModel, name)
}

func (svc *ReportAccountReportOverdueService) GetByIds(ids []int) (*types.ReportAccountReportOverdues, error) {
	r := &types.ReportAccountReportOverdues{}
	return r, svc.client.getByIds(types.ReportAccountReportOverdueModel, ids, r)
}

func (svc *ReportAccountReportOverdueService) GetByName(name string) (*types.ReportAccountReportOverdues, error) {
	r := &types.ReportAccountReportOverdues{}
	return r, svc.client.getByName(types.ReportAccountReportOverdueModel, name, r)
}

func (svc *ReportAccountReportOverdueService) GetByField(field string, value string) (*types.ReportAccountReportOverdues, error) {
	r := &types.ReportAccountReportOverdues{}
	return r, svc.client.getByField(types.ReportAccountReportOverdueModel, field, value, r)
}

func (svc *ReportAccountReportOverdueService) GetAll() (*types.ReportAccountReportOverdues, error) {
	r := &types.ReportAccountReportOverdues{}
	return r, svc.client.getAll(types.ReportAccountReportOverdueModel, r)
}

func (svc *ReportAccountReportOverdueService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ReportAccountReportOverdueModel, fields, relations)
}

func (svc *ReportAccountReportOverdueService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportAccountReportOverdueModel, ids, fields, relations)
}

func (svc *ReportAccountReportOverdueService) Delete(ids []int) error {
	return svc.client.delete(types.ReportAccountReportOverdueModel, ids)
}
