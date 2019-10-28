package api

import (
	"github.com/morezig/go-odoo/types"
)

type IrActionsReportService struct {
	client *Client
}

func NewIrActionsReportService(c *Client) *IrActionsReportService {
	return &IrActionsReportService{client: c}
}

func (svc *IrActionsReportService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrActionsReportModel, name)
}

func (svc *IrActionsReportService) GetByIds(ids []int64) (*types.IrActionsReports, error) {
	i := &types.IrActionsReports{}
	return i, svc.client.getByIds(types.IrActionsReportModel, ids, i)
}

func (svc *IrActionsReportService) GetByName(name string) (*types.IrActionsReports, error) {
	i := &types.IrActionsReports{}
	return i, svc.client.getByName(types.IrActionsReportModel, name, i)
}

func (svc *IrActionsReportService) GetByField(field string, value string) (*types.IrActionsReports, error) {
	i := &types.IrActionsReports{}
	return i, svc.client.getByField(types.IrActionsReportModel, field, value, i)
}

func (svc *IrActionsReportService) GetAll() (*types.IrActionsReports, error) {
	i := &types.IrActionsReports{}
	return i, svc.client.getAll(types.IrActionsReportModel, i)
}

func (svc *IrActionsReportService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrActionsReportModel, fields, relations)
}

func (svc *IrActionsReportService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrActionsReportModel, ids, fields, relations)
}

func (svc *IrActionsReportService) Delete(ids []int64) error {
	return svc.client.delete(types.IrActionsReportModel, ids)
}
