package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseImportImportService struct {
	client *Client
}

func NewBaseImportImportService(c *Client) *BaseImportImportService {
	return &BaseImportImportService{client: c}
}

func (svc *BaseImportImportService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BaseImportImportModel, name)
}

func (svc *BaseImportImportService) GetByIds(ids []int) (*types.BaseImportImports, error) {
	b := &types.BaseImportImports{}
	return b, svc.client.getByIds(types.BaseImportImportModel, ids, b)
}

func (svc *BaseImportImportService) GetByName(name string) (*types.BaseImportImports, error) {
	b := &types.BaseImportImports{}
	return b, svc.client.getByName(types.BaseImportImportModel, name, b)
}

func (svc *BaseImportImportService) GetAll() (*types.BaseImportImports, error) {
	b := &types.BaseImportImports{}
	return b, svc.client.getAll(types.BaseImportImportModel, b)
}

func (svc *BaseImportImportService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BaseImportImportModel, fields, relations)
}

func (svc *BaseImportImportService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseImportImportModel, ids, fields, relations)
}

func (svc *BaseImportImportService) Delete(ids []int) error {
	return svc.client.delete(types.BaseImportImportModel, ids)
}
