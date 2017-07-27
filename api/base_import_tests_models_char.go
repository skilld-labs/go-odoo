package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseImportTestsModelsCharService struct {
	client *Client
}

func NewBaseImportTestsModelsCharService(c *Client) *BaseImportTestsModelsCharService {
	return &BaseImportTestsModelsCharService{client: c}
}

func (svc *BaseImportTestsModelsCharService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BaseImportTestsModelsCharModel, name)
}

func (svc *BaseImportTestsModelsCharService) GetByIds(ids []int) (*types.BaseImportTestsModelsChars, error) {
	b := &types.BaseImportTestsModelsChars{}
	return b, svc.client.getByIds(types.BaseImportTestsModelsCharModel, ids, b)
}

func (svc *BaseImportTestsModelsCharService) GetByName(name string) (*types.BaseImportTestsModelsChars, error) {
	b := &types.BaseImportTestsModelsChars{}
	return b, svc.client.getByName(types.BaseImportTestsModelsCharModel, name, b)
}

func (svc *BaseImportTestsModelsCharService) GetAll() (*types.BaseImportTestsModelsChars, error) {
	b := &types.BaseImportTestsModelsChars{}
	return b, svc.client.getAll(types.BaseImportTestsModelsCharModel, b)
}

func (svc *BaseImportTestsModelsCharService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BaseImportTestsModelsCharModel, fields, relations)
}

func (svc *BaseImportTestsModelsCharService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseImportTestsModelsCharModel, ids, fields, relations)
}

func (svc *BaseImportTestsModelsCharService) Delete(ids []int) error {
	return svc.client.delete(types.BaseImportTestsModelsCharModel, ids)
}
