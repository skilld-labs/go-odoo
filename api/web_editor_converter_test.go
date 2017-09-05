package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type WebEditorConverterTestService struct {
	client *Client
}

func NewWebEditorConverterTestService(c *Client) *WebEditorConverterTestService {
	return &WebEditorConverterTestService{client: c}
}

func (svc *WebEditorConverterTestService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.WebEditorConverterTestModel, name)
}

func (svc *WebEditorConverterTestService) GetByIds(ids []int) (*types.WebEditorConverterTests, error) {
	w := &types.WebEditorConverterTests{}
	return w, svc.client.getByIds(types.WebEditorConverterTestModel, ids, w)
}

func (svc *WebEditorConverterTestService) GetByName(name string) (*types.WebEditorConverterTests, error) {
	w := &types.WebEditorConverterTests{}
	return w, svc.client.getByName(types.WebEditorConverterTestModel, name, w)
}

func (svc *WebEditorConverterTestService) GetByField(field string, value string) (*types.WebEditorConverterTests, error) {
	w := &types.WebEditorConverterTests{}
	return w, svc.client.getByField(types.WebEditorConverterTestModel, field, value, w)
}

func (svc *WebEditorConverterTestService) GetAll() (*types.WebEditorConverterTests, error) {
	w := &types.WebEditorConverterTests{}
	return w, svc.client.getAll(types.WebEditorConverterTestModel, w)
}

func (svc *WebEditorConverterTestService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.WebEditorConverterTestModel, fields, relations)
}

func (svc *WebEditorConverterTestService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.WebEditorConverterTestModel, ids, fields, relations)
}

func (svc *WebEditorConverterTestService) Delete(ids []int) error {
	return svc.client.delete(types.WebEditorConverterTestModel, ids)
}
