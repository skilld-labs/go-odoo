package api

import (
	"github.com/morezig/go-odoo/types"
)

type AccountCommonPartnerReportService struct {
	client *Client
}

func NewAccountCommonPartnerReportService(c *Client) *AccountCommonPartnerReportService {
	return &AccountCommonPartnerReportService{client: c}
}

func (svc *AccountCommonPartnerReportService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountCommonPartnerReportModel, name)
}

func (svc *AccountCommonPartnerReportService) GetByIds(ids []int64) (*types.AccountCommonPartnerReports, error) {
	a := &types.AccountCommonPartnerReports{}
	return a, svc.client.getByIds(types.AccountCommonPartnerReportModel, ids, a)
}

func (svc *AccountCommonPartnerReportService) GetByName(name string) (*types.AccountCommonPartnerReports, error) {
	a := &types.AccountCommonPartnerReports{}
	return a, svc.client.getByName(types.AccountCommonPartnerReportModel, name, a)
}

func (svc *AccountCommonPartnerReportService) GetByField(field string, value string) (*types.AccountCommonPartnerReports, error) {
	a := &types.AccountCommonPartnerReports{}
	return a, svc.client.getByField(types.AccountCommonPartnerReportModel, field, value, a)
}

func (svc *AccountCommonPartnerReportService) GetAll() (*types.AccountCommonPartnerReports, error) {
	a := &types.AccountCommonPartnerReports{}
	return a, svc.client.getAll(types.AccountCommonPartnerReportModel, a)
}

func (svc *AccountCommonPartnerReportService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountCommonPartnerReportModel, fields, relations)
}

func (svc *AccountCommonPartnerReportService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountCommonPartnerReportModel, ids, fields, relations)
}

func (svc *AccountCommonPartnerReportService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountCommonPartnerReportModel, ids)
}
