package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseImportTestsModelsM2oRelatedService struct {
	client *Client
}

func NewBaseImportTestsModelsM2oRelatedService(c *Client) *BaseImportTestsModelsM2oRelatedService {
	return &BaseImportTestsModelsM2oRelatedService{client: c}
}

func (svc *BaseImportTestsModelsM2oRelatedService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.BaseImportTestsModelsM2oRelatedModel, name)
}

func (svc *BaseImportTestsModelsM2oRelatedService) GetByIds(ids []int64) (*types.BaseImportTestsModelsM2oRelateds, error) {
	b := &types.BaseImportTestsModelsM2oRelateds{}
	return b, svc.client.getByIds(types.BaseImportTestsModelsM2oRelatedModel, ids, b)
}

func (svc *BaseImportTestsModelsM2oRelatedService) GetByName(name string) (*types.BaseImportTestsModelsM2oRelateds, error) {
	b := &types.BaseImportTestsModelsM2oRelateds{}
	return b, svc.client.getByName(types.BaseImportTestsModelsM2oRelatedModel, name, b)
}

func (svc *BaseImportTestsModelsM2oRelatedService) GetByField(field string, value string) (*types.BaseImportTestsModelsM2oRelateds, error) {
	b := &types.BaseImportTestsModelsM2oRelateds{}
	return b, svc.client.getByField(types.BaseImportTestsModelsM2oRelatedModel, field, value, b)
}

func (svc *BaseImportTestsModelsM2oRelatedService) GetAll() (*types.BaseImportTestsModelsM2oRelateds, error) {
	b := &types.BaseImportTestsModelsM2oRelateds{}
	return b, svc.client.getAll(types.BaseImportTestsModelsM2oRelatedModel, b)
}

func (svc *BaseImportTestsModelsM2oRelatedService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.BaseImportTestsModelsM2oRelatedModel, fields, relations)
}

func (svc *BaseImportTestsModelsM2oRelatedService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseImportTestsModelsM2oRelatedModel, ids, fields, relations)
}

func (svc *BaseImportTestsModelsM2oRelatedService) Delete(ids []int64) error {
	return svc.client.delete(types.BaseImportTestsModelsM2oRelatedModel, ids)
}
