package api

import (
	"github.com/morezig/go-odoo/types"
)

type MailStatisticsReportService struct {
	client *Client
}

func NewMailStatisticsReportService(c *Client) *MailStatisticsReportService {
	return &MailStatisticsReportService{client: c}
}

func (svc *MailStatisticsReportService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.MailStatisticsReportModel, name)
}

func (svc *MailStatisticsReportService) GetByIds(ids []int64) (*types.MailStatisticsReports, error) {
	m := &types.MailStatisticsReports{}
	return m, svc.client.getByIds(types.MailStatisticsReportModel, ids, m)
}

func (svc *MailStatisticsReportService) GetByName(name string) (*types.MailStatisticsReports, error) {
	m := &types.MailStatisticsReports{}
	return m, svc.client.getByName(types.MailStatisticsReportModel, name, m)
}

func (svc *MailStatisticsReportService) GetByField(field string, value string) (*types.MailStatisticsReports, error) {
	m := &types.MailStatisticsReports{}
	return m, svc.client.getByField(types.MailStatisticsReportModel, field, value, m)
}

func (svc *MailStatisticsReportService) GetAll() (*types.MailStatisticsReports, error) {
	m := &types.MailStatisticsReports{}
	return m, svc.client.getAll(types.MailStatisticsReportModel, m)
}

func (svc *MailStatisticsReportService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.MailStatisticsReportModel, fields, relations)
}

func (svc *MailStatisticsReportService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.MailStatisticsReportModel, ids, fields, relations)
}

func (svc *MailStatisticsReportService) Delete(ids []int64) error {
	return svc.client.delete(types.MailStatisticsReportModel, ids)
}
