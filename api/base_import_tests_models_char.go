package api

import (
	"github.com/morezig/go-odoo/types"
)

type BaseImportTestsModelsCharService struct {
	client *Client
}

func NewBaseImportTestsModelsCharService(c *Client) *BaseImportTestsModelsCharService {
	return &BaseImportTestsModelsCharService{client: c}
}

func (svc *BaseImportTestsModelsCharService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.BaseImportTestsModelsCharModel, name)
}

func (svc *BaseImportTestsModelsCharService) GetByIds(ids []int64) (*types.BaseImportTestsModelsChars, error) {
	b := &types.BaseImportTestsModelsChars{}
	return b, svc.client.getByIds(types.BaseImportTestsModelsCharModel, ids, b)
}

func (svc *BaseImportTestsModelsCharService) GetByName(name string) (*types.BaseImportTestsModelsChars, error) {
	b := &types.BaseImportTestsModelsChars{}
	return b, svc.client.getByName(types.BaseImportTestsModelsCharModel, name, b)
}

func (svc *BaseImportTestsModelsCharService) GetByField(field string, value string) (*types.BaseImportTestsModelsChars, error) {
	b := &types.BaseImportTestsModelsChars{}
	return b, svc.client.getByField(types.BaseImportTestsModelsCharModel, field, value, b)
}

func (svc *BaseImportTestsModelsCharService) GetAll() (*types.BaseImportTestsModelsChars, error) {
	b := &types.BaseImportTestsModelsChars{}
	return b, svc.client.getAll(types.BaseImportTestsModelsCharModel, b)
}

func (svc *BaseImportTestsModelsCharService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.BaseImportTestsModelsCharModel, fields, relations)
}

func (svc *BaseImportTestsModelsCharService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseImportTestsModelsCharModel, ids, fields, relations)
}

func (svc *BaseImportTestsModelsCharService) Delete(ids []int64) error {
	return svc.client.delete(types.BaseImportTestsModelsCharModel, ids)
}
