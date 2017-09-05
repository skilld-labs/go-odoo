package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseImportTestsModelsCharRequiredService struct {
	client *Client
}

func NewBaseImportTestsModelsCharRequiredService(c *Client) *BaseImportTestsModelsCharRequiredService {
	return &BaseImportTestsModelsCharRequiredService{client: c}
}

func (svc *BaseImportTestsModelsCharRequiredService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BaseImportTestsModelsCharRequiredModel, name)
}

func (svc *BaseImportTestsModelsCharRequiredService) GetByIds(ids []int) (*types.BaseImportTestsModelsCharRequireds, error) {
	b := &types.BaseImportTestsModelsCharRequireds{}
	return b, svc.client.getByIds(types.BaseImportTestsModelsCharRequiredModel, ids, b)
}

func (svc *BaseImportTestsModelsCharRequiredService) GetByName(name string) (*types.BaseImportTestsModelsCharRequireds, error) {
	b := &types.BaseImportTestsModelsCharRequireds{}
	return b, svc.client.getByName(types.BaseImportTestsModelsCharRequiredModel, name, b)
}

func (svc *BaseImportTestsModelsCharRequiredService) GetByField(field string, value string) (*types.BaseImportTestsModelsCharRequireds, error) {
	b := &types.BaseImportTestsModelsCharRequireds{}
	return b, svc.client.getByName(types.BaseImportTestsModelsCharRequiredModel, field, value, b)
}

func (svc *BaseImportTestsModelsCharRequiredService) GetAll() (*types.BaseImportTestsModelsCharRequireds, error) {
	b := &types.BaseImportTestsModelsCharRequireds{}
	return b, svc.client.getAll(types.BaseImportTestsModelsCharRequiredModel, b)
}

func (svc *BaseImportTestsModelsCharRequiredService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BaseImportTestsModelsCharRequiredModel, fields, relations)
}

func (svc *BaseImportTestsModelsCharRequiredService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseImportTestsModelsCharRequiredModel, ids, fields, relations)
}

func (svc *BaseImportTestsModelsCharRequiredService) Delete(ids []int) error {
	return svc.client.delete(types.BaseImportTestsModelsCharRequiredModel, ids)
}
