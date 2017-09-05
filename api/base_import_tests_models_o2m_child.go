package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseImportTestsModelsO2mChildService struct {
	client *Client
}

func NewBaseImportTestsModelsO2mChildService(c *Client) *BaseImportTestsModelsO2mChildService {
	return &BaseImportTestsModelsO2mChildService{client: c}
}

func (svc *BaseImportTestsModelsO2mChildService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BaseImportTestsModelsO2mChildModel, name)
}

func (svc *BaseImportTestsModelsO2mChildService) GetByIds(ids []int) (*types.BaseImportTestsModelsO2mChilds, error) {
	b := &types.BaseImportTestsModelsO2mChilds{}
	return b, svc.client.getByIds(types.BaseImportTestsModelsO2mChildModel, ids, b)
}

func (svc *BaseImportTestsModelsO2mChildService) GetByName(name string) (*types.BaseImportTestsModelsO2mChilds, error) {
	b := &types.BaseImportTestsModelsO2mChilds{}
	return b, svc.client.getByName(types.BaseImportTestsModelsO2mChildModel, name, b)
}

func (svc *BaseImportTestsModelsO2mChildService) GetByField(field string, value string) (*types.BaseImportTestsModelsO2mChilds, error) {
	b := &types.BaseImportTestsModelsO2mChilds{}
	return b, svc.client.getByField(types.BaseImportTestsModelsO2mChildModel, field, value, b)
}

func (svc *BaseImportTestsModelsO2mChildService) GetAll() (*types.BaseImportTestsModelsO2mChilds, error) {
	b := &types.BaseImportTestsModelsO2mChilds{}
	return b, svc.client.getAll(types.BaseImportTestsModelsO2mChildModel, b)
}

func (svc *BaseImportTestsModelsO2mChildService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BaseImportTestsModelsO2mChildModel, fields, relations)
}

func (svc *BaseImportTestsModelsO2mChildService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseImportTestsModelsO2mChildModel, ids, fields, relations)
}

func (svc *BaseImportTestsModelsO2mChildService) Delete(ids []int) error {
	return svc.client.delete(types.BaseImportTestsModelsO2mChildModel, ids)
}
