package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type ReportBaseReportIrmodulereferenceService struct {
	client *Client
}

func NewReportBaseReportIrmodulereferenceService(c *Client) *ReportBaseReportIrmodulereferenceService {
	return &ReportBaseReportIrmodulereferenceService{client: c}
}

func (svc *ReportBaseReportIrmodulereferenceService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.ReportBaseReportIrmodulereferenceModel, name)
}

func (svc *ReportBaseReportIrmodulereferenceService) GetByIds(ids []int64) (*types.ReportBaseReportIrmodulereferences, error) {
	r := &types.ReportBaseReportIrmodulereferences{}
	return r, svc.client.getByIds(types.ReportBaseReportIrmodulereferenceModel, ids, r)
}

func (svc *ReportBaseReportIrmodulereferenceService) GetByName(name string) (*types.ReportBaseReportIrmodulereferences, error) {
	r := &types.ReportBaseReportIrmodulereferences{}
	return r, svc.client.getByName(types.ReportBaseReportIrmodulereferenceModel, name, r)
}

func (svc *ReportBaseReportIrmodulereferenceService) GetByField(field string, value string) (*types.ReportBaseReportIrmodulereferences, error) {
	r := &types.ReportBaseReportIrmodulereferences{}
	return r, svc.client.getByField(types.ReportBaseReportIrmodulereferenceModel, field, value, r)
}

func (svc *ReportBaseReportIrmodulereferenceService) GetAll() (*types.ReportBaseReportIrmodulereferences, error) {
	r := &types.ReportBaseReportIrmodulereferences{}
	return r, svc.client.getAll(types.ReportBaseReportIrmodulereferenceModel, r)
}

func (svc *ReportBaseReportIrmodulereferenceService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.ReportBaseReportIrmodulereferenceModel, fields, relations)
}

func (svc *ReportBaseReportIrmodulereferenceService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.ReportBaseReportIrmodulereferenceModel, ids, fields, relations)
}

func (svc *ReportBaseReportIrmodulereferenceService) Delete(ids []int64) error {
	return svc.client.delete(types.ReportBaseReportIrmodulereferenceModel, ids)
}
