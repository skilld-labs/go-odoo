package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseLanguageExportService struct {
	client *Client
}

func NewBaseLanguageExportService(c *Client) *BaseLanguageExportService {
	return &BaseLanguageExportService{client: c}
}

func (svc *BaseLanguageExportService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BaseLanguageExportModel, name)
}

func (svc *BaseLanguageExportService) GetByIds(ids []int) (*types.BaseLanguageExports, error) {
	b := &types.BaseLanguageExports{}
	return b, svc.client.getByIds(types.BaseLanguageExportModel, ids, b)
}

func (svc *BaseLanguageExportService) GetByName(name string) (*types.BaseLanguageExports, error) {
	b := &types.BaseLanguageExports{}
	return b, svc.client.getByName(types.BaseLanguageExportModel, name, b)
}

func (svc *BaseLanguageExportService) GetByField(field string, value string) (*types.BaseLanguageExports, error) {
	b := &types.BaseLanguageExports{}
	return b, svc.client.getByName(types.BaseLanguageExportModel, field, value, b)
}

func (svc *BaseLanguageExportService) GetAll() (*types.BaseLanguageExports, error) {
	b := &types.BaseLanguageExports{}
	return b, svc.client.getAll(types.BaseLanguageExportModel, b)
}

func (svc *BaseLanguageExportService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BaseLanguageExportModel, fields, relations)
}

func (svc *BaseLanguageExportService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseLanguageExportModel, ids, fields, relations)
}

func (svc *BaseLanguageExportService) Delete(ids []int) error {
	return svc.client.delete(types.BaseLanguageExportModel, ids)
}
