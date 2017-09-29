package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseImportTestsModelsCharNoreadonlyService struct {
	client *Client
}

func NewBaseImportTestsModelsCharNoreadonlyService(c *Client) *BaseImportTestsModelsCharNoreadonlyService {
	return &BaseImportTestsModelsCharNoreadonlyService{client: c}
}

func (svc *BaseImportTestsModelsCharNoreadonlyService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.BaseImportTestsModelsCharNoreadonlyModel, name)
}

func (svc *BaseImportTestsModelsCharNoreadonlyService) GetByIds(ids []int64) (*types.BaseImportTestsModelsCharNoreadonlys, error) {
	b := &types.BaseImportTestsModelsCharNoreadonlys{}
	return b, svc.client.getByIds(types.BaseImportTestsModelsCharNoreadonlyModel, ids, b)
}

func (svc *BaseImportTestsModelsCharNoreadonlyService) GetByName(name string) (*types.BaseImportTestsModelsCharNoreadonlys, error) {
	b := &types.BaseImportTestsModelsCharNoreadonlys{}
	return b, svc.client.getByName(types.BaseImportTestsModelsCharNoreadonlyModel, name, b)
}

func (svc *BaseImportTestsModelsCharNoreadonlyService) GetByField(field string, value string) (*types.BaseImportTestsModelsCharNoreadonlys, error) {
	b := &types.BaseImportTestsModelsCharNoreadonlys{}
	return b, svc.client.getByField(types.BaseImportTestsModelsCharNoreadonlyModel, field, value, b)
}

func (svc *BaseImportTestsModelsCharNoreadonlyService) GetAll() (*types.BaseImportTestsModelsCharNoreadonlys, error) {
	b := &types.BaseImportTestsModelsCharNoreadonlys{}
	return b, svc.client.getAll(types.BaseImportTestsModelsCharNoreadonlyModel, b)
}

func (svc *BaseImportTestsModelsCharNoreadonlyService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.BaseImportTestsModelsCharNoreadonlyModel, fields, relations)
}

func (svc *BaseImportTestsModelsCharNoreadonlyService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseImportTestsModelsCharNoreadonlyModel, ids, fields, relations)
}

func (svc *BaseImportTestsModelsCharNoreadonlyService) Delete(ids []int64) error {
	return svc.client.delete(types.BaseImportTestsModelsCharNoreadonlyModel, ids)
}
