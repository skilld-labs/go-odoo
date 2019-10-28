package api

import (
	"github.com/morezig/go-odoo/types"
)

type IrExportsService struct {
	client *Client
}

func NewIrExportsService(c *Client) *IrExportsService {
	return &IrExportsService{client: c}
}

func (svc *IrExportsService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.IrExportsModel, name)
}

func (svc *IrExportsService) GetByIds(ids []int64) (*types.IrExportss, error) {
	i := &types.IrExportss{}
	return i, svc.client.getByIds(types.IrExportsModel, ids, i)
}

func (svc *IrExportsService) GetByName(name string) (*types.IrExportss, error) {
	i := &types.IrExportss{}
	return i, svc.client.getByName(types.IrExportsModel, name, i)
}

func (svc *IrExportsService) GetByField(field string, value string) (*types.IrExportss, error) {
	i := &types.IrExportss{}
	return i, svc.client.getByField(types.IrExportsModel, field, value, i)
}

func (svc *IrExportsService) GetAll() (*types.IrExportss, error) {
	i := &types.IrExportss{}
	return i, svc.client.getAll(types.IrExportsModel, i)
}

func (svc *IrExportsService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.IrExportsModel, fields, relations)
}

func (svc *IrExportsService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.IrExportsModel, ids, fields, relations)
}

func (svc *IrExportsService) Delete(ids []int64) error {
	return svc.client.delete(types.IrExportsModel, ids)
}
