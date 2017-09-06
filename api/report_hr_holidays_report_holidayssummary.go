package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ReportHrHolidaysReportHolidayssummaryService struct {
	client *Client
}

func NewReportHrHolidaysReportHolidayssummaryService(c *Client) *ReportHrHolidaysReportHolidayssummaryService {
	return &ReportHrHolidaysReportHolidayssummaryService{client: c}
}

func (svc *ReportHrHolidaysReportHolidayssummaryService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ReportHrHolidaysReportHolidayssummaryModel, name)
}

func (svc *ReportHrHolidaysReportHolidayssummaryService) GetByIds(ids []int) (*types.ReportHrHolidaysReportHolidayssummarys, error) {
	r := &types.ReportHrHolidaysReportHolidayssummarys{}
	return r, svc.client.getByIds(types.ReportHrHolidaysReportHolidayssummaryModel, ids, r)
}

func (svc *ReportHrHolidaysReportHolidayssummaryService) GetByName(name string) (*types.ReportHrHolidaysReportHolidayssummarys, error) {
	r := &types.ReportHrHolidaysReportHolidayssummarys{}
	return r, svc.client.getByName(types.ReportHrHolidaysReportHolidayssummaryModel, name, r)
}

func (svc *ReportHrHolidaysReportHolidayssummaryService) GetByField(field string, value string) (*types.ReportHrHolidaysReportHolidayssummarys, error) {
	r := &types.ReportHrHolidaysReportHolidayssummarys{}
	return r, svc.client.getByField(types.ReportHrHolidaysReportHolidayssummaryModel, field, value, r)
}

func (svc *ReportHrHolidaysReportHolidayssummaryService) GetAll() (*types.ReportHrHolidaysReportHolidayssummarys, error) {
	r := &types.ReportHrHolidaysReportHolidayssummarys{}
	return r, svc.client.getAll(types.ReportHrHolidaysReportHolidayssummaryModel, r)
}

func (svc *ReportHrHolidaysReportHolidayssummaryService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ReportHrHolidaysReportHolidayssummaryModel, fields, relations)
}

func (svc *ReportHrHolidaysReportHolidayssummaryService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportHrHolidaysReportHolidayssummaryModel, ids, fields, relations)
}

func (svc *ReportHrHolidaysReportHolidayssummaryService) Delete(ids []int) error {
	return svc.client.delete(types.ReportHrHolidaysReportHolidayssummaryModel, ids)
}
