package api

import (
	"github.com/morezig/go-odoo/types"
)

type BaseImportTestsModelsPreviewService struct {
	client *Client
}

func NewBaseImportTestsModelsPreviewService(c *Client) *BaseImportTestsModelsPreviewService {
	return &BaseImportTestsModelsPreviewService{client: c}
}

func (svc *BaseImportTestsModelsPreviewService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.BaseImportTestsModelsPreviewModel, name)
}

func (svc *BaseImportTestsModelsPreviewService) GetByIds(ids []int64) (*types.BaseImportTestsModelsPreviews, error) {
	b := &types.BaseImportTestsModelsPreviews{}
	return b, svc.client.getByIds(types.BaseImportTestsModelsPreviewModel, ids, b)
}

func (svc *BaseImportTestsModelsPreviewService) GetByName(name string) (*types.BaseImportTestsModelsPreviews, error) {
	b := &types.BaseImportTestsModelsPreviews{}
	return b, svc.client.getByName(types.BaseImportTestsModelsPreviewModel, name, b)
}

func (svc *BaseImportTestsModelsPreviewService) GetByField(field string, value string) (*types.BaseImportTestsModelsPreviews, error) {
	b := &types.BaseImportTestsModelsPreviews{}
	return b, svc.client.getByField(types.BaseImportTestsModelsPreviewModel, field, value, b)
}

func (svc *BaseImportTestsModelsPreviewService) GetAll() (*types.BaseImportTestsModelsPreviews, error) {
	b := &types.BaseImportTestsModelsPreviews{}
	return b, svc.client.getAll(types.BaseImportTestsModelsPreviewModel, b)
}

func (svc *BaseImportTestsModelsPreviewService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.BaseImportTestsModelsPreviewModel, fields, relations)
}

func (svc *BaseImportTestsModelsPreviewService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseImportTestsModelsPreviewModel, ids, fields, relations)
}

func (svc *BaseImportTestsModelsPreviewService) Delete(ids []int64) error {
	return svc.client.delete(types.BaseImportTestsModelsPreviewModel, ids)
}
