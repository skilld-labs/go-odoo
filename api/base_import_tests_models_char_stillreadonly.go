package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseImportTestsModelsCharStillreadonlyService struct {
	client *Client
}

func NewBaseImportTestsModelsCharStillreadonlyService(c *Client) *BaseImportTestsModelsCharStillreadonlyService {
	return &BaseImportTestsModelsCharStillreadonlyService{client: c}
}

func (svc *BaseImportTestsModelsCharStillreadonlyService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BaseImportTestsModelsCharStillreadonlyModel, name)
}

func (svc *BaseImportTestsModelsCharStillreadonlyService) GetByIds(ids []int) (*types.BaseImportTestsModelsCharStillreadonlys, error) {
	b := &types.BaseImportTestsModelsCharStillreadonlys{}
	return b, svc.client.getByIds(types.BaseImportTestsModelsCharStillreadonlyModel, ids, b)
}

func (svc *BaseImportTestsModelsCharStillreadonlyService) GetByName(name string) (*types.BaseImportTestsModelsCharStillreadonlys, error) {
	b := &types.BaseImportTestsModelsCharStillreadonlys{}
	return b, svc.client.getByName(types.BaseImportTestsModelsCharStillreadonlyModel, name, b)
}

func (svc *BaseImportTestsModelsCharStillreadonlyService) GetAll() (*types.BaseImportTestsModelsCharStillreadonlys, error) {
	b := &types.BaseImportTestsModelsCharStillreadonlys{}
	return b, svc.client.getAll(types.BaseImportTestsModelsCharStillreadonlyModel, b)
}

func (svc *BaseImportTestsModelsCharStillreadonlyService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BaseImportTestsModelsCharStillreadonlyModel, fields, relations)
}

func (svc *BaseImportTestsModelsCharStillreadonlyService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseImportTestsModelsCharStillreadonlyModel, ids, fields, relations)
}

func (svc *BaseImportTestsModelsCharStillreadonlyService) Delete(ids []int) error {
	return svc.client.delete(types.BaseImportTestsModelsCharStillreadonlyModel, ids)
}
