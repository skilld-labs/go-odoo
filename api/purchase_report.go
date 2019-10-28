package api

import (
	"github.com/morezig/go-odoo/types"
)

type PurchaseReportService struct {
	client *Client
}

func NewPurchaseReportService(c *Client) *PurchaseReportService {
	return &PurchaseReportService{client: c}
}

func (svc *PurchaseReportService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.PurchaseReportModel, name)
}

func (svc *PurchaseReportService) GetByIds(ids []int64) (*types.PurchaseReports, error) {
	p := &types.PurchaseReports{}
	return p, svc.client.getByIds(types.PurchaseReportModel, ids, p)
}

func (svc *PurchaseReportService) GetByName(name string) (*types.PurchaseReports, error) {
	p := &types.PurchaseReports{}
	return p, svc.client.getByName(types.PurchaseReportModel, name, p)
}

func (svc *PurchaseReportService) GetByField(field string, value string) (*types.PurchaseReports, error) {
	p := &types.PurchaseReports{}
	return p, svc.client.getByField(types.PurchaseReportModel, field, value, p)
}

func (svc *PurchaseReportService) GetAll() (*types.PurchaseReports, error) {
	p := &types.PurchaseReports{}
	return p, svc.client.getAll(types.PurchaseReportModel, p)
}

func (svc *PurchaseReportService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.PurchaseReportModel, fields, relations)
}

func (svc *PurchaseReportService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.PurchaseReportModel, ids, fields, relations)
}

func (svc *PurchaseReportService) Delete(ids []int64) error {
	return svc.client.delete(types.PurchaseReportModel, ids)
}
