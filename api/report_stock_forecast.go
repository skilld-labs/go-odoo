package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ReportStockForecastService struct {
	client *Client
}

func NewReportStockForecastService(c *Client) *ReportStockForecastService {
	return &ReportStockForecastService{client: c}
}

func (svc *ReportStockForecastService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ReportStockForecastModel, name)
}

func (svc *ReportStockForecastService) GetByIds(ids []int64) (*types.ReportStockForecasts, error) {
	r := &types.ReportStockForecasts{}
	return r, svc.client.getByIds(types.ReportStockForecastModel, ids, r)
}

func (svc *ReportStockForecastService) GetByName(name string) (*types.ReportStockForecasts, error) {
	r := &types.ReportStockForecasts{}
	return r, svc.client.getByName(types.ReportStockForecastModel, name, r)
}

func (svc *ReportStockForecastService) GetByField(field string, value string) (*types.ReportStockForecasts, error) {
	r := &types.ReportStockForecasts{}
	return r, svc.client.getByField(types.ReportStockForecastModel, field, value, r)
}

func (svc *ReportStockForecastService) GetAll() (*types.ReportStockForecasts, error) {
	r := &types.ReportStockForecasts{}
	return r, svc.client.getAll(types.ReportStockForecastModel, r)
}

func (svc *ReportStockForecastService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ReportStockForecastModel, fields, relations)
}

func (svc *ReportStockForecastService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportStockForecastModel, ids, fields, relations)
}

func (svc *ReportStockForecastService) Delete(ids []int64) error {
	return svc.client.delete(types.ReportStockForecastModel, ids)
}
