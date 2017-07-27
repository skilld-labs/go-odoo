package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type AccountCommonAccountReportService struct {
	client *Client
}

func NewAccountCommonAccountReportService(c *Client) *AccountCommonAccountReportService {
	return &AccountCommonAccountReportService{client: c}
}

func (svc *AccountCommonAccountReportService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.AccountCommonAccountReportModel, name)
}

func (svc *AccountCommonAccountReportService) GetByIds(ids []int) (*types.AccountCommonAccountReports, error) {
	a := &types.AccountCommonAccountReports{}
	return a, svc.client.getByIds(types.AccountCommonAccountReportModel, ids, a)
}

func (svc *AccountCommonAccountReportService) GetByName(name string) (*types.AccountCommonAccountReports, error) {
	a := &types.AccountCommonAccountReports{}
	return a, svc.client.getByName(types.AccountCommonAccountReportModel, name, a)
}

func (svc *AccountCommonAccountReportService) GetAll() (*types.AccountCommonAccountReports, error) {
	a := &types.AccountCommonAccountReports{}
	return a, svc.client.getAll(types.AccountCommonAccountReportModel, a)
}

func (svc *AccountCommonAccountReportService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.AccountCommonAccountReportModel, fields, relations)
}

func (svc *AccountCommonAccountReportService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.AccountCommonAccountReportModel, ids, fields, relations)
}

func (svc *AccountCommonAccountReportService) Delete(ids []int) error {
	return svc.client.delete(types.AccountCommonAccountReportModel, ids)
}
