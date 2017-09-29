package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type SaleReportService struct {
	client *Client
}

func NewSaleReportService(c *Client) *SaleReportService {
	return &SaleReportService{client: c}
}

func (svc *SaleReportService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.SaleReportModel, name)
}

func (svc *SaleReportService) GetByIds(ids []int64) (*types.SaleReports, error) {
	s := &types.SaleReports{}
	return s, svc.client.getByIds(types.SaleReportModel, ids, s)
}

func (svc *SaleReportService) GetByName(name string) (*types.SaleReports, error) {
	s := &types.SaleReports{}
	return s, svc.client.getByName(types.SaleReportModel, name, s)
}

func (svc *SaleReportService) GetByField(field string, value string) (*types.SaleReports, error) {
	s := &types.SaleReports{}
	return s, svc.client.getByField(types.SaleReportModel, field, value, s)
}

func (svc *SaleReportService) GetAll() (*types.SaleReports, error) {
	s := &types.SaleReports{}
	return s, svc.client.getAll(types.SaleReportModel, s)
}

func (svc *SaleReportService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.SaleReportModel, fields, relations)
}

func (svc *SaleReportService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.SaleReportModel, ids, fields, relations)
}

func (svc *SaleReportService) Delete(ids []int64) error {
	return svc.client.delete(types.SaleReportModel, ids)
}
