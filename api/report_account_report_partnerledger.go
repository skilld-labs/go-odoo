package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ReportAccountReportPartnerledgerService struct {
	client *Client
}

func NewReportAccountReportPartnerledgerService(c *Client) *ReportAccountReportPartnerledgerService {
	return &ReportAccountReportPartnerledgerService{client: c}
}

func (svc *ReportAccountReportPartnerledgerService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ReportAccountReportPartnerledgerModel, name)
}

func (svc *ReportAccountReportPartnerledgerService) GetByIds(ids []int64) (*types.ReportAccountReportPartnerledgers, error) {
	r := &types.ReportAccountReportPartnerledgers{}
	return r, svc.client.getByIds(types.ReportAccountReportPartnerledgerModel, ids, r)
}

func (svc *ReportAccountReportPartnerledgerService) GetByName(name string) (*types.ReportAccountReportPartnerledgers, error) {
	r := &types.ReportAccountReportPartnerledgers{}
	return r, svc.client.getByName(types.ReportAccountReportPartnerledgerModel, name, r)
}

func (svc *ReportAccountReportPartnerledgerService) GetByField(field string, value string) (*types.ReportAccountReportPartnerledgers, error) {
	r := &types.ReportAccountReportPartnerledgers{}
	return r, svc.client.getByField(types.ReportAccountReportPartnerledgerModel, field, value, r)
}

func (svc *ReportAccountReportPartnerledgerService) GetAll() (*types.ReportAccountReportPartnerledgers, error) {
	r := &types.ReportAccountReportPartnerledgers{}
	return r, svc.client.getAll(types.ReportAccountReportPartnerledgerModel, r)
}

func (svc *ReportAccountReportPartnerledgerService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ReportAccountReportPartnerledgerModel, fields, relations)
}

func (svc *ReportAccountReportPartnerledgerService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportAccountReportPartnerledgerModel, ids, fields, relations)
}

func (svc *ReportAccountReportPartnerledgerService) Delete(ids []int64) error {
	return svc.client.delete(types.ReportAccountReportPartnerledgerModel, ids)
}
