package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ReportService struct {
	client *Client
}

func NewReportService(c *Client) *ReportService {
	return &ReportService{client: c}
}

func (svc *ReportService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ReportModel, name)
}

func (svc *ReportService) GetByIds(ids []int64) (*types.Reports, error) {
	r := &types.Reports{}
	return r, svc.client.getByIds(types.ReportModel, ids, r)
}

func (svc *ReportService) GetByName(name string) (*types.Reports, error) {
	r := &types.Reports{}
	return r, svc.client.getByName(types.ReportModel, name, r)
}

func (svc *ReportService) GetByField(field string, value string) (*types.Reports, error) {
	r := &types.Reports{}
	return r, svc.client.getByField(types.ReportModel, field, value, r)
}

func (svc *ReportService) GetAll() (*types.Reports, error) {
	r := &types.Reports{}
	return r, svc.client.getAll(types.ReportModel, r)
}

func (svc *ReportService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ReportModel, fields, relations)
}

func (svc *ReportService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportModel, ids, fields, relations)
}

func (svc *ReportService) Delete(ids []int64) error {
	return svc.client.delete(types.ReportModel, ids)
}
