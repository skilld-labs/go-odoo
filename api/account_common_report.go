package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountCommonReportService struct {
	client *Client
}

func NewAccountCommonReportService(c *Client) *AccountCommonReportService {
	return &AccountCommonReportService{client: c}
}

func (svc *AccountCommonReportService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountCommonReportModel, name)
}

func (svc *AccountCommonReportService) GetByIds(ids []int) (*types.AccountCommonReports, error) {
	a := &types.AccountCommonReports{}
	return a, svc.client.getByIds(types.AccountCommonReportModel, ids, a)
}

func (svc *AccountCommonReportService) GetByName(name string) (*types.AccountCommonReports, error) {
	a := &types.AccountCommonReports{}
	return a, svc.client.getByName(types.AccountCommonReportModel, name, a)
}

func (svc *AccountCommonReportService) GetByField(field string, value string) (*types.AccountCommonReports, error) {
	a := &types.AccountCommonReports{}
	return a, svc.client.getByName(types.AccountCommonReportModel, field, value, a)
}

func (svc *AccountCommonReportService) GetAll() (*types.AccountCommonReports, error) {
	a := &types.AccountCommonReports{}
	return a, svc.client.getAll(types.AccountCommonReportModel, a)
}

func (svc *AccountCommonReportService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountCommonReportModel, fields, relations)
}

func (svc *AccountCommonReportService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountCommonReportModel, ids, fields, relations)
}

func (svc *AccountCommonReportService) Delete(ids []int) error {
	return svc.client.delete(types.AccountCommonReportModel, ids)
}
