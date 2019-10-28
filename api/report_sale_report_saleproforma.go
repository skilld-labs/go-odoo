package api

import (
	"github.com/morezig/go-odoo/types"
)

type ReportSaleReportSaleproformaService struct {
	client *Client
}

func NewReportSaleReportSaleproformaService(c *Client) *ReportSaleReportSaleproformaService {
	return &ReportSaleReportSaleproformaService{client: c}
}

func (svc *ReportSaleReportSaleproformaService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ReportSaleReportSaleproformaModel, name)
}

func (svc *ReportSaleReportSaleproformaService) GetByIds(ids []int64) (*types.ReportSaleReportSaleproformas, error) {
	r := &types.ReportSaleReportSaleproformas{}
	return r, svc.client.getByIds(types.ReportSaleReportSaleproformaModel, ids, r)
}

func (svc *ReportSaleReportSaleproformaService) GetByName(name string) (*types.ReportSaleReportSaleproformas, error) {
	r := &types.ReportSaleReportSaleproformas{}
	return r, svc.client.getByName(types.ReportSaleReportSaleproformaModel, name, r)
}

func (svc *ReportSaleReportSaleproformaService) GetByField(field string, value string) (*types.ReportSaleReportSaleproformas, error) {
	r := &types.ReportSaleReportSaleproformas{}
	return r, svc.client.getByField(types.ReportSaleReportSaleproformaModel, field, value, r)
}

func (svc *ReportSaleReportSaleproformaService) GetAll() (*types.ReportSaleReportSaleproformas, error) {
	r := &types.ReportSaleReportSaleproformas{}
	return r, svc.client.getAll(types.ReportSaleReportSaleproformaModel, r)
}

func (svc *ReportSaleReportSaleproformaService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ReportSaleReportSaleproformaModel, fields, relations)
}

func (svc *ReportSaleReportSaleproformaService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportSaleReportSaleproformaModel, ids, fields, relations)
}

func (svc *ReportSaleReportSaleproformaService) Delete(ids []int64) error {
	return svc.client.delete(types.ReportSaleReportSaleproformaModel, ids)
}
