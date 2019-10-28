package api

import (
	"github.com/morezig/go-odoo/types"
)

type BaseImportImportService struct {
	client *Client
}

func NewBaseImportImportService(c *Client) *BaseImportImportService {
	return &BaseImportImportService{client: c}
}

func (svc *BaseImportImportService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.BaseImportImportModel, name)
}

func (svc *BaseImportImportService) GetByIds(ids []int64) (*types.BaseImportImports, error) {
	b := &types.BaseImportImports{}
	return b, svc.client.getByIds(types.BaseImportImportModel, ids, b)
}

func (svc *BaseImportImportService) GetByName(name string) (*types.BaseImportImports, error) {
	b := &types.BaseImportImports{}
	return b, svc.client.getByName(types.BaseImportImportModel, name, b)
}

func (svc *BaseImportImportService) GetByField(field string, value string) (*types.BaseImportImports, error) {
	b := &types.BaseImportImports{}
	return b, svc.client.getByField(types.BaseImportImportModel, field, value, b)
}

func (svc *BaseImportImportService) GetAll() (*types.BaseImportImports, error) {
	b := &types.BaseImportImports{}
	return b, svc.client.getAll(types.BaseImportImportModel, b)
}

func (svc *BaseImportImportService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.BaseImportImportModel, fields, relations)
}

func (svc *BaseImportImportService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseImportImportModel, ids, fields, relations)
}

func (svc *BaseImportImportService) Delete(ids []int64) error {
	return svc.client.delete(types.BaseImportImportModel, ids)
}
