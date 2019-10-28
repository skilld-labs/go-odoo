package api

import (
	"github.com/morezig/go-odoo/types"
)

type BaseImportTestsModelsCharStatesService struct {
	client *Client
}

func NewBaseImportTestsModelsCharStatesService(c *Client) *BaseImportTestsModelsCharStatesService {
	return &BaseImportTestsModelsCharStatesService{client: c}
}

func (svc *BaseImportTestsModelsCharStatesService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.BaseImportTestsModelsCharStatesModel, name)
}

func (svc *BaseImportTestsModelsCharStatesService) GetByIds(ids []int64) (*types.BaseImportTestsModelsCharStatess, error) {
	b := &types.BaseImportTestsModelsCharStatess{}
	return b, svc.client.getByIds(types.BaseImportTestsModelsCharStatesModel, ids, b)
}

func (svc *BaseImportTestsModelsCharStatesService) GetByName(name string) (*types.BaseImportTestsModelsCharStatess, error) {
	b := &types.BaseImportTestsModelsCharStatess{}
	return b, svc.client.getByName(types.BaseImportTestsModelsCharStatesModel, name, b)
}

func (svc *BaseImportTestsModelsCharStatesService) GetByField(field string, value string) (*types.BaseImportTestsModelsCharStatess, error) {
	b := &types.BaseImportTestsModelsCharStatess{}
	return b, svc.client.getByField(types.BaseImportTestsModelsCharStatesModel, field, value, b)
}

func (svc *BaseImportTestsModelsCharStatesService) GetAll() (*types.BaseImportTestsModelsCharStatess, error) {
	b := &types.BaseImportTestsModelsCharStatess{}
	return b, svc.client.getAll(types.BaseImportTestsModelsCharStatesModel, b)
}

func (svc *BaseImportTestsModelsCharStatesService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.BaseImportTestsModelsCharStatesModel, fields, relations)
}

func (svc *BaseImportTestsModelsCharStatesService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseImportTestsModelsCharStatesModel, ids, fields, relations)
}

func (svc *BaseImportTestsModelsCharStatesService) Delete(ids []int64) error {
	return svc.client.delete(types.BaseImportTestsModelsCharStatesModel, ids)
}
