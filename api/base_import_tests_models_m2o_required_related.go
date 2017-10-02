package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseImportTestsModelsM2oRequiredRelatedService struct {
	client *Client
}

func NewBaseImportTestsModelsM2oRequiredRelatedService(c *Client) *BaseImportTestsModelsM2oRequiredRelatedService {
	return &BaseImportTestsModelsM2oRequiredRelatedService{client: c}
}

func (svc *BaseImportTestsModelsM2oRequiredRelatedService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.BaseImportTestsModelsM2oRequiredRelatedModel, name)
}

func (svc *BaseImportTestsModelsM2oRequiredRelatedService) GetByIds(ids []int64) (*types.BaseImportTestsModelsM2oRequiredRelateds, error) {
	b := &types.BaseImportTestsModelsM2oRequiredRelateds{}
	return b, svc.client.getByIds(types.BaseImportTestsModelsM2oRequiredRelatedModel, ids, b)
}

func (svc *BaseImportTestsModelsM2oRequiredRelatedService) GetByName(name string) (*types.BaseImportTestsModelsM2oRequiredRelateds, error) {
	b := &types.BaseImportTestsModelsM2oRequiredRelateds{}
	return b, svc.client.getByName(types.BaseImportTestsModelsM2oRequiredRelatedModel, name, b)
}

func (svc *BaseImportTestsModelsM2oRequiredRelatedService) GetByField(field string, value string) (*types.BaseImportTestsModelsM2oRequiredRelateds, error) {
	b := &types.BaseImportTestsModelsM2oRequiredRelateds{}
	return b, svc.client.getByField(types.BaseImportTestsModelsM2oRequiredRelatedModel, field, value, b)
}

func (svc *BaseImportTestsModelsM2oRequiredRelatedService) GetAll() (*types.BaseImportTestsModelsM2oRequiredRelateds, error) {
	b := &types.BaseImportTestsModelsM2oRequiredRelateds{}
	return b, svc.client.getAll(types.BaseImportTestsModelsM2oRequiredRelatedModel, b)
}

func (svc *BaseImportTestsModelsM2oRequiredRelatedService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.BaseImportTestsModelsM2oRequiredRelatedModel, fields, relations)
}

func (svc *BaseImportTestsModelsM2oRequiredRelatedService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseImportTestsModelsM2oRequiredRelatedModel, ids, fields, relations)
}

func (svc *BaseImportTestsModelsM2oRequiredRelatedService) Delete(ids []int64) error {
	return svc.client.delete(types.BaseImportTestsModelsM2oRequiredRelatedModel, ids)
}
