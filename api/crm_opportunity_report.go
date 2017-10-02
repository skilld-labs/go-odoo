package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type CrmOpportunityReportService struct {
	client *Client
}

func NewCrmOpportunityReportService(c *Client) *CrmOpportunityReportService {
	return &CrmOpportunityReportService{client: c}
}

func (svc *CrmOpportunityReportService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.CrmOpportunityReportModel, name)
}

func (svc *CrmOpportunityReportService) GetByIds(ids []int64) (*types.CrmOpportunityReports, error) {
	c := &types.CrmOpportunityReports{}
	return c, svc.client.getByIds(types.CrmOpportunityReportModel, ids, c)
}

func (svc *CrmOpportunityReportService) GetByName(name string) (*types.CrmOpportunityReports, error) {
	c := &types.CrmOpportunityReports{}
	return c, svc.client.getByName(types.CrmOpportunityReportModel, name, c)
}

func (svc *CrmOpportunityReportService) GetByField(field string, value string) (*types.CrmOpportunityReports, error) {
	c := &types.CrmOpportunityReports{}
	return c, svc.client.getByField(types.CrmOpportunityReportModel, field, value, c)
}

func (svc *CrmOpportunityReportService) GetAll() (*types.CrmOpportunityReports, error) {
	c := &types.CrmOpportunityReports{}
	return c, svc.client.getAll(types.CrmOpportunityReportModel, c)
}

func (svc *CrmOpportunityReportService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.CrmOpportunityReportModel, fields, relations)
}

func (svc *CrmOpportunityReportService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.CrmOpportunityReportModel, ids, fields, relations)
}

func (svc *CrmOpportunityReportService) Delete(ids []int64) error {
	return svc.client.delete(types.CrmOpportunityReportModel, ids)
}
