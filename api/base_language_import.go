package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseLanguageImportService struct {
	client *Client
}

func NewBaseLanguageImportService(c *Client) *BaseLanguageImportService {
	return &BaseLanguageImportService{client: c}
}

func (svc *BaseLanguageImportService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BaseLanguageImportModel, name)
}

func (svc *BaseLanguageImportService) GetByIds(ids []int) (*types.BaseLanguageImports, error) {
	b := &types.BaseLanguageImports{}
	return b, svc.client.getByIds(types.BaseLanguageImportModel, ids, b)
}

func (svc *BaseLanguageImportService) GetByName(name string) (*types.BaseLanguageImports, error) {
	b := &types.BaseLanguageImports{}
	return b, svc.client.getByName(types.BaseLanguageImportModel, name, b)
}

func (svc *BaseLanguageImportService) GetAll() (*types.BaseLanguageImports, error) {
	b := &types.BaseLanguageImports{}
	return b, svc.client.getAll(types.BaseLanguageImportModel, b)
}

func (svc *BaseLanguageImportService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BaseLanguageImportModel, fields, relations)
}

func (svc *BaseLanguageImportService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseLanguageImportModel, ids, fields, relations)
}

func (svc *BaseLanguageImportService) Delete(ids []int) error {
	return svc.client.delete(types.BaseLanguageImportModel, ids)
}
