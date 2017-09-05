package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseLanguageInstallService struct {
	client *Client
}

func NewBaseLanguageInstallService(c *Client) *BaseLanguageInstallService {
	return &BaseLanguageInstallService{client: c}
}

func (svc *BaseLanguageInstallService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.BaseLanguageInstallModel, name)
}

func (svc *BaseLanguageInstallService) GetByIds(ids []int) (*types.BaseLanguageInstalls, error) {
	b := &types.BaseLanguageInstalls{}
	return b, svc.client.getByIds(types.BaseLanguageInstallModel, ids, b)
}

func (svc *BaseLanguageInstallService) GetByName(name string) (*types.BaseLanguageInstalls, error) {
	b := &types.BaseLanguageInstalls{}
	return b, svc.client.getByName(types.BaseLanguageInstallModel, name, b)
}

func (svc *BaseLanguageInstallService) GetByField(field string, value string) (*types.BaseLanguageInstalls, error) {
	b := &types.BaseLanguageInstalls{}
	return b, svc.client.getByName(types.BaseLanguageInstallModel, field, value, b)
}

func (svc *BaseLanguageInstallService) GetAll() (*types.BaseLanguageInstalls, error) {
	b := &types.BaseLanguageInstalls{}
	return b, svc.client.getAll(types.BaseLanguageInstallModel, b)
}

func (svc *BaseLanguageInstallService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.BaseLanguageInstallModel, fields, relations)
}

func (svc *BaseLanguageInstallService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseLanguageInstallModel, ids, fields, relations)
}

func (svc *BaseLanguageInstallService) Delete(ids []int) error {
	return svc.client.delete(types.BaseLanguageInstallModel, ids)
}
