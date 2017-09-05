package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountingReportService struct {
	client *Client
}

func NewAccountingReportService(c *Client) *AccountingReportService {
	return &AccountingReportService{client: c}
}

func (svc *AccountingReportService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountingReportModel, name)
}

func (svc *AccountingReportService) GetByIds(ids []int) (*types.AccountingReports, error) {
	a := &types.AccountingReports{}
	return a, svc.client.getByIds(types.AccountingReportModel, ids, a)
}

func (svc *AccountingReportService) GetByName(name string) (*types.AccountingReports, error) {
	a := &types.AccountingReports{}
	return a, svc.client.getByName(types.AccountingReportModel, name, a)
}

func (svc *AccountingReportService) GetByField(field string, value string) (*types.AccountingReports, error) {
	a := &types.AccountingReports{}
	return a, svc.client.getByField(types.AccountingReportModel, field, value, a)
}

func (svc *AccountingReportService) GetAll() (*types.AccountingReports, error) {
	a := &types.AccountingReports{}
	return a, svc.client.getAll(types.AccountingReportModel, a)
}

func (svc *AccountingReportService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountingReportModel, fields, relations)
}

func (svc *AccountingReportService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountingReportModel, ids, fields, relations)
}

func (svc *AccountingReportService) Delete(ids []int) error {
	return svc.client.delete(types.AccountingReportModel, ids)
}
