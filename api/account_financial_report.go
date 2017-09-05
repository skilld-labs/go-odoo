package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountFinancialReportService struct {
	client *Client
}

func NewAccountFinancialReportService(c *Client) *AccountFinancialReportService {
	return &AccountFinancialReportService{client: c}
}

func (svc *AccountFinancialReportService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountFinancialReportModel, name)
}

func (svc *AccountFinancialReportService) GetByIds(ids []int) (*types.AccountFinancialReports, error) {
	a := &types.AccountFinancialReports{}
	return a, svc.client.getByIds(types.AccountFinancialReportModel, ids, a)
}

func (svc *AccountFinancialReportService) GetByName(name string) (*types.AccountFinancialReports, error) {
	a := &types.AccountFinancialReports{}
	return a, svc.client.getByName(types.AccountFinancialReportModel, name, a)
}

func (svc *AccountFinancialReportService) GetByField(field string, value string) (*types.AccountFinancialReports, error) {
	a := &types.AccountFinancialReports{}
	return a, svc.client.getByField(types.AccountFinancialReportModel, field, value, a)
}

func (svc *AccountFinancialReportService) GetAll() (*types.AccountFinancialReports, error) {
	a := &types.AccountFinancialReports{}
	return a, svc.client.getAll(types.AccountFinancialReportModel, a)
}

func (svc *AccountFinancialReportService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountFinancialReportModel, fields, relations)
}

func (svc *AccountFinancialReportService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountFinancialReportModel, ids, fields, relations)
}

func (svc *AccountFinancialReportService) Delete(ids []int) error {
	return svc.client.delete(types.AccountFinancialReportModel, ids)
}
