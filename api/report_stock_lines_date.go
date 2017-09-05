package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ReportStockLinesDateService struct {
	client *Client
}

func NewReportStockLinesDateService(c *Client) *ReportStockLinesDateService {
	return &ReportStockLinesDateService{client: c}
}

func (svc *ReportStockLinesDateService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ReportStockLinesDateModel, name)
}

func (svc *ReportStockLinesDateService) GetByIds(ids []int) (*types.ReportStockLinesDates, error) {
	r := &types.ReportStockLinesDates{}
	return r, svc.client.getByIds(types.ReportStockLinesDateModel, ids, r)
}

func (svc *ReportStockLinesDateService) GetByName(name string) (*types.ReportStockLinesDates, error) {
	r := &types.ReportStockLinesDates{}
	return r, svc.client.getByName(types.ReportStockLinesDateModel, name, r)
}

func (svc *ReportStockLinesDateService) GetByField(field string, value string) (*types.ReportStockLinesDates, error) {
	r := &types.ReportStockLinesDates{}
	return r, svc.client.getByField(types.ReportStockLinesDateModel, field, value, r)
}

func (svc *ReportStockLinesDateService) GetAll() (*types.ReportStockLinesDates, error) {
	r := &types.ReportStockLinesDates{}
	return r, svc.client.getAll(types.ReportStockLinesDateModel, r)
}

func (svc *ReportStockLinesDateService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ReportStockLinesDateModel, fields, relations)
}

func (svc *ReportStockLinesDateService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportStockLinesDateModel, ids, fields, relations)
}

func (svc *ReportStockLinesDateService) Delete(ids []int) error {
	return svc.client.delete(types.ReportStockLinesDateModel, ids)
}
