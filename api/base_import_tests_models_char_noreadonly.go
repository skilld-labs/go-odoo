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

func (svc *BaseImportTestsModelsCharNoreadonlyService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BaseImportTestsModelsCharNoreadonlyModel, name)
}

func (svc *BaseImportTestsModelsCharNoreadonlyService) GetByIds(ids []int) (*types.BaseImportTestsModelsCharNoreadonlys, error) {
	b := &types.BaseImportTestsModelsCharNoreadonlys{}
	return b, svc.client.getByIds(types.BaseImportTestsModelsCharNoreadonlyModel, ids, b)
}

func (svc *BaseImportTestsModelsCharNoreadonlyService) GetByName(name string) (*types.BaseImportTestsModelsCharNoreadonlys, error) {
	b := &types.BaseImportTestsModelsCharNoreadonlys{}
	return b, svc.client.getByName(types.BaseImportTestsModelsCharNoreadonlyModel, name, b)
}

func (svc *BaseImportTestsModelsCharNoreadonlyService) GetAll() (*types.BaseImportTestsModelsCharNoreadonlys, error) {
	b := &types.BaseImportTestsModelsCharNoreadonlys{}
	return b, svc.client.getAll(types.BaseImportTestsModelsCharNoreadonlyModel, b)
}

func (svc *BaseImportTestsModelsCharNoreadonlyService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BaseImportTestsModelsCharNoreadonlyModel, fields, relations)
}

func (svc *BaseImportTestsModelsCharNoreadonlyService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseImportTestsModelsCharNoreadonlyModel, ids, fields, relations)
}

func (svc *BaseImportTestsModelsCharNoreadonlyService) Delete(ids []int) error {
	return svc.client.delete(types.BaseImportTestsModelsCharNoreadonlyModel, ids)
}
