package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseImportTestsModelsO2mService struct {
	client *Client
}

func NewBaseImportTestsModelsO2mService(c *Client) *BaseImportTestsModelsO2mService {
	return &BaseImportTestsModelsO2mService{client: c}
}

func (svc *BaseImportTestsModelsO2mService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.BaseImportTestsModelsO2mModel, name)
}

func (svc *BaseImportTestsModelsO2mService) GetByIds(ids []int64) (*types.BaseImportTestsModelsO2ms, error) {
	b := &types.BaseImportTestsModelsO2ms{}
	return b, svc.client.getByIds(types.BaseImportTestsModelsO2mModel, ids, b)
}

func (svc *BaseImportTestsModelsO2mService) GetByName(name string) (*types.BaseImportTestsModelsO2ms, error) {
	b := &types.BaseImportTestsModelsO2ms{}
	return b, svc.client.getByName(types.BaseImportTestsModelsO2mModel, name, b)
}

func (svc *BaseImportTestsModelsO2mService) GetByField(field string, value string) (*types.BaseImportTestsModelsO2ms, error) {
	b := &types.BaseImportTestsModelsO2ms{}
	return b, svc.client.getByField(types.BaseImportTestsModelsO2mModel, field, value, b)
}

func (svc *BaseImportTestsModelsO2mService) GetAll() (*types.BaseImportTestsModelsO2ms, error) {
	b := &types.BaseImportTestsModelsO2ms{}
	return b, svc.client.getAll(types.BaseImportTestsModelsO2mModel, b)
}

func (svc *BaseImportTestsModelsO2mService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.BaseImportTestsModelsO2mModel, fields, relations)
}

func (svc *BaseImportTestsModelsO2mService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseImportTestsModelsO2mModel, ids, fields, relations)
}

func (svc *BaseImportTestsModelsO2mService) Delete(ids []int64) error {
	return svc.client.delete(types.BaseImportTestsModelsO2mModel, ids)
}
