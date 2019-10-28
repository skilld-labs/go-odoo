package api

import (
	"github.com/morezig/go-odoo/types"
)

type AccountCommonAccountReportService struct {
	client *Client
}

func NewAccountCommonAccountReportService(c *Client) *AccountCommonAccountReportService {
	return &AccountCommonAccountReportService{client: c}
}

func (svc *AccountCommonAccountReportService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.AccountCommonAccountReportModel, name)
}

func (svc *AccountCommonAccountReportService) GetByIds(ids []int64) (*types.AccountCommonAccountReports, error) {
	a := &types.AccountCommonAccountReports{}
	return a, svc.client.getByIds(types.AccountCommonAccountReportModel, ids, a)
}

func (svc *AccountCommonAccountReportService) GetByName(name string) (*types.AccountCommonAccountReports, error) {
	a := &types.AccountCommonAccountReports{}
	return a, svc.client.getByName(types.AccountCommonAccountReportModel, name, a)
}

func (svc *AccountCommonAccountReportService) GetByField(field string, value string) (*types.AccountCommonAccountReports, error) {
	a := &types.AccountCommonAccountReports{}
	return a, svc.client.getByField(types.AccountCommonAccountReportModel, field, value, a)
}

func (svc *AccountCommonAccountReportService) GetAll() (*types.AccountCommonAccountReports, error) {
	a := &types.AccountCommonAccountReports{}
	return a, svc.client.getAll(types.AccountCommonAccountReportModel, a)
}

func (svc *AccountCommonAccountReportService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.AccountCommonAccountReportModel, fields, relations)
}

func (svc *AccountCommonAccountReportService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountCommonAccountReportModel, ids, fields, relations)
}

func (svc *AccountCommonAccountReportService) Delete(ids []int64) error {
	return svc.client.delete(types.AccountCommonAccountReportModel, ids)
}
