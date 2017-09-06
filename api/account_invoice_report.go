package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountInvoiceReportService struct {
	client *Client
}

func NewAccountInvoiceReportService(c *Client) *AccountInvoiceReportService {
	return &AccountInvoiceReportService{client: c}
}

func (svc *AccountInvoiceReportService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountInvoiceReportModel, name)
}

func (svc *AccountInvoiceReportService) GetByIds(ids []int) (*types.AccountInvoiceReports, error) {
	a := &types.AccountInvoiceReports{}
	return a, svc.client.getByIds(types.AccountInvoiceReportModel, ids, a)
}

func (svc *AccountInvoiceReportService) GetByName(name string) (*types.AccountInvoiceReports, error) {
	a := &types.AccountInvoiceReports{}
	return a, svc.client.getByName(types.AccountInvoiceReportModel, name, a)
}

func (svc *AccountInvoiceReportService) GetByField(field string, value string) (*types.AccountInvoiceReports, error) {
	a := &types.AccountInvoiceReports{}
	return a, svc.client.getByField(types.AccountInvoiceReportModel, field, value, a)
}

func (svc *AccountInvoiceReportService) GetAll() (*types.AccountInvoiceReports, error) {
	a := &types.AccountInvoiceReports{}
	return a, svc.client.getAll(types.AccountInvoiceReportModel, a)
}

func (svc *AccountInvoiceReportService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountInvoiceReportModel, fields, relations)
}

func (svc *AccountInvoiceReportService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountInvoiceReportModel, ids, fields, relations)
}

func (svc *AccountInvoiceReportService) Delete(ids []int) error {
	return svc.client.delete(types.AccountInvoiceReportModel, ids)
}
