package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ReportAccountReportGeneralledgerService struct {
	client *Client
}

func NewReportAccountReportGeneralledgerService(c *Client) *ReportAccountReportGeneralledgerService {
	return &ReportAccountReportGeneralledgerService{client: c}
}

func (svc *ReportAccountReportGeneralledgerService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.ReportAccountReportGeneralledgerModel, name)
}

func (svc *ReportAccountReportGeneralledgerService) GetByIds(ids []int) (*types.ReportAccountReportGeneralledgers, error) {
	r := &types.ReportAccountReportGeneralledgers{}
	return r, svc.client.getByIds(types.ReportAccountReportGeneralledgerModel, ids, r)
}

func (svc *ReportAccountReportGeneralledgerService) GetByName(name string) (*types.ReportAccountReportGeneralledgers, error) {
	r := &types.ReportAccountReportGeneralledgers{}
	return r, svc.client.getByName(types.ReportAccountReportGeneralledgerModel, name, r)
}

func (svc *ReportAccountReportGeneralledgerService) GetByField(field string, value string) (*types.ReportAccountReportGeneralledgers, error) {
	r := &types.ReportAccountReportGeneralledgers{}
	return r, svc.client.getByName(types.ReportAccountReportGeneralledgerModel, field, value, r)
}

func (svc *ReportAccountReportGeneralledgerService) GetAll() (*types.ReportAccountReportGeneralledgers, error) {
	r := &types.ReportAccountReportGeneralledgers{}
	return r, svc.client.getAll(types.ReportAccountReportGeneralledgerModel, r)
}

func (svc *ReportAccountReportGeneralledgerService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.ReportAccountReportGeneralledgerModel, fields, relations)
}

func (svc *ReportAccountReportGeneralledgerService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportAccountReportGeneralledgerModel, ids, fields, relations)
}

func (svc *ReportAccountReportGeneralledgerService) Delete(ids []int) error {
	return svc.client.delete(types.ReportAccountReportGeneralledgerModel, ids)
}
