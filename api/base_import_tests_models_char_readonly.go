package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseImportTestsModelsCharReadonlyService struct {
	client *Client
}

func NewBaseImportTestsModelsCharReadonlyService(c *Client) *BaseImportTestsModelsCharReadonlyService {
	return &BaseImportTestsModelsCharReadonlyService{client: c}
}

func (svc *BaseImportTestsModelsCharReadonlyService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.BaseImportTestsModelsCharReadonlyModel, name)
}

func (svc *BaseImportTestsModelsCharReadonlyService) GetByIds(ids []int64) (*types.BaseImportTestsModelsCharReadonlys, error) {
	b := &types.BaseImportTestsModelsCharReadonlys{}
	return b, svc.client.getByIds(types.BaseImportTestsModelsCharReadonlyModel, ids, b)
}

func (svc *BaseImportTestsModelsCharReadonlyService) GetByName(name string) (*types.BaseImportTestsModelsCharReadonlys, error) {
	b := &types.BaseImportTestsModelsCharReadonlys{}
	return b, svc.client.getByName(types.BaseImportTestsModelsCharReadonlyModel, name, b)
}

func (svc *BaseImportTestsModelsCharReadonlyService) GetByField(field string, value string) (*types.BaseImportTestsModelsCharReadonlys, error) {
	b := &types.BaseImportTestsModelsCharReadonlys{}
	return b, svc.client.getByField(types.BaseImportTestsModelsCharReadonlyModel, field, value, b)
}

func (svc *BaseImportTestsModelsCharReadonlyService) GetAll() (*types.BaseImportTestsModelsCharReadonlys, error) {
	b := &types.BaseImportTestsModelsCharReadonlys{}
	return b, svc.client.getAll(types.BaseImportTestsModelsCharReadonlyModel, b)
}

func (svc *BaseImportTestsModelsCharReadonlyService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.BaseImportTestsModelsCharReadonlyModel, fields, relations)
}

func (svc *BaseImportTestsModelsCharReadonlyService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseImportTestsModelsCharReadonlyModel, ids, fields, relations)
}

func (svc *BaseImportTestsModelsCharReadonlyService) Delete(ids []int64) error {
	return svc.client.delete(types.BaseImportTestsModelsCharReadonlyModel, ids)
}
