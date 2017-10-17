package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type StockTraceabilityReportService struct {
	client *Client
}

func NewStockTraceabilityReportService(c *Client) *StockTraceabilityReportService {
	return &StockTraceabilityReportService{client: c}
}

func (svc *StockTraceabilityReportService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.StockTraceabilityReportModel, name)
}

func (svc *StockTraceabilityReportService) GetByIds(ids []int) (*types.StockTraceabilityReports, error) {
	s := &types.StockTraceabilityReports{}
	return s, svc.client.getByIds(types.StockTraceabilityReportModel, ids, s)
}

func (svc *StockTraceabilityReportService) GetByName(name string) (*types.StockTraceabilityReports, error) {
	s := &types.StockTraceabilityReports{}
	return s, svc.client.getByName(types.StockTraceabilityReportModel, name, s)
}

func (svc *StockTraceabilityReportService) GetByField(field string, value string) (*types.StockTraceabilityReports, error) {
	s := &types.StockTraceabilityReports{}
	return s, svc.client.getByField(types.StockTraceabilityReportModel, field, value, s)
}

func (svc *StockTraceabilityReportService) GetAll() (*types.StockTraceabilityReports, error) {
	s := &types.StockTraceabilityReports{}
	return s, svc.client.getAll(types.StockTraceabilityReportModel, s)
}

func (svc *StockTraceabilityReportService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.StockTraceabilityReportModel, fields, relations)
}

func (svc *StockTraceabilityReportService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.StockTraceabilityReportModel, ids, fields, relations)
}

func (svc *StockTraceabilityReportService) Delete(ids []int) error {
	return svc.client.delete(types.StockTraceabilityReportModel, ids)
}
