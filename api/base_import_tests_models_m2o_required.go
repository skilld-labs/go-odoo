package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseImportTestsModelsM2oRequiredService struct {
	client *Client
}

func NewBaseImportTestsModelsM2oRequiredService(c *Client) *BaseImportTestsModelsM2oRequiredService {
	return &BaseImportTestsModelsM2oRequiredService{client: c}
}

func (svc *BaseImportTestsModelsM2oRequiredService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BaseImportTestsModelsM2oRequiredModel, name)
}

func (svc *BaseImportTestsModelsM2oRequiredService) GetByIds(ids []int) (*types.BaseImportTestsModelsM2oRequireds, error) {
	b := &types.BaseImportTestsModelsM2oRequireds{}
	return b, svc.client.getByIds(types.BaseImportTestsModelsM2oRequiredModel, ids, b)
}

func (svc *BaseImportTestsModelsM2oRequiredService) GetByName(name string) (*types.BaseImportTestsModelsM2oRequireds, error) {
	b := &types.BaseImportTestsModelsM2oRequireds{}
	return b, svc.client.getByName(types.BaseImportTestsModelsM2oRequiredModel, name, b)
}

func (svc *BaseImportTestsModelsM2oRequiredService) GetByField(field string, value string) (*types.BaseImportTestsModelsM2oRequireds, error) {
	b := &types.BaseImportTestsModelsM2oRequireds{}
	return b, svc.client.getByName(types.BaseImportTestsModelsM2oRequiredModel, field, value, b)
}

func (svc *BaseImportTestsModelsM2oRequiredService) GetAll() (*types.BaseImportTestsModelsM2oRequireds, error) {
	b := &types.BaseImportTestsModelsM2oRequireds{}
	return b, svc.client.getAll(types.BaseImportTestsModelsM2oRequiredModel, b)
}

func (svc *BaseImportTestsModelsM2oRequiredService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BaseImportTestsModelsM2oRequiredModel, fields, relations)
}

func (svc *BaseImportTestsModelsM2oRequiredService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseImportTestsModelsM2oRequiredModel, ids, fields, relations)
}

func (svc *BaseImportTestsModelsM2oRequiredService) Delete(ids []int) error {
	return svc.client.delete(types.BaseImportTestsModelsM2oRequiredModel, ids)
}
