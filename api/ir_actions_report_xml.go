package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type IrActionsReportXmlService struct {
	client *Client
}

func NewIrActionsReportXmlService(c *Client) *IrActionsReportXmlService {
	return &IrActionsReportXmlService{client: c}
}

func (svc *IrActionsReportXmlService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.IrActionsReportXmlModel, name)
}

func (svc *IrActionsReportXmlService) GetByIds(ids []int) (*types.IrActionsReportXmls, error) {
	i := &types.IrActionsReportXmls{}
	return i, svc.client.getByIds(types.IrActionsReportXmlModel, ids, i)
}

func (svc *IrActionsReportXmlService) GetByName(name string) (*types.IrActionsReportXmls, error) {
	i := &types.IrActionsReportXmls{}
	return i, svc.client.getByName(types.IrActionsReportXmlModel, name, i)
}

func (svc *IrActionsReportXmlService) GetByField(field string, value string) (*types.IrActionsReportXmls, error) {
	i := &types.IrActionsReportXmls{}
	return i, svc.client.getByName(types.IrActionsReportXmlModel, field, value, i)
}

func (svc *IrActionsReportXmlService) GetAll() (*types.IrActionsReportXmls, error) {
	i := &types.IrActionsReportXmls{}
	return i, svc.client.getAll(types.IrActionsReportXmlModel, i)
}

func (svc *IrActionsReportXmlService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.IrActionsReportXmlModel, fields, relations)
}

func (svc *IrActionsReportXmlService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrActionsReportXmlModel, ids, fields, relations)
}

func (svc *IrActionsReportXmlService) Delete(ids []int) error {
	return svc.client.delete(types.IrActionsReportXmlModel, ids)
}
