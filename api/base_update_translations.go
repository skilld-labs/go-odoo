package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type BaseUpdateTranslationsService struct {
	client *Client
}

func NewBaseUpdateTranslationsService(c *Client) *BaseUpdateTranslationsService {
	return &BaseUpdateTranslationsService{client: c}
}

func (svc *BaseUpdateTranslationsService) GetIdsByName(name string) ([]int64, error) {
	return svc.client.getIdsByName(types.BaseUpdateTranslationsModel, name)
}

func (svc *BaseUpdateTranslationsService) GetByIds(ids []int64) (*types.BaseUpdateTranslationss, error) {
	b := &types.BaseUpdateTranslationss{}
	return b, svc.client.getByIds(types.BaseUpdateTranslationsModel, ids, b)
}

func (svc *BaseUpdateTranslationsService) GetByName(name string) (*types.BaseUpdateTranslationss, error) {
	b := &types.BaseUpdateTranslationss{}
	return b, svc.client.getByName(types.BaseUpdateTranslationsModel, name, b)
}

func (svc *BaseUpdateTranslationsService) GetByField(field string, value string) (*types.BaseUpdateTranslationss, error) {
	b := &types.BaseUpdateTranslationss{}
	return b, svc.client.getByField(types.BaseUpdateTranslationsModel, field, value, b)
}

func (svc *BaseUpdateTranslationsService) GetAll() (*types.BaseUpdateTranslationss, error) {
	b := &types.BaseUpdateTranslationss{}
	return b, svc.client.getAll(types.BaseUpdateTranslationsModel, b)
}

func (svc *BaseUpdateTranslationsService) Create(fields map[string]interface{}, relations *types.Relations) (int64, error) {
	return svc.client.create(types.BaseUpdateTranslationsModel, fields, relations)
}

func (svc *BaseUpdateTranslationsService) Update(ids []int64, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.BaseUpdateTranslationsModel, ids, fields, relations)
}

func (svc *BaseUpdateTranslationsService) Delete(ids []int64) error {
	return svc.client.delete(types.BaseUpdateTranslationsModel, ids)
}
