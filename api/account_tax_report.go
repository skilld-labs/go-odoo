package api

import (
	"github.com/morezig/go-odoo/types"
)

type AccountTaxReportService struct {
	client *Client
}

func NewAccountTaxReportService(c *Client) *AccountTaxReportService {
	return &AccountTaxReportService{client: c}
}

func (svc *AccountTaxReportService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountTaxReportModel, name)
}

func (svc *AccountTaxReportService) GetByIds(ids []int64) (*types.AccountTaxReports, error) {
	a := &types.AccountTaxReports{}
	return a, svc.client.getByIds(types.AccountTaxReportModel, ids, a)
}

func (svc *AccountTaxReportService) GetByName(name string) (*types.AccountTaxReports, error) {
	a := &types.AccountTaxReports{}
	return a, svc.client.getByName(types.AccountTaxReportModel, name, a)
}

func (svc *AccountTaxReportService) GetByField(field string, value string) (*types.AccountTaxReports, error) {
	a := &types.AccountTaxReports{}
	return a, svc.client.getByField(types.AccountTaxReportModel, field, value, a)
}

func (svc *AccountTaxReportService) GetAll() (*types.AccountTaxReports, error) {
	a := &types.AccountTaxReports{}
	return a, svc.client.getAll(types.AccountTaxReportModel, a)
}

func (svc *AccountTaxReportService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountTaxReportModel, fields, relations)
}

func (svc *AccountTaxReportService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountTaxReportModel, ids, fields, relations)
}

func (svc *AccountTaxReportService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountTaxReportModel, ids)
}
