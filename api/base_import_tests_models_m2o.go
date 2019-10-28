package api

import (
	"github.com/morezig/go-odoo/types"
)

type BaseImportTestsModelsM2oService struct {
	client *Client
}

func NewBaseImportTestsModelsM2oService(c *Client) *BaseImportTestsModelsM2oService {
	return &BaseImportTestsModelsM2oService{client: c}
}

func (svc *BaseImportTestsModelsM2oService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.BaseImportTestsModelsM2oModel, name)
}

func (svc *BaseImportTestsModelsM2oService) GetByIds(ids []int64) (*types.BaseImportTestsModelsM2os, error) {
	b := &types.BaseImportTestsModelsM2os{}
	return b, svc.client.getByIds(types.BaseImportTestsModelsM2oModel, ids, b)
}

func (svc *BaseImportTestsModelsM2oService) GetByName(name string) (*types.BaseImportTestsModelsM2os, error) {
	b := &types.BaseImportTestsModelsM2os{}
	return b, svc.client.getByName(types.BaseImportTestsModelsM2oModel, name, b)
}

func (svc *BaseImportTestsModelsM2oService) GetByField(field string, value string) (*types.BaseImportTestsModelsM2os, error) {
	b := &types.BaseImportTestsModelsM2os{}
	return b, svc.client.getByField(types.BaseImportTestsModelsM2oModel, field, value, b)
}

func (svc *BaseImportTestsModelsM2oService) GetAll() (*types.BaseImportTestsModelsM2os, error) {
	b := &types.BaseImportTestsModelsM2os{}
	return b, svc.client.getAll(types.BaseImportTestsModelsM2oModel, b)
}

func (svc *BaseImportTestsModelsM2oService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.BaseImportTestsModelsM2oModel, fields, relations)
}

func (svc *BaseImportTestsModelsM2oService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseImportTestsModelsM2oModel, ids, fields, relations)
}

func (svc *BaseImportTestsModelsM2oService) Delete(ids []int64) error {
	return svc.client.delete(types.BaseImportTestsModelsM2oModel, ids)
}
